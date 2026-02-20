package patterns

import (
	"sync"
	"time"
)

// Advanced Pattern 3: Pipeline with Circuit Breaker
type CircuitBreakerPipeline struct {
	stages      []*CircuitBreakerStage
	breakers    []*PipelineCircuitBreaker
	mu          sync.RWMutex
}

type CircuitBreakerStage struct {
	name        string
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) (ProcessedData, error)
	breaker     *PipelineCircuitBreaker
}

type PipelineCircuitBreaker struct {
	failures    int64
	successes   int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
}

func NewCircuitBreakerPipeline() *CircuitBreakerPipeline {
	return &CircuitBreakerPipeline{
		stages:   make([]*CircuitBreakerStage, 0),
		breakers: make([]*PipelineCircuitBreaker, 0),
	}
}

func (cbp *CircuitBreakerPipeline) AddStage(name string, threshold int64, timeout time.Duration, processFunc func(ProcessedData) (ProcessedData, error)) {
	breaker := &PipelineCircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
	}
	
	stage := &CircuitBreakerStage{
		name:        name,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		breaker:     breaker,
	}
	
	// Start stage worker
	go cbp.stageWorker(stage)
	
	cbp.stages = append(cbp.stages, stage)
	cbp.breakers = append(cbp.breakers, breaker)
}

func (cbp *CircuitBreakerPipeline) stageWorker(stage *CircuitBreakerStage) {
	for data := range stage.input {
		if stage.breaker.Allow() {
			processed, err := stage.processFunc(data)
			if err != nil {
				stage.breaker.RecordFailure()
				// Skip this item or handle error
				continue
			} else {
				stage.breaker.RecordSuccess()
				stage.output <- processed
			}
		} else {
			// Circuit breaker is open, skip or use fallback
			continue
		}
	}
}

func (pcb *PipelineCircuitBreaker) Allow() bool {
	pcb.mu.RLock()
	defer pcb.mu.RUnlock()
	
	if pcb.state == 0 { // closed
		return true
	} else if pcb.state == 1 { // open
		if time.Since(pcb.lastFailure) > pcb.timeout {
			pcb.state = 2 // half-open
			return true
		}
		return false
	} else { // half-open
		return true
	}
}

func (pcb *PipelineCircuitBreaker) RecordSuccess() {
	pcb.mu.Lock()
	defer pcb.mu.Unlock()
	
	pcb.successes++
	if pcb.state == 2 { // half-open
		pcb.state = 0 // closed
		pcb.failures = 0
	}
}

func (pcb *PipelineCircuitBreaker) RecordFailure() {
	pcb.mu.Lock()
	defer pcb.mu.Unlock()
	
	pcb.failures++
	pcb.lastFailure = time.Now()
	
	if pcb.failures >= pcb.threshold {
		pcb.state = 1 // open
	}
}

func (cbp *CircuitBreakerPipeline) Submit(data ProcessedData) {
	if len(cbp.stages) > 0 {
		cbp.stages[0].input <- data
	}
}

func (cbp *CircuitBreakerPipeline) Close() {
	for _, stage := range cbp.stages {
		close(stage.input)
	}
}