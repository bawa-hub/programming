package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function exits
	
	fmt.Printf("Worker %d: Starting work\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d: Finished work\n", id)
}

func main() {
	fmt.Println("=== Sync.WaitGroup Example ===")
	
	var wg sync.WaitGroup
	
	// Start 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter for each goroutine
		go worker(i, &wg)
	}
	
	fmt.Println("All workers started, waiting for completion...")
	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All workers completed!")
}
