# 16 - Atomic Operations

## 🎯 Learning Objectives
- Understand what atomic operations are and why they're useful
- Learn how to use the `sync/atomic` package
- Master different atomic operations (Add, Load, Store, etc.)
- Practice with atomic operations for lock-free programming
- Understand when to use atomic operations vs mutexes

## 📚 Theory

### What are Atomic Operations?

**Atomic operations** are operations that are performed as a single, indivisible unit. They cannot be interrupted by other goroutines.

**Key characteristics:**
- **Indivisible**: Cannot be interrupted
- **Lock-free**: No mutexes needed
- **Fast**: Very efficient
- **Limited**: Only work with specific types

### Why Use Atomic Operations?

**Benefits:**
1. **Performance**: Faster than mutexes
2. **Lock-free**: No deadlock risk
3. **Simplicity**: Simple operations
4. **Efficiency**: Very low overhead

**Limitations:**
1. **Limited types**: Only works with specific types
2. **Simple operations**: Only basic operations
3. **No complex logic**: Cannot do complex operations

### Atomic Types

Go provides atomic operations for:
- `int32`
- `int64`
- `uint32`
- `uint64`
- `uintptr`
- `unsafe.Pointer`

## 💻 Code Examples

### Example 1: Basic Atomic Operations

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func basicAtomicOperations() {
    fmt.Println("=== Basic Atomic Operations ===")
    
    var counter int64
    
    // Multiple goroutines increment counter atomically
    for i := 0; i < 1000; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d (should be 1000)\n", counter)
}

func main() {
    basicAtomicOperations()
}
```

**Run this code:**
```bash
go run 16-atomic-operations.go
```

### Example 2: Atomic Load and Store

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func atomicLoadAndStore() {
    fmt.Println("=== Atomic Load and Store ===")
    
    var value int64
    
    // Goroutine 1: stores values
    go func() {
        for i := 1; i <= 5; i++ {
            atomic.StoreInt64(&value, int64(i))
            fmt.Printf("Stored: %d\n", i)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Goroutine 2: loads values
    go func() {
        for i := 0; i < 10; i++ {
            current := atomic.LoadInt64(&value)
            fmt.Printf("Loaded: %d\n", current)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    atomicLoadAndStore()
}
```

