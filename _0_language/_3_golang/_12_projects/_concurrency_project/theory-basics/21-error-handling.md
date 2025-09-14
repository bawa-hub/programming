# 21 - Error Handling in Concurrent Code

## 🎯 Learning Objectives
- Understand how to handle errors in concurrent code
- Learn different error handling patterns
- Master error propagation in goroutines
- Practice with error handling best practices
- Understand when to use different error handling approaches

## 📚 Theory

### Why is Error Handling Important in Concurrent Code?

**Challenges:**
1. **Errors can occur in any goroutine**: Need to handle errors from multiple sources
2. **Error propagation**: Errors need to be passed back to the main goroutine
3. **Resource cleanup**: Need to clean up resources when errors occur
4. **Context cancellation**: Need to cancel other operations when errors occur

**Benefits of proper error handling:**
1. **Reliability**: Code continues to work even when errors occur
2. **Debugging**: Easier to identify and fix issues
3. **Resource management**: Proper cleanup of resources
4. **User experience**: Better error messages and recovery

### Error Handling Patterns

1. **Error channels**: Pass errors through channels
2. **Error groups**: Use sync.WaitGroup with error handling
3. **Context cancellation**: Cancel operations on errors
4. **Panic recovery**: Recover from panics in goroutines

## 💻 Code Examples

### Example 1: Basic Error Handling with Channels

```go
package main

import (
    "fmt"
    "time"
)

func basicErrorHandling() {
    fmt.Println("=== Basic Error Handling with Channels ===")
    
    // Create channels for results and errors
    results := make(chan int, 10)
    errors := make(chan error, 10)
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        go func(id int) {
            // Simulate work that might fail
            time.Sleep(time.Duration(id) * 500 * time.Millisecond)
            
            if id == 2 {
                errors <- fmt.Errorf("worker %d failed", id)
                return
            }
            
            results <- id * 10
        }(i)
    }
    
    // Collect results and errors
    for i := 0; i < 3; i++ {
        select {
        case result := <-results:
            fmt.Printf("Result: %d\n", result)
        case err := <-errors:
            fmt.Printf("Error: %v\n", err)
        }
    }
}

func main() {
    basicErrorHandling()
}
```

**Run this code:**
```bash
go run 21-error-handling.go
```

### Example 2: Error Handling with WaitGroup

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func errorHandlingWithWaitGroup() {
    fmt.Println("=== Error Handling with WaitGroup ===")
    
    var wg sync.WaitGroup
    errors := make(chan error, 10)
    
    // Start multiple goroutines
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Simulate work that might fail
            time.Sleep(time.Duration(id) * 200 * time.Millisecond)
            
            if id%3 == 0 {
                errors <- fmt.Errorf("worker %d failed", id)
                return
            }
            
            fmt.Printf("Worker %d completed successfully\n", id)
        }(i)
    }
    
    // Wait for all goroutines to complete
    go func() {
        wg.Wait()
        close(errors)
    }()
    
    // Collect errors
    for err := range errors {
        fmt.Printf("Error: %v\n", err)
    }
}

func main() {
    errorHandlingWithWaitGroup()
}
```

### Example 3: Error Handling with Context

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func errorHandlingWithContext() {
    fmt.Println("=== Error Handling with Context ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Create channels for results and errors
    results := make(chan int, 10)
    errors := make(chan error, 10)
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        go func(id int) {
            select {
            case <-time.After(time.Duration(id) * 800 * time.Millisecond):
                if id == 2 {
                    errors <- fmt.Errorf("worker %d failed", id)
                    return
                }
                results <- id * 10
            case <-ctx.Done():
                errors <- fmt.Errorf("worker %d cancelled: %v", id, ctx.Err())
                return
            }
        }(i)
    }
    
    // Collect results and errors
    for i := 0; i < 3; i++ {
        select {
        case result := <-results:
            fmt.Printf("Result: %d\n", result)
        case err := <-errors:
            fmt.Printf("Error: %v\n", err)
        case <-ctx.Done():
            fmt.Printf("Context cancelled: %v\n", ctx.Err())
            return
        }
    }
}

func main() {
    errorHandlingWithContext()
}
```

### Example 4: Error Handling with Panic Recovery

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func errorHandlingWithPanicRecovery() {
    fmt.Println("=== Error Handling with Panic Recovery ===")
    
    var wg sync.WaitGroup
    errors := make(chan error, 10)
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            defer func() {
                if r := recover(); r != nil {
                    errors <- fmt.Errorf("worker %d panicked: %v", id, r)
                }
            }()
            
            // Simulate work that might panic
            time.Sleep(time.Duration(id) * 300 * time.Millisecond)
            
            if id == 2 {
                panic(fmt.Sprintf("worker %d panicked", id))
            }
            
            fmt.Printf("Worker %d completed successfully\n", id)
        }(i)
    }
    
    // Wait for all goroutines to complete
    go func() {
        wg.Wait()
        close(errors)
    }()
    
    // Collect errors
    for err := range errors {
        fmt.Printf("Error: %v\n", err)
    }
}

