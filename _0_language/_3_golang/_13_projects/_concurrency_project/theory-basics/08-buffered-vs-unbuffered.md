# 08 - Buffered vs Unbuffered Channels

## 🎯 Learning Objectives
- Understand the difference between buffered and unbuffered channels
- Learn when to use each type
- Master blocking behavior differences
- Practice with real-world examples
- Understand performance implications

## 📚 Theory

### Unbuffered Channels (Default)

**Behavior**: Synchronous communication
- Sender blocks until receiver is ready
- Receiver blocks until sender is ready
- Perfect synchronization between goroutines
- Zero capacity: `make(chan int)`

### Buffered Channels

**Behavior**: Asynchronous communication
- Sender only blocks when buffer is full
- Receiver only blocks when buffer is empty
- Can hold multiple values before blocking
- Non-zero capacity: `make(chan int, 5)`

### Key Differences

| Feature | Unbuffered | Buffered |
|---------|------------|----------|
| Blocking | Always blocks | Only when full/empty |
| Synchronization | Perfect sync | Loose coupling |
| Performance | Slower | Faster |
| Use Case | Coordination | Queuing |

## 💻 Code Examples

### Example 1: Basic Comparison

```go
package main

import (
    "fmt"
    "time"
)

func unbufferedExample() {
    fmt.Println("=== Unbuffered Channel ===")
    
    ch := make(chan string) // Unbuffered
    
    // This will block until someone receives
    go func() {
        fmt.Println("Sending message...")
        ch <- "Hello"
        fmt.Println("Message sent")
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Receiving message...")
    msg := <-ch
    fmt.Printf("Received: %s\n", msg)
}

func bufferedExample() {
    fmt.Println("\n=== Buffered Channel ===")
    
    ch := make(chan string, 2) // Buffered with capacity 2
    
    // These won't block because buffer has space
    fmt.Println("Sending messages...")
    ch <- "Hello"
    ch <- "World"
    fmt.Println("Messages sent (no blocking)")
    
    // Receive them
    fmt.Println("Receiving messages...")
    fmt.Printf("Received: %s\n", <-ch)
    fmt.Printf("Received: %s\n", <-ch)
}

func main() {
    unbufferedExample()
    bufferedExample()
}
```

**Run this code:**
```bash
go run 08-buffered-vs-unbuffered.go
```

