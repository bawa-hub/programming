# Composite Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Composite vs Decorator vs Bridge](#composite-vs-decorator-vs-bridge)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Composite pattern is a structural design pattern that lets you compose objects into tree structures to represent part-whole hierarchies. It allows clients to treat individual objects and compositions of objects uniformly.

## Problem Statement

**When to use Composite Pattern?**
- When you need to represent part-whole hierarchies
- When you want clients to treat individual and composite objects uniformly
- When you need to build tree-like structures
- When you want to add new types of components easily

**Common Scenarios:**
- File system structures (files and folders)
- UI component hierarchies (buttons, panels, windows)
- Organization structures (departments, employees)
- Mathematical expressions (operators, operands)
- Menu systems (menus, menu items)

## Solution

The Composite pattern provides:
1. **Component Interface** - Defines common operations for both leaf and composite objects
2. **Leaf** - Represents individual objects that don't have children
3. **Composite** - Represents objects that have children
4. **Client** - Uses the component interface to work with objects

## Implementation Approaches

### 1. Basic Composite
- Simple interface with common operations
- Leaf and Composite implement the interface
- Client treats them uniformly

### 2. Composite with Operations
- Additional operations for composite-specific functionality
- Add, remove, get children operations
- More complex but more flexible

### 3. Composite with Visitor
- Visitor pattern for operations on composite structures
- More flexible for complex operations
- Better separation of concerns

## Composite vs Decorator vs Bridge

### Composite
- **Purpose**: Treats individual and composite objects uniformly
- **Focus**: Tree structure representation
- **Use Case**: Part-whole hierarchies

### Decorator
- **Purpose**: Adds new functionality to objects
- **Focus**: Behavior extension
- **Use Case**: Adding features dynamically

### Bridge
- **Purpose**: Separates abstraction from implementation
- **Focus**: Structural separation
- **Use Case**: Independent evolution of abstractions and implementations

## Pros and Cons

### ✅ Pros
- **Uniform Treatment**: Clients treat individual and composite objects uniformly
- **Easy to Add**: Easy to add new types of components
- **Flexible Structure**: Can build complex tree structures
- **Recursive Operations**: Natural for recursive operations

### ❌ Cons
- **Complexity**: Can make the design overly general
- **Type Safety**: Can lose type safety
- **Performance**: May have performance overhead
- **Understanding**: Can be hard to understand initially

## Real-world Examples

1. **File System**: Files and folders
2. **UI Components**: Buttons, panels, windows
3. **Organization Structure**: Departments, employees
4. **Mathematical Expressions**: Operators, operands
5. **Menu Systems**: Menus, menu items

## Interview Questions

### Basic Level
1. What is the Composite pattern?
2. When would you use Composite pattern?
3. What is the difference between Leaf and Composite?

### Intermediate Level
1. How do you implement Composite pattern?
2. What are the benefits of using Composite pattern?
3. How do you handle operations on composite structures?

### Advanced Level
1. How would you implement Composite with Visitor pattern?
2. How do you handle type safety in Composite pattern?
3. How would you implement Composite pattern with generics?

## Code Structure

```go
// Component interface
type Component interface {
    Operation() string
    Add(component Component)
    Remove(component Component)
    GetChild(index int) Component
}

// Leaf
type Leaf struct {
    name string
}

func (l *Leaf) Operation() string {
    return l.name
}

// Composite
type Composite struct {
    name     string
    children []Component
}

func (c *Composite) Operation() string {
    result := c.name + " ["
    for i, child := range c.children {
        if i > 0 {
            result += ", "
        }
        result += child.Operation()
    }
    result += "]"
    return result
}
```

## Next Steps

After mastering Composite pattern, move to:
- **Decorator Pattern** - For adding new functionality to objects
- **Facade Pattern** - For providing a simplified interface to a complex subsystem
- **Flyweight Pattern** - For sharing state between objects

---

**Remember**: Composite pattern is perfect for representing tree structures where you want to treat individual and composite objects uniformly. It's the foundation of many UI frameworks! 🚀
