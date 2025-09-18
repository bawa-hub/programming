package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Exercise 1: Basic Context Implementation
func Exercise1() {
	fmt.Println("\nExercise 1: Basic Context Implementation")
	fmt.Println("=======================================")
	
	// TODO: Create a context with timeout and process it
	// 1. Create a context with 3-second timeout
	// 2. Start a goroutine that does work
	// 3. Check for cancellation in the goroutine
	// 4. Handle timeout properly
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  Exercise 1: Cancelled after %d iterations: %v\n", i, ctx.Err())
				return
			default:
				fmt.Printf("  Exercise 1: Working... %d\n", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
		fmt.Println("  Exercise 1: Completed all work")
	}()
	
	time.Sleep(4 * time.Second)
	fmt.Println("Exercise 1 completed")
}

// Exercise 2: Context with Values
func Exercise2() {
	fmt.Println("\nExercise 2: Context with Values")
	fmt.Println("===============================")
	
	// TODO: Create a context with multiple values
	// 1. Create a context with user information
	// 2. Add request ID and trace ID
	// 3. Pass context through multiple functions
	// 4. Extract and use values
	
	type userKey string
	const userIDKey userKey = "userID"
	const requestIDKey userKey = "requestID"
	const traceIDKey userKey = "traceID"
	
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, "user-123")
	ctx = context.WithValue(ctx, requestIDKey, "req-456")
	ctx = context.WithValue(ctx, traceIDKey, "trace-789")
	
	processUserRequest(ctx, string(userIDKey), string(requestIDKey), string(traceIDKey))
	fmt.Println("Exercise 2 completed")
}

func processUserRequest(ctx context.Context, userIDKey, requestIDKey, traceIDKey string) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)
	traceID := ctx.Value(traceIDKey)
	
	fmt.Printf("  Exercise 2: Processing request %s for user %s (trace: %s)\n", 
		requestID, userID, traceID)
	
	// Simulate work
	time.Sleep(1 * time.Second)
	fmt.Println("  Exercise 2: Request processed successfully")
}

// Exercise 3: Context Cancellation Chain
func Exercise3() {
	fmt.Println("\nExercise 3: Context Cancellation Chain")
	fmt.Println("======================================")
	
	// TODO: Create a context hierarchy and test cancellation
	// 1. Create parent context with timeout
	// 2. Create child context with different timeout
	// 3. Create grandchild context
	// 4. Test cancellation propagation
	
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer parentCancel()
	
	childCtx, childCancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer childCancel()
	
	grandchildCtx := context.WithValue(childCtx, "level", "grandchild")
	
	// Start goroutines for each level
	go processLevel("Parent", parentCtx)
	go processLevel("Child", childCtx)
	go processLevel("Grandchild", grandchildCtx)
	
	time.Sleep(4 * time.Second)
	fmt.Println("Exercise 3 completed")
}

func processLevel(name string, ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Exercise 3: %s cancelled: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("  Exercise 3: %s working... %d\n", name, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("  Exercise 3: %s completed\n", name)
}

// Exercise 4: Context with Deadline
func Exercise4() {
	fmt.Println("\nExercise 4: Context with Deadline")
	fmt.Println("=================================")
	
	// TODO: Create a context with deadline and handle it
	// 1. Set a deadline 2 seconds from now
	// 2. Start work that might exceed deadline
	// 3. Handle deadline exceeded error
	// 4. Demonstrate proper cleanup
	
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					fmt.Printf("  Exercise 4: Deadline exceeded after %d iterations\n", i)
				} else {
					fmt.Printf("  Exercise 4: Cancelled after %d iterations: %v\n", i, ctx.Err())
				}
				return
			default:
				fmt.Printf("  Exercise 4: Working... %d\n", i)
				time.Sleep(400 * time.Millisecond)
			}
		}
		fmt.Println("  Exercise 4: Work completed")
	}()
	
	time.Sleep(3 * time.Second)
	fmt.Println("Exercise 4 completed")
}

