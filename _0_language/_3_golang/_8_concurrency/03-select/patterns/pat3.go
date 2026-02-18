package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 3: Select-based Load Balancer
type SelectLoadBalancer struct {
	workers []chan Job
	current int64
	mu      sync.Mutex
}

type Job struct {
	ID   int
	Data interface{}
}

func NewSelectLoadBalancer(workers int) *SelectLoadBalancer {
	lb := &SelectLoadBalancer{
		workers: make([]chan Job, workers),
	}
	
	for i := 0; i < workers; i++ {
		lb.workers[i] = make(chan Job, 10)
		go lb.worker(i)
	}
	
	return lb
}

func (lb *SelectLoadBalancer) worker(id int) {
	for job := range lb.workers[id] {
		fmt.Printf("Worker %d: Processing job %d\n", id, job.ID)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func (lb *SelectLoadBalancer) Submit(job Job) {
	// Round-robin selection
	current := atomic.AddInt64(&lb.current, 1)
	worker := int(current) % len(lb.workers)
	
	select {
	case lb.workers[worker] <- job:
		// Job submitted successfully
	default:
		// Worker is busy, try next worker
		nextWorker := (int(current) + 1) % len(lb.workers)
		select {
		case lb.workers[nextWorker] <- job:
			// Job submitted to next worker
		default:
			// All workers busy, drop job
			fmt.Printf("All workers busy, dropping job %d\n", job.ID)
		}
	}
}

func (lb *SelectLoadBalancer) Stop() {
	for _, worker := range lb.workers {
		close(worker)
	}
}