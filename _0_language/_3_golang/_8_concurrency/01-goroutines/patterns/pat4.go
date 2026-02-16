package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 4: Goroutine with Circuit Breaker
type CircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
}

func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.RLock()
	state := atomic.LoadInt32(&cb.state)
	cb.mu.RUnlock()
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			atomic.StoreInt32(&cb.state, 2) // half-open
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.threshold {
			atomic.StoreInt32(&cb.state, 1) // open
		}
		return err
	}
	
	// Success
	cb.failures = 0
	atomic.StoreInt32(&cb.state, 0) // closed
	return nil
}