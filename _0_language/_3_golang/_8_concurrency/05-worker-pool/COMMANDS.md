# 🚀 Worker Pool Pattern - Quick Commands

## 📋 Basic Commands

### **Run Examples**
```bash
# Basic worker pool examples
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

### **Basic Worker Pool**
```bash
go run . basic | grep -A 10 "Basic Worker Pool"
```

### **Buffered Worker Pool**
```bash
go run . basic | grep -A 10 "Buffered Worker Pool"
```

### **Dynamic Worker Pool**
```bash
go run . basic | grep -A 10 "Dynamic Worker Pool"
```

### **Priority Worker Pool**
```bash
go run . basic | grep -A 10 "Priority Worker Pool"
```

### **Worker Pool with Results**
```bash
go run . basic | grep -A 10 "Worker Pool with Results"
```

### **Worker Pool with Error Handling**
```bash
go run . basic | grep -A 10 "Worker Pool with Error Handling"
```

### **Worker Pool with Timeout**
```bash
go run . basic | grep -A 10 "Worker Pool with Timeout"
```

### **Worker Pool with Rate Limiting**
```bash
go run . basic | grep -A 10 "Worker Pool with Rate Limiting"
```

### **Worker Pool with Metrics**
```bash
go run . basic | grep -A 10 "Worker Pool with Metrics"
```

### **Pipeline Worker Pool**
```bash
go run . basic | grep -A 10 "Pipeline Worker Pool"
```

### **Performance Comparison**
```bash
go run . basic | grep -A 5 "Performance Comparison"
```

### **Common Pitfalls**
```bash
go run . basic | grep -A 20 "Common Pitfalls"
```

## 🧪 Exercise Commands

### **Exercise 1: Basic Worker Pool**
```bash
go run . exercises | grep -A 10 "Exercise 1"
```

### **Exercise 2: Buffered Worker Pool**
```bash
go run . exercises | grep -A 10 "Exercise 2"
```

### **Exercise 3: Dynamic Worker Pool**
```bash
go run . exercises | grep -A 10 "Exercise 3"
```

### **Exercise 4: Priority Worker Pool**
```bash
go run . exercises | grep -A 10 "Exercise 4"
```

### **Exercise 5: Worker Pool with Results**
```bash
go run . exercises | grep -A 10 "Exercise 5"
```

### **Exercise 6: Worker Pool with Error Handling**
```bash
go run . exercises | grep -A 10 "Exercise 6"
```

### **Exercise 7: Worker Pool with Timeout**
```bash
go run . exercises | grep -A 10 "Exercise 7"
```

### **Exercise 8: Worker Pool with Rate Limiting**
```bash
go run . exercises | grep -A 10 "Exercise 8"
```

### **Exercise 9: Worker Pool with Metrics**
```bash
go run . exercises | grep -A 10 "Exercise 9"
```

### **Exercise 10: Pipeline Worker Pool**
```bash
go run . exercises | grep -A 10 "Exercise 10"
```

## 🚀 Advanced Pattern Commands

### **Pattern 1: Work Stealing Worker Pool**
```bash
go run . advanced | grep -A 10 "Work Stealing Worker Pool"
```

### **Pattern 2: Adaptive Worker Pool**
```bash
go run . advanced | grep -A 10 "Adaptive Worker Pool"
```

### **Pattern 3: Circuit Breaker Worker Pool**
```bash
go run . advanced | grep -A 10 "Circuit Breaker Worker Pool"
```

### **Pattern 4: Priority Queue Worker Pool**
```bash
go run . advanced | grep -A 10 "Priority Queue Worker Pool"
```

### **Pattern 5: Load Balancing Worker Pool**
```bash
go run . advanced | grep -A 10 "Load Balancing Worker Pool"
```

### **Pattern 6: Batch Processing Worker Pool**
```bash
go run . advanced | grep -A 10 "Batch Processing Worker Pool"
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
go run -race . basic > /dev/null && echo "✅ Race detection passed"
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
go run -race . basic > /dev/null && echo "✅ Race detection: PASS" || echo "❌ Race detection: FAIL"

# Test static analysis
go vet . > /dev/null && echo "✅ Static analysis: PASS" || echo "❌ Static analysis: FAIL"
```

## 📝 Output Examples

### **Expected Basic Output**
```
🚀 Worker Pool Pattern Examples
===============================
1. Basic Worker Pool
===================
Results:
Worker 0 processed job 0
  Job 0: Processed: Job 0 (took 42.209µs, worker 0)
```

### **Expected Exercise Output**
```
Exercise 1: Basic Worker Pool
=============================
Exercise 1 Results:
  Job 0: Exercise1: Exercise Job 0 (took 1.542µs, worker 1)
```

### **Expected Advanced Output**
```
🚀 Advanced Worker Pool Patterns
=================================

1. Work Stealing Worker Pool:
  Job 0: Work Stealing: Work Stealing Job 0 (worker 0)
```

## 🎉 Success Indicators

- ✅ All examples run without errors
- ✅ Race detection passes cleanly
- ✅ Performance comparisons show expected results
- ✅ No deadlocks or hangs
- ✅ Proper worker pool behavior
- ✅ All tests pass

**🚀 Ready for Pipeline Pattern!**
