package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Example 1: Basic Race Condition
func basicRaceCondition() {
	fmt.Println("\n1. Basic Race Condition")
	fmt.Println("=======================")
	
	var counter int
	
	// Start multiple goroutines that modify the same variable
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++ // Race condition!
			}
		}()
	}
	
	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	
	fmt.Printf("  Counter value: %d (expected: 10000, but may be different due to race)\n", counter)
	fmt.Println("  This demonstrates a race condition - the result is unpredictable")
}

// Example 2: Race Condition with Maps
func raceConditionWithMaps() {
	fmt.Println("\n2. Race Condition with Maps")
	fmt.Println("===========================")
	
	fmt.Println("  This example demonstrates race conditions with maps")
	fmt.Println("  Maps are not thread-safe and will cause panics with concurrent access")
	fmt.Println("  Run with 'go run -race .' to see race detection in action")
	
	// Instead of actually causing a panic, we'll just explain the concept
	fmt.Println("  ❌ Bad: Concurrent map access without synchronization")
	fmt.Println("  ✅ Good: Use sync.Map or mutex protection")
	fmt.Println("  ✅ Good: Use channels for communication")
}

// Example 3: Fixing Race Condition with Mutex
func fixingRaceConditionWithMutex() {
	fmt.Println("\n3. Fixing Race Condition with Mutex")
	fmt.Println("===================================")
	
	var counter int
	var mu sync.Mutex
	
	// Start multiple goroutines that modify the same variable
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++ // Now protected by mutex
				mu.Unlock()
			}
		}()
	}
	
	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	
	fmt.Printf("  Counter value: %d (expected: 10000)\n", counter)
	fmt.Println("  This demonstrates how mutex fixes the race condition")
}

// Example 4: Fixing Race Condition with Atomic Operations
func fixingRaceConditionWithAtomic() {
	fmt.Println("\n4. Fixing Race Condition with Atomic Operations")
	fmt.Println("==============================================")
	
	var counter int64
	
	// Start multiple goroutines that modify the same variable
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // Atomic operation
			}
		}()
	}
	
	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	
	fmt.Printf("  Counter value: %d (expected: 10000)\n", atomic.LoadInt64(&counter))
	fmt.Println("  This demonstrates how atomic operations fix the race condition")
}

