package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool demonstrates a basic worker pool pattern
func WorkerPool() {
	fmt.Println("=== Basic Worker Pool Pattern ===")
	
	// Create job queue
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				result := job * job
				fmt.Printf("Worker %d: processing job %d -> %d\n", workerID, job, result)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Worker pool results:")
	for result := range results {
		fmt.Printf("  Result: %d\n", result)
	}
}

// AdvancedWorkerPool demonstrates an advanced worker pool with job types
func AdvancedWorkerPool() {
	fmt.Println("\n=== Advanced Worker Pool with Job Types ===")
	
	// Create job queue
	jobs := make(chan Job, 100)
	results := make(chan JobResult, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				result := processJob(workerID, job)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		jobTypes := []string{"add", "multiply", "square", "divide"}
		for i := 1; i <= 12; i++ {
			jobType := jobTypes[i%len(jobTypes)]
			jobs <- Job{
				ID:   i,
				Type: jobType,
				Data: i,
			}
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Advanced worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

// Job represents a job to be processed
type Job struct {
	ID   int
	Type string
	Data int
}

// JobResult represents the result of a job
type JobResult struct {
	JobID    int
	WorkerID int
	Result   int
	Error    error
}

// String returns a string representation of JobResult
func (jr JobResult) String() string {
	if jr.Error != nil {
		return fmt.Sprintf("Job %d (Worker %d): ERROR - %v", jr.JobID, jr.WorkerID, jr.Error)
	}
	return fmt.Sprintf("Job %d (Worker %d): %d", jr.JobID, jr.WorkerID, jr.Result)
}

// processJob processes a job and returns the result
func processJob(workerID int, job Job) JobResult {
	result := JobResult{
		JobID:    job.ID,
		WorkerID: workerID,
	}
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	switch job.Type {
	case "add":
		result.Result = job.Data + 10
	case "multiply":
		result.Result = job.Data * 2
	case "square":
		result.Result = job.Data * job.Data
	case "divide":
		if job.Data == 0 {
			result.Error = fmt.Errorf("division by zero")
		} else {
			result.Result = 100 / job.Data
		}
	default:
		result.Error = fmt.Errorf("unknown job type: %s", job.Type)
	}
	
	return result
}

// WorkerPoolWithMetrics demonstrates a worker pool with metrics
func WorkerPoolWithMetrics() {
	fmt.Println("\n=== Worker Pool with Metrics ===")
	
	// Create job queue
	jobs := make(chan Job, 100)
	results := make(chan JobResult, 100)
	
	// Create metrics
	metrics := &WorkerPoolMetrics{}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				result := processJob(workerID, job)
				duration := time.Since(start)
				
				metrics.RecordJob(workerID, duration, result.Error != nil)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- Job{
				ID:   i,
				Type: "square",
				Data: i,
			}
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
	
	// Print metrics
	metrics.Print()
}

// WorkerPoolMetrics tracks worker pool metrics
type WorkerPoolMetrics struct {
	workerStats map[int]*WorkerStats
	mutex       sync.RWMutex
}

// WorkerStats represents statistics for a worker
type WorkerStats struct {
	JobsProcessed int
	TotalTime     time.Duration
	ErrorCount    int
}

// NewWorkerPoolMetrics creates a new worker pool metrics
func NewWorkerPoolMetrics() *WorkerPoolMetrics {
	return &WorkerPoolMetrics{
		workerStats: make(map[int]*WorkerStats),
	}
}

// RecordJob records a job completion
func (wpm *WorkerPoolMetrics) RecordJob(workerID int, duration time.Duration, hasError bool) {
	wpm.mutex.Lock()
	defer wpm.mutex.Unlock()
	
	if wpm.workerStats[workerID] == nil {
		wpm.workerStats[workerID] = &WorkerStats{}
	}
	
	stats := wpm.workerStats[workerID]
	stats.JobsProcessed++
	stats.TotalTime += duration
	if hasError {
		stats.ErrorCount++
	}
}

// Print prints the metrics
func (wpm *WorkerPoolMetrics) Print() {
	wpm.mutex.RLock()
	defer wpm.mutex.RUnlock()
	
	fmt.Println("Worker Pool Metrics:")
	for workerID, stats := range wpm.workerStats {
		avgTime := stats.TotalTime / time.Duration(stats.JobsProcessed)
		fmt.Printf("  Worker %d: %d jobs, avg time: %v, errors: %d\n", 
			workerID, stats.JobsProcessed, avgTime, stats.ErrorCount)
	}
}

// WorkerPoolWithPriority demonstrates a worker pool with priority jobs
func WorkerPoolWithPriority() {
	fmt.Println("\n=== Worker Pool with Priority Jobs ===")
	
	// Create priority job queue
	priorityJobs := make(chan PriorityJob, 100)
	results := make(chan JobResult, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range priorityJobs {
				result := processPriorityJob(workerID, job)
				results <- result
			}
		}(i)
	}
	
	// Send jobs with different priorities
	go func() {
		priorities := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
		for i := 1; i <= 9; i++ {
			priorityJobs <- PriorityJob{
				Job:      Job{ID: i, Type: "square", Data: i},
				Priority: priorities[i-1],
			}
		}
		close(priorityJobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Priority worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

// PriorityJob represents a job with priority
type PriorityJob struct {
	Job      Job
	Priority int
}

// processPriorityJob processes a priority job
func processPriorityJob(workerID int, job PriorityJob) JobResult {
	result := JobResult{
		JobID:    job.Job.ID,
		WorkerID: workerID,
	}
	
	// Higher priority jobs take less time
	duration := time.Duration(4-job.Priority) * 50 * time.Millisecond
	time.Sleep(duration)
	
	result.Result = job.Job.Data * job.Job.Data
	return result
}

// WorkerPoolWithBackpressure demonstrates a worker pool with backpressure
func WorkerPoolWithBackpressure() {
	fmt.Println("\n=== Worker Pool with Backpressure ===")
	
	// Create job queue with small buffer
	jobs := make(chan Job, 2)
	results := make(chan JobResult, 2)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Slow processing to cause backpressure
				time.Sleep(500 * time.Millisecond)
				result := processJob(workerID, job)
				results <- result
			}
		}(i)
	}
	
	// Send jobs (will cause backpressure)
	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Printf("Sending job %d\n", i)
			jobs <- Job{
				ID:   i,
				Type: "square",
				Data: i,
			}
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Backpressure worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

// WorkerPoolWithGracefulShutdown demonstrates a worker pool with graceful shutdown
func WorkerPoolWithGracefulShutdown() {
	fmt.Println("\n=== Worker Pool with Graceful Shutdown ===")
	
	// Create job queue
	jobs := make(chan Job, 100)
	results := make(chan JobResult, 100)
	shutdown := make(chan struct{})
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case job := <-jobs:
					result := processJob(workerID, job)
					results <- result
				case <-shutdown:
					fmt.Printf("Worker %d shutting down\n", workerID)
					return
				}
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- Job{
				ID:   i,
				Type: "square",
				Data: i,
			}
		}
	}()
	
	// Shutdown after some time
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Initiating graceful shutdown...")
		close(shutdown)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Graceful shutdown worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

// WorkerPoolWithLoadBalancing demonstrates a worker pool with load balancing
func WorkerPoolWithLoadBalancing() {
	fmt.Println("\n=== Worker Pool with Load Balancing ===")
	
	// Create job queue
	jobs := make(chan Job, 100)
	results := make(chan JobResult, 100)
	
	// Create load balancer
	balancer := NewLoadBalancer(3)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				result := processJob(workerID, job)
				duration := time.Since(start)
				
				balancer.RecordJob(workerID, duration)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 15; i++ {
			jobs <- Job{
				ID:   i,
				Type: "square",
				Data: i,
			}
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Load balancing worker pool results:")
	for result := range results {
		fmt.Printf("  %s\n", result)
	}
	
	// Print load balancing stats
	balancer.Print()
}

// LoadBalancer represents a load balancer for workers
type LoadBalancer struct {
	workerLoads map[int]time.Duration
	mutex       sync.RWMutex
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(workerCount int) *LoadBalancer {
	return &LoadBalancer{
		workerLoads: make(map[int]time.Duration),
	}
}

// RecordJob records a job completion for load balancing
func (lb *LoadBalancer) RecordJob(workerID int, duration time.Duration) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	lb.workerLoads[workerID] += duration
}

// Print prints the load balancing statistics
func (lb *LoadBalancer) Print() {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	fmt.Println("Load Balancing Statistics:")
	for workerID, load := range lb.workerLoads {
		fmt.Printf("  Worker %d: %v total load\n", workerID, load)
	}
}
