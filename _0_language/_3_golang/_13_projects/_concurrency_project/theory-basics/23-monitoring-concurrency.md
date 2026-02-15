# 23 - Monitoring Concurrency

## 🎯 Learning Objectives
- Understand how to monitor concurrent applications
- Learn about different monitoring techniques
- Master debugging concurrent code
- Practice with monitoring tools and techniques
- Understand when to use different monitoring approaches

## 📚 Theory

### Why Monitor Concurrent Applications?

**Challenges of concurrent applications:**
1. **Race conditions**: Hard to detect and debug
2. **Deadlocks**: Can cause applications to hang
3. **Resource leaks**: Goroutines that never exit
4. **Performance issues**: Poor concurrency patterns
5. **Non-deterministic behavior**: Results vary between runs

**Benefits of monitoring:**
1. **Early detection**: Catch issues before they become problems
2. **Performance optimization**: Identify bottlenecks
3. **Debugging**: Understand what's happening
4. **Resource management**: Track resource usage
5. **Reliability**: Ensure application stability

### Monitoring Techniques

1. **Logging**: Record events and state changes
2. **Metrics**: Measure performance and behavior
3. **Tracing**: Track requests across goroutines
4. **Profiling**: Analyze performance bottlenecks
5. **Debugging**: Step through concurrent code

## 💻 Code Examples

### Example 1: Basic Logging for Concurrency

```go
package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

func basicLoggingForConcurrency() {
    fmt.Println("=== Basic Logging for Concurrency ===")
    
    // Configure logging
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    
    var wg sync.WaitGroup
    
    // Start multiple goroutines with logging
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            log.Printf("Worker %d started", workerID)
            
            for j := 1; j <= 3; j++ {
                log.Printf("Worker %d: processing item %d", workerID, j)
                time.Sleep(500 * time.Millisecond)
            }
            
            log.Printf("Worker %d completed", workerID)
        }(i)
    }
    
    wg.Wait()
    log.Println("All workers completed")
}

func main() {
    basicLoggingForConcurrency()
}
```

**Run this code:**
```bash
go run 23-monitoring-concurrency.go
```

### Example 2: Metrics Collection

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type Metrics struct {
    RequestsTotal    int64
    RequestsSuccess  int64
    RequestsFailed   int64
    ProcessingTime   int64
    ActiveGoroutines int64
}

func (m *Metrics) IncrementRequests() {
    atomic.AddInt64(&m.RequestsTotal, 1)
}

func (m *Metrics) IncrementSuccess() {
    atomic.AddInt64(&m.RequestsSuccess, 1)
}

func (m *Metrics) IncrementFailed() {
    atomic.AddInt64(&m.RequestsFailed, 1)
}

func (m *Metrics) AddProcessingTime(duration time.Duration) {
    atomic.AddInt64(&m.ProcessingTime, int64(duration))
}

func (m *Metrics) SetActiveGoroutines(count int64) {
    atomic.StoreInt64(&m.ActiveGoroutines, count)
}

func (m *Metrics) GetStats() map[string]int64 {
    return map[string]int64{
        "requests_total":     atomic.LoadInt64(&m.RequestsTotal),
        "requests_success":   atomic.LoadInt64(&m.RequestsSuccess),
        "requests_failed":    atomic.LoadInt64(&m.RequestsFailed),
        "processing_time_ms": atomic.LoadInt64(&m.ProcessingTime) / int64(time.Millisecond),
        "active_goroutines":  atomic.LoadInt64(&m.ActiveGoroutines),
    }
}

func metricsCollection() {
    fmt.Println("=== Metrics Collection ===")
    
    metrics := &Metrics{}
    
    // Start metrics reporter
    go func() {
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()
        
        for range ticker.C {
            stats := metrics.GetStats()
            fmt.Printf("Metrics: %+v\n", stats)
        }
    }()
    
    // Simulate work with metrics
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            metrics.IncrementRequests()
            metrics.SetActiveGoroutines(int64(5 - workerID + 1))
            
            start := time.Now()
            
            // Simulate work
            time.Sleep(time.Duration(workerID) * 200 * time.Millisecond)
            
            if workerID%3 == 0 {
                metrics.IncrementFailed()
                fmt.Printf("Worker %d failed\n", workerID)
            } else {
                metrics.IncrementSuccess()
                fmt.Printf("Worker %d succeeded\n", workerID)
            }
            
            metrics.AddProcessingTime(time.Since(start))
        }(i)
    }
    
    wg.Wait()
    metrics.SetActiveGoroutines(0)
    
    // Final stats
    fmt.Println("Final metrics:")
    stats := metrics.GetStats()
    for key, value := range stats {
        fmt.Printf("  %s: %d\n", key, value)
    }
}

