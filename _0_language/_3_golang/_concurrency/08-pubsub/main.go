package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Message represents a message in the pub-sub system
type Message struct {
	ID        string
	Topic     string
	Data      interface{}
	Timestamp time.Time
	Metadata  map[string]string
}

// Broker manages topics and message routing
type Broker struct {
	topics      map[string]chan Message
	subscribers map[string][]func(Message)
	mutex       sync.RWMutex
}

// NewBroker creates a new message broker
func NewBroker() *Broker {
	return &Broker{
		topics:      make(map[string]chan Message),
		subscribers: make(map[string][]func(Message)),
	}
}

// Publish sends a message to a topic
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
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout publishing to topic %s", topic)
	}
}

// Subscribe adds a subscriber to a topic
func (b *Broker) Subscribe(topic string, handler func(Message)) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	// Create topic if it doesn't exist
	if _, exists := b.topics[topic]; !exists {
		b.topics[topic] = make(chan Message, 100)
	}
	
	// Add subscriber
	b.subscribers[topic] = append(b.subscribers[topic], handler)
	
	// Start message processing for this subscriber
	go func() {
		for msg := range b.topics[topic] {
			handler(msg)
		}
	}()
}

// CreateTopic creates a new topic
func (b *Broker) CreateTopic(topic string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if _, exists := b.topics[topic]; !exists {
		b.topics[topic] = make(chan Message, 100)
		b.subscribers[topic] = make([]func(Message), 0)
	}
}

// GetTopicCount returns the number of topics
func (b *Broker) GetTopicCount() int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return len(b.topics)
}

// GetSubscriberCount returns the number of subscribers for a topic
func (b *Broker) GetSubscriberCount(topic string) int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return len(b.subscribers[topic])
}

