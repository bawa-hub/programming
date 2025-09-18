# 🚀 Database Concurrency Commands

## Quick Reference Commands

### Basic Commands
```bash
# Run all tests
./quick_test_fast.sh

# Run specific examples
go run . basic      # Basic database concurrency examples
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
go mod init database-concurrency-examples

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
# Enable connection pool logging
go run . basic | grep -i "connection"

# Check transaction state
go run . basic | grep -i "transaction"

# Check deadlock prevention
go run . basic | grep -i "deadlock"

# Check performance metrics
go run . basic | grep -i "performance"
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
# Monitor connection pool stats
watch -n 1 'go run . basic | grep "Pool stats"'

# Monitor transaction operations
watch -n 1 'go run . basic | grep "Transaction"'

# Monitor deadlock prevention
watch -n 1 'go run . basic | grep "Deadlock"'

# Monitor performance metrics
watch -n 1 'go run . basic | grep "Performance"'
```

### Advanced Commands
```bash
# Run with specific connection pool size
CONNECTION_POOL_SIZE=20 go run . basic

# Run with specific transaction timeout
TRANSACTION_TIMEOUT=30s go run . basic

# Run with specific isolation level
ISOLATION_LEVEL=serializable go run . basic

# Run with specific deadlock timeout
DEADLOCK_TIMEOUT=10s go run . basic
```

### Cleanup Commands
```bash
# Clean build artifacts
go clean

# Clean module cache
go clean -modcache

# Remove test files
rm -f *.prof
rm -f database_test
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
- `go run . basic | grep -i "connection"` - Connection pooling
- `go run . basic | grep -i "transaction"` - Transaction management
- `go run . basic | grep -i "deadlock"` - Deadlock prevention
- `go run . basic | grep -i "performance"` - Performance metrics

### 5. Monitoring
- `watch -n 1 'go run . basic | grep "Pool stats"'` - Connection pool
- `watch -n 1 'go run . basic | grep "Transaction"'` - Transactions
- `watch -n 1 'go run . basic | grep "Deadlock"'` - Deadlock prevention
- `watch -n 1 'go run . basic | grep "Performance"'` - Performance

---

## 🔧 Environment Variables

### Connection Pool
```bash
# Set connection pool size
export CONNECTION_POOL_SIZE=20
export MAX_IDLE_CONNECTIONS=10

# Set connection lifetime
export CONNECTION_MAX_LIFETIME=1h
export CONNECTION_MAX_IDLE_TIME=30m
```

### Transaction Management
```bash
# Set transaction timeout
export TRANSACTION_TIMEOUT=30s
export TRANSACTION_RETRY_COUNT=3

# Set isolation level
export ISOLATION_LEVEL=read_committed
export ISOLATION_LEVEL=repeatable_read
export ISOLATION_LEVEL=serializable
```

### Deadlock Prevention
```bash
# Set deadlock timeout
export DEADLOCK_TIMEOUT=10s
export DEADLOCK_RETRY_COUNT=3

# Set lock timeout
export LOCK_TIMEOUT=5s
export LOCK_RETRY_DELAY=100ms
```

### Read Replicas
```bash
# Set read replica configuration
export READ_REPLICA_COUNT=3
export READ_REPLICA_WEIGHT=1.0

# Set load balancing strategy
export LOAD_BALANCING_STRATEGY=round_robin
export LOAD_BALANCING_STRATEGY=least_connections
```

---

## 📊 Output Examples

### Basic Examples Output
```
🗄️ Database Concurrency Examples
=================================

1. Basic Database Connection Management
======================================
  Database manager created successfully
  Max connections: 10
  Max idle connections: 5
  Basic database connection management completed
```

### Exercises Output
```
💪 Database Concurrency Exercises
=================================

Exercise 1: Implement Connection Pooling
=======================================
  Testing connection pooling...
  Acquired connection 1: conn-1234567890
  Released connection 1
  Pool statistics:
    Open connections: 10
    In use: 5
    Idle: 5
  Exercise 1: Connection pooling completed
```

### Advanced Patterns Output
```
🚀 Advanced Database Concurrency Patterns
========================================

1. Multi-Database Connection Pool
  Got connection from db1: conn-1234567890

2. Distributed Transaction Manager
  Distributed transaction dist-tx-1 begun
  Distributed transaction dist-tx-1 committed
```

---

## 🎉 Success Indicators

### ✅ All Tests Pass
```
🎉 Fast tests passed! Database Concurrency is working!
Ready to move to the next topic!
```

### ✅ Performance Metrics
```
  Connection Pool: < 10ms
  Transaction Management: < 5ms
  Read Replicas: < 20ms
  Database Sharding: < 25ms
```

### ✅ Connection Pooling
```
  Acquired connection 1: conn-1234567890
  Released connection 1
  Pool stats - Open: 10, InUse: 5, Idle: 5
```

---

## 🚀 Ready to Test!

Use these commands to test your Database Concurrency knowledge:

1. **Start with basic examples**: `go run . basic`
2. **Run exercises**: `go run . exercises`
3. **Test advanced patterns**: `go run . advanced`
4. **Run all tests**: `./quick_test_fast.sh`
5. **Check performance**: `go run . all | grep Performance`

Happy testing! 💪

