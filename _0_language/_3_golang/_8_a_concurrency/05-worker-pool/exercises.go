package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

// Exercise 3: Dynamic Worker Pool
// Create a worker pool that can adjust the number of workers based on workload.
func Exercise3() {
	fmt.Println("\nExercise 3: Dynamic Worker Pool")
	fmt.Println("==============================")
	
	const minWorkers = 2
	const maxWorkers = 4
	const numJobs = 12
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Dynamic worker pool
	pool := &ExerciseDynamicWorkerPool{
		workers:    minWorkers,
		minWorkers: minWorkers,
		maxWorkers: maxWorkers,
		jobs:       jobs,
		results:    results,
	}
	
	// Start initial workers
	pool.startWorkers()
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Dynamic Job %d", i),
			}
			
			// Adjust workers based on queue size
			pool.adjustWorkers()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		pool.wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 3 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

type ExerciseDynamicWorkerPool struct {
	workers    int
	minWorkers int
	maxWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.RWMutex
}

func (dwp *ExerciseDynamicWorkerPool) startWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	for i := 0; i < dwp.workers; i++ {
		dwp.wg.Add(1)
		go dwp.worker(i)
	}
}

func (dwp *ExerciseDynamicWorkerPool) worker(id int) {
	defer dwp.wg.Done()
	
	for job := range dwp.jobs {
		start := time.Now()
		
		// Simulate work
		time.Sleep(time.Duration(job.ID*50) * time.Millisecond)
		
		result := Result{
			JobID:    job.ID,
			Data:     fmt.Sprintf("Dynamic: %s", job.Data),
			Duration: time.Since(start),
			WorkerID: id,
		}
		
		dwp.results <- result
		fmt.Printf("Worker %d processed dynamic job %d\n", id, job.ID)
	}
}

func (dwp *ExerciseDynamicWorkerPool) adjustWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	queueSize := len(dwp.jobs)
	
	if queueSize > 2 && dwp.workers < dwp.maxWorkers {
		// Add worker
		dwp.workers++
		dwp.wg.Add(1)
		go dwp.worker(dwp.workers - 1)
		fmt.Printf("Added worker, total: %d\n", dwp.workers)
	}
}

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

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Worker Pool Exercises")
	fmt.Println("====================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}
