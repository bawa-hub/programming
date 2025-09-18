package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// GOD-LEVEL CONCEPT 5: Advanced Optimization Techniques
// Mastering performance optimization for concurrent systems

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: Advanced Optimization Techniques ===")
	
	// 1. Object Pooling
	demonstrateObjectPooling()
	
	// 2. Batching Strategies
	demonstrateBatching()
	
	// 3. NUMA Awareness
	demonstrateNUMAwareness()
	
	// 4. CPU Cache Optimization
	demonstrateCacheOptimization()
	
	// 5. Memory Layout Optimization
	demonstrateMemoryLayout()
	
	// 6. Lock-Free Optimization
	demonstrateLockFreeOptimization()
	
	// 7. Goroutine Optimization
	demonstrateGoroutineOptimization()
	
	// 8. Channel Optimization
	demonstrateChannelOptimization()
}

// Demonstrate Object Pooling
func demonstrateObjectPooling() {
	fmt.Println("\n=== 1. OBJECT POOLING ===")
	
	fmt.Println(`
🔄 Object Pooling:
• Reuse objects instead of allocating
• Reduces GC pressure
• Improves performance
• Use sync.Pool for thread-safe pooling
`)

	// Basic object pooling
	basicObjectPooling()
	
	// Advanced object pooling
	advancedObjectPooling()
	
	// Custom object pooling
	customObjectPooling()
}

func basicObjectPooling() {
	fmt.Println("\n--- Basic Object Pooling ---")
	
	const iterations = 1000000
	
	// Without pooling
	start := time.Now()
	for i := 0; i < iterations; i++ {
		// Allocate new slice each time
		data := make([]int, 1000)
		for j := range data {
			data[j] = j
		}
		_ = data
	}
	withoutPoolDuration := time.Since(start)
	
	// With pooling
	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, 0, 1000)
		},
	}
	
	start = time.Now()
	for i := 0; i < iterations; i++ {
		// Get from pool
		data := pool.Get().([]int)
		data = data[:0] // Reset length
		
		// Use data
		for j := 0; j < 1000; j++ {
			data = append(data, j)
		}
		
		// Return to pool
		pool.Put(data)
	}
	withPoolDuration := time.Since(start)
	
	fmt.Printf("Without pooling: %v\n", withoutPoolDuration)
	fmt.Printf("With pooling:    %v\n", withPoolDuration)
	fmt.Printf("Pooling is %.2fx faster\n", float64(withoutPoolDuration)/float64(withPoolDuration))
}

func advancedObjectPooling() {
	fmt.Println("\n--- Advanced Object Pooling ---")
	
	// Pool with different object types
	type Worker struct {
		ID       int
		Data     []int
		Metadata map[string]interface{}
	}
	
	pool := sync.Pool{
		New: func() interface{} {
			return &Worker{
				Data:     make([]int, 0, 1000),
				Metadata: make(map[string]interface{}),
			}
		},
	}
	
	const iterations = 100000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		worker := pool.Get().(*Worker)
		
		// Reset worker
		worker.ID = i
		worker.Data = worker.Data[:0]
		for k := range worker.Metadata {
			delete(worker.Metadata, k)
		}
		
		// Use worker
		for j := 0; j < 100; j++ {
			worker.Data = append(worker.Data, j)
		}
		worker.Metadata["processed"] = true
		
		// Return to pool
		pool.Put(worker)
	}
	duration := time.Since(start)
	
	fmt.Printf("Advanced pooling: %v\n", duration)
	fmt.Println("💡 Advanced pooling with complex objects")
}

func customObjectPooling() {
	fmt.Println("\n--- Custom Object Pooling ---")
	
	// Custom pool with size limits
	type CustomPool struct {
		pool chan interface{}
		new  func() interface{}
	}
	
	NewCustomPool := func(size int, newFunc func() interface{}) *CustomPool {
		return &CustomPool{
			pool: make(chan interface{}, size),
			new:  newFunc,
		}
	}
	
	Get := func(p *CustomPool) interface{} {
		select {
		case obj := <-p.pool:
			return obj
		default:
			return p.new()
		}
	}
	
	Put := func(p *CustomPool, obj interface{}) {
		select {
		case p.pool <- obj:
		default:
			// Pool is full, discard object
		}
	}
	
	// Create custom pool
	customPool := NewCustomPool(100, func() interface{} {
		return make([]int, 0, 1000)
	})
	
	const iterations = 100000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		data := Get(customPool).([]int)
		data = data[:0]
		
		for j := 0; j < 100; j++ {
			data = append(data, j)
		}
		
		Put(customPool, data)
	}
	duration := time.Since(start)
	
	fmt.Printf("Custom pooling: %v\n", duration)
	fmt.Println("💡 Custom pool with size limits")
}

