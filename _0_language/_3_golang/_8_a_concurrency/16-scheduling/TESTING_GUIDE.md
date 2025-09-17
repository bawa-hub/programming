# 🧪 Advanced Scheduling Testing Guide

## Overview
This guide provides comprehensive testing instructions for the Advanced Scheduling topic, covering Go scheduler internals, work stealing, preemption, and scheduler-aware programming.

---

## 🚀 Quick Start

### Automated Testing
```bash
# Run all tests
./quick_test.sh

# Run specific test categories
go run . basic      # Basic scheduling examples
go run . exercises  # Hands-on exercises
go run . advanced   # Advanced patterns
go run . all        # Everything
```

---

## 📋 Test Categories

### 1. Basic Examples (20 examples)
- **Scheduler Components**: Understanding G, M, P model
- **Goroutine Lifecycle**: Creation, execution, termination
- **Scheduler Statistics**: Runtime metrics and monitoring
- **Work Stealing**: Basic work distribution algorithms
- **Preemption**: Cooperative and forced preemption
- **GOMAXPROCS**: Processor configuration and tuning
- **CPU Affinity**: Processor binding simulation
- **Performance Tuning**: Optimization strategies

### 2. Hands-on Exercises (10 exercises)
- **Exercise 1**: Implement Basic Work Stealing
- **Exercise 2**: Implement GOMAXPROCS Tuning
- **Exercise 3**: Implement Cooperative Yielding
- **Exercise 4**: Implement Work Distribution
- **Exercise 5**: Implement Scheduler Statistics
- **Exercise 6**: Implement Context Switching Analysis
- **Exercise 7**: Implement Work Stealing Performance
- **Exercise 8**: Implement CPU Affinity Simulation
- **Exercise 9**: Implement Scheduler Contention Avoidance
- **Exercise 10**: Implement Performance Optimization

### 3. Advanced Patterns (12 patterns)
- **Advanced Work Stealing Queue**: Lock-free work distribution
- **Scheduler-Aware Worker Pool**: Optimized worker management
- **CPU-Aware Load Balancer**: Intelligent load distribution
- **Scheduler Statistics Monitor**: Real-time monitoring
- **Adaptive Scheduler**: Dynamic scaling
- **Work Stealing with Priority**: Priority-based work distribution
- **Scheduler-Aware Rate Limiter**: Intelligent rate limiting
- **Scheduler-Aware Circuit Breaker**: Fault tolerance
- **Scheduler-Aware Metrics Collector**: Performance metrics
- **Scheduler-Aware Event Bus**: Event-driven architecture
- **Scheduler-Aware Web Server**: High-performance web server
- **Scheduler-Aware Message Queue**: Message distribution

---

## 🔧 Testing Commands

### Compilation Tests
```bash
# Basic compilation
go build .

# Static analysis
go vet .

# Race detection
go run -race . basic
```

### Performance Tests
```bash
# Benchmark basic examples
go run . basic | grep "Performance"

# Benchmark exercises
go run . exercises | grep "Performance"

# Benchmark advanced patterns
go run . advanced | grep "Performance"
```

### Memory Tests
```bash
# Memory profiling
go run . basic
go tool pprof -http=:8080 mem.prof

# Goroutine profiling
go run . basic
go tool pprof -http=:8080 goroutine.prof
```

---

## 📊 Expected Results

### Basic Examples
- All 20 examples should run without errors
- Scheduler statistics should be displayed
- Work stealing should demonstrate load balancing
- GOMAXPROCS changes should show performance impact

### Exercises
- All 10 exercises should complete successfully
- Work stealing queue should handle concurrent access
- GOMAXPROCS tuning should show performance differences
- Context switching analysis should show overhead patterns

### Advanced Patterns
- All 12 patterns should run without errors
- Work stealing should demonstrate priority handling
- Adaptive scheduler should show dynamic scaling
- Circuit breaker should demonstrate fault tolerance

---

## 🐛 Troubleshooting

### Common Issues

#### 1. Compilation Errors
```bash
# Check Go version
go version

# Clean module cache
go clean -modcache

# Rebuild
go build .
```

#### 2. Race Detection Failures
```bash
# Run with more detailed output
go run -race . basic 2>&1 | grep -i race

# Check for data races in specific functions
go run -race . exercises 2>&1 | grep -i race
```

#### 3. Performance Issues
```bash
# Check GOMAXPROCS
go run . basic | grep GOMAXPROCS

# Monitor goroutine count
go run . basic | grep "goroutines"
```

#### 4. Memory Issues
```bash
# Check memory usage
go run . basic
go tool pprof -http=:8080 mem.prof

# Check for memory leaks
go run . advanced
go tool pprof -http=:8080 mem.prof
```

---

## 📈 Performance Benchmarks

### Expected Performance Ranges

#### Work Stealing Performance
- **Small work (10 items)**: < 1ms
- **Medium work (100 items)**: < 10ms
- **Large work (1000 items)**: < 100ms

#### Context Switching Overhead
- **1 goroutine**: < 1ms
- **10 goroutines**: < 5ms
- **100 goroutines**: < 50ms
- **1000 goroutines**: < 500ms

#### GOMAXPROCS Impact
- **1 processor**: Baseline
- **2 processors**: 1.5-2x improvement
- **4 processors**: 2-4x improvement
- **8 processors**: 2-8x improvement (diminishing returns)

---

## 🔍 Debugging Tips

### 1. Scheduler Debugging
```bash
# Enable scheduler tracing
GODEBUG=schedtrace=1000 go run . basic

# Enable scheduler detail
GODEBUG=scheddetail=1 go run . basic
```

### 2. Goroutine Debugging
```bash
# Show goroutine stack traces
go run . basic
go tool pprof -http=:8080 goroutine.prof

# Monitor goroutine count
watch -n 1 'go run . basic | grep goroutines'
```

### 3. Memory Debugging
```bash
# Show memory allocation
go run . basic
go tool pprof -http=:8080 mem.prof

# Check for memory leaks
go run . advanced
go tool pprof -http=:8080 mem.prof
```

---

## 📚 Learning Objectives

After completing all tests, you should understand:

1. **Go Scheduler Internals**: How G, M, P model works
2. **Work Stealing**: Efficient work distribution algorithms
3. **Preemption**: Cooperative and forced preemption mechanisms
4. **GOMAXPROCS**: Processor configuration and tuning
5. **Performance Optimization**: Scheduler-aware programming
6. **Advanced Patterns**: Real-world scheduling applications

---

## 🎯 Success Criteria

### Basic Level
- [ ] All basic examples run without errors
- [ ] Understand scheduler components
- [ ] Can implement basic work stealing
- [ ] Know how to configure GOMAXPROCS

### Intermediate Level
- [ ] All exercises complete successfully
- [ ] Can implement work distribution strategies
- [ ] Understand context switching overhead
- [ ] Can optimize goroutine count

### Advanced Level
- [ ] All advanced patterns work correctly
- [ ] Can implement scheduler-aware applications
- [ ] Understand performance implications
- [ ] Can debug scheduler issues

---

## 🚀 Next Steps

1. **Complete all tests** in this guide
2. **Experiment** with different configurations
3. **Apply** techniques to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** advanced scheduling patterns

Ready to become an Advanced Scheduling expert? Let's test! 💪

