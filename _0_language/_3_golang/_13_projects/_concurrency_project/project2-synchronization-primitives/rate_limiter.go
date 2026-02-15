package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter limits the number of requests per time window
type RateLimiter struct {
	requests   int
	window     time.Duration
	requestsCh chan time.Time
	mutex      sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(requests int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests:   requests,
		window:     window,
		requestsCh: make(chan time.Time, requests*2), // Buffer for burst
	}
	
	// Start cleanup goroutine
	go rl.cleanup()
	
	return rl
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	
	// Check if we have space in the current window
	select {
	case rl.requestsCh <- now:
		return true
	default:
		return false
	}
}

// cleanup removes old requests from the window
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window / 10) // Check 10 times per window
	defer ticker.Stop()
	
	for range ticker.C {
		rl.mutex.Lock()
		cutoff := time.Now().Add(-rl.window)
		
		// Remove old requests
		for {
			select {
			case reqTime := <-rl.requestsCh:
				if reqTime.After(cutoff) {
					// Put it back if it's still valid
					select {
					case rl.requestsCh <- reqTime:
					default:
						// Channel is full, request is too old
					}
					break
				}
			default:
				// No more requests to check
				break
			}
		}
		rl.mutex.Unlock()
	}
}

// GetStats returns current rate limiter statistics
func (rl *RateLimiter) GetStats() (int, int) {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	current := len(rl.requestsCh)
	return current, rl.requests
}

// AtomicRateLimiter uses atomic operations for better performance
type AtomicRateLimiter struct {
	requests   int64
	window     time.Duration
	tokens     int64
	lastUpdate time.Time
	mutex      sync.Mutex
}

// NewAtomicRateLimiter creates a new atomic rate limiter
func NewAtomicRateLimiter(requests int64, window time.Duration) *AtomicRateLimiter {
	return &AtomicRateLimiter{
		requests:   requests,
		window:     window,
		tokens:     requests,
		lastUpdate: time.Now(),
	}
}

// Allow checks if a request is allowed (atomic version)
func (arl *AtomicRateLimiter) Allow() bool {
	arl.mutex.Lock()
	defer arl.mutex.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(arl.lastUpdate)
	
	// Add tokens based on elapsed time
	tokensToAdd := int64(elapsed / (arl.window / time.Duration(arl.requests)))
	if tokensToAdd > 0 {
		arl.tokens = arl.requests // Reset to max
		arl.lastUpdate = now
	}
	
	if arl.tokens > 0 {
		arl.tokens--
		return true
	}
	
	return false
}

// GetStats returns current atomic rate limiter statistics
func (arl *AtomicRateLimiter) GetStats() (int64, int64) {
	arl.mutex.Lock()
	defer arl.mutex.Unlock()
	
	return arl.tokens, arl.requests
}

// DemonstrateRateLimiter demonstrates the rate limiter
func DemonstrateRateLimiter() {
	fmt.Println("=== Rate Limiter Demonstration ===")
	
	// Create rate limiter: 5 requests per second
	rl := NewRateLimiter(5, time.Second)
	
	fmt.Println("Testing rate limiter (5 requests per second):")
	
	// Try to make 10 requests quickly
	for i := 0; i < 10; i++ {
		allowed := rl.Allow()
		current, max := rl.GetStats()
		fmt.Printf("Request %d: %v (tokens: %d/%d)\n", i+1, allowed, current, max)
		time.Sleep(100 * time.Millisecond)
	}
	
	// Wait for window to reset
	fmt.Println("\nWaiting for window to reset...")
	time.Sleep(2 * time.Second)
	
	// Try again
	fmt.Println("Trying again after reset:")
	for i := 0; i < 5; i++ {
		allowed := rl.Allow()
		current, max := rl.GetStats()
		fmt.Printf("Request %d: %v (tokens: %d/%d)\n", i+1, allowed, current, max)
		time.Sleep(100 * time.Millisecond)
	}
}

// DemonstrateAtomicRateLimiter demonstrates the atomic rate limiter
func DemonstrateAtomicRateLimiter() {
	fmt.Println("\n=== Atomic Rate Limiter Demonstration ===")
	
	// Create atomic rate limiter: 3 requests per second
	arl := NewAtomicRateLimiter(3, time.Second)
	
	fmt.Println("Testing atomic rate limiter (3 requests per second):")
	
	// Try to make 8 requests quickly
	for i := 0; i < 8; i++ {
		allowed := arl.Allow()
		tokens, max := arl.GetStats()
		fmt.Printf("Request %d: %v (tokens: %d/%d)\n", i+1, allowed, tokens, max)
		time.Sleep(200 * time.Millisecond)
	}
}

// BurstRateLimiter allows burst requests up to a limit
type BurstRateLimiter struct {
	requests   int
	window     time.Duration
	tokens     int
	lastRefill time.Time
	mutex      sync.Mutex
}

// NewBurstRateLimiter creates a new burst rate limiter
func NewBurstRateLimiter(requests int, window time.Duration) *BurstRateLimiter {
	return &BurstRateLimiter{
		requests:   requests,
		window:     window,
		tokens:     requests,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request is allowed (burst version)
func (brl *BurstRateLimiter) Allow() bool {
	brl.mutex.Lock()
	defer brl.mutex.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(brl.lastRefill)
	
	// Refill tokens based on elapsed time
	tokensToAdd := int(elapsed / (brl.window / time.Duration(brl.requests)))
	if tokensToAdd > 0 {
		brl.tokens = brl.requests // Allow full burst
		brl.lastRefill = now
	}
	
	if brl.tokens > 0 {
		brl.tokens--
		return true
	}
	
	return false
}

// GetStats returns current burst rate limiter statistics
func (brl *BurstRateLimiter) GetStats() (int, int) {
	brl.mutex.Lock()
	defer brl.mutex.Unlock()
	
	return brl.tokens, brl.requests
}

// DemonstrateBurstRateLimiter demonstrates the burst rate limiter
func DemonstrateBurstRateLimiter() {
	fmt.Println("\n=== Burst Rate Limiter Demonstration ===")
	
	// Create burst rate limiter: 2 requests per second, but allows burst
	brl := NewBurstRateLimiter(2, time.Second)
	
	fmt.Println("Testing burst rate limiter (2 requests per second, allows burst):")
	
	// Try to make 5 requests quickly (should allow burst)
	for i := 0; i < 5; i++ {
		allowed := brl.Allow()
		tokens, max := brl.GetStats()
		fmt.Printf("Request %d: %v (tokens: %d/%d)\n", i+1, allowed, tokens, max)
		time.Sleep(50 * time.Millisecond)
	}
	
	// Wait and try again
	fmt.Println("\nWaiting for refill...")
	time.Sleep(1 * time.Second)
	
	fmt.Println("Trying again after refill:")
	for i := 0; i < 3; i++ {
		allowed := brl.Allow()
		tokens, max := brl.GetStats()
		fmt.Printf("Request %d: %v (tokens: %d/%d)\n", i+1, allowed, tokens, max)
		time.Sleep(100 * time.Millisecond)
	}
}
