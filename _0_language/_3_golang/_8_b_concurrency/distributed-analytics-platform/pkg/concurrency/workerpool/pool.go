package workerpool

import (
	"context"
	"sync"
)

// Pool represents a worker pool
type Pool struct {
	workers    int
	tasks      chan Task
	results    chan Result
	quit       chan struct{}
	wg         sync.WaitGroup
	workerFunc WorkerFunc
	mu         sync.RWMutex
	isRunning  bool
}

// Task represents a unit of work
type Task struct {
	ID      string
	Data    interface{}
	Handler func(interface{}) (interface{}, error)
}

// Result represents the outcome of a task
type Result struct {
	TaskID string
	Data   interface{}
	Error  error
}

// WorkerFunc is the function that a worker executes
type WorkerFunc func(ctx context.Context, task Task) (Result, error)

// NewPool creates a new worker pool
func NewPool(workers int) *Pool {
	return &Pool{
		workers: workers,
		tasks:   make(chan Task, workers*2), // Buffer for tasks
		results: make(chan Result, workers*2), // Buffer for results
		quit:    make(chan struct{}),
	}
}

// Start starts the worker pool
func (p *Pool) Start() {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.isRunning {
		return
	}
	
	p.isRunning = true
	
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

// Stop stops the worker pool
func (p *Pool) Stop() {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if !p.isRunning {
		return
	}
	
	p.isRunning = false
	close(p.quit)
	p.wg.Wait()
	close(p.results)
}

// Submit submits a task to the worker pool
func (p *Pool) Submit(task Task) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	if !p.isRunning {
		return ErrPoolNotRunning
	}
	
	select {
	case p.tasks <- task:
		return nil
	case <-p.quit:
		return ErrPoolStopped
	default:
		return ErrPoolFull
	}
}

// Results returns the results channel
func (p *Pool) Results() <-chan Result {
	return p.results
}

// GetStats returns pool statistics
func (p *Pool) GetStats() PoolStats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	return PoolStats{
		Workers:     p.workers,
		TaskQueue:   len(p.tasks),
		ResultQueue: len(p.results),
		IsRunning:   p.isRunning,
	}
}

// worker is a worker goroutine
func (p *Pool) worker(id int) {
	defer p.wg.Done()
	
	ctx := context.Background()
	
	for {
		select {
		case task := <-p.tasks:
			// Process task
			result := Result{
				TaskID: task.ID,
			}
			
			if task.Handler != nil {
				data, err := task.Handler(task.Data)
				result.Data = data
				result.Error = err
			} else if p.workerFunc != nil {
				data, err := p.workerFunc(ctx, task)
				result.Data = data
				result.Error = err
			}
			
			// Send result
			select {
			case p.results <- result:
			case <-p.quit:
				return
			}
			
		case <-p.quit:
			return
		}
	}
}

// PoolStats represents pool statistics
type PoolStats struct {
	Workers     int
	TaskQueue   int
	ResultQueue int
	IsRunning   bool
}

// Pool errors
var (
	ErrPoolNotRunning = &PoolError{msg: "pool is not running"}
	ErrPoolStopped    = &PoolError{msg: "pool is stopped"}
	ErrPoolFull       = &PoolError{msg: "pool is full"}
)

// PoolError represents a pool-related error
type PoolError struct {
	msg string
}

func (e *PoolError) Error() string {
	return e.msg
}