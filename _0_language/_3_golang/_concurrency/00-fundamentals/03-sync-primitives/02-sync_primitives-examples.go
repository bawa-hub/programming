package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// PART 1: BASIC SYNC PRIMITIVES
// ============================================================================

// Example 1: Mutex - Basic Usage
// Mutex provides mutual exclusion - only one goroutine can hold the lock at a time
func example1() {
	fmt.Println("\n=== Example 1: Mutex - Basic Usage ===")
	
	var mu sync.Mutex
	counter := 0
	
	// Worker function that increments counter
	worker := func(id int) {
		mu.Lock()         // Acquire lock
		defer mu.Unlock() // Release lock when function exits
		
		// Critical section - only one goroutine can execute this at a time
		counter++
		fmt.Printf("Worker %d: Counter = %d\n", id, counter)
	}
	
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final counter value: %d\n", counter)
}

// Example 2: Mutex - Without Protection (Race Condition)
// This shows what happens without mutex protection
func example2() {
	fmt.Println("\n=== Example 2: Mutex - Without Protection (Race Condition) ===")
	
	counter := 0
	
	// Worker function WITHOUT mutex protection
	worker := func(id int) {
		// This is unsafe - multiple goroutines can modify counter simultaneously
		counter++
		fmt.Printf("Worker %d: Counter = %d\n", id, counter)
	}
	
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final counter value: %d (may be incorrect due to race condition)\n", counter)
}

// Example 3: RWMutex - Read-Write Lock
// RWMutex allows multiple readers OR one writer, but not both
func example3() {
	fmt.Println("\n=== Example 3: RWMutex - Read-Write Lock ===")
	
	var rwmu sync.RWMutex
	data := make(map[string]int)
	
	// Writer function
	writer := func(key string, value int) {
		rwmu.Lock()         // Exclusive lock for writing
		defer rwmu.Unlock()
		
		data[key] = value
		fmt.Printf("Writer: Set %s = %d\n", key, value)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}
	
	// Reader function
	reader := func(id int) {
		rwmu.RLock()         // Shared lock for reading
		defer rwmu.RUnlock()
		
		fmt.Printf("Reader %d: Data = %v\n", id, data)
		time.Sleep(30 * time.Millisecond) // Simulate work
	}
	
	// Start readers and writers
	go writer("key1", 100)
	go reader(1)
	go reader(2)
	go writer("key2", 200)
	go reader(3)
	
	time.Sleep(300 * time.Millisecond)
}

// Example 4: WaitGroup - Wait for Goroutines
// WaitGroup waits for a collection of goroutines to finish
func example4() {
	fmt.Println("\n=== Example 4: WaitGroup - Wait for Goroutines ===")
	
	var wg sync.WaitGroup
	
	// Worker function
	worker := func(id int) {
		defer wg.Done() // Decrement counter when done
		
		fmt.Printf("Worker %d: Starting work\n", id)
		time.Sleep(100 * time.Millisecond) // Simulate work
		fmt.Printf("Worker %d: Finished work\n", id)
	}
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)    // Increment counter
		go worker(i)
	}
	
	// Wait for all workers to complete
	fmt.Println("Main: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers finished!")
}

// Example 5: Once - One-Time Execution
// Once ensures a function is executed only once
func example5() {
	fmt.Println("\n=== Example 5: Once - One-Time Execution ===")
	
	var once sync.Once
	initialized := false
	
	// Initialization function
	initFunc := func() {
		fmt.Println("Initializing...")
		time.Sleep(100 * time.Millisecond)
		initialized = true
		fmt.Println("Initialization complete")
	}
	
	// Multiple goroutines trying to initialize
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Attempting to initialize\n", id)
			once.Do(initFunc) // This will only execute once
			fmt.Printf("Goroutine %d: Initialized = %t\n", id, initialized)
		}(i)
	}
	
	time.Sleep(200 * time.Millisecond)
}

