package exercises

import (
	"context"
	"fmt"
	"time"
)

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