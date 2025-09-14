# 17 - Context Package

## 🎯 Learning Objectives
- Understand what the context package is and why it's important
- Learn how to use context for cancellation and timeouts
- Master different context types and their use cases
- Practice with context patterns and best practices
- Understand context propagation and inheritance

## 📚 Theory

### What is the Context Package?

The **context** package provides a way to carry request-scoped values, cancellation signals, and timeouts across API boundaries and between processes.

**Key characteristics:**
- **Cancellation**: Can signal when operations should stop
- **Timeouts**: Can set deadlines for operations
- **Values**: Can carry request-scoped values
- **Immutable**: Contexts are immutable and thread-safe

### Why Do We Need Context?

**Problems without context:**
1. **No cancellation**: Operations can't be stopped
2. **No timeouts**: Operations can hang forever
3. **No request tracing**: Can't track requests across goroutines
4. **Resource leaks**: Goroutines never exit

**Benefits of context:**
1. **Cancellation**: Stop operations when needed
2. **Timeouts**: Prevent operations from hanging
3. **Request tracing**: Track requests across goroutines
4. **Resource management**: Clean up resources properly

### Context Types

1. **Background**: Empty context, never cancelled
2. **TODO**: Empty context, should be replaced
3. **WithCancel**: Can be cancelled
4. **WithTimeout**: Cancelled after timeout
5. **WithDeadline**: Cancelled at specific time
6. **WithValue**: Carries key-value pairs

## 💻 Code Examples

### Example 1: Basic Context Usage

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func basicContext() {
    fmt.Println("=== Basic Context Usage ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Simulate work
    go func() {
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("Work cancelled")
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(500 * time.Millisecond)
            }
        }
        fmt.Println("Work completed")
    }()
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}

func main() {
    basicContext()
}
```

**Run this code:**
```bash
go run 17-context-package.go
```

### Example 2: Context with Cancellation

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func contextWithCancellation() {
    fmt.Println("=== Context with Cancellation ===")
    
    // Create context that can be cancelled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Start worker
    go func() {
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Worker cancelled: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Worker: %d\n", i)
                time.Sleep(300 * time.Millisecond)
            }
        }
    }()
    
    // Cancel after 1 second
    time.Sleep(1 * time.Second)
    fmt.Println("Cancelling context...")
    cancel()
    
    time.Sleep(500 * time.Millisecond)
}

func main() {
    contextWithCancellation()
}
```

### Example 3: Context with Timeout

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func contextWithTimeout() {
    fmt.Println("=== Context with Timeout ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Start worker
    go func() {
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Worker timed out: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Worker: %d\n", i)
                time.Sleep(300 * time.Millisecond)
            }
        }
    }()
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}

func main() {
    contextWithTimeout()
}
```

### Example 4: Context with Deadline

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func contextWithDeadline() {
    fmt.Println("=== Context with Deadline ===")
    
    // Create context with deadline
    deadline := time.Now().Add(2 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    
    // Start worker
    go func() {
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Worker deadline reached: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Worker: %d\n", i)
                time.Sleep(400 * time.Millisecond)
            }
        }
    }()
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}

func main() {
    contextWithDeadline()
}
```

