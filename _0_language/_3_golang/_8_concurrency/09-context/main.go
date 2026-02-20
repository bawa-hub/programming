package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// User represents a user in the system
type User struct {
	ID    string
	Email string
	Role  string
}

// Order represents an order in the system
type Order struct {
	ID     string
	UserID string
	Amount float64
	Status string
}

// contextKey is a custom type for context keys
type contextKey string

const (
	userKey    contextKey = "user"
	requestKey contextKey = "requestID"
	traceKey   contextKey = "traceID"
)

// Example 1: Basic Context Creation
func basicContextCreation() {
	fmt.Println("\n1. Basic Context Creation")
	fmt.Println("=========================")
	
	// Background context
	ctx := context.Background()
	fmt.Printf("  Background context: %v\n", ctx)
	
	// TODO context
	todoCtx := context.TODO()
	fmt.Printf("  TODO context: %v\n", todoCtx)
	
	// Context with cancellation
	cancelCtx, cancel := context.WithCancel(ctx)
	fmt.Printf("  Cancel context: %v\n", cancelCtx)
	cancel() // Cancel the context
	fmt.Printf("  After cancel: %v\n", cancelCtx.Err())
	
	// Context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	fmt.Printf("  Timeout context: %v\n", timeoutCtx)
	cancel() // Cancel immediately for demo
	
	// Context with deadline
	deadline := time.Now().Add(3 * time.Second)
	deadlineCtx, cancel := context.WithDeadline(ctx, deadline)
	fmt.Printf("  Deadline context: %v\n", deadlineCtx)
	cancel() // Cancel immediately for demo
	
	// Context with value
	valueCtx := context.WithValue(ctx, "key", "value")
	fmt.Printf("  Value context: %v\n", valueCtx)
	fmt.Printf("  Value: %v\n", valueCtx.Value("key"))
}

// Example 2: Context Cancellation
func contextCancellation() {
	fmt.Println("\n2. Context Cancellation")
	fmt.Println("======================")
	
	ctx, cancel := context.WithCancel(context.Background())
	
	// Start a goroutine that checks for cancellation
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  Goroutine cancelled: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("  Working... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
		fmt.Println("  Goroutine completed")
	}()
	
	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("  Cancelling context...")
	cancel()
	
	// Wait for cancellation to propagate
	time.Sleep(100 * time.Millisecond)
}

// Example 3: Context Timeout
func contextTimeout() {
	fmt.Println("\n3. Context Timeout")
	fmt.Println("==================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Simulate work that might take longer than timeout
	go func() {
		for i := 0; i < 20; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  Work cancelled due to timeout: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("  Working... %d\n", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
		fmt.Println("  Work completed")
	}()
	
	// Wait for timeout
	time.Sleep(3 * time.Second)
}

// Example 4: Context Deadline
func contextDeadline() {
	fmt.Println("\n4. Context Deadline")
	fmt.Println("===================")
	
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// Simulate work
	go func() {
		for i := 0; i < 20; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  Work cancelled due to deadline: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("  Working... %d\n", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
		fmt.Println("  Work completed")
	}()
	
	// Wait for deadline
	time.Sleep(3 * time.Second)
}

// Example 5: Context Values
func contextValues() {
	fmt.Println("\n5. Context Values")
	fmt.Println("=================")
	
	// Create context with values
	ctx := context.Background()
	ctx = context.WithValue(ctx, userKey, "john@example.com")
	ctx = context.WithValue(ctx, requestKey, "req-12345")
	ctx = context.WithValue(ctx, traceKey, "trace-67890")
	
	// Extract values
	user := ctx.Value(userKey)
	requestID := ctx.Value(requestKey)
	traceID := ctx.Value(traceKey)
	
	fmt.Printf("  User: %v\n", user)
	fmt.Printf("  Request ID: %v\n", requestID)
	fmt.Printf("  Trace ID: %v\n", traceID)
	
	// Non-existent key
	nonExistent := ctx.Value("non-existent")
	fmt.Printf("  Non-existent key: %v\n", nonExistent)
}

// Example 6: Context Propagation
func contextPropagation() {
	fmt.Println("\n6. Context Propagation")
	fmt.Println("======================")
	
	// Create root context
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, "level", "root")
	
	// Create child context
	childCtx, cancel := context.WithTimeout(rootCtx, 3*time.Second)
	defer cancel()
	childCtx = context.WithValue(childCtx, "level", "child")
	
	// Create grandchild context
	grandchildCtx := context.WithValue(childCtx, "level", "grandchild")
	
	// Process each context
	processContext("Root", rootCtx)
	processContext("Child", childCtx)
	processContext("Grandchild", grandchildCtx)
	
	// Cancel child context
	fmt.Println("  Cancelling child context...")
	cancel()
	
	// Check grandchild context
	processContext("Grandchild after cancel", grandchildCtx)
}

func processContext(name string, ctx context.Context) {
	fmt.Printf("  %s context: level=%v, err=%v\n", 
		name, ctx.Value("level"), ctx.Err())
}

// Example 7: HTTP Request with Context
func httpRequestWithContext() {
	fmt.Println("\n7. HTTP Request with Context")
	fmt.Println("============================")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		log.Printf("  Error creating request: %v", err)
		return
	}
	
	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("  Request timed out")
		} else {
			fmt.Printf("  Request failed: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()
	
	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("  Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("  Response status: %s\n", resp.Status)
	fmt.Printf("  Response body length: %d bytes\n", len(body))
}

// Example 8: Database Operations with Context
func databaseOperationsWithContext() {
	fmt.Println("\n8. Database Operations with Context")
	fmt.Println("===================================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Simulate database operations
	user, err := getUserByID(ctx, "123")
	if err != nil {
		fmt.Printf("  Error getting user: %v\n", err)
		return
	}
	
	fmt.Printf("  User: %+v\n", user)
	
	// Update user
	err = updateUser(ctx, user)
	if err != nil {
		fmt.Printf("  Error updating user: %v\n", err)
		return
	}
	
	fmt.Println("  User updated successfully")
}

func getUserByID(ctx context.Context, id string) (*User, error) {
	// Simulate database query
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(2 * time.Second):
		return &User{ID: id, Email: "john@example.com", Role: "admin"}, nil
	}
}

func updateUser(ctx context.Context, user *User) error {
	// Simulate database update
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(1 * time.Second):
		user.Email = "john.updated@example.com"
		return nil
	}
}

// Example 9: Multiple Goroutines with Context
func multipleGoroutinesWithContext() {
	fmt.Println("\n9. Multiple Goroutines with Context")
	fmt.Println("===================================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Create multiple goroutines
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 10; j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("  Goroutine %d cancelled: %v\n", id, ctx.Err())
					return
				default:
					fmt.Printf("  Goroutine %d working... %d\n", id, j)
					time.Sleep(200 * time.Millisecond)
				}
			}
			fmt.Printf("  Goroutine %d completed\n", id)
		}(i)
	}
	
	// Wait for all goroutines
	wg.Wait()
}

