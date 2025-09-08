# 09 - Channel Closing Patterns

## 🎯 Learning Objectives
- Understand when and how to close channels
- Learn proper channel closing patterns
- Master detecting closed channels
- Practice with different closing scenarios
- Understand the impact of closing channels

## 📚 Theory

### Why Close Channels?

**Reasons to close channels:**
1. **Signal completion**: No more data will be sent
2. **Enable range loops**: `for value := range ch` works
3. **Prevent goroutine leaks**: Receivers can exit
4. **Clean shutdown**: Graceful termination

### Channel Closing Rules

1. **Only close once**: Closing a closed channel panics
2. **Only sender closes**: Receiver should not close
3. **Close when done**: Close after sending all data
4. **Check if closed**: Use `value, ok := <-ch`

### Channel States After Closing

- **Sending**: Panics immediately
- **Receiving**: Returns zero value and `ok = false`
- **Range loop**: Exits when channel is closed
- **Select**: Case becomes disabled

## 💻 Code Examples

### Example 1: Basic Channel Closing

```go
package main

import (
    "fmt"
    "time"
)

func basicClosing() {
    fmt.Println("=== Basic Channel Closing ===")
    
    ch := make(chan string)
    
    // Sender
    go func() {
        ch <- "Hello"
        ch <- "World"
        close(ch) // Close when done
        fmt.Println("Sender closed channel")
    }()
    
    // Receiver
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
    }
    fmt.Println("Receiver finished")
}

func main() {
    basicClosing()
}
```

**Run this code:**
```bash
go run 09-channel-closing.go
```

### Example 2: Detecting Closed Channels

```go
package main

import (
    "fmt"
    "time"
)

func detectClosedChannels() {
    fmt.Println("=== Detecting Closed Channels ===")
    
    ch := make(chan int)
    
    // Sender
    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
            time.Sleep(500 * time.Millisecond)
        }
        close(ch)
    }()
    
    // Receiver with closed channel detection
    for {
        value, ok := <-ch
        if !ok {
            fmt.Println("Channel is closed")
            break
        }
        fmt.Printf("Received: %d\n", value)
    }
}

func main() {
    detectClosedChannels()
}
```

### Example 3: Multiple Receivers

```go
package main

import (
    "fmt"
    "time"
)

func multipleReceivers() {
    fmt.Println("=== Multiple Receivers ===")
    
    ch := make(chan string)
    
    // Sender
    go func() {
        messages := []string{"Hello", "World", "from", "Go"}
        for _, msg := range messages {
            ch <- msg
            time.Sleep(300 * time.Millisecond)
        }
        close(ch)
    }()
    
    // Multiple receivers
    for i := 1; i <= 3; i++ {
        go func(id int) {
            for msg := range ch {
                fmt.Printf("Receiver %d: %s\n", id, msg)
                time.Sleep(200 * time.Millisecond)
            }
            fmt.Printf("Receiver %d finished\n", id)
        }(i)
    }
    
    time.Sleep(3 * time.Second)
}

func main() {
    multipleReceivers()
}
```

### Example 4: Closing from Different Goroutines

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func closingFromDifferentGoroutines() {
    fmt.Println("=== Closing from Different Goroutines ===")
    
    ch := make(chan string)
    var wg sync.WaitGroup
    
    // Multiple senders
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 1; j <= 2; j++ {
                msg := fmt.Sprintf("Message %d from sender %d", j, id)
                ch <- msg
                time.Sleep(300 * time.Millisecond)
            }
        }(i)
    }
    
    // Close channel when all senders are done
    go func() {
        wg.Wait()
        close(ch)
        fmt.Println("Channel closed after all senders finished")
    }()
    
    // Receiver
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
    }
}

func main() {
    closingFromDifferentGoroutines()
}
```

### Example 5: Safe Channel Closing

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeChannel struct {
    ch   chan string
    once sync.Once
}

func NewSafeChannel() *SafeChannel {
    return &SafeChannel{
        ch: make(chan string),
    }
}

func (sc *SafeChannel) Send(msg string) {
    sc.ch <- msg
}

func (sc *SafeChannel) Close() {
    sc.once.Do(func() {
        close(sc.ch)
    })
}

func (sc *SafeChannel) Receive() <-chan string {
    return sc.ch
}

func safeClosingExample() {
    fmt.Println("=== Safe Channel Closing ===")
    
    sc := NewSafeChannel()
    
    // Sender
    go func() {
        for i := 1; i <= 3; i++ {
            sc.Send(fmt.Sprintf("Message %d", i))
            time.Sleep(300 * time.Millisecond)
        }
        sc.Close()
    }()
    
    // Receiver
    for msg := range sc.Receive() {
        fmt.Printf("Received: %s\n", msg)
    }
    
    // Try to close again (safe)
    sc.Close()
    sc.Close()
    fmt.Println("Multiple closes are safe")
}

func main() {
    safeClosingExample()
}
```

