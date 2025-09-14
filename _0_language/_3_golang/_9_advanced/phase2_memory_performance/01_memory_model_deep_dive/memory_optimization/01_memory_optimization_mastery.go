package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// 🧠 MEMORY OPTIMIZATION MASTERY
// Understanding advanced memory optimization techniques in Go

func main() {
	fmt.Println("🧠 MEMORY OPTIMIZATION MASTERY")
	fmt.Println("==============================")
	fmt.Println()

	// 1. Stack vs Heap Allocation
	stackVsHeapAllocation()
	fmt.Println()

	// 2. Escape Analysis
	escapeAnalysis()
	fmt.Println()

	// 3. Memory Pooling Techniques
	memoryPoolingTechniques()
	fmt.Println()

	// 4. Zero-Allocation Programming
	zeroAllocationProgramming()
	fmt.Println()

	// 5. String Building Optimization
	stringBuildingOptimization()
	fmt.Println()

	// 6. Slice Pre-allocation Strategies
	slicePreallocationStrategies()
	fmt.Println()

	// 7. Memory Pool Patterns
	memoryPoolPatterns()
	fmt.Println()

	// 8. Advanced Memory Techniques
	advancedMemoryTechniques()
	fmt.Println()

	// 9. Performance Comparison
	performanceComparison()
	fmt.Println()

	// 10. Best Practices
	memoryOptimizationBestPractices()
}

// 1. Stack vs Heap Allocation
func stackVsHeapAllocation() {
	fmt.Println("1. Stack vs Heap Allocation:")
	fmt.Println("Understanding Go's memory allocation decisions...")

	// Stack allocation example
	stackExample()
	
	// Heap allocation example
	heapExample()
	
	// Mixed allocation example
	mixedAllocationExample()
}

func stackExample() {
	fmt.Println("  📊 Stack allocation example:")
	
	// This will likely be allocated on the stack
	var localInt int = 42
	var localString string = "stack allocated"
	var localSlice []int = make([]int, 10)
	
	fmt.Printf("    Local int: %d (likely stack)\n", localInt)
	fmt.Printf("    Local string: %s (likely stack)\n", localString)
	fmt.Printf("    Local slice: %v (slice header on stack, data on heap)\n", localSlice[:3])
}

func heapExample() {
	fmt.Println("  📊 Heap allocation example:")
	
	// This will likely be allocated on the heap
	heapData := createHeapData()
	fmt.Printf("    Heap data: %+v\n", heapData)
}

func createHeapData() *Data {
	// Returning a pointer forces heap allocation
	return &Data{
		ID:   123,
		Name: "heap allocated",
		Data: make([]byte, 1000),
	}
}

type Data struct {
	ID   int
	Name string
	Data []byte
}

func mixedAllocationExample() {
	fmt.Println("  📊 Mixed allocation example:")
	
	// Some on stack, some on heap
	stackVar := 42
	heapVar := &stackVar // Pointer to stack variable
	
	fmt.Printf("    Stack var: %d\n", stackVar)
	fmt.Printf("    Heap pointer: %d\n", *heapVar)
}

// 2. Escape Analysis
func escapeAnalysis() {
	fmt.Println("2. Escape Analysis:")
	fmt.Println("Understanding how Go determines stack vs heap allocation...")

	// Demonstrate escape analysis
	escapeAnalysisExamples()
	
	// Show how to analyze escapes
	analyzeEscapes()
}

func escapeAnalysisExamples() {
	fmt.Println("  📊 Escape analysis examples:")
	
	// This escapes to heap (returned)
	escaped := returnEscaped()
	fmt.Printf("    Escaped data: %+v\n", escaped)
	
	// This doesn't escape (local use only)
	localUse()
	
	// This escapes to heap (stored in global)
	storeInGlobal()
}

func returnEscaped() *Data {
	// This escapes because it's returned
	return &Data{ID: 1, Name: "escaped"}
}

func localUse() {
	// This doesn't escape - only used locally
	data := Data{ID: 2, Name: "local"}
	fmt.Printf("    Local data: %+v\n", data)
}

