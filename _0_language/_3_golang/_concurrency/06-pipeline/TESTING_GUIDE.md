# Pipeline Pattern Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Pipeline Pattern topic, covering basic examples, exercises, and advanced patterns.

## Test Structure

### 1. Basic Examples (`go run . basic`)
Tests fundamental pipeline patterns:
- **Basic Pipeline**: Simple 3-stage pipeline
- **Buffered Pipeline**: Pipeline with buffered channels
- **Fan-Out/Fan-In**: Parallel processing with worker distribution
- **Error Handling**: Pipeline with error propagation
- **Timeout**: Pipeline with timeout controls
- **Rate Limiting**: Pipeline with rate limiting
- **Metrics**: Pipeline with performance monitoring
- **Backpressure**: Pipeline with backpressure handling
- **Performance Comparison**: Sequential vs pipeline performance
- **Common Pitfalls**: Educational examples of mistakes

### 2. Exercises (`go run . exercises`)
Hands-on exercises covering:
- **Exercise 1**: Basic Pipeline Implementation
- **Exercise 2**: Buffered Pipeline with Custom Buffer Sizes
- **Exercise 3**: Fan-Out/Fan-In with Dynamic Workers
- **Exercise 4**: Error Handling with Recovery
- **Exercise 5**: Timeout and Cancellation
- **Exercise 6**: Rate Limiting with Different Strategies
- **Exercise 7**: Metrics Collection and Analysis
- **Exercise 8**: Backpressure with Adaptive Buffering
- **Exercise 9**: Circuit Breaker Integration
- **Exercise 10**: Caching Layer Integration

### 3. Advanced Patterns (`go run . advanced`)
Advanced pipeline implementations:
- **Adaptive Pipeline**: Self-adjusting pipeline based on load
- **Pipeline with Circuit Breaker**: Fault-tolerant pipeline
- **Pipeline with Caching**: Performance-optimized pipeline
- **Pipeline with Load Balancing**: Distributed processing
- **Pipeline with Monitoring**: Real-time monitoring and alerting

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
- Race detection (with educational race conditions)
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
**Note**: The error handling example intentionally contains race conditions for educational purposes.

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
- Pipeline processing results with stage information
- Performance metrics and timing
- Error handling demonstrations
- Rate limiting behavior
- Backpressure handling

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
- Monitoring metrics

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

### 4. Buffer Sizing
**Symptom**: Poor performance or blocking
**Solution**: Size buffers based on expected throughput and processing time

### 5. Error Handling
**Symptom**: Panics or unhandled errors
**Solution**: Implement proper error propagation and recovery mechanisms

## Performance Expectations

### Basic Pipeline
- Should process 10 items in ~1.5 seconds
- 3-stage pipeline with 50ms per stage

### Buffered Pipeline
- Should show improved throughput over unbuffered
- Buffer size should match processing capacity

### Fan-Out/Fan-In
- Should utilize multiple workers effectively
- Load should be distributed evenly

### Error Handling
- Should gracefully handle errors without crashing
- Should continue processing successful items

### Rate Limiting
- Should respect rate limits (e.g., 2 items per second)
- Should show controlled processing speed

## Debugging Tips

### 1. Add Logging
```go
fmt.Printf("Stage %s processing item %d\n", stageName, item.ID)
```

### 2. Use Timeouts
```go
select {
case result := <-output:
    // Process result
case <-time.After(5 * time.Second):
    fmt.Println("Timeout waiting for result")
}
```

### 3. Monitor Goroutines
```go
fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
```

### 4. Check Channel States
```go
select {
case data := <-channel:
    // Process data
default:
    fmt.Println("Channel is empty or closed")
}
```

## Success Criteria

✅ All basic examples run without errors
✅ All exercises complete successfully
✅ All advanced patterns demonstrate expected behavior
✅ Code compiles without warnings
✅ Static analysis passes
✅ Race detection shows only educational race conditions
✅ Performance meets expected benchmarks
✅ Error handling works correctly
✅ Documentation is complete and accurate

## Next Steps

After completing all tests successfully:
1. Review the code to understand the patterns
2. Experiment with different configurations
3. Try implementing your own pipeline variations
4. Move on to the next topic: Fan-Out/Fan-In Pattern

## Troubleshooting

If tests fail:
1. Check the error messages carefully
2. Ensure all dependencies are installed
3. Verify Go version compatibility
4. Check for syntax errors
5. Review the implementation against the requirements

For additional help, refer to the main README.md or the individual exercise comments.

