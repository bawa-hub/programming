# Distributed Real-Time Analytics Platform Makefile

.PHONY: help build test clean run docker-build docker-run

# Default target
help: ## Show this help message
	@echo "Distributed Real-Time Analytics Platform"
	@echo "========================================"
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build all binaries
	@echo "Building all binaries..."
	@go build -o bin/server ./cmd/server
	@go build -o bin/worker ./cmd/worker
	@go build -o bin/client ./cmd/client
	@echo "Build complete!"

build-server: ## Build server binary
	@echo "Building server..."
	@go build -o bin/server ./cmd/server

build-worker: ## Build worker binary
	@echo "Building worker..."
	@go build -o bin/worker ./cmd/worker

build-client: ## Build client binary
	@echo "Building client..."
	@go build -o bin/client ./cmd/client

# Test targets
test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	@go test -race -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-performance: ## Run performance tests
	@echo "Running performance tests..."
	@go test -bench=. -benchmem ./tests/performance/...

test-chaos: ## Run chaos engineering tests
	@echo "Running chaos tests..."
	@go test -v ./tests/chaos/...

# Development targets
run: ## Run the platform (requires dependencies)
	@echo "Starting platform..."
	@docker-compose up -d
	@echo "Waiting for services to start..."
	@sleep 10
	@./bin/server --config configs/server.yaml &
	@./bin/worker --config configs/worker.yaml &
	@echo "Platform started! Dashboard: http://localhost:8080"

run-server: ## Run only the server
	@echo "Starting server..."
	@./bin/server --config configs/server.yaml

run-worker: ## Run only a worker
	@echo "Starting worker..."
	@./bin/worker --config configs/worker.yaml

# Docker targets
docker-build: ## Build Docker images
	@echo "Building Docker images..."
	@docker build -t analytics-server -f Dockerfile.server .
	@docker build -t analytics-worker -f Dockerfile.worker .
	@docker build -t analytics-client -f Dockerfile.client .

docker-run: ## Run with Docker Compose
	@echo "Starting with Docker Compose..."
	@docker-compose up -d

docker-stop: ## Stop Docker Compose
	@echo "Stopping Docker Compose..."
	@docker-compose down

# Utility targets
clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean

deps: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

# Benchmark targets
bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./pkg/concurrency/...

bench-cpu: ## Run CPU benchmarks
	@echo "Running CPU benchmarks..."
	@go test -bench=. -benchmem -cpuprofile=cpu.prof ./pkg/concurrency/...

bench-mem: ## Run memory benchmarks
	@echo "Running memory benchmarks..."
	@go test -bench=. -benchmem -memprofile=mem.prof ./pkg/concurrency/...

# Profiling targets
profile-cpu: ## Generate CPU profile
	@echo "Generating CPU profile..."
	@go tool pprof cpu.prof

profile-mem: ## Generate memory profile
	@echo "Generating memory profile..."
	@go tool pprof mem.prof

# Documentation targets
docs: ## Generate documentation
	@echo "Generating documentation..."
	@go doc -all ./... > docs/api/godoc.txt

# Development setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	@go mod download
	@go mod tidy
	@mkdir -p bin
	@echo "Development environment ready!"

# Production targets
prod-build: ## Build for production
	@echo "Building for production..."
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/worker ./cmd/worker
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/client ./cmd/client

# Monitoring targets
monitor: ## Start monitoring dashboard
	@echo "Starting monitoring dashboard..."
	@open http://localhost:8080/monitor

logs: ## View logs
	@echo "Viewing logs..."
	@docker-compose logs -f

# Security targets
security: ## Run security checks
	@echo "Running security checks..."
	@gosec ./...

# All targets
all: clean deps fmt lint test build ## Run all checks and build
