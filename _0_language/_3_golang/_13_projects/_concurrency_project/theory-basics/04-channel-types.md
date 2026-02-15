# 04 - Channel Types

## 🎯 Learning Objectives
- Understand different types of channels
- Learn when to use buffered vs unbuffered channels
- Understand channel capacity and blocking behavior
- Practice with different channel patterns
- Learn about channel zero values

## 📚 Theory

### Types of Channels

Go has several types of channels, each with different behaviors:

1. **Unbuffered Channels** (default)
2. **Buffered Channels**
3. **Nil Channels**
4. **Closed Channels**

### Unbuffered Channels

**Default behavior** - synchronous communication:
- Sender blocks until receiver is ready
- Receiver blocks until sender is ready
- Perfect synchronization between goroutines

```go
ch := make(chan int) // Unbuffered
```

### Buffered Channels

**Asynchronous communication** - can hold multiple values:
- Sender only blocks when buffer is full
- Receiver only blocks when buffer is empty
- Better performance for certain patterns

```go
ch := make(chan int, 5) // Buffered with capacity 5
```

### Nil Channels

**Special state** - cannot send or receive:
- Sending to nil channel blocks forever
- Receiving from nil channel blocks forever
- Useful for disabling select cases

```go
var ch chan int // nil channel
```

### Closed Channels

**End of data** - no more values will be sent:
- Sending to closed channel panics
- Receiving from closed channel returns zero value
- Use `close(ch)` to close a channel

## 💻 Code Examples

