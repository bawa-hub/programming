package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// ContextPool manages a pool of contexts
type ContextPool struct {
	contexts chan context.Context
	factory  func() context.Context
	mutex    sync.RWMutex
}

// ContextMiddleware represents a context middleware function
type ContextMiddleware func(context.Context) context.Context

// ContextChain represents a chain of context middlewares
type ContextChain struct {
	middlewares []ContextMiddleware
}

// ContextMetrics tracks context usage metrics
type ContextMetrics struct {
	Created     int64
	Cancelled   int64
	TimedOut    int64
	ValueLookups int64
	Errors      int64
}

// ContextTracer provides tracing for context operations
type ContextTracer struct {
	operations []string
	mutex      sync.Mutex
}

// ContextValidator validates context state
type ContextValidator struct {
	rules []func(context.Context) error
}

// ContextCache caches context values
type ContextCache struct {
	cache map[string]interface{}
	mutex sync.RWMutex
}

// ContextRateLimiter limits context operations
type ContextRateLimiter struct {
	limit    int
	interval time.Duration
	tokens   int64
	lastTime time.Time
	mutex    sync.Mutex
}

// CircuitBreaker implements circuit breaker pattern
type CircuitBreaker struct {
	failureCount    int64
	successCount    int64
	threshold       int64
	timeout         time.Duration
	state           CircuitState
	lastFailureTime time.Time
	mutex           sync.RWMutex
}

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// NewContextPool creates a new context pool
func NewContextPool(size int, factory func() context.Context) *ContextPool {
	pool := &ContextPool{
		contexts: make(chan context.Context, size),
		factory:  factory,
	}
	
	// Pre-populate pool
	for i := 0; i < size; i++ {
		pool.contexts <- factory()
	}
	
	return pool
}

// Get gets a context from the pool
func (p *ContextPool) Get() context.Context {
	select {
	case ctx := <-p.contexts:
		return ctx
	default:
		return p.factory()
	}
}

// Put puts a context back into the pool
func (p *ContextPool) Put(ctx context.Context) {
	select {
	case p.contexts <- ctx:
	default:
		// Pool is full, discard context
	}
}

// NewContextChain creates a new context chain
func NewContextChain() *ContextChain {
	return &ContextChain{
		middlewares: make([]ContextMiddleware, 0),
	}
}

// Add adds a middleware to the chain
func (c *ContextChain) Add(middleware ContextMiddleware) {
	c.middlewares = append(c.middlewares, middleware)
}

// Execute executes the chain on a context
func (c *ContextChain) Execute(ctx context.Context) context.Context {
	for _, middleware := range c.middlewares {
		ctx = middleware(ctx)
	}
	return ctx
}

// NewContextMetrics creates new context metrics
func NewContextMetrics() *ContextMetrics {
	return &ContextMetrics{}
}

// RecordCreation records a context creation
func (m *ContextMetrics) RecordCreation() {
	atomic.AddInt64(&m.Created, 1)
}

// RecordCancellation records a context cancellation
func (m *ContextMetrics) RecordCancellation() {
	atomic.AddInt64(&m.Cancelled, 1)
}

// RecordTimeout records a context timeout
func (m *ContextMetrics) RecordTimeout() {
	atomic.AddInt64(&m.TimedOut, 1)
}

// RecordValueLookup records a value lookup
func (m *ContextMetrics) RecordValueLookup() {
	atomic.AddInt64(&m.ValueLookups, 1)
}

// RecordError records an error
func (m *ContextMetrics) RecordError() {
	atomic.AddInt64(&m.Errors, 1)
}

// GetStats returns current metrics
func (m *ContextMetrics) GetStats() map[string]int64 {
	return map[string]int64{
		"created":      atomic.LoadInt64(&m.Created),
		"cancelled":    atomic.LoadInt64(&m.Cancelled),
		"timed_out":    atomic.LoadInt64(&m.TimedOut),
		"value_lookups": atomic.LoadInt64(&m.ValueLookups),
		"errors":       atomic.LoadInt64(&m.Errors),
	}
}

// NewContextTracer creates a new context tracer
func NewContextTracer() *ContextTracer {
	return &ContextTracer{
		operations: make([]string, 0),
	}
}

// Trace records an operation
func (t *ContextTracer) Trace(operation string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	
	t.operations = append(t.operations, fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), operation))
}