func main() {
    metricsCollection()
}
```

### Example 3: Tracing Concurrent Operations

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Trace struct {
    ID        string
    StartTime time.Time
    EndTime   time.Time
    Events    []Event
    mutex     sync.Mutex
}

type Event struct {
    Timestamp time.Time
    Message   string
    Data      map[string]interface{}
}

func NewTrace(id string) *Trace {
    return &Trace{
        ID:        id,
        StartTime: time.Now(),
        Events:    make([]Event, 0),
    }
}

func (t *Trace) AddEvent(message string, data map[string]interface{}) {
    t.mutex.Lock()
    defer t.mutex.Unlock()
    
    t.Events = append(t.Events, Event{
        Timestamp: time.Now(),
        Message:   message,
        Data:      data,
    })
}

func (t *Trace) Finish() {
    t.EndTime = time.Now()
}

func (t *Trace) Duration() time.Duration {
    if t.EndTime.IsZero() {
        return time.Since(t.StartTime)
    }
    return t.EndTime.Sub(t.StartTime)
}

func (t *Trace) Print() {
    fmt.Printf("Trace %s (duration: %v):\n", t.ID, t.Duration())
    for _, event := range t.Events {
        fmt.Printf("  %v: %s", event.Timestamp.Format("15:04:05.000"), event.Message)
        if event.Data != nil {
            fmt.Printf(" %+v", event.Data)
        }
        fmt.Println()
    }
}

func tracingConcurrentOperations() {
    fmt.Println("=== Tracing Concurrent Operations ===")
    
    // Create trace
    trace := NewTrace("concurrent-operation")
    trace.AddEvent("Starting concurrent operation", nil)
    
    var wg sync.WaitGroup
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            trace.AddEvent(fmt.Sprintf("Worker %d started", workerID), map[string]interface{}{
                "worker_id": workerID,
            })
            
            for j := 1; j <= 2; j++ {
                trace.AddEvent(fmt.Sprintf("Worker %d processing item %d", workerID, j), map[string]interface{}{
                    "worker_id": workerID,
                    "item":      j,
                })
                time.Sleep(300 * time.Millisecond)
            }
            
            trace.AddEvent(fmt.Sprintf("Worker %d completed", workerID), map[string]interface{}{
                "worker_id": workerID,
            })
        }(i)
    }
    
    wg.Wait()
    trace.AddEvent("All workers completed", nil)
    trace.Finish()
    
    // Print trace
    trace.Print()
}

func main() {
    tracingConcurrentOperations()
}
```

### Example 4: Debugging Race Conditions

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func debuggingRaceConditions() {
    fmt.Println("=== Debugging Race Conditions ===")
    
    // This code has a race condition
    var counter int
    var mutex sync.Mutex
    
    // Start multiple goroutines
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            // This is safe - using mutex
            mutex.Lock()
            counter++
            mutex.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter value: %d (should be 1000)\n", counter)
    
    // This code has a race condition (commented out to avoid actual race)
    /*
    var counter2 int
    for i := 0; i < 1000; i++ {
        go func() {
            counter2++ // Race condition!
        }()
    }
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter2 value: %d (unpredictable)\n", counter2)
    */
}

func main() {
    debuggingRaceConditions()
}
```

### Example 5: Performance Profiling

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func performanceProfiling() {
    fmt.Println("=== Performance Profiling ===")
    
    // Start profiling
    start := time.Now()
    startMem := getMemStats()
    
    // Simulate work
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            // Simulate CPU work
            sum := 0
            for j := 0; j < 1000; j++ {
                sum += j
            }
        }()
    }
    
    wg.Wait()
    
    // End profiling
    end := time.Now()
    endMem := getMemStats()
    
    // Print results
    fmt.Printf("Duration: %v\n", end.Sub(start))
    fmt.Printf("Memory before: %+v\n", startMem)
    fmt.Printf("Memory after: %+v\n", endMem)
    fmt.Printf("Memory delta: %d bytes\n", endMem.Alloc-startMem.Alloc)
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
}

func getMemStats() runtime.MemStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m
}

func main() {
    performanceProfiling()
}
```

