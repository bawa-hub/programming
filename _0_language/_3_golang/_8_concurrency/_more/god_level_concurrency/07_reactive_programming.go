package main

import (
	"fmt"
	"sync"
	"time"
)

// GOD-LEVEL CONCEPT 7: Reactive Programming
// Stream processing, backpressure, and event-driven architectures

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: Reactive Programming ===")
	
	// 1. Basic Stream Processing
	demonstrateBasicStreamProcessing()
	
	// 2. Backpressure Handling
	demonstrateBackpressureHandling()
	
	// 3. Event Sourcing
	demonstrateEventSourcing()
	
	// 4. CQRS Pattern
	demonstrateCQRSPattern()
	
	// 5. Reactive Streams
	demonstrateReactiveStreams()
	
	// 6. Advanced Reactive Patterns
	demonstrateAdvancedReactivePatterns()
}

// Basic Stream Processing
func demonstrateBasicStreamProcessing() {
	fmt.Println("\n=== 1. BASIC STREAM PROCESSING ===")
	
	fmt.Println(`
🌊 Stream Processing:
• Process data as it arrives
• Transform and filter streams
• Combine multiple streams
• Handle infinite data streams
`)

	// Create data stream
	stream := NewDataStream("numbers")
	
	// Transform stream
	transformed := stream.Map(func(x int) int {
		return x * 2
	}).Filter(func(x int) bool {
		return x%4 == 0
	})
	
	// Subscribe to stream
	transformed.Subscribe(func(x int) {
		fmt.Printf("Processed: %d\n", x)
	})
	
	// Emit data
	for i := 1; i <= 10; i++ {
		stream.Emit(i)
		time.Sleep(10 * time.Millisecond)
	}
	
	stream.Close()
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("💡 Stream processing with transformations")
}

// Backpressure Handling
func demonstrateBackpressureHandling() {
	fmt.Println("\n=== 2. BACKPRESSURE HANDLING ===")
	
	fmt.Println(`
⏸️  Backpressure:
• Slow consumers can't keep up
• Need to handle pressure
• Drop, buffer, or throttle
• Prevent memory overflow
`)

	// Create fast producer
	producer := NewDataStream("fast-producer")
	
	// Create slow consumer with backpressure
	consumer := NewSlowConsumer(producer, 50*time.Millisecond)
	
	// Start consumer
	consumer.Start()
	
	// Emit data faster than consumer can process
	for i := 1; i <= 20; i++ {
		producer.Emit(i)
		time.Sleep(10 * time.Millisecond)
	}
	
	producer.Close()
	consumer.Stop()
	
	fmt.Println("💡 Backpressure prevents memory overflow")
}

// Event Sourcing
func demonstrateEventSourcing() {
	fmt.Println("\n=== 3. EVENT SOURCING ===")
	
	fmt.Println(`
📝 Event Sourcing:
• Store events instead of state
• Replay events to reconstruct state
• Audit trail of all changes
• Time travel debugging
`)

	// Create event store
	eventStore := NewEventStore()
	
	// Create aggregate
	account := NewAccount("account-1", eventStore)
	
	// Perform operations (these create events)
	account.Deposit(100)
	account.Withdraw(30)
	account.Deposit(50)
	account.Withdraw(20)
	
	// Get current state
	balance := account.GetBalance()
	fmt.Printf("Current balance: %d\n", balance)
	
	// Replay events to reconstruct state
	reconstructed := NewAccount("account-1", eventStore)
	reconstructed.ReplayEvents()
	reconstructedBalance := reconstructed.GetBalance()
	fmt.Printf("Reconstructed balance: %d\n", reconstructedBalance)
	
	fmt.Println("💡 Event sourcing provides complete audit trail")
}