// Example 6: Cond - Condition Variables
// Cond provides a way for goroutines to wait for a condition
func example6() {
	fmt.Println("\n=== Example 6: Cond - Condition Variables ===")
	
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	// Waiter goroutine
	go func() {
		mu.Lock()
		for !ready {
			fmt.Println("Waiter: Waiting for condition...")
			cond.Wait() // Wait for condition to be signaled
		}
		fmt.Println("Waiter: Condition met! Proceeding...")
		mu.Unlock()
	}()
	
	// Signaler goroutine
	go func() {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		ready = true
		fmt.Println("Signaler: Condition is now true, signaling...")
		cond.Signal() // Signal one waiter
		mu.Unlock()
	}()
	
	time.Sleep(500 * time.Millisecond)
}

// ============================================================================
// PART 2: ADVANCED SYNC PRIMITIVES
// ============================================================================

// Example 7: Atomic Operations
// Atomic operations are lock-free and thread-safe
func example7() {
	fmt.Println("\n=== Example 7: Atomic Operations ===")
	
	var counter int64
	
	// Worker function using atomic operations
	worker := func(id int) {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1) // Atomic increment
		}
		fmt.Printf("Worker %d: Completed\n", id)
	}
	
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
}

// Example 8: Pool - Object Pooling
// Pool provides a way to reuse objects
func example8() {
	fmt.Println("\n=== Example 8: Pool - Object Pooling ===")
	
	// Create a pool of buffers
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024) // Create new buffer
		},
	}
	
	// Get buffer from pool
	buf1 := pool.Get().([]byte)
	fmt.Printf("Got buffer 1: len=%d, cap=%d\n", len(buf1), cap(buf1))
	
	// Use buffer
	copy(buf1, []byte("Hello, World!"))
	fmt.Printf("Buffer 1 content: %s\n", string(buf1[:13]))
	
	// Return buffer to pool
	pool.Put(buf1)
	
	// Get buffer from pool again (should be reused)
	buf2 := pool.Get().([]byte)
	fmt.Printf("Got buffer 2: len=%d, cap=%d\n", len(buf2), cap(buf2))
	fmt.Printf("Buffer 2 content: %s\n", string(buf2[:13])) // Same buffer!
	
	// Return buffer to pool
	pool.Put(buf2)
}

// Example 9: Map - Concurrent Map
// Map is a concurrent map implementation
func example9() {
	fmt.Println("\n=== Example 9: Map - Concurrent Map ===")
	
	var m sync.Map
	
	// Store values
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")
	
	// Load values
	if value, ok := m.Load("key1"); ok {
		fmt.Printf("key1 = %s\n", value)
	}
	
	// Load or store (atomic)
	if value, loaded := m.LoadOrStore("key4", "value4"); loaded {
		fmt.Printf("key4 already existed: %s\n", value)
	} else {
		fmt.Printf("key4 was stored: %s\n", value)
	}
	
	// Delete value
	m.Delete("key2")
	
	// Range over all key-value pairs
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%s = %s\n", key, value)
		return true // Continue iteration
	})
}

// Example 10: Semaphore - Counting Semaphore
// Semaphore limits the number of concurrent operations
func example10() {
	fmt.Println("\n=== Example 10: Semaphore - Counting Semaphore ===")
	
	// Create semaphore with capacity 2
	semaphore := make(chan struct{}, 2)
	
	// Worker function
	worker := func(id int) {
		// Acquire semaphore
		semaphore <- struct{}{}
		defer func() { <-semaphore }() // Release semaphore
		
		fmt.Printf("Worker %d: Acquired semaphore\n", id)
		time.Sleep(200 * time.Millisecond) // Simulate work
		fmt.Printf("Worker %d: Released semaphore\n", id)
	}
	
	// Start 5 workers (only 2 can run concurrently)
	for i := 1; i <= 5; i++ {
		go worker(i)
	}
	
	time.Sleep(1000 * time.Millisecond)
}

// ============================================================================
// PART 3: COMMON PATTERNS
// ============================================================================

