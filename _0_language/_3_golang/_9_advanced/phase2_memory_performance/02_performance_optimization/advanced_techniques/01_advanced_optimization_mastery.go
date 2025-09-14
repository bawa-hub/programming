package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

// ⚡ ADVANCED OPTIMIZATION TECHNIQUES
// Mastering advanced performance optimization techniques

func main() {
	fmt.Println("⚡ ADVANCED OPTIMIZATION TECHNIQUES")
	fmt.Println("===================================")
	fmt.Println()

	// 1. SIMD Operations
	simdOperations()
	fmt.Println()

	// 2. Cache-Friendly Data Structures
	cacheFriendlyDataStructures()
	fmt.Println()

	// 3. Branch Prediction Optimization
	branchPredictionOptimization()
	fmt.Println()

	// 4. Compiler Optimizations
	compilerOptimizations()
	fmt.Println()

	// 5. Assembly and Low-Level Optimization
	assemblyAndLowLevelOptimization()
	fmt.Println()

	// 6. Vector Operations
	vectorOperations()
	fmt.Println()

	// 7. Memory Access Patterns
	memoryAccessPatterns()
	fmt.Println()

	// 8. Performance Profiling
	performanceProfiling()
	fmt.Println()

	// 9. Micro-optimizations
	microOptimizations()
	fmt.Println()

	// 10. Best Practices
	advancedOptimizationBestPractices()
}

// 1. SIMD Operations
func simdOperations() {
	fmt.Println("1. SIMD Operations:")
	fmt.Println("Understanding SIMD (Single Instruction, Multiple Data)...")

	// Demonstrate vector operations
	vectorOperationsExample()
	
	// Show SIMD-like operations in Go
	simdLikeOperations()
	
	// Performance comparison
	simdPerformanceComparison()
}

func vectorOperationsExample() {
	fmt.Println("  📊 Vector operations example:")
	
	// Simulate vector addition
	a := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	b := []float64{8, 7, 6, 5, 4, 3, 2, 1}
	
	// Scalar addition (traditional)
	scalarResult := make([]float64, len(a))
	for i := range a {
		scalarResult[i] = a[i] + b[i]
	}
	fmt.Printf("    Scalar addition: %v\n", scalarResult)
	
	// Vector addition (SIMD-like)
	vectorResult := vectorAdd(a, b)
	fmt.Printf("    Vector addition: %v\n", vectorResult)
}

func vectorAdd(a, b []float64) []float64 {
	result := make([]float64, len(a))
	
	// Process 4 elements at a time (SIMD-like)
	for i := 0; i < len(a); i += 4 {
		end := i + 4
		if end > len(a) {
			end = len(a)
		}
		
		for j := i; j < end; j++ {
			result[j] = a[j] + b[j]
		}
	}
	
	return result
}

func simdLikeOperations() {
	fmt.Println("  📊 SIMD-like operations in Go:")
	
	// Matrix multiplication (SIMD-like)
	matrixMultiplicationExample()
	
	// Dot product (SIMD-like)
	dotProductExample()
}

func matrixMultiplicationExample() {
	fmt.Println("    Matrix multiplication:")
	
	// 2x2 matrices
	a := [][]float64{{1, 2}, {3, 4}}
	b := [][]float64{{5, 6}, {7, 8}}
	
	result := matrixMultiply(a, b)
	fmt.Printf("      Result: %v\n", result)
}

