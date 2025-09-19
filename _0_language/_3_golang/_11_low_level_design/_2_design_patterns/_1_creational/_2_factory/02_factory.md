# Factory Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Types of Factory Patterns](#types-of-factory-patterns)
5. [Implementation Approaches](#implementation-approaches)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Factory pattern provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. It's one of the most widely used creational patterns.

## Problem Statement

**When to use Factory Pattern?**
- When you don't know the exact types of objects your code should work with
- When you want to provide a library of products and reveal only their interfaces
- When you want to extend your code with new product types without changing existing code
- When you want to centralize object creation logic

**Common Scenarios:**
- Creating different types of database connections
- Creating different types of UI components
- Creating different types of payment processors
- Creating different types of file readers/writers

## Solution

The Factory pattern provides:
1. **Product Interface** - Common interface for all products
2. **Concrete Products** - Specific implementations of the interface
3. **Factory Interface** - Interface for creating products
4. **Concrete Factories** - Specific implementations of the factory

## Types of Factory Patterns

### 1. Simple Factory
- Single factory class creates all products
- Uses a parameter to determine which product to create
- Simplest form of factory pattern

### 2. Factory Method
- Abstract factory class with abstract factory method
- Each concrete factory creates specific products
- Follows Open/Closed Principle

### 3. Abstract Factory
- Factory of factories
- Creates families of related products
- More complex but more flexible

## Implementation Approaches

### 1. Simple Factory
```go
type Product interface {
    GetName() string
}

type ProductFactory struct{}

func (pf *ProductFactory) CreateProduct(productType string) Product {
    switch productType {
    case "A":
        return &ProductA{}
    case "B":
        return &ProductB{}
    default:
        return nil
    }
}
```

### 2. Factory Method
```go
type Factory interface {
    CreateProduct() Product
}

type FactoryA struct{}
func (f *FactoryA) CreateProduct() Product {
    return &ProductA{}
}
```

### 3. Abstract Factory
```go
type AbstractFactory interface {
    CreateProductA() ProductA
    CreateProductB() ProductB
}
```

## Pros and Cons

### ✅ Pros
- **Loose Coupling**: Client code doesn't depend on concrete classes
- **Single Responsibility**: Object creation logic is centralized
- **Open/Closed Principle**: Easy to add new product types
- **Code Reusability**: Factory logic can be reused

### ❌ Cons
- **Complexity**: Can make code more complex
- **Many Classes**: May result in many factory classes
- **Indirection**: Adds a layer of indirection

## Real-world Examples

1. **Database Factory**: Create different database connections
2. **UI Component Factory**: Create different UI components
3. **Payment Processor Factory**: Create different payment processors
4. **File Reader Factory**: Create different file readers
5. **Logger Factory**: Create different logger types

## Interview Questions

### Basic Level
1. What is the Factory pattern?
2. What are the different types of Factory patterns?
3. When would you use Factory pattern?

### Intermediate Level
1. What is the difference between Simple Factory and Factory Method?
2. How does Factory pattern follow Open/Closed Principle?
3. What is the difference between Factory Method and Abstract Factory?

### Advanced Level
1. How would you implement a Factory pattern for a plugin system?
2. How do you handle error cases in Factory pattern?
3. How would you implement a Factory pattern with dependency injection?

## Code Structure

```go
// Product interface
type Product interface {
    GetName() string
    GetPrice() float64
}

// Concrete products
type ProductA struct{}
type ProductB struct{}

// Factory interface
type Factory interface {
    CreateProduct() Product
}

// Concrete factories
type FactoryA struct{}
type FactoryB struct{}
```

## Next Steps

After mastering Factory pattern, move to:
- **Abstract Factory Pattern** - For creating families of related products
- **Builder Pattern** - For constructing complex objects step by step
- **Prototype Pattern** - For creating objects by cloning existing instances

---

**Remember**: Factory pattern is about creating objects without specifying their exact classes. It's a fundamental pattern for building flexible and maintainable code! 🚀
