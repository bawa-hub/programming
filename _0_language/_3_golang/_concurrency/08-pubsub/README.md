# Level 1, Topic 8: Pub-Sub Pattern

## 🎯 Learning Objectives

By the end of this topic, you will master:
- **Pub-Sub Architecture**: Understanding the publish-subscribe messaging pattern
- **Event-Driven Design**: Building loosely coupled, event-driven systems
- **Message Routing**: Implementing topic-based and content-based routing
- **Subscriber Management**: Managing multiple subscribers and subscriptions
- **Message Persistence**: Handling message durability and delivery guarantees
- **Error Handling**: Managing failures in distributed messaging
- **Performance Optimization**: Scaling pub-sub systems for high throughput
- **Advanced Patterns**: Dead letter queues, message ordering, and transactions

## 📚 Theory Deep Dive

### What is Pub-Sub?

The Publish-Subscribe (Pub-Sub) pattern is a messaging pattern where:
1. **Publishers**: Send messages to topics without knowing who will receive them
2. **Subscribers**: Subscribe to topics and receive messages without knowing who sent them
3. **Broker**: Acts as an intermediary, routing messages from publishers to subscribers

```
Publisher → Topic → Broker → Subscribers
```

### Core Concepts

#### 1. Publishers
- **Purpose**: Send messages to specific topics
- **Benefits**: Decoupled from subscribers, scalable
- **Challenges**: Message delivery guarantees, error handling

#### 2. Subscribers
- **Purpose**: Receive messages from subscribed topics
- **Benefits**: Decoupled from publishers, can process at their own pace
- **Challenges**: Message ordering, duplicate handling

#### 3. Topics
- **Purpose**: Logical channels for message routing
- **Benefits**: Organized message flow, filtering capabilities
- **Challenges**: Topic management, access control

#### 4. Broker
- **Purpose**: Routes messages between publishers and subscribers
- **Benefits**: Centralized message management, reliability
- **Challenges**: Single point of failure, scalability

### Pattern Variations

#### 1. Topic-Based Pub-Sub
```go
// Publishers send to specific topics
publisher.Publish("user.events", message)

// Subscribers subscribe to topics
subscriber.Subscribe("user.events", handler)
```

#### 2. Content-Based Pub-Sub
```go
// Subscribers filter by message content
subscriber.Subscribe(func(msg Message) bool {
    return msg.Type == "user.created" && msg.UserID > 1000
}, handler)
```

#### 3. Type-Based Pub-Sub
```go
// Subscribers filter by message type
subscriber.Subscribe(UserCreatedEvent{}, handler)
subscriber.Subscribe(OrderPlacedEvent{}, handler)
```

#### 4. Wildcard Pub-Sub
```go
// Subscribers can use wildcards
subscriber.Subscribe("user.*", handler)  // All user events
subscriber.Subscribe("*.created", handler)  // All created events
```

### Channel Patterns

#### 1. Single Topic, Multiple Subscribers
```go
topic := make(chan Message, 100)
go subscriber1(topic)
go subscriber2(topic)
go subscriber3(topic)
```

#### 2. Multiple Topics, Single Subscriber
```go
userTopic := make(chan Message, 100)
orderTopic := make(chan Message, 100)
go subscriber(userTopic, orderTopic)
```

#### 3. Fan-Out Pattern
```go
func fanOut(input <-chan Message, outputs []chan Message) {
    for msg := range input {
        for _, output := range outputs {
            select {
            case output <- msg:
            default:
                // Handle full channel
            }
        }
    }
}
```

### Message Types

#### 1. Event Messages
```go
type Event struct {
    ID        string
    Type      string
    Timestamp time.Time
    Data      interface{}
    Metadata  map[string]string
}
```

#### 2. Command Messages
```go
type Command struct {
    ID        string
    Type      string
    Timestamp time.Time
    Payload   interface{}
    ReplyTo   string
}
```

#### 3. Query Messages
```go
type Query struct {
    ID        string
    Type      string
    Timestamp time.Time
    Criteria  map[string]interface{}
    ReplyTo   string
}
```

### Error Handling Strategies

#### 1. Retry Logic
```go
func publishWithRetry(topic string, msg Message, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        if err := broker.Publish(topic, msg); err == nil {
            return nil
        }
        time.Sleep(time.Duration(i) * time.Second)
    }
    return fmt.Errorf("failed after %d retries", maxRetries)
}
```

#### 2. Dead Letter Queue
```go
type DeadLetterQueue struct {
    queue chan Message
    dlq   chan Message
}

func (dlq *DeadLetterQueue) Process() {
    for msg := range dlq.queue {
        if err := processMessage(msg); err != nil {
            dlq.dlq <- msg
        }
    }
}
```

