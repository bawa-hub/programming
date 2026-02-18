package patterns

import (
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 2: Select-based Rate Limiter
type SelectRateLimiter struct {
	tokens   chan struct{}
	rate     time.Duration
	burst    int
	stopCh   chan bool
	mu       sync.RWMutex
	requests int64
	allowed  int64
}

func NewSelectRateLimiter(rate time.Duration, burst int) *SelectRateLimiter {
	rl := &SelectRateLimiter{
		tokens: make(chan struct{}, burst),
		rate:   rate,
		burst:  burst,
		stopCh: make(chan bool),
	}
	
	// Fill initial tokens
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}
	
	go rl.refill()
	return rl
}

func (rl *SelectRateLimiter) refill() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
				// Token added
			default:
				// Bucket is full
			}
		case <-rl.stopCh:
			return
		}
	}
}

func (rl *SelectRateLimiter) Allow() bool {
	atomic.AddInt64(&rl.requests, 1)
	
	select {
	case <-rl.tokens:
		atomic.AddInt64(&rl.allowed, 1)
		return true
	default:
		return false
	}
}

func (rl *SelectRateLimiter) Stats() (requests, allowed int64) {
	return atomic.LoadInt64(&rl.requests), atomic.LoadInt64(&rl.allowed)
}

func (rl *SelectRateLimiter) Stop() {
	close(rl.stopCh)
}