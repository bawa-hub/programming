package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// 🧠 GARBAGE COLLECTION MASTERY
// Understanding Go's garbage collector and memory management

func main() {
	fmt.Println("🧠 GARBAGE COLLECTION MASTERY")
	fmt.Println("=============================")

	// 1. Understanding GC Basics
	fmt.Println("\n1. GC Basics:")
	gcBasics()

	// 2. GC Tuning and Configuration
	fmt.Println("\n2. GC Tuning:")
	gcTuning()

	// 3. Memory Allocation Patterns
	fmt.Println("\n3. Memory Allocation Patterns:")
	memoryAllocationPatterns()

	// 4. Stack vs Heap Allocation
	fmt.Println("\n4. Stack vs Heap Allocation:")
	stackVsHeapAllocation()

	// 5. Escape Analysis
	fmt.Println("\n5. Escape Analysis:")
	escapeAnalysis()

	// 6. Memory Pressure and GC Impact
	fmt.Println("\n6. Memory Pressure and GC Impact:")
	memoryPressureImpact()

	// 7. GC Best Practices
	fmt.Println("\n7. GC Best Practices:")
	gcBestPractices()
}

// GC BASICS: Understanding garbage collection fundamentals
func gcBasics() {
	fmt.Println("Understanding GC basics...")
	
	// Get current GC stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  📊 Current Memory Stats:\n")
	fmt.Printf("    Alloc: %d bytes (%.2f MB)\n", m.Alloc, float64(m.Alloc)/1024/1024)
	fmt.Printf("    TotalAlloc: %d bytes (%.2f MB)\n", m.TotalAlloc, float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("    Sys: %d bytes (%.2f MB)\n", m.Sys, float64(m.Sys)/1024/1024)
	fmt.Printf("    NumGC: %d\n", m.NumGC)
	fmt.Printf("    GCCPUFraction: %.6f\n", m.GCCPUFraction)
	
	// Force a GC cycle
	fmt.Println("  🔄 Forcing GC cycle...")
	runtime.GC()
	
	// Get stats after GC
	runtime.ReadMemStats(&m)
	fmt.Printf("  📊 After GC:\n")
	fmt.Printf("    Alloc: %d bytes (%.2f MB)\n", m.Alloc, float64(m.Alloc)/1024/1024)
	fmt.Printf("    NumGC: %d\n", m.NumGC)
}

// GC TUNING: Understanding GC configuration
func gcTuning() {
	fmt.Println("Understanding GC tuning...")
	
	// Get current GC target percentage
	targetPercent := debug.SetGCPercent(100)
	fmt.Printf("  📊 Previous GC target: %d%%\n", targetPercent)
	
	// Set new GC target
	newTarget := debug.SetGCPercent(50)
	fmt.Printf("  📊 New GC target: %d%%\n", newTarget)
	
	// Restore original target
	debug.SetGCPercent(targetPercent)
	fmt.Printf("  📊 Restored GC target: %d%%\n", targetPercent)
	
	// Memory limit (Go 1.19+)
	fmt.Println("  📊 Memory limit features:")
	fmt.Println("    - SetMemoryLimit() for soft memory limit")
	fmt.Println("    - GOMEMLIMIT environment variable")
	fmt.Println("    - Automatic GC tuning based on limit")
}

// MEMORY ALLOCATION PATTERNS: Understanding allocation behavior
func memoryAllocationPatterns() {
	fmt.Println("Understanding memory allocation patterns...")
	
	// Pattern 1: Small allocations
	fmt.Println("  📊 Pattern 1: Small allocations")
	smallAllocations()
	
	// Pattern 2: Large allocations
	fmt.Println("  📊 Pattern 2: Large allocations")
	largeAllocations()
	
	// Pattern 3: Slice growth patterns
	fmt.Println("  📊 Pattern 3: Slice growth patterns")
	sliceGrowthPatterns()
	
	// Pattern 4: String concatenation
	fmt.Println("  📊 Pattern 4: String concatenation")
	stringConcatenationPatterns()
}

func smallAllocations() {
	// Create many small objects
	objects := make([]*SmallObject, 1000)
	for i := 0; i < 1000; i++ {
		objects[i] = &SmallObject{
			ID:    i,
			Value: fmt.Sprintf("object-%d", i),
		}
	}
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("    Small objects allocated: %d\n", len(objects))
	fmt.Printf("    Memory used: %d bytes\n", m.Alloc)
}

func largeAllocations() {
	// Create large objects
	largeObjects := make([]*LargeObject, 10)
	for i := 0; i < 10; i++ {
		largeObjects[i] = &LargeObject{
			ID:    i,
			Data:  make([]byte, 1024*1024), // 1MB each
		}
	}
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("    Large objects allocated: %d\n", len(largeObjects))
	fmt.Printf("    Memory used: %d bytes (%.2f MB)\n", m.Alloc, float64(m.Alloc)/1024/1024)
}

func sliceGrowthPatterns() {
	// Demonstrate slice growth
	slice := make([]int, 0, 1)
	fmt.Printf("    Initial capacity: %d\n", cap(slice))
	
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		if i%100 == 0 {
			fmt.Printf("    Length: %d, Capacity: %d\n", len(slice), cap(slice))
		}
	}
}

