package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	ex "gsync/exercies"
	pat "gsync/patterns"
)

// Example 1: Basic Mutex
func basicMutex() {
	fmt.Println("1. Basic Mutex")
	fmt.Println("==============")
	
	var mu sync.Mutex
	var counter int
	
	// Start multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

// Example 2: RWMutex (Read-Write Mutex)
func rwMutex() {
	fmt.Println("\n2. RWMutex (Read-Write Mutex)")
	fmt.Println("=============================")
	
	var rwmu sync.RWMutex
	data := make(map[string]int)
	
	// Initialize data
	data["key1"] = 1
	data["key2"] = 2
	data["key3"] = 3
	
	// Multiple readers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				rwmu.RLock()
				value := data["key1"]
				fmt.Printf("Reader %d: key1 = %d\n", id, value)
				rwmu.RUnlock()
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			rwmu.Lock()
			data["key1"]++
			fmt.Printf("Writer: updated key1 to %d\n", data["key1"])
			rwmu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Printf("Final value: %d\n", data["key1"])
}

// Example 3: WaitGroup
func waitGroup() {
	fmt.Println("\n3. WaitGroup")
	fmt.Println("============")
	
	var wg sync.WaitGroup
	results := make(chan int, 5)
	
	// Start workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate work
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			result := id * id
			results <- result
			fmt.Printf("Worker %d completed with result %d\n", id, result)
		}(i)
	}
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("  Result: %d\n", result)
	}
}

// Example 4: Once (One-Time Execution)
func onceExample() {
	fmt.Println("\n4. Once (One-Time Execution)")
	fmt.Println("=============================")
	
	var once sync.Once
	var instance *ex.Singleton
	var wg sync.WaitGroup
	
	// Multiple goroutines trying to create singleton
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				instance = &ex.Singleton{ID: id, Created: time.Now()}
				fmt.Printf("Goroutine %d created singleton\n", id)
			})
			fmt.Printf("Goroutine %d got instance: %+v\n", id, instance)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final instance: %+v\n", instance)
}

// Example 5: Cond (Condition Variables)
func condExample() {
	fmt.Println("\n5. Cond (Condition Variables)")
	fmt.Println("=============================")
	
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var ready bool
	var wg sync.WaitGroup
	
	// Waiters
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			for !ready {
				fmt.Printf("Waiter %d: waiting for condition\n", id)
				cond.Wait()
			}
			fmt.Printf("Waiter %d: condition met!\n", id)
			mu.Unlock()
		}(i)
	}
	
	// Signaler
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		mu.Lock()
		ready = true
		fmt.Println("Signaler: condition is ready, broadcasting...")
		cond.Broadcast()
		mu.Unlock()
	}()
	
	wg.Wait()
}

// Example 6: Atomic Operations
func atomicExample() {
	fmt.Println("\n6. Atomic Operations")
	fmt.Println("===================")
	
	var counter int64
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
	
	// Compare and swap
	oldValue := atomic.LoadInt64(&counter)
	newValue := oldValue + 100
	if atomic.CompareAndSwapInt64(&counter, oldValue, newValue) {
		fmt.Printf("CAS successful: %d -> %d\n", oldValue, newValue)
	} else {
		fmt.Println("CAS failed")
	}
}

// Example 7: Concurrent Map
func concurrentMap() {
	fmt.Println("\n7. Concurrent Map")
	fmt.Println("=================")
	
	var m sync.Map
	var wg sync.WaitGroup
	
	// Store values
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			m.Store(key, value)
			fmt.Printf("Stored: %s = %s\n", key, value)
		}(i)
	}
	
	// Load values
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if value, ok := m.Load(key); ok {
				fmt.Printf("Loaded: %s = %s\n", key, value)
			} else {
				fmt.Printf("Key %s not found\n", key)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Range over map
	fmt.Println("All key-value pairs:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %s = %s\n", key, value)
		return true
	})
}

// Example 8: Object Pool
func objectPool() {
	fmt.Println("\n8. Object Pool")
	fmt.Println("==============")
	
	var pool = sync.Pool{
		New: func() interface{} {
			return &ex.Buffer{ID: time.Now().UnixNano()}
		},
	}
	
	var wg sync.WaitGroup
	
	// Get objects from pool
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Get buffer from pool
			buf := pool.Get().(*ex.Buffer)
			fmt.Printf("Worker %d got buffer %d\n", id, buf.ID)
			
			// Use buffer
			buf.WriteString(fmt.Sprintf("Data from worker %d", id))
			time.Sleep(100 * time.Millisecond)
			
			// Put buffer back to pool
			buf.Reset()
			pool.Put(buf)
			fmt.Printf("Worker %d returned buffer %d to pool\n", id, buf.ID)
		}(i)
	}
	
	wg.Wait()
}




// Example 9: Performance Comparison
func performanceComparison() {
	fmt.Println("\n9. Performance Comparison")
	fmt.Println("=========================")
	
	const iterations = 1000000
	
	// Mutex performance
	start := time.Now()
	var mu sync.Mutex
	var counter1 int
	for i := 0; i < iterations; i++ {
		mu.Lock()
		counter1++
		mu.Unlock()
	}
	mutexTime := time.Since(start)
	
	// Atomic performance
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v\n", mutexTime)
	fmt.Printf("Atomic time: %v\n", atomicTime)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}

