package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Example 1: Basic Memory Statistics
func basicMemoryStatistics() {
	fmt.Println("\n1. Basic Memory Statistics")
	fmt.Println("=========================")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  Allocated: %d bytes\n", m.Alloc)
	fmt.Printf("  Total Allocated: %d bytes\n", m.TotalAlloc)
	fmt.Printf("  System Memory: %d bytes\n", m.Sys)
	fmt.Printf("  GC Cycles: %d\n", m.NumGC)
	fmt.Printf("  GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
	fmt.Printf("  Heap Size: %d bytes\n", m.HeapSys)
	fmt.Printf("  Heap Allocated: %d bytes\n", m.HeapAlloc)
	fmt.Printf("  Heap Objects: %d\n", m.HeapObjects)
	
	fmt.Println("  Basic memory statistics completed")
}

// Example 2: GC Tuning
func gcTuning() {
	fmt.Println("\n2. GC Tuning")
	fmt.Println("============")
	
	// Set GC percentage
	debug.SetGCPercent(200)
	fmt.Println("  GC percentage set to 200%")
	
	// Force immediate GC
	runtime.GC()
	fmt.Println("  Forced GC cycle")
	
	// Get current GC percentage
	currentPercent := debug.SetGCPercent(-1)
	fmt.Printf("  Current GC percentage: %d\n", currentPercent)
	
	// Reset to default
	debug.SetGCPercent(100)
	fmt.Println("  GC percentage reset to default")
	
	fmt.Println("  GC tuning completed")
}

// Example 3: Memory Allocation Patterns
func memoryAllocationPatterns() {
	fmt.Println("\n3. Memory Allocation Patterns")
	fmt.Println("=============================")
	
	var m1, m2 runtime.MemStats
	
	// Test small allocations
	runtime.ReadMemStats(&m1)
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1)
	}
	runtime.ReadMemStats(&m2)
	smallAlloc := m2.Alloc - m1.Alloc
	
	// Test large allocations
	runtime.ReadMemStats(&m1)
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024)
	}
	runtime.ReadMemStats(&m2)
	largeAlloc := m2.Alloc - m1.Alloc
	
	fmt.Printf("  Small allocations: %d bytes\n", smallAlloc)
	fmt.Printf("  Large allocations: %d bytes\n", largeAlloc)
	
	fmt.Println("  Memory allocation patterns completed")
}

// Example 4: Stack vs Heap Allocation
func stackVsHeapAllocation() {
	fmt.Println("\n4. Stack vs Heap Allocation")
	fmt.Println("===========================")
	
	// Stack allocation (local variable)
	localVar := 42
	fmt.Printf("  Local variable: %d\n", localVar)
	
	// Heap allocation (escapes to heap)
	globalVar := &localVar
	fmt.Printf("  Global variable: %d\n", *globalVar)
	
	// Demonstrate escape analysis
	escaped := escapeAnalysis()
	fmt.Printf("  Escaped value: %d\n", *escaped)
	
	fmt.Println("  Stack vs heap allocation completed")
}

func escapeAnalysis() *int {
	// This escapes to heap because it's returned
	value := 42
	return &value
}

// Example 5: Memory Pool
func memoryPool() {
	fmt.Println("\n5. Memory Pool")
	fmt.Println("==============")
	
	pool := NewMemoryPool(1024)
	
	// Get objects from pool
	obj1 := pool.Get()
	obj2 := pool.Get()
	
	fmt.Printf("  Allocated objects: %v, %v\n", obj1, obj2)
	
	// Return objects to pool
	pool.Put(obj1)
	pool.Put(obj2)
	
	// Get object again (should reuse)
	obj3 := pool.Get()
	fmt.Printf("  Reused object: %v\n", obj3)
	
	fmt.Println("  Memory pool completed")
}

type MemoryPool struct {
	pool sync.Pool
	size int
}

func NewMemoryPool(size int) *MemoryPool {
	return &MemoryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
		size: size,
	}
}