#### 3. Circuit Breaker
```go
type CircuitBreaker struct {
    failureCount int
    threshold    int
    timeout      time.Duration
    state        State
}

func (cb *CircuitBreaker) Publish(topic string, msg Message) error {
    if cb.state == Open {
        return fmt.Errorf("circuit breaker is open")
    }
    // Publish message
}
```

### Performance Considerations

#### 1. Message Batching
```go
type BatchedPublisher struct {
    batchSize int
    timeout   time.Duration
    messages  []Message
    mutex     sync.Mutex
}

func (bp *BatchedPublisher) Publish(msg Message) {
    bp.mutex.Lock()
    defer bp.mutex.Unlock()
    
    bp.messages = append(bp.messages, msg)
    if len(bp.messages) >= bp.batchSize {
        bp.flush()
    }
}
```

#### 2. Message Compression
```go
func compressMessage(msg Message) ([]byte, error) {
    data, err := json.Marshal(msg)
    if err != nil {
        return nil, err
    }
    
    var buf bytes.Buffer
    writer := gzip.NewWriter(&buf)
    writer.Write(data)
    writer.Close()
    
    return buf.Bytes(), nil
}
```

#### 3. Message Persistence
```go
type PersistentBroker struct {
    storage Storage
    topics  map[string]chan Message
}

func (pb *PersistentBroker) Publish(topic string, msg Message) error {
    // Persist message
    if err := pb.storage.Store(topic, msg); err != nil {
        return err
    }
    
    // Send to subscribers
    select {
    case pb.topics[topic] <- msg:
    default:
        return fmt.Errorf("topic %s is full", topic)
    }
    
    return nil
}
```

### Advanced Patterns

#### 1. Message Ordering
```go
type OrderedBroker struct {
    topics    map[string]chan Message
    sequences map[string]int64
    mutex     sync.Mutex
}

func (ob *OrderedBroker) Publish(topic string, msg Message) error {
    ob.mutex.Lock()
    defer ob.mutex.Unlock()
    
    sequence := ob.sequences[topic]
    msg.Sequence = sequence
    ob.sequences[topic]++
    
    return ob.sendToTopic(topic, msg)
}
```

#### 2. Message Deduplication
```go
type DeduplicatingBroker struct {
    topics      map[string]chan Message
    seen        map[string]bool
    mutex       sync.Mutex
    cleanupTime time.Duration
}

func (db *DeduplicatingBroker) Publish(topic string, msg Message) error {
    db.mutex.Lock()
    defer db.mutex.Unlock()
    
    if db.seen[msg.ID] {
        return nil // Duplicate, ignore
    }
    
    db.seen[msg.ID] = true
    return db.sendToTopic(topic, msg)
}
```

#### 3. Message Routing
```go
type Router struct {
    routes map[string][]string
    broker *Broker
}

func (r *Router) Route(topic string, msg Message) error {
    targets := r.routes[topic]
    for _, target := range targets {
        if err := r.broker.Publish(target, msg); err != nil {
            return err
        }
    }
    return nil
}
```

#### 4. Message Filtering
```go
type FilteringBroker struct {
    topics   map[string]chan Message
    filters  map[string]func(Message) bool
    mutex    sync.RWMutex
}

func (fb *FilteringBroker) Subscribe(topic string, filter func(Message) bool, handler func(Message)) {
    fb.mutex.Lock()
    defer fb.mutex.Unlock()
    
    fb.filters[topic] = filter
    
    go func() {
        for msg := range fb.topics[topic] {
            if filter(msg) {
                handler(msg)
            }
        }
    }()
}
```

## 🏗️ Implementation Patterns

### Basic Pub-Sub
```go
type Broker struct {
    topics map[string]chan Message
    mutex  sync.RWMutex
}

func NewBroker() *Broker {
    return &Broker{
        topics: make(map[string]chan Message),
    }
}

func (b *Broker) Publish(topic string, msg Message) error {
    b.mutex.RLock()
    ch, exists := b.topics[topic]
    b.mutex.RUnlock()
    
    if !exists {
        return fmt.Errorf("topic %s does not exist", topic)
    }
    
    select {
    case ch <- msg:
        return nil
    default:
        return fmt.Errorf("topic %s is full", topic)
    }
}

func (b *Broker) Subscribe(topic string, handler func(Message)) {
    b.mutex.Lock()
    defer b.mutex.Unlock()
    
    if _, exists := b.topics[topic]; !exists {
        b.topics[topic] = make(chan Message, 100)
    }
    
    go func() {
        for msg := range b.topics[topic] {
            handler(msg)
        }
    }()
}
```

