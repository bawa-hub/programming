# 18 - Worker Pools

## 🎯 Learning Objectives
- Understand what worker pools are and why they're useful
- Learn how to implement worker pools in Go
- Master different worker pool patterns
- Practice with worker pool best practices
- Understand when to use worker pools

## 📚 Theory

### What is a Worker Pool?

A **worker pool** is a pattern where a fixed number of worker goroutines process jobs from a queue.

**Key characteristics:**
- **Fixed number of workers**: Controlled concurrency
- **Job queue**: Jobs are queued and processed by workers
- **Load balancing**: Work is distributed evenly
- **Resource management**: Limits resource usage

### Why Use Worker Pools?

**Benefits:**
1. **Controlled concurrency**: Limit number of concurrent operations
2. **Resource management**: Prevent resource exhaustion
3. **Load balancing**: Distribute work evenly
4. **Backpressure**: Handle high load gracefully

**Use cases:**
1. **Web scraping**: Limit concurrent requests
2. **File processing**: Process files in parallel
3. **API calls**: Limit concurrent API requests
4. **Database operations**: Control database connections

## 💻 Code Examples

### Example 1: Basic Worker Pool

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicWorkerPool() {
    fmt.Println("=== Basic Worker Pool ===")
    
    // Create job queue
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Start workers
    numWorkers := 3
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d\n", workerID, job)
                time.Sleep(500 * time.Millisecond) // Simulate work
                results <- job * 2
            }
        }(i)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 5; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    basicWorkerPool()
}
```

**Run this code:**
```bash
go run 18-worker-pools.go
```

### Example 2: Worker Pool with Different Job Types

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID     int
    Type   string
    Data   string
    Result chan<- string
}

func workerPoolWithJobTypes() {
    fmt.Println("=== Worker Pool with Different Job Types ===")
    
    // Create job queue
    jobs := make(chan Job, 10)
    
    // Start workers
    numWorkers := 3
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d (type: %s)\n", workerID, job.ID, job.Type)
                
                // Process based on job type
                var result string
                switch job.Type {
                case "email":
                    result = fmt.Sprintf("Email sent: %s", job.Data)
                case "sms":
                    result = fmt.Sprintf("SMS sent: %s", job.Data)
                case "push":
                    result = fmt.Sprintf("Push notification sent: %s", job.Data)
                default:
                    result = fmt.Sprintf("Unknown job type: %s", job.Type)
                }
                
                time.Sleep(500 * time.Millisecond) // Simulate work
                job.Result <- result
            }
        }(i)
    }
    
    // Send jobs
    go func() {
        jobTypes := []string{"email", "sms", "push", "email", "sms"}
        for i, jobType := range jobTypes {
            resultChan := make(chan string, 1)
            job := Job{
                ID:     i + 1,
                Type:   jobType,
                Data:   fmt.Sprintf("data for job %d", i+1),
                Result: resultChan,
            }
            jobs <- job
            
            // Wait for result
            result := <-resultChan
            fmt.Printf("Job %d completed: %s\n", job.ID, result)
        }
        close(jobs)
    }()
    
    wg.Wait()
}

func main() {
    workerPoolWithJobTypes()
}
```

### Example 3: Worker Pool with Context

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func workerPoolWithContext() {
    fmt.Println("=== Worker Pool with Context ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // Create job queue
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Start workers
    numWorkers := 3
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    fmt.Printf("Worker %d processing job %d\n", workerID, job)
                    time.Sleep(500 * time.Millisecond)
                    results <- job * 2
                case <-ctx.Done():
                    fmt.Printf("Worker %d cancelled: %v\n", workerID, ctx.Err())
                    return
                }
            }
        }(i)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 10; i++ {
            select {
            case jobs <- i:
                fmt.Printf("Sent job %d\n", i)
            case <-ctx.Done():
                close(jobs)
                return
            }
            time.Sleep(200 * time.Millisecond)
        }
        close(jobs)
    }()
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    workerPoolWithContext()
}
```

### Example 4: Worker Pool with Rate Limiting

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func workerPoolWithRateLimiting() {
    fmt.Println("=== Worker Pool with Rate Limiting ===")
    
    // Create job queue
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Rate limiter
    rateLimiter := time.NewTicker(200 * time.Millisecond)
    defer rateLimiter.Stop()
    
    // Start workers
    numWorkers := 2
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                <-rateLimiter.C // Wait for rate limiter
                fmt.Printf("Worker %d processing job %d\n", workerID, job)
                time.Sleep(300 * time.Millisecond) // Simulate work
                results <- job * 2
            }
        }(i)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 5; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    workerPoolWithRateLimiting()
}
```

