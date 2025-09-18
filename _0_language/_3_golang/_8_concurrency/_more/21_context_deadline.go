package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context, taskID int) {
	fmt.Printf("Task %d: Starting long-running task\n", taskID)
	
	// Check if we have a deadline
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("Task %d: Must complete by %v\n", taskID, deadline)
	}
	
	// Simulate work with periodic checks
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d: Cancelled at step %d, reason: %v\n", taskID, i, ctx.Err())
			return
		default:
			fmt.Printf("Task %d: Working... step %d\n", taskID, i+1)
			time.Sleep(300 * time.Millisecond)
		}
	}
	
	fmt.Printf("Task %d: Completed successfully!\n", taskID)
}

func main() {
	fmt.Println("=== Context Deadline Example ===")
	
	// Create context with deadline (2 seconds from now)
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// Start multiple tasks
	for i := 1; i <= 3; i++ {
		go longRunningTask(ctx, i)
	}
	
	// Wait for all tasks to complete or timeout
	time.Sleep(3 * time.Second)
	
	fmt.Println("All tasks processed!")
}