### Example 6: Monitoring Goroutine Leaks

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func monitoringGoroutineLeaks() {
    fmt.Println("=== Monitoring Goroutine Leaks ===")
    
    // Monitor goroutine count
    go func() {
        ticker := time.NewTicker(500 * time.Millisecond)
        defer ticker.Stop()
        
        for range ticker.C {
            count := runtime.NumGoroutine()
            fmt.Printf("Goroutines: %d\n", count)
        }
    }()
    
    // Start some goroutines
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            fmt.Printf("Worker %d started\n", workerID)
            time.Sleep(2 * time.Second)
            fmt.Printf("Worker %d completed\n", workerID)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
    
    // Wait a bit to see goroutine count
    time.Sleep(1 * time.Second)
}

func main() {
    monitoringGoroutineLeaks()
}
```

### Example 7: Health Checks

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type HealthChecker struct {
    isHealthy int32
    mutex     sync.RWMutex
    checks    map[string]func() bool
}

func NewHealthChecker() *HealthChecker {
    return &HealthChecker{
        checks: make(map[string]func() bool),
    }
}

func (hc *HealthChecker) AddCheck(name string, check func() bool) {
    hc.mutex.Lock()
    defer hc.mutex.Unlock()
    hc.checks[name] = check
}

func (hc *HealthChecker) IsHealthy() bool {
    return atomic.LoadInt32(&hc.isHealthy) == 1
}

func (hc *HealthChecker) CheckHealth() {
    hc.mutex.RLock()
    defer hc.mutex.RUnlock()
    
    allHealthy := true
    for name, check := range hc.checks {
        if !check() {
            fmt.Printf("Health check failed: %s\n", name)
            allHealthy = false
        }
    }
    
    if allHealthy {
        atomic.StoreInt32(&hc.isHealthy, 1)
        fmt.Println("All health checks passed")
    } else {
        atomic.StoreInt32(&hc.isHealthy, 0)
        fmt.Println("Some health checks failed")
    }
}

func (hc *HealthChecker) StartMonitoring(interval time.Duration) {
    go func() {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()
        
        for range ticker.C {
            hc.CheckHealth()
        }
    }()
}

func healthChecks() {
    fmt.Println("=== Health Checks ===")
    
    hc := NewHealthChecker()
    
    // Add health checks
    hc.AddCheck("database", func() bool {
        // Simulate database check
        time.Sleep(100 * time.Millisecond)
        return true
    })
    
    hc.AddCheck("cache", func() bool {
        // Simulate cache check
        time.Sleep(50 * time.Millisecond)
        return true
    })
    
    hc.AddCheck("api", func() bool {
        // Simulate API check
        time.Sleep(200 * time.Millisecond)
        return true
    })
    
    // Start monitoring
    hc.StartMonitoring(1 * time.Second)
    
    // Simulate work
    time.Sleep(5 * time.Second)
    
    // Check final health
    if hc.IsHealthy() {
        fmt.Println("System is healthy")
    } else {
        fmt.Println("System is unhealthy")
    }
}

func main() {
    healthChecks()
}
```

### Example 8: Monitoring Best Practices

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"
    "sync/atomic"
    "time"
)

type Monitor struct {
    metrics    map[string]int64
    mutex      sync.RWMutex
    ctx        context.Context
    cancel     context.CancelFunc
    logLevel   int
}

func NewMonitor() *Monitor {
    ctx, cancel := context.WithCancel(context.Background())
    return &Monitor{
        metrics:  make(map[string]int64),
        ctx:      ctx,
        cancel:   cancel,
        logLevel: 1,
    }
}

func (m *Monitor) SetLogLevel(level int) {
    m.logLevel = level
}

func (m *Monitor) Log(level int, message string, data map[string]interface{}) {
    if level <= m.logLevel {
        log.Printf("[%d] %s %+v", level, message, data)
    }
}

func (m *Monitor) IncrementMetric(name string) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.metrics[name]++
}

func (m *Monitor) SetMetric(name string, value int64) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.metrics[name] = value
}

func (m *Monitor) GetMetrics() map[string]int64 {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    result := make(map[string]int64)
    for k, v := range m.metrics {
        result[k] = v
    }
    return result
}

