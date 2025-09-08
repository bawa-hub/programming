# Project 3: Common Concurrency Patterns

## 🎯 Learning Objectives
- Master fan-in and fan-out patterns
- Learn pipeline and generator patterns
- Practice with worker pools and job queues
- Build a real-time data processing system
- Understand graceful shutdown patterns

## 📚 What You'll Build

### 1. Fan-In/Fan-Out System
A system that distributes work across multiple workers and collects results.

### 2. Pipeline Processing
A multi-stage pipeline for data transformation.

### 3. Generator Pattern
A generator that produces data on-demand.

### 4. Worker Pool with Job Queue
A scalable worker pool that processes jobs from a queue.

### 5. Graceful Shutdown System
A system that can shut down gracefully while finishing current work.

## 🚀 Getting Started

```bash
# Run all examples
go run main.go

# Run specific components
go run main.go fan_in_out.go
go run main.go pipeline.go
go run main.go generator.go
go run main.go worker_pool.go
go run main.go graceful_shutdown.go

# Run tests
go test -v

# Run with race detection
go run -race main.go
```

## 📁 Project Structure

- `main.go` - Main entry point with all examples
- `fan_in_out.go` - Fan-in/fan-out pattern implementation
- `pipeline.go` - Pipeline processing implementation
- `generator.go` - Generator pattern implementation
- `worker_pool.go` - Worker pool with job queue
- `graceful_shutdown.go` - Graceful shutdown implementation
- `*_test.go` - Unit tests for each component

## 🎯 Key Concepts Covered

1. **Fan-In Pattern**: Collecting results from multiple sources
2. **Fan-Out Pattern**: Distributing work to multiple workers
3. **Pipeline Pattern**: Multi-stage data processing
4. **Generator Pattern**: On-demand data production
5. **Worker Pool Pattern**: Scalable concurrent processing
6. **Graceful Shutdown**: Clean termination of concurrent systems

## 🏋️ Exercises

1. **Fan-In/Fan-Out**: Distribute work and collect results
2. **Pipeline Processing**: Multi-stage data transformation
3. **Generator Pattern**: On-demand data production
4. **Worker Pool**: Scalable job processing
5. **Graceful Shutdown**: Clean system termination
6. **Performance Optimization**: Optimize concurrent patterns

## 🎯 Success Criteria

After completing this project, you should be able to:
- ✅ Implement fan-in and fan-out patterns
- ✅ Build multi-stage pipelines
- ✅ Create generators for on-demand data
- ✅ Design scalable worker pools
- ✅ Implement graceful shutdown
- ✅ Optimize concurrent patterns for performance

## 🚀 Next Steps

After mastering common patterns, move on to:
- **Project 4**: Advanced Concurrency
- **Project 5**: Real-world Application

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!