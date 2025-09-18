# 🚀 Synchronization Primitives: The Foundation of Safe Concurrency

## 📚 Table of Contents
1. [What Are Synchronization Primitives?](#what-are-synchronization-primitives)
2. [Mutex (Mutual Exclusion)](#mutex-mutual-exclusion)
3. [RWMutex (Read-Write Mutex)](#rwmutex-read-write-mutex)
4. [WaitGroup](#waitgroup)
5. [Once (One-Time Execution)](#once-one-time-execution)
6. [Cond (Condition Variables)](#cond-condition-variables)
7. [Atomic Operations](#atomic-operations)
8. [Map (Concurrent Map)](#map-concurrent-map)
9. [Pool (Object Pool)](#pool-object-pool)
10. [Performance Considerations](#performance-considerations)
11. [Common Patterns](#common-patterns)
12. [Best Practices](#best-practices)
13. [Common Pitfalls](#common-pitfalls)
14. [Exercises](#exercises)

---

## 🎯 What Are Synchronization Primitives?

**Synchronization primitives** are tools that help coordinate goroutines and protect shared resources from race conditions. They ensure that only one goroutine can access a shared resource at a time, or that goroutines can wait for specific conditions.

### Key Characteristics:
- **Thread-safe**: Can be used safely across multiple goroutines
- **Blocking**: Some primitives block until conditions are met
- **Efficient**: Optimized for high-performance concurrent access
- **Essential**: Required for safe concurrent programming

### Types of Synchronization Primitives:
- **Mutex**: Mutual exclusion for shared resources
- **RWMutex**: Read-write locks for multiple readers
- **WaitGroup**: Wait for goroutines to complete
- **Once**: Execute code exactly once
- **Cond**: Wait for specific conditions
- **Atomic**: Lock-free operations
- **Map**: Concurrent map operations
- **Pool**: Object pooling for efficiency

---

## 🔒 Mutex (Mutual Exclusion)

A **mutex** ensures that only one goroutine can access a shared resource at a time.

### Basic Usage:
```go
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

### Key Methods:
- `Lock()`: Acquire the mutex (blocks if already locked)
- `Unlock()`: Release the mutex
- `TryLock()`: Try to acquire without blocking (Go 1.18+)

### Characteristics:
- **Exclusive access**: Only one goroutine at a time
- **Recursive**: Same goroutine can lock multiple times
- **Fair**: First-come, first-served ordering
- **Memory barrier**: Ensures memory visibility

---

## 📖 RWMutex (Read-Write Mutex)

An **RWMutex** allows multiple readers or one writer, but not both simultaneously.

### Basic Usage:
```go
var rwmu sync.RWMutex
var data map[string]int

func read(key string) int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data[key]
}

func write(key string, value int) {
    rwmu.Lock()
    defer rwmu.Unlock()
    data[key] = value
}
```

### Key Methods:
- `RLock()`: Acquire read lock
- `RUnlock()`: Release read lock
- `Lock()`: Acquire write lock
- `Unlock()`: Release write lock
- `TryRLock()`: Try to acquire read lock (Go 1.18+)
- `TryLock()`: Try to acquire write lock (Go 1.18+)

### Characteristics:
- **Multiple readers**: Many goroutines can read simultaneously
- **Single writer**: Only one goroutine can write
- **Reader priority**: Readers have priority over writers
- **Writer starvation**: Writers may be starved by continuous readers

---

## ⏳ WaitGroup

A **WaitGroup** waits for a collection of goroutines to finish.

### Basic Usage:
```go
var wg sync.WaitGroup

func main() {
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(i)
    }
    wg.Wait() // Wait for all workers to complete
}

func worker(id int) {
    defer wg.Done()
    // Do work
}
```

### Key Methods:
- `Add(delta int)`: Add delta to the counter
- `Done()`: Decrement the counter by 1
- `Wait()`: Block until counter reaches zero

### Characteristics:
- **Counter-based**: Tracks number of goroutines
- **Blocking**: Wait() blocks until counter is zero
- **One-time use**: Cannot be reused after Wait() returns
- **No negative values**: Counter cannot go below zero

---

## 🎯 Once (One-Time Execution)

A **Once** ensures that a function is executed exactly once.

### Basic Usage:
```go
var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

### Key Methods:
- `Do(f func())`: Execute function exactly once

### Characteristics:
- **One-time execution**: Function runs only once
- **Thread-safe**: Safe to call from multiple goroutines
- **Blocking**: Other goroutines wait for first execution
- **No parameters**: Function cannot take parameters

---

## 🔔 Cond (Condition Variables)

A **Cond** provides a way for goroutines to wait for specific conditions.

### Basic Usage:
```go
var mu sync.Mutex
var cond = sync.NewCond(&mu)
var ready bool

func waitForReady() {
    mu.Lock()
    defer mu.Unlock()
    for !ready {
        cond.Wait()
    }
}

func signalReady() {
    mu.Lock()
    defer mu.Unlock()
    ready = true
    cond.Signal() // or cond.Broadcast()
}
```

### Key Methods:
- `Wait()`: Wait for condition (must hold lock)
- `Signal()`: Wake one waiting goroutine
- `Broadcast()`: Wake all waiting goroutines

### Characteristics:
- **Condition-based**: Wait for specific conditions
- **Lock required**: Must hold associated mutex
- **Spurious wakeups**: May wake up without condition being true
- **Broadcast vs Signal**: Choose based on number of waiters

---

## ⚛️ Atomic Operations

**Atomic operations** provide lock-free access to shared variables.

### Basic Usage:
```go
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}

func getValue() int64 {
    return atomic.LoadInt64(&counter)
}
```

### Key Functions:
- `AddInt64()`, `AddInt32()`: Atomic addition
- `LoadInt64()`, `LoadInt32()`: Atomic load
- `StoreInt64()`, `StoreInt32()`: Atomic store
- `SwapInt64()`, `SwapInt32()`: Atomic swap
- `CompareAndSwapInt64()`: Atomic compare-and-swap

### Characteristics:
- **Lock-free**: No mutex required
- **Hardware support**: Uses CPU atomic instructions
- **Limited types**: Only specific integer types
- **Memory ordering**: Provides memory barriers

---

## 🗺️ Map (Concurrent Map)

A **Map** provides a concurrent map implementation.

### Basic Usage:
```go
var m sync.Map

func main() {
    m.Store("key1", "value1")
    m.Store("key2", "value2")
    
    if value, ok := m.Load("key1"); ok {
        fmt.Println(value)
    }
    
    m.Delete("key1")
}
```

### Key Methods:
- `Store(key, value)`: Store key-value pair
- `Load(key)`: Load value for key
- `Delete(key)`: Delete key
- `LoadOrStore(key, value)`: Load or store
- `LoadAndDelete(key)`: Load and delete
- `Range(f func(key, value) bool)`: Range over map

### Characteristics:
- **Thread-safe**: Safe for concurrent access
- **Interface{}**: Keys and values are interface{}
- **No locking**: Uses lock-free techniques
- **Performance**: Optimized for concurrent access

---

## 🏊 Pool (Object Pool)

A **Pool** provides a pool of reusable objects.

### Basic Usage:
```go
var pool = sync.Pool{
    New: func() interface{} {
        return &Buffer{}
    },
}

func getBuffer() *Buffer {
    return pool.Get().(*Buffer)
}

func putBuffer(buf *Buffer) {
    buf.Reset()
    pool.Put(buf)
}
```

### Key Methods:
- `Get() interface{}`: Get object from pool
- `Put(x interface{})`: Put object back to pool

### Characteristics:
- **Object reuse**: Reduces allocation overhead
- **Automatic cleanup**: Objects are garbage collected
- **Thread-safe**: Safe for concurrent access
- **No size limit**: Pool can grow and shrink

---

## 📊 Performance Considerations

### 1. **Mutex vs RWMutex**
- **Mutex**: Better for write-heavy workloads
- **RWMutex**: Better for read-heavy workloads
- **Overhead**: RWMutex has higher overhead than Mutex

### 2. **Lock Contention**
- **High contention**: Many goroutines competing for same lock
- **Low contention**: Few goroutines using locks
- **Solution**: Reduce critical sections, use finer-grained locks

### 3. **Atomic vs Mutex**
- **Atomic**: Better for simple operations on single variables
- **Mutex**: Better for complex operations on multiple variables
- **Performance**: Atomic operations are generally faster

### 4. **Memory Ordering**
- **Acquire**: Read operations acquire memory ordering
- **Release**: Write operations release memory ordering
- **Sequential consistency**: Default memory ordering in Go

---

## 🎨 Common Patterns

### 1. **Singleton Pattern**
```go
var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

### 2. **Worker Pool Pattern**
```go
func workerPool(jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }
    
    wg.Wait()
}
```

### 3. **Rate Limiting Pattern**
```go
type RateLimiter struct {
    tokens chan struct{}
    rate   time.Duration
}

func (rl *RateLimiter) Allow() bool {
    select {
    case <-rl.tokens:
        return true
    default:
        return false
    }
}
```

### 4. **Cache Pattern**
```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, ok := c.data[key]
    return value, ok
}
```

---

## ✅ Best Practices

### 1. **Always Use defer with Mutex**
```go
// ✅ Good
mu.Lock()
defer mu.Unlock()
// Do work

// ❌ Bad
mu.Lock()
// Do work
mu.Unlock() // May not execute if panic occurs
```

### 2. **Minimize Critical Sections**
```go
// ✅ Good - minimal critical section
mu.Lock()
value := sharedData
mu.Unlock()
process(value) // Outside critical section

// ❌ Bad - large critical section
mu.Lock()
value := sharedData
process(value) // Inside critical section
mu.Unlock()
```

### 3. **Use RWMutex for Read-Heavy Workloads**
```go
// ✅ Good - many readers, few writers
var rwmu sync.RWMutex
var data map[string]int

func read(key string) int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data[key]
}
```

### 4. **Check WaitGroup Counter**
```go
// ✅ Good - check counter before waiting
if wg.counter > 0 {
    wg.Wait()
}
```

### 5. **Use Atomic for Simple Operations**
```go
// ✅ Good - simple counter
var counter int64
atomic.AddInt64(&counter, 1)

// ❌ Bad - complex operation
var mu sync.Mutex
var counter int64
mu.Lock()
counter++
mu.Unlock()
```

---

## ⚠️ Common Pitfalls

### 1. **Deadlock**
```go
// ❌ Wrong - potential deadlock
var mu1, mu2 sync.Mutex

func func1() {
    mu1.Lock()
    mu2.Lock()
    // Do work
    mu2.Unlock()
    mu1.Unlock()
}

func func2() {
    mu2.Lock()
    mu1.Lock()
    // Do work
    mu1.Unlock()
    mu2.Unlock()
}

// ✅ Correct - consistent lock ordering
func func1() {
    mu1.Lock()
    mu2.Lock()
    // Do work
    mu2.Unlock()
    mu1.Unlock()
}

func func2() {
    mu1.Lock() // Same order as func1
    mu2.Lock()
    // Do work
    mu2.Unlock()
    mu1.Unlock()
}
```

### 2. **Race Conditions**
```go
// ❌ Wrong - race condition
var counter int
go func() { counter++ }()
go func() { counter++ }()

// ✅ Correct - use mutex
var mu sync.Mutex
var counter int
go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

### 3. **WaitGroup Misuse**
```go
// ❌ Wrong - negative counter
var wg sync.WaitGroup
wg.Done() // Counter becomes negative

// ✅ Correct - call Add before Done
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
```

### 4. **Mutex Leaks**
```go
// ❌ Wrong - mutex not unlocked
mu.Lock()
// Do work
// Forgot to unlock

// ✅ Correct - use defer
mu.Lock()
defer mu.Unlock()
// Do work
```

### 5. **Double Unlock**
```go
// ❌ Wrong - double unlock
mu.Lock()
defer mu.Unlock()
mu.Unlock() // Panic!

// ✅ Correct - unlock only once
mu.Lock()
defer mu.Unlock()
// Do work
```

---

## 🧪 Exercises

### Exercise 1: Basic Mutex
Create a counter protected by a mutex.

### Exercise 2: RWMutex
Implement a thread-safe cache using RWMutex.

### Exercise 3: WaitGroup
Use WaitGroup to wait for multiple goroutines.

### Exercise 4: Once
Implement a singleton pattern using Once.

### Exercise 5: Cond
Use condition variables to coordinate goroutines.

### Exercise 6: Atomic Operations
Implement a counter using atomic operations.

### Exercise 7: Concurrent Map
Use sync.Map for thread-safe map operations.

### Exercise 8: Object Pool
Implement an object pool using sync.Pool.

---

## 🎯 Key Takeaways

1. **Use mutexes for shared resources** - protect critical sections
2. **Use RWMutex for read-heavy workloads** - allow multiple readers
3. **Use WaitGroup to coordinate goroutines** - wait for completion
4. **Use Once for one-time initialization** - thread-safe singletons
5. **Use Cond for complex coordination** - wait for conditions
6. **Use atomic operations for simple variables** - lock-free performance
7. **Use sync.Map for concurrent maps** - thread-safe map operations
8. **Use sync.Pool for object reuse** - reduce allocation overhead

---

## 🚀 Next Steps

Ready for the next topic? Let's move on to **Worker Pool Pattern** where you'll learn how to efficiently manage goroutines!

**Run the examples in this directory to see synchronization primitives in action!**
