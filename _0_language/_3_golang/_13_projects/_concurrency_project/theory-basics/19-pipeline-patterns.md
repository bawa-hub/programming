# 19 - Pipeline Patterns

## 🎯 Learning Objectives
- Understand what pipeline patterns are and why they're useful
- Learn how to implement different pipeline patterns
- Master fan-in and fan-out patterns
- Practice with pipeline best practices
- Understand when to use pipeline patterns

## 📚 Theory

### What are Pipeline Patterns?

**Pipeline patterns** are a way to process data through multiple stages, where each stage transforms the data and passes it to the next stage.

**Key characteristics:**
- **Multiple stages**: Data flows through multiple processing stages
- **Transformation**: Each stage transforms the data
- **Concurrent processing**: Stages can run concurrently
- **Streaming**: Data flows continuously through the pipeline

### Why Use Pipeline Patterns?

**Benefits:**
1. **Modularity**: Each stage is independent
2. **Concurrency**: Stages can run in parallel
3. **Scalability**: Easy to add/remove stages
4. **Reusability**: Stages can be reused
5. **Testability**: Each stage can be tested independently

**Use cases:**
1. **Data processing**: ETL pipelines
2. **Image processing**: Multiple transformation stages
3. **Text processing**: Multiple analysis stages
4. **Stream processing**: Real-time data processing

## 💻 Code Examples

### Example 1: Basic Pipeline Pattern

```go
package main

import (
    "fmt"
    "time"
)

func basicPipeline() {
    fmt.Println("=== Basic Pipeline Pattern ===")
    
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        for num := range numbers {
            square := num * num
            fmt.Printf("Squaring: %d -> %d\n", num, square)
            squares <- square
        }
        close(squares)
    }()
    
    // Stage 3: Add 10
    results := make(chan int)
    go func() {
        for square := range squares {
            result := square + 10
            fmt.Printf("Adding 10: %d -> %d\n", square, result)
            results <- result
        }
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Final result: %d\n", result)
    }
}

func main() {
    basicPipeline()
}
```

**Run this code:**
```bash
go run 19-pipeline-patterns.go
```

### Example 2: Pipeline with Multiple Stages

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

func pipelineWithMultipleStages() {
    fmt.Println("=== Pipeline with Multiple Stages ===")
    
    // Stage 1: Generate strings
    strings := make(chan string)
    go func() {
        words := []string{"hello", "world", "golang", "concurrency", "pipeline"}
        for _, word := range words {
            strings <- word
        }
        close(strings)
    }()
    
    // Stage 2: Convert to uppercase
    uppercase := make(chan string)
    go func() {
        for str := range strings {
            upper := strings.ToUpper(str)
            fmt.Printf("Uppercase: %s -> %s\n", str, upper)
            uppercase <- upper
        }
        close(uppercase)
    }()
    
    // Stage 3: Add prefix
    prefixed := make(chan string)
    go func() {
        for str := range uppercase {
            prefixedStr := "PREFIX_" + str
            fmt.Printf("Add prefix: %s -> %s\n", str, prefixedStr)
            prefixed <- prefixedStr
        }
        close(prefixed)
    }()
    
    // Stage 4: Add suffix
    results := make(chan string)
    go func() {
        for str := range prefixed {
            result := str + "_SUFFIX"
            fmt.Printf("Add suffix: %s -> %s\n", str, result)
            results <- result
        }
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Final result: %s\n", result)
    }
}

func main() {
    pipelineWithMultipleStages()
}
```

### Example 3: Fan-out Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanOutPattern() {
    fmt.Println("=== Fan-out Pattern ===")
    
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
    fanOutPattern()
}
```

### Example 4: Fan-in Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func fanInPattern() {
    fmt.Println("=== Fan-in Pattern ===")
    
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
    fanInPattern()
}
```

### Example 5: Pipeline with Error Handling

```go
package main

import (
    "fmt"
    "time"
)

type PipelineResult struct {
    Value int
    Error error
}

func pipelineWithErrorHandling() {
    fmt.Println("=== Pipeline with Error Handling ===")
    
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    // Stage 2: Square numbers (with error simulation)
    squares := make(chan PipelineResult)
    go func() {
        defer close(squares)
        for num := range numbers {
            if num%3 == 0 {
                squares <- PipelineResult{
                    Value: 0,
                    Error: fmt.Errorf("error processing %d", num),
                }
            } else {
                square := num * num
                fmt.Printf("Squaring: %d -> %d\n", num, square)
                squares <- PipelineResult{
                    Value: square,
                    Error: nil,
                }
            }
        }
    }()
    
    // Stage 3: Add 10 (only if no error)
    results := make(chan PipelineResult)
    go func() {
        defer close(results)
        for result := range squares {
            if result.Error != nil {
                fmt.Printf("Error in pipeline: %v\n", result.Error)
                results <- result
            } else {
                finalResult := result.Value + 10
                fmt.Printf("Adding 10: %d -> %d\n", result.Value, finalResult)
                results <- PipelineResult{
                    Value: finalResult,
                    Error: nil,
                }
            }
        }
    }()
    
    // Collect results
    for result := range results {
        if result.Error != nil {
            fmt.Printf("Final error: %v\n", result.Error)
        } else {
            fmt.Printf("Final result: %d\n", result.Value)
        }
    }
}

