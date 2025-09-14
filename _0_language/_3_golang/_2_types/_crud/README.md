# 🎯 Golang Data Types Mastery - CRUD Application

A comprehensive Go application that demonstrates mastery of all Go data types through practical CRUD operations. This project serves as a complete learning resource for understanding Go's type system, from basic primitives to advanced generics.

## 🚀 Features

### Data Types Covered

1. **Primitive Types** - All basic Go types with comprehensive examples
2. **Arrays & Slices** - Dynamic arrays with advanced operations
3. **Strings** - Text processing and manipulation
4. **Structs** - Data modeling and composition
5. **Interfaces** - Polymorphism and abstraction
6. **Maps** - Key-value storage and operations
7. **Pointers** - Memory management and references
8. **Generics** - Type parameters and constraints

### CRUD Operations

Each data type includes complete CRUD (Create, Read, Update, Delete) operations:

- **Create** - Initialize and populate data structures
- **Read** - Display and analyze data
- **Update** - Modify existing data
- **Delete** - Remove data and clean up

### Advanced Features

- **Type Safety** - Compile-time type checking
- **Memory Management** - Efficient memory usage patterns
- **Performance Analysis** - Benchmarking and optimization
- **Error Handling** - Robust error management
- **Concurrency** - Goroutines and channels
- **Reflection** - Runtime type inspection
- **Validation** - Data validation patterns

## 📁 Project Structure

```
golang-crud-mastery/
├── main.go              # Application entry point
├── app.go               # Main application logic
├── primitives.go        # Primitive types examples
├── arrays_slices.go     # Arrays and slices examples
├── strings.go           # String manipulation examples
├── structs.go           # Struct and data modeling examples
├── interfaces.go        # Interface and polymorphism examples
├── maps.go              # Map and key-value examples
├── pointers.go          # Pointer and memory management examples
├── generics.go          # Generic types and functions examples
├── go.mod               # Go module definition
└── README.md            # This file
```

## 🛠️ Installation

### Prerequisites

- Go 1.21 or later
- Git

### Setup

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd golang-crud-mastery
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run .
   ```

## 🎮 Usage

### Interactive Menu

The application provides an interactive menu system:

```
🎯 GOLANG DATA TYPES MASTERY - CRUD APPLICATION
================================================
Choose a data type to explore:

1.  Primitive Types (int, float, bool, string, rune, byte)
2.  Arrays & Slices
3.  Strings & Text Processing
4.  Structs & Data Modeling
5.  Interfaces & Polymorphism
6.  Maps & Key-Value Storage
7.  Pointers & Memory Management
8.  Generics & Type Parameters
9.  Advanced Demonstrations
10. Run All Examples
11. Exit
```

### Data Type Menus

Each data type has its own submenu with specific operations:

- **CRUD Operations** - Create, Read, Update, Delete
- **Advanced Features** - Type-specific demonstrations
- **Performance Analysis** - Benchmarking and optimization
- **Best Practices** - Recommended patterns and techniques

## 📚 Learning Path

### 1. Primitive Types

Learn about Go's basic data types:

- **Integer Types**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Types**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floating Point**: `float32`, `float64`
- **Complex Numbers**: `complex64`, `complex128`
- **Boolean**: `bool`
- **String**: `string`
- **Character Types**: `rune`, `byte`

**Key Concepts:**
- Type conversion
- Zero values
- Constants and iota
- Type inference

### 2. Arrays & Slices

Master dynamic arrays and collections:

- **Arrays**: Fixed-size collections
- **Slices**: Dynamic arrays
- **2D Arrays**: Multi-dimensional data
- **Slice Operations**: Append, copy, filter, map, reduce

**Key Concepts:**
- Memory layout
- Capacity and length
- Slice expressions
- Slice tricks and patterns

### 3. Strings

Comprehensive text processing:

- **String Creation**: Literals, raw strings, concatenation
- **String Operations**: Search, replace, split, join
- **Regular Expressions**: Pattern matching and validation
- **String Formatting**: Printf, Sprintf, string builders

**Key Concepts:**
- Immutability
- Unicode support
- String conversion
- Performance optimization

### 4. Structs

Data modeling and composition:

- **Basic Structs**: Fields, methods, tags
- **Embedded Structs**: Composition and inheritance
- **Anonymous Structs**: Inline data structures
- **Struct Methods**: Value and pointer receivers

**Key Concepts:**
- Encapsulation
- Method promotion
- JSON serialization
- Struct validation

### 5. Interfaces

Polymorphism and abstraction:

- **Basic Interfaces**: Method sets
- **Interface Composition**: Combining interfaces
- **Type Assertions**: Runtime type checking
- **Empty Interface**: `interface{}`

**Key Concepts:**
- Implicit implementation
- Interface satisfaction
- Type switches
- Interface values

### 6. Maps

Key-value storage and operations:

- **Basic Maps**: Creation, access, modification
- **Map Operations**: Iteration, filtering, transformation
- **Nested Maps**: Complex data structures
- **Map Patterns**: Caching, counting, grouping

**Key Concepts:**
- Hash tables
- Key constraints
- Map iteration order
- Memory efficiency

### 7. Pointers

Memory management and references:

- **Basic Pointers**: Addresses and dereferencing
- **Pointer Arithmetic**: Limited in Go
- **Pointer Passing**: Value vs reference semantics
- **Pointer Safety**: Avoiding common pitfalls

**Key Concepts:**
- Memory addresses
- Reference semantics
- Pointer receivers
- Garbage collection

### 8. Generics

Type parameters and constraints:

- **Generic Functions**: Type-parameterized functions
- **Generic Types**: Type-parameterized structs
- **Type Constraints**: Interface constraints
- **Type Inference**: Automatic type deduction

**Key Concepts:**
- Type parameters
- Constraint interfaces
- Type erasure
- Performance implications

## 🔧 Advanced Features

### Performance Analysis

- **Benchmarking**: Compare different approaches
- **Memory Profiling**: Analyze memory usage
- **CPU Profiling**: Identify performance bottlenecks
- **Optimization Techniques**: Best practices for performance

### Error Handling

- **Error Types**: Custom error types
- **Error Wrapping**: Error context and chaining
- **Panic and Recover**: Exception handling
- **Error Patterns**: Common error handling strategies

### Concurrency

- **Goroutines**: Lightweight concurrency
- **Channels**: Communication between goroutines
- **Select Statements**: Non-blocking operations
- **Synchronization**: Mutexes, wait groups, and atomic operations

### Testing

- **Unit Tests**: Test individual functions
- **Integration Tests**: Test component interactions
- **Benchmark Tests**: Performance testing
- **Table-Driven Tests**: Parameterized testing

## 📖 Examples

### Primitive Types Example

```go
// Create primitive values
var intVal int = 42
var floatVal float64 = 3.14159
var stringVal string = "Hello, World!"
var boolVal bool = true