var globalData *Data

func storeInGlobal() {
	// This escapes because it's stored in global variable
	globalData = &Data{ID: 3, Name: "global"}
	fmt.Printf("    Global data: %+v\n", globalData)
}

func analyzeEscapes() {
	fmt.Println("  📊 To analyze escapes, use: go build -gcflags='-m'")
	fmt.Println("  📊 Look for 'escapes to heap' in compiler output")
}

// 3. Memory Pooling Techniques
func memoryPoolingTechniques() {
	fmt.Println("3. Memory Pooling Techniques:")
	fmt.Println("Understanding memory pooling patterns...")

	// Sync.Pool usage
	syncPoolExample()
	
	// Custom memory pools
	customMemoryPoolExample()
	
	// Pool lifecycle management
	poolLifecycleManagement()
}

func syncPoolExample() {
	fmt.Println("  📊 Sync.Pool example:")
	
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	
	// Get from pool
	buf1 := pool.Get().([]byte)
	fmt.Printf("    Got buffer from pool: len=%d\n", len(buf1))
	
	// Use buffer
	copy(buf1, []byte("Hello, World!"))
	fmt.Printf("    Buffer content: %s\n", string(buf1[:13]))
	
	// Return to pool
	pool.Put(buf1)
	fmt.Println("    Buffer returned to pool")
	
	// Get again (might be same buffer)
	buf2 := pool.Get().([]byte)
	fmt.Printf("    Got buffer again: len=%d\n", len(buf2))
	pool.Put(buf2)
}

func customMemoryPoolExample() {
	fmt.Println("  📊 Custom memory pool example:")
	
	pool := NewCustomPool(1024)
	
	// Allocate from pool
	buf := pool.Get()
	fmt.Printf("    Allocated from custom pool: len=%d\n", len(buf))
	
	// Use buffer
	copy(buf, []byte("Custom pool data"))
	fmt.Printf("    Buffer content: %s\n", string(buf[:16]))
	
	// Return to pool
	pool.Put(buf)
	fmt.Println("    Buffer returned to custom pool")
}

type CustomPool struct {
	pool sync.Pool
	size int
}

func NewCustomPool(size int) *CustomPool {
	return &CustomPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
		size: size,
	}
}

func (p *CustomPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *CustomPool) Put(buf []byte) {
	if len(buf) == p.size {
		p.pool.Put(buf)
	}
}

func poolLifecycleManagement() {
	fmt.Println("  📊 Pool lifecycle management:")
	
	// Demonstrate pool cleanup
	pool := NewManagedPool()
	
	// Use pool
	buf := pool.Get()
	fmt.Printf("    Got buffer: %p\n", buf)
	pool.Put(buf)
	
	// Cleanup pool
	pool.Cleanup()
	fmt.Println("    Pool cleaned up")
}

type ManagedPool struct {
	pool    sync.Pool
	active  map[uintptr]bool
	mu      sync.Mutex
	cleaned bool
}

func NewManagedPool() *ManagedPool {
	return &ManagedPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 1024)
			},
		},
		active: make(map[uintptr]bool),
	}
}

func (p *ManagedPool) Get() []byte {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.cleaned {
		return nil
	}
	
	buf := p.pool.Get().([]byte)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	p.active[ptr] = true
	return buf
}

func (p *ManagedPool) Put(buf []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.cleaned {
		return
	}
	
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	delete(p.active, ptr)
	p.pool.Put(buf)
}

func (p *ManagedPool) Cleanup() {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	p.cleaned = true
	p.active = nil
}

// 4. Zero-Allocation Programming
func zeroAllocationProgramming() {
	fmt.Println("4. Zero-Allocation Programming:")
	fmt.Println("Techniques to minimize allocations...")

	// String building without allocations
	stringBuildingZeroAlloc()
	
	// Slice operations without allocations
	sliceOperationsZeroAlloc()
	
	// Avoiding allocations in hot paths
	hotPathOptimization()
}

