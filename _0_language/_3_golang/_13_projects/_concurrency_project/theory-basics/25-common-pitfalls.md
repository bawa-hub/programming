# 25 - Common Pitfalls

## 🎯 Learning Objectives
- Understand common pitfalls in Go concurrency
- Learn how to avoid and fix these pitfalls
- Master debugging techniques for concurrent code
- Practice with pitfall prevention strategies
- Understand when to use different approaches

## 📚 Theory

### Why Do Pitfalls Occur?

**Common reasons:**
1. **Misunderstanding concurrency**: Not understanding how goroutines work
2. **Race conditions**: Not protecting shared data
3. **Deadlocks**: Circular dependencies
4. **Resource leaks**: Goroutines that never exit
5. **Poor patterns**: Using wrong concurrency patterns

**Impact of pitfalls:**
1. **Bugs**: Non-deterministic behavior
2. **Performance issues**: Poor performance
3. **Resource leaks**: Memory and goroutine leaks
4. **System instability**: Crashes and hangs
5. **Hard to debug**: Difficult to reproduce and fix

### Common Pitfall Categories

1. **Race conditions**: Data races and memory corruption
2. **Deadlocks**: Circular waiting
3. **Goroutine leaks**: Goroutines that never exit
4. **Channel misuse**: Incorrect channel usage
5. **Synchronization issues**: Wrong synchronization patterns

## 💻 Code Examples

### Example 1: Race Conditions

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func raceConditions() {
    fmt.Println("=== Race Conditions ===")
    
    // BAD: Race condition
    fmt.Println("Bad approach - race condition:")
    var counter int
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // Race condition!
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d (unpredictable)\n", counter)
    
    // GOOD: Use mutex
    fmt.Println("\nGood approach - use mutex:")
    counter = 0
    var mutex sync.Mutex
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mutex.Lock()
            counter++
            mutex.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d (should be 1000)\n", counter)
}

func main() {
    raceConditions()
}
```

**Run this code:**
```bash
go run 25-common-pitfalls.go
```

### Example 2: Deadlocks

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func deadlocks() {
    fmt.Println("=== Deadlocks ===")
    
    // BAD: Deadlock
    fmt.Println("Bad approach - deadlock:")
    var mutex1, mutex2 sync.Mutex
    
    go func() {
        mutex1.Lock()
        fmt.Println("Goroutine 1: acquired mutex1")
        time.Sleep(100 * time.Millisecond)
        mutex2.Lock()
        fmt.Println("Goroutine 1: acquired mutex2")
        mutex2.Unlock()
        mutex1.Unlock()
    }()
    
    go func() {
        mutex2.Lock()
        fmt.Println("Goroutine 2: acquired mutex2")
        time.Sleep(100 * time.Millisecond)
        mutex1.Lock()
        fmt.Println("Goroutine 2: acquired mutex1")
        mutex1.Unlock()
        mutex2.Unlock()
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Deadlock example completed")
    
    // GOOD: Same order
    fmt.Println("\nGood approach - same order:")
    mutex1 = sync.Mutex{}
    mutex2 = sync.Mutex{}
    
    go func() {
        mutex1.Lock()
        fmt.Println("Goroutine 1: acquired mutex1")
        time.Sleep(100 * time.Millisecond)
        mutex2.Lock()
        fmt.Println("Goroutine 1: acquired mutex2")
        mutex2.Unlock()
        mutex1.Unlock()
    }()
    
    go func() {
        mutex1.Lock()
        fmt.Println("Goroutine 2: acquired mutex1")
        time.Sleep(100 * time.Millisecond)
        mutex2.Lock()
        fmt.Println("Goroutine 2: acquired mutex2")
        mutex2.Unlock()
        mutex1.Unlock()
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("No deadlock example completed")
}

func main() {
    deadlocks()
}
```

