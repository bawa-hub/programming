package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Concurrency Learning - Project 4 ===")
	fmt.Println("Advanced Concurrency")
	fmt.Println()

	// Exercise 1: Advanced Worker Pool
	fmt.Println("Exercise 1: Advanced Worker Pool")
	demonstrateAdvancedWorkerPool()
	fmt.Println()

	// Exercise 2: Context-Based Cancellation
	fmt.Println("Exercise 2: Context-Based Cancellation")
	demonstrateContextSystem()
	fmt.Println()

	// Exercise 3: Pipeline with Backpressure
	fmt.Println("Exercise 3: Pipeline with Backpressure")
	demonstratePipelineBackpressure()
	fmt.Println()

	// Exercise 4: Concurrent Data Structures
	fmt.Println("Exercise 4: Concurrent Data Structures")
	demonstrateConcurrentDataStructures()
	fmt.Println()

	// Exercise 5: Memory Pool System
	fmt.Println("Exercise 5: Memory Pool System")
	demonstrateMemoryPool()
	fmt.Println()

	fmt.Println("=== All advanced concurrency exercises completed! ===")
	fmt.Println()
	fmt.Println("Run specific components:")
	fmt.Println("  go run main.go advanced_worker_pool.go")
	fmt.Println("  go run main.go context_system.go")
	fmt.Println("  go run main.go pipeline_backpressure.go")
	fmt.Println("  go run main.go concurrent_data_structures.go")
	fmt.Println("  go run main.go memory_pool.go")
}

func demonstrateAdvancedWorkerPool() {
	// Create advanced worker pool
	pool := NewAdvancedWorkerPool(3, 10)
	pool.Start()
	
	// Submit jobs
	for i := 1; i <= 10; i++ {
		job := &AdvancedJob{
			ID:   i,
			Type: "process",
			Data: i,
		}
		pool.SubmitJob(job)
	}
	
	// Let it run
	time.Sleep(2 * time.Second)
	
	// Shutdown
	pool.Shutdown()
}

func demonstrateContextSystem() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Create context-based system
	system := NewContextBasedSystem(ctx)
	system.Start()
	
	// Let it run
	time.Sleep(1 * time.Second)
	
	// Shutdown will happen automatically due to context timeout
	<-system.Done()
}

func demonstratePipelineBackpressure() {
	// Create pipeline with backpressure
	pipeline := NewBackpressurePipeline(3, 5)
	pipeline.Start()
	
	// Send data
	for i := 1; i <= 20; i++ {
		pipeline.Send(i)
	}
	
	// Let it process
	time.Sleep(3 * time.Second)
	
	// Shutdown
	pipeline.Shutdown()
}

func demonstrateConcurrentDataStructures() {
	// Test concurrent data structures
	testConcurrentMap()
	testConcurrentQueue()
	testConcurrentStack()
}

func demonstrateMemoryPool() {
	// Create memory pool
	pool := NewMemoryPool(10)
	
	// Use objects from pool
	objects := make([]*PooledObject, 0)
	for i := 0; i < 5; i++ {
		obj := pool.Get()
		obj.Data = i
		objects = append(objects, obj)
	}
	
	// Return objects to pool
	for _, obj := range objects {
		pool.Put(obj)
	}
	
	// Shutdown pool
	pool.Close()
}

// AdvancedJob represents an advanced job
type AdvancedJob struct {
	ID       int
	Type     string
	Data     interface{}
	Priority int
	Created  time.Time
}

// AdvancedWorkerPool represents an advanced worker pool
type AdvancedWorkerPool struct {
	workers     int
	maxWorkers  int
	jobs        chan *AdvancedJob
	results     chan *JobResult
	shutdown    chan struct{}
	done        chan struct{}
	wg          sync.WaitGroup
	metrics     *WorkerPoolMetrics
	healthCheck *HealthChecker
}

// NewAdvancedWorkerPool creates a new advanced worker pool
func NewAdvancedWorkerPool(initialWorkers, maxWorkers int) *AdvancedWorkerPool {
	return &AdvancedWorkerPool{
		workers:     initialWorkers,
		maxWorkers:  maxWorkers,
		jobs:        make(chan *AdvancedJob, 100),
		results:     make(chan *JobResult, 100),
		shutdown:    make(chan struct{}),
		done:        make(chan struct{}),
		metrics:     NewWorkerPoolMetrics(),
		healthCheck: NewHealthChecker(),
	}
}

// Start starts the worker pool
func (awp *AdvancedWorkerPool) Start() {
	// Start workers
	for i := 0; i < awp.workers; i++ {
		awp.wg.Add(1)
		go awp.worker(i)
	}
	
	// Start result collector
	awp.wg.Add(1)
	go awp.resultCollector()
	
	// Start health checker
	awp.wg.Add(1)
	go awp.healthChecker()
}