// GetOperations returns all operations
func (t *ContextTracer) GetOperations() []string {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	
	return append([]string(nil), t.operations...)
}

// NewContextValidator creates a new context validator
func NewContextValidator() *ContextValidator {
	return &ContextValidator{
		rules: make([]func(context.Context) error, 0),
	}
}

// AddRule adds a validation rule
func (v *ContextValidator) AddRule(rule func(context.Context) error) {
	v.rules = append(v.rules, rule)
}

// Validate validates a context
func (v *ContextValidator) Validate(ctx context.Context) error {
	for _, rule := range v.rules {
		if err := rule(ctx); err != nil {
			return err
		}
	}
	return nil
}

// NewContextCache creates a new context cache
func NewContextCache() *ContextCache {
	return &ContextCache{
		cache: make(map[string]interface{}),
	}
}

// Get gets a value from cache
func (c *ContextCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	value, exists := c.cache[key]
	return value, exists
}

// Set sets a value in cache
func (c *ContextCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.cache[key] = value
}

// NewContextRateLimiter creates a new rate limiter
func NewContextRateLimiter(limit int, interval time.Duration) *ContextRateLimiter {
	return &ContextRateLimiter{
		limit:    limit,
		interval: interval,
		tokens:   int64(limit),
		lastTime: time.Now(),
	}
}

// Allow checks if an operation is allowed
func (r *ContextRateLimiter) Allow() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(r.lastTime)
	
	// Add tokens based on elapsed time
	tokensToAdd := int64(elapsed / r.interval)
	if tokensToAdd > 0 {
		r.tokens = min(r.tokens+tokensToAdd, int64(r.limit))
		r.lastTime = now
	}
	
	if r.tokens > 0 {
		r.tokens--
		return true
	}
	
	return false
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// GetState returns the current circuit breaker state
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// RecordSuccess records a successful operation
func (cb *CircuitBreaker) RecordSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.successCount++
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failureCount = 0
	}
}

// RecordFailure records a failed operation
func (cb *CircuitBreaker) RecordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	
	if cb.failureCount >= cb.threshold {
		cb.state = StateOpen
	}
}

// ShouldAllow checks if the operation should be allowed
func (cb *CircuitBreaker) ShouldAllow() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
			return true
		}
		return false
	}
	
	return true
}

// Advanced Pattern 1: Context Pool
func contextPoolExample() {
	fmt.Println("\n1. Context Pool")
	fmt.Println("===============")
	
	// Create context pool
	factory := func() context.Context {
		return context.WithValue(context.Background(), "pooled", true)
	}
	pool := NewContextPool(5, factory)
	
	// Use contexts from pool
	for i := 0; i < 10; i++ {
		ctx := pool.Get()
		fmt.Printf("  Got context: %v\n", ctx.Value("pooled"))
		
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		
		// Return context to pool
		pool.Put(ctx)
	}
	
	fmt.Println("Context pool example completed")
}

// Advanced Pattern 2: Context Middleware Chain
func contextMiddlewareChainExample() {
	fmt.Println("\n2. Context Middleware Chain")
	fmt.Println("===========================")
	
	// Create middleware chain
	chain := NewContextChain()
	
	// Add middlewares
	chain.Add(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "middleware1", "processed")
	})
	
	chain.Add(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "middleware2", "processed")
	})
	
	chain.Add(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "middleware3", "processed")
	})
	
	// Execute chain
	ctx := context.Background()
	ctx = chain.Execute(ctx)
	
	// Check results
	fmt.Printf("  Middleware 1: %v\n", ctx.Value("middleware1"))
	fmt.Printf("  Middleware 2: %v\n", ctx.Value("middleware2"))
	fmt.Printf("  Middleware 3: %v\n", ctx.Value("middleware3"))
	
	fmt.Println("Context middleware chain example completed")
}

