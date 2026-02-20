package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Worker represents a worker in the pool
type Worker struct {
	ID          int
	Input       chan WorkItem
	Output      chan Result
	IsActive    bool
	StartTime   time.Time
	LastActivity time.Time
}

// AdaptivePool manages workers dynamically based on load
type AdaptivePool struct {
	minWorkers     int
	maxWorkers     int
	currentWorkers int
	loadThreshold  float64
	workers        []Worker
	mutex          sync.RWMutex
	metrics        *PoolMetrics
}

// PoolMetrics tracks pool performance
type PoolMetrics struct {
	TotalProcessed    int64
	TotalErrors       int64
	AverageLatency    time.Duration
	ThroughputPerSec  float64
	WorkerUtilization float64
	QueueLength       int64
	LastUpdate        time.Time
}

// CircuitBreaker implements circuit breaker pattern
type CircuitBreaker struct {
	failureCount    int64
	successCount    int64
	threshold       int64
	timeout         time.Duration
	state           CircuitState
	lastFailureTime time.Time
	mutex           sync.RWMutex
}

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// CacheLayer provides caching functionality
type CacheLayer struct {
	cache    map[string]interface{}
	mutex    sync.RWMutex
	hitCount int64
	missCount int64
	maxSize  int
}

// LoadBalancer distributes work across workers
type AdvancedLoadBalancer struct {
	workers      []Worker
	loadMetrics  map[int]float64
	roundRobin   int
	mutex        sync.RWMutex
	healthChecks map[int]bool
}

// MonitoringSystem provides real-time monitoring
type MonitoringSystem struct {
	metrics     *PoolMetrics
	alerts      []Alert
	subscribers []chan Alert
	mutex       sync.RWMutex
}

type Alert struct {
	Type      string
	Message   string
	Severity  string
	Timestamp time.Time
}

// Advanced Fan-Out/Fan-In with Adaptive Scaling
func adaptiveFanOutFanIn(input <-chan WorkItem, minWorkers, maxWorkers int) <-chan Result {
	output := make(chan Result, maxWorkers*2)
	
	pool := &AdaptivePool{
		minWorkers:     minWorkers,
		maxWorkers:     maxWorkers,
		currentWorkers: minWorkers,
		loadThreshold:  0.8,
		workers:        make([]Worker, 0, maxWorkers),
		metrics:        &PoolMetrics{},
	}
	
	// Start with minimum workers
	for i := 0; i < minWorkers; i++ {
		worker := Worker{
			ID:        i,
			Input:     make(chan WorkItem, 10),
			Output:    output,
			IsActive:  true,
			StartTime: time.Now(),
		}
		pool.workers = append(pool.workers, worker)
		go adaptiveWorker(worker, pool)
	}
	
	// Adaptive scaling goroutine
	go pool.adaptiveScaling()
	
	// Work distributor
	go pool.distributeWork(input)
	
	return output
}

// Adaptive worker that reports metrics
func adaptiveWorker(w Worker, pool *AdaptivePool) {
	defer close(w.Output)
	
	for item := range w.Input {
		start := time.Now()
		
		// Simulate work processing
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		// Simulate occasional errors
		var err error
		if rand.Float32() < 0.1 {
			err = fmt.Errorf("processing failed for item %d", item.ID)
			atomic.AddInt64(&pool.metrics.TotalErrors, 1)
		} else {
			atomic.AddInt64(&pool.metrics.TotalProcessed, 1)
		}
		
		// Send result
		result := Result{
			ID:        item.ID,
			Processed: fmt.Sprintf("Adaptive Worker %d processed: %s", w.ID, item.Data),
			WorkerID:  w.ID,
			Duration:  time.Since(start),
			Error:     err,
			Timestamp: time.Now(),
		}
		
		w.Output <- result
		w.LastActivity = time.Now()
		
		// Update metrics
		pool.updateMetrics(result)
	}
}

// Adaptive scaling based on load
func (p *AdaptivePool) adaptiveScaling() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		p.mutex.Lock()
		
		// Calculate current load
		load := p.calculateLoad()
		
		// Scale up if load is high
		if load > p.loadThreshold && p.currentWorkers < p.maxWorkers {
			p.scaleUp()
		}
		
		// Scale down if load is low
		if load < p.loadThreshold/2 && p.currentWorkers > p.minWorkers {
			p.scaleDown()
		}
		
		p.mutex.Unlock()
	}
}

