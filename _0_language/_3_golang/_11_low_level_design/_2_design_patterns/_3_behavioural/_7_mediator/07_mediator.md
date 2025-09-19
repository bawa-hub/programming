# Mediator Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Mediator vs Observer vs Facade](#mediator-vs-observer-vs-facade)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Mediator pattern is a behavioral design pattern that defines how a set of objects interact. It promotes loose coupling by keeping objects from referring to each other explicitly and lets you vary their interaction independently.

## Problem Statement

**When to use Mediator Pattern?**
- When you have a set of objects that communicate in complex ways
- When you want to reduce coupling between objects
- When you want to centralize control over object interactions
- When you want to make object interactions more maintainable

**Common Scenarios:**
- Chat rooms (users communicate through mediator)
- Air traffic control (aircraft communicate through controller)
- GUI components (components communicate through mediator)
- Event systems (events are routed through mediator)
- Microservices communication

## Solution

The Mediator pattern provides:
1. **Mediator Interface** - Defines the interface for communicating with colleague objects
2. **Concrete Mediator** - Implements cooperative behavior by coordinating colleague objects
3. **Colleague Classes** - Know their mediator and communicate with other colleagues through it
4. **Client** - Uses the mediator to coordinate colleague objects

## Implementation Approaches

### 1. Basic Mediator
- Simple mediator with direct communication
- Colleagues communicate through mediator
- Easy to understand and implement

### 2. Mediator with Events
- Uses event system for communication
- More flexible but more complex
- Better for complex interactions

### 3. Mediator with State
- Mediator maintains state
- More powerful but more complex
- Useful for complex state management

## Mediator vs Observer vs Facade

### Mediator
- **Purpose**: Encapsulates how objects interact
- **Focus**: Many-to-many communication
- **Use Case**: Complex object interactions

### Observer
- **Purpose**: Notifies multiple objects about changes
- **Focus**: One-to-many communication
- **Use Case**: Event-driven systems

### Facade
- **Purpose**: Provides a simplified interface to a complex subsystem
- **Focus**: Interface simplification
- **Use Case**: Hiding complexity

## Pros and Cons

### ✅ Pros
- **Loose Coupling**: Reduces coupling between objects
- **Centralized Control**: Centralizes control over object interactions
- **Easy to Extend**: Easy to add new colleagues
- **Single Responsibility**: Mediator has one responsibility

### ❌ Cons
- **Complexity**: Can make code more complex
- **Single Point of Failure**: Mediator can become a bottleneck
- **God Object**: Mediator can become too complex
- **Performance**: May have performance overhead

## Real-world Examples

1. **Chat Rooms**: Users communicate through chat room mediator
2. **Air Traffic Control**: Aircraft communicate through controller
3. **GUI Components**: Components communicate through mediator
4. **Event Systems**: Events are routed through mediator
5. **Microservices**: Services communicate through API gateway

## Interview Questions

### Basic Level
1. What is the Mediator pattern?
2. When would you use Mediator pattern?
3. What is the difference between Mediator and Observer pattern?

### Intermediate Level
1. How do you implement Mediator pattern?
2. What are the benefits of using Mediator pattern?
3. How do you handle complex interactions in Mediator pattern?

### Advanced Level
1. How would you implement Mediator with events?
2. How do you handle mediator scalability?
3. How would you implement Mediator pattern with generics?

## Code Structure

```go
// Mediator interface
type Mediator interface {
    Notify(sender Colleague, event string)
}

// Colleague interface
type Colleague interface {
    SetMediator(mediator Mediator)
    Notify(event string)
}

// Concrete Mediator
type ConcreteMediator struct {
    colleagues []Colleague
}

func (cm *ConcreteMediator) Notify(sender Colleague, event string) {
    for _, colleague := range cm.colleagues {
        if colleague != sender {
            colleague.Notify(event)
        }
    }
}
```

## Next Steps

After mastering Mediator pattern, move to:
- **Memento Pattern** - For saving and restoring object state
- **Interpreter Pattern** - For defining grammar and interpreting expressions
- **Chain of Responsibility Pattern** - For handling requests

---

**Remember**: Mediator pattern is perfect for managing complex object interactions. It's like having a traffic controller for your objects! 🚀
