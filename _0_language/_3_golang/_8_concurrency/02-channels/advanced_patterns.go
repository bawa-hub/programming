package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Channel-based State Machine
type StateMachine struct {
	state    int32
	stateCh  chan int32
	actionCh chan func()
	quitCh   chan bool
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		state:    0,
		stateCh:  make(chan int32, 1),
		actionCh: make(chan func(), 10),
		quitCh:   make(chan bool),
	}
	
	sm.stateCh <- 0 // Initial state
	go sm.run()
	return sm
}

func (sm *StateMachine) run() {
	for {
		select {
		case newState := <-sm.stateCh:
			atomic.StoreInt32(&sm.state, newState)
			fmt.Printf("State changed to: %d\n", newState)
		case action := <-sm.actionCh:
			action()
		case <-sm.quitCh:
			return
		}
	}
}

func (sm *StateMachine) SetState(state int32) {
	sm.stateCh <- state
}

func (sm *StateMachine) GetState() int32 {
	return atomic.LoadInt32(&sm.state)
}

func (sm *StateMachine) DoAction(action func()) {
	sm.actionCh <- action
}

func (sm *StateMachine) Stop() {
	close(sm.quitCh)
}

// Advanced Pattern 2: Channel-based Rate Limiter
type RateLimiter struct {
	tokens   chan struct{}
	rate     time.Duration
	burst    int
	stopCh   chan bool
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	rl := &RateLimiter{
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

func (rl *RateLimiter) refill() {
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

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Stop() {
	close(rl.stopCh)
}

// Advanced Pattern 3: Channel-based Circuit Breaker
type CircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
	stateCh     chan int32
}

func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	cb := &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
		stateCh:   make(chan int32, 1),
	}
	
	cb.stateCh <- 0 // Initial state
	go cb.monitor()
	return cb
}

func (cb *CircuitBreaker) monitor() {
	for state := range cb.stateCh {
		atomic.StoreInt32(&cb.state, state)
		fmt.Printf("Circuit breaker state: %d\n", state)
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.RLock()
	state := atomic.LoadInt32(&cb.state)
	cb.mu.RUnlock()
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.stateCh <- 2 // half-open
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.threshold {
			cb.stateCh <- 1 // open
		}
		return err
	}
	
	// Success
	cb.failures = 0
	cb.stateCh <- 0 // closed
	return nil
}

func (cb *CircuitBreaker) Close() {
	close(cb.stateCh)
}

// Advanced Pattern 4: Channel-based Event Bus
type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(topic string) <-chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	ch := make(chan interface{}, 10)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	
	for _, ch := range eb.subscribers[topic] {
		select {
		case ch <- data:
		default:
			// Channel is full, skip
		}
	}
}

func (eb *EventBus) Unsubscribe(topic string, ch <-chan interface{}) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	subscribers := eb.subscribers[topic]
	for i, sub := range subscribers {
		if sub == ch {
			eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(sub)
			break
		}
	}
}

// Advanced Pattern 5: Channel-based Worker Pool with Priority
type PriorityWorkerPool struct {
	workers    int
	highJobs   chan Job
	normalJobs chan Job
	lowJobs    chan Job
	quitCh     chan bool
	wg         sync.WaitGroup
}

type Job struct {
	ID       int
	Priority int // 1: high, 2: normal, 3: low
	Data     interface{}
	Result   chan interface{}
}

func NewPriorityWorkerPool(workers int) *PriorityWorkerPool {
	return &PriorityWorkerPool{
		workers:    workers,
		highJobs:   make(chan Job, 100),
		normalJobs: make(chan Job, 100),
		lowJobs:    make(chan Job, 100),
		quitCh:     make(chan bool),
	}
}

func (pwp *PriorityWorkerPool) Start() {
	for i := 0; i < pwp.workers; i++ {
		pwp.wg.Add(1)
		go pwp.worker(i)
	}
}

func (pwp *PriorityWorkerPool) worker(id int) {
	defer pwp.wg.Done()
	
	for {
		select {
		case job := <-pwp.highJobs:
			pwp.processJob(id, job)
		case job := <-pwp.normalJobs:
			pwp.processJob(id, job)
		case job := <-pwp.lowJobs:
			pwp.processJob(id, job)
		case <-pwp.quitCh:
			return
		}
	}
}

