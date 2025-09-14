# 11 - Race Conditions

## 🎯 Learning Objectives
- Understand what race conditions are and why they occur
- Learn how to detect race conditions in Go
- Master the Go race detector
- Practice identifying and fixing race conditions
- Understand the impact of race conditions

## 📚 Theory

### What is a Race Condition?

A **race condition** occurs when two or more goroutines access shared data concurrently, and at least one of them modifies the data. The outcome depends on the timing of the operations.

**Key characteristics:**
- **Non-deterministic**: Results vary between runs
- **Timing-dependent**: Depends on execution order
- **Hard to reproduce**: May work most of the time
- **Data corruption**: Can lead to incorrect results

### Why Do Race Conditions Occur?

1. **Shared data**: Multiple goroutines access the same variable
2. **Concurrent access**: At least one goroutine modifies the data
3. **No synchronization**: No protection against concurrent access
4. **Timing**: Operations happen in unpredictable order

### Go Race Detector

Go provides a built-in race detector that can detect race conditions at runtime:
- **Enable**: `go run -race` or `go test -race`
- **Performance**: Slows down execution significantly
- **Memory**: Uses more memory
- **Accuracy**: Very good at detecting races

## 💻 Code Examples

### Example 1: Basic Race Condition

```go
package main

import (
    "fmt"
    "time"
)

func basicRaceCondition() {
    fmt.Println("=== Basic Race Condition ===")
    
    counter := 0
    
    // Multiple goroutines increment counter
    for i := 0; i < 1000; i++ {
        go func() {
            counter++ // Race condition!
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d (should be 1000)\n", counter)
}

func main() {
    basicRaceCondition()
}
```

**Run this code:**
```bash
# Run without race detection
go run 11-race-conditions.go

# Run with race detection
go run -race 11-race-conditions.go
```

### Example 2: Race Condition with Shared Data

```go
package main

import (
    "fmt"
    "time"
)

func raceWithSharedData() {
    fmt.Println("=== Race Condition with Shared Data ===")
    
    var shared int
    
    // Goroutine 1: increments
    go func() {
        for i := 0; i < 1000; i++ {
            shared++ // Race condition!
        }
    }()
    
    // Goroutine 2: decrements
    go func() {
        for i := 0; i < 1000; i++ {
            shared-- // Race condition!
        }
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Shared value: %d (should be 0)\n", shared)
}

func main() {
    raceWithSharedData()
}
```

### Example 3: Race Condition with Slices

```go
package main

import (
    "fmt"
    "time"
)

func raceWithSlices() {
    fmt.Println("=== Race Condition with Slices ===")
    
    slice := make([]int, 0)
    
    // Multiple goroutines append to slice
    for i := 0; i < 100; i++ {
        go func(i int) {
            slice = append(slice, i) // Race condition!
        }(i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Slice length: %d (should be 100)\n", len(slice))
    fmt.Printf("First few elements: %v\n", slice[:min(10, len(slice))])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    raceWithSlices()
}
```

### Example 4: Race Condition with Maps

```go
package main

import (
    "fmt"
    "time"
)

func raceWithMaps() {
    fmt.Println("=== Race Condition with Maps ===")
    
    m := make(map[string]int)
    
    // Multiple goroutines write to map
    for i := 0; i < 100; i++ {
        go func(i int) {
            key := fmt.Sprintf("key%d", i)
            m[key] = i // Race condition!
        }(i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Map size: %d (should be 100)\n", len(m))
}

func main() {
    raceWithMaps()
}
```

### Example 5: Race Condition with Structs

```go
package main

import (
    "fmt"
    "time"
)

type Counter struct {
    value int
}

func (c *Counter) Increment() {
    c.value++ // Race condition!
}

func (c *Counter) Value() int {
    return c.value
}

func raceWithStructs() {
    fmt.Println("=== Race Condition with Structs ===")
    
    counter := &Counter{}
    
    // Multiple goroutines increment counter
    for i := 0; i < 1000; i++ {
        go func() {
            counter.Increment() // Race condition!
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d (should be 1000)\n", counter.Value())
}

func main() {
    raceWithStructs()
}
```

