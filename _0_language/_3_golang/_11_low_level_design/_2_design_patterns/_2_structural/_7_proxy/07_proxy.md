# Proxy Pattern

## Table of Contents
1. [Introduction](#introduction)
2. [Problem Statement](#problem-statement)
3. [Solution](#solution)
4. [Implementation Approaches](#implementation-approaches)
5. [Proxy vs Decorator vs Adapter](#proxy-vs-decorator-vs-adapter)
6. [Pros and Cons](#pros-and-cons)
7. [Real-world Examples](#real-world-examples)
8. [Interview Questions](#interview-questions)

## Introduction

The Proxy pattern is a structural design pattern that provides a placeholder or surrogate for another object to control access to it. It acts as an intermediary between the client and the real object.

## Problem Statement

**When to use Proxy Pattern?**
- When you need to control access to an object
- When you want to add functionality before or after the core functionality
- When you need to lazy load expensive objects
- When you want to add security, caching, or logging

**Common Scenarios:**
- Virtual proxies (lazy loading)
- Protection proxies (access control)
- Remote proxies (network communication)
- Caching proxies (performance optimization)
- Logging proxies (monitoring)

## Solution

The Proxy pattern provides:
1. **Subject Interface** - Defines the common interface for RealSubject and Proxy
2. **Real Subject** - The actual object that the proxy represents
3. **Proxy** - Controls access to the real subject and may be responsible for creating and deleting it
4. **Client** - Uses the subject interface to work with both real subject and proxy

## Implementation Approaches

### 1. Virtual Proxy
- Lazy loading of expensive objects
- Creates real object only when needed
- Useful for large objects or resources

### 2. Protection Proxy
- Controls access to sensitive objects
- Implements access control logic
- Useful for security and permissions

### 3. Remote Proxy
- Represents objects in different address spaces
- Handles network communication
- Useful for distributed systems

### 4. Caching Proxy
- Caches results of expensive operations
- Reduces redundant computations
- Useful for performance optimization

## Proxy vs Decorator vs Adapter

### Proxy
- **Purpose**: Controls access to objects
- **Focus**: Access control and lazy loading
- **Use Case**: Security, caching, lazy loading

### Decorator
- **Purpose**: Adds new functionality to objects
- **Focus**: Behavior extension
- **Use Case**: Adding features dynamically

### Adapter
- **Purpose**: Makes incompatible interfaces work together
- **Focus**: Interface conversion
- **Use Case**: Integrating existing code

## Pros and Cons

### ✅ Pros
- **Access Control**: Can control access to objects
- **Lazy Loading**: Can defer object creation until needed
- **Performance**: Can cache results and optimize performance
- **Security**: Can add security checks and validation

### ❌ Cons
- **Complexity**: Can make code more complex
- **Indirection**: Adds a layer of indirection
- **Performance Overhead**: May add overhead for simple operations
- **Debugging**: Can make debugging more difficult

## Real-world Examples

1. **Virtual Proxies**: Lazy loading of images, documents
2. **Protection Proxies**: Access control for sensitive operations
3. **Remote Proxies**: Database connections, web services
4. **Caching Proxies**: API responses, database queries
5. **Logging Proxies**: Method call logging, performance monitoring

## Interview Questions

### Basic Level
1. What is the Proxy pattern?
2. When would you use Proxy pattern?
3. What are the different types of proxies?

### Intermediate Level
1. How do you implement Proxy pattern?
2. What are the benefits of using Proxy pattern?
3. How do you handle lazy loading in Proxy pattern?

### Advanced Level
1. How would you implement a caching proxy?
2. How do you handle thread safety in Proxy pattern?
3. How would you implement Proxy pattern with generics?

## Code Structure

```go
// Subject interface
type Subject interface {
    Request() string
}

// Real Subject
type RealSubject struct{}

func (rs *RealSubject) Request() string {
    return "RealSubject request"
}

// Proxy
type Proxy struct {
    realSubject *RealSubject
}

func (p *Proxy) Request() string {
    if p.realSubject == nil {
        p.realSubject = &RealSubject{}
    }
    return p.realSubject.Request()
}
```

## Next Steps

After mastering Proxy pattern, move to:
- **Behavioral Patterns** - For communication between objects
- **Observer Pattern** - For notifying multiple objects about changes
- **Strategy Pattern** - For selecting algorithms at runtime

---

**Remember**: Proxy pattern is perfect for controlling access to objects and adding functionality like caching, security, or lazy loading. It's like having a gatekeeper for your objects! 🚀
