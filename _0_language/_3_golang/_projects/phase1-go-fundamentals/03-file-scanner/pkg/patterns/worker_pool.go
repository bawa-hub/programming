package patterns

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// WorkerPool represents a pool of workers for concurrent processing
type WorkerPool struct {
	workerCount int
	jobQueue    chan Job
	resultQueue chan Result
	errorQueue  chan Error
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	mu          sync.RWMutex
	running     bool
	stats       *PoolStats
}

// Job represents a unit of work
type Job interface {
	ID() string
	Process() (Result, error)
	Timeout() time.Duration
}

// Result represents the result of a job
type Result interface {
	JobID() string
	Data() interface{}
	Duration() time.Duration
	Success() bool
}

// Error represents an error that occurred during job processing
type Error struct {
	JobID    string
	Error    error
	Time     time.Time
	WorkerID int
}

// PoolStats represents statistics about the worker pool
type PoolStats struct {
	TotalJobs     int64
	CompletedJobs int64
	FailedJobs    int64
	ActiveWorkers int
	QueueSize     int
	StartTime     time.Time
	EndTime       time.Time
	mu            sync.RWMutex
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	if workerCount <= 0 {
		workerCount = runtime.NumCPU()
	}
	if queueSize <= 0 {
		queueSize = workerCount * 2
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	
	return &WorkerPool{
		workerCount: workerCount,
		jobQueue:    make(chan Job, queueSize),
		resultQueue: make(chan Result, queueSize),
		errorQueue:  make(chan Error, queueSize),
		ctx:         ctx,
		cancel:      cancel,
		stats: &PoolStats{
			StartTime: time.Now(),
		},
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	
	if wp.running {
		return
	}
	
	wp.running = true
	
	// Start workers
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	
	if !wp.running {
		return
	}
	
	wp.running = false
	wp.cancel()
	close(wp.jobQueue)
	wp.wg.Wait()
	close(wp.resultQueue)
	close(wp.errorQueue)
	
	wp.stats.mu.Lock()
	wp.stats.EndTime = time.Now()
	wp.stats.ActiveWorkers = 0
	wp.stats.QueueSize = 0
	wp.stats.mu.Unlock()
}

// Submit submits a job to the worker pool
func (wp *WorkerPool) Submit(job Job) error {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	
	if !wp.running {
		return fmt.Errorf("worker pool is not running")
	}
	
	select {
	case wp.jobQueue <- job:
		wp.stats.mu.Lock()
		wp.stats.TotalJobs++
		wp.stats.QueueSize = len(wp.jobQueue)
		wp.stats.mu.Unlock()
		return nil
	case <-wp.ctx.Done():
		return fmt.Errorf("worker pool is shutting down")
	default:
		return fmt.Errorf("job queue is full")
	}
}

// GetResult returns the next result
func (wp *WorkerPool) GetResult() (Result, bool) {
	result, ok := <-wp.resultQueue
	return result, ok
}

// GetError returns the next error
func (wp *WorkerPool) GetError() (Error, bool) {
	error, ok := <-wp.errorQueue
	return error, ok
}

// GetStats returns current pool statistics
func (wp *WorkerPool) GetStats() PoolStats {
	wp.stats.mu.RLock()
	defer wp.stats.mu.RUnlock()
	
	stats := *wp.stats
	stats.ActiveWorkers = wp.workerCount
	stats.QueueSize = len(wp.jobQueue)
	
	return stats
}

// WaitForCompletion waits for all jobs to complete
func (wp *WorkerPool) WaitForCompletion() {
	wp.wg.Wait()
}

// IsRunning returns true if the worker pool is running
func (wp *WorkerPool) IsRunning() bool {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	return wp.running
}

// worker is the main worker function
func (wp *WorkerPool) worker(workerID int) {
	defer wp.wg.Done()
	
	for {
		select {
		case job, ok := <-wp.jobQueue:
			if !ok {
				return
			}
			
			wp.processJob(workerID, job)
			
		case <-wp.ctx.Done():
			return
		}
	}
}

// processJob processes a single job
func (wp *WorkerPool) processJob(workerID int, job Job) {
	start := time.Now()
	
	// Process the job
	result, err := job.Process()
	
	_ = time.Since(start)
	
	// Update statistics
	wp.stats.mu.Lock()
	wp.stats.QueueSize = len(wp.jobQueue)
	wp.stats.mu.Unlock()
	
	if err != nil {
		// Send error
		select {
		case wp.errorQueue <- Error{
			JobID:    job.ID(),
			Error:    err,
			Time:     time.Now(),
			WorkerID: workerID,
		}:
		case <-wp.ctx.Done():
			return
		}
		
		wp.stats.mu.Lock()
		wp.stats.FailedJobs++
		wp.stats.mu.Unlock()
	} else {
		// Send result
		select {
		case wp.resultQueue <- result:
		case <-wp.ctx.Done():
			return
		}
		
		wp.stats.mu.Lock()
		wp.stats.CompletedJobs++
		wp.stats.mu.Unlock()
	}
}

// Pipeline represents a processing pipeline
type Pipeline struct {
	stages []Stage
	mu     sync.RWMutex
}

// Stage represents a pipeline stage
type Stage interface {
	Name() string
	Process(input interface{}) (interface{}, error)
	Concurrency() int
}

// NewPipeline creates a new pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{
		stages: make([]Stage, 0),
	}
}