// Example 5: Happens-Before with Channels
func happensBeforeWithChannels() {
	fmt.Println("\n5. Happens-Before with Channels")
	fmt.Println("===============================")
	
	var x int
	ch := make(chan int)
	
	// Goroutine 1: writes x, then sends on channel
	go func() {
		x = 42
		ch <- 1 // This establishes happens-before relationship
	}()
	
	// Goroutine 2: receives from channel, then reads x
	go func() {
		<-ch // This establishes happens-before relationship
		fmt.Printf("  x = %d (guaranteed to be 42)\n", x)
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Channel operations establish happens-before relationships")
}

// Example 6: Happens-Before with Mutex
func happensBeforeWithMutex() {
	fmt.Println("\n6. Happens-Before with Mutex")
	fmt.Println("============================")
	
	var x int
	var mu sync.Mutex
	
	// Goroutine 1: locks, writes x, unlocks
	go func() {
		mu.Lock()
		x = 42
		mu.Unlock() // Unlock happens before next Lock
	}()
	
	// Goroutine 2: locks, reads x, unlocks
	go func() {
		mu.Lock() // This establishes happens-before relationship
		fmt.Printf("  x = %d (guaranteed to be 42)\n", x)
		mu.Unlock()
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Mutex operations establish happens-before relationships")
}

// Example 7: Atomic Operations and Memory Ordering
func atomicOperationsAndMemoryOrdering() {
	fmt.Println("\n7. Atomic Operations and Memory Ordering")
	fmt.Println("=======================================")
	
	var x, y int64
	
	// Goroutine 1: writes x, then y
	go func() {
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 2)
	}()
	
	// Goroutine 2: reads y, then x
	go func() {
		for atomic.LoadInt64(&y) != 2 {
			// Wait for y to be set
		}
		// x is guaranteed to be 1 here due to sequential consistency
		fmt.Printf("  x = %d, y = %d (x guaranteed to be 1)\n", 
			atomic.LoadInt64(&x), atomic.LoadInt64(&y))
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Atomic operations provide sequential consistency")
}

// Example 8: Visibility Guarantees
func visibilityGuarantees() {
	fmt.Println("\n8. Visibility Guarantees")
	fmt.Println("========================")
	
	// Without synchronization - no visibility guarantee
	fmt.Println("  Without synchronization:")
	var x1 int
	go func() {
		x1 = 42
	}()
	go func() {
		// x1 might be 0 or 42 - no guarantee
		fmt.Printf("    x1 = %d (unpredictable)\n", x1)
	}()
	time.Sleep(100 * time.Millisecond)
	
	// With atomic operations - visibility guarantee
	fmt.Println("  With atomic operations:")
	var x2 int64
	go func() {
		atomic.StoreInt64(&x2, 42)
	}()
	go func() {
		// x2 is guaranteed to be 42
		fmt.Printf("    x2 = %d (guaranteed to be 42)\n", atomic.LoadInt64(&x2))
	}()
	time.Sleep(100 * time.Millisecond)
}

// Example 9: Race Detection Example
func raceDetectionExample() {
	fmt.Println("\n9. Race Detection Example")
	fmt.Println("=========================")
	
	var counter int
	
	// This will trigger race detection when run with -race
	for i := 0; i < 100; i++ {
		go func() {
			counter++ // Race condition
		}()
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Counter: %d\n", counter)
	fmt.Println("  Run with 'go run -race .' to see race detection in action")
}

// Example 10: Safe Counter with Mutex
func safeCounterWithMutex() {
	fmt.Println("\n10. Safe Counter with Mutex")
	fmt.Println("===========================")
	
	type SafeCounter struct {
		mu      sync.Mutex
		counter int
	}
	
	sc := &SafeCounter{}
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				sc.mu.Lock()
				sc.counter++
				sc.mu.Unlock()
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Safe counter value: %d (expected: 10000)\n", sc.counter)
}

// Example 11: Safe Counter with Atomic Operations
func safeCounterWithAtomic() {
	fmt.Println("\n11. Safe Counter with Atomic Operations")
	fmt.Println("======================================")
	
	type AtomicCounter struct {
		counter int64
	}
	
	ac := &AtomicCounter{}
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&ac.counter, 1)
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Atomic counter value: %d (expected: 10000)\n", atomic.LoadInt64(&ac.counter))
}

// Example 12: Compare and Swap
func compareAndSwap() {
	fmt.Println("\n12. Compare and Swap")
	fmt.Println("===================")
	
	var value int64 = 42
	
	// Try to change 42 to 100
	go func() {
		if atomic.CompareAndSwapInt64(&value, 42, 100) {
			fmt.Println("  Successfully changed value to 100")
		} else {
			fmt.Println("  Failed to change value to 100")
		}
	}()
	
	// Try to change 42 to 200
	go func() {
		if atomic.CompareAndSwapInt64(&value, 42, 200) {
			fmt.Println("  Successfully changed value to 200")
		} else {
			fmt.Println("  Failed to change value to 200")
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Final value: %d\n", atomic.LoadInt64(&value))
}

// Example 13: Atomic Flag
func atomicFlag() {
	fmt.Println("\n13. Atomic Flag")
	fmt.Println("===============")
	
	var flag int32
	
	// Set flag after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		atomic.StoreInt32(&flag, 1)
		fmt.Println("  Flag set to 1")
	}()
	
	// Wait for flag
	fmt.Println("  Waiting for flag...")
	for atomic.LoadInt32(&flag) == 0 {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("  Flag is now set!")
}

// Example 14: Performance Comparison
func performanceComparison() {
	fmt.Println("\n14. Performance Comparison")
	fmt.Println("==========================")
	
	iterations := 1000000
	
	// Mutex performance
	start := time.Now()
	var mu sync.Mutex
	var counter1 int
	for i := 0; i < iterations; i++ {
		mu.Lock()
		counter1++
		mu.Unlock()
	}
	mutexDuration := time.Since(start)
	
	// Atomic performance
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicDuration := time.Since(start)
	
	fmt.Printf("  Mutex: %v for %d operations\n", mutexDuration, iterations)
	fmt.Printf("  Atomic: %v for %d operations\n", atomicDuration, iterations)
	fmt.Printf("  Atomic is %.2fx faster\n", float64(mutexDuration)/float64(atomicDuration))
}

// Example 15: Memory Model Best Practices
func memoryModelBestPractices() {
	fmt.Println("\n15. Memory Model Best Practices")
	fmt.Println("===============================")
	
	fmt.Println("  ✅ Best Practices:")
	fmt.Println("    - Use atomic operations for simple counters/flags")
	fmt.Println("    - Use mutexes for complex data structures")
	fmt.Println("    - Use channels for communication patterns")
	fmt.Println("    - Always run race detector in development")
	fmt.Println("    - Avoid data races at all costs")
	
	fmt.Println("\n  ❌ Common Mistakes:")
	fmt.Println("    - Accessing shared data without synchronization")
	fmt.Println("    - Using mutexes for simple atomic operations")
	fmt.Println("    - Ignoring race detector warnings")
	fmt.Println("    - Assuming sequential consistency without synchronization")
	
	// Demonstrate good practice
	var counter int64
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // Good: atomic operation
		}()
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n  Example: Atomic counter = %d\n", atomic.LoadInt64(&counter))
}

// Example 16: Lock-Free Data Structure
func lockFreeDataStructure() {
	fmt.Println("\n16. Lock-Free Data Structure")
	fmt.Println("============================")
	
	// Simple lock-free counter using atomic operations
	type LockFreeCounter struct {
		value int64
	}
	
	counter := &LockFreeCounter{}
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter.value, 1)
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Lock-free counter value: %d (expected: 10000)\n", atomic.LoadInt64(&counter.value))
}

// Example 17: False Sharing Example
func falseSharingExample() {
	fmt.Println("\n17. False Sharing Example")
	fmt.Println("=========================")
	
	// Bad: False sharing
	type BadCounter struct {
		counter1 int64
		counter2 int64 // Same cache line as counter1
	}
	
	badCounter := &BadCounter{}
	
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(&badCounter.counter1, 1)
		atomic.AddInt64(&badCounter.counter2, 1)
	}
	badDuration := time.Since(start)
	
	// Good: Avoid false sharing with padding
	type GoodCounter struct {
		counter1 int64
		_        [7]int64 // Padding to avoid false sharing
		counter2 int64
	}
	
	goodCounter := &GoodCounter{}
	
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(&goodCounter.counter1, 1)
		atomic.AddInt64(&goodCounter.counter2, 1)
	}
	goodDuration := time.Since(start)
	
	fmt.Printf("  Bad (false sharing): %v\n", badDuration)
	fmt.Printf("  Good (no false sharing): %v\n", goodDuration)
	fmt.Printf("  Improvement: %.2fx faster\n", float64(badDuration)/float64(goodDuration))
}

