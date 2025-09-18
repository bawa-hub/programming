package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Exercise 1: Basic Mutex
// Create a counter protected by a mutex.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Mutex")
	fmt.Println("=======================")
	
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
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

// Cache type for Exercise 2
type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Exercise 2: RWMutex
// Implement a thread-safe cache using RWMutex.
func Exercise2() {
	fmt.Println("\nExercise 2: RWMutex")
	fmt.Println("===================")
	
	cache := &Cache{
		data: make(map[string]int),
	}
	
	// Set values
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	
	// Multiple readers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				value := cache.Get("key1")
				fmt.Printf("Reader %d: key1 = %d\n", id, value)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			cache.Set("key1", cache.Get("key1")+1)
			fmt.Printf("Writer: updated key1 to %d\n", cache.Get("key1"))
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Printf("Final value: %d\n", cache.Get("key1"))
}

// Exercise 3: WaitGroup
// Use WaitGroup to wait for multiple goroutines.
func Exercise3() {
	fmt.Println("\nExercise 3: WaitGroup")
	fmt.Println("=====================")
	
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

// Exercise 4: Once
// Implement a singleton pattern using Once.
func Exercise4() {
	fmt.Println("\nExercise 4: Once")
	fmt.Println("================")
	
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

// Exercise 5: Cond
// Use condition variables to coordinate goroutines.
func Exercise5() {
	fmt.Println("\nExercise 5: Cond")
	fmt.Println("================")
	
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

// Exercise 6: Atomic Operations
// Implement a counter using atomic operations.
func Exercise6() {
	fmt.Println("\nExercise 6: Atomic Operations")
	fmt.Println("============================")
	
	var counter int64
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
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

// Exercise 7: Concurrent Map
// Use sync.Map for thread-safe map operations.
func Exercise7() {
	fmt.Println("\nExercise 7: Concurrent Map")
	fmt.Println("=========================")
	
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

// Exercise 8: Object Pool
// Implement an object pool using sync.Pool.
func Exercise8() {
	fmt.Println("\nExercise 8: Object Pool")
	fmt.Println("=======================")
	
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

// Exercise 9: Deadlock Prevention
func Exercise9() {
	fmt.Println("\nExercise 9: Deadlock Prevention")
	fmt.Println("===============================")
	
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

// Exercise 10: Performance Comparison
func Exercise10() {
	fmt.Println("\nExercise 10: Performance Comparison")
	fmt.Println("===================================")
	
	const iterations = 100000
	
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

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Synchronization Exercises")
	fmt.Println("========================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}
