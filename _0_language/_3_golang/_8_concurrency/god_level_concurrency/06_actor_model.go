package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// GOD-LEVEL CONCEPT 6: Actor Model Implementation
// Message-passing architectures for scalable concurrent systems

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: Actor Model Implementation ===")
	
	// 1. Basic Actor System
	demonstrateBasicActorSystem()
	
	// 2. Actor Supervision
	demonstrateActorSupervision()
	
	// 3. Actor Communication Patterns
	demonstrateActorCommunication()
	
	// 4. Fault Isolation and Recovery
	demonstrateFaultIsolation()
	
	// 5. Actor Pool and Load Balancing
	demonstrateActorPool()
	
	// 6. Advanced Actor Patterns
	demonstrateAdvancedActorPatterns()
}

// Basic Actor System
func demonstrateBasicActorSystem() {
	fmt.Println("\n=== 1. BASIC ACTOR SYSTEM ===")
	
	fmt.Println(`
🎭 Actor Model:
• Actors are independent units of computation
• Communicate only through messages
• No shared state between actors
• Fault isolation and recovery
`)

	// Create a simple actor
	actor := NewActor("calculator")
	
	// Send messages to actor
	actor.Send(AddMessage{Value: 10})
	actor.Send(AddMessage{Value: 5})
	actor.Send(GetValueMessage{})
	
	// Wait for processing
	time.Sleep(100 * time.Millisecond)
	
	// Stop actor
	actor.Stop()
	
	fmt.Println("💡 Basic actor system with message passing")
}

// Actor Supervision
func demonstrateActorSupervision() {
	fmt.Println("\n=== 2. ACTOR SUPERVISION ===")
	
	fmt.Println(`
👨‍💼 Actor Supervision:
• Supervisors monitor child actors
• Restart failed actors
• Implement supervision strategies
• Fault tolerance and recovery
`)

	// Create supervisor
	supervisor := NewSupervisor("main-supervisor")
	
	// Create child actors
	child1 := NewActor("worker-1")
	child2 := NewActor("worker-2")
	
	// Register children with supervisor
	supervisor.AddChild(child1)
	supervisor.AddChild(child2)
	
	// Start supervisor
	supervisor.Start()
	
	// Send work to children
	child1.Send(WorkMessage{Task: "task-1"})
	child2.Send(WorkMessage{Task: "task-2"})
	
	// Simulate failure
	child1.Send(PanicMessage{})
	
	// Wait for recovery
	time.Sleep(200 * time.Millisecond)
	
	// Stop supervisor
	supervisor.Stop()
	
	fmt.Println("💡 Supervisor restarts failed actors")
}

// Actor Communication Patterns
func demonstrateActorCommunication() {
	fmt.Println("\n=== 3. ACTOR COMMUNICATION PATTERNS ===")
	
	fmt.Println(`
📡 Communication Patterns:
• Request-Response
• Publish-Subscribe
• Ask Pattern
• Tell Pattern
`)

	// Request-Response pattern
	requestResponsePattern()
	
	// Publish-Subscribe pattern
	publishSubscribePattern()
	
	// Ask pattern
	askPattern()
}

func requestResponsePattern() {
	fmt.Println("\n--- Request-Response Pattern ---")
	
	// Create service actor
	service := NewActor("service")
	service.SetBehavior(func(msg Message) {
		switch m := msg.(type) {
		case RequestMessage:
			// Process request and send response
			response := ResponseMessage{
				RequestID: m.RequestID,
				Result:    fmt.Sprintf("Processed: %s", m.Data),
			}
			service.Send(response)
		}
	})
	
	// Create client actor
	client := NewActor("client")
	client.SetBehavior(func(msg Message) {
		switch m := msg.(type) {
		case ResponseMessage:
			fmt.Printf("Client received: %s\n", m.Result)
		}
	})
	
	// Send request
	request := RequestMessage{
		RequestID: "req-1",
		Data:      "Hello World",
	}
	service.Send(request)
	
	time.Sleep(50 * time.Millisecond)
	service.Stop()
	client.Stop()
}

