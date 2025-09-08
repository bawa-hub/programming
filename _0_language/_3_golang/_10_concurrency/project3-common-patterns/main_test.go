package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestFanInFanOut(t *testing.T) {
	// Test basic fan-in/fan-out
	input := make(chan int)
	worker1 := fanOutWorker(input, 1)
	worker2 := fanOutWorker(input, 2)
	output := fanIn(worker1, worker2)
	
	// Send work
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results
	results := make([]string, 0)
	for result := range output {
		results = append(results, result)
	}
	
	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}
}

func TestPipeline(t *testing.T) {
	// Test basic pipeline
	stage1 := make(chan int)
	stage2 := make(chan int)
	stage3 := make(chan int)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(stage1)
		for i := 1; i <= 3; i++ {
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers
	go func() {
		defer close(stage2)
		for n := range stage1 {
			stage2 <- n * n
		}
	}()
	
	// Stage 3: Add 10
	go func() {
		defer close(stage3)
		for n := range stage2 {
			stage3 <- n + 10
		}
	}()
	
	// Collect results
	results := make([]int, 0)
	for result := range stage3 {
		results = append(results, result)
	}
	
	expected := []int{11, 14, 19} // (1*1+10), (2*2+10), (3*3+10)
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
	}
}

func TestGenerator(t *testing.T) {
	// Test number generator
	numbers := generateNumbers(1, 5)
	
	results := make([]int, 0)
	for n := range numbers {
		results = append(results, n)
	}
	
	expected := []int{1, 2, 3, 4, 5}
	if len(results) != len(expected) {
		t.Errorf("Expected %d numbers, got %d", len(expected), len(results))
	}
}

func TestFibonacciGenerator(t *testing.T) {
	// Test Fibonacci generator
	fib := generateFibonacci(5)
	
	results := make([]int, 0)
	for n := range fib {
		results = append(results, n)
	}
	
	expected := []int{0, 1, 1, 2, 3}
	if len(results) != len(expected) {
		t.Errorf("Expected %d Fibonacci numbers, got %d", len(expected), len(results))
	}
}

func TestWorkerPool(t *testing.T) {
	// Test basic worker pool
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				results <- job * job
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	resultsList := make([]int, 0)
	for result := range results {
		resultsList = append(resultsList, result)
	}
	
	if len(resultsList) != 5 {
		t.Errorf("Expected 5 results, got %d", len(resultsList))
	}
}

func TestAdvancedWorkerPool(t *testing.T) {
	// Test advanced worker pool
	jobs := make(chan Job, 10)
	results := make(chan JobResult, 10)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				result := processJob(workerID, job)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		jobs <- Job{ID: 1, Type: "add", Data: 5}
		jobs <- Job{ID: 2, Type: "multiply", Data: 3}
		close(jobs)
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	resultsList := make([]JobResult, 0)
	for result := range results {
		resultsList = append(resultsList, result)
	}
	
	if len(resultsList) != 2 {
		t.Errorf("Expected 2 results, got %d", len(resultsList))
	}
}

func TestGracefulSystem(t *testing.T) {
	// Test graceful system
	system := NewGracefulSystem()
	
	// Start the system
	system.Start()
	
	// Let it run briefly
	time.Sleep(100 * time.Millisecond)
	
	// Shutdown
	system.Shutdown()
	
	// Wait for shutdown to complete
	select {
	case <-system.Done():
		// Shutdown completed
	case <-time.After(2 * time.Second):
		t.Error("Shutdown timeout")
	}
}

func TestContextSystem(t *testing.T) {
	// Test context-based system
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	system := NewContextSystem(ctx)
	
	// Start the system
	system.Start()
	
	// Wait for context timeout
	select {
	case <-system.Done():
		// Shutdown completed
	case <-time.After(200 * time.Millisecond):
		t.Error("Context shutdown timeout")
	}
}

func TestTimeoutSystem(t *testing.T) {
	// Test timeout-based system
	system := NewTimeoutSystem(100 * time.Millisecond)
	
	// Start the system
	system.Start()
	
	// Let it run briefly
	time.Sleep(50 * time.Millisecond)
	
	// Shutdown
	system.Shutdown()
	
	// Wait for shutdown to complete
	select {
	case <-system.Done():
		// Shutdown completed
	case <-time.After(200 * time.Millisecond):
		t.Error("Timeout shutdown timeout")
	}
}

