# Quick Start Guide

## 🚀 Get Started in 5 Minutes

### 1. Prerequisites
```bash
# Check Go version (1.19+ required)
go version

# If not installed, download from https://golang.org/dl/
```

### 2. Initialize Your Learning Environment
```bash
# Navigate to your concurrency directory
cd /Users/vikramkumar/CS/_1_dsa/_0_language/_3_golang/_concurrency

# Initialize Go module
go mod init concurrency-learning

# Create your first project
cd project1-basic-goroutines
```

### 3. Run Your First Concurrent Program
```bash
# Run the basic examples
go run main.go

# Run the calculator
go run main.go calculator.go

# Run tests
go test -v

# Run with race detection
go run -race main.go calculator.go
```

### 4. Explore Examples
```bash
# Run basic operations examples
go run examples/basic_operations.go

# Run concurrent calculations
go run examples/concurrent_calculations.go

# Run channel patterns
go run examples/channel_patterns.go
```

## 📚 Learning Path

### Week 1-2: Project 1 - Basic Goroutines & Channels
- **Goal**: Understand goroutines and channels
- **Time**: 2-3 hours per day
- **Focus**: Hands-on coding, not just reading

### Week 3-4: Project 2 - Synchronization Primitives
- **Goal**: Master mutexes, waitgroups, atomic operations
- **Time**: 2-3 hours per day
- **Focus**: Thread safety and race conditions

### Week 5-6: Project 3 - Common Patterns
- **Goal**: Learn fan-in/fan-out, pipelines, generators
- **Time**: 2-3 hours per day
- **Focus**: Pattern recognition and application

### Week 7-8: Project 4 - Advanced Concurrency
- **Goal**: Worker pools, context, graceful shutdown
- **Time**: 2-3 hours per day
- **Focus**: Production-ready patterns

### Week 9-10: Project 5 - Real-world Application
- **Goal**: Build complete concurrent application
- **Time**: 3-4 hours per day
- **Focus**: End-to-end development

## 🛠️ Essential Commands

```bash
# Run with race detection (always use this!)
go run -race main.go

# Run tests with race detection
go test -race

# Run benchmarks
go test -bench=.

# Memory profiling
go test -memprofile=mem.prof

# CPU profiling
go test -cpuprofile=cpu.prof

# View profiles
go tool pprof mem.prof
go tool pprof cpu.prof
```

## 🎯 Success Tips

1. **Code Every Day**: Don't just read, write code
2. **Use Race Detection**: Always run with `-race` flag
3. **Write Tests**: Test your concurrent code thoroughly
4. **Benchmark**: Measure performance improvements
5. **Experiment**: Try different approaches and patterns
6. **Debug**: Learn to debug concurrent programs
7. **Read Code**: Study other concurrent Go projects

## 🆘 Common Issues

### Race Conditions
```bash
# Always run with race detection
go run -race main.go
```

### Deadlocks
- Use `go tool trace` to debug
- Check for circular channel dependencies
- Use timeouts in select statements

### Memory Leaks
- Close channels when done
- Use context for cancellation
- Monitor goroutine counts

## 📖 Additional Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)

## 🎉 Ready to Start?

```bash
cd project1-basic-goroutines
go run main.go
```

Happy coding! 🚀
