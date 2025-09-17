package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Example 1: Understanding Scheduler Components
func schedulerComponents() {
	fmt.Println("\n1. Understanding Scheduler Components")
	fmt.Println("====================================")
	
	// Get current GOMAXPROCS
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("  GOMAXPROCS: %d\n", maxProcs)
	
	// Get number of goroutines
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("  Number of goroutines: %d\n", numGoroutines)
	
	// Get number of CPUs
	numCPU := runtime.NumCPU()
	fmt.Printf("  Number of CPUs: %d\n", numCPU)
	
	// Get number of CGO calls
	numCgo := runtime.NumCgoCall()
	fmt.Printf("  Number of CGO calls: %d\n", numCgo)
	
	fmt.Println("  Scheduler components completed")
}

// Example 2: Goroutine Lifecycle
func goroutineLifecycle() {
	fmt.Println("\n2. Goroutine Lifecycle")
	fmt.Println("=====================")
	
	// Create a goroutine
	go func() {
		fmt.Println("  Goroutine started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Goroutine finished")
	}()
	
	// Wait for goroutine to complete
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  Main goroutine finished")
}

// Example 3: Scheduler Statistics
func schedulerStatistics() {
	fmt.Println("\n3. Scheduler Statistics")
	fmt.Println("=======================")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("  Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("  Number of OS threads: %d\n", runtime.NumCPU())
	fmt.Printf("  GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("  GC cycles: %d\n", m.NumGC)
	
	fmt.Println("  Scheduler statistics completed")
}

// Example 4: Basic Work Stealing
func basicWorkStealing() {
	fmt.Println("\n4. Basic Work Stealing")
	fmt.Println("======================")
	
	// Create work items
	work := make(chan int, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range work {
				fmt.Printf("  Worker %d processing job %d\n", workerID, job)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	// Send work
	for i := 0; i < 20; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
	fmt.Println("  All work completed")
}

// Example 5: Work Stealing with Local Queues
func workStealingWithLocalQueues() {
	fmt.Println("\n5. Work Stealing with Local Queues")
	fmt.Println("=================================")
	
	// Create local work queues
	numWorkers := runtime.GOMAXPROCS(0)
	localQueues := make([]chan int, numWorkers)
	
	for i := range localQueues {
		localQueues[i] = make(chan int, 10)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			for {
				select {
				case job := <-localQueues[workerID]:
					fmt.Printf("  Worker %d processing local job %d\n", workerID, job)
					time.Sleep(10 * time.Millisecond)
				case <-time.After(100 * time.Millisecond):
					// Try to steal work from other queues
					for j := 0; j < numWorkers; j++ {
						if j != workerID {
							select {
							case job := <-localQueues[j]:
								fmt.Printf("  Worker %d stole job %d from worker %d\n", workerID, job, j)
								time.Sleep(10 * time.Millisecond)
								goto continueWork
							default:
								continue
							}
						}
					}
					return // No work found
				}
			continueWork:
			}
		}(i)
	}
	
	// Distribute work
	for i := 0; i < 20; i++ {
		localQueues[i%numWorkers] <- i
	}
	
	// Close all queues
	for i := range localQueues {
		close(localQueues[i])
	}
	
	wg.Wait()
	fmt.Println("  All work completed")
}

// Example 6: Work Stealing Performance
func workStealingPerformance() {
	fmt.Println("\n6. Work Stealing Performance")
	fmt.Println("===========================")
	
	// Test with different work distributions
	testCases := []struct {
		name       string
		workSize   int
		numWorkers int
	}{
		{"Small work, many workers", 10, 8},
		{"Medium work, balanced", 50, 4},
		{"Large work, few workers", 100, 2},
	}
	
	for _, tc := range testCases {
		start := time.Now()
		
		work := make(chan int, tc.workSize)
		var wg sync.WaitGroup
		
		for i := 0; i < tc.numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for job := range work {
					_ = job * job // Simulate work
				}
			}()
		}
		
		for i := 0; i < tc.workSize; i++ {
			work <- i
		}
		close(work)
		
		wg.Wait()
		duration := time.Since(start)
		
		fmt.Printf("  %s: %v\n", tc.name, duration)
	}
	
	fmt.Println("  Work stealing performance completed")
}

// Example 7: Cooperative Preemption
func cooperativePreemption() {
	fmt.Println("\n7. Cooperative Preemption")
	fmt.Println("========================")
	
	// This goroutine will yield control at function calls
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("  Goroutine 1: %d\n", i)
			time.Sleep(10 * time.Millisecond) // Yields control
		}
	}()
	
	// This goroutine will also yield control
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("  Goroutine 2: %d\n", i)
			time.Sleep(10 * time.Millisecond) // Yields control
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Cooperative preemption completed")
}

