package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID       int
	Data     string
	Priority int
	Timeout  time.Duration
}

// Result represents the result of processing a job
type Result struct {
	JobID    int
	Data     string
	Error    error
	Duration time.Duration
	WorkerID int
}

// Exercise 1: Basic Worker Pool
// Create a basic worker pool that processes tasks from a channel.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Worker Pool")
	fmt.Println("=============================")
	
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
					Data:     fmt.Sprintf("Exercise1: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				results <- result
				fmt.Printf("Worker %d processed job %d\n", workerID, job.ID)
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Exercise Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 1 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}