// Example 10: Context Chain
func contextChain() {
	fmt.Println("\n10. Context Chain")
	fmt.Println("=================")
	
	// Create chain of contexts
	ctx := context.Background()
	
	// Chain operations
	ctx = withRequestID(ctx)
	ctx = withUser(ctx, "john@example.com")
	ctx = withTimeout(ctx, 2*time.Second)
	ctx = withTrace(ctx, "trace-12345")
	
	// Process chain
	processChain(ctx)
}

func withRequestID(ctx context.Context) context.Context {
	requestID := fmt.Sprintf("req-%d", time.Now().UnixNano())
	return context.WithValue(ctx, requestKey, requestID)
}

func withUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func withTimeout(ctx context.Context, timeout time.Duration) context.Context {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	_ = cancel // We're not using cancel in this demo
	return newCtx
}

func withTrace(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceKey, traceID)
}

func processChain(ctx context.Context) {
	// Extract all values
	requestID := ctx.Value(requestKey)
	user := ctx.Value(userKey)
	traceID := ctx.Value(traceKey)
	
	fmt.Printf("  Processing chain: %s, %s, %s\n", requestID, user, traceID)
	
	// Simulate work
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Chain cancelled: %v\n", ctx.Err())
			return
		default:
			fmt.Printf("  Chain working... %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}
	fmt.Println("  Chain completed")
}

// Example 11: Context Middleware
func contextMiddleware() {
	fmt.Println("\n11. Context Middleware")
	fmt.Println("======================")
	
	// Create middleware chain
	ctx := context.Background()
	ctx = addRequestID(ctx)
	ctx = addUser(ctx, "john@example.com")
	ctx = addTimeout(ctx, 2*time.Second)
	
	// Process request
	processRequest(ctx)
}

func addRequestID(ctx context.Context) context.Context {
	requestID := generateRequestID()
	return context.WithValue(ctx, requestKey, requestID)
}

func addUser(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, userKey, email)
}

func addTimeout(ctx context.Context, timeout time.Duration) context.Context {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	_ = cancel // We're not using cancel in this demo
	return newCtx
}

func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}

