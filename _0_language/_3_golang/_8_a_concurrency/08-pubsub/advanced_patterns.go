package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// AdvancedBroker extends the basic broker with advanced features
type AdvancedBroker struct {
	*Broker
	messageStore    map[string][]Message
	deadLetterQueue chan Message
	metrics         *BrokerMetrics
	circuitBreaker  *CircuitBreaker
	mutex           sync.RWMutex
}

// BrokerMetrics tracks broker performance
type BrokerMetrics struct {
	MessagesPublished   int64
	MessagesDelivered   int64
	MessagesFailed      int64
	MessagesDropped     int64
	AverageLatency      time.Duration
	ThroughputPerSecond float64
	LastUpdate          time.Time
}

// CircuitBreaker implements circuit breaker pattern
type CircuitBreaker struct {
	failureCount    int64
	successCount    int64
	threshold       int64
	timeout         time.Duration
	state           CircuitState
	lastFailureTime time.Time
	mutex           sync.RWMutex
}

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// BatchedPublisher implements message batching
type BatchedPublisher struct {
	broker     *AdvancedBroker
	batchSize  int
	timeout    time.Duration
	messages   []Message
	mutex      sync.Mutex
	ticker     *time.Ticker
}

// MessageCompressor handles message compression
type MessageCompressor struct {
	compressionLevel int
}

// PersistentStorage simulates message persistence
type PersistentStorage struct {
	storage map[string][]Message
	mutex   sync.RWMutex
}

// LoadBalancer distributes messages across multiple brokers
type LoadBalancer struct {
	brokers []*AdvancedBroker
	current int
	mutex   sync.Mutex
}

// MessageRouter implements advanced message routing
type MessageRouter struct {
	routes map[string][]string
	broker *AdvancedBroker
}

// DeadLetterQueue handles failed messages
type DeadLetterQueue struct {
	queue   chan Message
	handler func(Message)
}

// NewAdvancedBroker creates an advanced broker
func NewAdvancedBroker() *AdvancedBroker {
	return &AdvancedBroker{
		Broker:          NewBroker(),
		messageStore:    make(map[string][]Message),
		deadLetterQueue: make(chan Message, 1000),
		metrics:         &BrokerMetrics{},
		circuitBreaker:  NewCircuitBreaker(5, 30*time.Second),
	}
}

// PublishWithPersistence publishes a message with persistence
func (ab *AdvancedBroker) PublishWithPersistence(topic string, msg Message) error {
	// Check circuit breaker
	if ab.circuitBreaker.GetState() == StateOpen {
		return fmt.Errorf("circuit breaker is open")
	}
	
	// Store message
	ab.mutex.Lock()
	ab.messageStore[topic] = append(ab.messageStore[topic], msg)
	ab.mutex.Unlock()
	
	// Publish message
	err := ab.Publish(topic, msg)
	if err != nil {
		ab.circuitBreaker.RecordFailure()
		atomic.AddInt64(&ab.metrics.MessagesFailed, 1)
		return err
	}
	
	ab.circuitBreaker.RecordSuccess()
	atomic.AddInt64(&ab.metrics.MessagesPublished, 1)
	return nil
}

// SubscribeWithFilter subscribes with message filtering
func (ab *AdvancedBroker) SubscribeWithFilter(topic string, filter func(Message) bool, handler func(Message)) {
	ab.Subscribe(topic, func(msg Message) {
		if filter(msg) {
			handler(msg)
			atomic.AddInt64(&ab.metrics.MessagesDelivered, 1)
		} else {
			atomic.AddInt64(&ab.metrics.MessagesDropped, 1)
		}
	})
}

// ReplayMessages replays stored messages for a topic
func (ab *AdvancedBroker) ReplayMessages(topic string, handler func(Message)) {
	ab.mutex.RLock()
	messages := ab.messageStore[topic]
	ab.mutex.RUnlock()
	
	for _, msg := range messages {
		handler(msg)
	}
}

// GetMetrics returns broker metrics
func (ab *AdvancedBroker) GetMetrics() *BrokerMetrics {
	ab.mutex.RLock()
	defer ab.mutex.RUnlock()
	
	now := time.Now()
	if ab.metrics.LastUpdate.IsZero() {
		ab.metrics.LastUpdate = now
	}
	
	elapsed := now.Sub(ab.metrics.LastUpdate)
	if elapsed > 0 {
		ab.metrics.ThroughputPerSecond = float64(ab.metrics.MessagesPublished) / elapsed.Seconds()
	}
	
	return ab.metrics
}