### Persistent Pub-Sub
```go
type PersistentBroker struct {
    broker  *Broker
    storage Storage
}

func (pb *PersistentBroker) Publish(topic string, msg Message) error {
    // Persist message
    if err := pb.storage.Store(topic, msg); err != nil {
        return err
    }
    
    // Send to subscribers
    return pb.broker.Publish(topic, msg)
}

func (pb *PersistentBroker) Subscribe(topic string, handler func(Message)) {
    // Subscribe to new messages
    pb.broker.Subscribe(topic, handler)
    
    // Replay persisted messages
    go func() {
        messages, err := pb.storage.GetMessages(topic)
        if err != nil {
            return
        }
        
        for _, msg := range messages {
            handler(msg)
        }
    }()
}
```

### Clustered Pub-Sub
```go
type ClusteredBroker struct {
    localBroker *Broker
    peers       []*Peer
    mutex       sync.RWMutex
}

func (cb *ClusteredBroker) Publish(topic string, msg Message) error {
    // Publish locally
    if err := cb.localBroker.Publish(topic, msg); err != nil {
        return err
    }
    
    // Replicate to peers
    cb.mutex.RLock()
    peers := cb.peers
    cb.mutex.RUnlock()
    
    for _, peer := range peers {
        go peer.Publish(topic, msg)
    }
    
    return nil
}
```

## 🎯 Use Cases

### 1. Event-Driven Architecture
- **User Events**: User registration, login, profile updates
- **Order Events**: Order placed, paid, shipped, delivered
- **System Events**: Service started, stopped, health checks

### 2. Microservices Communication
- **Service Discovery**: Service registration and discovery
- **Configuration Updates**: Dynamic configuration changes
- **Health Monitoring**: Service health status updates

### 3. Real-time Applications
- **Chat Applications**: Real-time messaging
- **Live Updates**: Stock prices, sports scores
- **Notifications**: Push notifications, alerts

### 4. Data Processing
- **ETL Pipelines**: Extract, transform, load operations
- **Stream Processing**: Real-time data processing
- **Event Sourcing**: Event store and replay

## ⚡ Performance Optimization

### 1. Message Batching
```go
type BatchedPublisher struct {
    batchSize int
    timeout   time.Duration
    messages  []Message
    mutex     sync.Mutex
    broker    *Broker
}

func (bp *BatchedPublisher) Publish(msg Message) {
    bp.mutex.Lock()
    defer bp.mutex.Unlock()
    
    bp.messages = append(bp.messages, msg)
    if len(bp.messages) >= bp.batchSize {
        bp.flush()
    }
}

func (bp *BatchedPublisher) flush() {
    if len(bp.messages) == 0 {
        return
    }
    
    batch := make([]Message, len(bp.messages))
    copy(batch, bp.messages)
    bp.messages = bp.messages[:0]
    
    go func() {
        for _, msg := range batch {
            bp.broker.Publish(msg.Topic, msg)
        }
    }()
}
```

### 2. Connection Pooling
```go
type ConnectionPool struct {
    connections chan *Connection
    factory     func() *Connection
    maxSize     int
}

func (cp *ConnectionPool) Get() *Connection {
    select {
    case conn := <-cp.connections:
        return conn
    default:
        return cp.factory()
    }
}

func (cp *ConnectionPool) Put(conn *Connection) {
    select {
    case cp.connections <- conn:
    default:
        // Pool is full, close connection
        conn.Close()
    }
}
```

### 3. Message Compression
```go
func compressMessage(msg Message) ([]byte, error) {
    data, err := json.Marshal(msg)
    if err != nil {
        return nil, err
    }
    
    var buf bytes.Buffer
    writer := gzip.NewWriter(&buf)
    writer.Write(data)
    writer.Close()
    
    return buf.Bytes(), nil
}

func decompressMessage(data []byte) (Message, error) {
    reader, err := gzip.NewReader(bytes.NewReader(data))
    if err != nil {
        return Message{}, err
    }
    defer reader.Close()
    
    var msg Message
    err = json.NewDecoder(reader).Decode(&msg)
    return msg, err
}
```

## 🚨 Common Pitfalls

### 1. Message Loss
```go
// ❌ Wrong: Unbuffered channel can lose messages
topic := make(chan Message)

// ✅ Correct: Buffered channel prevents message loss
topic := make(chan Message, 1000)
```

### 2. Deadlocks
```go
// ❌ Wrong: Blocking operations can cause deadlocks
func (b *Broker) Publish(topic string, msg Message) {
    b.topics[topic] <- msg // Can block
}

// ✅ Correct: Non-blocking with timeout
func (b *Broker) Publish(topic string, msg Message) error {
    select {
    case b.topics[topic] <- msg:
        return nil
    case <-time.After(5 * time.Second):
        return fmt.Errorf("timeout publishing to %s", topic)
    }
}
```