// Calculate current load
func (p *AdaptivePool) calculateLoad() float64 {
	if len(p.workers) == 0 {
		return 0
	}
	
	var totalLoad float64
	for _, worker := range p.workers {
		if worker.IsActive {
			// Calculate load based on queue length and processing time
			queueLength := len(worker.Input)
			processingTime := time.Since(worker.LastActivity)
			load := float64(queueLength) + processingTime.Seconds()
			totalLoad += load
		}
	}
	
	return totalLoad / float64(len(p.workers))
}

// Scale up by adding workers
func (p *AdaptivePool) scaleUp() {
	if p.currentWorkers >= p.maxWorkers {
		return
	}
	
	worker := Worker{
		ID:        p.currentWorkers,
		Input:     make(chan WorkItem, 10),
		Output:    p.workers[0].Output, // Use same output channel
		IsActive:  true,
		StartTime: time.Now(),
	}
	
	p.workers = append(p.workers, worker)
	p.currentWorkers++
	
	go adaptiveWorker(worker, p)
	
	fmt.Printf("Scaled up to %d workers (load: %.2f)\n", p.currentWorkers, p.calculateLoad())
}

// Scale down by removing workers
func (p *AdaptivePool) scaleDown() {
	if p.currentWorkers <= p.minWorkers {
		return
	}
	
	// Find least active worker
	leastActiveIndex := 0
	leastActiveTime := p.workers[0].LastActivity
	
	for i, worker := range p.workers {
		if worker.LastActivity.Before(leastActiveTime) {
			leastActiveIndex = i
			leastActiveTime = worker.LastActivity
		}
	}
	
	// Remove worker
	close(p.workers[leastActiveIndex].Input)
	p.workers[leastActiveIndex].IsActive = false
	p.currentWorkers--
	
	// Remove from slice
	p.workers = append(p.workers[:leastActiveIndex], p.workers[leastActiveIndex+1:]...)
	
	fmt.Printf("Scaled down to %d workers (load: %.2f)\n", p.currentWorkers, p.calculateLoad())
}

// Distribute work to workers
func (p *AdaptivePool) distributeWork(input <-chan WorkItem) {
	defer func() {
		for _, worker := range p.workers {
			close(worker.Input)
		}
	}()
	
	workerIndex := 0
	for item := range input {
		p.mutex.RLock()
		if len(p.workers) > 0 {
			p.workers[workerIndex].Input <- item
			workerIndex = (workerIndex + 1) % len(p.workers)
		}
		p.mutex.RUnlock()
	}
}

// Update metrics
func (p *AdaptivePool) updateMetrics(result Result) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	
	p.metrics.LastUpdate = time.Now()
	
	// Update average latency
	if p.metrics.TotalProcessed > 0 {
		totalLatency := p.metrics.AverageLatency * time.Duration(p.metrics.TotalProcessed-1)
		p.metrics.AverageLatency = (totalLatency + result.Duration) / time.Duration(p.metrics.TotalProcessed)
	} else {
		p.metrics.AverageLatency = result.Duration
	}
	
	// Update throughput
	elapsed := time.Since(p.metrics.LastUpdate)
	if elapsed > 0 {
		p.metrics.ThroughputPerSec = float64(p.metrics.TotalProcessed) / elapsed.Seconds()
	}
	
	// Update worker utilization
	activeWorkers := 0
	for _, worker := range p.workers {
		if worker.IsActive {
			activeWorkers++
		}
	}
	p.metrics.WorkerUtilization = float64(activeWorkers) / float64(p.currentWorkers) * 100
}

// Circuit Breaker Implementation
func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     StateClosed,
	}
}

// Execute with circuit breaker
func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	// Check if circuit is open
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	// Execute function
	err := fn()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()
		
		if cb.failureCount >= cb.threshold {
			cb.state = StateOpen
		}
	} else {
		cb.successCount++
		if cb.state == StateHalfOpen {
			cb.state = StateClosed
			cb.failureCount = 0
		}
	}
	
	return err
}

// Get circuit breaker state
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// Cache Layer Implementation
func NewCacheLayer(maxSize int) *CacheLayer {
	return &CacheLayer{
		cache:   make(map[string]interface{}),
		maxSize: maxSize,
	}
}

// Get from cache
func (c *CacheLayer) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	value, exists := c.cache[key]
	if exists {
		atomic.AddInt64(&c.hitCount, 1)
	} else {
		atomic.AddInt64(&c.missCount, 1)
	}
	
	return value, exists
}

// Set in cache
func (c *CacheLayer) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	// Check if cache is full
	if len(c.cache) >= c.maxSize {
		// Remove oldest entry (simple implementation)
		for k := range c.cache {
			delete(c.cache, k)
			break
		}
	}
	
	c.cache[key] = value
}