func matrixMultiply(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(b[0])
	result := make([][]float64, rows)
	
	for i := range result {
		result[i] = make([]float64, cols)
		for j := range result[i] {
			for k := range a[i] {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	
	return result
}

func dotProductExample() {
	fmt.Println("    Dot product:")
	
	a := []float64{1, 2, 3, 4, 5}
	b := []float64{5, 4, 3, 2, 1}
	
	result := dotProduct(a, b)
	fmt.Printf("      Dot product: %.2f\n", result)
}

func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}
	
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	
	return sum
}

func simdPerformanceComparison() {
	fmt.Println("  📊 SIMD performance comparison:")
	
	const size = 1000000
	a := make([]float64, size)
	b := make([]float64, size)
	
	// Initialize with random data
	for i := range a {
		a[i] = rand.Float64()
		b[i] = rand.Float64()
	}
	
	// Scalar addition
	start := time.Now()
	scalarResult := make([]float64, size)
	for i := range a {
		scalarResult[i] = a[i] + b[i]
	}
	scalarTime := time.Since(start)
	
	// Vector addition (SIMD-like)
	start = time.Now()
	vectorResult := vectorAdd(a, b)
	vectorTime := time.Since(start)
	
	fmt.Printf("    Scalar time: %v\n", scalarTime)
	fmt.Printf("    Vector time: %v\n", vectorTime)
	fmt.Printf("    Speedup: %.2fx\n", float64(scalarTime)/float64(vectorTime))
	
	// Verify results are the same
	_ = vectorResult
}

// 2. Cache-Friendly Data Structures
func cacheFriendlyDataStructures() {
	fmt.Println("2. Cache-Friendly Data Structures:")
	fmt.Println("Understanding cache hierarchy and data layout...")

	// Demonstrate cache-friendly vs cache-unfriendly layouts
	cacheLayoutComparison()
	
	// Show data structure optimization
	dataStructureOptimization()
	
	// Demonstrate locality of reference
	localityOfReference()
}

func cacheLayoutComparison() {
	fmt.Println("  📊 Cache layout comparison:")
	
	// Array of Structs (AoS) - less cache friendly
	type PointAoS struct {
		X, Y, Z float64
		Color   int32
	}
	
	// Struct of Arrays (SoA) - more cache friendly
	type PointsSoA struct {
		X     []float64
		Y     []float64
		Z     []float64
		Color []int32
	}
	
	const count = 1000000
	
	// AoS performance
	start := time.Now()
	pointsAoS := make([]PointAoS, count)
	for i := range pointsAoS {
		pointsAoS[i] = PointAoS{
			X:     float64(i),
			Y:     float64(i * 2),
			Z:     float64(i * 3),
			Color: int32(i % 256),
		}
	}
	aosTime := time.Since(start)
	
	// SoA performance
	start = time.Now()
	pointsSoA := PointsSoA{
		X:     make([]float64, count),
		Y:     make([]float64, count),
		Z:     make([]float64, count),
		Color: make([]int32, count),
	}
	for i := 0; i < count; i++ {
		pointsSoA.X[i] = float64(i)
		pointsSoA.Y[i] = float64(i * 2)
		pointsSoA.Z[i] = float64(i * 3)
		pointsSoA.Color[i] = int32(i % 256)
	}
	soaTime := time.Since(start)
	
	fmt.Printf("    AoS (Array of Structs): %v\n", aosTime)
	fmt.Printf("    SoA (Struct of Arrays): %v\n", soaTime)
	fmt.Printf("    SoA is %.2fx faster\n", float64(aosTime)/float64(soaTime))
}

func dataStructureOptimization() {
	fmt.Println("  📊 Data structure optimization:")
	
	// Demonstrate padding and alignment
	paddingAndAlignment()
	
	// Show cache line optimization
	cacheLineOptimization()
}

func paddingAndAlignment() {
	fmt.Println("    Padding and alignment:")
	
	// Struct with padding
	type PaddedStruct struct {
		Field1 int8   // 1 byte
		Field2 int32  // 4 bytes (3 bytes padding)
		Field3 int8   // 1 byte (3 bytes padding)
		Field4 int64  // 8 bytes
	}
	
	// Struct without padding (packed)
	type PackedStruct struct {
		Field1 int8   // 1 byte
		Field3 int8   // 1 byte
		Field2 int32  // 4 bytes
		Field4 int64  // 8 bytes
	}
	
	padded := PaddedStruct{}
	packed := PackedStruct{}
	
	fmt.Printf("      Padded struct size: %d bytes\n", unsafe.Sizeof(padded))
	fmt.Printf("      Packed struct size: %d bytes\n", unsafe.Sizeof(packed))
	fmt.Printf("      Memory saved: %d bytes\n", unsafe.Sizeof(padded)-unsafe.Sizeof(packed))
}

func cacheLineOptimization() {
	fmt.Println("    Cache line optimization:")
	
	// Typical cache line size is 64 bytes
	const cacheLineSize = 64
	const elementsPerLine = cacheLineSize / 8 // 8 bytes per int64
	
	fmt.Printf("      Cache line size: %d bytes\n", cacheLineSize)
	fmt.Printf("      Elements per cache line: %d\n", elementsPerLine)
	fmt.Println("      Access elements in cache line order for better performance")
}

func localityOfReference() {
	fmt.Println("  📊 Locality of reference:")
	
	// Demonstrate spatial locality
	spatialLocalityExample()
	
	// Demonstrate temporal locality
	temporalLocalityExample()
}

func spatialLocalityExample() {
	fmt.Println("    Spatial locality example:")
	
	const size = 10000
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	
	// Good: Row-major order (spatial locality)
	start := time.Now()
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += matrix[i][j]
		}
	}
	rowMajorTime := time.Since(start)
	
	// Bad: Column-major order (poor spatial locality)
	start = time.Now()
	sum = 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			sum += matrix[i][j]
		}
	}
	columnMajorTime := time.Since(start)
	
	fmt.Printf("      Row-major order: %v\n", rowMajorTime)
	fmt.Printf("      Column-major order: %v\n", columnMajorTime)
	fmt.Printf("      Row-major is %.2fx faster\n", float64(columnMajorTime)/float64(rowMajorTime))
}

