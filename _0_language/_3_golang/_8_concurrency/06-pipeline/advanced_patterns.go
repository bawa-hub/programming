package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Adaptive Pipeline
type AdaptivePipeline struct {
	stages       []*PipelineStage
	metrics      *AdaptiveMetrics
	mu           sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
}

type PipelineStage struct {
	name        string
	workers     int
	minWorkers  int
	maxWorkers  int
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	mu          sync.RWMutex
}

type AdaptiveMetrics struct {
	stageMetrics map[string]*StageMetrics
	mu           sync.RWMutex
}

type StageMetrics struct {
	processedItems int64
	processingTime time.Duration
	queueSize      int64
	workerUtilization float64
}

func NewAdaptivePipeline() *AdaptivePipeline {
	ctx, cancel := context.WithCancel(context.Background())
	
	pipeline := &AdaptivePipeline{
		stages:  make([]*PipelineStage, 0),
		metrics: &AdaptiveMetrics{
			stageMetrics: make(map[string]*StageMetrics),
		},
		ctx:    ctx,
		cancel: cancel,
	}
	
	go pipeline.adaptiveController()
	return pipeline
}

func (ap *AdaptivePipeline) AddStage(name string, minWorkers, maxWorkers int, processFunc func(ProcessedData) ProcessedData) {
	stage := &PipelineStage{
		name:        name,
		workers:     minWorkers,
		minWorkers:  minWorkers,
		maxWorkers:  maxWorkers,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
	}
	
	ap.stages = append(ap.stages, stage)
	ap.metrics.stageMetrics[name] = &StageMetrics{}
	
	// Start workers
	ap.startStageWorkers(stage)
}

func (ap *AdaptivePipeline) startStageWorkers(stage *PipelineStage) {
	for i := 0; i < stage.workers; i++ {
		go ap.stageWorker(stage, i)
	}
}

func (ap *AdaptivePipeline) stageWorker(stage *PipelineStage, workerID int) {
	for {
		select {
		case data, ok := <-stage.input:
			if !ok {
				return
			}
			
			start := time.Now()
			processed := stage.processFunc(data)
			duration := time.Since(start)
			
			// Update metrics
			ap.updateStageMetrics(stage.name, duration)
			
			// Send to next stage
			if len(ap.stages) > 0 {
				nextStageIndex := ap.getNextStageIndex(stage)
				if nextStageIndex < len(ap.stages) {
					ap.stages[nextStageIndex].input <- processed
				}
			}
			
		case <-ap.ctx.Done():
			return
		}
	}
}

func (ap *AdaptivePipeline) getNextStageIndex(currentStage *PipelineStage) int {
	for i, stage := range ap.stages {
		if stage == currentStage {
			return i + 1
		}
	}
	return -1
}

func (ap *AdaptivePipeline) updateStageMetrics(stageName string, duration time.Duration) {
	ap.metrics.mu.Lock()
	defer ap.metrics.mu.Unlock()
	
	if metrics, exists := ap.metrics.stageMetrics[stageName]; exists {
		metrics.processedItems++
		metrics.processingTime += duration
	}
}

func (ap *AdaptivePipeline) adaptiveController() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			ap.adjustStages()
		case <-ap.ctx.Done():
			return
		}
	}
}

func (ap *AdaptivePipeline) adjustStages() {
	ap.mu.Lock()
	defer ap.mu.Unlock()
	
	for _, stage := range ap.stages {
		queueSize := len(stage.input)
		utilization := ap.getStageUtilization(stage.name)
		
		if queueSize > 5 && stage.workers < stage.maxWorkers && utilization > 0.8 {
			// Add worker
			stage.workers++
			go ap.stageWorker(stage, stage.workers-1)
			fmt.Printf("Added worker to stage %s, total: %d\n", stage.name, stage.workers)
		} else if queueSize == 0 && stage.workers > stage.minWorkers && utilization < 0.3 {
			// Remove worker (simplified - in real implementation, you'd need to signal workers to stop)
			fmt.Printf("Would remove worker from stage %s, total: %d\n", stage.name, stage.workers)
		}
	}
}

