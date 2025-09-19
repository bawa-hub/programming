# Adapter Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Adapter vs Decorator vs Facade](#adapter-vs-decorator-vs-facade)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Adapter pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It acts as a bridge between two incompatible interfaces by wrapping an existing class with a new interface.

## Problem Statement

**When to use Adapter Pattern?**
- When you need to use an existing class whose interface doesn't match what you need
- When you want to create a reusable class that cooperates with unrelated classes
- When you need to integrate third-party libraries with different interfaces
- When you want to make legacy code work with new code

**Common Scenarios:**
- Integrating third-party payment processors
- Using legacy database systems with new code
- Adapting different data formats (JSON, XML, CSV)
- Integrating different API versions
- Using different UI frameworks

## Solution

The Adapter pattern provides:
1. **Target Interface** - The interface that the client expects
2. **Adaptee** - The existing class that needs to be adapted
3. **Adapter** - The class that adapts the adaptee to the target interface
4. **Client** - Uses the target interface

## Implementation Approaches

### 1. Object Adapter
- Uses composition to wrap the adaptee
- More flexible and reusable
- Can adapt multiple adaptees

### 2. Class Adapter
- Uses inheritance to adapt the adaptee
- Less flexible but more direct
- Can only adapt one adaptee

### 3. Two-way Adapter
- Can adapt in both directions
- More complex but more flexible
- Useful for bidirectional communication

## Adapter vs Decorator vs Facade

### Adapter
- **Purpose**: Makes incompatible interfaces work together
- **Focus**: Interface conversion
- **Use Case**: Integrating existing code

### Decorator
- **Purpose**: Adds new functionality to objects
- **Focus**: Behavior extension
- **Use Case**: Adding features dynamically

### Facade
- **Purpose**: Provides a simplified interface to a complex subsystem
- **Focus**: Interface simplification
- **Use Case**: Hiding complexity

## Pros and Cons

### ✅ Pros
- **Interface Compatibility**: Makes incompatible interfaces work together
- **Reusability**: Can reuse existing code without modification
- **Single Responsibility**: Separates interface conversion from business logic
- **Open/Closed Principle**: Can add new adapters without changing existing code

### ❌ Cons
- **Complexity**: Can make code more complex
- **Performance Overhead**: Additional layer of indirection
- **Many Classes**: May result in many adapter classes
- **Tight Coupling**: Can create tight coupling between adaptee and adapter

## Real-world Examples

1. **Payment Gateway Integration**: Adapt different payment processors
2. **Database Abstraction**: Adapt different database systems
3. **API Integration**: Adapt different API versions
4. **File Format Conversion**: Adapt different file formats
5. **Legacy System Integration**: Adapt legacy systems to new interfaces

## Interview Questions

### Basic Level
1. What is the Adapter pattern?
2. When would you use Adapter pattern?
3. What is the difference between Object Adapter and Class Adapter?

### Intermediate Level
1. How do you implement an Adapter pattern?
2. What are the benefits of using Adapter pattern?
3. How do you handle multiple adaptees in an Adapter?

### Advanced Level
1. How would you implement a two-way Adapter?
2. How do you handle error cases in Adapter pattern?
3. How would you implement Adapter pattern with generics?

## Code Structure

```go
// Target interface
type Target interface {
    Request() string
}

// Adaptee
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
    return "Adaptee specific request"
}

// Adapter
type Adapter struct {
    adaptee *Adaptee
}

func (a *Adapter) Request() string {
    return a.adaptee.SpecificRequest()
}
```

## Next Steps

After mastering Adapter pattern, move to:
- **Bridge Pattern** - For separating abstraction from implementation
- **Composite Pattern** - For treating individual and composite objects uniformly
- **Decorator Pattern** - For adding new functionality to objects

---

**Remember**: Adapter pattern is perfect for integrating existing code with new interfaces. It's a bridge that makes incompatible systems work together! 🚀
