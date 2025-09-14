# 10 - Timeout Patterns

## 🎯 Learning Objectives
- Understand why timeouts are important in concurrent programming
- Learn different timeout patterns and when to use them
- Master the `time.After` function
- Practice with timeout scenarios
- Learn about timeout best practices

## 📚 Theory

### Why Do We Need Timeouts?

**Problems without timeouts:**
1. **Deadlocks**: Goroutines can block forever
2. **Resource leaks**: Goroutines never exit
3. **Poor user experience**: Operations hang indefinitely
4. **System instability**: Unresponsive applications

**Benefits of timeouts:**
1. **Prevent deadlocks**: Operations fail fast
2. **Resource management**: Clean up blocked goroutines
3. **Better UX**: Users get feedback quickly
4. **System reliability**: Graceful degradation

### Timeout Patterns

1. **Simple timeout**: Wait with a timeout
2. **Timeout with retry**: Retry on timeout
3. **Timeout with fallback**: Use alternative when timeout
4. **Timeout with cancellation**: Cancel ongoing operations

## 💻 Code Examples

### Example 1: Basic Timeout Pattern

```go
package main

import (
    "fmt"
    "time"
)

func basicTimeout() {
    fmt.Println("=== Basic Timeout Pattern ===")
    
    ch := make(chan string)
    
    // Simulate slow operation
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "Slow operation completed"
    }()
    
    // Wait with timeout
    select {
    case result := <-ch:
        fmt.Printf("Received: %s\n", result)
    case <-time.After(2 * time.Second):
        fmt.Println("Operation timed out")
    }
}

func main() {
    basicTimeout()
}
```

**Run this code:**
```bash
go run 10-timeout-patterns.go
```

### Example 2: Timeout with Retry

```go
package main

import (
    "fmt"
    "time"
)

func timeoutWithRetry() {
    fmt.Println("=== Timeout with Retry ===")
    
    maxRetries := 3
    timeout := 1 * time.Second
    
    for attempt := 1; attempt <= maxRetries; attempt++ {
        fmt.Printf("Attempt %d: ", attempt)
        
        ch := make(chan string)
        
        // Simulate operation that might fail
        go func() {
            time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
            ch <- "Operation completed"
        }()
        
        select {
        case result := <-ch:
            fmt.Printf("Success: %s\n", result)
            return
        case <-time.After(timeout):
            fmt.Println("Timeout")
            if attempt == maxRetries {
                fmt.Println("Max retries reached")
            }
        }
    }
}

func main() {
    timeoutWithRetry()
}
```

### Example 3: Timeout with Fallback

```go
package main

import (
    "fmt"
    "time"
)

func timeoutWithFallback() {
    fmt.Println("=== Timeout with Fallback ===")
    
    ch := make(chan string)
    
    // Simulate slow operation
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "Primary operation completed"
    }()
    
    select {
    case result := <-ch:
        fmt.Printf("Primary result: %s\n", result)
    case <-time.After(2 * time.Second):
        fmt.Println("Primary operation timed out, using fallback")
        fallbackResult := "Fallback operation completed"
        fmt.Printf("Fallback result: %s\n", fallbackResult)
    }
}

func main() {
    timeoutWithFallback()
}
```

### Example 4: Multiple Timeouts

```go
package main

import (
    "fmt"
    "time"
)

func multipleTimeouts() {
    fmt.Println("=== Multiple Timeouts ===")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start two operations with different timeouts
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Operation 1 completed"
    }()
    
    go func() {
        time.Sleep(3 * time.Second)
        ch2 <- "Operation 2 completed"
    }()
    
    // Wait for either operation or timeout
    select {
    case result1 := <-ch1:
        fmt.Printf("Received from ch1: %s\n", result1)
    case result2 := <-ch2:
        fmt.Printf("Received from ch2: %s\n", result2)
    case <-time.After(2 * time.Second):
        fmt.Println("Both operations timed out")
    }
}

func main() {
    multipleTimeouts()
}
```