// Get cache statistics
func (c *CacheLayer) GetStats() (int64, int64, float64) {
	hitCount := atomic.LoadInt64(&c.hitCount)
	missCount := atomic.LoadInt64(&c.missCount)
	
	total := hitCount + missCount
	hitRate := float64(0)
	if total > 0 {
		hitRate = float64(hitCount) / float64(total) * 100
	}
	
	return hitCount, missCount, hitRate
}

// Advanced Load Balancer
func NewAdvancedLoadBalancer(numWorkers int) *AdvancedLoadBalancer {
	workers := make([]Worker, numWorkers)
	loadMetrics := make(map[int]float64)
	healthChecks := make(map[int]bool)
	
	for i := 0; i < numWorkers; i++ {
		workers[i] = Worker{
			ID:        i,
			Input:     make(chan WorkItem, 10),
			Output:    make(chan Result, 10),
			IsActive:  true,
			StartTime: time.Now(),
		}
		loadMetrics[i] = 0
		healthChecks[i] = true
	}
	
	return &AdvancedLoadBalancer{
		workers:      workers,
		loadMetrics:  loadMetrics,
		healthChecks: healthChecks,
	}
}

// Get least loaded worker
func (lb *AdvancedLoadBalancer) GetLeastLoadedWorker() int {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	minLoad := lb.loadMetrics[0]
	minIndex := 0
	
	for i, load := range lb.loadMetrics {
		if lb.healthChecks[i] && load < minLoad {
			minLoad = load
			minIndex = i
		}
	}
	
	return minIndex
}

// Update worker load
func (lb *AdvancedLoadBalancer) UpdateLoad(workerIndex int, load float64) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	lb.loadMetrics[workerIndex] += load
}

// Health check
func (lb *AdvancedLoadBalancer) HealthCheck(workerIndex int) bool {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	return lb.healthChecks[workerIndex]
}

// Set worker health
func (lb *AdvancedLoadBalancer) SetHealth(workerIndex int, healthy bool) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	lb.healthChecks[workerIndex] = healthy
}

// Monitoring System
func NewMonitoringSystem() *MonitoringSystem {
	return &MonitoringSystem{
		metrics:     &PoolMetrics{},
		alerts:      make([]Alert, 0),
		subscribers: make([]chan Alert, 0),
	}
}

// Subscribe to alerts
func (ms *MonitoringSystem) Subscribe() <-chan Alert {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	
	alertChan := make(chan Alert, 10)
	ms.subscribers = append(ms.subscribers, alertChan)
	return alertChan
}

// Send alert
func (ms *MonitoringSystem) SendAlert(alertType, message, severity string) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	
	alert := Alert{
		Type:      alertType,
		Message:   message,
		Severity:  severity,
		Timestamp: time.Now(),
	}
	
	ms.alerts = append(ms.alerts, alert)
	
	// Send to subscribers
	for _, subscriber := range ms.subscribers {
		select {
		case subscriber <- alert:
		default:
			// Subscriber is full, skip
		}
	}
}

// Update metrics
func (ms *MonitoringSystem) UpdateMetrics(metrics *PoolMetrics) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	
	ms.metrics = metrics
	
	// Check for alerts
	if metrics.ErrorRate() > 10 {
		ms.SendAlert("HIGH_ERROR_RATE", "Error rate is above 10%", "WARNING")
	}
	
	if metrics.ThroughputPerSec < 1 {
		ms.SendAlert("LOW_THROUGHPUT", "Throughput is below 1 item/sec", "WARNING")
	}
	
	if metrics.WorkerUtilization < 50 {
		ms.SendAlert("LOW_UTILIZATION", "Worker utilization is below 50%", "INFO")
	}
}

// Get error rate
func (pm *PoolMetrics) ErrorRate() float64 {
	if pm.TotalProcessed == 0 {
		return 0
	}
	return float64(pm.TotalErrors) / float64(pm.TotalProcessed) * 100
}

// Advanced Fan-Out/Fan-In with Circuit Breaker
func circuitBreakerFanOutFanIn(input <-chan WorkItem, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers*2)
	cb := NewCircuitBreaker(5, 30*time.Second)
	
	// Create workers
	workers := make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = Worker{
			ID:        i,
			Input:     make(chan WorkItem, 10),
			Output:    output,
			IsActive:  true,
			StartTime: time.Now(),
		}
	}
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go circuitBreakerWorker(workers[i], cb)
	}
	
	// Fan-out with circuit breaker
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workers[i].Input)
			}
		}()
		
		workerIndex := 0
		for item := range input {
			workers[workerIndex].Input <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	return output
}

