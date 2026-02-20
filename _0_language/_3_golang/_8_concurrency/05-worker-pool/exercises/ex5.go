package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 5: Worker Pool with Results
// Create a worker pool that collects and processes results from workers.
func Exercise5() {
	fmt.Println("\nExercise 5: Worker Pool with Results")
	fmt.Println("====================================")
	
	const numWorkers = 3
	const numJobs = 8
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				
				// Simulate work
				time.Sleep(time.Duration(job.ID*50) * time.Millisecond)
				
				result := Result{
					JobID:    job.ID,
					Data:     fmt.Sprintf("Result: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				results <- result
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Result Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Process results as they arrive
	fmt.Println("Exercise 5 Results:")
	for result := range results {
		fmt.Printf("  Processing result for job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
		
		// Simulate result processing
		time.Sleep(20 * time.Millisecond)
	}
}