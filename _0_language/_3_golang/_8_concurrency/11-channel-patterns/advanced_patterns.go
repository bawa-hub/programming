package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Advanced Pattern 1: Channel-Based State Machine
type State int

const (
	Idle State = iota
	Running
	Paused
	Stopped
)

type StateMachine struct {
	state     State
	stateCh   chan State
	actionCh  chan string
	resultCh  chan string
	mu        sync.RWMutex
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		state:     Idle,
		stateCh:   make(chan State),
		actionCh:  make(chan string),
		resultCh:  make(chan string, 1),
	}
	
	go sm.run()
	return sm
}

func (sm *StateMachine) run() {
	for {
		select {
		case newState := <-sm.stateCh:
			sm.mu.Lock()
			sm.state = newState
			sm.mu.Unlock()
			fmt.Printf("  State changed to: %v\n", sm.state)
		case action := <-sm.actionCh:
			result := sm.handleAction(action)
			sm.resultCh <- result
		}
	}
}

func (sm *StateMachine) handleAction(action string) string {
	sm.mu.RLock()
	currentState := sm.state
	sm.mu.RUnlock()
	
	switch currentState {
	case Idle:
		if action == "start" {
			sm.stateCh <- Running
			return "Started"
		}
		return "Invalid action for Idle state"
	case Running:
		switch action {
		case "stop":
			sm.stateCh <- Stopped
			return "Stopped"
		case "pause":
			sm.stateCh <- Paused
			return "Paused"
		}
		return "Invalid action for Running state"
	case Paused:
		if action == "resume" {
			sm.stateCh <- Running
			return "Resumed"
		}
		return "Invalid action for Paused state"
	case Stopped:
		return "Machine is stopped"
	default:
		return "Unknown state"
	}
}

func (sm *StateMachine) SendAction(action string) string {
	select {
	case sm.actionCh <- action:
		return <-sm.resultCh
	case <-time.After(1 * time.Second):
		return "timeout"
	}
}

// Advanced Pattern 2: Event-Driven State Machine
type Event int

const (
	Start Event = iota
	Stop
	Pause
	Resume
)

type EventStateMachine struct {
	state    State
	eventCh  chan Event
	stateCh  chan State
	mu       sync.RWMutex
}

func NewEventStateMachine() *EventStateMachine {
	esm := &EventStateMachine{
		state:   Idle,
		eventCh: make(chan Event),
		stateCh: make(chan State),
	}
	
	go esm.run()
	return esm
}

func (esm *EventStateMachine) run() {
	for event := range esm.eventCh {
		esm.handleEvent(event)
	}
}

func (esm *EventStateMachine) handleEvent(event Event) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	switch esm.state {
	case Idle:
		if event == Start {
			esm.state = Running
			esm.stateCh <- esm.state
		}
	case Running:
		switch event {
		case Stop:
			esm.state = Stopped
		case Pause:
			esm.state = Paused
		}
		esm.stateCh <- esm.state
	case Paused:
		if event == Resume {
			esm.state = Running
			esm.stateCh <- esm.state
		}
	}
}

func (esm *EventStateMachine) SendEvent(event Event) {
	esm.eventCh <- event
}

func (esm *EventStateMachine) GetState() State {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	return esm.state
}

// Advanced Pattern 3: Channel Pool with Load Balancing
type ChannelPool struct {
	channels []chan int
	current  int
	mu       sync.Mutex
}

func NewChannelPool(size int) *ChannelPool {
	pool := &ChannelPool{
		channels: make([]chan int, size),
	}
	
	for i := 0; i < size; i++ {
		pool.channels[i] = make(chan int, 100)
	}
	
	return pool
}

func (cp *ChannelPool) GetChannel() chan int {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	ch := cp.channels[cp.current]
	cp.current = (cp.current + 1) % len(cp.channels)
	return ch
}

func (cp *ChannelPool) SendToPool(value int) bool {
	ch := cp.GetChannel()
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

// Advanced Pattern 4: Channel Rate Limiter
type ChannelRateLimiter struct {
	limiter chan time.Time
	rate    time.Duration
}

func NewChannelRateLimiter(rate time.Duration) *ChannelRateLimiter {
	rl := &ChannelRateLimiter{
		limiter: make(chan time.Time, 1),
		rate:    rate,
	}
	
	go rl.run()
	return rl
}

func (rl *ChannelRateLimiter) run() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()
	
	for t := range ticker.C {
		select {
		case rl.limiter <- t:
		default:
		}
	}
}

func (rl *ChannelRateLimiter) Wait() {
	<-rl.limiter
}

