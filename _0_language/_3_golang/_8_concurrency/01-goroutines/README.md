# 🚀 Goroutines Deep Dive: The Foundation of Go Concurrency

## 📚 Table of Contents
1. [What Are Goroutines?](#what-are-goroutines)
2. [Goroutine Lifecycle](#goroutine-lifecycle)
3. [M:N Threading Model](#mn-threading-model)
4. [Stack Management](#stack-management)
5. [Goroutine Scheduling](#goroutine-scheduling)
6. [Performance Characteristics](#performance-characteristics)
7. [Common Patterns](#common-patterns)
8. [Best Practices](#best-practices)
9. [Common Pitfalls](#common-pitfalls)
10. [Exercises](#exercises)

---

## 🎯 What Are Goroutines?

A **goroutine** is a lightweight thread managed by the Go runtime. Think of it as a function that can run concurrently with other functions.

### Key Characteristics:
- **Lightweight**: Start with only 2KB stack (grows as needed)
- **Cheap**: Can create millions of goroutines
- **Cooperative**: Uses cooperative multitasking
- **Managed**: Go runtime handles scheduling and lifecycle

### Basic Syntax:
```go
go functionName()        // Anonymous function
go func() { ... }()      // Inline function
go method()              // Method call
```

---

## 🔄 Goroutine Lifecycle

### 1. **Creation**
```go
go func() {
    fmt.Println("Hello from goroutine!")
}()
```

### 2. **Execution**
- Goroutine runs concurrently with main thread
- Go scheduler manages execution
- Can run on any available OS thread

### 3. **Termination**
- Function returns → goroutine ends
- Main function ends → all goroutines terminated
- Panic in goroutine → goroutine ends (unless recovered)

### 4. **Blocking States**
- Channel operations
- System calls
- Network I/O
- Mutex operations

---

## 🧵 M:N Threading Model

Go uses an **M:N threading model**:

- **M Goroutines** : **N OS Threads**
- **M >> N** (many more goroutines than threads)
- **GOMAXPROCS** controls N (number of OS threads)

### Components:
- **G (Goroutine)**: The goroutine itself
- **M (Machine)**: OS thread
- **P (Processor)**: Logical processor (context for running goroutines)

### Benefits:
- **Efficiency**: Thousands of goroutines on few threads
- **Scalability**: Better resource utilization
- **Simplicity**: No need to manage threads manually

---

## 📚 Stack Management

### Initial Stack Size:
- **2KB** per goroutine (very small!)
- **Grows dynamically** as needed
- **Shrinks** when possible (garbage collection)

### Stack Growth:
```go
func recursiveFunction(n int) {
    if n > 0 {
        recursiveFunction(n - 1)  // Stack grows
    }
}
```

### Stack vs Heap:
- **Stack**: Local variables, function parameters
- **Heap**: Dynamically allocated memory
- **Escape Analysis**: Compiler decides stack vs heap

---

## ⚙️ Goroutine Scheduling

### Work Stealing Algorithm:
1. **P** (Processor) has a local run queue
2. When local queue is empty, steal from other P's
3. **Global run queue** for new goroutines
4. **Network poller** for I/O operations

### Scheduling Points:
- Channel operations
- System calls
- Function calls (occasionally)
- Garbage collection
- Time slices (10ms)

### Preemption:
- **Cooperative**: Goroutines yield voluntarily
- **Preemptive**: Go 1.14+ has limited preemption
- **Sysmon**: Background thread monitors long-running goroutines

---

## 📊 Performance Characteristics

### Creation Overhead:
```go
// Very cheap to create
for i := 0; i < 1000000; i++ {
    go func() {
        // Do work
    }()
}
```

### Memory Usage:
- **2KB** initial stack
- **Grows to 1GB** maximum
- **Automatic shrinking** when possible

### Context Switching:
- **Much faster** than OS threads
- **User-space** scheduling
- **No kernel involvement** for most operations

---

## 🎨 Common Patterns

### 1. **Fire and Forget**
```go
go func() {
    // Do work in background
    processData()
}()
```

### 2. **Wait for Completion**
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
wg.Wait()
```

### 3. **Goroutine Pool**
```go
// Limit concurrent goroutines
semaphore := make(chan struct{}, 10)
for i := 0; i < 100; i++ {
    go func() {
        semaphore <- struct{}{} // Acquire
        defer func() { <-semaphore }() // Release
        // Do work
    }()
}
```

### 4. **Fan-Out Pattern**
```go
// Start multiple workers
for i := 0; i < numWorkers; i++ {
    go worker(i)
}
```

---

## ✅ Best Practices

### 1. **Always Use `go` Keyword**
```go
// ✅ Correct
go processData()

// ❌ Wrong - will block
processData()
```

### 2. **Handle Goroutine Lifecycle**
```go
// ✅ Use WaitGroup or channels
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

### 3. **Avoid Goroutine Leaks**
```go
// ✅ Use context for cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    select {
    case <-ctx.Done():
        return
    case <-time.After(5 * time.Second):
        // work
    }
}()
```

### 4. **Don't Share Memory**
```go
// ❌ Wrong - race condition
var counter int
go func() { counter++ }()
go func() { counter++ }()

// ✅ Correct - use channels
ch := make(chan int)
go func() { ch <- 1 }()
go func() { ch <- 1 }()
```

---

## ⚠️ Common Pitfalls

### 1. **Goroutine Leaks**
```go
// ❌ Leak - goroutine never exits
go func() {
    for {
        // infinite loop
    }
}()
```

### 2. **Race Conditions**
```go
// ❌ Race condition
var counter int
go func() { counter++ }()
go func() { counter++ }()
fmt.Println(counter) // Unpredictable result
```

### 3. **Blocking Main Thread**
```go
// ❌ Blocks main thread
go func() {
    time.Sleep(5 * time.Second)
}()
// Main thread exits before goroutine completes
```

### 4. **Variable Capture in Loops**
```go
// ❌ All goroutines see same value
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i) // Always prints 3
    }()
}

// ✅ Correct - pass value
for i := 0; i < 3; i++ {
    go func(i int) {
        fmt.Println(i) // Prints 0, 1, 2
    }(i)
}
```

---

## 🧪 Exercises

### Exercise 1: Basic Goroutines
Create a program that starts 5 goroutines, each printing a unique number.

### Exercise 2: Goroutine Synchronization
Use WaitGroup to wait for 3 goroutines to complete.

### Exercise 3: Goroutine Pool
Implement a worker pool with 3 workers processing 10 jobs.

### Exercise 4: Goroutine Communication
Create 2 goroutines that communicate through a channel.

### Exercise 5: Goroutine Lifecycle
Implement a goroutine that can be started, paused, and stopped.

---

## 🎯 Key Takeaways

1. **Goroutines are lightweight** - create them liberally
2. **Use channels for communication** - don't share memory
3. **Handle lifecycle properly** - avoid leaks
4. **Understand scheduling** - it affects performance
5. **Practice patterns** - they're the building blocks

---

## 🚀 Next Steps

Ready for the next topic? Let's move on to **Channels Fundamentals** where you'll learn how goroutines communicate!

**Run the examples in this directory to see goroutines in action!**
