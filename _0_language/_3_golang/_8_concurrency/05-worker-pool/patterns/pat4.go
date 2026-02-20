package patterns

import (
	"fmt"
	"sync"
	"time"
)

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