func temporalLocalityExample() {
	fmt.Println("    Temporal locality example:")
	
	// Demonstrate temporal locality with repeated access
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	// Good: Access same data multiple times
	start := time.Now()
	sum := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			sum += data[i%100] // Access same elements repeatedly
		}
	}
	temporalTime := time.Since(start)
	
	// Bad: Access different data each time
	start = time.Now()
	sum = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			sum += data[j] // Access different elements
		}
	}
	nonTemporalTime := time.Since(start)
	
	fmt.Printf("      Temporal locality: %v\n", temporalTime)
	fmt.Printf("      Non-temporal: %v\n", nonTemporalTime)
	fmt.Printf("      Temporal is %.2fx faster\n", float64(nonTemporalTime)/float64(temporalTime))
}

// 3. Branch Prediction Optimization
func branchPredictionOptimization() {
	fmt.Println("3. Branch Prediction Optimization:")
	fmt.Println("Understanding branch prediction and optimization...")

	// Demonstrate branch prediction impact
	branchPredictionImpact()
	
	// Show branchless programming techniques
	branchlessProgramming()
	
	// Demonstrate branch optimization
	branchOptimization()
}

func branchPredictionImpact() {
	fmt.Println("  📊 Branch prediction impact:")
	
	// Create sorted and unsorted data
	sortedData := make([]int, 100000)
	unsortedData := make([]int, 100000)
	
	for i := range sortedData {
		sortedData[i] = i
		unsortedData[i] = rand.Intn(100000)
	}
	
	// Test with sorted data (good branch prediction)
	start := time.Now()
	sum := 0
	for _, v := range sortedData {
		if v < 50000 {
			sum += v
		}
	}
	sortedTime := time.Since(start)
	
	// Test with unsorted data (poor branch prediction)
	start = time.Now()
	sum = 0
	for _, v := range unsortedData {
		if v < 50000 {
			sum += v
		}
	}
	unsortedTime := time.Since(start)
	
	fmt.Printf("    Sorted data (good prediction): %v\n", sortedTime)
	fmt.Printf("    Unsorted data (poor prediction): %v\n", unsortedTime)
	fmt.Printf("    Sorted is %.2fx faster\n", float64(unsortedTime)/float64(sortedTime))
}

func branchlessProgramming() {
	fmt.Println("  📊 Branchless programming techniques:")
	
	// Traditional branch
	a, b := 10, 20
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Printf("    Traditional max: %d\n", max)
	
	// Branchless max
	max = a - ((a - b) & ((a - b) >> 31))
	fmt.Printf("    Branchless max: %d\n", max)
	
	// Traditional conditional
	value := 15
	var result int
	if value > 10 {
		result = value * 2
	} else {
		result = value
	}
	fmt.Printf("    Traditional conditional: %d\n", result)
	
	// Branchless conditional
	mask := (value - 10) >> 31 // -1 if value <= 10, 0 if value > 10
	result = value + ((value * 2 - value) & ^mask)
	fmt.Printf("    Branchless conditional: %d\n", result)
}

func branchOptimization() {
	fmt.Println("  📊 Branch optimization:")
	
	// Demonstrate branch reordering
	branchReorderingExample()
	
	// Show branch elimination
	branchEliminationExample()
}

