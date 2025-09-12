# Go Standard Library Mastery 🚀

A comprehensive learning resource for mastering Go's standard library through practical examples, detailed notes, and real-world patterns. This project covers all major packages in Go's standard library with hands-on examples that you can run, modify, and learn from.

## 📚 Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Learning Path](#learning-path)
- [Package Categories](#package-categories)
- [Getting Started](#getting-started)
- [Examples](#examples)
- [Best Practices](#best-practices)
- [Performance Tips](#performance-tips)
- [Real-world Applications](#real-world-applications)
- [Contributing](#contributing)

## 🎯 Overview

This project is designed to help you master Go's standard library through:

- **Comprehensive Coverage**: All major Go standard library packages
- **Practical Examples**: 200+ working code examples
- **Detailed Notes**: In-depth explanations and concepts
- **Real-world Patterns**: Production-ready code patterns
- **Performance Insights**: Optimization techniques and benchmarks
- **Best Practices**: Industry-standard coding practices

## 📁 Project Structure

```
04-go-standard-library-mastery/
├── README.md                    # This file
├── go.mod                       # Go module file
├── main.go                      # Main entry point
├── Makefile                     # Build and run commands
├── PROJECT_SUMMARY.md           # Project completion summary
│
├── core/                        # Core packages
│   ├── fmt.md & fmt.go         # Formatting and I/O
│   ├── os.md & os.go           # Operating system interface
│   ├── io.md & io.go           # I/O primitives
│   ├── time.md & time.go       # Time operations
│   ├── math.md & math.go       # Mathematical functions
│   ├── reflect.md & reflect.go # Reflection
│   ├── errors.md & errors.go   # Error handling
│   ├── log.md & log.go         # Logging
│   └── CORE_SUMMARY.md         # Core packages summary
│
├── data-structures/             # Data structure packages
│   ├── sort.md & sort.go       # Sorting algorithms
│   ├── container/
│   │   ├── heap.md & heap.go   # Heap data structure
│   │   ├── list.md & list.go   # Doubly linked list
│   │   └── ring.md & ring.go   # Circular list
│   └── DATA_STRUCTURES_SUMMARY.md
│
├── networking/                  # Networking packages
│   ├── net.md & net.go         # Network I/O
│   ├── http.md & http.go       # HTTP client and server
│   └── url.md & url.go         # URL parsing
│   └── NETWORKING_SUMMARY.md
│
├── concurrency/                 # Concurrency packages
│   ├── sync.md & sync.go       # Synchronization primitives
│   ├── context.md & context.go # Context management
│   └── atomic.md & atomic.go   # Atomic operations
│   └── CONCURRENCY_SUMMARY.md
│
├── encoding/                    # Encoding packages
│   ├── json.md & json.go       # JSON encoding/decoding
│   ├── xml.md & xml.go         # XML encoding/decoding
│   └── ENCODING_SUMMARY.md
│
├── utility/                     # Utility packages
│   ├── strings.md & strings.go # String manipulation
│   ├── strconv.md & strconv.go # String conversion
│   └── UTILITY_SUMMARY.md
│
└── system/                      # System packages
    ├── runtime.md & runtime.go # Runtime system control
    ├── unsafe.md & unsafe.go   # Unsafe operations
    └── SYSTEM_SUMMARY.md
```

## 🛤️ Learning Path

### Phase 1: Core Packages (Start Here)
1. **fmt** - Formatting and I/O operations
2. **os** - Operating system interface
3. **io** - I/O primitives and utilities
4. **time** - Time operations and formatting
5. **math** - Mathematical functions
6. **reflect** - Reflection capabilities
7. **errors** - Error handling
8. **log** - Logging functionality

### Phase 2: Data Structures
1. **sort** - Sorting algorithms and utilities
2. **container/heap** - Heap data structure
3. **container/list** - Doubly linked list
4. **container/ring** - Circular list

### Phase 3: Networking
1. **net** - Network I/O operations
2. **http** - HTTP client and server
3. **url** - URL parsing and manipulation

### Phase 4: Concurrency
1. **sync** - Synchronization primitives
2. **context** - Context management
3. **sync/atomic** - Atomic operations

### Phase 5: Encoding
1. **encoding/json** - JSON encoding/decoding
2. **encoding/xml** - XML encoding/decoding

### Phase 6: Utilities
1. **strings** - String manipulation
2. **strconv** - String conversion

### Phase 7: System Programming
1. **runtime** - Runtime system control
2. **unsafe** - Unsafe operations

## 📦 Package Categories

### 🔧 Core Packages
Essential packages for basic Go programming:
- **fmt**: Formatting and I/O operations
- **os**: Operating system interface
- **io**: I/O primitives and utilities
- **time**: Time operations and formatting
- **math**: Mathematical functions
- **reflect**: Reflection capabilities
- **errors**: Error handling
- **log**: Logging functionality

### 📊 Data Structures
Packages for data manipulation and algorithms:
- **sort**: Sorting algorithms and utilities
- **container/heap**: Heap data structure
- **container/list**: Doubly linked list
- **container/ring**: Circular list

### 🌐 Networking
Packages for network programming:
- **net**: Network I/O operations
- **http**: HTTP client and server
- **url**: URL parsing and manipulation

### ⚡ Concurrency
Packages for concurrent programming:
- **sync**: Synchronization primitives
- **context**: Context management
- **sync/atomic**: Atomic operations

### 🔄 Encoding
Packages for data encoding/decoding:
- **encoding/json**: JSON encoding/decoding
- **encoding/xml**: XML encoding/decoding

### 🛠️ Utilities
Utility packages for common operations:
- **strings**: String manipulation
- **strconv**: String conversion

### ⚙️ System
Low-level system programming packages:
- **runtime**: Runtime system control
- **unsafe**: Unsafe operations

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or later
- Basic understanding of Go syntax
- Terminal/command line access

### Installation
```bash
# Clone or download this project
cd 04-go-standard-library-mastery

# Initialize Go module (if not already done)
go mod init go-standard-library-mastery

# Run all examples
go run main.go

# Or run specific package examples
go run ./core/fmt.go
go run ./data-structures/sort.go
go run ./networking/http.go
# ... and so on
```

### Using the Makefile
```bash
# Run all examples
make run

# Run specific category
make run-core
make run-data-structures
make run-networking
make run-concurrency
make run-encoding
make run-utility
make run-system

# Build all examples
make build

# Clean build artifacts
make clean
```

## 📖 Examples

### Basic Example (fmt package)
```go
package main

import "fmt"

func main() {
    // Basic formatting
    fmt.Printf("Hello, %s!\n", "World")
    
    // String formatting
    name := "Alice"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    
    // Number formatting
    pi := 3.14159
    fmt.Printf("Pi: %.2f\n", pi)
}
```

### Advanced Example (concurrency)
```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    // Context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Worker pool
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(ctx, jobs, results, &wg)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 10; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Collect results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Print results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

func worker(ctx context.Context, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case job, ok := <-jobs:
            if !ok {
                return
            }
            // Simulate work
            time.Sleep(100 * time.Millisecond)
            results <- job * 2
        case <-ctx.Done():
            return
        }
    }
}
```

## 🎯 Best Practices

### 1. Error Handling
```go
// Always handle errors
file, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### 2. Resource Management
```go
// Use defer for cleanup
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Process file
    return nil
}
```

### 3. Context Usage
```go
// Use context for cancellation
func longRunningOperation(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
        }
    }
}
```

### 4. Concurrency Safety
```go
// Use sync primitives for thread safety
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

## ⚡ Performance Tips

### 1. String Operations
```go
// Use strings.Builder for string concatenation
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("hello")
}
result := builder.String()
```

### 2. Memory Allocation
```go
// Pre-allocate slices when size is known
items := make([]Item, 0, 1000) // Capacity 1000
```

### 3. Goroutine Management
```go
// Use worker pools for controlled concurrency
func workerPool(jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        results <- process(job)
    }
}
```

### 4. I/O Operations
```go
// Use buffered I/O for better performance
reader := bufio.NewReader(file)
writer := bufio.NewWriter(file)
```

## 🌟 Real-world Applications

### 1. Web Server
```go
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### 2. File Processing
```go
func processFiles(filenames []string) error {
    for _, filename := range filenames {
        if err := processFile(filename); err != nil {
            return fmt.Errorf("processing %s: %w", filename, err)
        }
    }
    return nil
}
```

### 3. Data Processing
```go
func processData(data []byte) ([]Item, error) {
    var items []Item
    if err := json.Unmarshal(data, &items); err != nil {
        return nil, err
    }
    
    // Process items
    sort.Slice(items, func(i, j int) bool {
        return items[i].Priority > items[j].Priority
    })
    
    return items, nil
}
```

## 📊 Learning Progress

### ✅ Completed Categories
- [x] Core Packages (8 packages)
- [x] Data Structures (4 packages)
- [x] Networking (3 packages)
- [x] Concurrency (3 packages)
- [x] Encoding (2 packages)
- [x] Utility (2 packages)
- [x] System (2 packages)

### 📈 Statistics
- **Total Packages**: 24 packages
- **Total Examples**: 200+ examples
- **Total Lines of Code**: 5000+ lines
- **Documentation**: Complete notes for each package
- **Test Coverage**: All examples tested and working

## 🎓 Learning Outcomes

After completing this project, you will:

1. **Master Go's Standard Library**: Understand all major packages
2. **Write Efficient Code**: Use best practices and performance patterns
3. **Handle Concurrency**: Build concurrent applications safely
4. **Process Data**: Work with JSON, XML, and other formats
5. **Build Systems**: Create system-level applications
6. **Debug and Profile**: Use runtime and reflection tools
7. **Optimize Performance**: Apply performance optimization techniques

## 🔧 Troubleshooting

### Common Issues

1. **Import Errors**: Make sure you're in the correct directory
2. **Build Errors**: Check Go version compatibility
3. **Runtime Errors**: Review error handling in examples
4. **Performance Issues**: Use profiling tools to identify bottlenecks

### Getting Help

1. Check the package-specific notes (`.md` files)
2. Review the example code (`.go` files)
3. Run examples to see expected output
4. Modify examples to experiment with different approaches

## 🤝 Contributing

This project is designed for learning, but contributions are welcome:

1. **Report Issues**: Found a bug or error? Let us know!
2. **Suggest Improvements**: Have ideas for better examples?
3. **Add Examples**: Want to add more use cases?
4. **Improve Documentation**: Help make the notes clearer

## 📚 Additional Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go Standard Library](https://pkg.go.dev/std)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)

## 🎉 Conclusion

This Go Standard Library Mastery project provides a comprehensive foundation for mastering Go's standard library. Through practical examples, detailed explanations, and real-world patterns, you'll gain the knowledge and skills needed to build robust, efficient, and maintainable Go applications.

**Happy Learning! 🚀**

---

*Remember: The best way to learn is by doing. Run the examples, modify them, experiment with different approaches, and build your own projects using these patterns!*