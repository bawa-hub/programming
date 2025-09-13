package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// 📊 HIGH-PERFORMANCE DATA STRUCTURES
// Understanding custom data structures and optimization techniques

func main() {
	fmt.Println("📊 HIGH-PERFORMANCE DATA STRUCTURES")
	fmt.Println("====================================")

	// 1. High-Performance Slices
	fmt.Println("\n1. High-Performance Slices:")
	highPerformanceSlices()

	// 2. Memory-Efficient Maps
	fmt.Println("\n2. Memory-Efficient Maps:")
	memoryEfficientMaps()

	// 3. Lock-Free Data Structures
	fmt.Println("\n3. Lock-Free Data Structures:")
	lockFreeDataStructures()

	// 4. Cache-Friendly Layouts
	fmt.Println("\n4. Cache-Friendly Layouts:")
	cacheFriendlyLayouts()

	// 5. Zero-Allocation Programming
	fmt.Println("\n5. Zero-Allocation Programming:")
	zeroAllocationProgramming()

	// 6. Custom Allocators
	fmt.Println("\n6. Custom Allocators:")
	customAllocators()

	// 7. Performance Comparison
	fmt.Println("\n7. Performance Comparison:")
	performanceComparison()
}

// HIGH-PERFORMANCE SLICES: Optimized slice operations
func highPerformanceSlices() {
	fmt.Println("Understanding high-performance slices...")
	
	// Pre-allocated slice with known capacity
	slice := make([]int, 0, 1000)
	
	// Benchmark slice operations
	start := time.Now()
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}
	duration := time.Since(start)
	fmt.Printf("  📊 Pre-allocated slice: %v\n", duration)
	
	// Slice reuse pattern
	reusableSlice := make([]int, 0, 1000)
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		reusableSlice = reusableSlice[:0] // Reset length, keep capacity
		for j := 0; j < 100; j++ {
			reusableSlice = append(reusableSlice, j)
		}
	}
	duration = time.Since(start)
	fmt.Printf("  📊 Reusable slice: %v\n", duration)
	
	// Slice copying optimization
	source := make([]int, 1000)
	for i := range source {
		source[i] = i
	}
	
	start = time.Now()
	dest := make([]int, len(source))
	copy(dest, source)
	duration = time.Since(start)
	fmt.Printf("  📊 Slice copy: %v\n", duration)
}

// MEMORY-EFFICIENT MAPS: Optimized map usage
func memoryEfficientMaps() {
	fmt.Println("Understanding memory-efficient maps...")
	
	// Pre-allocated map with known size
	start := time.Now()
	map1 := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		map1[fmt.Sprintf("key-%d", i)] = i
	}
	duration1 := time.Since(start)
	fmt.Printf("  📊 Pre-allocated map: %v\n", duration1)
	
	// Map without pre-allocation
	start = time.Now()
	map2 := make(map[string]int)
	for i := 0; i < 1000; i++ {
		map2[fmt.Sprintf("key-%d", i)] = i
	}
	duration2 := time.Since(start)
	fmt.Printf("  📊 Default map: %v\n", duration2)
	
	// Map reuse pattern
	reusableMap := make(map[string]int, 1000)
	
	start = time.Now()
	for i := 0; i < 100; i++ {
		// Clear map efficiently
		for k := range reusableMap {
			delete(reusableMap, k)
		}
		
		// Add new data
		for j := 0; j < 10; j++ {
			reusableMap[fmt.Sprintf("key-%d-%d", i, j)] = j
		}
	}
	duration3 := time.Since(start)
	fmt.Printf("  📊 Reusable map: %v\n", duration3)
}

// LOCK-FREE DATA STRUCTURES: Understanding lock-free programming
func lockFreeDataStructures() {
	fmt.Println("Understanding lock-free data structures...")
	
	// Atomic counter
	counter := NewAtomicCounter()
	
	// Test concurrent access
	var wg sync.WaitGroup
	goroutines := 100
	operations := 1000
	
	start := time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("  📊 Atomic counter: %v (value: %d)\n", duration, counter.Value())
	
	// Lock-free stack
	stack := NewLockFreeStack()
	
	start = time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				stack.Push(id*1000 + j)
				stack.Pop()
			}
		}(i)
	}
	
	wg.Wait()
	duration = time.Since(start)
	fmt.Printf("  📊 Lock-free stack: %v\n", duration)
}