// AddStage adds a stage to the pipeline
func (p *Pipeline) AddStage(stage Stage) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.stages = append(p.stages, stage)
}

// Process processes data through the pipeline
func (p *Pipeline) Process(input interface{}) (interface{}, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	result := input
	var err error
	
	for _, stage := range p.stages {
		result, err = stage.Process(result)
		if err != nil {
			return nil, fmt.Errorf("stage %s failed: %w", stage.Name(), err)
		}
	}
	
	return result, nil
}

// ProcessConcurrent processes data through the pipeline concurrently
func (p *Pipeline) ProcessConcurrent(inputs []interface{}) ([]interface{}, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	if len(p.stages) == 0 {
		return inputs, nil
	}
	
	// Create worker pools for each stage
	pools := make([]*WorkerPool, len(p.stages))
	for i, stage := range p.stages {
		pools[i] = NewWorkerPool(stage.Concurrency(), len(inputs))
		pools[i].Start()
	}
	
	// Clean up pools
	defer func() {
		for _, pool := range pools {
			pool.Stop()
		}
	}()
	
	// Process through stages
	currentInputs := inputs
	for stageIndex, stage := range p.stages {
		pool := pools[stageIndex]
		
		// Submit jobs for current stage
		for _, input := range currentInputs {
			job := &PipelineJob{
				Stage: stage,
				Input: input,
			}
			if err := pool.Submit(job); err != nil {
				return nil, fmt.Errorf("failed to submit job to stage %s: %w", stage.Name(), err)
			}
		}
		
		// Collect results
		results := make([]interface{}, 0, len(currentInputs))
		for i := 0; i < len(currentInputs); i++ {
			result, ok := pool.GetResult()
			if !ok {
				return nil, fmt.Errorf("failed to get result from stage %s", stage.Name())
			}
			results = append(results, result.Data())
		}
		
		currentInputs = results
	}
	
	return currentInputs, nil
}

// PipelineJob represents a job in the pipeline
type PipelineJob struct {
	Stage Stage
	Input interface{}
}

// ID returns the job ID
func (pj *PipelineJob) ID() string {
	return fmt.Sprintf("pipeline_%s_%p", pj.Stage.Name(), pj.Input)
}

// Process processes the job
func (pj *PipelineJob) Process() (Result, error) {
	start := time.Now()
	
	result, err := pj.Stage.Process(pj.Input)
	if err != nil {
		return nil, err
	}
	
	return &PipelineResult{
		JobIDValue:    pj.ID(),
		DataValue:     result,
		DurationValue: time.Since(start),
		SuccessValue:  true,
	}, nil
}

// Timeout returns the job timeout
func (pj *PipelineJob) Timeout() time.Duration {
	return 30 * time.Second // Default timeout
}

// PipelineResult represents a pipeline result
type PipelineResult struct {
	JobIDValue    string
	DataValue     interface{}
	DurationValue time.Duration
	SuccessValue  bool
}

// JobID returns the job ID
func (pr *PipelineResult) JobID() string {
	return pr.JobIDValue
}

// Data returns the result data
func (pr *PipelineResult) Data() interface{} {
	return pr.DataValue
}

// Duration returns the processing duration
func (pr *PipelineResult) Duration() time.Duration {
	return pr.DurationValue
}

// Success returns true if the job was successful
func (pr *PipelineResult) Success() bool {
	return pr.SuccessValue
}

// RateLimiter represents a rate limiter
type RateLimiter struct {
	rate     time.Duration
	burst    int
	tokens   int
	lastTime time.Time
	mu       sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		rate:     rate,
		burst:    burst,
		tokens:   burst,
		lastTime: time.Now(),
	}
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(rl.lastTime)
	
	// Add tokens based on elapsed time
	tokensToAdd := int(elapsed / rl.rate)
	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.burst {
			rl.tokens = rl.burst
		}
		rl.lastTime = now
	}
	
	// Check if we have tokens
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	
	return false
}

// Wait waits for a token to become available
func (rl *RateLimiter) Wait() {
	for !rl.Allow() {
		time.Sleep(rl.rate)
	}
}

// Semaphore represents a counting semaphore
type Semaphore struct {
	permits int
	channel chan struct{}
}

// NewSemaphore creates a new semaphore
func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{
		permits: permits,
		channel: make(chan struct{}, permits),
	}
}

// Acquire acquires a permit
func (s *Semaphore) Acquire() {
	s.channel <- struct{}{}
}

// Release releases a permit
func (s *Semaphore) Release() {
	<-s.channel
}

// TryAcquire tries to acquire a permit without blocking
func (s *Semaphore) TryAcquire() bool {
	select {
	case s.channel <- struct{}{}:
		return true
	default:
		return false
	}
}

// Available returns the number of available permits
func (s *Semaphore) Available() int {
	return s.permits - len(s.channel)
}

// AcquireWithTimeout acquires a permit with timeout
func (s *Semaphore) AcquireWithTimeout(timeout time.Duration) bool {
	select {
	case s.channel <- struct{}{}:
		return true
	case <-time.After(timeout):
		return false
	}
}
