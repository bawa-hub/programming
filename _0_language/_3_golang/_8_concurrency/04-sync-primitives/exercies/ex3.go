package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 3: WaitGroup
// Use WaitGroup to wait for multiple goroutines.
func Exercise3() {
	fmt.Println("\nExercise 3: WaitGroup")
	fmt.Println("=====================")
	
	var wg sync.WaitGroup
	results := make(chan int, 5)
	
	// Start workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate work
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			result := id * id
			results <- result
			fmt.Printf("Worker %d completed with result %d\n", id, result)
		}(i)
	}
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("  Result: %d\n", result)
	}
}