### Example 1: Unbuffered vs Buffered Channels

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateUnbuffered() {
    fmt.Println("=== Unbuffered Channel ===")
    
    ch := make(chan int) // Unbuffered
    
    // This will block until someone receives
    go func() {
        fmt.Println("Sending 1...")
        ch <- 1
        fmt.Println("Sent 1")
        
        fmt.Println("Sending 2...")
        ch <- 2
        fmt.Println("Sent 2")
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Receiving...")
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
}

func demonstrateBuffered() {
    fmt.Println("\n=== Buffered Channel ===")
    
    ch := make(chan int, 2) // Buffered with capacity 2
    
    // These won't block because buffer has space
    fmt.Println("Sending 1...")
    ch <- 1
    fmt.Println("Sent 1")
    
    fmt.Println("Sending 2...")
    ch <- 2
    fmt.Println("Sent 2")
    
    // This would block because buffer is full
    // ch <- 3 // Would block here
    
    fmt.Println("Receiving...")
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
}

func main() {
    demonstrateUnbuffered()
    demonstrateBuffered()
}
```

**Run this code:**
```bash
go run 04-channel-types.go
```

### Example 2: Channel Capacity and Blocking

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateCapacity() {
    fmt.Println("=== Channel Capacity Demo ===")
    
    // Create buffered channel with capacity 3
    ch := make(chan string, 3)
    
    fmt.Printf("Channel capacity: %d\n", cap(ch))
    fmt.Printf("Channel length: %d\n", len(ch))
    
    // Fill the buffer
    ch <- "First"
    ch <- "Second"
    ch <- "Third"
    
    fmt.Printf("After sending 3 items - Length: %d\n", len(ch))
    
    // Try to send one more (this would block)
    go func() {
        fmt.Println("Trying to send 4th item...")
        ch <- "Fourth" // This will block until space is available
        fmt.Println("4th item sent!")
    }()
    
    time.Sleep(1 * time.Second)
    
    // Receive one item to make space
    fmt.Printf("Received: %s\n", <-ch)
    fmt.Printf("After receiving 1 item - Length: %d\n", len(ch))
    
    time.Sleep(1 * time.Second) // Let the 4th item be sent
    fmt.Printf("Received: %s\n", <-ch)
}

func main() {
    demonstrateCapacity()
}
```

### Example 3: Nil Channels

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateNilChannels() {
    fmt.Println("=== Nil Channels Demo ===")
    
    var ch chan int // nil channel
    
    // Sending to nil channel blocks forever
    go func() {
        fmt.Println("Trying to send to nil channel...")
        ch <- 42 // This will block forever
        fmt.Println("This will never print")
    }()
    
    // Receiving from nil channel blocks forever
    go func() {
        fmt.Println("Trying to receive from nil channel...")
        value := <-ch // This will block forever
        fmt.Printf("Received: %d\n", value) // This will never print
    }()
    
    time.Sleep(2 * time.Second)
    fmt.Println("Nil channels block forever")
}

func main() {
    demonstrateNilChannels()
}
```

### Example 4: Closed Channels

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateClosedChannels() {
    fmt.Println("=== Closed Channels Demo ===")
    
    ch := make(chan int, 3)
    
    // Send some data
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Close the channel
    close(ch)
    fmt.Println("Channel closed")
    
    // Can still receive data that was sent before closing
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
    
    // Receiving from closed channel returns zero value
    value, ok := <-ch
    fmt.Printf("Received from closed channel: %d, ok: %v\n", value, ok)
    
    // Sending to closed channel panics
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic recovered: %v\n", r)
        }
    }()
    
    ch <- 4 // This will panic
}

func main() {
    demonstrateClosedChannels()
}
```

### Example 5: Channel Zero Values

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateZeroValues() {
    fmt.Println("=== Channel Zero Values Demo ===")
    
    // Zero value of channel is nil
    var ch chan int
    fmt.Printf("Zero value channel: %v\n", ch)
    fmt.Printf("Is nil: %v\n", ch == nil)
    
    // Create actual channel
    ch = make(chan int, 2)
    fmt.Printf("After make: %v\n", ch)
    fmt.Printf("Is nil: %v\n", ch == nil)
    
    // Send and receive
    ch <- 42
    value := <-ch
    fmt.Printf("Value: %d\n", value)
}

func main() {
    demonstrateZeroValues()
}
```

### Example 6: Practical Use Cases

```go
package main

import (
    "fmt"
    "time"
)

// Unbuffered channel for synchronization
func synchronizationExample() {
    fmt.Println("=== Synchronization Example ===")
    
    done := make(chan bool) // Unbuffered for perfect sync
    
    go func() {
        fmt.Println("Worker starting...")
        time.Sleep(1 * time.Second)
        fmt.Println("Worker finished")
        done <- true
    }()
    
    fmt.Println("Waiting for worker...")
    <-done
    fmt.Println("Worker completed!")
}

// Buffered channel for queuing
func queuingExample() {
    fmt.Println("\n=== Queuing Example ===")
    
    jobs := make(chan string, 5) // Buffered for queuing
    
    // Producer
    go func() {
        for i := 1; i <= 5; i++ {
            job := fmt.Sprintf("Job %d", i)
            jobs <- job
            fmt.Printf("Queued: %s\n", job)
        }
        close(jobs)
    }()
    
    // Consumer
    for job := range jobs {
        fmt.Printf("Processing: %s\n", job)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    synchronizationExample()
    queuingExample()
}
```

## 🧪 Key Concepts to Remember

1. **Unbuffered channels**: Perfect synchronization, both sides must be ready
2. **Buffered channels**: Can hold multiple values, better for queuing
3. **Nil channels**: Block forever, useful for disabling select cases
4. **Closed channels**: No more data, returns zero value when empty
5. **Capacity**: `cap(ch)` gives buffer size, `len(ch)` gives current length
6. **Zero value**: `var ch chan int` creates nil channel

## 🎯 When to Use Which Type

### Use Unbuffered When:
- You need perfect synchronization
- Sender and receiver must coordinate
- You want to ensure data is received before continuing

### Use Buffered When:
- You want to decouple sender and receiver
- You need to queue multiple items
- You want better performance for certain patterns

### Use Nil When:
- You want to disable a select case
- You need to temporarily disable communication

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

2. **Sending to closed channel**:
   ```go
   // BAD - panic
   ch := make(chan int)
   close(ch)
   ch <- 42 // Panic!
   ```

3. **Not checking if channel is closed**:
   ```go
   // BAD - might get zero value
   value := <-ch
   
   // GOOD - check if closed
   value, ok := <-ch
   if !ok {
       fmt.Println("Channel closed")
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a printer queue:
- Use a buffered channel to queue print jobs
- A printer goroutine processes jobs from the queue
- Add 5 print jobs to the queue
- Show the order of processing

**Hint**: Use a buffered channel for the queue and show how jobs are queued and processed.

## 🚀 Next Steps

Now that you understand channel types, let's learn about **channel operations** in detail in the next file: `05-channel-operations.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