// Example 10: Deadlock Prevention
func deadlockPrevention() {
	fmt.Println("\n10. Deadlock Prevention")
	fmt.Println("=======================")
	
	var mu1, mu2 sync.Mutex
	
	// Function that acquires locks in order
	acquireLocks := func(id int, mu1, mu2 *sync.Mutex) {
		mu1.Lock()
		fmt.Printf("Goroutine %d: acquired lock1\n", id)
		time.Sleep(10 * time.Millisecond)
		
		mu2.Lock()
		fmt.Printf("Goroutine %d: acquired lock2\n", id)
		time.Sleep(10 * time.Millisecond)
		
		mu2.Unlock()
		fmt.Printf("Goroutine %d: released lock2\n", id)
		mu1.Unlock()
		fmt.Printf("Goroutine %d: released lock1\n", id)
	}
	
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			acquireLocks(id, &mu1, &mu2)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All goroutines completed without deadlock")
}

// Example 11: Race Condition Detection
func raceConditionDetection() {
	fmt.Println("\n11. Race Condition Detection")
	fmt.Println("============================")
	
	// This example demonstrates a race condition
	// Run with: go run -race . to see race detection
	var counter int
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // Race condition!
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter value: %d (should be 3000, but may not be due to race condition)\n", counter)
	
	// Fixed version with mutex
	var mu sync.Mutex
	var safeCounter int
	wg = sync.WaitGroup{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				safeCounter++
				mu.Unlock()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Safe counter value: %d (correct)\n", safeCounter)
}

// Example 12: Common Pitfalls
func commonPitfalls() {
	fmt.Println("\n12. Common Pitfalls")
	fmt.Println("===================")
	
	// Pitfall 1: Forgetting to unlock
	fmt.Println("Pitfall 1: Forgetting to unlock")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// mu.Lock()")
	fmt.Println("// // Do work")
	fmt.Println("// // Forgot to unlock!")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// mu.Lock()")
	fmt.Println("// defer mu.Unlock()")
	fmt.Println("// // Do work")
	
	// Pitfall 2: Double unlock
	fmt.Println("\nPitfall 2: Double unlock")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// mu.Lock()")
	fmt.Println("// defer mu.Unlock()")
	fmt.Println("// mu.Unlock() // Panic!")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// mu.Lock()")
	fmt.Println("// defer mu.Unlock()")
	fmt.Println("// // Do work")
	
	// Pitfall 3: WaitGroup negative counter
	fmt.Println("\nPitfall 3: WaitGroup negative counter")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// var wg sync.WaitGroup")
	fmt.Println("// wg.Done() // Panic! Counter becomes negative")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// var wg sync.WaitGroup")
	fmt.Println("// wg.Add(1)")
	fmt.Println("// go func() { defer wg.Done(); /* work */ }()")
	
	// Pitfall 4: Using sync.Map incorrectly
	fmt.Println("\nPitfall 4: Using sync.Map incorrectly")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// var m sync.Map")
	fmt.Println("// m.Store(\"key\", \"value\")")
	fmt.Println("// if v, ok := m.Load(\"key\"); ok {")
	fmt.Println("//     v = v.(string) + \" modified\" // Modifying value")
	fmt.Println("// }")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// var m sync.Map")
	fmt.Println("// m.Store(\"key\", \"value\")")
	fmt.Println("// if v, ok := m.Load(\"key\"); ok {")
	fmt.Println("//     newValue := v.(string) + \" modified\"")
	fmt.Println("//     m.Store(\"key\", newValue)")
	fmt.Println("// }")
}

// Utility function to show sync info
func showSyncInfo() {
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}


func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "basic":
		runBasicExamples()
	case "exercises":
		RunAllExercises()
	case "advanced":
		RunAdvancedPatterns()
	case "all":
		runBasicExamples()
		fmt.Println("\n" + "==================================================")
		RunAllExercises()
		fmt.Println("\n" + "==================================================")
		RunAdvancedPatterns()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func showUsage() {
	fmt.Println("🚀 Synchronization Primitives - Usage")
	fmt.Println("=====================================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic synchronization examples")
	fmt.Println("  exercises - Run all exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  all       - Run everything")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println("  go run . all")
}

func runBasicExamples() {
	fmt.Println("🚀 Synchronization Primitives Examples")
	fmt.Println("======================================")

	// Example 1: Basic Mutex
	basicMutex()

	// Example 2: RWMutex (Read-Write Mutex)
	rwMutex()

	// Example 3: WaitGroup
	waitGroup()

	// Example 4: Once (One-Time Execution)
	onceExample()

	// Example 5: Cond (Condition Variables)
	condExample()

	// Example 6: Atomic Operations
	atomicExample()

	// Example 7: Concurrent Map
	concurrentMap()

	// Example 8: Object Pool
	objectPool()

	// Example 9: Performance Comparison
	performanceComparison()

	// Example 10: Deadlock Prevention
	deadlockPrevention()

	// Example 11: Race Condition Detection
	raceConditionDetection()

	// Example 12: Common Pitfalls
	commonPitfalls()
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Synchronization Exercises")
	fmt.Println("========================================")
	
	ex.Exercise1()
	ex.Exercise2()
	ex.Exercise3()
	ex.Exercise4()
	ex.Exercise5()
	ex.Exercise6()
	ex.Exercise7()
	ex.Exercise8()
	ex.Exercise9()
	ex.Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Synchronization Patterns")
	fmt.Println("====================================")
	
	// Pattern 1: Thread-Safe Counter with Metrics
	fmt.Println("\n1. Thread-Safe Counter with Metrics:")
	counter := pat.NewSafeCounter()
	
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				counter.Increment(fmt.Sprintf("key%d", id))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Counters: %v\n", counter.GetAllCounters())
	fmt.Printf("Metrics: %v\n", counter.GetMetrics())
	
	// Pattern 2: Priority RWMutex
	fmt.Println("\n2. Priority RWMutex:")
	prw := pat.NewPriorityRWMutex()
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			prw.RLock()
			fmt.Printf("Reader %d: reading\n", id)
			time.Sleep(100 * time.Millisecond)
			prw.RUnlock()
		}(i)
	}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		prw.Lock()
		fmt.Println("Writer: writing")
		time.Sleep(100 * time.Millisecond)
		prw.Unlock()
	}()
	
	wg.Wait()
	readers, writers, readerWait, writerWait := prw.GetStats()
	fmt.Printf("Stats: readers=%d, writers=%d, readerWait=%d, writerWait=%d\n", readers, writers, readerWait, writerWait)
	
	// Pattern 3: WaitGroup with Timeout
	fmt.Println("\n3. WaitGroup with Timeout:")
	twg := pat.NewTimeoutWaitGroup(500 * time.Millisecond)
	
	for i := 0; i < 3; i++ {
		twg.Add(1)
		go func(id int) {
			defer twg.Done()
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Worker %d completed\n", id)
		}(i)
	}
	
	if err := twg.Wait(); err != nil {
		fmt.Printf("WaitGroup error: %v\n", err)
	} else {
		fmt.Println("All workers completed")
	}
	
	// Pattern 4: Once with Error Handling
	fmt.Println("\n4. Once with Error Handling:")
	so := &pat.SafeOnce{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := so.Do(func() error {
				fmt.Printf("Initializing from goroutine %d\n", id)
				if id == 1 {
					return fmt.Errorf("initialization failed")
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Goroutine %d: error: %v\n", id, err)
			} else {
				fmt.Printf("Goroutine %d: success\n", id)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Pattern 5: Condition Variable with Timeout
	fmt.Println("\n5. Condition Variable with Timeout:")
	tc := pat.NewTimeoutCond()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		if tc.WaitWithTimeout(1 * time.Second) {
			fmt.Println("Condition met within timeout")
		} else {
			fmt.Println("Timeout waiting for condition")
		}
	}()
	
	time.Sleep(500 * time.Millisecond)
	tc.Signal()
	wg.Wait()
	
	// Pattern 6: Atomic Counter with Statistics
	fmt.Println("\n6. Atomic Counter with Statistics:")
	ac := &pat.AtomicCounter{}
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ac.Increment()
			}
		}(i)
	}
	
	wg.Wait()
	value, increments, decrements, resets := ac.GetStats()
	fmt.Printf("Counter stats: value=%d, increments=%d, decrements=%d, resets=%d\n", value, increments, decrements, resets)
	
	// Pattern 7: Concurrent Map with Statistics
	fmt.Println("\n7. Concurrent Map with Statistics:")
	sm := pat.NewStatsMap()
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			sm.Store(key, fmt.Sprintf("value%d", id))
			sm.Load(key)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Map stats: %v\n", sm.GetStats())
	
	// Pattern 8: Object Pool with Statistics
	fmt.Println("\n8. Object Pool with Statistics:")
	sp := pat.NewStatsPool(func() interface{} {
		return &ex.Buffer{ID: time.Now().UnixNano()}
	})
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			buf := sp.Get().(*ex.Buffer)
			time.Sleep(10 * time.Millisecond)
			sp.Put(buf)
		}(i)
	}
	
	wg.Wait()
	created, reused, returned := sp.GetStats()
	fmt.Printf("Pool stats: created=%d, reused=%d, returned=%d\n", created, reused, returned)
	
	// Pattern 9: Barrier Synchronization
	fmt.Println("\n9. Barrier Synchronization:")
	barrier := pat.NewBarrier(3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: before barrier\n", id)
			phase := barrier.Wait()
			fmt.Printf("Goroutine %d: after barrier phase %d\n", id, phase)
		}(i)
	}
	
	wg.Wait()
	
	// Pattern 10: Semaphore
	fmt.Println("\n10. Semaphore:")
	sem := pat.NewSemaphore(2)
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			fmt.Printf("Goroutine %d: acquired semaphore\n", id)
			time.Sleep(100 * time.Millisecond)
			sem.Release()
			fmt.Printf("Goroutine %d: released semaphore\n", id)
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}