func (mp *MemoryPool) Get() []byte {
	return mp.pool.Get().([]byte)
}

func (mp *MemoryPool) Put(buf []byte) {
	if len(buf) == mp.size {
		mp.pool.Put(buf)
	}
}

// Example 6: String Optimization
func stringOptimization() {
	fmt.Println("\n6. String Optimization")
	fmt.Println("======================")
	
	// Bad: String concatenation
	start := time.Now()
	var badString string
	for i := 0; i < 1000; i++ {
		badString += "a"
	}
	badTime := time.Since(start)
	
	// Good: String builder
	start = time.Now()
	var goodString strings.Builder
	goodString.Grow(1000) // Pre-allocate capacity
	for i := 0; i < 1000; i++ {
		goodString.WriteString("a")
	}
	goodTime := time.Since(start)
	
	fmt.Printf("  Bad string length: %d, time: %v\n", len(badString), badTime)
	fmt.Printf("  Good string length: %d, time: %v\n", goodString.Len(), goodTime)
	
	fmt.Println("  String optimization completed")
}

// Example 7: Slice Pre-allocation
func slicePreAllocation() {
	fmt.Println("\n7. Slice Pre-allocation")
	fmt.Println("======================")
	
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
	
	fmt.Printf("  Bad slice: len=%d, cap=%d, time=%v\n", len(badSlice), cap(badSlice), badTime)
	fmt.Printf("  Good slice: len=%d, cap=%d, time=%v\n", len(goodSlice), cap(goodSlice), goodTime)
	
	fmt.Println("  Slice pre-allocation completed")
}

// Example 8: Map Optimization
func mapOptimization() {
	fmt.Println("\n8. Map Optimization")
	fmt.Println("==================")
	
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
	
	fmt.Printf("  Bad map: size=%d, time=%v\n", len(badMap), badTime)
	fmt.Printf("  Good map: size=%d, time=%v\n", len(goodMap), goodTime)
	
	fmt.Println("  Map optimization completed")
}

// Example 9: Memory Leak Detection
func memoryLeakDetection() {
	fmt.Println("\n9. Memory Leak Detection")
	fmt.Println("=======================")
	
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
	
	fmt.Printf("  Before allocation: %d bytes\n", m1.Alloc)
	fmt.Printf("  After allocation: %d bytes\n", m2.Alloc)
	fmt.Printf("  After GC: %d bytes\n", m3.Alloc)
	
	if m3.Alloc < m2.Alloc {
		fmt.Println("  No memory leak detected")
	} else {
		fmt.Println("  Potential memory leak detected")
	}
	
	fmt.Println("  Memory leak detection completed")
}

