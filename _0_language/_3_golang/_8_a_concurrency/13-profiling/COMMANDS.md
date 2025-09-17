# Profiling & Benchmarking Commands

## Quick Reference

### Basic Commands
```bash
# Compile the code
go build .

# Run basic examples
go run . basic

# Run exercises
go run . exercises

# Run advanced patterns
go run . advanced

# Run all examples
go run . all
```

### Profiling Commands
```bash
# CPU profiling
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof

# Memory profiling
go run -memprofile=mem.prof . basic
go tool pprof mem.prof

# Goroutine profiling
go run -goroutineprofile=goroutine.prof . basic
go tool pprof goroutine.prof

# Block profiling
go run -blockprofile=block.prof . basic
go tool pprof block.prof

# Mutex profiling
go run -mutexprofile=mutex.prof . basic
go tool pprof mutex.prof
```

### Benchmark Commands
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run with memory info
go test -bench=. -benchmem

# Run multiple times
go test -bench=. -count=5
```

### Testing Commands
```bash
# Run all tests
./quick_test.sh

# Static analysis
go vet .

# Race detection
go run -race . basic

# Memory profiling
go run -memprofile=mem.prof . basic
go tool pprof mem.prof

# CPU profiling
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

## Command Details

### 1. Basic Examples (`go run . basic`)
Runs 20 basic profiling and benchmarking examples:
- Basic CPU Profiling
- Basic Memory Profiling
- Basic Goroutine Profiling
- Basic Block Profiling
- Basic Mutex Profiling
- Basic Benchmarking
- Memory Allocation Analysis
- Goroutine Analysis
- CPU Usage Analysis
- Memory Leak Detection
- Performance Comparison
- Profiling with HTTP Server
- Custom Profiling
- Memory Pool Usage
- Goroutine Pool
- Channel Performance
- Select Performance
- Mutex vs Channel Performance
- Memory Efficiency
- Performance Monitoring

### 2. Exercises (`go run . exercises`)
Runs 10 hands-on exercises:
- Implement CPU Profiling
- Implement Memory Profiling
- Implement Goroutine Profiling
- Implement Block Profiling
- Implement Mutex Profiling
- Implement Benchmarking
- Implement Memory Analysis
- Implement Performance Comparison
- Implement Memory Pool
- Implement Goroutine Pool

### 3. Advanced Patterns (`go run . advanced`)
Runs 10 advanced patterns:
- Real-time Performance Monitoring
- Profiling with Context
- Custom Profiling
- Performance Profiler
- Memory Profiler
- Goroutine Profiler
- Performance Dashboard
- Performance Optimizer
- Memory Optimizer
- Performance Tester

### 4. All Examples (`go run . all`)
Runs all examples and exercises in sequence:
- Basic examples
- Exercises
- Advanced patterns

## Profiling Commands

### CPU Profiling
```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof . basic

# Analyze CPU profile
go tool pprof cpu.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by CPU usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph

# Web interface
go tool pprof -http=:8080 cpu.prof
```

### Memory Profiling
```bash
# Generate memory profile
go run -memprofile=mem.prof . basic

# Analyze memory profile
go tool pprof mem.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 functions by memory usage
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
# (pprof) png            # Generate PNG graph

# Web interface
go tool pprof -http=:8080 mem.prof
```

### Goroutine Profiling
```bash
# Generate goroutine profile
go run -goroutineprofile=goroutine.prof . basic

# Analyze goroutine profile
go tool pprof goroutine.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 goroutines
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

### Block Profiling
```bash
# Generate block profile
go run -blockprofile=block.prof . basic

# Analyze block profile
go tool pprof block.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 blocking operations
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

### Mutex Profiling
```bash
# Generate mutex profile
go run -mutexprofile=mutex.prof . basic

# Analyze mutex profile
go tool pprof mutex.prof

# Interactive commands in pprof:
# (pprof) top10          # Top 10 mutex contentions
# (pprof) list function  # Show source code for function
# (pprof) web            # Open web interface
```

## Benchmark Commands

### Run Benchmarks
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFunction

# Run with memory info
go test -bench=. -benchmem

# Run multiple times
go test -bench=. -count=5

# Run with profiling
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
```

### Benchmark Analysis
```bash
# Compare benchmarks
go test -bench=. -benchmem > before.txt
# Make changes
go test -bench=. -benchmem > after.txt
# Compare results
```

## Testing Commands

### Quick Test Script
```bash
./quick_test.sh
```
Runs all tests automatically:
1. Basic compilation
2. Static analysis
3. Basic examples
4. Exercises
5. Advanced patterns
6. All examples
7. Benchmark tests

### Manual Testing
```bash
# Test individual components
go run . basic
go run . exercises
go run . advanced

# Test with different parameters
go run . all
```

## Development Commands

### Format Code
```bash
go fmt .
```

### Check for Unused Imports
```bash
go mod tidy
```

### Run with Verbose Output
```bash
go run -v . basic
```

### Run with Timeout
```bash
timeout 30s go run . basic
```

## File Structure Commands

### List Files
```bash
ls -la
```

### View File Contents
```bash
cat main.go
cat exercises.go
cat advanced_patterns.go
```

### Check File Permissions
```bash
ls -la *.sh
```

## Module Commands

### Initialize Module
```bash
go mod init profiling
```

### Add Dependencies
```bash
go get <package>
```

### Update Dependencies
```bash
go get -u <package>
```

### Clean Module
```bash
go mod tidy
```

## Git Commands (if using version control)

### Check Status
```bash
git status
```

### Add Files
```bash
git add .
```

### Commit Changes
```bash
git commit -m "Add profiling implementation"
```

### View History
```bash
git log --oneline
```

## Troubleshooting Commands

### Check Go Version
```bash
go version
```

### Check Module Status
```bash
go mod verify
```

### Clean Build Cache
```bash
go clean -cache
```

### Check for Updates
```bash
go list -u -m all
```

## Example Usage

### Run Basic Examples
```bash
cd 13-profiling
go run . basic
```

### Run All Tests
```bash
cd 13-profiling
./quick_test.sh
```

### Run with Profiling
```bash
cd 13-profiling
go run -cpuprofile=cpu.prof . all
go tool pprof cpu.prof
```

### Debug Performance Issues
```bash
cd 13-profiling
go run -race . basic
```

## Tips

1. **Always run tests before moving to next topic**
2. **Use profiling to identify performance bottlenecks**
3. **Use benchmarks to measure performance improvements**
4. **Check static analysis for code quality**
5. **Use timeouts to prevent hanging processes**

## Common Issues

### Permission Denied
```bash
chmod +x quick_test.sh
```

### Module Not Found
```bash
go mod tidy
```

### Profile Generation Issues
```bash
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

### Memory Issues
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

## Next Steps

After running all commands successfully:
1. Review the output and understand the patterns
2. Experiment with different profiling scenarios
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects

