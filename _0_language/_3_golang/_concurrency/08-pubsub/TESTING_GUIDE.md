# Pub-Sub Pattern Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Pub-Sub Pattern implementation.

## Test Structure

### 1. Basic Compilation Test
```bash
go build .
```
- **Purpose**: Verify the code compiles without errors
- **Expected**: Clean compilation with no errors

### 2. Static Analysis Test
```bash
go vet .
```
- **Purpose**: Check for common programming errors
- **Expected**: No warnings or errors

### 3. Basic Examples Test
```bash
go run . basic
```
- **Purpose**: Verify basic pub-sub functionality
- **Expected**: All 10 basic examples run successfully

### 4. Exercises Test
```bash
go run . exercises
```
- **Purpose**: Verify hands-on exercises work correctly
- **Expected**: All 10 exercises complete successfully

### 5. Advanced Patterns Test
```bash
go run . advanced
```
- **Purpose**: Verify advanced patterns implementation
- **Expected**: All 10 advanced patterns run successfully

### 6. Race Detection Test
```bash
go run -race . basic
```
- **Purpose**: Detect data races in concurrent code
- **Expected**: No race conditions detected

### 7. All Examples Test
```bash
go run . all
```
- **Purpose**: Run all examples and exercises together
- **Expected**: Complete execution without errors

## Test Categories

### Basic Examples
1. **Basic Pub-Sub**: Simple publisher-subscriber pattern
2. **Multiple Subscribers**: Multiple subscribers per topic
3. **Topic-Based Filtering**: Message filtering by topic
4. **Content-Based Filtering**: Message filtering by content
5. **Error Handling**: Error handling and recovery
6. **Message Ordering**: Sequential message processing
7. **Performance Test**: Throughput and latency testing
8. **Wildcard Subscriptions**: Hierarchical topic subscriptions
9. **Message Persistence**: Message storage and replay
10. **Common Pitfalls**: Educational examples of mistakes

### Exercises
1. **Basic Implementation**: Create a simple pub-sub system
2. **Multiple Topics**: Handle multiple topics and subscribers
3. **Message Filtering**: Implement content-based filtering
4. **Error Handling**: Add retry logic and error recovery
5. **Message Ordering**: Ensure message sequence integrity
6. **Message Batching**: Implement message batching
7. **Message Persistence**: Add message storage
8. **Message Routing**: Route messages between topics
9. **Message Deduplication**: Prevent duplicate message processing
10. **Context and Cancellation**: Handle context cancellation

### Advanced Patterns
1. **Adaptive Scaling**: Scale based on load
2. **Circuit Breaker**: Prevent cascade failures
3. **Message Batching**: Optimize message delivery
4. **Message Compression**: Reduce message size
5. **Load Balancing**: Distribute load across brokers
6. **Message Routing**: Advanced routing strategies
7. **Dead Letter Queue**: Handle failed messages
8. **Metrics and Monitoring**: Track system performance
9. **Message Ordering**: Guarantee message sequence
10. **Message Deduplication**: Prevent duplicate processing

## Performance Expectations

### Basic Examples
- **Execution Time**: < 5 seconds
- **Memory Usage**: < 50MB
- **Throughput**: > 100 messages/second

### Exercises
- **Execution Time**: < 10 seconds
- **Memory Usage**: < 100MB
- **Success Rate**: 100% completion

### Advanced Patterns
- **Execution Time**: < 15 seconds
- **Memory Usage**: < 200MB
- **Error Rate**: < 5% (intentional failures)

## Troubleshooting

### Common Issues

#### 1. Compilation Errors
- **Symptom**: `go build .` fails
- **Cause**: Syntax errors or missing imports
- **Solution**: Check code syntax and import statements

#### 2. Race Conditions
- **Symptom**: `go run -race .` reports data races
- **Cause**: Concurrent access to shared variables
- **Solution**: Use proper synchronization primitives

#### 3. Deadlocks
- **Symptom**: Program hangs indefinitely
- **Cause**: Circular wait conditions
- **Solution**: Review channel operations and goroutine synchronization

#### 4. Memory Leaks
- **Symptom**: Memory usage grows continuously
- **Cause**: Goroutines not properly cleaned up
- **Solution**: Ensure proper context cancellation and cleanup

#### 5. Message Loss
- **Symptom**: Messages not delivered
- **Cause**: Unbuffered channels or blocking operations
- **Solution**: Use buffered channels and non-blocking operations

### Debug Commands

#### Verbose Output
```bash
go run -v . basic
```

#### Race Detection
```bash
go run -race . basic
```

#### Memory Profiling
```bash
go run -memprofile=mem.prof . basic
go tool pprof mem.prof
```

#### CPU Profiling
```bash
go run -cpuprofile=cpu.prof . basic
go tool pprof cpu.prof
```

## Test Automation

### Quick Test Script
```bash
./quick_test.sh
```
- Runs all tests automatically
- Provides clear pass/fail status
- Includes timeout protection

### Manual Testing
```bash
# Test individual components
go run . basic
go run . exercises
go run . advanced

# Test with different parameters
go run . all
```

## Success Criteria

### ✅ All Tests Pass
- Compilation successful
- Static analysis clean
- All examples run without errors
- No race conditions detected
- Performance within expected ranges

### ✅ Code Quality
- Clean, readable code
- Proper error handling
- Good documentation
- No memory leaks
- Efficient algorithms

### ✅ Learning Objectives
- Understand pub-sub patterns
- Implement basic functionality
- Handle advanced scenarios
- Apply best practices
- Avoid common pitfalls

## Next Steps

After successful testing:
1. Review the code and understand the patterns
2. Experiment with different configurations
3. Try implementing your own variations
4. Move to the next topic in the curriculum
5. Apply learnings to real-world projects

## Support

If you encounter issues:
1. Check the error messages carefully
2. Review the code for common mistakes
3. Use the debug commands provided
4. Refer to the Go concurrency documentation
5. Ask for help in the learning community

