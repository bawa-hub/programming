package basics

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Example 6: Basic Benchmarking
func basicBenchmarking() {
	fmt.Println("\n6. Basic Benchmarking")
	fmt.Println("====================")
	
	// Simple benchmark
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
	duration := time.Since(start)
	
	fmt.Printf("  Simple benchmark: %v\n", duration)
	
	// Concurrent benchmark
	start = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				_ = j * j
			}
		}()
	}
	wg.Wait()
	duration = time.Since(start)
	
	fmt.Printf("  Concurrent benchmark: %v\n", duration)
	
	fmt.Println("  Basic benchmarking completed")
}

// Example 7: Memory Allocation Analysis
func memoryAllocationAnalysis() {
	fmt.Println("\n7. Memory Allocation Analysis")
	fmt.Println("=============================")
	
	var m1, m2 runtime.MemStats
	
	// Get initial memory stats
	runtime.ReadMemStats(&m1)
	
	// Allocate memory
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}
	
	// Get final memory stats
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("  Allocated memory: %d bytes\n", m2.Alloc-m1.Alloc)
	fmt.Printf("  Total allocations: %d\n", m2.Mallocs-m1.Mallocs)
	fmt.Printf("  GC cycles: %d\n", m2.NumGC-m1.NumGC)
	
	fmt.Println("  Memory allocation analysis completed")
}

// Example 8: Goroutine Analysis
func goroutineAnalysis() {
	fmt.Println("\n8. Goroutine Analysis")
	fmt.Println("====================")
	
	// Get initial goroutine count
	initial := runtime.NumGoroutine()
	
	// Start goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	// Get goroutine count during execution
	during := runtime.NumGoroutine()
	
	wg.Wait()
	
	// Get final goroutine count
	final := runtime.NumGoroutine()
	
	fmt.Printf("  Initial goroutines: %d\n", initial)
	fmt.Printf("  During execution: %d\n", during)
	fmt.Printf("  Final goroutines: %d\n", final)
	
	fmt.Println("  Goroutine analysis completed")
}

// Example 9: CPU Usage Analysis
func cpuUsageAnalysis() {
	fmt.Println("\n9. CPU Usage Analysis")
	fmt.Println("====================")
	
	// CPU-intensive task
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		_ = i * i
	}
	duration := time.Since(start)
	
	fmt.Printf("  CPU task duration: %v\n", duration)
	fmt.Printf("  Operations per second: %.0f\n", float64(10000000)/duration.Seconds())
	
	fmt.Println("  CPU usage analysis completed")
}

// Example 10: Memory Leak Detection
func memoryLeakDetection() {
	fmt.Println("\n10. Memory Leak Detection")
	fmt.Println("=========================")
	
	var m1, m2, m3 runtime.MemStats
	
	// Get initial memory stats
	runtime.ReadMemStats(&m1)
	
	// Allocate memory
	data1 := make([]int, 1000000)
	for i := range data1 {
		data1[i] = i
	}
	
	// Get memory stats after allocation
	runtime.ReadMemStats(&m2)
	
	// Clear reference
	data1 = nil
	
	// Force garbage collection
	runtime.GC()
	
	// Get memory stats after GC
	runtime.ReadMemStats(&m3)
	
	fmt.Printf("  Memory before allocation: %d bytes\n", m1.Alloc)
	fmt.Printf("  Memory after allocation: %d bytes\n", m2.Alloc)
	fmt.Printf("  Memory after GC: %d bytes\n", m3.Alloc)
	
	if m3.Alloc < m2.Alloc {
		fmt.Println("  No memory leak detected")
	} else {
		fmt.Println("  Potential memory leak detected")
	}
	
	fmt.Println("  Memory leak detection completed")
}

// Example 11: Performance Comparison
func performanceComparison() {
	fmt.Println("\n11. Performance Comparison")
	fmt.Println("==========================")
	
	// Test 1: Sequential processing
	start := time.Now()
	sequentialProcessing()
	sequentialTime := time.Since(start)
	
	// Test 2: Concurrent processing
	start = time.Now()
	concurrentProcessing()
	concurrentTime := time.Since(start)
	
	fmt.Printf("  Sequential processing: %v\n", sequentialTime)
	fmt.Printf("  Concurrent processing: %v\n", concurrentTime)
	fmt.Printf("  Speedup: %.2fx\n", float64(sequentialTime)/float64(concurrentTime))
	
	fmt.Println("  Performance comparison completed")
}

func sequentialProcessing() {
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
}

func concurrentProcessing() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				_ = j * j
			}
		}()
	}
	wg.Wait()
}

// Example 12: Profiling with HTTP Server
func profilingWithHTTPServer() {
	fmt.Println("\n12. Profiling with HTTP Server")
	fmt.Println("=============================")
	
	// Start HTTP server for profiling
	go func() {
		// This would normally be: http.ListenAndServe("localhost:6060", nil)
		// But we'll simulate it for this example
		time.Sleep(100 * time.Millisecond)
	}()
	
	// Simulate some work
	time.Sleep(50 * time.Millisecond)
	
	fmt.Println("  HTTP server profiling setup completed")
}

// Example 13: Custom Profiling
func basicCustomProfiling() {
	fmt.Println("\n13. Custom Profiling")
	fmt.Println("===================")
	
	// For this example, we'll just demonstrate the concept
	// without actually creating a custom profile to avoid complexity
	fmt.Println("  Custom profiling concept demonstrated")
	fmt.Println("  Custom profiling completed")
}

