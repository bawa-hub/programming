# 🚀 Microservices Communication Commands

## Quick Reference Commands

### Basic Commands
```bash
# Run all tests
./quick_test_fast.sh

# Run specific examples
go run . basic      # Basic microservices examples
go run . exercises  # Hands-on exercises
go run . advanced   # Advanced patterns
go run . all        # Everything

# Compilation and analysis
go build .          # Compile
go vet .            # Static analysis
go run -race . basic # Race detection
```

### Development Commands
```bash
# Initialize module
go mod init microservices-examples

# Add dependencies
go get <package>

# Clean build
go clean
go build .

# Run with profiling
go run . basic
go tool pprof -http=:8080 mem.prof
```

### Testing Commands
```bash
# Run tests
go test .

# Run benchmarks
go test -bench=.

# Run with coverage
go test -cover .

# Run with race detection
go test -race .
```

### Debugging Commands
```bash
# Enable service discovery logging
go run . basic | grep -i "service"

# Check load balancer state
go run . basic | grep -i "balancer"

# Check circuit breaker state
go run . basic | grep -i "circuit"

# Check rate limiter state
go run . basic | grep -i "rate"
```

### Performance Commands
```bash
# CPU profiling
go run . basic
go tool pprof -http=:8080 cpu.prof

# Memory profiling
go run . basic
go tool pprof -http=:8080 mem.prof

# Goroutine profiling
go run . basic
go tool pprof -http=:8080 goroutine.prof

# Block profiling
go run . basic
go tool pprof -http=:8080 block.prof
```

### Monitoring Commands
```bash
# Monitor service discovery
watch -n 1 'go run . basic | grep "Service found"'

# Monitor load balancing
watch -n 1 'go run . basic | grep "routed to"'

# Monitor circuit breaker
watch -n 1 'go run . basic | grep "Circuit breaker"'

# Monitor rate limiting
watch -n 1 'go run . basic | grep "Rate limited"'
```

### Advanced Commands
```bash
# Run with specific service discovery
SERVICE_DISCOVERY=consul go run . basic

# Run with specific load balancer
LOAD_BALANCER=round_robin go run . basic

# Run with specific circuit breaker
CIRCUIT_BREAKER=threshold=5 go run . basic

# Run with specific rate limiter
RATE_LIMITER=10/s go run . basic
```

### Cleanup Commands
```bash
# Clean build artifacts
go clean

# Clean module cache
go clean -modcache

# Remove test files
rm -f *.prof
rm -f microservices_test
```

### Help Commands
```bash
# Show help
go run . help

# Show available commands
go run . --help

# Show version
go version

# Show environment
go env
```

---

## 🎯 Command Categories

### 1. Basic Testing
- `./quick_test_fast.sh` - Run all tests
- `go run . basic` - Basic examples
- `go run . exercises` - Hands-on exercises
- `go run . advanced` - Advanced patterns

### 2. Compilation & Analysis
- `go build .` - Compile
- `go vet .` - Static analysis
- `go run -race . basic` - Race detection

### 3. Performance Testing
- `go run . basic | grep Performance` - Performance metrics
- `go tool pprof -http=:8080 cpu.prof` - CPU profiling
- `go tool pprof -http=:8080 mem.prof` - Memory profiling

### 4. Debugging
- `go run . basic | grep -i "service"` - Service discovery
- `go run . basic | grep -i "balancer"` - Load balancing
- `go run . basic | grep -i "circuit"` - Circuit breaker
- `go run . basic | grep -i "rate"` - Rate limiting

### 5. Monitoring
- `watch -n 1 'go run . basic | grep "Service found"'` - Service discovery
- `watch -n 1 'go run . basic | grep "routed to"'` - Load balancing
- `watch -n 1 'go run . basic | grep "Circuit breaker"'` - Circuit breaker
- `watch -n 1 'go run . basic | grep "Rate limited"'` - Rate limiting

---

## 🔧 Environment Variables

### Service Discovery
```bash
# Set service discovery backend
export SERVICE_DISCOVERY=consul
export CONSUL_ADDR=localhost:8500

# Set service registry
export SERVICE_REGISTRY=etcd
export ETCD_ENDPOINTS=localhost:2379
```

### Load Balancing
```bash
# Set load balancer algorithm
export LOAD_BALANCER=round_robin
export LOAD_BALANCER=least_connections
export LOAD_BALANCER=random

# Set load balancer weights
export LOAD_BALANCER_WEIGHTS=1,2,3
```

### Circuit Breaker
```bash
# Set circuit breaker threshold
export CIRCUIT_BREAKER_THRESHOLD=5

# Set circuit breaker timeout
export CIRCUIT_BREAKER_TIMEOUT=30s

# Set circuit breaker half-open timeout
export CIRCUIT_BREAKER_HALF_OPEN_TIMEOUT=10s
```

### Rate Limiting
```bash
# Set rate limiter rate
export RATE_LIMITER_RATE=10

# Set rate limiter interval
export RATE_LIMITER_INTERVAL=1s

# Set rate limiter burst
export RATE_LIMITER_BURST=20
```

---

## 📊 Output Examples

### Basic Examples Output
```
🏗️ Microservices Communication Examples
=======================================

1. Basic Microservices Architecture
===================================
  Processing order for user user-123, product product-456
  User retrieved: John Doe
  Order created: order-123
  Payment processed: payment-123
  Order confirmed successfully
  Basic microservices architecture completed
```

### Exercises Output
```
💪 Microservices Communication Exercises
=======================================

Exercise 1: Implement Basic Service Discovery
=============================================
  Testing service discovery...
  Found instance: localhost:8081
  Found instance: localhost:8082
  Exercise 1: Basic service discovery completed
```

### Advanced Patterns Output
```
🚀 Advanced Microservices Patterns
==================================

1. Service Mesh with Sidecar Proxy
  client -> user-service (via sidecar): get-user
  client -> user-service (via sidecar): get-user
  client -> user-service (via sidecar): get-user
```

---

## 🎉 Success Indicators

### ✅ All Tests Pass
```
🎉 Fast tests passed! Microservices Communication is working!
Ready to move to the next topic!
```

### ✅ Performance Metrics
```
  Service Discovery: < 5ms
  Load Balancing: < 1ms
  Circuit Breaker: < 10ms
  Rate Limiting: < 1ms
```

### ✅ Service Discovery
```
  Registered service: user-service at localhost:8081
  Service found: localhost:8081
  Service found: localhost:8082
```

---

## 🚀 Ready to Test!

Use these commands to test your Microservices Communication knowledge:

1. **Start with basic examples**: `go run . basic`
2. **Run exercises**: `go run . exercises`
3. **Test advanced patterns**: `go run . advanced`
4. **Run all tests**: `./quick_test_fast.sh`
5. **Check performance**: `go run . all | grep Performance`

Happy testing! 💪

