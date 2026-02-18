package patterns

import "time"

// Advanced Pattern 2: Channel-based Rate Limiter
type RateLimiter struct {
	tokens   chan struct{}
	rate     time.Duration
	burst    int
	stopCh   chan bool
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	rl := &RateLimiter{
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

func (rl *RateLimiter) refill() {
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

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Stop() {
	close(rl.stopCh)
}