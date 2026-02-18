package patterns

import "sync"

// Advanced Pattern 4: Channel-based Event Bus
type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(topic string) <-chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	ch := make(chan interface{}, 10)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	
	for _, ch := range eb.subscribers[topic] {
		select {
		case ch <- data:
		default:
			// Channel is full, skip
		}
	}
}

func (eb *EventBus) Unsubscribe(topic string, ch <-chan interface{}) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	subscribers := eb.subscribers[topic]
	for i, sub := range subscribers {
		if sub == ch {
			eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(sub)
			break
		}
	}
}
