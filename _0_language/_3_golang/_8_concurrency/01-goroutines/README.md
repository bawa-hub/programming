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

# 🧪 Goroutines Deep Dive - Complete Testing Guide

## 📋 Prerequisites
Make sure you're in the correct directory:
```bash
cd /Users/vikramkumar/CS/_1_dsa/_0_language/_3_golang/_concurrency/01-goroutines
```

## 🚀 Basic Testing Commands

### 1. **Test Basic Examples**
```bash
go run . basic
```
**Expected Output:** 8 examples covering basic goroutine concepts, performance comparison, and common pitfalls.

### 2. **Test All Exercises**
```bash
go run . exercises
```
**Expected Output:** 8 hands-on exercises including worker pools, communication patterns, and error handling.

### 3. **Test Advanced Patterns**
```bash
go run . advanced
```
**Expected Output:** 7 advanced patterns including dynamic pools, circuit breakers, and graceful shutdown.

### 4. **Test Everything Together**
```bash
go run . all
```
**Expected Output:** All basic examples, exercises, and advanced patterns in sequence.

### 5. **Test Help/Usage**
```bash
go run .
```
**Expected Output:** Usage information and available commands.

## 🔍 Advanced Testing Commands

### 6. **Race Detection Testing**
```bash
go run -race . basic
```
**Expected Output:** Should show intentional race conditions in educational examples (this is good for learning!).

### 7. **Race Detection on Exercises**
```bash
go run -race . exercises
```
**Expected Output:** Should be race-free (exercises demonstrate proper patterns).

### 8. **Race Detection on Advanced Patterns**
```bash
go run -race . advanced
```
**Expected Output:** Should be race-free (advanced patterns use proper synchronization).

### 9. **Build Testing**
```bash
go build .
```
**Expected Output:** Should compile without errors.

### 10. **Lint Testing**
```bash
go vet .
```
**Expected Output:** Should pass without warnings.

## 🎯 Individual Function Testing

### 11. **Test Specific Examples (if you want to modify code)**
You can test individual functions by creating a simple test file:

```bash
# Create a test file
cat > test_individual.go << 'EOF'
package main

import (
    "fmt"
    "time"
)

func main() {
    // Test individual functions
    fmt.Println("Testing basicGoroutine:")
    basicGoroutine()
    
    time.Sleep(100 * time.Millisecond)
    
    fmt.Println("\nTesting multipleGoroutines:")
    multipleGoroutines()
    
    time.Sleep(1 * time.Second)
}
EOF

# Run the test
go run test_individual.go

# Clean up
rm test_individual.go
```

## 🔧 Performance Testing

### 12. **Benchmark Testing**
```bash
# Create a benchmark file
cat > benchmark_test.go << 'EOF'
package main

import (
    "sync"
    "testing"
)

func BenchmarkGoroutineCreation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var wg sync.WaitGroup
        wg.Add(1)
        go func() {
            defer wg.Done()
            // Simulate work
            _ = i * i
        }()
        wg.Wait()
    }
}

func BenchmarkFunctionCall(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Simulate work
        _ = i * i
    }
}
EOF

# Run benchmarks
go test -bench=.

# Clean up
rm benchmark_test.go
```

### 13. **Memory Profiling**
```bash
# Run with memory profiling
go run . basic -memprofile=mem.prof

# Analyze memory profile (if you have pprof installed)
go tool pprof mem.prof
```

## 🐛 Debugging Commands

### 14. **Verbose Testing**
```bash
go run -v . basic
```

### 15. **Debug Information**
```bash
# Show goroutine information
go run . basic 2>&1 | grep -i goroutine
```

### 16. **Trace Analysis**
```bash
# Run with trace
go run . basic -trace=trace.out

# Analyze trace (if you have trace viewer)
go tool trace trace.out
```

## 📊 Expected Test Results

### ✅ **Successful Test Indicators:**
- All commands run without compilation errors
- Basic examples show goroutine behavior
- Exercises demonstrate proper synchronization
- Advanced patterns show production-ready code
- Race detection shows intentional educational race conditions
- Performance comparison shows goroutine overhead

### ⚠️ **Expected Warnings (These are GOOD for learning):**
- Race detection warnings in basic examples (intentional)
- Performance overhead warnings (expected)
- Goroutine count variations (normal)

## 🎯 Testing Checklist

- [ ] `go run . basic` - Basic examples work
- [ ] `go run . exercises` - All exercises complete
- [ ] `go run . advanced` - Advanced patterns work
- [ ] `go run . all` - Everything runs together
- [ ] `go run .` - Help shows correctly
- [ ] `go run -race . basic` - Race detection works
- [ ] `go run -race . exercises` - Exercises are race-free
- [ ] `go run -race . advanced` - Advanced patterns are race-free
- [ ] `go build .` - Compiles without errors
- [ ] `go vet .` - Passes static analysis

