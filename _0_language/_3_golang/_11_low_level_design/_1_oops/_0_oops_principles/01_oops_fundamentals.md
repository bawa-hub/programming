# Object-Oriented Programming (OOPS) Fundamentals

## Table of Contents
1. [Introduction](#introduction)
2. [Four Pillars of OOPS](#four-pillars-of-oops)
3. [SOLID Principles](#solid-principles)
4. [Composition vs Inheritance](#composition-vs-inheritance)
5. [Interview Questions](#interview-questions)

## Introduction

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects" which contain data (attributes) and code (methods). It's the foundation of modern software design and is crucial for Low Level Design interviews.

### Why OOPS Matters in LLD Interviews?

1. **Design Thinking**: OOPS helps you think in terms of real-world entities
2. **Code Organization**: Proper encapsulation and modularity
3. **Maintainability**: Easy to modify and extend code
4. **Reusability**: Write once, use many times
5. **Interview Focus**: Most LLD questions test your OOPS understanding

## Four Pillars of OOPS

### 1. Encapsulation 🏠

**Definition**: Bundling data and methods that work on that data within one unit (class).

**Key Points**:
- Data hiding using access modifiers (private, protected, public)
- Provides controlled access to data through methods
- Prevents external code from directly accessing internal state

**Benefits**:
- Data integrity
- Security
- Easier maintenance
- Clear interface

**Interview Tip**: Always ask about access modifiers and data hiding in your design.

### 2. Inheritance 👨‍👩‍👧‍👦

**Definition**: Mechanism where a new class is derived from an existing class.

**Key Points**:
- "IS-A" relationship
- Code reusability
- Method overriding
- Single and multiple inheritance

**Benefits**:
- Code reuse
- Polymorphism
- Hierarchical organization

**Interview Tip**: Be careful about deep inheritance hierarchies. Prefer composition over inheritance.

### 3. Polymorphism 🔄

**Definition**: Ability of objects of different types to be treated as objects of a common type.

**Types**:
- **Compile-time Polymorphism**: Method overloading
- **Runtime Polymorphism**: Method overriding

**Benefits**:
- Flexibility
- Extensibility
- Interface consistency

**Interview Tip**: Polymorphism is crucial for designing extensible systems.

### 4. Abstraction 🎭

**Definition**: Hiding complex implementation details and showing only essential features.

**Key Points**:
- Abstract classes and interfaces
- Focus on "what" not "how"
- Contract definition

**Benefits**:
- Simplifies complex systems
- Reduces complexity
- Better maintainability

**Interview Tip**: Use interfaces to define contracts between different components.

## SOLID Principles

SOLID principles are five design principles that make software designs more understandable, flexible, and maintainable.

### 1. Single Responsibility Principle (SRP) 🎯

**Definition**: A class should have only one reason to change.

**Example**: A User class should only handle user data, not email sending.

**Benefits**:
- Easier testing
- Better maintainability
- Clearer code

### 2. Open/Closed Principle (OCP) 🔓

**Definition**: Software entities should be open for extension but closed for modification.

**Example**: Use interfaces and abstract classes to allow extension without changing existing code.

**Benefits**:
- Extensibility
- Stability
- Reduced risk

### 3. Liskov Substitution Principle (LSP) 🔄

**Definition**: Objects of a superclass should be replaceable with objects of a subclass without breaking the application.

**Example**: If Bird is a superclass, any Bird subclass should be usable wherever Bird is expected.

**Benefits**:
- Proper inheritance
- Polymorphism works correctly
- Design consistency

### 4. Interface Segregation Principle (ISP) 🔌

**Definition**: Clients should not be forced to depend on interfaces they don't use.

**Example**: Instead of one large interface, create multiple smaller, focused interfaces.

**Benefits**:
- Loose coupling
- Better maintainability
- Clearer contracts

### 5. Dependency Inversion Principle (DIP) ⬆️

**Definition**: High-level modules should not depend on low-level modules. Both should depend on abstractions.

**Example**: Use dependency injection and interfaces.

**Benefits**:
- Loose coupling
- Testability
- Flexibility

## Composition vs Inheritance

### When to Use Inheritance? 🧬

- **IS-A relationship**: Car IS-A Vehicle
- **Code reuse**: Common behavior across classes
- **Polymorphism**: Need to treat objects uniformly

### When to Use Composition? 🧩

- **HAS-A relationship**: Car HAS-A Engine
- **Flexibility**: Can change behavior at runtime
- **Multiple behaviors**: Can combine multiple components

### The Favor Composition Rule

**Prefer composition over inheritance** because:
- More flexible
- Easier to test
- Less coupling
- Better for complex hierarchies

## Interview Questions

### Basic Level
1. What are the four pillars of OOPS?
2. Explain encapsulation with an example.
3. What is the difference between method overloading and overriding?
4. What is an abstract class?

### Intermediate Level
1. Explain SOLID principles with examples.
2. When would you use composition over inheritance?
3. How do you achieve multiple inheritance in Go?
4. What is the difference between abstract class and interface?

### Advanced Level
1. Design a system using SOLID principles.
2. How would you refactor a class that violates SRP?
3. Explain the Liskov Substitution Principle with a real-world example.
4. Design a plugin system using OOPS principles.

## Next Steps

After mastering OOPS fundamentals, we'll move to:
1. **Design Patterns** - Applying OOPS principles in real-world scenarios
2. **System Design** - Using OOPS for large-scale systems
3. **Interview Practice** - Solving LLD problems using OOPS

---

**Remember**: OOPS is not just about syntax, it's about thinking in terms of objects and their relationships. Master this foundation, and you'll excel in LLD interviews! 🚀
