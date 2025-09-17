package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Exercise 1: Implement Basic Memory Pool
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Basic Memory Pool")
	fmt.Println("=====================================")
	
	// TODO: Implement a basic memory pool
	// 1. Create a memory pool for byte slices
	// 2. Implement Get and Put methods
	// 3. Test with concurrent access
	
	pool := NewExerciseMemoryPool(1024)
	
	// Test basic operations
	buf1 := pool.Get()
	buf2 := pool.Get()
	
	fmt.Printf("  Exercise 1: Allocated buffers: %d, %d bytes\n", len(buf1), len(buf2))
	
	// Return buffers
	pool.Put(buf1)
	pool.Put(buf2)
	
	// Get buffer again (should reuse)
	buf3 := pool.Get()
	fmt.Printf("  Exercise 1: Reused buffer: %d bytes\n", len(buf3))
	
	fmt.Println("  Exercise 1: Basic memory pool completed")
}

type ExerciseMemoryPool struct {
	pool sync.Pool
	size int
}

func NewExerciseMemoryPool(size int) *ExerciseMemoryPool {
	return &ExerciseMemoryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
		size: size,
	}
}

func (emp *ExerciseMemoryPool) Get() []byte {
	return emp.pool.Get().([]byte)
}

func (emp *ExerciseMemoryPool) Put(buf []byte) {
	if len(buf) == emp.size {
		emp.pool.Put(buf)
	}
}

// Exercise 2: Implement String Optimization
func Exercise2() {
	fmt.Println("\nExercise 2: Implement String Optimization")
	fmt.Println("=========================================")
	
	// TODO: Implement string optimization
	// 1. Compare string concatenation vs string builder
	// 2. Measure performance difference
	// 3. Use pre-allocation for better performance
	
	// Bad: String concatenation
	start := time.Now()
	var badString string
	for i := 0; i < 1000; i++ {
		badString += "a"
	}
	badTime := time.Since(start)
	
	// Good: String builder with pre-allocation
	start = time.Now()
	var goodString strings.Builder
	goodString.Grow(1000) // Pre-allocate capacity
	for i := 0; i < 1000; i++ {
		goodString.WriteString("a")
	}
	goodTime := time.Since(start)
	
	fmt.Printf("  Exercise 2: Bad string: len=%d, time=%v\n", len(badString), badTime)
	fmt.Printf("  Exercise 2: Good string: len=%d, time=%v\n", goodString.Len(), goodTime)
	fmt.Printf("  Exercise 2: Speedup: %.2fx\n", float64(badTime)/float64(goodTime))
	
	fmt.Println("  Exercise 2: String optimization completed")
}

// Exercise 3: Implement Slice Pre-allocation
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Slice Pre-allocation")
	fmt.Println("=========================================")
	
	// TODO: Implement slice pre-allocation
	// 1. Compare growing slice vs pre-allocated slice
	// 2. Measure performance difference
	// 3. Use appropriate capacity for better performance
	
	// Bad: Growing slice
	start := time.Now()
	var badSlice []int
	for i := 0; i < 1000; i++ {
		badSlice = append(badSlice, i)
	}
	badTime := time.Since(start)
	
	// Good: Pre-allocated slice
	start = time.Now()
	goodSlice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		goodSlice = append(goodSlice, i)
	}
	goodTime := time.Since(start)
	
	fmt.Printf("  Exercise 3: Bad slice: len=%d, cap=%d, time=%v\n", len(badSlice), cap(badSlice), badTime)
	fmt.Printf("  Exercise 3: Good slice: len=%d, cap=%d, time=%v\n", len(goodSlice), cap(goodSlice), goodTime)
	fmt.Printf("  Exercise 3: Speedup: %.2fx\n", float64(badTime)/float64(goodTime))
	
	fmt.Println("  Exercise 3: Slice pre-allocation completed")
}

// Exercise 4: Implement Map Optimization
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Map Optimization")
	fmt.Println("=====================================")
	
	// TODO: Implement map optimization
	// 1. Compare growing map vs pre-sized map
	// 2. Measure performance difference
	// 3. Use appropriate initial size for better performance
	
	// Bad: Growing map
	start := time.Now()
	badMap := make(map[int]string)
	for i := 0; i < 1000; i++ {
		badMap[i] = "value"
	}
	badTime := time.Since(start)
	
	// Good: Pre-sized map
	start = time.Now()
	goodMap := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		goodMap[i] = "value"
	}
	goodTime := time.Since(start)
	
	fmt.Printf("  Exercise 4: Bad map: size=%d, time=%v\n", len(badMap), badTime)
	fmt.Printf("  Exercise 4: Good map: size=%d, time=%v\n", len(goodMap), goodTime)
	fmt.Printf("  Exercise 4: Speedup: %.2fx\n", float64(badTime)/float64(goodTime))
	
	fmt.Println("  Exercise 4: Map optimization completed")
}

