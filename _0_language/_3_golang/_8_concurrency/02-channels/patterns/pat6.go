package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 6: Channel-based Load Balancer
type LoadBalancer struct {
	workers []chan Job
	current int
	mu      sync.Mutex
}

func NewLoadBalancer(workers int) *LoadBalancer {
	lb := &LoadBalancer{
		workers: make([]chan Job, workers),
	}
	
	for i := 0; i < workers; i++ {
		lb.workers[i] = make(chan Job, 10)
		go lb.worker(i)
	}
	
	return lb
}

func (lb *LoadBalancer) worker(id int) {
	for job := range lb.workers[id] {
		fmt.Printf("Worker %d: Processing job %d\n", id, job.ID)
		time.Sleep(100 * time.Millisecond) // Simulate work
		job.Result <- fmt.Sprintf("Result from worker %d for job %d", id, job.ID)
	}
}

func (lb *LoadBalancer) Submit(job Job) {
	lb.mu.Lock()
	worker := lb.current
	lb.current = (lb.current + 1) % len(lb.workers)
	lb.mu.Unlock()
	
	lb.workers[worker] <- job
}

func (lb *LoadBalancer) Stop() {
	for _, worker := range lb.workers {
		close(worker)
	}
}