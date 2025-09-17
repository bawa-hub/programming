# Lock-Free Programming Commands

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
Runs 20 basic lock-free programming examples:
- Basic Atomic Operations
- Atomic Value Operations
- Atomic Boolean Operations
- Compare and Swap Operations
- CAS Loops
- Memory Ordering
- Lock-Free Counter
- Lock-Free Stack
- Lock-Free Queue
- Lock-Free Ring Buffer
- Performance Comparison
- Atomic Pointer Operations
- Memory Pool
- Lock-Free Hash Table
- Work Stealing Queue
- Reference Counting
- ABA Problem Demonstration
- Lock-Free Allocator
- Concurrent Testing
- Stress Testing

### 2. Exercises (`go run . exercises`)
Runs 10 hands-on exercises:
- Implement Atomic Counter
- Implement Lock-Free Stack
- Implement Lock-Free Queue
- Implement Lock-Free Ring Buffer
- Implement Lock-Free Hash Table
- Implement Memory Pool
- Implement Reference Counting
- Implement Work Stealing Queue
- Implement Performance Comparison
- Implement Stress Testing

### 3. Advanced Patterns (`go run . advanced`)
Runs 10 advanced patterns:
- Lock-Free Skip List
- Lock-Free Hash Table
- Work Stealing Deque
- Lock-Free Memory Allocator
- Lock-Free Reference Counting
- Lock-Free Cache
- Lock-Free Priority Queue
- Lock-Free Bounded Queue
- Lock-Free Set
- Lock-Free Thread Pool

### 4. All Examples (`go run . all`)
Runs all examples and exercises in sequence:
- Basic examples
- Exercises
- Advanced patterns

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
7. Race detection

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

### Run with Race Detection
```bash
go run -race . basic
```

## Profiling Commands

### Memory Profiling
```bash
# Generate memory profile
go run -memprofile=mem.prof . basic

# Analyze memory profile
go tool pprof mem.prof

# Web interface
go tool pprof -http=:8080 mem.prof
```

### CPU Profiling
```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof . basic

# Analyze CPU profile
go tool pprof cpu.prof

# Web interface
go tool pprof -http=:8080 cpu.prof
```

## Debugging Commands

### Race Detection
```bash
# Run with race detection
go run -race . basic

# Run specific examples with race detection
go run -race . exercises
go run -race . advanced
```

### Verbose Output
```bash
# Run with verbose output
go run -v . basic

# Run with race detection and verbose output
go run -race -v . basic
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
go mod init lock-free
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
git commit -m "Add lock-free programming implementation"
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
cd 14-lock-free
go run . basic
```

### Run All Tests
```bash
cd 14-lock-free
./quick_test.sh
```

### Run with Profiling
```bash
cd 14-lock-free
go run -cpuprofile=cpu.prof . all
go tool pprof cpu.prof
```

### Debug Performance Issues
```bash
cd 14-lock-free
go run -race . basic
```

## Tips

1. **Always run tests before moving to next topic**
2. **Use race detection to verify lock-free correctness**
3. **Profile performance to identify bottlenecks**
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

### Race Conditions
```bash
go run -race . basic
```

### Memory Issues
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

## Next Steps

After running all commands successfully:
1. Review the output and understand the patterns
2. Experiment with different lock-free algorithms
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects

