# Go Standard Library Mastery Makefile

.PHONY: help build test clean run-fmt run-os run-io run-time run-math run-reflect run-errors run-log run-all

# Default target
help:
	@echo "🚀 Go Standard Library Mastery"
	@echo "=============================="
	@echo ""
	@echo "Available targets:"
	@echo "  build        - Build all examples"
	@echo "  test         - Run all tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  run-fmt      - Run fmt package examples"
	@echo "  run-os       - Run os package examples"
	@echo "  run-io       - Run io package examples"
	@echo "  run-time     - Run time package examples"
	@echo "  run-math     - Run math package examples"
	@echo "  run-reflect  - Run reflect package examples"
	@echo "  run-errors   - Run errors package examples"
	@echo "  run-log      - Run log package examples"
	@echo "  run-all      - Run all package examples"
	@echo ""
	@echo "Examples:"
	@echo "  make run-fmt"
	@echo "  make run-all"
	@echo "  make test"

# Build all examples
build:
	@echo "Building all examples..."
	@go build -o build/fmt-examples ./core/fmt.go
	@go build -o build/os-examples ./core/os.go
	@go build -o build/io-examples ./core/io.go
	@echo "Build complete!"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./core/...
	@echo "Tests complete!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf build/
	@go clean
	@echo "Clean complete!"

# Run fmt package examples
run-fmt:
	@echo "Running fmt package examples..."
	@go run ./core/fmt.go

# Run os package examples
run-os:
	@echo "Running os package examples..."
	@go run ./core/os.go

# Run io package examples
run-io:
	@echo "Running io package examples..."
	@go run ./core/io.go

# Run time package examples
run-time:
	@echo "Running time package examples..."
	@go run ./core/time.go

# Run math package examples
run-math:
	@echo "Running math package examples..."
	@go run ./core/math.go

# Run reflect package examples
run-reflect:
	@echo "Running reflect package examples..."
	@go run ./core/reflect.go

# Run errors package examples
run-errors:
	@echo "Running errors package examples..."
	@go run ./core/errors.go

# Run log package examples
run-log:
	@echo "Running log package examples..."
	@go run ./core/log.go

# Run all package examples
run-all:
	@echo "Running all package examples..."
	@echo "=============================="
	@echo ""
	@echo "1. fmt package:"
	@echo "==============="
	@go run ./core/fmt.go
	@echo ""
	@echo "2. os package:"
	@echo "=============="
	@go run ./core/os.go
	@echo ""
	@echo "3. io package:"
	@echo "=============="
	@go run ./core/io.go
	@echo ""
	@echo "4. time package:"
	@echo "==============="
	@go run ./core/time.go
	@echo ""
	@echo "5. math package:"
	@echo "==============="
	@go run ./core/math.go
	@echo ""
	@echo "6. reflect package:"
	@echo "=================="
	@go run ./core/reflect.go
	@echo ""
	@echo "7. errors package:"
	@echo "================="
	@go run ./core/errors.go
	@echo ""
	@echo "8. log package:"
	@echo "=============="
	@go run ./core/log.go
	@echo ""
	@echo "🎉 All examples completed!"

# Create build directory
build-dir:
	@mkdir -p build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download
	@echo "Dependencies installed!"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Code formatted!"

# Lint code
lint:
	@echo "Linting code..."
	@go vet ./...
	@echo "Linting complete!"

# Run specific package with main.go
run-package:
	@if [ -z "$(PACKAGE)" ]; then \
		echo "Usage: make run-package PACKAGE=<package_name>"; \
		echo "Available packages: fmt, os, io, time, math, reflect, errors, log, all"; \
	else \
		go run main.go -package $(PACKAGE); \
	fi
