# 🚀 Advanced Scheduling Commands

## Quick Reference Commands

### Basic Commands
```bash
# Run all tests
./quick_test.sh

# Run specific examples
go run . basic      # Basic scheduling examples
go run . exercises  # Hands-on exercises
go run . advanced   # Advanced patterns
go run . all        # Everything

# Compilation and analysis
go build .          # Compile
go vet .            # Static analysis
go run -race . basic # Race detection
```

### Development Commands
```bash
# Initialize module
go mod init scheduling-examples

# Add dependencies
go get <package>

# Clean build
go clean
go build .

# Run with profiling
go run . basic
go tool pprof -http=:8080 mem.prof
```

### Testing Commands
```bash
# Run tests
go test .

# Run benchmarks
go test -bench=.

# Run with coverage
go test -cover .

# Run with race detection
go test -race .
```

### Debugging Commands
```bash
# Enable scheduler tracing
GODEBUG=schedtrace=1000 go run . basic

# Enable scheduler detail
GODEBUG=scheddetail=1 go run . basic

# Show goroutine stack traces
go run . basic
go tool pprof -http=:8080 goroutine.prof

# Monitor memory usage
go run . basic
go tool pprof -http=:8080 mem.prof
```

### Performance Commands
```bash
# CPU profiling
go run . basic
go tool pprof -http=:8080 cpu.prof

# Memory profiling
go run . basic
go tool pprof -http=:8080 mem.prof

# Goroutine profiling
go run . basic
go tool pprof -http=:8080 goroutine.prof

# Block profiling
go run . basic
go tool pprof -http=:8080 block.prof
```

### Monitoring Commands
```bash
# Monitor goroutine count
watch -n 1 'go run . basic | grep goroutines'

# Monitor memory usage
watch -n 1 'go run . basic | grep memory'

# Monitor performance
watch -n 1 'go run . basic | grep Performance'
```

### Advanced Commands
```bash
# Run with specific GOMAXPROCS
GOMAXPROCS=1 go run . basic
GOMAXPROCS=2 go run . basic
GOMAXPROCS=4 go run . basic

# Run with memory limit
go run . basic
ulimit -v 1000000

# Run with CPU limit
go run . basic
taskset -c 0,1 go run . basic
```

### Cleanup Commands
```bash
# Clean build artifacts
go clean

# Clean module cache
go clean -modcache

# Remove test files
rm -f *.prof
rm -f scheduling_test
```

### Help Commands
```bash
# Show help
go run . help

# Show available commands
go run . --help

# Show version
go version

# Show environment
go env
```

---

## 🎯 Command Categories

### 1. Basic Testing
- `./quick_test.sh` - Run all tests
- `go run . basic` - Basic examples
- `go run . exercises` - Hands-on exercises
- `go run . advanced` - Advanced patterns

### 2. Compilation & Analysis
- `go build .` - Compile
- `go vet .` - Static analysis
- `go run -race . basic` - Race detection

### 3. Performance Testing
- `go run . basic | grep Performance` - Performance metrics
- `go tool pprof -http=:8080 cpu.prof` - CPU profiling
- `go tool pprof -http=:8080 mem.prof` - Memory profiling

### 4. Debugging
- `GODEBUG=schedtrace=1000 go run . basic` - Scheduler tracing
- `GODEBUG=scheddetail=1 go run . basic` - Scheduler detail
- `go tool pprof -http=:8080 goroutine.prof` - Goroutine profiling

### 5. Monitoring
- `watch -n 1 'go run . basic | grep goroutines'` - Goroutine count
- `watch -n 1 'go run . basic | grep memory'` - Memory usage
- `watch -n 1 'go run . basic | grep Performance'` - Performance

---

## 🔧 Environment Variables

### Go Runtime
```bash
# Set GOMAXPROCS
export GOMAXPROCS=4

# Enable scheduler tracing
export GODEBUG=schedtrace=1000

# Enable scheduler detail
export GODEBUG=scheddetail=1

# Enable memory profiling
export GODEBUG=memprofilerate=1
```

### System Limits
```bash
# Set memory limit
ulimit -v 1000000

# Set CPU limit
taskset -c 0,1

# Set file descriptor limit
ulimit -n 10000
```

---

## 📊 Output Examples

### Basic Examples Output
```
⚙️ Advanced Scheduling Examples
===============================

1. Understanding Scheduler Components
====================================
  GOMAXPROCS: 4
  Number of goroutines: 1
  Number of CPUs: 4
  Number of CGO calls: 0
  Scheduler components completed
```

### Exercises Output
```
💪 Advanced Scheduling Exercises
===============================

Exercise 1: Implement Basic Work Stealing
=======================================
  Exercise 1: Pushed 0: true
  Exercise 1: Pushed 1: true
  Exercise 1: Popped: 1
  Exercise 1: Basic work stealing completed
```

### Advanced Patterns Output
```
🚀 Advanced Scheduling Patterns
==============================

1. Advanced Work Stealing Queue
  Popped: 0
  Popped: 1
  Popped: 2
  Popped: 3
  Popped: 4
```

---

## 🎉 Success Indicators

### ✅ All Tests Pass
```
🎉 All tests passed! Advanced Scheduling is ready!
Ready to move to the next topic!
```

### ✅ Performance Metrics
```
  Small work, many workers: 1.234ms
  Medium work, balanced: 5.678ms
  Large work, few workers: 12.345ms
```

### ✅ Scheduler Statistics
```
  GOMAXPROCS: 4
  Number of goroutines: 1
  Number of CPUs: 4
  GC cycles: 0
```

---

## 🚀 Ready to Test!

Use these commands to test your Advanced Scheduling knowledge:

1. **Start with basic examples**: `go run . basic`
2. **Run exercises**: `go run . exercises`
3. **Test advanced patterns**: `go run . advanced`
4. **Run all tests**: `./quick_test.sh`
5. **Check performance**: `go run . all | grep Performance`

Happy testing! 💪

