package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Goroutine Pool with Dynamic Scaling
type DynamicPool struct {
	workers    int
	maxWorkers int
	jobs       chan func()
	quit       chan bool
	wg         sync.WaitGroup
	mu         sync.RWMutex
}

func NewDynamicPool(maxWorkers int) *DynamicPool {
	return &DynamicPool{
		workers:    0,
		maxWorkers: maxWorkers,
		jobs:       make(chan func(), 100),
		quit:       make(chan bool),
	}
}

func (p *DynamicPool) Start() {
	go p.manager()
}

func (p *DynamicPool) manager() {
	for {
		select {
		case job := <-p.jobs:
			p.mu.RLock()
			workers := p.workers
			p.mu.RUnlock()
			
			if workers < p.maxWorkers {
				p.addWorker()
			}
			
			// Send job to worker
			select {
			case p.jobs <- job:
			default:
				// If jobs channel is full, create more workers
				if workers < p.maxWorkers {
					p.addWorker()
					p.jobs <- job
				}
			}
		case <-p.quit:
			return
		}
	}
}

func (p *DynamicPool) addWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.workers < p.maxWorkers {
		p.workers++
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *DynamicPool) worker() {
	defer p.wg.Done()
	
	for {
		select {
		case job := <-p.jobs:
			job()
		case <-p.quit:
			return
		}
	}
}

func (p *DynamicPool) Submit(job func()) {
	p.jobs <- job
}

func (p *DynamicPool) Stop() {
	close(p.quit)
	p.wg.Wait()
}

// Advanced Pattern 2: Goroutine with Context Cancellation
func ContextGoroutine() {
	fmt.Println("Advanced Pattern 2: Context Cancellation")
	fmt.Println("========================================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Start goroutine with context
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine: Context cancelled, exiting")
				return
			default:
				fmt.Println("Goroutine: Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	
	// Let it run for a bit
	time.Sleep(3 * time.Second)
}

// Advanced Pattern 3: Goroutine with Heartbeat
type HeartbeatGoroutine struct {
	heartbeat chan time.Time
	done      chan bool
	wg        sync.WaitGroup
}

func NewHeartbeatGoroutine() *HeartbeatGoroutine {
	return &HeartbeatGoroutine{
		heartbeat: make(chan time.Time),
		done:      make(chan bool),
	}
}

func (h *HeartbeatGoroutine) Start() {
	h.wg.Add(1)
	go h.run()
}

func (h *HeartbeatGoroutine) run() {
	defer h.wg.Done()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case t := <-ticker.C:
			select {
			case h.heartbeat <- t:
			default:
				// Heartbeat channel is full, skip
			}
		case <-h.done:
			return
		}
	}
}

func (h *HeartbeatGoroutine) GetHeartbeat() <-chan time.Time {
	return h.heartbeat
}

func (h *HeartbeatGoroutine) Stop() {
	close(h.done)
	h.wg.Wait()
}

// Advanced Pattern 4: Goroutine with Circuit Breaker
type CircuitBreaker struct {
	failures    int64
	threshold   int64
	timeout     time.Duration
	lastFailure time.Time
	state       int32 // 0: closed, 1: open, 2: half-open
	mu          sync.RWMutex
}

func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     0, // closed
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.RLock()
	state := atomic.LoadInt32(&cb.state)
	cb.mu.RUnlock()
	
	switch state {
	case 1: // open
		if time.Since(cb.lastFailure) > cb.timeout {
			atomic.StoreInt32(&cb.state, 2) // half-open
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
			atomic.StoreInt32(&cb.state, 1) // open
		}
		return err
	}
	
	// Success
	cb.failures = 0
	atomic.StoreInt32(&cb.state, 0) // closed
	return nil
}

// Advanced Pattern 5: Goroutine with Backpressure
type BackpressureGoroutine struct {
	input    chan interface{}
	output   chan interface{}
	pressure int
	maxPressure int
	mu       sync.RWMutex
}

func NewBackpressureGoroutine(maxPressure int) *BackpressureGoroutine {
	return &BackpressureGoroutine{
		input:       make(chan interface{}),
		output:      make(chan interface{}),
		pressure:    0,
		maxPressure: maxPressure,
	}
}

func (bp *BackpressureGoroutine) Start() {
	go bp.process()
}

