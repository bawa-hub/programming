package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 5: Channel-based Worker Pool with Priority
type PriorityWorkerPool struct {
	workers    int
	highJobs   chan Job
	normalJobs chan Job
	lowJobs    chan Job
	quitCh     chan bool
	wg         sync.WaitGroup
}

type Job struct {
	ID       int
	Priority int // 1: high, 2: normal, 3: low
	Data     interface{}
	Result   chan interface{}
}

func NewPriorityWorkerPool(workers int) *PriorityWorkerPool {
	return &PriorityWorkerPool{
		workers:    workers,
		highJobs:   make(chan Job, 100),
		normalJobs: make(chan Job, 100),
		lowJobs:    make(chan Job, 100),
		quitCh:     make(chan bool),
	}
}

func (pwp *PriorityWorkerPool) Start() {
	for i := 0; i < pwp.workers; i++ {
		pwp.wg.Add(1)
		go pwp.worker(i)
	}
}

func (pwp *PriorityWorkerPool) worker(id int) {
	defer pwp.wg.Done()
	
	for {
		select {
		case job := <-pwp.highJobs:
			pwp.processJob(id, job)
		case job := <-pwp.normalJobs:
			pwp.processJob(id, job)
		case job := <-pwp.lowJobs:
			pwp.processJob(id, job)
		case <-pwp.quitCh:
			return
		}
	}
}

func (pwp *PriorityWorkerPool) processJob(workerID int, job Job) {
	fmt.Printf("Worker %d: Processing job %d (priority %d)\n", workerID, job.ID, job.Priority)
	time.Sleep(100 * time.Millisecond) // Simulate work
	job.Result <- fmt.Sprintf("Result for job %d", job.ID)
}

func (pwp *PriorityWorkerPool) Submit(job Job) {
	switch job.Priority {
	case 1:
		pwp.highJobs <- job
	case 2:
		pwp.normalJobs <- job
	case 3:
		pwp.lowJobs <- job
	}
}

func (pwp *PriorityWorkerPool) Stop() {
	close(pwp.quitCh)
	pwp.wg.Wait()
}