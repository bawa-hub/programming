# 02 - Goroutines Basics

## 🎯 Learning Objectives
- Understand what goroutines are
- Learn how to create and manage goroutines
- Understand goroutine lifecycle
- Learn about goroutine scheduling
- Practice with real examples

## 📚 Theory

### What are Goroutines?

A **goroutine** is a lightweight thread managed by the Go runtime. Think of it as a function that can run independently alongside other functions.

**Key characteristics:**
- **Lightweight**: Only 2KB initial stack (vs 2MB for OS threads)
- **Managed by Go**: Go runtime handles scheduling
- **Cheap to create**: Can have millions of goroutines
- **Cooperative**: Goroutines yield control voluntarily

### How Goroutines Work

1. **Creation**: Use `go` keyword before a function call
2. **Scheduling**: Go runtime schedules them on OS threads
3. **Execution**: Run concurrently with other goroutines
4. **Completion**: End when function returns

### Goroutine vs Thread

| Feature | Goroutine | OS Thread |
|---------|-----------|-----------|
| Stack Size | 2KB (grows) | 2MB (fixed) |
| Creation Cost | Very cheap | Expensive |
| Scheduling | Go runtime | OS kernel |
| Context Switch | Fast | Slow |
| Max Count | Millions | Thousands |

## 💻 Code Examples

### Example 1: Basic Goroutine Creation

```go
package main

import (
    "fmt"
    "time"
)

// Simple function to run in goroutine
func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello %s! (iteration %d)\n", name, i+1)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    fmt.Println("=== Basic Goroutine Creation ===")
    
    // Run function in goroutine
    go sayHello("Alice")
    
    // Run another goroutine
    go sayHello("Bob")
    
    // Main function continues immediately
    fmt.Println("Main function continues...")
    
    // Wait for goroutines to complete
    time.Sleep(2 * time.Second)
    
    fmt.Println("Main function ends")
}
```

**Run this code:**
```bash
go run 02-goroutines-basics.go
```

### Example 2: Anonymous Functions in Goroutines

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== Anonymous Functions in Goroutines ===")
    
    // Method 1: Anonymous function
    go func() {
        fmt.Println("Anonymous function 1")
        time.Sleep(1 * time.Second)
        fmt.Println("Anonymous function 1 completed")
    }()
    
    // Method 2: Anonymous function with parameters
    go func(name string, count int) {
        for i := 0; i < count; i++ {
            fmt.Printf("Hello from %s (iteration %d)\n", name, i+1)
            time.Sleep(300 * time.Millisecond)
        }
    }("Charlie", 3)
    
    // Method 3: Assign to variable first
    worker := func(id int) {
        fmt.Printf("Worker %d starting\n", id)
        time.Sleep(800 * time.Millisecond)
        fmt.Printf("Worker %d finished\n", id)
    }
    
    go worker(1)
    go worker(2)
    
    // Wait for all goroutines
    time.Sleep(2 * time.Second)
    fmt.Println("All goroutines completed")
}
```

### Example 3: Multiple Goroutines

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Duration(job) * 100 * time.Millisecond) // Simulate work
        results <- job * 2 // Send result
    }
}

func main() {
    fmt.Println("=== Multiple Goroutines Working Together ===")
    
    // Create channels
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Start 3 workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Send jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Collect results
    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}
```

### Example 4: Goroutine Lifecycle and Synchronization

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func demonstrateGoroutineLifecycle() {
    fmt.Println("=== Goroutine Lifecycle Demo ===")
    
    var wg sync.WaitGroup
    
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment counter
        go func(id int) {
            defer wg.Done() // Decrement counter when done
            fmt.Printf("Goroutine %d started\n", id)
            time.Sleep(time.Duration(id) * 500 * time.Millisecond)
            fmt.Printf("Goroutine %d finished\n", id)
        }(i)
    }
    
    fmt.Println("All goroutines started, waiting for completion...")
    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("All goroutines completed")
}

func main() {
    demonstrateGoroutineLifecycle()
}
```

### Example 5: Goroutine Scheduling Demo

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func demonstrateScheduling() {
    fmt.Println("=== Goroutine Scheduling Demo ===")
    
    // Print number of CPUs
    fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
    
    // Start many goroutines
    for i := 1; i <= 10; i++ {
        go func(id int) {
            for j := 0; j < 3; j++ {
                fmt.Printf("Goroutine %d: iteration %d\n", id, j+1)
                runtime.Gosched() // Yield to other goroutines
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // Wait a bit and check goroutine count
    time.Sleep(500 * time.Millisecond)
    fmt.Printf("Number of goroutines after starting: %d\n", runtime.NumGoroutine())
    
    // Wait for completion
    time.Sleep(2 * time.Second)
    fmt.Printf("Number of goroutines after completion: %d\n", runtime.NumGoroutine())
}

func main() {
    demonstrateScheduling()
}
```

## 🧪 Key Concepts to Remember

1. **`go` keyword**: Creates a new goroutine
2. **Lightweight**: Very cheap to create and manage
3. **Concurrent execution**: Run alongside other goroutines
4. **Synchronization needed**: Main function might exit before goroutines finish
5. **Go runtime manages**: You don't manage OS threads directly

## 🎯 Common Pitfalls

1. **Main function exits too early**:
   ```go
   // BAD - main exits before goroutine finishes
   go sayHello("Alice")
   // main ends here, goroutine might not complete
   
   // GOOD - wait for goroutine
   go sayHello("Alice")
   time.Sleep(1 * time.Second) // or use WaitGroup
   ```

2. **Loop variable capture**:
   ```go
   // BAD - all goroutines see the same value
   for i := 0; i < 3; i++ {
       go func() {
           fmt.Println(i) // Always prints 3
       }()
   }
   
   // GOOD - pass variable as parameter
   for i := 0; i < 3; i++ {
       go func(id int) {
           fmt.Println(id) // Prints 0, 1, 2
       }(i)
   }
   ```

3. **Not handling goroutine errors**:
   ```go
   // BAD - error is lost
   go func() {
       if err := riskyOperation(); err != nil {
           // Error is lost!
       }
   }()
   
   // GOOD - handle errors properly
   go func() {
       if err := riskyOperation(); err != nil {
           // Send error to channel or log it
           errorChan <- err
       }
   }()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a bank with multiple tellers:
- 3 tellers (goroutines) serve customers
- Each customer takes a random amount of time (1-3 seconds)
- Serve 10 customers total
- Show which teller serves which customer and when

**Hint**: Use `time.Sleep()` with random durations, and `fmt.Printf()` to show the activity.

## 🚀 Next Steps

Now that you understand goroutines, let's learn about **channels** - how goroutines communicate with each other in the next file: `03-channels-introduction.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
