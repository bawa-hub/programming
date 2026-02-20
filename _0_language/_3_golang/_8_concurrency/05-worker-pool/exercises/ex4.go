package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 4: Priority Worker Pool
// Implement a worker pool that processes tasks based on priority.
func Exercise4() {
	fmt.Println("\nExercise 4: Priority Worker Pool")
	fmt.Println("===============================")
	
	const numWorkers = 3
	const numJobs = 8
	
	highPriority := make(chan Job, 4)
	lowPriority := make(chan Job, 4)
	results := make(chan Result, 8)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			for {
				select {
				case job, ok := <-highPriority:
					if !ok {
						highPriority = nil
					} else {
						start := time.Now()
						time.Sleep(time.Duration(job.ID*30) * time.Millisecond)
						
						result := Result{
							JobID:    job.ID,
							Data:     fmt.Sprintf("High Priority: %s", job.Data),
							Duration: time.Since(start),
							WorkerID: workerID,
						}
						
						results <- result
						fmt.Printf("Worker %d processed HIGH priority job %d\n", workerID, job.ID)
					}
				case job, ok := <-lowPriority:
					if !ok {
						lowPriority = nil
					} else {
						start := time.Now()
						time.Sleep(time.Duration(job.ID*50) * time.Millisecond)
						
						result := Result{
							JobID:    job.ID,
							Data:     fmt.Sprintf("Low Priority: %s", job.Data),
							Duration: time.Since(start),
							WorkerID: workerID,
						}
						
						results <- result
						fmt.Printf("Worker %d processed LOW priority job %d\n", workerID, job.ID)
					}
				}
				
				if highPriority == nil && lowPriority == nil {
					break
				}
			}
		}(i)
	}
	
	// Submit jobs with different priorities
	go func() {
		defer close(highPriority)
		defer close(lowPriority)
		
		for i := 0; i < numJobs; i++ {
			if i%3 == 0 {
				highPriority <- Job{
					ID:       i,
					Data:     fmt.Sprintf("High Priority Job %d", i),
					Priority: 1,
				}
			} else {
				lowPriority <- Job{
					ID:       i,
					Data:     fmt.Sprintf("Low Priority Job %d", i),
					Priority: 0,
				}
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 4 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}