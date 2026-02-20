package exercises

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

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