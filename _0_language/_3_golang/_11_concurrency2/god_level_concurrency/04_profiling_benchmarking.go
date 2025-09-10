package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// GOD-LEVEL CONCEPT 4: Profiling & Benchmarking
// Mastering performance measurement and optimization

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: Profiling & Benchmarking ===")
	
	// Start pprof server
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println("Access profiles at:")
		log.Println("  CPU: http://localhost:6060/debug/pprof/profile")
		log.Println("  Memory: http://localhost:6060/debug/pprof/heap")
		log.Println("  Goroutines: http://localhost:6060/debug/pprof/goroutine")
		log.Println("  Block: http://localhost:6060/debug/pprof/block")
		log.Println("  Mutex: http://localhost:6060/debug/pprof/mutex")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	// 1. CPU Profiling
	demonstrateCPUProfiling()
	
	// 2. Memory Profiling
	demonstrateMemoryProfiling()
	
	// 3. Goroutine Profiling
	demonstrateGoroutineProfiling()
	
	// 4. Block Profiling
	demonstrateBlockProfiling()
	
	// 5. Mutex Profiling
	demonstrateMutexProfiling()
	
	// 6. Benchmarking Techniques
	demonstrateBenchmarking()
	
	// 7. Performance Optimization Examples
	demonstratePerformanceOptimization()
	
	fmt.Println("\n=== PROFILING COMPLETE ===")
	fmt.Println("Check the pprof server at http://localhost:6060 for detailed profiles!")
	fmt.Println("Press Ctrl+C to exit...")
	
	// Keep the server running
	select {}
}

// Demonstrate CPU Profiling
func demonstrateCPUProfiling() {
	fmt.Println("\n=== 1. CPU PROFILING ===")
	
	fmt.Println(`
🔥 CPU Profiling:
• Identifies CPU bottlenecks
• Shows function call frequency
• Helps optimize hot paths
• Use: go tool pprof http://localhost:6060/debug/pprof/profile
`)

	// Create CPU-intensive workload
	cpuIntensiveWorkload()
}

func cpuIntensiveWorkload() {
	fmt.Println("\n--- CPU-Intensive Workload ---")
	
	const numGoroutines = 10
	const iterations = 1000000
	
	var wg sync.WaitGroup
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// CPU-intensive work
			for j := 0; j < iterations; j++ {
				// Simulate complex calculations
				result := 0
				for k := 0; k < 100; k++ {
					result += k * k * k
				}
				_ = result
			}
			
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("CPU workload completed in %v\n", duration)
	fmt.Println("💡 Check CPU profile at http://localhost:6060/debug/pprof/profile")
}

// Demonstrate Memory Profiling
func demonstrateMemoryProfiling() {
	fmt.Println("\n=== 2. MEMORY PROFILING ===")
	
	fmt.Println(`
🧠 Memory Profiling:
• Identifies memory allocations
• Shows memory usage patterns
• Helps find memory leaks
• Use: go tool pprof http://localhost:6060/debug/pprof/heap
`)

	// Create memory-intensive workload
	memoryIntensiveWorkload()
}

func memoryIntensiveWorkload() {
	fmt.Println("\n--- Memory-Intensive Workload ---")
	
	const numGoroutines = 5
	const iterations = 100000
	
	var wg sync.WaitGroup
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Memory-intensive work
			for j := 0; j < iterations; j++ {
				// Allocate large slices
				data := make([]int, 1000)
				for k := range data {
					data[k] = rand.Intn(1000)
				}
				
				// Process data
				sum := 0
				for _, v := range data {
					sum += v
				}
				_ = sum
			}
			
			fmt.Printf("Memory goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("Memory workload completed in %v\n", duration)
	fmt.Println("💡 Check memory profile at http://localhost:6060/debug/pprof/heap")
}

// Demonstrate Goroutine Profiling
func demonstrateGoroutineProfiling() {
	fmt.Println("\n=== 3. GOROUTINE PROFILING ===")
	
	fmt.Println(`
🔄 Goroutine Profiling:
• Shows goroutine states
• Identifies goroutine leaks
• Helps debug deadlocks
• Use: go tool pprof http://localhost:6060/debug/pprof/goroutine
`)

	// Create goroutine workload
	goroutineWorkload()
}

func goroutineWorkload() {
	fmt.Println("\n--- Goroutine Workload ---")
	
	const numGoroutines = 100
	
	// Create many goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// Simulate work with different durations
			duration := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(duration)
			fmt.Printf("Goroutine %d completed after %v\n", id, duration)
		}(i)
	}
	
	// Show current goroutine count
	fmt.Printf("Current goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("💡 Check goroutine profile at http://localhost:6060/debug/pprof/goroutine")
}