func branchReorderingExample() {
	fmt.Println("    Branch reordering:")
	
	// Common case first
	value := 5
	if value < 10 { // Common case
		fmt.Println("      Common case: value < 10")
	} else if value < 100 { // Less common
		fmt.Println("      Less common: value < 100")
	} else { // Rare case
		fmt.Println("      Rare case: value >= 100")
	}
}

func branchEliminationExample() {
	fmt.Println("    Branch elimination:")
	
	// Instead of multiple if-else
	values := []int{1, 2, 3, 4, 5}
	
	// Bad: Multiple branches
	for _, v := range values {
		if v == 1 {
			fmt.Println("      One")
		} else if v == 2 {
			fmt.Println("      Two")
		} else if v == 3 {
			fmt.Println("      Three")
		} else {
			fmt.Println("      Other")
		}
	}
	
	// Good: Lookup table (branch elimination)
	lookup := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	
	for _, v := range values {
		if result, ok := lookup[v]; ok {
			fmt.Printf("      %s\n", result)
		} else {
			fmt.Println("      Other")
		}
	}
}

// 4. Compiler Optimizations
func compilerOptimizations() {
	fmt.Println("4. Compiler Optimizations:")
	fmt.Println("Understanding Go compiler optimization flags...")

	// Demonstrate inlining
	inliningExample()
	
	// Show dead code elimination
	deadCodeEliminationExample()
	
	// Demonstrate loop optimization
	loopOptimizationExample()
}

func inliningExample() {
	fmt.Println("  📊 Inlining example:")
	
	// Small function that can be inlined
	result := add(10, 20)
	fmt.Printf("    Inlined function result: %d\n", result)
	
	// Large function that might not be inlined
	largeResult := largeFunction(100)
	fmt.Printf("    Large function result: %d\n", largeResult)
}

func add(a, b int) int {
	return a + b
}

func largeFunction(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += i * j
		}
	}
	return sum
}

func deadCodeEliminationExample() {
	fmt.Println("  📊 Dead code elimination:")
	
	// This code will be eliminated by the compiler
	if false {
		fmt.Println("    This will never execute")
	}
	
	// This code will be kept
	if true {
		fmt.Println("    This will execute")
	}
	
	// Unused variable (might be eliminated)
	unused := 42
	_ = unused
}

func loopOptimizationExample() {
	fmt.Println("  📊 Loop optimization:")
	
	// Loop unrolling example
	loopUnrollingExample()
	
	// Loop fusion example
	loopFusionExample()
}

func loopUnrollingExample() {
	fmt.Println("    Loop unrolling:")
	
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	// Manual loop unrolling
	start := time.Now()
	sum := 0
	for i := 0; i < len(data); i += 4 {
		sum += data[i]
		if i+1 < len(data) {
			sum += data[i+1]
		}
		if i+2 < len(data) {
			sum += data[i+2]
		}
		if i+3 < len(data) {
			sum += data[i+3]
		}
	}
	unrolledTime := time.Since(start)
	
	// Regular loop
	start = time.Now()
	sum = 0
	for _, v := range data {
		sum += v
	}
	regularTime := time.Since(start)
	
	fmt.Printf("      Unrolled loop: %v\n", unrolledTime)
	fmt.Printf("      Regular loop: %v\n", regularTime)
}

func loopFusionExample() {
	fmt.Println("    Loop fusion:")
	
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	// Separate loops
	start := time.Now()
	sum := 0
	for _, v := range data {
		sum += v
	}
	product := 1
	for _, v := range data {
		product *= v
	}
	separateTime := time.Since(start)
	
	// Fused loop
	start = time.Now()
	sum = 0
	product = 1
	for _, v := range data {
		sum += v
		product *= v
	}
	fusedTime := time.Since(start)
	
	fmt.Printf("      Separate loops: %v\n", separateTime)
	fmt.Printf("      Fused loop: %v\n", fusedTime)
}

// 5. Assembly and Low-Level Optimization
func assemblyAndLowLevelOptimization() {
	fmt.Println("5. Assembly and Low-Level Optimization:")
	fmt.Println("Understanding Go assembly and low-level techniques...")

	// Demonstrate unsafe operations
	unsafeOperations()
	
	// Show bit manipulation
	bitManipulation()
	
	// Demonstrate memory operations
	memoryOperations()
}

