package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
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
	var instance *Singleton
	var wg sync.WaitGroup
	
	// Multiple goroutines trying to create singleton
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				instance = &Singleton{ID: id, Created: time.Now()}
				fmt.Printf("Goroutine %d created singleton\n", id)
			})
			fmt.Printf("Goroutine %d got instance: %+v\n", id, instance)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final instance: %+v\n", instance)
}

type Singleton struct {
	ID      int
	Created time.Time
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
			return &Buffer{ID: time.Now().UnixNano()}
		},
	}
	
	var wg sync.WaitGroup
	
	// Get objects from pool
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Get buffer from pool
			buf := pool.Get().(*Buffer)
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

type Buffer struct {
	ID   int64
	Data string
}

func (b *Buffer) WriteString(s string) {
	b.Data += s
}

func (b *Buffer) Reset() {
	b.Data = ""
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