// Demonstrate Batching
func demonstrateBatching() {
	fmt.Println("\n=== 2. BATCHING STRATEGIES ===")
	
	fmt.Println(`
📦 Batching Strategies:
• Process multiple items together
• Reduce per-item overhead
• Improve throughput
• Balance latency vs throughput
`)

	// Basic batching
	basicBatching()
	
	// Time-based batching
	timeBasedBatching()
	
	// Size-based batching
	sizeBasedBatching()
	
	// Adaptive batching
	adaptiveBatching()
}

func basicBatching() {
	fmt.Println("\n--- Basic Batching ---")
	
	const totalItems = 100000
	const batchSize = 1000
	
	// Without batching
	start := time.Now()
	for i := 0; i < totalItems; i++ {
		// Process one item at a time
		processItem(i)
	}
	withoutBatchingDuration := time.Since(start)
	
	// With batching
	start = time.Now()
	for i := 0; i < totalItems; i += batchSize {
		batch := make([]int, 0, batchSize)
		for j := i; j < i+batchSize && j < totalItems; j++ {
			batch = append(batch, j)
		}
		processBatch(batch)
	}
	withBatchingDuration := time.Since(start)
	
	fmt.Printf("Without batching: %v\n", withoutBatchingDuration)
	fmt.Printf("With batching:    %v\n", withBatchingDuration)
	fmt.Printf("Batching is %.2fx faster\n", float64(withoutBatchingDuration)/float64(withBatchingDuration))
}

func timeBasedBatching() {
	fmt.Println("\n--- Time-Based Batching ---")
	
	// Time-based batcher
	type TimeBatcher struct {
		batch     []int
		timeout   time.Duration
		processor func([]int)
		timer     *time.Timer
		mu        sync.Mutex
	}
	
	NewTimeBatcher := func(timeout time.Duration, processor func([]int)) *TimeBatcher {
		tb := &TimeBatcher{
			batch:     make([]int, 0, 1000),
			timeout:   timeout,
			processor: processor,
		}
		return tb
	}
	
	Add := func(tb *TimeBatcher, item int) {
		tb.mu.Lock()
		defer tb.mu.Unlock()
		
		tb.batch = append(tb.batch, item)
		
		// Reset timer
		if tb.timer != nil {
			tb.timer.Stop()
		}
		tb.timer = time.AfterFunc(tb.timeout, func() {
			tb.mu.Lock()
			defer tb.mu.Unlock()
			
			if len(tb.batch) > 0 {
				tb.processor(tb.batch)
				tb.batch = tb.batch[:0]
			}
		})
	}
	
	// Create time-based batcher
	batcher := NewTimeBatcher(10*time.Millisecond, processBatch)
	
	const iterations = 10000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		Add(batcher, i)
		time.Sleep(1 * time.Microsecond) // Simulate work
	}
	
	// Flush remaining items
	batcher.mu.Lock()
	if len(batcher.batch) > 0 {
		batcher.processor(batcher.batch)
		batcher.batch = batcher.batch[:0]
	}
	batcher.mu.Unlock()
	duration := time.Since(start)
	
	fmt.Printf("Time-based batching: %v\n", duration)
	fmt.Println("💡 Time-based batching with automatic flush")
}

func sizeBasedBatching() {
	fmt.Println("\n--- Size-Based Batching ---")
	
	// Size-based batcher
	type SizeBatcher struct {
		batch     []int
		maxSize   int
		processor func([]int)
		mu        sync.Mutex
	}
	
	NewSizeBatcher := func(maxSize int, processor func([]int)) *SizeBatcher {
		return &SizeBatcher{
			batch:     make([]int, 0, maxSize),
			maxSize:   maxSize,
			processor: processor,
		}
	}
	
	Add := func(sb *SizeBatcher, item int) {
		sb.mu.Lock()
		defer sb.mu.Unlock()
		
		sb.batch = append(sb.batch, item)
		
		if len(sb.batch) >= sb.maxSize {
			sb.processor(sb.batch)
			sb.batch = sb.batch[:0]
		}
	}
	
	// Create size-based batcher
	batcher := NewSizeBatcher(1000, processBatch)
	
	const iterations = 10000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		Add(batcher, i)
	}
	
	// Process remaining items
	if len(batcher.batch) > 0 {
		processBatch(batcher.batch)
	}
	duration := time.Since(start)
	
	fmt.Printf("Size-based batching: %v\n", duration)
	fmt.Println("💡 Size-based batching with automatic processing")
}

