# 14 - WaitGroup

## 🎯 Learning Objectives
- Understand what WaitGroup is and why it's needed
- Learn how to use `sync.WaitGroup` to coordinate goroutines
- Master WaitGroup operations (Add, Done, Wait)
- Practice with different WaitGroup patterns
- Understand WaitGroup best practices

## 📚 Theory

### What is WaitGroup?

**WaitGroup** is a synchronization primitive that waits for a collection of goroutines to finish executing.

**Key characteristics:**
- **Counter-based**: Tracks number of goroutines
- **Blocking wait**: Main goroutine waits until counter reaches zero
- **Thread-safe**: Can be used by multiple goroutines
- **One-time use**: Cannot be reused after Wait returns

### Why Do We Need WaitGroup?

**Problem**: Main goroutine might exit before worker goroutines finish.

**Solution**: WaitGroup allows main goroutine to wait for all workers to complete.

### WaitGroup Operations

1. **Add**: `wg.Add(n)` - Increment counter by n
2. **Done**: `wg.Done()` - Decrement counter by 1
3. **Wait**: `wg.Wait()` - Block until counter reaches zero

## 💻 Code Examples

### Example 1: Basic WaitGroup Usage

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicWaitGroup() {
    fmt.Println("=== Basic WaitGroup Usage ===")
    
    var wg sync.WaitGroup
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment counter
        go func(id int) {
            defer wg.Done() // Decrement counter when done
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Duration(id) * 500 * time.Millisecond)
            fmt.Printf("Worker %d finished\n", id)
        }(i)
    }
    
    fmt.Println("Waiting for all workers to complete...")
    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All workers completed!")
}

func main() {
    basicWaitGroup()
}
```

**Run this code:**
```bash
go run 14-waitgroup.go
```

### Example 2: WaitGroup with Different Workloads

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupWithDifferentWorkloads() {
    fmt.Println("=== WaitGroup with Different Workloads ===")
    
    var wg sync.WaitGroup
    
    // Worker 1: Quick task
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Quick worker starting...")
        time.Sleep(500 * time.Millisecond)
        fmt.Println("Quick worker finished")
    }()
    
    // Worker 2: Medium task
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Medium worker starting...")
        time.Sleep(1 * time.Second)
        fmt.Println("Medium worker finished")
    }()
    
    // Worker 3: Long task
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Long worker starting...")
        time.Sleep(2 * time.Second)
        fmt.Println("Long worker finished")
    }()
    
    fmt.Println("Waiting for all workers...")
    wg.Wait()
    fmt.Println("All workers completed!")
}

func main() {
    waitGroupWithDifferentWorkloads()
}
```

### Example 3: WaitGroup with Results Collection

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupWithResults() {
    fmt.Println("=== WaitGroup with Results Collection ===")
    
    var wg sync.WaitGroup
    results := make([]int, 0)
    var mutex sync.Mutex
    
    // Start multiple workers
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Simulate work
            result := id * id
            fmt.Printf("Worker %d: computed %d\n", id, result)
            
            // Safely append result
            mutex.Lock()
            results = append(results, result)
            mutex.Unlock()
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("All results: %v\n", results)
}

func main() {
    waitGroupWithResults()
}
```

### Example 4: WaitGroup with Error Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupWithErrorHandling() {
    fmt.Println("=== WaitGroup with Error Handling ===")
    
    var wg sync.WaitGroup
    errors := make([]error, 0)
    var mutex sync.Mutex
    
    // Start multiple workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Simulate work that might fail
            if id == 2 {
                err := fmt.Errorf("worker %d failed", id)
                fmt.Printf("Worker %d: %v\n", id, err)
                
                mutex.Lock()
                errors = append(errors, err)
                mutex.Unlock()
                return
            }
            
            fmt.Printf("Worker %d: completed successfully\n", id)
        }(i)
    }
    
    wg.Wait()
    
    if len(errors) > 0 {
        fmt.Printf("Errors occurred: %v\n", errors)
    } else {
        fmt.Println("All workers completed successfully")
    }
}

func main() {
    waitGroupWithErrorHandling()
}
```

### Example 5: WaitGroup with Channels

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupWithChannels() {
    fmt.Println("=== WaitGroup with Channels ===")
    
    var wg sync.WaitGroup
    results := make(chan int, 5)
    
    // Start workers
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Simulate work
            result := id * 10
            fmt.Printf("Worker %d: sending %d\n", id, result)
            results <- result
        }(i)
    }
    
    // Close results channel when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    fmt.Println("Collecting results...")
    for result := range results {
        fmt.Printf("Received: %d\n", result)
    }
}

func main() {
    waitGroupWithChannels()
}
```

### Example 6: WaitGroup with Timeout

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupWithTimeout() {
    fmt.Println("=== WaitGroup with Timeout ===")
    
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Simulate work with different durations
            duration := time.Duration(id) * 2 * time.Second
            fmt.Printf("Worker %d: working for %v\n", id, duration)
            time.Sleep(duration)
            fmt.Printf("Worker %d: finished\n", id)
        }(i)
    }
    
    // Wait with timeout
    done := make(chan bool)
    go func() {
        wg.Wait()
        done <- true
    }()
    
    select {
    case <-done:
        fmt.Println("All workers completed")
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout reached, some workers may still be running")
    }
}

func main() {
    waitGroupWithTimeout()
}
```

