package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 2: Buffered Worker Pool
// Implement a worker pool with buffered channels for better performance.
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered Worker Pool")
	fmt.Println("===============================")
	
	const numWorkers = 3
	const numJobs = 8
	const bufferSize = 4
	
	jobs := make(chan Job, bufferSize)
	results := make(chan Result, bufferSize)
	
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
					Data:     fmt.Sprintf("Buffered: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				results <- result
				fmt.Printf("Worker %d processed buffered job %d\n", workerID, job.ID)
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Buffered Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 2 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}