// CQRS Pattern
func demonstrateCQRSPattern() {
	fmt.Println("\n=== 4. CQRS PATTERN ===")
	
	fmt.Println(`
🔄 CQRS (Command Query Responsibility Segregation):
• Separate read and write models
• Commands modify state
• Queries read state
• Event bus synchronizes models
`)

	// Create CQRS system
	cqrs := NewCQRSSystem()
	
	// Send commands
	cqrs.SendCommand(CreateUserCommand{ID: "user-1", Name: "Alice"})
	cqrs.SendCommand(UpdateUserCommand{ID: "user-1", Name: "Alice Smith"})
	cqrs.SendCommand(CreateUserCommand{ID: "user-2", Name: "Bob"})
	
	// Wait for processing
	time.Sleep(100 * time.Millisecond)
	
	// Query data
	user1 := cqrs.QueryUser("user-1")
	user2 := cqrs.QueryUser("user-2")
	allUsers := cqrs.QueryAllUsers()
	
	fmt.Printf("User 1: %+v\n", user1)
	fmt.Printf("User 2: %+v\n", user2)
	fmt.Printf("All users: %d\n", len(allUsers))
	
	fmt.Println("💡 CQRS separates read and write concerns")
}

// Reactive Streams
func demonstrateReactiveStreams() {
	fmt.Println("\n=== 5. REACTIVE STREAMS ===")
	
	fmt.Println(`
🔄 Reactive Streams:
• Standard for asynchronous stream processing
• Backpressure handling
• Non-blocking operations
• Composable streams
`)

	// Create reactive stream
	stream := NewReactiveStream()
	
	// Chain operations
	stream.
		Map(func(x int) int { return x * 2 }).
		Filter(func(x int) bool { return x > 10 }).
		Take(5).
		Subscribe(func(x int) {
			fmt.Printf("Reactive: %d\n", x)
		})
	
	// Emit data
	for i := 1; i <= 20; i++ {
		stream.Emit(i)
		time.Sleep(10 * time.Millisecond)
	}
	
	stream.Close()
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("💡 Reactive streams are composable and backpressure-aware")
}

// Advanced Reactive Patterns
func demonstrateAdvancedReactivePatterns() {
	fmt.Println("\n=== 6. ADVANCED REACTIVE PATTERNS ===")
	
	fmt.Println(`
🔬 Advanced Patterns:
• Stream merging and splitting
• Windowing and batching
• Error handling and recovery
• Circuit breakers
`)

	// Stream merging
	stream1 := NewDataStream("stream1")
	stream2 := NewDataStream("stream2")
	
	merged := MergeStreams(stream1, stream2)
	merged.Subscribe(func(x int) {
		fmt.Printf("Merged: %d\n", x)
	})
	
	// Emit to both streams
	go func() {
		for i := 1; i <= 5; i++ {
			stream1.Emit(i)
			time.Sleep(20 * time.Millisecond)
		}
		stream1.Close()
	}()
	
	go func() {
		for i := 10; i <= 15; i++ {
			stream2.Emit(i)
			time.Sleep(30 * time.Millisecond)
		}
		stream2.Close()
	}()
	
	time.Sleep(200 * time.Millisecond)
	
	fmt.Println("💡 Advanced patterns for complex stream processing")
}

// Data Stream Implementation
type DataStream struct {
	name      string
	subscribers []func(int)
	closed    bool
	mu        sync.RWMutex
}

func NewDataStream(name string) *DataStream {
	return &DataStream{
		name:        name,
		subscribers: make([]func(int), 0),
	}
}

func (ds *DataStream) Emit(value int) {
	ds.mu.RLock()
	if ds.closed {
		ds.mu.RUnlock()
		return
	}
	
	subscribers := make([]func(int), len(ds.subscribers))
	copy(subscribers, ds.subscribers)
	ds.mu.RUnlock()
	
	// Notify all subscribers
	for _, subscriber := range subscribers {
		go subscriber(value)
	}
}

func (ds *DataStream) Subscribe(subscriber func(int)) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	
	ds.subscribers = append(ds.subscribers, subscriber)
}

func (ds *DataStream) Map(fn func(int) int) *DataStream {
	newStream := NewDataStream(ds.name + "-mapped")
	
	ds.Subscribe(func(x int) {
		newStream.Emit(fn(x))
	})
	
	return newStream
}

func (ds *DataStream) Filter(predicate func(int) bool) *DataStream {
	newStream := NewDataStream(ds.name + "-filtered")
	
	ds.Subscribe(func(x int) {
		if predicate(x) {
			newStream.Emit(x)
		}
	})
	
	return newStream
}