func publishSubscribePattern() {
	fmt.Println("\n--- Publish-Subscribe Pattern ---")
	
	// Create publisher
	publisher := NewActor("publisher")
	
	// Create subscribers
	subscriber1 := NewActor("subscriber-1")
	subscriber2 := NewActor("subscriber-2")
	
	// Set up pub/sub
	pubsub := NewPubSub()
	pubsub.Subscribe("events", subscriber1)
	pubsub.Subscribe("events", subscriber2)
	
	// Publish events
	pubsub.Publish("events", EventMessage{Type: "user-login", Data: "user123"})
	pubsub.Publish("events", EventMessage{Type: "user-logout", Data: "user123"})
	
	time.Sleep(50 * time.Millisecond)
	publisher.Stop()
	subscriber1.Stop()
	subscriber2.Stop()
}

func askPattern() {
	fmt.Println("\n--- Ask Pattern ---")
	
	// Create actor
	actor := NewActor("ask-actor")
	actor.SetBehavior(func(msg Message) {
		switch m := msg.(type) {
		case AskMessage:
			// Process and respond
			response := AskResponseMessage{
				AskID:  m.AskID,
				Result: fmt.Sprintf("Processed: %s", m.Data),
			}
			actor.Send(response)
		}
	})
	
	// Ask for result
	result := actor.Ask(AskMessage{
		AskID: "ask-1",
		Data:  "Hello Ask",
	})
	
	fmt.Printf("Ask result: %s\n", result)
	actor.Stop()
}

// Fault Isolation and Recovery
func demonstrateFaultIsolation() {
	fmt.Println("\n=== 4. FAULT ISOLATION AND RECOVERY ===")
	
	fmt.Println(`
🛡️  Fault Isolation:
• Actors fail independently
• Supervisors handle failures
• Different restart strategies
• Circuit breaker pattern
`)

	// Create fault-tolerant system
	system := NewFaultTolerantSystem()
	
	// Add actors
	actor1 := NewActor("critical-actor")
	actor2 := NewActor("non-critical-actor")
	
	system.AddActor(actor1, true)  // Critical
	system.AddActor(actor2, false) // Non-critical
	
	// Start system
	system.Start()
	
	// Simulate failures
	actor1.Send(PanicMessage{})
	actor2.Send(PanicMessage{})
	
	// Wait for recovery
	time.Sleep(200 * time.Millisecond)
	
	// Stop system
	system.Stop()
	
	fmt.Println("💡 Fault isolation prevents cascade failures")
}

// Actor Pool and Load Balancing
func demonstrateActorPool() {
	fmt.Println("\n=== 5. ACTOR POOL AND LOAD BALANCING ===")
	
	fmt.Println(`
⚖️  Actor Pool:
• Pool of actors for load balancing
• Round-robin distribution
• Dynamic scaling
• Work distribution
`)

	// Create actor pool
	pool := NewActorPool("worker-pool", 3)
	
	// Start pool
	pool.Start()
	
	// Send work to pool
	for i := 0; i < 10; i++ {
		work := WorkMessage{Task: fmt.Sprintf("task-%d", i)}
		pool.Send(work)
	}
	
	// Wait for processing
	time.Sleep(100 * time.Millisecond)
	
	// Stop pool
	pool.Stop()
	
	fmt.Println("💡 Actor pool distributes work efficiently")
}

