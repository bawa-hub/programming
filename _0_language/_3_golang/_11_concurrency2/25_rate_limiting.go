package main

import (
	"fmt"
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

func main() {
	fmt.Println("=== Rate Limiting Pattern ===")
	
	// Create rate limiter: 2 operations per second
	limiter := NewRateLimiter(2, 1*time.Second)
	defer limiter.Stop()
	
	// Simulate API calls
	for i := 1; i <= 10; i++ {
		fmt.Printf("Attempting API call %d...\n", i)
		
		if limiter.Allow() {
			fmt.Printf("API call %d: SUCCESS\n", i)
			// Simulate API work
			time.Sleep(100 * time.Millisecond)
		} else {
			fmt.Printf("API call %d: RATE LIMITED\n", i)
			// Wait for next available slot
			limiter.Wait()
			fmt.Printf("API call %d: SUCCESS (after waiting)\n", i)
		}
	}
	
	fmt.Println("Rate limiting example completed!")
}
