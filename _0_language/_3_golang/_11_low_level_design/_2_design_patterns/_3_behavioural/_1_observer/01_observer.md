# Observer Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Observer vs Mediator vs Command](#observer-vs-mediator-vs-command)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Observer pattern is a behavioral design pattern that defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

## Problem Statement

**When to use Observer Pattern?**
- When you need to notify multiple objects about changes in another object
- When you want to decouple the subject from its observers
- When you need to implement event-driven systems
- When you want to implement publish-subscribe systems

**Common Scenarios:**
- Model-View-Controller (MVC) architecture
- Event handling systems
- Notification systems
- Stock price monitoring
- Weather station updates
- Social media feeds

## Solution

The Observer pattern provides:
1. **Subject Interface** - Defines the interface for attaching and detaching observers
2. **Concrete Subject** - Maintains a list of observers and notifies them of changes
3. **Observer Interface** - Defines the interface for objects that should be notified
4. **Concrete Observer** - Implements the observer interface and reacts to notifications

## Implementation Approaches

### 1. Basic Observer
- Simple subject-observer relationship
- Direct notification mechanism
- Easy to understand and implement

### 2. Observer with Events
- Uses event objects to pass data
- More flexible and extensible
- Better for complex scenarios

### 3. Observer with Filtering
- Observers can filter notifications
- More efficient for large systems
- Better performance

## Observer vs Mediator vs Command

### Observer
- **Purpose**: Notifies multiple objects about changes
- **Focus**: One-to-many communication
- **Use Case**: Event-driven systems

### Mediator
- **Purpose**: Encapsulates how objects interact
- **Focus**: Many-to-many communication
- **Use Case**: Complex object interactions

### Command
- **Purpose**: Encapsulates requests as objects
- **Focus**: Request handling
- **Use Case**: Undo/redo functionality

## Pros and Cons

### ✅ Pros
- **Loose Coupling**: Subject and observers are loosely coupled
- **Dynamic Relationships**: Can add/remove observers at runtime
- **Broadcast Communication**: One change can notify many observers
- **Open/Closed Principle**: Easy to add new observers

### ❌ Cons
- **Memory Leaks**: Can cause memory leaks if not properly managed
- **Performance**: Can be slow with many observers
- **Unexpected Updates**: Observers can be updated in unexpected order
- **Debugging**: Can be hard to debug with many observers

## Real-world Examples

1. **MVC Architecture**: Model notifies views of changes
2. **Event Systems**: GUI event handling
3. **Notification Systems**: Push notifications, email alerts
4. **Stock Monitoring**: Price change notifications
5. **Weather Updates**: Weather station data updates

## Interview Questions

### Basic Level
1. What is the Observer pattern?
2. When would you use Observer pattern?
3. What is the difference between Subject and Observer?

### Intermediate Level
1. How do you implement Observer pattern?
2. What are the benefits of using Observer pattern?
3. How do you handle memory leaks in Observer pattern?

### Advanced Level
1. How would you implement Observer with filtering?
2. How do you handle thread safety in Observer pattern?
3. How would you implement Observer pattern with generics?

## Code Structure

```go
// Observer interface
type Observer interface {
    Update(data interface{})
}

// Subject interface
type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify()
}

// Concrete Subject
type ConcreteSubject struct {
    observers []Observer
    state     interface{}
}

func (cs *ConcreteSubject) Attach(observer Observer) {
    cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) Notify() {
    for _, observer := range cs.observers {
        observer.Update(cs.state)
    }
}
```

## Next Steps

After mastering Observer pattern, move to:
- **Strategy Pattern** - For selecting algorithms at runtime
- **Command Pattern** - For encapsulating requests as objects
- **State Pattern** - For changing object behavior based on state

---

**Remember**: Observer pattern is perfect for implementing event-driven systems and maintaining loose coupling between objects. It's the foundation of many modern frameworks! 🚀
