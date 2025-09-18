package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Select-based Event Loop
type EventLoop struct {
	events    chan Event
	commands  chan Command
	quit      chan bool
	handlers  map[string]func(Event)
	mu        sync.RWMutex
}

type Event struct {
	Type string
	Data interface{}
}

type Command struct {
	Type string
	Data interface{}
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		events:   make(chan Event, 100),
		commands: make(chan Command, 100),
		quit:     make(chan bool),
		handlers: make(map[string]func(Event)),
	}
}

func (el *EventLoop) RegisterHandler(eventType string, handler func(Event)) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.handlers[eventType] = handler
}

func (el *EventLoop) EmitEvent(event Event) {
	select {
	case el.events <- event:
	default:
		// Event channel is full, drop event
	}
}

func (el *EventLoop) SendCommand(cmd Command) {
	select {
	case el.commands <- cmd:
	default:
		// Command channel is full, drop command
	}
}

func (el *EventLoop) Start() {
	go el.run()
}

func (el *EventLoop) run() {
	for {
		select {
		case event := <-el.events:
			el.handleEvent(event)
		case cmd := <-el.commands:
			el.handleCommand(cmd)
		case <-el.quit:
			return
		}
	}
}

func (el *EventLoop) handleEvent(event Event) {
	el.mu.RLock()
	handler, exists := el.handlers[event.Type]
	el.mu.RUnlock()
	
	if exists {
		handler(event)
	}
}

func (el *EventLoop) handleCommand(cmd Command) {
	switch cmd.Type {
	case "quit":
		el.quit <- true
	case "register":
		if data, ok := cmd.Data.(map[string]interface{}); ok {
			if eventType, ok := data["eventType"].(string); ok {
				el.RegisterHandler(eventType, func(e Event) {
					fmt.Printf("Dynamic handler for %s: %v\n", e.Type, e.Data)
				})
			}
		}
	}
}

func (el *EventLoop) Stop() {
	el.quit <- true
}

// Advanced Pattern 2: Select-based Rate Limiter
type SelectRateLimiter struct {
	tokens   chan struct{}
	rate     time.Duration
	burst    int
	stopCh   chan bool
	mu       sync.RWMutex
	requests int64
	allowed  int64
}

func NewSelectRateLimiter(rate time.Duration, burst int) *SelectRateLimiter {
	rl := &SelectRateLimiter{
		tokens: make(chan struct{}, burst),
		rate:   rate,
		burst:  burst,
		stopCh: make(chan bool),
	}
	
	// Fill initial tokens
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}
	
	go rl.refill()
	return rl
}

func (rl *SelectRateLimiter) refill() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
				// Token added
			default:
				// Bucket is full
			}
		case <-rl.stopCh:
			return
		}
	}
}

func (rl *SelectRateLimiter) Allow() bool {
	atomic.AddInt64(&rl.requests, 1)
	
	select {
	case <-rl.tokens:
		atomic.AddInt64(&rl.allowed, 1)
		return true
	default:
		return false
	}
}

func (rl *SelectRateLimiter) Stats() (requests, allowed int64) {
	return atomic.LoadInt64(&rl.requests), atomic.LoadInt64(&rl.allowed)
}

func (rl *SelectRateLimiter) Stop() {
	close(rl.stopCh)
}

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

// Advanced Pattern 4: Select-based Circuit Breaker
type SelectCircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	closed      int32 // 0: open, 1: closed
	mu          sync.RWMutex
	stateCh     chan int32
	requestCh   chan func() error
	resultCh    chan error
}

func NewSelectCircuitBreaker(threshold int64, timeout time.Duration) *SelectCircuitBreaker {
	cb := &SelectCircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
		stateCh:   make(chan int32, 1),
		requestCh: make(chan func() error, 100),
		resultCh:  make(chan error, 100),
	}
	
	cb.stateCh <- 0 // Initial state
	go cb.run()
	return cb
}

func (cb *SelectCircuitBreaker) run() {
	for {
		select {
		case state := <-cb.stateCh:
			atomic.StoreInt32(&cb.state, state)
			fmt.Printf("Circuit breaker state changed to: %d\n", state)
		case fn := <-cb.requestCh:
			cb.handleRequest(fn)
		}
		
		// Check if closed
		if atomic.LoadInt32(&cb.closed) == 1 {
			return
		}
	}
}

func (cb *SelectCircuitBreaker) handleRequest(fn func() error) {
	state := atomic.LoadInt32(&cb.state)
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.stateCh <- 2 // half-open
			cb.executeRequest(fn)
		} else {
			cb.resultCh <- fmt.Errorf("circuit breaker is open")
		}
	case 2: // half-open
		cb.executeRequest(fn)
	default: // closed
		cb.executeRequest(fn)
	}
}

