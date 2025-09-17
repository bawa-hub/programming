package main

import (
	"fmt"
	"sync"
	"time"
)

// CircuitState represents the state of the circuit breaker
type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	mu            sync.RWMutex
	state         CircuitState
	failureCount  int
	successCount  int
	maxFailures   int
	timeout       time.Duration
	lastFailTime  time.Time
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if circuit is open and timeout has passed
	if cb.state == StateOpen {
		if time.Since(cb.lastFailTime) > cb.timeout {
			cb.state = StateHalfOpen
			cb.successCount = 0
		} else {
			return fmt.Errorf("circuit breaker is OPEN")
		}
	}
	
	// Execute the function
	err := fn()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailTime = time.Now()
		
		if cb.state == StateHalfOpen || cb.failureCount >= cb.maxFailures {
			cb.state = StateOpen
		}
		
		return err
	}
	
	// Success
	cb.failureCount = 0
	if cb.state == StateHalfOpen {
		cb.successCount++
		if cb.successCount >= 2 { // Require 2 successes to close
			cb.state = StateClosed
		}
	}
	
	return nil
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}
