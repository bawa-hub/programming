# Context Package Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Context Package implementation.

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
- **Purpose**: Verify basic context functionality
- **Expected**: All 15 basic examples run successfully

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
1. **Basic Context Creation**: Different context types and creation methods
2. **Context Cancellation**: Manual cancellation and propagation
3. **Context Timeout**: Timeout-based cancellation
4. **Context Deadline**: Deadline-based cancellation
5. **Context Values**: Storing and retrieving values
6. **Context Propagation**: Passing context through call chains
7. **HTTP Request with Context**: Using context in HTTP requests
8. **Database Operations with Context**: Context in database operations
9. **Multiple Goroutines with Context**: Coordinating goroutines with context
10. **Context Chain**: Chaining context operations
11. **Context Middleware**: Middleware pattern for context
12. **Context Performance**: Performance testing and optimization
13. **Context Error Handling**: Proper error handling with context
14. **Context Values with Types**: Type-safe context values
15. **Context Best Practices**: Demonstrating best practices

### Exercises
1. **Basic Context Implementation**: Create context with timeout
2. **Context with Values**: Multiple values and propagation
3. **Context Cancellation Chain**: Hierarchy and cancellation
4. **Context with Deadline**: Deadline handling
5. **Context in HTTP Handler**: HTTP request processing
6. **Context with Multiple Goroutines**: Coordinating goroutines
7. **Context Value Propagation**: Passing values through functions
8. **Context with Error Handling**: Retry logic and error handling
9. **Context with Database Operations**: Database integration
10. **Context Performance Testing**: Performance measurement

### Advanced Patterns
1. **Context Pool**: Pooling contexts for performance
2. **Context Middleware Chain**: Middleware pattern implementation
3. **Context Metrics**: Tracking context usage metrics
4. **Context Tracing**: Tracing context operations
5. **Context Validation**: Validating context state
6. **Context Cache**: Caching context values
7. **Context Rate Limiting**: Rate limiting context operations
8. **Context with Circuit Breaker**: Circuit breaker pattern
9. **Context with Retry Logic**: Retry mechanisms
10. **Context with Load Balancing**: Load balancing with context

## Performance Expectations

### Basic Examples
- **Execution Time**: < 10 seconds
- **Memory Usage**: < 100MB
- **Context Creation**: < 1ms per context
- **Value Lookup**: < 100ns per lookup

### Exercises
- **Execution Time**: < 15 seconds
- **Memory Usage**: < 150MB
- **Success Rate**: 100% completion
- **Error Handling**: Proper error propagation

### Advanced Patterns
- **Execution Time**: < 20 seconds
- **Memory Usage**: < 200MB
- **Performance**: Optimized for production use
- **Scalability**: Handles high concurrency

## Troubleshooting

### Common Issues

#### 1. Compilation Errors
- **Symptom**: `go build .` fails
- **Cause**: Syntax errors or missing imports
- **Solution**: Check code syntax and import statements

#### 2. Context Leaks
- **Symptom**: Memory usage grows continuously
- **Cause**: Not calling cancel() on context
- **Solution**: Always call cancel() in defer statements

#### 3. Race Conditions
- **Symptom**: `go run -race .` reports data races
- **Cause**: Concurrent access to shared variables
- **Solution**: Use proper synchronization primitives

#### 4. Context Timeout Issues
- **Symptom**: Contexts timeout unexpectedly
- **Cause**: Incorrect timeout values or blocking operations
- **Solution**: Check timeout values and operation duration

#### 5. Value Context Issues
- **Symptom**: Context values not found
- **Cause**: Wrong key types or context hierarchy
- **Solution**: Use custom key types and check context hierarchy

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
- Understand context package
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
4. Refer to the Go context documentation
5. Ask for help in the learning community

