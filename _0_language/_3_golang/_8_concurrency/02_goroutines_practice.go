package main

import (
	"fmt"
	"time"
)

// Simple function that simulates work
func doWork(workerID int, taskCount int) {
	for i := 1; i <= taskCount; i++ {
		fmt.Printf("Worker %d: Starting task %d\n", workerID, i)
		time.Sleep(200 * time.Millisecond) // Simulate work
		fmt.Printf("Worker %d: Completed task %d\n", workerID, i)
	}
}

// Function that counts numbers
func countNumbers(name string, start, end int) {
	for i := start; i <= end; i++ {
		fmt.Printf("%s counting: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Goroutines Practice (No Channels) ===")
	
	// Example 1: Multiple workers doing different tasks
	fmt.Println("\n1. Multiple workers:")
	go doWork(1, 3)  // Worker 1 does 3 tasks
	go doWork(2, 2)  // Worker 2 does 2 tasks
	go doWork(3, 4)  // Worker 3 does 4 tasks
	
	// Wait for all workers to complete
	time.Sleep(2 * time.Second)
	
	// Example 2: Different functions running concurrently
	fmt.Println("\n2. Different functions concurrently:")
	go countNumbers("Counter A", 1, 5)
	go countNumbers("Counter B", 10, 15)
	
	// Wait for counters to complete
	time.Sleep(1 * time.Second)
	
	// Example 3: Anonymous functions with different delays
	fmt.Println("\n3. Anonymous functions with different speeds:")
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Fast worker: task %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Slow worker: task %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Wait for both to complete
	time.Sleep(1 * time.Second)
	
	fmt.Println("\nAll goroutines completed!")
}
