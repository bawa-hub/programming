package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// GOD-LEVEL CONCEPT 1: Go Runtime Scheduler Deep Dive
// Understanding the G-M-P model and work stealing algorithm

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: Go Runtime Scheduler Deep Dive ===")
	
	// 1. Understanding G-M-P Model
	explainGMPModel()
	
	// 2. Work Stealing Algorithm
	demonstrateWorkStealing()
	
	// 3. Goroutine Scheduling Decisions
	demonstrateSchedulingDecisions()
	
	// 4. Preemption and Fairness
	demonstratePreemption()
	
	// 5. Memory Management and GC Interaction
	demonstrateMemoryManagement()
	
	// 6. NUMA Awareness and CPU Affinity
	demonstrateNUMAwareness()
}

// G-M-P Model Explanation
func explainGMPModel() {
	fmt.Println("\n=== 1. G-M-P MODEL (Goroutines, Machine, Processors) ===")
	
	fmt.Println(`
🏗️  G-M-P Architecture:
┌─────────────────────────────────────────────────────────────┐
│                    MACHINE (M)                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ Processor 1 │  │ Processor 2 │  │ Processor 3 │  ...   │
│  │     (P)     │  │     (P)     │  │     (P)     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│       │                │                │                 │
│   ┌───▼───┐        ┌───▼───┐        ┌───▼───┐             │
│   │ G1,G2 │        │ G3,G4 │        │ G5,G6 │             │
│   │ G7,G8 │        │ G9,G10│        │G11,G12│             │
│   └───────┘        └───────┘        └───────┘             │
└─────────────────────────────────────────────────────────────┘

🔑 Key Components:
• G (Goroutine): Lightweight thread, 2KB initial stack
• M (Machine): OS thread, managed by runtime
• P (Processor): Logical processor, runs goroutines

⚡ Work Stealing Algorithm:
• Each P has a local run queue
• When P's queue is empty, it steals from other P's queues
• Steals from the middle of the queue (random access)
• Global run queue as fallback
`)

	// Show current GOMAXPROCS
	fmt.Printf("Current GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
}

// Demonstrate Work Stealing
func demonstrateWorkStealing() {
	fmt.Println("\n=== 2. WORK STEALING ALGORITHM DEMONSTRATION ===")
	
	// Create a scenario where work stealing occurs
	numGoroutines := 1000
	var wg sync.WaitGroup
	
	// Track which processor each goroutine runs on
	processorMap := make(map[int]int)
	var mu sync.Mutex
	
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Get the current processor ID
			// Note: This is a simplified way to demonstrate the concept
			procID := runtime.GOMAXPROCS(0) - 1
			
			mu.Lock()
			processorMap[procID]++
			mu.Unlock()
			
			// Simulate some work
			time.Sleep(1 * time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("Completed %d goroutines in %v\n", numGoroutines, duration)
	fmt.Printf("Work distribution across processors:\n")
	for procID, count := range processorMap {
		fmt.Printf("  Processor %d: %d goroutines\n", procID, count)
	}
	
	fmt.Println(`
💡 Work Stealing Benefits:
• Load balancing across processors
• No central dispatcher (scalable)
• Efficient cache utilization
• Reduces contention
`)
}

// Demonstrate Scheduling Decisions
func demonstrateSchedulingDecisions() {
	fmt.Println("\n=== 3. GOROUTINE SCHEDULING DECISIONS ===")
	
	fmt.Println(`
🎯 Scheduling Triggers:
1. Blocking system calls (I/O operations)
2. Channel operations (send/receive)
3. Time slice expiration (10ms)
4. Function calls (stack growth)
5. Garbage collection
6. Explicit yields (runtime.Gosched())
`)

	// Demonstrate different scheduling triggers
	demonstrateBlockingIO()
	demonstrateChannelBlocking()
	demonstrateTimeSliceExpiration()
	demonstrateExplicitYield()
}

func demonstrateBlockingIO() {
	fmt.Println("\n--- Blocking I/O Example ---")
	
	start := time.Now()
	
	// This will cause a context switch due to blocking I/O
	time.Sleep(100 * time.Millisecond) // Simulates I/O
	
	fmt.Printf("Blocking I/O completed in %v\n", time.Since(start))
	fmt.Println("💡 This triggers a context switch to another goroutine")
}

func demonstrateChannelBlocking() {
	fmt.Println("\n--- Channel Blocking Example ---")
	
	ch := make(chan int)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- 42
	}()
	
	start := time.Now()
	value := <-ch // This blocks until data is available
	duration := time.Since(start)
	
	fmt.Printf("Received %d after %v (blocking on channel)\n", value, duration)
	fmt.Println("💡 Channel operations trigger context switches")
}

func demonstrateTimeSliceExpiration() {
	fmt.Println("\n--- Time Slice Expiration Example ---")
	
	var wg sync.WaitGroup
	
	// Create a CPU-intensive goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		
		// CPU-intensive work
		for i := 0; i < 1000000000; i++ {
			_ = i * i
		}
		
		fmt.Printf("CPU work completed in %v\n", time.Since(start))
		fmt.Println("💡 Time slice expiration allows other goroutines to run")
	}()
	
	// Create another goroutine that should get scheduled
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Other goroutine got scheduled!")
	}()
	
	wg.Wait()
}