// Example 1: Basic Pub-Sub
func basicPubSubExample() {
	fmt.Println("\n1. Basic Pub-Sub")
	fmt.Println("================")
	
	broker := NewBroker()
	
	// Create topics
	broker.CreateTopic("user.events")
	broker.CreateTopic("order.events")
	
	// Subscribe to user events
	broker.Subscribe("user.events", func(msg Message) {
		fmt.Printf("  User Event: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Subscribe to order events
	broker.Subscribe("order.events", func(msg Message) {
		fmt.Printf("  Order Event: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	fmt.Println("Publishing messages:")
	
	// User events
	broker.Publish("user.events", Message{
		ID:        "user-1",
		Topic:     "user.events",
		Data:      "User registered",
		Timestamp: time.Now(),
		Metadata:  map[string]string{"type": "registration"},
	})
	
	broker.Publish("user.events", Message{
		ID:        "user-2",
		Topic:     "user.events",
		Data:      "User logged in",
		Timestamp: time.Now(),
		Metadata:  map[string]string{"type": "login"},
	})
	
	// Order events
	broker.Publish("order.events", Message{
		ID:        "order-1",
		Topic:     "order.events",
		Data:      "Order placed",
		Timestamp: time.Now(),
		Metadata:  map[string]string{"type": "placed"},
	})
	
	broker.Publish("order.events", Message{
		ID:        "order-2",
		Topic:     "order.events",
		Data:      "Order shipped",
		Timestamp: time.Now(),
		Metadata:  map[string]string{"type": "shipped"},
	})
	
	// Wait for messages to be processed
	time.Sleep(100 * time.Millisecond)
	
	fmt.Printf("\nBasic Pub-Sub completed: %d topics, %d user subscribers, %d order subscribers\n",
		broker.GetTopicCount(), broker.GetSubscriberCount("user.events"), broker.GetSubscriberCount("order.events"))
}

// Example 2: Multiple Subscribers
func multipleSubscribersExample() {
	fmt.Println("\n2. Multiple Subscribers")
	fmt.Println("=======================")
	
	broker := NewBroker()
	broker.CreateTopic("notifications")
	
	// Multiple subscribers for the same topic
	broker.Subscribe("notifications", func(msg Message) {
		fmt.Printf("  Email Service: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("notifications", func(msg Message) {
		fmt.Printf("  SMS Service: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("notifications", func(msg Message) {
		fmt.Printf("  Push Service: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	fmt.Println("Publishing notifications:")
	
	for i := 1; i <= 5; i++ {
		broker.Publish("notifications", Message{
			ID:        fmt.Sprintf("notif-%d", i),
			Topic:     "notifications",
			Data:      fmt.Sprintf("Notification %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "high"},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("\nMultiple Subscribers completed: %d subscribers\n", broker.GetSubscriberCount("notifications"))
}

// Example 3: Topic-Based Filtering
func topicBasedFilteringExample() {
	fmt.Println("\n3. Topic-Based Filtering")
	fmt.Println("=========================")
	
	broker := NewBroker()
	
	// Create multiple topics
	topics := []string{"user.created", "user.updated", "user.deleted", "order.placed", "order.shipped"}
	for _, topic := range topics {
		broker.CreateTopic(topic)
	}
	
	// Subscribe to specific topics
	broker.Subscribe("user.created", func(msg Message) {
		fmt.Printf("  User Created Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("user.updated", func(msg Message) {
		fmt.Printf("  User Updated Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("order.placed", func(msg Message) {
		fmt.Printf("  Order Placed Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages to different topics
	fmt.Println("Publishing messages to different topics:")
	
	messages := []struct {
		topic string
		data  string
	}{
		{"user.created", "New user John created"},
		{"user.updated", "User John updated profile"},
		{"user.deleted", "User John deleted"},
		{"order.placed", "Order #123 placed"},
		{"order.shipped", "Order #123 shipped"},
		{"user.created", "New user Jane created"},
		{"order.placed", "Order #124 placed"},
	}
	
	for i, msg := range messages {
		broker.Publish(msg.topic, Message{
			ID:        fmt.Sprintf("msg-%d", i+1),
			Topic:     msg.topic,
			Data:      msg.data,
			Timestamp: time.Now(),
			Metadata:  map[string]string{"source": "api"},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("\nTopic-Based Filtering completed: %d topics\n", broker.GetTopicCount())
}

// Example 4: Content-Based Filtering
func contentBasedFilteringExample() {
	fmt.Println("\n4. Content-Based Filtering")
	fmt.Println("===========================")
	
	broker := NewBroker()
	broker.CreateTopic("events")
	
	// Subscriber with content filtering
	broker.Subscribe("events", func(msg Message) {
		// Filter by metadata
		if priority, exists := msg.Metadata["priority"]; exists && priority == "high" {
			fmt.Printf("  High Priority Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	broker.Subscribe("events", func(msg Message) {
		// Filter by data content
		if data, ok := msg.Data.(string); ok && len(data) > 10 {
			fmt.Printf("  Long Message Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	broker.Subscribe("events", func(msg Message) {
		// Filter by message ID pattern
		if msg.ID[:4] == "user" {
			fmt.Printf("  User Event Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	// Publish messages with different characteristics
	fmt.Println("Publishing messages with different characteristics:")
	
	messages := []Message{
		{
			ID:        "user-001",
			Topic:     "events",
			Data:      "User action",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "high"},
		},
		{
			ID:        "system-001",
			Topic:     "events",
			Data:      "Short message",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "low"},
		},
		{
			ID:        "user-002",
			Topic:     "events",
			Data:      "This is a very long message that should trigger the long message handler",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "medium"},
		},
		{
			ID:        "order-001",
			Topic:     "events",
			Data:      "Order processing",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "high"},
		},
	}
	
	for _, msg := range messages {
		broker.Publish("events", msg)
	}
	
	// Wait for messages to be processed
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("\nContent-Based Filtering completed: %d subscribers\n", broker.GetSubscriberCount("events"))
}

// Example 5: Error Handling
func errorHandlingExample() {
	fmt.Println("\n5. Error Handling")
	fmt.Println("=================")
	
	broker := NewBroker()
	broker.CreateTopic("errors")
	
	// Subscriber that might fail
	broker.Subscribe("errors", func(msg Message) {
		// Simulate random failures
		if rand.Float32() < 0.3 {
			fmt.Printf("  ERROR: Failed to process message %s: %v\n", msg.ID, msg.Data)
			return
		}
		fmt.Printf("  SUCCESS: Processed message %s: %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	fmt.Println("Publishing messages with error handling:")
	
	for i := 1; i <= 10; i++ {
		broker.Publish("errors", Message{
			ID:        fmt.Sprintf("error-msg-%d", i),
			Topic:     "errors",
			Data:      fmt.Sprintf("Message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"retry": "true"},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(300 * time.Millisecond)
	
	fmt.Printf("\nError Handling completed: %d subscribers\n", broker.GetSubscriberCount("errors"))
}

// Example 6: Message Ordering
func basicMessageOrderingExample() {
	fmt.Println("\n6. Message Ordering")
	fmt.Println("===================")
	
	broker := NewBroker()
	broker.CreateTopic("ordered")
	
	// Subscriber that processes messages in order
	broker.Subscribe("ordered", func(msg Message) {
		fmt.Printf("  Ordered Message: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages in sequence
	fmt.Println("Publishing messages in sequence:")
	
	for i := 1; i <= 5; i++ {
		broker.Publish("ordered", Message{
			ID:        fmt.Sprintf("ordered-%d", i),
			Topic:     "ordered",
			Data:      fmt.Sprintf("Sequential message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"sequence": fmt.Sprintf("%d", i)},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("\nMessage Ordering completed: %d subscribers\n", broker.GetSubscriberCount("ordered"))
}

// Example 7: Performance Test
func performanceTestExample() {
	fmt.Println("\n7. Performance Test")
	fmt.Println("===================")
	
	broker := NewBroker()
	broker.CreateTopic("performance")
	
	// Multiple subscribers
	for i := 0; i < 3; i++ {
		broker.Subscribe("performance", func(msg Message) {
			// Simulate processing
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		})
	}
	
	// Publish many messages
	fmt.Println("Publishing 100 messages for performance test:")
	
	start := time.Now()
	
	for i := 1; i <= 100; i++ {
		broker.Publish("performance", Message{
			ID:        fmt.Sprintf("perf-%d", i),
			Topic:     "performance",
			Data:      fmt.Sprintf("Performance message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"test": "performance"},
		})
	}
	
	// Wait for all messages to be processed
	time.Sleep(2 * time.Second)
	
	duration := time.Since(start)
	
	fmt.Printf("\nPerformance Test completed:")
	fmt.Printf("  Messages: 100")
	fmt.Printf("  Duration: %v", duration)
	fmt.Printf("  Throughput: %.2f messages/sec\n", 100.0/duration.Seconds())
	fmt.Printf("  Subscribers: %d\n", broker.GetSubscriberCount("performance"))
}

// Example 8: Wildcard Subscriptions
func wildcardSubscriptionsExample() {
	fmt.Println("\n8. Wildcard Subscriptions")
	fmt.Println("=========================")
	
	broker := NewBroker()
	
	// Create topics with hierarchical structure
	topics := []string{
		"user.events.created",
		"user.events.updated",
		"user.events.deleted",
		"order.events.placed",
		"order.events.shipped",
		"system.events.startup",
		"system.events.shutdown",
	}
	
	for _, topic := range topics {
		broker.CreateTopic(topic)
	}
	
	// Subscribe to all user events (simulated wildcard)
	broker.Subscribe("user.events.created", func(msg Message) {
		fmt.Printf("  User Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	broker.Subscribe("user.events.updated", func(msg Message) {
		fmt.Printf("  User Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	broker.Subscribe("user.events.deleted", func(msg Message) {
		fmt.Printf("  User Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Subscribe to all order events
	broker.Subscribe("order.events.placed", func(msg Message) {
		fmt.Printf("  Order Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	broker.Subscribe("order.events.shipped", func(msg Message) {
		fmt.Printf("  Order Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Subscribe to all system events
	broker.Subscribe("system.events.startup", func(msg Message) {
		fmt.Printf("  System Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	broker.Subscribe("system.events.shutdown", func(msg Message) {
		fmt.Printf("  System Events Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages to different topics
	fmt.Println("Publishing messages to hierarchical topics:")
	
	for i, topic := range topics {
		broker.Publish(topic, Message{
			ID:        fmt.Sprintf("wildcard-%d", i+1),
			Topic:     topic,
			Data:      fmt.Sprintf("Message for %s", topic),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"category": topic[:strings.Index(topic, ".")]},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(300 * time.Millisecond)
	
	fmt.Printf("\nWildcard Subscriptions completed: %d topics\n", broker.GetTopicCount())
}

// Example 9: Message Persistence (Simulated)
func messagePersistenceExample() {
	fmt.Println("\n9. Message Persistence (Simulated)")
	fmt.Println("===================================")
	
	broker := NewBroker()
	broker.CreateTopic("persistent")
	
	// Simulate message persistence
	persistedMessages := make([]Message, 0)
	
	// Subscriber that "persists" messages
	broker.Subscribe("persistent", func(msg Message) {
		persistedMessages = append(persistedMessages, msg)
		fmt.Printf("  Persisted Message: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	fmt.Println("Publishing messages for persistence:")
	
	for i := 1; i <= 5; i++ {
		broker.Publish("persistent", Message{
			ID:        fmt.Sprintf("persist-%d", i),
			Topic:     "persistent",
			Data:      fmt.Sprintf("Persistent message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"persistent": "true"},
		})
	}
	
	// Wait for messages to be processed
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("\nMessage Persistence completed:")
	fmt.Printf("  Messages persisted: %d\n", len(persistedMessages))
	fmt.Printf("  Subscribers: %d\n", broker.GetSubscriberCount("persistent"))
}

// Example 10: Common Pitfalls
func commonPitfallsExample() {
	fmt.Println("\n10. Common Pitfalls")
	fmt.Println("===================")
	
	fmt.Println("Pitfall 1: Message Loss")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// topic := make(chan Message) // Unbuffered channel")
	fmt.Println("// topic <- msg // Can block and lose messages")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// topic := make(chan Message, 100) // Buffered channel")
	fmt.Println("// select {")
	fmt.Println("// case topic <- msg:")
	fmt.Println("// case <-time.After(5 * time.Second):")
	fmt.Println("//     return fmt.Errorf(\"timeout\")")
	fmt.Println("// }")
	
	fmt.Println("\nPitfall 2: Deadlocks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// func (b *Broker) Publish(topic string, msg Message) {")
	fmt.Println("//     b.topics[topic] <- msg // Can block")
	fmt.Println("// }")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// func (b *Broker) Publish(topic string, msg Message) error {")
	fmt.Println("//     select {")
	fmt.Println("//     case b.topics[topic] <- msg:")
	fmt.Println("//         return nil")
	fmt.Println("//     case <-time.After(5 * time.Second):")
	fmt.Println("//         return fmt.Errorf(\"timeout\")")
	fmt.Println("//     }")
	fmt.Println("// }")
	
	fmt.Println("\nPitfall 3: Memory Leaks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// go func() {")
	fmt.Println("//     for msg := range topic {")
	fmt.Println("//         handler(msg)")
	fmt.Println("//     }")
	fmt.Println("//     // Goroutine never exits")
	fmt.Println("// }()")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// go func() {")
	fmt.Println("//     for {")
	fmt.Println("//         select {")
	fmt.Println("//         case msg := <-topic:")
	fmt.Println("//             handler(msg)")
	fmt.Println("//         case <-ctx.Done():")
	fmt.Println("//             return")
	fmt.Println("//         }")
	fmt.Println("//     }")
	fmt.Println("// }()")
	
	fmt.Println("\nPitfall 4: Race Conditions")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// type Broker struct {")
	fmt.Println("//     topics map[string]chan Message")
	fmt.Println("//     // Missing mutex")
	fmt.Println("// }")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// type Broker struct {")
	fmt.Println("//     topics map[string]chan Message")
	fmt.Println("//     mutex  sync.RWMutex")
	fmt.Println("// }")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🚀 Pub-Sub Pattern Examples")
	fmt.Println("============================")
	
	basicPubSubExample()
	multipleSubscribersExample()
	topicBasedFilteringExample()
	contentBasedFilteringExample()
	errorHandlingExample()
	basicMessageOrderingExample()
	performanceTestExample()
	wildcardSubscriptionsExample()
	messagePersistenceExample()
	commonPitfallsExample()
}
