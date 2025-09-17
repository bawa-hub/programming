package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Advanced Work Stealing Queue
type AdvancedWorkStealingQueue struct {
	tasks    []interface{}
	head     int64
	tail     int64
	capacity int64
	mutex    sync.Mutex
}

func NewAdvancedWorkStealingQueue(capacity int) *AdvancedWorkStealingQueue {
	return &AdvancedWorkStealingQueue{
		tasks:    make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (wsq *AdvancedWorkStealingQueue) Push(task interface{}) bool {
	wsq.mutex.Lock()
	defer wsq.mutex.Unlock()
	
	currentTail := wsq.tail
	nextTail := (currentTail + 1) % wsq.capacity
	
	if nextTail == wsq.head {
		return false // Queue full
	}
	
	wsq.tasks[currentTail] = task
	wsq.tail = nextTail
	return true
}

func (wsq *AdvancedWorkStealingQueue) Pop() (interface{}, bool) {
	wsq.mutex.Lock()
	defer wsq.mutex.Unlock()
	
	currentTail := wsq.tail
	currentHead := wsq.head
	
	if currentHead == currentTail {
		return nil, false // Queue empty
	}
	
	// Try to pop from tail
	newTail := (currentTail - 1 + wsq.capacity) % wsq.capacity
	if newTail != currentHead {
		wsq.tail = newTail
		task := wsq.tasks[newTail]
		return task, true
	}
	
	return nil, false
}

func (wsq *AdvancedWorkStealingQueue) Steal() (interface{}, bool) {
	wsq.mutex.Lock()
	defer wsq.mutex.Unlock()
	
	currentHead := wsq.head
	currentTail := wsq.tail
	
	if currentHead == currentTail {
		return nil, false // Queue empty
	}
	
	// Try to steal from head
	newHead := (currentHead + 1) % wsq.capacity
	if newHead != currentTail {
		wsq.head = newHead
		task := wsq.tasks[currentHead]
		return task, true
	}
	
	return nil, false
}

// Advanced Pattern 2: Scheduler-Aware Worker Pool
type SchedulerAwareWorkerPool struct {
	workers    int
	workQueue  chan interface{}
	resultChan chan interface{}
	done       chan bool
	wg         sync.WaitGroup
}

func NewSchedulerAwareWorkerPool(workers int, queueSize int) *SchedulerAwareWorkerPool {
	return &SchedulerAwareWorkerPool{
		workers:    workers,
		workQueue:  make(chan interface{}, queueSize),
		resultChan: make(chan interface{}, queueSize),
		done:       make(chan bool),
	}
}

func (p *SchedulerAwareWorkerPool) Start(processor func(interface{}) interface{}) {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func(workerID int) {
			defer p.wg.Done()
			
			for {
				select {
				case work := <-p.workQueue:
					result := processor(work)
					select {
					case p.resultChan <- result:
					case <-p.done:
						return
					}
				case <-p.done:
					return
				}
			}
		}(i)
	}
}

func (p *SchedulerAwareWorkerPool) Submit(work interface{}) {
	p.workQueue <- work
}

func (p *SchedulerAwareWorkerPool) GetResult() interface{} {
	return <-p.resultChan
}

func (p *SchedulerAwareWorkerPool) Stop() {
	close(p.done)
	p.wg.Wait()
	close(p.workQueue)
	close(p.resultChan)
}

// Advanced Pattern 3: CPU-Aware Load Balancer
type CPUAwareLoadBalancer struct {
	queues     []chan interface{}
	numQueues  int
	roundRobin int64
}

func NewCPUAwareLoadBalancer(numQueues int, queueSize int) *CPUAwareLoadBalancer {
	queues := make([]chan interface{}, numQueues)
	for i := range queues {
		queues[i] = make(chan interface{}, queueSize)
	}
	
	return &CPUAwareLoadBalancer{
		queues:    queues,
		numQueues: numQueues,
	}
}

func (lb *CPUAwareLoadBalancer) Submit(work interface{}) {
	// Round-robin distribution
	index := atomic.AddInt64(&lb.roundRobin, 1) % int64(lb.numQueues)
	lb.queues[index] <- work
}

func (lb *CPUAwareLoadBalancer) GetQueue(index int) <-chan interface{} {
	return lb.queues[index]
}

func (lb *CPUAwareLoadBalancer) Close() {
	for _, queue := range lb.queues {
		close(queue)
	}
}

// Advanced Pattern 4: Scheduler Statistics Monitor
type SchedulerStatsMonitor struct {
	goroutineCount    int64
	contextSwitches   int64
	workSteals        int64
	preemptions       int64
	lastUpdate        time.Time
	mutex             sync.RWMutex
}

func NewSchedulerStatsMonitor() *SchedulerStatsMonitor {
	return &SchedulerStatsMonitor{
		lastUpdate: time.Now(),
	}
}

func (sm *SchedulerStatsMonitor) UpdateStats() {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	sm.goroutineCount = int64(runtime.NumGoroutine())
	sm.lastUpdate = time.Now()
}

func (sm *SchedulerStatsMonitor) GetStats() (int64, int64, int64, int64, time.Time) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	return sm.goroutineCount, sm.contextSwitches, sm.workSteals, sm.preemptions, sm.lastUpdate
}

