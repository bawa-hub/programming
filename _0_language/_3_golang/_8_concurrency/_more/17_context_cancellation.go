package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Cancelled! Reason: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("=== Context Cancellation Example ===")
	
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}
	
	// Let workers run for 2 seconds
	time.Sleep(2 * time.Second)
	
	// Cancel all workers
	fmt.Println("Cancelling all workers...")
	cancel()
	
	// Wait a bit to see the cancellation
	time.Sleep(1 * time.Second)
	
	fmt.Println("All workers stopped!")
}
