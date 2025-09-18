# 🧠 Level 3, Topic 2: Memory Model & Race Conditions

## 🚀 Overview
Understanding Go's memory model and race conditions is crucial for writing correct concurrent programs. This topic will take you from basic memory model concepts to advanced race-free programming patterns that will make you a memory model master.

---

## 📚 Table of Contents

1. [Memory Model Fundamentals](#memory-model-fundamentals)
2. [Happens-Before Relationships](#happens-before-relationships)
3. [Race Conditions](#race-conditions)
4. [Atomic Operations](#atomic-operations)
5. [Memory Ordering](#memory-ordering)
6. [Visibility Guarantees](#visibility-guarantees)
7. [Race Detection](#race-detection)
8. [Race-Free Patterns](#race-free-patterns)
9. [Performance Implications](#performance-implications)
10. [Advanced Memory Patterns](#advanced-memory-patterns)
11. [Best Practices](#best-practices)
12. [Common Pitfalls](#common-pitfalls)
13. [Real-World Applications](#real-world-applications)

---

## 🧠 Memory Model Fundamentals

### What is a Memory Model?

A memory model defines the behavior of concurrent programs by specifying:
- **When memory operations are visible** to other goroutines
- **The ordering of memory operations** across goroutines
- **What constitutes a data race** and its consequences
- **Guarantees provided** by synchronization primitives

### Go's Memory Model

Go's memory model is based on the **happens-before** relationship:
- If event A happens before event B, then A's effects are visible to B
- If two events don't have a happens-before relationship, they can execute in any order
- **Data races** occur when two goroutines access the same memory location concurrently, at least one is a write, and there's no happens-before relationship

### Key Principles

1. **Sequential Consistency**: Within a single goroutine, operations appear to execute in program order
2. **Happens-Before**: Synchronization operations establish happens-before relationships
3. **No Data Races**: Programs with data races have undefined behavior
4. **Atomic Operations**: Provide synchronization without explicit locks

---

## ⏰ Happens-Before Relationships

### 1. Program Order

```go
// Within a single goroutine, operations happen in program order
func programOrder() {
    x := 1        // 1
    y := 2        // 2
    z := x + y    // 3 (sees x=1, y=2)
}
```

### 2. Channel Operations

```go
// Sending on a channel happens before receiving from that channel
func channelHappensBefore() {
    ch := make(chan int, 1)
    
    go func() {
        x := 42
        ch <- x    // Send happens before receive
    }()
    
    y := <-ch      // Receive happens after send
    // y is guaranteed to be 42
}
```

### 3. Mutex Operations

```go
// Unlocking a mutex happens before locking it again
func mutexHappensBefore() {
    var mu sync.Mutex
    var x int
    
    go func() {
        mu.Lock()
        x = 1
        mu.Unlock()    // Unlock happens before next Lock
    }()
    
    mu.Lock()
    y := x             // y is guaranteed to see x=1
    mu.Unlock()
}
```

### 4. Once Operations

```go
// Once.Do() calls happen in order
func onceHappensBefore() {
    var once sync.Once
    var x int
    
    go func() {
        once.Do(func() {
            x = 1
        })
    }()
    
    once.Do(func() {
        // This will see x=1 if the other goroutine ran first
        fmt.Println(x)
    })
}
```

### 5. Package Init

```go
// Package initialization happens before main()
var globalVar = initGlobalVar()

func initGlobalVar() int {
    return 42
}

func main() {
    // globalVar is guaranteed to be 42
    fmt.Println(globalVar)
}
```

---

## 🏃 Race Conditions

### What is a Data Race?

A data race occurs when:
1. Two goroutines access the same memory location
2. At least one access is a write
3. There's no happens-before relationship between the accesses

### Race Condition Examples

#### 1. Simple Race Condition

```go
// ❌ RACE CONDITION
func raceCondition() {
    var x int
    
    go func() {
        x = 1    // Write
    }()
    
    go func() {
        fmt.Println(x)    // Read - RACE!
    }()
}
```

#### 2. Counter Race

```go
// ❌ RACE CONDITION
func counterRace() {
    var counter int
    
    for i := 0; i < 1000; i++ {
        go func() {
            counter++    // Read-modify-write - RACE!
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Println(counter)    // Unpredictable result
}
```

#### 3. Map Race

```go
// ❌ RACE CONDITION
func mapRace() {
    m := make(map[string]int)
    
    go func() {
        m["key"] = 1    // Write
    }()
    
    go func() {
        fmt.Println(m["key"])    // Read - RACE!
    }()
}
```

### Race Condition Consequences

1. **Undefined Behavior**: Programs with data races have undefined behavior
2. **Unpredictable Results**: Race conditions can cause unpredictable program behavior
3. **Memory Corruption**: Can lead to memory corruption and crashes
4. **Security Vulnerabilities**: Race conditions can be exploited for security attacks

---

## ⚛️ Atomic Operations

### What are Atomic Operations?

Atomic operations are operations that are performed as a single, indivisible unit:
- **Cannot be interrupted** by other goroutines
- **Provide synchronization** without explicit locks
- **Are lock-free** and generally faster than mutexes
- **Guarantee memory ordering** for the operation

### Atomic Types

```go
import "sync/atomic"

// Atomic integer types
var counter int64
var flag int32
var pointer unsafe.Pointer

// Atomic operations
atomic.AddInt64(&counter, 1)           // Add
atomic.LoadInt64(&counter)             // Load
atomic.StoreInt64(&counter, 42)        // Store
atomic.SwapInt64(&counter, 100)        // Swap
atomic.CompareAndSwapInt64(&counter, 42, 100)  // Compare and swap
```

### Atomic Counter Example

```go
// ✅ RACE-FREE with atomic operations
func atomicCounter() {
    var counter int64
    
    for i := 0; i < 1000; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)    // Atomic increment
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Println(atomic.LoadInt64(&counter))    // Always 1000
}
```

### Atomic Flag Example

```go
// ✅ RACE-FREE with atomic flag
func atomicFlag() {
    var flag int32
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        atomic.StoreInt32(&flag, 1)    // Set flag
    }()
    
    for atomic.LoadInt32(&flag) == 0 {
        // Wait for flag
        time.Sleep(10 * time.Millisecond)
    }
    
    fmt.Println("Flag is set!")
}
```

### Compare and Swap Example

```go
// ✅ RACE-FREE with compare and swap
func compareAndSwap() {
    var value int64 = 42
    
    go func() {
        // Try to change 42 to 100
        if atomic.CompareAndSwapInt64(&value, 42, 100) {
            fmt.Println("Successfully changed value to 100")
        } else {
            fmt.Println("Failed to change value")
        }
    }()
    
    go func() {
        // Try to change 42 to 200
        if atomic.CompareAndSwapInt64(&value, 42, 200) {
            fmt.Println("Successfully changed value to 200")
        } else {
            fmt.Println("Failed to change value")
        }
    }()
    
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Final value:", atomic.LoadInt64(&value))
}
```

---

## 🔄 Memory Ordering

### Sequential Consistency

Go provides **sequential consistency** for:
- **Atomic operations** on the same variable
- **Channel operations** on the same channel
- **Mutex operations** on the same mutex

### Memory Ordering Guarantees

#### 1. Atomic Operations

```go
// Atomic operations on the same variable are sequentially consistent
func atomicOrdering() {
    var x, y int64
    
    go func() {
        atomic.StoreInt64(&x, 1)
        atomic.StoreInt64(&y, 2)
    }()
    
    go func() {
        for atomic.LoadInt64(&y) != 2 {
            // Wait for y to be set
        }
        // x is guaranteed to be 1 here
        fmt.Println("x =", atomic.LoadInt64(&x))
    }()
}
```

#### 2. Channel Operations

```go
// Channel operations are sequentially consistent
func channelOrdering() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        ch1 <- 1
        ch2 <- 2
    }()
    
    go func() {
        <-ch2    // Receive from ch2
        // ch1 <- 1 is guaranteed to have happened
        fmt.Println("Received from ch1:", <-ch1)
    }()
}
```

#### 3. Mutex Operations

```go
// Mutex operations are sequentially consistent
func mutexOrdering() {
    var mu sync.Mutex
    var x, y int
    
    go func() {
        mu.Lock()
        x = 1
        y = 2
        mu.Unlock()
    }()
    
    go func() {
        mu.Lock()
        if y == 2 {
            // x is guaranteed to be 1 here
            fmt.Println("x =", x)
        }
        mu.Unlock()
    }()
}
```

---

## 👁️ Visibility Guarantees

### What are Visibility Guarantees?

Visibility guarantees ensure that:
- **Writes are visible** to reads in other goroutines
- **Synchronization operations** establish visibility
- **Memory barriers** prevent reordering

### Visibility Examples

#### 1. Without Synchronization

```go
// ❌ NO VISIBILITY GUARANTEE
func noVisibility() {
    var x int
    
    go func() {
        x = 42    // Write
    }()
    
    go func() {
        // x might be 0 or 42 - no guarantee
        fmt.Println(x)
    }()
}
```

#### 2. With Channel Synchronization

```go
// ✅ VISIBILITY GUARANTEE with channels
func channelVisibility() {
    var x int
    ch := make(chan int)
    
    go func() {
        x = 42
        ch <- 1    // Synchronization point
    }()
    
    go func() {
        <-ch       // Synchronization point
        // x is guaranteed to be 42
        fmt.Println(x)
    }()
}
```

#### 3. With Mutex Synchronization

```go
// ✅ VISIBILITY GUARANTEE with mutex
func mutexVisibility() {
    var x int
    var mu sync.Mutex
    
    go func() {
        mu.Lock()
        x = 42
        mu.Unlock()    // Synchronization point
    }()
    
    go func() {
        mu.Lock()      // Synchronization point
        // x is guaranteed to be 42
        fmt.Println(x)
        mu.Unlock()
    }()
}
```

#### 4. With Atomic Operations

```go
// ✅ VISIBILITY GUARANTEE with atomic
func atomicVisibility() {
    var x int64
    
    go func() {
        atomic.StoreInt64(&x, 42)    // Synchronization point
    }()
    
    go func() {
        // x is guaranteed to be 42
        fmt.Println(atomic.LoadInt64(&x))
    }()
}
```

---

## 🔍 Race Detection

### Go Race Detector

Go provides a built-in race detector:
```bash
go run -race program.go
go build -race program.go
go test -race ./...
```

### How Race Detection Works

1. **Instrumentation**: Compiler instruments memory accesses
2. **Runtime Monitoring**: Runtime monitors for concurrent accesses
3. **Race Reporting**: Reports races when detected
4. **Performance Impact**: 2-10x slower, 5-10x more memory

### Race Detection Examples

#### 1. Detecting Simple Races

```go
// Run with: go run -race race_example.go
func raceDetectionExample() {
    var counter int
    
    for i := 0; i < 1000; i++ {
        go func() {
            counter++    // Race detected here
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Println(counter)
}
```

#### 2. Detecting Map Races

```go
// Run with: go run -race map_race.go
func mapRaceDetection() {
    m := make(map[string]int)
    
    go func() {
        for i := 0; i < 1000; i++ {
            m[fmt.Sprintf("key%d", i)] = i    // Race detected here
        }
    }()
    
    go func() {
        for i := 0; i < 1000; i++ {
            fmt.Println(m[fmt.Sprintf("key%d", i)])    // Race detected here
        }
    }()
    
    time.Sleep(1 * time.Second)
}
```

### Race Detection Best Practices

1. **Run race detector** in CI/CD pipeline
2. **Test with race detector** regularly
3. **Fix races immediately** when detected
4. **Don't ignore race warnings**
5. **Use race detector** for debugging

---

## 🛡️ Race-Free Patterns

### 1. Mutex Pattern

```go
// ✅ RACE-FREE with mutex
type SafeCounter struct {
    mu      sync.Mutex
    counter int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    c.counter++
    c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.counter
}
```

### 2. Channel Pattern

```go
// ✅ RACE-FREE with channels
type ChannelCounter struct {
    counter int
    ch      chan int
}

func NewChannelCounter() *ChannelCounter {
    c := &ChannelCounter{
        ch: make(chan int),
    }
    
    go c.run()
    return c
}

func (c *ChannelCounter) run() {
    for {
        select {
        case <-c.ch:
            c.counter++
        case c.ch <- c.counter:
            // Return current value
        }
    }
}

func (c *ChannelCounter) Increment() {
    c.ch <- 1
}

func (c *ChannelCounter) Value() int {
    return <-c.ch
}
```

### 3. Atomic Pattern

```go
// ✅ RACE-FREE with atomic operations
type AtomicCounter struct {
    counter int64
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.counter, 1)
}

func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.counter)
}
```

### 4. Single Writer Pattern

```go
// ✅ RACE-FREE with single writer
type SingleWriterCounter struct {
    counter int
    ch      chan int
}

func NewSingleWriterCounter() *SingleWriterCounter {
    c := &SingleWriterCounter{
        ch: make(chan int, 1000),
    }
    
    go c.writer()
    return c
}

func (c *SingleWriterCounter) writer() {
    for {
        select {
        case <-c.ch:
            c.counter++
        }
    }
}

func (c *SingleWriterCounter) Increment() {
    c.ch <- 1
}

func (c *SingleWriterCounter) Value() int {
    return c.counter
}
```

---

## ⚡ Performance Implications

### Mutex vs Atomic vs Channel Performance

#### 1. Mutex Performance

```go
// Mutex: ~100ns per operation
func mutexPerformance() {
    var mu sync.Mutex
    var counter int
    
    start := time.Now()
    for i := 0; i < 1000000; i++ {
        mu.Lock()
        counter++
        mu.Unlock()
    }
    duration := time.Since(start)
    fmt.Printf("Mutex: %v\n", duration)
}
```

#### 2. Atomic Performance

```go
// Atomic: ~10ns per operation
func atomicPerformance() {
    var counter int64
    
    start := time.Now()
    for i := 0; i < 1000000; i++ {
        atomic.AddInt64(&counter, 1)
    }
    duration := time.Since(start)
    fmt.Printf("Atomic: %v\n", duration)
}
```

#### 3. Channel Performance

```go
// Channel: ~1000ns per operation
func channelPerformance() {
    ch := make(chan int, 1000)
    
    go func() {
        for i := 0; i < 1000000; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    start := time.Now()
    for range ch {
        // Process
    }
    duration := time.Since(start)
    fmt.Printf("Channel: %v\n", duration)
}
```

### Performance Guidelines

1. **Use atomic operations** for simple counters/flags
2. **Use mutexes** for complex data structures
3. **Use channels** for communication patterns
4. **Avoid unnecessary synchronization**
5. **Profile before optimizing**

---

## 🔬 Advanced Memory Patterns

### 1. Lock-Free Data Structures

```go
// Lock-free stack using atomic operations
type LockFreeStack struct {
    head unsafe.Pointer
}

type node struct {
    value int
    next  unsafe.Pointer
}

func (s *LockFreeStack) Push(value int) {
    n := &node{value: value}
    
    for {
        head := atomic.LoadPointer(&s.head)
        n.next = head
        
        if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(n)) {
            break
        }
    }
}

func (s *LockFreeStack) Pop() (int, bool) {
    for {
        head := atomic.LoadPointer(&s.head)
        if head == nil {
            return 0, false
        }
        
        n := (*node)(head)
        next := atomic.LoadPointer(&n.next)
        
        if atomic.CompareAndSwapPointer(&s.head, head, next) {
            return n.value, true
        }
    }
}
```

### 2. Memory Pool Pattern

```go
// Memory pool to avoid allocations
type MemoryPool struct {
    pool sync.Pool
}

func NewMemoryPool() *MemoryPool {
    return &MemoryPool{
        pool: sync.Pool{
            New: func() interface{} {
                return make([]byte, 1024)
            },
        },
    }
}

func (p *MemoryPool) Get() []byte {
    return p.pool.Get().([]byte)
}

func (p *MemoryPool) Put(buf []byte) {
    p.pool.Put(buf)
}
```

### 3. Double-Checked Locking

```go
// Double-checked locking pattern
type Singleton struct {
    data string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    if instance == nil {
        once.Do(func() {
            instance = &Singleton{data: "initialized"}
        })
    }
    return instance
}
```

---

## 🎯 Best Practices

### 1. Avoid Data Races

```go
// ❌ Bad: Data race
var globalCounter int

func increment() {
    globalCounter++    // Race condition
}

// ✅ Good: Use synchronization
var globalCounter int
var mu sync.Mutex

func increment() {
    mu.Lock()
    globalCounter++
    mu.Unlock()
}
```

### 2. Use Atomic Operations for Simple Cases

```go
// ❌ Bad: Mutex for simple counter
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Add() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}

// ✅ Good: Atomic for simple counter
type Counter struct {
    n int64
}

func (c *Counter) Add() {
    atomic.AddInt64(&c.n, 1)
}
```

### 3. Minimize Critical Sections

```go
// ❌ Bad: Large critical section
func badExample() {
    mu.Lock()
    // Do a lot of work
    result := expensiveComputation()
    data[result] = result
    mu.Unlock()
}

// ✅ Good: Small critical section
func goodExample() {
    result := expensiveComputation()    // Outside critical section
    
    mu.Lock()
    data[result] = result
    mu.Unlock()
}
```

### 4. Use Channels for Communication

```go
// ❌ Bad: Shared memory with mutex
type BadWorker struct {
    mu   sync.Mutex
    data map[string]int
}

// ✅ Good: Channels for communication
type GoodWorker struct {
    requests  chan string
    responses chan int
}

func (w *GoodWorker) Process() {
    for req := range w.requests {
        // Process request
        w.responses <- len(req)
    }
}
```

---

## ⚠️ Common Pitfalls

### 1. False Sharing

```go
// ❌ Bad: False sharing
type BadCounter struct {
    counter1 int64
    counter2 int64    // Same cache line as counter1
}

// ✅ Good: Avoid false sharing
type GoodCounter struct {
    counter1 int64
    _        [7]int64    // Padding to avoid false sharing
    counter2 int64
}
```

### 2. Race in Map Operations

```go
// ❌ Bad: Race in map operations
func badMapExample() {
    m := make(map[string]int)
    
    go func() {
        m["key"] = 1    // Race condition
    }()
    
    go func() {
        delete(m, "key")    // Race condition
    }()
}

// ✅ Good: Use sync.Map or mutex
func goodMapExample() {
    var m sync.Map
    
    go func() {
        m.Store("key", 1)    // Thread-safe
    }()
    
    go func() {
        m.Delete("key")    // Thread-safe
    }()
}
```

### 3. Race in Slice Operations

```go
// ❌ Bad: Race in slice operations
func badSliceExample() {
    s := make([]int, 0)
    
    go func() {
        s = append(s, 1)    // Race condition
    }()
    
    go func() {
        fmt.Println(len(s))    // Race condition
    }()
}

// ✅ Good: Use channels or mutex
func goodSliceExample() {
    ch := make(chan int, 100)
    
    go func() {
        ch <- 1    // Thread-safe
    }()
    
    go func() {
        fmt.Println(len(ch))    // Thread-safe
    }()
}
```

---

## 🌍 Real-World Applications

### 1. Web Server with Atomic Counters

```go
type WebServer struct {
    requestCount int64
    errorCount   int64
    startTime    time.Time
}

func (s *WebServer) HandleRequest() {
    atomic.AddInt64(&s.requestCount, 1)
    
    // Process request
    if err := s.processRequest(); err != nil {
        atomic.AddInt64(&s.errorCount, 1)
    }
}

func (s *WebServer) GetStats() (int64, int64) {
    return atomic.LoadInt64(&s.requestCount), atomic.LoadInt64(&s.errorCount)
}
```

### 2. Cache with Atomic Operations

```go
type AtomicCache struct {
    data map[string]interface{}
    mu   sync.RWMutex
    hits int64
    misses int64
}

func (c *AtomicCache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    value, exists := c.data[key]
    c.mu.RUnlock()
    
    if exists {
        atomic.AddInt64(&c.hits, 1)
    } else {
        atomic.AddInt64(&c.misses, 1)
    }
    
    return value, exists
}
```

### 3. Rate Limiter with Atomic Operations

```go
type RateLimiter struct {
    limit     int64
    interval  time.Duration
    tokens    int64
    lastReset time.Time
    mu        sync.Mutex
}

func (r *RateLimiter) Allow() bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    now := time.Now()
    if now.Sub(r.lastReset) >= r.interval {
        atomic.StoreInt64(&r.tokens, r.limit)
        r.lastReset = now
    }
    
    if atomic.LoadInt64(&r.tokens) > 0 {
        atomic.AddInt64(&r.tokens, -1)
        return true
    }
    
    return false
}
```

---

## 🎓 Summary

Understanding Go's memory model and race conditions is essential for writing correct concurrent programs. Key takeaways:

1. **Understand happens-before relationships** for correct synchronization
2. **Use atomic operations** for simple counters and flags
3. **Use mutexes** for complex data structures
4. **Use channels** for communication patterns
5. **Always run race detector** in development
6. **Avoid data races** at all costs
7. **Profile performance** before optimizing
8. **Follow Go's memory model** guidelines

Mastering memory model concepts will make you a more effective Go developer and help you build better concurrent applications! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Run race detector** on your code
3. **Experiment** with different synchronization patterns
4. **Move to the next topic** in the curriculum
5. **Apply** memory model knowledge to real-world projects

Ready to become a Memory Model master? Let's dive into the implementation! 💪