func (cb *SelectCircuitBreaker) executeRequest(fn func() error) {
	if fn == nil {
		select {
		case cb.resultCh <- fmt.Errorf("nil function"):
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
		return
	}
	
	err := fn()
	
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.threshold {
			select {
			case cb.stateCh <- 1: // open
			default:
			}
		}
		select {
		case cb.resultCh <- err:
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
	} else {
		// Success
		cb.failures = 0
		select {
		case cb.stateCh <- 0: // closed
		default:
		}
		select {
		case cb.resultCh <- nil:
		case <-time.After(100 * time.Millisecond):
			// Channel might be closed
		}
	}
}

func (cb *SelectCircuitBreaker) Call(fn func() error) error {
	if atomic.LoadInt32(&cb.closed) == 1 {
		return fmt.Errorf("circuit breaker is closed")
	}
	
	select {
	case cb.requestCh <- fn:
		return <-cb.resultCh
	case <-time.After(1 * time.Second):
		return fmt.Errorf("circuit breaker timeout")
	}
}

func (cb *SelectCircuitBreaker) Close() {
	atomic.StoreInt32(&cb.closed, 1)
	close(cb.stateCh)
	close(cb.requestCh)
	close(cb.resultCh)
}

// Advanced Pattern 5: Select-based Message Router
type SelectMessageRouter struct {
	routes   map[string][]chan interface{}
	mu       sync.RWMutex
	messageCh chan Message
	quitCh   chan bool
}

type Message struct {
	Topic string
	Data  interface{}
}

func NewSelectMessageRouter() *SelectMessageRouter {
	return &SelectMessageRouter{
		routes:    make(map[string][]chan interface{}),
		messageCh: make(chan Message, 100),
		quitCh:    make(chan bool),
	}
}

func (mr *SelectMessageRouter) Subscribe(topic string) <-chan interface{} {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	ch := make(chan interface{}, 10)
	mr.routes[topic] = append(mr.routes[topic], ch)
	return ch
}

func (mr *SelectMessageRouter) Publish(topic string, data interface{}) {
	select {
	case mr.messageCh <- Message{Topic: topic, Data: data}:
		// Message queued
	default:
		// Router is busy, drop message
	}
}

func (mr *SelectMessageRouter) Start() {
	go mr.run()
}

func (mr *SelectMessageRouter) run() {
	for {
		select {
		case msg := <-mr.messageCh:
			mr.routeMessage(msg)
		case <-mr.quitCh:
			return
		}
	}
}

func (mr *SelectMessageRouter) routeMessage(msg Message) {
	mr.mu.RLock()
	subscribers, exists := mr.routes[msg.Topic]
	mr.mu.RUnlock()
	
	if exists {
		for _, ch := range subscribers {
			select {
			case ch <- msg.Data:
				// Message sent
			default:
				// Subscriber is busy, skip
			}
		}
	}
}

func (mr *SelectMessageRouter) Stop() {
	close(mr.quitCh)
}

// Advanced Pattern 6: Select-based Worker Pool with Priority
type SelectPriorityPool struct {
	highJobs   chan Job
	normalJobs chan Job
	lowJobs    chan Job
	quitCh     chan bool
	wg         sync.WaitGroup
}

func NewSelectPriorityPool(workers int) *SelectPriorityPool {
	pool := &SelectPriorityPool{
		highJobs:   make(chan Job, 100),
		normalJobs: make(chan Job, 100),
		lowJobs:    make(chan Job, 100),
		quitCh:     make(chan bool),
	}
	
	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go pool.worker(i)
	}
	
	return pool
}

func (pool *SelectPriorityPool) worker(id int) {
	defer pool.wg.Done()
	
	for {
		select {
		case job := <-pool.highJobs:
			pool.processJob(id, job, "HIGH")
		case job := <-pool.normalJobs:
			pool.processJob(id, job, "NORMAL")
		case job := <-pool.lowJobs:
			pool.processJob(id, job, "LOW")
		case <-pool.quitCh:
			return
		}
	}
}

func (pool *SelectPriorityPool) processJob(workerID int, job Job, priority string) {
	fmt.Printf("Worker %d: Processing %s priority job %d\n", workerID, priority, job.ID)
	time.Sleep(100 * time.Millisecond) // Simulate work
}

func (pool *SelectPriorityPool) Submit(job Job, priority string) {
	switch priority {
	case "high":
		select {
		case pool.highJobs <- job:
		default:
			fmt.Printf("High priority queue full, dropping job %d\n", job.ID)
		}
	case "normal":
		select {
		case pool.normalJobs <- job:
		default:
			fmt.Printf("Normal priority queue full, dropping job %d\n", job.ID)
		}
	case "low":
		select {
		case pool.lowJobs <- job:
		default:
			fmt.Printf("Low priority queue full, dropping job %d\n", job.ID)
		}
	}
}

func (pool *SelectPriorityPool) Stop() {
	close(pool.quitCh)
	pool.wg.Wait()
}

// Advanced Pattern 7: Select-based Context Manager
type SelectContextManager struct {
	contexts map[string]context.Context
	cancels  map[string]context.CancelFunc
	mu       sync.RWMutex
	requestCh chan ContextRequest
	responseCh chan context.Context
}

type ContextRequest struct {
	ID   string
	Type string // "get", "create", "cancel"
}