func stringConcatenationPatterns() {
	// Bad: String concatenation in loop
	start := time.Now()
	var badResult string
	for i := 0; i < 1000; i++ {
		badResult += fmt.Sprintf("item-%d ", i)
	}
	badTime := time.Since(start)
	
	// Good: Using strings.Builder
	start = time.Now()
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("item-%d ", i))
	}
	goodResult := builder.String()
	goodTime := time.Since(start)
	
	fmt.Printf("    Bad concatenation time: %v\n", badTime)
	fmt.Printf("    Good concatenation time: %v\n", goodTime)
	fmt.Printf("    Results equal: %t\n", badResult == goodResult)
}

// STACK VS HEAP ALLOCATION: Understanding allocation decisions
func stackVsHeapAllocation() {
	fmt.Println("Understanding stack vs heap allocation...")
	
	// This will likely be allocated on stack
	stackAllocated := createStackObject()
	fmt.Printf("  📊 Stack allocated object: %+v\n", stackAllocated)
	
	// This will likely be allocated on heap
	heapAllocated := createHeapObject()
	fmt.Printf("  📊 Heap allocated object: %+v\n", heapAllocated)
	
	// Demonstrate escape analysis
	escapeAnalysisDemo()
}

func createStackObject() StackObject {
	// This should stay on stack
	return StackObject{
		Value: 42,
		Name:  "stack-object",
	}
}

func createHeapObject() *HeapObject {
	// This will escape to heap
	return &HeapObject{
		Value: 42,
		Name:  "heap-object",
	}
}

func escapeAnalysisDemo() {
	// This function demonstrates escape analysis
	obj := &EscapeObject{
		Value: 100,
		Name:  "escape-demo",
	}
	
	// The object escapes because we return a pointer to it
	// Go's escape analysis will move it to heap
	fmt.Printf("  📊 Escape object: %+v\n", obj)
}

// ESCAPE ANALYSIS: Understanding when objects escape to heap
func escapeAnalysis() {
	fmt.Println("Understanding escape analysis...")
	
	// Case 1: No escape - stays on stack
	noEscape()
	
	// Case 2: Escape due to return
	escapeReturn()
	
	// Case 3: Escape due to interface
	escapeInterface()
	
	// Case 4: Escape due to closure
	escapeClosure()
}

func noEscape() {
	// This object should stay on stack
	obj := struct {
		value int
		name  string
	}{
		value: 42,
		name:  "no-escape",
	}
	
	fmt.Printf("  📊 No escape object: %+v\n", obj)
}

func escapeReturn() *struct {
	value int
	name  string
} {
	// This object escapes to heap because we return a pointer
	return &struct {
		value int
		name  string
	}{
		value: 42,
		name:  "escape-return",
	}
}

func escapeInterface() {
	// This object escapes to heap because it's assigned to interface
	var iface interface{} = struct {
		value int
		name  string
	}{
		value: 42,
		name:  "escape-interface",
	}
	
	fmt.Printf("  📊 Interface object: %+v\n", iface)
}

func escapeClosure() {
	// This object escapes to heap because it's captured by closure
	obj := struct {
		value int
		name  string
	}{
		value: 42,
		name:  "escape-closure",
	}
	
	// Closure captures obj, causing it to escape
	func() {
		fmt.Printf("  📊 Closure object: %+v\n", obj)
	}()
}

// MEMORY PRESSURE IMPACT: Understanding GC under pressure
func memoryPressureImpact() {
	fmt.Println("Understanding memory pressure impact...")
	
	// Create memory pressure
	fmt.Println("  📊 Creating memory pressure...")
	createMemoryPressure()
	
	// Monitor GC behavior
	fmt.Println("  📊 Monitoring GC behavior...")
	monitorGCBehavior()
}