// Example 10: GC Pressure Analysis
func gcPressureAnalysis() {
	fmt.Println("\n10. GC Pressure Analysis")
	fmt.Println("========================")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Calculate GC pressure
	gcPressure := float64(m.PauseTotalNs) / float64(time.Since(time.Unix(0, int64(m.LastGC))).Nanoseconds())
	heapUtilization := float64(m.HeapAlloc) / float64(m.HeapSys)
	
	fmt.Printf("  GC Pressure: %.2f%%\n", gcPressure*100)
	fmt.Printf("  Heap Utilization: %.2f%%\n", heapUtilization*100)
	fmt.Printf("  GC Cycles: %d\n", m.NumGC)
	fmt.Printf("  GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
	
	fmt.Println("  GC pressure analysis completed")
}

// Example 11: Advanced Memory Pool
func advancedMemoryPool() {
	fmt.Println("\n11. Advanced Memory Pool")
	fmt.Println("=======================")
	
	pool := NewAdvancedMemoryPool()
	
	// Get buffers of different sizes
	buf1 := pool.Get(64)
	buf2 := pool.Get(128)
	buf3 := pool.Get(256)
	
	fmt.Printf("  Allocated buffers: %d, %d, %d bytes\n", len(buf1), len(buf2), len(buf3))
	
	// Return buffers
	pool.Put(buf1)
	pool.Put(buf2)
	pool.Put(buf3)
	
	// Get buffer again (should reuse)
	buf4 := pool.Get(64)
	fmt.Printf("  Reused buffer: %d bytes\n", len(buf4))
	
	fmt.Println("  Advanced memory pool completed")
}

type AdvancedMemoryPool struct {
	pools map[int]*sync.Pool
	mu    sync.RWMutex
}

func NewAdvancedMemoryPool() *AdvancedMemoryPool {
	return &AdvancedMemoryPool{
		pools: make(map[int]*sync.Pool),
	}
}

func (amp *AdvancedMemoryPool) Get(size int) []byte {
	amp.mu.RLock()
	pool, exists := amp.pools[size]
	amp.mu.RUnlock()
	
	if !exists {
		amp.mu.Lock()
		pool, exists = amp.pools[size]
		if !exists {
			pool = &sync.Pool{
				New: func() interface{} {
					return make([]byte, size)
				},
			}
			amp.pools[size] = pool
		}
		amp.mu.Unlock()
	}
	
	return pool.Get().([]byte)
}

func (amp *AdvancedMemoryPool) Put(buf []byte) {
	size := len(buf)
	amp.mu.RLock()
	pool, exists := amp.pools[size]
	amp.mu.RUnlock()
	
	if exists {
		pool.Put(buf)
	}
}

// Example 12: Object Reuse Pattern
func objectReusePattern() {
	fmt.Println("\n12. Object Reuse Pattern")
	fmt.Println("=======================")
	
	pool := NewObjectPool(10)
	
	// Get objects from pool
	obj1 := pool.Get()
	obj2 := pool.Get()
	
	fmt.Printf("  Got objects: %d, %d\n", obj1.ID, obj2.ID)
	
	// Use objects
	obj1.Data[0] = 1
	obj2.Data[0] = 2
	
	// Return objects to pool
	pool.Put(obj1)
	pool.Put(obj2)
	
	// Get objects again (should reuse)
	obj3 := pool.Get()
	obj4 := pool.Get()
	
	fmt.Printf("  Reused objects: %d, %d\n", obj3.ID, obj4.ID)
	
	fmt.Println("  Object reuse pattern completed")
}

type ReusableObject struct {
	ID        int
	Data      []byte
	Timestamp time.Time
	inUse     bool
	mu        sync.Mutex
}

type ObjectPool struct {
	objects []*ReusableObject
	free    chan *ReusableObject
	mu      sync.Mutex
}

func NewObjectPool(size int) *ObjectPool {
	pool := &ObjectPool{
		objects: make([]*ReusableObject, size),
		free:    make(chan *ReusableObject, size),
	}
	
	// Initialize objects
	for i := 0; i < size; i++ {
		obj := &ReusableObject{
			ID:    i,
			Data:  make([]byte, 1024),
			inUse: false,
		}
		pool.objects[i] = obj
		pool.free <- obj
	}
	
	return pool
}

func (op *ObjectPool) Get() *ReusableObject {
	obj := <-op.free
	obj.mu.Lock()
	obj.inUse = true
	obj.Timestamp = time.Now()
	obj.mu.Unlock()
	return obj
}

func (op *ObjectPool) Put(obj *ReusableObject) {
	obj.mu.Lock()
	obj.inUse = false
	obj.mu.Unlock()
	op.free <- obj
}

// Example 13: String Interning
func stringInterning() {
	fmt.Println("\n13. String Interning")
	fmt.Println("==================")
	
	interner := NewStringInterner()
	
	// Intern some strings
	str1 := interner.Intern("hello")
	str2 := interner.Intern("world")
	str3 := interner.Intern("hello") // Should return same as str1
	
	fmt.Printf("  str1: %p\n", &str1)
	fmt.Printf("  str2: %p\n", &str2)
	fmt.Printf("  str3: %p\n", &str3)
	fmt.Printf("  str1 == str3: %t\n", &str1 == &str3)
	
	fmt.Println("  String interning completed")
}

type StringInterner struct {
	strings map[string]string
	mu      sync.RWMutex
}

func NewStringInterner() *StringInterner {
	return &StringInterner{
		strings: make(map[string]string),
	}
}

func (si *StringInterner) Intern(s string) string {
	si.mu.RLock()
	if interned, exists := si.strings[s]; exists {
		si.mu.RUnlock()
		return interned
	}
	si.mu.RUnlock()
	
	si.mu.Lock()
	defer si.mu.Unlock()
	
	if interned, exists := si.strings[s]; exists {
		return interned
	}
	
	si.strings[s] = s
	return s
}

// Example 14: Memory Monitoring
func basicMemoryMonitoring() {
	fmt.Println("\n14. Memory Monitoring")
	fmt.Println("====================")
	
	monitor := NewBasicMemoryMonitor(1024*1024, func() {
		fmt.Println("  Memory threshold exceeded!")
	})
	
	// Simulate memory usage
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024)
		time.Sleep(1 * time.Millisecond)
	}
	
	// Stop monitor
	monitor.Stop()
	
	fmt.Println("  Memory monitoring completed")
}

