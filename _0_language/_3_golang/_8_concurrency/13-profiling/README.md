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


# Profiling & Benchmarking Commands

## Quick Reference

### Basic Commands
```bash
# Compile the code
go build .

# Run basic examples
go run . basic

# Run exercises
go run . exercises

# Run advanced patterns
go run . advanced

# Run all examples
go run . all
```

### Profiling Commands
```bash
# CPU profiling
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof

# Memory profiling
go run -memprofile=mem.prof . basic
go tool pprof mem.prof

# Goroutine profiling
go run -goroutineprofile=goroutine.prof . basic
go tool pprof goroutine.prof

# Block profiling
go run -blockprofile=block.prof . basic
go tool pprof block.prof

# Mutex profiling
go run -mutexprofile=mutex.prof . basic
go tool pprof mutex.prof
```

### Benchmark Commands
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run with memory info
go test -bench=. -benchmem

# Run multiple times
go test -bench=. -count=5
```

### Testing Commands
```bash
# Run all tests
./quick_test.sh

# Static analysis
go vet .

# Race detection
go run -race . basic

# Memory profiling
go run -memprofile=mem.prof . basic
go tool pprof mem.prof

# CPU profiling
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

## Command Details

### 1. Basic Examples (`go run . basic`)
Runs 20 basic profiling and benchmarking examples:
- Basic CPU Profiling
- Basic Memory Profiling
- Basic Goroutine Profiling
- Basic Block Profiling
- Basic Mutex Profiling
- Basic Benchmarking
- Memory Allocation Analysis
- Goroutine Analysis
- CPU Usage Analysis
- Memory Leak Detection
- Performance Comparison
- Profiling with HTTP Server
- Custom Profiling
- Memory Pool Usage
- Goroutine Pool
- Channel Performance
- Select Performance
- Mutex vs Channel Performance
- Memory Efficiency
- Performance Monitoring

### 2. Exercises (`go run . exercises`)
Runs 10 hands-on exercises:
- Implement CPU Profiling
- Implement Memory Profiling
- Implement Goroutine Profiling
- Implement Block Profiling
- Implement Mutex Profiling
- Implement Benchmarking
- Implement Memory Analysis
- Implement Performance Comparison
- Implement Memory Pool
- Implement Goroutine Pool

### 3. Advanced Patterns (`go run . advanced`)
Runs 10 advanced patterns:
- Real-time Performance Monitoring
- Profiling with Context
- Custom Profiling
- Performance Profiler
- Memory Profiler
- Goroutine Profiler
- Performance Dashboard
- Performance Optimizer
- Memory Optimizer
- Performance Tester

### 4. All Examples (`go run . all`)
Runs all examples and exercises in sequence:
- Basic examples
- Exercises
- Advanced patterns

## Profiling Commands

### CPU Profiling
```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof . basic

# Analyze CPU profile
go tool pprof cpu.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by CPU usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph

# Web interface
go tool pprof -http=:8080 cpu.prof
```

### Memory Profiling
```bash
# Generate memory profile
go run -memprofile=mem.prof . basic

# Analyze memory profile
go tool pprof mem.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by memory usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph

# Web interface
go tool pprof -http=:8080 mem.prof
```

### Goroutine Profiling
```bash
# Generate goroutine profile
go run -goroutineprofile=goroutine.prof . basic

# Analyze goroutine profile
go tool pprof goroutine.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 goroutines
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

### Block Profiling
```bash
# Generate block profile
go run -blockprofile=block.prof . basic

# Analyze block profile
go tool pprof block.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 blocking operations
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

### Mutex Profiling
```bash
# Generate mutex profile
go run -mutexprofile=mutex.prof . basic

# Analyze mutex profile
go tool pprof mutex.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 mutex contentions
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

## Benchmark Commands

### Run Benchmarks
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run with memory info
go test -bench=. -benchmem

# Run multiple times
go test -bench=. -count=5

# Run with profiling
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
```

### Benchmark Analysis
```bash
# Compare benchmarks
go test -bench=. -benchmem > before.txt
# Make changes
go test -bench=. -benchmem > after.txt
# Compare results
```

## Testing Commands

### Quick Test Script
```bash
./quick_test.sh
```
Runs all tests automatically:
1. Basic compilation
2. Static analysis
3. Basic examples
4. Exercises
5. Advanced patterns
6. All examples
7. Benchmark tests

### Manual Testing
```bash
# Test individual components
go run . basic
go run . exercises
go run . advanced

# Test with different parameters
go run . all
```

## Development Commands

### Format Code
```bash
go fmt .
```

### Check for Unused Imports
```bash
go mod tidy
```

### Run with Verbose Output
```bash
go run -v . basic
```

### Run with Timeout
```bash
timeout 30s go run . basic
```

## File Structure Commands

