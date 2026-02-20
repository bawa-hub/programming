package patterns

import (
	"fmt"
	"sync"
	"time"
)

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