type BasicMemoryMonitor struct {
	threshold uint64
	callback  func()
	mu        sync.Mutex
	quit      chan bool
}

func NewBasicMemoryMonitor(threshold uint64, callback func()) *BasicMemoryMonitor {
	mm := &BasicMemoryMonitor{
		threshold: threshold,
		callback:  callback,
		quit:      make(chan bool),
	}
	
	go mm.monitor()
	
	return mm
}

func (mm *BasicMemoryMonitor) monitor() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			
			if m.Alloc > mm.threshold {
				mm.mu.Lock()
				if mm.callback != nil {
					mm.callback()
				}
				mm.mu.Unlock()
			}
		case <-mm.quit:
			return
		}
	}
}

func (mm *BasicMemoryMonitor) Stop() {
	close(mm.quit)
}

// Example 15: Atomic Memory Counter
func atomicMemoryCounter() {
	fmt.Println("\n15. Atomic Memory Counter")
	fmt.Println("========================")
	
	counter := &AtomicMemoryCounter{}
	
	// Simulate allocations
	counter.Allocate(1024)
	counter.Allocate(2048)
	counter.Allocate(4096)
	
	allocated, freed, peak := counter.GetStats()
	fmt.Printf("  Allocated: %d bytes\n", allocated)
	fmt.Printf("  Freed: %d bytes\n", freed)
	fmt.Printf("  Peak: %d bytes\n", peak)
	
	// Simulate freeing
	counter.Free(1024)
	counter.Free(2048)
	
	allocated, freed, peak = counter.GetStats()
	fmt.Printf("  After freeing: allocated=%d, freed=%d, peak=%d\n", allocated, freed, peak)
	
	fmt.Println("  Atomic memory counter completed")
}

type AtomicMemoryCounter struct {
	allocated int64
	freed     int64
	peak      int64
}

func (amc *AtomicMemoryCounter) Allocate(size int) {
	atomic.AddInt64(&amc.allocated, int64(size))
	
	// Update peak
	for {
		current := atomic.LoadInt64(&amc.peak)
		newPeak := atomic.LoadInt64(&amc.allocated) - atomic.LoadInt64(&amc.freed)
		if newPeak <= current {
			break
		}
		if atomic.CompareAndSwapInt64(&amc.peak, current, newPeak) {
			break
		}
	}
}

func (amc *AtomicMemoryCounter) Free(size int) {
	atomic.AddInt64(&amc.freed, int64(size))
}

func (amc *AtomicMemoryCounter) GetStats() (int64, int64, int64) {
	return atomic.LoadInt64(&amc.allocated),
		atomic.LoadInt64(&amc.freed),
		atomic.LoadInt64(&amc.peak)
}

