package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 3: Channel-based Circuit Breaker
type CircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
	stateCh     chan int32
}

func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	cb := &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
		stateCh:   make(chan int32, 1),
	}
	
	cb.stateCh <- 0 // Initial state
	go cb.monitor()
	return cb
}

func (cb *CircuitBreaker) monitor() {
	for state := range cb.stateCh {
		atomic.StoreInt32(&cb.state, state)
		fmt.Printf("Circuit breaker state: %d\n", state)
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.RLock()
	state := atomic.LoadInt32(&cb.state)
	cb.mu.RUnlock()
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.stateCh <- 2 // half-open
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
			cb.stateCh <- 1 // open
		}
		return err
	}
	
	// Success
	cb.failures = 0
	cb.stateCh <- 0 // closed
	return nil
}

func (cb *CircuitBreaker) Close() {
	close(cb.stateCh)
}