package patterns

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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