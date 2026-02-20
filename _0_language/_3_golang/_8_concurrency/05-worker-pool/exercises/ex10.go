package exercises

import (
	"fmt"
	"time"
)

// Exercise 10: Pipeline Worker Pool
// Implement a worker pool that processes tasks through multiple stages.
func Exercise10() {
	fmt.Println("\nExercise 10: Pipeline Worker Pool")
	fmt.Println("=================================")
	
	const numWorkers = 3
	const numJobs = 8
	
	input := make(chan Job, numJobs)
	stage1 := make(chan Job, numJobs)
	stage2 := make(chan Job, numJobs)
	output := make(chan Result, numJobs)
	
	// Stage 1: Process input
	go func() {
		defer close(stage1)
		for job := range input {
			// Simulate stage 1 processing
			time.Sleep(30 * time.Millisecond)
			job.Data = fmt.Sprintf("Stage1: %s", job.Data)
			stage1 <- job
		}
	}()
	
	// Stage 2: Process stage1 output
	go func() {
		defer close(stage2)
		for job := range stage1 {
			// Simulate stage 2 processing
			time.Sleep(30 * time.Millisecond)
			job.Data = fmt.Sprintf("Stage2: %s", job.Data)
			stage2 <- job
		}
	}()
	
	// Stage 3: Final processing with workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range stage2 {
				start := time.Now()
				
				// Simulate final processing
				time.Sleep(30 * time.Millisecond)
				
				result := Result{
					JobID:    job.ID,
					Data:     fmt.Sprintf("Pipeline: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				output <- result
				fmt.Printf("Worker %d processed pipeline job %d\n", workerID, job.ID)
			}
		}(i)
	}
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Submit jobs
	go func() {
		defer close(input)
		for i := 0; i < numJobs; i++ {
			input <- Job{
				ID:   i,
				Data: fmt.Sprintf("Pipeline Job %d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 10 Results:")
	for result := range output {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}