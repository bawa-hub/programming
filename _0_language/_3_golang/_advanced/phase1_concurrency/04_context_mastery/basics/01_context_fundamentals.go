package main

import (
	"context"
	"fmt"
	"time"
)

// 🎯 CONTEXT FUNDAMENTALS MASTERY
// Understanding the core concepts of context in Go

func main() {
	fmt.Println("🎯 CONTEXT FUNDAMENTALS MASTERY")
	fmt.Println("===============================")

	// 1. Basic Context
	fmt.Println("\n1. Basic Context:")
	basicContext()

	// 2. Context with Cancellation
	fmt.Println("\n2. Context with Cancellation:")
	contextWithCancellation()

	// 3. Context with Timeout
	fmt.Println("\n3. Context with Timeout:")
	contextWithTimeout()

	// 4. Context with Deadline
	fmt.Println("\n4. Context with Deadline:")
	contextWithDeadline()

	// 5. Context with Values
	fmt.Println("\n5. Context with Values:")
	contextWithValues()

	// 6. Context Inheritance
	fmt.Println("\n6. Context Inheritance:")
	contextInheritance()

	// 7. Context Best Practices
	fmt.Println("\n7. Context Best Practices:")
	contextBestPractices()
}

// BASIC CONTEXT: Understanding context basics
func basicContext() {
	fmt.Println("Understanding basic context...")
	
	// Create a background context
	ctx := context.Background()
	fmt.Printf("  📊 Background context: %v\n", ctx)
	
	// Create a TODO context
	todoCtx := context.TODO()
	fmt.Printf("  📊 TODO context: %v\n", todoCtx)
	
	// Check if context is done
	select {
	case <-ctx.Done():
		fmt.Println("  ❌ Context is done")
	default:
		fmt.Println("  ✅ Context is active")
	}
	
	// Check context error
	if err := ctx.Err(); err != nil {
		fmt.Printf("  ❌ Context error: %v\n", err)
	} else {
		fmt.Println("  ✅ Context has no error")
	}
}

// CONTEXT WITH CANCELLATION: Manual cancellation
func contextWithCancellation() {
	fmt.Println("Understanding context with cancellation...")
	
	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Start a goroutine that does work
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  🧵 Goroutine: Context cancelled, stopping at iteration %d\n", i)
				return
			default:
				fmt.Printf("  🧵 Goroutine: Working... iteration %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
		fmt.Println("  🧵 Goroutine: Work completed")
	}()
	
	// Let it run for a bit
	time.Sleep(1 * time.Second)
	
	// Cancel the context
	fmt.Println("  🛑 Main: Cancelling context...")
	cancel()
	
	// Wait a bit to see the cancellation effect
	time.Sleep(500 * time.Millisecond)
}

// CONTEXT WITH TIMEOUT: Automatic timeout
func contextWithTimeout() {
	fmt.Println("Understanding context with timeout...")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// Start a goroutine that might take longer than timeout
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  🧵 Goroutine: Context done, stopping at iteration %d\n", i)
				return
			default:
				fmt.Printf("  🧵 Goroutine: Working... iteration %d\n", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
		fmt.Println("  🧵 Goroutine: Work completed")
	}()
	
	// Wait for context to timeout
	<-ctx.Done()
	
	// Check why context was done
	if err := ctx.Err(); err != nil {
		fmt.Printf("  ⏰ Context timed out: %v\n", err)
	}
}

// CONTEXT WITH DEADLINE: Specific deadline
func contextWithDeadline() {
	fmt.Println("Understanding context with deadline...")
	
	// Create context with deadline
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// Start a goroutine that does work
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  🧵 Goroutine: Context done, stopping at iteration %d\n", i)
				return
			default:
				fmt.Printf("  🧵 Goroutine: Working... iteration %d\n", i)
				time.Sleep(400 * time.Millisecond)
			}
		}
		fmt.Println("  🧵 Goroutine: Work completed")
	}()
	
	// Wait for context to be done
	<-ctx.Done()
	
	// Check why context was done
	if err := ctx.Err(); err != nil {
		fmt.Printf("  ⏰ Context deadline exceeded: %v\n", err)
	}
	
	// Check if deadline was exceeded
	if time.Now().After(deadline) {
		fmt.Println("  ⏰ Deadline was exceeded")
	}
}

