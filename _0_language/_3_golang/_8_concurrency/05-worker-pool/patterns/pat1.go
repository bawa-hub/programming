package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Work Stealing Worker Pool

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