# 🏗️ Level 5, Topic 1: Microservices Communication

## 🚀 Overview
Welcome to **Level 5: Real-World Architectures**! This is where we build production-ready concurrent systems. Microservices Communication is the foundation of modern distributed systems, covering gRPC, HTTP/2, service mesh, circuit breakers, and fault tolerance patterns.

---

## 📚 Table of Contents

1. [Microservices Fundamentals](#microservices-fundamentals)
2. [gRPC Communication](#grpc-communication)
3. [HTTP/2 and REST APIs](#http2-and-rest-apis)
4. [Service Discovery](#service-discovery)
5. [Load Balancing](#load-balancing)
6. [Circuit Breakers](#circuit-breakers)
7. [Service Mesh](#service-mesh)
8. [Fault Tolerance](#fault-tolerance)
9. [Message Queues](#message-queues)
10. [Event-Driven Architecture](#event-driven-architecture)
11. [Monitoring and Observability](#monitoring-and-observability)
12. [Security and Authentication](#security-and-authentication)
13. [Performance Optimization](#performance-optimization)
14. [Testing Strategies](#testing-strategies)
15. [Deployment Patterns](#deployment-patterns)
16. [Best Practices](#best-practices)

---

## 🏗️ Microservices Fundamentals

### What are Microservices?

Microservices are an architectural approach where applications are built as a collection of loosely coupled, independently deployable services. Each service is responsible for a specific business capability and communicates with other services through well-defined APIs.

### Key Characteristics

#### 1. Service Independence
- Each service can be developed, deployed, and scaled independently
- Technology stack can vary per service
- Database per service pattern

#### 2. Decentralized Governance
- No single point of failure
- Distributed decision making
- Team autonomy

#### 3. Fault Isolation
- Failure in one service doesn't bring down the entire system
- Graceful degradation
- Circuit breaker patterns

#### 4. Data Consistency
- Eventual consistency
- Saga patterns for distributed transactions
- CQRS (Command Query Responsibility Segregation)

### Microservices vs Monoliths

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Example 1: Monolithic Architecture
type MonolithicService struct {
	userService    *UserService
	orderService   *OrderService
	paymentService *PaymentService
}

func (m *MonolithicService) ProcessOrder(ctx context.Context, userID, productID string) error {
	// All services in one process
	user, err := m.userService.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	
	order, err := m.orderService.CreateOrder(ctx, userID, productID)
	if err != nil {
		return err
	}
	
	payment, err := m.paymentService.ProcessPayment(ctx, order.ID, user.CreditCard)
	if err != nil {
		return err
	}
	
	return m.orderService.ConfirmOrder(ctx, order.ID, payment.ID)
}

// Example 2: Microservices Architecture
type MicroservicesService struct {
	userClient    *UserServiceClient
	orderClient   *OrderServiceClient
	paymentClient *PaymentServiceClient
}

func (m *MicroservicesService) ProcessOrder(ctx context.Context, userID, productID string) error {
	// Services communicate over network
	user, err := m.userClient.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	
	order, err := m.orderClient.CreateOrder(ctx, userID, productID)
	if err != nil {
		return err
	}
	
	payment, err := m.paymentClient.ProcessPayment(ctx, order.ID, user.CreditCard)
	if err != nil {
		return err
	}
	
	return m.orderClient.ConfirmOrder(ctx, order.ID, payment.ID)
}

// Service definitions
type UserService struct{}
type OrderService struct{}
type PaymentService struct{}

type UserServiceClient struct{}
type OrderServiceClient struct{}
type PaymentServiceClient struct{}

func (u *UserService) GetUser(ctx context.Context, userID string) (*User, error) {
	return &User{ID: userID, CreditCard: "1234-5678-9012-3456"}, nil
}

func (o *OrderService) CreateOrder(ctx context.Context, userID, productID string) (*Order, error) {
	return &Order{ID: "order-123", UserID: userID, ProductID: productID}, nil
}

func (p *PaymentService) ProcessPayment(ctx context.Context, orderID, creditCard string) (*Payment, error) {
	return &Payment{ID: "payment-123", OrderID: orderID}, nil
}

func (o *OrderService) ConfirmOrder(ctx context.Context, orderID, paymentID string) error {
	return nil
}

// Data structures
type User struct {
	ID         string
	CreditCard string
}

type Order struct {
	ID        string
	UserID    string
	ProductID string
}

type Payment struct {
	ID      string
	OrderID string
}
```

---

## 🔌 gRPC Communication

### Understanding gRPC

gRPC is a high-performance, open-source RPC framework developed by Google. It uses HTTP/2 for transport, Protocol Buffers for serialization, and provides features like load balancing, health checking, and authentication.

### gRPC Benefits

#### 1. Performance
- HTTP/2 multiplexing
- Binary serialization with Protocol Buffers
- Streaming support
- Connection pooling

#### 2. Type Safety
- Strong typing with Protocol Buffers
- Code generation
- Interface contracts

#### 3. Cross-Language Support
- Works across multiple programming languages
- Consistent API across services

### gRPC Implementation

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Example 1: Basic gRPC Server
type UserServiceServer struct {
	users map[string]*User
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	user, exists := s.users[req.UserId]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	user := &User{
		Id:    req.Name + "-" + fmt.Sprintf("%d", time.Now().Unix()),
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[user.Id] = user
	return user, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	var users []*User
	for _, user := range s.users {
		users = append(users, user)
	}
	return &ListUsersResponse{Users: users}, nil
}

// Example 2: gRPC Client
type UserServiceClient struct {
	conn   *grpc.ClientConn
	client UserServiceClient
}

func NewUserServiceClient(address string) (*UserServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	
	client := NewUserServiceClient(conn)
	return &UserServiceClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *UserServiceClient) GetUser(ctx context.Context, userID string) (*User, error) {
	req := &GetUserRequest{UserId: userID}
	return c.client.GetUser(ctx, req)
}

func (c *UserServiceClient) CreateUser(ctx context.Context, name, email string) (*User, error) {
	req := &CreateUserRequest{
		Name:  name,
		Email: email,
	}
	return c.client.CreateUser(ctx, req)
}

func (c *UserServiceClient) Close() error {
	return c.conn.Close()
}

// Example 3: gRPC Streaming
func (s *UserServiceServer) StreamUsers(req *ListUsersRequest, stream UserService_StreamUsersServer) error {
	for _, user := range s.users {
		if err := stream.Send(user); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond) // Simulate processing
	}
	return nil
}

// Example 4: gRPC with Middleware
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	
	log.Printf("gRPC call: %s", info.FullMethod)
	
	resp, err := handler(ctx, req)
	
	log.Printf("gRPC call completed: %s, duration: %v", info.FullMethod, time.Since(start))
	
	return resp, err
}

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	
	userService := &UserServiceServer{
		users: make(map[string]*User),
	}
	
	RegisterUserServiceServer(s, userService)
	reflection.Register(s)
	
	log.Println("gRPC server starting on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

---

## 🌐 HTTP/2 and REST APIs

### HTTP/2 Benefits

#### 1. Multiplexing
- Multiple requests over single connection
- Reduced latency
- Better resource utilization

#### 2. Server Push
- Proactive resource delivery
- Reduced round trips
- Improved performance

#### 3. Header Compression
- HPACK compression
- Reduced bandwidth usage
- Better performance

### REST API Implementation

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Example 1: REST API Server
type RESTAPIServer struct {
	userService    *UserService
	orderService   *OrderService
	paymentService *PaymentService
}

func NewRESTAPIServer() *RESTAPIServer {
	return &RESTAPIServer{
		userService:    &UserService{},
		orderService:   &OrderService{},
		paymentService: &PaymentService{},
	}
}

func (s *RESTAPIServer) Start() {
	r := mux.NewRouter()
	
	// User endpoints
	r.HandleFunc("/users", s.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", s.GetUser).Methods("GET")
	r.HandleFunc("/users", s.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
	
	// Order endpoints
	r.HandleFunc("/orders", s.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", s.GetOrder).Methods("GET")
	r.HandleFunc("/orders", s.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}/confirm", s.ConfirmOrder).Methods("POST")
	
	// Payment endpoints
	r.HandleFunc("/payments", s.GetPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", s.GetPayment).Methods("GET")
	r.HandleFunc("/payments", s.ProcessPayment).Methods("POST")
	
	// Middleware
	r.Use(s.loggingMiddleware)
	r.Use(s.corsMiddleware)
	r.Use(s.rateLimitMiddleware)
	
	log.Println("REST API server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Example 2: HTTP Handlers
func (s *RESTAPIServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	users, err := s.userService.GetUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s *RESTAPIServer) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	userID := vars["id"]
	
	user, err := s.userService.GetUser(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *RESTAPIServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	user, err := s.userService.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Example 3: Middleware
func (s *RESTAPIServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		log.Printf("HTTP %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		log.Printf("HTTP %s %s completed in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func (s *RESTAPIServer) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func (s *RESTAPIServer) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simple rate limiting implementation
		// In production, use a proper rate limiter
		time.Sleep(10 * time.Millisecond)
		next.ServeHTTP(w, r)
	})
}

// Example 4: HTTP Client
type HTTPClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *HTTPClient) GetUser(ctx context.Context, userID string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, userID)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}
	
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (c *HTTPClient) CreateUser(ctx context.Context, name, email string) (*User, error) {
	url := fmt.Sprintf("%s/users", c.baseURL)
	
	reqBody := CreateUserRequest{
		Name:  name,
		Email: email,
	}
	
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}
	
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	
	return &user, nil
}
```

---

## 🔍 Service Discovery

### Understanding Service Discovery

Service discovery is the process of automatically detecting and registering services in a distributed system. It enables services to find and communicate with each other without hardcoded endpoints.

### Service Discovery Patterns

#### 1. Client-Side Discovery
- Client queries service registry
- Client selects service instance
- Client makes request directly

#### 2. Server-Side Discovery
- Client makes request to load balancer
- Load balancer queries service registry
- Load balancer routes request to service

### Service Discovery Implementation

```go
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Example 1: Service Registry
type ServiceRegistry struct {
	services map[string][]ServiceInstance
	mutex    sync.RWMutex
}

type ServiceInstance struct {
	ID       string
	Address  string
	Port     int
	Health   HealthStatus
	Metadata map[string]string
}

type HealthStatus int

const (
	Healthy HealthStatus = iota
	Unhealthy
	Unknown
)

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string][]ServiceInstance),
	}
}

func (sr *ServiceRegistry) Register(serviceName string, instance ServiceInstance) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	sr.services[serviceName] = append(sr.services[serviceName], instance)
	log.Printf("Registered service: %s at %s:%d", serviceName, instance.Address, instance.Port)
}

func (sr *ServiceRegistry) Deregister(serviceName, instanceID string) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	instances := sr.services[serviceName]
	for i, instance := range instances {
		if instance.ID == instanceID {
			sr.services[serviceName] = append(instances[:i], instances[i+1:]...)
			log.Printf("Deregistered service: %s instance %s", serviceName, instanceID)
			break
		}
	}
}

func (sr *ServiceRegistry) GetInstances(serviceName string) []ServiceInstance {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	instances := sr.services[serviceName]
	var healthyInstances []ServiceInstance
	
	for _, instance := range instances {
		if instance.Health == Healthy {
			healthyInstances = append(healthyInstances, instance)
		}
	}
	
	return healthyInstances
}

// Example 2: Service Discovery Client
type ServiceDiscoveryClient struct {
	registry *ServiceRegistry
	balancer LoadBalancer
}

func NewServiceDiscoveryClient(registry *ServiceRegistry) *ServiceDiscoveryClient {
	return &ServiceDiscoveryClient{
		registry: registry,
		balancer: NewRoundRobinBalancer(),
	}
}

func (sdc *ServiceDiscoveryClient) GetServiceInstance(serviceName string) (*ServiceInstance, error) {
	instances := sdc.registry.GetInstances(serviceName)
	if len(instances) == 0 {
		return nil, fmt.Errorf("no healthy instances found for service: %s", serviceName)
	}
	
	instance := sdc.balancer.SelectInstance(instances)
	return &instance, nil
}

// Example 3: Load Balancer
type LoadBalancer interface {
	SelectInstance(instances []ServiceInstance) ServiceInstance
}

type RoundRobinBalancer struct {
	current int
	mutex   sync.Mutex
}

func NewRoundRobinBalancer() *RoundRobinBalancer {
	return &RoundRobinBalancer{}
}

func (rr *RoundRobinBalancer) SelectInstance(instances []ServiceInstance) ServiceInstance {
	rr.mutex.Lock()
	defer rr.mutex.Unlock()
	
	if len(instances) == 0 {
		return ServiceInstance{}
	}
	
	instance := instances[rr.current%len(instances)]
	rr.current++
	return instance
}

type RandomBalancer struct{}

func NewRandomBalancer() *RandomBalancer {
	return &RandomBalancer{}
}

func (rb *RandomBalancer) SelectInstance(instances []ServiceInstance) ServiceInstance {
	if len(instances) == 0 {
		return ServiceInstance{}
	}
	
	// Simple random selection (in production, use crypto/rand)
	index := time.Now().UnixNano() % int64(len(instances))
	return instances[index]
}

// Example 4: Health Checker
type HealthChecker struct {
	registry *ServiceRegistry
	interval time.Duration
}

func NewHealthChecker(registry *ServiceRegistry, interval time.Duration) *HealthChecker {
	return &HealthChecker{
		registry: registry,
		interval: interval,
	}
}

func (hc *HealthChecker) Start() {
	ticker := time.NewTicker(hc.interval)
	go func() {
		for range ticker.C {
			hc.checkAllServices()
		}
	}()
}

func (hc *HealthChecker) checkAllServices() {
	hc.registry.mutex.RLock()
	services := make(map[string][]ServiceInstance)
	for name, instances := range hc.registry.services {
		services[name] = instances
	}
	hc.registry.mutex.RUnlock()
	
	for serviceName, instances := range services {
		for i, instance := range instances {
			healthy := hc.checkHealth(instance)
			if !healthy {
				hc.updateHealthStatus(serviceName, instance.ID, Unhealthy)
			} else {
				hc.updateHealthStatus(serviceName, instance.ID, Healthy)
			}
		}
	}
}

func (hc *HealthChecker) checkHealth(instance ServiceInstance) bool {
	// Simple health check - in production, make actual HTTP request
	// to health endpoint
	return time.Now().UnixNano()%2 == 0 // Simulate random health
}

func (hc *HealthChecker) updateHealthStatus(serviceName, instanceID string, health HealthStatus) {
	hc.registry.mutex.Lock()
	defer hc.registry.mutex.Unlock()
	
	instances := hc.registry.services[serviceName]
	for i, instance := range instances {
		if instance.ID == instanceID {
			instances[i].Health = health
			break
		}
	}
}
```

---

## ⚖️ Load Balancing

### Load Balancing Strategies

#### 1. Round Robin
- Distributes requests evenly
- Simple and predictable
- Good for stateless services

#### 2. Weighted Round Robin
- Assigns weights to instances
- Handles different capacity instances
- More sophisticated distribution

#### 3. Least Connections
- Routes to instance with fewest active connections
- Good for long-lived connections
- Considers current load

#### 4. Random
- Randomly selects instance
- Simple implementation
- Good for uniform distribution

### Load Balancing Implementation

```go
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Example 1: Advanced Load Balancer
type AdvancedLoadBalancer struct {
	instances    []ServiceInstance
	balancer     LoadBalancer
	healthChecker *HealthChecker
	mutex        sync.RWMutex
}

func NewAdvancedLoadBalancer() *AdvancedLoadBalancer {
	return &AdvancedLoadBalancer{
		balancer: NewRoundRobinBalancer(),
	}
}

func (alb *AdvancedLoadBalancer) AddInstance(instance ServiceInstance) {
	alb.mutex.Lock()
	defer alb.mutex.Unlock()
	
	alb.instances = append(alb.instances, instance)
	log.Printf("Added instance: %s", instance.ID)
}

func (alb *AdvancedLoadBalancer) RemoveInstance(instanceID string) {
	alb.mutex.Lock()
	defer alb.mutex.Unlock()
	
	for i, instance := range alb.instances {
		if instance.ID == instanceID {
			alb.instances = append(alb.instances[:i], alb.instances[i+1:]...)
			log.Printf("Removed instance: %s", instanceID)
			break
		}
	}
}

func (alb *AdvancedLoadBalancer) GetInstance() (*ServiceInstance, error) {
	alb.mutex.RLock()
	defer alb.mutex.RUnlock()
	
	healthyInstances := alb.getHealthyInstances()
	if len(healthyInstances) == 0 {
		return nil, fmt.Errorf("no healthy instances available")
	}
	
	instance := alb.balancer.SelectInstance(healthyInstances)
	return &instance, nil
}

func (alb *AdvancedLoadBalancer) getHealthyInstances() []ServiceInstance {
	var healthy []ServiceInstance
	for _, instance := range alb.instances {
		if instance.Health == Healthy {
			healthy = append(healthy, instance)
		}
	}
	return healthy
}

// Example 2: Weighted Round Robin Load Balancer
type WeightedRoundRobinBalancer struct {
	current int
	weights []int
	mutex   sync.Mutex
}

func NewWeightedRoundRobinBalancer(weights []int) *WeightedRoundRobinBalancer {
	return &WeightedRoundRobinBalancer{
		weights: weights,
	}
}

func (wrr *WeightedRoundRobinBalancer) SelectInstance(instances []ServiceInstance) ServiceInstance {
	wrr.mutex.Lock()
	defer wrr.mutex.Unlock()
	
	if len(instances) == 0 {
		return ServiceInstance{}
	}
	
	// Use weights to determine selection
	totalWeight := 0
	for _, weight := range wrr.weights {
		totalWeight += weight
	}
	
	if totalWeight == 0 {
		// Fallback to round robin
		instance := instances[wrr.current%len(instances)]
		wrr.current++
		return instance
	}
	
	// Weighted selection
	random := time.Now().UnixNano() % int64(totalWeight)
	currentWeight := 0
	
	for i, instance := range instances {
		if i < len(wrr.weights) {
			currentWeight += wrr.weights[i]
		} else {
			currentWeight += 1 // Default weight
		}
		
		if int64(currentWeight) > random {
			return instance
		}
	}
	
	// Fallback
	return instances[0]
}

// Example 3: Least Connections Load Balancer
type LeastConnectionsBalancer struct {
	connections map[string]int
	mutex       sync.Mutex
}

func NewLeastConnectionsBalancer() *LeastConnectionsBalancer {
	return &LeastConnectionsBalancer{
		connections: make(map[string]int),
	}
}

func (lc *LeastConnectionsBalancer) SelectInstance(instances []ServiceInstance) ServiceInstance {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	
	if len(instances) == 0 {
		return ServiceInstance{}
	}
	
	// Find instance with least connections
	minConnections := int(^uint(0) >> 1) // Max int
	var selected ServiceInstance
	
	for _, instance := range instances {
		connections := lc.connections[instance.ID]
		if connections < minConnections {
			minConnections = connections
			selected = instance
		}
	}
	
	// Increment connection count
	lc.connections[selected.ID]++
	
	return selected
}

func (lc *LeastConnectionsBalancer) ReleaseConnection(instanceID string) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	
	if connections := lc.connections[instanceID]; connections > 0 {
		lc.connections[instanceID]--
	}
}

// Example 4: Load Balancer with Circuit Breaker
type LoadBalancerWithCircuitBreaker struct {
	balancer      LoadBalancer
	circuitBreaker *CircuitBreaker
}

func NewLoadBalancerWithCircuitBreaker() *LoadBalancerWithCircuitBreaker {
	return &LoadBalancerWithCircuitBreaker{
		balancer:      NewRoundRobinBalancer(),
		circuitBreaker: NewCircuitBreaker(5, 30*time.Second),
	}
}

func (lb *LoadBalancerWithCircuitBreaker) SelectInstance(instances []ServiceInstance) ServiceInstance {
	// Use circuit breaker to protect against cascading failures
	return lb.circuitBreaker.Execute(func() ServiceInstance {
		return lb.balancer.SelectInstance(instances)
	})
}
```

---

## 🔌 Circuit Breakers

### Understanding Circuit Breakers

Circuit breakers are a design pattern used to prevent cascading failures in distributed systems. They monitor the success/failure rate of operations and "open" the circuit when failures exceed a threshold.

### Circuit Breaker States

#### 1. Closed State
- Normal operation
- Requests pass through
- Monitors failure rate

#### 2. Open State
- Circuit is open
- Requests fail fast
- No calls to downstream service

#### 3. Half-Open State
- Testing if service recovered
- Limited requests allowed
- Transitions based on results

### Circuit Breaker Implementation

```go
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Example 1: Basic Circuit Breaker
type CircuitBreaker struct {
	state         CircuitState
	failureCount  int
	successCount  int
	threshold     int
	timeout       time.Duration
	lastFailure   time.Time
	mutex         sync.RWMutex
}

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:    StateClosed,
		threshold: threshold,
		timeout:  timeout,
	}
}

func (cb *CircuitBreaker) Execute(operation func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	// Check if circuit should be opened
	if cb.state == StateClosed && cb.failureCount >= cb.threshold {
		cb.state = StateOpen
		cb.lastFailure = time.Now()
		log.Printf("Circuit breaker opened due to %d failures", cb.failureCount)
	}
	
	// Check if circuit should be half-opened
	if cb.state == StateOpen && time.Since(cb.lastFailure) > cb.timeout {
		cb.state = StateHalfOpen
		log.Println("Circuit breaker half-opened for testing")
	}
	
	// Execute operation based on state
	switch cb.state {
	case StateClosed, StateHalfOpen:
		err := operation()
		if err != nil {
			cb.failureCount++
			cb.lastFailure = time.Now()
			if cb.state == StateHalfOpen {
				cb.state = StateOpen
			}
			return err
		}
		
		cb.successCount++
		if cb.state == StateHalfOpen {
			cb.state = StateClosed
			cb.failureCount = 0
		}
		return nil
		
	case StateOpen:
		return fmt.Errorf("circuit breaker is open")
		
	default:
		return fmt.Errorf("unknown circuit breaker state")
	}
}

// Example 2: Advanced Circuit Breaker with Metrics
type AdvancedCircuitBreaker struct {
	state         CircuitState
	failureCount  int
	successCount  int
	threshold     int
	timeout       time.Duration
	lastFailure   time.Time
	metrics       *CircuitBreakerMetrics
	mutex         sync.RWMutex
}

type CircuitBreakerMetrics struct {
	TotalRequests    int64
	SuccessfulRequests int64
	FailedRequests   int64
	OpenDuration     time.Duration
	LastOpened       time.Time
}

func NewAdvancedCircuitBreaker(threshold int, timeout time.Duration) *AdvancedCircuitBreaker {
	return &AdvancedCircuitBreaker{
		state:    StateClosed,
		threshold: threshold,
		timeout:  timeout,
		metrics:  &CircuitBreakerMetrics{},
	}
}

func (acb *AdvancedCircuitBreaker) Execute(operation func() error) error {
	acb.mutex.Lock()
	defer acb.mutex.Unlock()
	
	acb.metrics.TotalRequests++
	
	// Check if circuit should be opened
	if acb.state == StateClosed && acb.failureCount >= acb.threshold {
		acb.state = StateOpen
		acb.lastFailure = time.Now()
		acb.metrics.LastOpened = time.Now()
		log.Printf("Circuit breaker opened due to %d failures", acb.failureCount)
	}
	
	// Check if circuit should be half-opened
	if acb.state == StateOpen && time.Since(acb.lastFailure) > acb.timeout {
		acb.state = StateHalfOpen
		acb.metrics.OpenDuration = time.Since(acb.metrics.LastOpened)
		log.Println("Circuit breaker half-opened for testing")
	}
	
	// Execute operation based on state
	switch acb.state {
	case StateClosed, StateHalfOpen:
		err := operation()
		if err != nil {
			acb.failureCount++
			acb.lastFailure = time.Now()
			acb.metrics.FailedRequests++
			if acb.state == StateHalfOpen {
				acb.state = StateOpen
			}
			return err
		}
		
		acb.successCount++
		acb.metrics.SuccessfulRequests++
		if acb.state == StateHalfOpen {
			acb.state = StateClosed
			acb.failureCount = 0
		}
		return nil
		
	case StateOpen:
		acb.metrics.FailedRequests++
		return fmt.Errorf("circuit breaker is open")
		
	default:
		return fmt.Errorf("unknown circuit breaker state")
	}
}

func (acb *AdvancedCircuitBreaker) GetMetrics() *CircuitBreakerMetrics {
	acb.mutex.RLock()
	defer acb.mutex.RUnlock()
	
	return &CircuitBreakerMetrics{
		TotalRequests:      acb.metrics.TotalRequests,
		SuccessfulRequests: acb.metrics.SuccessfulRequests,
		FailedRequests:     acb.metrics.FailedRequests,
		OpenDuration:       acb.metrics.OpenDuration,
		LastOpened:         acb.metrics.LastOpened,
	}
}

// Example 3: Circuit Breaker with Retry
type CircuitBreakerWithRetry struct {
	circuitBreaker *CircuitBreaker
	maxRetries     int
	retryDelay     time.Duration
}

func NewCircuitBreakerWithRetry(threshold int, timeout time.Duration, maxRetries int, retryDelay time.Duration) *CircuitBreakerWithRetry {
	return &CircuitBreakerWithRetry{
		circuitBreaker: NewCircuitBreaker(threshold, timeout),
		maxRetries:     maxRetries,
		retryDelay:     retryDelay,
	}
}

func (cbr *CircuitBreakerWithRetry) Execute(operation func() error) error {
	var lastErr error
	
	for i := 0; i <= cbr.maxRetries; i++ {
		err := cbr.circuitBreaker.Execute(operation)
		if err == nil {
			return nil
		}
		
		lastErr = err
		
		if i < cbr.maxRetries {
			time.Sleep(cbr.retryDelay)
		}
	}
	
	return fmt.Errorf("operation failed after %d retries: %v", cbr.maxRetries, lastErr)
}

// Example 4: Circuit Breaker with Fallback
type CircuitBreakerWithFallback struct {
	circuitBreaker *CircuitBreaker
	fallback       func() error
}

func NewCircuitBreakerWithFallback(threshold int, timeout time.Duration, fallback func() error) *CircuitBreakerWithFallback {
	return &CircuitBreakerWithFallback{
		circuitBreaker: NewCircuitBreaker(threshold, timeout),
		fallback:       fallback,
	}
}

func (cbf *CircuitBreakerWithFallback) Execute(operation func() error) error {
	err := cbf.circuitBreaker.Execute(operation)
	if err != nil {
		// Try fallback if circuit breaker is open
		if cbf.circuitBreaker.state == StateOpen {
			log.Println("Using fallback due to circuit breaker being open")
			return cbf.fallback()
		}
		return err
	}
	return nil
}
```

---

## 🎯 Summary

Microservices Communication is the foundation of modern distributed systems. Key takeaways:

1. **gRPC** provides high-performance, type-safe communication
2. **HTTP/2** offers multiplexing and better performance than HTTP/1.1
3. **Service Discovery** enables dynamic service location
4. **Load Balancing** distributes traffic across service instances
5. **Circuit Breakers** prevent cascading failures
6. **Fault Tolerance** ensures system resilience

This topic provides the foundation for building production-ready microservices! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different communication patterns
3. **Apply** techniques to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** microservices communication patterns

Ready to become a Microservices Communication expert? Let's dive into the implementation! 💪

