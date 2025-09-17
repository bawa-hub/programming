# 🧪 Select Statement Mastery - Complete Testing Guide

## 📋 Prerequisites
Make sure you're in the correct directory:
```bash
cd /Users/vikramkumar/CS/_1_dsa/_0_language/_3_golang/_concurrency/03-select
```

## 🚀 Basic Testing Commands

### 1. **Test Basic Examples**
```bash
go run . basic
```
**Expected Output:** 12 examples covering basic select operations, non-blocking operations, timeouts, priority handling, multiplexing, loops, tickers, context, performance, error handling, and common pitfalls.

### 2. **Test All Exercises**
```bash
go run . exercises
```
**Expected Output:** 10 hands-on exercises including basic select, non-blocking operations, timeouts, priority handling, multiplexing, loops, error handling, dynamic select, tickers, and context.

### 3. **Test Advanced Patterns**
```bash
go run . advanced
```
**Expected Output:** 7 advanced patterns including event loops, rate limiters, load balancers, circuit breakers, message routers, priority pools, and context managers.

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
**Expected Output:** Should be race-free (select statements provide proper synchronization).

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
    fmt.Println("Testing basicSelect:")
    basicSelect()
    
    time.Sleep(100 * time.Millisecond)
    
    fmt.Println("\nTesting nonBlockingOperations:")
    nonBlockingOperations()
    
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
    "testing"
    "time"
)

func BenchmarkSelect(b *testing.B) {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        for i := 0; i < b.N; i++ {
            ch1 <- i
        }
    }()
    
    go func() {
        for i := 0; i < b.N; i++ {
            ch2 <- i
        }
    }()
    
    for i := 0; i < b.N; i++ {
        select {
        case <-ch1:
        case <-ch2:
        }
    }
}

func BenchmarkSelectWithDefault(b *testing.B) {
    ch := make(chan int)
    
    for i := 0; i < b.N; i++ {
        select {
        case <-ch:
        default:
        }
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
- Basic examples show select statement behavior
- Exercises demonstrate proper select usage
- Advanced patterns show production-ready code
- Race detection shows no race conditions
- Performance examples show select efficiency

### ⚠️ **Expected Warnings (These are GOOD for learning):**
- Timeout examples showing expected timeouts
- Performance differences between different select patterns

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

echo "🧪 Running Quick Select Test Suite"
echo "=================================="

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

echo "6. Testing static analysis..."
go vet . > /dev/null 2>&1 && echo "✅ Static analysis: PASS" || echo "❌ Static analysis: FAIL"

echo "=================================="
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
| `go run . basic` | Select operations, non-blocking, timeouts, priority |
| `go run . exercises` | Proper select usage patterns |
| `go run . advanced` | Production-ready select patterns |
| `go run -race . basic` | Race-free select operations |
| `go run -race . exercises` | Race-free synchronization |
| `go run -race . advanced` | Race-free advanced patterns |
| `go build .` | Code compiles correctly |
| `go vet .` | Code follows Go best practices |

## 🏆 Success Criteria

Your select statement topic is ready when:
- ✅ All commands run without errors
- ✅ Race detection shows no race conditions
- ✅ Exercises demonstrate proper patterns
- ✅ Advanced patterns show production quality
- ✅ Code compiles and passes static analysis

## 🚀 Ready for Next Topic?

Once all tests pass, you're ready to move to **Level 1, Topic 4: Synchronization Primitives**!

**Type "NEXT" to continue your journey to Go concurrency mastery!**