func (bp *BackpressureGoroutine) process() {
	for item := range bp.input {
		bp.mu.Lock()
		if bp.pressure >= bp.maxPressure {
			bp.mu.Unlock()
			// Drop item or block
			continue
		}
		bp.pressure++
		bp.mu.Unlock()
		
		// Process item
		time.Sleep(100 * time.Millisecond) // Simulate work
		
		select {
		case bp.output <- item:
		default:
			// Output channel is full
		}
		
		bp.mu.Lock()
		bp.pressure--
		bp.mu.Unlock()
	}
}

func (bp *BackpressureGoroutine) Send(item interface{}) bool {
	select {
	case bp.input <- item:
		return true
	default:
		return false // Backpressure applied
	}
}

func (bp *BackpressureGoroutine) Receive() <-chan interface{} {
	return bp.output
}

// Advanced Pattern 6: Goroutine with Metrics
type MetricsGoroutine struct {
	processed int64
	errors    int64
	startTime time.Time
	mu        sync.RWMutex
}

func NewMetricsGoroutine() *MetricsGoroutine {
	return &MetricsGoroutine{
		startTime: time.Now(),
	}
}

func (m *MetricsGoroutine) RecordProcessed() {
	atomic.AddInt64(&m.processed, 1)
}

func (m *MetricsGoroutine) RecordError() {
	atomic.AddInt64(&m.errors, 1)
}

func (m *MetricsGoroutine) GetStats() (processed, errors int64, uptime time.Duration) {
	processed = atomic.LoadInt64(&m.processed)
	errors = atomic.LoadInt64(&m.errors)
	uptime = time.Since(m.startTime)
	return
}

// Advanced Pattern 7: Goroutine with Graceful Shutdown
type GracefulGoroutine struct {
	shutdown chan struct{}
	done     chan struct{}
	wg       sync.WaitGroup
}

func NewGracefulGoroutine() *GracefulGoroutine {
	return &GracefulGoroutine{
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
}

func (g *GracefulGoroutine) Start() {
	g.wg.Add(1)
	go g.run()
}

func (g *GracefulGoroutine) run() {
	defer g.wg.Done()
	defer close(g.done)
	
	for {
		select {
		case <-g.shutdown:
			fmt.Println("Goroutine: Received shutdown signal")
			return
		default:
			// Do work
			fmt.Println("Goroutine: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (g *GracefulGoroutine) Shutdown() {
	close(g.shutdown)
	g.wg.Wait()
}

func (g *GracefulGoroutine) Wait() {
	<-g.done
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Goroutine Patterns")
	fmt.Println("==============================")
	
	// Pattern 1: Dynamic Pool
	fmt.Println("\n1. Dynamic Goroutine Pool:")
	dynamicPool := NewDynamicPool(3)
	dynamicPool.Start()
	
	for i := 0; i < 10; i++ {
		dynamicPool.Submit(func() {
			fmt.Printf("Dynamic pool job %d\n", i)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	time.Sleep(2 * time.Second)
	dynamicPool.Stop()
	
	// Pattern 2: Context Cancellation
	ContextGoroutine()
	
	// Pattern 3: Heartbeat
	fmt.Println("\n3. Heartbeat Goroutine:")
	heartbeat := NewHeartbeatGoroutine()
	heartbeat.Start()
	
	go func() {
		for t := range heartbeat.GetHeartbeat() {
			fmt.Printf("Heartbeat: %v\n", t)
		}
	}()
	
	time.Sleep(3 * time.Second)
	heartbeat.Stop()
	
	// Pattern 4: Circuit Breaker
	fmt.Println("\n4. Circuit Breaker:")
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
	
	// Pattern 5: Backpressure
	fmt.Println("\n5. Backpressure Goroutine:")
	bp := NewBackpressureGoroutine(2)
	bp.Start()
	
	for i := 0; i < 5; i++ {
		success := bp.Send(i)
		if success {
			fmt.Printf("Sent item %d\n", i)
		} else {
			fmt.Printf("Failed to send item %d (backpressure)\n", i)
		}
	}
	
	// Pattern 6: Metrics
	fmt.Println("\n6. Metrics Goroutine:")
	metrics := NewMetricsGoroutine()
	
	for i := 0; i < 5; i++ {
		metrics.RecordProcessed()
		if i%2 == 0 {
			metrics.RecordError()
		}
	}
	
	processed, errors, uptime := metrics.GetStats()
	fmt.Printf("Processed: %d, Errors: %d, Uptime: %v\n", processed, errors, uptime)
	
	// Pattern 7: Graceful Shutdown
	fmt.Println("\n7. Graceful Shutdown:")
	graceful := NewGracefulGoroutine()
	graceful.Start()
	
	time.Sleep(2 * time.Second)
	graceful.Shutdown()
	graceful.Wait()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
