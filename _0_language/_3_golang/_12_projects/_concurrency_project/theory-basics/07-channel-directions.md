# 07 - Channel Directions

## 🎯 Learning Objectives
- Understand channel directions and their purpose
- Learn about send-only and receive-only channels
- Master function parameter channel directions
- Practice with channel direction patterns
- Understand when and why to use different directions

## 📚 Theory

### What are Channel Directions?

Channel directions specify what operations are allowed on a channel:
- **Bidirectional**: `chan T` - can both send and receive
- **Send-only**: `chan<- T` - can only send
- **Receive-only**: `<-chan T` - can only receive

### Why Use Channel Directions?

**Benefits:**
1. **Type Safety**: Prevents accidental wrong operations
2. **Clear Intent**: Makes code more readable
3. **Compile-time Checks**: Catches errors at compile time
4. **Better APIs**: Functions express their channel usage clearly

### Channel Direction Syntax

```go
// Bidirectional (default)
var ch chan int

// Send-only
var sendCh chan<- int

// Receive-only
var recvCh <-chan int

// Function parameters
func sender(ch chan<- int) { ... }
func receiver(ch <-chan int) { ... }
```

## 💻 Code Examples

### Example 1: Basic Channel Directions

```go
package main

import (
    "fmt"
    "time"
)

func basicDirections() {
    fmt.Println("=== Basic Channel Directions ===")
    
    // Create bidirectional channel
    ch := make(chan string)
    
    // Convert to send-only
    var sendCh chan<- string = ch
    
    // Convert to receive-only
    var recvCh <-chan string = ch
    
    // Send-only can only send
    go func() {
        sendCh <- "Hello from send-only"
        sendCh <- "Another message"
        close(ch) // Can close from send-only
    }()
    
    // Receive-only can only receive
    for msg := range recvCh {
        fmt.Printf("Received: %s\n", msg)
    }
}

func main() {
    basicDirections()
}
```

**Run this code:**
```bash
go run 07-channel-directions.go
```

### Example 2: Function Parameters with Directions

```go
package main

import (
    "fmt"
    "time"
)

// Function that only sends to channel
func producer(ch chan<- string) {
    messages := []string{"Hello", "World", "from", "Go"}
    
    for _, msg := range messages {
        fmt.Printf("Producing: %s\n", msg)
        ch <- msg
        time.Sleep(300 * time.Millisecond)
    }
    close(ch) // Can close from send-only
}

// Function that only receives from channel
func consumer(ch <-chan string) {
    for msg := range ch {
        fmt.Printf("Consuming: %s\n", msg)
        time.Sleep(200 * time.Millisecond)
    }
    fmt.Println("Consumer finished")
}

// Function that processes data between channels
func processor(input <-chan string, output chan<- string) {
    for msg := range input {
        processed := fmt.Sprintf("Processed: %s", msg)
        fmt.Printf("Processing: %s -> %s\n", msg, processed)
        output <- processed
        time.Sleep(100 * time.Millisecond)
    }
    close(output) // Close output when done
}

func main() {
    fmt.Println("=== Function Parameters with Directions ===")
    
    // Create channels
    input := make(chan string, 5)
    output := make(chan string, 5)
    
    // Start goroutines with specific directions
    go producer(input)
    go processor(input, output)
    go consumer(output)
    
    // Wait for all to complete
    time.Sleep(3 * time.Second)
}
```

### Example 3: Channel Direction Safety

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateDirectionSafety() {
    fmt.Println("=== Channel Direction Safety ===")
    
    ch := make(chan int)
    
    // This function can only send
    sendOnly := func(ch chan<- int) {
        ch <- 42
        ch <- 24
        close(ch)
        // ch <- 10 // This would cause compile error
    }
    
    // This function can only receive
    receiveOnly := func(ch <-chan int) {
        for value := range ch {
            fmt.Printf("Received: %d\n", value)
        }
        // ch <- 10 // This would cause compile error
    }
    
    go sendOnly(ch)
    receiveOnly(ch)
}

func main() {
    demonstrateDirectionSafety()
}
```

### Example 4: Pipeline with Channel Directions

```go
package main

import (
    "fmt"
    "time"
)

// Stage 1: Generate numbers
func generator(output chan<- int) {
    defer close(output)
    
    for i := 1; i <= 5; i++ {
        fmt.Printf("Generating: %d\n", i)
        output <- i
        time.Sleep(200 * time.Millisecond)
    }
}

// Stage 2: Square numbers
func squarer(input <-chan int, output chan<- int) {
    defer close(output)
    
    for value := range input {
        result := value * value
        fmt.Printf("Squaring: %d^2 = %d\n", value, result)
        output <- result
        time.Sleep(300 * time.Millisecond)
    }
}

// Stage 3: Add 10
func adder(input <-chan int, output chan<- int) {
    defer close(output)
    
    for value := range input {
        result := value + 10
        fmt.Printf("Adding: %d + 10 = %d\n", value, result)
        output <- result
        time.Sleep(200 * time.Millisecond)
    }
}

// Stage 4: Print results
func printer(input <-chan int) {
    for value := range input {
        fmt.Printf("Final result: %d\n", value)
    }
}

func main() {
    fmt.Println("=== Pipeline with Channel Directions ===")
    
    // Create channels for each stage
    numbers := make(chan int)
    squares := make(chan int)
    results := make(chan int)
    
    // Start pipeline stages
    go generator(numbers)
    go squarer(numbers, squares)
    go adder(squares, results)
    go printer(results)
    
    // Wait for pipeline to complete
    time.Sleep(5 * time.Second)
}
```

### Example 5: Worker Pool with Directions

```go
package main

