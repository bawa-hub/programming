# State Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [State vs Strategy vs Command](#state-vs-strategy-vs-command)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The State pattern is a behavioral design pattern that allows an object to alter its behavior when its internal state changes. The object will appear to change its class.

## Problem Statement

**When to use State Pattern?**
- When an object's behavior depends on its state
- When you have many conditional statements based on object state
- When you want to avoid large if-else or switch statements
- When you need to add new states easily

**Common Scenarios:**
- Vending machines (idle, processing, dispensing)
- Media players (playing, paused, stopped)
- Order processing (pending, confirmed, shipped, delivered)
- Game characters (idle, running, jumping, attacking)
- Traffic lights (red, yellow, green)

## Solution

The State pattern provides:
1. **State Interface** - Defines the interface for all concrete states
2. **Concrete States** - Implement the behavior associated with a particular state
3. **Context** - Maintains a reference to the current state and delegates state-specific behavior
4. **Client** - Uses the context to interact with the state machine

## Implementation Approaches

### 1. Basic State
- Simple state interface with one method
- Context holds current state
- Easy to understand and implement

### 2. State with Transitions
- States define valid transitions
- More robust state management
- Better for complex state machines

### 3. State with Actions
- States can perform actions on entry/exit
- More powerful but more complex
- Useful for complex state machines

## State vs Strategy vs Command

### State
- **Purpose**: Changes object behavior based on internal state
- **Focus**: State management
- **Use Case**: Object behavior changes with state

### Strategy
- **Purpose**: Encapsulates algorithms and makes them interchangeable
- **Focus**: Algorithm selection
- **Use Case**: Multiple ways to perform a task

### Command
- **Purpose**: Encapsulates requests as objects
- **Focus**: Request handling
- **Use Case**: Undo/redo functionality

## Pros and Cons

### ✅ Pros
- **Eliminates Conditionals**: Reduces if-else statements
- **Open/Closed Principle**: Easy to add new states
- **Single Responsibility**: Each state has one responsibility
- **State Transitions**: Clear state transition logic

### ❌ Cons
- **Complexity**: Can make code more complex
- **Many Classes**: May result in many state classes
- **State Management**: Can be tricky to manage state transitions
- **Memory Usage**: States can consume memory

## Real-world Examples

1. **Vending Machines**: Different states for different operations
2. **Media Players**: Play, pause, stop states
3. **Order Processing**: Different states for order lifecycle
4. **Game Characters**: Different states for character behavior
5. **Traffic Lights**: Different states for traffic control

## Interview Questions

### Basic Level
1. What is the State pattern?
2. When would you use State pattern?
3. What is the difference between State and Strategy pattern?

### Intermediate Level
1. How do you implement State pattern?
2. What are the benefits of using State pattern?
3. How do you handle state transitions?

### Advanced Level
1. How would you implement State with actions?
2. How do you handle state persistence?
3. How would you implement State pattern with generics?

## Code Structure

```go
// State interface
type State interface {
    Handle(context *Context)
    GetName() string
}

// Concrete States
type ConcreteStateA struct{}
type ConcreteStateB struct{}

func (csa *ConcreteStateA) Handle(context *Context) {
    // Handle state A
}

// Context
type Context struct {
    state State
}

func (c *Context) SetState(state State) {
    c.state = state
}

func (c *Context) Request() {
    c.state.Handle(c)
}
```

## Next Steps

After mastering State pattern, move to:
- **Template Method Pattern** - For defining algorithm skeleton
- **Visitor Pattern** - For adding operations to objects
- **Mediator Pattern** - For encapsulating object interactions

---

**Remember**: State pattern is perfect for managing object behavior that changes based on internal state. It's like having different personalities for the same object! 🚀
