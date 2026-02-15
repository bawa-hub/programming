# 20 - Fan-in and Fan-out Patterns

## 🎯 Learning Objectives
- Understand what fan-in and fan-out patterns are
- Learn how to implement fan-in and fan-out patterns
- Master different fan-in and fan-out scenarios
- Practice with fan-in and fan-out best practices
- Understand when to use these patterns

## 📚 Theory

### What are Fan-in and Fan-out Patterns?

**Fan-out** is a pattern where one input is distributed to multiple outputs (one-to-many).
**Fan-in** is a pattern where multiple inputs are collected into one output (many-to-one).

**Key characteristics:**
- **Fan-out**: Distribute work to multiple workers
- **Fan-in**: Collect results from multiple workers
- **Load balancing**: Distribute work evenly
- **Result aggregation**: Collect and combine results

### Why Use Fan-in and Fan-out?

**Benefits:**
1. **Parallel processing**: Process multiple items simultaneously
2. **Load balancing**: Distribute work evenly
3. **Scalability**: Easy to add/remove workers
4. **Result aggregation**: Combine results from multiple sources

**Use cases:**
1. **Data processing**: Process large datasets in parallel
2. **Web scraping**: Scrape multiple URLs simultaneously
3. **API calls**: Make multiple API calls in parallel
4. **File processing**: Process multiple files simultaneously

## 💻 Code Examples

