# Channel Patterns & Idioms Commands

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
Runs 20 basic channel patterns and idioms:
- Basic Channel Ownership Pattern
- Channel Factory Pattern
- Channel Wrapper Pattern
- Graceful Shutdown Pattern
- Nil Channel Tricks
- Channel Switching Pattern
- Channel Pipeline Pattern
- Channel Fan-Out Pattern
- Channel Fan-In Pattern
- Error Channel Pattern
- Channel Batching Pattern
- Channel Rate Limiting Pattern
- Channel Generator Pattern
- Channel Transformer Pattern
- Channel Filter Pattern
- Channel Accumulator Pattern
- Channel Pool Pattern
- Channel Mock Pattern
- Channel Test Helper
- Channel Anti-Patterns

### 2. Exercises (`go run . exercises`)
Runs 10 hands-on exercises:
- Implement Channel Ownership Pattern
- Implement Channel Factory Pattern
- Implement Channel Wrapper Pattern
- Implement Graceful Shutdown Pattern
- Implement Nil Channel Tricks
- Implement Channel Pipeline Pattern
- Implement Channel Fan-Out Pattern
- Implement Channel Fan-In Pattern
- Implement Error Channel Pattern
- Implement Channel Batching Pattern

### 3. Advanced Patterns (`go run . advanced`)
Runs 12 advanced patterns:
- Channel-Based State Machine
- Event-Driven State Machine
- Channel Pool with Load Balancing
- Channel Rate Limiter
- Channel Circuit Breaker
- Channel Message Router
- Channel Priority Queue
- Channel Event Bus
- Channel Work Stealing
- Channel Metrics Collector
- Web Server with Channel Patterns
- Message Queue with Channel Patterns

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
6. All examples

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
go mod init channel-patterns
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
git commit -m "Add channel patterns implementation"
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
cd 11-channel-patterns
go run . basic
```

### Run All Tests
```bash
cd 11-channel-patterns
./quick_test.sh
```

### Run with Profiling
```bash
cd 11-channel-patterns
go run -cpuprofile=cpu.prof . all
go tool pprof cpu.prof
```

### Debug Channel Issues
```bash
cd 11-channel-patterns
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

### Channel Deadlocks
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

