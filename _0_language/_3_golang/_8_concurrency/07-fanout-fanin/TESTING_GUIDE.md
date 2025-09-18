# Fan-Out/Fan-In Pattern Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Fan-Out/Fan-In Pattern topic, covering basic examples, exercises, and advanced patterns.

## Test Structure

### 1. Basic Examples (`go run . basic`)
Tests fundamental fan-out/fan-in patterns:
- **Basic Fan-Out/Fan-In**: Simple parallel processing with multiple workers
- **Buffered Fan-Out/Fan-In**: Pipeline with buffered channels for better performance
- **Priority-Based Fan-Out/Fan-In**: Work distribution based on priority levels
- **Performance Comparison**: Sequential vs parallel processing comparison

### 2. Exercises (`go run . exercises`)
Hands-on exercises covering:
- **Exercise 1**: Basic Fan-Out/Fan-In Implementation
- **Exercise 2**: Buffered Fan-Out/Fan-In with Custom Buffer Sizes
- **Exercise 3**: Fan-Out/Fan-In with Dynamic Workers
- **Exercise 4**: Error Handling with Recovery
- **Exercise 5**: Timeout and Cancellation
- **Exercise 6**: Rate Limiting with Different Strategies
- **Exercise 7**: Metrics Collection and Analysis
- **Exercise 8**: Backpressure with Adaptive Buffering
- **Exercise 9**: Circuit Breaker Integration
- **Exercise 10**: Caching Layer Integration

### 3. Advanced Patterns (`go run . advanced`)
Advanced fan-out/fan-in implementations:
- **Adaptive Scaling**: Self-adjusting worker pool based on load
- **Circuit Breaker**: Fault-tolerant processing with automatic recovery
- **Caching Layer**: Performance optimization with intelligent caching
- **Load Balancing**: Intelligent work distribution across workers
- **Monitoring System**: Real-time monitoring and alerting

## Testing Commands

### Quick Test Suite
```bash
./quick_test.sh
```
Runs all tests including:
- Basic examples
- Exercises
- Advanced patterns
- Compilation
- Race detection
- Static analysis

### Individual Test Commands

#### Basic Examples
```bash
go run . basic
```

#### Exercises
```bash
go run . exercises
```

#### Advanced Patterns
```bash
go run . advanced
```

#### Compilation Test
```bash
go build .
```

#### Race Detection
```bash
go run -race . basic
```

#### Static Analysis
```bash
go vet .
```

#### Performance Testing
```bash
go test -bench=. -benchmem
```

## Expected Outputs

### Basic Examples
- Parallel processing results with worker distribution
- Error handling demonstrations
- Performance metrics and timing
- Priority-based processing order
- Sequential vs parallel performance comparison

### Exercises
- Successful completion of all 10 exercises
- Proper error handling and recovery
- Performance improvements over basic implementations
- Correct implementation of advanced patterns

### Advanced Patterns
- Adaptive behavior under different loads
- Circuit breaker state changes
- Cache hit/miss ratios
- Load balancing distribution
- Monitoring metrics and alerts

## Common Issues and Solutions

### 1. Deadlocks
**Symptom**: Program hangs indefinitely
**Solution**: Ensure all channels are properly closed and buffered appropriately

### 2. Race Conditions
**Symptom**: `go run -race` reports data races
**Solution**: Use proper synchronization primitives or make variables local to goroutines

### 3. Channel Leaks
**Symptom**: Goroutines not terminating
**Solution**: Always close channels and use `defer close(channel)`

### 4. Worker Pool Management
**Symptom**: Workers not processing work or hanging
**Solution**: Ensure proper work distribution and channel management

### 5. Load Balancing
**Symptom**: Uneven work distribution
**Solution**: Implement proper load balancing algorithms

## Performance Expectations

### Basic Fan-Out/Fan-In
- Should process 20 items with 4 workers
- Processing time should be distributed across workers
- Error rate should be around 10%

### Buffered Fan-Out/Fan-In
- Should process 30 items with 6 workers
- Should show improved throughput over unbuffered
- Buffer size should match processing capacity

### Priority-Based Fan-Out/Fan-In
- Should process 25 items with priority-based distribution
- High priority items should be processed first
- Should maintain worker load balancing

### Performance Comparison
- Should show performance comparison between sequential and parallel processing
- Parallel processing should show speedup for CPU-bound work
- Efficiency should be calculated based on worker count

## Debugging Tips

### 1. Add Logging
```go
fmt.Printf("Worker %d processing item %d\n", workerID, item.ID)
```

### 2. Monitor Worker Activity
```go
fmt.Printf("Worker %d processed %d items\n", workerID, processedCount)
```

### 3. Check Channel States
```go
select {
case item := <-input:
    // Process item
default:
    fmt.Println("Input channel is empty")
}
```

### 4. Use Timeouts
```go
select {
case result := <-output:
    // Process result
case <-time.After(5 * time.Second):
    fmt.Println("Timeout waiting for result")
}
```

### 5. Monitor Goroutines
```go
fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
```

## Success Criteria

✅ All basic examples run without errors
✅ All exercises complete successfully
✅ All advanced patterns demonstrate expected behavior
✅ Code compiles without warnings
✅ Static analysis passes
✅ Race detection shows no race conditions
✅ Performance meets expected benchmarks
✅ Error handling works correctly
✅ Documentation is complete and accurate

## Next Steps

After completing all tests successfully:
1. Review the code to understand the patterns
2. Experiment with different worker counts and buffer sizes
3. Try implementing your own fan-out/fan-in variations
4. Move on to the next topic: Pub-Sub Pattern

## Troubleshooting

If tests fail:
1. Check the error messages carefully
2. Ensure all dependencies are installed
3. Verify Go version compatibility
4. Check for syntax errors
5. Review the implementation against the requirements

For additional help, refer to the main README.md or the individual exercise comments.