func stringBuildingZeroAlloc() {
	fmt.Println("  📊 String building without allocations:")
	
	// Bad: Multiple allocations
	bad := "Hello" + " " + "World" + "!"
	fmt.Printf("    Bad (multiple allocs): %s\n", bad)
	
	// Good: Pre-allocated string builder
	var builder strings.Builder
	builder.Grow(20) // Pre-allocate capacity
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteString("!")
	good := builder.String()
	fmt.Printf("    Good (pre-allocated): %s\n", good)
}

func sliceOperationsZeroAlloc() {
	fmt.Println("  📊 Slice operations without allocations:")
	
	// Bad: Growing slice multiple times
	bad := make([]int, 0)
	for i := 0; i < 10; i++ {
		bad = append(bad, i) // May cause multiple reallocations
	}
	fmt.Printf("    Bad (growing): %v\n", bad)
	
	// Good: Pre-allocated slice
	good := make([]int, 0, 10) // Pre-allocate capacity
	for i := 0; i < 10; i++ {
		good = append(good, i) // No reallocations
	}
	fmt.Printf("    Good (pre-allocated): %v\n", good)
}

func hotPathOptimization() {
	fmt.Println("  📊 Hot path optimization:")
	
	// Avoid allocations in frequently called functions
	for i := 0; i < 5; i++ {
		result := processWithoutAllocation(i)
		fmt.Printf("    Processed %d: %s\n", i, result)
	}
}

func processWithoutAllocation(value int) string {
	// Use pre-allocated buffer instead of creating new strings
	var buf [32]byte
	n := fmt.Sprintf(string(buf[:0]), "value-%d", value)
	return n
}

// 5. String Building Optimization
func stringBuildingOptimization() {
	fmt.Println("5. String Building Optimization:")
	fmt.Println("Optimizing string concatenation...")

	// Compare different string building methods
	compareStringBuildingMethods()
}

func compareStringBuildingMethods() {
	fmt.Println("  📊 String building method comparison:")
	
	parts := []string{"Hello", " ", "beautiful", " ", "world", "!"}
	
	// Method 1: Simple concatenation (worst)
	start := time.Now()
	result1 := ""
	for _, part := range parts {
		result1 += part
	}
	time1 := time.Since(start)
	fmt.Printf("    Simple concatenation: %s (took %v)\n", result1, time1)
	
	// Method 2: strings.Builder (better)
	start = time.Now()
	var builder strings.Builder
	builder.Grow(50) // Pre-allocate
	for _, part := range parts {
		builder.WriteString(part)
	}
	result2 := builder.String()
	time2 := time.Since(start)
	fmt.Printf("    strings.Builder: %s (took %v)\n", result2, time2)
	
	// Method 3: strings.Join (best for known parts)
	start = time.Now()
	result3 := strings.Join(parts, "")
	time3 := time.Since(start)
	fmt.Printf("    strings.Join: %s (took %v)\n", result3, time3)
}

// 6. Slice Pre-allocation Strategies
func slicePreallocationStrategies() {
	fmt.Println("6. Slice Pre-allocation Strategies:")
	fmt.Println("Optimizing slice operations...")

	// Demonstrate different pre-allocation strategies
	slicePreallocationExamples()
}

func slicePreallocationExamples() {
	fmt.Println("  📊 Slice pre-allocation examples:")
	
	// Strategy 1: Known size
	knownSize := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		knownSize = append(knownSize, i)
	}
	fmt.Printf("    Known size (100): len=%d, cap=%d\n", len(knownSize), cap(knownSize))
	
	// Strategy 2: Estimated size
	estimated := make([]int, 0, 50) // Estimate 50, actual might be different
	for i := 0; i < 75; i++ {
		estimated = append(estimated, i)
	}
	fmt.Printf("    Estimated size (50->75): len=%d, cap=%d\n", len(estimated), cap(estimated))
	
	// Strategy 3: Dynamic growth with factor
	dynamic := make([]int, 0, 4)
	for i := 0; i < 20; i++ {
		if len(dynamic) == cap(dynamic) {
			// Double capacity when full
			newSlice := make([]int, len(dynamic), cap(dynamic)*2)
			copy(newSlice, dynamic)
			dynamic = newSlice
		}
		dynamic = append(dynamic, i)
	}
	fmt.Printf("    Dynamic growth: len=%d, cap=%d\n", len(dynamic), cap(dynamic))
}

