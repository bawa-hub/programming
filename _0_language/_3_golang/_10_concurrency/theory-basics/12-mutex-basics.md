# 12 - Mutex Basics

## 🎯 Learning Objectives
- Understand what mutexes are and why they're needed
- Learn how to use `sync.Mutex` to protect shared data
- Master mutex locking and unlocking patterns
- Practice fixing race conditions with mutexes
- Understand mutex performance implications

## 📚 Theory

### What is a Mutex?

A **mutex** (mutual exclusion) is a synchronization primitive that ensures only one goroutine can access shared data at a time.

**Key characteristics:**
- **Mutual exclusion**: Only one goroutine can hold the lock
- **Blocking**: Other goroutines wait until lock is released
- **Reentrant**: Same goroutine can lock multiple times
- **Deadlock risk**: Can cause deadlocks if not used carefully

### Why Do We Need Mutexes?

**Problem**: Race conditions occur when multiple goroutines access shared data concurrently.

**Solution**: Mutexes provide exclusive access to shared data, preventing race conditions.

### Mutex Operations

1. **Lock**: `mutex.Lock()` - Acquire exclusive access
2. **Unlock**: `mutex.Unlock()` - Release exclusive access
3. **TryLock**: `mutex.TryLock()` - Try to acquire lock without blocking

## 💻 Code Examples

### Example 1: Basic Mutex Usage

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicMutex() {
    fmt.Println("=== Basic Mutex Usage ===")
    
    var counter int
    var mutex sync.Mutex
    
    // Multiple goroutines increment counter safely
    for i := 0; i < 1000; i++ {
        go func() {
            mutex.Lock()   // Lock before accessing shared data
            counter++      // Safe access to shared data
            mutex.Unlock() // Unlock after accessing
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d (should be 1000)\n", counter)
}

func main() {
    basicMutex()
}
```

**Run this code:**
```bash
go run 12-mutex-basics.go
```

### Example 2: Mutex with Struct Methods

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    value int
    mutex sync.Mutex
}

func (c *SafeCounter) Increment() {
    c.mutex.Lock()
    defer c.mutex.Unlock() // Ensure unlock even if panic occurs
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.value
}

func (c *SafeCounter) Add(amount int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.value += amount
}

func mutexWithStruct() {
    fmt.Println("=== Mutex with Struct Methods ===")
    
    counter := &SafeCounter{}
    
    // Multiple goroutines safely increment counter
    for i := 0; i < 1000; i++ {
        go func() {
            counter.Increment()
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d (should be 1000)\n", counter.Value())
}

func main() {
    mutexWithStruct()
}
```

### Example 3: Mutex with Slices

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeSlice struct {
    data  []int
    mutex sync.Mutex
}

func (s *SafeSlice) Append(value int) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.data = append(s.data, value)
}

func (s *SafeSlice) Length() int {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return len(s.data)
}

func (s *SafeSlice) Get(index int) (int, bool) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    if index < len(s.data) {
        return s.data[index], true
    }
    return 0, false
}