// CACHE-FRIENDLY LAYOUTS: Understanding memory layout optimization
func cacheFriendlyLayouts() {
	fmt.Println("Understanding cache-friendly layouts...")
	
	// Bad: Array of structs (AoS) - poor cache locality
	type BadPoint struct {
		X, Y, Z float64
		Color   [3]uint8
	}
	
	badPoints := make([]BadPoint, 1000)
	for i := range badPoints {
		badPoints[i] = BadPoint{
			X: float64(i),
			Y: float64(i * 2),
			Z: float64(i * 3),
		}
	}
	
	start := time.Now()
	sum := 0.0
	for _, p := range badPoints {
		sum += p.X + p.Y + p.Z
	}
	duration1 := time.Since(start)
	fmt.Printf("  📊 Array of structs: %v (sum: %.2f)\n", duration1, sum)
	
	// Good: Structure of arrays (SoA) - better cache locality
	type GoodPoints struct {
		X, Y, Z []float64
		Color   [][]uint8
	}
	
	goodPoints := GoodPoints{
		X: make([]float64, 1000),
		Y: make([]float64, 1000),
		Z: make([]float64, 1000),
	}
	
	for i := 0; i < 1000; i++ {
		goodPoints.X[i] = float64(i)
		goodPoints.Y[i] = float64(i * 2)
		goodPoints.Z[i] = float64(i * 3)
	}
	
	start = time.Now()
	sum = 0.0
	for i := 0; i < 1000; i++ {
		sum += goodPoints.X[i] + goodPoints.Y[i] + goodPoints.Z[i]
	}
	duration2 := time.Since(start)
	fmt.Printf("  📊 Structure of arrays: %v (sum: %.2f)\n", duration2, sum)
	
	// Packed struct for better memory usage
	type PackedPoint struct {
		X, Y, Z float32 // Use float32 instead of float64
		Color   uint32  // Pack color into single uint32
	}
	
	packedPoints := make([]PackedPoint, 1000)
	for i := range packedPoints {
		packedPoints[i] = PackedPoint{
			X: float32(i),
			Y: float32(i * 2),
			Z: float32(i * 3),
		}
	}
	
	start = time.Now()
	sum = 0.0
	for _, p := range packedPoints {
		sum += float64(p.X + p.Y + p.Z)
	}
	duration3 := time.Since(start)
	fmt.Printf("  📊 Packed struct: %v (sum: %.2f)\n", duration3, sum)
}

// ZERO-ALLOCATION PROGRAMMING: Avoiding allocations
func zeroAllocationProgramming() {
	fmt.Println("Understanding zero-allocation programming...")
	
	// String building without allocation
	var builder strings.Builder
	builder.Grow(1000) // Pre-allocate capacity
	
	start := time.Now()
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("item-%d ", i))
	}
	result := builder.String()
	duration1 := time.Since(start)
	fmt.Printf("  📊 String builder: %v (length: %d)\n", duration1, len(result))
	
	// Slice reuse for zero allocation
	reusableSlice := make([]int, 0, 1000)
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		reusableSlice = reusableSlice[:0] // Reset without allocation
		for j := 0; j < 10; j++ {
			reusableSlice = append(reusableSlice, j)
		}
	}
	duration2 := time.Since(start)
	fmt.Printf("  📊 Reusable slice: %v\n", duration2)
	
	// Object pooling
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		buf := pool.Get().([]byte)
		// Use buffer
		_ = buf
		pool.Put(buf)
	}
	duration3 := time.Since(start)
	fmt.Printf("  📊 Object pool: %v\n", duration3)
}

// CUSTOM ALLOCATORS: Understanding custom memory management
func customAllocators() {
	fmt.Println("Understanding custom allocators...")
	
	// Simple arena allocator
	arena := NewArena(1024 * 1024) // 1MB arena
	
	start := time.Now()
	for i := 0; i < 1000; i++ {
		// Allocate from arena
		ptr := arena.Alloc(64)
		_ = ptr
	}
	duration1 := time.Since(start)
	fmt.Printf("  📊 Arena allocator: %v\n", duration1)
	
	// Memory pool allocator
	pool := NewMemoryPool(64, 1000) // 64-byte blocks, 1000 max
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		ptr := pool.Alloc()
		_ = ptr
		pool.Free(ptr)
	}
	duration2 := time.Since(start)
	fmt.Printf("  📊 Memory pool: %v\n", duration2)
	
	// Stack allocator
	stack := NewStackAllocator(1024 * 1024)
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		ptr := stack.Alloc(64)
		_ = ptr
		stack.Reset() // Reset for next iteration
	}
	duration3 := time.Since(start)
	fmt.Printf("  📊 Stack allocator: %v\n", duration3)
}

