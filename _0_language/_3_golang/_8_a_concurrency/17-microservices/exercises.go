package main

import (
	"context"
	"fmt"
	"time"
)

// Exercise 1: Implement Basic Service Discovery
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Basic Service Discovery")
	fmt.Println("=============================================")
	
	// TODO: Implement basic service discovery
	// 1. Create a service registry
	// 2. Register multiple service instances
	// 3. Implement service discovery client
	// 4. Test service discovery
	
	registry := NewServiceRegistry()
	
	// Register user service instances
	userService1 := ServiceInstance{
		ID:      "user-service-1",
		Address: "localhost",
		Port:    8081,
		Health:  Healthy,
		Metadata: map[string]string{
			"version": "1.0.0",
			"region":  "us-east-1",
		},
	}
	
	userService2 := ServiceInstance{
		ID:      "user-service-2",
		Address: "localhost",
		Port:    8082,
		Health:  Healthy,
		Metadata: map[string]string{
			"version": "1.0.0",
			"region":  "us-west-1",
		},
	}
	
	registry.Register("user-service", userService1)
	registry.Register("user-service", userService2)
	
	// Create service discovery client
	discoveryClient := NewServiceDiscoveryClient(registry)
	
	// Test service discovery
	fmt.Println("  Testing service discovery...")
	
	for i := 0; i < 5; i++ {
		instance, err := discoveryClient.GetServiceInstance("user-service")
		if err != nil {
			fmt.Printf("    Error getting service instance: %v\n", err)
		} else {
			fmt.Printf("    Found instance: %s:%d\n", instance.Address, instance.Port)
		}
	}
	
	fmt.Println("  Exercise 1: Basic service discovery completed")
}

// Exercise 2: Implement Load Balancing
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Load Balancing")
	fmt.Println("====================================")
	
	// TODO: Implement load balancing
	// 1. Create multiple service instances
	// 2. Implement round-robin load balancer
	// 3. Test load distribution
	// 4. Compare with random load balancer
	
	// Create load balancer
	loadBalancer := NewAdvancedLoadBalancer()
	
	// Add service instances
	instances := []ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "instance-2", Address: "localhost", Port: 8082, Health: Healthy},
		{ID: "instance-3", Address: "localhost", Port: 8083, Health: Healthy},
		{ID: "instance-4", Address: "localhost", Port: 8084, Health: Healthy},
	}
	
	for _, instance := range instances {
		loadBalancer.AddInstance(instance)
	}
	
	// Test load balancing
	fmt.Println("  Testing load balancing...")
	
	distribution := make(map[string]int)
	for i := 0; i < 20; i++ {
		instance, err := loadBalancer.GetInstance()
		if err != nil {
			fmt.Printf("    Error getting instance: %v\n", err)
		} else {
			distribution[instance.ID]++
		}
	}
	
	fmt.Println("  Load distribution:")
	for instance, count := range distribution {
		fmt.Printf("    %s: %d requests\n", instance, count)
	}
	
	fmt.Println("  Exercise 2: Load balancing completed")
}

// Exercise 3: Implement Circuit Breaker
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Circuit Breaker")
	fmt.Println("=====================================")
	
	// TODO: Implement circuit breaker
	// 1. Create circuit breaker with threshold
	// 2. Simulate failures to trigger circuit breaker
	// 3. Test circuit breaker states
	// 4. Verify recovery behavior
	
	circuitBreaker := NewCircuitBreaker(3, 2*time.Second)
	
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
	
	// Simulate failures to trigger circuit breaker
	for i := 0; i < 5; i++ {
		err := circuitBreaker.Execute(func() error {
			fmt.Printf("    Operation %d: Failure\n", i+3)
			return fmt.Errorf("simulated failure")
		})
		if err != nil {
			fmt.Printf("    Operation %d failed: %v\n", i+3, err)
		}
	}
	
	// Wait for circuit breaker to half-open
	fmt.Println("  Waiting for circuit breaker to half-open...")
	time.Sleep(3 * time.Second)
	
	// Test recovery
	err := circuitBreaker.Execute(func() error {
		fmt.Println("    Recovery operation: Success")
		return nil
	})
	if err != nil {
		fmt.Printf("    Recovery operation failed: %v\n", err)
	}
	
	fmt.Println("  Exercise 3: Circuit breaker completed")
}