// Advanced Pattern 3: Context Metrics
func contextMetricsExample() {
	fmt.Println("\n3. Context Metrics")
	fmt.Println("==================")
	
	metrics := NewContextMetrics()
	
	// Simulate context operations
	for i := 0; i < 100; i++ {
		metrics.RecordCreation()
		
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(5))*time.Second)
		
		// Simulate work
		go func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			cancel()
		}()
		
		// Simulate value lookups
		for j := 0; j < 10; j++ {
			metrics.RecordValueLookup()
			_ = ctx.Value("key")
		}
		
		// Simulate errors
		if rand.Float32() < 0.1 {
			metrics.RecordError()
		}
		
		// Wait for context to be cancelled
		<-ctx.Done()
		
		if ctx.Err() == context.DeadlineExceeded {
			metrics.RecordTimeout()
		} else {
			metrics.RecordCancellation()
		}
	}
	
	// Display metrics
	stats := metrics.GetStats()
	fmt.Printf("  Created: %d\n", stats["created"])
	fmt.Printf("  Cancelled: %d\n", stats["cancelled"])
	fmt.Printf("  Timed out: %d\n", stats["timed_out"])
	fmt.Printf("  Value lookups: %d\n", stats["value_lookups"])
	fmt.Printf("  Errors: %d\n", stats["errors"])
	
	fmt.Println("Context metrics example completed")
}

// Advanced Pattern 4: Context Tracing
func contextTracingExample() {
	fmt.Println("\n4. Context Tracing")
	fmt.Println("==================")
	
	tracer := NewContextTracer()
	
	// Create context with tracing
	ctx := context.Background()
	ctx = context.WithValue(ctx, "tracer", tracer)
	
	// Simulate operations
	processWithTracing(ctx, "operation1")
	processWithTracing(ctx, "operation2")
	processWithTracing(ctx, "operation3")
	
	// Display trace
	operations := tracer.GetOperations()
	fmt.Println("  Trace operations:")
	for _, op := range operations {
		fmt.Printf("    %s\n", op)
	}
	
	fmt.Println("Context tracing example completed")
}

func processWithTracing(ctx context.Context, operation string) {
	tracer := ctx.Value("tracer").(*ContextTracer)
	tracer.Trace(fmt.Sprintf("Starting %s", operation))
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	tracer.Trace(fmt.Sprintf("Completed %s", operation))
}

// Advanced Pattern 5: Context Validation
func contextValidationExample() {
	fmt.Println("\n5. Context Validation")
	fmt.Println("=====================")
	
	validator := NewContextValidator()
	
	// Add validation rules
	validator.AddRule(func(ctx context.Context) error {
		if ctx.Value("userID") == nil {
			return fmt.Errorf("userID is required")
		}
		return nil
	})
	
	validator.AddRule(func(ctx context.Context) error {
		if ctx.Value("requestID") == nil {
			return fmt.Errorf("requestID is required")
		}
		return nil
	})
	
	validator.AddRule(func(ctx context.Context) error {
		if ctx.Err() != nil {
			return fmt.Errorf("context is cancelled: %v", ctx.Err())
		}
		return nil
	})
	
	// Test valid context
	validCtx := context.Background()
	validCtx = context.WithValue(validCtx, "userID", "user-123")
	validCtx = context.WithValue(validCtx, "requestID", "req-456")
	
	if err := validator.Validate(validCtx); err != nil {
		fmt.Printf("  Validation failed: %v\n", err)
	} else {
		fmt.Println("  Validation passed")
	}
	
	// Test invalid context
	invalidCtx := context.Background()
	invalidCtx = context.WithValue(invalidCtx, "userID", "user-123")
	// Missing requestID
	
	if err := validator.Validate(invalidCtx); err != nil {
		fmt.Printf("  Validation failed as expected: %v\n", err)
	}
	
	fmt.Println("Context validation example completed")
}

// Advanced Pattern 6: Context Cache
func contextCacheExample() {
	fmt.Println("\n6. Context Cache")
	fmt.Println("================")
	
	cache := NewContextCache()
	
	// Cache some values
	cache.Set("user:123", User{ID: "123", Email: "john@example.com", Role: "admin"})
	cache.Set("order:456", Order{ID: "456", UserID: "123", Amount: 99.99, Status: "pending"})
	
	// Retrieve values
	if user, exists := cache.Get("user:123"); exists {
		fmt.Printf("  Cached user: %+v\n", user)
	}
	
	if order, exists := cache.Get("order:456"); exists {
		fmt.Printf("  Cached order: %+v\n", order)
	}
	
	// Non-existent key
	if _, exists := cache.Get("user:999"); !exists {
		fmt.Println("  User 999 not found in cache")
	}
	
	fmt.Println("Context cache example completed")
}

