package ratelimit

import (
	"sync"
	"time"
)

// Limiter represents a rate limiter
type Limiter struct {
	limit    int
	interval time.Duration
	tokens   int
	lastTime time.Time
	mu       sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, interval time.Duration) *Limiter {
	return &Limiter{
		limit:    limit,
		interval: interval,
		tokens:   limit,
		lastTime: time.Now(),
	}
}

// Allow checks if a request is allowed
func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	
	// Add tokens based on elapsed time
	elapsed := now.Sub(l.lastTime)
	tokensToAdd := int(elapsed / l.interval)
	
	if tokensToAdd > 0 {
		l.tokens = min(l.limit, l.tokens+tokensToAdd)
		l.lastTime = now
	}
	
	// Check if we have tokens available
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	
	return false
}

// AllowN checks if n requests are allowed
func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	
	// Add tokens based on elapsed time
	elapsed := now.Sub(l.lastTime)
	tokensToAdd := int(elapsed / l.interval)
	
	if tokensToAdd > 0 {
		l.tokens = min(l.limit, l.tokens+tokensToAdd)
		l.lastTime = now
	}
	
	// Check if we have enough tokens
	if l.tokens >= n {
		l.tokens -= n
		return true
	}
	
	return false
}

// Reserve reserves a token
func (l *Limiter) Reserve() *Reservation {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	
	// Add tokens based on elapsed time
	elapsed := now.Sub(l.lastTime)
	tokensToAdd := int(elapsed / l.interval)
	
	if tokensToAdd > 0 {
		l.tokens = min(l.limit, l.tokens+tokensToAdd)
		l.lastTime = now
	}
	
	if l.tokens > 0 {
		l.tokens--
		return &Reservation{
			ok:      true,
			timeToAct: now,
		}
	}
	
	// Calculate when the next token will be available
	nextTokenTime := l.lastTime.Add(l.interval)
	return &Reservation{
		ok:        false,
		timeToAct: nextTokenTime,
	}
}

// Wait waits for a token to become available
func (l *Limiter) Wait() {
	reservation := l.Reserve()
	if !reservation.ok {
		time.Sleep(time.Until(reservation.timeToAct))
	}
}

// WaitN waits for n tokens to become available
func (l *Limiter) WaitN(n int) {
	for i := 0; i < n; i++ {
		l.Wait()
	}
}

// GetLimit returns the current limit
func (l *Limiter) GetLimit() int {
	return l.limit
}

// GetInterval returns the current interval
func (l *Limiter) GetInterval() time.Duration {
	return l.interval
}

// GetTokens returns the current number of tokens
func (l *Limiter) GetTokens() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(l.lastTime)
	tokensToAdd := int(elapsed / l.interval)
	
	if tokensToAdd > 0 {
		l.tokens = min(l.limit, l.tokens+tokensToAdd)
		l.lastTime = now
	}
	
	return l.tokens
}

// SetLimit sets the rate limit
func (l *Limiter) SetLimit(limit int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	l.limit = limit
	l.tokens = min(l.tokens, limit)
}

// SetInterval sets the interval
func (l *Limiter) SetInterval(interval time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	l.interval = interval
}

// Reservation represents a rate limit reservation
type Reservation struct {
	ok        bool
	timeToAct time.Time
}

// OK returns true if the reservation was successful
func (r *Reservation) OK() bool {
	return r.ok
}

// DelayFrom returns the delay from the given time
func (r *Reservation) DelayFrom(now time.Time) time.Duration {
	return r.timeToAct.Sub(now)
}

// TokenBucket represents a token bucket rate limiter
type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int
	lastRefill   time.Time
	mu           sync.Mutex
}

// NewTokenBucket creates a new token bucket
func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request is allowed
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	now := time.Now()
	
	// Add tokens based on elapsed time
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate
	
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
	
	// Check if we have tokens available
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	
	return false
}

// AllowN checks if n requests are allowed
func (tb *TokenBucket) AllowN(n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	now := time.Now()
	
	// Add tokens based on elapsed time
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate
	
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
	
	// Check if we have enough tokens
	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}
	
	return false
}