func (sm *SchedulerStatsMonitor) RecordContextSwitch() {
	atomic.AddInt64(&sm.contextSwitches, 1)
}

func (sm *SchedulerStatsMonitor) RecordWorkSteal() {
	atomic.AddInt64(&sm.workSteals, 1)
}

func (sm *SchedulerStatsMonitor) RecordPreemption() {
	atomic.AddInt64(&sm.preemptions, 1)
}

// Advanced Pattern 5: Adaptive Scheduler
type AdaptiveScheduler struct {
	workers        int
	workQueue      chan interface{}
	adaptiveWorkers int
	loadThreshold  int
	mutex          sync.RWMutex
}

func NewAdaptiveScheduler(initialWorkers int, queueSize int, loadThreshold int) *AdaptiveScheduler {
	return &AdaptiveScheduler{
		workers:        initialWorkers,
		workQueue:      make(chan interface{}, queueSize),
		adaptiveWorkers: initialWorkers,
		loadThreshold:  loadThreshold,
	}
}

func (as *AdaptiveScheduler) Start(processor func(interface{}) interface{}) {
	for i := 0; i < as.workers; i++ {
		go func(workerID int) {
			for work := range as.workQueue {
				processor(work)
				
				// Adaptive scaling based on queue length
				queueLen := len(as.workQueue)
				as.mutex.Lock()
				if queueLen > as.loadThreshold && as.adaptiveWorkers < as.workers*2 {
					as.adaptiveWorkers++
					go func(newWorkerID int) {
						for work := range as.workQueue {
							processor(work)
						}
					}(as.adaptiveWorkers)
				} else if queueLen < as.loadThreshold/2 && as.adaptiveWorkers > as.workers {
					as.adaptiveWorkers--
				}
				as.mutex.Unlock()
			}
		}(i)
	}
}

func (as *AdaptiveScheduler) Submit(work interface{}) {
	as.workQueue <- work
}

func (as *AdaptiveScheduler) Close() {
	close(as.workQueue)
}

// Advanced Pattern 6: Work Stealing with Priority
type PriorityWorkStealingQueue struct {
	highPriority chan interface{}
	lowPriority  chan interface{}
	capacity     int
}

func NewPriorityWorkStealingQueue(capacity int) *PriorityWorkStealingQueue {
	return &PriorityWorkStealingQueue{
		highPriority: make(chan interface{}, capacity),
		lowPriority:  make(chan interface{}, capacity),
		capacity:     capacity,
	}
}

func (pwsq *PriorityWorkStealingQueue) SubmitHighPriority(work interface{}) {
	select {
	case pwsq.highPriority <- work:
	default:
		// Fallback to low priority if high priority queue is full
		pwsq.lowPriority <- work
	}
}

func (pwsq *PriorityWorkStealingQueue) SubmitLowPriority(work interface{}) {
	pwsq.lowPriority <- work
}