// Worker with circuit breaker
func circuitBreakerWorker(w Worker, cb *CircuitBreaker) {
	defer close(w.Output)
	
	for item := range w.Input {
		start := time.Now()
		
		// Execute with circuit breaker
		err := cb.Execute(func() error {
			// Simulate work processing
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			
			// Simulate errors
			if rand.Float32() < 0.2 {
				return fmt.Errorf("processing failed for item %d", item.ID)
			}
			
			return nil
		})
		
		// Send result
		result := Result{
			ID:        item.ID,
			Processed: fmt.Sprintf("Circuit Breaker Worker %d processed: %s", w.ID, item.Data),
			WorkerID:  w.ID,
			Duration:  time.Since(start),
			Error:     err,
			Timestamp: time.Now(),
		}
		
		w.Output <- result
		w.LastActivity = time.Now()
	}
}

// Advanced Fan-Out/Fan-In with Caching
func cachingFanOutFanIn(input <-chan WorkItem, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers*2)
	cache := NewCacheLayer(1000)
	
	// Create workers
	workers := make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = Worker{
			ID:        i,
			Input:     make(chan WorkItem, 10),
			Output:    output,
			IsActive:  true,
			StartTime: time.Now(),
		}
	}
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go cachingWorker(workers[i], cache)
	}
	
	// Fan-out
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workers[i].Input)
			}
		}()
		
		workerIndex := 0
		for item := range input {
			workers[workerIndex].Input <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	return output
}

// Worker with caching
func cachingWorker(w Worker, cache *CacheLayer) {
	defer close(w.Output)
	
	for item := range w.Input {
		start := time.Now()
		
		// Check cache first
		cacheKey := fmt.Sprintf("item_%d", item.ID)
		if cached, exists := cache.Get(cacheKey); exists {
			// Return cached result
			result := Result{
				ID:        item.ID,
				Processed: fmt.Sprintf("Cached result for: %s", cached),
				WorkerID:  w.ID,
				Duration:  time.Since(start),
				Error:     nil,
				Timestamp: time.Now(),
			}
			w.Output <- result
			continue
		}
		
		// Process item
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		// Simulate errors
		var err error
		if rand.Float32() < 0.1 {
			err = fmt.Errorf("processing failed for item %d", item.ID)
		}
		
		// Cache result if successful
		if err == nil {
			cache.Set(cacheKey, item.Data)
		}
		
		// Send result
		result := Result{
			ID:        item.ID,
			Processed: fmt.Sprintf("Caching Worker %d processed: %s", w.ID, item.Data),
			WorkerID:  w.ID,
			Duration:  time.Since(start),
			Error:     err,
			Timestamp: time.Now(),
		}
		
		w.Output <- result
		w.LastActivity = time.Now()
	}
}

// Advanced Fan-Out/Fan-In with Load Balancing
func loadBalancingFanOutFanIn(input <-chan WorkItem, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers*2)
	lb := NewAdvancedLoadBalancer(numWorkers)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go loadBalancingWorker(lb.workers[i], lb)
	}
	
	// Fan-out with load balancing
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(lb.workers[i].Input)
			}
		}()
		
		for item := range input {
			workerIndex := lb.GetLeastLoadedWorker()
			lb.workers[workerIndex].Input <- item
			lb.UpdateLoad(workerIndex, 1)
		}
	}()
	
	return output
}

// Worker with load balancing
func loadBalancingWorker(w Worker, lb *AdvancedLoadBalancer) {
	defer close(w.Output)
	
	for item := range w.Input {
		start := time.Now()
		
		// Simulate work processing
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		// Simulate errors
		var err error
		if rand.Float32() < 0.1 {
			err = fmt.Errorf("processing failed for item %d", item.ID)
			lb.SetHealth(w.ID, false)
		} else {
			lb.SetHealth(w.ID, true)
		}
		
		// Send result
		result := Result{
			ID:        item.ID,
			Processed: fmt.Sprintf("Load Balanced Worker %d processed: %s", w.ID, item.Data),
			WorkerID:  w.ID,
			Duration:  time.Since(start),
			Error:     err,
			Timestamp: time.Now(),
		}
		
		w.Output <- result
		w.LastActivity = time.Now()
		
		// Update load
		lb.UpdateLoad(w.ID, -1)
	}
}

