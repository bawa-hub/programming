# Builder Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Pros and Cons](#pros-and-cons)
6. [Real-world Examples](#real-world-examples)
7. [Interview Questions](#interview-questions)

## Introduction

The Builder pattern is a creational design pattern that lets you construct complex objects step by step. It allows you to produce different types and representations of an object using the same construction code.

## Problem Statement

**When to use Builder Pattern?**
- When you need to create complex objects with many optional parameters
- When you want to create different representations of the same object
- When you want to avoid constructor parameter lists that are too long
- When you want to ensure that objects are always in a valid state

**Common Scenarios:**
- Building SQL queries
- Creating configuration objects
- Building HTTP requests
- Creating complex data structures
- Building UI components

## Solution

The Builder pattern provides:
1. **Builder Interface** - Defines the construction steps
2. **Concrete Builder** - Implements the construction steps
3. **Product** - The complex object being built
4. **Director** - Orchestrates the construction process (optional)

## Implementation Approaches

### 1. Basic Builder
- Simple builder with methods for each parameter
- Fluent interface for method chaining
- Build method to create the final object

### 2. Builder with Director
- Director class orchestrates the construction
- Builder focuses on construction steps
- Director knows the construction order

### 3. Fluent Builder
- Method chaining for better readability
- Returns builder instance from each method
- More intuitive API

## Pros and Cons

### ✅ Pros
- **Step-by-step Construction**: Build complex objects step by step
- **Reusable Construction Code**: Same construction code for different representations
- **Isolated Construction**: Construction logic is isolated from the product
- **Fluent Interface**: Method chaining for better readability
- **Immutable Objects**: Can build immutable objects

### ❌ Cons
- **Complexity**: Increases overall complexity of the code
- **Many Classes**: May result in many builder classes
- **Overhead**: Additional overhead for simple objects

## Real-world Examples

1. **SQL Query Builder**: Building complex SQL queries
2. **HTTP Request Builder**: Building HTTP requests with headers, body, etc.
3. **Configuration Builder**: Building application configuration
4. **UI Component Builder**: Building complex UI components
5. **Document Builder**: Building documents with different sections

## Interview Questions

### Basic Level
1. What is the Builder pattern?
2. When would you use Builder pattern?
3. What is the difference between Builder and Factory pattern?

### Intermediate Level
1. What is method chaining in Builder pattern?
2. How do you implement a fluent builder?
3. What is the role of Director in Builder pattern?

### Advanced Level
1. How would you implement a generic builder?
2. How do you handle validation in Builder pattern?
3. How would you implement a builder with optional parameters?

## Code Structure

```go
// Product
type Product struct {
    field1 string
    field2 int
    field3 bool
}

// Builder interface
type Builder interface {
    SetField1(value string) Builder
    SetField2(value int) Builder
    SetField3(value bool) Builder
    Build() *Product
}

// Concrete builder
type ConcreteBuilder struct {
    product *Product
}

func (cb *ConcreteBuilder) SetField1(value string) Builder {
    cb.product.field1 = value
    return cb
}

func (cb *ConcreteBuilder) Build() *Product {
    return cb.product
}
```

## Next Steps

After mastering Builder pattern, move to:
- **Prototype Pattern** - For creating objects by cloning existing instances
- **Abstract Factory Pattern** - For creating families of related products
- **Structural Patterns** - For organizing classes and objects

---

**Remember**: Builder pattern is perfect for creating complex objects with many optional parameters. It makes your code more readable and maintainable! 🚀
