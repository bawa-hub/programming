# 06 - Select Statement

## 🎯 Learning Objectives
- Understand what the select statement is and why it's important
- Learn how to use select for non-blocking operations
- Master timeout and default cases
- Practice with multiple channel operations
- Learn advanced select patterns

## 📚 Theory

### What is the Select Statement?

The **select** statement is like a switch statement, but for channels. It allows a goroutine to wait on multiple channel operations simultaneously.

**Key features:**
- **Non-blocking**: Can handle multiple channels at once
- **Random selection**: If multiple cases are ready, one is chosen randomly
- **Timeout support**: Can wait with a timeout
- **Default case**: Can be non-blocking

### Why Do We Need Select?

**Problem**: Sometimes you need to:
- Wait on multiple channels
- Handle timeouts
- Do non-blocking operations
- Prioritize certain channels

**Solution**: Select statement handles all these cases elegantly.

### Select Syntax

```go
select {
case value := <-ch1:
    // Handle value from ch1
case ch2 <- value:
    // Send value to ch2
case <-time.After(1 * time.Second):
    // Timeout after 1 second
default:
    // No channel is ready
}
```

## 💻 Code Examples

### Example 1: Basic Select Statement

```go
package main

import (
    "fmt"
    "time"
)

func basicSelect() {
    fmt.Println("=== Basic Select Statement ===")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start goroutines that will send to channels
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from channel 1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from channel 2"
    }()
    
    // Select will wait for the first available channel
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received from ch1: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received from ch2: %s\n", msg2)
        }
    }
}

func main() {
    basicSelect()
}
```

**Run this code:**
```bash
go run 06-select-statement.go
```

### Example 2: Select with Timeout

```go
package main

import (
    "fmt"
    "time"
)

func selectWithTimeout() {
    fmt.Println("=== Select with Timeout ===")
    
    ch := make(chan string)
    
    // Start a goroutine that takes time
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "Slow message"
    }()
    
    // Wait with timeout
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout! No message received")
    }
}

func main() {
    selectWithTimeout()
}
```

### Example 3: Non-blocking Select with Default

```go
package main

import (
    "fmt"
    "time"
)

func nonBlockingSelect() {
    fmt.Println("=== Non-blocking Select ===")
    
    ch := make(chan string)
    
    // Try to receive without blocking
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("No message available, using default")
    }
    
    // Send a message
    go func() {
        time.Sleep(500 * time.Millisecond)
        ch <- "Hello"
    }()
    
    // Try again after a short delay
    time.Sleep(1 * time.Second)
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("Still no message")
    }
}

func main() {
    nonBlockingSelect()
}
```

### Example 4: Multiple Channel Operations

```go
package main

import (
    "fmt"
    "time"
)

func multipleChannels() {
    fmt.Println("=== Multiple Channel Operations ===")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)
    
    // Start goroutines with different delays
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Fast message"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Medium message"
    }()
    
    go func() {
        time.Sleep(3 * time.Second)
        ch3 <- "Slow message"
    }()
    
    // Wait for all messages
    for i := 0; i < 3; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received from ch1: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received from ch2: %s\n", msg2)
        case msg3 := <-ch3:
            fmt.Printf("Received from ch3: %s\n", msg3)
        }
    }
}

func main() {
    multipleChannels()
}
```

### Example 5: Select in a Loop

```go
package main

import (
    "fmt"
    "time"
)

func selectInLoop() {
    fmt.Println("=== Select in a Loop ===")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    done := make(chan bool)
    
    // Start goroutines
    go func() {
        for i := 1; i <= 3; i++ {
            ch1 <- fmt.Sprintf("Message %d from ch1", i)
            time.Sleep(500 * time.Millisecond)
        }
        close(ch1)
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch2 <- fmt.Sprintf("Message %d from ch2", i)
            time.Sleep(700 * time.Millisecond)
        }
        close(ch2)
    }()
    
    // Process messages until both channels are closed
    for {
        select {
        case msg1, ok := <-ch1:
            if ok {
                fmt.Printf("Received from ch1: %s\n", msg1)
            } else {
                ch1 = nil // Disable this case
            }
        case msg2, ok := <-ch2:
            if ok {
                fmt.Printf("Received from ch2: %s\n", msg2)
            } else {
                ch2 = nil // Disable this case
            }
        case <-time.After(5 * time.Second):
            fmt.Println("Timeout reached")
            return
        }
        
        // Exit if both channels are closed
        if ch1 == nil && ch2 == nil {
            break
        }
    }
    
    fmt.Println("All channels closed")
}

func main() {
    selectInLoop()
}
```