### Example 6: Channel Closing Patterns

```go
package main

import (
    "fmt"
    "time"
)

// Pattern 1: Sender closes
func senderClosesPattern() {
    fmt.Println("=== Pattern 1: Sender Closes ===")
    
    ch := make(chan string)
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch <- fmt.Sprintf("Message %d", i)
        }
        close(ch) // Sender closes
    }()
    
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
    }
}

// Pattern 2: Done channel
func doneChannelPattern() {
    fmt.Println("\n=== Pattern 2: Done Channel ===")
    
    ch := make(chan string)
    done := make(chan bool)
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch <- fmt.Sprintf("Message %d", i)
            time.Sleep(300 * time.Millisecond)
        }
        done <- true
    }()
    
    go func() {
        <-done
        close(ch)
    }()
    
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
    }
}

// Pattern 3: Context cancellation
func contextPattern() {
    fmt.Println("\n=== Pattern 3: Context Cancellation ===")
    
    ch := make(chan string)
    
    go func() {
        for i := 1; i <= 5; i++ {
            select {
            case ch <- fmt.Sprintf("Message %d", i):
                time.Sleep(300 * time.Millisecond)
            case <-time.After(2 * time.Second):
                close(ch)
                return
            }
        }
        close(ch)
    }()
    
    for msg := range ch {
        fmt.Printf("Received: %s\n", msg)
    }
}

func main() {
    senderClosesPattern()
    doneChannelPattern()
    contextPattern()
}
```

### Example 7: Common Closing Mistakes

```go
package main

import (
    "fmt"
    "time"
)

func commonMistakes() {
    fmt.Println("=== Common Closing Mistakes ===")
    
    // Mistake 1: Closing from receiver
    fmt.Println("Mistake 1: Closing from receiver")
    ch1 := make(chan string)
    
    go func() {
        ch1 <- "Hello"
        close(ch1) // BAD: Receiver shouldn't close
    }()
    
    for msg := range ch1 {
        fmt.Printf("Received: %s\n", msg)
    }
    
    // Mistake 2: Sending to closed channel
    fmt.Println("\nMistake 2: Sending to closed channel")
    ch2 := make(chan string)
    close(ch2)
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic recovered: %v\n", r)
        }
    }()
    
    ch2 <- "This will panic" // BAD: Sending to closed channel
    
    // Mistake 3: Not checking if closed
    fmt.Println("\nMistake 3: Not checking if closed")
    ch3 := make(chan string)
    close(ch3)
    
    value := <-ch3
    fmt.Printf("Value from closed channel: '%s'\n", value) // Empty string
    
    // Good practice: Check if closed
    value, ok := <-ch3
    if !ok {
        fmt.Println("Channel is closed")
    }
}

func main() {
    commonMistakes()
}
```

## 🧪 Key Concepts to Remember

1. **Only close once**: Closing a closed channel panics
2. **Only sender closes**: Receiver should not close
3. **Close when done**: Close after sending all data
4. **Check if closed**: Use `value, ok := <-ch`
5. **Range loop exits**: When channel is closed
6. **Zero value returned**: When receiving from closed channel

## 🎯 Best Practices

1. **Close from sender**:
   ```go
   go func() {
       // Send data
       ch <- data
       close(ch) // Sender closes
   }()
   ```

2. **Check if closed**:
   ```go
   value, ok := <-ch
   if !ok {
       // Channel is closed
   }
   ```

3. **Use range loop**:
   ```go
   for value := range ch {
       // Process value
   }
   // Loop exits when channel is closed
   ```

4. **Safe closing with sync.Once**:
   ```go
   var once sync.Once
   once.Do(func() { close(ch) })
   ```

## 🎯 Common Pitfalls

1. **Closing from receiver**:
   ```go
   // BAD - receiver shouldn't close
   for msg := range ch {
       if msg == "done" {
           close(ch) // Panic!
       }
   }
   
   // GOOD - sender closes
   go func() {
       ch <- "done"
       close(ch)
   }()
   ```

2. **Sending to closed channel**:
   ```go
   // BAD - panic
   close(ch)
   ch <- "data" // Panic!
   
   // GOOD - don't send after closing
   ch <- "data"
   close(ch)
   ```

3. **Not checking if closed**:
   ```go
   // BAD - might get zero value
   value := <-ch
   
   // GOOD - check if closed
   value, ok := <-ch
   if !ok {
       // Handle closed channel
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a file processing system:
- Multiple goroutines process files and send results to a channel
- A collector goroutine receives results and writes to output
- Use proper channel closing to signal when all files are processed
- Handle the case where some files might fail to process

**Hint**: Use a done channel or WaitGroup to coordinate when all processors are finished, then close the results channel.

## 🚀 Next Steps

Now that you understand channel closing, let's learn about **timeout patterns** in the next file: `10-timeout-patterns.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
