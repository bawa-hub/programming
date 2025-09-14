package main

import (
	"fmt"
	"time"
)

// Worker function that processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Starting job %d\n", id, job)
		time.Sleep(time.Duration(job) * 200 * time.Millisecond) // Simulate work
		result := job * 2
		fmt.Printf("Worker %d: Completed job %d, result: %d\n", id, job, result)
		results <- result
	}
}

func main() {
	fmt.Println("=== Worker Pattern with Channels ===")
	
	// Create channels
	jobs := make(chan int, 5)    // Buffered channel for jobs
	results := make(chan int, 5) // Buffered channel for results
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	
	// Send jobs
	fmt.Println("Sending jobs...")
	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Printf("Sent job %d\n", j)
	}
	close(jobs) // Close jobs channel when done sending
	
	// Collect results
	fmt.Println("\nCollecting results...")
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Final result: %d\n", result)
	}
	
	fmt.Println("\nAll jobs completed!")
}
