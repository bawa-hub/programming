package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Exercise 1: Implement Basic Work Stealing
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Basic Work Stealing")
	fmt.Println("=======================================")
	
	// TODO: Implement basic work stealing
	// 1. Create a work stealing queue
	// 2. Implement push and pop operations
	// 3. Test with concurrent access
	
	wsq := NewExerciseWorkStealingQueue(10)
	
	// Test basic operations
	for i := 0; i < 5; i++ {
		success := wsq.Push(i)
		fmt.Printf("  Exercise 1: Pushed %d: %t\n", i, success)
	}
	
	// Pop work
	for i := 0; i < 5; i++ {
		work, ok := wsq.Pop()
		if ok {
			fmt.Printf("  Exercise 1: Popped: %v\n", work)
		} else {
			fmt.Println("  Exercise 1: No work available")
		}
	}
	
	fmt.Println("  Exercise 1: Basic work stealing completed")
}

type ExerciseWorkStealingQueue struct {
	tasks    []interface{}
	head     int64
	tail     int64
	capacity int64
}

func NewExerciseWorkStealingQueue(capacity int) *ExerciseWorkStealingQueue {
	return &ExerciseWorkStealingQueue{
		tasks:    make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (wsq *ExerciseWorkStealingQueue) Push(task interface{}) bool {
	currentTail := wsq.tail
	nextTail := (currentTail + 1) % wsq.capacity
	
	if nextTail == wsq.head {
		return false // Queue full
	}
	
	wsq.tasks[currentTail] = task
	wsq.tail = nextTail
	return true
}

func (wsq *ExerciseWorkStealingQueue) Pop() (interface{}, bool) {
	currentTail := wsq.tail
	currentHead := wsq.head
	
	if currentHead == currentTail {
		return nil, false // Queue empty
	}
	
	// Try to pop from tail
	newTail := (currentTail - 1 + wsq.capacity) % wsq.capacity
	if newTail != currentHead {
		wsq.tail = newTail
		task := wsq.tasks[newTail]
		return task, true
	}
	
	return nil, false
}

// Exercise 2: Implement GOMAXPROCS Tuning
func Exercise2() {
	fmt.Println("\nExercise 2: Implement GOMAXPROCS Tuning")
	fmt.Println("======================================")
	
	// TODO: Implement GOMAXPROCS tuning
	// 1. Test different GOMAXPROCS values
	// 2. Measure performance impact
	// 3. Find optimal configuration
	
	numCPU := runtime.NumCPU()
	testCases := []int{1, numCPU / 2, numCPU, numCPU * 2}
	
	for _, maxProcs := range testCases {
		if maxProcs > 0 {
			runtime.GOMAXPROCS(maxProcs)
			
			start := time.Now()
			testConcurrency(maxProcs)
			duration := time.Since(start)
			
			fmt.Printf("  Exercise 2: GOMAXPROCS=%d: %v\n", maxProcs, duration)
		}
	}
	
	// Reset to default
	runtime.GOMAXPROCS(0)
	fmt.Println("  Exercise 2: GOMAXPROCS tuning completed")
}

func testConcurrency(maxProcs int) {
	var wg sync.WaitGroup
	work := make(chan int, maxProcs*10)
	
	// Start workers
	for i := 0; i < maxProcs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range work {
				_ = job * job
			}
		}()
	}
	
	// Send work
	for i := 0; i < 1000; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
}

// Exercise 3: Implement Cooperative Yielding
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Cooperative Yielding")
	fmt.Println("========================================")
	
	// TODO: Implement cooperative yielding
	// 1. Create long-running computation
	// 2. Add yielding points
	// 3. Test with other goroutines
	
	// Long-running computation with yielding
	go func() {
		for i := 0; i < 1000000; i++ {
			if i%100000 == 0 {
				fmt.Printf("  Exercise 3: Computation: %d\n", i)
				runtime.Gosched() // Yield to other goroutines
			}
			_ = i * i
		}
		fmt.Println("  Exercise 3: Computation completed")
	}()
	
	// Other goroutines can run
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("  Exercise 3: Other work: %d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  Exercise 3: Cooperative yielding completed")
}