// Advanced Pattern 7: Context Rate Limiting
func contextRateLimitingExample() {
	fmt.Println("\n7. Context Rate Limiting")
	fmt.Println("========================")
	
	// Create rate limiter (5 operations per second)
	limiter := NewContextRateLimiter(5, time.Second)
	
	// Simulate operations
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("  Operation %d allowed\n", i+1)
		} else {
			fmt.Printf("  Operation %d rate limited\n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}
	
	fmt.Println("Context rate limiting example completed")
}

// Advanced Pattern 8: Context with Circuit Breaker
func contextCircuitBreakerExample() {
	fmt.Println("\n8. Context with Circuit Breaker")
	fmt.Println("===============================")
	
	// Simulate circuit breaker
	circuitBreaker := &CircuitBreaker{
		failureCount: 0,
		threshold:    3,
		timeout:      5 * time.Second,
		state:        StateClosed,
	}
	
	// Simulate operations
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		
		if circuitBreaker.ShouldAllow() {
			fmt.Printf("  Operation %d: Circuit breaker allows\n", i+1)
			
			// Simulate work
			go func(id int) {
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				cancel()
			}(i)
			
			// Wait for completion or timeout
			<-ctx.Done()
			
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("  Operation %d: Timeout\n", i+1)
				circuitBreaker.RecordFailure()
			} else {
				fmt.Printf("  Operation %d: Success\n", i+1)
				circuitBreaker.RecordSuccess()
			}
		} else {
			fmt.Printf("  Operation %d: Circuit breaker blocks\n", i+1)
			cancel() // Cancel the context since we're not using it
		}
		
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Println("Context circuit breaker example completed")
}

// Advanced Pattern 9: Context with Retry Logic
func contextRetryLogicExample() {
	fmt.Println("\n9. Context with Retry Logic")
	fmt.Println("===========================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Simulate operation with retry
	err := advancedOperationWithRetry(ctx, 3, 1*time.Second)
	if err != nil {
		fmt.Printf("  Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("  Operation succeeded")
	}
	
	fmt.Println("Context retry logic example completed")
}

func advancedOperationWithRetry(ctx context.Context, maxRetries int, backoff time.Duration) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("operation cancelled: %v", ctx.Err())
		default:
			fmt.Printf("  Attempt %d\n", attempt)
			
			// Simulate operation that might fail
			if rand.Float32() < 0.6 { // 60% chance of failure
				fmt.Printf("  Attempt %d failed\n", attempt)
				if attempt < maxRetries {
					// Wait before retry
					select {
					case <-ctx.Done():
						return fmt.Errorf("operation cancelled during backoff: %v", ctx.Err())
					case <-time.After(backoff * time.Duration(attempt)):
						continue
					}
				}
				return fmt.Errorf("operation failed after %d attempts", maxRetries)
			}
			
			fmt.Printf("  Attempt %d succeeded\n", attempt)
			return nil
		}
	}
	return nil
}

// Advanced Pattern 10: Context with Load Balancing
func contextLoadBalancingExample() {
	fmt.Println("\n10. Context with Load Balancing")
	fmt.Println("===============================")
	
	// Simulate load balancer
	loadBalancer := &LoadBalancer{
		services: []string{"service1", "service2", "service3"},
		current:  0,
	}
	
	// Simulate requests
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		
		service := loadBalancer.GetNextService()
		fmt.Printf("  Request %d: Using %s\n", i+1, service)
		
		// Simulate service call
		go func(id int, svc string) {
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			cancel()
		}(i, service)
		
		// Wait for completion
		<-ctx.Done()
		
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("  Request %d: Timeout\n", i+1)
		} else {
			fmt.Printf("  Request %d: Completed\n", i+1)
		}
		
		time.Sleep(200 * time.Millisecond)
	}
	
	fmt.Println("Context load balancing example completed")
}

// LoadBalancer simulates a load balancer
type LoadBalancer struct {
	services []string
	current  int
	mutex    sync.Mutex
}

// GetNextService returns the next service
func (lb *LoadBalancer) GetNextService() string {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	service := lb.services[lb.current]
	lb.current = (lb.current + 1) % len(lb.services)
	return service
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Context Patterns")
	fmt.Println("============================")
	
	contextPoolExample()
	contextMiddlewareChainExample()
	contextMetricsExample()
	contextTracingExample()
	contextValidationExample()
	contextCacheExample()
	contextRateLimitingExample()
	contextCircuitBreakerExample()
	contextRetryLogicExample()
	contextLoadBalancingExample()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