### Example 6: Race Condition Detection

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateRaceDetection() {
    fmt.Println("=== Race Condition Detection ===")
    
    var counter int
    
    // This will trigger race detection
    go func() {
        for i := 0; i < 1000; i++ {
            counter++
        }
    }()
    
    go func() {
        for i := 0; i < 1000; i++ {
            counter++
        }
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter: %d\n", counter)
}

func main() {
    demonstrateRaceDetection()
}
```

### Example 7: Race Condition with Channels

```go
package main

import (
    "fmt"
    "time"
)

func raceWithChannels() {
    fmt.Println("=== Race Condition with Channels ===")
    
    ch := make(chan int, 1)
    var counter int
    
    // Goroutine 1: sends values
    go func() {
        for i := 0; i < 100; i++ {
            ch <- i
        }
    }()
    
    // Goroutine 2: receives values and increments counter
    go func() {
        for i := 0; i < 100; i++ {
            <-ch
            counter++ // Race condition!
        }
    }()
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter: %d (should be 100)\n", counter)
}

func main() {
    raceWithChannels()
}
```

### Example 8: Race Condition with Global Variables

```go
package main

import (
    "fmt"
    "time"
)

var globalCounter int

func incrementGlobal() {
    globalCounter++ // Race condition!
}

func raceWithGlobalVariables() {
    fmt.Println("=== Race Condition with Global Variables ===")
    
    // Multiple goroutines access global variable
    for i := 0; i < 1000; i++ {
        go incrementGlobal()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Global counter: %d (should be 1000)\n", globalCounter)
}

func main() {
    raceWithGlobalVariables()
}
```

## 🧪 Key Concepts to Remember

1. **Race conditions are non-deterministic**: Results vary between runs
2. **Use race detector**: Always run with `go run -race`
3. **Shared data is dangerous**: Multiple goroutines accessing same data
4. **Timing matters**: Race conditions depend on execution order
5. **Hard to reproduce**: May work most of the time

## 🎯 How to Detect Race Conditions

1. **Use Go race detector**:
   ```bash
   go run -race main.go
   go test -race
   ```

2. **Look for warning messages**:
   ```
   WARNING: DATA RACE
   Read at 0x00c0000140a8 by goroutine 7:
   Previous write at 0x00c0000140a8 by goroutine 6:
   ```

3. **Test multiple times**: Race conditions are timing-dependent

## 🎯 Common Sources of Race Conditions

1. **Global variables**: Multiple goroutines accessing globals
2. **Shared structs**: Concurrent access to struct fields
3. **Slices and maps**: Concurrent modifications
4. **Channels**: Improper channel usage
5. **Pointers**: Concurrent pointer dereferencing

## 🎯 Impact of Race Conditions

1. **Data corruption**: Incorrect values
2. **Crashes**: Segmentation faults
3. **Inconsistent state**: Unpredictable behavior
4. **Security issues**: Data leaks
5. **Hard to debug**: Non-deterministic

## 🎯 Best Practices

1. **Always use race detector**:
   ```bash
   go run -race main.go
   go test -race
   ```

2. **Avoid shared data**: Use channels for communication
3. **Use synchronization**: Mutexes, atomic operations
4. **Test thoroughly**: Run tests multiple times
5. **Code review**: Look for shared data access

## 🎯 Common Pitfalls

1. **Not using race detector**:
   ```go
   // BAD - might miss race conditions
   go run main.go
   
   // GOOD - always use race detector
   go run -race main.go
   ```

2. **Ignoring race warnings**:
   ```go
   // BAD - ignore race detector warnings
   // WARNING: DATA RACE
   
   // GOOD - fix race conditions
   // Use mutexes or channels
   ```

3. **Thinking it works**: Race conditions may work most of the time

## 🏋️ Exercise

**Problem**: Create a program that demonstrates different types of race conditions:
- Race condition with integers
- Race condition with slices
- Race condition with maps
- Race condition with structs

Run each example with and without the race detector to see the difference.

**Hint**: Use the race detector to identify the exact lines where race conditions occur.

## 🚀 Next Steps

Now that you understand race conditions, let's learn about **mutexes** to fix them in the next file: `12-mutex-basics.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
