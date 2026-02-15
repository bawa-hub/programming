# Prototype Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Deep vs Shallow Copy](#deep-vs-shallow-copy)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Prototype pattern is a creational design pattern that lets you copy existing objects without making your code dependent on their classes. It's particularly useful when the cost of creating an object is expensive or when you need to create many similar objects.

## Problem Statement

**When to use Prototype Pattern?**
- When you need to create objects that are expensive to create
- When you want to avoid subclasses of an object creator
- When you need to create objects at runtime
- When you want to create objects that are similar to existing ones

**Common Scenarios:**
- Creating database records
- Creating UI components
- Creating game objects
- Creating configuration objects
- Creating document templates

## Solution

The Prototype pattern provides:
1. **Prototype Interface** - Defines the cloning method
2. **Concrete Prototypes** - Implement the cloning method
3. **Prototype Registry** - Manages available prototypes (optional)
4. **Client** - Uses prototypes to create new objects

## Implementation Approaches

### 1. Basic Prototype
- Simple interface with Clone method
- Concrete classes implement cloning
- Client creates objects by cloning

### 2. Prototype Registry
- Registry manages available prototypes
- Client requests prototypes by name
- More flexible and organized

### 3. Deep vs Shallow Copy
- Shallow copy: Copies object references
- Deep copy: Copies all nested objects
- Choose based on requirements

## Deep vs Shallow Copy

### Shallow Copy
- Copies only the immediate object
- References to other objects are shared
- Faster but can cause issues
- Use when objects are immutable

### Deep Copy
- Copies the object and all nested objects
- Each copy is completely independent
- Slower but safer
- Use when objects are mutable

## Pros and Cons

### ✅ Pros
- **Performance**: Avoids expensive object creation
- **Flexibility**: Create objects at runtime
- **Reduced Subclasses**: Avoids inheritance hierarchies
- **Dynamic Configuration**: Change prototypes at runtime

### ❌ Cons
- **Complexity**: Can be complex for deep copying
- **Circular References**: Can cause issues with circular references
- **Memory Usage**: May use more memory for deep copies
- **Implementation**: Can be tricky to implement correctly

## Real-world Examples

1. **Database Records**: Clone existing records for new entries
2. **UI Components**: Clone existing components for new instances
3. **Game Objects**: Clone game entities (enemies, items, etc.)
4. **Document Templates**: Clone document templates
5. **Configuration Objects**: Clone configuration for different environments

## Interview Questions

### Basic Level
1. What is the Prototype pattern?
2. When would you use Prototype pattern?
3. What is the difference between deep and shallow copy?

### Intermediate Level
1. How do you implement deep copying in Go?
2. What are the challenges with circular references?
3. How do you implement a prototype registry?

### Advanced Level
1. How would you implement prototype pattern with generics?
2. How do you handle memory management in prototype pattern?
3. How would you implement prototype pattern in a distributed system?

## Code Structure

```go
// Prototype interface
type Prototype interface {
    Clone() Prototype
    GetType() string
}

// Concrete prototype
type ConcretePrototype struct {
    field1 string
    field2 int
    field3 []string
}

func (cp *ConcretePrototype) Clone() Prototype {
    // Deep copy implementation
    return &ConcretePrototype{
        field1: cp.field1,
        field2: cp.field2,
        field3: append([]string{}, cp.field3...),
    }
}
```

## Next Steps

After mastering Prototype pattern, move to:
- **Abstract Factory Pattern** - For creating families of related products
- **Structural Patterns** - For organizing classes and objects
- **Behavioral Patterns** - For communication between objects

---

**Remember**: Prototype pattern is perfect for creating expensive objects or when you need to create many similar objects. Choose between deep and shallow copy based on your requirements! 🚀
