package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// 🔄 ADVANCED CHANNEL PATTERNS
// Understanding advanced channel patterns and state machines

func main() {
	fmt.Println("🔄 ADVANCED CHANNEL PATTERNS")
	fmt.Println("============================")

	// 1. Channel Multiplexing
	fmt.Println("\n1. Channel Multiplexing:")
	channelMultiplexing()

	// 2. Channel Demultiplexing
	fmt.Println("\n2. Channel Demultiplexing:")
	channelDemultiplexing()

	// 3. Channel-Based State Machines
	fmt.Println("\n3. Channel-Based State Machines:")
	channelBasedStateMachines()

	// 4. Fan-In Pattern
	fmt.Println("\n4. Fan-In Pattern:")
	fanInPattern()

	// 5. Fan-Out Pattern
	fmt.Println("\n5. Fan-Out Pattern:")
	fanOutPattern()

	// 6. Pipeline Pattern
	fmt.Println("\n6. Pipeline Pattern:")
	pipelinePattern()

	// 7. Worker Pool with Channels
	fmt.Println("\n7. Worker Pool with Channels:")
	workerPoolWithChannels()

	// 8. Rate Limiting with Channels
	fmt.Println("\n8. Rate Limiting with Channels:")
	rateLimitingWithChannels()

	// 9. Circuit Breaker Pattern
	fmt.Println("\n9. Circuit Breaker Pattern:")
	circuitBreakerPattern()

	// 10. Advanced Select Patterns
	fmt.Println("\n10. Advanced Select Patterns:")
	advancedSelectPatterns()
}

// CHANNEL MULTIPLEXING: Understanding channel multiplexing
func channelMultiplexing() {
	fmt.Println("Understanding channel multiplexing...")
	
	// Create multiple input channels
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	ch3 := make(chan string, 2)
	
	// Send data to channels
	go func() {
		ch1 <- "from channel 1"
		ch1 <- "from channel 1 again"
		close(ch1)
	}()
	
	go func() {
		ch2 <- "from channel 2"
		ch2 <- "from channel 2 again"
		close(ch2)
	}()
	
	go func() {
		ch3 <- "from channel 3"
		ch3 <- "from channel 3 again"
		close(ch3)
	}()
	
	// Multiplex channels
	multiplexed := multiplexChannels(ch1, ch2, ch3)
	
	// Read from multiplexed channel
	for msg := range multiplexed {
		fmt.Printf("  📊 Received: %s\n", msg)
	}
}