// worker processes jobs
func (awp *AdvancedWorkerPool) worker(id int) {
	defer awp.wg.Done()
	
	for {
		select {
		case job := <-awp.jobs:
			start := time.Now()
			result := awp.processJob(id, job)
			duration := time.Since(start)
			
			awp.metrics.RecordJob(id, duration, result.Error != nil)
			awp.results <- result
		case <-awp.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// processJob processes a job
func (awp *AdvancedWorkerPool) processJob(workerID int, job *AdvancedJob) *JobResult {
	result := &JobResult{
		JobID:    job.ID,
		WorkerID: workerID,
	}
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	// Process based on job type
	switch job.Type {
	case "process":
		if data, ok := job.Data.(int); ok {
			result.Result = data * data
		}
	default:
		result.Error = fmt.Errorf("unknown job type: %s", job.Type)
	}
	
	return result
}

// resultCollector collects results
func (awp *AdvancedWorkerPool) resultCollector() {
	defer awp.wg.Done()
	
	for {
		select {
		case result := <-awp.results:
			fmt.Printf("Result: Job %d -> %v\n", result.JobID, result.Result)
		case <-awp.shutdown:
			fmt.Println("Result collector shutting down")
			return
		}
	}
}

// healthChecker monitors worker health
func (awp *AdvancedWorkerPool) healthChecker() {
	defer awp.wg.Done()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			awp.healthCheck.Check()
		case <-awp.shutdown:
			fmt.Println("Health checker shutting down")
			return
		}
	}
}

// SubmitJob submits a job to the pool
func (awp *AdvancedWorkerPool) SubmitJob(job *AdvancedJob) {
	job.Created = time.Now()
	awp.jobs <- job
}

// Shutdown shuts down the worker pool
func (awp *AdvancedWorkerPool) Shutdown() {
	close(awp.shutdown)
	awp.wg.Wait()
	close(awp.done)
}

// Done returns a channel that signals when shutdown is complete
func (awp *AdvancedWorkerPool) Done() <-chan struct{} {
	return awp.done
}

// JobResult represents the result of a job
type JobResult struct {
	JobID    int
	WorkerID int
	Result   interface{}
	Error    error
}

// WorkerPoolMetrics tracks worker pool metrics
type WorkerPoolMetrics struct {
	workerStats map[int]*WorkerStats
	mutex       sync.RWMutex
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

// WorkerStats represents statistics for a worker
type WorkerStats struct {
	JobsProcessed int
	TotalTime     time.Duration
	ErrorCount    int
}

// HealthChecker monitors worker health
type HealthChecker struct {
	lastCheck time.Time
	mutex     sync.RWMutex
}

// NewHealthChecker creates a new health checker
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		lastCheck: time.Now(),
	}
}

// Check performs a health check
func (hc *HealthChecker) Check() {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()
	
	hc.lastCheck = time.Now()
	fmt.Printf("Health check at %v\n", hc.lastCheck)
}

// ContextBasedSystem represents a context-based system
type ContextBasedSystem struct {
	ctx     context.Context
	workers int
	jobs    chan int
	results chan int
	done    chan struct{}
	wg      sync.WaitGroup
}

// NewContextBasedSystem creates a new context-based system
func NewContextBasedSystem(ctx context.Context) *ContextBasedSystem {
	return &ContextBasedSystem{
		ctx:     ctx,
		workers: 3,
		jobs:    make(chan int, 100),
		results: make(chan int, 100),
		done:    make(chan struct{}),
	}
}

// Start starts the system
func (cbs *ContextBasedSystem) Start() {
	// Start workers
	for i := 0; i < cbs.workers; i++ {
		cbs.wg.Add(1)
		go cbs.worker(i)
	}
	
	// Start job producer
	cbs.wg.Add(1)
	go cbs.producer()
	
	// Start result collector
	cbs.wg.Add(1)
	go cbs.collector()
	
	// Start shutdown handler
	go cbs.shutdownHandler()
}

// worker processes jobs
func (cbs *ContextBasedSystem) worker(id int) {
	defer cbs.wg.Done()
	
	for {
		select {
		case job := <-cbs.jobs:
			// Process job
			time.Sleep(200 * time.Millisecond)
			result := job * job
			cbs.results <- result
		case <-cbs.ctx.Done():
			fmt.Printf("Worker %d shutting down due to context\n", id)
			return
		}
	}
}

// producer generates jobs
func (cbs *ContextBasedSystem) producer() {
	defer cbs.wg.Done()
	
	jobCount := 0
	for {
		select {
		case cbs.jobs <- jobCount:
			jobCount++
			time.Sleep(100 * time.Millisecond)
		case <-cbs.ctx.Done():
			fmt.Println("Producer shutting down due to context")
			return
		}
	}
}