## 🚀 Quick Test Script

Create this script for automated testing:

```bash
#!/bin/bash
# quick_test.sh

echo "🧪 Running Quick Goroutines Test Suite"
echo "======================================"

echo "1. Testing basic examples..."
go run . basic > /dev/null 2>&1 && echo "✅ Basic examples: PASS" || echo "❌ Basic examples: FAIL"

echo "2. Testing exercises..."
go run . exercises > /dev/null 2>&1 && echo "✅ Exercises: PASS" || echo "❌ Exercises: FAIL"

echo "3. Testing advanced patterns..."
go run . advanced > /dev/null 2>&1 && echo "✅ Advanced patterns: PASS" || echo "❌ Advanced patterns: FAIL"

echo "4. Testing compilation..."
go build . > /dev/null 2>&1 && echo "✅ Compilation: PASS" || echo "❌ Compilation: FAIL"

echo "5. Testing race detection..."
go run -race . basic > /dev/null 2>&1 && echo "✅ Race detection: PASS" || echo "❌ Race detection: FAIL"

echo "======================================"
echo "🎉 Test suite completed!"
```

Make it executable and run:
```bash
chmod +x quick_test.sh
./quick_test.sh
```

## 🎯 What Each Test Validates

| Command | Validates |
|---------|-----------|
| `go run . basic` | Goroutine creation, communication, lifecycle |
| `go run . exercises` | Proper synchronization patterns |
| `go run . advanced` | Production-ready concurrency patterns |
| `go run -race . basic` | Educational race conditions (intentional) |
| `go run -race . exercises` | Race-free synchronization |
| `go run -race . advanced` | Race-free advanced patterns |
| `go build .` | Code compiles correctly |
| `go vet .` | Code follows Go best practices |

## 🏆 Success Criteria

Your goroutines topic is ready when:
- ✅ All commands run without errors
- ✅ Race detection shows intentional educational races
- ✅ Exercises demonstrate proper patterns
- ✅ Advanced patterns show production quality
- ✅ Code compiles and passes static analysis

## 🚀 Ready for Next Topic?

Once all tests pass, you're ready to move to **Level 1, Topic 2: Channels Fundamentals**!

**Type "NEXT" to continue your journey to Go concurrency mastery!**

# 🚀 Goroutines Deep Dive - Command Reference

## 📋 Quick Reference Commands

### **Basic Testing**
```bash
# Test basic examples
go run . basic

# Test all exercises  
go run . exercises

# Test advanced patterns
go run . advanced

# Test everything together
go run . all

# Show help/usage
go run .
```

### **Advanced Testing**
```bash
# Race detection (shows intentional educational races)
go run -race . basic

# Race detection on exercises (should be race-free)
go run -race . exercises

# Race detection on advanced patterns (should be race-free)
go run -race . advanced

# Compilation test
go build .

# Static analysis (shows intentional educational warnings)
go vet .
```

### **Automated Testing**
```bash
# Run the quick test suite
./quick_test.sh

# Make script executable (if needed)
chmod +x quick_test.sh
```

### **Performance Testing**
```bash
# Run with verbose output
go run -v . basic

# Run with trace (creates trace.out)
go run . basic -trace=trace.out

# Run with memory profile
go run . basic -memprofile=mem.prof

# Analyze trace
go tool trace trace.out

# Analyze memory profile
go tool pprof mem.prof
```

### **Individual Testing**
```bash
# Test specific functions (create test file)
cat > test.go << 'EOF'
package main
import "time"
func main() {
    basicGoroutine()
    time.Sleep(100 * time.Millisecond)
}
EOF
go run test.go
rm test.go
```

## 🎯 Expected Results

| Command | Expected Result |
|---------|----------------|
| `go run . basic` | 8 examples with goroutine behavior |
| `go run . exercises` | 8 exercises with proper synchronization |
| `go run . advanced` | 7 advanced patterns working |
| `go run -race . basic` | Shows intentional race conditions |
| `go run -race . exercises` | No race conditions (clean) |
| `go run -race . advanced` | No race conditions (clean) |
| `go vet .` | Shows intentional variable capture warnings |
| `go build .` | Compiles successfully |

## 🏆 Success Indicators

✅ **All commands run without errors**  
✅ **Race detection shows intentional educational races**  
✅ **Exercises demonstrate proper patterns**  
✅ **Advanced patterns show production quality**  
✅ **Static analysis shows intentional warnings**  
✅ **Code compiles and builds successfully**

## 🚀 Ready for Next Topic?

Once all tests pass, you're ready for **Level 1, Topic 2: Channels Fundamentals**!

**Type "NEXT" to continue your journey to Go concurrency mastery!**