// Example 11: Producer-Consumer Pattern
// Using sync primitives for producer-consumer pattern
func example11() {
	fmt.Println("\n=== Example 11: Producer-Consumer Pattern ===")
	
	var mu sync.Mutex
	var wg sync.WaitGroup
	queue := make([]int, 0)
	
	// Producer
	producer := func(id int) {
		defer wg.Done()
		
		for i := 0; i < 3; i++ {
			mu.Lock()
			queue = append(queue, id*100+i)
			fmt.Printf("Producer %d: Produced %d\n", id, id*100+i)
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}
	
	// Consumer
	consumer := func(id int) {
		defer wg.Done()
		
		for i := 0; i < 3; i++ {
			mu.Lock()
			if len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]
				fmt.Printf("Consumer %d: Consumed %d\n", id, item)
			}
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	// Start producers and consumers
	wg.Add(2)
	go producer(1)
	go consumer(1)
	
	wg.Wait()
}

// Example 12: Worker Pool Pattern
// Using sync primitives for worker pool
func example12() {
	fmt.Println("\n=== Example 12: Worker Pool Pattern ===")
	
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Worker function
	worker := func(id int) {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Worker %d: Processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond)
			results <- job * 2
		}
	}
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}
	
	// Send jobs
	go func() {
		defer close(jobs)
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Example 13: Rate Limiting Pattern
// Using sync primitives for rate limiting
func example13() {
	fmt.Println("\n=== Example 13: Rate Limiting Pattern ===")
	
	var mu sync.Mutex
	rateLimiter := time.Tick(200 * time.Millisecond) // Allow 1 request per 200ms
	requestCount := 0
	
	// Request function
	makeRequest := func(id int) {
		<-rateLimiter // Wait for rate limit
		
		mu.Lock()
		requestCount++
		fmt.Printf("Request %d: Processing (total: %d)\n", id, requestCount)
		mu.Unlock()
		
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
	
	// Make requests
	for i := 1; i <= 5; i++ {
		go makeRequest(i)
	}
	
	time.Sleep(1500 * time.Millisecond)
}

// Example 14: Circuit Breaker Pattern
// Using sync primitives for circuit breaker
func example14() {
	fmt.Println("\n=== Example 14: Circuit Breaker Pattern ===")
	
	var mu sync.Mutex
	state := "CLOSED" // CLOSED, OPEN, HALF_OPEN
	failures := 0
	maxFailures := 3
	
	// Service call
	callService := func(id int) bool {
		mu.Lock()
		defer mu.Unlock()
		
		if state == "OPEN" {
			fmt.Printf("Request %d: Circuit breaker is OPEN - request rejected\n", id)
			return false
		}
		
		// Simulate service call (fails every 3rd request)
		success := id%3 != 0
		
		if success {
			fmt.Printf("Request %d: SUCCESS\n", id)
			failures = 0
			state = "CLOSED"
		} else {
			fmt.Printf("Request %d: FAILED\n", id)
			failures++
			if failures >= maxFailures {
				state = "OPEN"
				fmt.Printf("Circuit breaker OPENED after %d failures\n", failures)
			}
		}
		
		return success
	}
	
	// Make requests
	for i := 1; i <= 10; i++ {
		callService(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Example 15: Graceful Shutdown Pattern
// Using sync primitives for graceful shutdown
func example15() {
	fmt.Println("\n=== Example 15: Graceful Shutdown Pattern ===")
	
	var wg sync.WaitGroup
	shutdown := make(chan struct{})
	
	// Worker function
	worker := func(id int) {
		defer wg.Done()
		
		for {
			select {
			case <-shutdown:
				fmt.Printf("Worker %d: Shutting down gracefully\n", id)
				return
			default:
				fmt.Printf("Worker %d: Working...\n", id)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	
	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}
	
	// Let them work for a bit
	time.Sleep(300 * time.Millisecond)
	
	// Shutdown
	fmt.Println("Initiating graceful shutdown...")
	close(shutdown)
	
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers shut down gracefully")
}
