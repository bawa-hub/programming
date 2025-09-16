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
// COMPREHENSIVE SYNC PRIMITIVES EXAMPLES
// ============================================================================

// Example 1: Mutex - Basic Protection
func example1() {
	fmt.Println("\n=== Example 1: Mutex - Basic Protection ===")
	
	var mu sync.Mutex
	counter := 0
	
	// Worker function with mutex protection
	worker := func(id int) {
		for i := 0; i < 1000; i++ {
			mu.Lock()         // Acquire lock
			counter++         // Critical section
			mu.Unlock()       // Release lock
		}
		fmt.Printf("Worker %d: Completed\n", id)
	}
	
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final counter value: %d (should be 3000)\n", counter)
}

// Example 2: Mutex - Race Condition Without Protection
func example2() {
	fmt.Println("\n=== Example 2: Mutex - Race Condition Without Protection ===")
	
	counter := 0
	
	// Worker function WITHOUT mutex protection
	worker := func(id int) {
		for i := 0; i < 1000; i++ {
			counter++ // Race condition - unsafe!
		}
		fmt.Printf("Worker %d: Completed\n", id)
	}
	
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final counter value: %d (may be incorrect due to race condition)\n", counter)
}

// Example 3: RWMutex - Read-Write Lock
func example3() {
	fmt.Println("\n=== Example 3: RWMutex - Read-Write Lock ===")
	
	var rwmu sync.RWMutex
	data := make(map[string]int)
	
	// Writer function
	writer := func(id int) {
		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("key%d", i)
			value := id*100 + i
			
			rwmu.Lock()         // Exclusive lock for writing
			data[key] = value
			fmt.Printf("Writer %d: Set %s = %d\n", id, key, value)
			rwmu.Unlock()
			
			time.Sleep(10 * time.Millisecond)
		}
	}
	
	// Reader function
	reader := func(id int) {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("key%d", i%5)
			
			rwmu.RLock()         // Shared lock for reading
			value := data[key]
			fmt.Printf("Reader %d: Read %s = %d\n", id, key, value)
			rwmu.RUnlock()
			
			time.Sleep(5 * time.Millisecond)
		}
	}
	
	// Start readers and writers
	go writer(1)
	go reader(1)
	go reader(2)
	go reader(3)
	
	time.Sleep(200 * time.Millisecond)
}

// Example 4: WaitGroup - Wait for Goroutines
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
	
	// Start 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)    // Increment counter
		go worker(i)
	}
	
	// Wait for all workers to complete
	fmt.Println("Main: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers finished!")
}

// Example 5: Once - One-Time Execution
func example5() {
	fmt.Println("\n=== Example 5: Once - One-Time Execution ===")
	
	var once sync.Once
	initialized := false
	
	// Initialization function
	initFunc := func() {
		fmt.Println("Initializing expensive resource...")
		time.Sleep(100 * time.Millisecond)
		initialized = true
		fmt.Println("Initialization complete!")
	}
	
	// Multiple goroutines trying to initialize
	for i := 1; i <= 10; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Attempting to initialize\n", id)
			once.Do(initFunc) // This will only execute once
			fmt.Printf("Goroutine %d: Initialized = %t\n", id, initialized)
		}(i)
	}
	
	time.Sleep(200 * time.Millisecond)
}

// Example 6: Cond - Condition Variables
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

// Example 7: Pool - Object Pooling
func example7() {
	fmt.Println("\n=== Example 7: Pool - Object Pooling ===")
	
	// Create a pool of buffers
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new buffer")
			return make([]byte, 1024) // Create new buffer
		},
	}
	
	// Get buffer from pool
	fmt.Println("Getting buffer 1...")
	buf1 := pool.Get().([]byte)
	fmt.Printf("Got buffer 1: len=%d, cap=%d\n", len(buf1), cap(buf1))
	
	// Use buffer
	copy(buf1, []byte("Hello, World!"))
	fmt.Printf("Buffer 1 content: %s\n", string(buf1[:13]))
	
	// Return buffer to pool
	fmt.Println("Returning buffer 1 to pool...")
	pool.Put(buf1)
	
	// Get buffer from pool again (should be reused)
	fmt.Println("Getting buffer 2...")
	buf2 := pool.Get().([]byte)
	fmt.Printf("Got buffer 2: len=%d, cap=%d\n", len(buf2), cap(buf2))
	fmt.Printf("Buffer 2 content: %s\n", string(buf2[:13])) // Same buffer!
	
	// Return buffer to pool
	pool.Put(buf2)
}

