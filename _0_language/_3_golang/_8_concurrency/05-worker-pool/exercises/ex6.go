package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 6: Worker Pool with Error Handling
// Implement a worker pool that properly handles errors from workers.
func Exercise6() {
	fmt.Println("\nExercise 6: Worker Pool with Error Handling")
	fmt.Println("===========================================")
	
	const numWorkers = 3
	const numJobs = 8
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	errors := make(chan error, numJobs)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				
				// Simulate work with occasional errors
				time.Sleep(time.Duration(job.ID*50) * time.Millisecond)
				
				if job.ID%3 == 0 {
					// Simulate error
					errors <- fmt.Errorf("worker %d failed to process job %d", workerID, job.ID)
				} else {
					result := Result{
						JobID:    job.ID,
						Data:     fmt.Sprintf("Success: %s", job.Data),
						Duration: time.Since(start),
						WorkerID: workerID,
					}
					results <- result
				}
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Error Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()
	
	// Process results and errors
	fmt.Println("Exercise 6 Results:")
	for {
		select {
		case result, ok := <-results:
			if !ok {
				results = nil
			} else {
				fmt.Printf("  SUCCESS: Job %d: %s (took %v, worker %d)\n", 
					result.JobID, result.Data, result.Duration, result.WorkerID)
			}
		case err, ok := <-errors:
			if !ok {
				errors = nil
			} else {
				fmt.Printf("  ERROR: %v\n", err)
			}
		}
		
		if results == nil && errors == nil {
			break
		}
	}
}