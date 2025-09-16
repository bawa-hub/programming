package main

import (
	"context"
	"fmt"
	"time"
)

// Example 1: Basic Microservices Architecture
func basicMicroservicesArchitecture() {
	fmt.Println("\n1. Basic Microservices Architecture")
	fmt.Println("===================================")
	
	// Simulate microservices communication
	userService := &UserService{}
	orderService := &OrderService{}
	paymentService := &PaymentService{}
	
	// Process order through microservices
	ctx := context.Background()
	userID := "user-123"
	productID := "product-456"
	
	fmt.Printf("  Processing order for user %s, product %s\n", userID, productID)
	
	// Step 1: Get user information
	user, err := userService.GetUser(ctx, userID)
	if err != nil {
		fmt.Printf("  Error getting user: %v\n", err)
		return
	}
	fmt.Printf("  User retrieved: %s\n", user.Name)
	
	// Step 2: Create order
	order, err := orderService.CreateOrder(ctx, userID, productID)
	if err != nil {
		fmt.Printf("  Error creating order: %v\n", err)
		return
	}
	fmt.Printf("  Order created: %s\n", order.ID)
	
	// Step 3: Process payment
	payment, err := paymentService.ProcessPayment(ctx, order.ID, user.CreditCard)
	if err != nil {
		fmt.Printf("  Error processing payment: %v\n", err)
		return
	}
	fmt.Printf("  Payment processed: %s\n", payment.ID)
	
	// Step 4: Confirm order
	err = orderService.ConfirmOrder(ctx, order.ID, payment.ID)
	if err != nil {
		fmt.Printf("  Error confirming order: %v\n", err)
		return
	}
	fmt.Println("  Order confirmed successfully")
	
	fmt.Println("  Basic microservices architecture completed")
}

// Example 2: HTTP Client Communication
func httpClientCommunication() {
	fmt.Println("\n2. HTTP Client Communication")
	fmt.Println("============================")
	
	// Create HTTP client
	client := &HTTPClient{
		baseURL: "http://localhost:8080",
	}
	
	ctx := context.Background()
	
	// Test user operations
	fmt.Println("  Testing user operations...")
	
	// Create user
	user, err := client.CreateUser(ctx, "John Doe", "john@example.com")
	if err != nil {
		fmt.Printf("  Error creating user: %v\n", err)
	} else {
		fmt.Printf("  User created: %s\n", user.Name)
	}
	
	// Get user
	if user != nil {
		retrievedUser, err := client.GetUser(ctx, user.ID)
		if err != nil {
			fmt.Printf("  Error getting user: %v\n", err)
		} else {
			fmt.Printf("  User retrieved: %s\n", retrievedUser.Name)
		}
	}
	
	fmt.Println("  HTTP client communication completed")
}

// Example 3: Service Discovery
func serviceDiscovery() {
	fmt.Println("\n3. Service Discovery")
	fmt.Println("===================")
	
	// Create service registry
	registry := NewServiceRegistry()
	
	// Register services
	userService := ServiceInstance{
		ID:      "user-service-1",
		Address: "localhost",
		Port:    8081,
		Health:  Healthy,
		Metadata: map[string]string{
			"version": "1.0.0",
			"region":  "us-east-1",
		},
	}
	
	orderService := ServiceInstance{
		ID:      "order-service-1",
		Address: "localhost",
		Port:    8082,
		Health:  Healthy,
		Metadata: map[string]string{
			"version": "1.0.0",
			"region":  "us-east-1",
		},
	}
	
	registry.Register("user-service", userService)
	registry.Register("order-service", orderService)
	
	// Create service discovery client
	discoveryClient := NewServiceDiscoveryClient(registry)
	
	// Test service discovery
	fmt.Println("  Testing service discovery...")
	
	userInstance, err := discoveryClient.GetServiceInstance("user-service")
	if err != nil {
		fmt.Printf("  Error getting user service: %v\n", err)
	} else {
		fmt.Printf("  User service found: %s:%d\n", userInstance.Address, userInstance.Port)
	}
	
	orderInstance, err := discoveryClient.GetServiceInstance("order-service")
	if err != nil {
		fmt.Printf("  Error getting order service: %v\n", err)
	} else {
		fmt.Printf("  Order service found: %s:%d\n", orderInstance.Address, orderInstance.Port)
	}
	
	fmt.Println("  Service discovery completed")
}