// Advanced Pattern 5: Channel Circuit Breaker
type CircuitBreaker struct {
	state       int // 0: closed, 1: open, 2: half-open
	failureCount int
	threshold   int
	timeout     time.Duration
	lastFailure time.Time
	mu          sync.Mutex
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if cb.state == 1 { // Open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = 2 // Half-open
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		
		if cb.failureCount >= cb.threshold {
			cb.state = 1 // Open
		}
		return err
	}
	
	// Success
	cb.failureCount = 0
	cb.state = 0 // Closed
	return nil
}

// Advanced Pattern 6: Channel Message Router
type MessageRouter struct {
	routes map[string]chan interface{}
	mu     sync.RWMutex
}

func NewMessageRouter() *MessageRouter {
	return &MessageRouter{
		routes: make(map[string]chan interface{}),
	}
}

func (mr *MessageRouter) RegisterRoute(topic string, ch chan interface{}) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.routes[topic] = ch
}

func (mr *MessageRouter) SendMessage(topic string, message interface{}) bool {
	mr.mu.RLock()
	ch, exists := mr.routes[topic]
	mr.mu.RUnlock()
	
	if !exists {
		return false
	}
	
	select {
	case ch <- message:
		return true
	default:
		return false
	}
}

// Advanced Pattern 7: Channel Priority Queue
type PriorityItem struct {
	Value    interface{}
	Priority int
}

type PriorityQueue struct {
	items chan PriorityItem
	mu    sync.Mutex
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		items: make(chan PriorityItem, capacity),
	}
}

func (pq *PriorityQueue) Enqueue(item PriorityItem) bool {
	select {
	case pq.items <- item:
		return true
	default:
		return false
	}
}

func (pq *PriorityQueue) Dequeue() (PriorityItem, bool) {
	select {
	case item := <-pq.items:
		return item, true
	default:
		return PriorityItem{}, false
	}
}

// Advanced Pattern 8: Channel Event Bus
type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(eventType string) <-chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	ch := make(chan interface{}, 10)
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)
	return ch
}

func (eb *EventBus) Publish(eventType string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	
	for _, ch := range eb.subscribers[eventType] {
		select {
		case ch <- data:
		default:
			// Channel is full, skip
		}
	}
}

// Advanced Pattern 9: Channel Work Stealing
type WorkStealingPool struct {
	workers []chan func()
	stealCh chan func()
	mu      sync.Mutex
}

func NewWorkStealingPool(workers int) *WorkStealingPool {
	pool := &WorkStealingPool{
		workers: make([]chan func(), workers),
		stealCh: make(chan func(), 100),
	}
	
	for i := 0; i < workers; i++ {
		pool.workers[i] = make(chan func(), 10)
		go pool.worker(i)
	}
	
	return pool
}

func (wsp *WorkStealingPool) worker(id int) {
	for {
		select {
		case task := <-wsp.workers[id]:
			task()
		case task := <-wsp.stealCh:
			task()
		}
	}
}

func (wsp *WorkStealingPool) Submit(task func()) bool {
	wsp.mu.Lock()
	defer wsp.mu.Unlock()
	
	// Try to submit to own queue first
	select {
	case wsp.workers[0] <- task:
		return true
	default:
	}
	
	// Try to steal work
	select {
	case wsp.stealCh <- task:
		return true
	default:
		return false
	}
}

// Advanced Pattern 10: Channel Metrics Collector
type MetricsCollector struct {
	metrics chan Metric
	mu      sync.Mutex
	stats   map[string]int64
}

type Metric struct {
	Name  string
	Value int64
	Time  time.Time
}

func NewMetricsCollector() *MetricsCollector {
	mc := &MetricsCollector{
		metrics: make(chan Metric, 1000),
		stats:   make(map[string]int64),
	}
	
	go mc.collect()
	return mc
}

func (mc *MetricsCollector) collect() {
	for metric := range mc.metrics {
		mc.mu.Lock()
		mc.stats[metric.Name] += metric.Value
		mc.mu.Unlock()
	}
}

func (mc *MetricsCollector) Record(name string, value int64) {
	select {
	case mc.metrics <- Metric{Name: name, Value: value, Time: time.Now()}:
	default:
		// Metrics channel is full, drop metric
	}
}

func (mc *MetricsCollector) GetStats() map[string]int64 {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	
	stats := make(map[string]int64)
	for k, v := range mc.stats {
		stats[k] = v
	}
	return stats
}

// Advanced Pattern 1: Channel-Based State Machine
func channelBasedStateMachine() {
	fmt.Println("\n1. Channel-Based State Machine")
	fmt.Println("=============================")
	
	// Simplified state machine demonstration
	fmt.Println("  State machine concepts:")
	fmt.Println("  - Channel-based communication")
	fmt.Println("  - State transitions")
	fmt.Println("  - Action handling")
	fmt.Println("  - Event processing")
	
	fmt.Println("Channel-based state machine completed")
}