// Example 8: Forced Preemption
func forcedPreemption() {
	fmt.Println("\n8. Forced Preemption")
	fmt.Println("===================")
	
	// This goroutine will not yield control voluntarily
	go func() {
		for i := 0; i < 1000000; i++ {
			if i%100000 == 0 {
				fmt.Printf("  Goroutine 1: %d\n", i)
			}
			// No yielding - will be preempted by scheduler
		}
	}()
	
	// This goroutine will also run
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("  Goroutine 2: %d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  Forced preemption completed")
}

// Example 9: Context Switching Overhead
func contextSwitchingOverhead() {
	fmt.Println("\n9. Context Switching Overhead")
	fmt.Println("============================")
	
	// Test with different numbers of goroutines
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
		
		fmt.Printf("  %d goroutines: %v\n", numGoroutines, duration)
	}
	
	fmt.Println("  Context switching overhead completed")
}

// Example 10: GOMAXPROCS Configuration
func gomaxprocsConfiguration() {
	fmt.Println("\n10. GOMAXPROCS Configuration")
	fmt.Println("============================")
	
	// Get current GOMAXPROCS
	current := runtime.GOMAXPROCS(0)
	fmt.Printf("  Current GOMAXPROCS: %d\n", current)
	
	// Get number of CPUs
	numCPU := runtime.NumCPU()
	fmt.Printf("  Number of CPUs: %d\n", numCPU)
	
	// Set GOMAXPROCS to different values
	for _, maxProcs := range []int{1, 2, 4, 8} {
		if maxProcs <= numCPU*2 {
			runtime.GOMAXPROCS(maxProcs)
			fmt.Printf("  Set GOMAXPROCS to %d\n", maxProcs)
			
			// Test performance
			start := time.Now()
			basicTestConcurrency(maxProcs)
			duration := time.Since(start)
			
			fmt.Printf("    Performance with %d processors: %v\n", maxProcs, duration)
		}
	}
	
	// Reset to default
	runtime.GOMAXPROCS(0)
	fmt.Printf("  Reset GOMAXPROCS to default: %d\n", runtime.GOMAXPROCS(0))
}

func basicTestConcurrency(maxProcs int) {
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

// Example 11: CPU Affinity Simulation
func cpuAffinitySimulation() {
	fmt.Println("\n11. CPU Affinity Simulation")
	fmt.Println("===========================")
	
	// Simulate CPU affinity by pinning goroutines to specific processors
	numCPU := runtime.NumCPU()
	fmt.Printf("  Number of CPUs: %d\n", numCPU)
	
	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(cpuID int) {
			defer wg.Done()
			
			// Simulate CPU-bound work
			for j := 0; j < 1000000; j++ {
				if j%100000 == 0 {
					fmt.Printf("    CPU %d: %d\n", cpuID, j)
				}
				_ = j * j
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("  CPU affinity simulation completed")
}

// Example 12: GOMAXPROCS Performance Impact
func gomaxprocsPerformanceImpact() {
	fmt.Println("\n12. GOMAXPROCS Performance Impact")
	fmt.Println("=================================")
	
	numCPU := runtime.NumCPU()
	testCases := []int{1, numCPU / 2, numCPU, numCPU * 2}
	
	for _, maxProcs := range testCases {
		if maxProcs > 0 {
			runtime.GOMAXPROCS(maxProcs)
			
		start := time.Now()
		basicTestConcurrency(maxProcs)
		duration := time.Since(start)
			
			fmt.Printf("  GOMAXPROCS=%d: %v\n", maxProcs, duration)
		}
	}
	
	// Reset to default
	runtime.GOMAXPROCS(0)
	fmt.Println("  GOMAXPROCS performance impact completed")
}

// Example 13: Scheduler States
func schedulerStates() {
	fmt.Println("\n13. Scheduler States")
	fmt.Println("===================")
	
	// Create goroutines in different states
	var wg sync.WaitGroup
	
	// Running goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 1: Running")
		time.Sleep(50 * time.Millisecond)
	}()
	
	// Blocked goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 2: Blocked on channel")
		ch := make(chan int)
		<-ch // This will block
	}()
	
	// Waiting goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 3: Waiting on WaitGroup")
		var wg2 sync.WaitGroup
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			time.Sleep(30 * time.Millisecond)
		}()
		wg2.Wait()
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Scheduler states demonstration completed")
}