// Type conversion
floatFromInt := float64(intVal)
intFromFloat := int(floatVal)

// String operations
concatenated := stringVal + " " + strconv.Itoa(intVal)
```

### Arrays & Slices Example

```go
// Create slice
numbers := []int{1, 2, 3, 4, 5}

// Append elements
numbers = append(numbers, 6, 7, 8)

// Filter even numbers
evenNumbers := filter(numbers, func(x int) bool {
    return x%2 == 0
})

// Map to squares
squares := mapSlice(numbers, func(x int) int {
    return x * x
})
```

### Structs Example

```go
type Person struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// Create instance
person := Person{
    ID:    1,
    Name:  "Alice Johnson",
    Email: "alice@example.com",
    Age:   30,
}

// Method
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}
```

### Interfaces Example

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

### Maps Example

```go
// Create map
scores := make(map[string]int)
scores["Alice"] = 95
scores["Bob"] = 87
scores["Charlie"] = 92

// Iterate over map
for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}

// Check if key exists
if score, exists := scores["Alice"]; exists {
    fmt.Printf("Alice's score: %d\n", score)
}
```

### Pointers Example

```go
// Create pointer
value := 42
ptr := &value

// Dereference pointer
fmt.Printf("Value: %d\n", *ptr)

// Modify through pointer
*ptr = 100
fmt.Printf("New value: %d\n", value)
```

### Generics Example

```go
// Generic function
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Generic type
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}
```

## 🎯 Learning Objectives

After completing this project, you will:

1. **Master Go's Type System** - Understand all data types and their use cases
2. **Implement CRUD Operations** - Build complete data management systems
3. **Apply Best Practices** - Follow Go idioms and conventions
4. **Handle Errors Effectively** - Implement robust error handling
5. **Optimize Performance** - Write efficient and scalable code
6. **Use Concurrency** - Leverage goroutines and channels
7. **Write Generic Code** - Create reusable type-parameterized functions
8. **Test Thoroughly** - Implement comprehensive testing strategies

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Go team for creating an amazing language
- Go community for excellent documentation and resources
- Contributors who helped improve this project

## 📞 Support

If you have any questions or need help, please:

1. Check the documentation
2. Search existing issues
3. Create a new issue
4. Contact the maintainers

---

**Happy Coding! 🚀**

*Master Go's data types and build amazing applications!*
