package exercises

import (
	"context"
	"fmt"
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