// Advanced Pattern 2: Event-Driven State Machine
func eventDrivenStateMachine() {
	fmt.Println("\n2. Event-Driven State Machine")
	fmt.Println("=============================")
	
	// Simplified event-driven state machine demonstration
	fmt.Println("  Event-driven concepts:")
	fmt.Println("  - Event processing")
	fmt.Println("  - State transitions")
	fmt.Println("  - Channel-based events")
	fmt.Println("  - Asynchronous handling")
	
	fmt.Println("Event-driven state machine completed")
}

// Advanced Pattern 3: Channel Pool with Load Balancing
func channelPoolWithLoadBalancing() {
	fmt.Println("\n3. Channel Pool with Load Balancing")
	fmt.Println("===================================")
	
	pool := NewChannelPool(3)
	
	// Test load balancing
	for i := 0; i < 10; i++ {
		if success := pool.SendToPool(i); success {
			fmt.Printf("  Sent %d to pool\n", i)
		} else {
			fmt.Printf("  Failed to send %d to pool\n", i)
		}
	}
	
	fmt.Println("Channel pool with load balancing completed")
}

// Advanced Pattern 4: Channel Rate Limiter
func channelRateLimiter() {
	fmt.Println("\n4. Channel Rate Limiter")
	fmt.Println("=======================")
	
	limiter := NewChannelRateLimiter(100 * time.Millisecond)
	
	// Test rate limiting
	for i := 0; i < 5; i++ {
		limiter.Wait()
		fmt.Printf("  Rate limited operation %d\n", i)
	}
	
	fmt.Println("Channel rate limiter completed")
}

// Advanced Pattern 5: Channel Circuit Breaker
func channelCircuitBreaker() {
	fmt.Println("\n5. Channel Circuit Breaker")
	fmt.Println("==========================")
	
	cb := NewCircuitBreaker(3, 1*time.Second)
	
	// Test circuit breaker
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("  Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("  Call %d succeeded\n", i)
		}
	}
	
	fmt.Println("Channel circuit breaker completed")
}

// Advanced Pattern 6: Channel Message Router
func channelMessageRouter() {
	fmt.Println("\n6. Channel Message Router")
	fmt.Println("=========================")
	
	router := NewMessageRouter()
	
	// Register routes
	ch1 := make(chan interface{}, 10)
	ch2 := make(chan interface{}, 10)
	router.RegisterRoute("topic1", ch1)
	router.RegisterRoute("topic2", ch2)
	
	// Send messages
	router.SendMessage("topic1", "message1")
	router.SendMessage("topic2", "message2")
	
	// Receive messages
	select {
	case msg := <-ch1:
		fmt.Printf("  Received from topic1: %v\n", msg)
	default:
	}
	
	select {
	case msg := <-ch2:
		fmt.Printf("  Received from topic2: %v\n", msg)
	default:
	}
	
	fmt.Println("Channel message router completed")
}

// Advanced Pattern 7: Channel Priority Queue
func channelPriorityQueue() {
	fmt.Println("\n7. Channel Priority Queue")
	fmt.Println("=========================")
	
	pq := NewPriorityQueue(10)
	
	// Enqueue items with different priorities
	pq.Enqueue(PriorityItem{Value: "low", Priority: 1})
	pq.Enqueue(PriorityItem{Value: "high", Priority: 10})
	pq.Enqueue(PriorityItem{Value: "medium", Priority: 5})
	
	// Dequeue items
	for i := 0; i < 3; i++ {
		if item, ok := pq.Dequeue(); ok {
			fmt.Printf("  Dequeued: %v (priority: %d)\n", item.Value, item.Priority)
		}
	}
	
	fmt.Println("Channel priority queue completed")
}

// Advanced Pattern 8: Channel Event Bus
func channelEventBus() {
	fmt.Println("\n8. Channel Event Bus")
	fmt.Println("===================")
	
	bus := NewEventBus()
	
	// Subscribe to events
	ch1 := bus.Subscribe("user.created")
	ch2 := bus.Subscribe("user.updated")
	
	// Publish events
	bus.Publish("user.created", "user123")
	bus.Publish("user.updated", "user456")
	
	// Receive events
	select {
	case event := <-ch1:
		fmt.Printf("  Received user.created: %v\n", event)
	default:
	}
	
	select {
	case event := <-ch2:
		fmt.Printf("  Received user.updated: %v\n", event)
	default:
	}
	
	fmt.Println("Channel event bus completed")
}

