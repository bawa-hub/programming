package basic

import "fmt"

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
	GetState() interface{}
	SetState(state interface{})
}

// Concrete Subject
type ConcreteSubject struct {
	observers []Observer
	state     interface{}
	mu        sync.RWMutex
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
		state:     nil,
	}
}

func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.observers = append(cs.observers, observer)
	fmt.Printf("Observer %s attached\n", observer.GetID())
}

func (cs *ConcreteSubject) Detach(observer Observer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	for i, obs := range cs.observers {
		if obs.GetID() == observer.GetID() {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			fmt.Printf("Observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (cs *ConcreteSubject) Notify() {
	cs.mu.RLock()
	observers := make([]Observer, len(cs.observers))
	copy(observers, cs.observers)
	state := cs.state
	cs.mu.RUnlock()
	
	fmt.Printf("Notifying %d observers\n", len(observers))
	for _, observer := range observers {
		observer.Update(state)
	}
}

func (cs *ConcreteSubject) GetState() interface{} {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.state
}

func (cs *ConcreteSubject) SetState(state interface{}) {
	cs.mu.Lock()
	cs.state = state
	cs.mu.Unlock()
	fmt.Printf("Subject state changed to: %v\n", state)
	cs.Notify()
}