package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Work Stealing Worker Pool
type WorkStealingPool struct {
	workers    int
	jobs       chan Job
	results    chan Result
	workQueues []chan Job
	stealIndex int64
}

func NewWorkStealingPool(workers int) *WorkStealingPool {
	pool := &WorkStealingPool{
		workers:    workers,
		jobs:       make(chan Job, workers*10),
		results:    make(chan Result, workers*10),
		workQueues: make([]chan Job, workers),
	}
	
	for i := 0; i < workers; i++ {
		pool.workQueues[i] = make(chan Job, 10)
	}
	
	return pool
}

func (wsp *WorkStealingPool) Start() {
	var wg sync.WaitGroup
	
	for i := 0; i < wsp.workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			wsp.worker(workerID)
		}(i)
	}
	
	go func() {
		wg.Wait()
		close(wsp.results)
	}()
}

func (wsp *WorkStealingPool) worker(workerID int) {
	for {
		// Try to get work from own queue first
		select {
		case job, ok := <-wsp.workQueues[workerID]:
			if !ok {
				return
			}
			wsp.processJob(job, workerID)
		default:
			// Try to steal work from other queues
			if wsp.stealWork(workerID) {
				continue
			}
			
			// No work available, wait a bit
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func (wsp *WorkStealingPool) stealWork(workerID int) bool {
	for i := 0; i < wsp.workers; i++ {
		target := (workerID + i + 1) % wsp.workers
		select {
		case job, ok := <-wsp.workQueues[target]:
			if ok {
				wsp.processJob(job, workerID)
				return true
			}
		default:
			continue
		}
	}
	return false
}

func (wsp *WorkStealingPool) processJob(job Job, workerID int) {
	start := time.Now()
	
	// Simulate work
	time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
	
	result := Result{
		JobID:    job.ID,
		Data:     fmt.Sprintf("Work Stealing: %s", job.Data),
		Duration: time.Since(start),
		WorkerID: workerID,
	}
	
	wsp.results <- result
}

func (wsp *WorkStealingPool) Submit(job Job) {
	// Distribute work round-robin
	index := atomic.AddInt64(&wsp.stealIndex, 1) % int64(wsp.workers)
	select {
	case wsp.workQueues[index] <- job:
	default:
		// Queue is full, try to submit to jobs channel
		wsp.jobs <- job
	}
}

func (wsp *WorkStealingPool) Close() {
	for _, queue := range wsp.workQueues {
		close(queue)
	}
	close(wsp.jobs)
}

// Advanced Pattern 2: Adaptive Worker Pool
type AdaptiveWorkerPool struct {
	minWorkers    int
	maxWorkers    int
	currentWorkers int
	jobs          chan Job
	results       chan Result
	metrics       *AdaptiveMetrics
	mu            sync.RWMutex
	wg            sync.WaitGroup
	ctx           context.Context
	cancel        context.CancelFunc
}

type AdaptiveMetrics struct {
	processedTasks    int64
	averageLatency    time.Duration
	queueSize         int64
	workerUtilization float64
	mu                sync.RWMutex
}

func NewAdaptiveWorkerPool(minWorkers, maxWorkers int) *AdaptiveWorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	
	pool := &AdaptiveWorkerPool{
		minWorkers:    minWorkers,
		maxWorkers:    maxWorkers,
		currentWorkers: minWorkers,
		jobs:          make(chan Job, maxWorkers*10),
		results:       make(chan Result, maxWorkers*10),
		metrics:       &AdaptiveMetrics{},
		ctx:           ctx,
		cancel:        cancel,
	}
	
	pool.startWorkers()
	go pool.adaptiveController()
	
	return pool
}

func (awp *AdaptiveWorkerPool) startWorkers() {
	awp.mu.Lock()
	defer awp.mu.Unlock()
	
	for i := 0; i < awp.currentWorkers; i++ {
		awp.wg.Add(1)
		go awp.worker(i)
	}
}

func (awp *AdaptiveWorkerPool) worker(workerID int) {
	defer awp.wg.Done()
	
	for {
		select {
		case job, ok := <-awp.jobs:
			if !ok {
				return
			}
			
			start := time.Now()
			
			// Simulate work
			time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
			
			result := Result{
				JobID:    job.ID,
				Data:     fmt.Sprintf("Adaptive: %s", job.Data),
				Duration: time.Since(start),
				WorkerID: workerID,
			}
			
			awp.results <- result
			awp.metrics.recordTask(start, nil)
			
		case <-awp.ctx.Done():
			return
		}
	}
}

func (awp *AdaptiveWorkerPool) adaptiveController() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			awp.adjustWorkers()
		case <-awp.ctx.Done():
			return
		}
	}
}