func (ap *AdaptivePipeline) getStageUtilization(stageName string) float64 {
	ap.metrics.mu.RLock()
	defer ap.metrics.mu.RUnlock()
	
	if metrics, exists := ap.metrics.stageMetrics[stageName]; exists {
		return metrics.workerUtilization
	}
	return 0.0
}

func (ap *AdaptivePipeline) Submit(data ProcessedData) {
	if len(ap.stages) > 0 {
		ap.stages[0].input <- data
	}
}

func (ap *AdaptivePipeline) Close() {
	ap.cancel()
	for _, stage := range ap.stages {
		close(stage.input)
	}
}

// Advanced Pattern 2: Pipeline with Load Balancing
type LoadBalancedPipeline struct {
	stages    []*LoadBalancedStage
	balancers []*LoadBalancer
	mu        sync.RWMutex
}

type LoadBalancedStage struct {
	name        string
	workers     []*Worker
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	balancer    *LoadBalancer
}

type Worker struct {
	id          int
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	load        int64
	mu          sync.RWMutex
}

type LoadBalancer struct {
	workers []*Worker
	index   int64
	mu      sync.RWMutex
}

func NewLoadBalancedPipeline() *LoadBalancedPipeline {
	return &LoadBalancedPipeline{
		stages:    make([]*LoadBalancedStage, 0),
		balancers: make([]*LoadBalancer, 0),
	}
}

func (lbp *LoadBalancedPipeline) AddStage(name string, numWorkers int, processFunc func(ProcessedData) ProcessedData) {
	workers := make([]*Worker, numWorkers)
	balancer := &LoadBalancer{
		workers: workers,
		index:   0,
	}
	
	stage := &LoadBalancedStage{
		name:        name,
		workers:     workers,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		balancer:    balancer,
	}
	
	// Create workers
	for i := 0; i < numWorkers; i++ {
		worker := &Worker{
			id:          i,
			input:       make(chan ProcessedData, 10),
			output:      stage.output,
			processFunc: processFunc,
		}
		workers[i] = worker
		balancer.workers[i] = worker
		
		// Start worker
		go lbp.workerLoop(worker)
	}
	
	// Start load balancer
	go lbp.loadBalancerLoop(stage)
	
	lbp.stages = append(lbp.stages, stage)
	lbp.balancers = append(lbp.balancers, balancer)
}

func (lbp *LoadBalancedPipeline) workerLoop(worker *Worker) {
	for data := range worker.input {
		processed := worker.processFunc(data)
		
		// Update load
		worker.mu.Lock()
		worker.load++
		worker.mu.Unlock()
		
		worker.output <- processed
	}
}

func (lbp *LoadBalancedPipeline) loadBalancerLoop(stage *LoadBalancedStage) {
	for data := range stage.input {
		worker := stage.balancer.GetWorker()
		worker.input <- data
	}
}

func (lb *LoadBalancer) GetWorker() *Worker {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	// Round-robin selection
	worker := lb.workers[lb.index%int64(len(lb.workers))]
	lb.index++
	return worker
}

func (lbp *LoadBalancedPipeline) Submit(data ProcessedData) {
	if len(lbp.stages) > 0 {
		lbp.stages[0].input <- data
	}
}

func (lbp *LoadBalancedPipeline) Close() {
	for _, stage := range lbp.stages {
		close(stage.input)
		for _, worker := range stage.workers {
			close(worker.input)
		}
	}
}

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

// Advanced Pattern 4: Pipeline with Caching
type CachedPipeline struct {
	stages  []*CachedStage
	caches  []*PipelineCache
	mu      sync.RWMutex
}

type CachedStage struct {
	name        string
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	cache       *PipelineCache
}