// CONTEXT WITH VALUES: Request-scoped values
func contextWithValues() {
	fmt.Println("Understanding context with values...")
	
	// Create context with values
	ctx := context.WithValue(context.Background(), "userID", "12345")
	ctx = context.WithValue(ctx, "requestID", "req-67890")
	ctx = context.WithValue(ctx, "traceID", "trace-abcde")
	
	// Pass context to a function
	processRequest(ctx)
}

// processRequest demonstrates context value usage
func processRequest(ctx context.Context) {
	// Extract values from context
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	traceID := ctx.Value("traceID")
	
	fmt.Printf("  📊 User ID: %v\n", userID)
	fmt.Printf("  📊 Request ID: %v\n", requestID)
	fmt.Printf("  📊 Trace ID: %v\n", traceID)
	
	// Pass context to another function
	processData(ctx)
}

// processData demonstrates context propagation
func processData(ctx context.Context) {
	// Extract values from context
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	
	fmt.Printf("  🧵 Processing data for user %v, request %v\n", userID, requestID)
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("  🧵 Data processing completed")
}

// CONTEXT INHERITANCE: Context trees
func contextInheritance() {
	fmt.Println("Understanding context inheritance...")
	
	// Create parent context
	parentCtx := context.WithValue(context.Background(), "parent", "value")
	
	// Create child context with timeout
	childCtx, cancel := context.WithTimeout(parentCtx, 1*time.Second)
	defer cancel()
	
	// Create grandchild context with cancellation
	grandchildCtx, cancel2 := context.WithCancel(childCtx)
	defer cancel2()
	
	// Check inheritance
	fmt.Printf("  📊 Parent context value: %v\n", parentCtx.Value("parent"))
	fmt.Printf("  📊 Child context value: %v\n", childCtx.Value("parent"))
	fmt.Printf("  📊 Grandchild context value: %v\n", grandchildCtx.Value("parent"))
	
	// Check context hierarchy
	fmt.Printf("  📊 Parent context: %v\n", parentCtx)
	fmt.Printf("  📊 Child context: %v\n", childCtx)
	fmt.Printf("  📊 Grandchild context: %v\n", grandchildCtx)
	
	// Cancel grandchild context
	fmt.Println("  🛑 Cancelling grandchild context...")
	cancel2()
	
	// Check if parent contexts are affected
	select {
	case <-parentCtx.Done():
		fmt.Println("  ❌ Parent context is done")
	default:
		fmt.Println("  ✅ Parent context is still active")
	}
	
	select {
	case <-childCtx.Done():
		fmt.Println("  ❌ Child context is done")
	default:
		fmt.Println("  ✅ Child context is still active")
	}
	
	select {
	case <-grandchildCtx.Done():
		fmt.Println("  ❌ Grandchild context is done")
	default:
		fmt.Println("  ✅ Grandchild context is still active")
	}
}

// CONTEXT BEST PRACTICES: Following Go conventions
func contextBestPractices() {
	fmt.Println("Understanding context best practices...")
	
	// 1. Always pass context as first parameter
	fmt.Println("  📝 Best Practice 1: Context as first parameter")
	goodFunction(context.Background(), "data")
	
	// 2. Don't store context in structs
	fmt.Println("  📝 Best Practice 2: Don't store context in structs")
	
	// 3. Use context.TODO() when unsure
	fmt.Println("  📝 Best Practice 3: Use context.TODO() when unsure")
	unsureFunction(context.TODO())
	
	// 4. Always call cancel function
	fmt.Println("  📝 Best Practice 4: Always call cancel function")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Always call cancel
	
	// 5. Check context in long-running operations
	fmt.Println("  📝 Best Practice 5: Check context in long-running operations")
	longRunningOperation(ctx)
}

// goodFunction demonstrates proper context usage
func goodFunction(ctx context.Context, data string) {
	// Check if context is done
	select {
	case <-ctx.Done():
		fmt.Printf("  🧵 Function: Context done, not processing %s\n", data)
		return
	default:
		fmt.Printf("  🧵 Function: Processing %s\n", data)
	}
}

// unsureFunction demonstrates context.TODO() usage
func unsureFunction(ctx context.Context) {
	fmt.Println("  🧵 Function: Using context.TODO() for future context support")
}

// longRunningOperation demonstrates context checking in loops
func longRunningOperation(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  🧵 Operation: Context done, stopping at iteration %d\n", i)
			return
		default:
			fmt.Printf("  🧵 Operation: Working... iteration %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Println("  🧵 Operation: Completed successfully")
}
