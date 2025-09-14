package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Exercise 1: Basic Pub-Sub Implementation
func Exercise1() {
	fmt.Println("\nExercise 1: Basic Pub-Sub Implementation")
	fmt.Println("=========================================")
	
	// TODO: Implement basic pub-sub
	// 1. Create a broker
	// 2. Create topics
	// 3. Add subscribers
	// 4. Publish messages
	// 5. Handle message delivery
	
	broker := NewBroker()
	broker.CreateTopic("exercise1")
	
	// Add subscriber
	broker.Subscribe("exercise1", func(msg Message) {
		fmt.Printf("  Exercise1 Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	for i := 1; i <= 5; i++ {
		broker.Publish("exercise1", Message{
			ID:        fmt.Sprintf("ex1-%d", i),
			Topic:     "exercise1",
			Data:      fmt.Sprintf("Exercise 1 message %d", i),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Exercise 1 completed: %d subscribers\n", broker.GetSubscriberCount("exercise1"))
}

// Exercise 2: Multiple Topics and Subscribers
func Exercise2() {
	fmt.Println("\nExercise 2: Multiple Topics and Subscribers")
	fmt.Println("===========================================")
	
	// TODO: Implement multiple topics and subscribers
	// 1. Create multiple topics
	// 2. Add multiple subscribers per topic
	// 3. Publish messages to different topics
	// 4. Handle message distribution
	
	broker := NewBroker()
	
	// Create topics
	topics := []string{"topic1", "topic2", "topic3"}
	for _, topic := range topics {
		broker.CreateTopic(topic)
	}
	
	// Add multiple subscribers per topic
	for i, topic := range topics {
		for j := 0; j < 2; j++ {
			subscriberID := fmt.Sprintf("sub-%d-%d", i+1, j+1)
			broker.Subscribe(topic, func(msg Message) {
				fmt.Printf("  %s Handler: %s - %v\n", subscriberID, msg.ID, msg.Data)
			})
		}
	}
	
	// Publish messages
	for i, topic := range topics {
		for j := 1; j <= 3; j++ {
			broker.Publish(topic, Message{
				ID:        fmt.Sprintf("ex2-%d-%d", i+1, j),
				Topic:     topic,
				Data:      fmt.Sprintf("Exercise 2 message for %s", topic),
				Timestamp: time.Now(),
			})
		}
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Exercise 2 completed: %d topics\n", broker.GetTopicCount())
}

// Exercise 3: Message Filtering
func Exercise3() {
	fmt.Println("\nExercise 3: Message Filtering")
	fmt.Println("=============================")
	
	// TODO: Implement message filtering
	// 1. Create a topic with different message types
	// 2. Add subscribers with different filters
	// 3. Publish messages with different characteristics
	// 4. Verify filtering works correctly
	
	broker := NewBroker()
	broker.CreateTopic("filtered")
	
	// Add filtered subscribers
	broker.Subscribe("filtered", func(msg Message) {
		if priority, exists := msg.Metadata["priority"]; exists && priority == "high" {
			fmt.Printf("  High Priority Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	broker.Subscribe("filtered", func(msg Message) {
		if msg.ID[:4] == "user" {
			fmt.Printf("  User Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	broker.Subscribe("filtered", func(msg Message) {
		if data, ok := msg.Data.(string); ok && len(data) > 15 {
			fmt.Printf("  Long Message Handler: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	// Publish messages with different characteristics
	messages := []Message{
		{
			ID:        "user-001",
			Topic:     "filtered",
			Data:      "User action",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "high"},
		},
		{
			ID:        "system-001",
			Topic:     "filtered",
			Data:      "Short message",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "low"},
		},
		{
			ID:        "user-002",
			Topic:     "filtered",
			Data:      "This is a very long message that should trigger the long message handler",
			Timestamp: time.Now(),
			Metadata:  map[string]string{"priority": "medium"},
		},
	}
	
	for _, msg := range messages {
		broker.Publish("filtered", msg)
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Exercise 3 completed: %d subscribers\n", broker.GetSubscriberCount("filtered"))
}

// Exercise 4: Error Handling and Retry
func Exercise4() {
	fmt.Println("\nExercise 4: Error Handling and Retry")
	fmt.Println("====================================")
	
	// TODO: Implement error handling and retry logic
	// 1. Create a topic for error-prone messages
	// 2. Add subscriber with error handling
	// 3. Implement retry logic
	// 4. Handle failed messages
	
	broker := NewBroker()
	broker.CreateTopic("error-prone")
	
	retryCount := 0
	maxRetries := 3
	
	broker.Subscribe("error-prone", func(msg Message) {
		// Simulate random failures
		if rand.Float32() < 0.4 {
			retryCount++
			if retryCount <= maxRetries {
				fmt.Printf("  RETRY %d: Failed to process %s, retrying...\n", retryCount, msg.ID)
				// Simulate retry delay
				time.Sleep(50 * time.Millisecond)
				// Republish for retry
				broker.Publish("error-prone", msg)
			} else {
				fmt.Printf("  FAILED: Message %s failed after %d retries\n", msg.ID, maxRetries)
			}
		} else {
			fmt.Printf("  SUCCESS: Processed message %s\n", msg.ID)
		}
	})
	
	// Publish messages
	for i := 1; i <= 5; i++ {
		broker.Publish("error-prone", Message{
			ID:        fmt.Sprintf("error-msg-%d", i),
			Topic:     "error-prone",
			Data:      fmt.Sprintf("Error-prone message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"retry": "true"},
		})
	}
	
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Exercise 4 completed: %d subscribers\n", broker.GetSubscriberCount("error-prone"))
}

// Exercise 5: Message Ordering
func Exercise5() {
	fmt.Println("\nExercise 5: Message Ordering")
	fmt.Println("============================")
	
	// TODO: Implement message ordering
	// 1. Create a topic for ordered messages
	// 2. Add subscriber that maintains order
	// 3. Publish messages with sequence numbers
	// 4. Verify messages are processed in order
	
	broker := NewBroker()
	broker.CreateTopic("ordered")
	
	lastSequence := 0
	
	broker.Subscribe("ordered", func(msg Message) {
		sequence := 0
		if seq, exists := msg.Metadata["sequence"]; exists {
			fmt.Sscanf(seq, "%d", &sequence)
		}
		
		if sequence == lastSequence+1 {
			fmt.Printf("  Ordered: %s (sequence %d)\n", msg.ID, sequence)
			lastSequence = sequence
		} else {
			fmt.Printf("  Out of order: %s (expected %d, got %d)\n", msg.ID, lastSequence+1, sequence)
		}
	})
	
	// Publish messages in sequence
	for i := 1; i <= 5; i++ {
		broker.Publish("ordered", Message{
			ID:        fmt.Sprintf("ordered-%d", i),
			Topic:     "ordered",
			Data:      fmt.Sprintf("Ordered message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"sequence": fmt.Sprintf("%d", i)},
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Exercise 5 completed: %d subscribers\n", broker.GetSubscriberCount("ordered"))
}

// Exercise 6: Message Batching
func Exercise6() {
	fmt.Println("\nExercise 6: Message Batching")
	fmt.Println("============================")
	
	// TODO: Implement message batching
	// 1. Create a batched publisher
	// 2. Collect messages in batches
	// 3. Process batches when full or timeout
	// 4. Handle batch processing
	
	broker := NewBroker()
	broker.CreateTopic("batched")
	
	// Batched publisher
	batchSize := 3
	batchTimeout := 200 * time.Millisecond
	batch := make([]Message, 0, batchSize)
	var batchMutex sync.Mutex
	
	// Batch processor
	go func() {
		ticker := time.NewTicker(batchTimeout)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				batchMutex.Lock()
				if len(batch) > 0 {
					fmt.Printf("  Batch processed: %d messages\n", len(batch))
					for _, msg := range batch {
						broker.Publish("batched", msg)
					}
					batch = batch[:0]
				}
				batchMutex.Unlock()
			}
		}
	}()
	
	// Add subscriber
	broker.Subscribe("batched", func(msg Message) {
		fmt.Printf("  Batched Message: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Add messages to batch
	for i := 1; i <= 7; i++ {
		batchMutex.Lock()
		batch = append(batch, Message{
			ID:        fmt.Sprintf("batch-%d", i),
			Topic:     "batched",
			Data:      fmt.Sprintf("Batch message %d", i),
			Timestamp: time.Now(),
		})
		
		if len(batch) >= batchSize {
			fmt.Printf("  Batch full, processing: %d messages\n", len(batch))
			for _, msg := range batch {
				broker.Publish("batched", msg)
			}
			batch = batch[:0]
		}
		batchMutex.Unlock()
	}
	
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Exercise 6 completed: %d subscribers\n", broker.GetSubscriberCount("batched"))
}

// Exercise 7: Message Persistence
func Exercise7() {
	fmt.Println("\nExercise 7: Message Persistence")
	fmt.Println("===============================")
	
	// TODO: Implement message persistence
	// 1. Create a persistent broker
	// 2. Store messages before delivery
	// 3. Replay messages on startup
	// 4. Handle message durability
	
	broker := NewBroker()
	broker.CreateTopic("persistent")
	
	// Simulate message storage
	messageStore := make([]Message, 0)
	var storeMutex sync.Mutex
	
	// Persistent publisher
	broker.Subscribe("persistent", func(msg Message) {
		// Store message
		storeMutex.Lock()
		messageStore = append(messageStore, msg)
		storeMutex.Unlock()
		
		fmt.Printf("  Stored Message: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages
	for i := 1; i <= 5; i++ {
		broker.Publish("persistent", Message{
			ID:        fmt.Sprintf("persist-%d", i),
			Topic:     "persistent",
			Data:      fmt.Sprintf("Persistent message %d", i),
			Timestamp: time.Now(),
			Metadata:  map[string]string{"persistent": "true"},
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	
	// Simulate replay
	fmt.Println("  Replaying stored messages:")
	storeMutex.Lock()
	for _, msg := range messageStore {
		fmt.Printf("  Replayed: %s - %v\n", msg.ID, msg.Data)
	}
	storeMutex.Unlock()
	
	fmt.Printf("Exercise 7 completed: %d messages stored\n", len(messageStore))
}

// Exercise 8: Message Routing
func Exercise8() {
	fmt.Println("\nExercise 8: Message Routing")
	fmt.Println("===========================")
	
	// TODO: Implement message routing
	// 1. Create a router that routes messages between topics
	// 2. Set up routing rules
	// 3. Route messages based on content
	// 4. Handle routing failures
	
	broker := NewBroker()
	
	// Create source and destination topics
	broker.CreateTopic("source")
	broker.CreateTopic("user-events")
	broker.CreateTopic("order-events")
	broker.CreateTopic("system-events")
	
	// Router function
	broker.Subscribe("source", func(msg Message) {
		// Route based on message content
		if data, ok := msg.Data.(string); ok {
			if data[:4] == "user" {
				broker.Publish("user-events", msg)
				fmt.Printf("  Routed to user-events: %s\n", msg.ID)
			} else if data[:5] == "order" {
				broker.Publish("order-events", msg)
				fmt.Printf("  Routed to order-events: %s\n", msg.ID)
			} else {
				broker.Publish("system-events", msg)
				fmt.Printf("  Routed to system-events: %s\n", msg.ID)
			}
		}
	})
	
	// Destination handlers
	broker.Subscribe("user-events", func(msg Message) {
		fmt.Printf("  User Event Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("order-events", func(msg Message) {
		fmt.Printf("  Order Event Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	broker.Subscribe("system-events", func(msg Message) {
		fmt.Printf("  System Event Handler: %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages to source
	messages := []string{
		"user created",
		"order placed",
		"system startup",
		"user updated",
		"order shipped",
		"system shutdown",
	}
	
	for i, data := range messages {
		broker.Publish("source", Message{
			ID:        fmt.Sprintf("route-%d", i+1),
			Topic:     "source",
			Data:      data,
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Exercise 8 completed: %d topics\n", broker.GetTopicCount())
}

// Exercise 9: Message Deduplication
func Exercise9() {
	fmt.Println("\nExercise 9: Message Deduplication")
	fmt.Println("=================================")
	
	// TODO: Implement message deduplication
	// 1. Create a deduplicating broker
	// 2. Track message IDs
	// 3. Filter duplicate messages
	// 4. Handle duplicate detection
	
	broker := NewBroker()
	broker.CreateTopic("dedup")
	
	// Track seen message IDs
	seenMessages := make(map[string]bool)
	var seenMutex sync.Mutex
	
	broker.Subscribe("dedup", func(msg Message) {
		seenMutex.Lock()
		if seenMessages[msg.ID] {
			fmt.Printf("  DUPLICATE: Ignoring message %s\n", msg.ID)
			seenMutex.Unlock()
			return
		}
		seenMessages[msg.ID] = true
		seenMutex.Unlock()
		
		fmt.Printf("  UNIQUE: Processing message %s - %v\n", msg.ID, msg.Data)
	})
	
	// Publish messages (some duplicates)
	messageIDs := []string{"msg-1", "msg-2", "msg-1", "msg-3", "msg-2", "msg-4"}
	
	for i, id := range messageIDs {
		broker.Publish("dedup", Message{
			ID:        id,
			Topic:     "dedup",
			Data:      fmt.Sprintf("Message %d", i+1),
			Timestamp: time.Now(),
		})
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Exercise 9 completed: %d unique messages processed\n", len(seenMessages))
}

// Exercise 10: Context and Cancellation
func Exercise10() {
	fmt.Println("\nExercise 10: Context and Cancellation")
	fmt.Println("=====================================")
	
	// TODO: Implement context and cancellation
	// 1. Create a context-aware broker
	// 2. Handle context cancellation
	// 3. Graceful shutdown of subscribers
	// 4. Cleanup resources
	
	broker := NewBroker()
	broker.CreateTopic("context")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Context-aware subscriber
	broker.Subscribe("context", func(msg Message) {
		select {
		case <-ctx.Done():
			fmt.Printf("  Context cancelled, stopping processing\n")
			return
		default:
			fmt.Printf("  Context Message: %s - %v\n", msg.ID, msg.Data)
		}
	})
	
	// Publish messages
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				broker.Publish("context", Message{
					ID:        fmt.Sprintf("ctx-%d", i),
					Topic:     "context",
					Data:      fmt.Sprintf("Context message %d", i),
					Timestamp: time.Now(),
				})
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	// Wait for context timeout
	<-ctx.Done()
	fmt.Printf("Exercise 10 completed: Context cancelled after timeout\n")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Pub-Sub Pattern Exercises")
	fmt.Println("=============================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
