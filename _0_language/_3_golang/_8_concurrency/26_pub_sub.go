package main

import (
	"fmt"
	"sync"
	"time"
)

// Message represents a published message
type Message struct {
	Topic   string
	Content string
}

// Subscriber represents a message subscriber
type Subscriber struct {
	ID      string
	Channel chan Message
}

// PubSub implements publish-subscribe pattern
type PubSub struct {
	subscribers map[string][]*Subscriber
	mu          sync.RWMutex
}

// NewPubSub creates a new PubSub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]*Subscriber),
	}
}

// Subscribe adds a subscriber to a topic
func (ps *PubSub) Subscribe(topic string, subscriberID string) *Subscriber {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	subscriber := &Subscriber{
		ID:      subscriberID,
		Channel: make(chan Message, 10),
	}
	
	ps.subscribers[topic] = append(ps.subscribers[topic], subscriber)
	fmt.Printf("Subscriber %s subscribed to topic '%s'\n", subscriberID, topic)
	
	return subscriber
}

// Publish sends a message to all subscribers of a topic
func (ps *PubSub) Publish(topic string, content string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	message := Message{
		Topic:   topic,
		Content: content,
	}
	
	subscribers := ps.subscribers[topic]
	fmt.Printf("Publishing to topic '%s': %s (to %d subscribers)\n", topic, content, len(subscribers))
	
	for _, subscriber := range subscribers {
		select {
		case subscriber.Channel <- message:
		default:
			fmt.Printf("Subscriber %s channel is full, dropping message\n", subscriber.ID)
		}
	}
}

// Unsubscribe removes a subscriber from a topic
func (ps *PubSub) Unsubscribe(topic string, subscriberID string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	subscribers := ps.subscribers[topic]
	for i, sub := range subscribers {
		if sub.ID == subscriberID {
			ps.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(sub.Channel)
			fmt.Printf("Subscriber %s unsubscribed from topic '%s'\n", subscriberID, topic)
			return
		}
	}
}

func main() {
	fmt.Println("=== Pub/Sub Pattern ===")
	
	pubsub := NewPubSub()
	
	// Create subscribers
	sub1 := pubsub.Subscribe("news", "subscriber1")
	sub2 := pubsub.Subscribe("news", "subscriber2")
	sub3 := pubsub.Subscribe("sports", "subscriber3")
	
	// Start subscriber goroutines
	var wg sync.WaitGroup
	
	// Subscriber 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range sub1.Channel {
			fmt.Printf("Subscriber1 received: [%s] %s\n", msg.Topic, msg.Content)
		}
	}()
	
	// Subscriber 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range sub2.Channel {
			fmt.Printf("Subscriber2 received: [%s] %s\n", msg.Topic, msg.Content)
		}
	}()
	
	// Subscriber 3
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range sub3.Channel {
			fmt.Printf("Subscriber3 received: [%s] %s\n", msg.Topic, msg.Content)
		}
	}()
	
	// Publish messages
	time.Sleep(100 * time.Millisecond)
	pubsub.Publish("news", "Breaking: New Go version released!")
	time.Sleep(100 * time.Millisecond)
	pubsub.Publish("sports", "Team wins championship!")
	time.Sleep(100 * time.Millisecond)
	pubsub.Publish("news", "Weather update: Sunny day ahead")
	
	// Unsubscribe one subscriber
	time.Sleep(100 * time.Millisecond)
	pubsub.Unsubscribe("news", "subscriber2")
	
	// Publish more messages
	time.Sleep(100 * time.Millisecond)
	pubsub.Publish("news", "This message won't reach subscriber2")
	
	// Wait a bit then close
	time.Sleep(500 * time.Millisecond)
	
	// Close remaining channels
	close(sub1.Channel)
	close(sub3.Channel)
	
	wg.Wait()
	fmt.Println("Pub/Sub example completed!")
}
