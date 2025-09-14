# 13 - Read-Write Mutex (RWMutex)

## 🎯 Learning Objectives
- Understand what RWMutex is and when to use it
- Learn the difference between RLock and Lock
- Master RWMutex performance characteristics
- Practice with read-heavy vs write-heavy scenarios
- Understand RWMutex best practices

## 📚 Theory

### What is RWMutex?

**RWMutex** (Read-Write Mutex) allows multiple readers or one writer at a time, but not both simultaneously.

**Key characteristics:**
- **Multiple readers**: Many goroutines can read simultaneously
- **Single writer**: Only one goroutine can write at a time
- **Reader-writer exclusion**: Readers and writers are mutually exclusive
- **Better performance**: For read-heavy workloads

### RWMutex vs Mutex

| Feature | Mutex | RWMutex |
|---------|-------|---------|
| Readers | 1 at a time | Multiple |
| Writers | 1 at a time | 1 at a time |
| Performance | Lower | Higher (read-heavy) |
| Complexity | Simple | More complex |

### RWMutex Operations

1. **RLock**: `rwmutex.RLock()` - Acquire read lock
2. **RUnlock**: `rwmutex.RUnlock()` - Release read lock
3. **Lock**: `rwmutex.Lock()` - Acquire write lock
4. **Unlock**: `rwmutex.Unlock()` - Release write lock

## 💻 Code Examples

