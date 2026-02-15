# 24 - Performance Optimization

## 🎯 Learning Objectives
- Understand how to optimize concurrent Go applications
- Learn about different performance optimization techniques
- Master profiling and benchmarking
- Practice with performance optimization strategies
- Understand when to use different optimization approaches

## 📚 Theory

### Why Optimize Concurrent Applications?

**Performance bottlenecks in concurrent applications:**
1. **Lock contention**: Too much contention on mutexes
2. **Channel overhead**: Excessive channel operations
3. **Goroutine overhead**: Too many goroutines
4. **Memory allocation**: Excessive memory allocation
5. **CPU cache misses**: Poor data locality

**Benefits of optimization:**
1. **Better performance**: Faster execution
2. **Lower resource usage**: Less memory and CPU
3. **Better scalability**: Handle more load
4. **Cost efficiency**: Lower infrastructure costs
5. **User experience**: Faster response times

### Optimization Techniques

1. **Profiling**: Identify performance bottlenecks
2. **Benchmarking**: Measure performance improvements
3. **Memory optimization**: Reduce allocations
4. **CPU optimization**: Improve CPU usage
5. **Concurrency optimization**: Optimize concurrent patterns

## 💻 Code Examples

### Example 1: Basic Profiling

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func basicProfiling() {
    fmt.Println("=== Basic Profiling ===")
    
    // Start profiling
    start := time.Now()
    startMem := getMemStats()
    
    // Simulate work
    result := 0
    for i := 0; i < 1000000; i++ {
        result += i
    }
    
    // End profiling
    end := time.Now()
    endMem := getMemStats()
    
    // Print results
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Memory before: %d bytes\n", startMem.Alloc)
    fmt.Printf("Memory after: %d bytes\n", endMem.Alloc)
    fmt.Printf("Memory delta: %d bytes\n", endMem.Alloc-startMem.Alloc)
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
    fmt.Printf("Result: %d\n", result)
}

func getMemStats() runtime.MemStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m
}

func main() {
    basicProfiling()
}
```

**Run this code:**
```bash
go run 24-performance-optimization.go
```

### Example 2: Benchmarking

```go
package main

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

// Benchmark function
func BenchmarkConcurrentWork(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concurrentWork()
    }
}

func BenchmarkSequentialWork(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sequentialWork()
    }
}

func concurrentWork() {
    var wg sync.WaitGroup
    results := make(chan int, 10)
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            results <- id * 2
        }(i)
    }
    
    wg.Wait()
    close(results)
    
    sum := 0
    for result := range results {
        sum += result
    }
}

func sequentialWork() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i * 2
    }
}

func benchmarking() {
    fmt.Println("=== Benchmarking ===")
    
    // Run benchmarks
    fmt.Println("Running concurrent work benchmark...")
    start := time.Now()
    for i := 0; i < 1000; i++ {
        concurrentWork()
    }
    concurrentTime := time.Since(start)
    
    fmt.Println("Running sequential work benchmark...")
    start = time.Now()
    for i := 0; i < 1000; i++ {
        sequentialWork()
    }
    sequentialTime := time.Since(start)
    
    fmt.Printf("Concurrent time: %v\n", concurrentTime)
    fmt.Printf("Sequential time: %v\n", sequentialTime)
    fmt.Printf("Speedup: %.2fx\n", float64(sequentialTime)/float64(concurrentTime))
}

func main() {
    benchmarking()
}
```

### Example 3: Memory Optimization

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func memoryOptimization() {
    fmt.Println("=== Memory Optimization ===")
    
    // Bad: excessive allocations
    fmt.Println("Bad approach - excessive allocations:")
    start := time.Now()
    startMem := getMemStats()
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Bad: creating new slice for each goroutine
            data := make([]int, 1000)
            for j := range data {
                data[j] = j * id
            }
        }(i)
    }
    
    wg.Wait()
    end := time.Now()
    endMem := getMemStats()
    
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Memory delta: %d bytes\n", endMem.Alloc-startMem.Alloc)
    
    // Good: reuse objects
    fmt.Println("\nGood approach - reuse objects:")
    start = time.Now()
    startMem = getMemStats()
    
    // Reuse slice
    data := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Good: reuse existing slice
            for j := range data {
                data[j] = j * id
            }
        }(i)
    }
    
    wg.Wait()
    end = time.Now()
    endMem = getMemStats()
    
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Memory delta: %d bytes\n", endMem.Alloc-startMem.Alloc)
}

func getMemStats() runtime.MemStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m
}

func main() {
    memoryOptimization()
}
```