// Advanced Actor Patterns
func demonstrateAdvancedActorPatterns() {
	fmt.Println("\n=== 6. ADVANCED ACTOR PATTERNS ===")
	
	fmt.Println(`
🔬 Advanced Patterns:
• Actor hierarchies
• State machines
• Event sourcing
• CQRS pattern
`)

	// State machine actor
	stateMachine := NewStateMachineActor("state-machine")
	stateMachine.Start()
	
	// Send state transitions
	stateMachine.Send(StateTransitionMessage{To: "idle"})
	stateMachine.Send(StateTransitionMessage{To: "working"})
	stateMachine.Send(StateTransitionMessage{To: "idle"})
	
	time.Sleep(50 * time.Millisecond)
	stateMachine.Stop()
	
	fmt.Println("💡 Advanced patterns for complex systems")
}

// Message Types
type Message interface{}

type AddMessage struct {
	Value int
}

type GetValueMessage struct{}

type WorkMessage struct {
	Task string
}

type PanicMessage struct{}

type RequestMessage struct {
	RequestID string
	Data      string
}

type ResponseMessage struct {
	RequestID string
	Result    string
}

type EventMessage struct {
	Type string
	Data string
}

type AskMessage struct {
	AskID string
	Data  string
}

type AskResponseMessage struct {
	AskID  string
	Result string
}

type StateTransitionMessage struct {
	To string
}

// Basic Actor Implementation
type Actor struct {
	name     string
	mailbox  chan Message
	behavior func(Message)
	context  context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func NewActor(name string) *Actor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Actor{
		name:    name,
		mailbox: make(chan Message, 100),
		context: ctx,
		cancel:  cancel,
	}
}

func (a *Actor) SetBehavior(behavior func(Message)) {
	a.behavior = behavior
}

func (a *Actor) Send(msg Message) {
	select {
	case a.mailbox <- msg:
	case <-a.context.Done():
		// Actor is stopped
	}
}

func (a *Actor) Start() {
	a.wg.Add(1)
	go a.run()
}

func (a *Actor) run() {
	defer a.wg.Done()
	
	for {
		select {
		case msg := <-a.mailbox:
			if a.behavior != nil {
				func() {
					defer func() {
						if r := recover(); r != nil {
							log.Printf("Actor %s panicked: %v", a.name, r)
						}
					}()
					a.behavior(msg)
				}()
			}
		case <-a.context.Done():
			return
		}
	}
}

func (a *Actor) Stop() {
	a.cancel()
	a.wg.Wait()
}

func (a *Actor) Ask(msg AskMessage) string {
	responseCh := make(chan AskResponseMessage, 1)
	
	// Set up response handler
	originalBehavior := a.behavior
	a.behavior = func(m Message) {
		switch resp := m.(type) {
		case AskResponseMessage:
			if resp.AskID == msg.AskID {
				responseCh <- resp
				return
			}
		}
		if originalBehavior != nil {
			originalBehavior(m)
		}
	}
	
	// Send ask message
	a.Send(msg)
	
	// Wait for response
	select {
	case response := <-responseCh:
		return response.Result
	case <-time.After(5 * time.Second):
		return "timeout"
	}
}

// Supervisor Implementation
type Supervisor struct {
	name     string
	children []*Actor
	context  context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func NewSupervisor(name string) *Supervisor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Supervisor{
		name:    name,
		context: ctx,
		cancel:  cancel,
	}
}

func (s *Supervisor) AddChild(actor *Actor) {
	s.children = append(s.children, actor)
}

func (s *Supervisor) Start() {
	for _, child := range s.children {
		child.Start()
	}
	
	s.wg.Add(1)
	go s.monitor()
}

func (s *Supervisor) monitor() {
	defer s.wg.Done()
	
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			// Check if any child needs restart
			for _, child := range s.children {
				select {
				case <-child.context.Done():
					// Child stopped, restart it
					log.Printf("Supervisor restarting child: %s", child.name)
					child.Start()
				default:
					// Child is running
				}
			}
		case <-s.context.Done():
			return
		}
	}
}

func (s *Supervisor) Stop() {
	s.cancel()
	for _, child := range s.children {
		child.Stop()
	}
	s.wg.Wait()
}

// PubSub Implementation
type PubSub struct {
	subscribers map[string][]*Actor
	mu          sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]*Actor),
	}
}