// Exercise 5: Implement Memory Leak Detection
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Memory Leak Detection")
	fmt.Println("==========================================")
	
	// TODO: Implement memory leak detection
	// 1. Allocate memory and track usage
	// 2. Clear references and force GC
	// 3. Check if memory was properly freed
	
	var m1, m2, m3 runtime.MemStats
	
	// Get initial stats
	runtime.ReadMemStats(&m1)
	
	// Allocate memory
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024)
	}
	
	// Get stats after allocation
	runtime.ReadMemStats(&m2)
	
	// Clear references
	data = nil
	
	// Force GC
	runtime.GC()
	
	// Get stats after GC
	runtime.ReadMemStats(&m3)
	
	fmt.Printf("  Exercise 5: Before allocation: %d bytes\n", m1.Alloc)
	fmt.Printf("  Exercise 5: After allocation: %d bytes\n", m2.Alloc)
	fmt.Printf("  Exercise 5: After GC: %d bytes\n", m3.Alloc)
	
	if m3.Alloc < m2.Alloc {
		fmt.Println("  Exercise 5: No memory leak detected")
	} else {
		fmt.Println("  Exercise 5: Potential memory leak detected")
	}
	
	fmt.Println("  Exercise 5: Memory leak detection completed")
}

// Exercise 6: Implement GC Pressure Analysis
func Exercise6() {
	fmt.Println("\nExercise 6: Implement GC Pressure Analysis")
	fmt.Println("=========================================")
	
	// TODO: Implement GC pressure analysis
	// 1. Calculate GC pressure metrics
	// 2. Analyze heap utilization
	// 3. Monitor GC frequency and pause times
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Calculate GC pressure
	gcPressure := float64(m.PauseTotalNs) / float64(time.Since(time.Unix(0, int64(m.LastGC))).Nanoseconds())
	heapUtilization := float64(m.HeapAlloc) / float64(m.HeapSys)
	
	fmt.Printf("  Exercise 6: GC Pressure: %.2f%%\n", gcPressure*100)
	fmt.Printf("  Exercise 6: Heap Utilization: %.2f%%\n", heapUtilization*100)
	fmt.Printf("  Exercise 6: GC Cycles: %d\n", m.NumGC)
	fmt.Printf("  Exercise 6: GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
	
	fmt.Println("  Exercise 6: GC pressure analysis completed")
}

// Exercise 7: Implement Advanced Memory Pool
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Advanced Memory Pool")
	fmt.Println("=========================================")
	
	// TODO: Implement advanced memory pool
	// 1. Create pool for different buffer sizes
	// 2. Use sync.Pool for each size
	// 3. Implement thread-safe Get and Put methods
	
	pool := NewExerciseAdvancedMemoryPool()
	
	// Test different buffer sizes
	buf1 := pool.Get(64)
	buf2 := pool.Get(128)
	buf3 := pool.Get(256)
	
	fmt.Printf("  Exercise 7: Allocated buffers: %d, %d, %d bytes\n", len(buf1), len(buf2), len(buf3))
	
	// Return buffers
	pool.Put(buf1)
	pool.Put(buf2)
	pool.Put(buf3)
	
	// Get buffers again (should reuse)
	buf4 := pool.Get(64)
	buf5 := pool.Get(128)
	
	fmt.Printf("  Exercise 7: Reused buffers: %d, %d bytes\n", len(buf4), len(buf5))
	
	fmt.Println("  Exercise 7: Advanced memory pool completed")
}

type ExerciseAdvancedMemoryPool struct {
	pools map[int]*sync.Pool
	mu    sync.RWMutex
}

func NewExerciseAdvancedMemoryPool() *ExerciseAdvancedMemoryPool {
	return &ExerciseAdvancedMemoryPool{
		pools: make(map[int]*sync.Pool),
	}
}

func (eamp *ExerciseAdvancedMemoryPool) Get(size int) []byte {
	eamp.mu.RLock()
	pool, exists := eamp.pools[size]
	eamp.mu.RUnlock()
	
	if !exists {
		eamp.mu.Lock()
		pool, exists = eamp.pools[size]
		if !exists {
			pool = &sync.Pool{
				New: func() interface{} {
					return make([]byte, size)
				},
			}
			eamp.pools[size] = pool
		}
		eamp.mu.Unlock()
	}
	
	return pool.Get().([]byte)
}

func (eamp *ExerciseAdvancedMemoryPool) Put(buf []byte) {
	size := len(buf)
	eamp.mu.RLock()
	pool, exists := eamp.pools[size]
	eamp.mu.RUnlock()
	
	if exists {
		pool.Put(buf)
	}
}

// Exercise 8: Implement Object Reuse Pattern
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Object Reuse Pattern")
	fmt.Println("=========================================")
	
	// TODO: Implement object reuse pattern
	// 1. Create reusable objects with pool
	// 2. Implement Get and Put methods
	// 3. Test object reuse and thread safety
	
	pool := NewExerciseObjectPool(5)
	
	// Get objects from pool
	obj1 := pool.Get()
	obj2 := pool.Get()
	
	fmt.Printf("  Exercise 8: Got objects: %d, %d\n", obj1.ID, obj2.ID)
	
	// Use objects
	obj1.Data[0] = 1
	obj2.Data[0] = 2
	
	// Return objects to pool
	pool.Put(obj1)
	pool.Put(obj2)
	
	// Get objects again (should reuse)
	obj3 := pool.Get()
	obj4 := pool.Get()
	
	fmt.Printf("  Exercise 8: Reused objects: %d, %d\n", obj3.ID, obj4.ID)
	
	fmt.Println("  Exercise 8: Object reuse pattern completed")
}