// NewCircuitBreaker creates a circuit breaker
func NewCircuitBreaker(threshold int64, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		state:     StateClosed,
	}
}

// GetState returns the current circuit breaker state
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// RecordSuccess records a successful operation
func (cb *CircuitBreaker) RecordSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.successCount++
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failureCount = 0
	}
}

// RecordFailure records a failed operation
func (cb *CircuitBreaker) RecordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	
	if cb.failureCount >= cb.threshold {
		cb.state = StateOpen
	}
}

// ShouldAllow checks if the operation should be allowed
func (cb *CircuitBreaker) ShouldAllow() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
			return true
		}
		return false
	}
	
	return true
}

// NewBatchedPublisher creates a batched publisher
func NewBatchedPublisher(broker *AdvancedBroker, batchSize int, timeout time.Duration) *BatchedPublisher {
	bp := &BatchedPublisher{
		broker:    broker,
		batchSize: batchSize,
		timeout:   timeout,
		messages:  make([]Message, 0, batchSize),
		ticker:    time.NewTicker(timeout),
	}
	
	// Start batch processor
	go bp.processBatches()
	
	return bp
}

// Publish adds a message to the batch
func (bp *BatchedPublisher) Publish(topic string, msg Message) {
	bp.mutex.Lock()
	defer bp.mutex.Unlock()
	
	msg.Topic = topic
	bp.messages = append(bp.messages, msg)
	
	if len(bp.messages) >= bp.batchSize {
		bp.flush()
	}
}

// processBatches processes batches on timeout
func (bp *BatchedPublisher) processBatches() {
	for range bp.ticker.C {
		bp.mutex.Lock()
		if len(bp.messages) > 0 {
			bp.flush()
		}
		bp.mutex.Unlock()
	}
}

// flush sends the current batch
func (bp *BatchedPublisher) flush() {
	if len(bp.messages) == 0 {
		return
	}
	
	batch := make([]Message, len(bp.messages))
	copy(batch, bp.messages)
	bp.messages = bp.messages[:0]
	
	// Send batch
	for _, msg := range batch {
		bp.broker.PublishWithPersistence(msg.Topic, msg)
	}
	
	fmt.Printf("  Batch sent: %d messages\n", len(batch))
}

// NewMessageCompressor creates a message compressor
func NewMessageCompressor(level int) *MessageCompressor {
	return &MessageCompressor{
		compressionLevel: level,
	}
}

// Compress compresses a message
func (mc *MessageCompressor) Compress(msg Message) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	
	var buf bytes.Buffer
	writer, err := gzip.NewWriterLevel(&buf, mc.compressionLevel)
	if err != nil {
		return nil, err
	}
	
	writer.Write(data)
	writer.Close()
	
	return buf.Bytes(), nil
}

// Decompress decompresses a message
func (mc *MessageCompressor) Decompress(data []byte) (Message, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return Message{}, err
	}
	defer reader.Close()
	
	var msg Message
	err = json.NewDecoder(reader).Decode(&msg)
	return msg, err
}

// NewPersistentStorage creates persistent storage
func NewPersistentStorage() *PersistentStorage {
	return &PersistentStorage{
		storage: make(map[string][]Message),
	}
}

// Store stores a message
func (ps *PersistentStorage) Store(topic string, msg Message) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	
	ps.storage[topic] = append(ps.storage[topic], msg)
}

// GetMessages retrieves messages for a topic
func (ps *PersistentStorage) GetMessages(topic string) []Message {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()
	
	return ps.storage[topic]
}

// NewLoadBalancer creates a load balancer
func NewLoadBalancer(brokers []*AdvancedBroker) *LoadBalancer {
	return &LoadBalancer{
		brokers: brokers,
	}
}

// Publish publishes to the least loaded broker
func (lb *LoadBalancer) Publish(topic string, msg Message) error {
	lb.mutex.Lock()
	broker := lb.brokers[lb.current]
	lb.current = (lb.current + 1) % len(lb.brokers)
	lb.mutex.Unlock()
	
	return broker.PublishWithPersistence(topic, msg)
}

