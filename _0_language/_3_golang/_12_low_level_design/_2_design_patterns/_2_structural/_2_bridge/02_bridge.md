# Bridge Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Bridge vs Adapter vs Decorator](#bridge-vs-adapter-vs-decorator)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Bridge pattern is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

## Problem Statement

**When to use Bridge Pattern?**
- When you want to avoid a permanent binding between an abstraction and its implementation
- When you want to extend both abstractions and implementations independently
- When you want to share an implementation among multiple objects
- When you want to hide implementation details from clients

**Common Scenarios:**
- Different UI frameworks (Windows, Mac, Linux)
- Different database systems (MySQL, PostgreSQL, MongoDB)
- Different payment processors (Stripe, PayPal, Square)
- Different notification channels (Email, SMS, Push)

## Solution

The Bridge pattern provides:
1. **Abstraction** - Defines the interface for the control part
2. **Refined Abstraction** - Extends the abstraction with additional functionality
3. **Implementor** - Defines the interface for the implementation classes
4. **Concrete Implementor** - Implements the implementor interface

## Implementation Approaches

### 1. Basic Bridge
- Simple abstraction and implementation separation
- One abstraction, multiple implementations
- Clear separation of concerns

### 2. Multiple Abstractions
- Multiple abstractions using the same implementations
- More flexible but more complex
- Better for complex hierarchies

### 3. Dynamic Bridge
- Runtime switching of implementations
- More flexible but requires careful design
- Useful for plugin systems

## Bridge vs Adapter vs Decorator

### Bridge
- **Purpose**: Separates abstraction from implementation
- **Focus**: Structural separation
- **Use Case**: Independent evolution of abstractions and implementations

### Adapter
- **Purpose**: Makes incompatible interfaces work together
- **Focus**: Interface conversion
- **Use Case**: Integrating existing code

### Decorator
- **Purpose**: Adds new functionality to objects
- **Focus**: Behavior extension
- **Use Case**: Adding features dynamically

## Pros and Cons

### ✅ Pros
- **Separation of Concerns**: Separates abstraction from implementation
- **Independence**: Abstractions and implementations can evolve independently
- **Reusability**: Implementations can be shared among multiple abstractions
- **Open/Closed Principle**: Easy to add new abstractions and implementations

### ❌ Cons
- **Complexity**: Can make code more complex
- **Indirection**: Adds a layer of indirection
- **Many Classes**: May result in many classes
- **Understanding**: Can be hard to understand initially

## Real-world Examples

1. **UI Framework**: Different UI components for different platforms
2. **Database Abstraction**: Different database implementations
3. **Payment Processing**: Different payment processors
4. **Notification System**: Different notification channels
5. **File System**: Different file system implementations

## Interview Questions

### Basic Level
1. What is the Bridge pattern?
2. When would you use Bridge pattern?
3. What is the difference between Bridge and Adapter pattern?

### Intermediate Level
1. How do you implement Bridge pattern?
2. What are the benefits of using Bridge pattern?
3. How do you handle multiple implementations in Bridge pattern?

### Advanced Level
1. How would you implement a dynamic Bridge?
2. How do you handle state management in Bridge pattern?
3. How would you implement Bridge pattern with generics?

## Code Structure

```go
// Implementor interface
type Implementor interface {
    OperationImpl() string
}

// Concrete Implementor
type ConcreteImplementorA struct{}

func (cia *ConcreteImplementorA) OperationImpl() string {
    return "ConcreteImplementorA"
}

// Abstraction
type Abstraction struct {
    implementor Implementor
}

func (a *Abstraction) Operation() string {
    return a.implementor.OperationImpl()
}
```

## Next Steps

After mastering Bridge pattern, move to:
- **Composite Pattern** - For treating individual and composite objects uniformly
- **Decorator Pattern** - For adding new functionality to objects
- **Facade Pattern** - For providing a simplified interface to a complex subsystem

---

**Remember**: Bridge pattern is perfect for separating abstraction from implementation. It allows both to evolve independently and makes your code more flexible! 🚀
