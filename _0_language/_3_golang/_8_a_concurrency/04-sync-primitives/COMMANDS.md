# 🚀 Synchronization Primitives - Quick Commands

## 📋 Basic Commands

### **Run Examples**
```bash
# Basic synchronization examples
go run . basic

# All exercises
go run . exercises

# Advanced patterns
go run . advanced

# Everything
go run . all
```

### **Testing Commands**
```bash
# Quick test suite
./quick_test.sh

# Compilation test
go build .

# Race detection
go run -race . basic

# Static analysis
go vet .
```

## 🔍 Individual Examples

### **Basic Mutex**
```bash
go run . basic | grep -A 10 "Basic Mutex"
```

### **RWMutex**
```bash
go run . basic | grep -A 15 "RWMutex"
```

### **WaitGroup**
```bash
go run . basic | grep -A 10 "WaitGroup"
```

### **Once**
```bash
go run . basic | grep -A 10 "Once"
```

### **Condition Variables**
```bash
go run . basic | grep -A 10 "Cond"
```

### **Atomic Operations**
```bash
go run . basic | grep -A 10 "Atomic"
```

### **Concurrent Map**
```bash
go run . basic | grep -A 10 "Concurrent Map"
```

### **Object Pool**
```bash
go run . basic | grep -A 10 "Object Pool"
```

### **Performance Comparison**
```bash
go run . basic | grep -A 5 "Performance"
```

### **Deadlock Prevention**
```bash
go run . basic | grep -A 10 "Deadlock"
```

### **Race Condition Detection**
```bash
go run . basic | grep -A 10 "Race Condition"
```

### **Common Pitfalls**
```bash
go run . basic | grep -A 20 "Common Pitfalls"
```

## 🧪 Exercise Commands

### **Exercise 1: Basic Mutex**
```bash
go run . exercises | grep -A 10 "Exercise 1"
```

### **Exercise 2: RWMutex**
```bash
go run . exercises | grep -A 15 "Exercise 2"
```

### **Exercise 3: WaitGroup**
```bash
go run . exercises | grep -A 10 "Exercise 3"
```

### **Exercise 4: Once**
```bash
go run . exercises | grep -A 10 "Exercise 4"
```

### **Exercise 5: Cond**
```bash
go run . exercises | grep -A 10 "Exercise 5"
```

### **Exercise 6: Atomic Operations**
```bash
go run . exercises | grep -A 10 "Exercise 6"
```

### **Exercise 7: Concurrent Map**
```bash
go run . exercises | grep -A 10 "Exercise 7"
```

### **Exercise 8: Object Pool**
```bash
go run . exercises | grep -A 10 "Exercise 8"
```

### **Exercise 9: Deadlock Prevention**
```bash
go run . exercises | grep -A 10 "Exercise 9"
```

### **Exercise 10: Performance Comparison**
```bash
go run . exercises | grep -A 10 "Exercise 10"
```

## 🚀 Advanced Pattern Commands

### **Pattern 1: Thread-Safe Counter**
```bash
go run . advanced | grep -A 10 "Thread-Safe Counter"
```

### **Pattern 2: Priority RWMutex**
```bash
go run . advanced | grep -A 10 "Priority RWMutex"
```

### **Pattern 3: WaitGroup with Timeout**
```bash
go run . advanced | grep -A 10 "WaitGroup with Timeout"
```

### **Pattern 4: Once with Error Handling**
```bash
go run . advanced | grep -A 10 "Once with Error Handling"
```

### **Pattern 5: Condition Variable with Timeout**
```bash
go run . advanced | grep -A 10 "Condition Variable with Timeout"
```

### **Pattern 6: Atomic Counter with Statistics**
```bash
go run . advanced | grep -A 10 "Atomic Counter with Statistics"
```

### **Pattern 7: Concurrent Map with Statistics**
```bash
go run . advanced | grep -A 10 "Concurrent Map with Statistics"
```

### **Pattern 8: Object Pool with Statistics**
```bash
go run . advanced | grep -A 10 "Object Pool with Statistics"
```

### **Pattern 9: Barrier Synchronization**
```bash
go run . advanced | grep -A 10 "Barrier Synchronization"
```

### **Pattern 10: Semaphore**
```bash
go run . advanced | grep -A 10 "Semaphore"
```

## 🔧 Debugging Commands

### **Verbose Output**
```bash
go run -v . basic
```

### **Race Detection with Details**
```bash
go run -race . basic 2>&1 | grep -A 5 "WARNING"
```

### **Static Analysis with Details**
```bash
go vet . -v
```

### **Build with Details**
```bash
go build -v .
```

## 📊 Performance Commands

### **CPU Profiling**
```bash
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

### **Memory Profiling**
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

### **Benchmark Testing**
```bash
go test -bench=. -benchmem
```

## 🎯 Quick Verification

### **Check All Examples Work**
```bash
go run . all > /dev/null && echo "✅ All examples work"
```

### **Check Race Detection**
```bash
go run -race . basic 2>&1 | grep -q "WARNING: DATA RACE" && echo "✅ Race detection working"
```

### **Check Compilation**
```bash
go build . && echo "✅ Compilation successful"
```

### **Check Static Analysis**
```bash
go vet . && echo "✅ Static analysis passed"
```

## 🚀 Quick Test Suite

### **Run All Tests**
```bash
./quick_test.sh
```

### **Test Individual Components**
```bash
# Test basic examples
go run . basic > /dev/null && echo "✅ Basic: PASS" || echo "❌ Basic: FAIL"

# Test exercises
go run . exercises > /dev/null && echo "✅ Exercises: PASS" || echo "❌ Exercises: FAIL"

# Test advanced patterns
go run . advanced > /dev/null && echo "✅ Advanced: PASS" || echo "❌ Advanced: FAIL"

# Test compilation
go build . > /dev/null && echo "✅ Compilation: PASS" || echo "❌ Compilation: FAIL"

# Test race detection
go run -race . basic 2>&1 | grep -q "WARNING: DATA RACE" && echo "✅ Race detection: PASS" || echo "❌ Race detection: FAIL"

# Test static analysis
go vet . > /dev/null && echo "✅ Static analysis: PASS" || echo "❌ Static analysis: FAIL"
```

## 📝 Output Examples

### **Expected Basic Output**
```
🚀 Synchronization Primitives Examples
======================================
1. Basic Mutex
==============
Goroutine 4 completed
Goroutine 2 completed
Goroutine 0 completed
Goroutine 1 completed
Goroutine 3 completed
Final counter value: 5000
```

### **Expected Exercise Output**
```
Exercise 1: Basic Mutex
=======================
Goroutine 2 completed
Goroutine 0 completed
Goroutine 1 completed
Final counter value: 300
```

### **Expected Advanced Output**
```
🚀 Advanced Synchronization Patterns
====================================

1. Thread-Safe Counter with Metrics:
Counters: map[key0:10 key1:10 key2:10 key3:10 key4:10]
Metrics: map[increments:50]
```

## 🎉 Success Indicators

- ✅ All examples run without errors
- ✅ Race detection identifies intentional race
- ✅ Performance comparisons show expected results
- ✅ No deadlocks or hangs
- ✅ Proper synchronization behavior
- ✅ All tests pass

**🚀 Ready for Worker Pool Pattern!**
