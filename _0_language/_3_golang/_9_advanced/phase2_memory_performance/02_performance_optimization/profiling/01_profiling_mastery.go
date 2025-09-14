package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sort"
	"strings"
	"sync"
	"time"
)

// ⚡ PROFILING MASTERY
// Understanding Go profiling and performance measurement

func main() {
	fmt.Println("⚡ PROFILING MASTERY")
	fmt.Println("===================")

	// 1. CPU Profiling
	fmt.Println("\n1. CPU Profiling:")
	cpuProfiling()

	// 2. Memory Profiling
	fmt.Println("\n2. Memory Profiling:")
	memoryProfiling()

	// 3. Goroutine Profiling
	fmt.Println("\n3. Goroutine Profiling:")
	goroutineProfiling()

	// 4. Block Profiling
	fmt.Println("\n4. Block Profiling:")
	blockProfiling()

	// 5. Mutex Profiling
	fmt.Println("\n5. Mutex Profiling:")
	mutexProfiling()

	// 6. Trace Analysis
	fmt.Println("\n6. Trace Analysis:")
	traceAnalysis()

	// 7. Performance Measurement
	fmt.Println("\n7. Performance Measurement:")
	performanceMeasurement()
}

// CPU PROFILING: Understanding CPU usage patterns
func cpuProfiling() {
	fmt.Println("Understanding CPU profiling...")
	
	// Start CPU profiling
	fmt.Println("  📊 Starting CPU profiling...")
	
	// Simulate CPU-intensive work
	cpuIntensiveWork()
	
	// Stop CPU profiling
	fmt.Println("  📊 CPU profiling completed")
	fmt.Println("  💡 Use: go tool pprof cpu.prof")
}

