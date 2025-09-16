# 🧠 Level 4, Topic 3: Memory Management

## 🚀 Overview
Mastering memory management is crucial for building high-performance concurrent Go applications. This topic will take you from basic garbage collection understanding to advanced memory optimization techniques, making you an expert in memory-efficient concurrent programming.

---

## 📚 Table of Contents

1. [Memory Management Fundamentals](#memory-management-fundamentals)
2. [Garbage Collection in Go](#garbage-collection-in-go)
3. [Memory Allocation Patterns](#memory-allocation-patterns)
4. [Memory Pools and Object Reuse](#memory-pools-and-object-reuse)
5. [Memory Profiling and Analysis](#memory-profiling-and-analysis)
6. [Memory Optimization Techniques](#memory-optimization-techniques)
7. [Concurrent Memory Management](#concurrent-memory-management)
8. [Memory Leak Detection](#memory-leak-detection)
9. [Performance Tuning](#performance-tuning)
10. [Real-World Applications](#real-world-applications)
11. [Advanced Techniques](#advanced-techniques)
12. [Best Practices](#best-practices)

---

## 🧠 Memory Management Fundamentals

### What is Memory Management?

Memory management in Go involves understanding how the Go runtime allocates, uses, and reclaims memory. Unlike languages with manual memory management, Go uses automatic garbage collection, but understanding the underlying mechanisms is crucial for performance optimization.

### Key Concepts

#### 1. Memory Layout
- **Stack**: Fast, automatic allocation/deallocation
- **Heap**: Dynamic allocation, managed by GC
- **Data Segment**: Global variables and constants
- **Code Segment**: Program instructions

#### 2. Allocation Strategies
- **Stack Allocation**: Local variables, function parameters
- **Heap Allocation**: Dynamic data, shared objects
- **Escape Analysis**: Compiler determines allocation location

#### 3. Garbage Collection
- **Mark and Sweep**: Traditional GC algorithm
- **Generational GC**: Go's approach
- **Concurrent GC**: Non-blocking collection

### Memory Management Goals

1. **Efficiency**: Minimize allocation overhead
2. **Performance**: Reduce GC pressure
3. **Reliability**: Prevent memory leaks
4. **Scalability**: Handle large workloads

---

## 🗑️ Garbage Collection in Go

### Go's Garbage Collector

Go uses a concurrent, tri-color, mark-and-sweep garbage collector that runs alongside your program with minimal pauses.

#### GC Phases

```go
package main

import (
    "runtime"
    "time"
)

// Example 1: Understanding GC Phases
func gcPhases() {
    // Force a GC cycle
    runtime.GC()
    
    // Get GC statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("GC Cycles: %d\n", m.NumGC)
    fmt.Printf("GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
    fmt.Printf("Last GC: %v\n", time.Unix(0, int64(m.LastGC)))
}

// Example 2: GC Tuning
func gcTuning() {
    // Set GC percentage (default is 100)
    // Higher values = more memory usage, less GC pressure
    // Lower values = less memory usage, more GC pressure
    debug.SetGCPercent(200)
    
    // Force immediate GC
    runtime.GC()
    
    // Get current GC percentage
    fmt.Printf("GC Percentage: %d\n", debug.SetGCPercent(-1))
}
```

### GC Performance Monitoring

```go
package main

import (
    "runtime"
    "time"
)

// Example 3: GC Performance Monitoring
func gcPerformanceMonitoring() {
    var m1, m2 runtime.MemStats
    
    // Get initial stats
    runtime.ReadMemStats(&m1)
    
    // Do some work
    for i := 0; i < 1000000; i++ {
        _ = make([]byte, 1024)
    }
    
    // Get final stats
    runtime.ReadMemStats(&m2)
    
    fmt.Printf("Allocated: %d bytes\n", m2.Alloc-m1.Alloc)
    fmt.Printf("GC Cycles: %d\n", m2.NumGC-m1.NumGC)
    fmt.Printf("GC Pause: %v\n", time.Duration(m2.PauseTotalNs-m1.PauseTotalNs))
}

// Example 4: GC Pressure Analysis
func gcPressureAnalysis() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    // Calculate GC pressure
    gcPressure := float64(m.PauseTotalNs) / float64(time.Since(time.Unix(0, int64(m.LastGC))).Nanoseconds())
    
    fmt.Printf("GC Pressure: %.2f%%\n", gcPressure*100)
    fmt.Printf("Heap Size: %d bytes\n", m.HeapSys)
    fmt.Printf("Heap Allocated: %d bytes\n", m.HeapAlloc)
}
```

---

## 📊 Memory Allocation Patterns

### Understanding Allocation Patterns

```go
package main

import (
    "runtime"
    "unsafe"
)

// Example 1: Stack vs Heap Allocation
func stackVsHeapAllocation() {
    // Stack allocation (local variable)
    localVar := 42
    
    // Heap allocation (escapes to heap)
    globalVar := &localVar
    
    fmt.Printf("Local: %d\n", localVar)
    fmt.Printf("Global: %d\n", *globalVar)
}

// Example 2: Escape Analysis
func escapeAnalysis() {
    // This escapes to heap because it's returned
    return &struct{ value int }{value: 42}
}

// Example 3: Allocation Size Impact
func allocationSizeImpact() {
    var m1, m2 runtime.MemStats
    
    runtime.ReadMemStats(&m1)
    
    // Small allocations
    for i := 0; i < 1000; i++ {
        _ = make([]byte, 1)
    }
    
    runtime.ReadMemStats(&m2)
    smallAlloc := m2.Alloc - m1.Alloc
    
    runtime.ReadMemStats(&m1)
    
    // Large allocations
    for i := 0; i < 1000; i++ {
        _ = make([]byte, 1024)
    }
    
    runtime.ReadMemStats(&m2)
    largeAlloc := m2.Alloc - m1.Alloc
    
    fmt.Printf("Small allocations: %d bytes\n", smallAlloc)
    fmt.Printf("Large allocations: %d bytes\n", largeAlloc)
}
```

### Memory Allocation Strategies

```go
package main

import (
    "runtime"
    "unsafe"
)

// Example 4: Pre-allocation Strategy
func preAllocationStrategy() {
    // Bad: Growing slice
    var badSlice []int
    for i := 0; i < 1000; i++ {
        badSlice = append(badSlice, i)
    }
    
    // Good: Pre-allocated slice
    goodSlice := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        goodSlice = append(goodSlice, i)
    }
    
    fmt.Printf("Bad slice len: %d, cap: %d\n", len(badSlice), cap(badSlice))
    fmt.Printf("Good slice len: %d, cap: %d\n", len(goodSlice), cap(goodSlice))
}

// Example 5: Object Pooling
func objectPooling() {
    // Create object pool
    pool := &sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024)
        },
    }
    
    // Get object from pool
    obj := pool.Get().([]byte)
    
    // Use object
    for i := range obj {
        obj[i] = byte(i % 256)
    }
    
    // Return object to pool
    pool.Put(obj)
    
    fmt.Println("Object pooled successfully")
}
```

---

## 🏊 Memory Pools and Object Reuse

### Basic Memory Pool

```go
package main

import (
    "sync"
    "unsafe"
)

// Example 1: Basic Memory Pool
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

// Example 2: Advanced Memory Pool
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
```

### Object Reuse Patterns

```go
package main

import (
    "sync"
    "time"
)

// Example 3: Object Reuse Pattern
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

// Example 4: String Interning
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
```

---

## 📈 Memory Profiling and Analysis

### Memory Profiling Tools

```go
package main

import (
    "os"
    "runtime"
    "runtime/pprof"
    "time"
)

// Example 1: Memory Profiling
func memoryProfiling() {
    // Create memory profile
    f, err := os.Create("mem.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    // Do some memory-intensive work
    for i := 0; i < 1000000; i++ {
        _ = make([]byte, 1024)
    }
    
    // Write memory profile
    runtime.GC()
    if err := pprof.WriteHeapProfile(f); err != nil {
        panic(err)
    }
    
    fmt.Println("Memory profile written to mem.prof")
}

// Example 2: Memory Statistics
func memoryStatistics() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Allocated: %d bytes\n", m.Alloc)
    fmt.Printf("Total Allocated: %d bytes\n", m.TotalAlloc)
    fmt.Printf("System Memory: %d bytes\n", m.Sys)
    fmt.Printf("GC Cycles: %d\n", m.NumGC)
    fmt.Printf("GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
    fmt.Printf("Heap Size: %d bytes\n", m.HeapSys)
    fmt.Printf("Heap Allocated: %d bytes\n", m.HeapAlloc)
    fmt.Printf("Heap Objects: %d\n", m.HeapObjects)
}

// Example 3: Memory Leak Detection
func memoryLeakDetection() {
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
    
    fmt.Printf("Before allocation: %d bytes\n", m1.Alloc)
    fmt.Printf("After allocation: %d bytes\n", m2.Alloc)
    fmt.Printf("After GC: %d bytes\n", m3.Alloc)
    
    if m3.Alloc < m2.Alloc {
        fmt.Println("No memory leak detected")
    } else {
        fmt.Println("Potential memory leak detected")
    }
}
```

### Advanced Memory Analysis

```go
package main

import (
    "runtime"
    "time"
)

// Example 4: Memory Growth Analysis
func memoryGrowthAnalysis() {
    var m1, m2 runtime.MemStats
    
    runtime.ReadMemStats(&m1)
    
    // Simulate memory growth
    var data [][]byte
    for i := 0; i < 1000; i++ {
        data = append(data, make([]byte, 1024))
        if i%100 == 0 {
            runtime.ReadMemStats(&m2)
            fmt.Printf("Iteration %d: %d bytes\n", i, m2.Alloc-m1.Alloc)
        }
    }
    
    runtime.ReadMemStats(&m2)
    fmt.Printf("Total growth: %d bytes\n", m2.Alloc-m1.Alloc)
}

// Example 5: GC Pressure Analysis
func gcPressureAnalysis() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    // Calculate various metrics
    gcPressure := float64(m.PauseTotalNs) / float64(time.Since(time.Unix(0, int64(m.LastGC))).Nanoseconds())
    heapUtilization := float64(m.HeapAlloc) / float64(m.HeapSys)
    
    fmt.Printf("GC Pressure: %.2f%%\n", gcPressure*100)
    fmt.Printf("Heap Utilization: %.2f%%\n", heapUtilization*100)
    fmt.Printf("GC Frequency: %.2f Hz\n", float64(m.NumGC)/time.Since(time.Unix(0, int64(m.LastGC))).Seconds())
}
```

---

## ⚡ Memory Optimization Techniques

### Allocation Optimization

```go
package main

import (
    "runtime"
    "sync"
)

// Example 1: Slice Pre-allocation
func slicePreAllocation() {
    // Bad: Growing slice
    var badSlice []int
    for i := 0; i < 1000; i++ {
        badSlice = append(badSlice, i)
    }
    
    // Good: Pre-allocated slice
    goodSlice := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        goodSlice = append(goodSlice, i)
    }
    
    fmt.Printf("Bad slice: len=%d, cap=%d\n", len(badSlice), cap(badSlice))
    fmt.Printf("Good slice: len=%d, cap=%d\n", len(goodSlice), cap(goodSlice))
}

// Example 2: String Optimization
func stringOptimization() {
    // Bad: String concatenation
    var badString string
    for i := 0; i < 1000; i++ {
        badString += "a"
    }
    
    // Good: String builder
    var goodString strings.Builder
    goodString.Grow(1000) // Pre-allocate capacity
    for i := 0; i < 1000; i++ {
        goodString.WriteString("a")
    }
    
    fmt.Printf("Bad string length: %d\n", len(badString))
    fmt.Printf("Good string length: %d\n", goodString.Len())
}

// Example 3: Map Optimization
func mapOptimization() {
    // Bad: Growing map
    badMap := make(map[int]string)
    for i := 0; i < 1000; i++ {
        badMap[i] = "value"
    }
    
    // Good: Pre-sized map
    goodMap := make(map[int]string, 1000)
    for i := 0; i < 1000; i++ {
        goodMap[i] = "value"
    }
    
    fmt.Printf("Bad map size: %d\n", len(badMap))
    fmt.Printf("Good map size: %d\n", len(goodMap))
}
```

### Memory Pool Optimization

```go
package main

import (
    "sync"
    "unsafe"
)

// Example 4: Optimized Memory Pool
type OptimizedMemoryPool struct {
    pools []sync.Pool
    sizes []int
    mu    sync.RWMutex
}

func NewOptimizedMemoryPool() *OptimizedMemoryPool {
    // Pre-defined sizes for common allocations
    sizes := []int{64, 128, 256, 512, 1024, 2048, 4096, 8192}
    pools := make([]sync.Pool, len(sizes))
    
    for i, size := range sizes {
        size := size // Capture for closure
        pools[i] = sync.Pool{
            New: func() interface{} {
                return make([]byte, size)
            },
        }
    }
    
    return &OptimizedMemoryPool{
        pools: pools,
        sizes: sizes,
    }
}

func (omp *OptimizedMemoryPool) Get(size int) []byte {
    // Find appropriate pool
    for i, poolSize := range omp.sizes {
        if size <= poolSize {
            return omp.pools[i].Get().([]byte)[:size]
        }
    }
    
    // Fallback to direct allocation
    return make([]byte, size)
}

func (omp *OptimizedMemoryPool) Put(buf []byte) {
    size := len(buf)
    
    // Find appropriate pool
    for i, poolSize := range omp.sizes {
        if size <= poolSize {
            omp.pools[i].Put(buf)
            return
        }
    }
    
    // Fallback: let GC handle it
}

// Example 5: Lock-Free Memory Pool
type LockFreeMemoryPool struct {
    freeList unsafe.Pointer
    size     int
}

func NewLockFreeMemoryPool(size int) *LockFreeMemoryPool {
    return &LockFreeMemoryPool{size: size}
}

func (lfmp *LockFreeMemoryPool) Get() unsafe.Pointer {
    for {
        current := atomic.LoadPointer(&lfmp.freeList)
        if current == nil {
            return unsafe.Pointer(&make([]byte, lfmp.size)[0])
        }
        
        next := *(*unsafe.Pointer)(current)
        if atomic.CompareAndSwapPointer(&lfmp.freeList, current, next) {
            return current
        }
    }
}

func (lfmp *LockFreeMemoryPool) Put(ptr unsafe.Pointer) {
    for {
        current := atomic.LoadPointer(&lfmp.freeList)
        *(*unsafe.Pointer)(ptr) = current
        
        if atomic.CompareAndSwapPointer(&lfmp.freeList, current, ptr) {
            break
        }
    }
}
```

---

## 🔄 Concurrent Memory Management

### Thread-Safe Memory Management

```go
package main

import (
    "sync"
    "sync/atomic"
)

// Example 1: Thread-Safe Memory Pool
type ThreadSafeMemoryPool struct {
    pools map[int]*sync.Pool
    mu    sync.RWMutex
}

func NewThreadSafeMemoryPool() *ThreadSafeMemoryPool {
    return &ThreadSafeMemoryPool{
        pools: make(map[int]*sync.Pool),
    }
}

func (tsmp *ThreadSafeMemoryPool) Get(size int) []byte {
    tsmp.mu.RLock()
    pool, exists := tsmp.pools[size]
    tsmp.mu.RUnlock()
    
    if !exists {
        tsmp.mu.Lock()
        pool, exists = tsmp.pools[size]
        if !exists {
            pool = &sync.Pool{
                New: func() interface{} {
                    return make([]byte, size)
                },
            }
            tsmp.pools[size] = pool
        }
        tsmp.mu.Unlock()
    }
    
    return pool.Get().([]byte)
}

func (tsmp *ThreadSafeMemoryPool) Put(buf []byte) {
    size := len(buf)
    tsmp.mu.RLock()
    pool, exists := tsmp.pools[size]
    tsmp.mu.RUnlock()
    
    if exists {
        pool.Put(buf)
    }
}

// Example 2: Atomic Memory Counter
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
```

### Memory Management in Concurrent Code

```go
package main

import (
    "sync"
    "time"
)

// Example 3: Concurrent Memory Management
type ConcurrentMemoryManager struct {
    pools    map[int]*sync.Pool
    mu       sync.RWMutex
    counter  *AtomicMemoryCounter
    cleaner  *MemoryCleaner
}

func NewConcurrentMemoryManager() *ConcurrentMemoryManager {
    manager := &ConcurrentMemoryManager{
        pools:   make(map[int]*sync.Pool),
        counter: &AtomicMemoryCounter{},
        cleaner: NewMemoryCleaner(),
    }
    
    // Start cleaner
    go manager.cleaner.Run()
    
    return manager
}

func (cmm *ConcurrentMemoryManager) Get(size int) []byte {
    cmm.mu.RLock()
    pool, exists := cmm.pools[size]
    cmm.mu.RUnlock()
    
    if !exists {
        cmm.mu.Lock()
        pool, exists = cmm.pools[size]
        if !exists {
            pool = &sync.Pool{
                New: func() interface{} {
                    return make([]byte, size)
                },
            }
            cmm.pools[size] = pool
        }
        cmm.mu.Unlock()
    }
    
    buf := pool.Get().([]byte)
    cmm.counter.Allocate(size)
    return buf
}

func (cmm *ConcurrentMemoryManager) Put(buf []byte) {
    size := len(buf)
    cmm.mu.RLock()
    pool, exists := cmm.pools[size]
    cmm.mu.RUnlock()
    
    if exists {
        pool.Put(buf)
        cmm.counter.Free(size)
    }
}

// Example 4: Memory Cleaner
type MemoryCleaner struct {
    quit chan bool
}

func NewMemoryCleaner() *MemoryCleaner {
    return &MemoryCleaner{
        quit: make(chan bool),
    }
}

func (mc *MemoryCleaner) Run() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            runtime.GC()
        case <-mc.quit:
            return
        }
    }
}

func (mc *MemoryCleaner) Stop() {
    close(mc.quit)
}
```

---

## 🔍 Memory Leak Detection

### Memory Leak Detection Techniques

```go
package main

import (
    "runtime"
    "time"
)

// Example 1: Basic Memory Leak Detection
func basicMemoryLeakDetection() {
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
    
    fmt.Printf("Before allocation: %d bytes\n", m1.Alloc)
    fmt.Printf("After allocation: %d bytes\n", m2.Alloc)
    fmt.Printf("After GC: %d bytes\n", m3.Alloc)
    
    if m3.Alloc < m2.Alloc {
        fmt.Println("No memory leak detected")
    } else {
        fmt.Println("Potential memory leak detected")
    }
}

// Example 2: Advanced Memory Leak Detection
type MemoryLeakDetector struct {
    snapshots []MemorySnapshot
    mu        sync.Mutex
}

type MemorySnapshot struct {
    timestamp time.Time
    stats     runtime.MemStats
}

func NewMemoryLeakDetector() *MemoryLeakDetector {
    return &MemoryLeakDetector{
        snapshots: make([]MemorySnapshot, 0),
    }
}

func (mld *MemoryLeakDetector) TakeSnapshot() {
    mld.mu.Lock()
    defer mld.mu.Unlock()
    
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    snapshot := MemorySnapshot{
        timestamp: time.Now(),
        stats:     m,
    }
    
    mld.snapshots = append(mld.snapshots, snapshot)
    
    // Keep only last 10 snapshots
    if len(mld.snapshots) > 10 {
        mld.snapshots = mld.snapshots[1:]
    }
}

func (mld *MemoryLeakDetector) DetectLeak() bool {
    mld.mu.Lock()
    defer mld.mu.Unlock()
    
    if len(mld.snapshots) < 2 {
        return false
    }
    
    // Check if memory is growing consistently
    for i := 1; i < len(mld.snapshots); i++ {
        if mld.snapshots[i].stats.Alloc <= mld.snapshots[i-1].stats.Alloc {
            return false
        }
    }
    
    return true
}

// Example 3: Memory Leak Prevention
type MemoryLeakPrevention struct {
    objects map[interface{}]time.Time
    mu      sync.RWMutex
    ttl     time.Duration
}

func NewMemoryLeakPrevention(ttl time.Duration) *MemoryLeakPrevention {
    mlp := &MemoryLeakPrevention{
        objects: make(map[interface{}]time.Time),
        ttl:     ttl,
    }
    
    // Start cleanup goroutine
    go mlp.cleanup()
    
    return mlp
}

func (mlp *MemoryLeakPrevention) Register(obj interface{}) {
    mlp.mu.Lock()
    defer mlp.mu.Unlock()
    
    mlp.objects[obj] = time.Now()
}

func (mlp *MemoryLeakPrevention) Unregister(obj interface{}) {
    mlp.mu.Lock()
    defer mlp.mu.Unlock()
    
    delete(mlp.objects, obj)
}

func (mlp *MemoryLeakPrevention) cleanup() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        mlp.mu.Lock()
        now := time.Now()
        for obj, timestamp := range mlp.objects {
            if now.Sub(timestamp) > mlp.ttl {
                delete(mlp.objects, obj)
                // Log or handle expired object
            }
        }
        mlp.mu.Unlock()
    }
}
```

---

## ⚙️ Performance Tuning

### GC Tuning

```go
package main

import (
    "runtime"
    "runtime/debug"
)

// Example 1: GC Tuning
func gcTuning() {
    // Set GC percentage
    // Higher values = more memory usage, less GC pressure
    // Lower values = less memory usage, more GC pressure
    debug.SetGCPercent(200)
    
    // Force immediate GC
    runtime.GC()
    
    // Get current GC percentage
    currentPercent := debug.SetGCPercent(-1)
    fmt.Printf("Current GC percentage: %d\n", currentPercent)
}

// Example 2: Memory Limit
func memoryLimit() {
    // Set memory limit (Go 1.19+)
    // This will trigger GC when memory usage exceeds the limit
    debug.SetMemoryLimit(100 * 1024 * 1024) // 100MB
    
    fmt.Println("Memory limit set to 100MB")
}

// Example 3: GC Statistics
func gcStatistics() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("GC Cycles: %d\n", m.NumGC)
    fmt.Printf("GC Pause Total: %v\n", time.Duration(m.PauseTotalNs))
    fmt.Printf("GC Pause Average: %v\n", time.Duration(m.PauseTotalNs)/time.Duration(m.NumGC))
    fmt.Printf("Last GC: %v\n", time.Unix(0, int64(m.LastGC)))
}
```

### Memory Optimization Strategies

```go
package main

import (
    "runtime"
    "sync"
)

// Example 4: Memory Optimization Strategies
func memoryOptimizationStrategies() {
    // Strategy 1: Use sync.Pool for frequently allocated objects
    pool := &sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024)
        },
    }
    
    // Strategy 2: Pre-allocate slices with known capacity
    slice := make([]int, 0, 1000)
    
    // Strategy 3: Use string builder for string concatenation
    var builder strings.Builder
    builder.Grow(1000)
    
    // Strategy 4: Reuse objects instead of creating new ones
    obj := &struct{ value int }{value: 42}
    // Reuse obj instead of creating new ones
    
    fmt.Println("Memory optimization strategies applied")
}

// Example 5: Memory Monitoring
type MemoryMonitor struct {
    threshold int64
    callback  func()
    mu        sync.Mutex
}

func NewMemoryMonitor(threshold int64, callback func()) *MemoryMonitor {
    mm := &MemoryMonitor{
        threshold: threshold,
        callback:  callback,
    }
    
    go mm.monitor()
    
    return mm
}

func (mm *MemoryMonitor) monitor() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        if m.Alloc > mm.threshold {
            mm.mu.Lock()
            if mm.callback != nil {
                mm.callback()
            }
            mm.mu.Unlock()
        }
    }
}
```

---

## 🌍 Real-World Applications

### Web Server Memory Management

```go
package main

import (
    "net/http"
    "sync"
    "time"
)

// Example 1: Web Server Memory Management
type WebServerMemoryManager struct {
    pools    map[int]*sync.Pool
    mu       sync.RWMutex
    counter  *AtomicMemoryCounter
}

func NewWebServerMemoryManager() *WebServerMemoryManager {
    return &WebServerMemoryManager{
        pools:   make(map[int]*sync.Pool),
        counter: &AtomicMemoryCounter{},
    }
}

func (wsmm *WebServerMemoryManager) GetBuffer(size int) []byte {
    wsmm.mu.RLock()
    pool, exists := wsmm.pools[size]
    wsmm.mu.RUnlock()
    
    if !exists {
        wsmm.mu.Lock()
        pool, exists = wsmm.pools[size]
        if !exists {
            pool = &sync.Pool{
                New: func() interface{} {
                    return make([]byte, size)
                },
            }
            wsmm.pools[size] = pool
        }
        wsmm.mu.Unlock()
    }
    
    buf := pool.Get().([]byte)
    wsmm.counter.Allocate(size)
    return buf
}

func (wsmm *WebServerMemoryManager) PutBuffer(buf []byte) {
    size := len(buf)
    wsmm.mu.RLock()
    pool, exists := wsmm.pools[size]
    wsmm.mu.RUnlock()
    
    if exists {
        pool.Put(buf)
        wsmm.counter.Free(size)
    }
}

func (wsmm *WebServerMemoryManager) Handler(w http.ResponseWriter, r *http.Request) {
    // Get buffer from pool
    buf := wsmm.GetBuffer(1024)
    defer wsmm.PutBuffer(buf)
    
    // Use buffer
    response := "Hello, World!"
    copy(buf, response)
    
    w.Write(buf[:len(response)])
}

// Example 2: Database Connection Pool
type DatabaseConnectionPool struct {
    connections chan *Connection
    factory     func() *Connection
    mu          sync.Mutex
    size        int
    maxSize     int
}

type Connection struct {
    ID   int
    Data []byte
}

func NewDatabaseConnectionPool(maxSize int, factory func() *Connection) *DatabaseConnectionPool {
    return &DatabaseConnectionPool{
        connections: make(chan *Connection, maxSize),
        factory:     factory,
        maxSize:     maxSize,
    }
}

func (dcp *DatabaseConnectionPool) Get() *Connection {
    select {
    case conn := <-dcp.connections:
        return conn
    default:
        dcp.mu.Lock()
        if dcp.size < dcp.maxSize {
            dcp.size++
            dcp.mu.Unlock()
            return dcp.factory()
        }
        dcp.mu.Unlock()
        
        // Wait for available connection
        return <-dcp.connections
    }
}

func (dcp *DatabaseConnectionPool) Put(conn *Connection) {
    select {
    case dcp.connections <- conn:
        // Connection returned to pool
    default:
        // Pool is full, let GC handle it
    }
}
```

### Cache Memory Management

```go
package main

import (
    "sync"
    "time"
)

// Example 3: Cache Memory Management
type CacheMemoryManager struct {
    cache    map[string]*CacheEntry
    mu       sync.RWMutex
    maxSize  int
    ttl      time.Duration
    cleaner  *CacheCleaner
}

type CacheEntry struct {
    value     interface{}
    timestamp time.Time
    size      int
}

func NewCacheMemoryManager(maxSize int, ttl time.Duration) *CacheMemoryManager {
    cmm := &CacheMemoryManager{
        cache:   make(map[string]*CacheEntry),
        maxSize: maxSize,
        ttl:     ttl,
        cleaner: NewCacheCleaner(),
    }
    
    go cmm.cleaner.Run(cmm)
    
    return cmm
}

func (cmm *CacheMemoryManager) Get(key string) (interface{}, bool) {
    cmm.mu.RLock()
    entry, exists := cmm.cache[key]
    cmm.mu.RUnlock()
    
    if !exists {
        return nil, false
    }
    
    if time.Since(entry.timestamp) > cmm.ttl {
        cmm.Delete(key)
        return nil, false
    }
    
    return entry.value, true
}

func (cmm *CacheMemoryManager) Set(key string, value interface{}) {
    cmm.mu.Lock()
    defer cmm.mu.Unlock()
    
    // Check if we need to evict
    if len(cmm.cache) >= cmm.maxSize {
        cmm.evict()
    }
    
    entry := &CacheEntry{
        value:     value,
        timestamp: time.Now(),
        size:      calculateSize(value),
    }
    
    cmm.cache[key] = entry
}

func (cmm *CacheMemoryManager) Delete(key string) {
    cmm.mu.Lock()
    delete(cmm.cache, key)
    cmm.mu.Unlock()
}

func (cmm *CacheMemoryManager) evict() {
    // Simple LRU eviction
    var oldestKey string
    var oldestTime time.Time
    
    for key, entry := range cmm.cache {
        if oldestKey == "" || entry.timestamp.Before(oldestTime) {
            oldestKey = key
            oldestTime = entry.timestamp
        }
    }
    
    if oldestKey != "" {
        delete(cmm.cache, oldestKey)
    }
}

func calculateSize(value interface{}) int {
    // Simple size calculation
    switch v := value.(type) {
    case string:
        return len(v)
    case []byte:
        return len(v)
    case int:
        return 8
    default:
        return 0
    }
}

// Example 4: Cache Cleaner
type CacheCleaner struct {
    quit chan bool
}

func NewCacheCleaner() *CacheCleaner {
    return &CacheCleaner{
        quit: make(chan bool),
    }
}

func (cc *CacheCleaner) Run(cmm *CacheMemoryManager) {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            cc.cleanup(cmm)
        case <-cc.quit:
            return
        }
    }
}

func (cc *CacheCleaner) cleanup(cmm *CacheMemoryManager) {
    cmm.mu.Lock()
    defer cmm.mu.Unlock()
    
    now := time.Now()
    for key, entry := range cmm.cache {
        if now.Sub(entry.timestamp) > cmm.ttl {
            delete(cmm.cache, key)
        }
    }
}

func (cc *CacheCleaner) Stop() {
    close(cc.quit)
}
```

---

## 🎓 Summary

Mastering memory management is essential for building high-performance concurrent Go applications. Key takeaways:

1. **Understand GC behavior** and how to tune it
2. **Use memory pools** for frequently allocated objects
3. **Pre-allocate slices and maps** when possible
4. **Monitor memory usage** and detect leaks
5. **Optimize allocation patterns** for better performance
6. **Use profiling tools** to identify bottlenecks
7. **Implement proper cleanup** to prevent leaks
8. **Tune GC parameters** for your workload

Memory management provides the foundation for building efficient, scalable concurrent systems! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different memory optimization techniques
3. **Apply** techniques to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** advanced memory management patterns

Ready to become a Memory Management expert? Let's dive into the implementation! 💪

