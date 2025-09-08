# 05 - Channel Operations

## 🎯 Learning Objectives
- Master all channel operations (send, receive, close)
- Understand channel blocking behavior
- Learn about channel state checking
- Practice with advanced channel patterns
- Understand channel safety and best practices

## 📚 Theory

### Channel Operations

Channels support three main operations:

1. **Send Operation**: `ch <- value`
2. **Receive Operation**: `value := <-ch`
3. **Close Operation**: `close(ch)`

### Send Operation (`ch <- value`)

- Sends a value to the channel
- Blocks if channel is full (buffered) or no receiver (unbuffered)
- Panics if channel is closed
- Returns immediately if channel has space

### Receive Operation (`value := <-ch`)

- Receives a value from the channel
- Blocks if channel is empty
- Returns zero value if channel is closed
- Can check if channel is closed with `value, ok := <-ch`

### Close Operation (`close(ch)`)

- Closes the channel
- Sending to closed channel panics
- Receiving from closed channel returns zero value
- Can only close a channel once

## 💻 Code Examples

### Example 1: Basic Send and Receive

```go
package main

import (
    "fmt"
    "time"
)

func basicSendReceive() {
    fmt.Println("=== Basic Send and Receive ===")
    
    ch := make(chan string)
    
    // Send in goroutine
    go func() {
        fmt.Println("Sending message...")
        ch <- "Hello, World!"
        fmt.Println("Message sent")
    }()
    
    // Receive in main
    fmt.Println("Waiting for message...")
    message := <-ch
    fmt.Printf("Received: %s\n", message)
}

func main() {
    basicSendReceive()
}
```

**Run this code:**
```bash
go run 05-channel-operations.go
```

### Example 2: Channel State Checking

```go
package main

import (
    "fmt"
    "time"
)

func channelStateChecking() {
    fmt.Println("=== Channel State Checking ===")
    
    ch := make(chan int, 2)
    
    // Send some data
    ch <- 1
    ch <- 2
    
    // Check channel state
    fmt.Printf("Channel length: %d\n", len(ch))
    fmt.Printf("Channel capacity: %d\n", cap(ch))
    
    // Receive with state check
    for i := 0; i < 3; i++ {
        value, ok := <-ch
        if ok {
            fmt.Printf("Received: %d\n", value)
        } else {
            fmt.Println("Channel is closed")
        }
    }
}

func main() {
    channelStateChecking()
}
```

### Example 3: Channel Closing Patterns

```go
package main

import (
    "fmt"
    "time"
)

func sender(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Sending: %d\n", i)
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }
    close(ch) // Close when done
    fmt.Println("Sender finished")
}

func receiver(ch <-chan int) {
    for {
        value, ok := <-ch
        if !ok {
            fmt.Println("Channel closed, receiver stopping")
            break
        }
        fmt.Printf("Received: %d\n", value)
        time.Sleep(300 * time.Millisecond)
    }
}

func main() {
    fmt.Println("=== Channel Closing Patterns ===")
    
    ch := make(chan int, 3)
    
    go sender(ch)
    go receiver(ch)
    
    time.Sleep(4 * time.Second)
}
```

### Example 4: Range Over Channel

```go
package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- string) {
    messages := []string{"Hello", "World", "from", "Go", "channels"}
    
    for _, msg := range messages {
        fmt.Printf("Producing: %s\n", msg)
        ch <- msg
        time.Sleep(300 * time.Millisecond)
    }
    close(ch) // Must close for range to work
}

func consumer(ch <-chan string) {
    fmt.Println("=== Using Range Over Channel ===")
    for msg := range ch {
        fmt.Printf("Consuming: %s\n", msg)
        time.Sleep(200 * time.Millisecond)
    }
    fmt.Println("No more messages")
}

func main() {
    ch := make(chan string, 3)
    
    go producer(ch)
    consumer(ch)
}
```

### Example 5: Channel Safety and Error Handling

```go
package main

import (
    "fmt"
    "time"
)

func safeChannelOperations() {
    fmt.Println("=== Safe Channel Operations ===")
    
    ch := make(chan int, 2)
    
    // Safe sending
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Panic recovered: %v\n", r)
            }
        }()
        
        ch <- 1
        ch <- 2
        close(ch)
        
        // This will panic - sending to closed channel
        ch <- 3
    }()
    
    time.Sleep(1 * time.Second)
    
    // Safe receiving
    for {
        value, ok := <-ch
        if !ok {
            fmt.Println("Channel closed, stopping receiver")
            break
        }
        fmt.Printf("Received: %d\n", value)
    }
}

func main() {
    safeChannelOperations()
}
```

### Example 6: Advanced Channel Patterns