func unsafeOperations() {
	fmt.Println("  📊 Unsafe operations:")
	
	// Convert between types using unsafe
	var i int32 = 42
	ptr := unsafe.Pointer(&i)
	floatPtr := (*float32)(ptr)
	fmt.Printf("    Int32 to Float32: %f\n", *floatPtr)
	
	// Get size of types
	fmt.Printf("    Size of int32: %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("    Size of float32: %d bytes\n", unsafe.Sizeof(*floatPtr))
}

func bitManipulation() {
	fmt.Println("  📊 Bit manipulation:")
	
	// Fast power of 2 check
	value := 16
	isPowerOfTwo := value&(value-1) == 0
	fmt.Printf("    %d is power of 2: %t\n", value, isPowerOfTwo)
	
	// Fast absolute value
	negative := -42
	abs := (negative ^ (negative >> 31)) - (negative >> 31)
	fmt.Printf("    Absolute value of %d: %d\n", negative, abs)
	
	// Count set bits
	count := 0
	for n := 255; n > 0; n &= n - 1 {
		count++
	}
	fmt.Printf("    Number of set bits in 255: %d\n", count)
}

func memoryOperations() {
	fmt.Println("  📊 Memory operations:")
	
	// Copy memory efficiently
	src := []byte("Hello, World!")
	dst := make([]byte, len(src))
	copy(dst, src)
	fmt.Printf("    Copied: %s\n", string(dst))
	
	// Zero memory
	zeroSlice := make([]byte, 100)
	for i := range zeroSlice {
		zeroSlice[i] = 0
	}
	fmt.Printf("    Zeroed %d bytes\n", len(zeroSlice))
}

// 6. Vector Operations
func vectorOperations() {
	fmt.Println("6. Vector Operations:")
	fmt.Println("Advanced vector operations and optimizations...")

	// Vector math operations
	vectorMathOperations()
	
	// Vector reduction operations
	vectorReductionOperations()
	
	// Vector comparison operations
	vectorComparisonOperations()
}

func vectorMathOperations() {
	fmt.Println("  📊 Vector math operations:")
	
	// Vector addition
	a := []float64{1, 2, 3, 4, 5}
	b := []float64{5, 4, 3, 2, 1}
	c := vectorAdd(a, b)
	fmt.Printf("    Vector addition: %v + %v = %v\n", a, b, c)
	
	// Vector multiplication
	d := vectorMultiply(a, b)
	fmt.Printf("    Vector multiplication: %v * %v = %v\n", a, b, d)
	
	// Vector scaling
	scaled := vectorScale(a, 2.0)
	fmt.Printf("    Vector scaling: %v * 2.0 = %v\n", a, scaled)
}

func vectorMultiply(a, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] * b[i]
	}
	return result
}

func vectorScale(a []float64, scale float64) []float64 {
	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] * scale
	}
	return result
}

func vectorReductionOperations() {
	fmt.Println("  📊 Vector reduction operations:")
	
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Sum
	sum := vectorSum(data)
	fmt.Printf("    Sum: %.2f\n", sum)
	
	// Product
	product := vectorProduct(data)
	fmt.Printf("    Product: %.2f\n", product)
	
	// Maximum
	max := vectorMax(data)
	fmt.Printf("    Maximum: %.2f\n", max)
	
	// Minimum
	min := vectorMin(data)
	fmt.Printf("    Minimum: %.2f\n", min)
}

func vectorSum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func vectorProduct(data []float64) float64 {
	product := 1.0
	for _, v := range data {
		product *= v
	}
	return product
}

