package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 9: Worker Pool with Metrics
// Create a worker pool that collects and reports performance metrics.
func Exercise9() {
	fmt.Println("\nExercise 9: Worker Pool with Metrics")
	fmt.Println("====================================")
	
	const numWorkers = 3
	const numJobs = 8
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	metrics := &ExerciseWorkerPoolMetrics{}
	
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
					Data:     fmt.Sprintf("Metrics: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				
				results <- result
				metrics.recordTask(start, nil)
				fmt.Printf("Worker %d processed metrics job %d\n", workerID, job.ID)
			}
		}(i)
	}
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Metrics Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 9 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
	
	// Print metrics
	fmt.Printf("\nExercise 9 Metrics:\n")
	fmt.Printf("  Processed Tasks: %d\n", metrics.getProcessedTasks())
	fmt.Printf("  Average Processing Time: %v\n", metrics.getAverageProcessingTime())
	fmt.Printf("  Error Count: %d\n", metrics.getErrorCount())
}

type ExerciseWorkerPoolMetrics struct {
	processedTasks int64
	processingTime time.Duration
	errorCount     int64
	mu             sync.RWMutex
}

func (wpm *ExerciseWorkerPoolMetrics) recordTask(start time.Time, err error) {
	wpm.mu.Lock()
	defer wpm.mu.Unlock()
	
	wpm.processedTasks++
	wpm.processingTime += time.Since(start)
	if err != nil {
		wpm.errorCount++
	}
}

func (wpm *ExerciseWorkerPoolMetrics) getProcessedTasks() int64 {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	return wpm.processedTasks
}

func (wpm *ExerciseWorkerPoolMetrics) getAverageProcessingTime() time.Duration {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	if wpm.processedTasks == 0 {
		return 0
	}
	return wpm.processingTime / time.Duration(wpm.processedTasks)
}

func (wpm *ExerciseWorkerPoolMetrics) getErrorCount() int64 {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	return wpm.errorCount
}