func (awp *AdaptiveWorkerPool) adjustWorkers() {
	awp.mu.Lock()
	defer awp.mu.Unlock()
	
	queueSize := len(awp.jobs)
	utilization := awp.metrics.getWorkerUtilization()
	
	if queueSize > 5 && awp.currentWorkers < awp.maxWorkers && utilization > 0.8 {
		// Add worker
		awp.currentWorkers++
		awp.wg.Add(1)
		go awp.worker(awp.currentWorkers - 1)
		fmt.Printf("Added worker, total: %d\n", awp.currentWorkers)
	} else if queueSize == 0 && awp.currentWorkers > awp.minWorkers && utilization < 0.3 {
		// Remove worker (simplified - in real implementation, you'd need to signal workers to stop)
		fmt.Printf("Would remove worker, total: %d\n", awp.currentWorkers)
	}
}

func (awp *AdaptiveWorkerPool) Submit(job Job) {
	select {
	case awp.jobs <- job:
	case <-awp.ctx.Done():
		return
	}
}

func (awp *AdaptiveWorkerPool) Close() {
	awp.cancel()
	close(awp.jobs)
	awp.wg.Wait()
	close(awp.results)
}

func (am *AdaptiveMetrics) recordTask(start time.Time, err error) {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	am.processedTasks++
	latency := time.Since(start)
	am.averageLatency = (am.averageLatency + latency) / 2
}

func (am *AdaptiveMetrics) getWorkerUtilization() float64 {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.workerUtilization
}

// Advanced Pattern 3: Circuit Breaker Worker Pool
type CircuitBreakerWorkerPool struct {
	workers      int
	jobs         chan Job
	results      chan Result
	errors       chan error
	breaker      *CircuitBreaker
	mu           sync.RWMutex
	wg           sync.WaitGroup
}

type CircuitBreaker struct {
	failures    int64
	successes   int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
}

func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
	}
}

func (cb *CircuitBreaker) Allow() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	
	if cb.state == 0 { // closed
		return true
	} else if cb.state == 1 { // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = 2 // half-open
			return true
		}
		return false
	} else { // half-open
		return true
	}
}

func (cb *CircuitBreaker) RecordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	cb.successes++
	if cb.state == 2 { // half-open
		cb.state = 0 // closed
		cb.failures = 0
	}
}

func (cb *CircuitBreaker) RecordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	cb.failures++
	cb.lastFailure = time.Now()
	
	if cb.failures >= cb.threshold {
		cb.state = 1 // open
	}
}

func NewCircuitBreakerWorkerPool(workers int, threshold int64, timeout time.Duration) *CircuitBreakerWorkerPool {
	return &CircuitBreakerWorkerPool{
		workers: workers,
		jobs:    make(chan Job, workers*10),
		results: make(chan Result, workers*10),
		errors:  make(chan error, workers*10),
		breaker: NewCircuitBreaker(threshold, timeout),
	}
}

func (cbwp *CircuitBreakerWorkerPool) Start() {
	for i := 0; i < cbwp.workers; i++ {
		cbwp.wg.Add(1)
		go cbwp.worker(i)
	}
	
	go func() {
		cbwp.wg.Wait()
		close(cbwp.results)
		close(cbwp.errors)
	}()
}