func NewSelectContextManager() *SelectContextManager {
	cm := &SelectContextManager{
		contexts:   make(map[string]context.Context),
		cancels:    make(map[string]context.CancelFunc),
		requestCh:  make(chan ContextRequest, 100),
		responseCh: make(chan context.Context, 100),
	}
	
	go cm.run()
	return cm
}

func (cm *SelectContextManager) run() {
	for {
		select {
		case req := <-cm.requestCh:
			cm.handleRequest(req)
		}
	}
}

func (cm *SelectContextManager) handleRequest(req ContextRequest) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	switch req.Type {
	case "get":
		if ctx, exists := cm.contexts[req.ID]; exists {
			cm.responseCh <- ctx
		} else {
			cm.responseCh <- nil
		}
	case "create":
		ctx, cancel := context.WithCancel(context.Background())
		cm.contexts[req.ID] = ctx
		cm.cancels[req.ID] = cancel
		cm.responseCh <- ctx
	case "cancel":
		if cancel, exists := cm.cancels[req.ID]; exists {
			cancel()
			delete(cm.contexts, req.ID)
			delete(cm.cancels, req.ID)
		}
		cm.responseCh <- nil
	}
}

func (cm *SelectContextManager) GetContext(id string) context.Context {
	cm.requestCh <- ContextRequest{ID: id, Type: "get"}
	return <-cm.responseCh
}

func (cm *SelectContextManager) CreateContext(id string) context.Context {
	cm.requestCh <- ContextRequest{ID: id, Type: "create"}
	return <-cm.responseCh
}

func (cm *SelectContextManager) CancelContext(id string) {
	cm.requestCh <- ContextRequest{ID: id, Type: "cancel"}
	<-cm.responseCh
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Select Patterns")
	fmt.Println("===========================")
	
	// Pattern 1: Event Loop
	fmt.Println("\n1. Select-based Event Loop:")
	eventLoop := NewEventLoop()
	eventLoop.RegisterHandler("test", func(e Event) {
		fmt.Printf("Event received: %s - %v\n", e.Type, e.Data)
	})
	eventLoop.Start()
	
	eventLoop.EmitEvent(Event{Type: "test", Data: "Hello World"})
	eventLoop.SendCommand(Command{Type: "quit"})
	time.Sleep(100 * time.Millisecond)
	
	// Pattern 2: Rate Limiter
	fmt.Println("\n2. Select-based Rate Limiter:")
	rateLimiter := NewSelectRateLimiter(100*time.Millisecond, 3)
	
	for i := 0; i < 10; i++ {
		if rateLimiter.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i)
		}
	}
	
	requests, allowed := rateLimiter.Stats()
	fmt.Printf("Stats: %d requests, %d allowed\n", requests, allowed)
	rateLimiter.Stop()
	
	// Pattern 3: Load Balancer
	fmt.Println("\n3. Select-based Load Balancer:")
	loadBalancer := NewSelectLoadBalancer(3)
	
	for i := 0; i < 10; i++ {
		loadBalancer.Submit(Job{ID: i, Data: fmt.Sprintf("Data %d", i)})
	}
	
	time.Sleep(1 * time.Second)
	loadBalancer.Stop()
	
	// Pattern 4: Circuit Breaker
	fmt.Println("\n4. Select-based Circuit Breaker:")
	circuitBreaker := NewSelectCircuitBreaker(3, 1*time.Second)
	
	for i := 0; i < 5; i++ {
		err := circuitBreaker.Call(func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i)
		}
	}
	circuitBreaker.Close()
	
	// Pattern 5: Message Router
	fmt.Println("\n5. Select-based Message Router:")
	router := NewSelectMessageRouter()
	router.Start()
	
	sub1 := router.Subscribe("topic1")
	sub2 := router.Subscribe("topic1")
	
	go func() {
		for msg := range sub1 {
			fmt.Printf("Subscriber 1: %v\n", msg)
		}
	}()
	
	go func() {
		for msg := range sub2 {
			fmt.Printf("Subscriber 2: %v\n", msg)
		}
	}()
	
	router.Publish("topic1", "Message 1")
	router.Publish("topic1", "Message 2")
	
	time.Sleep(100 * time.Millisecond)
	router.Stop()
	
	// Pattern 6: Priority Pool
	fmt.Println("\n6. Select-based Priority Pool:")
	priorityPool := NewSelectPriorityPool(2)
	
	for i := 0; i < 5; i++ {
		priorities := []string{"high", "normal", "low"}
		priority := priorities[i%3]
		priorityPool.Submit(Job{ID: i, Data: fmt.Sprintf("Data %d", i)}, priority)
	}
	
	time.Sleep(1 * time.Second)
	priorityPool.Stop()
	
	// Pattern 7: Context Manager
	fmt.Println("\n7. Select-based Context Manager:")
	contextManager := NewSelectContextManager()
	
	ctx := contextManager.CreateContext("test")
	if ctx != nil {
		fmt.Println("Context created successfully")
	}
	
	// Test context cancellation
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Context still active")
	}
	
	contextManager.CancelContext("test")
	
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled after explicit cancel")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Context still active")
	}
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
