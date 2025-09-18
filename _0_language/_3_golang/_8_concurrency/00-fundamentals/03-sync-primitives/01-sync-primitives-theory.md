# 🔒 Go Sync Primitives - Complete Theory Guide
## Everything in the sync package explained

### 📚 Table of Contents

1. [Introduction to Sync Primitives](#introduction)
2. [Mutex - Mutual Exclusion](#mutex)
3. [RWMutex - Read-Write Mutex](#rwmutex)
4. [WaitGroup - Wait for Goroutines](#waitgroup)
5. [Once - One-Time Execution](#once)
6. [Cond - Condition Variables](#cond)
7. [Pool - Object Pooling](#pool)
8. [Map - Concurrent Map](#map)
9. [Atomic Operations](#atomic)
10. [Best Practices](#best-practices)
11. [Common Patterns](#common-patterns)
12. [Performance Considerations](#performance)
13. [Troubleshooting](#troubleshooting)

---

## Introduction to Sync Primitives {#introduction}

### What are Sync Primitives?
Sync primitives are low-level synchronization mechanisms that help coordinate access to shared resources in concurrent programs. They prevent race conditions and ensure data consistency.

### Why Do We Need Them?
- **Race Conditions**: Multiple goroutines accessing shared data simultaneously
- **Data Consistency**: Ensuring data remains valid during concurrent access
- **Coordination**: Synchronizing goroutines to work together
- **Resource Protection**: Preventing unauthorized access to critical sections

### Go's Philosophy
> "Don't communicate by sharing memory; share memory by communicating" - Go's concurrency philosophy

However, when you must share memory, Go provides sync primitives to do it safely.

---

## Mutex - Mutual Exclusion {#mutex}

### What is a Mutex?
A Mutex (mutual exclusion) is a lock that ensures only one goroutine can access a shared resource at a time.

### Key Concepts
- **Critical Section**: Code that accesses shared resources
- **Lock/Unlock**: Acquiring and releasing the mutex
- **Blocking**: Other goroutines wait until mutex is released

### Basic Usage
```go
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()         // Acquire lock
    defer mu.Unlock() // Release lock when function exits
    counter++         // Critical section
}
```

### Mutex States
1. **Unlocked**: No goroutine holds the lock
2. **Locked**: One goroutine holds the lock
3. **Blocked**: Other goroutines waiting for the lock

### Important Rules
- Always unlock in the same goroutine that locked
- Use `defer mu.Unlock()` to prevent forgetting to unlock
- Don't try to lock an already locked mutex in the same goroutine (deadlock)

### Performance Characteristics
- **Fast**: Very efficient for short critical sections
- **Blocking**: Other goroutines wait (don't consume CPU)
- **Fair**: FIFO ordering for waiting goroutines

---

## RWMutex - Read-Write Mutex {#rwmutex}

### What is an RWMutex?
A read-write mutex allows multiple readers OR one writer, but not both simultaneously.

### Key Concepts
- **Read Lock**: Multiple goroutines can hold read locks
- **Write Lock**: Only one goroutine can hold write lock
- **Priority**: Writers have priority over readers

### Basic Usage
```go
var rwmu sync.RWMutex
var data map[string]int

func read(key string) int {
    rwmu.RLock()         // Acquire read lock
    defer rwmu.RUnlock() // Release read lock
    return data[key]     // Safe to read
}

func write(key string, value int) {
    rwmu.Lock()         // Acquire write lock
    defer rwmu.Unlock() // Release write lock
    data[key] = value   // Safe to write
}
```

### When to Use RWMutex
- **Read-heavy workloads**: Many readers, few writers
- **Large data structures**: Expensive to copy
- **Caching scenarios**: Frequently read, occasionally updated

### Performance Characteristics
- **Readers**: Very fast, can run concurrently
- **Writers**: Slower, blocks all readers
- **Memory**: More memory overhead than regular Mutex

---

## WaitGroup - Wait for Goroutines {#waitgroup}

### What is a WaitGroup?
A WaitGroup waits for a collection of goroutines to finish executing.

### Key Concepts
- **Add(n)**: Increment counter by n
- **Done()**: Decrement counter by 1
- **Wait()**: Block until counter reaches 0

### Basic Usage
```go
var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done() // Decrement counter when done
    // Do work...
}

func main() {
    for i := 0; i < 3; i++ {
        wg.Add(1)    // Increment counter
        go worker(i)
    }
    wg.Wait() // Wait for all workers to finish
}
```

### Important Rules
- **Add before goroutine**: Call Add() before starting goroutine
- **Done in goroutine**: Call Done() inside the goroutine
- **Don't reuse**: Don't reuse WaitGroup after Wait() returns

### Common Patterns
- **Worker pools**: Wait for all workers to complete
- **Batch processing**: Wait for all items to be processed
- **Cleanup**: Wait for cleanup goroutines to finish

---

## Once - One-Time Execution {#once}

### What is Once?
Once ensures a function is executed only once, even if called from multiple goroutines.

### Key Concepts
- **Do(f)**: Execute function f only once
- **Thread-safe**: Safe to call from multiple goroutines
- **Idempotent**: Multiple calls to Do() are safe

### Basic Usage
```go
var once sync.Once
var initialized bool

func init() {
    once.Do(func() {
        // This will only execute once
        initialized = true
        fmt.Println("Initialized")
    })
}
```

### When to Use Once
- **Singleton initialization**: Initialize global variables once
- **Lazy initialization**: Initialize on first use
- **Resource setup**: Set up resources only once

### Performance Characteristics
- **Fast**: Very efficient after first call
- **Atomic**: Uses atomic operations internally
- **Memory**: Minimal memory overhead

---

## Cond - Condition Variables {#cond}

### What is Cond?
A condition variable allows goroutines to wait for a condition to become true.

### Key Concepts
- **Wait()**: Wait for condition to be signaled
- **Signal()**: Wake up one waiting goroutine
- **Broadcast()**: Wake up all waiting goroutines
- **Mutex**: Must be associated with a mutex

### Basic Usage
```go
var mu sync.Mutex
cond := sync.NewCond(&mu)
ready := false

// Waiter
go func() {
    mu.Lock()
    for !ready {
        cond.Wait() // Wait for condition
    }
    // Condition is true, proceed
    mu.Unlock()
}()

// Signaler
go func() {
    mu.Lock()
    ready = true
    cond.Signal() // Wake up waiter
    mu.Unlock()
}()
```

### When to Use Cond
- **Producer-Consumer**: Wait for data to be available
- **Resource waiting**: Wait for resources to become available
- **State changes**: Wait for state to change

### Important Rules
- **Lock before Wait()**: Must hold mutex before calling Wait()
- **Loop around Wait()**: Always use a loop to check condition
- **Signal after change**: Signal after changing the condition

---

## Pool - Object Pooling {#pool}

### What is a Pool?
A Pool provides a way to reuse objects, reducing garbage collection pressure.

### Key Concepts
- **Get()**: Get object from pool
- **Put()**: Return object to pool
- **New**: Function to create new objects when pool is empty

### Basic Usage
```go
var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024) // Create new buffer
    },
}

func useBuffer() {
    buf := pool.Get().([]byte) // Get from pool
    defer pool.Put(buf)        // Return to pool
    
    // Use buffer...
}
```

### When to Use Pool
- **Frequent allocation**: Objects created/destroyed frequently
- **Expensive objects**: Objects that are expensive to create
- **Memory pressure**: Reduce garbage collection

### Performance Characteristics
- **Memory efficient**: Reduces allocations
- **GC friendly**: Reduces garbage collection pressure
- **Thread-safe**: Safe for concurrent access

---

## Map - Concurrent Map {#map}

### What is a Map?
A Map is a concurrent map implementation that's safe for concurrent use.

### Key Concepts
- **Store(key, value)**: Store key-value pair
- **Load(key)**: Load value for key
- **Delete(key)**: Delete key-value pair
- **Range(f)**: Iterate over all key-value pairs

### Basic Usage
```go
var m sync.Map

// Store values
m.Store("key1", "value1")
m.Store("key2", "value2")

// Load values
if value, ok := m.Load("key1"); ok {
    fmt.Println(value)
}

// Delete values
m.Delete("key2")

// Range over all pairs
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%v = %v\n", key, value)
    return true // Continue iteration
})
```

### When to Use Map
- **Concurrent access**: Multiple goroutines accessing map
- **Read-heavy**: More reads than writes
- **Large maps**: Maps with many key-value pairs

### Performance Characteristics
- **Read-optimized**: Very fast reads
- **Write-optimized**: Fast writes
- **Memory**: More memory overhead than regular map

---

## Atomic Operations {#atomic}

### What are Atomic Operations?
Atomic operations are lock-free operations that are guaranteed to complete without interruption.

### Key Concepts
- **Load**: Atomically load a value
- **Store**: Atomically store a value
- **Add**: Atomically add to a value
- **CompareAndSwap**: Atomically compare and swap

### Basic Usage
```go
var counter int64

// Atomic operations
atomic.AddInt64(&counter, 1)           // Increment
value := atomic.LoadInt64(&counter)    // Load
atomic.StoreInt64(&counter, 100)       // Store

// Compare and swap
old := atomic.LoadInt64(&counter)
new := old + 1
if atomic.CompareAndSwapInt64(&counter, old, new) {
    // Successfully updated
}
```

### When to Use Atomic
- **Counters**: Simple counters and flags
- **Lock-free algorithms**: When you need lock-free code
- **Performance critical**: When mutex overhead is too high

### Performance Characteristics
- **Very fast**: Faster than mutex for simple operations
- **Lock-free**: No blocking or waiting
- **Limited**: Only works with specific types

---

## Best Practices {#best-practices}

### 1. Always Use defer for Unlock
```go
// Good
mu.Lock()
defer mu.Unlock()
// ... critical section

// Bad
mu.Lock()
// ... critical section
mu.Unlock() // Might be forgotten
```

### 2. Keep Critical Sections Short
```go
// Good
mu.Lock()
data[key] = value
mu.Unlock()

// Bad
mu.Lock()
// ... lots of work ...
mu.Unlock()
```

### 3. Don't Hold Locks Across Function Calls
```go
// Good
mu.Lock()
value := data[key]
mu.Unlock()
process(value)

// Bad
mu.Lock()
process(data[key]) // process might take long
mu.Unlock()
```

### 4. Use RWMutex for Read-Heavy Workloads
```go
// Good for read-heavy
var rwmu sync.RWMutex
var data map[string]int

func read(key string) int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data[key]
}
```

### 5. Use Once for One-Time Initialization
```go
var once sync.Once
var instance *MyType

func GetInstance() *MyType {
    once.Do(func() {
        instance = &MyType{}
    })
    return instance
}
```

---

## Common Patterns {#common-patterns}

### 1. Producer-Consumer with Mutex
```go
var mu sync.Mutex
var queue []int
var wg sync.WaitGroup

func producer() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        mu.Lock()
        queue = append(queue, i)
        mu.Unlock()
    }
}

func consumer() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        mu.Lock()
        if len(queue) > 0 {
            item := queue[0]
            queue = queue[1:]
            // Process item
        }
        mu.Unlock()
    }
}
```

### 2. Worker Pool with WaitGroup
```go
func workerPool(jobs <-chan int, results chan<- int) {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- job * 2
            }
        }()
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 3. Rate Limiting with Mutex
```go
type RateLimiter struct {
    mu       sync.Mutex
    lastTime time.Time
    interval time.Duration
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    if now.Sub(rl.lastTime) >= rl.interval {
        rl.lastTime = now
        return true
    }
    return false
}
```

---

## Performance Considerations {#performance}

### 1. Mutex vs RWMutex
- **Mutex**: Use for balanced read/write workloads
- **RWMutex**: Use for read-heavy workloads (80%+ reads)

### 2. Atomic vs Mutex
- **Atomic**: Use for simple operations (counters, flags)
- **Mutex**: Use for complex operations or multiple variables

### 3. Pool vs New
- **Pool**: Use for frequently allocated objects
- **New**: Use for rarely allocated objects

### 4. Map vs Regular Map + Mutex
- **Map**: Use for concurrent access
- **Regular Map + Mutex**: Use for simple cases

---

## Troubleshooting {#troubleshooting}

### 1. Deadlocks
**Problem**: Program hangs, all goroutines are blocked
**Solution**: 
- Check for lock ordering issues
- Use `go run -race` to detect race conditions
- Avoid holding locks across function calls

### 2. Race Conditions
**Problem**: Data corruption, unpredictable behavior
**Solution**:
- Use mutex to protect shared data
- Use atomic operations for simple cases
- Run with `-race` flag

### 3. Performance Issues
**Problem**: Slow concurrent performance
**Solution**:
- Profile with `go tool pprof`
- Reduce critical section size
- Consider using RWMutex for read-heavy workloads

### 4. Memory Leaks
**Problem**: Memory usage keeps growing
**Solution**:
- Use Pool for frequently allocated objects
- Ensure proper cleanup in defer statements
- Use context for cancellation

---

## Summary

Sync primitives are essential for writing safe, concurrent Go programs. Choose the right primitive for your use case:

- **Mutex**: General-purpose locking
- **RWMutex**: Read-heavy workloads
- **WaitGroup**: Wait for goroutines
- **Once**: One-time initialization
- **Cond**: Wait for conditions
- **Pool**: Object reuse
- **Map**: Concurrent map
- **Atomic**: Lock-free operations

Remember: "Don't communicate by sharing memory; share memory by communicating" - but when you must share memory, use sync primitives to do it safely!
