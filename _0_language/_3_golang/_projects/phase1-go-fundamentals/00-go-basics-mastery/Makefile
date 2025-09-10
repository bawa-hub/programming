# Go Basics Mastery Makefile

.PHONY: run run-primitives run-arrays run-structs run-interfaces run-pointers run-examples build clean help

# Default target
all: run

# Run all modules
run:
	@echo "Running all Go Basics Mastery modules..."
	@go run main.go -module all

# Run specific modules
run-primitives:
	@echo "Running primitives module..."
	@go run main.go -module primitives

run-arrays:
	@echo "Running arrays and slices module..."
	@go run main.go -module arrays

run-structs:
	@echo "Running structs module..."
	@go run main.go -module structs

run-interfaces:
	@echo "Running interfaces module..."
	@go run main.go -module interfaces

run-pointers:
	@echo "Running pointers module..."
	@go run main.go -module pointers

run-examples:
	@echo "Running examples module..."
	@go run main.go -module examples

# Build the application
build:
	@echo "Building Go Basics Mastery..."
	@go build -o go-basics-mastery main.go
	@echo "Build complete: go-basics-mastery"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f go-basics-mastery
	@echo "Clean complete"

# Show help
help:
	@echo "Go Basics Mastery - Available targets:"
	@echo "  run              - Run all modules"
	@echo "  run-primitives   - Run primitives module"
	@echo "  run-arrays       - Run arrays and slices module"
	@echo "  run-structs      - Run structs module"
	@echo "  run-interfaces   - Run interfaces module"
	@echo "  run-pointers     - Run pointers module"
	@echo "  run-examples     - Run examples module"
	@echo "  build            - Build the application"
	@echo "  test             - Run tests"
	@echo "  fmt              - Format code"
	@echo "  lint             - Run linter"
	@echo "  clean            - Clean build artifacts"
	@echo "  help             - Show this help"
	@echo ""
	@echo "Learning Path:"
	@echo "  1. make run-primitives  # Start with basic data types"
	@echo "  2. make run-arrays      # Learn collections"
	@echo "  3. make run-structs     # Understand structs and methods"
	@echo "  4. make run-interfaces  # Master polymorphism"
	@echo "  5. make run-pointers    # Learn memory management"
	@echo "  6. make run-examples    # See everything combined"
	@echo "  7. make run             # Run everything"