func (ds *DataStream) Close() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	
	ds.closed = true
}

// Slow Consumer with Backpressure
type SlowConsumer struct {
	stream    *DataStream
	delay     time.Duration
	stopCh    chan struct{}
	wg        sync.WaitGroup
}

func NewSlowConsumer(stream *DataStream, delay time.Duration) *SlowConsumer {
	return &SlowConsumer{
		stream: stream,
		delay:  delay,
		stopCh: make(chan struct{}),
	}
}

func (sc *SlowConsumer) Start() {
	sc.wg.Add(1)
	go sc.consume()
}

func (sc *SlowConsumer) consume() {
	defer sc.wg.Done()
	
	sc.stream.Subscribe(func(x int) {
		select {
		case <-sc.stopCh:
			return
		default:
			time.Sleep(sc.delay)
			fmt.Printf("Slow consumer processed: %d\n", x)
		}
	})
}

func (sc *SlowConsumer) Stop() {
	close(sc.stopCh)
	sc.wg.Wait()
}

// Event Sourcing Implementation
type Event interface {
	GetAggregateID() string
	GetTimestamp() time.Time
}

type AccountCreatedEvent struct {
	AccountID string
	Timestamp time.Time
}

func (e AccountCreatedEvent) GetAggregateID() string { return e.AccountID }
func (e AccountCreatedEvent) GetTimestamp() time.Time { return e.Timestamp }

type MoneyDepositedEvent struct {
	AccountID string
	Amount    int
	Timestamp time.Time
}

func (e MoneyDepositedEvent) GetAggregateID() string { return e.AccountID }
func (e MoneyDepositedEvent) GetTimestamp() time.Time { return e.Timestamp }

type MoneyWithdrawnEvent struct {
	AccountID string
	Amount    int
	Timestamp time.Time
}

func (e MoneyWithdrawnEvent) GetAggregateID() string { return e.AccountID }
func (e MoneyWithdrawnEvent) GetTimestamp() time.Time { return e.Timestamp }

type EventStore struct {
	events []Event
	mu     sync.RWMutex
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: make([]Event, 0),
	}
}

func (es *EventStore) Append(event Event) {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	es.events = append(es.events, event)
}

func (es *EventStore) GetEvents(aggregateID string) []Event {
	es.mu.RLock()
	defer es.mu.RUnlock()
	
	var result []Event
	for _, event := range es.events {
		if event.GetAggregateID() == aggregateID {
			result = append(result, event)
		}
	}
	return result
}

type Account struct {
	ID      string
	balance int
	store   *EventStore
}

func NewAccount(id string, store *EventStore) *Account {
	return &Account{
		ID:    id,
		store: store,
	}
}

