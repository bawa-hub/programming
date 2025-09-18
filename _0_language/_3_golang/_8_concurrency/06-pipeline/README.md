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