func (pwsq *PriorityWorkStealingQueue) GetWork() (interface{}, bool) {
	// Try high priority first
	select {
	case work := <-pwsq.highPriority:
		return work, true
	default:
		// Try low priority
		select {
		case work := <-pwsq.lowPriority:
			return work, true
		default:
			return nil, false
		}
	}
}

func (pwsq *PriorityWorkStealingQueue) Close() {
	close(pwsq.highPriority)
	close(pwsq.lowPriority)
}

// Advanced Pattern 7: Scheduler-Aware Rate Limiter
type SchedulerAwareRateLimiter struct {
	rate     int
	interval time.Duration
	tokens   int64
	lastRefill time.Time
	mutex    sync.Mutex
}

func NewSchedulerAwareRateLimiter(rate int, interval time.Duration) *SchedulerAwareRateLimiter {
	return &SchedulerAwareRateLimiter{
		rate:      rate,
		interval:  interval,
		tokens:    int64(rate),
		lastRefill: time.Now(),
	}
}

func (rl *SchedulerAwareRateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	
	// Refill tokens based on elapsed time
	if elapsed >= rl.interval {
		rl.tokens = int64(rl.rate)
		rl.lastRefill = now
	}
	
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	
	return false
}

func (rl *SchedulerAwareRateLimiter) Wait() {
	for !rl.Allow() {
		runtime.Gosched() // Yield to scheduler
		time.Sleep(rl.interval / time.Duration(rl.rate))
	}
}

// Advanced Pattern 8: Scheduler-Aware Circuit Breaker
type SchedulerAwareCircuitBreaker struct {
	state         int32 // 0: closed, 1: open, 2: half-open
	failureCount  int64
	successCount  int64
	threshold     int64
	timeout       time.Duration
	lastFailure   time.Time
	mutex         sync.RWMutex
}

const (
	StateClosed = iota
	StateOpen
	StateHalfOpen
)

func NewSchedulerAwareCircuitBreaker(threshold int64, timeout time.Duration) *SchedulerAwareCircuitBreaker {
	return &SchedulerAwareCircuitBreaker{
		state:    StateClosed,
		threshold: threshold,
		timeout:  timeout,
	}
}

func (cb *SchedulerAwareCircuitBreaker) Execute(operation func() error) error {
	cb.mutex.RLock()
	state := cb.state
	cb.mutex.RUnlock()
	
	switch state {
	case StateOpen:
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.mutex.Lock()
			cb.state = StateHalfOpen
			cb.mutex.Unlock()
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	case StateHalfOpen:
		// Allow one request to test
	case StateClosed:
		// Normal operation
	}
	
	err := operation()
	
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		if cb.failureCount >= cb.threshold {
			cb.state = StateOpen
		}
		return err
	}
	
	cb.successCount++
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failureCount = 0
	}
	
	return nil
}

// Advanced Pattern 9: Scheduler-Aware Metrics Collector
type SchedulerAwareMetricsCollector struct {
	metrics map[string]int64
	mutex   sync.RWMutex
}

func NewSchedulerAwareMetricsCollector() *SchedulerAwareMetricsCollector {
	return &SchedulerAwareMetricsCollector{
		metrics: make(map[string]int64),
	}
}

func (mc *SchedulerAwareMetricsCollector) Increment(key string) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.metrics[key]++
}

func (mc *SchedulerAwareMetricsCollector) Add(key string, value int64) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.metrics[key] += value
}

func (mc *SchedulerAwareMetricsCollector) Get(key string) int64 {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	return mc.metrics[key]
}

func (mc *SchedulerAwareMetricsCollector) GetAll() map[string]int64 {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range mc.metrics {
		result[k] = v
	}
	return result
}

// Advanced Pattern 10: Scheduler-Aware Event Bus
type SchedulerAwareEventBus struct {
	subscribers map[string][]chan interface{}
	mutex       sync.RWMutex
}

func NewSchedulerAwareEventBus() *SchedulerAwareEventBus {
	return &SchedulerAwareEventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *SchedulerAwareEventBus) Subscribe(event string) <-chan interface{} {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()
	
	ch := make(chan interface{}, 10)
	eb.subscribers[event] = append(eb.subscribers[event], ch)
	return ch
}