type ExerciseReusableObject struct {
	ID        int
	Data      []byte
	Timestamp time.Time
	inUse     bool
	mu        sync.Mutex
}

type ExerciseObjectPool struct {
	objects []*ExerciseReusableObject
	free    chan *ExerciseReusableObject
	mu      sync.Mutex
}

func NewExerciseObjectPool(size int) *ExerciseObjectPool {
	pool := &ExerciseObjectPool{
		objects: make([]*ExerciseReusableObject, size),
		free:    make(chan *ExerciseReusableObject, size),
	}
	
	// Initialize objects
	for i := 0; i < size; i++ {
		obj := &ExerciseReusableObject{
			ID:    i,
			Data:  make([]byte, 1024),
			inUse: false,
		}
		pool.objects[i] = obj
		pool.free <- obj
	}
	
	return pool
}

func (eop *ExerciseObjectPool) Get() *ExerciseReusableObject {
	obj := <-eop.free
	obj.mu.Lock()
	obj.inUse = true
	obj.Timestamp = time.Now()
	obj.mu.Unlock()
	return obj
}

func (eop *ExerciseObjectPool) Put(obj *ExerciseReusableObject) {
	obj.mu.Lock()
	obj.inUse = false
	obj.mu.Unlock()
	eop.free <- obj
}

// Exercise 9: Implement Memory Monitoring
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Memory Monitoring")
	fmt.Println("======================================")
	
	// TODO: Implement memory monitoring
	// 1. Create memory monitor with threshold
	// 2. Monitor memory usage periodically
	// 3. Trigger callback when threshold exceeded
	
	monitor := NewExerciseMemoryMonitor(1024*1024, func() {
		fmt.Println("  Exercise 9: Memory threshold exceeded!")
	})
	
	// Simulate memory usage
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024)
		time.Sleep(1 * time.Millisecond)
	}
	
	// Stop monitor
	monitor.Stop()
	
	fmt.Println("  Exercise 9: Memory monitoring completed")
}

type ExerciseMemoryMonitor struct {
	threshold uint64
	callback  func()
	mu        sync.Mutex
	quit      chan bool
}

func NewExerciseMemoryMonitor(threshold uint64, callback func()) *ExerciseMemoryMonitor {
	emm := &ExerciseMemoryMonitor{
		threshold: threshold,
		callback:  callback,
		quit:      make(chan bool),
	}
	
	go emm.monitor()
	
	return emm
}

func (emm *ExerciseMemoryMonitor) monitor() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			
			if m.Alloc > emm.threshold {
				emm.mu.Lock()
				if emm.callback != nil {
					emm.callback()
				}
				emm.mu.Unlock()
			}
		case <-emm.quit:
			return
		}
	}
}

func (emm *ExerciseMemoryMonitor) Stop() {
	close(emm.quit)
}

// Exercise 10: Implement Performance Comparison
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Performance Comparison")
	fmt.Println("============================================")
	
	// TODO: Implement performance comparison
	// 1. Compare different memory allocation strategies
	// 2. Measure performance metrics
	// 3. Analyze results and provide recommendations
	
	// Test 1: Direct allocation
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1024)
	}
	directTime := time.Since(start)
	
	// Test 2: Memory pool
	pool := NewExerciseMemoryPool(1024)
	start = time.Now()
	for i := 0; i < 10000; i++ {
		buf := pool.Get()
		_ = buf
		pool.Put(buf)
	}
	poolTime := time.Since(start)
	
	// Test 3: Pre-allocated slice
	start = time.Now()
	slice := make([]byte, 0, 10000*1024)
	for i := 0; i < 10000; i++ {
		slice = append(slice, make([]byte, 1024)...)
	}
	sliceTime := time.Since(start)
	
	fmt.Printf("  Exercise 10: Direct allocation: %v\n", directTime)
	fmt.Printf("  Exercise 10: Memory pool: %v\n", poolTime)
	fmt.Printf("  Exercise 10: Pre-allocated slice: %v\n", sliceTime)
	
	// Calculate speedups
	directVsPool := float64(directTime) / float64(poolTime)
	directVsSlice := float64(directTime) / float64(sliceTime)
	
	fmt.Printf("  Exercise 10: Pool vs Direct: %.2fx\n", directVsPool)
	fmt.Printf("  Exercise 10: Slice vs Direct: %.2fx\n", directVsSlice)
	
	fmt.Println("  Exercise 10: Performance comparison completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Memory Management Exercises")
	fmt.Println("==============================")
	
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
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
