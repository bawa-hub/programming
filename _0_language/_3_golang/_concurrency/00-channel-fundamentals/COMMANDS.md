# 🚀 Channel Fundamentals - Quick Commands

## 📚 Overview
This project provides a comprehensive understanding of Go channels without any complex patterns. Every concept is explained with clear examples and detailed comments.

## 🏃‍♂️ Quick Start

### Run All Examples
```bash
# Run everything
go run .

# Run specific concepts
go run . basic
go run . types
go run . operations
go run . behavior
go run . patterns
go run . pitfalls
```

### Individual Files
```bash
# Run specific files
go run 01-basic-concepts.go
go run 02-channel-types.go
go run 03-operations.go
go run 04-behavior.go
go run 05-patterns.go
go run 06-pitfalls.go
```

## 🔧 Development Commands

### Compilation
```bash
# Compile all files
go build .

# Compile specific file
go build 01-basic-concepts.go
```

### Testing
```bash
# Run with race detection
go run -race .

# Run with race detection on specific file
go run -race 01-basic-concepts.go
```

### Static Analysis
```bash
# Check for common mistakes
go vet .

# Check specific file
go vet 01-basic-concepts.go
```

### Formatting
```bash
# Format all files
go fmt .

# Format specific file
go fmt 01-basic-concepts.go
```

## 📖 Learning Path

### 1. Start Here
```bash
go run . basic
```
**Learn**: What are channels, synchronization, directionality, lifecycle

### 2. Channel Types
```bash
go run . types
```
**Learn**: Unbuffered vs buffered, capacity, nil channels

### 3. Operations
```bash
go run . operations
```
**Learn**: Send/receive, closing, range, non-blocking operations

### 4. Behavior
```bash
go run . behavior
```
**Learn**: Blocking/non-blocking, states, deadlocks, performance

### 5. Patterns
```bash
go run . patterns
```
**Learn**: Common patterns, worker, pipeline, fan-out/in, select

### 6. Pitfalls
```bash
go run . pitfalls
```
**Learn**: Common mistakes, how to avoid them, best practices

## 🎯 Key Concepts Covered

### Basic Concepts
- What are channels
- Channel synchronization
- Channel as signal
- Channel directionality
- Channel lifecycle

### Channel Types
- Unbuffered channels
- Buffered channels
- Channel capacity
- Nil channels

### Operations
- Sending data
- Receiving data
- Closing channels
- Range over channels
- Non-blocking operations

### Behavior
- Blocking behavior
- Non-blocking behavior
- Channel states
- Deadlocks
- Performance characteristics

### Patterns
- Send/receive pattern
- Worker pattern
- Pipeline pattern
- Fan-out pattern
- Fan-in pattern
- Timeout pattern
- Select pattern
- Signal pattern
- Quit pattern

### Pitfalls
- Deadlock pitfalls
- Closed channel pitfalls
- Buffer pitfalls
- Goroutine lifecycle pitfalls
- Select pitfalls
- Memory leak pitfalls
- Race condition pitfalls
- Performance pitfalls

## 🚨 Important Notes

### Race Detection
Always use race detection when testing:
```bash
go run -race .
```

### Static Analysis
Always check for common mistakes:
```bash
go vet .
```

### Best Practices
- Always check if channel is closed when receiving
- Use buffered channels when you want to decouple sender/receiver
- Use unbuffered channels when you need tight synchronization
- Close channels to signal completion
- Use select with default for non-blocking operations
- Use sync.WaitGroup to wait for goroutines
- Use context.Context for cancellation

## 🎉 Next Steps

After completing this project:
1. Practice with the examples
2. Try building your own channel-based programs
3. Use `go run -race` to test for race conditions
4. Use `go vet` to check for common mistakes
5. Explore the main concurrency curriculum

## 📚 File Structure

```
00-channel-fundamentals/
├── README.md              # Comprehensive guide
├── main.go                # Entry point with command handling
├── 01-basic-concepts.go   # Basic channel concepts
├── 02-channel-types.go    # Channel types and capacity
├── 03-operations.go       # Channel operations
├── 04-behavior.go         # Channel behavior
├── 05-patterns.go         # Common patterns
├── 06-pitfalls.go         # Common pitfalls
├── go.mod                 # Go module file
└── COMMANDS.md            # This file
```

## 🔍 Debugging Tips

### Common Issues
- **Deadlock**: Check for nil channels, circular dependencies
- **Panic**: Check for sending to closed channel, closing already closed channel
- **Race conditions**: Use `go run -race` to detect
- **Memory leaks**: Check for goroutine leaks, not closing channels

### Debugging Commands
```bash
# Check for race conditions
go run -race .

# Check for common mistakes
go vet .

# Profile memory usage
go run -memprofile=mem.prof .
go tool pprof mem.prof

# Profile CPU usage
go run -cpuprofile=cpu.prof .
go tool pprof cpu.prof
```

---

**Remember**: Channels are Go's way of communicating between goroutines. Think of them as pipes that allow data to flow from one goroutine to another safely and efficiently.