### Example 5: Timeout with Context

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func timeoutWithContext() {
    fmt.Println("=== Timeout with Context ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    ch := make(chan string)
    
    // Simulate operation
    go func() {
        time.Sleep(3 * time.Second)
        select {
        case ch <- "Operation completed":
        case <-ctx.Done():
            fmt.Println("Operation cancelled")
            return
        }
    }()
    
    select {
    case result := <-ch:
        fmt.Printf("Received: %s\n", result)
    case <-ctx.Done():
        fmt.Printf("Context cancelled: %v\n", ctx.Err())
    }
}

func main() {
    timeoutWithContext()
}
```

### Example 6: Timeout with Heartbeat

```go
package main

import (
    "fmt"
    "time"
)

func timeoutWithHeartbeat() {
    fmt.Println("=== Timeout with Heartbeat ===")
    
    ch := make(chan string)
    heartbeat := time.NewTicker(500 * time.Millisecond)
    defer heartbeat.Stop()
    
    // Simulate long-running operation
    go func() {
        for i := 1; i <= 10; i++ {
            time.Sleep(300 * time.Millisecond)
            if i == 5 {
                ch <- "Operation completed"
                return
            }
        }
    }()
    
    timeout := time.After(3 * time.Second)
    
    for {
        select {
        case result := <-ch:
            fmt.Printf("Operation completed: %s\n", result)
            return
        case <-heartbeat.C:
            fmt.Println("Heartbeat...")
        case <-timeout:
            fmt.Println("Operation timed out")
            return
        }
    }
}

func main() {
    timeoutWithHeartbeat()
}
```

### Example 7: Timeout with Graceful Shutdown

```go
package main

import (
    "fmt"
    "time"
)

func timeoutWithGracefulShutdown() {
    fmt.Println("=== Timeout with Graceful Shutdown ===")
    
    ch := make(chan string)
    done := make(chan bool)
    
    // Long-running operation
    go func() {
        for i := 1; i <= 10; i++ {
            select {
            case <-done:
                fmt.Println("Operation cancelled")
                return
            default:
                fmt.Printf("Working... step %d\n", i)
                time.Sleep(500 * time.Millisecond)
            }
        }
        ch <- "Operation completed"
    }()
    
    // Wait with timeout
    select {
    case result := <-ch:
        fmt.Printf("Received: %s\n", result)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout reached, shutting down gracefully")
        close(done) // Signal cancellation
        time.Sleep(100 * time.Millisecond) // Give time to clean up
    }
}

func main() {
    timeoutWithGracefulShutdown()
}
```

### Example 8: Timeout Best Practices

```go
package main

import (
    "fmt"
    "time"
)

func timeoutBestPractices() {
    fmt.Println("=== Timeout Best Practices ===")
    
    // Practice 1: Use appropriate timeout duration
    shortTimeout := 100 * time.Millisecond
    longTimeout := 5 * time.Second
    
    // Practice 2: Timeout with cleanup
    ch := make(chan string)
    
    go func() {
        defer func() {
            fmt.Println("Cleaning up resources")
        }()
        
        time.Sleep(2 * time.Second)
        ch <- "Operation completed"
    }()
    
    select {
    case result := <-ch:
        fmt.Printf("Received: %s\n", result)
    case <-time.After(shortTimeout):
        fmt.Println("Short timeout reached")
    case <-time.After(longTimeout):
        fmt.Println("Long timeout reached")
    }
    
    // Practice 3: Timeout with error handling
    fmt.Println("\n=== Timeout with Error Handling ===")
    
    result, err := performOperationWithTimeout(1 * time.Second)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %s\n", result)
    }
}

func performOperationWithTimeout(timeout time.Duration) (string, error) {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Operation completed"
    }()
    
    select {
    case result := <-ch:
        return result, nil
    case <-time.After(timeout):
        return "", fmt.Errorf("operation timed out after %v", timeout)
    }
}

func main() {
    timeoutBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **`time.After(duration)`**: Creates a channel that sends after duration
2. **Select with timeout**: Use select to wait with timeout
3. **Timeout patterns**: Simple, retry, fallback, cancellation
4. **Context timeouts**: Use context for more complex timeout scenarios
5. **Heartbeat pattern**: Regular status checks with timeout
6. **Graceful shutdown**: Clean up resources on timeout

## 🎯 Common Use Cases

1. **Network operations**: HTTP requests, database queries
2. **File operations**: Reading/writing files
3. **External services**: API calls, third-party integrations
4. **User interactions**: UI timeouts, user input
5. **System operations**: Process management, resource allocation

## 🎯 Best Practices

1. **Choose appropriate timeout duration**:
   ```go
   // BAD - too short
   timeout := 1 * time.Millisecond
   
   // GOOD - reasonable timeout
   timeout := 5 * time.Second
   ```

2. **Handle timeout errors**:
   ```go
   select {
   case result := <-ch:
       return result, nil
   case <-time.After(timeout):
       return "", fmt.Errorf("timeout after %v", timeout)
   }
   ```

3. **Clean up resources on timeout**:
   ```go
   defer func() {
       // Cleanup code
   }()
   ```

4. **Use context for complex timeouts**:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), timeout)
   defer cancel()
   ```

## 🎯 Common Pitfalls

1. **Too short timeout**:
   ```go
   // BAD - might timeout too early
   timeout := 1 * time.Millisecond
   
   // GOOD - reasonable timeout
   timeout := 1 * time.Second
   ```

2. **Not handling timeout errors**:
   ```go
   // BAD - ignore timeout
   select {
   case result := <-ch:
       // Handle result
   case <-time.After(timeout):
       // Do nothing
   }
   
   // GOOD - handle timeout
   select {
   case result := <-ch:
       // Handle result
   case <-time.After(timeout):
       return fmt.Errorf("timeout")
   }
   ```

3. **Resource leaks on timeout**:
   ```go
   // BAD - resources not cleaned up
   go func() {
       // Long operation
   }()
   
   // GOOD - cleanup on timeout
   go func() {
       defer cleanup()
       // Long operation
   }()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a distributed system:
- Multiple services send heartbeats
- If a service doesn't send a heartbeat within 2 seconds, it's considered down
- Implement timeout detection for each service
- Show how to handle service failures gracefully

**Hint**: Use a map to track last heartbeat time for each service and use timeouts to detect failures.

## 🚀 Next Steps

Now that you understand timeout patterns, let's learn about **race conditions** in the next file: `11-race-conditions.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