type PipelineCache struct {
	data   map[string]ProcessedData
	hits   int64
	misses int64
	mu     sync.RWMutex
}

func NewCachedPipeline() *CachedPipeline {
	return &CachedPipeline{
		stages: make([]*CachedStage, 0),
		caches: make([]*PipelineCache, 0),
	}
}

func (cp *CachedPipeline) AddStage(name string, processFunc func(ProcessedData) ProcessedData) {
	cache := &PipelineCache{
		data: make(map[string]ProcessedData),
	}
	
	stage := &CachedStage{
		name:        name,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		cache:       cache,
	}
	
	// Start stage worker
	go cp.stageWorker(stage)
	
	cp.stages = append(cp.stages, stage)
	cp.caches = append(cp.caches, cache)
}

func (cp *CachedPipeline) stageWorker(stage *CachedStage) {
	for data := range stage.input {
		// Check cache first
		cacheKey := fmt.Sprintf("%s_%s", stage.name, data.Key)
		if cached, found := stage.cache.Get(cacheKey); found {
			stage.output <- cached
		} else {
			processed := stage.processFunc(data)
			stage.cache.Set(cacheKey, processed)
			stage.output <- processed
		}
	}
}

func (pc *PipelineCache) Get(key string) (ProcessedData, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	
	if value, found := pc.data[key]; found {
		atomic.AddInt64(&pc.hits, 1)
		return value, true
	}
	atomic.AddInt64(&pc.misses, 1)
	return ProcessedData{}, false
}

func (pc *PipelineCache) Set(key string, value ProcessedData) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.data[key] = value
}

func (pc *PipelineCache) GetHits() int64 {
	return atomic.LoadInt64(&pc.hits)
}

func (pc *PipelineCache) GetMisses() int64 {
	return atomic.LoadInt64(&pc.misses)
}

func (cp *CachedPipeline) Submit(data ProcessedData) {
	if len(cp.stages) > 0 {
		cp.stages[0].input <- data
	}
}

func (cp *CachedPipeline) Close() {
	for _, stage := range cp.stages {
		close(stage.input)
	}
}

// Advanced Pattern 5: Pipeline with Metrics and Monitoring
type MonitoredPipeline struct {
	stages   []*MonitoredStage
	metrics  *PipelineMetrics
	mu       sync.RWMutex
}

type MonitoredStage struct {
	name        string
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	metrics     *StageMetrics
}

type PipelineMetrics struct {
	totalItems     int64
	totalTime      time.Duration
	stageMetrics   map[string]*StageMetrics
	mu             sync.RWMutex
}

func NewMonitoredPipeline() *MonitoredPipeline {
	return &MonitoredPipeline{
		stages: make([]*MonitoredStage, 0),
		metrics: &PipelineMetrics{
			stageMetrics: make(map[string]*StageMetrics),
		},
	}
}

func (mp *MonitoredPipeline) AddStage(name string, processFunc func(ProcessedData) ProcessedData) {
	stageMetrics := &StageMetrics{}
	
	stage := &MonitoredStage{
		name:        name,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		metrics:     stageMetrics,
	}
	
	// Start stage worker
	go mp.stageWorker(stage)
	
	mp.stages = append(mp.stages, stage)
	mp.metrics.stageMetrics[name] = stageMetrics
}

func (mp *MonitoredPipeline) stageWorker(stage *MonitoredStage) {
	for data := range stage.input {
		start := time.Now()
		processed := stage.processFunc(data)
		duration := time.Since(start)
		
		// Update metrics
		mp.updateStageMetrics(stage.name, duration)
		
		stage.output <- processed
	}
}

func (mp *MonitoredPipeline) updateStageMetrics(stageName string, duration time.Duration) {
	mp.metrics.mu.Lock()
	defer mp.metrics.mu.Unlock()
	
	mp.metrics.totalItems++
	mp.metrics.totalTime += duration
	
	if stageMetrics, exists := mp.metrics.stageMetrics[stageName]; exists {
		stageMetrics.processedItems++
		stageMetrics.processingTime += duration
	}
}