func cpuIntensiveWork() {
	// Simulate CPU-intensive work
	data := make([]int, 1000000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	
	// Sort the data (CPU intensive)
	sort.Ints(data)
	
	// Search for specific values
	for i := 0; i < 1000; i++ {
		searchValue := rand.Intn(1000)
		_ = sort.SearchInts(data, searchValue)
	}
}

// MEMORY PROFILING: Understanding memory usage patterns
func memoryProfiling() {
	fmt.Println("Understanding memory profiling...")
	
	// Start memory profiling
	fmt.Println("  📊 Starting memory profiling...")
	
	// Simulate memory-intensive work
	memoryIntensiveWork()
	
	// Stop memory profiling
	fmt.Println("  📊 Memory profiling completed")
	fmt.Println("  💡 Use: go tool pprof mem.prof")
}

func memoryIntensiveWork() {
	// Create large data structures
	largeSlice := make([][]byte, 1000)
	for i := range largeSlice {
		largeSlice[i] = make([]byte, 1024) // 1KB each
	}
	
	// Simulate string operations
	var builder strings.Builder
	for i := 0; i < 10000; i++ {
		builder.WriteString(fmt.Sprintf("item-%d ", i))
	}
	_ = builder.String()
	
	// Create maps
	largeMap := make(map[string]int, 10000)
	for i := 0; i < 10000; i++ {
		largeMap[fmt.Sprintf("key-%d", i)] = i
	}
}

// GOROUTINE PROFILING: Understanding goroutine usage
func goroutineProfiling() {
	fmt.Println("Understanding goroutine profiling...")
	
	// Start goroutine profiling
	fmt.Println("  📊 Starting goroutine profiling...")
	
	// Create multiple goroutines
	goroutineIntensiveWork()
	
	// Stop goroutine profiling
	fmt.Println("  📊 Goroutine profiling completed")
	fmt.Println("  💡 Use: go tool pprof goroutine.prof")
}

func goroutineIntensiveWork() {
	// Create multiple goroutines
	for i := 0; i < 100; i++ {
		go func(id int) {
			// Simulate work
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			
			// Create more goroutines
			for j := 0; j < 10; j++ {
				go func(subID int) {
					time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
				}(j)
			}
		}(i)
	}
	
	// Wait for goroutines to complete
	time.Sleep(2 * time.Second)
}

// BLOCK PROFILING: Understanding blocking operations
func blockProfiling() {
	fmt.Println("Understanding block profiling...")
	
	// Start block profiling
	fmt.Println("  📊 Starting block profiling...")
	
	// Simulate blocking operations
	blockingWork()
	
	// Stop block profiling
	fmt.Println("  📊 Block profiling completed")
	fmt.Println("  💡 Use: go tool pprof block.prof")
}

func blockingWork() {
	// Simulate blocking operations
	ch := make(chan int, 10)
	
	// Producer goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	// Consumer goroutine
	go func() {
		for range ch {
			// Simulate processing
			time.Sleep(1 * time.Millisecond)
		}
	}()
	
	// Wait for completion
	time.Sleep(1 * time.Second)
}

// MUTEX PROFILING: Understanding mutex contention
func mutexProfiling() {
	fmt.Println("Understanding mutex profiling...")
	
	// Start mutex profiling
	fmt.Println("  📊 Starting mutex profiling...")
	
	// Simulate mutex contention
	mutexContentionWork()
	
	// Stop mutex profiling
	fmt.Println("  📊 Mutex profiling completed")
	fmt.Println("  💡 Use: go tool pprof mutex.prof")
}

func mutexContentionWork() {
	var mu sync.Mutex
	var counter int
	
	// Create multiple goroutines that contend for the mutex
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	
	// Wait for completion
	time.Sleep(2 * time.Second)
}

// TRACE ANALYSIS: Understanding execution traces
func traceAnalysis() {
	fmt.Println("Understanding trace analysis...")
	
	// Start tracing
	fmt.Println("  📊 Starting trace analysis...")
	
	// Simulate traced work
	tracedWork()
	
	// Stop tracing
	fmt.Println("  📊 Trace analysis completed")
	fmt.Println("  💡 Use: go tool trace trace.out")
}

func tracedWork() {
	// Simulate work that will be traced
	for i := 0; i < 100; i++ {
		go func(id int) {
			// Simulate work
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			
			// Simulate channel operations
			ch := make(chan int, 1)
			ch <- id
			<-ch
		}(i)
	}
	
	// Wait for completion
	time.Sleep(1 * time.Second)
}

// PERFORMANCE MEASUREMENT: Understanding performance metrics
func performanceMeasurement() {
	fmt.Println("Understanding performance measurement...")
	
	// 1. Basic timing
	fmt.Println("  📊 Basic timing:")
	basicTiming()
	
	// 2. Memory measurement
	fmt.Println("  📊 Memory measurement:")
	memoryMeasurement()
	
	// 3. CPU measurement
	fmt.Println("  📊 CPU measurement:")
	cpuMeasurement()
	
	// 4. Benchmarking
	fmt.Println("  📊 Benchmarking:")
	benchmarking()
}

func basicTiming() {
	// Measure function execution time
	start := time.Now()
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	duration := time.Since(start)
	fmt.Printf("    Execution time: %v\n", duration)
}

func memoryMeasurement() {
	// Measure memory usage
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	// Allocate memory
	data := make([]byte, 1024*1024) // 1MB
	_ = data
	
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("    Memory allocated: %d bytes\n", m2.Alloc-m1.Alloc)
	fmt.Printf("    Total allocations: %d\n", m2.Mallocs-m1.Mallocs)
}

func cpuMeasurement() {
	// Measure CPU usage
	start := time.Now()
	
	// Simulate CPU work
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	
	duration := time.Since(start)
	fmt.Printf("    CPU work completed in: %v\n", duration)
	fmt.Printf("    Result: %d\n", sum)
}

func benchmarking() {
	// Simple benchmarking
	iterations := 1000
	
	// Benchmark function 1
	start := time.Now()
	for i := 0; i < iterations; i++ {
		function1()
	}
	duration1 := time.Since(start)
	
	// Benchmark function 2
	start = time.Now()
	for i := 0; i < iterations; i++ {
		function2()
	}
	duration2 := time.Since(start)
	
	fmt.Printf("    Function 1: %v per iteration\n", duration1/time.Duration(iterations))
	fmt.Printf("    Function 2: %v per iteration\n", duration2/time.Duration(iterations))
	fmt.Printf("    Speedup: %.2fx\n", float64(duration1)/float64(duration2))
}

func function1() {
	// Simple function
	_ = make([]int, 100)
}

func function2() {
	// Optimized function
	_ = make([]int, 0, 100)
}

// PROFILING UTILITIES: Helper functions for profiling
func startCPUProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}

func startMemoryProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	runtime.GC()
	pprof.WriteHeapProfile(f)
	f.Close()
}

func startGoroutineProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	pprof.Lookup("goroutine").WriteTo(f, 0)
	f.Close()
}

func startBlockProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	pprof.Lookup("block").WriteTo(f, 0)
	f.Close()
}

func startMutexProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	pprof.Lookup("mutex").WriteTo(f, 0)
	f.Close()
}

func startTrace(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	trace.Start(f)
}

func stopTrace() {
	trace.Stop()
}