// Example 4: Load Balancing
func loadBalancing() {
	fmt.Println("\n4. Load Balancing")
	fmt.Println("=================")
	
	// Create load balancer
	loadBalancer := NewAdvancedLoadBalancer()
	
	// Add service instances
	instances := []ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "instance-2", Address: "localhost", Port: 8082, Health: Healthy},
		{ID: "instance-3", Address: "localhost", Port: 8083, Health: Healthy},
	}
	
	for _, instance := range instances {
		loadBalancer.AddInstance(instance)
	}
	
	// Test load balancing
	fmt.Println("  Testing load balancing...")
	
	for i := 0; i < 10; i++ {
		instance, err := loadBalancer.GetInstance()
		if err != nil {
			fmt.Printf("  Error getting instance: %v\n", err)
		} else {
			fmt.Printf("  Request %d routed to: %s\n", i+1, instance.ID)
		}
	}
	
	fmt.Println("  Load balancing completed")
}

// Example 5: Circuit Breaker
func circuitBreaker() {
	fmt.Println("\n5. Circuit Breaker")
	fmt.Println("==================")
	
	// Create circuit breaker
	circuitBreaker := NewCircuitBreaker(3, 5*time.Second)
	
	// Test circuit breaker
	fmt.Println("  Testing circuit breaker...")
	
	// Simulate some successful operations
	for i := 0; i < 2; i++ {
		err := circuitBreaker.Execute(func() error {
			fmt.Printf("    Operation %d: Success\n", i+1)
			return nil
		})
		if err != nil {
			fmt.Printf("    Operation %d failed: %v\n", i+1, err)
		}
	}
	
	// Simulate some failures to trigger circuit breaker
	for i := 0; i < 4; i++ {
		err := circuitBreaker.Execute(func() error {
			fmt.Printf("    Operation %d: Failure\n", i+3)
			return fmt.Errorf("simulated failure")
		})
		if err != nil {
			fmt.Printf("    Operation %d failed: %v\n", i+3, err)
		}
	}
	
	fmt.Println("  Circuit breaker completed")
}