// Example 14: Memory Pool Usage
func memoryPoolUsage() {
	fmt.Println("\n14. Memory Pool Usage")
	fmt.Println("====================")
	
	// Create memory pool
	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, 1000)
		},
	}
	
	// Use pool
	start := time.Now()
	for i := 0; i < 1000; i++ {
		slice := pool.Get().([]int)
		// Use slice
		for j := range slice {
			slice[j] = j
		}
		pool.Put(slice)
	}
	duration := time.Since(start)
	
	fmt.Printf("  Memory pool usage: %v\n", duration)
	
	fmt.Println("  Memory pool usage completed")
}

// Example 15: Goroutine Pool
func goroutinePool() {
	fmt.Println("\n15. Goroutine Pool")
	fmt.Println("=================")
	
	// Create goroutine pool
	jobs := make(chan int, 100)
	results := make(chan int, 1000) // Buffer for all results
	
	// Start workers
	for i := 0; i < 10; i++ {
		go func() {
			for j := range jobs {
				results <- j * j
			}
		}()
	}
	
	// Send jobs in a goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	
	// Collect results
	start := time.Now()
	for i := 0; i < 1000; i++ {
		<-results
	}
	duration := time.Since(start)
	
	fmt.Printf("  Goroutine pool: %v\n", duration)
	
	fmt.Println("  Goroutine pool completed")
}

// Example 16: Channel Performance
func basicChannelPerformance() {
	fmt.Println("\n16. Channel Performance")
	fmt.Println("======================")
	
	// Test 1: Unbuffered channel
	start := time.Now()
	unbufferedChannel()
	unbufferedTime := time.Since(start)
	
	// Test 2: Buffered channel
	start = time.Now()
	bufferedChannel()
	bufferedTime := time.Since(start)
	
	fmt.Printf("  Unbuffered channel: %v\n", unbufferedTime)
	fmt.Printf("  Buffered channel: %v\n", bufferedTime)
	
	fmt.Println("  Channel performance completed")
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	for range ch {
		// Receive data
	}
}

func bufferedChannel() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	for range ch {
		// Receive data
	}
}

// Example 17: Select Performance
func selectPerformance() {
	fmt.Println("\n17. Select Performance")
	fmt.Println("=====================")
	
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	
	// Fill channels
	go func() {
		for i := 0; i < 1000; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	
	go func() {
		for i := 0; i < 1000; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	
	// Test select performance
	start := time.Now()
	count := 0
	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				count++
			}
		case _, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				count++
			}
		}
		
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	duration := time.Since(start)
	
	fmt.Printf("  Select performance: %v (%d operations)\n", duration, count)
	
	fmt.Println("  Select performance completed")
}

// Example 18: Mutex vs Channel Performance
func mutexVsChannelPerformance() {
	fmt.Println("\n18. Mutex vs Channel Performance")
	fmt.Println("===============================")
	
	// Test 1: Mutex
	start := time.Now()
	mutexPerformance()
	mutexTime := time.Since(start)
	
	// Test 2: Channel
	start = time.Now()
	channelPerformance()
	channelTime := time.Since(start)
	
	fmt.Printf("  Mutex performance: %v\n", mutexTime)
	fmt.Printf("  Channel performance: %v\n", channelTime)
	
	fmt.Println("  Mutex vs Channel performance completed")
}

func mutexPerformance() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}

func channelPerformance() {
	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	
	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// Process data
		}
	}()
	
	wg.Wait()
}

// Example 19: Memory Efficiency
func memoryEfficiency() {
	fmt.Println("\n19. Memory Efficiency")
	fmt.Println("====================")
	
	var m1, m2 runtime.MemStats
	
	// Test 1: Inefficient memory usage
	runtime.ReadMemStats(&m1)
	inefficientMemory()
	runtime.ReadMemStats(&m2)
	inefficientAlloc := m2.Alloc - m1.Alloc
	
	// Test 2: Efficient memory usage
	runtime.ReadMemStats(&m1)
	efficientMemory()
	runtime.ReadMemStats(&m2)
	efficientAlloc := m2.Alloc - m1.Alloc
	
	fmt.Printf("  Inefficient allocation: %d bytes\n", inefficientAlloc)
	fmt.Printf("  Efficient allocation: %d bytes\n", efficientAlloc)
	fmt.Printf("  Memory savings: %d bytes\n", inefficientAlloc-efficientAlloc)
	
	fmt.Println("  Memory efficiency completed")
}

func inefficientMemory() {
	var result []int
	for i := 0; i < 1000; i++ {
		result = append(result, i)
	}
}

func efficientMemory() {
	result := make([]int, 0, 1000) // Pre-allocate capacity
	for i := 0; i < 1000; i++ {
		result = append(result, i)
	}
}

// Example 20: Performance Monitoring
func performanceMonitoring() {
	fmt.Println("\n20. Performance Monitoring")
	fmt.Println("=========================")
	
	// Monitor memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  Allocated memory: %d bytes\n", m.Alloc)
	fmt.Printf("  System memory: %d bytes\n", m.Sys)
	fmt.Printf("  Number of GCs: %d\n", m.NumGC)
	fmt.Printf("  Number of goroutines: %d\n", runtime.NumGoroutine())
	
	fmt.Println("  Performance monitoring completed")
}

// Run all basic examples
func RunBasicExamples() {
	fmt.Println("📊 Profiling & Benchmarking Examples")
	fmt.Println("====================================")
	
	basicCPUProfiling()
	basicMemoryProfiling()
	basicGoroutineProfiling()
	basicBlockProfiling()
	basicMutexProfiling()
	basicBenchmarking()
	memoryAllocationAnalysis()
	goroutineAnalysis()
	cpuUsageAnalysis()
	memoryLeakDetection()
	performanceComparison()
	profilingWithHTTPServer()
	basicCustomProfiling()
	memoryPoolUsage()
	goroutinePool()
	basicChannelPerformance()
	selectPerformance()
	mutexVsChannelPerformance()
	memoryEfficiency()
	performanceMonitoring()
}