### Example 7: WaitGroup Best Practices

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Worker struct {
    ID    int
    wg    *sync.WaitGroup
    mutex *sync.Mutex
}

func (w *Worker) DoWork() {
    defer w.wg.Done() // Always defer Done()
    
    fmt.Printf("Worker %d: starting work\n", w.ID)
    
    // Simulate work
    time.Sleep(time.Duration(w.ID) * 500 * time.Millisecond)
    
    fmt.Printf("Worker %d: finished work\n", w.ID)
}

func waitGroupBestPractices() {
    fmt.Println("=== WaitGroup Best Practices ===")
    
    var wg sync.WaitGroup
    var mutex sync.Mutex
    
    // Create workers
    workers := make([]*Worker, 0)
    for i := 1; i <= 3; i++ {
        worker := &Worker{
            ID:    i,
            wg:    &wg,
            mutex: &mutex,
        }
        workers = append(workers, worker)
    }
    
    // Start all workers
    for _, worker := range workers {
        wg.Add(1) // Add before starting goroutine
        go worker.DoWork()
    }
    
    fmt.Println("Waiting for all workers...")
    wg.Wait()
    fmt.Println("All workers completed!")
}

func main() {
    waitGroupBestPractices()
}
```

### Example 8: WaitGroup Common Mistakes

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func waitGroupCommonMistakes() {
    fmt.Println("=== WaitGroup Common Mistakes ===")
    
    // Mistake 1: Not calling Add
    fmt.Println("Mistake 1: Not calling Add")
    var wg1 sync.WaitGroup
    go func() {
        defer wg1.Done()
        fmt.Println("Worker 1")
    }()
    wg1.Wait() // This will block forever
    
    // Mistake 2: Calling Add after starting goroutine
    fmt.Println("\nMistake 2: Calling Add after starting goroutine")
    var wg2 sync.WaitGroup
    go func() {
        defer wg2.Done()
        fmt.Println("Worker 2")
    }()
    wg2.Add(1) // Too late!
    wg2.Wait()
    
    // Mistake 3: Not calling Done
    fmt.Println("\nMistake 3: Not calling Done")
    var wg3 sync.WaitGroup
    wg3.Add(1)
    go func() {
        fmt.Println("Worker 3")
        // Forgot to call Done()
    }()
    wg3.Wait() // This will block forever
    
    // Correct usage
    fmt.Println("\nCorrect usage:")
    var wg4 sync.WaitGroup
    wg4.Add(1)
    go func() {
        defer wg4.Done()
        fmt.Println("Worker 4")
    }()
    wg4.Wait()
    fmt.Println("All done!")
}

func main() {
    waitGroupCommonMistakes()
}
```

## 🧪 Key Concepts to Remember

1. **Add before starting**: Call `Add()` before starting goroutine
2. **Done when finished**: Call `Done()` when goroutine completes
3. **Wait blocks**: `Wait()` blocks until counter reaches zero
4. **One-time use**: Cannot reuse WaitGroup after Wait returns
5. **Thread-safe**: Can be used by multiple goroutines

## 🎯 When to Use WaitGroup

1. **Coordinate goroutines**: Wait for multiple workers to complete
2. **Collect results**: Gather results from multiple goroutines
3. **Synchronize operations**: Ensure all operations finish before proceeding
4. **Resource cleanup**: Wait for cleanup operations to complete

## 🎯 Best Practices

1. **Always defer Done()**:
   ```go
   wg.Add(1)
   go func() {
       defer wg.Done() // Ensures Done is called even if panic occurs
       // Work
   }()
   ```

2. **Add before starting goroutine**:
   ```go
   // BAD - race condition
   go func() {
       wg.Add(1) // Too late!
       defer wg.Done()
   }()
   
   // GOOD - add before starting
   wg.Add(1)
   go func() {
       defer wg.Done()
   }()
   ```

3. **Use channels for results**:
   ```go
   results := make(chan int, 10)
   go func() {
       wg.Wait()
       close(results)
   }()
   ```

## 🎯 Common Pitfalls

1. **Not calling Add**:
   ```go
   // BAD - will block forever
   go func() {
       defer wg.Done()
   }()
   wg.Wait()
   
   // GOOD - call Add
   wg.Add(1)
   go func() {
       defer wg.Done()
   }()
   wg.Wait()
   ```

2. **Not calling Done**:
   ```go
   // BAD - will block forever
   wg.Add(1)
   go func() {
       // Forgot to call Done()
   }()
   wg.Wait()
   
   // GOOD - call Done
   wg.Add(1)
   go func() {
       defer wg.Done()
   }()
   wg.Wait()
   ```

3. **Reusing WaitGroup**:
   ```go
   // BAD - reusing WaitGroup
   wg.Wait()
   wg.Add(1) // This will panic!
   
   // GOOD - create new WaitGroup
   wg2 := sync.WaitGroup{}
   wg2.Add(1)
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a file processing system:
- Multiple goroutines process different files
- Each goroutine reports its progress
- Use WaitGroup to wait for all files to be processed
- Collect and display the results from all goroutines

**Hint**: Use a struct to hold file information and results, and use channels or a shared slice to collect results.

## 🚀 Next Steps

Now that you understand WaitGroup, let's learn about **sync.Once** in the next file: `15-once-sync.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
