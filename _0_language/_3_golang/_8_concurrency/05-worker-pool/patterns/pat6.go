package patterns

import (
	"fmt"
	"sync"
	"time"
)

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