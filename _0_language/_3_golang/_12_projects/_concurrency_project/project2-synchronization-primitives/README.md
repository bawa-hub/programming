# Project 2: Synchronization Primitives

## 🎯 Learning Objectives
- Master mutexes, RWMutex, and WaitGroup
- Learn atomic operations and sync.Once
- Practice with real-world synchronization scenarios
- Build a thread-safe cache system
- Understand performance implications of different primitives

## 📚 What You'll Build

### 1. Thread-Safe Cache System
A high-performance cache that can handle concurrent reads and writes safely.

### 2. Rate Limiter
A rate limiter that controls the number of requests per second.

### 3. Connection Pool
A pool of database connections with proper synchronization.

### 4. Atomic Counter
A high-performance counter using atomic operations.

## 🚀 Getting Started

```bash
# Run all examples
go run main.go

# Run specific components
go run main.go cache.go
go run main.go rate_limiter.go
go run main.go connection_pool.go
go run main.go atomic_counter.go

# Run tests
go test -v

# Run with race detection
go run -race main.go
```

## 📁 Project Structure

- `main.go` - Main entry point with all examples
- `cache.go` - Thread-safe cache implementation
- `rate_limiter.go` - Rate limiter with mutex and atomic operations
- `connection_pool.go` - Database connection pool
- `atomic_counter.go` - High-performance atomic counter
- `*_test.go` - Unit tests for each component

## 🎯 Key Concepts Covered

1. **Mutex vs RWMutex**: When to use each
2. **WaitGroup**: Coordinating goroutines
3. **Atomic Operations**: Lock-free programming
4. **sync.Once**: One-time initialization
5. **Performance Comparison**: Different approaches
6. **Race Condition Prevention**: Best practices

## 🏋️ Exercises

1. **Basic Mutex Usage**: Protect shared data
2. **RWMutex Optimization**: Optimize for read-heavy workloads
3. **WaitGroup Coordination**: Wait for multiple goroutines
4. **Atomic Operations**: High-performance counters
5. **sync.Once**: Singleton pattern implementation
6. **Performance Testing**: Benchmark different approaches

## 🎯 Success Criteria

After completing this project, you should be able to:
- ✅ Choose the right synchronization primitive for each scenario
- ✅ Implement thread-safe data structures
- ✅ Use atomic operations for high-performance code
- ✅ Prevent race conditions and deadlocks
- ✅ Optimize concurrent code for performance
- ✅ Debug synchronization issues

## 🚀 Next Steps

After mastering synchronization primitives, move on to:
- **Project 3**: Common Concurrency Patterns
- **Project 4**: Advanced Concurrency
- **Project 5**: Real-world Application

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!