### Example 5: Context with Values

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func contextWithValues() {
    fmt.Println("=== Context with Values ===")
    
    // Create context with values
    ctx := context.WithValue(context.Background(), "userID", "12345")
    ctx = context.WithValue(ctx, "requestID", "req-001")
    
    // Start worker
    go func() {
        for i := 0; i < 5; i++ {
            select {
            case <-ctx.Done():
                return
            default:
                userID := ctx.Value("userID")
                requestID := ctx.Value("requestID")
                fmt.Printf("Worker %d: userID=%s, requestID=%s\n", i, userID, requestID)
                time.Sleep(300 * time.Millisecond)
            }
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    contextWithValues()
}
```

### Example 6: Context Propagation

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker1(ctx context.Context) {
    fmt.Println("Worker 1 starting")
    defer fmt.Println("Worker 1 finished")
    
    // Pass context to worker2
    go worker2(ctx)
    
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker 1 cancelled: %v\n", ctx.Err())
            return
        default:
            fmt.Printf("Worker 1: %d\n", i)
            time.Sleep(300 * time.Millisecond)
        }
    }
}

func worker2(ctx context.Context) {
    fmt.Println("Worker 2 starting")
    defer fmt.Println("Worker 2 finished")
    
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker 2 cancelled: %v\n", ctx.Err())
            return
        default:
            fmt.Printf("Worker 2: %d\n", i)
            time.Sleep(400 * time.Millisecond)
        }
    }
}

func contextPropagation() {
    fmt.Println("=== Context Propagation ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Start worker1
    go worker1(ctx)
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}

func main() {
    contextPropagation()
}
```

### Example 7: Context with HTTP-like Operations

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func simulateHTTPRequest(ctx context.Context, url string) error {
    fmt.Printf("Making request to %s\n", url)
    
    // Simulate network delay
    select {
    case <-time.After(1 * time.Second):
        fmt.Printf("Request to %s completed\n", url)
        return nil
    case <-ctx.Done():
        fmt.Printf("Request to %s cancelled: %v\n", url, ctx.Err())
        return ctx.Err()
    }
}

func contextWithHTTPOperations() {
    fmt.Println("=== Context with HTTP-like Operations ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // Make multiple requests
    urls := []string{"https://api1.com", "https://api2.com", "https://api3.com"}
    
    for _, url := range urls {
        go func(u string) {
            if err := simulateHTTPRequest(ctx, u); err != nil {
                fmt.Printf("Error: %v\n", err)
            }
        }(url)
    }
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("All requests done: %v\n", ctx.Err())
}

func main() {
    contextWithHTTPOperations()
}
```

### Example 8: Context Best Practices

```go
package main

import (
    "context"
    "fmt"
    "time"
)

type Service struct {
    name string
}

func (s *Service) Process(ctx context.Context, data string) error {
    fmt.Printf("Service %s processing: %s\n", s.name, data)
    
    // Check if context is cancelled
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
    
    // Simulate work
    time.Sleep(500 * time.Millisecond)
    
    // Check again before finishing
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        fmt.Printf("Service %s completed: %s\n", s.name, data)
        return nil
    }
}

func contextBestPractices() {
    fmt.Println("=== Context Best Practices ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Create services
    service1 := &Service{name: "Service1"}
    service2 := &Service{name: "Service2"}
    
    // Process data with services
    go func() {
        if err := service1.Process(ctx, "data1"); err != nil {
            fmt.Printf("Service1 error: %v\n", err)
        }
    }()
    
    go func() {
        if err := service2.Process(ctx, "data2"); err != nil {
            fmt.Printf("Service2 error: %v\n", err)
        }
    }()
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("All services done: %v\n", ctx.Err())
}

func main() {
    contextBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Context is immutable**: Cannot be modified after creation
2. **Context is thread-safe**: Can be used by multiple goroutines
3. **Context carries values**: Can store request-scoped data
4. **Context provides cancellation**: Can signal when to stop
5. **Context provides timeouts**: Can set deadlines

## 🎯 When to Use Context

1. **HTTP requests**: Pass context through request handlers
2. **Database operations**: Set timeouts for queries
3. **External API calls**: Handle cancellations and timeouts
4. **Long-running operations**: Allow cancellation
5. **Request tracing**: Pass request ID through call chain

## 🎯 Best Practices

1. **Always pass context as first parameter**:
   ```go
   func Process(ctx context.Context, data string) error {
       // Process data
   }
   ```

2. **Check context cancellation**:
   ```go
   select {
   case <-ctx.Done():
       return ctx.Err()
   default:
       // Continue processing
   }
   ```

3. **Don't store context in structs**:
   ```go
   // BAD - storing context in struct
   type Service struct {
       ctx context.Context
   }
   
   // GOOD - pass context as parameter
   func (s *Service) Process(ctx context.Context) error {
       // Process
   }
   ```

4. **Use context.WithTimeout for timeouts**:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   ```

## 🎯 Common Pitfalls

1. **Not checking context cancellation**:
   ```go
   // BAD - doesn't check for cancellation
   func Process(ctx context.Context) {
       for i := 0; i < 1000; i++ {
           // Long operation
       }
   }
   
   // GOOD - checks for cancellation
   func Process(ctx context.Context) {
       for i := 0; i < 1000; i++ {
           select {
           case <-ctx.Done():
               return
           default:
               // Long operation
           }
       }
   }
   ```

2. **Storing context in structs**:
   ```go
   // BAD - storing context in struct
   type Service struct {
       ctx context.Context
   }
   
   // GOOD - pass context as parameter
   func (s *Service) Process(ctx context.Context) error {
       // Process
   }
   ```

3. **Not calling cancel**:
   ```go
   // BAD - not calling cancel
   ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
   
   // GOOD - call cancel
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a web scraper:
- Use context to handle timeouts and cancellation
- Multiple goroutines scrape different URLs
- Show how context can be used to cancel all operations
- Implement proper error handling with context

**Hint**: Use context.WithTimeout to set a deadline for all scraping operations and check for cancellation in each goroutine.

## 🚀 Next Steps

Now that you understand the context package, let's learn about **worker pools** in the next file: `18-worker-pools.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