// Exercise 4: Implement Work Distribution
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Work Distribution")
	fmt.Println("======================================")
	
	// TODO: Implement work distribution
	// 1. Test different distribution strategies
	// 2. Measure performance
	// 3. Compare approaches
	
	strategies := []struct {
		name string
		fn   func(int)
	}{
		{"Equal distribution", equalDistribution},
		{"Chunked distribution", chunkedDistribution},
		{"Dynamic distribution", dynamicDistribution},
	}
	
	for _, strategy := range strategies {
		start := time.Now()
		strategy.fn(1000)
		duration := time.Since(start)
		
		fmt.Printf("  Exercise 4: %s: %v\n", strategy.name, duration)
	}
	
	fmt.Println("  Exercise 4: Work distribution completed")
}

func equalDistribution(workSize int) {
	var wg sync.WaitGroup
	work := make(chan int, workSize)
	
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range work {
				_ = job * job
			}
		}()
	}
	
	for i := 0; i < workSize; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
}

func chunkedDistribution(workSize int) {
	numWorkers := runtime.GOMAXPROCS(0)
	chunkSize := workSize / numWorkers
	
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				_ = j * j
			}
		}(i*chunkSize, (i+1)*chunkSize)
	}
	
	wg.Wait()
}

func dynamicDistribution(workSize int) {
	var wg sync.WaitGroup
	work := make(chan int, workSize)
	
	// Start more workers than processors
	for i := 0; i < runtime.GOMAXPROCS(0)*2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range work {
				_ = job * job
			}
		}()
	}
	
	for i := 0; i < workSize; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
}

// Exercise 5: Implement Scheduler Statistics
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Scheduler Statistics")
	fmt.Println("========================================")
	
	// TODO: Implement scheduler statistics
	// 1. Monitor goroutine count
	// 2. Track GOMAXPROCS changes
	// 3. Measure performance metrics
	
	// Get initial statistics
	numGoroutines := runtime.NumGoroutine()
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	
	fmt.Printf("  Exercise 5: Initial goroutines: %d\n", numGoroutines)
	fmt.Printf("  Exercise 5: GOMAXPROCS: %d\n", maxProcs)
	fmt.Printf("  Exercise 5: CPUs: %d\n", numCPU)
	
	// Create some goroutines
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	// Get statistics after creating goroutines
	numGoroutines = runtime.NumGoroutine()
	fmt.Printf("  Exercise 5: After creating goroutines: %d\n", numGoroutines)
	
	wg.Wait()
	
	// Get final statistics
	numGoroutines = runtime.NumGoroutine()
	fmt.Printf("  Exercise 5: Final goroutines: %d\n", numGoroutines)
	
	fmt.Println("  Exercise 5: Scheduler statistics completed")
}

// Exercise 6: Implement Context Switching Analysis
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Context Switching Analysis")
	fmt.Println("===============================================")
	
	// TODO: Implement context switching analysis
	// 1. Test with different numbers of goroutines
	// 2. Measure context switching overhead
	// 3. Find optimal goroutine count
	
	testCases := []int{1, 10, 100, 1000}
	
	for _, numGoroutines := range testCases {
		start := time.Now()
		
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for j := 0; j < 1000; j++ {
					_ = id * j
				}
			}(i)
		}
		
		wg.Wait()
		duration := time.Since(start)
		
		fmt.Printf("  Exercise 6: %d goroutines: %v\n", numGoroutines, duration)
	}
	
	fmt.Println("  Exercise 6: Context switching analysis completed")
}

// Exercise 7: Implement Work Stealing Performance
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Work Stealing Performance")
	fmt.Println("==============================================")
	
	// TODO: Implement work stealing performance test
	// 1. Test with different work sizes
	// 2. Measure work stealing efficiency
	// 3. Compare with non-work-stealing approach
	
	testCases := []struct {
		name     string
		workSize int
		workers  int
	}{
		{"Small work", 100, 4},
		{"Medium work", 1000, 4},
		{"Large work", 10000, 4},
	}
	
	for _, tc := range testCases {
		start := time.Now()
		
		work := make(chan int, tc.workSize)
		var wg sync.WaitGroup
		
		for i := 0; i < tc.workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for job := range work {
					_ = job * job
				}
			}()
		}
		
		for i := 0; i < tc.workSize; i++ {
			work <- i
		}
		close(work)
		
		wg.Wait()
		duration := time.Since(start)
		
		fmt.Printf("  Exercise 7: %s: %v\n", tc.name, duration)
	}
	
	fmt.Println("  Exercise 7: Work stealing performance completed")
}

