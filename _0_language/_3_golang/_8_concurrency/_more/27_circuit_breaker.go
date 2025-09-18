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
			fmt.Println("Circuit breaker: Moving to HALF-OPEN state")
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
			fmt.Printf("Circuit breaker: Moving to OPEN state (failures: %d)\n", cb.failureCount)
		}
		
		return err
	}
	
	// Success
	cb.failureCount = 0
	if cb.state == StateHalfOpen {
		cb.successCount++
		if cb.successCount >= 2 { // Require 2 successes to close
			cb.state = StateClosed
			fmt.Println("Circuit breaker: Moving to CLOSED state")
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

// Simulate a service that sometimes fails
func simulateServiceCall(shouldFail bool) error {
	if shouldFail {
		return fmt.Errorf("service unavailable")
	}
	return nil
}

func main() {
	fmt.Println("=== Circuit Breaker Pattern ===")
	
	// Create circuit breaker: 3 failures, 2 second timeout
	cb := NewCircuitBreaker(3, 2*time.Second)
	
	// Simulate service calls
	scenarios := []bool{
		true,  // fail
		true,  // fail
		true,  // fail - should open circuit
		true,  // fail - circuit open
		true,  // fail - circuit open
		false, // success - but circuit still open
		false, // success - but circuit still open
		false, // success - should close circuit
		false, // success - circuit closed
	}
	
	for i, shouldFail := range scenarios {
		fmt.Printf("\nCall %d: ", i+1)
		
		err := cb.Call(func() error {
			return simulateServiceCall(shouldFail)
		})
		
		if err != nil {
			fmt.Printf("FAILED - %v\n", err)
		} else {
			fmt.Printf("SUCCESS\n")
		}
		
		state := cb.GetState()
		var stateStr string
		switch state {
		case StateClosed:
			stateStr = "CLOSED"
		case StateOpen:
			stateStr = "OPEN"
		case StateHalfOpen:
			stateStr = "HALF-OPEN"
		}
		fmt.Printf("Circuit state: %s\n", stateStr)
		
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Println("\nCircuit breaker example completed!")
}
