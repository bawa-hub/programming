# 🚀 Worker Pool Pattern: Efficient Concurrent Task Processing

## 📚 Table of Contents
1. [What is the Worker Pool Pattern?](#what-is-the-worker-pool-pattern)
2. [Basic Worker Pool](#basic-worker-pool)
3. [Buffered Worker Pool](#buffered-worker-pool)
4. [Dynamic Worker Pool](#dynamic-worker-pool)
5. [Priority Worker Pool](#priority-worker-pool)
6. [Worker Pool with Results](#worker-pool-with-results)
7. [Worker Pool with Error Handling](#worker-pool-with-error-handling)
8. [Worker Pool with Timeout](#worker-pool-with-timeout)
9. [Worker Pool with Rate Limiting](#worker-pool-with-rate-limiting)
10. [Worker Pool with Metrics](#worker-pool-with-metrics)
11. [Performance Considerations](#performance-considerations)
12. [Common Patterns](#common-patterns)
13. [Best Practices](#best-practices)
14. [Common Pitfalls](#common-pitfalls)
15. [Exercises](#exercises)

---

## 🎯 What is the Worker Pool Pattern?

The **Worker Pool Pattern** is a concurrency design pattern that manages a fixed number of worker goroutines to process tasks from a queue. It's one of the most efficient ways to handle concurrent workloads in Go.

### Key Characteristics:
- **Fixed number of workers**: Pre-created goroutines that process tasks
- **Task queue**: Channel-based queue for distributing work
- **Load balancing**: Tasks are distributed evenly among workers
- **Resource control**: Limits concurrent operations to prevent resource exhaustion
- **Scalability**: Can handle high-volume workloads efficiently

### Benefits:
- **Memory efficient**: Reuses goroutines instead of creating new ones
- **CPU efficient**: Optimal number of workers for available cores
- **Predictable performance**: Fixed resource usage
- **Easy to reason about**: Clear separation of concerns
- **Backpressure handling**: Queue size limits prevent memory issues

---

## 🏗️ Basic Worker Pool

A basic worker pool consists of:
1. **Worker goroutines** that process tasks
2. **Task channel** for distributing work
3. **WaitGroup** for synchronization
4. **Task submission** mechanism

### Basic Structure:
```go
func workerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
}
```

### Key Components:
- **Jobs channel**: Input channel for tasks
- **Results channel**: Output channel for results
- **Worker goroutines**: Process tasks concurrently
- **WaitGroup**: Waits for all workers to complete
- **Channel closing**: Signals completion

---

## 📦 Buffered Worker Pool

A buffered worker pool uses buffered channels to improve performance and handle bursts of work.

### Benefits:
- **Better throughput**: Buffers smooth out work distribution
- **Reduced blocking**: Workers don't block on channel operations
- **Burst handling**: Can handle sudden spikes in workload
- **Memory control**: Buffer size limits memory usage

### Implementation:
```go
func bufferedWorkerPool(jobs <-chan Job, results chan<- Result, bufferSize int) {
    var wg sync.WaitGroup
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                select {
                case results <- result:
                case <-time.After(timeout):
                    // Handle timeout
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
}
```

---

## 🔄 Dynamic Worker Pool

A dynamic worker pool can adjust the number of workers based on workload or system conditions.

### Features:
- **Adaptive scaling**: Adjusts worker count based on load
- **Resource monitoring**: Monitors CPU, memory, or queue size
- **Graceful scaling**: Adds/removes workers smoothly
- **Performance optimization**: Maintains optimal worker count

### Implementation:
```go
type DynamicWorkerPool struct {
    workers    int
    minWorkers int
    maxWorkers int
    jobs       chan Job
    results    chan Result
    mu         sync.RWMutex
}

func (dwp *DynamicWorkerPool) adjustWorkers() {
    dwp.mu.Lock()
    defer dwp.mu.Unlock()
    
    queueSize := len(dwp.jobs)
    
    if queueSize > threshold && dwp.workers < dwp.maxWorkers {
        // Add workers
        dwp.addWorker()
    } else if queueSize < lowThreshold && dwp.workers > dwp.minWorkers {
        // Remove workers
        dwp.removeWorker()
    }
}
```

---

## ⭐ Priority Worker Pool

A priority worker pool processes tasks based on priority levels.

### Features:
- **Priority queues**: Multiple queues for different priority levels
- **Priority processing**: High-priority tasks are processed first
- **Fair scheduling**: Prevents starvation of low-priority tasks
- **Dynamic priorities**: Can change task priorities at runtime

### Implementation:
```go
type PriorityWorkerPool struct {
    highPriority chan Job
    lowPriority  chan Job
    results      chan Result
    workers      int
}

func (pwp *PriorityWorkerPool) processJobs() {
    for {
        select {
        case job := <-pwp.highPriority:
            result := process(job)
            pwp.results <- result
        case job := <-pwp.lowPriority:
            result := process(job)
            pwp.results <- result
        case <-time.After(timeout):
            // Handle timeout
        }
    }
}
```

---

## 📊 Worker Pool with Results

A worker pool that collects and processes results from workers.

### Features:
- **Result collection**: Gathers results from all workers
- **Result processing**: Can process results as they arrive
- **Error handling**: Handles errors from individual workers
- **Result aggregation**: Combines results from multiple workers

### Implementation:
```go
func workerPoolWithResults(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    
    // Wait for all workers to complete
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Process results
    for result := range results {
        handleResult(result)
    }
}
```

---

## ⚠️ Worker Pool with Error Handling

A worker pool that properly handles errors from workers.

### Features:
- **Error collection**: Collects errors from workers
- **Error recovery**: Can recover from worker failures
- **Error reporting**: Reports errors to monitoring systems
- **Graceful degradation**: Continues processing despite errors

### Implementation:
```go
type WorkerPoolWithErrors struct {
    jobs    chan Job
    results chan Result
    errors  chan error
    workers int
}

func (wpe *WorkerPoolWithErrors) processWithErrors() {
    var wg sync.WaitGroup
    
    for i := 0; i < wpe.workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range wpe.jobs {
                result, err := processWithError(job)
                if err != nil {
                    wpe.errors <- err
                } else {
                    wpe.results <- result
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(wpe.results)
    close(wpe.errors)
}
```

---

## ⏰ Worker Pool with Timeout

A worker pool that handles timeouts for individual tasks and overall processing.

### Features:
- **Task timeouts**: Individual tasks can timeout
- **Overall timeout**: Entire worker pool can timeout
- **Timeout handling**: Graceful handling of timeouts
- **Resource cleanup**: Cleans up resources on timeout

### Implementation:
```go
func workerPoolWithTimeout(jobs <-chan Job, results chan<- Result, timeout time.Duration) {
    var wg sync.WaitGroup
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    result := processWithContext(ctx, job)
                    results <- result
                case <-ctx.Done():
                    return
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
}
```

---

## 🚦 Worker Pool with Rate Limiting

A worker pool that limits the rate of task processing.

### Features:
- **Rate limiting**: Controls the rate of task processing
- **Burst handling**: Allows bursts up to a limit
- **Adaptive limiting**: Can adjust rate based on conditions
- **Backpressure**: Slows down when rate limit is exceeded

### Implementation:
```go
type RateLimitedWorkerPool struct {
    jobs       chan Job
    results    chan Result
    rateLimiter *rate.Limiter
    workers    int
}

func (rlwp *RateLimitedWorkerPool) processWithRateLimit() {
    var wg sync.WaitGroup
    
    for i := 0; i < rlwp.workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range rlwp.jobs {
                // Wait for rate limiter
                rlwp.rateLimiter.Wait(context.Background())
                
                result := process(job)
                rlwp.results <- result
            }
        }(i)
    }
    
    wg.Wait()
    close(rlwp.results)
}
```

---

## 📈 Worker Pool with Metrics

A worker pool that collects and reports metrics about its performance.

### Features:
- **Performance metrics**: Tracks processing times, throughput
- **Resource metrics**: Monitors CPU, memory usage
- **Error metrics**: Tracks error rates and types
- **Real-time monitoring**: Provides real-time metrics

### Implementation:
```go
type MetricsWorkerPool struct {
    jobs     chan Job
    results  chan Result
    workers  int
    metrics  *Metrics
}

type Metrics struct {
    ProcessedTasks int64
    ProcessingTime time.Duration
    ErrorCount     int64
    mu             sync.RWMutex
}

func (mwp *MetricsWorkerPool) processWithMetrics() {
    var wg sync.WaitGroup
    
    for i := 0; i < mwp.workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range mwp.jobs {
                start := time.Now()
                
                result, err := process(job)
                if err != nil {
                    mwp.metrics.incrementErrorCount()
                } else {
                    mwp.results <- result
                }
                
                mwp.metrics.recordProcessingTime(time.Since(start))
                mwp.metrics.incrementProcessedTasks()
            }
        }(i)
    }
    
    wg.Wait()
    close(mwp.results)
}
```

---

## 📊 Performance Considerations

### 1. **Worker Count Optimization**
- **CPU-bound tasks**: Number of workers = number of CPU cores
- **I/O-bound tasks**: Number of workers = 2-4x number of CPU cores
- **Mixed workloads**: Start with CPU cores, adjust based on performance

### 2. **Channel Buffer Sizing**
- **Small buffers**: Better for memory usage, can cause blocking
- **Large buffers**: Better for throughput, uses more memory
- **Dynamic sizing**: Adjust based on workload characteristics

### 3. **Memory Management**
- **Object pooling**: Reuse objects to reduce GC pressure
- **Channel management**: Close channels properly to prevent leaks
- **Resource cleanup**: Ensure all resources are cleaned up

### 4. **Load Balancing**
- **Work stealing**: Workers can steal work from other workers
- **Round-robin**: Distribute work evenly among workers
- **Priority-based**: Process high-priority work first

---

## 🎨 Common Patterns

### 1. **Pipeline Pattern**
```go
func pipelineWorkerPool(input <-chan Job, output chan<- Result) {
    // Stage 1: Process input
    stage1 := make(chan ProcessedJob, bufferSize)
    go func() {
        defer close(stage1)
        for job := range input {
            stage1 <- processStage1(job)
        }
    }()
    
    // Stage 2: Process stage1 output
    stage2 := make(chan Result, bufferSize)
    go func() {
        defer close(stage2)
        for job := range stage1 {
            stage2 <- processStage2(job)
        }
    }()
    
    // Collect results
    for result := range stage2 {
        output <- result
    }
}
```

### 2. **Fan-Out/Fan-In Pattern**
```go
func fanOutFanInWorkerPool(input <-chan Job, output chan<- Result) {
    // Fan-out: Distribute work to multiple workers
    workers := make([]chan Job, numWorkers)
    for i := range workers {
        workers[i] = make(chan Job)
    }
    
    // Start workers
    var wg sync.WaitGroup
    for i, worker := range workers {
        wg.Add(1)
        go func(id int, jobs <-chan Job) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
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
        
        for job := range input {
            // Round-robin distribution
            workers[job.ID%numWorkers] <- job
        }
    }()
    
    wg.Wait()
    close(output)
}
```

### 3. **Circuit Breaker Pattern**
```go
type CircuitBreakerWorkerPool struct {
    jobs      chan Job
    results   chan Result
    breaker   *CircuitBreaker
    workers   int
}

func (cbwp *CircuitBreakerWorkerPool) processWithCircuitBreaker() {
    var wg sync.WaitGroup
    
    for i := 0; i < cbwp.workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range cbwp.jobs {
                if cbwp.breaker.Allow() {
                    result, err := process(job)
                    if err != nil {
                        cbwp.breaker.RecordFailure()
                    } else {
                        cbwp.breaker.RecordSuccess()
                        cbwp.results <- result
                    }
                } else {
                    // Circuit breaker is open, skip job
                    cbwp.results <- Result{Error: "Circuit breaker open"}
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(cbwp.results)
}
```

---

## ✅ Best Practices

### 1. **Proper Resource Management**
```go
// ✅ Good - proper cleanup
func workerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}

// ❌ Bad - no cleanup
func badWorkerPool(jobs <-chan Job, results chan<- Result) {
    for i := 0; i < numWorkers; i++ {
        go func(id int) {
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    // No cleanup, results channel never closed
}
```

### 2. **Error Handling**
```go
// ✅ Good - proper error handling
func workerPoolWithErrors(jobs <-chan Job, results chan<- Result, errors chan<- error) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result, err := process(job)
                if err != nil {
                    errors <- err
                } else {
                    results <- result
                }
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
        close(errors)
    }()
}
```

### 3. **Context Usage**
```go
// ✅ Good - using context for cancellation
func workerPoolWithContext(ctx context.Context, jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    result := processWithContext(ctx, job)
                    results <- result
                case <-ctx.Done():
                    return
                }
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 4. **Metrics and Monitoring**
```go
// ✅ Good - collecting metrics
type WorkerPoolMetrics struct {
    ProcessedTasks int64
    ProcessingTime time.Duration
    ErrorCount     int64
    mu             sync.RWMutex
}

func (wpm *WorkerPoolMetrics) recordTask(start time.Time, err error) {
    wpm.mu.Lock()
    defer wpm.mu.Unlock()
    
    wpm.ProcessedTasks++
    wpm.ProcessingTime += time.Since(start)
    if err != nil {
        wpm.ErrorCount++
    }
}
```

---

## ⚠️ Common Pitfalls

### 1. **Goroutine Leaks**
```go
// ❌ Wrong - goroutine leak
func badWorkerPool(jobs <-chan Job, results chan<- Result) {
    for i := 0; i < numWorkers; i++ {
        go func(id int) {
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    // No WaitGroup, no cleanup
}

// ✅ Correct - proper cleanup
func goodWorkerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 2. **Channel Deadlocks**
```go
// ❌ Wrong - potential deadlock
func badWorkerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result // Can block if results channel is full
            }
        }(i)
    }
    
    wg.Wait() // This will block forever if results channel is full
    close(results)
}

// ✅ Correct - non-blocking send
func goodWorkerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                select {
                case results <- result:
                case <-time.After(timeout):
                    // Handle timeout
                }
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 3. **Incorrect Worker Count**
```go
// ❌ Wrong - too many workers
func badWorkerPool(jobs <-chan Job, results chan<- Result) {
    for i := 0; i < 1000; i++ { // Too many workers
        go func(id int) {
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
}

// ✅ Correct - optimal worker count
func goodWorkerPool(jobs <-chan Job, results chan<- Result) {
    numWorkers := runtime.NumCPU() // Optimal for CPU-bound tasks
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job)
                results <- result
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 4. **Missing Error Handling**
```go
// ❌ Wrong - no error handling
func badWorkerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                result := process(job) // Can panic
                results <- result
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
}

// ✅ Correct - proper error handling
func goodWorkerPool(jobs <-chan Job, results chan<- Result, errors chan<- error) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                func() {
                    defer func() {
                        if r := recover(); r != nil {
                            errors <- fmt.Errorf("worker %d panicked: %v", id, r)
                        }
                    }()
                    
                    result, err := process(job)
                    if err != nil {
                        errors <- err
                    } else {
                        results <- result
                    }
                }()
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
        close(errors)
    }()
}
```

---

## 🧪 Exercises

### Exercise 1: Basic Worker Pool
Create a basic worker pool that processes tasks from a channel.

### Exercise 2: Buffered Worker Pool
Implement a worker pool with buffered channels for better performance.

### Exercise 3: Dynamic Worker Pool
Create a worker pool that can adjust the number of workers based on workload.

### Exercise 4: Priority Worker Pool
Implement a worker pool that processes tasks based on priority.

### Exercise 5: Worker Pool with Results
Create a worker pool that collects and processes results from workers.

### Exercise 6: Worker Pool with Error Handling
Implement a worker pool that properly handles errors from workers.

### Exercise 7: Worker Pool with Timeout
Create a worker pool that handles timeouts for individual tasks.

### Exercise 8: Worker Pool with Rate Limiting
Implement a worker pool that limits the rate of task processing.

### Exercise 9: Worker Pool with Metrics
Create a worker pool that collects and reports performance metrics.

### Exercise 10: Pipeline Worker Pool
Implement a worker pool that processes tasks through multiple stages.

---

## 🎯 Key Takeaways

1. **Use worker pools for concurrent task processing** - efficient and predictable
2. **Choose optimal worker count** - based on workload characteristics
3. **Handle errors properly** - collect and process errors from workers
4. **Use context for cancellation** - graceful shutdown and timeout handling
5. **Collect metrics** - monitor performance and resource usage
6. **Avoid common pitfalls** - goroutine leaks, deadlocks, incorrect worker count
7. **Implement proper cleanup** - close channels and wait for completion
8. **Consider different patterns** - pipeline, fan-out/fan-in, circuit breaker

---

## 🚀 Next Steps

Ready for the next topic? Let's move on to **Pipeline Pattern** where you'll learn how to process data through multiple stages!

**Run the examples in this directory to see worker pools in action!**