### 3. Memory Leaks
```go
// ❌ Wrong: Subscribers not properly cleaned up
func (b *Broker) Subscribe(topic string, handler func(Message)) {
    go func() {
        for msg := range b.topics[topic] {
            handler(msg)
        }
        // Goroutine never exits
    }()
}

// ✅ Correct: Proper cleanup with context
func (b *Broker) Subscribe(ctx context.Context, topic string, handler func(Message)) {
    go func() {
        for {
            select {
            case msg := <-b.topics[topic]:
                handler(msg)
            case <-ctx.Done():
                return
            }
        }
    }()
}
```

### 4. Race Conditions
```go
// ❌ Wrong: Shared state without synchronization
type Broker struct {
    topics map[string]chan Message
    // Missing mutex
}

// ✅ Correct: Proper synchronization
type Broker struct {
    topics map[string]chan Message
    mutex  sync.RWMutex
}
```

## 🔧 Testing Strategies

### 1. Unit Testing
```go
func TestBrokerPublish(t *testing.T) {
    broker := NewBroker()
    topic := "test.topic"
    
    // Create topic
    broker.topics[topic] = make(chan Message, 1)
    
    // Publish message
    msg := Message{ID: "1", Data: "test"}
    err := broker.Publish(topic, msg)
    
    assert.NoError(t, err)
    
    // Verify message was sent
    select {
    case received := <-broker.topics[topic]:
        assert.Equal(t, msg, received)
    case <-time.After(1 * time.Second):
        t.Fatal("message not received")
    }
}
```

### 2. Integration Testing
```go
func TestPubSubIntegration(t *testing.T) {
    broker := NewBroker()
    topic := "integration.test"
    
    var received []Message
    broker.Subscribe(topic, func(msg Message) {
        received = append(received, msg)
    })
    
    // Publish multiple messages
    for i := 0; i < 10; i++ {
        msg := Message{ID: fmt.Sprintf("%d", i), Data: fmt.Sprintf("test%d", i)}
        broker.Publish(topic, msg)
    }
    
    // Wait for messages
    time.Sleep(100 * time.Millisecond)
    
    assert.Len(t, received, 10)
}
```

### 3. Performance Testing
```go
func BenchmarkBrokerPublish(b *testing.B) {
    broker := NewBroker()
    topic := "benchmark.topic"
    broker.topics[topic] = make(chan Message, 1000)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        msg := Message{ID: fmt.Sprintf("%d", i), Data: "benchmark"}
        broker.Publish(topic, msg)
    }
}
```

## 📊 Monitoring and Metrics

### 1. Key Metrics
- **Throughput**: Messages per second
- **Latency**: Message delivery time
- **Error Rate**: Failed message deliveries
- **Queue Length**: Pending messages per topic
- **Subscriber Count**: Active subscribers per topic

### 2. Health Checks
```go
type HealthChecker struct {
    broker *Broker
    topics []string
}

func (hc *HealthChecker) CheckHealth() error {
    for _, topic := range hc.topics {
        if _, exists := hc.broker.topics[topic]; !exists {
            return fmt.Errorf("topic %s does not exist", topic)
        }
    }
    return nil
}
```

### 3. Alerting
```go
type AlertManager struct {
    thresholds map[string]float64
    notifiers  []Notifier
}

func (am *AlertManager) CheckMetrics(metrics Metrics) {
    if metrics.ErrorRate > am.thresholds["error_rate"] {
        am.sendAlert("High error rate detected")
    }
    
    if metrics.QueueLength > am.thresholds["queue_length"] {
        am.sendAlert("Queue length exceeded threshold")
    }
}
```

## 🎓 Advanced Topics

### 1. Message Ordering
- **FIFO Ordering**: First in, first out
- **Partitioned Ordering**: Order within partitions
- **Global Ordering**: Global message order

### 2. Message Durability
- **At Most Once**: Messages may be lost
- **At Least Once**: Messages may be duplicated
- **Exactly Once**: Messages delivered exactly once

### 3. Message Filtering
- **Topic Filtering**: Filter by topic name
- **Content Filtering**: Filter by message content
- **Type Filtering**: Filter by message type

### 4. Message Routing
- **Direct Routing**: Direct topic-to-topic routing
- **Pattern Routing**: Pattern-based routing
- **Conditional Routing**: Conditional message routing

## 🚀 Next Steps

After mastering Pub-Sub patterns:
1. **Practice**: Implement various pub-sub scenarios
2. **Optimize**: Tune performance and reliability
3. **Monitor**: Add comprehensive monitoring and alerting
4. **Scale**: Design for high-scale distributed systems
5. **Next Topic**: Move to Event Loop Pattern

## 📚 Additional Resources

- **Go Concurrency Patterns**: https://golang.org/doc/effective_go.html#concurrency
- **Message Queue Patterns**: https://gobyexample.com/channels
- **Event-Driven Architecture**: https://martinfowler.com/articles/201701-event-driven.html
- **Performance Optimization**: https://golang.org/doc/diagnostics.html

---

**Ready to become a Pub-Sub master? Let's dive into the implementations!** 🚀
