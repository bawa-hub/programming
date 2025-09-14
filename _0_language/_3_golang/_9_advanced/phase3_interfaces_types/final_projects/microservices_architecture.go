// 🏗️ MICROSERVICES ARCHITECTURE DEMONSTRATION
// A comprehensive microservices system showcasing advanced Go patterns
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// ============================================================================
// CORE INTERFACES FOR MICROSERVICES
// ============================================================================

// Service discovery interface
type ServiceRegistry interface {
	Register(service ServiceInfo) error
	Deregister(serviceID string) error
	Discover(serviceName string) ([]ServiceInfo, error)
	HealthCheck(serviceID string) error
}

// Circuit breaker interface
type CircuitBreaker interface {
	Call(fn func() (interface{}, error)) (interface{}, error)
	GetState() CircuitState
	Reset()
}

// Load balancer interface
type LoadBalancer interface {
	SelectService(services []ServiceInfo) ServiceInfo
	UpdateServices(services []ServiceInfo)
}

// Message queue interface
type MessageQueue interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string, handler MessageHandler) error
	Unsubscribe(topic string, handler MessageHandler) error
}

// Configuration management interface
type ConfigManager interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Watch(key string, callback func(interface{})) error
}

// ============================================================================
// DATA STRUCTURES
// ============================================================================

type ServiceInfo struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Host        string            `json:"host"`
	Port        int               `json:"port"`
	Health      ServiceHealth     `json:"health"`
	Metadata    map[string]string `json:"metadata"`
	LastSeen    time.Time         `json:"last_seen"`
}

type ServiceHealth struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Metrics   map[string]float64 `json:"metrics"`
}

type CircuitState string

const (
	CircuitStateClosed   CircuitState = "closed"
	CircuitStateOpen     CircuitState = "open"
	CircuitStateHalfOpen CircuitState = "half-open"
)

type MessageHandler func(topic string, message []byte) error

// ============================================================================
// SERVICE REGISTRY IMPLEMENTATION
// ============================================================================

type ServiceRegistryImpl struct {
	services map[string]ServiceInfo
	mu       sync.RWMutex
}

func NewServiceRegistry() ServiceRegistry {
	return &ServiceRegistryImpl{
		services: make(map[string]ServiceInfo),
	}
}

func (sr *ServiceRegistryImpl) Register(service ServiceInfo) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	
	service.LastSeen = time.Now()
	sr.services[service.ID] = service
	
	fmt.Printf("✅ Service registered: %s (%s) at %s:%d\n", 
		service.Name, service.ID, service.Host, service.Port)
	
	return nil
}

func (sr *ServiceRegistryImpl) Deregister(serviceID string) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	
	if service, exists := sr.services[serviceID]; exists {
		delete(sr.services, serviceID)
		fmt.Printf("❌ Service deregistered: %s (%s)\n", service.Name, serviceID)
	}
	
	return nil
}

func (sr *ServiceRegistryImpl) Discover(serviceName string) ([]ServiceInfo, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	
	var services []ServiceInfo
	for _, service := range sr.services {
		if service.Name == serviceName && service.Health.Status == "healthy" {
			services = append(services, service)
		}
	}
	
	return services, nil
}

func (sr *ServiceRegistryImpl) HealthCheck(serviceID string) error {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	
	service, exists := sr.services[serviceID]
	if !exists {
		return fmt.Errorf("service not found")
	}
	
	// Simulate health check
	if time.Since(service.LastSeen) > 30*time.Second {
		service.Health.Status = "unhealthy"
		service.Health.Message = "Service not responding"
	} else {
		service.Health.Status = "healthy"
		service.Health.Message = "Service is healthy"
	}
	
	service.Health.Timestamp = time.Now()
	sr.services[serviceID] = service
	
	return nil
}

// ============================================================================
// CIRCUIT BREAKER IMPLEMENTATION
// ============================================================================

type CircuitBreakerImpl struct {
	name           string
	maxFailures    int
	timeout        time.Duration
	state          CircuitState
	failureCount   int
	lastFailTime   time.Time
	mu             sync.RWMutex
}

func NewCircuitBreaker(name string, maxFailures int, timeout time.Duration) CircuitBreaker {
	return &CircuitBreakerImpl{
		name:        name,
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       CircuitStateClosed,
	}
}