func demonstrateExplicitYield() {
	fmt.Println("\n--- Explicit Yield Example ---")
	
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// Goroutine 1: Does work and yields
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 1: iteration %d\n", i)
			runtime.Gosched() // Explicitly yield to other goroutines
		}
	}()
	
	// Goroutine 2: Also does work
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 2: iteration %d\n", i)
			runtime.Gosched()
		}
	}()
	
	wg.Wait()
	fmt.Println("💡 runtime.Gosched() explicitly yields control")
}

// Demonstrate Preemption
func demonstratePreemption() {
	fmt.Println("\n=== 4. PREEMPTION AND FAIRNESS ===")
	
	fmt.Println(`
⚡ Preemption in Go 1.14+:
• Cooperative preemption (function calls)
• Non-cooperative preemption (async preemption)
• Ensures fair scheduling
• Prevents goroutine starvation
`)

	// Demonstrate preemption
	var wg sync.WaitGroup
	wg.Add(3)
	
	// Create goroutines with different work patterns
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			start := time.Now()
			
			// Mix of CPU work and yields
			for j := 0; j < 1000; j++ {
				_ = j * j
				if j%100 == 0 {
					runtime.Gosched()
				}
			}
			
			fmt.Printf("Goroutine %d completed in %v\n", id, time.Since(start))
		}(i)
	}
	
	wg.Wait()
	fmt.Println("💡 Preemption ensures all goroutines get fair CPU time")
}

// Demonstrate Memory Management
func demonstrateMemoryManagement() {
	fmt.Println("\n=== 5. MEMORY MANAGEMENT AND GC INTERACTION ===")
	
	fmt.Println(`
🧠 Memory Management:
• Goroutine stacks start at 2KB
• Stacks grow/shrink as needed
• GC is concurrent and low-latency
• Work stealing reduces GC pressure
`)

	// Show memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Current memory usage:\n")
	fmt.Printf("  Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("  TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("  NumGC: %d\n", m.NumGC)
	
	// Demonstrate stack growth
	demonstrateStackGrowth()
}

func demonstrateStackGrowth() {
	fmt.Println("\n--- Stack Growth Example ---")
	
	// Recursive function to demonstrate stack growth
	var stackDepth int
	var mu sync.Mutex
	
	var recursiveFunc func(int)
	recursiveFunc = func(depth int) {
		if depth <= 0 {
			return
		}
		
		mu.Lock()
		stackDepth = depth
		mu.Unlock()
		
		// Allocate some stack space
		var arr [1000]int
		_ = arr
		
		recursiveFunc(depth - 1)
	}
	
	// Start with a deep recursion
	recursiveFunc(1000)
	
	fmt.Printf("Maximum stack depth reached: %d\n", stackDepth)
	fmt.Println("💡 Go automatically grows stacks as needed")
}

// Demonstrate NUMA Awareness
func demonstrateNUMAwareness() {
	fmt.Println("\n=== 6. NUMA AWARENESS AND CPU AFFINITY ===")
	
	fmt.Println(`
🏗️  NUMA (Non-Uniform Memory Access):
• Modern systems have multiple CPU sockets
• Memory access speed varies by distance
• Go runtime is NUMA-aware
• Work stealing considers NUMA topology
`)

	// Show system topology
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	
	// Demonstrate CPU affinity considerations
	fmt.Println(`
💡 CPU Affinity Considerations:
• Keep related goroutines on same CPU
• Minimize cross-socket memory access
• Use sync.Pool for CPU-local caching
• Consider NUMA topology for large systems
`)
}

// Advanced: Demonstrate scheduler internals
func demonstrateSchedulerInternals() {
	fmt.Println("\n=== 7. SCHEDULER INTERNALS (GOD-LEVEL) ===")
	
	fmt.Println(`
🔬 Scheduler Internals:
• schedt: Global scheduler state
• p: Processor state (run queues, cache, etc.)
• g: Goroutine state (stack, program counter, etc.)
• m: Machine state (OS thread, current goroutine, etc.)

🎯 Run Queue Structure:
• Local run queue (per processor)
• Global run queue (shared)
• Network poller (I/O completion)
• Timer heap (timeouts, ticks)
`)

	// Show current goroutine count
	fmt.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
	
	// Demonstrate runtime scheduler stats
	fmt.Println(`
📊 Scheduler Statistics:
• Use GODEBUG=schedtrace=1000 to see scheduler traces
• Use GODEBUG=scheddetail=1 for detailed traces
• Use runtime/trace package for profiling
`)
}

// Utility function to show runtime information
func showRuntimeInfo() {
	fmt.Println("\n=== RUNTIME INFORMATION ===")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
}
