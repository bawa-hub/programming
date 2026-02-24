package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	rate       float64       // tokens per second
	capacity   float64       // max tokens
	tokens     float64       // current tokens
	lastRefill time.Time
	mu         sync.Mutex
}

func NewRateLimiter(rate float64, capacity float64) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()

	// Refill tokens
	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}

	rl.lastRefill = now

	if rl.tokens >= 1 {
		rl.tokens -= 1
		return true
	}

	return false
}

func main() {
	limiter := NewRateLimiter(2, 5) // 2 req/sec, burst 5

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Allowed")
		} else {
			fmt.Println("Blocked")
		}
		time.Sleep(300 * time.Millisecond)
	}
}