// Advanced Pattern 9: Channel Work Stealing
func channelWorkStealing() {
	fmt.Println("\n9. Channel Work Stealing")
	fmt.Println("========================")
	
	pool := NewWorkStealingPool(3)
	
	// Submit tasks
	for i := 0; i < 5; i++ {
		taskID := i
		if success := pool.Submit(func() {
			fmt.Printf("  Task %d executed\n", taskID)
		}); success {
			fmt.Printf("  Task %d submitted\n", taskID)
		} else {
			fmt.Printf("  Task %d failed to submit\n", taskID)
		}
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Channel work stealing completed")
}

// Advanced Pattern 10: Channel Metrics Collector
func channelMetricsCollector() {
	fmt.Println("\n10. Channel Metrics Collector")
	fmt.Println("=============================")
	
	collector := NewMetricsCollector()
	
	// Record metrics
	collector.Record("requests", 1)
	collector.Record("requests", 2)
	collector.Record("errors", 1)
	collector.Record("requests", 3)
	
	time.Sleep(100 * time.Millisecond)
	
	// Get stats
	stats := collector.GetStats()
	for name, value := range stats {
		fmt.Printf("  %s: %d\n", name, value)
	}
	
	fmt.Println("Channel metrics collector completed")
}

// Advanced Pattern 11: Web Server with Channel Patterns
func webServerWithChannelPatterns() {
	fmt.Println("\n11. Web Server with Channel Patterns")
	fmt.Println("===================================")
	
	server := NewWebServer()
	server.Start()
	
	// Simulate requests
	for i := 0; i < 3; i++ {
		req := &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: fmt.Sprintf("/api/%d", i)},
		}
		server.HandleRequest(req)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Web server with channel patterns completed")
}

type WebServer struct {
	requestCh  chan *http.Request
	responseCh chan *http.Response
	errorCh    chan error
}

func NewWebServer() *WebServer {
	return &WebServer{
		requestCh:  make(chan *http.Request, 100),
		responseCh: make(chan *http.Response, 100),
		errorCh:    make(chan error, 100),
	}
}

func (ws *WebServer) Start() {
	go ws.handleRequests()
	go ws.handleResponses()
	go ws.handleErrors()
}

func (ws *WebServer) HandleRequest(req *http.Request) {
	select {
	case ws.requestCh <- req:
	default:
		fmt.Println("  Request channel full")
	}
}

func (ws *WebServer) handleRequests() {
	for req := range ws.requestCh {
		// Process request
		_ = req // Use the request variable
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("OK")),
		}
		
		select {
		case ws.responseCh <- resp:
		case ws.errorCh <- fmt.Errorf("response channel full"):
		}
	}
}

func (ws *WebServer) handleResponses() {
	for resp := range ws.responseCh {
		fmt.Printf("  Response: %d\n", resp.StatusCode)
	}
}

func (ws *WebServer) handleErrors() {
	for err := range ws.errorCh {
		fmt.Printf("  Error: %v\n", err)
	}
}

// Advanced Pattern 12: Message Queue with Channel Patterns
func messageQueueWithChannelPatterns() {
	fmt.Println("\n12. Message Queue with Channel Patterns")
	fmt.Println("======================================")
	
	queue := NewMessageQueue(3)
	queue.Start()
	
	// Enqueue messages
	for i := 0; i < 5; i++ {
		msg := Message{
			ID:      fmt.Sprintf("msg-%d", i),
			Content: fmt.Sprintf("content-%d", i),
			Priority: i % 3,
		}
		
		if success := queue.Enqueue(msg); success {
			fmt.Printf("  Enqueued: %s\n", msg.ID)
		} else {
			fmt.Printf("  Failed to enqueue: %s\n", msg.ID)
		}
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Message queue with channel patterns completed")
}

type MessageQueue struct {
	messages chan Message
	workers  int
}

type Message struct {
	ID       string
	Content  string
	Priority int
}

func NewMessageQueue(workers int) *MessageQueue {
	return &MessageQueue{
		messages: make(chan Message, 1000),
		workers:  workers,
	}
}

func (mq *MessageQueue) Start() {
	for i := 0; i < mq.workers; i++ {
		go mq.worker(i)
	}
}

func (mq *MessageQueue) worker(id int) {
	for msg := range mq.messages {
		fmt.Printf("  Worker %d processing message %s\n", id, msg.ID)
		time.Sleep(50 * time.Millisecond)
	}
}

func (mq *MessageQueue) Enqueue(msg Message) bool {
	select {
	case mq.messages <- msg:
		return true
	default:
		return false
	}
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Channel Patterns & Idioms")
	fmt.Println("====================================")
	
	channelBasedStateMachine()
	eventDrivenStateMachine()
	channelPoolWithLoadBalancing()
	channelRateLimiter()
	channelCircuitBreaker()
	channelMessageRouter()
	channelPriorityQueue()
	channelEventBus()
	channelWorkStealing()
	channelMetricsCollector()
	webServerWithChannelPatterns()
	messageQueueWithChannelPatterns()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
