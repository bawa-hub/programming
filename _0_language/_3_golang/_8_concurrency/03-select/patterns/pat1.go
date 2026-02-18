package patterns

import (
	"fmt"
	"sync"
)

// Advanced Pattern 1: Select-based Event Loop
type EventLoop struct {
	events    chan Event
	commands  chan Command
	quit      chan bool
	handlers  map[string]func(Event)
	mu        sync.RWMutex
}

type Event struct {
	Type string
	Data interface{}
}

type Command struct {
	Type string
	Data interface{}
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		events:   make(chan Event, 100),
		commands: make(chan Command, 100),
		quit:     make(chan bool),
		handlers: make(map[string]func(Event)),
	}
}

func (el *EventLoop) RegisterHandler(eventType string, handler func(Event)) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.handlers[eventType] = handler
}

func (el *EventLoop) EmitEvent(event Event) {
	select {
	case el.events <- event:
	default:
		// Event channel is full, drop event
	}
}

func (el *EventLoop) SendCommand(cmd Command) {
	select {
	case el.commands <- cmd:
	default:
		// Command channel is full, drop command
	}
}

func (el *EventLoop) Start() {
	go el.run()
}

func (el *EventLoop) run() {
	for {
		select {
		case event := <-el.events:
			el.handleEvent(event)
		case cmd := <-el.commands:
			el.handleCommand(cmd)
		case <-el.quit:
			return
		}
	}
}

func (el *EventLoop) handleEvent(event Event) {
	el.mu.RLock()
	handler, exists := el.handlers[event.Type]
	el.mu.RUnlock()
	
	if exists {
		handler(event)
	}
}

func (el *EventLoop) handleCommand(cmd Command) {
	switch cmd.Type {
	case "quit":
		el.quit <- true
	case "register":
		if data, ok := cmd.Data.(map[string]interface{}); ok {
			if eventType, ok := data["eventType"].(string); ok {
				el.RegisterHandler(eventType, func(e Event) {
					fmt.Printf("Dynamic handler for %s: %v\n", e.Type, e.Data)
				})
			}
		}
	}
}

func (el *EventLoop) Stop() {
	el.quit <- true
}