```go
package main

import (
    "fmt"
    "time"
)

// Fan-out pattern: distribute work to multiple workers
func fanOut(input <-chan int, outputs []chan<- int) {
    defer func() {
        for _, ch := range outputs {
            close(ch)
        }
    }()
    
    for value := range input {
        for _, ch := range outputs {
            ch <- value
        }
    }
}

// Fan-in pattern: collect results from multiple workers
func fanIn(inputs []<-chan int, output chan<- int) {
    defer close(output)
    
    for _, input := range inputs {
        go func(ch <-chan int) {
            for value := range ch {
                output <- value
            }
        }(input)
    }
}

func worker(id int, input <-chan int, output chan<- int) {
    defer close(output)
    
    for value := range input {
        result := value * value // Square the value
        fmt.Printf("Worker %d: %d^2 = %d\n", id, value, result)
        output <- result
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    fmt.Println("=== Advanced Channel Patterns ===")
    
    // Create channels
    input := make(chan int, 10)
    outputs := make([]chan int, 3)
    for i := range outputs {
        outputs[i] = make(chan int, 5)
    }
    
    // Start fan-out
    go fanOut(input, outputs)
    
    // Start workers
    workerOutputs := make([]chan int, 3)
    for i := range workerOutputs {
        workerOutputs[i] = make(chan int, 5)
        go worker(i+1, outputs[i], workerOutputs[i])
    }
    
    // Start fan-in
    finalOutput := make(chan int, 10)
    go fanIn(workerOutputs, finalOutput)
    
    // Send input data
    go func() {
        for i := 1; i <= 5; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Collect results
    for result := range finalOutput {
        fmt.Printf("Final result: %d\n", result)
    }
}
```

### Example 7: Channel as Signal

```go
package main

import (
    "fmt"
    "time"
)

func signalPatterns() {
    fmt.Println("=== Channel as Signal ===")
    
    // Signal channel (no data, just signal)
    done := make(chan struct{})
    
    // Worker that signals completion
    go func() {
        fmt.Println("Worker starting...")
        time.Sleep(2 * time.Second)
        fmt.Println("Worker finished")
        done <- struct{}{} // Send signal
    }()
    
    // Wait for signal
    fmt.Println("Waiting for worker...")
    <-done
    fmt.Println("Worker completed!")
}

func timeoutPattern() {
    fmt.Println("\n=== Timeout Pattern ===")
    
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
    signalPatterns()
    timeoutPattern()
}
```

## 🧪 Key Concepts to Remember

1. **Send operation**: `ch <- value` (blocks if full/closed)
2. **Receive operation**: `value := <-ch` (blocks if empty)
3. **Close operation**: `close(ch)` (can only close once)
4. **State checking**: `value, ok := <-ch` (check if closed)
5. **Range over channel**: `for value := range ch` (until closed)
6. **Channel length**: `len(ch)` (current items)
7. **Channel capacity**: `cap(ch)` (max items)

## 🎯 Best Practices

1. **Always close channels when done**:
   ```go
   defer close(ch)
   ```

2. **Check if channel is closed**:
   ```go
   value, ok := <-ch
   if !ok {
       // Channel is closed
   }
   ```

3. **Use buffered channels for queuing**:
   ```go
   jobs := make(chan Job, 100) // Queue up to 100 jobs
   ```

4. **Use unbuffered channels for synchronization**:
   ```go
   done := make(chan bool) // Perfect sync
   ```

5. **Handle panics from closed channels**:
   ```go
   defer func() {
       if r := recover(); r != nil {
           // Handle panic
       }
   }()
   ```

## 🎯 Common Pitfalls

1. **Sending to closed channel**:
   ```go
   // BAD - panic
   close(ch)
   ch <- value
   
   // GOOD - don't send after closing
   ch <- value
   close(ch)
   ```

2. **Not checking channel state**:
   ```go
   // BAD - might get zero value
   value := <-ch
   
   // GOOD - check if closed
   value, ok := <-ch
   if !ok {
       // Handle closed channel
   }
   ```

3. **Deadlock with unbuffered channels**:
   ```go
   // BAD - deadlock
   ch := make(chan int)
   ch <- 42
   
   // GOOD - have receiver ready
   ch := make(chan int)
   go func() { ch <- 42 }()
   value := <-ch
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a message queue system:
- A producer sends messages to a queue
- Multiple consumers process messages from the queue
- Use proper channel closing and state checking
- Handle the case when all messages are processed

**Hint**: Use a buffered channel for the queue, close it when done, and use range to consume messages.

## 🚀 Next Steps

Now that you understand channel operations, let's learn about the **select statement** in the next file: `06-select-statement.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