// Example 6: Retry Pattern
func retryPattern() {
	fmt.Println("\n6. Retry Pattern")
	fmt.Println("================")
	
	// Create retry mechanism
	retry := NewRetry(3, 1*time.Second)
	
	// Test retry pattern
	fmt.Println("  Testing retry pattern...")
	
	err := retry.Execute(func() error {
		fmt.Println("    Attempting operation...")
		// Simulate random failure
		if time.Now().UnixNano()%3 == 0 {
			return nil // Success
		}
		return fmt.Errorf("temporary failure")
	})
	
	if err != nil {
		fmt.Printf("  Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("  Operation succeeded")
	}
	
	fmt.Println("  Retry pattern completed")
}

// Example 7: Timeout Pattern
func timeoutPattern() {
	fmt.Println("\n7. Timeout Pattern")
	fmt.Println("==================")
	
	// Test timeout pattern
	fmt.Println("  Testing timeout pattern...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Simulate long-running operation
	done := make(chan error, 1)
	go func() {
		time.Sleep(3 * time.Second) // This will exceed timeout
		done <- nil
	}()
	
	select {
	case <-ctx.Done():
		fmt.Printf("  Operation timed out: %v\n", ctx.Err())
	case err := <-done:
		if err != nil {
			fmt.Printf("  Operation failed: %v\n", err)
		} else {
			fmt.Println("  Operation completed")
		}
	}
	
	fmt.Println("  Timeout pattern completed")
}

// Example 8: Bulkhead Pattern
func bulkheadPattern() {
	fmt.Println("\n8. Bulkhead Pattern")
	fmt.Println("===================")
	
	// Create bulkhead with separate thread pools
	bulkhead := NewBulkhead(2, 2) // 2 threads for each pool
	
	// Test bulkhead pattern
	fmt.Println("  Testing bulkhead pattern...")
	
	// Submit tasks to different pools
	for i := 0; i < 4; i++ {
		poolID := i % 2
		taskID := i + 1
		
		bulkhead.Submit(poolID, func() {
			fmt.Printf("    Task %d executing in pool %d\n", taskID, poolID)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	// Wait for completion
	bulkhead.Wait()
	
	fmt.Println("  Bulkhead pattern completed")
}

// Example 9: Rate Limiting
func rateLimiting() {
	fmt.Println("\n9. Rate Limiting")
	fmt.Println("================")
	
	// Create rate limiter
	rateLimiter := NewRateLimiter(5, time.Second) // 5 requests per second
	
	// Test rate limiting
	fmt.Println("  Testing rate limiting...")
	
	for i := 0; i < 10; i++ {
		allowed := rateLimiter.Allow()
		if allowed {
			fmt.Printf("    Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("    Request %d: Rate limited\n", i+1)
		}
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Println("  Rate limiting completed")
}

// Example 10: Health Checks
func healthChecks() {
	fmt.Println("\n10. Health Checks")
	fmt.Println("=================")
	
	// Create health checker
	healthChecker := NewHealthChecker()
	
	// Add services to check
	healthChecker.AddService("user-service", "http://localhost:8081/health")
	healthChecker.AddService("order-service", "http://localhost:8082/health")
	healthChecker.AddService("payment-service", "http://localhost:8083/health")
	
	// Test health checks
	fmt.Println("  Testing health checks...")
	
	services := healthChecker.CheckAll()
	for name, status := range services {
		fmt.Printf("    %s: %s\n", name, status)
	}
	
	fmt.Println("  Health checks completed")
}

// Example 11: Service Mesh Simulation
func serviceMeshSimulation() {
	fmt.Println("\n11. Service Mesh Simulation")
	fmt.Println("===========================")
	
	// Create service mesh
	mesh := NewBasicServiceMesh()
	
	// Add services to mesh
	mesh.AddService("user-service", "localhost:8081")
	mesh.AddService("order-service", "localhost:8082")
	mesh.AddService("payment-service", "localhost:8083")
	
	// Test service mesh communication
	fmt.Println("  Testing service mesh communication...")
	
	// Simulate service-to-service communication
	mesh.SendRequest("user-service", "order-service", "get-user", map[string]interface{}{
		"user_id": "123",
	})
	
	mesh.SendRequest("order-service", "payment-service", "process-payment", map[string]interface{}{
		"order_id": "456",
		"amount":   100.00,
	})
	
	fmt.Println("  Service mesh simulation completed")
}

// Example 12: Event-Driven Communication
func eventDrivenCommunication() {
	fmt.Println("\n12. Event-Driven Communication")
	fmt.Println("==============================")
	
	// Create event bus
	eventBus := NewEventBus()
	
	// Subscribe to events
	eventBus.Subscribe("user.created", func(event Event) {
		fmt.Printf("  User created event received: %v\n", event.Data)
	})
	
	eventBus.Subscribe("order.created", func(event Event) {
		fmt.Printf("  Order created event received: %v\n", event.Data)
	})
	
	// Publish events
	fmt.Println("  Publishing events...")
	
	eventBus.Publish("user.created", map[string]interface{}{
		"user_id": "123",
		"name":    "John Doe",
		"email":   "john@example.com",
	})
	
	eventBus.Publish("order.created", map[string]interface{}{
		"order_id": "456",
		"user_id":  "123",
		"amount":   100.00,
	})
	
	fmt.Println("  Event-driven communication completed")
}

// Example 13: Message Queue Communication
func messageQueueCommunication() {
	fmt.Println("\n13. Message Queue Communication")
	fmt.Println("===============================")
	
	// Create message queue
	messageQueue := NewMessageQueue()
	
	// Create producer
	producer := NewProducer(messageQueue)
	
	// Create consumer
	consumer := NewConsumer(messageQueue)
	
	// Start consumer
	go consumer.Start()
	
	// Send messages
	fmt.Println("  Sending messages...")
	
	for i := 0; i < 5; i++ {
		message := Message{
			ID:      fmt.Sprintf("msg-%d", i+1),
			Content: fmt.Sprintf("Message %d", i+1),
			Topic:   "test-topic",
		}
		
		producer.Send(message)
		fmt.Printf("    Sent message: %s\n", message.ID)
	}
	
	// Wait for processing
	time.Sleep(500 * time.Millisecond)
	
	fmt.Println("  Message queue communication completed")
}

// Example 14: Distributed Tracing
func distributedTracing() {
	fmt.Println("\n14. Distributed Tracing")
	fmt.Println("=======================")
	
	// Create tracer
	tracer := NewTracer()
	
	// Start trace
	span := tracer.StartSpan("microservices-operation")
	defer span.Finish()
	
	// Add child spans
	userSpan := tracer.StartChildSpan(span, "get-user")
	time.Sleep(100 * time.Millisecond)
	userSpan.Finish()
	
	orderSpan := tracer.StartChildSpan(span, "create-order")
	time.Sleep(150 * time.Millisecond)
	orderSpan.Finish()
	
	paymentSpan := tracer.StartChildSpan(span, "process-payment")
	time.Sleep(200 * time.Millisecond)
	paymentSpan.Finish()
	
	// Log trace
	tracer.LogSpan(span)
	
	fmt.Println("  Distributed tracing completed")
}

// Example 15: Service Monitoring
func serviceMonitoring() {
	fmt.Println("\n15. Service Monitoring")
	fmt.Println("======================")
	
	// Create monitor
	monitor := NewServiceMonitor()
	
	// Add services to monitor
	monitor.AddService("user-service", "localhost:8081")
	monitor.AddService("order-service", "localhost:8082")
	monitor.AddService("payment-service", "localhost:8083")
	
	// Start monitoring
	go monitor.Start()
	
	// Simulate some activity
	time.Sleep(1 * time.Second)
	
	// Get metrics
	metrics := monitor.GetMetrics()
	fmt.Printf("  Services monitored: %d\n", len(metrics))
	
	for service, metric := range metrics {
		fmt.Printf("    %s: %d requests, %v avg response time\n", 
			service, metric.RequestCount, metric.AvgResponseTime)
	}
	
	fmt.Println("  Service monitoring completed")
}

// Example 16: Fault Injection
func faultInjection() {
	fmt.Println("\n16. Fault Injection")
	fmt.Println("===================")
	
	// Create fault injector
	faultInjector := NewFaultInjector()
	
	// Configure faults
	faultInjector.AddFault("user-service", "timeout", 0.3) // 30% timeout
	faultInjector.AddFault("order-service", "error", 0.2)  // 20% error
	
	// Test fault injection
	fmt.Println("  Testing fault injection...")
	
	for i := 0; i < 10; i++ {
		// Test user service
		err := faultInjector.InjectFault("user-service", func() error {
			fmt.Printf("    User service call %d\n", i+1)
			return nil
		})
		if err != nil {
			fmt.Printf("    User service fault: %v\n", err)
		}
		
		// Test order service
		err = faultInjector.InjectFault("order-service", func() error {
			fmt.Printf("    Order service call %d\n", i+1)
			return nil
		})
		if err != nil {
			fmt.Printf("    Order service fault: %v\n", err)
		}
	}
	
	fmt.Println("  Fault injection completed")
}

// Example 17: Service Versioning
func serviceVersioning() {
	fmt.Println("\n17. Service Versioning")
	fmt.Println("======================")
	
	// Create versioned service registry
	registry := NewVersionedServiceRegistry()
	
	// Register different versions
	registry.Register("user-service", "1.0.0", "localhost:8081")
	registry.Register("user-service", "2.0.0", "localhost:8082")
	registry.Register("order-service", "1.0.0", "localhost:8083")
	registry.Register("order-service", "1.1.0", "localhost:8084")
	
	// Test version selection
	fmt.Println("  Testing version selection...")
	
	// Get latest version
	instance, err := registry.GetLatest("user-service")
	if err != nil {
		fmt.Printf("  Error getting latest version: %v\n", err)
	} else {
		fmt.Printf("  Latest user-service: %s\n", instance.Address)
	}
	
	// Get specific version
	instance, err = registry.GetVersion("order-service", "1.0.0")
	if err != nil {
		fmt.Printf("  Error getting specific version: %v\n", err)
	} else {
		fmt.Printf("  Order-service v1.0.0: %s\n", instance.Address)
	}
	
	fmt.Println("  Service versioning completed")
}

// Example 18: Service Mesh Security
func serviceMeshSecurity() {
	fmt.Println("\n18. Service Mesh Security")
	fmt.Println("=========================")
	
	// Create secure service mesh
	secureMesh := NewSecureServiceMesh()
	
	// Add services with security policies
	secureMesh.AddService("user-service", "localhost", 8081, "us-east-1", "1.0.0")
	secureMesh.AddService("order-service", "localhost", 8082, "us-east-1", "1.0.0")
	secureMesh.AddService("payment-service", "localhost", 8083, "us-east-1", "1.0.0")
	
	// Test secure communication
	fmt.Println("  Testing secure communication...")
	
	// Internal service communication
	err := secureMesh.SendSecureRequest("user-service", "order-service", "get-user", map[string]interface{}{
		"user_id": "123",
	})
	if err != nil {
		fmt.Printf("  Internal communication failed: %v\n", err)
	} else {
		fmt.Println("  Internal communication successful")
	}
	
	// External service communication
	err = secureMesh.SendSecureRequest("order-service", "payment-service", "process-payment", map[string]interface{}{
		"order_id": "456",
		"amount":   100.00,
	})
	if err != nil {
		fmt.Printf("  External communication failed: %v\n", err)
	} else {
		fmt.Println("  External communication successful")
	}
	
	fmt.Println("  Service mesh security completed")
}

// Example 19: Service Mesh Observability
func serviceMeshObservability() {
	fmt.Println("\n19. Service Mesh Observability")
	fmt.Println("==============================")
	
	// Create observable service mesh
	observableMesh := NewObservableServiceMesh()
	
	// Add services
	observableMesh.AddService("user-service", "localhost", 8081)
	observableMesh.AddService("order-service", "localhost", 8082)
	observableMesh.AddService("payment-service", "localhost", 8083)
	
	// Start observability (simplified)
	time.Sleep(100 * time.Millisecond)
	
	// Simulate some traffic
	fmt.Println("  Simulating traffic...")
	
	for i := 0; i < 5; i++ {
		observableMesh.SendRequest("user-service", "order-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	// Wait for metrics collection
	time.Sleep(1 * time.Second)
	
	// Get observability data
	metrics := observableMesh.GetMetrics()
	fmt.Printf("  Total requests: %d\n", metrics.TotalRequests)
	fmt.Printf("  Average latency: %v\n", metrics.AvgLatency)
	fmt.Printf("  Error rate: %.2f%%\n", metrics.ErrorRate*100)
	
	fmt.Println("  Service mesh observability completed")
}

// Example 20: Microservices Testing
func microservicesTesting() {
	fmt.Println("\n20. Microservices Testing")
	fmt.Println("=========================")
	
	// Create test suite
	testSuite := NewMicroservicesTestSuite()
	
	// Add test cases
	testSuite.AddTest("user-service", func() error {
		fmt.Println("    Testing user service...")
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	
	testSuite.AddTest("order-service", func() error {
		fmt.Println("    Testing order service...")
		time.Sleep(150 * time.Millisecond)
		return nil
	})
	
	testSuite.AddTest("payment-service", func() error {
		fmt.Println("    Testing payment service...")
		time.Sleep(200 * time.Millisecond)
		return nil
	})
	
	// Run tests
	fmt.Println("  Running microservices tests...")
	
	results := testSuite.RunTests()
	
	for service, result := range results {
		if result.Error != nil {
			fmt.Printf("    %s: FAILED - %v\n", service, result.Error)
		} else {
			fmt.Printf("    %s: PASSED (%v)\n", service, result.Duration)
		}
	}
	
	fmt.Println("  Microservices testing completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🏗️ Microservices Communication Examples")
	fmt.Println("=======================================")
	
	basicMicroservicesArchitecture()
	httpClientCommunication()
	serviceDiscovery()
	loadBalancing()
	circuitBreaker()
	retryPattern()
	timeoutPattern()
	bulkheadPattern()
	rateLimiting()
	healthChecks()
	serviceMeshSimulation()
	eventDrivenCommunication()
	messageQueueCommunication()
	distributedTracing()
	serviceMonitoring()
	faultInjection()
	serviceVersioning()
	serviceMeshSecurity()
	serviceMeshObservability()
	microservicesTesting()
}