// CHANNEL DEMULTIPLEXING: Understanding channel demultiplexing
func channelDemultiplexing() {
	fmt.Println("Understanding channel demultiplexing...")
	
	// Create input channel
	input := make(chan int, 10)
	
	// Send data
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Demultiplex to multiple channels
	evenCh, oddCh := demultiplexChannels(input)
	
	// Read from demultiplexed channels
	go func() {
		for num := range evenCh {
			fmt.Printf("  📊 Even: %d\n", num)
		}
	}()
	
	go func() {
		for num := range oddCh {
			fmt.Printf("  📊 Odd: %d\n", num)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// CHANNEL-BASED STATE MACHINES: Understanding state machines with channels
func channelBasedStateMachines() {
	fmt.Println("Understanding channel-based state machines...")
	
	// Create state machine
	sm := NewStateMachine()
	
	// Send events
	events := []Event{
		{Type: "start"},
		{Type: "process"},
		{Type: "complete"},
		{Type: "error"},
		{Type: "retry"},
		{Type: "complete"},
	}
	
	for _, event := range events {
		sm.SendEvent(event)
		time.Sleep(100 * time.Millisecond)
	}
	
	sm.Close()
}

// FAN-IN PATTERN: Understanding fan-in pattern
func fanInPattern() {
	fmt.Println("Understanding fan-in pattern...")
	
	// Create multiple input channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	// Send data to channels
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- fmt.Sprintf("Channel 1: %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch2)
		for i := 0; i < 3; i++ {
			ch2 <- fmt.Sprintf("Channel 2: %d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch3)
		for i := 0; i < 3; i++ {
			ch3 <- fmt.Sprintf("Channel 3: %d", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Fan-in channels
	fanInCh := fanIn(ch1, ch2, ch3)
	
	// Read from fan-in channel
	for msg := range fanInCh {
		fmt.Printf("  📊 Fan-in: %s\n", msg)
	}
}

// FAN-OUT PATTERN: Understanding fan-out pattern
func fanOutPattern() {
	fmt.Println("Understanding fan-out pattern...")
	
	// Create input channel
	input := make(chan int, 10)
	
	// Send data
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()
	
	// Fan-out to multiple workers
	workers := 3
	outputs := fanOut(input, workers)
	
	// Read from all outputs
	var wg sync.WaitGroup
	for i, output := range outputs {
		wg.Add(1)
		go func(id int, ch <-chan int) {
			defer wg.Done()
			for num := range ch {
				fmt.Printf("  📊 Worker %d: %d\n", id, num)
			}
		}(i, output)
	}
	
	wg.Wait()
}

// PIPELINE PATTERN: Understanding pipeline pattern
func pipelinePattern() {
	fmt.Println("Understanding pipeline pattern...")
	
	// Create pipeline stages
	stage1 := make(chan int, 10)
	stage2 := make(chan int, 10)
	stage3 := make(chan int, 10)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers
	go func() {
		defer close(stage2)
		for num := range stage1 {
			stage2 <- num * num
		}
	}()
	
	// Stage 3: Add 100
	go func() {
		defer close(stage3)
		for num := range stage2 {
			stage3 <- num + 100
		}
	}()
	
	// Read final results
	for result := range stage3 {
		fmt.Printf("  📊 Pipeline result: %d\n", result)
	}
}

// WORKER POOL WITH CHANNELS: Understanding worker pool with channels
func workerPoolWithChannels() {
	fmt.Println("Understanding worker pool with channels...")
	
	// Create worker pool
	pool := NewChannelWorkerPool(3)
	
	// Add jobs
	for i := 0; i < 10; i++ {
		pool.AddJob(Job{
			ID:   i,
			Data: fmt.Sprintf("job-%d", i),
		})
	}
	
	// Start workers
	pool.Start()
	
	// Wait for completion
	pool.Wait()
	
	fmt.Printf("  📊 All jobs completed! Processed: %d\n", pool.ProcessedCount())
}

// RATE LIMITING WITH CHANNELS: Understanding rate limiting with channels
func rateLimitingWithChannels() {
	fmt.Println("Understanding rate limiting with channels...")
	
	// Create rate limiter
	limiter := NewRateLimiter(2, time.Second) // 2 requests per second
	
	// Create workers that need rate limiting
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				limiter.Wait()
				fmt.Printf("  📊 Worker %d: Request %d (rate limited)\n", id, j)
			}
		}(i)
	}
	
	time.Sleep(3 * time.Second)
}

// CIRCUIT BREAKER PATTERN: Understanding circuit breaker pattern
func circuitBreakerPattern() {
	fmt.Println("Understanding circuit breaker pattern...")
	
	// Create circuit breaker
	breaker := NewCircuitBreaker(3, time.Second)
	
	// Test circuit breaker
	for i := 0; i < 10; i++ {
		err := breaker.Call(func() error {
			// Simulate API call
			if i < 5 {
				return fmt.Errorf("API error")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("  📊 Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("  📊 Call %d succeeded\n", i)
		}
		
		time.Sleep(200 * time.Millisecond)
	}
}

// ADVANCED SELECT PATTERNS: Understanding advanced select patterns
func advancedSelectPatterns() {
	fmt.Println("Understanding advanced select patterns...")
	
	// Pattern 1: Priority select
	fmt.Println("  📊 Pattern 1: Priority select")
	prioritySelect()
	
	// Pattern 2: Timeout select
	fmt.Println("  📊 Pattern 2: Timeout select")
	timeoutSelect()
	
	// Pattern 3: Non-blocking select
	fmt.Println("  📊 Pattern 3: Non-blocking select")
	nonBlockingSelect()
	
	// Pattern 4: Dynamic select
	fmt.Println("  📊 Pattern 4: Dynamic select")
	dynamicSelect()
}

func prioritySelect() {
	highPriority := make(chan string, 1)
	lowPriority := make(chan string, 1)
	
	// Send to low priority first
	go func() {
		lowPriority <- "low priority message"
	}()
	
	// Send to high priority after delay
	go func() {
		time.Sleep(50 * time.Millisecond)
		highPriority <- "high priority message"
	}()
	
	// Select with priority
	select {
	case msg := <-highPriority:
		fmt.Printf("    📊 High priority: %s\n", msg)
	case msg := <-lowPriority:
		fmt.Printf("    📊 Low priority: %s\n", msg)
	}
}

func timeoutSelect() {
	ch := make(chan string)
	
	// Send message after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "message"
	}()
	
	// Select with timeout
	select {
	case msg := <-ch:
		fmt.Printf("    📊 Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("    📊 Timeout!\n")
	}
}

func nonBlockingSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Non-blocking select
	select {
	case msg := <-ch1:
		fmt.Printf("    📊 Ch1: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("    📊 Ch2: %s\n", msg)
	default:
		fmt.Printf("    📊 No message available\n")
	}
}

func dynamicSelect() {
	channels := []chan string{
		make(chan string),
		make(chan string),
		make(chan string),
	}
	
	// Send to random channel
	go func() {
		for i := 0; i < 3; i++ {
			ch := channels[rand.Intn(len(channels))]
			ch <- fmt.Sprintf("message %d", i)
		}
	}()
	
	// Dynamic select
	cases := make([]reflect.SelectCase, len(channels))
	for i, ch := range channels {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}
	
	for i := 0; i < 3; i++ {
		chosen, value, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("    📊 Channel %d: %s\n", chosen, value.String())
		}
	}
}

// IMPLEMENTATIONS

// Job type for worker pools
type Job struct {
	ID   int
	Data string
}

// Channel multiplexing
func multiplexChannels(chs ...<-chan string) <-chan string {
	out := make(chan string)
	
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		
		for _, ch := range chs {
			wg.Add(1)
			go func(c <-chan string) {
				defer wg.Done()
				for msg := range c {
					out <- msg
				}
			}(ch)
		}
		
		wg.Wait()
	}()
	
	return out
}

// Channel demultiplexing
func demultiplexChannels(input <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		
		for num := range input {
			if num%2 == 0 {
				evenCh <- num
			} else {
				oddCh <- num
			}
		}
	}()
	
	return evenCh, oddCh
}

// State Machine
type StateMachine struct {
	state     string
	events    chan Event
	responses chan string
	done      chan bool
}

type Event struct {
	Type string
	Data interface{}
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		state:     "idle",
		events:    make(chan Event),
		responses: make(chan string),
		done:      make(chan bool),
	}
	
	go sm.run()
	return sm
}

func (sm *StateMachine) run() {
	for {
		select {
		case event := <-sm.events:
			sm.handleEvent(event)
		case <-sm.done:
			return
		}
	}
}

func (sm *StateMachine) handleEvent(event Event) {
	oldState := sm.state
	
	switch sm.state {
	case "idle":
		if event.Type == "start" {
			sm.state = "running"
		}
	case "running":
		if event.Type == "process" {
			sm.state = "processing"
		} else if event.Type == "error" {
			sm.state = "error"
		}
	case "processing":
		if event.Type == "complete" {
			sm.state = "idle"
		} else if event.Type == "error" {
			sm.state = "error"
		}
	case "error":
		if event.Type == "retry" {
			sm.state = "running"
		}
	}
	
	fmt.Printf("  📊 State: %s -> %s (event: %s)\n", oldState, sm.state, event.Type)
}

func (sm *StateMachine) SendEvent(event Event) {
	sm.events <- event
}

func (sm *StateMachine) Close() {
	close(sm.done)
}

// Fan-in
func fanIn(chs ...<-chan string) <-chan string {
	out := make(chan string)
	
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		
		for _, ch := range chs {
			wg.Add(1)
			go func(c <-chan string) {
				defer wg.Done()
				for msg := range c {
					out <- msg
				}
			}(ch)
		}
		
		wg.Wait()
	}()
	
	return out
}

// Fan-out
func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	
	for i := 0; i < workers; i++ {
		output := make(chan int)
		outputs[i] = output
		
		go func(out chan<- int) {
			defer close(out)
			for num := range input {
				out <- num
			}
		}(output)
	}
	
	return outputs
}

// Channel Worker Pool
type ChannelWorkerPool struct {
	workers   int
	jobs      chan Job
	results   chan Job
	done      chan bool
	mu        sync.Mutex
	processed int
	completed bool
}

func NewChannelWorkerPool(workers int) *ChannelWorkerPool {
	return &ChannelWorkerPool{
		workers: workers,
		jobs:    make(chan Job, 100),
		results: make(chan Job, 100),
		done:    make(chan bool, workers),
	}
}

func (p *ChannelWorkerPool) AddJob(job Job) {
	p.jobs <- job
}

func (p *ChannelWorkerPool) Start() {
	for i := 0; i < p.workers; i++ {
		go p.worker(i)
	}
}

func (p *ChannelWorkerPool) worker(id int) {
	for job := range p.jobs {
		fmt.Printf("  📊 Worker %d processing job %d: %s\n", id, job.ID, job.Data)
		time.Sleep(100 * time.Millisecond)
		
		p.mu.Lock()
		p.processed++
		p.mu.Unlock()
		
		p.results <- job
	}
	p.done <- true
}

func (p *ChannelWorkerPool) Wait() {
	close(p.jobs)
	
	// Wait for all workers to finish
	for i := 0; i < p.workers; i++ {
		<-p.done
	}
	
	close(p.results)
}

func (p *ChannelWorkerPool) ProcessedCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.processed
}

// Rate Limiter
type RateLimiter struct {
	rate   int
	period time.Duration
	tokens chan struct{}
	ticker *time.Ticker
}

func NewRateLimiter(rate int, period time.Duration) *RateLimiter {
	rl := &RateLimiter{
		rate:   rate,
		period: period,
		tokens: make(chan struct{}, rate),
	}
	
	// Fill initial tokens
	for i := 0; i < rate; i++ {
		rl.tokens <- struct{}{}
	}
	
	// Start token refill
	rl.ticker = time.NewTicker(period / time.Duration(rate))
	go rl.refill()
	
	return rl
}

func (rl *RateLimiter) refill() {
	for range rl.ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}
	}
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}

// Circuit Breaker
type CircuitBreaker struct {
	failures   int
	threshold  int
	timeout    time.Duration
	lastFail   time.Time
	state      string // "closed", "open", "half-open"
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     "closed",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == "open" {
		if time.Since(cb.lastFail) > cb.timeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	if err != nil {
		cb.failures++
		cb.lastFail = time.Now()
		
		if cb.failures >= cb.threshold {
			cb.state = "open"
		}
		
		return err
	}
	
	cb.failures = 0
	cb.state = "closed"
	return nil
}
