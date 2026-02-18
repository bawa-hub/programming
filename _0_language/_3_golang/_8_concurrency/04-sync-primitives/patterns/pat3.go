package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 3: WaitGroup with Timeout
type TimeoutWaitGroup struct {
	wg      sync.WaitGroup
	timeout time.Duration
	done    chan struct{}
}

func NewTimeoutWaitGroup(timeout time.Duration) *TimeoutWaitGroup {
	return &TimeoutWaitGroup{
		timeout: timeout,
		done:    make(chan struct{}),
	}
}

func (twg *TimeoutWaitGroup) Add(delta int) {
	twg.wg.Add(delta)
}

func (twg *TimeoutWaitGroup) Done() {
	twg.wg.Done()
}

func (twg *TimeoutWaitGroup) Wait() error {
	go func() {
		twg.wg.Wait()
		close(twg.done)
	}()
	
	select {
	case <-twg.done:
		return nil
	case <-time.After(twg.timeout):
		return fmt.Errorf("waitgroup timeout after %v", twg.timeout)
	}
}