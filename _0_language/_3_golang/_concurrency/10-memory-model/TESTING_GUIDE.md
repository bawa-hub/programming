# Memory Model & Race Conditions Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Memory Model & Race Conditions implementation.

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
- **Purpose**: Verify basic memory model functionality
- **Expected**: All 20 basic examples run successfully

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
- **Expected**: All 12 advanced patterns run successfully

### 6. Race Detection Test
```bash
go run -race . basic
```
- **Purpose**: Detect data races in concurrent code
- **Expected**: May find intentional race conditions for educational purposes

### 7. All Examples Test
```bash
go run . all
```
- **Purpose**: Run all examples and exercises together
- **Expected**: Complete execution without errors

## Test Categories

### Basic Examples
1. **Basic Race Condition**: Demonstrates race conditions
2. **Race Condition with Maps**: Shows map race conditions
3. **Fixing Race Condition with Mutex**: Mutex solution
4. **Fixing Race Condition with Atomic**: Atomic solution
5. **Happens-Before with Channels**: Channel synchronization
6. **Happens-Before with Mutex**: Mutex synchronization
7. **Atomic Operations and Memory Ordering**: Atomic memory ordering
8. **Visibility Guarantees**: Memory visibility
9. **Race Detection Example**: Race detector demonstration
10. **Safe Counter with Mutex**: Mutex-based counter
11. **Safe Counter with Atomic**: Atomic-based counter
12. **Compare and Swap**: CAS operations
13. **Atomic Flag**: Atomic flag operations
14. **Performance Comparison**: Mutex vs atomic performance
15. **Memory Model Best Practices**: Best practices demonstration
16. **Lock-Free Data Structure**: Lock-free counter
17. **False Sharing Example**: False sharing demonstration
18. **Memory Barrier Example**: Memory barriers
19. **Race Condition in Slice Operations**: Slice race conditions
20. **Safe Slice Operations**: Safe slice handling

### Exercises
1. **Fix Race Condition with Mutex**: Implement mutex solution
2. **Fix Race Condition with Atomic**: Implement atomic solution
3. **Implement Safe Map Operations**: Safe map handling
4. **Implement Atomic Flag**: Atomic flag implementation
5. **Implement Compare and Swap**: CAS implementation
6. **Implement Happens-Before with Channels**: Channel synchronization
7. **Implement Happens-Before with Mutex**: Mutex synchronization
8. **Implement Memory Ordering with Atomic**: Atomic memory ordering
9. **Implement Safe Counter with Multiple Operations**: Multi-operation counter
10. **Implement Performance Comparison**: Performance testing

### Advanced Patterns
1. **Lock-Free Stack**: Lock-free stack implementation
2. **Memory Pool**: Memory pool for performance
3. **Double-Checked Locking**: Singleton pattern
4. **Lock-Free Counter**: Lock-free counter
5. **Lock-Free Ring Buffer**: Lock-free ring buffer
6. **False Sharing Prevention**: False sharing avoidance
7. **Memory Barrier with Atomic Operations**: Memory barriers
8. **Lock-Free Hash Table**: Lock-free hash table
9. **Lock-Free Queue**: Lock-free queue
10. **Performance Comparison**: Comprehensive performance testing
11. **Memory Model Validation**: Memory model validation
12. **Lock-Free Producer-Consumer**: Producer-consumer pattern

## Performance Expectations

### Basic Examples
- **Execution Time**: < 15 seconds
- **Memory Usage**: < 200MB
- **Race Detection**: May find intentional races
- **Performance**: Demonstrates mutex vs atomic differences

### Exercises
- **Execution Time**: < 20 seconds
- **Memory Usage**: < 250MB
- **Success Rate**: 100% completion
- **Race-Free**: All exercises should be race-free

### Advanced Patterns
- **Execution Time**: < 25 seconds
- **Memory Usage**: < 300MB
- **Performance**: Optimized lock-free implementations
- **Scalability**: Handles high concurrency

## Troubleshooting

### Common Issues

#### 1. Compilation Errors
- **Symptom**: `go build .` fails
- **Cause**: Syntax errors or missing imports
- **Solution**: Check code syntax and import statements

#### 2. Race Conditions
- **Symptom**: `go run -race .` reports data races
- **Cause**: Intentional race conditions for educational purposes
- **Solution**: Some races are intentional; focus on understanding the concepts

#### 3. Panic in Map Operations
- **Symptom**: "concurrent map read and map write" panic
- **Cause**: Race condition in map operations
- **Solution**: Use sync.Map or mutex protection

#### 4. Slice Race Conditions
- **Symptom**: Panic in slice operations
- **Cause**: Race condition in slice operations
- **Solution**: Use channels or mutex protection

#### 5. False Sharing
- **Symptom**: Poor performance with atomic operations
- **Cause**: False sharing between variables
- **Solution**: Add padding between variables

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
- Includes race detection

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
- Race detection may find intentional races
- Performance within expected ranges

### ✅ Code Quality
- Clean, readable code
- Proper error handling
- Good documentation
- Race-free implementations
- Efficient algorithms

### ✅ Learning Objectives
- Understand memory model concepts
- Implement race-free code
- Handle advanced patterns
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
4. Refer to the Go memory model documentation
5. Ask for help in the learning community