func adaptiveBatching() {
	fmt.Println("\n--- Adaptive Batching ---")
	
	// Adaptive batcher that adjusts batch size based on performance
	type AdaptiveBatcher struct {
		batch       []int
		minSize     int
		maxSize     int
		currentSize int
		processor   func([]int)
		mu          sync.Mutex
	}
	
	NewAdaptiveBatcher := func(minSize, maxSize int, processor func([]int)) *AdaptiveBatcher {
		return &AdaptiveBatcher{
			batch:       make([]int, 0, maxSize),
			minSize:     minSize,
			maxSize:     maxSize,
			currentSize: minSize,
			processor:   processor,
		}
	}
	
	Add := func(ab *AdaptiveBatcher, item int) {
		ab.mu.Lock()
		defer ab.mu.Unlock()
		
		ab.batch = append(ab.batch, item)
		
		if len(ab.batch) >= ab.currentSize {
			start := time.Now()
			ab.processor(ab.batch)
			duration := time.Since(start)
			
			// Adjust batch size based on performance
			if duration < 1*time.Millisecond {
				ab.currentSize = min(ab.currentSize*2, ab.maxSize)
			} else if duration > 5*time.Millisecond {
				ab.currentSize = max(ab.currentSize/2, ab.minSize)
			}
			
			ab.batch = ab.batch[:0]
		}
	}
	
	// Create adaptive batcher
	batcher := NewAdaptiveBatcher(100, 1000, processBatch)
	
	const iterations = 10000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		Add(batcher, i)
	}
	
	// Process remaining items
	if len(batcher.batch) > 0 {
		processBatch(batcher.batch)
	}
	duration := time.Since(start)
	
	fmt.Printf("Adaptive batching: %v\n", duration)
	fmt.Println("💡 Adaptive batching adjusts size based on performance")
}

// Demonstrate NUMA Awareness
func demonstrateNUMAwareness() {
	fmt.Println("\n=== 3. NUMA AWARENESS ===")
	
	fmt.Println(`
🏗️  NUMA Awareness:
• Non-Uniform Memory Access
• Memory access speed varies by distance
• Optimize for CPU topology
• Use CPU-local data structures
`)

	// CPU topology information
	cpuTopology()
	
	// NUMA-aware data structures
	numaAwareDataStructures()
	
	// CPU-local caching
	cpuLocalCaching()
}

func cpuTopology() {
	fmt.Println("\n--- CPU Topology ---")
	
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	
	// Show CPU information
	fmt.Println(`
💡 NUMA Considerations:
• Keep related data on same CPU
• Minimize cross-socket memory access
• Use CPU-local caches
• Consider NUMA topology for large systems
`)
}

func numaAwareDataStructures() {
	fmt.Println("\n--- NUMA-Aware Data Structures ---")
	
	// Per-CPU data structures
	numCPUs := runtime.NumCPU()
	perCPUData := make([][]int, numCPUs)
	
	for i := range perCPUData {
		perCPUData[i] = make([]int, 0, 1000)
	}
	
	const iterations = 100000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		// Use CPU-local data
		cpuID := i % numCPUs
		perCPUData[cpuID] = append(perCPUData[cpuID], i)
	}
	duration := time.Since(start)
	
	fmt.Printf("NUMA-aware data structures: %v\n", duration)
	fmt.Println("💡 Per-CPU data structures reduce cross-socket access")
}

func cpuLocalCaching() {
	fmt.Println("\n--- CPU-Local Caching ---")
	
	// CPU-local cache using sync.Pool
	numCPUs := runtime.NumCPU()
	caches := make([]sync.Pool, numCPUs)
	
	for i := range caches {
		caches[i] = sync.Pool{
			New: func() interface{} {
				return make([]int, 0, 1000)
			},
		}
	}
	
	const iterations = 100000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		// Get CPU-local cache
		cpuID := i % numCPUs
		cache := caches[cpuID].Get().([]int)
		cache = cache[:0]
		
		// Use cache
		for j := 0; j < 100; j++ {
			cache = append(cache, j)
		}
		
		// Return to CPU-local cache
		caches[cpuID].Put(cache)
	}
	duration := time.Since(start)
	
	fmt.Printf("CPU-local caching: %v\n", duration)
	fmt.Println("💡 CPU-local caches improve memory access patterns")
}

