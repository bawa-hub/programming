# File System Scanner - Advanced Concurrency Project

A powerful file system scanning and analysis tool that demonstrates advanced Go concurrency patterns, file system operations, and system programming concepts.

## 🎯 Project Overview

This File System Scanner is designed to help you master advanced Go concepts through hands-on practice with real-world system programming challenges. It implements a high-performance file system scanner using sophisticated concurrency patterns.

## ✨ Features

### Core Scanning Capabilities
- **Concurrent File System Scanning** with configurable worker pools
- **Advanced File Analysis** with detailed metadata extraction
- **Duplicate File Detection** using hash-based comparison
- **File Hash Calculation** (MD5, SHA1, SHA256)
- **Comprehensive Statistics** and analytics
- **Progress Reporting** with real-time updates
- **Error Handling** with detailed error reporting

### Advanced Concurrency Features
- **Worker Pool Pattern** for efficient resource utilization
- **Rate Limiting** to prevent system overload
- **Semaphore-based** concurrency control
- **Pipeline Processing** for complex data flows
- **Context-based** cancellation and timeout handling
- **Memory-efficient** streaming processing

### Analysis and Reporting
- **File Type Analysis** with visual indicators
- **Extension-based** categorization
- **Size Distribution** analysis
- **Time-based** file analysis
- **Permission Statistics** tracking
- **Depth Analysis** for directory structures
- **Smart Recommendations** for cleanup

### Export and Integration
- **Multiple Export Formats** (JSON, CSV, TXT)
- **Interactive CLI** with command completion
- **Batch Processing** capabilities
- **Configurable Options** for different use cases

## 🏗️ Architecture

The application demonstrates clean architecture principles with advanced concurrency patterns:

```
file-scanner/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/               # Internal packages
│   ├── cli/               # Interactive command-line interface
│   ├── scanner/           # Core scanning engine
│   └── analytics/         # Analysis and reporting
├── pkg/                   # Public packages
│   ├── models/            # Data models and structures
│   ├── patterns/          # Concurrency patterns
│   └── utils/             # Utility functions
├── testdata/              # Test data directory
└── build/                 # Build output
```

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or later
- Make (optional, for build automation)

### Installation

1. **Build the application**
   ```bash
   make build
   # or
   go build -o build/file-scanner ./cmd
   ```

2. **Run in interactive mode**
   ```bash
   make run
   # or
   ./build/file-scanner -interactive
   ```

3. **Scan a directory**
   ```bash
   ./build/file-scanner -path /path/to/directory -verbose
   ```

### Quick Start

1. **Create test data and run demo**
   ```bash
   make demo
   ```

2. **Run performance tests**
   ```bash
   make perf-test
   ```

3. **Run comprehensive tests**
   ```bash
   make test-coverage
   ```

## 🎮 Usage Examples

### Command Line Usage

```bash
# Basic scan
./file-scanner -path /home/user

# Advanced scan with options
./file-scanner -path /var/log -depth 5 -duplicates -hashes -verbose

# Export results
./file-scanner -path /tmp -export json -output scan_results.json

# Interactive mode
./file-scanner -interactive
```

### Interactive Commands

```bash
# Start interactive mode
scanner> scan /path/to/directory

# Scan with options
scanner> scan /home/user -depth 10 -duplicates -hashes

# Analyze results
scanner> analyze

# Filter results
scanner> filter type regular
scanner> filter size 1MB-100MB

# Search files
scanner> search "*.txt"

# Show statistics
scanner> stats

# Export results
scanner> export json results.json

# Show help
scanner> help
```

## 📚 Advanced Go Concepts Demonstrated

### 1. Concurrency Patterns
- **Worker Pool Pattern**: Efficient task distribution
- **Pipeline Processing**: Sequential data transformation
- **Rate Limiting**: Controlled resource usage
- **Semaphores**: Concurrency control
- **Context Management**: Cancellation and timeouts

### 2. File System Operations
- **Directory Traversal**: Recursive scanning
- **File Metadata**: Comprehensive information extraction
- **Hash Calculation**: Cryptographic hashing
- **Error Handling**: Robust error recovery
- **Memory Management**: Efficient resource usage

