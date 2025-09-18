# ⚙️ Level 4, Topic 4: Advanced Scheduling

## 🚀 Overview
Mastering Go's advanced scheduling is the pinnacle of concurrent programming expertise. This topic will take you deep into Go's scheduler internals, work stealing algorithms, preemption mechanisms, and scheduler-aware programming techniques.

---

## 📚 Table of Contents

1. [Go Scheduler Fundamentals](#go-scheduler-fundamentals)
2. [Work Stealing Algorithm](#work-stealing-algorithm)
3. [Preemption and Context Switching](#preemption-and-context-switching)
4. [GOMAXPROCS and CPU Affinity](#gomaxprocs-and-cpu-affinity)
5. [Scheduler Internals](#scheduler-internals)
6. [Performance Tuning](#performance-tuning)
7. [Scheduler-Aware Programming](#scheduler-aware-programming)
8. [Runtime Scheduling](#runtime-scheduling)
9. [Advanced Patterns](#advanced-patterns)
10. [Real-World Applications](#real-world-applications)
11. [Debugging and Profiling](#debugging-and-profiling)
12. [Best Practices](#best-practices)

---

## ⚙️ Go Scheduler Fundamentals

### What is the Go Scheduler?

The Go scheduler is a sophisticated runtime component that manages goroutines and their execution on OS threads. It implements an M:N threading model where M goroutines are multiplexed onto N OS threads.

### Key Components

#### 1. Goroutines (G)
- Lightweight threads managed by the Go runtime
- Stack size starts at 2KB and grows as needed
- Managed by the scheduler, not the OS

#### 2. OS Threads (M)
- Actual operating system threads
- Managed by the OS kernel
- Limited by GOMAXPROCS setting

#### 3. Logical Processors (P)
- Context for scheduling goroutines
- Number equals GOMAXPROCS
- Each P has a local run queue

#### 4. Global Run Queue
- Shared queue for all processors
- Used when local queues are empty

### Scheduler Architecture

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

// Example 1: Understanding Scheduler Components
func schedulerComponents() {
    fmt.Println("Go Scheduler Components")
    fmt.Println("======================")
    
    // Get current GOMAXPROCS
    maxProcs := runtime.GOMAXPROCS(0)
    fmt.Printf("GOMAXPROCS: %d\n", maxProcs)
    
    // Get number of goroutines
    numGoroutines := runtime.NumGoroutine()
    fmt.Printf("Number of goroutines: %d\n", numGoroutines)
    
    // Get number of CPUs
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    
    // Get number of CGO calls
    numCgo := runtime.NumCgoCall()
    fmt.Printf("Number of CGO calls: %d\n", numCgo)
}

// Example 2: Goroutine Lifecycle
func goroutineLifecycle() {
    fmt.Println("\nGoroutine Lifecycle")
    fmt.Println("==================")
    
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
    fmt.Println("\nScheduler Statistics")
    fmt.Println("===================")
    
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
    fmt.Printf("Number of OS threads: %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("GC cycles: %d\n", m.NumGC)
}
```

---

## 🔄 Work Stealing Algorithm

### Understanding Work Stealing

Work stealing is the core algorithm used by Go's scheduler to distribute work efficiently across processors. When a processor's local queue is empty, it "steals" work from other processors.

### Work Stealing Implementation

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: Basic Work Stealing
func basicWorkStealing() {
    fmt.Println("\nBasic Work Stealing")
    fmt.Println("==================")
    
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

// Example 2: Work Stealing with Local Queues
func workStealingWithLocalQueues() {
    fmt.Println("\nWork Stealing with Local Queues")
    fmt.Println("==============================")
    
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

// Example 3: Work Stealing Performance
func workStealingPerformance() {
    fmt.Println("\nWork Stealing Performance")
    fmt.Println("========================")
    
    // Test with different work distributions
    testCases := []struct {
        name        string
        workSize    int
        numWorkers  int
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
}
```

---

## ⏰ Preemption and Context Switching

### Understanding Preemption

Preemption is the mechanism by which the scheduler can interrupt a running goroutine to give other goroutines a chance to run. Go uses cooperative preemption in most cases.

### Preemption Examples

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

// Example 1: Cooperative Preemption
func cooperativePreemption() {
    fmt.Println("\nCooperative Preemption")
    fmt.Println("=====================")
    
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

// Example 2: Forced Preemption
func forcedPreemption() {
    fmt.Println("\nForced Preemption")
    fmt.Println("================")
    
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

// Example 3: Context Switching Overhead
func contextSwitchingOverhead() {
    fmt.Println("\nContext Switching Overhead")
    fmt.Println("=========================")
    
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
}
```

---

## 🔧 GOMAXPROCS and CPU Affinity

### Understanding GOMAXPROCS

GOMAXPROCS controls the maximum number of OS threads that can execute Go code simultaneously. It defaults to the number of CPU cores.

### GOMAXPROCS Examples

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: GOMAXPROCS Configuration
func gomaxprocsConfiguration() {
    fmt.Println("\nGOMAXPROCS Configuration")
    fmt.Println("=======================")
    
    // Get current GOMAXPROCS
    current := runtime.GOMAXPROCS(0)
    fmt.Printf("Current GOMAXPROCS: %d\n", current)
    
    // Get number of CPUs
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    
    // Set GOMAXPROCS to different values
    for _, maxProcs := range []int{1, 2, 4, 8} {
        if maxProcs <= numCPU*2 {
            runtime.GOMAXPROCS(maxProcs)
            fmt.Printf("Set GOMAXPROCS to %d\n", maxProcs)
            
            // Test performance
            start := time.Now()
            testConcurrency(maxProcs)
            duration := time.Since(start)
            
            fmt.Printf("  Performance with %d processors: %v\n", maxProcs, duration)
        }
    }
    
    // Reset to default
    runtime.GOMAXPROCS(0)
    fmt.Printf("Reset GOMAXPROCS to default: %d\n", runtime.GOMAXPROCS(0))
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

// Example 2: CPU Affinity Simulation
func cpuAffinitySimulation() {
    fmt.Println("\nCPU Affinity Simulation")
    fmt.Println("======================")
    
    // Simulate CPU affinity by pinning goroutines to specific processors
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    
    var wg sync.WaitGroup
    for i := 0; i < numCPU; i++ {
        wg.Add(1)
        go func(cpuID int) {
            defer wg.Done()
            
            // Simulate CPU-bound work
            for j := 0; j < 1000000; j++ {
                if j%100000 == 0 {
                    fmt.Printf("  CPU %d: %d\n", cpuID, j)
                }
                _ = j * j
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Println("  CPU affinity simulation completed")
}

// Example 3: GOMAXPROCS Performance Impact
func gomaxprocsPerformanceImpact() {
    fmt.Println("\nGOMAXPROCS Performance Impact")
    fmt.Println("============================")
    
    numCPU := runtime.NumCPU()
    testCases := []int{1, numCPU / 2, numCPU, numCPU * 2}
    
    for _, maxProcs := range testCases {
        if maxProcs > 0 {
            runtime.GOMAXPROCS(maxProcs)
            
            start := time.Now()
            testConcurrency(maxProcs)
            duration := time.Since(start)
            
            fmt.Printf("  GOMAXPROCS=%d: %v\n", maxProcs, duration)
        }
    }
    
    // Reset to default
    runtime.GOMAXPROCS(0)
}
```

---

## 🔍 Scheduler Internals

### Understanding Scheduler Internals

The Go scheduler uses several internal data structures and algorithms to manage goroutines efficiently.

### Scheduler Internals Examples

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: Scheduler States
func schedulerStates() {
    fmt.Println("\nScheduler States")
    fmt.Println("===============")
    
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

// Example 2: Run Queue Management
func runQueueManagement() {
    fmt.Println("\nRun Queue Management")
    fmt.Println("===================")
    
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

// Example 3: Scheduler Statistics
func schedulerStatistics() {
    fmt.Println("\nScheduler Statistics")
    fmt.Println("===================")
    
    // Get runtime statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
    fmt.Printf("Number of OS threads: %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("GC cycles: %d\n", m.NumGC)
    fmt.Printf("Heap size: %d bytes\n", m.HeapSys)
    fmt.Printf("Stack size: %d bytes\n", m.StackSys)
}
```

---

## ⚡ Performance Tuning

### Scheduler Performance Tuning

Optimizing scheduler performance involves understanding how to work with the scheduler rather than against it.

### Performance Tuning Examples

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: Optimal Goroutine Count
func optimalGoroutineCount() {
    fmt.Println("\nOptimal Goroutine Count")
    fmt.Println("======================")
    
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
}

// Example 2: Work Distribution Optimization
func workDistributionOptimization() {
    fmt.Println("\nWork Distribution Optimization")
    fmt.Println("=============================")
    
    // Test different work distribution strategies
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
        
        fmt.Printf("  %s: %v\n", strategy.name, duration)
    }
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

// Example 3: Memory vs CPU Bound Work
func memoryVsCpuBoundWork() {
    fmt.Println("\nMemory vs CPU Bound Work")
    fmt.Println("=======================")
    
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
```

---

## 🎯 Scheduler-Aware Programming

### Writing Scheduler-Aware Code

Scheduler-aware programming involves writing code that works well with Go's scheduler rather than fighting against it.

### Scheduler-Aware Examples

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: Cooperative Yielding
func cooperativeYielding() {
    fmt.Println("\nCooperative Yielding")
    fmt.Println("===================")
    
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

// Example 2: Avoiding Scheduler Contention
func avoidingSchedulerContention() {
    fmt.Println("\nAvoiding Scheduler Contention")
    fmt.Println("============================")
    
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
}

// Example 3: Work Stealing Optimization
func workStealingOptimization() {
    fmt.Println("\nWork Stealing Optimization")
    fmt.Println("=========================")
    
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
```

---

## 🔧 Runtime Scheduling

### Understanding Runtime Scheduling

The Go runtime provides several functions for interacting with the scheduler and controlling goroutine execution.

### Runtime Scheduling Examples

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Example 1: Runtime Scheduler Control
func runtimeSchedulerControl() {
    fmt.Println("\nRuntime Scheduler Control")
    fmt.Println("========================")
    
    // Get current GOMAXPROCS
    current := runtime.GOMAXPROCS(0)
    fmt.Printf("Current GOMAXPROCS: %d\n", current)
    
    // Set GOMAXPROCS
    newValue := runtime.GOMAXPROCS(4)
    fmt.Printf("Set GOMAXPROCS to 4, previous value: %d\n", newValue)
    
    // Get number of goroutines
    numGoroutines := runtime.NumGoroutine()
    fmt.Printf("Number of goroutines: %d\n", numGoroutines)
    
    // Get number of CPUs
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    
    // Reset GOMAXPROCS
    runtime.GOMAXPROCS(current)
    fmt.Printf("Reset GOMAXPROCS to %d\n", current)
}

// Example 2: Goroutine Yielding
func goroutineYielding() {
    fmt.Println("\nGoroutine Yielding")
    fmt.Println("=================")
    
    // Long-running computation
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
    
    // Other work
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Printf("  Other work: %d\n", i)
            time.Sleep(10 * time.Millisecond)
        }
    }()
    
    time.Sleep(200 * time.Millisecond)
    fmt.Println("  Goroutine yielding completed")
}

// Example 3: Scheduler Statistics
func schedulerStatistics() {
    fmt.Println("\nScheduler Statistics")
    fmt.Println("===================")
    
    // Get runtime statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
    fmt.Printf("Number of OS threads: %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("GC cycles: %d\n", m.NumGC)
    fmt.Printf("Heap size: %d bytes\n", m.HeapSys)
    fmt.Printf("Stack size: %d bytes\n", m.StackSys)
}

// Example 4: Goroutine Profiling
func goroutineProfiling() {
    fmt.Println("\nGoroutine Profiling")
    fmt.Println("==================")
    
    // Create some goroutines
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            time.Sleep(100 * time.Millisecond)
        }(i)
    }
    
    // Get goroutine count
    numGoroutines := runtime.NumGoroutine()
    fmt.Printf("Number of goroutines: %d\n", numGoroutines)
    
    wg.Wait()
    
    // Get final goroutine count
    numGoroutines = runtime.NumGoroutine()
    fmt.Printf("Final number of goroutines: %d\n", numGoroutines)
}
```

---

## 🎓 Summary

Mastering Go's advanced scheduling is essential for building high-performance concurrent applications. Key takeaways:

1. **Understand the scheduler** and its components (G, M, P)
2. **Use work stealing** effectively for load balancing
3. **Configure GOMAXPROCS** appropriately for your workload
4. **Write scheduler-aware code** that cooperates with the scheduler
5. **Profile and tune** scheduler performance
6. **Avoid common pitfalls** like too many goroutines
7. **Use runtime functions** to control and monitor the scheduler

Advanced scheduling provides the foundation for building ultra-high-performance concurrent systems! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different scheduling strategies
3. **Apply** techniques to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** advanced scheduling patterns

Ready to become an Advanced Scheduling expert? Let's dive into the implementation! 💪