### Example 3: Atomic Compare and Swap

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func atomicCompareAndSwap() {
    fmt.Println("=== Atomic Compare and Swap ===")
    
    var value int64 = 10
    
    // Multiple goroutines try to update value
    for i := 1; i <= 3; i++ {
        go func(id int) {
            for j := 0; j < 3; j++ {
                oldValue := atomic.LoadInt64(&value)
                newValue := oldValue + 1
                
                if atomic.CompareAndSwapInt64(&value, oldValue, newValue) {
                    fmt.Printf("Goroutine %d: updated %d -> %d\n", id, oldValue, newValue)
                } else {
                    fmt.Printf("Goroutine %d: failed to update (value changed)\n", id)
                }
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Final value: %d\n", atomic.LoadInt64(&value))
}

func main() {
    atomicCompareAndSwap()
}
```

### Example 4: Atomic Operations vs Mutex

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func atomicVsMutex() {
    fmt.Println("=== Atomic Operations vs Mutex ===")
    
    const iterations = 1000000
    
    // Test with atomic operations
    start := time.Now()
    var atomicCounter int64
    for i := 0; i < iterations; i++ {
        atomic.AddInt64(&atomicCounter, 1)
    }
    atomicTime := time.Since(start)
    
    // Test with mutex
    start = time.Now()
    var mutexCounter int64
    var mutex sync.Mutex
    for i := 0; i < iterations; i++ {
        mutex.Lock()
        mutexCounter++
        mutex.Unlock()
    }
    mutexTime := time.Since(start)
    
    fmt.Printf("Atomic time: %v\n", atomicTime)
    fmt.Printf("Mutex time: %v\n", mutexTime)
    fmt.Printf("Atomic speedup: %.2fx\n", float64(mutexTime)/float64(atomicTime))
}

func main() {
    atomicVsMutex()
}
```

### Example 5: Atomic Operations with Multiple Goroutines

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func atomicWithMultipleGoroutines() {
    fmt.Println("=== Atomic Operations with Multiple Goroutines ===")
    
    var counter int64
    
    // Multiple goroutines increment counter
    for i := 0; i < 10; i++ {
        go func(id int) {
            for j := 0; j < 100; j++ {
                atomic.AddInt64(&counter, 1)
                if j%20 == 0 {
                    fmt.Printf("Goroutine %d: counter = %d\n", id, atomic.LoadInt64(&counter))
                }
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Final counter: %d\n", atomic.LoadInt64(&counter))
}

func main() {
    atomicWithMultipleGoroutines()
}
```

### Example 6: Atomic Operations with Different Types

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func atomicWithDifferentTypes() {
    fmt.Println("=== Atomic Operations with Different Types ===")
    
    // int32
    var int32Value int32
    atomic.StoreInt32(&int32Value, 42)
    fmt.Printf("int32 value: %d\n", atomic.LoadInt32(&int32Value))
    
    // int64
    var int64Value int64
    atomic.StoreInt64(&int64Value, 123456789)
    fmt.Printf("int64 value: %d\n", atomic.LoadInt64(&int64Value))
    
    // uint32
    var uint32Value uint32
    atomic.StoreUint32(&uint32Value, 100)
    fmt.Printf("uint32 value: %d\n", atomic.LoadUint32(&uint32Value))
    
    // uint64
    var uint64Value uint64
    atomic.StoreUint64(&uint64Value, 987654321)
    fmt.Printf("uint64 value: %d\n", atomic.LoadUint64(&uint64Value))
}

func main() {
    atomicWithDifferentTypes()
}
```

### Example 7: Atomic Operations for Flags

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func atomicFlags() {
    fmt.Println("=== Atomic Operations for Flags ===")
    
    var flag int32 // 0 = false, 1 = true
    
    // Goroutine 1: sets flag
    go func() {
        time.Sleep(500 * time.Millisecond)
        atomic.StoreInt32(&flag, 1)
        fmt.Println("Flag set to true")
    }()
    
    // Goroutine 2: checks flag
    go func() {
        for i := 0; i < 10; i++ {
            if atomic.LoadInt32(&flag) == 1 {
                fmt.Println("Flag is true")
                break
            }
            fmt.Println("Flag is false, waiting...")
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    atomicFlags()
}
```

### Example 8: Atomic Operations Best Practices

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

type AtomicCounter struct {
    value int64
}

func NewAtomicCounter() *AtomicCounter {
    return &AtomicCounter{}
}

func (ac *AtomicCounter) Increment() {
    atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Decrement() {
    atomic.AddInt64(&ac.value, -1)
}

func (ac *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&ac.value)
}

func (ac *AtomicCounter) Set(value int64) {
    atomic.StoreInt64(&ac.value, value)
}

func (ac *AtomicCounter) CompareAndSet(old, new int64) bool {
    return atomic.CompareAndSwapInt64(&ac.value, old, new)
}

func atomicBestPractices() {
    fmt.Println("=== Atomic Operations Best Practices ===")
    
    counter := NewAtomicCounter()
    
    // Multiple goroutines use counter
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 10; j++ {
                counter.Increment()
                fmt.Printf("Goroutine %d: counter = %d\n", id, counter.Value())
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Final counter value: %d\n", counter.Value())
}

func main() {
    atomicBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Atomic operations are indivisible**: Cannot be interrupted
2. **Lock-free**: No mutexes needed
3. **Fast**: Very efficient
4. **Limited types**: Only works with specific types
5. **Simple operations**: Only basic operations

## 🎯 When to Use Atomic Operations

1. **Simple counters**: Incrementing/decrementing counters
2. **Flags**: Boolean flags
3. **Simple state**: Basic state management
4. **Performance critical**: When you need maximum performance
5. **Lock-free programming**: Avoiding mutexes

## 🎯 When NOT to Use Atomic Operations

1. **Complex operations**: Operations that require multiple steps
2. **Complex data structures**: Slices, maps, structs
3. **Error handling**: Operations that can fail
4. **Complex logic**: Operations that require complex logic

## 🎯 Best Practices

1. **Use for simple operations**:
   ```go
   // GOOD - simple counter
   atomic.AddInt64(&counter, 1)
   
   // BAD - complex operation
   atomic.AddInt64(&counter, complexCalculation())
   ```

2. **Use appropriate types**:
   ```go
   // GOOD - use int64 for large values
   var counter int64
   atomic.AddInt64(&counter, 1)
   
   // BAD - use int32 for large values
   var counter int32
   atomic.AddInt32(&counter, 1) // Might overflow
   ```

3. **Handle overflow**:
   ```go
   // GOOD - check for overflow
   if atomic.LoadInt64(&counter) < maxValue {
       atomic.AddInt64(&counter, 1)
   }
   ```

## 🎯 Common Pitfalls

1. **Using atomic operations for complex operations**:
   ```go
   // BAD - atomic operations can't do complex operations
   atomic.AddInt64(&counter, calculateComplexValue())
   
   // GOOD - use atomic operations for simple operations
   atomic.AddInt64(&counter, 1)
   ```

2. **Not handling overflow**:
   ```go
   // BAD - might overflow
   atomic.AddInt32(&counter, 1)
   
   // GOOD - check for overflow
   if atomic.LoadInt32(&counter) < maxInt32 {
       atomic.AddInt32(&counter, 1)
   }
   ```

3. **Using atomic operations for complex data structures**:
   ```go
   // BAD - atomic operations don't work with slices
   atomic.StorePointer(&slice, unsafe.Pointer(&newSlice))
   
   // GOOD - use mutexes for complex data structures
   mutex.Lock()
   slice = newSlice
   mutex.Unlock()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a rate limiter:
- Use atomic operations to track the number of requests
- Implement a simple rate limiter that allows only 10 requests per second
- Show how atomic operations can be used for simple state management
- Compare performance with a mutex-based implementation

**Hint**: Use atomic operations to increment a counter and check if it exceeds the limit.

## 🚀 Next Steps

Now that you understand atomic operations, let's learn about **context package** in the next file: `17-context-package.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