func (cbwp *CircuitBreakerWorkerPool) worker(workerID int) {
	defer cbwp.wg.Done()
	
	for job := range cbwp.jobs {
		if cbwp.breaker.Allow() {
			start := time.Now()
			
			// Simulate work with occasional failures
			time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
			
			if job.ID%5 == 0 {
				// Simulate failure
				cbwp.breaker.RecordFailure()
				cbwp.errors <- fmt.Errorf("worker %d failed to process job %d", workerID, job.ID)
			} else {
				cbwp.breaker.RecordSuccess()
				result := Result{
					JobID:    job.ID,
					Data:     fmt.Sprintf("Circuit Breaker: %s", job.Data),
					Duration: time.Since(start),
					WorkerID: workerID,
				}
				cbwp.results <- result
			}
		} else {
			// Circuit breaker is open
			cbwp.errors <- fmt.Errorf("circuit breaker open, job %d rejected", job.ID)
		}
	}
}

func (cbwp *CircuitBreakerWorkerPool) Submit(job Job) {
	cbwp.jobs <- job
}

func (cbwp *CircuitBreakerWorkerPool) Close() {
	close(cbwp.jobs)
}

// Advanced Pattern 4: Priority Queue Worker Pool
type PriorityQueueWorkerPool struct {
	workers      int
	highPriority chan Job
	lowPriority  chan Job
	results      chan Result
	mu           sync.RWMutex
	wg           sync.WaitGroup
}

func NewPriorityQueueWorkerPool(workers int) *PriorityQueueWorkerPool {
	return &PriorityQueueWorkerPool{
		workers:      workers,
		highPriority: make(chan Job, workers*5),
		lowPriority:  make(chan Job, workers*5),
		results:      make(chan Result, workers*10),
	}
}

func (pqwp *PriorityQueueWorkerPool) Start() {
	for i := 0; i < pqwp.workers; i++ {
		pqwp.wg.Add(1)
		go pqwp.worker(i)
	}
	
	go func() {
		pqwp.wg.Wait()
		close(pqwp.results)
	}()
}

func (pqwp *PriorityQueueWorkerPool) worker(workerID int) {
	defer pqwp.wg.Done()
	
	for {
		select {
		case job, ok := <-pqwp.highPriority:
			if !ok {
				pqwp.highPriority = nil
			} else {
				pqwp.processJob(job, workerID, "HIGH")
			}
		case job, ok := <-pqwp.lowPriority:
			if !ok {
				pqwp.lowPriority = nil
			} else {
				pqwp.processJob(job, workerID, "LOW")
			}
		}
		
		if pqwp.highPriority == nil && pqwp.lowPriority == nil {
			break
		}
	}
}

func (pqwp *PriorityQueueWorkerPool) processJob(job Job, workerID int, priority string) {
	start := time.Now()
	
	// Simulate work
	time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
	
	result := Result{
		JobID:    job.ID,
		Data:     fmt.Sprintf("Priority %s: %s", priority, job.Data),
		Duration: time.Since(start),
		WorkerID: workerID,
	}
	
	pqwp.results <- result
}

func (pqwp *PriorityQueueWorkerPool) Submit(job Job) {
	if job.Priority > 0 {
		pqwp.highPriority <- job
	} else {
		pqwp.lowPriority <- job
	}
}

func (pqwp *PriorityQueueWorkerPool) Close() {
	close(pqwp.highPriority)
	close(pqwp.lowPriority)
}

// Advanced Pattern 5: Load Balancing Worker Pool
type LoadBalancingWorkerPool struct {
	workers    int
	jobs       chan Job
	results    chan Result
	loads      []int64
	mu         sync.RWMutex
	wg         sync.WaitGroup
}

func NewLoadBalancingWorkerPool(workers int) *LoadBalancingWorkerPool {
	return &LoadBalancingWorkerPool{
		workers: workers,
		jobs:    make(chan Job, workers*10),
		results: make(chan Result, workers*10),
		loads:   make([]int64, workers),
	}
}

func (lbwp *LoadBalancingWorkerPool) Start() {
	for i := 0; i < lbwp.workers; i++ {
		lbwp.wg.Add(1)
		go lbwp.worker(i)
	}
	
	go func() {
		lbwp.wg.Wait()
		close(lbwp.results)
	}()
}

