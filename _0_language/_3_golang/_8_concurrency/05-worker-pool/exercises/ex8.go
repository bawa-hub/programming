package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 8: Worker Pool with Rate Limiting
// Implement a worker pool that limits the rate of task processing.
func Exercise8() {
	fmt.Println("\nExercise 8: Worker Pool with Rate Limiting")
	fmt.Println("=========================================")
	
	const numWorkers = 3
	const numJobs = 8
	const rateLimit = 3 // jobs per second
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Rate limiter
	rateLimiter := time.NewTicker(time.Second / rateLimit)
	defer rateLimiter.Stop()
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Wait for rate limiter
				<-rateLimiter.C
				
				start := time.Now()
				
				// Simulate work
				time.Sleep(100 * time.Millisecond)
				
				result := Result{
					JobID:    job.ID,
					Data:     fmt.Sprintf("Rate Limited: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				results <- result
				fmt.Printf("Worker %d processed rate-limited job %d\n", workerID, job.ID)
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Rate Limited Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 8 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}