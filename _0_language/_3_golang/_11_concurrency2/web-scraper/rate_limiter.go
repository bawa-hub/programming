package main

import (
	"time"
)

// RateLimiter controls the rate of operations
type RateLimiter struct {
	ticker *time.Ticker
	limit  chan struct{}
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		ticker: time.NewTicker(duration),
		limit:  make(chan struct{}, rate),
	}
	
	// Fill the limit channel initially
	for i := 0; i < rate; i++ {
		rl.limit <- struct{}{}
	}
	
	// Start the refill goroutine
	go rl.refill()
	
	return rl
}

// refill periodically refills the limit channel
func (rl *RateLimiter) refill() {
	for range rl.ticker.C {
		select {
		case rl.limit <- struct{}{}:
		default:
			// Channel is full, skip
		}
	}
}

// Allow checks if an operation is allowed
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.limit:
		return true
	default:
		return false
	}
}

// Wait blocks until an operation is allowed
func (rl *RateLimiter) Wait() {
	<-rl.limit
}

// Stop stops the rate limiter
func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}