### Example 1: Basic Fan-out Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicFanOut() {
    fmt.Println("=== Basic Fan-out Pattern ===")
    
    // Input channel
    input := make(chan int)
    
    // Multiple output channels
    output1 := make(chan int)
    output2 := make(chan int)
    output3 := make(chan int)
    
    // Fan-out: distribute input to multiple outputs
    go func() {
        defer close(output1)
        defer close(output2)
        defer close(output3)
        
        for num := range input {
            fmt.Printf("Fan-out: distributing %d\n", num)
            output1 <- num
            output2 <- num
            output3 <- num
        }
    }()
    
    // Send input data
    go func() {
        for i := 1; i <= 3; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Process outputs
    var wg sync.WaitGroup
    
    // Worker 1
    wg.Add(1)
    go func() {
        defer wg.Done()
        for num := range output1 {
            result := num * 2
            fmt.Printf("Worker 1: %d -> %d\n", num, result)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Worker 2
    wg.Add(1)
    go func() {
        defer wg.Done()
        for num := range output2 {
            result := num * 3
            fmt.Printf("Worker 2: %d -> %d\n", num, result)
            time.Sleep(300 * time.Millisecond)
        }
    }()
    
    // Worker 3
    wg.Add(1)
    go func() {
        defer wg.Done()
        for num := range output3 {
            result := num * 4
            fmt.Printf("Worker 3: %d -> %d\n", num, result)
            time.Sleep(400 * time.Millisecond)
        }
    }()
    
    wg.Wait()
}

func main() {
    basicFanOut()
}
```

**Run this code:**
```bash
go run 20-fan-in-fan-out.go
```

### Example 2: Basic Fan-in Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicFanIn() {
    fmt.Println("=== Basic Fan-in Pattern ===")
    
    // Multiple input channels
    input1 := make(chan int)
    input2 := make(chan int)
    input3 := make(chan int)
    
    // Single output channel
    output := make(chan int)
    
    // Fan-in: collect from multiple inputs
    go func() {
        defer close(output)
        
        for {
            select {
            case num, ok := <-input1:
                if !ok {
                    input1 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from input1\n", num)
                    output <- num
                }
            case num, ok := <-input2:
                if !ok {
                    input2 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from input2\n", num)
                    output <- num
                }
            case num, ok := <-input3:
                if !ok {
                    input3 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from input3\n", num)
                    output <- num
                }
            }
            
            // Exit when all inputs are closed
            if input1 == nil && input2 == nil && input3 == nil {
                break
            }
        }
    }()
    
    // Send data to inputs
    go func() {
        for i := 1; i <= 3; i++ {
            input1 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(input1)
    }()
    
    go func() {
        for i := 4; i <= 6; i++ {
            input2 <- i
            time.Sleep(150 * time.Millisecond)
        }
        close(input2)
    }()
    
    go func() {
        for i := 7; i <= 9; i++ {
            input3 <- i
            time.Sleep(200 * time.Millisecond)
        }
        close(input3)
    }()
    
    // Collect results
    for result := range output {
        fmt.Printf("Final result: %d\n", result)
    }
}

func main() {
    basicFanIn()
}
```

### Example 3: Fan-out with Worker Pool

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanOutWithWorkerPool() {
    fmt.Println("=== Fan-out with Worker Pool ===")
    
    // Input channel
    input := make(chan int, 10)
    
    // Output channel
    output := make(chan int, 10)
    
    // Number of workers
    numWorkers := 3
    
    // Start workers
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for num := range input {
                result := num * workerID
                fmt.Printf("Worker %d: %d -> %d\n", workerID, num, result)
                output <- result
                time.Sleep(200 * time.Millisecond)
            }
        }(i)
    }
    
    // Close output when all workers are done
    go func() {
        wg.Wait()
        close(output)
    }()
    
    // Send input data
    go func() {
        for i := 1; i <= 6; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Collect results
    for result := range output {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    fanOutWithWorkerPool()
}
```

### Example 4: Fan-in with Multiple Sources

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanInWithMultipleSources() {
    fmt.Println("=== Fan-in with Multiple Sources ===")
    
    // Multiple input channels
    inputs := make([]<-chan int, 3)
    for i := range inputs {
        ch := make(chan int)
        inputs[i] = ch
        
        // Send data to each input
        go func(id int, ch chan<- int) {
            for j := 1; j <= 3; j++ {
                value := id*10 + j
                ch <- value
                time.Sleep(time.Duration(id+1) * 100 * time.Millisecond)
            }
            close(ch)
        }(i, ch)
    }
    
    // Fan-in: collect from all inputs
    output := make(chan int)
    go func() {
        defer close(output)
        
        var wg sync.WaitGroup
        for _, input := range inputs {
            wg.Add(1)
            go func(ch <-chan int) {
                defer wg.Done()
                for num := range ch {
                    fmt.Printf("Fan-in: received %d\n", num)
                    output <- num
                }
            }(input)
        }
        wg.Wait()
    }()
    
    // Collect results
    for result := range output {
        fmt.Printf("Final result: %d\n", result)
    }
}

func main() {
    fanInWithMultipleSources()
}
```

### Example 5: Fan-out and Fan-in Combined

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanOutAndFanInCombined() {
    fmt.Println("=== Fan-out and Fan-in Combined ===")
    
    // Input channel
    input := make(chan int)
    
    // Fan-out: distribute to multiple workers
    worker1 := make(chan int)
    worker2 := make(chan int)
    worker3 := make(chan int)
    
    go func() {
        defer close(worker1)
        defer close(worker2)
        defer close(worker3)
        
        for num := range input {
            fmt.Printf("Fan-out: distributing %d\n", num)
            worker1 <- num
            worker2 <- num
            worker3 <- num
        }
    }()
    
    // Workers process data
    results1 := make(chan int)
    results2 := make(chan int)
    results3 := make(chan int)
    
    var wg sync.WaitGroup
    
    // Worker 1
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(results1)
        for num := range worker1 {
            result := num * 2
            fmt.Printf("Worker 1: %d -> %d\n", num, result)
            results1 <- result
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Worker 2
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(results2)
        for num := range worker2 {
            result := num * 3
            fmt.Printf("Worker 2: %d -> %d\n", num, result)
            results2 <- result
            time.Sleep(300 * time.Millisecond)
        }
    }()
    
    // Worker 3
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(results3)
        for num := range worker3 {
            result := num * 4
            fmt.Printf("Worker 3: %d -> %d\n", num, result)
            results3 <- result
            time.Sleep(400 * time.Millisecond)
        }
    }()
    
    // Fan-in: collect results
    output := make(chan int)
    go func() {
        defer close(output)
        
        for {
            select {
            case result, ok := <-results1:
                if !ok {
                    results1 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from worker1\n", result)
                    output <- result
                }
            case result, ok := <-results2:
                if !ok {
                    results2 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from worker2\n", result)
                    output <- result
                }
            case result, ok := <-results3:
                if !ok {
                    results3 = nil
                } else {
                    fmt.Printf("Fan-in: received %d from worker3\n", result)
                    output <- result
                }
            }
            
            // Exit when all results are closed
            if results1 == nil && results2 == nil && results3 == nil {
                break
            }
        }
    }()
    
    // Send input data
    go func() {
        for i := 1; i <= 3; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Wait for workers to complete
    go func() {
        wg.Wait()
    }()
    
    // Collect final results
    for result := range output {
        fmt.Printf("Final result: %d\n", result)
    }
}

func main() {
    fanOutAndFanInCombined()
}
```

### Example 6: Fan-out with Rate Limiting

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanOutWithRateLimiting() {
    fmt.Println("=== Fan-out with Rate Limiting ===")
    
    // Input channel
    input := make(chan int)
    
    // Output channel
    output := make(chan int)
    
    // Rate limiter
    rateLimiter := time.NewTicker(100 * time.Millisecond)
    defer rateLimiter.Stop()
    
    // Number of workers
    numWorkers := 3
    
    // Start workers
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for num := range input {
                <-rateLimiter.C // Wait for rate limiter
                result := num * workerID
                fmt.Printf("Worker %d: %d -> %d\n", workerID, num, result)
                output <- result
            }
        }(i)
    }
    
    // Close output when all workers are done
    go func() {
        wg.Wait()
        close(output)
    }()
    
    // Send input data
    go func() {
        for i := 1; i <= 6; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Collect results
    for result := range output {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    fanOutWithRateLimiting()
}
```

### Example 7: Fan-in with Error Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Result struct {
    Value int
    Error error
}

func fanInWithErrorHandling() {
    fmt.Println("=== Fan-in with Error Handling ===")
    
    // Multiple input channels
    input1 := make(chan int)
    input2 := make(chan int)
    input3 := make(chan int)
    
    // Single output channel
    output := make(chan Result)
    
    // Fan-in: collect from multiple inputs
    go func() {
        defer close(output)
        
        for {
            select {
            case num, ok := <-input1:
                if !ok {
                    input1 = nil
                } else {
                    result := processNumber(num, 1)
                    fmt.Printf("Fan-in: received result from input1: %v\n", result)
                    output <- result
                }
            case num, ok := <-input2:
                if !ok {
                    input2 = nil
                } else {
                    result := processNumber(num, 2)
                    fmt.Printf("Fan-in: received result from input2: %v\n", result)
                    output <- result
                }
            case num, ok := <-input3:
                if !ok {
                    input3 = nil
                } else {
                    result := processNumber(num, 3)
                    fmt.Printf("Fan-in: received result from input3: %v\n", result)
                    output <- result
                }
            }
            
            // Exit when all inputs are closed
            if input1 == nil && input2 == nil && input3 == nil {
                break
            }
        }
    }()
    
    // Send data to inputs
    go func() {
        for i := 1; i <= 3; i++ {
            input1 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(input1)
    }()
    
    go func() {
        for i := 4; i <= 6; i++ {
            input2 <- i
            time.Sleep(150 * time.Millisecond)
        }
        close(input2)
    }()
    
    go func() {
        for i := 7; i <= 9; i++ {
            input3 <- i
            time.Sleep(200 * time.Millisecond)
        }
        close(input3)
    }()
    
    // Collect results
    for result := range output {
        if result.Error != nil {
            fmt.Printf("Error: %v\n", result.Error)
        } else {
            fmt.Printf("Final result: %d\n", result.Value)
        }
    }
}

func processNumber(num, workerID int) Result {
    time.Sleep(time.Duration(workerID) * 100 * time.Millisecond)
    
    // Simulate error for certain numbers
    if num%3 == 0 {
        return Result{
            Value: 0,
            Error: fmt.Errorf("error processing %d", num),
        }
    }
    
    return Result{
        Value: num * workerID,
        Error: nil,
    }
}

func main() {
    fanInWithErrorHandling()
}
```

### Example 8: Fan-in and Fan-out Best Practices

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type FanOutFanIn struct {
    input  <-chan int
    output chan<- int
    ctx    context.Context
}

func NewFanOutFanIn(ctx context.Context, input <-chan int, output chan<- int) *FanOutFanIn {
    return &FanOutFanIn{
        input:  input,
        output: output,
        ctx:    ctx,
    }
}

func (f *FanOutFanIn) Process(numWorkers int) {
    // Fan-out: distribute to workers
    workerInputs := make([]chan int, numWorkers)
    for i := range workerInputs {
        workerInputs[i] = make(chan int)
    }
    
    // Distribute input to workers
    go func() {
        defer func() {
            for _, ch := range workerInputs {
                close(ch)
            }
        }()
        
        for num := range f.input {
            for _, ch := range workerInputs {
                select {
                case ch <- num:
                case <-f.ctx.Done():
                    return
                }
            }
        }
    }()
    
    // Fan-in: collect results from workers
    var wg sync.WaitGroup
    for i, workerInput := range workerInputs {
        wg.Add(1)
        go func(workerID int, input <-chan int) {
            defer wg.Done()
            for num := range input {
                select {
                case f.output <- num * (workerID + 1):
                case <-f.ctx.Done():
                    return
                }
            }
        }(i, workerInput)
    }
    
    // Close output when all workers are done
    go func() {
        wg.Wait()
        close(f.output)
    }()
}

func fanInFanOutBestPractices() {
    fmt.Println("=== Fan-in and Fan-out Best Practices ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // Create input and output channels
    input := make(chan int)
    output := make(chan int)
    
    // Create fan-out/fan-in processor
    processor := NewFanOutFanIn(ctx, input, output)
    
    // Start processing
    processor.Process(3)
    
    // Send input data
    go func() {
        for i := 1; i <= 6; i++ {
            select {
            case input <- i:
                fmt.Printf("Sent: %d\n", i)
            case <-ctx.Done():
                close(input)
                return
            }
            time.Sleep(200 * time.Millisecond)
        }
        close(input)
    }()
    
    // Collect results
    for result := range output {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    fanInFanOutBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Fan-out**: Distribute one input to multiple outputs
2. **Fan-in**: Collect multiple inputs into one output
3. **Load balancing**: Distribute work evenly
4. **Result aggregation**: Collect and combine results
5. **Concurrent processing**: Process multiple items simultaneously

## 🎯 When to Use Fan-in and Fan-out

1. **Parallel processing**: When you need to process multiple items simultaneously
2. **Load balancing**: When you need to distribute work evenly
3. **Result aggregation**: When you need to combine results from multiple sources
4. **Scalability**: When you need to scale processing

## 🎯 Best Practices

1. **Use appropriate number of workers**:
   ```go
   // GOOD - reasonable number of workers
   numWorkers := runtime.NumCPU()
   
   // BAD - too many workers
   numWorkers := 1000
   ```

2. **Handle errors properly**:
   ```go
   // GOOD - handle errors
   if err := processItem(item); err != nil {
       results <- Result{Error: err}
   }
   ```

3. **Use context for cancellation**:
   ```go
   // GOOD - use context
   select {
   case output <- result:
   case <-ctx.Done():
       return
   }
   ```

4. **Close channels properly**:
   ```go
   // GOOD - close channels
   defer close(output)
   ```

## 🎯 Common Pitfalls

1. **Not closing channels**:
   ```go
   // BAD - channels not closed
   func fanOut(input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           for num := range input {
               output <- num * 2
           }
           // Forgot to close output
       }()
       return output
   }
   ```

2. **Not handling errors**:
   ```go
   // BAD - ignore errors
   func process(input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           defer close(output)
           for num := range input {
               result := num * 2 // Might fail
               output <- result
           }
       }()
       return output
   }
   ```

3. **Not using context**:
   ```go
   // BAD - no cancellation
   func process(input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           defer close(output)
           for num := range input {
               output <- num * 2
           }
       }()
       return output
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a web scraper with fan-out and fan-in:
- Use fan-out to distribute URLs to multiple workers
- Each worker scrapes URLs and returns results
- Use fan-in to collect results from all workers
- Handle errors and timeouts properly
- Show how fan-out and fan-in can improve performance

**Hint**: Use channels to distribute URLs and collect results, and implement proper error handling and cancellation.

## 🚀 Next Steps

Now that you understand fan-in and fan-out patterns, let's learn about **error handling in concurrent code** in the next file: `21-error-handling.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