// Example 8: Map - Concurrent Map
func example8() {
	fmt.Println("\n=== Example 8: Map - Concurrent Map ===")
	
	var m sync.Map
	
	// Store values
	fmt.Println("Storing values...")
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")
	
	// Load values
	fmt.Println("Loading values...")
	if value, ok := m.Load("key1"); ok {
		fmt.Printf("key1 = %s\n", value)
	}
	
	// Load or store (atomic)
	fmt.Println("LoadOrStore...")
	if value, loaded := m.LoadOrStore("key4", "value4"); loaded {
		fmt.Printf("key4 already existed: %s\n", value)
	} else {
		fmt.Printf("key4 was stored: %s\n", value)
	}
	
	// Try to load non-existent key
	if value, ok := m.Load("nonexistent"); ok {
		fmt.Printf("nonexistent = %s\n", value)
	} else {
		fmt.Println("nonexistent key not found")
	}
	
	// Delete value
	fmt.Println("Deleting key2...")
	m.Delete("key2")
	
	// Range over all key-value pairs
	fmt.Println("All key-value pairs:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %s = %s\n", key, value)
		return true // Continue iteration
	})
}

// Example 9: Atomic Operations
func example9() {
	fmt.Println("\n=== Example 9: Atomic Operations ===")
	
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
	
	// Demonstrate CompareAndSwap
	fmt.Println("\nDemonstrating CompareAndSwap...")
	old := atomic.LoadInt64(&counter)
	new := old + 1
	if atomic.CompareAndSwapInt64(&counter, old, new) {
		fmt.Printf("Successfully updated counter from %d to %d\n", old, new)
	} else {
		fmt.Println("CompareAndSwap failed")
	}
}

// Example 10: Semaphore - Counting Semaphore
func example10() {
	fmt.Println("\n=== Example 10: Semaphore - Counting Semaphore ===")
	
	// Create semaphore with capacity 2
	semaphore := make(chan struct{}, 2)
	
	// Worker function
	worker := func(id int) {
		// Acquire semaphore
		fmt.Printf("Worker %d: Trying to acquire semaphore\n", id)
		semaphore <- struct{}{}
		defer func() { 
			<-semaphore 
			fmt.Printf("Worker %d: Released semaphore\n", id)
		}()
		
		fmt.Printf("Worker %d: Acquired semaphore\n", id)
		time.Sleep(200 * time.Millisecond) // Simulate work
		fmt.Printf("Worker %d: Work completed\n", id)
	}
	
	// Start 5 workers (only 2 can run concurrently)
	for i := 1; i <= 5; i++ {
		go worker(i)
	}
	
	time.Sleep(1000 * time.Millisecond)
}

