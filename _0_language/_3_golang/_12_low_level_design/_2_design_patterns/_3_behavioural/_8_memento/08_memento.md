# Memento Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Memento vs Command vs State](#memento-vs-command-vs-state)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Memento pattern is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation. It provides a way to capture and externalize an object's internal state.

## Problem Statement

**When to use Memento Pattern?**
- When you need to save and restore object state
- When you want to implement undo/redo functionality
- When you need to implement checkpoints in applications
- When you want to implement rollback functionality

**Common Scenarios:**
- Text editors (undo/redo)
- Games (save/load game state)
- Database transactions (rollback)
- Configuration management
- Version control systems

## Solution

The Memento pattern provides:
1. **Originator** - The object whose state needs to be saved
2. **Memento** - Stores the internal state of the originator
3. **Caretaker** - Manages and stores mementos
4. **Client** - Uses the caretaker to save and restore state

## Implementation Approaches

### 1. Basic Memento
- Simple memento with state storage
- Originator creates and restores mementos
- Easy to understand and implement

### 2. Memento with State Management
- Memento manages state transitions
- More robust state management
- Better for complex scenarios

### 3. Memento with History
- Memento maintains history of states
- More powerful but more complex
- Useful for complex undo/redo

## Memento vs Command vs State

### Memento
- **Purpose**: Saves and restores object state
- **Focus**: State preservation
- **Use Case**: Undo/redo functionality

### Command
- **Purpose**: Encapsulates requests as objects
- **Focus**: Request handling
- **Use Case**: Undo/redo functionality

### State
- **Purpose**: Changes object behavior based on internal state
- **Focus**: State management
- **Use Case**: Object behavior changes with state

## Pros and Cons

### ✅ Pros
- **State Preservation**: Saves and restores object state
- **Encapsulation**: Doesn't expose object's internal state
- **Undo/Redo**: Easy to implement undo/redo functionality
- **Checkpoints**: Provides checkpoint functionality

### ❌ Cons
- **Memory Usage**: Can consume significant memory
- **Complexity**: Can make code more complex
- **State Management**: Can be tricky to manage state
- **Performance**: May have performance overhead

## Real-world Examples

1. **Text Editors**: Undo/redo functionality
2. **Games**: Save/load game state
3. **Database Transactions**: Rollback functionality
4. **Configuration Management**: Save/restore settings
5. **Version Control**: Commit/rollback changes

## Interview Questions

### Basic Level
1. What is the Memento pattern?
2. When would you use Memento pattern?
3. What is the difference between Memento and Command pattern?

### Intermediate Level
1. How do you implement Memento pattern?
2. What are the benefits of using Memento pattern?
3. How do you handle memory management in Memento pattern?

### Advanced Level
1. How would you implement Memento with generics?
2. How do you handle complex state serialization?
3. How would you implement Memento pattern with persistence?

## Code Structure

```go
// Memento interface
type Memento interface {
    GetState() interface{}
}

// Originator
type Originator struct {
    state interface{}
}

func (o *Originator) CreateMemento() Memento {
    return &ConcreteMemento{state: o.state}
}

func (o *Originator) RestoreMemento(memento Memento) {
    o.state = memento.GetState()
}

// Concrete Memento
type ConcreteMemento struct {
    state interface{}
}

func (cm *ConcreteMemento) GetState() interface{} {
    return cm.state
}

// Caretaker
type Caretaker struct {
    mementos []Memento
}

func (c *Caretaker) AddMemento(memento Memento) {
    c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) Memento {
    return c.mementos[index]
}
```

## Next Steps

After mastering Memento pattern, move to:
- **Interpreter Pattern** - For defining grammar and interpreting expressions
- **Chain of Responsibility Pattern** - For handling requests
- **Iterator Pattern** - For traversing collections

---

**Remember**: Memento pattern is perfect for implementing undo/redo functionality and saving object state. It's like having a time machine for your objects! 🚀