func (pwp *PriorityWorkerPool) processJob(workerID int, job Job) {
	fmt.Printf("Worker %d: Processing job %d (priority %d)\n", workerID, job.ID, job.Priority)
	time.Sleep(100 * time.Millisecond) // Simulate work
	job.Result <- fmt.Sprintf("Result for job %d", job.ID)
}

func (pwp *PriorityWorkerPool) Submit(job Job) {
	switch job.Priority {
	case 1:
		pwp.highJobs <- job
	case 2:
		pwp.normalJobs <- job
	case 3:
		pwp.lowJobs <- job
	}
}

func (pwp *PriorityWorkerPool) Stop() {
	close(pwp.quitCh)
	pwp.wg.Wait()
}

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

// Advanced Pattern 7: Channel-based Context with Cancellation
type ChannelContext struct {
	done   chan struct{}
	cancel func()
	mu     sync.RWMutex
}

func NewChannelContext() *ChannelContext {
	ctx := &ChannelContext{
		done: make(chan struct{}),
	}
	
	ctx.cancel = func() {
		close(ctx.done)
	}
	
	return ctx
}

func (ctx *ChannelContext) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *ChannelContext) Cancel() {
	ctx.cancel()
}

func (ctx *ChannelContext) WithTimeout(timeout time.Duration) *ChannelContext {
	newCtx := &ChannelContext{
		done: make(chan struct{}),
	}
	
	newCtx.cancel = func() {
		close(newCtx.done)
	}
	
	go func() {
		select {
		case <-time.After(timeout):
			newCtx.cancel()
		case <-ctx.done:
			newCtx.cancel()
		}
	}()
	
	return newCtx
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Channel Patterns")
	fmt.Println("============================")
	
	// Pattern 1: State Machine
	fmt.Println("\n1. Channel-based State Machine:")
	sm := NewStateMachine()
	sm.DoAction(func() {
		fmt.Printf("Current state: %d\n", sm.GetState())
	})
	sm.SetState(1)
	sm.DoAction(func() {
		fmt.Printf("Current state: %d\n", sm.GetState())
	})
	sm.Stop()
	
	// Pattern 2: Rate Limiter
	fmt.Println("\n2. Channel-based Rate Limiter:")
	rl := NewRateLimiter(100*time.Millisecond, 3)
	for i := 0; i < 5; i++ {
		if rl.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i)
		}
	}
	rl.Stop()
	
	// Pattern 3: Circuit Breaker
	fmt.Println("\n3. Channel-based Circuit Breaker:")
	cb := NewCircuitBreaker(3, 1*time.Second)
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
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
	cb.Close()
	
	// Pattern 4: Event Bus
	fmt.Println("\n4. Channel-based Event Bus:")
	eb := NewEventBus()
	
	// Subscribe to events
	ch1 := eb.Subscribe("events")
	ch2 := eb.Subscribe("events")
	
	// Listen for events
	go func() {
		for event := range ch1 {
			fmt.Printf("Subscriber 1 received: %v\n", event)
		}
	}()
	
	go func() {
		for event := range ch2 {
			fmt.Printf("Subscriber 2 received: %v\n", event)
		}
	}()
	
	// Publish events
	eb.Publish("events", "Event 1")
	eb.Publish("events", "Event 2")
	
	time.Sleep(100 * time.Millisecond)
	
	// Pattern 5: Priority Worker Pool
	fmt.Println("\n5. Channel-based Priority Worker Pool:")
	pwp := NewPriorityWorkerPool(2)
	pwp.Start()
	
	// Submit jobs with different priorities
	for i := 0; i < 5; i++ {
		job := Job{
			ID:       i,
			Priority: (i % 3) + 1, // 1, 2, 3
			Data:     fmt.Sprintf("Data %d", i),
			Result:   make(chan interface{}, 1),
		}
		pwp.Submit(job)
	}
	
	time.Sleep(1 * time.Second)
	pwp.Stop()
	
	// Pattern 6: Load Balancer
	fmt.Println("\n6. Channel-based Load Balancer:")
	lb := NewLoadBalancer(3)
	
	for i := 0; i < 5; i++ {
		job := Job{
			ID:     i,
			Data:   fmt.Sprintf("Data %d", i),
			Result: make(chan interface{}, 1),
		}
		lb.Submit(job)
	}
	
	time.Sleep(1 * time.Second)
	lb.Stop()
	
	// Pattern 7: Channel-based Context
	fmt.Println("\n7. Channel-based Context:")
	ctx := NewChannelContext()
	
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
		}
	}()
	
	time.Sleep(1 * time.Second)
	ctx.Cancel()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