### List Files
```bash
ls -la
```

### View File Contents
```bash
cat main.go
cat exercises.go
cat advanced_patterns.go
```

### Check File Permissions
```bash
ls -la *.sh
```

## Module Commands

### Initialize Module
```bash
go mod init profiling
```

### Add Dependencies
```bash
go get <package>
```

### Update Dependencies
```bash
go get -u <package>
```

### Clean Module
```bash
go mod tidy
```

## Git Commands (if using version control)

### Check Status
```bash
git status
```

### Add Files
```bash
git add .
```

### Commit Changes
```bash
git commit -m "Add profiling implementation"
```

### View History
```bash
git log --oneline
```

## Troubleshooting Commands

### Check Go Version
```bash
go version
```

### Check Module Status
```bash
go mod verify
```

### Clean Build Cache
```bash
go clean -cache
```

### Check for Updates
```bash
go list -u -m all
```

## Example Usage

### Run Basic Examples
```bash
cd 13-profiling
go run . basic
```

### Run All Tests
```bash
cd 13-profiling
./quick_test.sh
```

### Run with Profiling
```bash
cd 13-profiling
go run -cpuprofile=cpu.prof . all
go tool pprof cpu.prof
```

### Debug Performance Issues
```bash
cd 13-profiling
go run -race . basic
```

## Tips

1. **Always run tests before moving to next topic**
2. **Use profiling to identify performance bottlenecks**
3. **Use benchmarks to measure performance improvements**
4. **Check static analysis for code quality**
5. **Use timeouts to prevent hanging processes**

## Common Issues

### Permission Denied
```bash
chmod +x quick_test.sh
```

### Module Not Found
```bash
go mod tidy
```

### Profile Generation Issues
```bash
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

### Memory Issues
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

## Next Steps

After running all commands successfully:
1. Review the output and understand the patterns
2. Experiment with different profiling scenarios
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects


# Profiling & Benchmarking Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Profiling & Benchmarking implementation.

## Test Structure

### 1. Basic Compilation Test
```bash
go build .
```
- **Purpose**: Verify the code compiles without errors
- **Expected**: Clean compilation with no errors

### 2. Static Analysis Test
```bash
go vet .
```
- **Purpose**: Check for common programming errors
- **Expected**: No warnings or errors

### 3. Basic Examples Test
```bash
go run . basic
```
- **Purpose**: Verify basic profiling and benchmarking functionality
- **Expected**: All 20 basic examples run successfully

### 4. Exercises Test
```bash
go run . exercises
```
- **Purpose**: Verify hands-on exercises work correctly
- **Expected**: All 10 exercises complete successfully

### 5. Advanced Patterns Test
```bash
go run . advanced
```
- **Purpose**: Verify advanced patterns implementation
- **Expected**: All 10 advanced patterns run successfully

### 6. All Examples Test
```bash
go run . all
```
- **Purpose**: Run all examples and exercises together
- **Expected**: Complete execution without errors

### 7. Benchmark Tests
```bash
go test -bench=. -benchmem
```
- **Purpose**: Run benchmark tests
- **Expected**: All benchmarks complete successfully

## Test Categories

### Basic Examples
1. **Basic CPU Profiling**: CPU profiling with file output
2. **Basic Memory Profiling**: Memory profiling with heap analysis
3. **Basic Goroutine Profiling**: Goroutine profiling and analysis
4. **Basic Block Profiling**: Block profiling for synchronization
5. **Basic Mutex Profiling**: Mutex profiling for contention analysis
6. **Basic Benchmarking**: Simple and concurrent benchmarking
7. **Memory Allocation Analysis**: Memory allocation tracking
8. **Goroutine Analysis**: Goroutine count monitoring
9. **CPU Usage Analysis**: CPU performance measurement
10. **Memory Leak Detection**: Memory leak identification
11. **Performance Comparison**: Sequential vs concurrent performance
12. **Profiling with HTTP Server**: HTTP-based profiling setup
13. **Custom Profiling**: Custom profile creation
14. **Memory Pool Usage**: Memory pool performance
15. **Goroutine Pool**: Goroutine pool implementation
16. **Channel Performance**: Channel performance comparison
17. **Select Performance**: Select statement performance
18. **Mutex vs Channel Performance**: Synchronization comparison
19. **Memory Efficiency**: Memory allocation optimization
20. **Performance Monitoring**: Real-time performance monitoring

### Exercises
1. **Implement CPU Profiling**: CPU profiling implementation
2. **Implement Memory Profiling**: Memory profiling implementation
3. **Implement Goroutine Profiling**: Goroutine profiling implementation
4. **Implement Block Profiling**: Block profiling implementation
5. **Implement Mutex Profiling**: Mutex profiling implementation
6. **Implement Benchmarking**: Benchmarking implementation
7. **Implement Memory Analysis**: Memory analysis implementation
8. **Implement Performance Comparison**: Performance comparison
9. **Implement Memory Pool**: Memory pool implementation
10. **Implement Goroutine Pool**: Goroutine pool implementation

### Advanced Patterns
1. **Real-time Performance Monitoring**: Live performance monitoring
2. **Profiling with Context**: Context-aware profiling
3. **Custom Profiling**: Custom profiling implementation
4. **Performance Profiler**: Performance measurement system
5. **Memory Profiler**: Memory allocation tracking
6. **Goroutine Profiler**: Goroutine monitoring
7. **Performance Dashboard**: Web-based performance dashboard
8. **Performance Optimizer**: Automatic performance optimization
9. **Memory Optimizer**: Memory usage optimization
10. **Performance Tester**: Performance testing framework

## Performance Expectations

### Basic Examples
- **Execution Time**: < 20 seconds
- **Memory Usage**: < 300MB
- **CPU Usage**: < 50% average
- **Profile Generation**: All profiles created successfully

### Exercises
- **Execution Time**: < 25 seconds
- **Memory Usage**: < 350MB
- **Success Rate**: 100% completion
- **Profile Quality**: All profiles valid and readable

### Advanced Patterns
- **Execution Time**: < 30 seconds
- **Memory Usage**: < 400MB
- **Performance**: Efficient profiling
- **Scalability**: Handles high load

## Profiling Commands

### CPU Profiling
```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof . basic

