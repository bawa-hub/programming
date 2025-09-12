package actor

import (
	"context"
	"sync"
	"time"
)

// Actor represents an actor in the actor model
type Actor struct {
	id       string
	inbox    chan Message
	behavior func(Message) ActorBehavior
	state    interface{}
	mu       sync.RWMutex
	stopCh   chan struct{}
	ctx      context.Context
	cancel   context.CancelFunc
}

// Message represents a message sent between actors
type Message struct {
	From    string
	To      string
	Type    string
	Payload interface{}
	ReplyTo chan Message
}

// ActorBehavior represents the result of processing a message
type ActorBehavior struct {
	NewState interface{}
	Messages []Message
	Spawn    []Actor
	Stop     bool
}

// NewActor creates a new actor with the given ID and behavior function
func NewActor(id string, behavior func(Message) ActorBehavior) *Actor {
	ctx, cancel := context.WithCancel(context.Background())
	
	actor := &Actor{
		id:       id,
		inbox:    make(chan Message, 1000), // Buffered channel for performance
		behavior: behavior,
		state:    nil,
		stopCh:   make(chan struct{}),
		ctx:      ctx,
		cancel:   cancel,
	}
	
	// Start the actor's message processing loop
	go actor.run()
	
	return actor
}

// run is the main message processing loop
func (a *Actor) run() {
	defer close(a.stopCh)
	
	for {
		select {
		case msg := <-a.inbox:
			// Process the message
			behavior := a.behavior(msg)
			
			// Update state
			a.mu.Lock()
			if behavior.NewState != nil {
				a.state = behavior.NewState
			}
			a.mu.Unlock()
			
			// Send messages to other actors
			for _, outgoingMsg := range behavior.Messages {
				// Set the sender
				outgoingMsg.From = a.id
				// Send message (this would typically go through an actor system)
				// For now, we'll just log it
				_ = outgoingMsg
			}
			
			// Spawn new actors
			for _, newActor := range behavior.Spawn {
				// Start the new actor
				go newActor.run()
			}
			
			// Check if we should stop
			if behavior.Stop {
				return
			}
			
		case <-a.ctx.Done():
			return
		}
	}
}

// Send sends a message to this actor
func (a *Actor) Send(msg Message) error {
	select {
	case a.inbox <- msg:
		return nil
	case <-a.ctx.Done():
		return ErrActorStopped
	default:
		return ErrInboxFull
	}
}

// SendWithReply sends a message and waits for a reply
func (a *Actor) SendWithReply(msg Message, timeout time.Duration) (Message, error) {
	replyCh := make(chan Message, 1)
	msg.ReplyTo = replyCh
	
	if err := a.Send(msg); err != nil {
		return Message{}, err
	}
	
	select {
	case reply := <-replyCh:
		return reply, nil
	case <-time.After(timeout):
		return Message{}, ErrTimeout
	case <-a.ctx.Done():
		return Message{}, ErrActorStopped
	}
}

// GetState returns the current state of the actor
func (a *Actor) GetState() interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.state
}

// SetState sets the state of the actor
func (a *Actor) SetState(state interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.state = state
}

// Stop stops the actor
func (a *Actor) Stop() {
	a.cancel()
	<-a.stopCh
}

// ID returns the actor's ID
func (a *Actor) ID() string {
	return a.id
}

// IsRunning returns true if the actor is running
func (a *Actor) IsRunning() bool {
	select {
	case <-a.ctx.Done():
		return false
	default:
		return true
	}
}

// ActorSystem manages a collection of actors
type ActorSystem struct {
	actors map[string]*Actor
	mu     sync.RWMutex
}

// NewActorSystem creates a new actor system
func NewActorSystem() *ActorSystem {
	return &ActorSystem{
		actors: make(map[string]*Actor),
	}
}

// Spawn creates a new actor in the system
func (as *ActorSystem) Spawn(id string, behavior func(Message) ActorBehavior) *Actor {
	as.mu.Lock()
	defer as.mu.Unlock()
	
	actor := NewActor(id, behavior)
	as.actors[id] = actor
	
	return actor
}

// GetActor returns an actor by ID
func (as *ActorSystem) GetActor(id string) (*Actor, bool) {
	as.mu.RLock()
	defer as.mu.RUnlock()
	
	actor, exists := as.actors[id]
	return actor, exists
}

// Send sends a message to an actor by ID
func (as *ActorSystem) Send(actorID string, msg Message) error {
	as.mu.RLock()
	actor, exists := as.actors[actorID]
	as.mu.RUnlock()
	
	if !exists {
		return ErrActorNotFound
	}
	
	return actor.Send(msg)
}

// Stop stops all actors in the system
func (as *ActorSystem) Stop() {
	as.mu.Lock()
	defer as.mu.Unlock()
	
	for _, actor := range as.actors {
		actor.Stop()
	}
	
	as.actors = make(map[string]*Actor)
}

// StopActor stops a specific actor
func (as *ActorSystem) StopActor(id string) error {
	as.mu.Lock()
	defer as.mu.Unlock()
	
	actor, exists := as.actors[id]
	if !exists {
		return ErrActorNotFound
	}
	
	actor.Stop()
	delete(as.actors, id)
	
	return nil
}

// ListActors returns a list of all actor IDs
func (as *ActorSystem) ListActors() []string {
	as.mu.RLock()
	defer as.mu.RUnlock()
	
	ids := make([]string, 0, len(as.actors))
	for id := range as.actors {
		ids = append(ids, id)
	}
	
	return ids
}

// Actor errors
var (
	ErrActorStopped    = &ActorError{msg: "actor is stopped"}
	ErrInboxFull       = &ActorError{msg: "actor inbox is full"}
	ErrTimeout         = &ActorError{msg: "message timeout"}
	ErrActorNotFound   = &ActorError{msg: "actor not found"}
)

// ActorError represents an actor-related error
type ActorError struct {
	msg string
}

func (e *ActorError) Error() string {
	return e.msg
}