func main() {
    errorHandlingWithPanicRecovery()
}
```

### Example 5: Error Handling with Retry

```go
package main

import (
    "fmt"
    "time"
)

func errorHandlingWithRetry() {
    fmt.Println("=== Error Handling with Retry ===")
    
    maxRetries := 3
    baseDelay := 100 * time.Millisecond
    
    for i := 1; i <= 3; i++ {
        fmt.Printf("Attempting operation %d\n", i)
        
        success := retryOperation(i, maxRetries, baseDelay)
        if success {
            fmt.Printf("Operation %d succeeded\n", i)
        } else {
            fmt.Printf("Operation %d failed after %d retries\n", i, maxRetries)
        }
    }
}

func retryOperation(operationID, maxRetries int, baseDelay time.Duration) bool {
    for attempt := 1; attempt <= maxRetries; attempt++ {
        fmt.Printf("  Attempt %d for operation %d\n", attempt, operationID)
        
        if err := performOperation(operationID); err != nil {
            fmt.Printf("  Error: %v\n", err)
            if attempt < maxRetries {
                delay := baseDelay * time.Duration(attempt)
                fmt.Printf("  Retrying in %v...\n", delay)
                time.Sleep(delay)
            }
        } else {
            return true
        }
    }
    return false
}

func performOperation(id int) error {
    time.Sleep(200 * time.Millisecond)
    
    // Simulate failure for certain operations
    if id == 2 {
        return fmt.Errorf("operation %d failed", id)
    }
    
    return nil
}

func main() {
    errorHandlingWithRetry()
}
```

### Example 6: Error Handling with Circuit Breaker

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type CircuitBreaker struct {
    maxFailures int
    timeout     time.Duration
    failures    int
    lastFailure time.Time
    mutex       sync.Mutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        maxFailures: maxFailures,
        timeout:     timeout,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    
    // Check if circuit is open
    if cb.failures >= cb.maxFailures {
        if time.Since(cb.lastFailure) < cb.timeout {
            return fmt.Errorf("circuit breaker is open")
        }
        // Reset failures after timeout
        cb.failures = 0
    }
    
    // Call the function
    err := fn()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        return err
    }
    
    // Reset failures on success
    cb.failures = 0
    return nil
}

func errorHandlingWithCircuitBreaker() {
    fmt.Println("=== Error Handling with Circuit Breaker ===")
    
    cb := NewCircuitBreaker(2, 1*time.Second)
    
    for i := 1; i <= 5; i++ {
        fmt.Printf("Attempt %d: ", i)
        
        err := cb.Call(func() error {
            return performOperation(i)
        })
        
        if err != nil {
            fmt.Printf("Error: %v\n", err)
        } else {
            fmt.Printf("Success\n")
        }
        
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    errorHandlingWithCircuitBreaker()
}
```

### Example 7: Error Handling with Error Groups

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type ErrorGroup struct {
    wg     sync.WaitGroup
    errors chan error
    mutex  sync.Mutex
}

func NewErrorGroup() *ErrorGroup {
    return &ErrorGroup{
        errors: make(chan error, 10),
    }
}

func (eg *ErrorGroup) Go(fn func() error) {
    eg.wg.Add(1)
    go func() {
        defer eg.wg.Done()
        if err := fn(); err != nil {
            eg.errors <- err
        }
    }()
}

func (eg *ErrorGroup) Wait() []error {
    go func() {
        eg.wg.Wait()
        close(eg.errors)
    }()
    
    var errors []error
    for err := range eg.errors {
        errors = append(errors, err)
    }
    return errors
}

func errorHandlingWithErrorGroups() {
    fmt.Println("=== Error Handling with Error Groups ===")
    
    eg := NewErrorGroup()
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        eg.Go(func() error {
            id := i
            time.Sleep(time.Duration(id) * 400 * time.Millisecond)
            
            if id == 2 {
                return fmt.Errorf("worker %d failed", id)
            }
            
            fmt.Printf("Worker %d completed successfully\n", id)
            return nil
        })
    }
    
    // Wait for all goroutines and collect errors
    errors := eg.Wait()
    
    if len(errors) > 0 {
        fmt.Printf("Errors occurred: %v\n", errors)
    } else {
        fmt.Println("All workers completed successfully")
    }
}