// NewMessageRouter creates a message router
func NewMessageRouter(broker *AdvancedBroker) *MessageRouter {
	return &MessageRouter{
		routes: make(map[string][]string),
		broker: broker,
	}
}

// AddRoute adds a routing rule
func (mr *MessageRouter) AddRoute(source, destination string) {
	mr.routes[source] = append(mr.routes[source], destination)
}

// Route routes a message
func (mr *MessageRouter) Route(topic string, msg Message) error {
	destinations := mr.routes[topic]
	for _, dest := range destinations {
		if err := mr.broker.PublishWithPersistence(dest, msg); err != nil {
			return err
		}
	}
	return nil
}

// NewDeadLetterQueue creates a dead letter queue
func NewDeadLetterQueue(handler func(Message)) *DeadLetterQueue {
	dlq := &DeadLetterQueue{
		queue:   make(chan Message, 1000),
		handler: handler,
	}
	
	// Start processing
	go dlq.process()
	
	return dlq
}

// Add adds a message to the dead letter queue
func (dlq *DeadLetterQueue) Add(msg Message) {
	select {
	case dlq.queue <- msg:
	default:
		// Queue is full, drop message
	}
}

// process processes messages in the dead letter queue
func (dlq *DeadLetterQueue) process() {
	for msg := range dlq.queue {
		dlq.handler(msg)
	}
}