// Demonstrate Cache Optimization
func demonstrateCacheOptimization() {
	fmt.Println("\n=== 4. CPU CACHE OPTIMIZATION ===")
	
	fmt.Println(`
⚡ CPU Cache Optimization:
• Optimize for cache line size (64 bytes)
• Reduce cache misses
• Improve data locality
• Use cache-friendly data structures
`)

	// Cache line size optimization
	cacheLineOptimization()
	
	// Data locality optimization
	dataLocalityOptimization()
	
	// Cache-friendly algorithms
	cacheFriendlyAlgorithms()
}

func cacheLineOptimization() {
	fmt.Println("\n--- Cache Line Optimization ---")
	
	const iterations = 1000000
	
	// Bad: False sharing
	type BadCounter struct {
		counter1 int64
		counter2 int64
		counter3 int64
		counter4 int64
	}
	
	badCounter := &BadCounter{}
	var wg sync.WaitGroup
	
	start := time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(counter *int64) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(counter, 1)
			}
		}(&badCounter.counter1)
	}
	wg.Wait()
	badDuration := time.Since(start)
	
	// Good: Cache line padding
	type GoodCounter struct {
		counter1 int64
		_        [7]int64 // Padding to next cache line
		counter2 int64
		_        [7]int64 // Padding to next cache line
		counter3 int64
		_        [7]int64 // Padding to next cache line
		counter4 int64
		_        [7]int64 // Padding to next cache line
	}
	
	goodCounter := &GoodCounter{}
	
	start = time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(counter *int64) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(counter, 1)
			}
		}(&goodCounter.counter1)
	}
	wg.Wait()
	goodDuration := time.Since(start)
	
	fmt.Printf("False sharing: %v\n", badDuration)
	fmt.Printf("Cache padding: %v\n", goodDuration)
	fmt.Printf("Cache padding is %.2fx faster\n", float64(badDuration)/float64(goodDuration))
}

func dataLocalityOptimization() {
	fmt.Println("\n--- Data Locality Optimization ---")
	
	const size = 10000
	
	// Bad: Random access pattern
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = i + j
		}
	}
	
	start := time.Now()
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += matrix[i][j]
		}
	}
	rowMajorDuration := time.Since(start)
	
	// Good: Sequential access pattern
	start = time.Now()
	sum = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += matrix[i][j]
		}
	}
	sequentialDuration := time.Since(start)
	
	fmt.Printf("Row-major access: %v\n", rowMajorDuration)
	fmt.Printf("Sequential access: %v\n", sequentialDuration)
	fmt.Println("💡 Sequential access improves cache locality")
}

func cacheFriendlyAlgorithms() {
	fmt.Println("\n--- Cache-Friendly Algorithms ---")
	
	const size = 1000
	
	// Create test data
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	
	// Bad: Bubble sort (cache-unfriendly)
	data1 := make([]int, len(data))
	copy(data1, data)
	
	start := time.Now()
	bubbleSort(data1)
	bubbleSortDuration := time.Since(start)
	
	// Good: Quick sort (cache-friendly)
	data2 := make([]int, len(data))
	copy(data2, data)
	
	start = time.Now()
	quickSort(data2)
	quickSortDuration := time.Since(start)
	
	fmt.Printf("Bubble sort: %v\n", bubbleSortDuration)
	fmt.Printf("Quick sort:  %v\n", quickSortDuration)
	fmt.Printf("Quick sort is %.2fx faster\n", float64(bubbleSortDuration)/float64(quickSortDuration))
}

// Demonstrate Memory Layout Optimization
func demonstrateMemoryLayout() {
	fmt.Println("\n=== 5. MEMORY LAYOUT OPTIMIZATION ===")
	
	fmt.Println(`
🧠 Memory Layout Optimization:
• Optimize struct layout
• Reduce memory usage
• Improve cache performance
• Use appropriate data types
`)

	// Struct layout optimization
	structLayoutOptimization()
	
	// Memory alignment
	memoryAlignment()
	
	// Slice vs Array performance
	sliceVsArrayPerformance()
}