// 7. Memory Pool Patterns
func memoryPoolPatterns() {
	fmt.Println("7. Memory Pool Patterns:")
	fmt.Println("Advanced memory pooling patterns...")

	// Object pool pattern
	objectPoolPattern()
	
	// Buffer pool pattern
	bufferPoolPattern()
	
	// Connection pool pattern
	connectionPoolPattern()
}

func objectPoolPattern() {
	fmt.Println("  📊 Object pool pattern:")
	
	pool := NewObjectPool()
	
	// Get objects from pool
	obj1 := pool.Get()
	obj2 := pool.Get()
	
	fmt.Printf("    Got objects: %+v, %+v\n", obj1, obj2)
	
	// Use objects
	obj1.Process("task1")
	obj2.Process("task2")
	
	// Return to pool
	pool.Put(obj1)
	pool.Put(obj2)
	fmt.Println("    Objects returned to pool")
}

type Object struct {
	ID   int
	Data string
}

func (o *Object) Process(task string) {
	o.Data = task
	fmt.Printf("    Processing %s with object %d\n", task, o.ID)
}

type ObjectPool struct {
	pool sync.Pool
	nextID int
	mu    sync.Mutex
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &Object{}
			},
		},
	}
}

func (p *ObjectPool) Get() *Object {
	p.mu.Lock()
	p.nextID++
	id := p.nextID
	p.mu.Unlock()
	
	obj := p.pool.Get().(*Object)
	obj.ID = id
	obj.Data = ""
	return obj
}

func (p *ObjectPool) Put(obj *Object) {
	obj.ID = 0
	obj.Data = ""
	p.pool.Put(obj)
}

func bufferPoolPattern() {
	fmt.Println("  📊 Buffer pool pattern:")
	
	pool := NewBufferPool(1024)
	
	// Get buffers
	buf1 := pool.Get()
	buf2 := pool.Get()
	
	fmt.Printf("    Got buffers: len=%d, len=%d\n", len(buf1), len(buf2))
	
	// Use buffers
	copy(buf1, []byte("Buffer 1 data"))
	copy(buf2, []byte("Buffer 2 data"))
	
	fmt.Printf("    Buffer 1: %s\n", string(buf1[:13]))
	fmt.Printf("    Buffer 2: %s\n", string(buf2[:13]))
	
	// Return buffers
	pool.Put(buf1)
	pool.Put(buf2)
	fmt.Println("    Buffers returned to pool")
}

type BufferPool struct {
	pool sync.Pool
	size int
}

func NewBufferPool(size int) *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
		size: size,
	}
}

func (p *BufferPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *BufferPool) Put(buf []byte) {
	if len(buf) == p.size {
		p.pool.Put(buf)
	}
}

func connectionPoolPattern() {
	fmt.Println("  📊 Connection pool pattern:")
	
	pool := NewConnectionPool(3)
	
	// Get connections
	conn1 := pool.Get()
	conn2 := pool.Get()
	
	fmt.Printf("    Got connections: %s, %s\n", conn1.ID, conn2.ID)
	
	// Use connections
	conn1.Execute("SELECT 1")
	conn2.Execute("SELECT 2")
	
	// Return connections
	pool.Put(conn1)
	pool.Put(conn2)
	fmt.Println("    Connections returned to pool")
}

type Connection struct {
	ID   string
	Busy bool
}

func (c *Connection) Execute(query string) {
	fmt.Printf("    Executing %s on connection %s\n", query, c.ID)
}

type ConnectionPool struct {
	pool   sync.Pool
	active map[string]*Connection
	mu     sync.Mutex
	nextID int
}

func NewConnectionPool(maxSize int) *ConnectionPool {
	return &ConnectionPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &Connection{}
			},
		},
		active: make(map[string]*Connection),
	}
}

