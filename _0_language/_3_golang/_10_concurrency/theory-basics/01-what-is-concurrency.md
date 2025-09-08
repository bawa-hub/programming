# 01 - What is Concurrency?

## 🎯 Learning Objectives
- Understand what concurrency is
- Learn the difference between concurrency and parallelism
- See why concurrency is important
- Understand Go's approach to concurrency

## 📚 Theory

### What is Concurrency?

**Concurrency** is the ability of a program to deal with multiple things at once. Think of it like a chef who can:
- Start cooking pasta
- While pasta is boiling, chop vegetables
- While vegetables are cooking, prepare the sauce
- All at the same time, not one after another

### Concurrency vs Parallelism

**Concurrency** = Managing multiple tasks (can be on one CPU)
**Parallelism** = Actually doing multiple tasks simultaneously (needs multiple CPUs)

**Example:**
- **Concurrency**: A single chef managing multiple dishes
- **Parallelism**: Multiple chefs each cooking different dishes

### Why is Concurrency Important?

1. **Better Resource Utilization**: Don't waste time waiting
2. **Responsive Applications**: UI doesn't freeze while processing
3. **Scalability**: Handle more users/requests
4. **Efficiency**: Get more work done in less time

### Go's Approach: Goroutines

Go makes concurrency easy with **goroutines** - lightweight threads that are:
- Very cheap to create (2KB stack vs 2MB for OS threads)
- Managed by Go runtime (not OS)
- Can have millions of them running

## 💻 Code Examples

### Example 1: Sequential vs Concurrent

```go
package main

import (
    "fmt"
    "time"
)

// Sequential approach - one after another
func sequentialWork() {
    fmt.Println("Sequential Work:")
    
    // Task 1
    fmt.Println("  Starting task 1...")
    time.Sleep(1 * time.Second) // Simulate work
    fmt.Println("  Task 1 completed")
    
    // Task 2
    fmt.Println("  Starting task 2...")
    time.Sleep(1 * time.Second) // Simulate work
    fmt.Println("  Task 2 completed")
    
    // Task 3
    fmt.Println("  Starting task 3...")
    time.Sleep(1 * time.Second) // Simulate work
    fmt.Println("  Task 3 completed")
    
    fmt.Println("All tasks completed sequentially")
}

// Concurrent approach - all at once
func concurrentWork() {
    fmt.Println("\nConcurrent Work:")
    
    // Start all tasks at the same time
    go func() {
        fmt.Println("  Starting task 1...")
        time.Sleep(1 * time.Second)
        fmt.Println("  Task 1 completed")
    }()
    
    go func() {
        fmt.Println("  Starting task 2...")
        time.Sleep(1 * time.Second)
        fmt.Println("  Task 2 completed")
    }()
    
    go func() {
        fmt.Println("  Starting task 3...")
        time.Sleep(1 * time.Second)
        fmt.Println("  Task 3 completed")
    }()
    
    // Wait for all goroutines to complete
    time.Sleep(2 * time.Second)
    fmt.Println("All tasks completed concurrently")
}

func main() {
    // Run sequential version
    start := time.Now()
    sequentialWork()
    sequentialTime := time.Since(start)
    
    // Run concurrent version
    start = time.Now()
    concurrentWork()
    concurrentTime := time.Since(start)
    
    fmt.Printf("\nTime comparison:\n")
    fmt.Printf("Sequential: %v\n", sequentialTime)
    fmt.Printf("Concurrent: %v\n", concurrentTime)
    fmt.Printf("Speedup: %.2fx\n", float64(sequentialTime)/float64(concurrentTime))
}
```

**Run this code:**
```bash
go run 01-what-is-concurrency.go
```

### Example 2: Real-World Analogy

```go
package main

import (
    "fmt"
    "time"
)

// Simulate different types of work
func downloadFile(filename string, size int) {
    fmt.Printf("📥 Downloading %s (%d MB)...\n", filename, size)
    time.Sleep(time.Duration(size) * 100 * time.Millisecond) // Simulate download time
    fmt.Printf("✅ %s downloaded successfully\n", filename)
}

func processData(data string) {
    fmt.Printf("🔄 Processing %s...\n", data)
    time.Sleep(500 * time.Millisecond) // Simulate processing time
    fmt.Printf("✅ %s processed successfully\n", data)
}

func sendEmail(recipient string) {
    fmt.Printf("📧 Sending email to %s...\n", recipient)
    time.Sleep(200 * time.Millisecond) // Simulate email sending
    fmt.Printf("✅ Email sent to %s\n", recipient)
}

func main() {
    fmt.Println("=== Sequential Approach ===")
    start := time.Now()
    
    // Do everything one by one
    downloadFile("document.pdf", 5)
    downloadFile("image.jpg", 3)
    downloadFile("video.mp4", 10)
    
    processData("user_data")
    processData("analytics_data")
    
    sendEmail("user@example.com")
    sendEmail("admin@example.com")
    
    fmt.Printf("Sequential time: %v\n", time.Since(start))
    
    fmt.Println("\n=== Concurrent Approach ===")
    start = time.Now()
    
    // Do everything concurrently
    go downloadFile("document.pdf", 5)
    go downloadFile("image.jpg", 3)
    go downloadFile("video.mp4", 10)
    
    go processData("user_data")
    go processData("analytics_data")
    
    go sendEmail("user@example.com")
    go sendEmail("admin@example.com")
    
    // Wait for all goroutines to complete
    time.Sleep(2 * time.Second)
    
    fmt.Printf("Concurrent time: %v\n", time.Since(start))
}
```

### Example 3: Goroutine Lifecycle

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateGoroutineLifecycle() {
    fmt.Println("=== Goroutine Lifecycle Demo ===")
    
    // Create a goroutine
    go func() {
        fmt.Println("🚀 Goroutine started")
        
        // Do some work
        for i := 1; i <= 3; i++ {
            fmt.Printf("  Working... step %d\n", i)
            time.Sleep(500 * time.Millisecond)
        }
        
        fmt.Println("✅ Goroutine finished")
    }()
    
    fmt.Println("Main function continues...")
    
    // Give goroutine time to complete
    time.Sleep(2 * time.Second)
    
    fmt.Println("Main function ends")
}

func main() {
    demonstrateGoroutineLifecycle()
}
```

## 🧪 Key Concepts to Remember

1. **Concurrency is about structure** - organizing your program to handle multiple tasks
2. **Parallelism is about execution** - actually doing multiple things at once
3. **Goroutines are lightweight** - you can have millions of them
4. **Go manages goroutines** - you don't manage OS threads directly
5. **Concurrency improves efficiency** - better resource utilization

## 🎯 Common Pitfalls

1. **Forgetting to wait** - Main function might exit before goroutines finish
2. **Not understanding the difference** - Concurrency ≠ Parallelism
3. **Overusing goroutines** - Not everything needs to be concurrent

## 🏋️ Exercise

**Problem**: Create a program that simulates a restaurant kitchen where:
- 3 chefs are cooking different dishes concurrently
- Each dish takes different amounts of time
- Show the difference between sequential and concurrent cooking

**Hint**: Use `time.Sleep()` to simulate cooking time, and `fmt.Printf()` to show what each chef is doing.

## 🚀 Next Steps

Now that you understand what concurrency is, let's learn about **goroutines** in detail in the next file: `02-goroutines-basics.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