func main() {
    pipelineWithErrorHandling()
}
```

### Example 6: Pipeline with Context

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func pipelineWithContext() {
    fmt.Println("=== Pipeline with Context ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        defer close(numbers)
        for i := 1; i <= 10; i++ {
            select {
            case numbers <- i:
                fmt.Printf("Generated: %d\n", i)
            case <-ctx.Done():
                fmt.Printf("Generator cancelled: %v\n", ctx.Err())
                return
            }
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Stage 2: Process numbers
    results := make(chan int)
    go func() {
        defer close(results)
        for {
            select {
            case num, ok := <-numbers:
                if !ok {
                    return
                }
                result := num * num
                fmt.Printf("Processed: %d -> %d\n", num, result)
                results <- result
            case <-ctx.Done():
                fmt.Printf("Processor cancelled: %v\n", ctx.Err())
                return
            }
        }
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Final result: %d\n", result)
    }
}

func main() {
    pipelineWithContext()
}
```

### Example 7: Pipeline Best Practices

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Pipeline struct {
    stages []func(<-chan interface{}) <-chan interface{}
}

func NewPipeline() *Pipeline {
    return &Pipeline{
        stages: make([]func(<-chan interface{}) <-chan interface{}, 0),
    }
}

func (p *Pipeline) AddStage(stage func(<-chan interface{}) <-chan interface{}) {
    p.stages = append(p.stages, stage)
}

func (p *Pipeline) Run(input <-chan interface{}) <-chan interface{} {
    current := input
    for _, stage := range p.stages {
        current = stage(current)
    }
    return current
}

func pipelineBestPractices() {
    fmt.Println("=== Pipeline Best Practices ===")
    
    // Create pipeline
    pipeline := NewPipeline()
    
    // Add stages
    pipeline.AddStage(func(input <-chan interface{}) <-chan interface{} {
        output := make(chan interface{})
        go func() {
            defer close(output)
            for value := range input {
                if num, ok := value.(int); ok {
                    result := num * 2
                    fmt.Printf("Stage 1: %d -> %d\n", num, result)
                    output <- result
                }
            }
        }()
        return output
    })
    
    pipeline.AddStage(func(input <-chan interface{}) <-chan interface{} {
        output := make(chan interface{})
        go func() {
            defer close(output)
            for value := range input {
                if num, ok := value.(int); ok {
                    result := num + 10
                    fmt.Printf("Stage 2: %d -> %d\n", num, result)
                    output <- result
                }
            }
        }()
        return output
    })
    
    // Create input
    input := make(chan interface{})
    go func() {
        defer close(input)
        for i := 1; i <= 5; i++ {
            input <- i
        }
    }()
    
    // Run pipeline
    output := pipeline.Run(input)
    
    // Collect results
    for result := range output {
        fmt.Printf("Final result: %v\n", result)
    }
}

func main() {
    pipelineBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Multiple stages**: Data flows through multiple processing stages
2. **Transformation**: Each stage transforms the data
3. **Concurrent processing**: Stages can run concurrently
4. **Streaming**: Data flows continuously through the pipeline
5. **Modularity**: Each stage is independent

## 🎯 When to Use Pipeline Patterns

1. **Data processing**: When you need to process data through multiple stages
2. **Stream processing**: When you need to process data continuously
3. **Modular design**: When you want to break down complex processing
4. **Concurrent processing**: When you want to process data concurrently

## 🎯 Best Practices

1. **Keep stages simple**:
   ```go
   // GOOD - simple stage
   func square(input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           defer close(output)
           for num := range input {
               output <- num * num
           }
       }()
       return output
   }
   ```

2. **Handle errors properly**:
   ```go
   // GOOD - handle errors
   func process(input <-chan int) <-chan Result {
       output := make(chan Result)
       go func() {
           defer close(output)
           for num := range input {
               if err := validate(num); err != nil {
                   output <- Result{Error: err}
               } else {
                   output <- Result{Value: num * 2}
               }
           }
       }()
       return output
   }
   ```

3. **Use context for cancellation**:
   ```go
   // GOOD - use context
   func process(ctx context.Context, input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           defer close(output)
           for {
               select {
               case num := <-input:
                   output <- num * 2
               case <-ctx.Done():
                   return
               }
           }
       }()
       return output
   }
   ```

## 🎯 Common Pitfalls

1. **Not closing channels**:
   ```go
   // BAD - channels not closed
   func process(input <-chan int) <-chan int {
       output := make(chan int)
       go func() {
           for num := range input {
               output <- num * 2
           }
           // Forgot to close output
       }()
       return output
   }
   
   // GOOD - close channels
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

**Problem**: Create a program that simulates a data processing pipeline:
- Stage 1: Read data from a source
- Stage 2: Validate the data
- Stage 3: Transform the data
- Stage 4: Store the data
- Use fan-out to process data in parallel
- Use fan-in to collect results
- Handle errors and timeouts properly

**Hint**: Use channels to connect the stages and implement proper error handling and cancellation.

## 🚀 Next Steps

Now that you understand pipeline patterns, let's learn about **fan-in and fan-out patterns** in the next file: `20-fan-in-fan-out.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
