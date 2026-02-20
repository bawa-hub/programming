package patterns

import (
	"fmt"
	"sync"
	"time"
)

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