func structLayoutOptimization() {
	fmt.Println("\n--- Struct Layout Optimization ---")
	
	// Bad: Inefficient struct layout
	type BadStruct struct {
		flag1 bool
		value int64
		flag2 bool
		value2 int64
		flag3 bool
	}
	
	// Good: Efficient struct layout
	type GoodStruct struct {
		value  int64
		value2 int64
		flag1  bool
		flag2  bool
		flag3  bool
	}
	
	const iterations = 1000000
	
	// Test bad struct
	badStructs := make([]BadStruct, iterations)
	start := time.Now()
	for i := range badStructs {
		badStructs[i] = BadStruct{
			flag1:  true,
			value:  int64(i),
			flag2:  false,
			value2: int64(i * 2),
			flag3:  true,
		}
	}
	badDuration := time.Since(start)
	
	// Test good struct
	goodStructs := make([]GoodStruct, iterations)
	start = time.Now()
	for i := range goodStructs {
		goodStructs[i] = GoodStruct{
			value:  int64(i),
			value2: int64(i * 2),
			flag1:  true,
			flag2:  false,
			flag3:  true,
		}
	}
	goodDuration := time.Since(start)
	
	fmt.Printf("Bad struct layout: %v\n", badDuration)
	fmt.Printf("Good struct layout: %v\n", goodDuration)
	fmt.Printf("Good layout is %.2fx faster\n", float64(badDuration)/float64(goodDuration))
}

func memoryAlignment() {
	fmt.Println("\n--- Memory Alignment ---")
	
	// Show struct sizes
	type UnalignedStruct struct {
		flag bool
		value int64
	}
	
	type AlignedStruct struct {
		value int64
		flag  bool
	}
	
	fmt.Printf("Unaligned struct size: %d bytes\n", unsafe.Sizeof(UnalignedStruct{}))
	fmt.Printf("Aligned struct size: %d bytes\n", unsafe.Sizeof(AlignedStruct{}))
	fmt.Println("💡 Proper alignment reduces memory usage and improves performance")
}

func sliceVsArrayPerformance() {
	fmt.Println("\n--- Slice vs Array Performance ---")
	
	const size = 1000
	const iterations = 100000
	
	// Array performance
	var array [size]int
	start := time.Now()
	for i := 0; i < iterations; i++ {
		for j := 0; j < size; j++ {
			array[j] = j
		}
	}
	arrayDuration := time.Since(start)
	
	// Slice performance
	slice := make([]int, size)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		for j := 0; j < size; j++ {
			slice[j] = j
		}
	}
	sliceDuration := time.Since(start)
	
	fmt.Printf("Array: %v\n", arrayDuration)
	fmt.Printf("Slice: %v\n", sliceDuration)
	fmt.Printf("Array is %.2fx faster\n", float64(sliceDuration)/float64(arrayDuration))
}

// Demonstrate Lock-Free Optimization
func demonstrateLockFreeOptimization() {
	fmt.Println("\n=== 6. LOCK-FREE OPTIMIZATION ===")
	
	fmt.Println(`
🔓 Lock-Free Optimization:
• Use atomic operations
• Avoid mutex contention
• Improve scalability
• Handle ABA problem
`)

	// Atomic operations optimization
	atomicOptimization()
	
	// Lock-free data structures
	lockFreeDataStructures()
	
	// Memory ordering optimization
	memoryOrderingOptimization()
}

func atomicOptimization() {
	fmt.Println("\n--- Atomic Operations Optimization ---")
	
	const iterations = 1000000
	const numGoroutines = 10
	
	// Atomic counter
	var atomicCounter int64
	var wg sync.WaitGroup
	
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicDuration := time.Since(start)
	
	// Mutex counter
	var mutexCounter int64
	var mu sync.Mutex
	
	start = time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexDuration := time.Since(start)
	
	fmt.Printf("Atomic: %v (%d ops)\n", atomicDuration, atomicCounter)
	fmt.Printf("Mutex:  %v (%d ops)\n", mutexDuration, mutexCounter)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexDuration)/float64(atomicDuration))
}

