package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 6: Select-based Worker Pool with Priority
type SelectPriorityPool struct {
	highJobs   chan Job
	normalJobs chan Job
	lowJobs    chan Job
	quitCh     chan bool
	wg         sync.WaitGroup
}

func NewSelectPriorityPool(workers int) *SelectPriorityPool {
	pool := &SelectPriorityPool{
		highJobs:   make(chan Job, 100),
		normalJobs: make(chan Job, 100),
		lowJobs:    make(chan Job, 100),
		quitCh:     make(chan bool),
	}
	
	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go pool.worker(i)
	}
	
	return pool
}

func (pool *SelectPriorityPool) worker(id int) {
	defer pool.wg.Done()
	
	for {
		select {
		case job := <-pool.highJobs:
			pool.processJob(id, job, "HIGH")
		case job := <-pool.normalJobs:
			pool.processJob(id, job, "NORMAL")
		case job := <-pool.lowJobs:
			pool.processJob(id, job, "LOW")
		case <-pool.quitCh:
			return
		}
	}
}

func (pool *SelectPriorityPool) processJob(workerID int, job Job, priority string) {
	fmt.Printf("Worker %d: Processing %s priority job %d\n", workerID, priority, job.ID)
	time.Sleep(100 * time.Millisecond) // Simulate work
}

func (pool *SelectPriorityPool) Submit(job Job, priority string) {
	switch priority {
	case "high":
		select {
		case pool.highJobs <- job:
		default:
			fmt.Printf("High priority queue full, dropping job %d\n", job.ID)
		}
	case "normal":
		select {
		case pool.normalJobs <- job:
		default:
			fmt.Printf("Normal priority queue full, dropping job %d\n", job.ID)
		}
	case "low":
		select {
		case pool.lowJobs <- job:
		default:
			fmt.Printf("Low priority queue full, dropping job %d\n", job.ID)
		}
	}
}

func (pool *SelectPriorityPool) Stop() {
	close(pool.quitCh)
	pool.wg.Wait()
}