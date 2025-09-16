package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
	"time"
)

// Exercise 1: Implement CPU Profiling
func Exercise1() {
	fmt.Println("\nExercise 1: Implement CPU Profiling")
	fmt.Println("==================================")
	
	// TODO: Implement CPU profiling
	// 1. Create a CPU profile file
	// 2. Start CPU profiling
	// 3. Run a CPU-intensive task
	// 4. Stop CPU profiling
	
	f, err := os.Create("exercise_cpu.prof")
	if err != nil {
		fmt.Printf("  Exercise 1: Error creating CPU profile: %v\n", err)
		return
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("  Exercise 1: Error starting CPU profile: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

	// CPU-intensive task
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
	
	fmt.Println("  Exercise 1: CPU profiling completed")
}

// Exercise 2: Implement Memory Profiling
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Memory Profiling")
	fmt.Println("=====================================")
	
	// TODO: Implement memory profiling
	// 1. Create a memory profile file
	// 2. Run a memory-intensive task
	// 3. Write memory profile
	
	f, err := os.Create("exercise_mem.prof")
	if err != nil {
		fmt.Printf("  Exercise 2: Error creating memory profile: %v\n", err)
		return
	}
	defer f.Close()

	// Memory-intensive task
	var data [][]int
	for i := 0; i < 1000; i++ {
		slice := make([]int, 1000)
		data = append(data, slice)
	}

	runtime.GC() // Force garbage collection
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("  Exercise 2: Error writing memory profile: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 2: Memory profiling completed")
}

// Exercise 3: Implement Goroutine Profiling
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Goroutine Profiling")
	fmt.Println("=======================================")
	
	// TODO: Implement goroutine profiling
	// 1. Create a goroutine profile file
	// 2. Start multiple goroutines
	// 3. Write goroutine profile
	
	f, err := os.Create("exercise_goroutine.prof")
	if err != nil {
		fmt.Printf("  Exercise 3: Error creating goroutine profile: %v\n", err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	
	if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
		fmt.Printf("  Exercise 3: Error writing goroutine profile: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 3: Goroutine profiling completed")
}

// Exercise 4: Implement Block Profiling
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Block Profiling")
	fmt.Println("====================================")
	
	// TODO: Implement block profiling
	// 1. Create a block profile file
	// 2. Start block profiling
	// 3. Create blocking operations
	// 4. Write block profile
	
	f, err := os.Create("exercise_block.prof")
	if err != nil {
		fmt.Printf("  Exercise 4: Error creating block profile: %v\n", err)
		return
	}
	defer f.Close()

	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)

	ch := make(chan int)
	var wg sync.WaitGroup
	
	// Start goroutines that will block
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := <-ch
			_ = data
		}()
	}
	
	// Send data to unblock
	for i := 0; i < 5; i++ {
		ch <- i
	}
	
	wg.Wait()
	close(ch)
	
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Printf("  Exercise 4: Error writing block profile: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 4: Block profiling completed")
}

// Exercise 5: Implement Mutex Profiling
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Mutex Profiling")
	fmt.Println("====================================")
	
	// TODO: Implement mutex profiling
	// 1. Create a mutex profile file
	// 2. Start mutex profiling
	// 3. Create mutex contention
	// 4. Write mutex profile
	
	f, err := os.Create("exercise_mutex.prof")
	if err != nil {
		fmt.Printf("  Exercise 5: Error creating mutex profile: %v\n", err)
		return
	}
	defer f.Close()

	runtime.SetMutexProfileFraction(1)
	defer runtime.SetMutexProfileFraction(0)

	var mu sync.Mutex
	var wg sync.WaitGroup
	
	// Start multiple goroutines that will contend for mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				time.Sleep(1 * time.Microsecond)
				mu.Unlock()
			}
		}()
	}
	
	wg.Wait()
	
	if err := pprof.Lookup("mutex").WriteTo(f, 0); err != nil {
		fmt.Printf("  Exercise 5: Error writing mutex profile: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 5: Mutex profiling completed")
}

// Exercise 6: Implement Benchmarking
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Benchmarking")
	fmt.Println("=================================")
	
	// TODO: Implement benchmarking
	// 1. Create a benchmark function
	// 2. Run the benchmark
	// 3. Measure performance
	
	// Simple benchmark
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
	duration := time.Since(start)
	
	fmt.Printf("  Exercise 6: Benchmark duration: %v\n", duration)
	fmt.Printf("  Exercise 6: Operations per second: %.0f\n", float64(1000000)/duration.Seconds())
	
	fmt.Println("  Exercise 6: Benchmarking completed")
}

// Exercise 7: Implement Memory Analysis
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Memory Analysis")
	fmt.Println("====================================")
	
	// TODO: Implement memory analysis
	// 1. Get initial memory stats
	// 2. Allocate memory
	// 3. Get final memory stats
	// 4. Calculate memory usage
	
	var m1, m2 runtime.MemStats
	
	runtime.ReadMemStats(&m1)
	
	// Allocate memory
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}
	
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("  Exercise 7: Allocated memory: %d bytes\n", m2.Alloc-m1.Alloc)
	fmt.Printf("  Exercise 7: Total allocations: %d\n", m2.Mallocs-m1.Mallocs)
	fmt.Printf("  Exercise 7: GC cycles: %d\n", m2.NumGC-m1.NumGC)
	
	fmt.Println("  Exercise 7: Memory analysis completed")
}

// Exercise 8: Implement Performance Comparison
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Performance Comparison")
	fmt.Println("==========================================")
	
	// TODO: Implement performance comparison
	// 1. Test sequential processing
	// 2. Test concurrent processing
	// 3. Compare performance
	
	// Sequential processing
	start := time.Now()
	sequentialTask()
	sequentialTime := time.Since(start)
	
	// Concurrent processing
	start = time.Now()
	concurrentTask()
	concurrentTime := time.Since(start)
	
	fmt.Printf("  Exercise 8: Sequential time: %v\n", sequentialTime)
	fmt.Printf("  Exercise 8: Concurrent time: %v\n", concurrentTime)
	fmt.Printf("  Exercise 8: Speedup: %.2fx\n", float64(sequentialTime)/float64(concurrentTime))
	
	fmt.Println("  Exercise 8: Performance comparison completed")
}

func sequentialTask() {
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
}

func concurrentTask() {
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

// Exercise 9: Implement Memory Pool
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Memory Pool")
	fmt.Println("===============================")
	
	// TODO: Implement memory pool
	// 1. Create a memory pool
	// 2. Use the pool for allocations
	// 3. Measure performance improvement
	
	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, 1000)
		},
	}
	
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
	
	fmt.Printf("  Exercise 9: Memory pool duration: %v\n", duration)
	
	fmt.Println("  Exercise 9: Memory pool completed")
}

// Exercise 10: Implement Goroutine Pool
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Goroutine Pool")
	fmt.Println("====================================")
	
	// TODO: Implement goroutine pool
	// 1. Create a goroutine pool
	// 2. Use the pool for processing
	// 3. Measure performance
	
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
	
	fmt.Printf("  Exercise 10: Goroutine pool duration: %v\n", duration)
	
	fmt.Println("  Exercise 10: Goroutine pool completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Profiling & Benchmarking Exercises")
	fmt.Println("====================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}

// Benchmark functions for testing
func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i * i
	}
}

func BenchmarkConcurrentFunction(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = 1 * 1
		}
	})
}

func BenchmarkWithSetup(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = data[i%len(data)]
	}
}

func BenchmarkWithCleanup(b *testing.B) {
	data := make([]int, 1000)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = data[i%len(data)]
	}
	
	b.StopTimer()
	
	data = nil
}