// Exercise 8: Implement CPU Affinity Simulation
func Exercise8() {
	fmt.Println("\nExercise 8: Implement CPU Affinity Simulation")
	fmt.Println("============================================")
	
	// TODO: Implement CPU affinity simulation
	// 1. Pin goroutines to specific processors
	// 2. Simulate CPU-bound work
	// 3. Measure performance impact
	
	numCPU := runtime.NumCPU()
	fmt.Printf("  Exercise 8: Number of CPUs: %d\n", numCPU)
	
	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(cpuID int) {
			defer wg.Done()
			
			// Simulate CPU-bound work
			for j := 0; j < 1000000; j++ {
				if j%100000 == 0 {
					fmt.Printf("    Exercise 8: CPU %d: %d\n", cpuID, j)
				}
				_ = j * j
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("  Exercise 8: CPU affinity simulation completed")
}

// Exercise 9: Implement Scheduler Contention Avoidance
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Scheduler Contention Avoidance")
	fmt.Println("==================================================")
	
	// TODO: Implement scheduler contention avoidance
	// 1. Compare too many vs optimal goroutines
	// 2. Measure contention impact
	// 3. Find optimal goroutine count
	
	// Bad: Too many goroutines
	fmt.Println("  Exercise 9: Bad: Too many goroutines")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = id * id
		}(i)
	}
	wg.Wait()
	badDuration := time.Since(start)
	
	// Good: Optimal number of goroutines
	fmt.Println("  Exercise 9: Good: Optimal number of goroutines")
	start = time.Now()
	numWorkers := runtime.GOMAXPROCS(0)
	work := make(chan int, 10000)
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range work {
				_ = job * job
			}
		}()
	}
	
	for i := 0; i < 10000; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
	goodDuration := time.Since(start)
	
	fmt.Printf("  Exercise 9: Bad approach: %v\n", badDuration)
	fmt.Printf("  Exercise 9: Good approach: %v\n", goodDuration)
	fmt.Printf("  Exercise 9: Improvement: %.2fx\n", float64(badDuration)/float64(goodDuration))
	
	fmt.Println("  Exercise 9: Scheduler contention avoidance completed")
}

// Exercise 10: Implement Performance Optimization
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Performance Optimization")
	fmt.Println("==============================================")
	
	// TODO: Implement performance optimization
	// 1. Test different optimization strategies
	// 2. Measure performance improvements
	// 3. Compare approaches
	
	// Test 1: Basic approach
	start := time.Now()
	basicApproach(1000)
	basicDuration := time.Since(start)
	
	// Test 2: Optimized approach
	start = time.Now()
	optimizedApproach(1000)
	optimizedDuration := time.Since(start)
	
	fmt.Printf("  Exercise 10: Basic approach: %v\n", basicDuration)
	fmt.Printf("  Exercise 10: Optimized approach: %v\n", optimizedDuration)
	fmt.Printf("  Exercise 10: Improvement: %.2fx\n", float64(basicDuration)/float64(optimizedDuration))
	
	fmt.Println("  Exercise 10: Performance optimization completed")
}

func basicApproach(workSize int) {
	var wg sync.WaitGroup
	for i := 0; i < workSize; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = id * id
		}(i)
	}
	wg.Wait()
}

func optimizedApproach(workSize int) {
	numWorkers := runtime.GOMAXPROCS(0)
	work := make(chan int, workSize)
	var wg sync.WaitGroup
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range work {
				_ = job * job
			}
		}()
	}
	
	for i := 0; i < workSize; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Advanced Scheduling Exercises")
	fmt.Println("===============================")
	
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

