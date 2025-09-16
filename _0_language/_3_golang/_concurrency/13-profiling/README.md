# 📊 Level 4, Topic 1: Profiling & Benchmarking

## 🚀 Overview
Mastering profiling and benchmarking is essential for building high-performance concurrent Go applications. This topic will take you from basic profiling techniques to advanced optimization strategies that will make you a performance expert.

---

## 📚 Table of Contents

1. [Profiling Fundamentals](#profiling-fundamentals)
2. [CPU Profiling](#cpu-profiling)
3. [Memory Profiling](#memory-profiling)
4. [Goroutine Profiling](#goroutine-profiling)
5. [Block Profiling](#block-profiling)
6. [Mutex Profiling](#mutex-profiling)
7. [Benchmarking](#benchmarking)
8. [Performance Optimization](#performance-optimization)
9. [Profiling Tools](#profiling-tools)
10. [Real-World Profiling](#real-world-profiling)
11. [Advanced Techniques](#advanced-techniques)
12. [Performance Monitoring](#performance-monitoring)

---

## 📊 Profiling Fundamentals

### What is Profiling?

Profiling is the process of analyzing program execution to identify performance bottlenecks, memory usage patterns, and optimization opportunities. In Go, profiling helps you:

- **Identify hot spots** in your code
- **Find memory leaks** and excessive allocations
- **Optimize goroutine usage** and concurrency patterns
- **Measure performance** improvements
- **Debug performance issues** in production

### Go Profiling Types

#### 1. CPU Profiling
- **Purpose**: Identify CPU-intensive operations
- **Tool**: `go tool pprof`
- **Use Case**: Finding performance bottlenecks

#### 2. Memory Profiling
- **Purpose**: Analyze memory usage and allocations
- **Tool**: `go tool pprof`
- **Use Case**: Finding memory leaks and optimization opportunities

#### 3. Goroutine Profiling
- **Purpose**: Analyze goroutine usage and blocking
- **Tool**: `go tool pprof`
- **Use Case**: Debugging concurrency issues

#### 4. Block Profiling
- **Purpose**: Identify blocking operations
- **Tool**: `go tool pprof`
- **Use Case**: Finding synchronization bottlenecks

#### 5. Mutex Profiling
- **Purpose**: Analyze mutex contention
- **Tool**: `go tool pprof`
- **Use Case**: Optimizing synchronization

---

## 🔥 CPU Profiling

### Basic CPU Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "time"
)

func cpuIntensiveTask() {
    for i := 0; i < 1000000; i++ {
        // CPU-intensive operation
        _ = i * i
    }
}

func main() {
    // Create CPU profile file
    f, err := os.Create("cpu.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Start CPU profiling
    if err := pprof.StartCPUProfile(f); err != nil {
        panic(err)
    }
    defer pprof.StopCPUProfile()

    // Run CPU-intensive task
    cpuIntensiveTask()
}
```

### CPU Profiling with HTTP Server

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
)

func main() {
    // Start HTTP server for profiling
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()

    // Your application code
    runApplication()
}
```

### CPU Profiling Analysis

```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof main.go

# Analyze CPU profile
go tool pprof cpu.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by CPU usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph
```

---

## 🧠 Memory Profiling

### Basic Memory Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "time"
)

func memoryIntensiveTask() {
    var data [][]int
    
    for i := 0; i < 1000; i++ {
        // Allocate memory
        slice := make([]int, 1000)
        data = append(data, slice)
    }
}

func main() {
    // Create memory profile file
    f, err := os.Create("mem.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Run memory-intensive task
    memoryIntensiveTask()

    // Write memory profile
    runtime.GC() // Force garbage collection
    if err := pprof.WriteHeapProfile(f); err != nil {
        panic(err)
    }
}
```

### Memory Profiling with HTTP Server

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
    "runtime"
)

func main() {
    // Start HTTP server for profiling
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()

    // Your application code
    runApplication()
}
```

### Memory Profiling Analysis

```bash
# Generate memory profile
go run -memprofile=mem.prof main.go

# Analyze memory profile
go tool pprof mem.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by memory usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph
```

---

## 🔄 Goroutine Profiling

### Basic Goroutine Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "sync"
    "time"
)

func goroutineTask(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        time.Sleep(1 * time.Millisecond)
    }
}

func main() {
    // Create goroutine profile file
    f, err := os.Create("goroutine.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var wg sync.WaitGroup
    
    // Start multiple goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go goroutineTask(i, &wg)
    }
    
    // Wait for completion
    wg.Wait()
    
    // Write goroutine profile
    if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
        panic(err)
    }
}
```

### Goroutine Profiling Analysis

```bash
# Generate goroutine profile
go run -goroutineprofile=goroutine.prof main.go

# Analyze goroutine profile
go tool pprof goroutine.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 goroutines
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

---

## 🚫 Block Profiling

### Basic Block Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "sync"
    "time"
)

func blockingTask(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // Blocking operation
    data := <-ch
    _ = data
}

func main() {
    // Create block profile file
    f, err := os.Create("block.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Start block profiling
    runtime.SetBlockProfileRate(1)
    defer runtime.SetBlockProfileRate(0)

    ch := make(chan int)
    var wg sync.WaitGroup
    
    // Start goroutines that will block
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go blockingTask(ch, &wg)
    }
    
    // Send data to unblock
    for i := 0; i < 5; i++ {
        ch <- i
    }
    
    wg.Wait()
    close(ch)
    
    // Write block profile
    if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
        panic(err)
    }
}
```

### Block Profiling Analysis

```bash
# Generate block profile
go run -blockprofile=block.prof main.go

# Analyze block profile
go tool pprof block.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 blocking operations
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

---

## 🔒 Mutex Profiling

### Basic Mutex Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "sync"
    "time"
)

func mutexTask(mu *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        mu.Lock()
        // Critical section
        time.Sleep(1 * time.Microsecond)
        mu.Unlock()
    }
}

func main() {
    // Create mutex profile file
    f, err := os.Create("mutex.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Start mutex profiling
    runtime.SetMutexProfileFraction(1)
    defer runtime.SetMutexProfileFraction(0)

    var mu sync.Mutex
    var wg sync.WaitGroup
    
    // Start multiple goroutines that will contend for mutex
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go mutexTask(&mu, &wg)
    }
    
    wg.Wait()
    
    // Write mutex profile
    if err := pprof.Lookup("mutex").WriteTo(f, 0); err != nil {
        panic(err)
    }
}
```

### Mutex Profiling Analysis

```bash
# Generate mutex profile
go run -mutexprofile=mutex.prof main.go

# Analyze mutex profile
go tool pprof mutex.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 mutex contentions
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

---

## ⚡ Benchmarking

### Basic Benchmarking

```go
package main

import (
    "testing"
    "time"
)

func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Function to benchmark
        time.Sleep(1 * time.Microsecond)
    }
}

func BenchmarkConcurrentFunction(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Function to benchmark
            time.Sleep(1 * time.Microsecond)
        }
    })
}
```

### Benchmarking with Setup

```go
package main

import (
    "testing"
    "sync"
)

func BenchmarkWithSetup(b *testing.B) {
    // Setup code
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    
    b.ResetTimer() // Reset timer after setup
    
    for i := 0; i < b.N; i++ {
        // Function to benchmark
        _ = data[i%len(data)]
    }
}

func BenchmarkWithCleanup(b *testing.B) {
    // Setup code
    data := make([]int, 1000)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        // Function to benchmark
        _ = data[i%len(data)]
    }
    
    b.StopTimer()
    
    // Cleanup code
    data = nil
}
```

### Benchmarking Analysis

```bash
# Run benchmarks
go test -bench=.

# Run benchmarks with profiling
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run benchmarks multiple times
go test -bench=. -count=5

# Run benchmarks with memory allocation info
go test -bench=. -benchmem
```

---

## 🚀 Performance Optimization

### 1. CPU Optimization

```go
// Before: Inefficient CPU usage
func inefficientCPU(data []int) int {
    result := 0
    for i := 0; i < len(data); i++ {
        for j := 0; j < len(data); j++ {
            result += data[i] * data[j]
        }
    }
    return result
}

// After: Optimized CPU usage
func optimizedCPU(data []int) int {
    result := 0
    sum := 0
    for _, v := range data {
        sum += v
    }
    for _, v := range data {
        result += v * sum
    }
    return result
}
```

### 2. Memory Optimization

```go
// Before: Inefficient memory usage
func inefficientMemory(data []int) []int {
    result := make([]int, 0)
    for _, v := range data {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}

// After: Optimized memory usage
func optimizedMemory(data []int) []int {
    result := make([]int, 0, len(data)) // Pre-allocate capacity
    for _, v := range data {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}
```

### 3. Goroutine Optimization

```go
// Before: Inefficient goroutine usage
func inefficientGoroutines(data []int) {
    var wg sync.WaitGroup
    for _, v := range data {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            // Process val
            _ = val * val
        }(v)
    }
    wg.Wait()
}

// After: Optimized goroutine usage
func optimizedGoroutines(data []int) {
    const numWorkers = runtime.NumCPU()
    jobs := make(chan int, len(data))
    results := make(chan int, len(data))
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func() {
            for val := range jobs {
                results <- val * val
            }
        }()
    }
    
    // Send jobs
    for _, v := range data {
        jobs <- v
    }
    close(jobs)
    
    // Collect results
    for i := 0; i < len(data); i++ {
        <-results
    }
}
```

---

## 🛠️ Profiling Tools

### 1. pprof Command Line

```bash
# CPU profiling
go tool pprof cpu.prof

# Memory profiling
go tool pprof mem.prof

# Goroutine profiling
go tool pprof goroutine.prof

# Block profiling
go tool pprof block.prof

# Mutex profiling
go tool pprof mutex.prof
```

### 2. pprof Web Interface

```bash
# Start web interface
go tool pprof -http=:8080 cpu.prof

# Open in browser
open http://localhost:8080
```

### 3. pprof Interactive Commands

```bash
# Top functions
(pprof) top10

# Show source code
(pprof) list functionName

# Show call graph
(pprof) web

# Generate PNG
(pprof) png

# Show memory allocation
(pprof) alloc_space

# Show CPU usage
(pprof) cpu
```

### 4. pprof Comparison

```bash
# Compare two profiles
go tool pprof -base cpu1.prof cpu2.prof

# Compare with diff
go tool pprof -diff_base cpu1.prof cpu2.prof
```

---

## 🌍 Real-World Profiling

### 1. Web Server Profiling

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    // Start profiling server
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Your web server
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Your handler logic
    time.Sleep(10 * time.Millisecond)
    w.Write([]byte("Hello, World!"))
}
```

### 2. Database Profiling

```go
package main

import (
    "database/sql"
    "net/http"
    _ "net/http/pprof"
    _ "github.com/lib/pq"
)

func main() {
    // Start profiling server
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Database operations
    db, err := sql.Open("postgres", "connection_string")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Your application logic
    runApplication(db)
}
```

### 3. Microservice Profiling

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    // Start profiling server
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Microservice logic
    for {
        processRequest()
        time.Sleep(100 * time.Millisecond)
    }
}

func processRequest() {
    // Process request logic
    time.Sleep(50 * time.Millisecond)
}
```

---

## 🔬 Advanced Techniques

### 1. Custom Profiling

```go
package main

import (
    "os"
    "runtime/pprof"
    "time"
)

func customProfiling() {
    // Create custom profile
    prof := pprof.NewProfile("custom")
    
    // Add samples
    for i := 0; i < 1000; i++ {
        prof.Add(1)
        time.Sleep(1 * time.Millisecond)
    }
    
    // Write profile
    f, _ := os.Create("custom.prof")
    defer f.Close()
    prof.WriteTo(f, 0)
}
```

### 2. Profiling with Context

```go
package main

import (
    "context"
    "runtime/pprof"
    "time"
)

func profilingWithContext() {
    // Create context with labels
    ctx := pprof.WithLabels(context.Background(), pprof.Labels("function", "profilingWithContext"))
    
    // Add labels
    pprof.SetGoroutineLabels(ctx)
    
    // Your code
    time.Sleep(100 * time.Millisecond)
}
```

### 3. Profiling with Metrics

```go
package main

import (
    "net/http"
    "runtime"
    "time"
)

func profilingWithMetrics() {
    // Start metrics collection
    go func() {
        for {
            // Collect metrics
            var m runtime.MemStats
            runtime.ReadMemStats(&m)
            
            // Log metrics
            log.Printf("Alloc: %d, Sys: %d, NumGC: %d", m.Alloc, m.Sys, m.NumGC)
            
            time.Sleep(1 * time.Second)
        }
    }()
    
    // Your application logic
    runApplication()
}
```

---

## 📈 Performance Monitoring

### 1. Real-time Monitoring

```go
package main

import (
    "net/http"
    "runtime"
    "time"
)

func realTimeMonitoring() {
    // Start monitoring server
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte(fmt.Sprintf("alloc_bytes %d\n", m.Alloc)))
        w.Write([]byte(fmt.Sprintf("sys_bytes %d\n", m.Sys)))
        w.Write([]byte(fmt.Sprintf("num_gc %d\n", m.NumGC)))
    })
    
    go http.ListenAndServe(":8080", nil)
    
    // Your application
    runApplication()
}
```

### 2. Performance Alerts

```go
package main

import (
    "log"
    "runtime"
    "time"
)

func performanceAlerts() {
    go func() {
        for {
            var m runtime.MemStats
            runtime.ReadMemStats(&m)
            
            // Check memory usage
            if m.Alloc > 100*1024*1024 { // 100MB
                log.Printf("High memory usage: %d bytes", m.Alloc)
            }
            
            // Check GC frequency
            if m.NumGC > 100 {
                log.Printf("High GC frequency: %d", m.NumGC)
            }
            
            time.Sleep(1 * time.Second)
        }
    }()
    
    // Your application
    runApplication()
}
```

### 3. Performance Dashboard

```go
package main

import (
    "html/template"
    "net/http"
    "runtime"
    "time"
)

func performanceDashboard() {
    http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        data := struct {
            Alloc      uint64
            Sys        uint64
            NumGC      uint32
            NumGoroutine int
        }{
            Alloc:      m.Alloc,
            Sys:        m.Sys,
            NumGC:      m.NumGC,
            NumGoroutine: runtime.NumGoroutine(),
        }
        
        tmpl := template.Must(template.New("dashboard").Parse(`
            <h1>Performance Dashboard</h1>
            <p>Memory Allocated: {{.Alloc}} bytes</p>
            <p>System Memory: {{.Sys}} bytes</p>
            <p>GC Count: {{.NumGC}}</p>
            <p>Goroutines: {{.NumGoroutine}}</p>
        `))
        
        tmpl.Execute(w, data)
    })
    
    go http.ListenAndServe(":8080", nil)
    
    // Your application
    runApplication()
}
```

---

## 🎓 Summary

Mastering profiling and benchmarking is essential for building high-performance concurrent Go applications. Key takeaways:

1. **Understand profiling types** and their use cases
2. **Use CPU profiling** to identify performance bottlenecks
3. **Use memory profiling** to find memory leaks and optimization opportunities
4. **Use goroutine profiling** to debug concurrency issues
5. **Use block and mutex profiling** to find synchronization bottlenecks
6. **Write effective benchmarks** for performance measurement
7. **Apply optimization techniques** based on profiling results
8. **Use profiling tools** effectively for analysis
9. **Implement real-world profiling** in production systems
10. **Monitor performance** continuously

Profiling and benchmarking provide the foundation for building high-performance systems! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different profiling techniques
3. **Apply** profiling to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** advanced profiling techniques

Ready to become a Profiling & Benchmarking expert? Let's dive into the implementation! 💪