### Example 4: CPU Optimization

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func cpuOptimization() {
    fmt.Println("=== CPU Optimization ===")
    
    // Bad: too many goroutines
    fmt.Println("Bad approach - too many goroutines:")
    start := time.Now()
    
    var wg sync.WaitGroup
    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Simple work
            _ = id * 2
        }(i)
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    
    // Good: appropriate number of goroutines
    fmt.Println("\nGood approach - appropriate number of goroutines:")
    start = time.Now()
    
    numWorkers := runtime.NumCPU()
    jobs := make(chan int, 1000)
    results := make(chan int, 1000)
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- job * 2
            }
        }()
    }
    
    // Send jobs
    go func() {
        for i := 0; i < 10000; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Close results when done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for range results {
        // Process result
    }
    
    end = time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Workers: %d\n", numWorkers)
}

func main() {
    cpuOptimization()
}
```

### Example 5: Channel Optimization

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func channelOptimization() {
    fmt.Println("=== Channel Optimization ===")
    
    // Bad: unbuffered channels
    fmt.Println("Bad approach - unbuffered channels:")
    start := time.Now()
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Bad: unbuffered channel
            ch := make(chan int)
            go func() {
                ch <- id * 2
            }()
            result := <-ch
            _ = result
        }(i)
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    
    // Good: buffered channels
    fmt.Println("\nGood approach - buffered channels:")
    start = time.Now()
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Good: buffered channel
            ch := make(chan int, 1)
            ch <- id * 2
            result := <-ch
            _ = result
        }(i)
    }
    
    wg.Wait()
    end = time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
}

func main() {
    channelOptimization()
}
```

### Example 6: Lock Optimization

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func lockOptimization() {
    fmt.Println("=== Lock Optimization ===")
    
    // Bad: coarse-grained locking
    fmt.Println("Bad approach - coarse-grained locking:")
    start := time.Now()
    
    var counter int
    var mutex sync.Mutex
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            // Bad: hold lock for long time
            mutex.Lock()
            for j := 0; j < 100; j++ {
                counter++
            }
            mutex.Unlock()
        }()
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Counter: %d\n", counter)
    
    // Good: fine-grained locking
    fmt.Println("\nGood approach - fine-grained locking:")
    start = time.Now()
    
    counter = 0
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            // Good: hold lock for short time
            for j := 0; j < 100; j++ {
                mutex.Lock()
                counter++
                mutex.Unlock()
            }
        }()
    }
    
    wg.Wait()
    end = time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Counter: %d\n", counter)
}

func main() {
    lockOptimization()
}
```

### Example 7: Atomic Operations Optimization

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func atomicOptimization() {
    fmt.Println("=== Atomic Operations Optimization ===")
    
    // Bad: using mutex for simple counter
    fmt.Println("Bad approach - using mutex for simple counter:")
    start := time.Now()
    
    var counter int64
    var mutex sync.Mutex
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                mutex.Lock()
                counter++
                mutex.Unlock()
            }
        }()
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Counter: %d\n", counter)
    
    // Good: using atomic operations
    fmt.Println("\nGood approach - using atomic operations:")
    start = time.Now()
    
    counter = 0
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                atomic.AddInt64(&counter, 1)
            }
        }()
    }
    
    wg.Wait()
    end = time.Now()
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Counter: %d\n", counter)
}

func main() {
    atomicOptimization()
}
```

### Example 8: Performance Optimization Best Practices

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

type OptimizedWorker struct {
    id       int
    jobs     <-chan int
    results  chan<- int
    wg       *sync.WaitGroup
}

