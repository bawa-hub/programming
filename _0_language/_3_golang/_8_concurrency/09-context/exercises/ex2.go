package exercises

import (
	"context"
	"fmt"
	"time"
)

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