### 3. Data Structures and Algorithms
- **Custom Data Types**: File and directory models
- **Hash Tables**: Duplicate detection
- **Sorting Algorithms**: File organization
- **Search Algorithms**: Pattern matching
- **Statistical Analysis**: Data aggregation

### 4. System Programming
- **File I/O**: Efficient file operations
- **Memory Management**: Zero-copy operations
- **Performance Optimization**: Concurrent processing
- **Resource Management**: Proper cleanup
- **Error Recovery**: Graceful failure handling

### 5. Advanced Go Features
- **Interfaces**: Polymorphic design
- **Goroutines**: Concurrent execution
- **Channels**: Communication between goroutines
- **Mutexes**: Thread-safe operations
- **Context Package**: Cancellation and timeouts
- **Reflection**: Dynamic type handling
- **Profiling**: Performance analysis

## 🧪 Testing and Benchmarking

### Comprehensive Test Suite
```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific package tests
make test-models
make test-scanner
make test-analytics

# Run benchmarks
make bench
```

### Performance Testing
```bash
# Performance comparison
make perf-test

# Memory profiling
make mem-test

# CPU profiling
make cpu-test
```

## 🔧 Development

### Build Commands
```bash
make build        # Build the application
make run          # Run in interactive mode
make test         # Run all tests
make fmt          # Format code
make lint         # Run linter
make clean        # Clean build artifacts
```

### Project Structure
- **Models**: Data structures and business logic
- **Scanner**: Core scanning engine with concurrency
- **Analytics**: Analysis and reporting engine
- **CLI**: Interactive command-line interface
- **Patterns**: Reusable concurrency patterns
- **Tests**: Comprehensive test coverage

## 📊 Performance Features

### Concurrency Optimization
- **Configurable Worker Pools**: Optimal resource utilization
- **Rate Limiting**: Prevents system overload
- **Memory-efficient Processing**: Streaming data handling
- **Context-aware Cancellation**: Responsive user control

### Memory Management
- **Zero-copy Operations**: Efficient data handling
- **Streaming Processing**: Large file support
- **Garbage Collection**: Optimized memory usage
- **Resource Pooling**: Reusable objects

### I/O Optimization
- **Concurrent File Operations**: Parallel processing
- **Buffered I/O**: Efficient data transfer
- **Error Recovery**: Robust failure handling
- **Progress Reporting**: Real-time feedback

## 🎓 Learning Outcomes

After working with this project, you will have mastered:

1. **Advanced Concurrency**: Worker pools, pipelines, rate limiting
2. **File System Programming**: I/O operations, metadata extraction
3. **Performance Optimization**: Profiling, benchmarking, memory management
4. **System Programming**: Low-level operations, resource management
5. **Error Handling**: Robust error recovery and reporting
6. **Testing**: Comprehensive test coverage and benchmarking
7. **CLI Development**: Interactive command-line interfaces
8. **Data Analysis**: Statistical analysis and reporting
9. **Memory Management**: Efficient resource utilization
10. **Real-world Applications**: Production-ready system software

## 🚀 Next Steps

This project provides a solid foundation for:
- **Distributed Systems**: Microservices and distributed processing
- **Database Systems**: Storage engines and query processing
- **Web Servers**: HTTP servers and API development
- **Cloud Computing**: Container orchestration and deployment
- **System Administration**: Automation and monitoring tools

## 🤝 Contributing

This is a learning project, but feel free to:
- Add new scanning features
- Implement additional concurrency patterns
- Enhance performance optimizations
- Add more export formats
- Improve error handling
- Add more comprehensive tests

## 📝 License

This project is for educational purposes and demonstrates advanced Go concurrency patterns.

---

**Happy Learning! 🎉**

This File System Scanner demonstrates that you've mastered advanced Go concepts and are ready to build complex, production-ready system software. The concurrency patterns and performance optimizations you've learned here will serve as the foundation for all your future Go development in system programming, distributed systems, and high-performance applications.