func main() {
    errorHandlingWithErrorGroups()
}
```

### Example 8: Error Handling Best Practices

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Worker struct {
    ID    int
    Error chan error
}

func (w *Worker) Process(ctx context.Context, data int) {
    defer func() {
        if r := recover(); r != nil {
            w.Error <- fmt.Errorf("worker %d panicked: %v", w.ID, r)
        }
    }()
    
    // Simulate work
    select {
    case <-time.After(time.Duration(data) * 200 * time.Millisecond):
        if data%3 == 0 {
            w.Error <- fmt.Errorf("worker %d failed processing %d", w.ID, data)
            return
        }
        fmt.Printf("Worker %d processed %d successfully\n", w.ID, data)
    case <-ctx.Done():
        w.Error <- fmt.Errorf("worker %d cancelled: %v", w.ID, ctx.Err())
    }
}

func errorHandlingBestPractices() {
    fmt.Println("=== Error Handling Best Practices ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Create workers
    workers := make([]*Worker, 3)
    for i := range workers {
        workers[i] = &Worker{
            ID:    i + 1,
            Error: make(chan error, 1),
        }
    }
    
    // Start workers
    var wg sync.WaitGroup
    for _, worker := range workers {
        wg.Add(1)
        go func(w *Worker) {
            defer wg.Done()
            for i := 1; i <= 3; i++ {
                w.Process(ctx, i)
            }
        }(worker)
    }
    
    // Collect errors
    go func() {
        wg.Wait()
        for _, worker := range workers {
            close(worker.Error)
        }
    }()
    
    // Process errors
    for _, worker := range workers {
        for err := range worker.Error {
            if err != nil {
                fmt.Printf("Error from worker %d: %v\n", worker.ID, err)
            }
        }
    }
}

func main() {
    errorHandlingBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Error channels**: Pass errors through channels
2. **Error groups**: Use WaitGroup with error handling
3. **Context cancellation**: Cancel operations on errors
4. **Panic recovery**: Recover from panics in goroutines
5. **Retry logic**: Retry operations on failure
6. **Circuit breaker**: Prevent cascading failures

## 🎯 When to Use Different Error Handling Approaches

1. **Error channels**: When you need to collect errors from multiple goroutines
2. **Error groups**: When you need to wait for multiple goroutines and collect errors
3. **Context cancellation**: When you need to cancel operations on errors
4. **Panic recovery**: When you need to handle panics in goroutines
5. **Retry logic**: When operations might fail temporarily
6. **Circuit breaker**: When you need to prevent cascading failures

## 🎯 Best Practices

1. **Always handle errors**:
   ```go
   // GOOD - handle errors
   if err := doSomething(); err != nil {
       // Handle error
   }
   
   // BAD - ignore errors
   doSomething()
   ```

2. **Use error channels**:
   ```go
   // GOOD - use error channels
   errors := make(chan error, 10)
   go func() {
       if err := doSomething(); err != nil {
           errors <- err
       }
   }()
   ```

3. **Use context for cancellation**:
   ```go
   // GOOD - use context
   select {
   case result := <-results:
       // Process result
   case <-ctx.Done():
       return ctx.Err()
   }
   ```

4. **Recover from panics**:
   ```go
   // GOOD - recover from panics
   defer func() {
       if r := recover(); r != nil {
           // Handle panic
       }
   }()
   ```

## 🎯 Common Pitfalls

1. **Not handling errors**:
   ```go
   // BAD - ignore errors
   go func() {
       doSomething() // Error is lost
   }()
   
   // GOOD - handle errors
   go func() {
       if err := doSomething(); err != nil {
           // Handle error
       }
   }()
   ```

2. **Not closing error channels**:
   ```go
   // BAD - error channel not closed
   errors := make(chan error)
   go func() {
       errors <- err
   }()
   
   // GOOD - close error channel
   errors := make(chan error)
   go func() {
       defer close(errors)
       errors <- err
   }()
   ```

3. **Not using context**:
   ```go
   // BAD - no cancellation
   go func() {
       for {
           // Long operation
       }
   }()
   
   // GOOD - use context
   go func() {
       for {
           select {
           case <-ctx.Done():
               return
           default:
               // Long operation
           }
       }
   }()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a distributed system with error handling:
- Multiple services process requests
- Each service can fail with different error types
- Implement proper error handling and recovery
- Use context for cancellation
- Show how to handle different types of errors

**Hint**: Use error channels to collect errors from different services and implement proper error handling and recovery.

## 🚀 Next Steps

Now that you understand error handling in concurrent code, let's learn about **graceful shutdown** in the next file: `22-graceful-shutdown.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
