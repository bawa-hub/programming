package main

import (
	"fmt"
	"sync"
	"time"
)

// 🔒 MUTEX PATTERNS MASTERY
// Understanding mutexes, RWMutexes, and synchronization patterns

func main() {
	fmt.Println("🔒 MUTEX PATTERNS MASTERY")
	fmt.Println("=========================")

	// 1. Basic Mutex
	fmt.Println("\n1. Basic Mutex:")
	basicMutex()

	// 2. RWMutex (Read-Write Mutex)
	fmt.Println("\n2. RWMutex (Read-Write Mutex):")
	rwMutex()

	// 3. Mutex with Struct
	fmt.Println("\n3. Mutex with Struct:")
	mutexWithStruct()

	// 4. Deadlock Prevention
	fmt.Println("\n4. Deadlock Prevention:")
	deadlockPrevention()

	// 5. Mutex Performance
	fmt.Println("\n5. Mutex Performance:")
	mutexPerformance()

	// 6. WaitGroup Patterns
	fmt.Println("\n6. WaitGroup Patterns:")
	waitGroupPatterns()

	// 7. Once Pattern
	fmt.Println("\n7. Once Pattern:")
	oncePattern()
}

// BASIC MUTEX: Protecting shared data
func basicMutex() {
	fmt.Println("Understanding basic mutex...")
	
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
			
			fmt.Printf("  🧵 Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  📊 Final counter value: %d (expected: 5000)\n", counter)
}

// RWMUTEX: Optimizing for read-heavy workloads
func rwMutex() {
	fmt.Println("Understanding RWMutex...")
	
	var data map[string]int
	var rwMutex sync.RWMutex
	var wg sync.WaitGroup
	
	// Initialize data
	data = make(map[string]int)
	data["count"] = 0
	
	// Start readers (multiple can read simultaneously)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 5; j++ {
				rwMutex.RLock()
				value := data["count"]
				fmt.Printf("  📖 Reader %d: count = %d\n", id, value)
				rwMutex.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	// Start writer (exclusive access)
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for j := 0; j < 3; j++ {
			rwMutex.Lock()
			data["count"]++
			fmt.Printf("  ✏️  Writer: updated count to %d\n", data["count"])
			rwMutex.Unlock()
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Printf("  📊 Final count: %d\n", data["count"])
}

// MUTEX WITH STRUCT: Embedding mutex in structs
func mutexWithStruct() {
	fmt.Println("Understanding mutex with struct...")
	
	// Create a thread-safe counter
	counter := NewSafeCounter()
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 5; j++ {
				counter.Increment()
				value := counter.Get()
				fmt.Printf("  🧵 Goroutine %d: counter = %d\n", id, value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  📊 Final counter value: %d\n", counter.Get())
}

// SafeCounter is a thread-safe counter
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// NewSafeCounter creates a new safe counter
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Increment increments the counter
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Get returns the current counter value
func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// DEADLOCK PREVENTION: Avoiding deadlocks
func deadlockPrevention() {
	fmt.Println("Understanding deadlock prevention...")
	
	// BAD: This can cause deadlock
	fmt.Println("  ❌ BAD: Potential deadlock example")
	badDeadlockExample()
	
	// GOOD: Proper lock ordering
	fmt.Println("  ✅ GOOD: Proper lock ordering")
	goodDeadlockPrevention()
}

// BAD: This can cause deadlock
func badDeadlockExample() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// Goroutine 1: locks mu1, then mu2
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("  🧵 Goroutine 1: locked mu1")
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		fmt.Println("  🧵 Goroutine 1: locked mu2")
		mu2.Unlock()
		mu1.Unlock()
	}()
	
	// Goroutine 2: locks mu2, then mu1 (DEADLOCK!)
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("  🧵 Goroutine 2: locked mu2")
		time.Sleep(100 * time.Millisecond)
		mu1.Lock()
		fmt.Println("  🧵 Goroutine 2: locked mu1")
		mu1.Unlock()
		mu2.Unlock()
	}()
	
	// Wait with timeout to avoid hanging
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("  ✅ No deadlock occurred")
	case <-time.After(2 * time.Second):
		fmt.Println("  ⚠️  Potential deadlock detected (timeout)")
	}
}