### Example 5: Worker Pool with Error Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type JobResult struct {
    JobID int
    Value int
    Error error
}

func workerPoolWithErrorHandling() {
    fmt.Println("=== Worker Pool with Error Handling ===")
    
    // Create job queue
    jobs := make(chan int, 10)
    results := make(chan JobResult, 10)
    
    // Start workers
    numWorkers := 3
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d\n", workerID, job)
                
                // Simulate work that might fail
                time.Sleep(500 * time.Millisecond)
                
                var result JobResult
                if job%3 == 0 {
                    result = JobResult{
                        JobID: job,
                        Value: 0,
                        Error: fmt.Errorf("job %d failed", job),
                    }
                } else {
                    result = JobResult{
                        JobID: job,
                        Value: job * 2,
                        Error: nil,
                    }
                }
                
                results <- result
            }
        }(i)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 6; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        if result.Error != nil {
            fmt.Printf("Job %d failed: %v\n", result.JobID, result.Error)
        } else {
            fmt.Printf("Job %d completed: %d\n", result.JobID, result.Value)
        }
    }
}

func main() {
    workerPoolWithErrorHandling()
}
```

### Example 6: Worker Pool with Priority Queue

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type PriorityJob struct {
    ID       int
    Priority int
    Data     string
}

func workerPoolWithPriorityQueue() {
    fmt.Println("=== Worker Pool with Priority Queue ===")
    
    // Create job queue
    jobs := make(chan PriorityJob, 10)
    results := make(chan string, 10)
    
    // Start workers
    numWorkers := 2
    var wg sync.WaitGroup
    
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d (priority: %d)\n", workerID, job.ID, job.Priority)
                time.Sleep(time.Duration(job.Priority) * 200 * time.Millisecond)
                result := fmt.Sprintf("Job %d completed (priority: %d)", job.ID, job.Priority)
                results <- result
            }
        }(i)
    }
    
    // Send jobs with different priorities
    go func() {
        jobs <- PriorityJob{ID: 1, Priority: 3, Data: "Low priority"}
        jobs <- PriorityJob{ID: 2, Priority: 1, Data: "High priority"}
        jobs <- PriorityJob{ID: 3, Priority: 2, Data: "Medium priority"}
        jobs <- PriorityJob{ID: 4, Priority: 1, Data: "High priority"}
        jobs <- PriorityJob{ID: 5, Priority: 3, Data: "Low priority"}
        close(jobs)
    }()
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %s\n", result)
    }
}

func main() {
    workerPoolWithPriorityQueue()
}
```

