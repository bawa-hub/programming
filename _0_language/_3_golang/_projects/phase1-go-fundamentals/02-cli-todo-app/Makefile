# CLI Todo App Makefile

.PHONY: build run test clean help install

# Variables
BINARY_NAME=cli-todo-app
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

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Run in memory mode (no persistence)
run-memory: build
	@echo "Running $(BINARY_NAME) in memory mode..."
	@./$(BUILD_DIR)/$(BINARY_NAME) -memory

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

test-storage:
	@echo "Running storage tests..."
	@go test -v ./internal/storage

test-todo:
	@echo "Running todo service tests..."
	@go test -v ./internal/todo

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
	@rm -rf data/
	@echo "Clean complete"

# Install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Create sample data
sample-data:
	@echo "Creating sample data..."
	@mkdir -p data
	@echo '{"todos":[{"id":1,"title":"Learn Go","description":"Master Go programming language","priority":2,"status":1,"tags":["programming","learning"],"created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}],"categories":[{"id":1,"name":"Learning","description":"Educational tasks","color":"blue","created_at":"2024-01-01T00:00:00Z"}],"next_id":2}' > data/todos.json
	@echo "Sample data created in data/todos.json"

# Demo the application
demo: build sample-data
	@echo "Running demo..."
	@echo "This will show you how to use the CLI Todo App"
	@echo "Commands to try:"
	@echo "  list"
	@echo "  add 'Complete Go tutorial'"
	@echo "  add 'Write tests' 'Add unit tests for all functions'"
	@echo "  priority 1 high"
	@echo "  due 1 2024-12-31"
	@echo "  tag 1 add urgent"
	@echo "  complete 1"
	@echo "  stats"
	@echo "  quit"
	@echo ""
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Show help
help:
	@echo "CLI Todo App - Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  run-memory   - Run in memory mode (no persistence)"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  test-models  - Run model tests only"
	@echo "  test-storage - Run storage tests only"
	@echo "  test-todo    - Run todo service tests only"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  bench        - Run benchmarks"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install the application"
	@echo "  sample-data  - Create sample data"
	@echo "  demo         - Run demo with sample data"
	@echo "  help         - Show this help"
	@echo ""
	@echo "This application demonstrates all Go fundamentals:"
	@echo "  • Primitive data types (int, string, bool, time.Time)"
	@echo "  • Arrays and slices for collections"
	@echo "  • Structs for data modeling"
	@echo "  • Interfaces for polymorphism"
	@echo "  • Pointers for memory management"
	@echo "  • Error handling and validation"
	@echo "  • File I/O and JSON processing"
	@echo "  • Concurrency with goroutines"
	@echo "  • Command-line argument parsing"
