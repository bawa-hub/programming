# Decorator Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Decorator vs Adapter vs Bridge](#decorator-vs-adapter-vs-bridge)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Decorator pattern is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors. It allows you to add new functionality to objects dynamically without altering their structure.

## Problem Statement

**When to use Decorator Pattern?**
- When you need to add new functionality to objects at runtime
- When you want to add features to objects without modifying their classes
- When you need to combine multiple features in different ways
- When you want to avoid creating a large number of subclasses

**Common Scenarios:**
- Adding features to UI components (borders, scrolling, etc.)
- Adding functionality to streams (compression, encryption, etc.)
- Adding features to coffee orders (milk, sugar, etc.)
- Adding features to text (bold, italic, underline, etc.)
- Adding features to HTTP requests (authentication, logging, etc.)

## Solution

The Decorator pattern provides:
1. **Component Interface** - Defines the interface for objects that can have responsibilities added dynamically
2. **Concrete Component** - Defines an object to which additional responsibilities can be attached
3. **Decorator** - Maintains a reference to a Component object and defines an interface that conforms to Component's interface
4. **Concrete Decorator** - Adds responsibilities to the component

## Implementation Approaches

### 1. Basic Decorator
- Simple decorator that wraps a component
- Adds one specific behavior
- Easy to understand and implement

### 2. Multiple Decorators
- Multiple decorators can be chained together
- Each decorator adds a specific behavior
- More flexible but more complex

### 3. Decorator with State
- Decorators that maintain state
- More complex but more powerful
- Useful for features that need to remember information

## Decorator vs Adapter vs Bridge

### Decorator
- **Purpose**: Adds new functionality to objects
- **Focus**: Behavior extension
- **Use Case**: Adding features dynamically

### Adapter
- **Purpose**: Makes incompatible interfaces work together
- **Focus**: Interface conversion
- **Use Case**: Integrating existing code

### Bridge
- **Purpose**: Separates abstraction from implementation
- **Focus**: Structural separation
- **Use Case**: Independent evolution of abstractions and implementations

## Pros and Cons

### ✅ Pros
- **Flexibility**: Add new functionality at runtime
- **Single Responsibility**: Each decorator has a single responsibility
- **Open/Closed Principle**: Open for extension, closed for modification
- **Composition**: Can combine multiple decorators

### ❌ Cons
- **Complexity**: Can make code more complex
- **Many Classes**: May result in many decorator classes
- **Ordering**: Order of decorators can matter
- **Debugging**: Can be hard to debug with many decorators

## Real-world Examples

1. **UI Components**: Adding borders, scrolling, etc.
2. **Streams**: Adding compression, encryption, etc.
3. **Coffee Orders**: Adding milk, sugar, etc.
4. **Text Formatting**: Adding bold, italic, etc.
5. **HTTP Requests**: Adding authentication, logging, etc.

## Interview Questions

### Basic Level
1. What is the Decorator pattern?
2. When would you use Decorator pattern?
3. What is the difference between Decorator and Inheritance?

### Intermediate Level
1. How do you implement Decorator pattern?
2. What are the benefits of using Decorator pattern?
3. How do you handle multiple decorators?

### Advanced Level
1. How would you implement Decorator with generics?
2. How do you handle decorator ordering?
3. How would you implement Decorator pattern with state?

## Code Structure

```go
// Component interface
type Component interface {
    Operation() string
}

// Concrete Component
type ConcreteComponent struct{}

func (cc *ConcreteComponent) Operation() string {
    return "ConcreteComponent"
}

// Decorator
type Decorator struct {
    component Component
}

func (d *Decorator) Operation() string {
    return d.component.Operation()
}

// Concrete Decorator
type ConcreteDecorator struct {
    Decorator
}

func (cd *ConcreteDecorator) Operation() string {
    return "ConcreteDecorator(" + cd.Decorator.Operation() + ")"
}
```

## Next Steps

After mastering Decorator pattern, move to:
- **Facade Pattern** - For providing a simplified interface to a complex subsystem
- **Flyweight Pattern** - For sharing state between objects
- **Proxy Pattern** - For controlling access to objects

---

**Remember**: Decorator pattern is perfect for adding new functionality to objects at runtime. It's like adding layers to an object without changing its core structure! 🚀