func (cb *CircuitBreakerImpl) Call(fn func() (interface{}, error)) (interface{}, error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if circuit should be opened
	if cb.state == CircuitStateOpen {
		if time.Since(cb.lastFailTime) > cb.timeout {
			cb.state = CircuitStateHalfOpen
			fmt.Printf("🔄 Circuit breaker %s: Moving to half-open state\n", cb.name)
		} else {
			return nil, fmt.Errorf("circuit breaker is open")
		}
	}
	
	// Execute the function
	result, err := fn()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailTime = time.Now()
		
		if cb.failureCount >= cb.maxFailures {
			cb.state = CircuitStateOpen
			fmt.Printf("🔴 Circuit breaker %s: Opened due to %d failures\n", cb.name, cb.failureCount)
		}
		
		return nil, err
	}
	
	// Success - reset failure count
	cb.failureCount = 0
	if cb.state == CircuitStateHalfOpen {
		cb.state = CircuitStateClosed
		fmt.Printf("🟢 Circuit breaker %s: Closed after successful call\n", cb.name)
	}
	
	return result, nil
}

func (cb *CircuitBreakerImpl) GetState() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

func (cb *CircuitBreakerImpl) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	cb.state = CircuitStateClosed
	cb.failureCount = 0
	fmt.Printf("🔄 Circuit breaker %s: Reset\n", cb.name)
}

// ============================================================================
// LOAD BALANCER IMPLEMENTATIONS
// ============================================================================

type RoundRobinLoadBalancer struct {
	current int
	mu      sync.Mutex
}

func NewRoundRobinLoadBalancer() LoadBalancer {
	return &RoundRobinLoadBalancer{}
}

func (rr *RoundRobinLoadBalancer) SelectService(services []ServiceInfo) ServiceInfo {
	if len(services) == 0 {
		return ServiceInfo{}
	}
	
	rr.mu.Lock()
	defer rr.mu.Unlock()
	
	service := services[rr.current%len(services)]
	rr.current++
	
	return service
}

func (rr *RoundRobinLoadBalancer) UpdateServices(services []ServiceInfo) {
	// No-op for round robin
}

type WeightedRoundRobinLoadBalancer struct {
	services []ServiceInfo
	weights  []int
	current  int
	mu       sync.Mutex
}

func NewWeightedRoundRobinLoadBalancer() LoadBalancer {
	return &WeightedRoundRobinLoadBalancer{
		weights: make([]int, 0),
	}
}

func (wrr *WeightedRoundRobinLoadBalancer) SelectService(services []ServiceInfo) ServiceInfo {
	if len(services) == 0 {
		return ServiceInfo{}
	}
	
	wrr.mu.Lock()
	defer wrr.mu.Unlock()
	
	// Use weighted selection
	totalWeight := 0
	for _, weight := range wrr.weights {
		totalWeight += weight
	}
	
	if totalWeight == 0 {
		// Fallback to round robin
		service := services[wrr.current%len(services)]
		wrr.current++
		return service
	}
	
	random := rand.Intn(totalWeight)
	currentWeight := 0
	
	for i, service := range services {
		currentWeight += wrr.weights[i]
		if random < currentWeight {
			return service
		}
	}
	
	// Fallback
	return services[0]
}

func (wrr *WeightedRoundRobinLoadBalancer) UpdateServices(services []ServiceInfo) {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()
	
	wrr.services = services
	wrr.weights = make([]int, len(services))
	
	// Assign weights based on service metadata
	for i, service := range services {
		if weight, exists := service.Metadata["weight"]; exists {
			fmt.Sscanf(weight, "%d", &wrr.weights[i])
		} else {
			wrr.weights[i] = 1 // Default weight
		}
	}
}

// ============================================================================
// MESSAGE QUEUE IMPLEMENTATION
// ============================================================================

type MessageQueueImpl struct {
	topics   map[string][]MessageHandler
	mu       sync.RWMutex
}

func NewMessageQueue() MessageQueue {
	return &MessageQueueImpl{
		topics: make(map[string][]MessageHandler),
	}
}

