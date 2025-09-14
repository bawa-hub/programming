package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 🧵 GOROUTINE LIFECYCLE MASTERY
// Understanding the complete lifecycle of goroutines

func main() {
	fmt.Println("🚀 GOROUTINE LIFECYCLE MASTERY")
	fmt.Println("================================")

	// 1. Basic Goroutine Creation and Execution
	fmt.Println("\n1. Basic Goroutine Creation:")
	basicGoroutine()

	// 2. Goroutine Scheduling and Runtime
	fmt.Println("\n2. Goroutine Scheduling:")
	goroutineScheduling()

	// 3. Goroutine Stack Management
	fmt.Println("\n3. Stack Management:")
	stackManagement()

	// 4. Goroutine Lifecycle with WaitGroup
	fmt.Println("\n4. Lifecycle with WaitGroup:")
	lifecycleWithWaitGroup()

	// 5. Goroutine Leak Prevention
	fmt.Println("\n5. Leak Prevention:")
	leakPrevention()
}

// Basic goroutine creation and execution
func basicGoroutine() {
	fmt.Println("Creating a basic goroutine...")
	
	go func() {
		fmt.Println("  🧵 Goroutine executing!")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  🧵 Goroutine completed!")
	}()

	// Give goroutine time to execute
	time.Sleep(200 * time.Millisecond)
}

// Understanding goroutine scheduling
func goroutineScheduling() {
	fmt.Println("Understanding goroutine scheduling...")
	
	// Get current number of goroutines
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("  Initial goroutines: %d\n", initialGoroutines)

	// Create multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  🧵 Goroutine %d executing\n", id)
			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	
	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("  Final goroutines: %d\n", finalGoroutines)
}

// Stack management and growth
func stackManagement() {
	fmt.Println("Understanding stack management...")
	
	// Get current stack size
	var stackSize int
	stackSize = 1024 * 1024 // 1MB
	fmt.Printf("  Default stack size: %d bytes\n", stackSize)

	// Demonstrate stack growth with recursion
	go func() {
		recurse(0)
	}()

	time.Sleep(100 * time.Millisecond)
}

// Recursive function to demonstrate stack growth
func recurse(depth int) {
	if depth > 1000 {
		fmt.Printf("  🧵 Stack depth reached: %d\n", depth)
		return
	}
	recurse(depth + 1)
}

// Goroutine lifecycle with proper synchronization
func lifecycleWithWaitGroup() {
	fmt.Println("Managing goroutine lifecycle with WaitGroup...")
	
	var wg sync.WaitGroup
	results := make(chan int, 3)

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  🧵 Worker %d starting\n", id)
			
			// Simulate work
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			
			fmt.Printf("  🧵 Worker %d completed\n", id)
			results <- id * 10
		}(i)
	}

	// Close results channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("  📊 Results:")
	for result := range results {
		fmt.Printf("    Result: %d\n", result)
	}
}

// Goroutine leak prevention
func leakPrevention() {
	fmt.Println("Preventing goroutine leaks...")
	
	// BAD: This creates a leak
	fmt.Println("  ❌ BAD: Goroutine leak example")
	leakyGoroutine()

	// GOOD: Proper cleanup
	fmt.Println("  ✅ GOOD: Proper cleanup")
	properCleanup()
}

// BAD: This creates a goroutine leak
func leakyGoroutine() {
	// This goroutine will run forever
	go func() {
		for {
			// This will run indefinitely
			time.Sleep(1 * time.Second)
			// No way to stop this goroutine!
		}
	}()
	
	// Give it time to start
	time.Sleep(100 * time.Millisecond)
	fmt.Println("    ⚠️  Leaky goroutine started (will run forever)")
}

// GOOD: Proper cleanup with context
func properCleanup() {
	// Use a channel to signal shutdown
	shutdown := make(chan struct{})
	
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Println("    🧵 Working...")
			case <-shutdown:
				fmt.Println("    🧵 Shutting down gracefully")
				return
			}
		}
	}()
	
	// Let it run for a bit
	time.Sleep(2 * time.Second)
	
	// Signal shutdown
	close(shutdown)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("    ✅ Goroutine cleaned up properly")
}