// Demonstrate Block Profiling
func demonstrateBlockProfiling() {
	fmt.Println("\n=== 4. BLOCK PROFILING ===")
	
	fmt.Println(`
🚧 Block Profiling:
• Shows blocking operations
• Identifies contention points
• Helps optimize synchronization
• Use: go tool pprof http://localhost:6060/debug/pprof/block
`)

	// Create blocking workload
	blockingWorkload()
}

func blockingWorkload() {
	fmt.Println("\n--- Blocking Workload ---")
	
	const numGoroutines = 10
	const iterations = 1000
	
	var mu sync.Mutex
	var wg sync.WaitGroup
	
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < iterations; j++ {
				// Simulate blocking operations
				mu.Lock()
				time.Sleep(1 * time.Millisecond) // Simulate work
				mu.Unlock()
			}
			
			fmt.Printf("Blocking goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("Blocking workload completed in %v\n", duration)
	fmt.Println("💡 Check block profile at http://localhost:6060/debug/pprof/block")
}

// Demonstrate Mutex Profiling
func demonstrateMutexProfiling() {
	fmt.Println("\n=== 5. MUTEX PROFILING ===")
	
	fmt.Println(`
🔒 Mutex Profiling:
• Shows mutex contention
• Identifies lock hotspots
• Helps optimize locking
• Use: go tool pprof http://localhost:6060/debug/pprof/mutex
`)

	// Create mutex contention workload
	mutexContentionWorkload()
}

func mutexContentionWorkload() {
	fmt.Println("\n--- Mutex Contention Workload ---")
	
	const numGoroutines = 20
	const iterations = 1000
	
	var mu sync.Mutex
	var counter int64
	var wg sync.WaitGroup
	
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < iterations; j++ {
				// High contention on mutex
				mu.Lock()
				counter++
				mu.Unlock()
			}
			
			fmt.Printf("Mutex goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("Mutex workload completed in %v (counter: %d)\n", duration, counter)
	fmt.Println("💡 Check mutex profile at http://localhost:6060/debug/pprof/mutex")
}

// Demonstrate Benchmarking
func demonstrateBenchmarking() {
	fmt.Println("\n=== 6. BENCHMARKING TECHNIQUES ===")
	
	fmt.Println(`
📊 Benchmarking:
• Measure performance accurately
• Compare different implementations
• Identify performance regressions
• Use: go test -bench=.
`)

	// Run various benchmarks
	runBenchmarks()
}

func runBenchmarks() {
	fmt.Println("\n--- Running Benchmarks ---")
	
	// Benchmark 1: Atomic vs Mutex
	benchmarkAtomicVsMutex()
	
	// Benchmark 2: Channel vs Mutex
	benchmarkChannelVsMutex()
	
	// Benchmark 3: Different data structures
	benchmarkDataStructures()
}

func benchmarkAtomicVsMutex() {
	fmt.Println("\n--- Atomic vs Mutex Benchmark ---")
	
	const iterations = 1000000
	const numGoroutines = 10
	
	// Atomic counter
	var atomicCounter int64
	start := time.Now()
	
	var wg sync.WaitGroup
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

func benchmarkChannelVsMutex() {
	fmt.Println("\n--- Channel vs Mutex Benchmark ---")
	
	const iterations = 100000
	const numGoroutines = 10
	
	// Channel-based communication
	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	
	start := time.Now()
	
	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < iterations; i++ {
			ch <- i
		}
	}()
	
	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// Process value
		}
	}()
	
	wg.Wait()
	channelDuration := time.Since(start)
	
	// Mutex-based communication
	var mu sync.Mutex
	var data []int
	
	start = time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/numGoroutines; j++ {
				mu.Lock()
				data = append(data, j)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexDuration := time.Since(start)
	
	fmt.Printf("Channel: %v\n", channelDuration)
	fmt.Printf("Mutex:   %v\n", mutexDuration)
	fmt.Printf("Channel is %.2fx faster\n", float64(mutexDuration)/float64(channelDuration))
}

func benchmarkDataStructures() {
	fmt.Println("\n--- Data Structures Benchmark ---")
	
	const iterations = 100000
	
	// Slice vs Map
	benchmarkSliceVsMap(iterations)
	
	// Different slice operations
	benchmarkSliceOperations(iterations)
}

func benchmarkSliceVsMap(iterations int) {
	fmt.Println("\n--- Slice vs Map ---")
	
	// Slice
	start := time.Now()
	slice := make([]int, 0, iterations)
	for i := 0; i < iterations; i++ {
		slice = append(slice, i)
	}
	sliceDuration := time.Since(start)
	
	// Map
	start = time.Now()
	m := make(map[int]int)
	for i := 0; i < iterations; i++ {
		m[i] = i
	}
	mapDuration := time.Since(start)
	
	fmt.Printf("Slice: %v\n", sliceDuration)
	fmt.Printf("Map:   %v\n", mapDuration)
	fmt.Printf("Slice is %.2fx faster\n", float64(mapDuration)/float64(sliceDuration))
}

func benchmarkSliceOperations(iterations int) {
	fmt.Println("\n--- Slice Operations ---")
	
	// Pre-allocated slice
	start := time.Now()
	slice1 := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		slice1[i] = i
	}
	preAllocDuration := time.Since(start)
	
	// Dynamic slice
	start = time.Now()
	slice2 := make([]int, 0)
	for i := 0; i < iterations; i++ {
		slice2 = append(slice2, i)
	}
	dynamicDuration := time.Since(start)
	
	fmt.Printf("Pre-allocated: %v\n", preAllocDuration)
	fmt.Printf("Dynamic:       %v\n", dynamicDuration)
	fmt.Printf("Pre-allocated is %.2fx faster\n", float64(dynamicDuration)/float64(preAllocDuration))
}