func lockFreeDataStructures() {
	fmt.Println("\n--- Lock-Free Data Structures ---")
	
	// Lock-free stack
	stack := NewLockFreeStack()
	const operations = 100000
	
	start := time.Now()
	for i := 0; i < operations; i++ {
		stack.Push(i)
	}
	for i := 0; i < operations; i++ {
		stack.Pop()
	}
	stackDuration := time.Since(start)
	
	// Lock-based stack
	lockStack := NewLockBasedStack()
	
	start = time.Now()
	for i := 0; i < operations; i++ {
		lockStack.Push(i)
	}
	for i := 0; i < operations; i++ {
		lockStack.Pop()
	}
	lockStackDuration := time.Since(start)
	
	fmt.Printf("Lock-free stack: %v\n", stackDuration)
	fmt.Printf("Lock-based stack: %v\n", lockStackDuration)
	fmt.Printf("Lock-free is %.2fx faster\n", float64(lockStackDuration)/float64(stackDuration))
}

func memoryOrderingOptimization() {
	fmt.Println("\n--- Memory Ordering Optimization ---")
	
	// Demonstrate different memory ordering
	var x, y int64
	
	// Relaxed ordering
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.StoreInt64(&x, int64(i))
		atomic.LoadInt64(&y)
	}
	relaxedDuration := time.Since(start)
	
	// Sequential ordering
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.StoreInt64(&x, int64(i))
		atomic.LoadInt64(&y)
	}
	sequentialDuration := time.Since(start)
	
	fmt.Printf("Relaxed ordering: %v\n", relaxedDuration)
	fmt.Printf("Sequential ordering: %v\n", sequentialDuration)
	fmt.Println("💡 Choose appropriate memory ordering for your use case")
}

// Demonstrate Goroutine Optimization
func demonstrateGoroutineOptimization() {
	fmt.Println("\n=== 7. GOROUTINE OPTIMIZATION ===")
	
	fmt.Println(`
🔄 Goroutine Optimization:
• Use appropriate number of goroutines
• Implement worker pools
• Avoid goroutine leaks
• Optimize goroutine communication
`)

	// Worker pool optimization
	workerPoolOptimization()
	
	// Goroutine communication optimization
	goroutineCommunicationOptimization()
	
	// Goroutine lifecycle optimization
	goroutineLifecycleOptimization()
}

func workerPoolOptimization() {
	fmt.Println("\n--- Worker Pool Optimization ---")
	
	const totalWork = 100000
	numWorkers := runtime.NumCPU()
	
	// Without worker pool
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < totalWork; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			processWork(id)
		}(i)
	}
	wg.Wait()
	withoutPoolDuration := time.Since(start)
	
	// With worker pool
	start = time.Now()
	jobs := make(chan int, totalWork)
	results := make(chan int, totalWork)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobs {
				results <- processWork(job)
			}
		}()
	}
	
	// Send jobs
	go func() {
		for i := 0; i < totalWork; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	
	// Collect results
	for i := 0; i < totalWork; i++ {
		<-results
	}
	withPoolDuration := time.Since(start)
	
	fmt.Printf("Without worker pool: %v\n", withoutPoolDuration)
	fmt.Printf("With worker pool:    %v\n", withPoolDuration)
	fmt.Printf("Worker pool is %.2fx faster\n", float64(withoutPoolDuration)/float64(withPoolDuration))
}

func goroutineCommunicationOptimization() {
	fmt.Println("\n--- Goroutine Communication Optimization ---")
	
	const iterations = 100000
	
	// Channel communication
	start := time.Now()
	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// Process value
		}
	}()
	wg.Wait()
	channelDuration := time.Since(start)
	
	// Shared memory communication
	start = time.Now()
	var counter int64
	var mu sync.Mutex
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			mu.Lock()
			_ = counter
			mu.Unlock()
		}
	}()
	wg.Wait()
	sharedMemoryDuration := time.Since(start)
	
	fmt.Printf("Channel communication: %v\n", channelDuration)
	fmt.Printf("Shared memory:         %v\n", sharedMemoryDuration)
	fmt.Printf("Channel is %.2fx faster\n", float64(sharedMemoryDuration)/float64(channelDuration))
}

func goroutineLifecycleOptimization() {
	fmt.Println("\n--- Goroutine Lifecycle Optimization ---")
	
	// Demonstrate proper goroutine cleanup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Start goroutine with context
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine cleaned up properly")
				return
			default:
				// Do work
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
	
	// Simulate work
	time.Sleep(10 * time.Millisecond)
	
	// Cancel context
	cancel()
	
	fmt.Println("💡 Use context for proper goroutine lifecycle management")
}

