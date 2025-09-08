package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// GracefulShutdown demonstrates basic graceful shutdown
func GracefulShutdown() {
	fmt.Println("=== Basic Graceful Shutdown ===")
	
	// Create a system that can be shut down gracefully
	system := NewGracefulSystem()
	
	// Start the system
	system.Start()
	
	// Let it run for a bit
	time.Sleep(2 * time.Second)
	
	// Shutdown gracefully
	fmt.Println("Initiating graceful shutdown...")
	system.Shutdown()
	
	// Wait for shutdown to complete
	<-system.Done()
	fmt.Println("System shutdown complete")
}

// GracefulSystem represents a system that can be shut down gracefully
type GracefulSystem struct {
	workers    int
	jobs       chan string
	results    chan string
	shutdown   chan struct{}
	done       chan struct{}
	wg         sync.WaitGroup
}

// NewGracefulSystem creates a new graceful system
func NewGracefulSystem() *GracefulSystem {
	return &GracefulSystem{
		workers:  3,
		jobs:    make(chan string, 100),
		results: make(chan string, 100),
		shutdown: make(chan struct{}),
		done:    make(chan struct{}),
	}
}

// Start starts the system
func (gs *GracefulSystem) Start() {
	// Start workers
	for i := 0; i < gs.workers; i++ {
		gs.wg.Add(1)
		go gs.worker(i)
	}
	
	// Start job producer
	gs.wg.Add(1)
	go gs.producer()
	
	// Start result collector
	gs.wg.Add(1)
	go gs.collector()
}