// Demonstrate Performance Optimization
func demonstratePerformanceOptimization() {
	fmt.Println("\n=== 7. PERFORMANCE OPTIMIZATION EXAMPLES ===")
	
	fmt.Println(`
⚡ Performance Optimization:
• Identify bottlenecks
• Apply optimization techniques
• Measure improvements
• Iterate and refine
`)

	// Optimization examples
	optimizationExamples()
}

func optimizationExamples() {
	fmt.Println("\n--- Optimization Examples ---")
	
	// Example 1: String concatenation
	optimizeStringConcatenation()
	
	// Example 2: Memory allocation
	optimizeMemoryAllocation()
	
	// Example 3: Goroutine usage
	optimizeGoroutineUsage()
}

func optimizeStringConcatenation() {
	fmt.Println("\n--- String Concatenation Optimization ---")
	
	const iterations = 10000
	
	// Inefficient: String concatenation
	start := time.Now()
	result1 := ""
	for i := 0; i < iterations; i++ {
		result1 += fmt.Sprintf("item-%d ", i)
	}
	inefficientDuration := time.Since(start)
	
	// Efficient: strings.Builder
	start = time.Now()
	var builder strings.Builder
	builder.Grow(iterations * 10) // Pre-allocate capacity
	for i := 0; i < iterations; i++ {
		builder.WriteString(fmt.Sprintf("item-%d ", i))
	}
	_ = builder.String() // Use the result
	efficientDuration := time.Since(start)
	
	fmt.Printf("String concatenation: %v\n", inefficientDuration)
	fmt.Printf("strings.Builder:      %v\n", efficientDuration)
	fmt.Printf("strings.Builder is %.2fx faster\n", float64(inefficientDuration)/float64(efficientDuration))
}

func optimizeMemoryAllocation() {
	fmt.Println("\n--- Memory Allocation Optimization ---")
	
	const iterations = 100000
	
	// Inefficient: Allocate in loop
	start := time.Now()
	var data []int
	for i := 0; i < iterations; i++ {
		data = append(data, i)
	}
	inefficientDuration := time.Since(start)
	
	// Efficient: Pre-allocate
	start = time.Now()
	data2 := make([]int, 0, iterations) // Pre-allocate capacity
	for i := 0; i < iterations; i++ {
		data2 = append(data2, i)
	}
	efficientDuration := time.Since(start)
	
	fmt.Printf("Dynamic allocation: %v\n", inefficientDuration)
	fmt.Printf("Pre-allocated:      %v\n", efficientDuration)
	fmt.Printf("Pre-allocated is %.2fx faster\n", float64(inefficientDuration)/float64(efficientDuration))
}

func optimizeGoroutineUsage() {
	fmt.Println("\n--- Goroutine Usage Optimization ---")
	
	const iterations = 100000
	
	// Inefficient: Too many goroutines
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = id * id // Simple work
		}(i)
	}
	wg.Wait()
	inefficientDuration := time.Since(start)
	
	// Efficient: Worker pool
	start = time.Now()
	numWorkers := runtime.NumCPU()
	jobs := make(chan int, iterations)
	results := make(chan int, iterations)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobs {
				results <- job * job
			}
		}()
	}
	
	// Send jobs
	go func() {
		for i := 0; i < iterations; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	
	// Collect results
	for i := 0; i < iterations; i++ {
		<-results
	}
	efficientDuration := time.Since(start)
	
	fmt.Printf("Many goroutines: %v\n", inefficientDuration)
	fmt.Printf("Worker pool:      %v\n", efficientDuration)
	fmt.Printf("Worker pool is %.2fx faster\n", float64(inefficientDuration)/float64(efficientDuration))
}