func (ps *PubSub) Subscribe(topic string, actor *Actor) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	ps.subscribers[topic] = append(ps.subscribers[topic], actor)
}

func (ps *PubSub) Publish(topic string, msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	if subscribers, exists := ps.subscribers[topic]; exists {
		for _, actor := range subscribers {
			actor.Send(msg)
		}
	}
}

// Fault Tolerant System
type FaultTolerantSystem struct {
	actors     map[string]*Actor
	critical   map[string]bool
	supervisor *Supervisor
}

func NewFaultTolerantSystem() *FaultTolerantSystem {
	return &FaultTolerantSystem{
		actors:   make(map[string]*Actor),
		critical: make(map[string]bool),
	}
}

func (fts *FaultTolerantSystem) AddActor(actor *Actor, critical bool) {
	fts.actors[actor.name] = actor
	fts.critical[actor.name] = critical
}

func (fts *FaultTolerantSystem) Start() {
	fts.supervisor = NewSupervisor("fault-tolerant-supervisor")
	
	for _, actor := range fts.actors {
		fts.supervisor.AddChild(actor)
	}
	
	fts.supervisor.Start()
}

func (fts *FaultTolerantSystem) Stop() {
	if fts.supervisor != nil {
		fts.supervisor.Stop()
	}
}

// Actor Pool Implementation
type ActorPool struct {
	name     string
	actors   []*Actor
	roundRobin int
	mu       sync.Mutex
}

func NewActorPool(name string, size int) *ActorPool {
	pool := &ActorPool{
		name:   name,
		actors: make([]*Actor, size),
	}
	
	for i := 0; i < size; i++ {
		actor := NewActor(fmt.Sprintf("%s-worker-%d", name, i))
		actor.SetBehavior(func(msg Message) {
			switch m := msg.(type) {
			case WorkMessage:
				fmt.Printf("Processing %s\n", m.Task)
				time.Sleep(10 * time.Millisecond) // Simulate work
			}
		})
		pool.actors[i] = actor
	}
	
	return pool
}

func (ap *ActorPool) Start() {
	for _, actor := range ap.actors {
		actor.Start()
	}
}

func (ap *ActorPool) Send(msg Message) {
	ap.mu.Lock()
	actor := ap.actors[ap.roundRobin%len(ap.actors)]
	ap.roundRobin++
	ap.mu.Unlock()
	
	actor.Send(msg)
}

func (ap *ActorPool) Stop() {
	for _, actor := range ap.actors {
		actor.Stop()
	}
}

// State Machine Actor
type StateMachineActor struct {
	name      string
	state     string
	actor     *Actor
	behaviors map[string]func(Message)
}

func NewStateMachineActor(name string) *StateMachineActor {
	sma := &StateMachineActor{
		name:      name,
		state:     "idle",
		behaviors: make(map[string]func(Message)),
	}
	
	sma.actor = NewActor(name)
	sma.actor.SetBehavior(sma.handleMessage)
	
	// Define state behaviors
	sma.behaviors["idle"] = func(msg Message) {
		fmt.Printf("State machine in idle state\n")
	}
	
	sma.behaviors["working"] = func(msg Message) {
		fmt.Printf("State machine in working state\n")
	}
	
	return sma
}

func (sma *StateMachineActor) Start() {
	sma.actor.Start()
}

func (sma *StateMachineActor) Send(msg Message) {
	sma.actor.Send(msg)
}

func (sma *StateMachineActor) Stop() {
	sma.actor.Stop()
}

func (sma *StateMachineActor) handleMessage(msg Message) {
	switch m := msg.(type) {
	case StateTransitionMessage:
		sma.state = m.To
		fmt.Printf("State transitioned to: %s\n", sma.state)
	}
	
	if behavior, exists := sma.behaviors[sma.state]; exists {
		behavior(msg)
	}
}
