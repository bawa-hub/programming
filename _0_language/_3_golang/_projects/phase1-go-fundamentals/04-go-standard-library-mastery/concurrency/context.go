package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Custom types for demonstration
type User struct {
	ID   int
	Name string
}

type key string

const (
	userKey    key = "user"
	requestID  key = "request_id"
	traceID    key = "trace_id"
	timeoutKey key = "timeout"
)

// Custom context with additional functionality
type customContext struct {
	context.Context
	customValue string
	mu          sync.RWMutex
}

func (c *customContext) Value(key interface{}) interface{} {
	if key == "custom" {
		c.mu.RLock()
		defer c.mu.RUnlock()
		return c.customValue
	}
	return c.Context.Value(key)
}

func (c *customContext) SetCustomValue(value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.customValue = value
}

// CircuitBreaker for demonstrating context usage
type CircuitBreaker struct {
	failures  int
	threshold int
	timeout   time.Duration
	lastFail  time.Time
	state     string // "closed", "open", "half-open"
}

func (cb *CircuitBreaker) Call(ctx context.Context, fn func() error) error {
	if cb.state == "open" {
		if time.Since(cb.lastFail) > cb.timeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	if err != nil {
		cb.failures++
		cb.lastFail = time.Now()
		if cb.failures >= cb.threshold {
			cb.state = "open"
		}
		return err
	}
	
	cb.failures = 0
	cb.state = "closed"
	return nil
}

// Service for demonstrating context usage
type Service struct {
	name string
}

func (s *Service) ProcessRequest(ctx context.Context, data string) error {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Add service name to context
	ctx = context.WithValue(ctx, "service", s.name)
	
	// Simulate processing time
	select {
	case <-time.After(100 * time.Millisecond):
		// Processing completed
		fmt.Printf("Service %s processed: %s\n", s.name, data)
		return nil
	case <-ctx.Done():
		fmt.Printf("Service %s cancelled: %v\n", s.name, ctx.Err())
		return ctx.Err()
	}
}

// Database service for demonstrating context usage
type Database struct {
	name string
}

func (db *Database) Query(ctx context.Context, query string) (string, error) {
	// Check context deadline
	if deadline, ok := ctx.Deadline(); ok {
		if time.Until(deadline) < 50*time.Millisecond {
			return "", context.DeadlineExceeded
		}
	}

	// Simulate database query
	select {
	case <-time.After(200 * time.Millisecond):
		return fmt.Sprintf("Result for query: %s", query), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// HTTP client for demonstrating context usage
type HTTPClient struct {
	timeout time.Duration
}

func (c *HTTPClient) Get(ctx context.Context, url string) (string, error) {
	// Create context with client timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	// Simulate HTTP request
	select {
	case <-time.After(150 * time.Millisecond):
		return fmt.Sprintf("Response from %s", url), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	fmt.Println("🚀 Go context Package Mastery Examples")
	fmt.Println("======================================")

	// 1. Basic Context Operations
	fmt.Println("\n1. Basic Context Operations:")
	
	// Background context
	bgCtx := context.Background()
	fmt.Printf("Background context: %v\n", bgCtx)
	
	// TODO context
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)
	
	// Context with values
	ctxWithValue := context.WithValue(bgCtx, userKey, User{ID: 1, Name: "John"})
	user := ctxWithValue.Value(userKey).(User)
	fmt.Printf("User from context: %+v\n", user)

	// 2. Context with Cancellation
	fmt.Println("\n2. Context with Cancellation:")
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Start goroutine that respects context
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine cancelled: %v\n", ctx.Err())
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine completed normally")
		}
	}()
	
	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond) // Give goroutine time to respond

	// 3. Context with Timeout
	fmt.Println("\n3. Context with Timeout:")
	
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	// Simulate work that might exceed timeout
	select {
	case <-timeoutCtx.Done():
		fmt.Printf("Operation timed out: %v\n", timeoutCtx.Err())
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed")
	}

	// 4. Context with Deadline
	fmt.Println("\n4. Context with Deadline:")
	
	deadline := time.Now().Add(300 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// Check deadline
	if d, ok := deadlineCtx.Deadline(); ok {
		fmt.Printf("Deadline: %v\n", d)
	}
	
	// Simulate work
	select {
	case <-deadlineCtx.Done():
		fmt.Printf("Operation exceeded deadline: %v\n", deadlineCtx.Err())
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Operation completed")
	}

	// 5. Context with Values
	fmt.Println("\n5. Context with Values:")
	
	// Create context with multiple values
	ctx = context.Background()
	ctx = context.WithValue(ctx, userKey, User{ID: 1, Name: "Alice"})
	ctx = context.WithValue(ctx, requestID, "req-123")
	ctx = context.WithValue(ctx, traceID, "trace-456")
	
	// Retrieve values
	user = ctx.Value(userKey).(User)
	requestID := ctx.Value(requestID).(string)
	traceID := ctx.Value(traceID).(string)
	
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Request ID: %s\n", requestID)
	fmt.Printf("Trace ID: %s\n", traceID)

	// 6. Context Propagation
	fmt.Println("\n6. Context Propagation:")
	
	service := &Service{name: "UserService"}
	
	// Process request with context
	err := service.ProcessRequest(ctx, "user data")
	if err != nil {
		log.Printf("Error processing request: %v", err)
	}
	
	// Check if context was cancelled
	select {
	case <-ctx.Done():
		fmt.Printf("Context was cancelled: %v\n", ctx.Err())
	default:
		fmt.Println("Context is still active")
	}

	// 7. Context Chain
	fmt.Println("\n7. Context Chain:")
	
	// Create context chain
	chainCtx := context.Background()
	chainCtx = context.WithValue(chainCtx, "step", "1")
	chainCtx = context.WithTimeout(chainCtx, 1*time.Second)
	chainCtx = context.WithValue(chainCtx, "step", "2")
	chainCtx = context.WithValue(chainCtx, "step", "3")
	
	// Process chain
	processChain := func(ctx context.Context, step string) {
		select {
		case <-ctx.Done():
			fmt.Printf("Step %s cancelled: %v\n", step, ctx.Err())
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("Step %s completed\n", step)
		}
	}
	
	processChain(chainCtx, "1")
	processChain(chainCtx, "2")
	processChain(chainCtx, "3")

	// 8. Context with Database Operations
	fmt.Println("\n8. Context with Database Operations:")
	
	db := &Database{name: "PostgreSQL"}
	
	// Query with timeout
	queryCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	result, err := db.Query(queryCtx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("Database query failed: %v\n", err)
	} else {
		fmt.Printf("Database result: %s\n", result)
	}

	// 9. Context with HTTP Operations
	fmt.Println("\n9. Context with HTTP Operations:")
	
	client := &HTTPClient{timeout: 100 * time.Millisecond}
	
	// HTTP request with context
	httpCtx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	
	response, err := client.Get(httpCtx, "https://api.example.com/users")
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
	} else {
		fmt.Printf("HTTP response: %s\n", response)
	}

	// 10. Context with Goroutine Pool
	fmt.Println("\n10. Context with Goroutine Pool:")
	
	// Create worker pool with context
	workerPool := func(ctx context.Context, jobs <-chan string, results chan<- string) {
		var wg sync.WaitGroup
		
		// Start workers
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for {
					select {
					case job := <-jobs:
						select {
						case <-ctx.Done():
							fmt.Printf("Worker %d cancelled\n", workerID)
							return
						case <-time.After(50 * time.Millisecond):
							results <- fmt.Sprintf("Worker %d processed: %s", workerID, job)
						}
					case <-ctx.Done():
						fmt.Printf("Worker %d cancelled\n", workerID)
						return
					}
				}
			}(i)
		}
		
		// Wait for all workers to finish
		go func() {
			wg.Wait()
			close(results)
		}()
	}
	
	// Test worker pool
	poolCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	jobs := make(chan string, 10)
	results := make(chan string, 10)
	
	// Start worker pool
	workerPool(poolCtx, jobs, results)
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < 10; i++ {
			jobs <- fmt.Sprintf("Job %d", i)
		}
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Result: %s\n", result)
	}

	// 11. Context with Custom Values
	fmt.Println("\n11. Context with Custom Values:")
	
	// Create custom context
	customCtx := &customContext{
		Context:     context.Background(),
		customValue: "initial",
	}
	
	// Set custom value
	customCtx.SetCustomValue("updated")
	
	// Get custom value
	value := customCtx.Value("custom").(string)
	fmt.Printf("Custom value: %s\n", value)

	// 12. Context Error Handling
	fmt.Println("\n12. Context Error Handling:")
	
	// Test different context errors
	testContexts := []struct {
		name string
		ctx  context.Context
	}{
		{"Cancelled", func() context.Context {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			return ctx
		}()},
		{"Timeout", func() context.Context {
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Millisecond)
			time.Sleep(10 * time.Millisecond)
			return ctx
		}()},
		{"Deadline", func() context.Context {
			ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Millisecond))
			time.Sleep(10 * time.Millisecond)
			return ctx
		}()},
	}
	
	for _, tc := range testContexts {
		select {
		case <-tc.ctx.Done():
			fmt.Printf("%s context error: %v\n", tc.name, tc.ctx.Err())
		default:
			fmt.Printf("%s context is still active\n", tc.name)
		}
	}

	// 13. Context with Middleware
	fmt.Println("\n13. Context with Middleware:")
	
	// Middleware function
	withTimeout := func(ctx context.Context, timeout time.Duration) context.Context {
		ctx, _ = context.WithTimeout(ctx, timeout)
		return ctx
	}
	
	withValue := func(ctx context.Context, key, value interface{}) context.Context {
		return context.WithValue(ctx, key, value)
	}
	
	// Apply middleware
	middlewareCtx := context.Background()
	middlewareCtx = withTimeout(middlewareCtx, 200*time.Millisecond)
	middlewareCtx = withValue(middlewareCtx, "middleware", "applied")
	
	// Use middleware context
	select {
	case <-middlewareCtx.Done():
		fmt.Printf("Middleware context cancelled: %v\n", middlewareCtx.Err())
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("Middleware context active, value: %s\n", middlewareCtx.Value("middleware"))
	}

	// 14. Context with Request Tracing
	fmt.Println("\n14. Context with Request Tracing:")
	
	// Simulate request tracing
	traceCtx := context.Background()
	traceCtx = context.WithValue(traceCtx, "trace_id", "trace-123")
	traceCtx = context.WithValue(traceCtx, "span_id", "span-456")
	traceCtx = context.WithValue(traceCtx, "parent_id", "parent-789")
	
	// Process request with tracing
	processWithTracing := func(ctx context.Context, operation string) {
		traceID := ctx.Value("trace_id")
		spanID := ctx.Value("span_id")
		fmt.Printf("Operation %s: trace_id=%s, span_id=%s\n", operation, traceID, spanID)
	}
	
	processWithTracing(traceCtx, "database_query")
	processWithTracing(traceCtx, "http_request")
	processWithTracing(traceCtx, "cache_lookup")

	// 15. Context with Rate Limiting
	fmt.Println("\n15. Context with Rate Limiting:")
	
	// Rate limiter with context
	rateLimiter := func(ctx context.Context, rate time.Duration) <-chan struct{} {
		ticker := time.NewTicker(rate)
		ch := make(chan struct{})
		
		go func() {
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					select {
					case ch <- struct{}{}:
					case <-ctx.Done():
						close(ch)
						return
					}
				case <-ctx.Done():
					close(ch)
					return
				}
			}
		}()
		
		return ch
	}
	
	// Test rate limiter
	rateCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	rateCh := rateLimiter(rateCtx, 100*time.Millisecond)
	
	count := 0
	for range rateCh {
		count++
		fmt.Printf("Rate limited operation %d\n", count)
		if count >= 5 {
			break
		}
	}

	// 16. Context with Circuit Breaker
	fmt.Println("\n16. Context with Circuit Breaker:")
	
	// Test circuit breaker
	breaker := &CircuitBreaker{
		threshold: 3,
		timeout:   100 * time.Millisecond,
		state:     "closed",
	}
	
	// Simulate some calls
	for i := 0; i < 5; i++ {
		err := breaker.Call(ctx, func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("Call %d failed: %v (state: %s)\n", i+1, err, breaker.state)
		} else {
			fmt.Printf("Call %d succeeded (state: %s)\n", i+1, breaker.state)
		}
	}

	// 17. Context with Timeout Hierarchy
	fmt.Println("\n17. Context with Timeout Hierarchy:")
	
	// Create context hierarchy with different timeouts
	parentCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	childCtx, cancel := context.WithTimeout(parentCtx, 500*time.Millisecond)
	defer cancel()
	
	grandchildCtx, cancel := context.WithTimeout(childCtx, 200*time.Millisecond)
	defer cancel()
	
	// Test timeout hierarchy
	timeoutTest := func(ctx context.Context, name string) {
		select {
		case <-ctx.Done():
			fmt.Printf("%s context cancelled: %v\n", name, ctx.Err())
		case <-time.After(300 * time.Millisecond):
			fmt.Printf("%s context completed\n", name)
		}
	}
	
	go timeoutTest(parentCtx, "Parent")
	go timeoutTest(childCtx, "Child")
	go timeoutTest(grandchildCtx, "Grandchild")
	
	time.Sleep(600 * time.Millisecond)

	// 18. Context with Graceful Shutdown
	fmt.Println("\n18. Context with Graceful Shutdown:")
	
	// Simulate graceful shutdown
	shutdownCtx, cancel := context.WithCancel(context.Background())
	
	// Start services
	services := []string{"HTTP", "Database", "Cache", "Queue"}
	var wg sync.WaitGroup
	
	for _, service := range services {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			select {
			case <-shutdownCtx.Done():
				fmt.Printf("Service %s shutting down gracefully\n", name)
			case <-time.After(2 * time.Second):
				fmt.Printf("Service %s completed\n", name)
			}
		}(service)
	}
	
	// Simulate shutdown signal after 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown signal received")
	cancel()
	
	// Wait for all services to shutdown
	wg.Wait()
	fmt.Println("All services shutdown complete")

	fmt.Println("\n🎉 context Package Mastery Complete!")
}