// Example 11: Producer-Consumer with Mutex
func example11() {
	fmt.Println("\n=== Example 11: Producer-Consumer with Mutex ===")
	
	var mu sync.Mutex
	var wg sync.WaitGroup
	queue := make([]int, 0)
	
	// Producer function
	producer := func(id int) {
		defer wg.Done()
		
		for i := 0; i < 3; i++ {
			mu.Lock()
			item := id*100 + i
			queue = append(queue, item)
			fmt.Printf("Producer %d: Produced %d\n", id, item)
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}
	
	// Consumer function
	consumer := func(id int) {
		defer wg.Done()
		
		for i := 0; i < 3; i++ {
			mu.Lock()
			if len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]
				fmt.Printf("Consumer %d: Consumed %d\n", id, item)
			} else {
				fmt.Printf("Consumer %d: No items available\n", id)
			}
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	// Start producers and consumers
	wg.Add(4)
	go producer(1)
	go producer(2)
	go consumer(1)
	go consumer(2)
	
	wg.Wait()
}

// Example 12: Worker Pool with WaitGroup
func example12() {
	fmt.Println("\n=== Example 12: Worker Pool with WaitGroup ===")
	
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

// Example 13: Rate Limiting with Mutex
func example13() {
	fmt.Println("\n=== Example 13: Rate Limiting with Mutex ===")
	
	var mu sync.Mutex
	lastTime := time.Now()
	interval := 100 * time.Millisecond
	
	// Request function
	makeRequest := func(id int) {
		mu.Lock()
		now := time.Now()
		if now.Sub(lastTime) >= interval {
			lastTime = now
			fmt.Printf("Request %d: Allowed\n", id)
		} else {
			fmt.Printf("Request %d: Rate limited\n", id)
		}
		mu.Unlock()
	}
	
	// Make requests
	for i := 1; i <= 10; i++ {
		go makeRequest(i)
		time.Sleep(50 * time.Millisecond)
	}
	
	time.Sleep(500 * time.Millisecond)
}

// Example 14: Circuit Breaker with Mutex
func example14() {
	fmt.Println("\n=== Example 14: Circuit Breaker with Mutex ===")
	
	var mu sync.Mutex
	state := "CLOSED" // CLOSED, OPEN, HALF_OPEN
	failures := 0
	maxFailures := 3
	
	// Service call function
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

// Example 15: Graceful Shutdown with WaitGroup
func example15() {
	fmt.Println("\n=== Example 15: Graceful Shutdown with WaitGroup ===")
	
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

// ============================================================================
// MAIN FUNCTION
// ============================================================================

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	arg := os.Args[1]
	
	if arg == "all" {
		runAllExamples()
		return
	}
	
	if arg == "basic" {
		runBasicExamples()
		return
	}
	
	if arg == "advanced" {
		runAdvancedExamples()
		return
	}
	
	if arg == "patterns" {
		runPatternExamples()
		return
	}
	
	// Try to parse as example number
	exampleNum, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid argument: %s\n", arg)
		showUsage()
		return
	}
	
	if exampleNum < 1 || exampleNum > 15 {
		fmt.Printf("Example number must be between 1 and 15, got: %d\n", exampleNum)
		showUsage()
		return
	}
	
	runExample(exampleNum)
}

func showUsage() {
	fmt.Println("🔒 Go Sync Primitives - Complete Examples")
	fmt.Println("=========================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run sync-primitives-examples.go <example_number>  # Run specific example (1-15)")
	fmt.Println("  go run sync-primitives-examples.go basic             # Run basic primitives (1-6)")
	fmt.Println("  go run sync-primitives-examples.go advanced          # Run advanced primitives (7-10)")
	fmt.Println("  go run sync-primitives-examples.go patterns          # Run common patterns (11-15)")
	fmt.Println("  go run sync-primitives-examples.go all               # Run all examples")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  1  - Mutex - Basic Protection")
	fmt.Println("  2  - Mutex - Race Condition Without Protection")
	fmt.Println("  3  - RWMutex - Read-Write Lock")
	fmt.Println("  4  - WaitGroup - Wait for Goroutines")
	fmt.Println("  5  - Once - One-Time Execution")
	fmt.Println("  6  - Cond - Condition Variables")
	fmt.Println("  7  - Pool - Object Pooling")
	fmt.Println("  8  - Map - Concurrent Map")
	fmt.Println("  9  - Atomic Operations")
	fmt.Println("  10 - Semaphore - Counting Semaphore")
	fmt.Println("  11 - Producer-Consumer with Mutex")
	fmt.Println("  12 - Worker Pool with WaitGroup")
	fmt.Println("  13 - Rate Limiting with Mutex")
	fmt.Println("  14 - Circuit Breaker with Mutex")
	fmt.Println("  15 - Graceful Shutdown with WaitGroup")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run sync-primitives-examples.go 1    # Run example 1")
	fmt.Println("  go run sync-primitives-examples.go 5    # Run example 5")
	fmt.Println("  go run sync-primitives-examples.go basic # Run basic primitives")
	fmt.Println("  go run sync-primitives-examples.go all  # Run all examples")
}

func runAllExamples() {
	fmt.Println("🚀 Running All Sync Primitive Examples")
	fmt.Println("======================================")
	
	examples := getAllExamples()
	
	for i, example := range examples {
		fmt.Printf("\n--- Example %d ---\n", i+1)
		example()
	}
	
	fmt.Println("\n🎉 All examples completed!")
}

func runBasicExamples() {
	fmt.Println("🔒 Running Basic Sync Primitives (1-6)")
	fmt.Println("======================================")
	
	examples := getBasicExamples()
	
	for i, example := range examples {
		fmt.Printf("\n--- Example %d ---\n", i+1)
		example()
	}
	
	fmt.Println("\n✅ Basic primitives completed!")
}

func runAdvancedExamples() {
	fmt.Println("🔧 Running Advanced Sync Primitives (7-10)")
	fmt.Println("==========================================")
	
	examples := getAdvancedExamples()
	
	for i, example := range examples {
		fmt.Printf("\n--- Example %d ---\n", i+7)
		example()
	}
	
	fmt.Println("\n✅ Advanced primitives completed!")
}

func runPatternExamples() {
	fmt.Println("🎯 Running Common Patterns (11-15)")
	fmt.Println("==================================")
	
	examples := getPatternExamples()
	
	for i, example := range examples {
		fmt.Printf("\n--- Example %d ---\n", i+11)
		example()
	}
	
	fmt.Println("\n✅ Common patterns completed!")
}

func runExample(exampleNum int) {
	fmt.Printf("🔒 Running Example %d\n", exampleNum)
	fmt.Println("==================")
	
	examples := getAllExamples()
	
	if exampleNum < 1 || exampleNum > len(examples) {
		fmt.Printf("Example %d not found\n", exampleNum)
		return
	}
	
	examples[exampleNum-1]()
	fmt.Printf("\n✅ Example %d completed!\n", exampleNum)
}

func getAllExamples() []func() {
	return []func(){
		example1, example2, example3, example4, example5,
		example6, example7, example8, example9, example10,
		example11, example12, example13, example14, example15,
	}
}

func getBasicExamples() []func() {
	return []func(){
		example1, example2, example3, example4, example5, example6,
	}
}

func getAdvancedExamples() []func() {
	return []func(){
		example7, example8, example9, example10,
	}
}

func getPatternExamples() []func() {
	return []func(){
		example11, example12, example13, example14, example15,
	}
}
