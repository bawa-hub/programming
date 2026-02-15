package event

import (
	"fmt"
	"sync"
)

type EventBus struct {
	observers map[string][]EventObserver
	mu        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		observers: make(map[string][]EventObserver),
	}
}

func (eb *EventBus) Subscribe(observer EventObserver) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	for _, eventType := range observer.GetEventTypes() {
		eb.observers[eventType] = append(eb.observers[eventType], observer)
	}
	fmt.Printf("Event observer %s subscribed to events\n", observer.GetID())
}

func (eb *EventBus) Unsubscribe(observer EventObserver) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	for eventType, observers := range eb.observers {
		for i, obs := range observers {
			if obs.GetID() == observer.GetID() {
				eb.observers[eventType] = append(observers[:i], observers[i+1:]...)
				break
			}
		}
	}
	fmt.Printf("Event observer %s unsubscribed from events\n", observer.GetID())
}

func (eb *EventBus) Publish(event *Event) {
	eb.mu.RLock()
	observers := eb.observers[event.Type]
	eb.mu.RUnlock()
	
	fmt.Printf("Publishing event: %s\n", event)
	for _, observer := range observers {
		observer.HandleEvent(event)
	}
}