### Example 1: Basic RWMutex Usage

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicRWMutex() {
    fmt.Println("=== Basic RWMutex Usage ===")
    
    var counter int
    var rwmutex sync.RWMutex
    
    // Multiple readers
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 3; j++ {
                rwmutex.RLock() // Read lock
                value := counter
                rwmutex.RUnlock() // Read unlock
                fmt.Printf("Reader %d: counter = %d\n", id, value)
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // One writer
    go func() {
        for i := 0; i < 3; i++ {
            rwmutex.Lock() // Write lock
            counter++
            fmt.Printf("Writer: counter = %d\n", counter)
            rwmutex.Unlock() // Write unlock
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    basicRWMutex()
}
```

**Run this code:**
```bash
go run 13-rwmutex.go
```

### Example 2: RWMutex with Struct

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeData struct {
    data   map[string]int
    rwmutex sync.RWMutex
}

func NewSafeData() *SafeData {
    return &SafeData{
        data: make(map[string]int),
    }
}

func (sd *SafeData) Get(key string) (int, bool) {
    sd.rwmutex.RLock() // Read lock
    defer sd.rwmutex.RUnlock()
    value, exists := sd.data[key]
    return value, exists
}

func (sd *SafeData) Set(key string, value int) {
    sd.rwmutex.Lock() // Write lock
    defer sd.rwmutex.Unlock()
    sd.data[key] = value
}

func (sd *SafeData) Delete(key string) {
    sd.rwmutex.Lock() // Write lock
    defer sd.rwmutex.Unlock()
    delete(sd.data, key)
}

func (sd *SafeData) Size() int {
    sd.rwmutex.RLock() // Read lock
    defer sd.rwmutex.RUnlock()
    return len(sd.data)
}

func rwmutexWithStruct() {
    fmt.Println("=== RWMutex with Struct ===")
    
    data := NewSafeData()
    
    // Multiple readers
    for i := 0; i < 3; i++ {
        go func(id int) {
            for j := 0; j < 5; j++ {
                key := fmt.Sprintf("key%d", j)
                if value, exists := data.Get(key); exists {
                    fmt.Printf("Reader %d: %s = %d\n", id, key, value)
                } else {
                    fmt.Printf("Reader %d: %s not found\n", id, key)
                }
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // One writer
    go func() {
        for i := 0; i < 5; i++ {
            key := fmt.Sprintf("key%d", i)
            data.Set(key, i*10)
            fmt.Printf("Writer: Set %s = %d\n", key, i*10)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    rwmutexWithStruct()
}
```

### Example 3: Performance Comparison

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func performanceComparison() {
    fmt.Println("=== Performance Comparison ===")
    
    const iterations = 100000
    const readers = 10
    
    // Test with regular Mutex
    start := time.Now()
    var counter1 int
    var mutex sync.Mutex
    
    for i := 0; i < readers; i++ {
        go func() {
            for j := 0; j < iterations/readers; j++ {
                mutex.Lock()
                _ = counter1
                mutex.Unlock()
            }
        }()
    }
    
    time.Sleep(1 * time.Second)
    mutexTime := time.Since(start)
    
    // Test with RWMutex
    start = time.Now()
    var counter2 int
    var rwmutex sync.RWMutex
    
    for i := 0; i < readers; i++ {
        go func() {
            for j := 0; j < iterations/readers; j++ {
                rwmutex.RLock()
                _ = counter2
                rwmutex.RUnlock()
            }
        }()
    }
    
    time.Sleep(1 * time.Second)
    rwmutexTime := time.Since(start)
    
    fmt.Printf("Mutex time: %v\n", mutexTime)
    fmt.Printf("RWMutex time: %v\n", rwmutexTime)
    fmt.Printf("RWMutex speedup: %.2fx\n", float64(mutexTime)/float64(rwmutexTime))
}

func main() {
    performanceComparison()
}
```

### Example 4: Read-Heavy Workload

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    data   map[string]string
    rwmutex sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) (string, bool) {
    c.rwmutex.RLock()
    defer c.rwmutex.RUnlock()
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Set(key, value string) {
    c.rwmutex.Lock()
    defer c.rwmutex.Unlock()
    c.data[key] = value
}

func readHeavyWorkload() {
    fmt.Println("=== Read-Heavy Workload ===")
    
    cache := NewCache()
    
    // Initialize cache
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("key%d", i)
        cache.Set(key, fmt.Sprintf("value%d", i))
    }
    
    // Multiple readers
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 20; j++ {
                key := fmt.Sprintf("key%d", j%10)
                if value, exists := cache.Get(key); exists {
                    fmt.Printf("Reader %d: %s = %s\n", id, key, value)
                }
                time.Sleep(50 * time.Millisecond)
            }
        }(i)
    }
    
    // Occasional writer
    go func() {
        for i := 0; i < 3; i++ {
            key := fmt.Sprintf("key%d", i)
            value := fmt.Sprintf("updated_value%d", i)
            cache.Set(key, value)
            fmt.Printf("Writer: Updated %s = %s\n", key, value)
            time.Sleep(500 * time.Millisecond)
        }
    }()
    
    time.Sleep(3 * time.Second)
}

func main() {
    readHeavyWorkload()
}
```

### Example 5: Write-Heavy Workload

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func writeHeavyWorkload() {
    fmt.Println("=== Write-Heavy Workload ===")
    
    var counter int
    var rwmutex sync.RWMutex
    
    // Multiple writers
    for i := 0; i < 3; i++ {
        go func(id int) {
            for j := 0; j < 5; j++ {
                rwmutex.Lock() // Write lock
                counter++
                fmt.Printf("Writer %d: counter = %d\n", id, counter)
                rwmutex.Unlock() // Write unlock
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // Occasional reader
    go func() {
        for i := 0; i < 10; i++ {
            rwmutex.RLock() // Read lock
            value := counter
            rwmutex.RUnlock() // Read unlock
            fmt.Printf("Reader: counter = %d\n", value)
            time.Sleep(150 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    writeHeavyWorkload()
}
```

### Example 6: RWMutex with TryLock

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func rwmutexWithTryLock() {
    fmt.Println("=== RWMutex with TryLock ===")
    
    var counter int
    var rwmutex sync.RWMutex
    
    // Goroutine 1: holds write lock
    go func() {
        rwmutex.Lock()
        fmt.Println("Writer 1: Lock acquired")
        time.Sleep(2 * time.Second)
        rwmutex.Unlock()
        fmt.Println("Writer 1: Lock released")
    }()
    
    // Goroutine 2: tries to acquire read lock
    go func() {
        time.Sleep(500 * time.Millisecond)
        if rwmutex.TryRLock() {
            fmt.Println("Reader 1: Read lock acquired with TryRLock")
            _ = counter
            rwmutex.RUnlock()
        } else {
            fmt.Println("Reader 1: Could not acquire read lock")
        }
    }()
    
    // Goroutine 3: tries to acquire write lock
    go func() {
        time.Sleep(1 * time.Second)
        if rwmutex.TryLock() {
            fmt.Println("Writer 2: Write lock acquired with TryLock")
            counter++
            rwmutex.Unlock()
        } else {
            fmt.Println("Writer 2: Could not acquire write lock")
        }
    }()
    
    time.Sleep(3 * time.Second)
}

func main() {
    rwmutexWithTryLock()
}
```

### Example 7: RWMutex Best Practices

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Database struct {
    data   map[string]interface{}
    rwmutex sync.RWMutex
}

func NewDatabase() *Database {
    return &Database{
        data: make(map[string]interface{}),
    }
}

func (db *Database) Read(key string) (interface{}, bool) {
    db.rwmutex.RLock()
    defer db.rwmutex.RUnlock()
    value, exists := db.data[key]
    return value, exists
}

func (db *Database) Write(key string, value interface{}) {
    db.rwmutex.Lock()
    defer db.rwmutex.Unlock()
    db.data[key] = value
}

func (db *Database) ReadAll() map[string]interface{} {
    db.rwmutex.RLock()
    defer db.rwmutex.RUnlock()
    
    // Create a copy to avoid holding lock too long
    result := make(map[string]interface{})
    for k, v := range db.data {
        result[k] = v
    }
    return result
}

func rwmutexBestPractices() {
    fmt.Println("=== RWMutex Best Practices ===")
    
    db := NewDatabase()
    
    // Initialize data
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("key%d", i)
        db.Write(key, fmt.Sprintf("value%d", i))
    }
    
    // Multiple readers
    for i := 0; i < 3; i++ {
        go func(id int) {
            for j := 0; j < 5; j++ {
                key := fmt.Sprintf("key%d", j)
                if value, exists := db.Read(key); exists {
                    fmt.Printf("Reader %d: %s = %v\n", id, key, value)
                }
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // One writer
    go func() {
        for i := 0; i < 3; i++ {
            key := fmt.Sprintf("newkey%d", i)
            value := fmt.Sprintf("newvalue%d", i)
            db.Write(key, value)
            fmt.Printf("Writer: Set %s = %s\n", key, value)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
}

func main() {
    rwmutexBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **RWMutex allows multiple readers**: Better for read-heavy workloads
2. **Only one writer at a time**: Writers are mutually exclusive
3. **Readers and writers are exclusive**: Can't read while writing
4. **Better performance**: For read-heavy scenarios
5. **More complex**: Requires understanding of read/write patterns

## 🎯 When to Use RWMutex

1. **Read-heavy workloads**: Many readers, few writers
2. **Caching systems**: Frequently read, occasionally updated
3. **Configuration data**: Read often, updated rarely
4. **Lookup tables**: Read frequently, updated infrequently

## 🎯 When NOT to Use RWMutex

1. **Write-heavy workloads**: Regular mutex might be better
2. **Simple scenarios**: Regular mutex is simpler
3. **Equal read/write**: No performance benefit
4. **Complex locking**: Can lead to deadlocks

## 🎯 Best Practices

1. **Use RLock for reading**:
   ```go
   rwmutex.RLock()
   defer rwmutex.RUnlock()
   // Read data
   ```

2. **Use Lock for writing**:
   ```go
   rwmutex.Lock()
   defer rwmutex.Unlock()
   // Write data
   ```

3. **Keep critical sections short**:
   ```go
   // BAD - long critical section
   rwmutex.RLock()
   // Long computation
   rwmutex.RUnlock()
   
   // GOOD - short critical section
   rwmutex.RLock()
   // Quick read
   rwmutex.RUnlock()
   // Do computation outside
   ```

## 🎯 Common Pitfalls

1. **Using Lock instead of RLock**:
   ```go
   // BAD - unnecessary blocking
   rwmutex.Lock()
   value := data[key]
   rwmutex.Unlock()
   
   // GOOD - allows multiple readers
   rwmutex.RLock()
   value := data[key]
   rwmutex.RUnlock()
   ```

2. **Deadlock with multiple locks**:
   ```go
   // BAD - potential deadlock
   rwmutex1.RLock()
   rwmutex2.Lock()
   // ... code ...
   rwmutex2.Unlock()
   rwmutex1.RUnlock()
   ```

3. **Not understanding read/write exclusion**:
   ```go
   // BAD - readers block writers
   rwmutex.RLock()
   // Long read operation
   rwmutex.RUnlock()
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a file system cache:
- Multiple goroutines read files from cache
- One goroutine occasionally updates the cache
- Use RWMutex to optimize for read-heavy workload
- Show the performance difference between Mutex and RWMutex

**Hint**: Use a map to simulate the cache and measure the time it takes to perform many read operations.

## 🚀 Next Steps

Now that you understand RWMutex, let's learn about **WaitGroup** in the next file: `14-waitgroup.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