// worker processes jobs
func (gs *GracefulSystem) worker(id int) {
	defer gs.wg.Done()
	
	for {
		select {
		case job := <-gs.jobs:
			// Process job
			time.Sleep(500 * time.Millisecond)
			result := fmt.Sprintf("Worker %d processed: %s", id, job)
			gs.results <- result
		case <-gs.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// producer generates jobs
func (gs *GracefulSystem) producer() {
	defer gs.wg.Done()
	
	jobCount := 0
	for {
		select {
		case gs.jobs <- fmt.Sprintf("job_%d", jobCount):
			jobCount++
			time.Sleep(200 * time.Millisecond)
		case <-gs.shutdown:
			fmt.Println("Producer shutting down")
			return
		}
	}
}

// collector collects results
func (gs *GracefulSystem) collector() {
	defer gs.wg.Done()
	
	for {
		select {
		case result := <-gs.results:
			fmt.Printf("Result: %s\n", result)
		case <-gs.shutdown:
			fmt.Println("Collector shutting down")
			return
		}
	}
}

// Shutdown initiates graceful shutdown
func (gs *GracefulSystem) Shutdown() {
	close(gs.shutdown)
	
	// Wait for all workers to finish
	go func() {
		gs.wg.Wait()
		close(gs.done)
	}()
}

// Done returns a channel that signals when shutdown is complete
func (gs *GracefulSystem) Done() <-chan struct{} {
	return gs.done
}

// ContextBasedShutdown demonstrates graceful shutdown using context
func ContextBasedShutdown() {
	fmt.Println("\n=== Context-Based Graceful Shutdown ===")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Create system with context
	system := NewContextSystem(ctx)
	
	// Start the system
	system.Start()
	
	// Let it run
	time.Sleep(2 * time.Second)
	
	// Shutdown will happen automatically due to context timeout
	fmt.Println("Waiting for context timeout...")
	<-system.Done()
	fmt.Println("Context-based shutdown complete")
}

// ContextSystem represents a system that uses context for shutdown
type ContextSystem struct {
	ctx     context.Context
	workers int
	jobs    chan string
	results chan string
	done    chan struct{}
	wg      sync.WaitGroup
}

// NewContextSystem creates a new context-based system
func NewContextSystem(ctx context.Context) *ContextSystem {
	return &ContextSystem{
		ctx:     ctx,
		workers: 3,
		jobs:    make(chan string, 100),
		results: make(chan string, 100),
		done:    make(chan struct{}),
	}
}

// Start starts the system
func (cs *ContextSystem) Start() {
	// Start workers
	for i := 0; i < cs.workers; i++ {
		cs.wg.Add(1)
		go cs.worker(i)
	}
	
	// Start job producer
	cs.wg.Add(1)
	go cs.producer()
	
	// Start result collector
	cs.wg.Add(1)
	go cs.collector()
	
	// Start shutdown handler
	go cs.shutdownHandler()
}

// worker processes jobs
func (cs *ContextSystem) worker(id int) {
	defer cs.wg.Done()
	
	for {
		select {
		case job := <-cs.jobs:
			// Process job
			time.Sleep(500 * time.Millisecond)
			result := fmt.Sprintf("Worker %d processed: %s", id, job)
			cs.results <- result
		case <-cs.ctx.Done():
			fmt.Printf("Worker %d shutting down due to context\n", id)
			return
		}
	}
}

// producer generates jobs
func (cs *ContextSystem) producer() {
	defer cs.wg.Done()
	
	jobCount := 0
	for {
		select {
		case cs.jobs <- fmt.Sprintf("job_%d", jobCount):
			jobCount++
			time.Sleep(200 * time.Millisecond)
		case <-cs.ctx.Done():
			fmt.Println("Producer shutting down due to context")
			return
		}
	}
}

// collector collects results
func (cs *ContextSystem) collector() {
	defer cs.wg.Done()
	
	for {
		select {
		case result := <-cs.results:
			fmt.Printf("Result: %s\n", result)
		case <-cs.ctx.Done():
			fmt.Println("Collector shutting down due to context")
			return
		}
	}
}

// shutdownHandler handles shutdown
func (cs *ContextSystem) shutdownHandler() {
	<-cs.ctx.Done()
	fmt.Println("Context cancelled, initiating shutdown...")
	cs.wg.Wait()
	close(cs.done)
}

// Done returns a channel that signals when shutdown is complete
func (cs *ContextSystem) Done() <-chan struct{} {
	return cs.done
}

// GracefulShutdownWithTimeout demonstrates graceful shutdown with timeout
func GracefulShutdownWithTimeout() {
	fmt.Println("\n=== Graceful Shutdown with Timeout ===")
	
	// Create system
	system := NewTimeoutSystem(5 * time.Second)
	
	// Start the system
	system.Start()
	
	// Let it run
	time.Sleep(2 * time.Second)
	
	// Shutdown with timeout
	fmt.Println("Initiating graceful shutdown with timeout...")
	system.Shutdown()
	
	// Wait for shutdown to complete
	<-system.Done()
	fmt.Println("Timeout-based shutdown complete")
}

// TimeoutSystem represents a system with timeout-based shutdown
type TimeoutSystem struct {
	timeout  time.Duration
	workers  int
	jobs     chan string
	results  chan string
	shutdown chan struct{}
	done     chan struct{}
	wg       sync.WaitGroup
}

// NewTimeoutSystem creates a new timeout-based system
func NewTimeoutSystem(timeout time.Duration) *TimeoutSystem {
	return &TimeoutSystem{
		timeout:  timeout,
		workers:  3,
		jobs:     make(chan string, 100),
		results:  make(chan string, 100),
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
}

// Start starts the system
func (ts *TimeoutSystem) Start() {
	// Start workers
	for i := 0; i < ts.workers; i++ {
		ts.wg.Add(1)
		go ts.worker(i)
	}
	
	// Start job producer
	ts.wg.Add(1)
	go ts.producer()
	
	// Start result collector
	ts.wg.Add(1)
	go ts.collector()
}

// worker processes jobs
func (ts *TimeoutSystem) worker(id int) {
	defer ts.wg.Done()
	
	for {
		select {
		case job := <-ts.jobs:
			// Process job
			time.Sleep(500 * time.Millisecond)
			result := fmt.Sprintf("Worker %d processed: %s", id, job)
			ts.results <- result
		case <-ts.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// producer generates jobs
func (ts *TimeoutSystem) producer() {
	defer ts.wg.Done()
	
	jobCount := 0
	for {
		select {
		case ts.jobs <- fmt.Sprintf("job_%d", jobCount):
			jobCount++
			time.Sleep(200 * time.Millisecond)
		case <-ts.shutdown:
			fmt.Println("Producer shutting down")
			return
		}
	}
}

// collector collects results
func (ts *TimeoutSystem) collector() {
	defer ts.wg.Done()
	
	for {
		select {
		case result := <-ts.results:
			fmt.Printf("Result: %s\n", result)
		case <-ts.shutdown:
			fmt.Println("Collector shutting down")
			return
		}
	}
}

// Shutdown initiates graceful shutdown with timeout
func (ts *TimeoutSystem) Shutdown() {
	close(ts.shutdown)
	
	// Wait for shutdown with timeout
	done := make(chan struct{})
	go func() {
		ts.wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("Graceful shutdown completed")
	case <-time.After(ts.timeout):
		fmt.Println("Shutdown timeout reached, forcing exit")
	}
	
	close(ts.done)
}

// Done returns a channel that signals when shutdown is complete
func (ts *TimeoutSystem) Done() <-chan struct{} {
	return ts.done
}

// GracefulShutdownWithSignal demonstrates graceful shutdown with signal handling
func GracefulShutdownWithSignal() {
	fmt.Println("\n=== Graceful Shutdown with Signal Handling ===")
	
	// Create system
	system := NewSignalSystem()
	
	// Start the system
	system.Start()
	
	// Let it run
	time.Sleep(2 * time.Second)
	
	// Simulate signal (in real app, this would be SIGINT/SIGTERM)
	fmt.Println("Simulating signal...")
	system.Signal()
	
	// Wait for shutdown to complete
	<-system.Done()
	fmt.Println("Signal-based shutdown complete")
}

// SignalSystem represents a system with signal-based shutdown
type SignalSystem struct {
	workers  int
	jobs     chan string
	results  chan string
	shutdown chan struct{}
	signal   chan struct{}
	done     chan struct{}
	wg       sync.WaitGroup
}

// NewSignalSystem creates a new signal-based system
func NewSignalSystem() *SignalSystem {
	return &SignalSystem{
		workers:  3,
		jobs:     make(chan string, 100),
		results:  make(chan string, 100),
		shutdown: make(chan struct{}),
		signal:   make(chan struct{}),
		done:     make(chan struct{}),
	}
}

// Start starts the system
func (ss *SignalSystem) Start() {
	// Start workers
	for i := 0; i < ss.workers; i++ {
		ss.wg.Add(1)
		go ss.worker(i)
	}
	
	// Start job producer
	ss.wg.Add(1)
	go ss.producer()
	
	// Start result collector
	ss.wg.Add(1)
	go ss.collector()
	
	// Start signal handler
	ss.wg.Add(1)
	go ss.signalHandler()
}

// worker processes jobs
func (ss *SignalSystem) worker(id int) {
	defer ss.wg.Done()
	
	for {
		select {
		case job := <-ss.jobs:
			// Process job
			time.Sleep(500 * time.Millisecond)
			result := fmt.Sprintf("Worker %d processed: %s", id, job)
			ss.results <- result
		case <-ss.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// producer generates jobs
func (ss *SignalSystem) producer() {
	defer ss.wg.Done()
	
	jobCount := 0
	for {
		select {
		case ss.jobs <- fmt.Sprintf("job_%d", jobCount):
			jobCount++
			time.Sleep(200 * time.Millisecond)
		case <-ss.shutdown:
			fmt.Println("Producer shutting down")
			return
		}
	}
}

// collector collects results
func (ss *SignalSystem) collector() {
	defer ss.wg.Done()
	
	for {
		select {
		case result := <-ss.results:
			fmt.Printf("Result: %s\n", result)
		case <-ss.shutdown:
			fmt.Println("Collector shutting down")
			return
		}
	}
}

// signalHandler handles signals
func (ss *SignalSystem) signalHandler() {
	defer ss.wg.Done()
	
	<-ss.signal
	fmt.Println("Signal received, initiating graceful shutdown...")
	close(ss.shutdown)
	ss.wg.Wait()
	close(ss.done)
}

// Signal simulates a signal
func (ss *SignalSystem) Signal() {
	close(ss.signal)
}

// Done returns a channel that signals when shutdown is complete
func (ss *SignalSystem) Done() <-chan struct{} {
	return ss.done
}