// Exercise 4: Implement Retry Pattern
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Retry Pattern")
	fmt.Println("===================================")
	
	// TODO: Implement retry pattern
	// 1. Create retry mechanism with exponential backoff
	// 2. Test retry with different failure scenarios
	// 3. Implement retry with jitter
	// 4. Test retry limits
	
	retry := NewRetry(3, 1*time.Second)
	
	fmt.Println("  Testing retry pattern...")
	
	attempt := 0
	err := retry.Execute(func() error {
		attempt++
		fmt.Printf("    Attempt %d\n", attempt)
		
		// Simulate random failure
		if attempt < 3 {
			return fmt.Errorf("temporary failure")
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("  Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("  Operation succeeded")
	}
	
	fmt.Println("  Exercise 4: Retry pattern completed")
}

// Exercise 5: Implement Timeout Pattern
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Timeout Pattern")
	fmt.Println("=====================================")
	
	// TODO: Implement timeout pattern
	// 1. Create timeout context
	// 2. Test timeout with different durations
	// 3. Implement timeout with fallback
	// 4. Test timeout propagation
	
	fmt.Println("  Testing timeout pattern...")
	
	// Test 1: Operation completes before timeout
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()
	
	done1 := make(chan error, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		done1 <- nil
	}()
	
	select {
	case <-ctx1.Done():
		fmt.Printf("    Test 1 timed out: %v\n", ctx1.Err())
	case err := <-done1:
		if err != nil {
			fmt.Printf("    Test 1 failed: %v\n", err)
		} else {
			fmt.Println("    Test 1 completed successfully")
		}
	}
	
	// Test 2: Operation times out
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()
	
	done2 := make(chan error, 1)
	go func() {
		time.Sleep(1 * time.Second)
		done2 <- nil
	}()
	
	select {
	case <-ctx2.Done():
		fmt.Printf("    Test 2 timed out: %v\n", ctx2.Err())
	case err := <-done2:
		if err != nil {
			fmt.Printf("    Test 2 failed: %v\n", err)
		} else {
			fmt.Println("    Test 2 completed successfully")
		}
	}
	
	fmt.Println("  Exercise 5: Timeout pattern completed")
}

// Exercise 6: Implement Bulkhead Pattern
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Bulkhead Pattern")
	fmt.Println("======================================")
	
	// TODO: Implement bulkhead pattern
	// 1. Create bulkhead with separate thread pools
	// 2. Test isolation between pools
	// 3. Test resource limits
	// 4. Test failure isolation
	
	bulkhead := NewBulkhead(3, 2) // 3 pools, 2 tasks per pool
	
	fmt.Println("  Testing bulkhead pattern...")
	
	// Submit tasks to different pools
	for i := 0; i < 9; i++ {
		poolID := i % 3
		taskID := i + 1
		
		bulkhead.Submit(poolID, func() {
			fmt.Printf("    Task %d executing in pool %d\n", taskID, poolID)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	// Wait for completion
	bulkhead.Wait()
	
	fmt.Println("  Exercise 6: Bulkhead pattern completed")
}

// Exercise 7: Implement Rate Limiting
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Rate Limiting")
	fmt.Println("===================================")
	
	// TODO: Implement rate limiting
	// 1. Create rate limiter with different rates
	// 2. Test rate limiting behavior
	// 3. Implement sliding window rate limiter
	// 4. Test burst capacity
	
	rateLimiter := NewRateLimiter(5, time.Second) // 5 requests per second
	
	fmt.Println("  Testing rate limiting...")
	
	allowed := 0
	blocked := 0
	
	for i := 0; i < 15; i++ {
		if rateLimiter.Allow() {
			allowed++
			fmt.Printf("    Request %d: Allowed\n", i+1)
		} else {
			blocked++
			fmt.Printf("    Request %d: Rate limited\n", i+1)
		}
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Printf("  Allowed: %d, Blocked: %d\n", allowed, blocked)
	
	fmt.Println("  Exercise 7: Rate limiting completed")
}

// Exercise 8: Implement Health Checks
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Health Checks")
	fmt.Println("===================================")
	
	// TODO: Implement health checks
	// 1. Create health checker
	// 2. Add services to check
	// 3. Implement health check endpoints
	// 4. Test health check aggregation
	
	healthChecker := NewHealthChecker()
	
	// Add services to check
	services := []struct {
		name string
		url  string
	}{
		{"user-service", "http://localhost:8081/health"},
		{"order-service", "http://localhost:8082/health"},
		{"payment-service", "http://localhost:8083/health"},
		{"notification-service", "http://localhost:8084/health"},
	}
	
	for _, service := range services {
		healthChecker.AddService(service.name, service.url)
	}
	
	fmt.Println("  Testing health checks...")
	
	results := healthChecker.CheckAll()
	for name, status := range results {
		fmt.Printf("    %s: %s\n", name, status)
	}
	
	fmt.Println("  Exercise 8: Health checks completed")
}

// Exercise 9: Implement Event-Driven Communication
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Event-Driven Communication")
	fmt.Println("================================================")
	
	// TODO: Implement event-driven communication
	// 1. Create event bus
	// 2. Subscribe to events
	// 3. Publish events
	// 4. Test event propagation
	
	eventBus := NewEventBus()
	
	// Subscribe to events
	eventBus.Subscribe("user.created", func(event Event) {
		fmt.Printf("    User created: %v\n", event.Data)
	})
	
	eventBus.Subscribe("order.created", func(event Event) {
		fmt.Printf("    Order created: %v\n", event.Data)
	})
	
	eventBus.Subscribe("payment.processed", func(event Event) {
		fmt.Printf("    Payment processed: %v\n", event.Data)
	})
	
	fmt.Println("  Testing event-driven communication...")
	
	// Publish events
	events := []struct {
		eventType string
		data      map[string]interface{}
	}{
		{
			"user.created",
			map[string]interface{}{
				"user_id": "123",
				"name":    "John Doe",
				"email":   "john@example.com",
			},
		},
		{
			"order.created",
			map[string]interface{}{
				"order_id": "456",
				"user_id":  "123",
				"amount":   100.00,
			},
		},
		{
			"payment.processed",
			map[string]interface{}{
				"payment_id": "789",
				"order_id":   "456",
				"amount":     100.00,
			},
		},
	}
	
	for _, event := range events {
		eventBus.Publish(event.eventType, event.data)
		time.Sleep(50 * time.Millisecond) // Allow time for event processing
	}
	
	fmt.Println("  Exercise 9: Event-driven communication completed")
}

// Exercise 10: Implement Message Queue Communication
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Message Queue Communication")
	fmt.Println("=================================================")
	
	// TODO: Implement message queue communication
	// 1. Create message queue
	// 2. Implement producer and consumer
	// 3. Test message processing
	// 4. Test message ordering
	
	messageQueue := NewMessageQueue()
	
	// Create producer
	producer := NewProducer(messageQueue)
	
	// Create consumer
	consumer := NewConsumer(messageQueue)
	
	// Start consumer
	go consumer.Start()
	
	fmt.Println("  Testing message queue communication...")
	
	// Send messages
	for i := 0; i < 10; i++ {
		message := Message{
			ID:      fmt.Sprintf("msg-%d", i+1),
			Content: fmt.Sprintf("Message %d", i+1),
			Topic:   "test-topic",
		}
		
		producer.Send(message)
		fmt.Printf("    Sent message: %s\n", message.ID)
		time.Sleep(50 * time.Millisecond)
	}
	
	// Wait for processing
	time.Sleep(1 * time.Second)
	
	fmt.Println("  Exercise 10: Message queue communication completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Microservices Communication Exercises")
	fmt.Println("=======================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
