# Go CLI Calculator Makefile

.PHONY: build test run clean help install

# Variables
BINARY_NAME=cli-calculator
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

# Run the application in interactive mode
run:
	@echo "Running $(BINARY_NAME) in interactive mode..."
	@go run $(MAIN_PACKAGE) -i

# Run the application with an expression
run-expr:
	@echo "Running $(BINARY_NAME) with expression..."
	@go run $(MAIN_PACKAGE) -expr "2 + 3 * 4"

# Install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

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

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  run          - Run in interactive mode"
	@echo "  run-expr     - Run with example expression"
	@echo "  install      - Install the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  bench        - Run benchmarks"
	@echo "  help         - Show this help"