// Example 16: Memory Growth Analysis
func memoryGrowthAnalysis() {
	fmt.Println("\n16. Memory Growth Analysis")
	fmt.Println("=========================")
	
	var m1, m2 runtime.MemStats
	
	runtime.ReadMemStats(&m1)
	
	// Simulate memory growth
	var data [][]byte
	for i := 0; i < 1000; i++ {
		data = append(data, make([]byte, 1024))
		if i%100 == 0 {
			runtime.ReadMemStats(&m2)
			fmt.Printf("  Iteration %d: %d bytes\n", i, m2.Alloc-m1.Alloc)
		}
	}
	
	runtime.ReadMemStats(&m2)
	fmt.Printf("  Total growth: %d bytes\n", m2.Alloc-m1.Alloc)
	
	fmt.Println("  Memory growth analysis completed")
}

// Example 17: GC Statistics
func gcStatistics() {
	fmt.Println("\n17. GC Statistics")
	fmt.Println("=================")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  GC Cycles: %d\n", m.NumGC)
	fmt.Printf("  GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
	if m.NumGC > 0 {
		fmt.Printf("  GC Pause Average: %v\n", time.Duration(m.PauseTotalNs)/time.Duration(m.NumGC))
	}
	fmt.Printf("  Last GC: %v\n", time.Unix(0, int64(m.LastGC)))
	fmt.Printf("  Next GC: %d bytes\n", m.NextGC)
	
	fmt.Println("  GC statistics completed")
}

// Example 18: Memory Limit
func memoryLimit() {
	fmt.Println("\n18. Memory Limit")
	fmt.Println("===============")
	
	// Set memory limit (Go 1.19+)
	debug.SetMemoryLimit(100 * 1024 * 1024) // 100MB
	fmt.Println("  Memory limit set to 100MB")
	
	// Try to allocate more than limit
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024*1024) // 1MB each
		if i%100 == 0 {
			fmt.Printf("  Allocated %d MB\n", i+1)
		}
	}
	
	// Reset to default
	debug.SetMemoryLimit(0)
	fmt.Println("  Memory limit reset to default")
	
	fmt.Println("  Memory limit completed")
}

// Example 19: Performance Comparison
func performanceComparison() {
	fmt.Println("\n19. Performance Comparison")
	fmt.Println("=========================")
	
	// Test without memory pool
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1024)
	}
	noPoolTime := time.Since(start)
	
	// Test with memory pool
	pool := NewMemoryPool(1024)
	start = time.Now()
	for i := 0; i < 10000; i++ {
		buf := pool.Get()
		_ = buf
		pool.Put(buf)
	}
	poolTime := time.Since(start)
	
	fmt.Printf("  Without pool: %v\n", noPoolTime)
	fmt.Printf("  With pool: %v\n", poolTime)
	fmt.Printf("  Speedup: %.2fx\n", float64(noPoolTime)/float64(poolTime))
	
	fmt.Println("  Performance comparison completed")
}

// Example 20: Memory Profiling
func basicMemoryProfiling() {
	fmt.Println("\n20. Memory Profiling")
	fmt.Println("===================")
	
	// Simulate memory-intensive work
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1024)
	}
	
	// Get memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  Allocated: %d bytes\n", m.Alloc)
	fmt.Printf("  Total Allocated: %d bytes\n", m.TotalAlloc)
	fmt.Printf("  GC Cycles: %d\n", m.NumGC)
	fmt.Printf("  Heap Objects: %d\n", m.HeapObjects)
	
	fmt.Println("  Memory profiling completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🧠 Memory Management Examples")
	fmt.Println("=============================")
	
	basicMemoryStatistics()
	gcTuning()
	memoryAllocationPatterns()
	stackVsHeapAllocation()
	memoryPool()
	stringOptimization()
	slicePreAllocation()
	mapOptimization()
	memoryLeakDetection()
	gcPressureAnalysis()
	advancedMemoryPool()
	objectReusePattern()
	stringInterning()
	basicMemoryMonitoring()
	atomicMemoryCounter()
	memoryGrowthAnalysis()
	gcStatistics()
	memoryLimit()
	performanceComparison()
	basicMemoryProfiling()
}