func (p *ConnectionPool) Get() *Connection {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	p.nextID++
	id := fmt.Sprintf("conn-%d", p.nextID)
	
	conn := p.pool.Get().(*Connection)
	conn.ID = id
	conn.Busy = true
	
	p.active[id] = conn
	return conn
}

func (p *ConnectionPool) Put(conn *Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	conn.Busy = false
	delete(p.active, conn.ID)
	p.pool.Put(conn)
}

// 8. Advanced Memory Techniques
func advancedMemoryTechniques() {
	fmt.Println("8. Advanced Memory Techniques:")
	fmt.Println("Advanced memory optimization techniques...")

	// Memory alignment
	memoryAlignment()
	
	// Cache-friendly data structures
	cacheFriendlyStructures()
	
	// Memory-mapped files
	memoryMappedFiles()
}

func memoryAlignment() {
	fmt.Println("  📊 Memory alignment:")
	
	// Show struct alignment
	type AlignedStruct struct {
		Field1 int32  // 4 bytes
		Field2 int64  // 8 bytes
		Field3 int32  // 4 bytes
	}
	
	type PackedStruct struct {
		Field1 int32  // 4 bytes
		Field3 int32  // 4 bytes
		Field2 int64  // 8 bytes
	}
	
	aligned := AlignedStruct{}
	packed := PackedStruct{}
	
	fmt.Printf("    Aligned struct size: %d bytes\n", unsafe.Sizeof(aligned))
	fmt.Printf("    Packed struct size: %d bytes\n", unsafe.Sizeof(packed))
}

func cacheFriendlyStructures() {
	fmt.Println("  📊 Cache-friendly structures:")
	
	// Array of structs (AoS) - less cache friendly
	type PointAoS struct {
		X, Y, Z float64
		Color   int32
	}
	
	pointsAoS := make([]PointAoS, 1000)
	fmt.Printf("    Array of Structs: %d points, %d bytes per point\n", 
		len(pointsAoS), unsafe.Sizeof(PointAoS{}))
	
	// Struct of arrays (SoA) - more cache friendly
	type PointsSoA struct {
		X     []float64
		Y     []float64
		Z     []float64
		Color []int32
	}
	
	pointsSoA := PointsSoA{
		X:     make([]float64, 1000),
		Y:     make([]float64, 1000),
		Z:     make([]float64, 1000),
		Color: make([]int32, 1000),
	}
	fmt.Printf("    Struct of Arrays: %d points, %d bytes per point\n", 
		len(pointsSoA.X), unsafe.Sizeof(pointsSoA.X[0])*4)
}

func memoryMappedFiles() {
	fmt.Println("  📊 Memory-mapped files:")
	fmt.Println("    Note: Memory-mapped files require os package")
	fmt.Println("    Use mmap for large file processing")
	fmt.Println("    Reduces memory usage for large datasets")
}

// 9. Performance Comparison
func performanceComparison() {
	fmt.Println("9. Performance Comparison:")
	fmt.Println("Comparing different memory optimization techniques...")

	// Compare allocation strategies
	compareAllocationStrategies()
	
	// Compare string building methods
	compareStringBuildingPerformance()
	
	// Compare slice operations
	compareSliceOperations()
}

func compareAllocationStrategies() {
	fmt.Println("  📊 Allocation strategy comparison:")
	
	const iterations = 100000
	
	// Strategy 1: No pre-allocation
	start := time.Now()
	var slice1 []int
	for i := 0; i < iterations; i++ {
		slice1 = append(slice1, i)
	}
	time1 := time.Since(start)
	fmt.Printf("    No pre-allocation: %v\n", time1)
	
	// Strategy 2: Pre-allocated
	start = time.Now()
	slice2 := make([]int, 0, iterations)
	for i := 0; i < iterations; i++ {
		slice2 = append(slice2, i)
	}
	time2 := time.Since(start)
	fmt.Printf("    Pre-allocated: %v\n", time2)
	
	// Strategy 3: Exact size
	start = time.Now()
	slice3 := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		slice3[i] = i
	}
	time3 := time.Since(start)
	fmt.Printf("    Exact size: %v\n", time3)
	
	fmt.Printf("    Pre-allocated is %.2fx faster than no pre-allocation\n", 
		float64(time1)/float64(time2))
}