func (mp *MonitoredPipeline) GetMetrics() *PipelineMetrics {
	mp.metrics.mu.RLock()
	defer mp.metrics.mu.RUnlock()
	
	// Return a copy of metrics
	return &PipelineMetrics{
		totalItems:   mp.metrics.totalItems,
		totalTime:    mp.metrics.totalTime,
		stageMetrics: mp.metrics.stageMetrics,
	}
}

func (mp *MonitoredPipeline) Submit(data ProcessedData) {
	if len(mp.stages) > 0 {
		mp.stages[0].input <- data
	}
}

func (mp *MonitoredPipeline) Close() {
	for _, stage := range mp.stages {
		close(stage.input)
	}
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Pipeline Patterns")
	fmt.Println("=============================")
	
	// Pattern 1: Adaptive Pipeline
	fmt.Println("\n1. Adaptive Pipeline:")
	adaptivePipeline := NewAdaptivePipeline()
	
	// Add stages
	adaptivePipeline.AddStage("stage1", 2, 5, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Adaptive Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	adaptivePipeline.AddStage("stage2", 2, 5, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Adaptive Stage2: %s", data.Value),
			Key:   data.Key,
			Stage: "stage2",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		adaptivePipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Adaptive Item %d", i),
			Key:   fmt.Sprintf("adaptive_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(2 * time.Second)
	adaptivePipeline.Close()
	
	// Pattern 2: Load Balanced Pipeline
	fmt.Println("\n2. Load Balanced Pipeline:")
	loadBalancedPipeline := NewLoadBalancedPipeline()
	
	loadBalancedPipeline.AddStage("stage1", 3, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Load Balanced Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		loadBalancedPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Load Balanced Item %d", i),
			Key:   fmt.Sprintf("load_balanced_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	loadBalancedPipeline.Close()
	
	// Pattern 3: Circuit Breaker Pipeline
	fmt.Println("\n3. Circuit Breaker Pipeline:")
	circuitBreakerPipeline := NewCircuitBreakerPipeline()
	
	circuitBreakerPipeline.AddStage("stage1", 3, 500*time.Millisecond, func(data ProcessedData) (ProcessedData, error) {
		time.Sleep(50 * time.Millisecond)
		if data.ID%4 == 0 {
			return ProcessedData{}, fmt.Errorf("stage1 failed for item %d", data.ID)
		}
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Circuit Breaker Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}, nil
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		circuitBreakerPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Circuit Breaker Item %d", i),
			Key:   fmt.Sprintf("circuit_breaker_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	circuitBreakerPipeline.Close()
	
	// Pattern 4: Cached Pipeline
	fmt.Println("\n4. Cached Pipeline:")
	cachedPipeline := NewCachedPipeline()
	
	cachedPipeline.AddStage("stage1", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Cached Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		cachedPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Cached Item %d", i),
			Key:   fmt.Sprintf("cached_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	cachedPipeline.Close()
	
	// Pattern 5: Monitored Pipeline
	fmt.Println("\n5. Monitored Pipeline:")
	monitoredPipeline := NewMonitoredPipeline()
	
	monitoredPipeline.AddStage("stage1", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Monitored Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	monitoredPipeline.AddStage("stage2", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Monitored Stage2: %s", data.Value),
			Key:   data.Key,
			Stage: "stage2",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		monitoredPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Monitored Item %d", i),
			Key:   fmt.Sprintf("monitored_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	
	// Print metrics
	metrics := monitoredPipeline.GetMetrics()
	fmt.Printf("  Total Items Processed: %d\n", metrics.totalItems)
	fmt.Printf("  Total Processing Time: %v\n", metrics.totalTime)
	
	monitoredPipeline.Close()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