func processRequest(ctx context.Context) {
	// Extract values
	requestID := ctx.Value(requestKey)
	user := ctx.Value(userKey)
	
	fmt.Printf("  Processing request: %s for user %s\n", requestID, user)
	
	// Simulate work
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Request cancelled: %v\n", ctx.Err())
			return
		default:
			fmt.Printf("  Processing... %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}
	fmt.Println("  Request completed")
}

// Example 12: Context Performance
func contextPerformance() {
	fmt.Println("\n12. Context Performance")
	fmt.Println("=======================")
	
	// Test context creation performance
	start := time.Now()
	for i := 0; i < 100000; i++ {
		ctx := context.WithValue(context.Background(), "key", i)
		_ = ctx.Value("key")
	}
	duration := time.Since(start)
	fmt.Printf("  Context creation time: %v\n", duration)
	
	// Test context cancellation performance
	ctx, cancel := context.WithCancel(context.Background())
	
	// Create many goroutines
	for i := 0; i < 1000; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// Do work
					time.Sleep(1 * time.Millisecond)
				}
			}
		}(i)
	}
	
	// Measure cancellation time
	start = time.Now()
	cancel()
	duration = time.Since(start)
	fmt.Printf("  Context cancellation time: %v\n", duration)
	
	// Wait for cancellation to propagate
	time.Sleep(100 * time.Millisecond)
}

// Example 13: Context Error Handling
func contextErrorHandling() {
	fmt.Println("\n13. Context Error Handling")
	fmt.Println("==========================")
	
	// Test different error scenarios
	testContextError("Timeout", func() context.Context {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		_ = cancel // We're not using cancel in this demo
		return ctx
	})
	
	testContextError("Deadline", func() context.Context {
		deadline := time.Now().Add(1 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		_ = cancel // We're not using cancel in this demo
		return ctx
	})
	
	testContextError("Cancellation", func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately
		return ctx
	})
}

func testContextError(name string, createCtx func() context.Context) {
	ctx := createCtx()
	
	// Simulate work
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  %s error: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("  %s working... %d\n", name, i)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Printf("  %s completed\n", name)
}

// Example 14: Context Values with Types
func contextValuesWithTypes() {
	fmt.Println("\n14. Context Values with Types")
	fmt.Println("=============================")
	
	// Create context with typed values
	ctx := context.Background()
	
	user := User{ID: "123", Email: "john@example.com", Role: "admin"}
	order := Order{ID: "456", UserID: "123", Amount: 99.99, Status: "pending"}
	
	ctx = context.WithValue(ctx, userKey, user)
	ctx = context.WithValue(ctx, "order", order)
	
	// Extract typed values
	if user, ok := ctx.Value(userKey).(User); ok {
		fmt.Printf("  User: %+v\n", user)
	}
	
	if order, ok := ctx.Value("order").(Order); ok {
		fmt.Printf("  Order: %+v\n", order)
	}
	
	// Type assertion failure
	if user, ok := ctx.Value("order").(User); !ok {
		fmt.Printf("  Type assertion failed: %v\n", user)
	}
}

// Example 15: Context Best Practices
func contextBestPractices() {
	fmt.Println("\n15. Context Best Practices")
	fmt.Println("==========================")
	
	fmt.Println("  ✅ Good practices:")
	fmt.Println("    - Context as first parameter")
	fmt.Println("    - Always check ctx.Done() in long operations")
	fmt.Println("    - Use custom key types for context values")
	fmt.Println("    - Always call cancel() to prevent leaks")
	fmt.Println("    - Use appropriate context types")
	
	fmt.Println("\n  ❌ Bad practices:")
	fmt.Println("    - Storing context in structs")
	fmt.Println("    - Not checking context in loops")
	fmt.Println("    - Using string keys for context values")
	fmt.Println("    - Forgetting to call cancel()")
	fmt.Println("    - Using TODO context in production")
	
	// Demonstrate good practice
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel
	
	// Check context in loop
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Good practice: cancelled after %d iterations\n", i)
			return
		default:
			// Do work
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Println("  Good practice: completed all iterations")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🚀 Context Package Examples")
	fmt.Println("============================")
	
	basicContextCreation()
	contextCancellation()
	contextTimeout()
	contextDeadline()
	contextValues()
	contextPropagation()
	httpRequestWithContext()
	databaseOperationsWithContext()
	multipleGoroutinesWithContext()
	contextChain()
	contextMiddleware()
	contextPerformance()
	contextErrorHandling()
	contextValuesWithTypes()
	contextBestPractices()
}


func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "basic":
		runBasicExamples()
	case "exercises":
		runExercises()
	case "advanced":
		runAdvancedPatterns()
	case "all":
		runBasicExamples()
		runExercises()
		runAdvancedPatterns()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("🚀 Context Package Commands")
	fmt.Println("===========================")
	fmt.Println("")
	fmt.Println("Usage: go run . <command>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic context examples")
	fmt.Println("  exercises - Run hands-on exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  all       - Run all examples and exercises")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println("  go run . all")
}