func (m *Monitor) StartReporting(interval time.Duration) {
    go func() {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()
        
        for {
            select {
            case <-ticker.C:
                metrics := m.GetMetrics()
                m.Log(1, "Metrics report", map[string]interface{}{
                    "metrics": metrics,
                })
            case <-m.ctx.Done():
                return
            }
        }
    }()
}

func (m *Monitor) Stop() {
    m.cancel()
}

func monitoringBestPractices() {
    fmt.Println("=== Monitoring Best Practices ===")
    
    // Create monitor
    monitor := NewMonitor()
    monitor.SetLogLevel(1)
    
    // Start reporting
    monitor.StartReporting(1 * time.Second)
    
    // Simulate work with monitoring
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            monitor.Log(2, "Worker started", map[string]interface{}{
                "worker_id": workerID,
            })
            
            monitor.IncrementMetric("workers_started")
            
            for j := 1; j <= 3; j++ {
                monitor.Log(3, "Processing item", map[string]interface{}{
                    "worker_id": workerID,
                    "item":      j,
                })
                
                time.Sleep(500 * time.Millisecond)
                monitor.IncrementMetric("items_processed")
            }
            
            monitor.Log(2, "Worker completed", map[string]interface{}{
                "worker_id": workerID,
            })
            
            monitor.IncrementMetric("workers_completed")
        }(i)
    }
    
    wg.Wait()
    
    // Final metrics
    monitor.Log(1, "Final metrics", map[string]interface{}{
        "metrics": monitor.GetMetrics(),
    })
    
    // Stop monitor
    monitor.Stop()
}

func main() {
    monitoringBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Logging**: Record events and state changes
2. **Metrics**: Measure performance and behavior
3. **Tracing**: Track requests across goroutines
4. **Profiling**: Analyze performance bottlenecks
5. **Debugging**: Step through concurrent code
6. **Health checks**: Monitor system health

## 🎯 When to Use Different Monitoring Approaches

1. **Logging**: When you need to record events and debug issues
2. **Metrics**: When you need to measure performance and behavior
3. **Tracing**: When you need to track requests across goroutines
4. **Profiling**: When you need to analyze performance bottlenecks
5. **Health checks**: When you need to monitor system health

## 🎯 Best Practices

1. **Use structured logging**:
   ```go
   // GOOD - structured logging
   log.Printf("Worker %d processed item %d", workerID, itemID)
   
   // BAD - unstructured logging
   log.Printf("Worker processed item")
   ```

2. **Use atomic operations for metrics**:
   ```go
   // GOOD - atomic operations
   atomic.AddInt64(&counter, 1)
   
   // BAD - race condition
   counter++
   ```

3. **Monitor goroutine count**:
   ```go
   // GOOD - monitor goroutines
   go func() {
       ticker := time.NewTicker(1 * time.Second)
       for range ticker.C {
           fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
       }
   }()
   ```

4. **Use context for cancellation**:
   ```go
   // GOOD - use context
   select {
   case <-ctx.Done():
       return
   default:
       // Work
   }
   ```

## 🎯 Common Pitfalls

1. **Not monitoring goroutine leaks**:
   ```go
   // BAD - no monitoring
   go func() {
       // Long running operation
   }()
   
   // GOOD - monitor goroutines
   go func() {
       defer func() {
           // Cleanup
       }()
       // Long running operation
   }()
   ```

2. **Not using atomic operations for metrics**:
   ```go
   // BAD - race condition
   var counter int
   counter++
   
   // GOOD - atomic operations
   var counter int64
   atomic.AddInt64(&counter, 1)
   ```

3. **Not handling errors in monitoring**:
   ```go
   // BAD - ignore errors
   go func() {
       // Monitoring code
   }()
   
   // GOOD - handle errors
   go func() {
       defer func() {
           if r := recover(); r != nil {
               // Handle panic
           }
       }()
       // Monitoring code
   }()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a concurrent system with comprehensive monitoring:
- Multiple services that can fail
- Metrics collection for performance
- Health checks for system status
- Tracing for debugging
- Proper error handling and logging

**Hint**: Use a struct to represent the system and implement monitoring methods for different aspects.

## 🚀 Next Steps

Now that you understand monitoring concurrency, let's learn about **performance optimization** in the next file: `24-performance-optimization.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