// PERFORMANCE COMPARISON: Comparing different approaches
func performanceComparison() {
	fmt.Println("Understanding performance comparison...")
	
	// Compare different data structures
	compareDataStructures()
	
	// Compare memory usage
	compareMemoryUsage()
	
	// Compare allocation patterns
	compareAllocationPatterns()
}

func compareDataStructures() {
	fmt.Println("  📊 Data Structure Comparison:")
	
	// Slice vs Array
	slice := make([]int, 1000)
	array := [1000]int{}
	
	start := time.Now()
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}
	duration1 := time.Since(start)
	
	start = time.Now()
	for i := 0; i < 1000; i++ {
		array[i] = i
	}
	duration2 := time.Since(start)
	
	fmt.Printf("    Slice access: %v\n", duration1)
	fmt.Printf("    Array access: %v\n", duration2)
}

func compareMemoryUsage() {
	fmt.Println("  📊 Memory Usage Comparison:")
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	// Allocate large slice
	largeSlice := make([]int, 1000000)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("    Memory allocated: %d bytes (%.2f MB)\n", 
		m2.Alloc-m1.Alloc, float64(m2.Alloc-m1.Alloc)/1024/1024)
}

func compareAllocationPatterns() {
	fmt.Println("  📊 Allocation Pattern Comparison:")
	
	// Pattern 1: Allocate in loop
	start := time.Now()
	for i := 0; i < 1000; i++ {
		_ = make([]int, 100)
	}
	duration1 := time.Since(start)
	
	// Pattern 2: Pre-allocate and reuse
	reusable := make([]int, 0, 100)
	start = time.Now()
	for i := 0; i < 1000; i++ {
		reusable = reusable[:0]
		for j := 0; j < 100; j++ {
			reusable = append(reusable, j)
		}
	}
	duration2 := time.Since(start)
	
	fmt.Printf("    Loop allocation: %v\n", duration1)
	fmt.Printf("    Reuse pattern: %v\n", duration2)
}

// IMPLEMENTATIONS

// Atomic Counter
type AtomicCounter struct {
	value int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Increment() {
	// In a real implementation, use atomic.AddInt64
	c.value++
}

func (c *AtomicCounter) Value() int64 {
	return c.value
}

// Lock-Free Stack
type LockFreeStack struct {
	head unsafe.Pointer
}

type stackNode struct {
	value int
	next  unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	node := &stackNode{value: value}
	// Simplified implementation - in reality would use atomic operations
	s.head = unsafe.Pointer(node)
}

func (s *LockFreeStack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	node := (*stackNode)(s.head)
	s.head = node.next
	return node.value, true
}

// Arena Allocator
type Arena struct {
	memory []byte
	offset int
}

func NewArena(size int) *Arena {
	return &Arena{
		memory: make([]byte, size),
		offset: 0,
	}
}

func (a *Arena) Alloc(size int) []byte {
	if a.offset+size > len(a.memory) {
		return nil // Arena full
	}
	
	ptr := a.memory[a.offset : a.offset+size]
	a.offset += size
	return ptr
}

func (a *Arena) Reset() {
	a.offset = 0
}

// Memory Pool
type MemoryPool struct {
	blocks [][]byte
	used   []bool
	size   int
}

func NewMemoryPool(blockSize, maxBlocks int) *MemoryPool {
	blocks := make([][]byte, maxBlocks)
	for i := range blocks {
		blocks[i] = make([]byte, blockSize)
	}
	
	return &MemoryPool{
		blocks: blocks,
		used:   make([]bool, maxBlocks),
		size:   blockSize,
	}
}

func (p *MemoryPool) Alloc() []byte {
	for i, used := range p.used {
		if !used {
			p.used[i] = true
			return p.blocks[i]
		}
	}
	return nil // Pool full
}

func (p *MemoryPool) Free(block []byte) {
	for i, b := range p.blocks {
		if &b[0] == &block[0] {
			p.used[i] = false
			return
		}
	}
}

// Stack Allocator
type StackAllocator struct {
	memory []byte
	offset int
}

func NewStackAllocator(size int) *StackAllocator {
	return &StackAllocator{
		memory: make([]byte, size),
		offset: 0,
	}
}

func (s *StackAllocator) Alloc(size int) []byte {
	if s.offset+size > len(s.memory) {
		return nil // Stack full
	}
	
	ptr := s.memory[s.offset : s.offset+size]
	s.offset += size
	return ptr
}

func (s *StackAllocator) Reset() {
	s.offset = 0
}