### Example 2: Blocking Behavior Demo

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateBlocking() {
    fmt.Println("=== Blocking Behavior Demo ===")
    
    // Unbuffered channel
    unbuffered := make(chan int)
    
    // This will block forever (deadlock)
    go func() {
        fmt.Println("Trying to send to unbuffered channel...")
        unbuffered <- 42
        fmt.Println("This will never print")
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Unbuffered channel blocks sender")
    
    // Buffered channel
    buffered := make(chan int, 1)
    
    // This won't block
    go func() {
        fmt.Println("Sending to buffered channel...")
        buffered <- 42
        fmt.Println("Sent to buffered channel")
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Buffered channel doesn't block sender")
    
    // Receive from buffered channel
    value := <-buffered
    fmt.Printf("Received: %d\n", value)
}

func main() {
    demonstrateBlocking()
}
```

### Example 3: Performance Comparison

```go
package main

import (
    "fmt"
    "time"
)

func performanceTest() {
    fmt.Println("=== Performance Comparison ===")
    
    const iterations = 1000
    
    // Test unbuffered channel
    start := time.Now()
    unbuffered := make(chan int)
    
    go func() {
        for i := 0; i < iterations; i++ {
            unbuffered <- i
        }
        close(unbuffered)
    }()
    
    for range unbuffered {
        // Just receive
    }
    
    unbufferedTime := time.Since(start)
    
    // Test buffered channel
    start = time.Now()
    buffered := make(chan int, 100) // Buffer size 100
    
    go func() {
        for i := 0; i < iterations; i++ {
            buffered <- i
        }
        close(buffered)
    }()
    
    for range buffered {
        // Just receive
    }
    
    bufferedTime := time.Since(start)
    
    fmt.Printf("Unbuffered time: %v\n", unbufferedTime)
    fmt.Printf("Buffered time: %v\n", bufferedTime)
    fmt.Printf("Speedup: %.2fx\n", float64(unbufferedTime)/float64(bufferedTime))
}

func main() {
    performanceTest()
}
```

### Example 4: Real-World Use Cases

```go
package main

import (
    "fmt"
    "time"
)

// Unbuffered for synchronization
func synchronizationExample() {
    fmt.Println("=== Synchronization Example ===")
    
    done := make(chan bool) // Unbuffered for perfect sync
    
    go func() {
        fmt.Println("Worker starting...")
        time.Sleep(2 * time.Second)
        fmt.Println("Worker finished")
        done <- true
    }()
    
    fmt.Println("Waiting for worker...")
    <-done
    fmt.Println("Worker completed!")
}

// Buffered for queuing
func queuingExample() {
    fmt.Println("\n=== Queuing Example ===")
    
    jobs := make(chan string, 5) // Buffered for queuing
    
    // Producer
    go func() {
        for i := 1; i <= 10; i++ {
            job := fmt.Sprintf("Job %d", i)
            jobs <- job
            fmt.Printf("Queued: %s\n", job)
            time.Sleep(100 * time.Millisecond)
        }
        close(jobs)
    }()
    
    // Consumer
    for job := range jobs {
        fmt.Printf("Processing: %s\n", job)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    synchronizationExample()
    queuingExample()
}
```

### Example 5: Buffer Size Impact

```go
package main

import (
    "fmt"
    "time"
)

func testBufferSize(bufferSize int) time.Duration {
    ch := make(chan int, bufferSize)
    
    start := time.Now()
    
    // Producer
    go func() {
        for i := 0; i < 1000; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    // Consumer
    for range ch {
        // Just receive
    }
    
    return time.Since(start)
}

func main() {
    fmt.Println("=== Buffer Size Impact ===")
    
    sizes := []int{0, 1, 10, 100, 1000}
    
    for _, size := range sizes {
        duration := testBufferSize(size)
        fmt.Printf("Buffer size %d: %v\n", size, duration)
    }
}
```

### Example 6: Producer-Consumer Pattern

```go
package main

import (
    "fmt"
    "time"
)

func producerConsumerExample() {
    fmt.Println("=== Producer-Consumer Pattern ===")
    
    // Use buffered channel for decoupling
    buffer := make(chan string, 3)
    
    // Producer
    go func() {
        items := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
        
        for _, item := range items {
            fmt.Printf("Producing: %s\n", item)
            buffer <- item
            time.Sleep(300 * time.Millisecond)
        }
        close(buffer)
    }()
    
    // Consumer
    for item := range buffer {
        fmt.Printf("Consuming: %s\n", item)
        time.Sleep(500 * time.Millisecond)
    }
    
    fmt.Println("All items processed")
}

func main() {
    producerConsumerExample()
}
```

### Example 7: When to Use Each Type

```go
package main

import (
    "fmt"
    "time"
)

// Use unbuffered for coordination
func coordinationExample() {
    fmt.Println("=== Coordination Example ===")
    
    ready := make(chan bool) // Unbuffered
    
    go func() {
        fmt.Println("Preparing data...")
        time.Sleep(1 * time.Second)
        fmt.Println("Data ready")
        ready <- true
    }()
    
    fmt.Println("Waiting for data to be ready...")
    <-ready
    fmt.Println("Data is ready, proceeding...")
}

// Use buffered for batching
func batchingExample() {
    fmt.Println("\n=== Batching Example ===")
    
    batch := make(chan string, 5) // Buffered for batching
    
    go func() {
        for i := 1; i <= 10; i++ {
            item := fmt.Sprintf("Item %d", i)
            batch <- item
            fmt.Printf("Added to batch: %s\n", item)
            
            if i%5 == 0 {
                fmt.Println("Batch full, processing...")
                time.Sleep(500 * time.Millisecond)
            }
        }
        close(batch)
    }()
    
    for item := range batch {
        fmt.Printf("Processing: %s\n", item)
    }
}

func main() {
    coordinationExample()
    batchingExample()
}
```

## 🧪 Key Concepts to Remember

1. **Unbuffered**: Perfect synchronization, always blocks
2. **Buffered**: Loose coupling, only blocks when full/empty
3. **Performance**: Buffered is faster for high-throughput
4. **Use cases**: Unbuffered for coordination, buffered for queuing
5. **Buffer size**: Affects performance and memory usage

## 🎯 When to Use Which

### Use Unbuffered When:
- You need perfect synchronization
- Sender and receiver must coordinate
- You want to ensure data is received before continuing
- You need to signal completion

### Use Buffered When:
- You want to decouple sender and receiver
- You need to queue multiple items
- You want better performance
- You need to handle bursts of data

## 🎯 Common Pitfalls

1. **Deadlock with unbuffered channels**:
   ```go
   // BAD - deadlock
   ch := make(chan int)
   ch <- 42 // Blocks forever
   
   // GOOD - have receiver ready
   ch := make(chan int)
   go func() { ch <- 42 }()
   value := <-ch
   ```

2. **Wrong buffer size**:
   ```go
   // BAD - too small buffer
   ch := make(chan int, 1)
   ch <- 1
   ch <- 2 // Blocks here
   
   // GOOD - appropriate buffer size
   ch := make(chan int, 10)
   ch <- 1
   ch <- 2 // No blocking
   ```

3. **Not understanding blocking behavior**:
   ```go
   // BAD - might block unexpectedly
   ch := make(chan int)
   go func() { ch <- 42 }()
   // If goroutine doesn't start, this blocks
   value := <-ch
   
   // GOOD - use buffered or ensure goroutine starts
   ch := make(chan int, 1)
   go func() { ch <- 42 }()
   value := <-ch
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a message queue system:
- Use buffered channels to queue messages
- Multiple producers send messages to the queue
- Multiple consumers process messages from the queue
- Show how buffer size affects performance
- Compare with unbuffered approach

**Hint**: Use different buffer sizes and measure the time it takes to process all messages.

## 🚀 Next Steps

Now that you understand buffered vs unbuffered channels, let's learn about **channel closing patterns** in the next file: `09-channel-closing.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
