package patterns

import (
	"sync"
	"time"
)

// Advanced Pattern 5: Condition Variable with Timeout
type TimeoutCond struct {
	mu    sync.Mutex
	cond  *sync.Cond
	ready bool
}

func NewTimeoutCond() *TimeoutCond {
	tc := &TimeoutCond{}
	tc.cond = sync.NewCond(&tc.mu)
	return tc
}

func (tc *TimeoutCond) Wait() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	for !tc.ready {
		tc.cond.Wait()
	}
}

func (tc *TimeoutCond) WaitWithTimeout(timeout time.Duration) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	
	done := make(chan struct{})
	go func() {
		for !tc.ready {
			tc.cond.Wait()
		}
		close(done)
	}()
	
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (tc *TimeoutCond) Signal() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.ready = true
	tc.cond.Signal()
}

func (tc *TimeoutCond) Broadcast() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.ready = true
	tc.cond.Broadcast()
}