func compareStringBuildingPerformance() {
	fmt.Println("  📊 String building performance comparison:")
	
	const iterations = 10000
	parts := []string{"Hello", " ", "World", " ", "!", " ", "Go", " ", "is", " ", "awesome"}
	
	// Method 1: Simple concatenation
	start := time.Now()
	result1 := ""
	for i := 0; i < iterations; i++ {
		for _, part := range parts {
			result1 += part
		}
	}
	time1 := time.Since(start)
	fmt.Printf("    Simple concatenation: %v\n", time1)
	
	// Method 2: strings.Builder
	start = time.Now()
	for i := 0; i < iterations; i++ {
		var builder strings.Builder
		builder.Grow(100)
		for _, part := range parts {
			builder.WriteString(part)
		}
		_ = builder.String()
	}
	time2 := time.Since(start)
	fmt.Printf("    strings.Builder: %v\n", time2)
	
	// Method 3: strings.Join
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = strings.Join(parts, "")
	}
	time3 := time.Since(start)
	fmt.Printf("    strings.Join: %v\n", time3)
	
	fmt.Printf("    strings.Builder is %.2fx faster than concatenation\n", 
		float64(time1)/float64(time2))
}

func compareSliceOperations() {
	fmt.Println("  📊 Slice operations performance comparison:")
	
	const size = 100000
	
	// Method 1: Growing slice
	start := time.Now()
	slice1 := make([]int, 0)
	for i := 0; i < size; i++ {
		slice1 = append(slice1, i)
	}
	time1 := time.Since(start)
	fmt.Printf("    Growing slice: %v\n", time1)
	
	// Method 2: Pre-allocated slice
	start = time.Now()
	slice2 := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice2 = append(slice2, i)
	}
	time2 := time.Since(start)
	fmt.Printf("    Pre-allocated slice: %v\n", time2)
	
	// Method 3: Direct assignment
	start = time.Now()
	slice3 := make([]int, size)
	for i := 0; i < size; i++ {
		slice3[i] = i
	}
	time3 := time.Since(start)
	fmt.Printf("    Direct assignment: %v\n", time3)
	
	fmt.Printf("    Pre-allocated is %.2fx faster than growing\n", 
		float64(time1)/float64(time2))
}

// 10. Best Practices
func memoryOptimizationBestPractices() {
	fmt.Println("10. Memory Optimization Best Practices:")
	fmt.Println("Best practices for memory optimization...")

	fmt.Println("  📝 Best Practice 1: Profile before optimizing")
	fmt.Println("    - Use pprof to identify memory bottlenecks")
	fmt.Println("    - Focus on hot paths and frequent allocations")
	
	fmt.Println("  📝 Best Practice 2: Pre-allocate when possible")
	fmt.Println("    - Use make([]T, 0, capacity) for slices")
	fmt.Println("    - Use strings.Builder with Grow() for strings")
	
	fmt.Println("  📝 Best Practice 3: Use object pools for frequent allocations")
	fmt.Println("    - sync.Pool for temporary objects")
	fmt.Println("    - Custom pools for specific use cases")
	
	fmt.Println("  📝 Best Practice 4: Avoid unnecessary heap allocations")
	fmt.Println("    - Keep data on stack when possible")
	fmt.Println("    - Avoid returning pointers to local data")
	
	fmt.Println("  📝 Best Practice 5: Use appropriate data structures")
	fmt.Println("    - Choose cache-friendly layouts")
	fmt.Println("    - Consider memory alignment")
	
	fmt.Println("  📝 Best Practice 6: Monitor memory usage")
	fmt.Println("    - Use runtime.MemStats for monitoring")
	fmt.Println("    - Set memory limits with debug.SetMemoryLimit")
	
	fmt.Println("  📝 Best Practice 7: Optimize for your use case")
	fmt.Println("    - Different optimizations for different scenarios")
	fmt.Println("    - Balance memory usage vs performance")
}