// Advanced Fan-Out/Fan-In with Monitoring
func monitoringFanOutFanIn(input <-chan WorkItem, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers*2)
	monitoring := NewMonitoringSystem()
	
	// Create workers
	workers := make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = Worker{
			ID:        i,
			Input:     make(chan WorkItem, 10),
			Output:    output,
			IsActive:  true,
			StartTime: time.Now(),
		}
	}
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go monitoringWorker(workers[i], monitoring)
	}
	
	// Fan-out
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workers[i].Input)
			}
		}()
		
		workerIndex := 0
		for item := range input {
			workers[workerIndex].Input <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	// Start monitoring
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		
		for range ticker.C {
			// Update metrics
			metrics := &PoolMetrics{
				TotalProcessed:    int64(rand.Intn(100)),
				TotalErrors:       int64(rand.Intn(10)),
				AverageLatency:    time.Duration(rand.Intn(100)) * time.Millisecond,
				ThroughputPerSec:  float64(rand.Intn(50)),
				WorkerUtilization: float64(rand.Intn(100)),
				LastUpdate:        time.Now(),
			}
			
			monitoring.UpdateMetrics(metrics)
		}
	}()
	
	return output
}

// Worker with monitoring
func monitoringWorker(w Worker, monitoring *MonitoringSystem) {
	defer close(w.Output)
	
	for item := range w.Input {
		start := time.Now()
		
		// Simulate work processing
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		// Simulate errors
		var err error
		if rand.Float32() < 0.1 {
			err = fmt.Errorf("processing failed for item %d", item.ID)
		}
		
		// Send result
		result := Result{
			ID:        item.ID,
			Processed: fmt.Sprintf("Monitoring Worker %d processed: %s", w.ID, item.Data),
			WorkerID:  w.ID,
			Duration:  time.Since(start),
			Error:     err,
			Timestamp: time.Now(),
		}
		
		w.Output <- result
		w.LastActivity = time.Now()
	}
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Fan-Out/Fan-In Patterns")
	fmt.Println("===================================")
	
	// Example 1: Adaptive Scaling
	fmt.Println("\n1. Adaptive Scaling Fan-Out/Fan-In")
	fmt.Println("===================================")
	
	const numItems = 50
	const minWorkers = 2
	const maxWorkers = 8
	
	input := make(chan WorkItem, numItems)
	output := adaptiveFanOutFanIn(input, minWorkers, maxWorkers)
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Adaptive Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nAdaptive scaling completed: %d items processed\n", len(results))
	
	// Example 2: Circuit Breaker
	fmt.Println("\n2. Circuit Breaker Fan-Out/Fan-In")
	fmt.Println("==================================")
	
	input2 := make(chan WorkItem, 30)
	output2 := circuitBreakerFanOutFanIn(input2, 4)
	
	// Send work items
	go func() {
		defer close(input2)
		for i := 0; i < 30; i++ {
			input2 <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Circuit Breaker Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results2 []Result
	for result := range output2 {
		results2 = append(results2, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nCircuit breaker completed: %d items processed\n", len(results2))
	
	// Example 3: Caching
	fmt.Println("\n3. Caching Fan-Out/Fan-In")
	fmt.Println("==========================")
	
	input3 := make(chan WorkItem, 25)
	output3 := cachingFanOutFanIn(input3, 3)
	
	// Send work items
	go func() {
		defer close(input3)
		for i := 0; i < 25; i++ {
			input3 <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Caching Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results3 []Result
	for result := range output3 {
		results3 = append(results3, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nCaching completed: %d items processed\n", len(results3))
	
	// Example 4: Load Balancing
	fmt.Println("\n4. Load Balancing Fan-Out/Fan-In")
	fmt.Println("=================================")
	
	input4 := make(chan WorkItem, 35)
	output4 := loadBalancingFanOutFanIn(input4, 5)
	
	// Send work items
	go func() {
		defer close(input4)
		for i := 0; i < 35; i++ {
			input4 <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Load Balancing Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results4 []Result
	for result := range output4 {
		results4 = append(results4, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nLoad balancing completed: %d items processed\n", len(results4))
	
	// Example 5: Monitoring
	fmt.Println("\n5. Monitoring Fan-Out/Fan-In")
	fmt.Println("=============================")
	
	input5 := make(chan WorkItem, 20)
	output5 := monitoringFanOutFanIn(input5, 3)
	
	// Send work items
	go func() {
		defer close(input5)
		for i := 0; i < 20; i++ {
			input5 <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Monitoring Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results5 []Result
	for result := range output5 {
		results5 = append(results5, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nMonitoring completed: %d items processed\n", len(results5))
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