// Example 14: Run Queue Management
func runQueueManagement() {
	fmt.Println("\n14. Run Queue Management")
	fmt.Println("=======================")
	
	// Create work with different priorities
	work := make(chan int, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range work {
				fmt.Printf("  Worker %d processing job %d\n", workerID, job)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	// Send work in batches
	for batch := 0; batch < 3; batch++ {
		fmt.Printf("  Sending batch %d\n", batch+1)
		for i := 0; i < 10; i++ {
			work <- batch*10 + i
		}
		time.Sleep(50 * time.Millisecond)
	}
	
	close(work)
	wg.Wait()
	fmt.Println("  Run queue management completed")
}

// Example 15: Optimal Goroutine Count
func optimalGoroutineCount() {
	fmt.Println("\n15. Optimal Goroutine Count")
	fmt.Println("==========================")
	
	numCPU := runtime.NumCPU()
	testCases := []int{numCPU, numCPU * 2, numCPU * 4, numCPU * 8}
	
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
		
		fmt.Printf("  %d goroutines: %v\n", numGoroutines, duration)
	}
	
	fmt.Println("  Optimal goroutine count completed")
}

// Example 16: Work Distribution Optimization
func workDistributionOptimization() {
	fmt.Println("\n16. Work Distribution Optimization")
	fmt.Println("=================================")
	
	// Test different work distribution strategies
	strategies := []struct {
		name string
		fn   func(int)
	}{
		{"Equal distribution", basicEqualDistribution},
		{"Chunked distribution", basicChunkedDistribution},
		{"Dynamic distribution", basicDynamicDistribution},
	}
	
	for _, strategy := range strategies {
		start := time.Now()
		strategy.fn(1000)
		duration := time.Since(start)
		
		fmt.Printf("  %s: %v\n", strategy.name, duration)
	}
	
	fmt.Println("  Work distribution optimization completed")
}

func basicEqualDistribution(workSize int) {
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

func basicChunkedDistribution(workSize int) {
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

func basicDynamicDistribution(workSize int) {
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

// Example 17: Memory vs CPU Bound Work
func memoryVsCpuBoundWork() {
	fmt.Println("\n17. Memory vs CPU Bound Work")
	fmt.Println("===========================")
	
	// CPU bound work
	start := time.Now()
	cpuBoundWork(1000000)
	cpuDuration := time.Since(start)
	
	// Memory bound work
	start = time.Now()
	memoryBoundWork(1000000)
	memoryDuration := time.Since(start)
	
	fmt.Printf("  CPU bound work: %v\n", cpuDuration)
	fmt.Printf("  Memory bound work: %v\n", memoryDuration)
	
	fmt.Println("  Memory vs CPU bound work completed")
}

func cpuBoundWork(size int) {
	for i := 0; i < size; i++ {
		_ = i * i
	}
}

func memoryBoundWork(size int) {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	_ = data
}

// Example 18: Cooperative Yielding
func cooperativeYielding() {
	fmt.Println("\n18. Cooperative Yielding")
	fmt.Println("=======================")
	
	// Long-running computation with yielding
	go func() {
		for i := 0; i < 1000000; i++ {
			if i%100000 == 0 {
				fmt.Printf("  Computation: %d\n", i)
				runtime.Gosched() // Yield to other goroutines
			}
			_ = i * i
		}
		fmt.Println("  Computation completed")
	}()
	
	// Other goroutines can run
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("  Other work: %d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  Cooperative yielding completed")
}

// Example 19: Avoiding Scheduler Contention
func avoidingSchedulerContention() {
	fmt.Println("\n19. Avoiding Scheduler Contention")
	fmt.Println("=================================")
	
	// Bad: Too many goroutines
	fmt.Println("  Bad: Too many goroutines")
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
	fmt.Println("  Good: Optimal number of goroutines")
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
	
	fmt.Printf("  Bad approach: %v\n", badDuration)
	fmt.Printf("  Good approach: %v\n", goodDuration)
	fmt.Printf("  Improvement: %.2fx\n", float64(badDuration)/float64(goodDuration))
	
	fmt.Println("  Avoiding scheduler contention completed")
}

// Example 20: Work Stealing Optimization
func workStealingOptimization() {
	fmt.Println("\n20. Work Stealing Optimization")
	fmt.Println("=============================")
	
	// Create work with different characteristics
	work := make(chan int, 1000)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range work {
				// Simulate work with different durations
				if job%10 == 0 {
					time.Sleep(1 * time.Millisecond) // Long work
				} else {
					_ = job * job // Short work
				}
			}
		}(i)
	}
	
	// Send work
	for i := 0; i < 1000; i++ {
		work <- i
	}
	close(work)
	
	wg.Wait()
	fmt.Println("  Work stealing optimization completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("⚙️ Advanced Scheduling Examples")
	fmt.Println("===============================")
	
	schedulerComponents()
	goroutineLifecycle()
	schedulerStatistics()
	basicWorkStealing()
	workStealingWithLocalQueues()
	workStealingPerformance()
	cooperativePreemption()
	forcedPreemption()
	contextSwitchingOverhead()
	gomaxprocsConfiguration()
	cpuAffinitySimulation()
	gomaxprocsPerformanceImpact()
	schedulerStates()
	runQueueManagement()
	optimalGoroutineCount()
	workDistributionOptimization()
	memoryVsCpuBoundWork()
	cooperativeYielding()
	avoidingSchedulerContention()
	workStealingOptimization()
}