### Example 6: Select for Sending and Receiving

```go
package main

import (
    "fmt"
    "time"
)

func selectSendReceive() {
    fmt.Println("=== Select for Sending and Receiving ===")
    
    input := make(chan string, 2)
    output := make(chan string, 2)
    
    // Start a processor
    go func() {
        for msg := range input {
            processed := fmt.Sprintf("Processed: %s", msg)
            output <- processed
        }
        close(output)
    }()
    
    // Send some messages
    go func() {
        input <- "Hello"
        input <- "World"
        close(input)
    }()
    
    // Process with select
    for {
        select {
        case msg, ok := <-input:
            if ok {
                fmt.Printf("Received input: %s\n", msg)
            } else {
                input = nil // Disable input case
            }
        case result, ok := <-output:
            if ok {
                fmt.Printf("Received output: %s\n", result)
            } else {
                output = nil // Disable output case
            }
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout")
            return
        }
        
        if input == nil && output == nil {
            break
        }
    }
}

func main() {
    selectSendReceive()
}
```

### Example 7: Advanced Select Patterns

```go
package main

import (
    "fmt"
    "time"
)

func advancedSelectPatterns() {
    fmt.Println("=== Advanced Select Patterns ===")
    
    // Pattern 1: Priority channel
    priority := make(chan string)
    normal := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        normal <- "Normal message"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        priority <- "Priority message"
    }()
    
    // Always check priority first
    for i := 0; i < 2; i++ {
        select {
        case msg := <-priority:
            fmt.Printf("Priority: %s\n", msg)
        case msg := <-normal:
            fmt.Printf("Normal: %s\n", msg)
        }
    }
    
    // Pattern 2: Heartbeat
    heartbeat := time.NewTicker(500 * time.Millisecond)
    defer heartbeat.Stop()
    
    work := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        work <- "Work completed"
    }()
    
    fmt.Println("\n=== Heartbeat Pattern ===")
    for {
        select {
        case <-heartbeat.C:
            fmt.Println("Heartbeat...")
        case msg := <-work:
            fmt.Printf("Work: %s\n", msg)
            return
        case <-time.After(5 * time.Second):
            fmt.Println("Timeout")
            return
        }
    }
}

func main() {
    advancedSelectPatterns()
}
```

## 🧪 Key Concepts to Remember

1. **Select waits on multiple channels**: Can handle multiple operations
2. **Random selection**: If multiple cases are ready, one is chosen randomly
3. **Timeout support**: Use `time.After()` for timeouts
4. **Default case**: Makes select non-blocking
5. **Nil channels**: Disable cases by setting channel to nil
6. **Can send and receive**: Both operations are supported

## 🎯 Common Use Cases

1. **Timeout handling**: Wait with a timeout
2. **Non-blocking operations**: Use default case
3. **Multiple channel processing**: Handle multiple channels
4. **Priority channels**: Check important channels first
5. **Heartbeat patterns**: Regular status checks

## 🎯 Common Pitfalls

1. **Infinite loop without exit condition**:
   ```go
   // BAD - infinite loop
   for {
       select {
       case msg := <-ch:
           // Process message
       }
   }
   
   // GOOD - have exit condition
   for {
       select {
       case msg, ok := <-ch:
           if !ok {
               return // Exit when channel closed
           }
       case <-time.After(5 * time.Second):
           return // Exit on timeout
       }
   }
   ```

2. **Not handling closed channels**:
   ```go
   // BAD - might get zero value
   case msg := <-ch:
   
   // GOOD - check if closed
   case msg, ok := <-ch:
       if !ok {
           ch = nil // Disable case
       }
   ```

3. **Deadlock with no default**:
   ```go
   // BAD - might deadlock
   select {
   case msg := <-ch: // If ch is never ready
   }
   
   // GOOD - have timeout or default
   select {
   case msg := <-ch:
   case <-time.After(1 * time.Second):
       // Handle timeout
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a server handling multiple client requests:
- Use select to handle requests from multiple clients
- Implement a timeout for each request
- Show how select can prioritize certain types of requests
- Handle the case when clients disconnect

**Hint**: Use different channels for different clients and use select to process them.

## 🚀 Next Steps

Now that you understand the select statement, let's learn about **channel directions** in the next file: `07-channel-directions.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