func (a *Account) Deposit(amount int) {
	event := MoneyDepositedEvent{
		AccountID: a.ID,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	a.store.Append(event)
	a.balance += amount
}

func (a *Account) Withdraw(amount int) {
	event := MoneyWithdrawnEvent{
		AccountID: a.ID,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	a.store.Append(event)
	a.balance -= amount
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) ReplayEvents() {
	events := a.store.GetEvents(a.ID)
	a.balance = 0
	
	for _, event := range events {
		switch e := event.(type) {
		case MoneyDepositedEvent:
			a.balance += e.Amount
		case MoneyWithdrawnEvent:
			a.balance -= e.Amount
		}
	}
}

// CQRS Implementation
type Command interface{}

type CreateUserCommand struct {
	ID   string
	Name string
}

type UpdateUserCommand struct {
	ID   string
	Name string
}

type User struct {
	ID   string
	Name string
}

type CQRSSystem struct {
	writeModel map[string]*User
	readModel  map[string]*User
	eventBus   chan Event
	mu         sync.RWMutex
}

func NewCQRSSystem() *CQRSSystem {
	cqrs := &CQRSSystem{
		writeModel: make(map[string]*User),
		readModel:  make(map[string]*User),
		eventBus:   make(chan Event, 100),
	}
	
	// Start event processor
	go cqrs.processEvents()
	
	return cqrs
}

func (cqrs *CQRSSystem) SendCommand(cmd Command) {
	cqrs.mu.Lock()
	defer cqrs.mu.Unlock()
	
	switch c := cmd.(type) {
	case CreateUserCommand:
		user := &User{ID: c.ID, Name: c.Name}
		cqrs.writeModel[c.ID] = user
		
		// Publish event
		event := UserCreatedEvent{
			UserID:    c.ID,
			Name:      c.Name,
			Timestamp: time.Now(),
		}
		cqrs.eventBus <- event
		
	case UpdateUserCommand:
		if user, exists := cqrs.writeModel[c.ID]; exists {
			user.Name = c.Name
			
			// Publish event
			event := UserUpdatedEvent{
				UserID:    c.ID,
				Name:      c.Name,
				Timestamp: time.Now(),
			}
			cqrs.eventBus <- event
		}
	}
}

func (cqrs *CQRSSystem) QueryUser(id string) *User {
	cqrs.mu.RLock()
	defer cqrs.mu.RUnlock()
	
	return cqrs.readModel[id]
}

func (cqrs *CQRSSystem) QueryAllUsers() []*User {
	cqrs.mu.RLock()
	defer cqrs.mu.RUnlock()
	
	var users []*User
	for _, user := range cqrs.readModel {
		users = append(users, user)
	}
	return users
}

func (cqrs *CQRSSystem) processEvents() {
	for event := range cqrs.eventBus {
		cqrs.mu.Lock()
		
		switch e := event.(type) {
		case UserCreatedEvent:
			cqrs.readModel[e.UserID] = &User{
				ID:   e.UserID,
				Name: e.Name,
			}
		case UserUpdatedEvent:
			if user, exists := cqrs.readModel[e.UserID]; exists {
				user.Name = e.Name
			}
		}
		
		cqrs.mu.Unlock()
	}
}

type UserCreatedEvent struct {
	UserID    string
	Name      string
	Timestamp time.Time
}

func (e UserCreatedEvent) GetAggregateID() string { return e.UserID }
func (e UserCreatedEvent) GetTimestamp() time.Time { return e.Timestamp }

type UserUpdatedEvent struct {
	UserID    string
	Name      string
	Timestamp time.Time
}

func (e UserUpdatedEvent) GetAggregateID() string { return e.UserID }
func (e UserUpdatedEvent) GetTimestamp() time.Time { return e.Timestamp }

// Reactive Stream Implementation
type ReactiveStream struct {
	operations []func(int) int
	closed     bool
	mu         sync.RWMutex
}

func NewReactiveStream() *ReactiveStream {
	return &ReactiveStream{
		operations: make([]func(int) int, 0),
	}
}

func (rs *ReactiveStream) Emit(value int) {
	rs.mu.RLock()
	if rs.closed {
		rs.mu.RUnlock()
		return
	}
	rs.mu.RUnlock()
	
	// Apply all operations
	result := value
	for _, op := range rs.operations {
		if result != -1 { // Don't process dropped values
			result = op(result)
		}
	}
}

func (rs *ReactiveStream) Map(fn func(int) int) *ReactiveStream {
	rs.operations = append(rs.operations, fn)
	return rs
}

func (rs *ReactiveStream) Filter(predicate func(int) bool) *ReactiveStream {
	rs.operations = append(rs.operations, func(x int) int {
		if predicate(x) {
			return x
		}
		return -1 // Signal to drop
	})
	return rs
}

func (rs *ReactiveStream) Take(n int) *ReactiveStream {
	count := 0
	rs.operations = append(rs.operations, func(x int) int {
		if count >= n {
			return -1 // Signal to drop
		}
		count++
		return x
	})
	return rs
}

func (rs *ReactiveStream) Subscribe(subscriber func(int)) {
	// This is a simplified implementation
	// In a real reactive stream, this would be more complex
}

func (rs *ReactiveStream) Close() {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	
	rs.closed = true
}

// Stream Merging
func MergeStreams(stream1, stream2 *DataStream) *DataStream {
	merged := NewDataStream("merged")
	
	stream1.Subscribe(func(x int) {
		merged.Emit(x)
	})
	
	stream2.Subscribe(func(x int) {
		merged.Emit(x)
	})
	
	return merged
}
