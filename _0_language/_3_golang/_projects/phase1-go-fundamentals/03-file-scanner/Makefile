# File System Scanner Makefile

.PHONY: build run test clean help install

# Variables
BINARY_NAME=file-scanner
BUILD_DIR=build
MAIN_PACKAGE=./cmd

# Default target
all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application in interactive mode
run: build
	@echo "Running $(BINARY_NAME) in interactive mode..."
	@./$(BUILD_DIR)/$(BINARY_NAME) -interactive

# Run with a specific path
run-scan: build
	@echo "Running $(BINARY_NAME) with test directory..."
	@./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -depth 3 -verbose

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests for specific package
test-models:
	@echo "Running model tests..."
	@go test -v ./pkg/models

test-scanner:
	@echo "Running scanner tests..."
	@go test -v ./internal/scanner

test-analytics:
	@echo "Running analytics tests..."
	@go test -v ./internal/analytics

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	@go test -bench=. ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@rm -rf testdata/
	@echo "Clean complete"

# Install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Create test data
test-data:
	@echo "Creating test data..."
	@mkdir -p testdata/{dir1,dir2,dir3/subdir}
	@echo "Test file 1" > testdata/file1.txt
	@echo "Test file 2" > testdata/file2.txt
	@echo "Test file 3" > testdata/dir1/file3.txt
	@echo "Test file 4" > testdata/dir2/file4.txt
	@echo "Test file 5" > testdata/dir3/subdir/file5.txt
	@echo "Test data created in testdata/"

# Demo the application
demo: build test-data
	@echo "Running demo..."
	@echo "This will demonstrate the File System Scanner capabilities"
	@echo ""
	@echo "Scanning testdata directory with various options:"
	@echo "1. Basic scan:"
	@./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -verbose
	@echo ""
	@echo "2. Scan with duplicates detection:"
	@./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -duplicates -verbose
	@echo ""
	@echo "3. Interactive mode:"
	@echo "Type 'scan ./testdata' to scan the test directory"
	@echo "Type 'help' for available commands"
	@echo "Type 'quit' to exit"
	@./$(BUILD_DIR)/$(BINARY_NAME) -interactive

# Performance test
perf-test: build test-data
	@echo "Running performance test..."
	@echo "Scanning testdata with different concurrency levels:"
	@echo ""
	@echo "1 worker:"
	@time ./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -concurrency 1 -duplicates
	@echo ""
	@echo "2 workers:"
	@time ./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -concurrency 2 -duplicates
	@echo ""
	@echo "4 workers:"
	@time ./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -concurrency 4 -duplicates
	@echo ""
	@echo "8 workers:"
	@time ./$(BUILD_DIR)/$(BINARY_NAME) -path ./testdata -concurrency 8 -duplicates

# Memory test
mem-test: build test-data
	@echo "Running memory test..."
	@go test -v -memprofile=mem.prof ./...
	@go tool pprof -web mem.prof
	@echo "Memory profile generated: mem.prof"

# CPU test
cpu-test: build test-data
	@echo "Running CPU test..."
	@go test -v -cpuprofile=cpu.prof ./...
	@go tool pprof -web cpu.prof
	@echo "CPU profile generated: cpu.prof"

# Show help
help:
	@echo "File System Scanner - Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run in interactive mode"
	@echo "  run-scan     - Run with test directory"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  test-models  - Run model tests only"
	@echo "  test-scanner - Run scanner tests only"
	@echo "  test-analytics- Run analytics tests only"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  bench        - Run benchmarks"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install the application"
	@echo "  test-data    - Create test data"
	@echo "  demo         - Run demo with test data"
	@echo "  perf-test    - Run performance tests"
	@echo "  mem-test     - Run memory tests"
	@echo "  cpu-test     - Run CPU tests"
	@echo "  help         - Show this help"
	@echo ""
	@echo "This application demonstrates advanced Go concepts:"
	@echo "  • Concurrency patterns (worker pools, goroutines)"
	@echo "  • File system operations and I/O"
	@echo "  • Data structures and algorithms"
	@echo "  • Error handling and recovery"
	@echo "  • Command-line interface development"
	@echo "  • Performance optimization"
	@echo "  • Memory management"
	@echo "  • Testing and benchmarking"