### Example 3: Goroutine Leaks

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func goroutineLeaks() {
    fmt.Println("=== Goroutine Leaks ===")
    
    // BAD: Goroutine leak
    fmt.Println("Bad approach - goroutine leak:")
    fmt.Printf("Goroutines before: %d\n", runtime.NumGoroutine())
    
    for i := 0; i < 5; i++ {
        go func(id int) {
            for {
                fmt.Printf("Goroutine %d running\n", id)
                time.Sleep(500 * time.Millisecond)
                // No exit condition - leak!
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutines after: %d\n", runtime.NumGoroutine())
    
    // GOOD: Proper exit condition
    fmt.Println("\nGood approach - proper exit condition:")
    fmt.Printf("Goroutines before: %d\n", runtime.NumGoroutine())
    
    done := make(chan bool)
    
    for i := 0; i < 5; i++ {
        go func(id int) {
            for {
                select {
                case <-done:
                    fmt.Printf("Goroutine %d exiting\n", id)
                    return
                default:
                    fmt.Printf("Goroutine %d running\n", id)
                    time.Sleep(500 * time.Millisecond)
                }
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    close(done) // Signal all goroutines to exit
    time.Sleep(500 * time.Millisecond)
    fmt.Printf("Goroutines after: %d\n", runtime.NumGoroutine())
}

func main() {
    goroutineLeaks()
}
```

### Example 4: Channel Misuse

```go
package main

import (
    "fmt"
    "time"
)

func channelMisuse() {
    fmt.Println("=== Channel Misuse ===")
    
    // BAD: Sending to closed channel
    fmt.Println("Bad approach - sending to closed channel:")
    ch := make(chan int)
    close(ch)
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic recovered: %v\n", r)
        }
    }()
    
    ch <- 42 // This will panic!
    
    // GOOD: Don't send to closed channel
    fmt.Println("\nGood approach - don't send to closed channel:")
    ch = make(chan int)
    
    go func() {
        ch <- 42
        close(ch)
    }()
    
    value := <-ch
    fmt.Printf("Received: %d\n", value)
    
    // BAD: Receiving from nil channel
    fmt.Println("\nBad approach - receiving from nil channel:")
    var nilCh chan int
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        fmt.Println("This will never print - nil channel blocks forever")
    }()
    
    // This will block forever
    // value := <-nilCh
    
    // GOOD: Use proper channel
    fmt.Println("Good approach - use proper channel:")
    ch = make(chan int)
    
    go func() {
        ch <- 42
    }()
    
    value = <-ch
    fmt.Printf("Received: %d\n", value)
}

func main() {
    channelMisuse()
}
```

### Example 5: Synchronization Issues

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func synchronizationIssues() {
    fmt.Println("=== Synchronization Issues ===")
    
    // BAD: Not waiting for goroutines
    fmt.Println("Bad approach - not waiting for goroutines:")
    for i := 0; i < 3; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d working\n", id)
            time.Sleep(500 * time.Millisecond)
            fmt.Printf("Goroutine %d done\n", id)
        }(i)
    }
    fmt.Println("Main function continues immediately")
    
    // GOOD: Use WaitGroup
    fmt.Println("\nGood approach - use WaitGroup:")
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d working\n", id)
            time.Sleep(500 * time.Millisecond)
            fmt.Printf("Goroutine %d done\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All goroutines completed")
}

func main() {
    synchronizationIssues()
}
```

### Example 6: Context Misuse

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func contextMisuse() {
    fmt.Println("=== Context Misuse ===")
    
    // BAD: Not checking context
    fmt.Println("Bad approach - not checking context:")
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Printf("Working... %d\n", i)
            time.Sleep(500 * time.Millisecond)
            // Not checking context - will continue after timeout
        }
    }()
    
    time.Sleep(2 * time.Second)
    
    // GOOD: Check context
    fmt.Println("\nGood approach - check context:")
    ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    go func() {
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Cancelled: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    contextMisuse()
}
```

### Example 7: Memory Leaks

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func memoryLeaks() {
    fmt.Println("=== Memory Leaks ===")
    
    // BAD: Memory leak
    fmt.Println("Bad approach - memory leak:")
    fmt.Printf("Memory before: %d bytes\n", getMemStats().Alloc)
    
    for i := 0; i < 1000; i++ {
        go func() {
            // Allocate memory that's never freed
            data := make([]byte, 1024)
            _ = data
            time.Sleep(1 * time.Second)
        }()
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Memory after: %d bytes\n", getMemStats().Alloc)
    
    // GOOD: Proper cleanup
    fmt.Println("\nGood approach - proper cleanup:")
    fmt.Printf("Memory before: %d bytes\n", getMemStats().Alloc)
    
    for i := 0; i < 1000; i++ {
        go func() {
            defer func() {
                // Cleanup code here
            }()
            
            // Allocate memory
            data := make([]byte, 1024)
            _ = data
            time.Sleep(1 * time.Second)
        }()
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Memory after: %d bytes\n", getMemStats().Alloc)
}

func getMemStats() runtime.MemStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m
}

func main() {
    memoryLeaks()
}
```

### Example 8: Pitfall Prevention Best Practices

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type SafeWorker struct {
    id      int
    ctx     context.Context
    wg      *sync.WaitGroup
    results chan<- int
    mutex   *sync.Mutex
    counter *int64
}

func (w *SafeWorker) Process() {
    defer w.wg.Done()
    
    for i := 0; i < 10; i++ {
        select {
        case <-w.ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", w.id, w.ctx.Err())
            return
        default:
            // Process work
            time.Sleep(100 * time.Millisecond)
            
            // Safely update shared counter
            w.mutex.Lock()
            *w.counter++
            w.mutex.Unlock()
            
            // Send result
            select {
            case w.results <- w.id * i:
            case <-w.ctx.Done():
                return
            }
        }
    }
    
    fmt.Printf("Worker %d completed\n", w.id)
}

func pitfallPreventionBestPractices() {
    fmt.Println("=== Pitfall Prevention Best Practices ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Create shared resources
    var counter int64
    var mutex sync.Mutex
    results := make(chan int, 100)
    
    // Create workers
    var wg sync.WaitGroup
    numWorkers := 3
    
    for i := 0; i < numWorkers; i++ {
        worker := &SafeWorker{
            id:      i,
            ctx:     ctx,
            wg:      &wg,
            results: results,
            mutex:   &mutex,
            counter: &counter,
        }
        
        wg.Add(1)
        go worker.Process()
    }
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
    
    fmt.Printf("Final counter: %d\n", counter)
}

func main() {
    pitfallPreventionBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Race conditions**: Protect shared data with mutexes
2. **Deadlocks**: Avoid circular dependencies
3. **Goroutine leaks**: Always have exit conditions
4. **Channel misuse**: Don't send to closed channels
5. **Synchronization issues**: Use WaitGroup properly
6. **Context misuse**: Always check context
7. **Memory leaks**: Clean up resources properly

## 🎯 How to Avoid Pitfalls

1. **Use race detector**: Always run with `go run -race`
2. **Use WaitGroup**: Wait for goroutines to complete
3. **Use context**: Handle cancellation properly
4. **Check channels**: Verify channel state
5. **Clean up resources**: Use defer for cleanup
6. **Test thoroughly**: Test concurrent code multiple times

## 🎯 Best Practices

1. **Always use race detector**:
   ```go
   // GOOD - use race detector
   go run -race main.go
   
   // BAD - don't use race detector
   go run main.go
   ```

2. **Use WaitGroup properly**:
   ```go
   // GOOD - use WaitGroup
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
       defer wg.Done()
       // Work
   }()
   wg.Wait()
   
   // BAD - don't wait
   go func() {
       // Work
   }()
   ```

3. **Check context**:
   ```go
   // GOOD - check context
   select {
   case <-ctx.Done():
       return
   default:
       // Work
   }
   
   // BAD - don't check context
   // Work
   ```

4. **Clean up resources**:
   ```go
   // GOOD - clean up resources
   defer func() {
       // Cleanup code
   }()
   
   // BAD - don't clean up
   // No cleanup
   ```

## 🎯 Common Pitfalls

1. **Race conditions**:
   ```go
   // BAD - race condition
   var counter int
   go func() {
       counter++
   }()
   
   // GOOD - use mutex
   var counter int
   var mutex sync.Mutex
   go func() {
       mutex.Lock()
       counter++
       mutex.Unlock()
   }()
   ```

2. **Deadlocks**:
   ```go
   // BAD - deadlock
   mutex1.Lock()
   mutex2.Lock()
   // ...
   mutex2.Unlock()
   mutex1.Unlock()
   
   // GOOD - same order
   mutex1.Lock()
   mutex2.Lock()
   // ...
   mutex1.Unlock()
   mutex2.Unlock()
   ```

3. **Goroutine leaks**:
   ```go
   // BAD - goroutine leak
   go func() {
       for {
           // No exit condition
       }
   }()
   
   // GOOD - exit condition
   go func() {
       for {
           select {
           case <-done:
               return
           default:
               // Work
           }
       }
   }()
   ```

## 🏋️ Exercise

**Problem**: Create a program that demonstrates and fixes common concurrency pitfalls:
- Show examples of race conditions, deadlocks, and goroutine leaks
- Implement proper fixes for each pitfall
- Use best practices to prevent pitfalls
- Show how to debug and test concurrent code

**Hint**: Create examples of each pitfall and then show how to fix them using proper concurrency patterns.

## 🚀 Next Steps

Congratulations! You've completed the Go concurrency theory basics. You now have a comprehensive understanding of:

- **Basic concepts**: Goroutines, channels, select statements
- **Synchronization**: Mutexes, WaitGroup, atomic operations
- **Advanced patterns**: Worker pools, pipelines, fan-in/fan-out
- **Error handling**: Proper error handling in concurrent code
- **Performance**: Optimization techniques and best practices
- **Monitoring**: How to monitor and debug concurrent applications
- **Pitfalls**: Common mistakes and how to avoid them

You're now ready to build production-ready concurrent applications in Go! 🎉

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