func vectorMax(data []float64) float64 {
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func vectorMin(data []float64) float64 {
	min := data[0]
	for _, v := range data[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func vectorComparisonOperations() {
	fmt.Println("  📊 Vector comparison operations:")
	
	a := []float64{1, 2, 3, 4, 5}
	b := []float64{1, 3, 2, 4, 6}
	
	// Element-wise comparison
	equal := vectorEqual(a, b)
	fmt.Printf("    Equal: %v\n", equal)
	
	// Greater than
	greater := vectorGreater(a, b)
	fmt.Printf("    Greater: %v\n", greater)
	
	// Less than
	less := vectorLess(a, b)
	fmt.Printf("    Less: %v\n", less)
}

func vectorEqual(a, b []float64) []bool {
	result := make([]bool, len(a))
	for i := range a {
		result[i] = a[i] == b[i]
	}
	return result
}

func vectorGreater(a, b []float64) []bool {
	result := make([]bool, len(a))
	for i := range a {
		result[i] = a[i] > b[i]
	}
	return result
}

func vectorLess(a, b []float64) []bool {
	result := make([]bool, len(a))
	for i := range a {
		result[i] = a[i] < b[i]
	}
	return result
}

// 7. Memory Access Patterns
func memoryAccessPatterns() {
	fmt.Println("7. Memory Access Patterns:")
	fmt.Println("Optimizing memory access patterns...")

	// Sequential vs random access
	sequentialVsRandomAccess()
	
	// Strided access patterns
	stridedAccessPatterns()
	
	// Memory prefetching
	memoryPrefetching()
}

func sequentialVsRandomAccess() {
	fmt.Println("  📊 Sequential vs random access:")
	
	const size = 1000000
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	
	// Sequential access
	start := time.Now()
	sum := 0
	for i := 0; i < size; i++ {
		sum += data[i]
	}
	sequentialTime := time.Since(start)
	
	// Random access
	indices := make([]int, size)
	for i := range indices {
		indices[i] = rand.Intn(size)
	}
	
	start = time.Now()
	sum = 0
	for _, idx := range indices {
		sum += data[idx]
	}
	randomTime := time.Since(start)
	
	fmt.Printf("    Sequential access: %v\n", sequentialTime)
	fmt.Printf("    Random access: %v\n", randomTime)
	fmt.Printf("    Sequential is %.2fx faster\n", float64(randomTime)/float64(sequentialTime))
}

func stridedAccessPatterns() {
	fmt.Println("  📊 Strided access patterns:")
	
	const size = 1000000
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	
	// Stride 1 (sequential)
	start := time.Now()
	sum := 0
	for i := 0; i < size; i += 1 {
		sum += data[i]
	}
	stride1Time := time.Since(start)
	
	// Stride 2
	start = time.Now()
	sum = 0
	for i := 0; i < size; i += 2 {
		sum += data[i]
	}
	stride2Time := time.Since(start)
	
	// Stride 4
	start = time.Now()
	sum = 0
	for i := 0; i < size; i += 4 {
		sum += data[i]
	}
	stride4Time := time.Since(start)
	
	fmt.Printf("    Stride 1: %v\n", stride1Time)
	fmt.Printf("    Stride 2: %v\n", stride2Time)
	fmt.Printf("    Stride 4: %v\n", stride4Time)
}

func memoryPrefetching() {
	fmt.Println("  📊 Memory prefetching:")
	fmt.Println("    Note: Go doesn't have explicit prefetch instructions")
	fmt.Println("    CPU automatically prefetches based on access patterns")
	fmt.Println("    Sequential access patterns benefit most from prefetching")
}

// 8. Performance Profiling
func performanceProfiling() {
	fmt.Println("8. Performance Profiling:")
	fmt.Println("Advanced performance profiling techniques...")

	// CPU profiling
	cpuProfiling()
	
	// Memory profiling
	memoryProfiling()
	
	// Goroutine profiling
	goroutineProfiling()
}

func cpuProfiling() {
	fmt.Println("  📊 CPU profiling:")
	fmt.Println("    Use: go tool pprof cpu.prof")
	fmt.Println("    Focus on hot functions and call graphs")
	fmt.Println("    Look for optimization opportunities")
}

func memoryProfiling() {
	fmt.Println("  📊 Memory profiling:")
	fmt.Println("    Use: go tool pprof mem.prof")
	fmt.Println("    Identify memory allocations and leaks")
	fmt.Println("    Optimize allocation patterns")
}

func goroutineProfiling() {
	fmt.Println("  📊 Goroutine profiling:")
	fmt.Println("    Use: go tool pprof goroutine.prof")
	fmt.Println("    Analyze goroutine usage and blocking")
	fmt.Println("    Optimize concurrency patterns")
}

// 9. Micro-optimizations
func microOptimizations() {
	fmt.Println("9. Micro-optimizations:")
	fmt.Println("Small but impactful optimizations...")

	// Function call optimization
	functionCallOptimization()
	
	// Variable optimization
	variableOptimization()
	
	// Expression optimization
	expressionOptimization()
}

func functionCallOptimization() {
	fmt.Println("  📊 Function call optimization:")
	
	// Inline small functions
	result := inlineAdd(10, 20)
	fmt.Printf("    Inline function: %d\n", result)
	
	// Avoid function calls in loops
	avoidFunctionCallsInLoops()
}

func inlineAdd(a, b int) int {
	return a + b
}

func avoidFunctionCallsInLoops() {
	fmt.Println("    Avoiding function calls in loops:")
	
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	// Bad: Function call in loop
	start := time.Now()
	sum := 0
	for _, v := range data {
		sum += expensiveFunction(v)
	}
	badTime := time.Since(start)
	
	// Good: Inline computation
	start = time.Now()
	sum = 0
	for _, v := range data {
		sum += v * v // Inline the computation
	}
	goodTime := time.Since(start)
	
	fmt.Printf("      Function calls: %v\n", badTime)
	fmt.Printf("      Inline computation: %v\n", goodTime)
	fmt.Printf("      Inline is %.2fx faster\n", float64(badTime)/float64(goodTime))
}

func expensiveFunction(x int) int {
	return x * x
}

func variableOptimization() {
	fmt.Println("  📊 Variable optimization:")
	
	// Use appropriate variable types
	var small int8 = 127
	var medium int32 = 2147483647
	var large int64 = 9223372036854775807
	
	fmt.Printf("    Small: %d (1 byte)\n", small)
	fmt.Printf("    Medium: %d (4 bytes)\n", medium)
	fmt.Printf("    Large: %d (8 bytes)\n", large)
	
	// Avoid unnecessary variable declarations
	unnecessaryVariableExample()
}

func unnecessaryVariableExample() {
	fmt.Println("    Avoiding unnecessary variables:")
	
	// Bad: Unnecessary variable
	result := 10 + 20
	fmt.Printf("      Result: %d\n", result)
	
	// Good: Direct computation
	fmt.Printf("      Direct: %d\n", 10+20)
}

func expressionOptimization() {
	fmt.Println("  📊 Expression optimization:")
	
	// Use bit operations for powers of 2
	multiplyByTwo := 42 << 1
	divideByTwo := 42 >> 1
	fmt.Printf("    42 * 2 = %d (bit shift)\n", multiplyByTwo)
	fmt.Printf("    42 / 2 = %d (bit shift)\n", divideByTwo)
	
	// Use multiplication instead of division when possible
	multiplyByHalf := 42 * 0.5
	fmt.Printf("    42 * 0.5 = %.2f\n", multiplyByHalf)
}

// 10. Best Practices
func advancedOptimizationBestPractices() {
	fmt.Println("10. Advanced Optimization Best Practices:")
	fmt.Println("Best practices for advanced optimization...")

	fmt.Println("  📝 Best Practice 1: Profile before optimizing")
	fmt.Println("    - Use pprof to identify bottlenecks")
	fmt.Println("    - Focus on hot paths and critical sections")
	
	fmt.Println("  📝 Best Practice 2: Understand your data")
	fmt.Println("    - Choose appropriate data structures")
	fmt.Println("    - Consider cache-friendly layouts")
	
	fmt.Println("  📝 Best Practice 3: Optimize for your use case")
	fmt.Println("    - Different optimizations for different scenarios")
	fmt.Println("    - Balance readability vs performance")
	
	fmt.Println("  📝 Best Practice 4: Use compiler optimizations")
	fmt.Println("    - Enable appropriate optimization flags")
	fmt.Println("    - Let the compiler do its job")
	
	fmt.Println("  📝 Best Practice 5: Measure and validate")
	fmt.Println("    - Always measure performance improvements")
	fmt.Println("    - Validate that optimizations work as expected")
	
	fmt.Println("  📝 Best Practice 6: Consider maintainability")
	fmt.Println("    - Don't sacrifice readability for minor gains")
	fmt.Println("    - Document complex optimizations")
	
	fmt.Println("  📝 Best Practice 7: Test on target hardware")
	fmt.Println("    - Performance varies across different systems")
	fmt.Println("    - Test on production-like environments")
}