func (eb *SchedulerAwareEventBus) Publish(event string, data interface{}) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()
	
	subscribers := eb.subscribers[event]
	for _, ch := range subscribers {
		select {
		case ch <- data:
		default:
			// Skip if channel is full
		}
	}
}

func (eb *SchedulerAwareEventBus) Close() {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()
	
	for _, subscribers := range eb.subscribers {
		for _, ch := range subscribers {
			close(ch)
		}
	}
}

// Advanced Pattern 11: Scheduler-Aware Web Server
type SchedulerAwareWebServer struct {
	workers    int
	requestChan chan interface{}
	responseChan chan interface{}
	done       chan bool
	wg         sync.WaitGroup
}

func NewSchedulerAwareWebServer(workers int, queueSize int) *SchedulerAwareWebServer {
	return &SchedulerAwareWebServer{
		workers:      workers,
		requestChan:  make(chan interface{}, queueSize),
		responseChan: make(chan interface{}, queueSize),
		done:         make(chan bool),
	}
}

func (ws *SchedulerAwareWebServer) Start(handler func(interface{}) interface{}) {
	for i := 0; i < ws.workers; i++ {
		ws.wg.Add(1)
		go func(workerID int) {
			defer ws.wg.Done()
			
			for {
				select {
				case request := <-ws.requestChan:
					response := handler(request)
					select {
					case ws.responseChan <- response:
					case <-ws.done:
						return
					}
				case <-ws.done:
					return
				}
			}
		}(i)
	}
}

func (ws *SchedulerAwareWebServer) HandleRequest(request interface{}) {
	ws.requestChan <- request
}

func (ws *SchedulerAwareWebServer) GetResponse() interface{} {
	return <-ws.responseChan
}

func (ws *SchedulerAwareWebServer) Stop() {
	close(ws.done)
	ws.wg.Wait()
	close(ws.requestChan)
	close(ws.responseChan)
}

// Advanced Pattern 12: Scheduler-Aware Message Queue
type SchedulerAwareMessageQueue struct {
	queues    map[string]chan interface{}
	mutex     sync.RWMutex
	capacity  int
}

func NewSchedulerAwareMessageQueue(capacity int) *SchedulerAwareMessageQueue {
	return &SchedulerAwareMessageQueue{
		queues:   make(map[string]chan interface{}),
		capacity: capacity,
	}
}

func (mq *SchedulerAwareMessageQueue) CreateQueue(name string) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	
	if _, exists := mq.queues[name]; !exists {
		mq.queues[name] = make(chan interface{}, mq.capacity)
	}
}

func (mq *SchedulerAwareMessageQueue) Publish(queueName string, message interface{}) error {
	mq.mutex.RLock()
	queue, exists := mq.queues[queueName]
	mq.mutex.RUnlock()
	
	if !exists {
		return fmt.Errorf("queue %s does not exist", queueName)
	}
	
	select {
	case queue <- message:
		return nil
	default:
		return fmt.Errorf("queue %s is full", queueName)
	}
}

func (mq *SchedulerAwareMessageQueue) Subscribe(queueName string) (<-chan interface{}, error) {
	mq.mutex.RLock()
	queue, exists := mq.queues[queueName]
	mq.mutex.RUnlock()
	
	if !exists {
		return nil, fmt.Errorf("queue %s does not exist", queueName)
	}
	
	return queue, nil
}

