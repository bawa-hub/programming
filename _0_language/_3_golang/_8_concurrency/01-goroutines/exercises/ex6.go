package exercises

import (
	"fmt"
	"time"
)

// Exercise 6: Advanced - Goroutine Pool with Rate Limiting
func Exercise6() {
	fmt.Println("\nExercise 6: Advanced - Rate Limited Pool")
	fmt.Println("=======================================")
	
	const numWorkers = 2
	const numJobs = 8
	const rateLimit = 2 // jobs per second
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	rateLimiter := time.Tick(time.Second / rateLimit)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go rateLimitedWorker(i, jobs, results, rateLimiter)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job %d completed with result: %d\n", r, result)
	}
}

func rateLimitedWorker(id int, jobs <-chan int, results chan<- int, rateLimiter <-chan time.Time) {
	for job := range jobs {
		<-rateLimiter // Wait for rate limit
		fmt.Printf("Worker %d: Processing job %d (rate limited)\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 3
	}
}