// GetTokens returns the current number of tokens
func (tb *TokenBucket) GetTokens() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate
	
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
	
	return tb.tokens
}

// SetCapacity sets the bucket capacity
func (tb *TokenBucket) SetCapacity(capacity int) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	tb.capacity = capacity
	tb.tokens = min(tb.tokens, capacity)
}

// SetRefillRate sets the refill rate
func (tb *TokenBucket) SetRefillRate(rate int) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	tb.refillRate = rate
}

// LeakyBucket represents a leaky bucket rate limiter
type LeakyBucket struct {
	capacity     int
	tokens       int
	leakRate     int
	lastLeak     time.Time
	mu           sync.Mutex
}

// NewLeakyBucket creates a new leaky bucket
func NewLeakyBucket(capacity, leakRate int) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		tokens:   0,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

// Allow checks if a request is allowed
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	now := time.Now()
	
	// Leak tokens based on elapsed time
	elapsed := now.Sub(lb.lastLeak)
	tokensToLeak := int(elapsed.Seconds()) * lb.leakRate
	
	if tokensToLeak > 0 {
		lb.tokens = max(0, lb.tokens-tokensToLeak)
		lb.lastLeak = now
	}
	
	// Check if we can add a token
	if lb.tokens < lb.capacity {
		lb.tokens++
		return true
	}
	
	return false
}

// AllowN checks if n requests are allowed
func (lb *LeakyBucket) AllowN(n int) bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	now := time.Now()
	
	// Leak tokens based on elapsed time
	elapsed := now.Sub(lb.lastLeak)
	tokensToLeak := int(elapsed.Seconds()) * lb.leakRate
	
	if tokensToLeak > 0 {
		lb.tokens = max(0, lb.tokens-tokensToLeak)
		lb.lastLeak = now
	}
	
	// Check if we can add n tokens
	if lb.tokens+n <= lb.capacity {
		lb.tokens += n
		return true
	}
	
	return false
}

// GetTokens returns the current number of tokens
func (lb *LeakyBucket) GetTokens() int {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(lb.lastLeak)
	tokensToLeak := int(elapsed.Seconds()) * lb.leakRate
	
	if tokensToLeak > 0 {
		lb.tokens = max(0, lb.tokens-tokensToLeak)
		lb.lastLeak = now
	}
	
	return lb.tokens
}

// SetCapacity sets the bucket capacity
func (lb *LeakyBucket) SetCapacity(capacity int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	lb.capacity = capacity
	lb.tokens = min(lb.tokens, capacity)
}

// SetLeakRate sets the leak rate
func (lb *LeakyBucket) SetLeakRate(rate int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	lb.leakRate = rate
}

// SlidingWindow represents a sliding window rate limiter
type SlidingWindow struct {
	windowSize time.Duration
	requests   []time.Time
	mu         sync.Mutex
}

// NewSlidingWindow creates a new sliding window rate limiter
func NewSlidingWindow(windowSize time.Duration) *SlidingWindow {
	return &SlidingWindow{
		windowSize: windowSize,
		requests:   make([]time.Time, 0),
	}
}

// Allow checks if a request is allowed
func (sw *SlidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	
	now := time.Now()
	
	// Remove old requests outside the window
	cutoff := now.Add(-sw.windowSize)
	for len(sw.requests) > 0 && sw.requests[0].Before(cutoff) {
		sw.requests = sw.requests[1:]
	}
	
	// Add current request
	sw.requests = append(sw.requests, now)
	
	return true
}

// GetRequestCount returns the number of requests in the current window
func (sw *SlidingWindow) GetRequestCount() int {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	
	now := time.Now()
	cutoff := now.Add(-sw.windowSize)
	
	count := 0
	for _, request := range sw.requests {
		if request.After(cutoff) {
			count++
		}
	}
	
	return count
}

// SetWindowSize sets the window size
func (sw *SlidingWindow) SetWindowSize(windowSize time.Duration) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	
	sw.windowSize = windowSize
}

// Helper functions

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