// Advanced Pattern 1: Adaptive Scaling
func adaptiveScalingExample() {
	fmt.Println("\n1. Adaptive Scaling")
	fmt.Println("===================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("adaptive")
	
	// Simulate adaptive scaling based on load
	currentLoad := 0
	
	broker.Subscribe("adaptive", func(msg Message) {
		currentLoad++
		fmt.Printf("  Processing message %s (load: %d)\n", msg.ID, currentLoad)
		
		// Simulate processing time based on load
		processingTime := time.Duration(currentLoad) * 10 * time.Millisecond
		time.Sleep(processingTime)
		
		currentLoad--
	})
	
	// Publish messages with varying load
	for i := 1; i <= 15; i++ {
		broker.PublishWithPersistence("adaptive", Message{
			ID:        fmt.Sprintf("adaptive-%d", i),
			Topic:     "adaptive",
			Data:      fmt.Sprintf("Adaptive message %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Adaptive scaling completed: %d messages processed\n", 15)
}

// Advanced Pattern 2: Circuit Breaker
func circuitBreakerExample() {
	fmt.Println("\n2. Circuit Breaker")
	fmt.Println("==================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("circuit")
	
	// Subscriber that fails frequently
	broker.Subscribe("circuit", func(msg Message) {
		// Simulate 60% failure rate
		if rand.Float32() < 0.6 {
			fmt.Printf("  FAILED: Message %s failed\n", msg.ID)
			return
		}
		fmt.Printf("  SUCCESS: Message %s processed\n", msg.ID)
	})
	
	// Publish messages
	for i := 1; i <= 10; i++ {
		err := broker.PublishWithPersistence("circuit", Message{
			ID:        fmt.Sprintf("circuit-%d", i),
			Topic:     "circuit",
			Data:      fmt.Sprintf("Circuit breaker message %d", i),
			Timestamp: time.Now(),
		})
		
		if err != nil {
			fmt.Printf("  Circuit breaker blocked message %d: %v\n", i, err)
		}
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Circuit breaker completed: State = %v\n", broker.circuitBreaker.GetState())
}

// Advanced Pattern 3: Message Batching
func messageBatchingExample() {
	fmt.Println("\n3. Message Batching")
	fmt.Println("===================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("batched")
	
	// Create batched publisher
	bp := NewBatchedPublisher(broker, 3, 200*time.Millisecond)
	
	// Add subscriber
	broker.Subscribe("batched", func(msg Message) {
		fmt.Printf("  Batched Message: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	for i := 1; i <= 7; i++ {
		bp.Publish("batched", Message{
			ID:        fmt.Sprintf("batch-%d", i),
			Topic:     "batched",
			Data:      fmt.Sprintf("Batch message %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Message batching completed\n")
}

// Advanced Pattern 4: Message Compression
func messageCompressionExample() {
	fmt.Println("\n4. Message Compression")
	fmt.Println("======================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("compressed")
	
	// Create compressor
	compressor := NewMessageCompressor(gzip.BestCompression)
	
	broker.Subscribe("compressed", func(msg Message) {
		// Simulate compression
		compressed, err := compressor.Compress(msg)
		if err != nil {
			fmt.Printf("  Compression failed for %s: %v\n", msg.ID, err)
			return
		}
		
		// Simulate decompression
		decompressed, err := compressor.Decompress(compressed)
		if err != nil {
			fmt.Printf("  Decompression failed for %s: %v\n", msg.ID, err)
			return
		}
		
		fmt.Printf("  Compressed Message: %s (compressed size: %d bytes)\n", 
			decompressed.ID, len(compressed))
	})
	
	// Publish messages
	for i := 1; i <= 5; i++ {
		broker.PublishWithPersistence("compressed", Message{
			ID:        fmt.Sprintf("compress-%d", i),
			Topic:     "compressed",
			Data:      fmt.Sprintf("This is a long message that will benefit from compression: %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Message compression completed\n")
}

// Advanced Pattern 5: Load Balancing
func loadBalancingExample() {
	fmt.Println("\n5. Load Balancing")
	fmt.Println("=================")
	
	// Create multiple brokers
	brokers := make([]*AdvancedBroker, 3)
	for i := 0; i < 3; i++ {
		brokers[i] = NewAdvancedBroker()
		brokers[i].CreateTopic("load-balanced")
	}
	
	// Create load balancer
	lb := NewLoadBalancer(brokers)
	
	// Add subscribers to each broker
	for i, broker := range brokers {
		broker.Subscribe("load-balanced", func(msg Message) {
			fmt.Printf("  Broker %d: %s - %v\n", i+1, msg.ID, msg.Data)
		})
	}
	
	// Publish messages through load balancer
	for i := 1; i <= 9; i++ {
		lb.Publish("load-balanced", Message{
			ID:        fmt.Sprintf("lb-%d", i),
			Topic:     "load-balanced",
			Data:      fmt.Sprintf("Load balanced message %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Load balancing completed: %d brokers\n", len(brokers))
}

// Advanced Pattern 6: Message Routing
func messageRoutingExample() {
	fmt.Println("\n6. Message Routing")
	fmt.Println("==================")
	
	broker := NewAdvancedBroker()
	
	// Create topics
	broker.CreateTopic("source")
	broker.CreateTopic("user-events")
	broker.CreateTopic("order-events")
	broker.CreateTopic("system-events")
	
	// Create router
	router := NewMessageRouter(broker)
	router.AddRoute("source", "user-events")
	router.AddRoute("source", "order-events")
	router.AddRoute("source", "system-events")
	
	// Add subscribers
	broker.Subscribe("user-events", func(msg Message) {
		fmt.Printf("  User Events: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("order-events", func(msg Message) {
		fmt.Printf("  Order Events: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("system-events", func(msg Message) {
		fmt.Printf("  System Events: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages to source
	messages := []string{
		"user created",
		"order placed",
		"system startup",
		"user updated",
		"order shipped",
	}
	
	for i, data := range messages {
		msg := Message{
			ID:        fmt.Sprintf("route-%d", i+1),
			Topic:     "source",
			Data:      data,
			Timestamp: time.Now(),
		}
		
		// Route message
		router.Route("source", msg)
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Message routing completed: %d topics\n", broker.GetTopicCount())
}

// Advanced Pattern 7: Dead Letter Queue
func deadLetterQueueExample() {
	fmt.Println("\n7. Dead Letter Queue")
	fmt.Println("====================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("dlq-test")
	
	// Create dead letter queue
	dlq := NewDeadLetterQueue(func(msg Message) {
		fmt.Printf("  DLQ Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Subscriber that fails for certain messages
	broker.Subscribe("dlq-test", func(msg Message) {
		// Simulate failure for messages with "fail" in data
		if data, ok := msg.Data.(string); ok && data[:4] == "fail" {
			fmt.Printf("  FAILED: Message %s failed, sending to DLQ\n", msg.ID)
			dlq.Add(msg)
			return
		}
		
		fmt.Printf("  SUCCESS: Message %s processed\n", msg.ID)
	})
	
	// Publish messages
	messages := []string{
		"success message 1",
		"fail message 1",
		"success message 2",
		"fail message 2",
		"success message 3",
	}
	
	for i, data := range messages {
		broker.PublishWithPersistence("dlq-test", Message{
			ID:        fmt.Sprintf("dlq-%d", i+1),
			Topic:     "dlq-test",
			Data:      data,
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Dead letter queue completed\n")
}

// Advanced Pattern 8: Metrics and Monitoring
func metricsMonitoringExample() {
	fmt.Println("\n8. Metrics and Monitoring")
	fmt.Println("=========================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("metrics")
	
	// Add subscriber
	broker.Subscribe("metrics", func(msg Message) {
		// Simulate processing
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	})
	
	// Publish messages
	for i := 1; i <= 20; i++ {
		broker.PublishWithPersistence("metrics", Message{
			ID:        fmt.Sprintf("metrics-%d", i),
			Topic:     "metrics",
			Data:      fmt.Sprintf("Metrics message %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(500 * time.Millisecond)
	
	// Display metrics
	metrics := broker.GetMetrics()
	fmt.Printf("  Messages Published: %d\n", metrics.MessagesPublished)
	fmt.Printf("  Messages Delivered: %d\n", metrics.MessagesDelivered)
	fmt.Printf("  Messages Failed: %d\n", metrics.MessagesFailed)
	fmt.Printf("  Messages Dropped: %d\n", metrics.MessagesDropped)
	fmt.Printf("  Throughput: %.2f messages/sec\n", metrics.ThroughputPerSecond)
	
	fmt.Printf("Metrics and monitoring completed\n")
}

// Advanced Pattern 9: Message Ordering
func messageOrderingExample() {
	fmt.Println("\n9. Message Ordering")
	fmt.Println("===================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("ordered")
	
	// Track message order
	expectedSequence := 1
	var orderMutex sync.Mutex
	
	broker.Subscribe("ordered", func(msg Message) {
		orderMutex.Lock()
		defer orderMutex.Unlock()
		
		sequence := 0
		if seq, exists := msg.Metadata["sequence"]; exists {
			fmt.Sscanf(seq, "%d", &sequence)
		}
		
		if sequence == expectedSequence {
			fmt.Printf("  Ordered: %s (sequence %d)\n", msg.ID, sequence)
			expectedSequence++
		} else {
			fmt.Printf("  Out of order: %s (expected %d, got %d)\n", 
				msg.ID, expectedSequence, sequence)
		}
	})
	
	// Publish messages in sequence
	for i := 1; i <= 5; i++ {
		broker.PublishWithPersistence("ordered", Message{
			ID:        fmt.Sprintf("ordered-%d", i),
			Topic:     "ordered",
			Data:      fmt.Sprintf("Ordered message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"sequence": fmt.Sprintf("%d", i)},
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Message ordering completed\n")
}

// Advanced Pattern 10: Message Deduplication
func messageDeduplicationExample() {
	fmt.Println("\n10. Message Deduplication")
	fmt.Println("=========================")
	
	broker := NewAdvancedBroker()
	broker.CreateTopic("dedup")
	
	// Track seen message IDs
	seenMessages := make(map[string]bool)
	var seenMutex sync.Mutex
	
	broker.Subscribe("dedup", func(msg Message) {
		seenMutex.Lock()
		defer seenMutex.Unlock()
		
		if seenMessages[msg.ID] {
			fmt.Printf("  DUPLICATE: Ignoring message %s\n", msg.ID)
			return
		}
		
		seenMessages[msg.ID] = true
		fmt.Printf("  UNIQUE: Processing message %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages (some duplicates)
	messageIDs := []string{"msg-1", "msg-2", "msg-1", "msg-3", "msg-2", "msg-4", "msg-1"}
	
	for i, id := range messageIDs {
		broker.PublishWithPersistence("dedup", Message{
			ID:        id,
			Topic:     "dedup",
			Data:      fmt.Sprintf("Message %d", i+1),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Message deduplication completed: %d unique messages\n", len(seenMessages))
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Pub-Sub Patterns")
	fmt.Println("=============================")
	
	adaptiveScalingExample()
	circuitBreakerExample()
	messageBatchingExample()
	messageCompressionExample()
	loadBalancingExample()
	messageRoutingExample()
	deadLetterQueueExample()
	metricsMonitoringExample()
	messageOrderingExample()
	messageDeduplicationExample()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