func (lbwp *LoadBalancingWorkerPool) worker(workerID int) {
	defer lbwp.wg.Done()
	
	for job := range lbwp.jobs {
		start := time.Now()
		
		// Simulate work
		time.Sleep(time.Duration(job.ID*10) * time.Millisecond)
		
		result := Result{
			JobID:    job.ID,
			Data:     fmt.Sprintf("Load Balanced: %s", job.Data),
			Duration: time.Since(start),
			WorkerID: workerID,
		}
		
		lbwp.results <- result
		
		// Update load
		lbwp.mu.Lock()
		lbwp.loads[workerID]++
		lbwp.mu.Unlock()
	}
}

func (lbwp *LoadBalancingWorkerPool) Submit(job Job) {
	// Find worker with least load
	lbwp.mu.RLock()
	minLoad := lbwp.loads[0]
	for i := 1; i < lbwp.workers; i++ {
		if lbwp.loads[i] < minLoad {
			minLoad = lbwp.loads[i]
		}
	}
	lbwp.mu.RUnlock()
	
	// Submit to least loaded worker (simplified - in real implementation, you'd route to specific workers)
	lbwp.jobs <- job
}

func (lbwp *LoadBalancingWorkerPool) Close() {
	close(lbwp.jobs)
}

// Advanced Pattern 6: Batch Processing Worker Pool
type BatchProcessingWorkerPool struct {
	workers    int
	jobs       chan Job
	results    chan Result
	batchSize  int
	batchTimeout time.Duration
	mu         sync.RWMutex
	wg         sync.WaitGroup
}

func NewBatchProcessingWorkerPool(workers, batchSize int, batchTimeout time.Duration) *BatchProcessingWorkerPool {
	return &BatchProcessingWorkerPool{
		workers:      workers,
		jobs:         make(chan Job, workers*10),
		results:      make(chan Result, workers*10),
		batchSize:    batchSize,
		batchTimeout: batchTimeout,
	}
}

func (bpwp *BatchProcessingWorkerPool) Start() {
	for i := 0; i < bpwp.workers; i++ {
		bpwp.wg.Add(1)
		go bpwp.worker(i)
	}
	
	go func() {
		bpwp.wg.Wait()
		close(bpwp.results)
	}()
}

func (bpwp *BatchProcessingWorkerPool) worker(workerID int) {
	defer bpwp.wg.Done()
	
	batch := make([]Job, 0, bpwp.batchSize)
	timer := time.NewTimer(bpwp.batchTimeout)
	defer timer.Stop()
	
	for {
		select {
		case job, ok := <-bpwp.jobs:
			if !ok {
				// Process remaining batch
				if len(batch) > 0 {
					bpwp.processBatch(batch, workerID)
				}
				return
			}
			
			batch = append(batch, job)
			if len(batch) >= bpwp.batchSize {
				bpwp.processBatch(batch, workerID)
				batch = batch[:0]
				timer.Reset(bpwp.batchTimeout)
			}
			
		case <-timer.C:
			if len(batch) > 0 {
				bpwp.processBatch(batch, workerID)
				batch = batch[:0]
			}
			timer.Reset(bpwp.batchTimeout)
		}
	}
}

func (bpwp *BatchProcessingWorkerPool) processBatch(batch []Job, workerID int) {
	start := time.Now()
	
	// Simulate batch processing
	time.Sleep(time.Duration(len(batch)*10) * time.Millisecond)
	
	for _, job := range batch {
		result := Result{
			JobID:    job.ID,
			Data:     fmt.Sprintf("Batch: %s", job.Data),
			Duration: time.Since(start),
			WorkerID: workerID,
		}
		bpwp.results <- result
	}
}

func (bpwp *BatchProcessingWorkerPool) Submit(job Job) {
	bpwp.jobs <- job
}