func (mq *MessageQueueImpl) Publish(topic string, message interface{}) error {
	mq.mu.RLock()
	handlers, exists := mq.topics[topic]
	mq.mu.RUnlock()
	
	if !exists {
		return fmt.Errorf("topic %s not found", topic)
	}
	
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	
	// Publish to all handlers
	for _, handler := range handlers {
		go func(h MessageHandler) {
			if err := h(topic, messageBytes); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}(handler)
	}
	
	fmt.Printf("📤 Message published to topic %s\n", topic)
	return nil
}

func (mq *MessageQueueImpl) Subscribe(topic string, handler MessageHandler) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	
	mq.topics[topic] = append(mq.topics[topic], handler)
	fmt.Printf("📥 Subscribed to topic %s\n", topic)
	return nil
}

func (mq *MessageQueueImpl) Unsubscribe(topic string, handler MessageHandler) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	
	handlers := mq.topics[topic]
	for i, h := range handlers {
		if &h == &handler {
			mq.topics[topic] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
	
	fmt.Printf("📤 Unsubscribed from topic %s\n", topic)
	return nil
}

// ============================================================================
// CONFIGURATION MANAGER IMPLEMENTATION
// ============================================================================

type ConfigManagerImpl struct {
	configs  map[string]interface{}
	watchers map[string][]func(interface{})
	mu       sync.RWMutex
}

func NewConfigManager() ConfigManager {
	return &ConfigManagerImpl{
		configs:  make(map[string]interface{}),
		watchers: make(map[string][]func(interface{})),
	}
}

func (cm *ConfigManagerImpl) Get(key string) (interface{}, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	value, exists := cm.configs[key]
	if !exists {
		return nil, fmt.Errorf("config key %s not found", key)
	}
	
	return value, nil
}

func (cm *ConfigManagerImpl) Set(key string, value interface{}) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	cm.configs[key] = value
	
	// Notify watchers
	if watchers, exists := cm.watchers[key]; exists {
		for _, watcher := range watchers {
			go watcher(value)
		}
	}
	
	fmt.Printf("⚙️  Config updated: %s = %v\n", key, value)
	return nil
}

func (cm *ConfigManagerImpl) Watch(key string, callback func(interface{})) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	cm.watchers[key] = append(cm.watchers[key], callback)
	fmt.Printf("👀 Watching config key: %s\n", key)
	return nil
}

// ============================================================================
// MICROSERVICE BASE CLASS
// ============================================================================

type Microservice struct {
	ID           string
	Name         string
	Version      string
	Host         string
	Port         int
	Registry     ServiceRegistry
	Config       ConfigManager
	MessageQueue MessageQueue
	Health       ServiceHealth
	mu           sync.RWMutex
}

func NewMicroservice(name, version, host string, port int, registry ServiceRegistry, config ConfigManager, mq MessageQueue) *Microservice {
	return &Microservice{
		ID:           fmt.Sprintf("%s-%d", name, time.Now().UnixNano()),
		Name:         name,
		Version:      version,
		Host:         host,
		Port:         port,
		Registry:     registry,
		Config:       config,
		MessageQueue: mq,
		Health: ServiceHealth{
			Status:    "healthy",
			Message:   "Service is running",
			Timestamp: time.Now(),
			Metrics:   make(map[string]float64),
		},
	}
}

func (ms *Microservice) Start() error {
	// Register with service registry
	serviceInfo := ServiceInfo{
		ID:       ms.ID,
		Name:     ms.Name,
		Version:  ms.Version,
		Host:     ms.Host,
		Port:     ms.Port,
		Health:   ms.Health,
		Metadata: map[string]string{
			"started_at": time.Now().Format(time.RFC3339),
		},
		LastSeen: time.Now(),
	}
	
	if err := ms.Registry.Register(serviceInfo); err != nil {
		return err
	}
	
	// Start health monitoring
	go ms.healthMonitor()
	
	// Start message processing
	go ms.messageProcessor()
	
	fmt.Printf("🚀 Microservice %s started on %s:%d\n", ms.Name, ms.Host, ms.Port)
	return nil
}

func (ms *Microservice) Stop() error {
	// Deregister from service registry
	if err := ms.Registry.Deregister(ms.ID); err != nil {
		return err
	}
	
	fmt.Printf("🛑 Microservice %s stopped\n", ms.Name)
	return nil
}