import (
    "fmt"
    "time"
)

// Job represents work to be done
type Job struct {
    ID     int
    Data   string
    Result chan<- string // Send-only channel for results
}

// Worker processes jobs
func worker(id int, jobs <-chan Job) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d: %s\n", id, job.ID, job.Data)
        
        // Simulate work
        time.Sleep(time.Duration(job.ID) * 200 * time.Millisecond)
        
        // Send result back
        result := fmt.Sprintf("Job %d completed by worker %d", job.ID, id)
        job.Result <- result
    }
}

// Job dispatcher
func dispatcher(jobs chan<- Job, results <-chan string) {
    // Create some jobs
    for i := 1; i <= 5; i++ {
        resultCh := make(chan string, 1)
        job := Job{
            ID:     i,
            Data:   fmt.Sprintf("Data for job %d", i),
            Result: resultCh,
        }
        
        jobs <- job
        
        // Wait for result
        result := <-resultCh
        fmt.Printf("Dispatcher received: %s\n", result)
    }
    close(jobs)
}

func main() {
    fmt.Println("=== Worker Pool with Directions ===")
    
    // Create channels
    jobs := make(chan Job, 10)
    results := make(chan string, 10)
    
    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs)
    }
    
    // Start dispatcher
    go dispatcher(jobs, results)
    
    // Wait for completion
    time.Sleep(3 * time.Second)
}
```

### Example 6: Channel Direction Conversion

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateConversion() {
    fmt.Println("=== Channel Direction Conversion ===")
    
    // Start with bidirectional channel
    ch := make(chan string)
    
    // Convert to send-only
    sendCh := (chan<- string)(ch)
    
    // Convert to receive-only
    recvCh := (<-chan string)(ch)
    
    // Send using send-only
    go func() {
        sendCh <- "Message 1"
        sendCh <- "Message 2"
        close(ch)
    }()
    
    // Receive using receive-only
    for msg := range recvCh {
        fmt.Printf("Received: %s\n", msg)
    }
}

func main() {
    demonstrateConversion()
}
```

### Example 7: Advanced Direction Patterns

```go
package main

import (
    "fmt"
    "time"
)

// Fan-out: distribute work to multiple workers
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

// Fan-in: collect results from multiple workers
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

// Worker that processes data
func worker(id int, input <-chan int, output chan<- int) {
    defer close(output)
    
    for value := range input {
        result := value * value
        fmt.Printf("Worker %d: %d^2 = %d\n", id, value, result)
        output <- result
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    fmt.Println("=== Advanced Direction Patterns ===")
    
    // Create channels
    input := make(chan int)
    workerInputs := make([]chan int, 3)
    workerOutputs := make([]chan int, 3)
    
    for i := range workerInputs {
        workerInputs[i] = make(chan int)
        workerOutputs[i] = make(chan int)
    }
    
    finalOutput := make(chan int)
    
    // Start fan-out
    go fanOut(input, workerInputs)
    
    // Start workers
    for i := range workerInputs {
        go worker(i+1, workerInputs[i], workerOutputs[i])
    }
    
    // Start fan-in
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

## 🧪 Key Concepts to Remember

1. **Send-only**: `chan<- T` - can only send and close
2. **Receive-only**: `<-chan T` - can only receive
3. **Bidirectional**: `chan T` - can both send and receive
4. **Type safety**: Prevents wrong operations at compile time
5. **Clear intent**: Makes code more readable
6. **Function parameters**: Express channel usage clearly

## 🎯 When to Use Which Direction

### Use Send-only (`chan<- T`) when:
- Function only sends data
- Function is a producer
- Function closes the channel

### Use Receive-only (`<-chan T`) when:
- Function only receives data
- Function is a consumer
- Function reads until channel is closed

### Use Bidirectional (`chan T`) when:
- Function needs both send and receive
- Function is a processor/transformer
- Function manages the channel lifecycle

## 🎯 Common Pitfalls

1. **Trying to send on receive-only channel**:
   ```go
   // BAD - compile error
   func consumer(ch <-chan int) {
       ch <- 42 // Error!
   }
   
   // GOOD - only receive
   func consumer(ch <-chan int) {
       value := <-ch
   }
   ```

2. **Trying to receive on send-only channel**:
   ```go
   // BAD - compile error
   func producer(ch chan<- int) {
       value := <-ch // Error!
   }
   
   // GOOD - only send
   func producer(ch chan<- int) {
       ch <- 42
   }
   ```

3. **Not understanding direction conversion**:
   ```go
   // BAD - can't convert back
   var sendCh chan<- int = make(chan int)
   var recvCh <-chan int = sendCh // Error!
   
   // GOOD - convert from bidirectional
   ch := make(chan int)
   var sendCh chan<- int = ch
   var recvCh <-chan int = ch
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a data processing pipeline:
- Stage 1: Generate random numbers (send-only output)
- Stage 2: Filter even numbers (receive-only input, send-only output)
- Stage 3: Square the numbers (receive-only input, send-only output)
- Stage 4: Print results (receive-only input)

Use proper channel directions for each stage and show how they prevent accidental wrong operations.

## 🚀 Next Steps

Now that you understand channel directions, let's learn about **buffered vs unbuffered channels** in the next file: `08-buffered-vs-unbuffered.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