// collector collects results
func (cbs *ContextBasedSystem) collector() {
	defer cbs.wg.Done()
	
	for {
		select {
		case result := <-cbs.results:
			fmt.Printf("Result: %d\n", result)
		case <-cbs.ctx.Done():
			fmt.Println("Collector shutting down due to context")
			return
		}
	}
}

// shutdownHandler handles shutdown
func (cbs *ContextBasedSystem) shutdownHandler() {
	<-cbs.ctx.Done()
	fmt.Println("Context cancelled, initiating shutdown...")
	cbs.wg.Wait()
	close(cbs.done)
}

// Done returns a channel that signals when shutdown is complete
func (cbs *ContextBasedSystem) Done() <-chan struct{} {
	return cbs.done
}

// BackpressurePipeline represents a pipeline with backpressure
type BackpressurePipeline struct {
	stages     int
	bufferSize int
	channels   []chan int
	shutdown   chan struct{}
	done       chan struct{}
	wg         sync.WaitGroup
}

// NewBackpressurePipeline creates a new backpressure pipeline
func NewBackpressurePipeline(stages, bufferSize int) *BackpressurePipeline {
	return &BackpressurePipeline{
		stages:     stages,
		bufferSize: bufferSize,
		channels:   make([]chan int, stages),
		shutdown:   make(chan struct{}),
		done:       make(chan struct{}),
	}
}

// Start starts the pipeline
func (bp *BackpressurePipeline) Start() {
	// Create channels
	for i := 0; i < bp.stages; i++ {
		bp.channels[i] = make(chan int, bp.bufferSize)
	}
	
	// Start stages
	for i := 0; i < bp.stages-1; i++ {
		bp.wg.Add(1)
		go bp.stage(i)
	}
	
	// Start final stage
	bp.wg.Add(1)
	go bp.finalStage()
	
	// Start shutdown handler
	go func() {
		<-bp.shutdown
		bp.wg.Wait()
		close(bp.done)
	}()
}

// stage processes data in a pipeline stage
func (bp *BackpressurePipeline) stage(stageID int) {
	defer bp.wg.Done()
	
	input := bp.channels[stageID]
	output := bp.channels[stageID+1]
	
	for {
		select {
		case data := <-input:
			// Process data
			time.Sleep(100 * time.Millisecond)
			result := data * 2
			
			// Send to next stage (with backpressure)
			select {
			case output <- result:
				fmt.Printf("Stage %d: %d -> %d\n", stageID, data, result)
			case <-bp.shutdown:
				return
			}
		case <-bp.shutdown:
			return
		}
	}
}

// finalStage processes data in the final stage
func (bp *BackpressurePipeline) finalStage() {
	defer bp.wg.Done()
	
	input := bp.channels[bp.stages-1]
	
	for {
		select {
		case data := <-input:
			// Process final data
			time.Sleep(100 * time.Millisecond)
			result := data + 10
			fmt.Printf("Final stage: %d -> %d\n", data, result)
		case <-bp.shutdown:
			return
		}
	}
}

// Send sends data to the pipeline
func (bp *BackpressurePipeline) Send(data int) {
	select {
	case bp.channels[0] <- data:
		// Data sent successfully
	case <-bp.shutdown:
		// Pipeline is shutting down
	default:
		// Backpressure - channel is full
		fmt.Printf("Backpressure: dropping data %d\n", data)
	}
}

// Shutdown shuts down the pipeline
func (bp *BackpressurePipeline) Shutdown() {
	select {
	case <-bp.shutdown:
		// Already closed
	default:
		close(bp.shutdown)
	}
	bp.wg.Wait()
	close(bp.done)
}

// Done returns a channel that signals when shutdown is complete
func (bp *BackpressurePipeline) Done() <-chan struct{} {
	return bp.done
}

// Concurrent data structures
func testConcurrentMap() {
	fmt.Println("Testing Concurrent Map...")
	cm := NewConcurrentMap()
	
	// Test concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			cm.Set(key, id)
			value, exists := cm.Get(key)
			if !exists || value != id {
				fmt.Printf("Error: expected %d, got %v\n", id, value)
			}
		}(i)
	}
	wg.Wait()
}

func testConcurrentQueue() {
	fmt.Println("Testing Concurrent Queue...")
	cq := NewConcurrentQueue()
	
	// Test concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cq.Enqueue(id)
			_, ok := cq.Dequeue()
			if !ok {
				fmt.Printf("Error: failed to dequeue %d\n", id)
			}
		}(i)
	}
	wg.Wait()
}

func testConcurrentStack() {
	fmt.Println("Testing Concurrent Stack...")
	cs := NewConcurrentStack()
	
	// Test concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cs.Push(id)
			_, ok := cs.Pop()
			if !ok {
				fmt.Printf("Error: failed to pop %d\n", id)
			}
		}(i)
	}
	wg.Wait()
}