func (ms *Microservice) healthMonitor() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		ms.mu.Lock()
		ms.Health.Timestamp = time.Now()
		ms.Health.Metrics["cpu_usage"] = rand.Float64() * 100
		ms.Health.Metrics["memory_usage"] = rand.Float64() * 100
		ms.mu.Unlock()
		
		// Update registry
		serviceInfo := ServiceInfo{
			ID:       ms.ID,
			Name:     ms.Name,
			Version:  ms.Version,
			Host:     ms.Host,
			Port:     ms.Port,
			Health:   ms.Health,
			Metadata: map[string]string{
				"last_health_check": time.Now().Format(time.RFC3339),
			},
			LastSeen: time.Now(),
		}
		
		ms.Registry.Register(serviceInfo)
	}
}

func (ms *Microservice) messageProcessor() {
	// Subscribe to service-specific topics
	topic := fmt.Sprintf("service.%s", ms.Name)
	ms.MessageQueue.Subscribe(topic, func(topic string, message []byte) error {
		fmt.Printf("📨 %s received message: %s\n", ms.Name, string(message))
		return nil
	})
}

func (ms *Microservice) GetHealth() ServiceHealth {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.Health
}

// ============================================================================
// SPECIFIC MICROSERVICE IMPLEMENTATIONS
// ============================================================================

type UserService struct {
	*Microservice
	users map[string]User
	mu    sync.RWMutex
}

func NewUserService(registry ServiceRegistry, config ConfigManager, mq MessageQueue) *UserService {
	ms := NewMicroservice("user-service", "1.0.0", "localhost", 8081, registry, config, mq)
	
	us := &UserService{
		Microservice: ms,
		users:        make(map[string]User),
	}
	
	// Subscribe to user events
	mq.Subscribe("user.events", us.handleUserEvent)
	
	return us
}

func (us *UserService) CreateUser(user User) error {
	us.mu.Lock()
	defer us.mu.Unlock()
	
	user.ID = fmt.Sprintf("user-%d", time.Now().UnixNano())
	user.CreatedAt = time.Now()
	us.users[user.ID] = user
	
	// Publish user created event
	event := map[string]interface{}{
		"event_type": "user.created",
		"user_id":    user.ID,
		"timestamp":  time.Now(),
	}
	
	us.MessageQueue.Publish("user.events", event)
	
	fmt.Printf("👤 User created: %s (%s)\n", user.Name, user.ID)
	return nil
}

func (us *UserService) GetUser(id string) (User, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()
	
	user, exists := us.users[id]
	if !exists {
		return User{}, fmt.Errorf("user not found")
	}
	
	return user, nil
}

func (us *UserService) handleUserEvent(topic string, message []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil {
		return err
	}
	
	fmt.Printf("📨 UserService received event: %s\n", event["event_type"])
	return nil
}

type OrderService struct {
	*Microservice
	orders map[string]Order
	mu     sync.RWMutex
}

func NewOrderService(registry ServiceRegistry, config ConfigManager, mq MessageQueue) *OrderService {
	ms := NewMicroservice("order-service", "1.0.0", "localhost", 8082, registry, config, mq)
	
	os := &OrderService{
		Microservice: ms,
		orders:       make(map[string]Order),
	}
	
	// Subscribe to order events
	mq.Subscribe("order.events", os.handleOrderEvent)
	
	return os
}

func (os *OrderService) CreateOrder(order Order) error {
	os.mu.Lock()
	defer os.mu.Unlock()
	
	order.ID = fmt.Sprintf("order-%d", time.Now().UnixNano())
	order.CreatedAt = time.Now()
	order.Status = "pending"
	os.orders[order.ID] = order
	
	// Publish order created event
	event := map[string]interface{}{
		"event_type": "order.created",
		"order_id":   order.ID,
		"user_id":    order.UserID,
		"total":      order.Total,
		"timestamp":  time.Now(),
	}
	
	os.MessageQueue.Publish("order.events", event)
	
	fmt.Printf("📦 Order created: %s for user %s (Total: $%.2f)\n", order.ID, order.UserID, order.Total)
	return nil
}

func (os *OrderService) GetOrder(id string) (Order, error) {
	os.mu.RLock()
	defer os.mu.RUnlock()
	
	order, exists := os.orders[id]
	if !exists {
		return Order{}, fmt.Errorf("order not found")
	}
	
	return order, nil
}

