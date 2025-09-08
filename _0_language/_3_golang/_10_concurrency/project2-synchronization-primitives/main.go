package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Concurrency Learning - Project 2 ===")
	fmt.Println("Synchronization Primitives")
	fmt.Println()

	// Exercise 1: Basic Mutex Usage
	fmt.Println("Exercise 1: Basic Mutex Usage")
	demonstrateBasicMutex()
	fmt.Println()

	// Exercise 2: RWMutex for Read-Heavy Workloads
	fmt.Println("Exercise 2: RWMutex for Read-Heavy Workloads")
	demonstrateRWMutex()
	fmt.Println()

	// Exercise 3: WaitGroup Coordination
	fmt.Println("Exercise 3: WaitGroup Coordination")
	demonstrateWaitGroup()
	fmt.Println()

	// Exercise 4: Atomic Operations
	fmt.Println("Exercise 4: Atomic Operations")
	demonstrateAtomicOperations()
	fmt.Println()

	// Exercise 5: sync.Once
	fmt.Println("Exercise 5: sync.Once")
	demonstrateSyncOnce()
	fmt.Println()

	// Exercise 6: Performance Comparison
	fmt.Println("Exercise 6: Performance Comparison")
	demonstratePerformanceComparison()
	fmt.Println()

	fmt.Println("=== All synchronization exercises completed! ===")
	fmt.Println()
	fmt.Println("Run specific components:")
	fmt.Println("  go run main.go cache.go")
	fmt.Println("  go run main.go rate_limiter.go")
	fmt.Println("  go run main.go connection_pool.go")
	fmt.Println("  go run main.go atomic_counter.go")
}

func demonstrateBasicMutex() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Start multiple goroutines that increment counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (should be 5000)\n", counter)
}

func demonstrateRWMutex() {
	var data map[string]int
	var rwmutex sync.RWMutex
	var wg sync.WaitGroup

	// Initialize data
	data = make(map[string]int)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		data[key] = i * 10
	}

	// Multiple readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				rwmutex.RLock()
				key := fmt.Sprintf("key%d", j)
				value := data[key]
				rwmutex.RUnlock()
				fmt.Printf("Reader %d: %s = %d\n", readerID, key, value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			rwmutex.Lock()
			key := fmt.Sprintf("key%d", i)
			data[key] = data[key] + 100
			fmt.Printf("Writer: Updated %s = %d\n", key, data[key])
			rwmutex.Unlock()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func demonstrateWaitGroup() {
	var wg sync.WaitGroup
	results := make(chan int, 10)

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", workerID)
			
			// Simulate work
			time.Sleep(time.Duration(workerID) * 500 * time.Millisecond)
			
			result := workerID * 100
			results <- result
			fmt.Printf("Worker %d completed with result: %d\n", workerID, result)
		}(i)
	}

	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("Collecting results:")
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}
}

func demonstrateAtomicOperations() {
	var counter int64
	var wg sync.WaitGroup

	// Start multiple goroutines that increment counter atomically
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// Atomic increment - no mutex needed!
				// atomic.AddInt64(&counter, 1)
				counter++ // This will have race conditions without atomic
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (may have race conditions)\n", counter)
	fmt.Println("Note: Use atomic.AddInt64(&counter, 1) for thread-safe increment")
}

func demonstrateSyncOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	// Multiple goroutines trying to initialize
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: attempting initialization\n", id)
			
			once.Do(func() {
				fmt.Printf("Goroutine %d: performing initialization\n", id)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Initialization completed")
			})
			
			fmt.Printf("Goroutine %d: initialization done\n", id)
		}(i)
	}

	wg.Wait()
}

func demonstratePerformanceComparison() {
	const iterations = 1000000

	// Test with mutex
	start := time.Now()
	var counter1 int64
	var mutex sync.Mutex
	for i := 0; i < iterations; i++ {
		mutex.Lock()
		counter1++
		mutex.Unlock()
	}
	mutexTime := time.Since(start)

	// Test with atomic (simulated)
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		// atomic.AddInt64(&counter2, 1)
		counter2++ // This will have race conditions
	}
	atomicTime := time.Since(start)

	fmt.Printf("Mutex approach: %v\n", mutexTime)
	fmt.Printf("Atomic approach: %v\n", atomicTime)
	fmt.Printf("Speedup: %.2fx\n", float64(mutexTime)/float64(atomicTime))
	fmt.Println("Note: Use atomic operations for simple operations like counters")
}