func (w *OptimizedWorker) Process() {
    defer w.wg.Done()
    
    for job := range w.jobs {
        // Process job
        result := job * 2
        w.results <- result
    }
}

func performanceOptimizationBestPractices() {
    fmt.Println("=== Performance Optimization Best Practices ===")
    
    // Configuration
    numWorkers := runtime.NumCPU()
    numJobs := 10000
    bufferSize := 100
    
    // Create channels with appropriate buffer size
    jobs := make(chan int, bufferSize)
    results := make(chan int, bufferSize)
    
    // Start workers
    var wg sync.WaitGroup
    workers := make([]*OptimizedWorker, numWorkers)
    
    for i := 0; i < numWorkers; i++ {
        workers[i] = &OptimizedWorker{
            id:      i,
            jobs:    jobs,
            results: results,
            wg:      &wg,
        }
        wg.Add(1)
        go workers[i].Process()
    }
    
    // Send jobs
    go func() {
        for i := 0; i < numJobs; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Close results when done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    start := time.Now()
    count := 0
    for range results {
        count++
    }
    end := time.Now()
    
    fmt.Printf("Processed %d jobs in %v\n", count, end.Sub(start))
    fmt.Printf("Workers: %d\n", numWorkers)
    fmt.Printf("Jobs per second: %.2f\n", float64(count)/end.Sub(start).Seconds())
}

func main() {
    performanceOptimizationBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Profiling**: Identify performance bottlenecks
2. **Benchmarking**: Measure performance improvements
3. **Memory optimization**: Reduce allocations
4. **CPU optimization**: Improve CPU usage
5. **Concurrency optimization**: Optimize concurrent patterns
6. **Lock optimization**: Minimize lock contention
7. **Atomic operations**: Use for simple operations

## 🎯 When to Use Different Optimization Techniques

1. **Profiling**: When you need to identify bottlenecks
2. **Benchmarking**: When you need to measure improvements
3. **Memory optimization**: When you have memory issues
4. **CPU optimization**: When you have CPU issues
5. **Concurrency optimization**: When you have concurrency issues

## 🎯 Best Practices

1. **Profile before optimizing**:
   ```go
   // GOOD - profile first
   go tool pprof cpu.prof
   
   // BAD - optimize without profiling
   // Guess what to optimize
   ```

2. **Use appropriate number of goroutines**:
   ```go
   // GOOD - use CPU count
   numWorkers := runtime.NumCPU()
   
   // BAD - too many goroutines
   numWorkers := 10000
   ```

3. **Use buffered channels**:
   ```go
   // GOOD - buffered channel
   ch := make(chan int, 100)
   
   // BAD - unbuffered channel
   ch := make(chan int)
   ```

4. **Use atomic operations for simple operations**:
   ```go
   // GOOD - atomic operations
   atomic.AddInt64(&counter, 1)
   
   // BAD - mutex for simple operation
   mutex.Lock()
   counter++
   mutex.Unlock()
   ```

## 🎯 Common Pitfalls

1. **Optimizing without profiling**:
   ```go
   // BAD - optimize without knowing what's slow
   // Guess what to optimize
   
   // GOOD - profile first
   go tool pprof cpu.prof
   ```

2. **Using too many goroutines**:
   ```go
   // BAD - too many goroutines
   for i := 0; i < 1000000; i++ {
       go func() {
           // Work
       }()
   }
   
   // GOOD - appropriate number
   numWorkers := runtime.NumCPU()
   for i := 0; i < numWorkers; i++ {
       go func() {
           // Work
       }()
   }
   ```

3. **Not using buffered channels**:
   ```go
   // BAD - unbuffered channel
   ch := make(chan int)
   
   // GOOD - buffered channel
   ch := make(chan int, 100)
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a high-performance concurrent system:
- Multiple workers process jobs
- Use profiling to identify bottlenecks
- Optimize performance using different techniques
- Measure improvements with benchmarking
- Show how to optimize for different scenarios

**Hint**: Use profiling tools to identify bottlenecks and then apply appropriate optimizations.

## 🚀 Next Steps

Now that you understand performance optimization, let's learn about **common pitfalls** in the next file: `25-common-pitfalls.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
