package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
	pl "pool/exercises"
	pt "pool/patterns"
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

// Example 1: Basic Worker Pool
func basicWorkerPool() {
	fmt.Println("1. Basic Worker Pool")
	fmt.Println("===================")
	
	const numWorkers = 3
	const numJobs = 10
	
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
				time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
				
				result := Result{
					JobID:    job.ID,
					Data:     fmt.Sprintf("Processed: %s", job.Data),
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
				Data: fmt.Sprintf("Job %d", i),
			}
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 2: Buffered Worker Pool
func bufferedWorkerPool() {
	fmt.Println("\n2. Buffered Worker Pool")
	fmt.Println("======================")
	
	const numWorkers = 3
	const numJobs = 10
	const bufferSize = 5
	
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
				time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
				
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
	fmt.Println("Buffered Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 3: Dynamic Worker Pool
func dynamicWorkerPool() {
	fmt.Println("\n3. Dynamic Worker Pool")
	fmt.Println("=====================")
	
	const minWorkers = 2
	const maxWorkers = 5
	const numJobs = 15
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Dynamic worker pool
	pool := &DynamicWorkerPool{
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
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		pool.wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Dynamic Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

type DynamicWorkerPool struct {
	workers    int
	minWorkers int
	maxWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.RWMutex
}

func (dwp *DynamicWorkerPool) startWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	for i := 0; i < dwp.workers; i++ {
		dwp.wg.Add(1)
		go dwp.worker(i)
	}
}

func (dwp *DynamicWorkerPool) worker(id int) {
	defer dwp.wg.Done()
	
	for job := range dwp.jobs {
		start := time.Now()
		
		// Simulate work
		time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
		
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

func (dwp *DynamicWorkerPool) adjustWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	queueSize := len(dwp.jobs)
	
	if queueSize > 3 && dwp.workers < dwp.maxWorkers {
		// Add worker
		dwp.workers++
		dwp.wg.Add(1)
		go dwp.worker(dwp.workers - 1)
		fmt.Printf("Added worker, total: %d\n", dwp.workers)
	} else if queueSize == 0 && dwp.workers > dwp.minWorkers {
		// Remove worker (simplified - in real implementation, you'd need to signal workers to stop)
		fmt.Printf("Would remove worker, total: %d\n", dwp.workers)
	}
}

// Example 4: Priority Worker Pool
func priorityWorkerPool() {
	fmt.Println("\n4. Priority Worker Pool")
	fmt.Println("======================")
	
	const numWorkers = 3
	const numJobs = 10
	
	highPriority := make(chan Job, 5)
	lowPriority := make(chan Job, 5)
	results := make(chan Result, 10)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			highPriorityOpen := true
			lowPriorityOpen := true
			
			for highPriorityOpen || lowPriorityOpen {
				select {
				case job, ok := <-highPriority:
					if !ok {
						highPriorityOpen = false
					} else {
						start := time.Now()
						time.Sleep(time.Duration(job.ID*5) * time.Millisecond)
						
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
						lowPriorityOpen = false
					} else {
						start := time.Now()
						time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
						
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
	fmt.Println("Priority Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 5: Worker Pool with Results
func workerPoolWithResults() {
	fmt.Println("\n5. Worker Pool with Results")
	fmt.Println("===========================")
	
	const numWorkers = 3
	const numJobs = 10
	
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
				time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
				
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
	fmt.Println("Processing results:")
	for result := range results {
		fmt.Printf("  Processing result for job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
		
		// Simulate result processing
		time.Sleep(10 * time.Millisecond)
	}
}

// Example 6: Worker Pool with Error Handling
func workerPoolWithErrorHandling() {
	fmt.Println("\n6. Worker Pool with Error Handling")
	fmt.Println("==================================")
	
	const numWorkers = 3
	const numJobs = 10
	
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
				time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
				
				if job.ID%4 == 0 {
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
	fmt.Println("Processing results and errors:")
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

// Example 7: Worker Pool with Timeout
func workerPoolWithTimeout() {
	fmt.Println("\n7. Worker Pool with Timeout")
	fmt.Println("===========================")
	
	const numWorkers = 3
	const numJobs = 10
	const timeout = 2 * time.Second
	
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
	fmt.Println("Timeout Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 8: Worker Pool with Rate Limiting
func workerPoolWithRateLimiting() {
	fmt.Println("\n8. Worker Pool with Rate Limiting")
	fmt.Println("=================================")
	
	const numWorkers = 3
	const numJobs = 10
	const rateLimit = 2 // jobs per second
	
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
	fmt.Println("Rate Limited Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 9: Worker Pool with Metrics
func workerPoolWithMetrics() {
	fmt.Println("\n9. Worker Pool with Metrics")
	fmt.Println("===========================")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	metrics := &WorkerPoolMetrics{}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				
				// Simulate work
				time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
				
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
	fmt.Println("Metrics Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
	
	// Print metrics
	fmt.Printf("\nMetrics:\n")
	fmt.Printf("  Processed Tasks: %d\n", metrics.getProcessedTasks())
	fmt.Printf("  Average Processing Time: %v\n", metrics.getAverageProcessingTime())
	fmt.Printf("  Error Count: %d\n", metrics.getErrorCount())
}

type WorkerPoolMetrics struct {
	processedTasks int64
	processingTime time.Duration
	errorCount     int64
	mu             sync.RWMutex
}

func (wpm *WorkerPoolMetrics) recordTask(start time.Time, err error) {
	wpm.mu.Lock()
	defer wpm.mu.Unlock()
	
	wpm.processedTasks++
	wpm.processingTime += time.Since(start)
	if err != nil {
		wpm.errorCount++
	}
}

func (wpm *WorkerPoolMetrics) getProcessedTasks() int64 {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	return wpm.processedTasks
}

func (wpm *WorkerPoolMetrics) getAverageProcessingTime() time.Duration {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	if wpm.processedTasks == 0 {
		return 0
	}
	return wpm.processingTime / time.Duration(wpm.processedTasks)
}

func (wpm *WorkerPoolMetrics) getErrorCount() int64 {
	wpm.mu.RLock()
	defer wpm.mu.RUnlock()
	return wpm.errorCount
}

// Example 10: Pipeline Worker Pool
func pipelineWorkerPool() {
	fmt.Println("\n10. Pipeline Worker Pool")
	fmt.Println("========================")
	
	const numWorkers = 3
	const numJobs = 10
	
	input := make(chan Job, numJobs)
	stage1 := make(chan Job, numJobs)
	stage2 := make(chan Job, numJobs)
	output := make(chan Result, numJobs)
	
	// Stage 1: Process input
	go func() {
		defer close(stage1)
		for job := range input {
			// Simulate stage 1 processing
			time.Sleep(50 * time.Millisecond)
			job.Data = fmt.Sprintf("Stage1: %s", job.Data)
			stage1 <- job
		}
	}()
	
	// Stage 2: Process stage1 output
	go func() {
		defer close(stage2)
		for job := range stage1 {
			// Simulate stage 2 processing
			time.Sleep(50 * time.Millisecond)
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
				time.Sleep(50 * time.Millisecond)
				
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
	fmt.Println("Pipeline Results:")
	for result := range output {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

// Example 11: Performance Comparison
func performanceComparison() {
	fmt.Println("\n11. Performance Comparison")
	fmt.Println("==========================")
	
	const numJobs = 1000
	
	// Sequential processing
	start := time.Now()
	for i := 0; i < numJobs; i++ {
		// Simulate work
		time.Sleep(1 * time.Millisecond)
	}
	sequentialTime := time.Since(start)
	
	// Worker pool processing
	start = time.Now()
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Simulate work
				time.Sleep(1 * time.Millisecond)
				results <- Result{JobID: job.ID, WorkerID: workerID}
			}
		}(i)
	}
	
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{ID: i}
		}
	}()
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for range results {
		// Just consume results
	}
	
	workerPoolTime := time.Since(start)
	
	fmt.Printf("Sequential processing: %v\n", sequentialTime)
	fmt.Printf("Worker pool processing: %v\n", workerPoolTime)
	fmt.Printf("Speedup: %.2fx\n", float64(sequentialTime)/float64(workerPoolTime))
}

// Example 12: Common Pitfalls
func commonPitfalls() {
	fmt.Println("\n12. Common Pitfalls")
	fmt.Println("==================")
	
	fmt.Println("Pitfall 1: Goroutine Leaks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// for i := 0; i < numWorkers; i++ {")
	fmt.Println("//     go func() { /* work */ }()")
	fmt.Println("// }")
	fmt.Println("// // No WaitGroup, no cleanup")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// var wg sync.WaitGroup")
	fmt.Println("// for i := 0; i < numWorkers; i++ {")
	fmt.Println("//     wg.Add(1)")
	fmt.Println("//     go func() { defer wg.Done(); /* work */ }()")
	fmt.Println("// }")
	fmt.Println("// wg.Wait()")
	
	fmt.Println("\nPitfall 2: Channel Deadlocks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// results <- result // Can block if channel is full")
	fmt.Println("// wg.Wait() // This will block forever")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// select {")
	fmt.Println("// case results <- result:")
	fmt.Println("// case <-time.After(timeout):")
	fmt.Println("//     // Handle timeout")
	fmt.Println("// }")
	
	fmt.Println("\nPitfall 3: Incorrect Worker Count")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// for i := 0; i < 1000; i++ { // Too many workers")
	fmt.Println("//     go func() { /* work */ }()")
	fmt.Println("// }")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// numWorkers := runtime.NumCPU() // Optimal for CPU-bound tasks")
	fmt.Println("// for i := 0; i < numWorkers; i++ {")
	fmt.Println("//     go func() { /* work */ }()")
	fmt.Println("// }")
	
	fmt.Println("\nPitfall 4: Missing Error Handling")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// result := process(job) // Can panic")
	fmt.Println("// results <- result")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// defer func() {")
	fmt.Println("//     if r := recover(); r != nil {")
	fmt.Printf("//         errors <- fmt.Errorf(\"worker panicked: %%v\", r)\n")
	fmt.Println("//     }")
	fmt.Println("// }()")
	fmt.Println("// result, err := process(job)")
	fmt.Println("// if err != nil { errors <- err } else { results <- result }")
}

// Utility function to show worker pool info
func showWorkerPoolInfo() {
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "basic":
		runBasicExamples()
	case "exercises":
		RunAllExercises()
	case "advanced":
		pt.RunAdvancedPatterns()
	case "all":
		runBasicExamples()
		fmt.Println("\n" + "==================================================")
		RunAllExercises()
		fmt.Println("\n" + "==================================================")
		pt.RunAdvancedPatterns()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func showUsage() {
	fmt.Println("🚀 Worker Pool Pattern - Usage")
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic worker pool examples")
	fmt.Println("  exercises - Run all exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  all       - Run everything")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println("  go run . all")
}

func runBasicExamples() {
	fmt.Println("🚀 Worker Pool Pattern Examples")
	fmt.Println("===============================")

	// Example 1: Basic Worker Pool
	basicWorkerPool()

	// Example 2: Buffered Worker Pool
	bufferedWorkerPool()

	// Example 3: Dynamic Worker Pool
	dynamicWorkerPool()

	// Example 4: Priority Worker Pool
	priorityWorkerPool()

	// Example 5: Worker Pool with Results
	workerPoolWithResults()

	// Example 6: Worker Pool with Error Handling
	workerPoolWithErrorHandling()

	// Example 7: Worker Pool with Timeout
	workerPoolWithTimeout()

	// Example 8: Worker Pool with Rate Limiting
	workerPoolWithRateLimiting()

	// Example 9: Worker Pool with Metrics
	workerPoolWithMetrics()

	// Example 10: Pipeline Worker Pool
	pipelineWorkerPool()

	// Example 11: Performance Comparison
	performanceComparison()

	// Example 12: Common Pitfalls
	commonPitfalls()
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Worker Pool Exercises")
	fmt.Println("====================================")
	
	pl.Exercise1()
	pl.Exercise2()
	pl.Exercise3()
	pl.Exercise4()
	pl.Exercise5()
	pl.Exercise6()
	pl.Exercise7()
	pl.Exercise8()
	pl.Exercise9()
	pl.Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}