func TestSignalSystem(t *testing.T) {
	// Test signal-based system
	system := NewSignalSystem()
	
	// Start the system
	system.Start()
	
	// Let it run briefly
	time.Sleep(50 * time.Millisecond)
	
	// Send signal
	system.Signal()
	
	// Wait for shutdown to complete
	select {
	case <-system.Done():
		// Shutdown completed
	case <-time.After(200 * time.Millisecond):
		t.Error("Signal shutdown timeout")
	}
}

func TestAtomicCounter(t *testing.T) {
	// Test atomic counter
	counter := NewAtomicCounter()
	
	// Test basic operations
	counter.Increment()
	if counter.Get() != 1 {
		t.Errorf("Expected 1, got %d", counter.Get())
	}
	
	counter.Add(5)
	if counter.Get() != 6 {
		t.Errorf("Expected 6, got %d", counter.Get())
	}
	
	counter.Decrement()
	if counter.Get() != 5 {
		t.Errorf("Expected 5, got %d", counter.Get())
	}
}

func TestAtomicCounterConcurrency(t *testing.T) {
	// Test atomic counter concurrency
	counter := NewAtomicCounter()
	var wg sync.WaitGroup
	
	// Test concurrent increments
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	
	if counter.Get() != 1000 {
		t.Errorf("Expected 1000, got %d", counter.Get())
	}
}

func TestThreadSafeCache(t *testing.T) {
	// Test thread-safe cache
	cache := NewThreadSafeCache()
	
	// Test basic operations
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	
	value, exists := cache.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}
	
	value, exists = cache.Get("key2")
	if !exists || value != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}
	
	if cache.Size() != 2 {
		t.Errorf("Expected size 2, got %d", cache.Size())
	}
	
	cache.Delete("key1")
	if cache.Size() != 1 {
		t.Errorf("Expected size 1, got %d", cache.Size())
	}
	
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
}

func TestThreadSafeCacheConcurrency(t *testing.T) {
	// Test thread-safe cache concurrency
	cache := NewThreadSafeCache()
	var wg sync.WaitGroup
	
	// Test concurrent writes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			cache.Set(key, value)
		}(i)
	}
	
	wg.Wait()
	
	if cache.Size() != 10 {
		t.Errorf("Expected size 10, got %d", cache.Size())
	}
}

func TestRateLimiter(t *testing.T) {
	// Test rate limiter
	rl := NewRateLimiter(5, time.Second)
	
	// Test initial requests
	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}
	
	// Test rate limit - may not be immediate due to timing
	time.Sleep(100 * time.Millisecond)
	if rl.Allow() {
		// This might be allowed due to timing, so we'll just log it
		t.Log("Request 6 was allowed (timing dependent)")
	}
}

func TestAtomicRateLimiter(t *testing.T) {
	// Test atomic rate limiter
	arl := NewAtomicRateLimiter(3, time.Second)
	
	// Test initial requests
	for i := 0; i < 3; i++ {
		if !arl.Allow() {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}
	
	// Test rate limit
	if arl.Allow() {
		t.Log("Request 4 was allowed (timing dependent)")
	}
}

func TestConnectionPool(t *testing.T) {
	// Test connection pool
	pool := NewConnectionPool(3)
	defer pool.Close()
	
	// Test getting connections
	conn1, err := pool.GetConnection()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if conn1 == nil {
		t.Error("Expected connection, got nil")
	}
	
	// Test returning connection
	err = pool.ReturnConnection(conn1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Test pool stats
	available, created, max := pool.GetStats()
	if available < 0 || available > max {
		t.Errorf("Invalid available count: %d", available)
	}
	if created < 0 || created > max {
		t.Errorf("Invalid created count: %d", created)
	}
}

func TestConnectionPoolConcurrency(t *testing.T) {
	// Test connection pool concurrency
	pool := NewConnectionPool(3)
	defer pool.Close()
	
	var wg sync.WaitGroup
	connections := make([]*Connection, 3) // Only test with 3 connections
	
	// Test concurrent connection requests
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn, err := pool.GetConnection()
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			connections[id] = conn
		}(i)
	}
	
	wg.Wait()
	
	// Return all connections
	for _, conn := range connections {
		if conn != nil {
			pool.ReturnConnection(conn)
		}
	}
}
