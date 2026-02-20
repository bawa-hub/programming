package exercises

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Exercise 7: Worker Pool with Timeout
// Create a worker pool that handles timeouts for individual tasks.
func Exercise7() {
	fmt.Println("\nExercise 7: Worker Pool with Timeout")
	fmt.Println("====================================")
	
	const numWorkers = 3
	const numJobs = 8
	const timeout = 1 * time.Second
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return
					}
					
					start := time.Now()
					
					// Simulate work with timeout
					select {
					case <-time.After(time.Duration(job.ID*100) * time.Millisecond):
						result := Result{
							JobID:    job.ID,
							Data:     fmt.Sprintf("Timeout: %s", job.Data),
							Duration: time.Since(start),
							WorkerID: workerID,
						}
						results <- result
						fmt.Printf("Worker %d processed timeout job %d\n", workerID, job.ID)
					case <-ctx.Done():
						fmt.Printf("Worker %d timed out on job %d\n", workerID, job.ID)
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			select {
			case jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Timeout Job %d", i),
			}:
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 7 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}