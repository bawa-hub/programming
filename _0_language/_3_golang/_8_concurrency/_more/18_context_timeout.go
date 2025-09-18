package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context, id int) error {
	fmt.Printf("Operation %d: Starting slow operation...\n", id)
	
	// Simulate slow work
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Operation %d: Cancelled after %d iterations\n", id, i)
			return ctx.Err()
		default:
			fmt.Printf("Operation %d: Working... step %d\n", id, i+1)
			time.Sleep(1 * time.Second)
		}
	}
	
	fmt.Printf("Operation %d: Completed successfully!\n", id)
	return nil
}

func main() {
	fmt.Println("=== Context Timeout Example ===")
	
	// Example 1: Operation completes before timeout
	fmt.Println("\n1. Operation with 3-second timeout (should complete):")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel1()
	
	go func() {
		err := slowOperation(ctx1, 1)
		if err != nil {
			fmt.Printf("Operation 1 failed: %v\n", err)
		}
	}()
	
	time.Sleep(4 * time.Second)
	
	// Example 2: Operation times out
	fmt.Println("\n2. Operation with 2-second timeout (should timeout):")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	
	go func() {
		err := slowOperation(ctx2, 2)
		if err != nil {
			fmt.Printf("Operation 2 failed: %v\n", err)
		}
	}()
	
	time.Sleep(3 * time.Second)
	
	fmt.Println("\nAll operations completed!")
}