// Exercise 5: Context in HTTP Handler
func Exercise5() {
	fmt.Println("\nExercise 5: Context in HTTP Handler")
	fmt.Println("===================================")
	
	// TODO: Simulate HTTP handler with context
	// 1. Create context from request
	// 2. Add request-scoped data
	// 3. Process request with timeout
	// 4. Handle cancellation
	
	// Simulate HTTP request context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "req-12345")
	ctx = context.WithValue(ctx, "userID", "user-67890")
	ctx = context.WithValue(ctx, "ip", "192.168.1.1")
	
	// Add timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	
	// Process request
	processHTTPRequest(ctx)
	fmt.Println("Exercise 5 completed")
}

func processHTTPRequest(ctx context.Context) {
	requestID := ctx.Value("requestID")
	userID := ctx.Value("userID")
	ip := ctx.Value("ip")
	
	fmt.Printf("  Exercise 5: Processing HTTP request %s for user %s from %s\n", 
		requestID, userID, ip)
	
	// Simulate work
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Exercise 5: Request cancelled after %d steps: %v\n", i, ctx.Err())
			return
		default:
			fmt.Printf("  Exercise 5: Processing step %d\n", i+1)
			time.Sleep(300 * time.Millisecond)
		}
	}
	fmt.Println("  Exercise 5: Request processed successfully")
}

// Exercise 6: Context with Multiple Goroutines
func Exercise6() {
	fmt.Println("\nExercise 6: Context with Multiple Goroutines")
	fmt.Println("===========================================")
	
	// TODO: Use context to coordinate multiple goroutines
	// 1. Create context with timeout
	// 2. Start multiple goroutines
	// 3. Use WaitGroup to wait for completion
	// 4. Handle cancellation in all goroutines
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	var wg sync.WaitGroup
	numGoroutines := 3
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 10; j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("  Exercise 6: Goroutine %d cancelled: %v\n", id, ctx.Err())
					return
				default:
					fmt.Printf("  Exercise 6: Goroutine %d working... %d\n", id, j)
					time.Sleep(200 * time.Millisecond)
				}
			}
			fmt.Printf("  Exercise 6: Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("Exercise 6 completed")
}

// Exercise 7: Context Value Propagation
func Exercise7() {
	fmt.Println("\nExercise 7: Context Value Propagation")
	fmt.Println("=====================================")
	
	// TODO: Create context with values and propagate through functions
	// 1. Create context with initial values
	// 2. Pass through multiple function calls
	// 3. Add values at each level
	// 4. Extract values at the end
	
	ctx := context.Background()
	ctx = addValue(ctx, "level", "root")
	ctx = addValue(ctx, "requestID", "req-001")
	
	ctx = processLevel1(ctx)
	ctx = processLevel2(ctx)
	ctx = processLevel3(ctx)
	
	// Extract all values
	fmt.Printf("  Exercise 7: Final values: %v\n", getAllValues(ctx))
	fmt.Println("Exercise 7 completed")
}

func addValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func processLevel1(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level1", "processed")
	ctx = addValue(ctx, "timestamp1", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 1 processed")
	return ctx
}

func processLevel2(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level2", "processed")
	ctx = addValue(ctx, "timestamp2", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 2 processed")
	return ctx
}

func processLevel3(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level3", "processed")
	ctx = addValue(ctx, "timestamp3", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 3 processed")
	return ctx
}

func getAllValues(ctx context.Context) map[string]interface{} {
	values := make(map[string]interface{})
	
	// Extract known values
	if level := ctx.Value("level"); level != nil {
		values["level"] = level
	}
	if requestID := ctx.Value("requestID"); requestID != nil {
		values["requestID"] = requestID
	}
	if level1 := ctx.Value("level1"); level1 != nil {
		values["level1"] = level1
	}
	if level2 := ctx.Value("level2"); level2 != nil {
		values["level2"] = level2
	}
	if level3 := ctx.Value("level3"); level3 != nil {
		values["level3"] = level3
	}
	
	return values
}

// Exercise 8: Context with Error Handling
func Exercise8() {
	fmt.Println("\nExercise 8: Context with Error Handling")
	fmt.Println("======================================")
	
	// TODO: Implement proper error handling with context
	// 1. Create context with timeout
	// 2. Simulate operations that can fail
	// 3. Handle different error types
	// 4. Implement retry logic with context
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Simulate operation with retry
	err := operationWithRetry(ctx, 3, 1*time.Second)
	if err != nil {
		fmt.Printf("  Exercise 8: Operation failed: %v\n", err)
	} else {
		fmt.Println("  Exercise 8: Operation succeeded")
	}
	
	fmt.Println("Exercise 8 completed")
}

func operationWithRetry(ctx context.Context, maxRetries int, backoff time.Duration) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("operation cancelled: %v", ctx.Err())
		default:
			fmt.Printf("  Exercise 8: Attempt %d\n", attempt)
			
			// Simulate operation that might fail
			if rand.Float32() < 0.7 { // 70% chance of failure
				fmt.Printf("  Exercise 8: Attempt %d failed\n", attempt)
				if attempt < maxRetries {
					time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
					continue
				}
				return fmt.Errorf("operation failed after %d attempts", maxRetries)
			}
			
			fmt.Printf("  Exercise 8: Attempt %d succeeded\n", attempt)
			return nil
		}
	}
	return nil
}

// Exercise 9: Context with Database Operations
func Exercise9() {
	fmt.Println("\nExercise 9: Context with Database Operations")
	fmt.Println("===========================================")
	
	// TODO: Simulate database operations with context
	// 1. Create context with timeout
	// 2. Simulate database queries
	// 3. Handle timeouts and cancellations
	// 4. Implement proper cleanup
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Simulate database operations
	user, err := getUserFromDB(ctx, "user-123")
	if err != nil {
		fmt.Printf("  Exercise 9: Error getting user: %v\n", err)
		return
	}
	
	fmt.Printf("  Exercise 9: User retrieved: %+v\n", user)
	
	// Update user
	err = updateUserInDB(ctx, user)
	if err != nil {
		fmt.Printf("  Exercise 9: Error updating user: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 9: User updated successfully")
	fmt.Println("Exercise 9 completed")
}

func getUserFromDB(ctx context.Context, userID string) (*User, error) {
	// Simulate database query
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("database query cancelled: %v", ctx.Err())
	case <-time.After(2 * time.Second):
		return &User{ID: userID, Email: "john@example.com", Role: "admin"}, nil
	}
}

func updateUserInDB(ctx context.Context, user *User) error {
	// Simulate database update
	select {
	case <-ctx.Done():
		return fmt.Errorf("database update cancelled: %v", ctx.Err())
	case <-time.After(1 * time.Second):
		user.Email = "john.updated@example.com"
		return nil
	}
}

// Exercise 10: Context Performance Testing
func Exercise10() {
	fmt.Println("\nExercise 10: Context Performance Testing")
	fmt.Println("=======================================")
	
	// TODO: Test context performance
	// 1. Measure context creation time
	// 2. Measure context value lookup time
	// 3. Measure context cancellation time
	// 4. Compare different context types
	
	// Test context creation performance
	start := time.Now()
	for i := 0; i < 100000; i++ {
		ctx := context.WithValue(context.Background(), "key", i)
		_ = ctx.Value("key")
	}
	duration := time.Since(start)
	fmt.Printf("  Exercise 10: Context creation time: %v\n", duration)
	
	// Test context cancellation performance
	ctx, cancel := context.WithCancel(context.Background())
	
	// Create many goroutines
	numGoroutines := 1000
	for i := 0; i < numGoroutines; i++ {
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
	fmt.Printf("  Exercise 10: Context cancellation time: %v\n", duration)
	
	// Wait for cancellation to propagate
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Exercise 10 completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Context Package Exercises")
	fmt.Println("============================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
