package patterns

import "sync"

// Advanced Pattern 5: Select-based Message Router
type SelectMessageRouter struct {
	routes   map[string][]chan interface{}
	mu       sync.RWMutex
	messageCh chan Message
	quitCh   chan bool
}

type Message struct {
	Topic string
	Data  interface{}
}

func NewSelectMessageRouter() *SelectMessageRouter {
	return &SelectMessageRouter{
		routes:    make(map[string][]chan interface{}),
		messageCh: make(chan Message, 100),
		quitCh:    make(chan bool),
	}
}

func (mr *SelectMessageRouter) Subscribe(topic string) <-chan interface{} {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	ch := make(chan interface{}, 10)
	mr.routes[topic] = append(mr.routes[topic], ch)
	return ch
}

func (mr *SelectMessageRouter) Publish(topic string, data interface{}) {
	select {
	case mr.messageCh <- Message{Topic: topic, Data: data}:
		// Message queued
	default:
		// Router is busy, drop message
	}
}

func (mr *SelectMessageRouter) Start() {
	go mr.run()
}

func (mr *SelectMessageRouter) run() {
	for {
		select {
		case msg := <-mr.messageCh:
			mr.routeMessage(msg)
		case <-mr.quitCh:
			return
		}
	}
}

func (mr *SelectMessageRouter) routeMessage(msg Message) {
	mr.mu.RLock()
	subscribers, exists := mr.routes[msg.Topic]
	mr.mu.RUnlock()
	
	if exists {
		for _, ch := range subscribers {
			select {
			case ch <- msg.Data:
				// Message sent
			default:
				// Subscriber is busy, skip
			}
		}
	}
}

func (mr *SelectMessageRouter) Stop() {
	close(mr.quitCh)
}