// Demonstrate Channel Optimization
func demonstrateChannelOptimization() {
	fmt.Println("\n=== 8. CHANNEL OPTIMIZATION ===")
	
	fmt.Println(`
📡 Channel Optimization:
• Use buffered channels appropriately
• Avoid channel leaks
• Optimize channel operations
• Use select efficiently
`)

	// Buffered vs unbuffered channels
	bufferedVsUnbufferedChannels()
	
	// Channel size optimization
	channelSizeOptimization()
	
	// Select optimization
	selectOptimization()
}

func bufferedVsUnbufferedChannels() {
	fmt.Println("\n--- Buffered vs Unbuffered Channels ---")
	
	const iterations = 100000
	
	// Unbuffered channel
	unbufferedCh := make(chan int)
	start := time.Now()
	
	go func() {
		for i := 0; i < iterations; i++ {
			unbufferedCh <- i
		}
		close(unbufferedCh)
	}()
	
	for range unbufferedCh {
		// Process value
	}
	unbufferedDuration := time.Since(start)
	
	// Buffered channel
	bufferedCh := make(chan int, 1000)
	start = time.Now()
	
	go func() {
		for i := 0; i < iterations; i++ {
			bufferedCh <- i
		}
		close(bufferedCh)
	}()
	
	for range bufferedCh {
		// Process value
	}
	bufferedDuration := time.Since(start)
	
	fmt.Printf("Unbuffered channel: %v\n", unbufferedDuration)
	fmt.Printf("Buffered channel:   %v\n", bufferedDuration)
	fmt.Printf("Buffered is %.2fx faster\n", float64(unbufferedDuration)/float64(bufferedDuration))
}

func channelSizeOptimization() {
	fmt.Println("\n--- Channel Size Optimization ---")
	
	const iterations = 100000
	
	// Test different buffer sizes
	bufferSizes := []int{0, 1, 10, 100, 1000}
	
	for _, size := range bufferSizes {
		ch := make(chan int, size)
		start := time.Now()
		
		go func() {
			for i := 0; i < iterations; i++ {
				ch <- i
			}
			close(ch)
		}()
		
		for range ch {
			// Process value
		}
		duration := time.Since(start)
		
		fmt.Printf("Buffer size %d: %v\n", size, duration)
	}
	
	fmt.Println("💡 Choose buffer size based on producer/consumer speed")
}

func selectOptimization() {
	fmt.Println("\n--- Select Optimization ---")
	
	const iterations = 100000
	
	// Multiple channels with select
	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)
	ch3 := make(chan int, 1000)
	
	// Start producers
	go func() {
		for i := 0; i < iterations; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	
	go func() {
		for i := 0; i < iterations; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	
	go func() {
		for i := 0; i < iterations; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	
	// Consumer with select
	start := time.Now()
	count := 0
	for count < iterations*3 {
		select {
		case <-ch1:
			count++
		case <-ch2:
			count++
		case <-ch3:
			count++
		}
	}
	duration := time.Since(start)
	
	fmt.Printf("Select with 3 channels: %v\n", duration)
	fmt.Println("💡 Select efficiently handles multiple channels")
}

// Helper functions
func processItem(item int) {
	_ = item * item
}

func processBatch(batch []int) {
	for _, item := range batch {
		_ = item * item
	}
}

func processWork(id int) int {
	return id * id
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	
	left, right := 0, len(arr)-1
	pivot := partition(arr, left, right)
	
	quickSort(arr[:pivot])
	quickSort(arr[pivot+1:])
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Lock-free stack implementation
type LockFreeStack struct {
	head unsafe.Pointer
}

type node struct {
	value int
	next  unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	newNode := &node{value: value}
	
	for {
		head := atomic.LoadPointer(&s.head)
		newNode.next = head
		
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}
		
		node := (*node)(head)
		next := atomic.LoadPointer(&node.next)
		
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return node.value, true
		}
	}
}

// Lock-based stack implementation
type LockBasedStack struct {
	items []int
	mu    sync.Mutex
}

func NewLockBasedStack() *LockBasedStack {
	return &LockBasedStack{}
}

func (s *LockBasedStack) Push(value int) {
	s.mu.Lock()
	s.items = append(s.items, value)
	s.mu.Unlock()
}

func (s *LockBasedStack) Pop() (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if len(s.items) == 0 {
		return 0, false
	}
	
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]
	return value, true
}

