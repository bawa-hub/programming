# 🧵 Concurrency Tutorial

## Introduction

This tutorial will guide you through building a distributed real-time analytics platform using advanced Go concurrency patterns. By the end of this tutorial, you'll understand how to implement production-grade concurrent systems.

## Prerequisites

- Go 1.21 or later
- Basic understanding of Go
- Familiarity with concurrent programming concepts

## Table of Contents

1. [Getting Started](#getting-started)
2. [Actor Model Implementation](#actor-model-implementation)
3. [Reactive Programming](#reactive-programming)
4. [Circuit Breaker Pattern](#circuit-breaker-pattern)
5. [Worker Pool Pattern](#worker-pool-pattern)
6. [Rate Limiting](#rate-limiting)
7. [Connection Pooling](#connection-pooling)
8. [Lock-Free Programming](#lock-free-programming)
9. [Event Sourcing](#event-sourcing)
10. [CQRS Pattern](#cqrs-pattern)
11. [Saga Pattern](#saga-pattern)
12. [MapReduce Implementation](#mapreduce-implementation)
13. [Testing Concurrent Code](#testing-concurrent-code)
14. [Performance Optimization](#performance-optimization)
15. [Production Deployment](#production-deployment)

---

## Getting Started

### Project Setup

First, let's set up our project structure:

```bash
# Create project directory
mkdir distributed-analytics-platform
cd distributed-analytics-platform

# Initialize Go module
go mod init github.com/your-username/distributed-analytics-platform

# Create directory structure
mkdir -p cmd/{server,worker,client}
mkdir -p internal/{server,analytics,worker,storage,messaging,monitoring}
mkdir -p pkg/{concurrency,models,utils,testing}
mkdir -p configs scripts tests docs
```

### Basic Configuration

Create a basic configuration file:

```yaml
# configs/server.yaml
server:
  address: "0.0.0.0"
  port: 8080
  read_timeout: "30s"
  write_timeout: "30s"
  idle_timeout: "60s"

analytics:
  workers: 10
  batch_size: 1000
  batch_timeout: "1s"
  cache_size: 10000

monitoring:
  metrics_port: 9090
  health_check_interval: "10s"

log:
  level: "info"
  format: "json"
```

---

## Actor Model Implementation

### Theory

The Actor Model treats actors as the fundamental unit of computation. Each actor:
- Has its own state
- Communicates only through messages
- Processes one message at a time
- Can create other actors

### Implementation

Let's implement a simple actor system:

```go
// pkg/concurrency/actor/actor.go
package actor

import (
    "context"
    "sync"
    "time"
)

type Actor struct {
    id       string
    inbox    chan Message
    behavior func(Message) ActorBehavior
    state    interface{}
    mu       sync.RWMutex
    stopCh   chan struct{}
    ctx      context.Context
    cancel   context.CancelFunc
}

type Message struct {
    From    string
    To      string
    Type    string
    Payload interface{}
    ReplyTo chan Message
}

type ActorBehavior struct {
    NewState interface{}
    Messages []Message
    Spawn    []Actor
    Stop     bool
}
```

### Usage Example

```go
// Create an analytics actor
analyticsActor := actor.NewActor("analytics", func(msg actor.Message) actor.ActorBehavior {
    switch msg.Type {
    case "process_event":
        // Process the event
        event := msg.Payload.(Event)
        result := processEvent(event)
        
        return actor.ActorBehavior{
            NewState: result,
            Messages: []actor.Message{{
                To: "dashboard",
                Type: "update_metrics",
                Payload: result,
            }},
        }
    }
    return actor.ActorBehavior{}
})

// Send a message
analyticsActor.Send(actor.Message{
    From: "api",
    Type: "process_event",
    Payload: eventData,
})
```

### Best Practices

1. **Keep actors small and focused**
2. **Use immutable state when possible**
3. **Handle errors gracefully**
4. **Implement proper shutdown**

---

## Reactive Programming

### Theory

Reactive Programming is about working with data streams and the propagation of change. It provides:
- Asynchronous processing
- Backpressure handling
- Composable operations
- Error handling

### Implementation

```go
// pkg/concurrency/reactive/stream.go
package reactive

import (
    "context"
    "sync"
)

type Stream[T any] struct {
    source    chan T
    operators []Operator[T]
    mu        sync.RWMutex
    stopCh    chan struct{}
    ctx       context.Context
    cancel    context.CancelFunc
}

type Operator[T any] interface {
    Process(input <-chan T) <-chan T
}

type MapOperator[T, U any] struct {
    mapper func(T) U
}

func (m *MapOperator[T, U]) Process(input <-chan T) <-chan U {
    output := make(chan U)
    go func() {
        defer close(output)
        for item := range input {
            select {
            case output <- m.mapper(item):
            case <-m.ctx.Done():
                return
            }
        }
    }()
    return output
}
```

### Usage Example

```go
// Create a reactive stream
stream := reactive.NewStream[int]()

// Chain operations
result := stream.
    Map(func(x int) int { return x * 2 }).
    Filter(func(x int) bool { return x > 10 }).
    Reduce(0, func(acc, x int) int { return acc + x })

// Emit data
stream.Emit(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Get result
total := <-result
```

### Best Practices

1. **Handle backpressure**
2. **Use context for cancellation**
3. **Compose operations carefully**
4. **Handle errors in streams**

---

## Circuit Breaker Pattern

### Theory

A Circuit Breaker prevents a failure from constantly recurring by:
- Monitoring calls to a service
- Opening the circuit when failures exceed threshold
- Allowing some calls through when circuit is half-open
- Closing circuit when service recovers

### Implementation

```go
// pkg/concurrency/circuitbreaker/breaker.go
package circuitbreaker

import (
    "sync"
    "time"
)

type CircuitBreaker struct {
    name            string
    maxRequests     uint32
    interval        time.Duration
    timeout         time.Duration
    readyToTrip     func(counts Counts) bool
    onStateChange   func(name string, from State, to State)
    
    mutex      sync.Mutex
    state      State
    generation uint64
    counts     Counts
    expiry     time.Time
}

type State int

const (
    StateClosed State = iota
    StateHalfOpen
    StateOpen
)

type Counts struct {
    Requests             uint32
    TotalSuccesses       uint32
    TotalFailures        uint32
    ConsecutiveSuccesses uint32
    ConsecutiveFailures  uint32
}
```

### Usage Example

```go
// Create circuit breaker
cb := circuitbreaker.NewCircuitBreaker("database", 5, 10*time.Second, 30*time.Second)

// Execute with circuit breaker
result, err := cb.Execute(func() (interface{}, error) {
    return database.Query("SELECT * FROM users")
})

if err != nil {
    // Handle error or use fallback
    return fallbackData, nil
}
```

### Best Practices

1. **Set appropriate thresholds**
2. **Monitor circuit breaker state**
3. **Implement fallback mechanisms**
4. **Test failure scenarios**

---

## Worker Pool Pattern

### Theory

A Worker Pool manages a fixed number of workers to process tasks from a queue. It provides:
- Resource management
- Load balancing
- Backpressure control
- Performance optimization

### Implementation

```go
// pkg/concurrency/workerpool/pool.go
package workerpool

import (
    "context"
    "sync"
)

type Pool struct {
    workers    int
    tasks      chan Task
    results    chan Result
    wg         sync.WaitGroup
    stopCh     chan struct{}
    mu         sync.RWMutex
    isRunning  bool
}

type Task struct {
    ID       string
    Data     interface{}
    Handler  func(interface{}) (interface{}, error)
}

type Result struct {
    TaskID string
    Data   interface{}
    Error  error
}
```

### Usage Example

```go
// Create worker pool
pool := workerpool.NewPool(10) // 10 workers

// Start pool
pool.Start()

// Submit tasks
for i := 0; i < 100; i++ {
    pool.Submit(workerpool.Task{
        ID: fmt.Sprintf("task-%d", i),
        Data: i,
        Handler: func(data interface{}) (interface{}, error) {
            // Process data
            return data.(int) * 2, nil
        },
    })
}

// Get results
for i := 0; i < 100; i++ {
    result := <-pool.Results()
    fmt.Printf("Task %s: %v\n", result.TaskID, result.Data)
}
```

### Best Practices

1. **Size pools appropriately**
2. **Handle task errors**
3. **Implement graceful shutdown**
4. **Monitor pool metrics**

---

## Rate Limiting

### Theory

Rate Limiting controls the rate of requests to prevent abuse and ensure fair resource usage. Common algorithms:
- Token Bucket
- Leaky Bucket
- Fixed Window
- Sliding Window

### Implementation

```go
// pkg/concurrency/ratelimit/limiter.go
package ratelimit

import (
    "sync"
    "time"
)

type Limiter struct {
    limit    int
    interval time.Duration
    tokens   int
    lastTime time.Time
    mu       sync.Mutex
}

type TokenBucket struct {
    capacity     int
    tokens       int
    refillRate   int
    lastRefill   time.Time
    mu           sync.Mutex
}
```

### Usage Example

```go
// Create rate limiter
limiter := ratelimit.NewRateLimiter(100, time.Second) // 100 requests per second

// Check if request is allowed
if limiter.Allow() {
    // Process request
    processRequest()
} else {
    // Rate limited
    return errors.New("rate limited")
}
```

### Best Practices

1. **Choose appropriate algorithm**
2. **Set reasonable limits**
3. **Monitor rate limiting metrics**
4. **Implement different limits per user**

---

## Connection Pooling

### Theory

Connection Pooling manages a pool of reusable connections to improve performance and resource utilization.

### Implementation

```go
// pkg/concurrency/pool/connection.go
package pool

import (
    "sync"
    "time"
)

type Pool struct {
    factory    func() (Connection, error)
    pool       chan Connection
    maxSize    int
    minSize    int
    mu         sync.RWMutex
    closed     bool
    closeCh    chan struct{}
}

type Connection interface {
    Query(string) (Result, error)
    Close() error
    Ping() error
}
```

### Usage Example

```go
// Create connection pool
pool := pool.NewPool(func() (Connection, error) {
    return database.Connect()
}, 10, 5) // max 10, min 5 connections

// Get connection
conn, err := pool.Get()
if err != nil {
    return err
}
defer pool.Put(conn)

// Use connection
result, err := conn.Query("SELECT * FROM users")
```

### Best Practices

1. **Size pools appropriately**
2. **Handle connection failures**
3. **Implement health checks**
4. **Monitor pool metrics**

---

## Lock-Free Programming

### Theory

Lock-Free Programming uses atomic operations instead of locks to ensure thread safety, providing:
- No blocking
- High performance
- Deadlock prevention
- Better scalability

### Implementation

```go
// pkg/concurrency/lockfree/queue.go
package lockfree

import (
    "sync/atomic"
    "unsafe"
)

type LockFreeQueue struct {
    head unsafe.Pointer
    tail unsafe.Pointer
}

type node struct {
    value interface{}
    next  unsafe.Pointer
}

func (q *LockFreeQueue) Enqueue(value interface{}) {
    newNode := &node{value: value}
    
    for {
        tail := atomic.LoadPointer(&q.tail)
        next := atomic.LoadPointer(&(*node)(tail).next)
        
        if tail == atomic.LoadPointer(&q.tail) {
            if next == nil {
                if atomic.CompareAndSwapPointer(&(*node)(tail).next, next, unsafe.Pointer(newNode)) {
                    break
                }
            } else {
                atomic.CompareAndSwapPointer(&q.tail, tail, next)
            }
        }
    }
    
    atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newNode))
}
```

### Usage Example

```go
// Create lock-free queue
queue := lockfree.NewLockFreeQueue()

// Enqueue items
for i := 0; i < 1000; i++ {
    queue.Enqueue(i)
}

// Dequeue items
for i := 0; i < 1000; i++ {
    value := queue.Dequeue()
    fmt.Println(value)
}
```

### Best Practices

1. **Understand memory ordering**
2. **Test thoroughly**
3. **Use atomic operations correctly**
4. **Handle ABA problem**

---

## Event Sourcing

### Theory

Event Sourcing stores changes as a sequence of events, providing:
- Complete history
- Audit trail
- State reconstruction
- Temporal queries

### Implementation

```go
// pkg/concurrency/eventsourcing/event.go
package eventsourcing

import (
    "sync"
    "time"
)

type Event struct {
    ID        string
    Type      string
    AggregateID string
    Data      interface{}
    Timestamp time.Time
    Version   int
}

type EventStore struct {
    events []Event
    mu     sync.RWMutex
}

type Aggregate struct {
    ID      string
    Version int
    State   interface{}
    Events  []Event
}
```

### Usage Example

```go
// Create event store
store := eventsourcing.NewEventStore()

// Create aggregate
user := eventsourcing.NewUserAggregate("user-123")

// Apply events
user.ApplyEvent(eventsourcing.Event{
    Type: "UserCreated",
    Data: UserCreatedData{Name: "John", Email: "john@example.com"},
})

user.ApplyEvent(eventsourcing.Event{
    Type: "UserUpdated",
    Data: UserUpdatedData{Name: "John Doe"},
})

// Save events
store.SaveEvents(user.ID, user.Events)
```

### Best Practices

1. **Design events carefully**
2. **Handle event versioning**
3. **Implement snapshots**
4. **Consider event migration**

---

## CQRS Pattern

### Theory

CQRS separates read and write operations:
- Commands: Write operations
- Queries: Read operations
- Independent optimization
- Separate scaling

### Implementation

```go
// pkg/concurrency/cqrs/cqrs.go
package cqrs

import (
    "sync"
)

type Command interface {
    Type() string
    Data() interface{}
}

type Query interface {
    Type() string
    Data() interface{}
}

type CommandHandler interface {
    Handle(Command) error
}

type QueryHandler interface {
    Handle(Query) (interface{}, error)
}

type CQRS struct {
    commandHandlers map[string]CommandHandler
    queryHandlers   map[string]QueryHandler
    mu              sync.RWMutex
}
```

### Usage Example

```go
// Create CQRS
cqrs := cqrs.NewCQRS()

// Register command handler
cqrs.RegisterCommandHandler("CreateUser", &CreateUserHandler{})

// Register query handler
cqrs.RegisterQueryHandler("GetUser", &GetUserHandler{})

// Execute command
err := cqrs.ExecuteCommand(CreateUserCommand{
    Name: "John",
    Email: "john@example.com",
})

// Execute query
user, err := cqrs.ExecuteQuery(GetUserQuery{
    ID: "user-123",
})
```

### Best Practices

1. **Keep commands and queries simple**
2. **Use appropriate data stores**
3. **Handle eventual consistency**
4. **Monitor performance**

---

## Saga Pattern

### Theory

The Saga Pattern manages distributed transactions using compensating actions:
- Compensating actions
- Eventual consistency
- Fault tolerance
- No distributed locks

### Implementation

```go
// pkg/concurrency/saga/saga.go
package saga

import (
    "sync"
)

type Saga struct {
    id       string
    steps    []Step
    status   Status
    mu       sync.RWMutex
}

type Step struct {
    name         string
    action       func() error
    compensation func() error
}

type Status int

const (
    StatusPending Status = iota
    StatusRunning
    StatusCompleted
    StatusFailed
    StatusCompensating
)
```

### Usage Example

```go
// Create saga
saga := saga.NewSaga("order-processing")

// Add steps
saga.AddStep(saga.Step{
    name: "reserve-inventory",
    action: func() error {
        return inventory.Reserve(productID, quantity)
    },
    compensation: func() error {
        return inventory.Release(productID, quantity)
    },
})

saga.AddStep(saga.Step{
    name: "charge-payment",
    action: func() error {
        return payment.Charge(customerID, amount)
    },
    compensation: func() error {
        return payment.Refund(customerID, amount)
    },
})

// Execute saga
err := saga.Execute()
```

### Best Practices

1. **Design compensating actions carefully**
2. **Handle partial failures**
3. **Implement idempotency**
4. **Monitor saga execution**

---

## MapReduce Implementation

### Theory

MapReduce processes large datasets by:
- Map phase: Transform input data
- Shuffle phase: Group by key
- Reduce phase: Aggregate values
- Distributed execution

### Implementation

```go
// pkg/concurrency/mapreduce/mapreduce.go
package mapreduce

import (
    "sync"
)

type MapReduce struct {
    mapper  Mapper
    reducer Reducer
    workers int
}

type Mapper interface {
    Map(input interface{}) []KeyValue
}

type Reducer interface {
    Reduce(key string, values []interface{}) interface{}
}

type KeyValue struct {
    Key   string
    Value interface{}
}
```

### Usage Example

```go
// Create MapReduce
mr := mapreduce.NewMapReduce(&WordCountMapper{}, &WordCountReducer{}, 10)

// Process data
data := []string{"hello world", "hello go", "world of go"}
result := mr.Process(data)

// Result contains word counts
for word, count := range result {
    fmt.Printf("%s: %d\n", word, count)
}
```

### Best Practices

1. **Design map and reduce functions carefully**
2. **Handle data skew**
3. **Implement fault tolerance**
4. **Monitor performance**

---

## Testing Concurrent Code

### Race Detection

Always run tests with race detection:

```bash
go test -race ./...
```

### Stress Testing

Test under high load:

```go
func TestConcurrentAccess(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // Test concurrent access
        }()
    }
    wg.Wait()
}
```

### Property-Based Testing

Test properties that should always hold:

```go
func TestMapReduceProperty(t *testing.T) {
    // Test that map-reduce preserves total count
    input := generateRandomData(1000)
    result := mapReduce.Process(input)
    
    total := 0
    for _, count := range result {
        total += count.(int)
    }
    
    assert.Equal(t, len(input), total)
}
```

### Best Practices

1. **Always use race detection**
2. **Test under load**
3. **Use property-based testing**
4. **Test error conditions**

---

## Performance Optimization

### Profiling

Use Go's built-in profiling tools:

```bash
# CPU profiling
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

### Benchmarking

Write benchmarks for critical paths:

```go
func BenchmarkActorMessagePassing(b *testing.B) {
    actor := NewActor("test", func(msg Message) ActorBehavior {
        return ActorBehavior{}
    })
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        actor.Send(Message{Type: "test"})
    }
}
```

### Best Practices

1. **Profile before optimizing**
2. **Write benchmarks**
3. **Measure, don't guess**
4. **Optimize hot paths**

---

## Production Deployment

### Configuration

Use environment-specific configurations:

```yaml
# configs/production.yaml
server:
  address: "0.0.0.0"
  port: 8080
  workers: 100

analytics:
  workers: 50
  batch_size: 10000
  cache_size: 100000
```

### Monitoring

Implement comprehensive monitoring:

```go
// Monitor actor system
func (as *ActorSystem) GetMetrics() Metrics {
    as.mu.RLock()
    defer as.mu.RUnlock()
    
    return Metrics{
        ActorCount: len(as.actors),
        MessageCount: as.messageCount,
        ErrorCount: as.errorCount,
    }
}
```

### Best Practices

1. **Use configuration management**
2. **Implement health checks**
3. **Monitor key metrics**
4. **Set up alerting**

---

## Conclusion

This tutorial has covered the essential concurrency patterns for building distributed systems in Go. By implementing these patterns correctly, you can build high-performance, fault-tolerant systems that scale to handle massive loads.

### Key Takeaways

1. **Choose the right pattern** for each use case
2. **Test thoroughly** with race detection and stress testing
3. **Monitor performance** and optimize hot paths
4. **Handle errors gracefully** and implement fallbacks
5. **Document everything** for future maintenance

### Next Steps

1. **Implement the patterns** in your own projects
2. **Experiment with combinations** of patterns
3. **Read the source code** of production systems
4. **Contribute to open source** projects
5. **Share your knowledge** with the community

---

*This tutorial demonstrates GOD-LEVEL Go concurrency skills through practical, real-world examples.*