func (os *OrderService) handleOrderEvent(topic string, message []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil {
		return err
	}
	
	fmt.Printf("📨 OrderService received event: %s\n", event["event_type"])
	return nil
}

type PaymentService struct {
	*Microservice
	payments map[string]Payment
	mu       sync.RWMutex
}

func NewPaymentService(registry ServiceRegistry, config ConfigManager, mq MessageQueue) *PaymentService {
	ms := NewMicroservice("payment-service", "1.0.0", "localhost", 8083, registry, config, mq)
	
	ps := &PaymentService{
		Microservice: ms,
		payments:     make(map[string]Payment),
	}
	
	// Subscribe to payment events
	mq.Subscribe("payment.events", ps.handlePaymentEvent)
	
	return ps
}

func (ps *PaymentService) ProcessPayment(payment Payment) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	payment.ID = fmt.Sprintf("payment-%d", time.Now().UnixNano())
	payment.CreatedAt = time.Now()
	payment.Status = "processing"
	ps.payments[payment.ID] = payment
	
	// Simulate payment processing
	time.Sleep(100 * time.Millisecond)
	
	// Randomly succeed or fail
	if rand.Float64() < 0.9 {
		payment.Status = "completed"
	} else {
		payment.Status = "failed"
	}
	
	ps.payments[payment.ID] = payment
	
	// Publish payment event
	event := map[string]interface{}{
		"event_type":  "payment.processed",
		"payment_id":  payment.ID,
		"order_id":    payment.OrderID,
		"amount":      payment.Amount,
		"status":      payment.Status,
		"timestamp":   time.Now(),
	}
	
	ps.MessageQueue.Publish("payment.events", event)
	
	fmt.Printf("💳 Payment processed: %s for order %s (Amount: $%.2f, Status: %s)\n", 
		payment.ID, payment.OrderID, payment.Amount, payment.Status)
	return nil
}

func (ps *PaymentService) handlePaymentEvent(topic string, message []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil {
		return err
	}
	
	fmt.Printf("📨 PaymentService received event: %s\n", event["event_type"])
	return nil
}

// ============================================================================
// DATA MODELS
// ============================================================================

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Payment struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// ============================================================================
// API GATEWAY
// ============================================================================

type APIGateway struct {
	registry     ServiceRegistry
	loadBalancer LoadBalancer
	circuitBreakers map[string]CircuitBreaker
	mu           sync.RWMutex
}

func NewAPIGateway(registry ServiceRegistry, loadBalancer LoadBalancer) *APIGateway {
	return &APIGateway{
		registry:        registry,
		loadBalancer:    loadBalancer,
		circuitBreakers: make(map[string]CircuitBreaker),
	}
}

