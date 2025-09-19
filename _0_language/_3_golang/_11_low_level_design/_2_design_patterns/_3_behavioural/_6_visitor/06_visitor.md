# Visitor Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Visitor vs Strategy vs Command](#visitor-vs-strategy-vs-command)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Visitor pattern is a behavioral design pattern that lets you define a new operation without changing the classes of the elements on which it operates. It separates the algorithm from the object structure.

## Problem Statement

**When to use Visitor Pattern?**
- When you need to add new operations to existing classes without modifying them
- When you have a stable object structure but frequently changing operations
- When you want to separate algorithms from object structure
- When you need to perform operations on a collection of different objects

**Common Scenarios:**
- Compiler design (AST traversal)
- Document processing (different document elements)
- Shopping cart calculations (different item types)
- Code analysis tools
- XML/JSON processing

## Solution

The Visitor pattern provides:
1. **Visitor Interface** - Declares the visit methods for each concrete element type
2. **Concrete Visitors** - Implement the visit methods for each element type
3. **Element Interface** - Defines the accept method that takes a visitor
4. **Concrete Elements** - Implement the accept method
5. **Object Structure** - Contains a collection of elements

## Implementation Approaches

### 1. Basic Visitor
- Simple visitor with visit methods for each element type
- Elements accept visitors
- Easy to understand and implement

### 2. Visitor with State
- Visitors can maintain state
- More powerful but more complex
- Useful for complex operations

### 3. Visitor with Return Values
- Visit methods can return values
- More flexible but more complex
- Useful for calculations and transformations

## Visitor vs Strategy vs Command

### Visitor
- **Purpose**: Adds new operations to existing classes
- **Focus**: Operation extension
- **Use Case**: Stable structure, changing operations

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
- **Open/Closed Principle**: Easy to add new operations
- **Single Responsibility**: Each visitor has one responsibility
- **Separation of Concerns**: Separates algorithms from structure
- **Flexibility**: Can add new operations without changing existing code

### ❌ Cons
- **Complexity**: Can make code more complex
- **Tight Coupling**: Visitors are tightly coupled to element types
- **Hard to Extend**: Adding new element types requires changing all visitors
- **Performance**: May have performance overhead

## Real-world Examples

1. **Compiler Design**: AST traversal for different operations
2. **Document Processing**: Different document elements
3. **Shopping Cart**: Different item types with different calculations
4. **Code Analysis**: Different code elements
5. **XML/JSON Processing**: Different node types

## Interview Questions

### Basic Level
1. What is the Visitor pattern?
2. When would you use Visitor pattern?
3. What is the difference between Visitor and Strategy pattern?

### Intermediate Level
1. How do you implement Visitor pattern?
2. What are the benefits of using Visitor pattern?
3. How do you handle adding new element types?

### Advanced Level
1. How would you implement Visitor with generics?
2. How do you handle complex visitor hierarchies?
3. How would you implement Visitor pattern with dependency injection?

## Code Structure

```go
// Visitor interface
type Visitor interface {
    VisitElementA(element *ElementA)
    VisitElementB(element *ElementB)
}

// Element interface
type Element interface {
    Accept(visitor Visitor)
}

// Concrete Element
type ElementA struct{}

func (ea *ElementA) Accept(visitor Visitor) {
    visitor.VisitElementA(ea)
}

// Concrete Visitor
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitElementA(element *ElementA) {
    // Visit ElementA
}

func (cv *ConcreteVisitor) VisitElementB(element *ElementB) {
    // Visit ElementB
}
```

## Next Steps

After mastering Visitor pattern, move to:
- **Mediator Pattern** - For encapsulating object interactions
- **Memento Pattern** - For saving and restoring object state
- **Interpreter Pattern** - For defining grammar and interpreting expressions

---

**Remember**: Visitor pattern is perfect for adding new operations to existing classes without modifying them. It's like having a tour guide for your object structure! 🚀
