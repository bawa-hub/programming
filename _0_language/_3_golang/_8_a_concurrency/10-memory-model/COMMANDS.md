# Memory Model & Race Conditions Commands

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

### Development Commands
```bash
# Format code
go fmt .

# Check for unused imports
go mod tidy

# Run with verbose output
go run -v . basic

# Run with timeout
timeout 30s go run . basic
```

## Command Details

### 1. Basic Examples (`go run . basic`)
Runs 20 basic memory model examples:
- Basic Race Condition
- Race Condition with Maps
- Fixing Race Condition with Mutex
- Fixing Race Condition with Atomic
- Happens-Before with Channels
- Happens-Before with Mutex
- Atomic Operations and Memory Ordering
- Visibility Guarantees
- Race Detection Example
- Safe Counter with Mutex
- Safe Counter with Atomic
- Compare and Swap
- Atomic Flag
- Performance Comparison
- Memory Model Best Practices
- Lock-Free Data Structure
- False Sharing Example
- Memory Barrier Example
- Race Condition in Slice Operations
- Safe Slice Operations

### 2. Exercises (`go run . exercises`)
Runs 10 hands-on exercises:
- Fix Race Condition with Mutex
- Fix Race Condition with Atomic
- Implement Safe Map Operations
- Implement Atomic Flag
- Implement Compare and Swap
- Implement Happens-Before with Channels
- Implement Happens-Before with Mutex
- Implement Memory Ordering with Atomic
- Implement Safe Counter with Multiple Operations
- Implement Performance Comparison

### 3. Advanced Patterns (`go run . advanced`)
Runs 12 advanced patterns:
- Lock-Free Stack
- Memory Pool
- Double-Checked Locking
- Lock-Free Counter
- Lock-Free Ring Buffer
- False Sharing Prevention
- Memory Barrier with Atomic Operations
- Lock-Free Hash Table
- Lock-Free Queue
- Performance Comparison
- Memory Model Validation
- Lock-Free Producer-Consumer

### 4. All Examples (`go run . all`)
Runs all examples and exercises in sequence:
- Basic examples
- Exercises
- Advanced patterns

## Testing Commands

### Quick Test Script (`./quick_test.sh`)
Automated testing script that runs:
1. Basic compilation
2. Static analysis
3. Basic examples
4. Exercises
5. Advanced patterns
6. Race detection
7. All examples

### Static Analysis (`go vet .`)
Checks for:
- Unused variables
- Unreachable code
- Incorrect printf formats
- Other common issues

### Race Detection (`go run -race . basic`)
Detects:
- Data races
- Concurrent access issues
- Unsafe operations

## Performance Commands

### Memory Profiling
```bash
# Generate memory profile
go run -memprofile=mem.prof . basic

# Analyze memory usage
go tool pprof mem.prof

# Interactive mode
go tool pprof -http=:8080 mem.prof
```

### CPU Profiling
```bash
# Generate CPU profile
go run -cpuprofile=cpu.prof . basic

# Analyze CPU usage
go tool pprof cpu.prof

# Interactive mode
go tool pprof -http=:8080 cpu.prof
```

## Debug Commands

### Verbose Output
```bash
go run -v . basic
```
Shows detailed execution information

### Timeout Protection
```bash
timeout 30s go run . basic
```
Prevents hanging processes

### Build with Debug Info
```bash
go build -gcflags="-N -l" .
```
Disables optimizations for debugging

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
go mod init memory-model
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
git commit -m "Add memory model implementation"
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
cd 10-memory-model
go run . basic
```

### Run All Tests
```bash
cd 10-memory-model
./quick_test.sh
```

### Run with Profiling
```bash
cd 10-memory-model
go run -cpuprofile=cpu.prof . all
go tool pprof cpu.prof
```

### Debug Race Conditions
```bash
cd 10-memory-model
go run -race . exercises
```

## Tips

1. **Always run tests before moving to next topic**
2. **Use race detection to catch concurrency issues**
3. **Profile performance for optimization**
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
2. Experiment with different configurations
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects

