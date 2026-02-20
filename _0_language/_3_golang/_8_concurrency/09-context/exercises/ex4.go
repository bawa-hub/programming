package exercises

import (
	"context"
	"fmt"
	"time"
)

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