# 🔒 Go Sync Primitives - Complete Guide
## From Basics to Mastery

### 📚 Overview
This guide covers all Go sync primitives with clear explanations, practical examples, and common patterns. Master these to write safe, concurrent Go programs.

---

## 🎯 Table of Contents

### **Part 1: Basic Sync Primitives**
1. **Mutex** - Mutual exclusion locks
2. **RWMutex** - Read-write locks
3. **WaitGroup** - Wait for goroutines to complete
4. **Once** - One-time execution
5. **Cond** - Condition variables

### **Part 2: Advanced Sync Primitives**
6. **Atomic** - Atomic operations
7. **Pool** - Object pooling
8. **Map** - Concurrent map
9. **Semaphore** - Counting semaphores
10. **Barrier** - Synchronization barriers

### **Part 3: Common Patterns**
11. **Producer-Consumer** - Using sync primitives
12. **Worker Pool** - Managing workers
13. **Rate Limiting** - Controlling access
14. **Circuit Breaker** - Fault tolerance
15. **Graceful Shutdown** - Clean termination

---

## 🚀 How to Use This Guide

### Run Individual Examples
```bash
go run sync_primitives.go 1    # Run example 1
go run sync_primitives.go 5    # Run example 5
```

### Run by Category
```bash
go run sync_primitives.go basic     # Basic primitives (1-5)
go run sync_primitives.go advanced  # Advanced primitives (6-10)
go run sync_primitives.go patterns  # Common patterns (11-15)
```

### Run All Examples
```bash
go run sync_primitives.go all
```

---

## 🔧 Testing Commands

```bash
# Compile and test
go build sync_primitives.go
go run sync_primitives.go 1

# Run with race detection
go run -race sync_primitives.go 1

# Check for common mistakes
go vet sync_primitives.go
```

---

## 💡 Key Concepts Covered

- **Mutex**: Protecting shared resources
- **RWMutex**: Optimizing read-heavy workloads
- **WaitGroup**: Coordinating goroutines
- **Once**: One-time initialization
- **Cond**: Waiting for conditions
- **Atomic**: Lock-free operations
- **Pool**: Object reuse
- **Map**: Concurrent data structures
- **Semaphore**: Resource limiting
- **Barrier**: Synchronization points

---

**Remember**: Sync primitives are the foundation of safe concurrency. Master these before moving to advanced patterns!
