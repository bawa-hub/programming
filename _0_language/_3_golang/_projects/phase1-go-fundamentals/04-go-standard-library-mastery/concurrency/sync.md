# sync Package - Synchronization Primitives 🔄

The `sync` package provides synchronization primitives such as mutexes, condition variables, and wait groups. It's essential for safe concurrent programming in Go.

## 🎯 Key Concepts

### 1. **Mutexes**
- `Mutex` - Mutual exclusion lock
- `RWMutex` - Reader-writer mutex
- `Lock()` - Acquire lock
- `Unlock()` - Release lock
- `RLock()` - Acquire read lock
- `RUnlock()` - Release read lock

### 2. **Wait Groups**
- `WaitGroup` - Wait for goroutines to complete
- `Add()` - Add goroutines to wait group
- `Done()` - Mark goroutine as done
- `Wait()` - Wait for all goroutines to complete

### 3. **Once**
- `Once` - Execute function only once
- `Do()` - Execute function if not already executed
- Thread-safe initialization
- Lazy initialization

### 4. **Condition Variables**
- `Cond` - Condition variable
- `NewCond()` - Create condition variable
- `Wait()` - Wait for condition
- `Signal()` - Signal one waiter
- `Broadcast()` - Signal all waiters

### 5. **Atomic Operations**
- `atomic` - Atomic operations package
- `AddInt64()` - Atomic addition
- `LoadInt64()` - Atomic load
- `StoreInt64()` - Atomic store
- `CompareAndSwapInt64()` - Atomic compare and swap

### 6. **Map**
- `Map` - Concurrent map
- `Load()` - Load value
- `Store()` - Store value
- `Delete()` - Delete value
- `Range()` - Range over map

## 🚀 Common Patterns

### Basic Mutex Usage
```go
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

### WaitGroup Pattern
```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        // Do work
    }(i)
}

wg.Wait()
```

### Once Pattern
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

### RWMutex Pattern
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

## ⚠️ Common Pitfalls

1. **Deadlock** - Locking mutexes in wrong order
2. **Race conditions** - Accessing shared data without locks
3. **Forgetting to unlock** - Always use defer
4. **Double locking** - Don't lock already locked mutex
5. **WaitGroup misuse** - Don't call Done() more than Add()

## 🎯 Best Practices

1. **Use defer** - Always unlock mutexes with defer
2. **Minimize lock scope** - Hold locks for minimal time
3. **Avoid nested locks** - Prevent deadlocks
4. **Use RWMutex** - For read-heavy workloads
5. **Use atomic operations** - For simple operations

## 🔍 Advanced Features

### Custom Synchronization
```go
type SafeCounter struct {
    mu    sync.RWMutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.count
}
```

### Worker Pool Pattern
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
    close(results)
}
```

## 📚 Real-world Applications

1. **Web Servers** - Concurrent request handling
2. **Data Processing** - Parallel data processing
3. **Caching** - Thread-safe caches
4. **Database Connections** - Connection pooling
5. **File Processing** - Concurrent file operations

## 🧠 Memory Tips

- **sync** = **S**ynchronization **Y**nchronization **C**ontrol
- **Mutex** = **M**utual **E**xclusion
- **RWMutex** = **R**eader-**W**riter **M**utex
- **WaitGroup** = **W**ait for **G**roup
- **Once** = **O**nce execution
- **Cond** = **C**ondition variable
- **Atomic** = **A**tomic operations
- **Map** = **M**ap for concurrency

Remember: The sync package is your gateway to safe concurrent programming in Go! 🎯