### Example 7: Worker Pool Best Practices

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type WorkerPool struct {
    workers    int
    jobs       chan Job
    results    chan Result
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

type Job struct {
    ID   int
    Data string
}

type Result struct {
    JobID int
    Value string
    Error error
}

func NewWorkerPool(workers int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    return &WorkerPool{
        workers: workers,
        jobs:    make(chan Job, 100),
        results: make(chan Result, 100),
        ctx:     ctx,
        cancel:  cancel,
    }
}

func (wp *WorkerPool) Start() {
    for i := 1; i <= wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    for {
        select {
        case job, ok := <-wp.jobs:
            if !ok {
                return
            }
            wp.processJob(id, job)
        case <-wp.ctx.Done():
            return
        }
    }
}

func (wp *WorkerPool) processJob(workerID int, job Job) {
    fmt.Printf("Worker %d processing job %d\n", workerID, job.ID)
    time.Sleep(500 * time.Millisecond) // Simulate work
    
    result := Result{
        JobID: job.ID,
        Value: fmt.Sprintf("Processed: %s", job.Data),
        Error: nil,
    }
    
    select {
    case wp.results <- result:
    case <-wp.ctx.Done():
    }
}

func (wp *WorkerPool) Submit(job Job) {
    select {
    case wp.jobs <- job:
    case <-wp.ctx.Done():
    }
}

func (wp *WorkerPool) GetResult() <-chan Result {
    return wp.results
}

func (wp *WorkerPool) Stop() {
    wp.cancel()
    close(wp.jobs)
    wp.wg.Wait()
    close(wp.results)
}

func workerPoolBestPractices() {
    fmt.Println("=== Worker Pool Best Practices ===")
    
    // Create worker pool
    pool := NewWorkerPool(3)
    pool.Start()
    
    // Submit jobs
    go func() {
        for i := 1; i <= 5; i++ {
            job := Job{
                ID:   i,
                Data: fmt.Sprintf("data-%d", i),
            }
            pool.Submit(job)
        }
    }()
    
    // Collect results
    go func() {
        for result := range pool.GetResult() {
            if result.Error != nil {
                fmt.Printf("Job %d failed: %v\n", result.JobID, result.Error)
            } else {
                fmt.Printf("Job %d completed: %s\n", result.JobID, result.Value)
            }
        }
    }()
    
    time.Sleep(3 * time.Second)
    pool.Stop()
}

func main() {
    workerPoolBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Fixed number of workers**: Control concurrency
2. **Job queue**: Jobs are queued and processed
3. **Load balancing**: Work is distributed evenly
4. **Resource management**: Limit resource usage
5. **Backpressure**: Handle high load gracefully

## 🎯 When to Use Worker Pools

1. **Controlled concurrency**: When you need to limit concurrent operations
2. **Resource management**: When you need to control resource usage
3. **Load balancing**: When you need to distribute work evenly
4. **Backpressure**: When you need to handle high load gracefully

## 🎯 Best Practices

1. **Use appropriate worker count**:
   ```go
   // GOOD - reasonable worker count
   numWorkers := runtime.NumCPU()
   
   // BAD - too many workers
   numWorkers := 1000
   ```

2. **Handle errors properly**:
   ```go
   // GOOD - handle errors
   if err := processJob(job); err != nil {
       results <- Result{Error: err}
   }
   ```

3. **Use context for cancellation**:
   ```go
   // GOOD - use context
   select {
   case job := <-jobs:
       processJob(job)
   case <-ctx.Done():
       return
   }
   ```

4. **Close channels properly**:
   ```go
   // GOOD - close channels
   defer close(jobs)
   defer close(results)
   ```

## 🎯 Common Pitfalls

1. **Too many workers**:
   ```go
   // BAD - too many workers
   numWorkers := 1000
   
   // GOOD - reasonable worker count
   numWorkers := runtime.NumCPU()
   ```

2. **Not handling errors**:
   ```go
   // BAD - ignore errors
   processJob(job)
   
   // GOOD - handle errors
   if err := processJob(job); err != nil {
       // Handle error
   }
   ```

3. **Not closing channels**:
   ```go
   // BAD - channels not closed
   jobs := make(chan Job)
   
   // GOOD - close channels
   defer close(jobs)
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a web scraper with worker pools:
- Use worker pools to limit concurrent requests
- Implement rate limiting to respect target site limits
- Handle errors and timeouts properly
- Show how worker pools can improve performance

**Hint**: Use a worker pool to process URLs and implement rate limiting to avoid overwhelming the target site.

## 🚀 Next Steps

Now that you understand worker pools, let's learn about **pipeline patterns** in the next file: `19-pipeline-patterns.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
