package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// AtomicCounter represents a thread-safe counter using atomic operations
type AtomicCounter struct {
	value int64
}

// NewAtomicCounter creates a new atomic counter
func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{value: 0}
}

// Increment increments the counter by 1
func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.value, 1)
}

// Decrement decrements the counter by 1
func (ac *AtomicCounter) Decrement() {
	atomic.AddInt64(&ac.value, -1)
}

// Add adds a value to the counter
func (ac *AtomicCounter) Add(delta int64) {
	atomic.AddInt64(&ac.value, delta)
}

// Get returns the current value
func (ac *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&ac.value)
}

// Set sets the counter to a specific value
func (ac *AtomicCounter) Set(value int64) {
	atomic.StoreInt64(&ac.value, value)
}

// CompareAndSwap compares and swaps the value
func (ac *AtomicCounter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&ac.value, old, new)
}

// String returns a string representation of the counter
func (ac *AtomicCounter) String() string {
	return fmt.Sprintf("AtomicCounter{value: %d}", ac.Get())
}

// DemonstrateAtomicCounter demonstrates the atomic counter
func DemonstrateAtomicCounter() {
	fmt.Println("=== Atomic Counter Demonstration ===")
	
	counter := NewAtomicCounter()
	var wg sync.WaitGroup
	
	// Multiple goroutines incrementing the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
				if j%200 == 0 {
					fmt.Printf("Worker %d: counter = %d\n", workerID, counter.Get())
				}
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d (should be 5000)\n", counter.Get())
}

// AtomicMap represents a thread-safe map using atomic operations
type AtomicMap struct {
	data map[string]int64
	mutex sync.RWMutex
}

// NewAtomicMap creates a new atomic map
func NewAtomicMap() *AtomicMap {
	return &AtomicMap{
		data: make(map[string]int64),
	}
}

// Increment increments the value for a key
func (am *AtomicMap) Increment(key string) {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	am.data[key]++
}

// Get returns the value for a key
func (am *AtomicMap) Get(key string) int64 {
	am.mutex.RLock()
	defer am.mutex.RUnlock()
	return am.data[key]
}

// Set sets the value for a key
func (am *AtomicMap) Set(key string, value int64) {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	am.data[key] = value
}

// GetAll returns a copy of all data
func (am *AtomicMap) GetAll() map[string]int64 {
	am.mutex.RLock()
	defer am.mutex.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range am.data {
		result[k] = v
	}
	return result
}

// DemonstrateAtomicMap demonstrates the atomic map
func DemonstrateAtomicMap() {
	fmt.Println("\n=== Atomic Map Demonstration ===")
	
	atomicMap := NewAtomicMap()
	var wg sync.WaitGroup
	
	// Multiple goroutines updating different keys
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	
	for i, key := range keys {
		wg.Add(1)
		go func(workerID int, k string) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomicMap.Increment(k)
				if j%200 == 0 {
					fmt.Printf("Worker %d: %s = %d\n", workerID, k, atomicMap.Get(k))
				}
			}
		}(i, key)
	}
	
	wg.Wait()
	
	fmt.Println("Final map values:")
	for key, value := range atomicMap.GetAll() {
		fmt.Printf("  %s: %d\n", key, value)
	}
}

// AtomicBool represents a thread-safe boolean using atomic operations
type AtomicBool struct {
	value int32
}

// NewAtomicBool creates a new atomic boolean
func NewAtomicBool() *AtomicBool {
	return &AtomicBool{value: 0}
}

// Set sets the boolean value
func (ab *AtomicBool) Set(value bool) {
	if value {
		atomic.StoreInt32(&ab.value, 1)
	} else {
		atomic.StoreInt32(&ab.value, 0)
	}
}

// Get returns the boolean value
func (ab *AtomicBool) Get() bool {
	return atomic.LoadInt32(&ab.value) == 1
}

// Toggle toggles the boolean value
func (ab *AtomicBool) Toggle() bool {
	for {
		old := atomic.LoadInt32(&ab.value)
		new := 1 - old
		if atomic.CompareAndSwapInt32(&ab.value, old, new) {
			return new == 1
		}
	}
}

// String returns a string representation of the atomic boolean
func (ab *AtomicBool) String() string {
	return fmt.Sprintf("AtomicBool{value: %v}", ab.Get())
}

// DemonstrateAtomicBool demonstrates the atomic boolean
func DemonstrateAtomicBool() {
	fmt.Println("\n=== Atomic Boolean Demonstration ===")
	
	atomicBool := NewAtomicBool()
	var wg sync.WaitGroup
	
	// Multiple goroutines toggling the boolean
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				value := atomicBool.Toggle()
				fmt.Printf("Worker %d: toggled to %v\n", workerID, value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final boolean value: %v\n", atomicBool.Get())
}

// Performance comparison between mutex and atomic operations
func DemonstratePerformanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")
	
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
	
	// Test with atomic operations
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex approach: %v (counter: %d)\n", mutexTime, counter1)
	fmt.Printf("Atomic approach: %v (counter: %d)\n", atomicTime, counter2)
	fmt.Printf("Atomic speedup: %.2fx\n", float64(mutexTime)/float64(atomicTime))
}

// AtomicCounterWithStats represents a counter with statistics
type AtomicCounterWithStats struct {
	value     int64
	increment int64
	decrement int64
}

// NewAtomicCounterWithStats creates a new atomic counter with stats
func NewAtomicCounterWithStats() *AtomicCounterWithStats {
	return &AtomicCounterWithStats{}
}

// Increment increments the counter and tracks stats
func (acs *AtomicCounterWithStats) Increment() {
	atomic.AddInt64(&acs.value, 1)
	atomic.AddInt64(&acs.increment, 1)
}

// Decrement decrements the counter and tracks stats
func (acs *AtomicCounterWithStats) Decrement() {
	atomic.AddInt64(&acs.value, -1)
	atomic.AddInt64(&acs.decrement, 1)
}

// Get returns the current value
func (acs *AtomicCounterWithStats) Get() int64 {
	return atomic.LoadInt64(&acs.value)
}

// GetStats returns the statistics
func (acs *AtomicCounterWithStats) GetStats() (int64, int64, int64) {
	return atomic.LoadInt64(&acs.value), 
		   atomic.LoadInt64(&acs.increment), 
		   atomic.LoadInt64(&acs.decrement)
}

// String returns a string representation
func (acs *AtomicCounterWithStats) String() string {
	value, inc, dec := acs.GetStats()
	return fmt.Sprintf("AtomicCounterWithStats{value: %d, increments: %d, decrements: %d}", 
		value, inc, dec)
}

// DemonstrateAtomicCounterWithStats demonstrates the counter with stats
func DemonstrateAtomicCounterWithStats() {
	fmt.Println("\n=== Atomic Counter with Stats Demonstration ===")
	
	counter := NewAtomicCounterWithStats()
	var wg sync.WaitGroup
	
	// Some goroutines increment
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("Worker %d: completed increments\n", workerID)
		}(i)
	}
	
	// Some goroutines decrement
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				counter.Decrement()
			}
			fmt.Printf("Worker %d: completed decrements\n", workerID)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter: %s\n", counter)
}
