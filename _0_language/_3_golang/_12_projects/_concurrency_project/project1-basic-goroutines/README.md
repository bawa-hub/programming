# Project 1: Basic Goroutines & Channels

## 🎯 Learning Objectives
- Understand what goroutines are and how they work
- Learn channel communication patterns
- Master the `select` statement
- Handle basic synchronization
- Implement concurrent calculations

## 📚 Concepts Covered
- Goroutines (`go` keyword)
- Channels (buffered/unbuffered)
- Select statements
- Channel direction (send-only, receive-only)
- Basic synchronization patterns

## 🏗️ Project: Multi-threaded Calculator

Build a calculator that can perform multiple operations concurrently using goroutines and channels.

### Features to Implement:
1. **Basic Operations**: Add, subtract, multiply, divide
2. **Concurrent Execution**: All operations run in separate goroutines
3. **Channel Communication**: Results passed through channels
4. **Error Handling**: Division by zero and other errors
5. **Result Collection**: Collect and display all results

### Project Structure:
```
project1-basic-goroutines/
├── README.md
├── main.go
├── calculator.go
├── calculator_test.go
└── examples/
    ├── basic_operations.go
    ├── concurrent_calculations.go
    └── channel_patterns.go
```

## 🚀 Getting Started

1. **Start with `main.go`** - Basic goroutine creation
2. **Implement `calculator.go`** - Core calculation logic
3. **Add tests** - Verify your implementation
4. **Explore examples** - See different patterns

## 📝 Exercises

### Exercise 1: Basic Goroutines
Create a simple program that prints numbers 1-10 using goroutines.

### Exercise 2: Channel Communication
Implement a producer-consumer pattern with channels.

### Exercise 3: Select Statement
Use `select` to handle multiple channel operations with timeouts.

### Exercise 4: Calculator Implementation
Build the main calculator with concurrent operations.

## 🧪 Testing
Run tests with: `go test -v`
Run benchmarks with: `go test -bench=.`

## 🎯 Success Criteria
- [ ] Can create and manage goroutines
- [ ] Understands channel communication
- [ ] Can use select statements effectively
- [ ] Implements error handling in concurrent code
- [ ] All tests pass
- [ ] Code is clean and well-documented

## 🔗 Next Steps
After completing this project, move to **Project 2: Synchronization Primitives** to learn about mutexes, waitgroups, and more advanced synchronization.
