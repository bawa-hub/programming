package main

import (
	"sync"
)

// PubSub implements publish-subscribe pattern
type PubSub struct {
	subscribers map[string][]chan Event
	mu          sync.RWMutex
}

// NewPubSub creates a new PubSub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe adds a subscriber to a topic
func (ps *PubSub) Subscribe(topic string) <-chan Event {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	ch := make(chan Event, 10)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	
	return ch
}

// Publish sends an event to all subscribers of a topic
func (ps *PubSub) Publish(topic string, event Event) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	subscribers := ps.subscribers[topic]
	for _, ch := range subscribers {
		select {
		case ch <- event:
		default:
			// Channel is full, skip
		}
	}
}

// Unsubscribe removes a subscriber from a topic
func (ps *PubSub) Unsubscribe(topic string, ch <-chan Event) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	subscribers := ps.subscribers[topic]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			ps.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(subscriber)
			return
		}
	}
}