func (mq *SchedulerAwareMessageQueue) Close() {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	
	for _, queue := range mq.queues {
		close(queue)
	}
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Scheduling Patterns")
	fmt.Println("==============================")
	
	// Pattern 1: Advanced Work Stealing Queue
	fmt.Println("\n1. Advanced Work Stealing Queue")
	wsq := NewAdvancedWorkStealingQueue(10)
	for i := 0; i < 5; i++ {
		wsq.Push(i)
	}
	for i := 0; i < 5; i++ {
		work, ok := wsq.Pop()
		if ok {
			fmt.Printf("  Popped: %v\n", work)
		}
	}
	
	// Pattern 2: Scheduler-Aware Worker Pool
	fmt.Println("\n2. Scheduler-Aware Worker Pool")
	pool := NewSchedulerAwareWorkerPool(4, 100)
	pool.Start(func(work interface{}) interface{} {
		return work.(int) * work.(int)
	})
	
	for i := 0; i < 10; i++ {
		pool.Submit(i)
	}
	
	for i := 0; i < 10; i++ {
		result := pool.GetResult()
		fmt.Printf("  Result: %v\n", result)
	}
	
	pool.Stop()
	
	// Pattern 3: CPU-Aware Load Balancer
	fmt.Println("\n3. CPU-Aware Load Balancer")
	balancer := NewCPUAwareLoadBalancer(4, 100)
	for i := 0; i < 10; i++ {
		balancer.Submit(i)
	}
	balancer.Close()
	
	// Pattern 4: Scheduler Statistics Monitor
	fmt.Println("\n4. Scheduler Statistics Monitor")
	monitor := NewSchedulerStatsMonitor()
	monitor.UpdateStats()
	goroutines, _, _, _, _ := monitor.GetStats()
	fmt.Printf("  Goroutines: %d\n", goroutines)
	
	// Pattern 5: Adaptive Scheduler
	fmt.Println("\n5. Adaptive Scheduler")
	scheduler := NewAdaptiveScheduler(4, 100, 50)
	scheduler.Start(func(work interface{}) interface{} {
		return work.(int) * work.(int)
	})
	
	for i := 0; i < 10; i++ {
		scheduler.Submit(i)
	}
	scheduler.Close()
	
	// Pattern 6: Work Stealing with Priority
	fmt.Println("\n6. Work Stealing with Priority")
	pwsq := NewPriorityWorkStealingQueue(10)
	pwsq.SubmitHighPriority("high priority work")
	pwsq.SubmitLowPriority("low priority work")
	
	work, ok := pwsq.GetWork()
	if ok {
		fmt.Printf("  Got work: %v\n", work)
	}
	pwsq.Close()
	
	// Pattern 7: Scheduler-Aware Rate Limiter
	fmt.Println("\n7. Scheduler-Aware Rate Limiter")
	limiter := NewSchedulerAwareRateLimiter(10, time.Second)
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			fmt.Printf("  Request %d allowed\n", i+1)
		}
	}
	
	// Pattern 8: Scheduler-Aware Circuit Breaker
	fmt.Println("\n8. Scheduler-Aware Circuit Breaker")
	breaker := NewSchedulerAwareCircuitBreaker(3, time.Second)
	err := breaker.Execute(func() error {
		return nil
	})
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Println("  Operation successful")
	}
	
	// Pattern 9: Scheduler-Aware Metrics Collector
	fmt.Println("\n9. Scheduler-Aware Metrics Collector")
	collector := NewSchedulerAwareMetricsCollector()
	collector.Increment("requests")
	collector.Add("bytes", 1024)
	fmt.Printf("  Requests: %d\n", collector.Get("requests"))
	
	// Pattern 10: Scheduler-Aware Event Bus
	fmt.Println("\n10. Scheduler-Aware Event Bus")
	eventBus := NewSchedulerAwareEventBus()
	subscriber := eventBus.Subscribe("test")
	eventBus.Publish("test", "test data")
	event := <-subscriber
	fmt.Printf("  Event: %v\n", event)
	eventBus.Close()
	
	// Pattern 11: Scheduler-Aware Web Server
	fmt.Println("\n11. Scheduler-Aware Web Server")
	server := NewSchedulerAwareWebServer(4, 100)
	server.Start(func(request interface{}) interface{} {
		return "response to " + request.(string)
	})
	
	server.HandleRequest("test request")
	response := server.GetResponse()
	fmt.Printf("  Response: %v\n", response)
	server.Stop()
	
	// Pattern 12: Scheduler-Aware Message Queue
	fmt.Println("\n12. Scheduler-Aware Message Queue")
	messageQueue := NewSchedulerAwareMessageQueue(100)
	messageQueue.CreateQueue("test-queue")
	messageQueue.Publish("test-queue", "test message")
	
	messageSubscriber, err := messageQueue.Subscribe("test-queue")
	if err == nil {
		message := <-messageSubscriber
		fmt.Printf("  Message: %v\n", message)
	}
	messageQueue.Close()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
}