// GOOD: Proper lock ordering prevents deadlocks
func goodDeadlockPrevention() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// Both goroutines lock in the same order
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("  🧵 Goroutine 1: locked mu1")
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		fmt.Println("  🧵 Goroutine 1: locked mu2")
		mu2.Unlock()
		mu1.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("  🧵 Goroutine 2: locked mu1")
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		fmt.Println("  🧵 Goroutine 2: locked mu2")
		mu2.Unlock()
		mu1.Unlock()
	}()
	
	wg.Wait()
	fmt.Println("  ✅ No deadlock occurred")
}

// MUTEX PERFORMANCE: Understanding mutex overhead
func mutexPerformance() {
	fmt.Println("Understanding mutex performance...")
	
	// Test without mutex
	fmt.Println("  📊 Testing without mutex:")
	start := time.Now()
	testWithoutMutex()
	noMutexTime := time.Since(start)
	fmt.Printf("  ⏱️  Time without mutex: %v\n", noMutexTime)
	
	// Test with mutex
	fmt.Println("  📊 Testing with mutex:")
	start = time.Now()
	testWithMutex()
	withMutexTime := time.Since(start)
	fmt.Printf("  ⏱️  Time with mutex: %v\n", withMutexTime)
	
	// Calculate overhead
	overhead := float64(withMutexTime-noMutexTime) / float64(noMutexTime) * 100
	fmt.Printf("  📈 Mutex overhead: %.2f%%\n", overhead)
}

func testWithoutMutex() {
	var counter int
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	
	wg.Wait()
}

func testWithMutex() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	
	wg.Wait()
}

// WAITGROUP PATTERNS: Coordinating goroutines
func waitGroupPatterns() {
	fmt.Println("Understanding WaitGroup patterns...")
	
	// Basic WaitGroup
	fmt.Println("  📊 Basic WaitGroup:")
	basicWaitGroup()
	
	// WaitGroup with error handling
	fmt.Println("  📊 WaitGroup with error handling:")
	waitGroupWithErrors()
	
	// WaitGroup with timeout
	fmt.Println("  📊 WaitGroup with timeout:")
	waitGroupWithTimeout()
}

func basicWaitGroup() {
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("    🧵 Worker %d starting\n", id)
			time.Sleep(time.Duration(id+1) * 100 * time.Millisecond)
			fmt.Printf("    🧵 Worker %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("    ✅ All workers completed")
}

func waitGroupWithErrors() {
	var wg sync.WaitGroup
	errors := make(chan error, 3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Simulate work that might fail
			if id == 1 {
				errors <- fmt.Errorf("worker %d failed", id)
				return
			}
			
			fmt.Printf("    🧵 Worker %d completed successfully\n", id)
		}(i)
	}
	
	wg.Wait()
	close(errors)
	
	// Check for errors
	for err := range errors {
		fmt.Printf("    ❌ Error: %v\n", err)
	}
}

func waitGroupWithTimeout() {
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("    🧵 Worker %d starting\n", id)
			time.Sleep(time.Duration(id+1) * 500 * time.Millisecond)
			fmt.Printf("    🧵 Worker %d completed\n", id)
		}(i)
	}
	
	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("    ✅ All workers completed")
	case <-time.After(1 * time.Second):
		fmt.Println("    ⏰ Timeout: Some workers may still be running")
	}
}

// ONCE PATTERN: One-time initialization
func oncePattern() {
	fmt.Println("Understanding Once pattern...")
	
	var once sync.Once
	var wg sync.WaitGroup
	
	// Multiple goroutines try to initialize
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			once.Do(func() {
				fmt.Printf("  🧵 Goroutine %d: Initializing (only once!)\n", id)
				time.Sleep(100 * time.Millisecond)
			})
			
			fmt.Printf("  🧵 Goroutine %d: After initialization\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("  ✅ Once pattern completed")
}