// Example 18: Memory Barrier Example
func basicMemoryBarrierExample() {
	fmt.Println("\n18. Memory Barrier Example")
	fmt.Println("==========================")
	
	var x, y int64
	
	// Goroutine 1: writes x, then y
	go func() {
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 1) // This acts as a memory barrier
	}()
	
	// Goroutine 2: reads y, then x
	go func() {
		for atomic.LoadInt64(&y) != 1 {
			// Wait for y to be set
		}
		// Due to sequential consistency, x is guaranteed to be 1
		fmt.Printf("  x = %d, y = %d (x guaranteed to be 1)\n", 
			atomic.LoadInt64(&x), atomic.LoadInt64(&y))
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Atomic operations provide memory barriers")
}

// Example 19: Race Condition in Slice Operations
func raceConditionInSliceOperations() {
	fmt.Println("\n19. Race Condition in Slice Operations")
	fmt.Println("=====================================")
	
	// This will likely cause a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Panic caught: %v\n", r)
			fmt.Println("  This demonstrates why slices are not thread-safe")
		}
	}()
	
	s := make([]int, 0, 1000)
	
	// Start goroutines that modify the slice
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				s = append(s, id*100+j) // Race condition!
			}
		}(i)
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Slice length: %d\n", len(s))
}

// Example 20: Safe Slice Operations
func safeSliceOperations() {
	fmt.Println("\n20. Safe Slice Operations")
	fmt.Println("========================")
	
	// Use channels for safe slice operations
	ch := make(chan int, 1000)
	
	// Start goroutines that send to channel
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				ch <- id*100 + j
			}
		}(i)
	}
	
	// Collect results in main goroutine
	var results []int
	go func() {
		for val := range ch {
			results = append(results, val)
		}
	}()
	
	time.Sleep(1 * time.Second)
	close(ch)
	time.Sleep(100 * time.Millisecond)
	
	fmt.Printf("  Safe slice length: %d\n", len(results))
	fmt.Println("  Using channels avoids race conditions")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🧠 Memory Model & Race Conditions Examples")
	fmt.Println("==========================================")
	
	basicRaceCondition()
	raceConditionWithMaps()
	fixingRaceConditionWithMutex()
	fixingRaceConditionWithAtomic()
	happensBeforeWithChannels()
	happensBeforeWithMutex()
	atomicOperationsAndMemoryOrdering()
	visibilityGuarantees()
	raceDetectionExample()
	safeCounterWithMutex()
	safeCounterWithAtomic()
	compareAndSwap()
	atomicFlag()
	performanceComparison()
	memoryModelBestPractices()
	lockFreeDataStructure()
	falseSharingExample()
	basicMemoryBarrierExample()
	raceConditionInSliceOperations()
	safeSliceOperations()
}
