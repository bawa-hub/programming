# Strategy Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Strategy vs State vs Command](#strategy-vs-state-vs-command)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Strategy pattern is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. It lets the algorithm vary independently from clients that use it.

## Problem Statement

**When to use Strategy Pattern?**
- When you have multiple ways to perform a task
- When you want to switch algorithms at runtime
- When you want to avoid conditional statements
- When you want to make algorithms easily extensible

**Common Scenarios:**
- Payment processing (different payment methods)
- Sorting algorithms (different sorting strategies)
- Compression algorithms (different compression methods)
- Validation strategies (different validation rules)
- Pricing strategies (different pricing models)

## Solution

The Strategy pattern provides:
1. **Strategy Interface** - Defines the interface for all concrete strategies
2. **Concrete Strategies** - Implement the algorithm using the strategy interface
3. **Context** - Uses a strategy to call the concrete algorithm
4. **Client** - Creates and configures the context with a specific strategy

## Implementation Approaches

### 1. Basic Strategy
- Simple interface with one method
- Context holds a reference to strategy
- Easy to understand and implement

### 2. Strategy with Parameters
- Strategies can accept parameters
- More flexible but more complex
- Better for complex algorithms

### 3. Strategy with State
- Strategies can maintain state
- More powerful but more complex
- Useful for stateful algorithms

## Strategy vs State vs Command

### Strategy
- **Purpose**: Encapsulates algorithms and makes them interchangeable
- **Focus**: Algorithm selection
- **Use Case**: Multiple ways to perform a task

### State
- **Purpose**: Changes object behavior based on internal state
- **Focus**: State management
- **Use Case**: Object behavior changes with state

### Command
- **Purpose**: Encapsulates requests as objects
- **Focus**: Request handling
- **Use Case**: Undo/redo functionality

## Pros and Cons

### ✅ Pros
- **Algorithm Interchangeability**: Easy to switch algorithms at runtime
- **Open/Closed Principle**: Easy to add new strategies
- **Eliminates Conditionals**: Reduces if-else statements
- **Single Responsibility**: Each strategy has one responsibility

### ❌ Cons
- **Increased Complexity**: Can make code more complex
- **Many Classes**: May result in many strategy classes
- **Client Awareness**: Clients must be aware of different strategies
- **Communication Overhead**: Strategies may need to communicate with context

## Real-world Examples

1. **Payment Processing**: Different payment methods (credit card, PayPal, etc.)
2. **Sorting Algorithms**: Different sorting strategies (quick sort, merge sort, etc.)
3. **Compression**: Different compression algorithms (ZIP, RAR, etc.)
4. **Validation**: Different validation strategies (email, phone, etc.)
5. **Pricing**: Different pricing strategies (discount, premium, etc.)

## Interview Questions

### Basic Level
1. What is the Strategy pattern?
2. When would you use Strategy pattern?
3. What is the difference between Strategy and State pattern?

### Intermediate Level
1. How do you implement Strategy pattern?
2. What are the benefits of using Strategy pattern?
3. How do you handle strategy selection at runtime?

### Advanced Level
1. How would you implement Strategy with dependency injection?
2. How do you handle strategy communication with context?
3. How would you implement Strategy pattern with generics?

## Code Structure

```go
// Strategy interface
type Strategy interface {
    Execute(data interface{}) interface{}
}

// Concrete Strategies
type ConcreteStrategyA struct{}
type ConcreteStrategyB struct{}

func (csa *ConcreteStrategyA) Execute(data interface{}) interface{} {
    return "Strategy A result"
}

// Context
type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) interface{} {
    return c.strategy.Execute(data)
}
```

## Next Steps

After mastering Strategy pattern, move to:
- **Command Pattern** - For encapsulating requests as objects
- **State Pattern** - For changing object behavior based on state
- **Template Method Pattern** - For defining algorithm skeleton

---

**Remember**: Strategy pattern is perfect for encapsulating algorithms and making them interchangeable. It's like having a toolbox of different tools for the same job! 🚀