func (bpwp *BatchProcessingWorkerPool) Close() {
	close(bpwp.jobs)
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Worker Pool Patterns")
	fmt.Println("=================================")
	
	// Pattern 1: Work Stealing Worker Pool
	fmt.Println("\n1. Work Stealing Worker Pool:")
	workStealingPool := NewWorkStealingPool(3)
	workStealingPool.Start()
	
	// Submit jobs
	for i := 0; i < 10; i++ {
		workStealingPool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("Work Stealing Job %d", i),
		})
	}
	
	// Collect results
	for i := 0; i < 10; i++ {
		select {
		case result := <-workStealingPool.results:
			fmt.Printf("  Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case <-time.After(1 * time.Second):
			break
		}
	}
	
	workStealingPool.Close()
	
	// Pattern 2: Adaptive Worker Pool
	fmt.Println("\n2. Adaptive Worker Pool:")
	adaptivePool := NewAdaptiveWorkerPool(2, 5)
	
	// Submit jobs
	for i := 0; i < 15; i++ {
		adaptivePool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("Adaptive Job %d", i),
		})
		time.Sleep(50 * time.Millisecond)
	}
	
	// Collect results
	for i := 0; i < 15; i++ {
		select {
		case result := <-adaptivePool.results:
			fmt.Printf("  Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case <-time.After(1 * time.Second):
			break
		}
	}
	
	adaptivePool.Close()
	
	// Pattern 3: Circuit Breaker Worker Pool
	fmt.Println("\n3. Circuit Breaker Worker Pool:")
	circuitBreakerPool := NewCircuitBreakerWorkerPool(3, 3, 1*time.Second)
	circuitBreakerPool.Start()
	
	// Submit jobs
	for i := 0; i < 10; i++ {
		circuitBreakerPool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("Circuit Breaker Job %d", i),
		})
	}
	
	// Collect results and errors
	for i := 0; i < 10; i++ {
		select {
		case result := <-circuitBreakerPool.results:
			fmt.Printf("  SUCCESS: Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case err := <-circuitBreakerPool.errors:
			fmt.Printf("  ERROR: %v\n", err)
		case <-time.After(1 * time.Second):
			break
		}
	}
	
	circuitBreakerPool.Close()
	
	// Pattern 4: Priority Queue Worker Pool
	fmt.Println("\n4. Priority Queue Worker Pool:")
	priorityPool := NewPriorityQueueWorkerPool(3)
	priorityPool.Start()
	
	// Submit jobs with different priorities
	for i := 0; i < 10; i++ {
		priorityPool.Submit(Job{
			ID:       i,
			Data:     fmt.Sprintf("Priority Job %d", i),
			Priority: i % 2, // Alternate priorities
		})
	}
	
	// Collect results
	for i := 0; i < 10; i++ {
		select {
		case result := <-priorityPool.results:
			fmt.Printf("  Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case <-time.After(1 * time.Second):
			break
		}
	}
	
	priorityPool.Close()
	
	// Pattern 5: Load Balancing Worker Pool
	fmt.Println("\n5. Load Balancing Worker Pool:")
	loadBalancingPool := NewLoadBalancingWorkerPool(3)
	loadBalancingPool.Start()
	
	// Submit jobs
	for i := 0; i < 10; i++ {
		loadBalancingPool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("Load Balanced Job %d", i),
		})
	}
	
	// Collect results
	for i := 0; i < 10; i++ {
		select {
		case result := <-loadBalancingPool.results:
			fmt.Printf("  Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case <-time.After(1 * time.Second):
			break
		}
	}
	
	loadBalancingPool.Close()
	
	// Pattern 6: Batch Processing Worker Pool
	fmt.Println("\n6. Batch Processing Worker Pool:")
	batchPool := NewBatchProcessingWorkerPool(3, 3, 500*time.Millisecond)
	batchPool.Start()
	
	// Submit jobs
	for i := 0; i < 10; i++ {
		batchPool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("Batch Job %d", i),
		})
	}
	
	// Collect results
	for i := 0; i < 10; i++ {
		select {
		case result := <-batchPool.results:
			fmt.Printf("  Job %d: %s (worker %d)\n", result.JobID, result.Data, result.WorkerID)
		case <-time.After(2 * time.Second):
			break
		}
	}
	
	batchPool.Close()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
