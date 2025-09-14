# Pipeline Pattern Commands Reference

## Quick Reference

### Run All Tests
```bash
./quick_test.sh
```

### Individual Commands

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

#### Compilation
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

## Detailed Commands

### 1. Basic Examples
```bash
# Run all basic examples
go run . basic

# Run specific example (if implemented)
go run . basic --example=1
go run . basic --example=2
```

### 2. Exercises
```bash
# Run all exercises
go run . exercises

# Run specific exercise (if implemented)
go run . exercises --exercise=1
go run . exercises --exercise=2
```

### 3. Advanced Patterns
```bash
# Run all advanced patterns
go run . advanced

# Run specific pattern (if implemented)
go run . advanced --pattern=adaptive
go run . advanced --pattern=circuit-breaker
```

### 4. Testing and Analysis

#### Compilation Test
```bash
# Basic compilation
go build .

# Cross-platform compilation
go build -o pipeline-linux GOOS=linux GOARCH=amd64 .
go build -o pipeline-windows GOOS=windows GOARCH=amd64 .
```

#### Race Detection
```bash
# Basic race detection
go run -race . basic

# Race detection with specific flags
go run -race -race=1 . basic
```

#### Static Analysis
```bash
# Basic static analysis
go vet .

# Static analysis with specific packages
go vet ./...

# Static analysis with additional checks
go vet -composites=false .
```

#### Performance Testing
```bash
# Basic benchmark
go test -bench=.

# Benchmark with memory profiling
go test -bench=. -benchmem

# Benchmark specific functions
go test -bench=BenchmarkPipeline -benchmem

# Benchmark with CPU profiling
go test -bench=. -cpuprofile=cpu.prof
```

### 5. Code Quality

#### Formatting
```bash
# Format code
go fmt .

# Format with specific options
go fmt -w .
```

#### Linting
```bash
# Install golangci-lint (if not installed)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Run specific linters
golangci-lint run --enable=gofmt,goimports,vet
```

#### Documentation
```bash
# Generate documentation
go doc .

# Generate documentation for specific function
go doc PipelinePattern

# Generate documentation in HTML
godoc -http=:6060
```

### 6. Debugging

#### Debug Build
```bash
# Build with debug information
go build -gcflags="all=-N -l" .

# Run with debug information
go run -gcflags="all=-N -l" . basic
```

#### Profiling
```bash
# CPU profiling
go run . basic -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling
go run . basic -memprofile=mem.prof
go tool pprof mem.prof

# Goroutine profiling
go run . basic -blockprofile=block.prof
go tool pprof block.prof
```

#### Trace Analysis
```bash
# Generate trace
go run . basic -trace=trace.out
go tool trace trace.out
```

### 7. Module Management

#### Dependencies
```bash
# Initialize module
go mod init pipeline-pattern

# Add dependencies
go get github.com/example/package

# Update dependencies
go get -u ./...

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

#### Version Management
```bash
# Check Go version
go version

# Check module versions
go list -m all

# Check for updates
go list -m -u all
```

### 8. Build and Deployment

#### Build Options
```bash
# Basic build
go build .

# Build with optimizations
go build -ldflags="-s -w" .

# Build with version information
go build -ldflags="-X main.Version=1.0.0" .

# Build for different architectures
go build -o pipeline-amd64 .
go build -o pipeline-arm64 .
```

#### Cross-Compilation
```bash
# Linux
GOOS=linux GOARCH=amd64 go build .

# Windows
GOOS=windows GOARCH=amd64 go build .

# macOS
GOOS=darwin GOARCH=amd64 go build .
```

### 9. Testing Specific Scenarios

#### Error Handling
```bash
# Test error handling
go run . basic --test-errors

# Test with specific error rates
go run . basic --error-rate=0.5
```

#### Performance Testing
```bash
# Test with different loads
go run . basic --load=100
go run . basic --load=1000

# Test with different buffer sizes
go run . basic --buffer-size=10
go run . basic --buffer-size=100
```

#### Stress Testing
```bash
# Run stress test
go run . basic --stress-test

# Run with timeout
timeout 30s go run . basic --stress-test
```

### 10. Monitoring and Metrics

#### Runtime Metrics
```bash
# Run with metrics
go run . basic --metrics

# Run with specific metrics
go run . basic --metrics=cpu,memory,goroutines
```

#### Logging
```bash
# Run with debug logging
go run . basic --debug

# Run with specific log level
go run . basic --log-level=info
go run . basic --log-level=debug
```

## Environment Variables

### Go-specific
```bash
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
```

### Pipeline-specific
```bash
export PIPELINE_BUFFER_SIZE=100
export PIPELINE_WORKER_COUNT=4
export PIPELINE_TIMEOUT=30s
```

## Common Issues and Solutions

### 1. Permission Denied
```bash
chmod +x quick_test.sh
```

### 2. Module Not Found
```bash
go mod tidy
go mod download
```

### 3. Race Detection Issues
```bash
# For educational race conditions, this is expected
go run -race . basic 2>&1 | grep -v "WARNING: DATA RACE"
```

### 4. Build Failures
```bash
# Clean and rebuild
go clean
go build .
```

### 5. Test Failures
```bash
# Run with verbose output
go run . basic -v

# Run with specific flags
go run . basic --help
```

## Tips and Best Practices

1. **Always run tests before committing**
2. **Use race detection in development**
3. **Profile performance-critical code**
4. **Keep dependencies up to date**
5. **Use proper error handling**
6. **Document complex patterns**
7. **Test with different loads**
8. **Monitor resource usage**
9. **Use proper logging**
10. **Follow Go conventions**

## Next Steps

After mastering these commands:
1. Experiment with different pipeline configurations
2. Try implementing your own patterns
3. Move on to the next topic: Fan-Out/Fan-In Pattern
4. Explore advanced Go concurrency features