# Analyze CPU profile
go tool pprof cpu.prof

# Web interface
go tool pprof -http=:8080 cpu.prof
```

### Memory Profiling
```bash
# Generate memory profile
go run -memprofile=mem.prof . basic

# Analyze memory profile
go tool pprof mem.prof

# Web interface
go tool pprof -http=:8080 mem.prof
```

### Goroutine Profiling
```bash
# Generate goroutine profile
go run -goroutineprofile=goroutine.prof . basic

# Analyze goroutine profile
go tool pprof goroutine.prof
```

### Block Profiling
```bash
# Generate block profile
go run -blockprofile=block.prof . basic

# Analyze block profile
go tool pprof block.prof
```

### Mutex Profiling
```bash
# Generate mutex profile
go run -mutexprofile=mutex.prof . basic

# Analyze mutex profile
go tool pprof mutex.prof
```

## Benchmark Commands

### Run Benchmarks
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run with memory info
go test -bench=. -benchmem

# Run multiple times
go test -bench=. -count=5
```

### Benchmark Analysis
```bash
# Compare benchmarks
go test -bench=. -benchmem > before.txt
# Make changes
go test -bench=. -benchmem > after.txt
# Compare results
```

## Troubleshooting

### Common Issues

#### 1. Compilation Errors
- **Symptom**: `go build .` fails
- **Cause**: Syntax errors or missing imports
- **Solution**: Check code syntax and import statements

#### 2. Profile Generation Errors
- **Symptom**: Profile files not created
- **Cause**: File permission issues or disk space
- **Solution**: Check file permissions and disk space

#### 3. Memory Issues
- **Symptom**: Out of memory errors
- **Cause**: Excessive memory allocation
- **Solution**: Reduce memory usage in examples

#### 4. Performance Issues
- **Symptom**: Slow execution
- **Cause**: Inefficient algorithms
- **Solution**: Optimize code or reduce workload

#### 5. Profile Analysis Issues
- **Symptom**: pprof commands fail
- **Cause**: Invalid profile files
- **Solution**: Regenerate profiles

### Debug Commands

#### Verbose Output
```bash
go run -v . basic
```

#### Race Detection
```bash
go run -race . basic
```

#### Memory Profiling
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

#### CPU Profiling
```bash
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

## Test Automation

### Quick Test Script
```bash
./quick_test.sh
```
- Runs all tests automatically
- Provides clear pass/fail status
- Includes comprehensive testing

### Manual Testing
```bash
# Test individual components
go run . basic
go run . exercises
go run . advanced

# Test with different parameters
go run . all
```

## Success Criteria

### ✅ All Tests Pass
- Compilation successful
- Static analysis clean
- All examples run without errors
- Performance within expected ranges

### ✅ Code Quality
- Clean, readable code
- Proper error handling
- Good documentation
- Efficient profiling
- Proper resource management

### ✅ Learning Objectives
- Understand profiling types
- Implement profiling techniques
- Use profiling tools effectively
- Apply optimization strategies
- Master benchmarking techniques

## Next Steps

After successful testing:
1. Review the code and understand the patterns
2. Experiment with different profiling scenarios
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects

## Support

If you encounter issues:
1. Check the error messages carefully
2. Review the code for common mistakes
3. Use the debug commands provided
4. Refer to the Go profiling documentation
5. Ask for help in the learning community

