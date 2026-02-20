# 🚀 Pipeline Pattern: Efficient Data Processing Through Stages

## 📚 Table of Contents
1. [What is the Pipeline Pattern?](#what-is-the-pipeline-pattern)
2. [Basic Pipeline](#basic-pipeline)
3. [Buffered Pipeline](#buffered-pipeline)
4. [Fan-Out/Fan-In Pipeline](#fan-outfan-in-pipeline)
5. [Pipeline with Error Handling](#pipeline-with-error-handling)
6. [Pipeline with Timeout](#pipeline-with-timeout)
7. [Pipeline with Rate Limiting](#pipeline-with-rate-limiting)
8. [Pipeline with Metrics](#pipeline-with-metrics)
9. [Pipeline with Backpressure](#pipeline-with-backpressure)
10. [Pipeline with Circuit Breaker](#pipeline-with-circuit-breaker)
11. [Pipeline with Load Balancing](#pipeline-with-load-balancing)
12. [Pipeline with Caching](#pipeline-with-caching)
13. [Performance Considerations](#performance-considerations)
14. [Common Patterns](#common-patterns)
15. [Best Practices](#best-practices)
16. [Common Pitfalls](#common-pitfalls)
17. [Exercises](#exercises)

---

## 🎯 What is the Pipeline Pattern?

The **Pipeline Pattern** is a concurrency design pattern that processes data through multiple stages, where each stage performs a specific transformation on the data. It's like an assembly line where data flows from one stage to the next.

### Key Characteristics:
- **Sequential stages**: Data flows through stages in order
- **Concurrent processing**: Each stage can process multiple items simultaneously
- **Streaming**: Data flows continuously through the pipeline
- **Transformations**: Each stage transforms data before passing to the next
- **Scalability**: Each stage can be scaled independently

### Benefits:
- **Parallel processing**: Multiple items can be processed simultaneously
- **Modularity**: Each stage is independent and can be modified separately
- **Scalability**: Each stage can be scaled independently
- **Efficiency**: Optimal resource utilization
- **Maintainability**: Clear separation of concerns

---

## 🏗️ Basic Pipeline

A basic pipeline consists of:
1. **Input stage**: Receives data from source
2. **Processing stages**: Transform data
3. **Output stage**: Delivers final results

### Basic Structure:
```go
func basicPipeline(input <-chan Data, output chan<- Result) {
    // Stage 1: Process input
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data)
        }
    }()
    
    // Stage 2: Process stage1 output
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            stage2 <- processStage2(data)
        }
    }()
    
    // Stage 3: Final processing
    go func() {
        defer close(output)
        for data := range stage2 {
            output <- processStage3(data)
        }
    }()
}
```

### Key Components:
- **Channels**: Connect stages together
- **Goroutines**: Each stage runs in its own goroutine
- **Buffering**: Channels can be buffered for better performance
- **Closing**: Proper channel closing for cleanup

---

## 📦 Buffered Pipeline

A buffered pipeline uses buffered channels to improve performance and handle bursts of data.

### Benefits:
- **Better throughput**: Buffers smooth out data flow
- **Reduced blocking**: Stages don't block on channel operations
- **Burst handling**: Can handle sudden spikes in data
- **Memory control**: Buffer size limits memory usage

### Implementation:
```go
func bufferedPipeline(input <-chan Data, output chan<- Result, bufferSize int) {
    // Stage 1 with buffering
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data)
        }
    }()
    
    // Stage 2 with buffering
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            stage2 <- processStage2(data)
        }
    }()
    
    // Final stage
    go func() {
        defer close(output)
        for data := range stage2 {
            output <- processStage3(data)
        }
    }()
}
```

---

## 🌊 Fan-Out/Fan-In Pipeline

A fan-out/fan-in pipeline distributes work to multiple workers and then collects results.

### Features:
- **Fan-out**: Distributes work to multiple workers
- **Fan-in**: Collects results from multiple workers
- **Load balancing**: Work is distributed evenly
- **Scalability**: Can scale each stage independently

### Implementation:
```go
func fanOutFanInPipeline(input <-chan Data, output chan<- Result, numWorkers int) {
    // Fan-out: Distribute work to multiple workers
    workers := make([]chan Data, numWorkers)
    for i := range workers {
        workers[i] = make(chan Data)
    }
    
    // Start workers
    var wg sync.WaitGroup
    for i, worker := range workers {
        wg.Add(1)
        go func(id int, jobs <-chan Data) {
            defer wg.Done()
            for data := range jobs {
                result := processData(data)
                output <- result
            }
        }(i, worker)
    }
    
    // Distribute work
    go func() {
        defer func() {
            for _, worker := range workers {
                close(worker)
            }
        }()
        
        for data := range input {
            // Round-robin distribution
            workers[data.ID%numWorkers] <- data
        }
    }()
    
    wg.Wait()
    close(output)
}
```

---

## ⚠️ Pipeline with Error Handling

A pipeline that properly handles errors from any stage.

### Features:
- **Error collection**: Collects errors from all stages
- **Error recovery**: Can recover from stage failures
- **Error reporting**: Reports errors to monitoring systems
- **Graceful degradation**: Continues processing despite errors

### Implementation:
```go
func pipelineWithErrorHandling(input <-chan Data, output chan<- Result, errors chan<- error) {
    // Stage 1 with error handling
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            processed, err := processStage1WithError(data)
            if err != nil {
                errors <- err
            } else {
                stage1 <- processed
            }
        }
    }()
    
    // Stage 2 with error handling
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            processed, err := processStage2WithError(data)
            if err != nil {
                errors <- err
            } else {
                stage2 <- processed
            }
        }
    }()
    
    // Final stage with error handling
    go func() {
        defer close(output)
        for data := range stage2 {
            result, err := processStage3WithError(data)
            if err != nil {
                errors <- err
            } else {
                output <- result
            }
        }
    }()
}
```

---

## ⏰ Pipeline with Timeout

A pipeline that handles timeouts for individual stages and overall processing.

### Features:
- **Stage timeouts**: Individual stages can timeout
- **Overall timeout**: Entire pipeline can timeout
- **Timeout handling**: Graceful handling of timeouts
- **Resource cleanup**: Cleans up resources on timeout

### Implementation:
```go
func pipelineWithTimeout(input <-chan Data, output chan<- Result, timeout time.Duration) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    // Stage 1 with timeout
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for {
            select {
            case data, ok := <-input:
                if !ok {
                    return
                }
                processed := processStage1WithContext(ctx, data)
                stage1 <- processed
            case <-ctx.Done():
                return
            }
        }
    }()
    
    // Stage 2 with timeout
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for {
            select {
            case data, ok := <-stage1:
                if !ok {
                    return
                }
                processed := processStage2WithContext(ctx, data)
                stage2 <- processed
            case <-ctx.Done():
                return
            }
        }
    }()
    
    // Final stage with timeout
    go func() {
        defer close(output)
        for {
            select {
            case data, ok := <-stage2:
                if !ok {
                    return
                }
                result := processStage3WithContext(ctx, data)
                output <- result
            case <-ctx.Done():
                return
            }
        }
    }()
}
```

---

## 🚦 Pipeline with Rate Limiting

A pipeline that limits the rate of data processing.

### Features:
- **Rate limiting**: Controls the rate of data processing
- **Burst handling**: Allows bursts up to a limit
- **Adaptive limiting**: Can adjust rate based on conditions
- **Backpressure**: Slows down when rate limit is exceeded

### Implementation:
```go
func pipelineWithRateLimiting(input <-chan Data, output chan<- Result, rateLimit int) {
    // Rate limiter
    rateLimiter := time.NewTicker(time.Second / time.Duration(rateLimit))
    defer rateLimiter.Stop()
    
    // Stage 1 with rate limiting
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            <-rateLimiter.C // Wait for rate limiter
            stage1 <- processStage1(data)
        }
    }()
    
    // Stage 2 with rate limiting
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            <-rateLimiter.C // Wait for rate limiter
            stage2 <- processStage2(data)
        }
    }()
    
    // Final stage with rate limiting
    go func() {
        defer close(output)
        for data := range stage2 {
            <-rateLimiter.C // Wait for rate limiter
            output <- processStage3(data)
        }
    }()
}
```

---

## 📈 Pipeline with Metrics

A pipeline that collects and reports metrics about its performance.

### Features:
- **Performance metrics**: Tracks processing times, throughput
- **Stage metrics**: Monitors each stage individually
- **Error metrics**: Tracks error rates and types
- **Real-time monitoring**: Provides real-time metrics

### Implementation:
```go
func pipelineWithMetrics(input <-chan Data, output chan<- Result) {
    metrics := &PipelineMetrics{}
    
    // Stage 1 with metrics
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            start := time.Now()
            processed := processStage1(data)
            metrics.recordStage("stage1", time.Since(start))
            stage1 <- processed
        }
    }()
    
    // Stage 2 with metrics
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            start := time.Now()
            processed := processStage2(data)
            metrics.recordStage("stage2", time.Since(start))
            stage2 <- processed
        }
    }()
    
    // Final stage with metrics
    go func() {
        defer close(output)
        for data := range stage2 {
            start := time.Now()
            result := processStage3(data)
            metrics.recordStage("stage3", time.Since(start))
            output <- result
        }
    }()
}
```

---

## 🔄 Pipeline with Backpressure

A pipeline that handles backpressure when downstream stages are slower.

### Features:
- **Backpressure detection**: Detects when downstream stages are slow
- **Flow control**: Adjusts input rate based on downstream capacity
- **Memory protection**: Prevents memory overflow
- **Adaptive processing**: Adjusts processing rate dynamically

### Implementation:
```go
func pipelineWithBackpressure(input <-chan Data, output chan<- Result) {
    // Stage 1 with backpressure
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            select {
            case stage1 <- processStage1(data):
                // Successfully sent
            case <-time.After(backpressureTimeout):
                // Backpressure detected, skip or retry
                continue
            }
        }
    }()
    
    // Stage 2 with backpressure
    stage2 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage2)
        for data := range stage1 {
            select {
            case stage2 <- processStage2(data):
                // Successfully sent
            case <-time.After(backpressureTimeout):
                // Backpressure detected, skip or retry
                continue
            }
        }
    }()
    
    // Final stage with backpressure
    go func() {
        defer close(output)
        for data := range stage2 {
            select {
            case output <- processStage3(data):
                // Successfully sent
            case <-time.After(backpressureTimeout):
                // Backpressure detected, skip or retry
                continue
            }
        }
    }()
}
```

---

## 🔌 Pipeline with Circuit Breaker

A pipeline that uses circuit breakers to handle failures gracefully.

### Features:
- **Circuit breaker**: Prevents cascading failures
- **Failure detection**: Detects when stages are failing
- **Recovery**: Automatically recovers when conditions improve
- **Fallback**: Provides fallback behavior when circuit is open

### Implementation:
```go
func pipelineWithCircuitBreaker(input <-chan Data, output chan<- Result) {
    // Circuit breakers for each stage
    breaker1 := NewCircuitBreaker(threshold, timeout)
    breaker2 := NewCircuitBreaker(threshold, timeout)
    breaker3 := NewCircuitBreaker(threshold, timeout)
    
    // Stage 1 with circuit breaker
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            if breaker1.Allow() {
                processed, err := processStage1WithError(data)
                if err != nil {
                    breaker1.RecordFailure()
                } else {
                    breaker1.RecordSuccess()
                    stage1 <- processed
                }
            } else {
                // Circuit breaker is open, skip or use fallback
                continue
            }
        }
    }()
    
    // Similar implementation for other stages...
}
```

---

## ⚖️ Pipeline with Load Balancing

A pipeline that balances load across multiple workers at each stage.

### Features:
- **Load balancing**: Distributes work evenly across workers
- **Worker health**: Monitors worker health and performance
- **Dynamic scaling**: Adjusts number of workers based on load
- **Fault tolerance**: Handles worker failures gracefully

### Implementation:
```go
func pipelineWithLoadBalancing(input <-chan Data, output chan<- Result, numWorkers int) {
    // Load balancer for each stage
    balancer1 := NewLoadBalancer(numWorkers)
    balancer2 := NewLoadBalancer(numWorkers)
    balancer3 := NewLoadBalancer(numWorkers)
    
    // Stage 1 with load balancing
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            worker := balancer1.GetWorker()
            processed := worker.Process(data)
            stage1 <- processed
        }
    }()
    
    // Similar implementation for other stages...
}
```

---

## 💾 Pipeline with Caching

A pipeline that caches results to improve performance.

### Features:
- **Result caching**: Caches results from expensive operations
- **Cache invalidation**: Invalidates cache when data changes
- **Cache warming**: Pre-loads cache with frequently used data
- **Cache metrics**: Monitors cache hit/miss rates

### Implementation:
```go
func pipelineWithCaching(input <-chan Data, output chan<- Result) {
    cache := NewCache()
    
    // Stage 1 with caching
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            // Check cache first
            if cached, found := cache.Get(data.Key); found {
                stage1 <- cached
            } else {
                processed := processStage1(data)
                cache.Set(data.Key, processed)
                stage1 <- processed
            }
        }
    }()
    
    // Similar implementation for other stages...
}
```

---

## 📊 Performance Considerations

### 1. **Buffer Sizing**
- **Small buffers**: Better for memory usage, can cause blocking
- **Large buffers**: Better for throughput, uses more memory
- **Dynamic sizing**: Adjust based on workload characteristics

### 2. **Stage Scaling**
- **CPU-bound stages**: Scale based on CPU cores
- **I/O-bound stages**: Scale based on I/O capacity
- **Mixed stages**: Balance between CPU and I/O requirements

### 3. **Memory Management**
- **Object pooling**: Reuse objects to reduce GC pressure
- **Channel management**: Close channels properly to prevent leaks
- **Resource cleanup**: Ensure all resources are cleaned up

### 4. **Load Balancing**
- **Round-robin**: Distribute work evenly
- **Weighted**: Distribute based on worker capacity
- **Adaptive**: Adjust based on worker performance

---

## 🎨 Common Patterns

### 1. **Producer-Consumer Pattern**
```go
func producerConsumerPipeline(input <-chan Data, output chan<- Result) {
    // Producer stage
    producer := make(chan Data, bufferSize)
    go func() {
        defer close(producer)
        for data := range input {
            producer <- data
        }
    }()
    
    // Consumer stage
    go func() {
        defer close(output)
        for data := range producer {
            result := processData(data)
            output <- result
        }
    }()
}
```

### 2. **Map-Reduce Pattern**
```go
func mapReducePipeline(input <-chan Data, output chan<- Result) {
    // Map stage
    mapped := make(chan MappedData, bufferSize)
    go func() {
        defer close(mapped)
        for data := range input {
            mapped <- mapData(data)
        }
    }()
    
    // Reduce stage
    go func() {
        defer close(output)
        for data := range mapped {
            result := reduceData(data)
            output <- result
        }
    }()
}
```

### 3. **Filter-Transform Pattern**
```go
func filterTransformPipeline(input <-chan Data, output chan<- Result) {
    // Filter stage
    filtered := make(chan Data, bufferSize)
    go func() {
        defer close(filtered)
        for data := range input {
            if shouldProcess(data) {
                filtered <- data
            }
        }
    }()
    
    // Transform stage
    go func() {
        defer close(output)
        for data := range filtered {
            result := transformData(data)
            output <- result
        }
    }()
}
```

---

## ✅ Best Practices

### 1. **Proper Resource Management**
```go
// ✅ Good - proper cleanup
func pipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data)
        }
    }()
    
    go func() {
        defer close(output)
        for data := range stage1 {
            output <- processStage2(data)
        }
    }()
}

// ❌ Bad - no cleanup
func badPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        for data := range input {
            stage1 <- processStage1(data)
        }
        // Forgot to close stage1
    }()
    
    go func() {
        for data := range stage1 {
            output <- processStage2(data)
        }
        // Forgot to close output
    }()
}
```

### 2. **Error Handling**
```go
// ✅ Good - proper error handling
func pipelineWithErrors(input <-chan Data, output chan<- Result, errors chan<- error) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            processed, err := processStage1WithError(data)
            if err != nil {
                errors <- err
            } else {
                stage1 <- processed
            }
        }
    }()
    
    go func() {
        defer close(output)
        for data := range stage1 {
            result, err := processStage2WithError(data)
            if err != nil {
                errors <- err
            } else {
                output <- result
            }
        }
    }()
}
```

### 3. **Context Usage**
```go
// ✅ Good - using context for cancellation
func pipelineWithContext(ctx context.Context, input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for {
            select {
            case data, ok := <-input:
                if !ok {
                    return
                }
                stage1 <- processStage1WithContext(ctx, data)
            case <-ctx.Done():
                return
            }
        }
    }()
    
    go func() {
        defer close(output)
        for {
            select {
            case data, ok := <-stage1:
                if !ok {
                    return
                }
                output <- processStage2WithContext(ctx, data)
            case <-ctx.Done():
                return
            }
        }
    }()
}
```

### 4. **Metrics and Monitoring**
```go
// ✅ Good - collecting metrics
type PipelineMetrics struct {
    Stage1Time time.Duration
    Stage2Time time.Duration
    Stage3Time time.Duration
    mu         sync.RWMutex
}

func (pm *PipelineMetrics) recordStage(stage string, duration time.Duration) {
    pm.mu.Lock()
    defer pm.mu.Unlock()
    
    switch stage {
    case "stage1":
        pm.Stage1Time += duration
    case "stage2":
        pm.Stage2Time += duration
    case "stage3":
        pm.Stage3Time += duration
    }
}
```

---

## ⚠️ Common Pitfalls

### 1. **Channel Leaks**
```go
// ❌ Wrong - channel leak
func badPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        for data := range input {
            stage1 <- processStage1(data)
        }
        // Forgot to close stage1
    }()
    
    go func() {
        for data := range stage1 {
            output <- processStage2(data)
        }
        // Forgot to close output
    }()
}

// ✅ Correct - proper cleanup
func goodPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data)
        }
    }()
    
    go func() {
        defer close(output)
        for data := range stage1 {
            output <- processStage2(data)
        }
    }()
}
```

### 2. **Deadlocks**
```go
// ❌ Wrong - potential deadlock
func badPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData) // Unbuffered
    go func() {
        for data := range input {
            stage1 <- processStage1(data) // Can block
        }
    }()
    
    go func() {
        for data := range stage1 {
            output <- processStage2(data) // Can block
        }
    }()
}

// ✅ Correct - buffered channels
func goodPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize) // Buffered
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data)
        }
    }()
    
    go func() {
        defer close(output)
        for data := range stage1 {
            output <- processStage2(data)
        }
    }()
}
```

### 3. **Incorrect Buffer Sizing**
```go
// ❌ Wrong - too small buffers
func badPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, 1) // Too small
    // ...
}

// ✅ Correct - appropriate buffer size
func goodPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize) // Appropriate size
    // ...
}
```

### 4. **Missing Error Handling**
```go
// ❌ Wrong - no error handling
func badPipeline(input <-chan Data, output chan<- Result) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            stage1 <- processStage1(data) // Can panic
        }
    }()
}

// ✅ Correct - proper error handling
func goodPipeline(input <-chan Data, output chan<- Result, errors chan<- error) {
    stage1 := make(chan ProcessedData, bufferSize)
    go func() {
        defer close(stage1)
        for data := range input {
            func() {
                defer func() {
                    if r := recover(); r != nil {
                        errors <- fmt.Errorf("stage1 panicked: %v", r)
                    }
                }()
                
                processed, err := processStage1WithError(data)
                if err != nil {
                    errors <- err
                } else {
                    stage1 <- processed
                }
            }()
        }
    }()
}
```

---

## 🧪 Exercises

### Exercise 1: Basic Pipeline
Create a basic pipeline with three stages.

### Exercise 2: Buffered Pipeline
Implement a pipeline with buffered channels.

### Exercise 3: Fan-Out/Fan-In Pipeline
Create a pipeline that distributes work to multiple workers.

### Exercise 4: Pipeline with Error Handling
Implement a pipeline that handles errors from any stage.

### Exercise 5: Pipeline with Timeout
Create a pipeline that handles timeouts for individual stages.

### Exercise 6: Pipeline with Rate Limiting
Implement a pipeline that limits the rate of data processing.

### Exercise 7: Pipeline with Metrics
Create a pipeline that collects and reports performance metrics.

### Exercise 8: Pipeline with Backpressure
Implement a pipeline that handles backpressure.

### Exercise 9: Pipeline with Circuit Breaker
Create a pipeline that uses circuit breakers.

### Exercise 10: Pipeline with Caching
Implement a pipeline that caches results.

---

## 🎯 Key Takeaways

1. **Use pipelines for data processing** - efficient and modular
2. **Choose appropriate buffer sizes** - balance memory and performance
3. **Handle errors properly** - collect and process errors from all stages
4. **Use context for cancellation** - graceful shutdown and timeout handling
5. **Collect metrics** - monitor performance and resource usage
6. **Avoid common pitfalls** - channel leaks, deadlocks, incorrect buffer sizing
7. **Implement proper cleanup** - close channels and wait for completion
8. **Consider different patterns** - fan-out/fan-in, map-reduce, filter-transform

---

## 🚀 Next Steps

Ready for the next topic? Let's move on to **Fan-Out/Fan-In Pattern** where you'll learn how to distribute work efficiently!

**Run the examples in this directory to see pipelines in action!**


# Pipeline Pattern Commands Reference

## Quick Reference

### Run All Tests
```bash
./quick_test.sh
```

### Individual Commands

#### Basic Examples
```bash
go run . basic
```

#### Exercises
```bash
go run . exercises
```

#### Advanced Patterns
```bash
go run . advanced
```

#### Compilation
```bash
go build .
```

#### Race Detection
```bash
go run -race . basic
```

#### Static Analysis
```bash
go vet .
```

#### Performance Testing
```bash
go test -bench=. -benchmem
```

## Detailed Commands

### 1. Basic Examples
```bash
# Run all basic examples
go run . basic

# Run specific example (if implemented)
go run . basic --example=1
go run . basic --example=2
```

### 2. Exercises
```bash
# Run all exercises
go run . exercises

# Run specific exercise (if implemented)
go run . exercises --exercise=1
go run . exercises --exercise=2
```

### 3. Advanced Patterns
```bash
# Run all advanced patterns
go run . advanced

# Run specific pattern (if implemented)
go run . advanced --pattern=adaptive
go run . advanced --pattern=circuit-breaker
```

### 4. Testing and Analysis

#### Compilation Test
```bash
# Basic compilation
go build .

# Cross-platform compilation
go build -o pipeline-linux GOOS=linux GOARCH=amd64 .
go build -o pipeline-windows GOOS=windows GOARCH=amd64 .
```

#### Race Detection
```bash
# Basic race detection
go run -race . basic

# Race detection with specific flags
go run -race -race=1 . basic
```

#### Static Analysis
```bash
# Basic static analysis
go vet .

# Static analysis with specific packages
go vet ./...

# Static analysis with additional checks
go vet -composites=false .
```

#### Performance Testing
```bash
# Basic benchmark
go test -bench=.

# Benchmark with memory profiling
go test -bench=. -benchmem

# Benchmark specific functions
go test -bench=BenchmarkPipeline -benchmem

# Benchmark with CPU profiling
go test -bench=. -cpuprofile=cpu.prof
```

### 5. Code Quality

#### Formatting
```bash
# Format code
go fmt .

# Format with specific options
go fmt -w .
```

#### Linting
```bash
# Install golangci-lint (if not installed)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Run specific linters
golangci-lint run --enable=gofmt,goimports,vet
```

#### Documentation
```bash
# Generate documentation
go doc .

# Generate documentation for specific function
go doc PipelinePattern

# Generate documentation in HTML
godoc -http=:6060
```

### 6. Debugging

#### Debug Build
```bash
# Build with debug information
go build -gcflags="all=-N -l" .

# Run with debug information
go run -gcflags="all=-N -l" . basic
```

#### Profiling
```bash
# CPU profiling
go run . basic -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling
go run . basic -memprofile=mem.prof
go tool pprof mem.prof

# Goroutine profiling
go run . basic -blockprofile=block.prof
go tool pprof block.prof
```

#### Trace Analysis
```bash
# Generate trace
go run . basic -trace=trace.out
go tool trace trace.out
```

### 7. Module Management

#### Dependencies
```bash
# Initialize module
go mod init pipeline-pattern

# Add dependencies
go get github.com/example/package

# Update dependencies
go get -u ./...

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

#### Version Management
```bash
# Check Go version
go version

# Check module versions
go list -m all

# Check for updates
go list -m -u all
```

### 8. Build and Deployment

#### Build Options
```bash
# Basic build
go build .

# Build with optimizations
go build -ldflags="-s -w" .

# Build with version information
go build -ldflags="-X main.Version=1.0.0" .

# Build for different architectures
go build -o pipeline-amd64 .
go build -o pipeline-arm64 .
```

#### Cross-Compilation
```bash
# Linux
GOOS=linux GOARCH=amd64 go build .

# Windows
GOOS=windows GOARCH=amd64 go build .

# macOS
GOOS=darwin GOARCH=amd64 go build .
```

### 9. Testing Specific Scenarios

#### Error Handling
```bash
# Test error handling
go run . basic --test-errors

# Test with specific error rates
go run . basic --error-rate=0.5
```

#### Performance Testing
```bash
# Test with different loads
go run . basic --load=100
go run . basic --load=1000

# Test with different buffer sizes
go run . basic --buffer-size=10
go run . basic --buffer-size=100
```

#### Stress Testing
```bash
# Run stress test
go run . basic --stress-test

# Run with timeout
timeout 30s go run . basic --stress-test
```

### 10. Monitoring and Metrics

#### Runtime Metrics
```bash
# Run with metrics
go run . basic --metrics

# Run with specific metrics
go run . basic --metrics=cpu,memory,goroutines
```

#### Logging
```bash
# Run with debug logging
go run . basic --debug

# Run with specific log level
go run . basic --log-level=info
go run . basic --log-level=debug
```

## Environment Variables

### Go-specific
```bash
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
```

### Pipeline-specific
```bash
export PIPELINE_BUFFER_SIZE=100
export PIPELINE_WORKER_COUNT=4
export PIPELINE_TIMEOUT=30s
```

## Common Issues and Solutions

### 1. Permission Denied
```bash
chmod +x quick_test.sh
```

### 2. Module Not Found
```bash
go mod tidy
go mod download
```

### 3. Race Detection Issues
```bash
# For educational race conditions, this is expected
go run -race . basic 2>&1 | grep -v "WARNING: DATA RACE"
```

### 4. Build Failures
```bash
# Clean and rebuild
go clean
go build .
```

### 5. Test Failures
```bash
# Run with verbose output
go run . basic -v

# Run with specific flags
go run . basic --help
```

## Tips and Best Practices

1. **Always run tests before committing**
2. **Use race detection in development**
3. **Profile performance-critical code**
4. **Keep dependencies up to date**
5. **Use proper error handling**
6. **Document complex patterns**
7. **Test with different loads**
8. **Monitor resource usage**
9. **Use proper logging**
10. **Follow Go conventions**

## Next Steps

After mastering these commands:
1. Experiment with different pipeline configurations
2. Try implementing your own patterns
3. Move on to the next topic: Fan-Out/Fan-In Pattern
4. Explore advanced Go concurrency features

# Pipeline Pattern Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Pipeline Pattern topic, covering basic examples, exercises, and advanced patterns.

## Test Structure

### 1. Basic Examples (`go run . basic`)
Tests fundamental pipeline patterns:
- **Basic Pipeline**: Simple 3-stage pipeline
- **Buffered Pipeline**: Pipeline with buffered channels
- **Fan-Out/Fan-In**: Parallel processing with worker distribution
- **Error Handling**: Pipeline with error propagation
- **Timeout**: Pipeline with timeout controls
- **Rate Limiting**: Pipeline with rate limiting
- **Metrics**: Pipeline with performance monitoring
- **Backpressure**: Pipeline with backpressure handling
- **Performance Comparison**: Sequential vs pipeline performance
- **Common Pitfalls**: Educational examples of mistakes

### 2. Exercises (`go run . exercises`)
Hands-on exercises covering:
- **Exercise 1**: Basic Pipeline Implementation
- **Exercise 2**: Buffered Pipeline with Custom Buffer Sizes
- **Exercise 3**: Fan-Out/Fan-In with Dynamic Workers
- **Exercise 4**: Error Handling with Recovery
- **Exercise 5**: Timeout and Cancellation
- **Exercise 6**: Rate Limiting with Different Strategies
- **Exercise 7**: Metrics Collection and Analysis
- **Exercise 8**: Backpressure with Adaptive Buffering
- **Exercise 9**: Circuit Breaker Integration
- **Exercise 10**: Caching Layer Integration

### 3. Advanced Patterns (`go run . advanced`)
Advanced pipeline implementations:
- **Adaptive Pipeline**: Self-adjusting pipeline based on load
- **Pipeline with Circuit Breaker**: Fault-tolerant pipeline
- **Pipeline with Caching**: Performance-optimized pipeline
- **Pipeline with Load Balancing**: Distributed processing
- **Pipeline with Monitoring**: Real-time monitoring and alerting

## Testing Commands

### Quick Test Suite
```bash
./quick_test.sh
```
Runs all tests including:
- Basic examples
- Exercises
- Advanced patterns
- Compilation
- Race detection (with educational race conditions)
- Static analysis

### Individual Test Commands

#### Basic Examples
```bash
go run . basic
```

#### Exercises
```bash
go run . exercises
```

#### Advanced Patterns
```bash
go run . advanced
```

#### Compilation Test
```bash
go build .
```

#### Race Detection
```bash
go run -race . basic
```
**Note**: The error handling example intentionally contains race conditions for educational purposes.

#### Static Analysis
```bash
go vet .
```

#### Performance Testing
```bash
go test -bench=. -benchmem
```

## Expected Outputs

### Basic Examples
- Pipeline processing results with stage information
- Performance metrics and timing
- Error handling demonstrations
- Rate limiting behavior
- Backpressure handling

### Exercises
- Successful completion of all 10 exercises
- Proper error handling and recovery
- Performance improvements over basic implementations
- Correct implementation of advanced patterns

### Advanced Patterns
- Adaptive behavior under different loads
- Circuit breaker state changes
- Cache hit/miss ratios
- Load balancing distribution
- Monitoring metrics

## Common Issues and Solutions

### 1. Deadlocks
**Symptom**: Program hangs indefinitely
**Solution**: Ensure all channels are properly closed and buffered appropriately

### 2. Race Conditions
**Symptom**: `go run -race` reports data races
**Solution**: Use proper synchronization primitives or make variables local to goroutines

### 3. Channel Leaks
**Symptom**: Goroutines not terminating
**Solution**: Always close channels and use `defer close(channel)`

### 4. Buffer Sizing
**Symptom**: Poor performance or blocking
**Solution**: Size buffers based on expected throughput and processing time

### 5. Error Handling
**Symptom**: Panics or unhandled errors
**Solution**: Implement proper error propagation and recovery mechanisms

## Performance Expectations

### Basic Pipeline
- Should process 10 items in ~1.5 seconds
- 3-stage pipeline with 50ms per stage

### Buffered Pipeline
- Should show improved throughput over unbuffered
- Buffer size should match processing capacity

### Fan-Out/Fan-In
- Should utilize multiple workers effectively
- Load should be distributed evenly

### Error Handling
- Should gracefully handle errors without crashing
- Should continue processing successful items

### Rate Limiting
- Should respect rate limits (e.g., 2 items per second)
- Should show controlled processing speed

## Debugging Tips

### 1. Add Logging
```go
fmt.Printf("Stage %s processing item %d\n", stageName, item.ID)
```

### 2. Use Timeouts
```go
select {
case result := <-output:
    // Process result
case <-time.After(5 * time.Second):
    fmt.Println("Timeout waiting for result")
}
```

### 3. Monitor Goroutines
```go
fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
```

### 4. Check Channel States
```go
select {
case data := <-channel:
    // Process data
default:
    fmt.Println("Channel is empty or closed")
}
```

## Success Criteria

✅ All basic examples run without errors
✅ All exercises complete successfully
✅ All advanced patterns demonstrate expected behavior
✅ Code compiles without warnings
✅ Static analysis passes
✅ Race detection shows only educational race conditions
✅ Performance meets expected benchmarks
✅ Error handling works correctly
✅ Documentation is complete and accurate

## Next Steps

After completing all tests successfully:
1. Review the code to understand the patterns
2. Experiment with different configurations
3. Try implementing your own pipeline variations
4. Move on to the next topic: Fan-Out/Fan-In Pattern

## Troubleshooting

If tests fail:
1. Check the error messages carefully
2. Ensure all dependencies are installed
3. Verify Go version compatibility
4. Check for syntax errors
5. Review the implementation against the requirements

For additional help, refer to the main README.md or the individual exercise comments.

