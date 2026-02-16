package patterns

import "sync"

// Advanced Pattern 1: Goroutine Pool with Dynamic Scaling
type DynamicPool struct {
	workers    int
	maxWorkers int
	jobs       chan func()
	quit       chan bool
	wg         sync.WaitGroup
	mu         sync.RWMutex
}

func NewDynamicPool(maxWorkers int) *DynamicPool {
	return &DynamicPool{
		workers:    0,
		maxWorkers: maxWorkers,
		jobs:       make(chan func(), 100),
		quit:       make(chan bool),
	}
}

func (p *DynamicPool) Start() {
	go p.manager()
}

func (p *DynamicPool) manager() {
	for {
		select {
		case job := <-p.jobs:
			p.mu.RLock()
			workers := p.workers
			p.mu.RUnlock()
			
			if workers < p.maxWorkers {
				p.addWorker()
			}
			
			// Send job to worker
			select {
			case p.jobs <- job:
			default:
				// If jobs channel is full, create more workers
				if workers < p.maxWorkers {
					p.addWorker()
					p.jobs <- job
				}
			}
		case <-p.quit:
			return
		}
	}
}

func (p *DynamicPool) addWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.workers < p.maxWorkers {
		p.workers++
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *DynamicPool) worker() {
	defer p.wg.Done()
	
	for {
		select {
		case job := <-p.jobs:
			job()
		case <-p.quit:
			return
		}
	}
}

func (p *DynamicPool) Submit(job func()) {
	p.jobs <- job
}

func (p *DynamicPool) Stop() {
	close(p.quit)
	p.wg.Wait()
}