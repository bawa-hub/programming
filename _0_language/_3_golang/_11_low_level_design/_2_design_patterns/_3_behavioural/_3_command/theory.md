# Command Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Command vs Strategy vs Observer](#command-vs-strategy-vs-observer)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Command pattern is a behavioral design pattern that encapsulates a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

## Problem Statement

**When to use Command Pattern?**
- When you need to parameterize objects with operations
- When you want to queue, log, or support undo operations
- When you need to support macro operations (composite commands)
- When you want to decouple the object that invokes the operation from the one that performs it

**Common Scenarios:**
- Undo/Redo functionality in text editors
- Macro recording and playback
- Queuing and logging requests
- Remote procedure calls
- GUI button actions

## Solution

The Command pattern provides:
1. **Command Interface** - Declares an interface for executing operations
2. **Concrete Commands** - Implement the command interface and bind to receiver objects
3. **Receiver** - Knows how to perform the operations associated with a request
4. **Invoker** - Asks the command to carry out the request
5. **Client** - Creates concrete command objects and sets their receivers

## Implementation Approaches

### 1. Basic Command
- Simple command with execute method
- Commands are stateless
- Easy to understand and implement

### 2. Command with Undo
- Commands support undo operations
- Commands maintain state
- More complex but more powerful

### 3. Macro Commands
- Commands can be composed
- Support for complex operations
- Better for complex scenarios

## Command vs Strategy vs Observer

### Command
- **Purpose**: Encapsulates requests as objects
- **Focus**: Request handling and queuing
- **Use Case**: Undo/redo functionality

### Strategy
- **Purpose**: Encapsulates algorithms and makes them interchangeable
- **Focus**: Algorithm selection
- **Use Case**: Multiple ways to perform a task

### Observer
- **Purpose**: Notifies multiple objects about changes
- **Focus**: Event notification
- **Use Case**: Event-driven systems

## Pros and Cons

### ✅ Pros
- **Decoupling**: Decouples the object that invokes the operation from the one that performs it
- **Undo/Redo**: Easy to implement undo and redo functionality
- **Queuing**: Can queue, log, and schedule requests
- **Macro Commands**: Can compose commands into macro commands

### ❌ Cons
- **Complexity**: Can make code more complex
- **Memory Usage**: Commands can consume memory
- **Overhead**: May add overhead for simple operations
- **State Management**: Can be tricky to manage command state

## Real-world Examples

1. **Text Editors**: Undo/redo functionality
2. **Macro Recording**: Recording and playing back operations
3. **Remote Controls**: Button actions
4. **Database Transactions**: Transaction commands
5. **GUI Applications**: Menu and button actions

## Interview Questions

### Basic Level
1. What is the Command pattern?
2. When would you use Command pattern?
3. What is the difference between Command and Strategy pattern?

### Intermediate Level
1. How do you implement Command pattern?
2. What are the benefits of using Command pattern?
3. How do you implement undo functionality in Command pattern?

### Advanced Level
1. How would you implement macro commands?
2. How do you handle command queuing and logging?
3. How would you implement Command pattern with generics?

## Code Structure

```go
// Command interface
type Command interface {
    Execute()
    Undo()
}

// Concrete Command
type ConcreteCommand struct {
    receiver *Receiver
    state    interface{}
}

func (cc *ConcreteCommand) Execute() {
    cc.receiver.Action(cc.state)
}

func (cc *ConcreteCommand) Undo() {
    cc.receiver.UndoAction(cc.state)
}

// Invoker
type Invoker struct {
    commands []Command
}

func (i *Invoker) ExecuteCommand(command Command) {
    command.Execute()
    i.commands = append(i.commands, command)
}
```

## Next Steps

After mastering Command pattern, move to:
- **State Pattern** - For changing object behavior based on state
- **Template Method Pattern** - For defining algorithm skeleton
- **Visitor Pattern** - For adding operations to objects

---

**Remember**: Command pattern is perfect for implementing undo/redo functionality and decoupling request senders from receivers. It's like having a remote control for your objects! 🚀
