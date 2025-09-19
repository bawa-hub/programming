# Template Method Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Template Method vs Strategy vs Factory](#template-method-vs-strategy-vs-factory)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Template Method pattern is a behavioral design pattern that defines the skeleton of an algorithm in a method, deferring some steps to subclasses. It lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

## Problem Statement

**When to use Template Method Pattern?**
- When you have an algorithm with invariant steps and variant steps
- When you want to avoid code duplication
- When you want to control the algorithm's structure
- When you want to provide a common interface for different implementations

**Common Scenarios:**
- Data processing pipelines
- Document generation
- Database operations
- Build processes
- Testing frameworks

## Solution

The Template Method pattern provides:
1. **Abstract Class** - Defines the template method and abstract methods
2. **Concrete Classes** - Implement the abstract methods
3. **Template Method** - Defines the algorithm structure
4. **Hook Methods** - Optional methods that can be overridden

## Implementation Approaches

### 1. Basic Template Method
- Simple template method with abstract steps
- Easy to understand and implement
- Good for simple algorithms

### 2. Template Method with Hooks
- Template method with optional hook methods
- More flexible but more complex
- Better for complex algorithms

### 3. Template Method with Parameters
- Template method that accepts parameters
- More flexible but more complex
- Better for parameterized algorithms

## Template Method vs Strategy vs Factory

### Template Method
- **Purpose**: Defines algorithm skeleton with customizable steps
- **Focus**: Algorithm structure
- **Use Case**: Common algorithm with variant steps

### Strategy
- **Purpose**: Encapsulates algorithms and makes them interchangeable
- **Focus**: Algorithm selection
- **Use Case**: Multiple ways to perform a task

### Factory
- **Purpose**: Creates objects without specifying their exact classes
- **Focus**: Object creation
- **Use Case**: Creating different types of objects

## Pros and Cons

### ✅ Pros
- **Code Reuse**: Eliminates code duplication
- **Control Structure**: Controls the algorithm's structure
- **Open/Closed Principle**: Easy to add new implementations
- **Consistent Interface**: Provides consistent interface

### ❌ Cons
- **Inheritance Dependency**: Relies on inheritance
- **Rigid Structure**: Can be rigid for complex algorithms
- **Debugging**: Can be hard to debug with many levels
- **Understanding**: Can be hard to understand initially

## Real-world Examples

1. **Data Processing**: ETL pipelines with different data sources
2. **Document Generation**: Different document formats
3. **Database Operations**: Different database types
4. **Build Processes**: Different build configurations
5. **Testing Frameworks**: Different test types

## Interview Questions

### Basic Level
1. What is the Template Method pattern?
2. When would you use Template Method pattern?
3. What is the difference between Template Method and Strategy pattern?

### Intermediate Level
1. How do you implement Template Method pattern?
2. What are the benefits of using Template Method pattern?
3. How do you handle hook methods in Template Method pattern?

### Advanced Level
1. How would you implement Template Method with generics?
2. How do you handle complex algorithms with Template Method?
3. How would you implement Template Method with dependency injection?

## Code Structure

```go
// Abstract class
type AbstractClass struct{}

func (ac *AbstractClass) TemplateMethod() {
    ac.Step1()
    ac.Step2()
    ac.Step3()
}

func (ac *AbstractClass) Step1() {
    // Default implementation
}

func (ac *AbstractClass) Step2() {
    // Abstract method - must be implemented
}

func (ac *AbstractClass) Step3() {
    // Default implementation
}

// Concrete class
type ConcreteClass struct {
    AbstractClass
}

func (cc *ConcreteClass) Step2() {
    // Concrete implementation
}
```

## Next Steps

After mastering Template Method pattern, move to:
- **Visitor Pattern** - For adding operations to objects
- **Mediator Pattern** - For encapsulating object interactions
- **Memento Pattern** - For saving and restoring object state

---

**Remember**: Template Method pattern is perfect for defining algorithm skeletons with customizable steps. It's like having a recipe with some steps that can be customized! 🚀