func createMemoryPressure() {
	// Allocate memory in chunks
	chunks := make([][]byte, 0, 100)
	
	for i := 0; i < 100; i++ {
		// Allocate 1MB chunks
		chunk := make([]byte, 1024*1024)
		chunks = append(chunks, chunk)
		
		if i%10 == 0 {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("    Chunk %d: Alloc=%.2f MB, NumGC=%d\n", 
				i, float64(m.Alloc)/1024/1024, m.NumGC)
		}
	}
	
	// Release some memory
	chunks = chunks[:50]
	runtime.GC()
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("    After cleanup: Alloc=%.2f MB, NumGC=%d\n", 
		float64(m.Alloc)/1024/1024, m.NumGC)
}

func monitorGCBehavior() {
	// Monitor GC for a period
	start := time.Now()
	lastGC := uint32(0)
	
	for time.Since(start) < 2*time.Second {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		
		if m.NumGC > lastGC {
			fmt.Printf("    GC cycle %d: PauseTotal=%v, PauseNs=%v\n", 
				m.NumGC, m.PauseTotalNs, m.PauseNs[(m.NumGC+255)%256])
			lastGC = m.NumGC
		}
		
		time.Sleep(10 * time.Millisecond)
	}
}

// GC BEST PRACTICES: Following Go GC best practices
func gcBestPractices() {
	fmt.Println("Understanding GC best practices...")
	
	// 1. Minimize allocations
	fmt.Println("  📝 Best Practice 1: Minimize allocations")
	minimizeAllocations()
	
	// 2. Use object pooling
	fmt.Println("  📝 Best Practice 2: Use object pooling")
	objectPooling()
	
	// 3. Avoid unnecessary string operations
	fmt.Println("  📝 Best Practice 3: Avoid unnecessary string operations")
	stringOptimization()
	
	// 4. Use appropriate data structures
	fmt.Println("  📝 Best Practice 4: Use appropriate data structures")
	dataStructureOptimization()
}

func minimizeAllocations() {
	// Bad: Allocating in loop
	badResult := make([]string, 0)
	for i := 0; i < 1000; i++ {
		badResult = append(badResult, fmt.Sprintf("item-%d", i))
	}
	
	// Good: Pre-allocate with known size
	goodResult := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		goodResult[i] = fmt.Sprintf("item-%d", i)
	}
	
	fmt.Printf("    Bad result length: %d\n", len(badResult))
	fmt.Printf("    Good result length: %d\n", len(goodResult))
}

func objectPooling() {
	// Create object pool
	pool := &sync.Pool{
		New: func() interface{} {
			return &PooledObject{
				Data: make([]byte, 1024),
			}
		},
	}
	
	// Get object from pool
	obj := pool.Get().(*PooledObject)
	obj.ID = 42
	obj.Data[0] = 1
	
	// Use object
	fmt.Printf("    Pooled object: ID=%d, Data[0]=%d\n", obj.ID, obj.Data[0])
	
	// Return to pool
	pool.Put(obj)
	
	// Get another object (might be the same one)
	obj2 := pool.Get().(*PooledObject)
	fmt.Printf("    Reused object: ID=%d, Data[0]=%d\n", obj2.ID, obj2.Data[0])
}

func stringOptimization() {
	// Bad: String concatenation
	badResult := ""
	for i := 0; i < 1000; i++ {
		badResult += fmt.Sprintf("item-%d ", i)
	}
	
	// Good: Using strings.Builder
	var builder strings.Builder
	builder.Grow(1000 * 10) // Pre-allocate capacity
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("item-%d ", i))
	}
	goodResult := builder.String()
	
	fmt.Printf("    Bad result length: %d\n", len(badResult))
	fmt.Printf("    Good result length: %d\n", len(goodResult))
}

func dataStructureOptimization() {
	// Use appropriate data structures
	// For small collections, use arrays
	smallArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// For dynamic collections, use slices with pre-allocation
	largeSlice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		largeSlice = append(largeSlice, i)
	}
	
	// For key-value lookups, use maps
	lookupMap := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		lookupMap[fmt.Sprintf("key-%d", i)] = i
	}
	
	fmt.Printf("    Small array length: %d\n", len(smallArray))
	fmt.Printf("    Large slice length: %d\n", len(largeSlice))
	fmt.Printf("    Lookup map size: %d\n", len(lookupMap))
}

// Data structures for examples
type SmallObject struct {
	ID    int
	Value string
}

type LargeObject struct {
	ID   int
	Data []byte
}

type StackObject struct {
	Value int
	Name  string
}

type HeapObject struct {
	Value int
	Name  string
}

type EscapeObject struct {
	Value int
	Name  string
}

type PooledObject struct {
	ID   int
	Data []byte
}