func (gw *APIGateway) RouteRequest(serviceName string, path string, method string) (interface{}, error) {
	// Discover services
	services, err := gw.registry.Discover(serviceName)
	if err != nil {
		return nil, fmt.Errorf("service discovery failed: %w", err)
	}
	
	if len(services) == 0 {
		return nil, fmt.Errorf("no healthy services found for %s", serviceName)
	}
	
	// Select service using load balancer
	service := gw.loadBalancer.SelectService(services)
	
	// Get or create circuit breaker
	gw.mu.Lock()
	cb, exists := gw.circuitBreakers[serviceName]
	if !exists {
		cb = NewCircuitBreaker(serviceName, 5, 30*time.Second)
		gw.circuitBreakers[serviceName] = cb
	}
	gw.mu.Unlock()
	
	// Call service through circuit breaker
	result, err := cb.Call(func() (interface{}, error) {
		// Simulate API call
		url := fmt.Sprintf("http://%s:%d%s", service.Host, service.Port, path)
		fmt.Printf("🌐 Routing %s %s to %s\n", method, path, url)
		
		// Simulate response
		time.Sleep(50 * time.Millisecond)
		
		return map[string]interface{}{
			"service": service.Name,
			"host":    service.Host,
			"port":    service.Port,
			"path":    path,
			"method":  method,
			"status":  "success",
		}, nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("service call failed: %w", err)
	}
	
	return result, nil
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🏗️ MICROSERVICES ARCHITECTURE DEMONSTRATION")
	fmt.Println("=============================================")
	fmt.Println()
	
	// Initialize infrastructure components
	registry := NewServiceRegistry()
	config := NewConfigManager()
	messageQueue := NewMessageQueue()
	loadBalancer := NewRoundRobinLoadBalancer()
	
	// Create API Gateway
	gateway := NewAPIGateway(registry, loadBalancer)
	
	// Create microservices
	userService := NewUserService(registry, config, messageQueue)
	orderService := NewOrderService(registry, config, messageQueue)
	paymentService := NewPaymentService(registry, config, messageQueue)
	
	// Start services
	fmt.Println("🚀 Starting microservices...")
	userService.Start()
	orderService.Start()
	paymentService.Start()
	
	// Wait for services to register
	time.Sleep(1 * time.Second)
	
	// Demonstrate service discovery
	fmt.Println("\n🔍 Service Discovery:")
	services, _ := registry.Discover("user-service")
	fmt.Printf("Found %d user services\n", len(services))
	
	services, _ = registry.Discover("order-service")
	fmt.Printf("Found %d order services\n", len(services))
	
	services, _ = registry.Discover("payment-service")
	fmt.Printf("Found %d payment services\n", len(services))
	
	// Demonstrate API Gateway routing
	fmt.Println("\n🌐 API Gateway Routing:")
	
	// Route user requests
	userResp, _ := gateway.RouteRequest("user-service", "/users", "GET")
	fmt.Printf("User service response: %v\n", userResp)
	
	orderResp, _ := gateway.RouteRequest("order-service", "/orders", "POST")
	fmt.Printf("Order service response: %v\n", orderResp)
	
	paymentResp, _ := gateway.RouteRequest("payment-service", "/payments", "POST")
	fmt.Printf("Payment service response: %v\n", paymentResp)
	
	// Demonstrate business operations
	fmt.Println("\n💼 Business Operations:")
	
	// Create a user
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	userService.CreateUser(user)
	
	// Create an order
	order := Order{
		UserID: "user-123",
		Items: []OrderItem{
			{ProductID: "prod-1", Quantity: 2, Price: 29.99},
			{ProductID: "prod-2", Quantity: 1, Price: 49.99},
		},
		Total: 109.97,
	}
	orderService.CreateOrder(order)
	
	// Process payment
	payment := Payment{
		OrderID: "order-456",
		Amount:  109.97,
		Method:  "credit_card",
	}
	paymentService.ProcessPayment(payment)
	
	// Demonstrate circuit breaker
	fmt.Println("\n🔴 Circuit Breaker Demonstration:")
	
	// Create a failing service call
	failingCB := NewCircuitBreaker("failing-service", 3, 10*time.Second)
	
	for i := 0; i < 5; i++ {
		_, err := failingCB.Call(func() (interface{}, error) {
			return nil, fmt.Errorf("service unavailable")
		})
		
		if err != nil {
			fmt.Printf("Call %d failed: %v (State: %s)\n", i+1, err, failingCB.GetState())
		}
	}
	
	// Demonstrate configuration management
	fmt.Println("\n⚙️ Configuration Management:")
	config.Set("database.url", "postgres://localhost:5432/mydb")
	config.Set("redis.url", "redis://localhost:6379")
	config.Set("log.level", "info")
	
	// Watch for config changes
	config.Watch("log.level", func(value interface{}) {
		fmt.Printf("📝 Log level changed to: %v\n", value)
	})
	
	// Update config
	config.Set("log.level", "debug")
	
	// Demonstrate message queue
	fmt.Println("\n📨 Message Queue Demonstration:")
	messageQueue.Publish("system.events", map[string]interface{}{
		"event_type": "system.startup",
		"timestamp":  time.Now(),
		"message":    "All services started successfully",
	})
	
	// Wait for messages to be processed
	time.Sleep(2 * time.Second)
	
	// Stop services
	fmt.Println("\n🛑 Stopping services...")
	userService.Stop()
	orderService.Stop()
	paymentService.Stop()
	
	fmt.Println("\n🎉 Microservices Architecture Demonstration Complete!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Service discovery and registration")
	fmt.Println("✅ Load balancing and routing")
	fmt.Println("✅ Circuit breaker patterns")
	fmt.Println("✅ Message queue and event handling")
	fmt.Println("✅ Configuration management")
	fmt.Println("✅ API Gateway implementation")
	fmt.Println("✅ Microservice communication")
}
