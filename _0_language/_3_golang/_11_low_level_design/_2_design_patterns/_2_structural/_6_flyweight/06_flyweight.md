# Flyweight Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Flyweight vs Singleton vs Prototype](#flyweight-vs-singleton-vs-prototype)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Flyweight pattern is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

## Problem Statement

**When to use Flyweight Pattern?**
- When you need to create a large number of similar objects
- When the cost of creating objects is high
- When you want to reduce memory usage
- When you can separate intrinsic and extrinsic state

**Common Scenarios:**
- Text editors (character formatting)
- Game development (trees, bullets, particles)
- GUI applications (icons, fonts)
- Database connections (connection pooling)
- Web applications (session management)

## Solution

The Flyweight pattern provides:
1. **Flyweight Interface** - Defines the interface through which flyweights can receive and act on extrinsic state
2. **Concrete Flyweight** - Implements the flyweight interface and stores intrinsic state
3. **Flyweight Factory** - Creates and manages flyweight objects
4. **Client** - Maintains references to flyweights and computes or stores extrinsic state

## Implementation Approaches

### 1. Basic Flyweight
- Simple flyweight with intrinsic state
- Factory manages flyweight instances
- Client maintains extrinsic state

### 2. Flyweight with State
- Flyweight with both intrinsic and extrinsic state
- More complex but more flexible
- Better for complex scenarios

### 3. Flyweight with Pooling
- Object pooling for flyweights
- More efficient memory management
- Better for high-performance scenarios

## Flyweight vs Singleton vs Prototype

### Flyweight
- **Purpose**: Shares state between objects to reduce memory usage
- **Focus**: Memory optimization
- **Use Case**: Large number of similar objects

### Singleton
- **Purpose**: Ensures only one instance exists
- **Focus**: Instance control
- **Use Case**: Single instance needed

### Prototype
- **Purpose**: Creates objects by cloning existing instances
- **Focus**: Object creation
- **Use Case**: Expensive object creation

## Pros and Cons

### ✅ Pros
- **Memory Efficiency**: Reduces memory usage by sharing state
- **Performance**: Improves performance by reducing object creation
- **Scalability**: Allows handling large numbers of objects
- **Flexibility**: Can separate intrinsic and extrinsic state

### ❌ Cons
- **Complexity**: Can make code more complex
- **State Management**: Can be tricky to manage state
- **Debugging**: Can be hard to debug with shared state
- **Thread Safety**: Requires careful handling in multi-threaded environments

## Real-world Examples

1. **Text Editors**: Character formatting (bold, italic, etc.)
2. **Game Development**: Trees, bullets, particles
3. **GUI Applications**: Icons, fonts, colors
4. **Database Connections**: Connection pooling
5. **Web Applications**: Session management

## Interview Questions

### Basic Level
1. What is the Flyweight pattern?
2. When would you use Flyweight pattern?
3. What is the difference between intrinsic and extrinsic state?

### Intermediate Level
1. How do you implement Flyweight pattern?
2. What are the benefits of using Flyweight pattern?
3. How do you handle state management in Flyweight pattern?

### Advanced Level
1. How would you implement Flyweight with object pooling?
2. How do you handle thread safety in Flyweight pattern?
3. How would you implement Flyweight pattern with generics?

## Code Structure

```go
// Flyweight interface
type Flyweight interface {
    Operation(extrinsicState string)
}

// Concrete Flyweight
type ConcreteFlyweight struct {
    intrinsicState string
}

func (cf *ConcreteFlyweight) Operation(extrinsicState string) {
    // Use both intrinsic and extrinsic state
}

// Flyweight Factory
type FlyweightFactory struct {
    flyweights map[string]Flyweight
}

func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
    if flyweight, exists := ff.flyweights[key]; exists {
        return flyweight
    }
    // Create new flyweight
    return ff.createFlyweight(key)
}
```

## Next Steps

After mastering Flyweight pattern, move to:
- **Proxy Pattern** - For controlling access to objects
- **Behavioral Patterns** - For communication between objects
- **Observer Pattern** - For notifying multiple objects about changes

---

**Remember**: Flyweight pattern is perfect for optimizing memory usage when you have many similar objects. It's like sharing a template and only storing the differences! 🚀
