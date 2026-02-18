package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 4: Select-based Circuit Breaker
type SelectCircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	closed      int32 // 0: open, 1: closed
	mu          sync.RWMutex
	stateCh     chan int32
	requestCh   chan func() error
	resultCh    chan error
}

func NewSelectCircuitBreaker(threshold int64, timeout time.Duration) *SelectCircuitBreaker {
	cb := &SelectCircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
		stateCh:   make(chan int32, 1),
		requestCh: make(chan func() error, 100),
		resultCh:  make(chan error, 100),
	}
	
	cb.stateCh <- 0 // Initial state
	go cb.run()
	return cb
}

func (cb *SelectCircuitBreaker) run() {
	for {
		select {
		case state := <-cb.stateCh:
			atomic.StoreInt32(&cb.state, state)
			fmt.Printf("Circuit breaker state changed to: %d\n", state)
		case fn := <-cb.requestCh:
			cb.handleRequest(fn)
		}
		
		// Check if closed
		if atomic.LoadInt32(&cb.closed) == 1 {
			return
		}
	}
}

func (cb *SelectCircuitBreaker) handleRequest(fn func() error) {
	state := atomic.LoadInt32(&cb.state)
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.stateCh <- 2 // half-open
			cb.executeRequest(fn)
		} else {
			cb.resultCh <- fmt.Errorf("circuit breaker is open")
		}
	case 2: // half-open
		cb.executeRequest(fn)
	default: // closed
		cb.executeRequest(fn)
	}
}

func (cb *SelectCircuitBreaker) executeRequest(fn func() error) {
	if fn == nil {
		select {
		case cb.resultCh <- fmt.Errorf("nil function"):
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
		return
	}
	
	err := fn()
	
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.threshold {
			select {
			case cb.stateCh <- 1: // open
			default:
			}
		}
		select {
		case cb.resultCh <- err:
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
	} else {
		// Success
		cb.failures = 0
		select {
		case cb.stateCh <- 0: // closed
		default:
		}
		select {
		case cb.resultCh <- nil:
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
	}
}

func (cb *SelectCircuitBreaker) Call(fn func() error) error {
	if atomic.LoadInt32(&cb.closed) == 1 {
		return fmt.Errorf("circuit breaker is closed")
	}
	
	select {
	case cb.requestCh <- fn:
		return <-cb.resultCh
	case <-time.After(1 * time.Second):
		return fmt.Errorf("circuit breaker timeout")
	}
}

func (cb *SelectCircuitBreaker) Close() {
	atomic.StoreInt32(&cb.closed, 1)
	close(cb.stateCh)
	close(cb.requestCh)
	close(cb.resultCh)
}