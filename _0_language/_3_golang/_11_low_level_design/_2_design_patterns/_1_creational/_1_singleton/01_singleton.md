# Singleton Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Pros and Cons](#pros-and-cons)
6. [Real-world Examples](#real-world-examples)
7. [Interview Questions](#interview-questions)

## Introduction

The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance. It's one of the most commonly used design patterns and is often the first pattern taught.

## Problem Statement

**When to use Singleton?**
- When you need exactly one instance of a class
- When you need global access to that instance
- When you want to control resource usage (database connections, file systems, etc.)
- When you need lazy initialization

**Common Scenarios:**
- Database connection managers
- Logger instances
- Configuration managers
- Cache managers
- Thread pools

## Solution

The Singleton pattern provides:
1. **Private constructor** - Prevents direct instantiation
2. **Static instance variable** - Holds the single instance
3. **Static getter method** - Provides access to the instance
4. **Thread safety** - Ensures only one instance in multi-threaded environments

## Implementation Approaches

### 1. Eager Initialization
- Instance created at class loading time
- Simple but may waste resources if not used

### 2. Lazy Initialization
- Instance created only when first requested
- More memory efficient but requires thread safety

### 3. Thread-Safe Lazy Initialization
- Uses synchronization to ensure thread safety
- Double-checked locking for better performance

### 4. Bill Pugh Solution (Static Inner Class)
- Uses static inner class for lazy initialization
- Thread-safe without synchronization overhead

## Pros and Cons

### ✅ Pros
- **Single Instance**: Guarantees only one instance exists
- **Global Access**: Provides global access point
- **Lazy Initialization**: Can be implemented for memory efficiency
- **Resource Control**: Controls access to shared resources

### ❌ Cons
- **Global State**: Can lead to hidden dependencies
- **Testing Difficulties**: Hard to unit test
- **Thread Safety**: Requires careful implementation
- **Violates SRP**: Can become a "god object"

## Real-world Examples

1. **Database Connection Pool**: Single pool manager
2. **Logger**: Single logging instance across application
3. **Configuration Manager**: Single config instance
4. **Cache Manager**: Single cache instance
5. **Thread Pool**: Single thread pool manager

## Interview Questions

### Basic Level
1. What is the Singleton pattern?
2. When would you use Singleton?
3. What are the different ways to implement Singleton?

### Intermediate Level
1. How do you make Singleton thread-safe?
2. What is the difference between eager and lazy initialization?
3. How do you prevent Singleton from being cloned?

### Advanced Level
1. What are the drawbacks of Singleton pattern?
2. How do you test Singleton classes?
3. How do you implement Singleton in a distributed system?
4. What is the Bill Pugh Singleton solution?

## Code Structure

```go
type Singleton struct {
    // fields
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

## Next Steps

After mastering Singleton, move to:
- **Factory Pattern** - For creating objects without specifying their exact classes
- **Builder Pattern** - For constructing complex objects step by step
- **Prototype Pattern** - For creating objects by cloning existing instances

---

**Remember**: Singleton is powerful but use it judiciously. It can make code harder to test and maintain if overused! 🚀