func mutexWithSlices() {
    fmt.Println("=== Mutex with Slices ===")
    
    slice := &SafeSlice{}
    
    // Multiple goroutines safely append to slice
    for i := 0; i < 100; i++ {
        go func(i int) {
            slice.Append(i)
        }(i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Slice length: %d (should be 100)\n", slice.Length())
}

func main() {
    mutexWithSlices()
}
```

### Example 4: Mutex with Maps

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeMap struct {
    data  map[string]int
    mutex sync.Mutex
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        data: make(map[string]int),
    }
}

func (m *SafeMap) Set(key string, value int) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.data[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    value, exists := m.data[key]
    return value, exists
}

func (m *SafeMap) Delete(key string) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    delete(m.data, key)
}

func (m *SafeMap) Size() int {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    return len(m.data)
}

func mutexWithMaps() {
    fmt.Println("=== Mutex with Maps ===")
    
    m := NewSafeMap()
    
    // Multiple goroutines safely access map
    for i := 0; i < 100; i++ {
        go func(i int) {
            key := fmt.Sprintf("key%d", i)
            m.Set(key, i)
        }(i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Map size: %d (should be 100)\n", m.Size())
}

func main() {
    mutexWithMaps()
}
```

### Example 5: Mutex with TryLock

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func mutexWithTryLock() {
    fmt.Println("=== Mutex with TryLock ===")
    
    var counter int
    var mutex sync.Mutex
    
    // Goroutine 1: holds lock for a while
    go func() {
        mutex.Lock()
        fmt.Println("Goroutine 1: Lock acquired")
        time.Sleep(2 * time.Second)
        mutex.Unlock()
        fmt.Println("Goroutine 1: Lock released")
    }()
    
    // Goroutine 2: tries to acquire lock
    go func() {
        time.Sleep(500 * time.Millisecond) // Wait a bit
        if mutex.TryLock() {
            fmt.Println("Goroutine 2: Lock acquired with TryLock")
            counter++
            mutex.Unlock()
        } else {
            fmt.Println("Goroutine 2: Could not acquire lock")
        }
    }()
    
    time.Sleep(3 * time.Second)
}

func main() {
    mutexWithTryLock()
}
```

### Example 6: Mutex Performance Comparison

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func performanceComparison() {
    fmt.Println("=== Mutex Performance Comparison ===")
    
    const iterations = 1000000
    
    // Without mutex (race condition)
    start := time.Now()
    var counter1 int
    for i := 0; i < iterations; i++ {
        counter1++
    }
    noMutexTime := time.Since(start)
    
    // With mutex
    start = time.Now()
    var counter2 int
    var mutex sync.Mutex
    for i := 0; i < iterations; i++ {
        mutex.Lock()
        counter2++
        mutex.Unlock()
    }
    mutexTime := time.Since(start)
    
    fmt.Printf("Without mutex: %v\n", noMutexTime)
    fmt.Printf("With mutex: %v\n", mutexTime)
    fmt.Printf("Overhead: %.2fx\n", float64(mutexTime)/float64(noMutexTime))
}

func main() {
    performanceComparison()
}
```

### Example 7: Mutex with Defer

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func mutexWithDefer() {
    fmt.Println("=== Mutex with Defer ===")
    
    var counter int
    var mutex sync.Mutex
    
    // Function that might panic
    increment := func() {
        mutex.Lock()
        defer mutex.Unlock() // Ensures unlock even if panic occurs
        
        counter++
        if counter == 5 {
            panic("Simulated panic")
        }
    }
    
    // Multiple goroutines call increment
    for i := 0; i < 10; i++ {
        go func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Printf("Panic recovered: %v\n", r)
                }
            }()
            increment()
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Counter value: %d\n", counter)
}

func main() {
    mutexWithDefer()
}
```

### Example 8: Mutex Best Practices

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type BankAccount struct {
    balance int
    mutex   sync.Mutex
}

func (ba *BankAccount) Deposit(amount int) {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount int) bool {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    
    if ba.balance >= amount {
        ba.balance -= amount
        return true
    }
    return false
}

func (ba *BankAccount) Balance() int {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    return ba.balance
}

func mutexBestPractices() {
    fmt.Println("=== Mutex Best Practices ===")
    
    account := &BankAccount{balance: 1000}
    
    // Multiple goroutines perform transactions
    for i := 0; i < 100; i++ {
        go func() {
            account.Deposit(10)
        }()
        
        go func() {
            account.Withdraw(5)
        }()
    }
    
    time.Sleep(1 * time.Second)
    fmt.Printf("Final balance: %d\n", account.Balance())
}

func main() {
    mutexBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Mutex provides exclusive access**: Only one goroutine at a time
2. **Always unlock**: Use `defer mutex.Unlock()` to ensure unlocking
3. **Protect shared data**: Lock before accessing, unlock after
4. **Performance cost**: Mutexes have overhead
5. **Deadlock risk**: Can cause deadlocks if not used carefully

## 🎯 When to Use Mutexes

1. **Shared data access**: When multiple goroutines access shared data
2. **Critical sections**: Code that must be executed atomically
3. **Data structures**: Protecting slices, maps, structs
4. **Resource management**: Controlling access to resources

## 🎯 Best Practices

1. **Use defer for unlocking**:
   ```go
   mutex.Lock()
   defer mutex.Unlock()
   // Access shared data
   ```

2. **Keep critical sections short**:
   ```go
   // BAD - long critical section
   mutex.Lock()
   // Long computation
   mutex.Unlock()
   
   // GOOD - short critical section
   // Do computation outside
   mutex.Lock()
   // Quick access to shared data
   mutex.Unlock()
   ```

3. **Don't forget to unlock**:
   ```go
   // BAD - might forget to unlock
   mutex.Lock()
   if condition {
       return // Forgot to unlock!
   }
   mutex.Unlock()
   
   // GOOD - use defer
   mutex.Lock()
   defer mutex.Unlock()
   if condition {
       return // Automatically unlocked
   }
   ```

## 🎯 Common Pitfalls

1. **Deadlock**:
   ```go
   // BAD - deadlock
   mutex1.Lock()
   mutex2.Lock()
   // ... code ...
   mutex1.Unlock()
   mutex2.Unlock()
   
   // GOOD - same order
   mutex1.Lock()
   mutex2.Lock()
   // ... code ...
   mutex2.Unlock()
   mutex1.Unlock()
   ```

2. **Not unlocking**:
   ```go
   // BAD - might not unlock
   mutex.Lock()
   if condition {
       return // Forgot to unlock!
   }
   mutex.Unlock()
   
   // GOOD - use defer
   mutex.Lock()
   defer mutex.Unlock()
   if condition {
       return // Automatically unlocked
   }
   ```

3. **Locking too much**:
   ```go
   // BAD - unnecessary locking
   mutex.Lock()
   for i := 0; i < 1000; i++ {
       // Do work
   }
   mutex.Unlock()
   
   // GOOD - lock only when needed
   for i := 0; i < 1000; i++ {
       mutex.Lock()
       // Quick access to shared data
       mutex.Unlock()
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a bank account with multiple transactions:
- Multiple goroutines perform deposits and withdrawals
- Use mutexes to protect the account balance
- Show how mutexes prevent race conditions
- Compare performance with and without mutexes

**Hint**: Use the race detector to verify that your mutex implementation prevents race conditions.

## 🚀 Next Steps

Now that you understand basic mutexes, let's learn about **read-write mutexes** in the next file: `13-rwmutex.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
