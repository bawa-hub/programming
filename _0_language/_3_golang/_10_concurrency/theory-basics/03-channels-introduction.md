# 03 - Channels Introduction

## 🎯 Learning Objectives
- Understand what channels are and why they exist
- Learn how to create and use channels
- Understand channel communication patterns
- Learn about channel blocking behavior
- Practice with real examples

## 📚 Theory

### What are Channels?

A **channel** is a communication mechanism that allows goroutines to send and receive data to each other. Think of it as a pipe where one goroutine can send data and another can receive it.

**Key characteristics:**
- **Type-safe**: Can only send/receive specific types
- **Synchronous by default**: Sender and receiver must be ready
- **Thread-safe**: Multiple goroutines can use the same channel safely
- **First-in-first-out**: Data comes out in the order it was sent

### Why Do We Need Channels?

**Problem**: Goroutines run independently, but often need to:
- Share data between goroutines
- Coordinate their work
- Signal when work is complete
- Pass results back to the main function

**Solution**: Channels provide a safe way to communicate between goroutines.

### Channel Syntax

```go
// Create a channel
ch := make(chan int)

// Send data to channel
ch <- 42

// Receive data from channel
value := <-ch

// Close channel when done
close(ch)
```

## 💻 Code Examples

### Example 1: Basic Channel Communication

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== Basic Channel Communication ===")
    
    // Create a channel for integers
    ch := make(chan int)
    
    // Goroutine that sends data
    go func() {
        fmt.Println("Sending data to channel...")
        ch <- 42
        fmt.Println("Data sent!")
    }()
    
    // Main goroutine receives data
    fmt.Println("Waiting to receive data...")
    value := <-ch
    fmt.Printf("Received: %d\n", value)
}
```

**Run this code:**
```bash
go run 03-channels-introduction.go
```

### Example 2: Channel as Communication Pipe

```go
package main

import (
    "fmt"
    "time"
)

// Function that sends data to channel
func sender(ch chan<- string) {
    messages := []string{"Hello", "World", "from", "Go", "channels"}
    
    for _, msg := range messages {
        fmt.Printf("Sending: %s\n", msg)
        ch <- msg
        time.Sleep(500 * time.Millisecond)
    }
    close(ch) // Close channel when done
}

// Function that receives data from channel
func receiver(ch <-chan string) {
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
        time.Sleep(300 * time.Millisecond)
    }
    fmt.Println("No more messages")
}

func main() {
    fmt.Println("=== Channel as Communication Pipe ===")
    
    // Create a channel for strings
    ch := make(chan string)
    
    // Start sender and receiver
    go sender(ch)
    go receiver(ch)
    
    // Wait for both to complete
    time.Sleep(3 * time.Second)
}
```

### Example 3: Channel Blocking Behavior

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateBlocking() {
    fmt.Println("=== Channel Blocking Behavior ===")
    
    ch := make(chan int)
    
    // This will block because no one is receiving
    fmt.Println("About to send (this will block)...")
    go func() {
        ch <- 42
        fmt.Println("Data sent (this won't print until someone receives)")
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Println("Now receiving...")
    value := <-ch
    fmt.Printf("Received: %d\n", value)
}

func demonstrateNonBlocking() {
    fmt.Println("\n=== Non-blocking with Buffered Channel ===")
    
    // Buffered channel - can hold 2 values
    ch := make(chan int, 2)
    
    // These won't block because channel has buffer
    ch <- 1
    ch <- 2
    fmt.Println("Sent 2 values to buffered channel")
    
    // Receive them
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
}

func main() {
    demonstrateBlocking()
    demonstrateNonBlocking()
}
```

### Example 4: Multiple Goroutines with Channels

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Duration(job) * 200 * time.Millisecond)
        results <- job * job // Send square of the job
    }
}

func main() {
    fmt.Println("=== Multiple Goroutines with Channels ===")
    
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
    close(jobs) // Close jobs channel
    
    // Collect results
    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}
```

### Example 5: Channel Directions

```go
package main

import (
    "fmt"
    "time"
)

// Function that only sends to channel
func sender(ch chan<- string) {
    for i := 1; i <= 3; i++ {
        msg := fmt.Sprintf("Message %d", i)
        fmt.Printf("Sending: %s\n", msg)
        ch <- msg
        time.Sleep(500 * time.Millisecond)
    }
    close(ch)
}

// Function that only receives from channel
func receiver(ch <-chan string) {
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
        time.Sleep(300 * time.Millisecond)
    }
}

// Function that can both send and receive
func processor(input <-chan string, output chan<- string) {
    for msg := range input {
        processed := fmt.Sprintf("Processed: %s", msg)
        fmt.Printf("Processing: %s -> %s\n", msg, processed)
        output <- processed
        time.Sleep(200 * time.Millisecond)
    }
    close(output)
}

func main() {
    fmt.Println("=== Channel Directions ===")
    
    // Create channels
    input := make(chan string)
    output := make(chan string)
    
    // Start goroutines with specific channel directions
    go sender(input)
    go processor(input, output)
    go receiver(output)
    
    // Wait for all to complete
    time.Sleep(3 * time.Second)
}
```

### Example 6: Channel as Signal

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Worker starting...")
    time.Sleep(2 * time.Second)
    fmt.Println("Worker finished")
    done <- true // Signal completion
}

func main() {
    fmt.Println("=== Channel as Signal ===")
    
    // Channel to signal completion
    done := make(chan bool)
    
    // Start worker
    go worker(done)
    
    // Wait for signal
    fmt.Println("Waiting for worker to complete...")
    <-done
    fmt.Println("Worker completed!")
}
```

## 🧪 Key Concepts to Remember

1. **Channels are typed**: `chan int`, `chan string`, etc.
2. **Synchronous by default**: Sender and receiver must be ready
3. **Use `<-` for send/receive**: `ch <- value` to send, `value := <-ch` to receive
4. **Close when done**: Use `close(ch)` to signal no more data
5. **Range over channel**: `for value := range ch` receives until closed
6. **Channel directions**: `chan<-` (send-only), `<-chan` (receive-only)

## 🎯 Common Pitfalls

1. **Deadlock - no receiver**:
   ```go
   // BAD - will deadlock
   ch := make(chan int)
   ch <- 42 // Blocks forever, no receiver
   
   // GOOD - have a receiver
   ch := make(chan int)
   go func() { ch <- 42 }()
   value := <-ch
   ```

2. **Deadlock - no sender**:
   ```go
   // BAD - will deadlock
   ch := make(chan int)
   value := <-ch // Blocks forever, no sender
   
   // GOOD - have a sender
   ch := make(chan int)
   go func() { ch <- 42 }()
   value := <-ch
   ```

3. **Sending to closed channel**:
   ```go
   // BAD - panic!
   ch := make(chan int)
   close(ch)
   ch <- 42 // Panic!
   
   // GOOD - don't send after closing
   ch := make(chan int)
   ch <- 42
   close(ch)
   ```

4. **Receiving from closed channel**:
   ```go
   // BAD - gets zero value
   ch := make(chan int)
   close(ch)
   value := <-ch // Gets 0, not an error
   
   // GOOD - check if channel is closed
   ch := make(chan int)
   close(ch)
   value, ok := <-ch
   if !ok {
       fmt.Println("Channel is closed")
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a restaurant order system:
- A chef (goroutine) prepares dishes and sends them to a channel
- A waiter (goroutine) receives dishes and serves them
- The chef prepares 5 different dishes
- Show the order of preparation and serving

**Hint**: Use channels to pass dish names between chef and waiter goroutines.

## 🚀 Next Steps

Now that you understand basic channels, let's learn about **different types of channels** in the next file: `04-channel-types.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
