package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Exercise 1: Implement Error Propagation
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Error Propagation")
	fmt.Println("======================================")
	
	// TODO: Implement error propagation across goroutines
	// 1. Create a function that processes data and can fail
	// 2. Use channels to propagate errors
	// 3. Handle errors in the main goroutine
	
	dataCh := make(chan int)
	errorCh := make(chan error)
	
	go func() {
		defer close(dataCh)
		defer close(errorCh)
		
		for i := 0; i < 5; i++ {
			if i == 3 {
				errorCh <- fmt.Errorf("processing failed at %d", i)
				return
			}
			dataCh <- i
		}
	}()
	
	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("  Exercise 1: Processed %d\n", data)
		case err := <-errorCh:
			fmt.Printf("  Exercise 1: Error %v\n", err)
			return
		}
	}
}

// Exercise 2: Implement Error Aggregation
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Error Aggregation")
	fmt.Println("======================================")
	
	// TODO: Implement error aggregation from multiple goroutines
	// 1. Create multiple goroutines that can fail
	// 2. Collect all errors
	// 3. Return aggregated error if any
	
	collector := &ErrorCollector{}
	
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := processItem(fmt.Sprintf("item%d", id)); err != nil {
				collector.Add(fmt.Errorf("item %d failed: %w", id, err))
			}
		}(i)
	}
	
	wg.Wait()
	
	if err := collector.Error(); err != nil {
		fmt.Printf("  Exercise 2: Aggregated errors: %v\n", err)
	} else {
		fmt.Println("  Exercise 2: All items processed successfully")
	}
}

// Exercise 3: Implement Panic Recovery
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Panic Recovery")
	fmt.Println("===================================")
	
	// TODO: Implement panic recovery in goroutines
	// 1. Create a goroutine that can panic
	// 2. Use defer and recover to handle panics
	// 3. Report panics as errors
	
	errorCh := make(chan error, 1)
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errorCh <- fmt.Errorf("panic recovered: %v", r)
			}
		}()
		
		// This will panic
		panic("exercise panic")
	}()
	
	select {
	case err := <-errorCh:
		fmt.Printf("  Exercise 3: Panic recovered: %v\n", err)
	case <-time.After(1 * time.Second):
		fmt.Println("  Exercise 3: No panic occurred")
	}
}

// Exercise 4: Implement Error Context
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Error Context")
	fmt.Println("==================================")
	
	// TODO: Implement error context preservation
	// 1. Create a function that can fail
	// 2. Wrap errors with context information
	// 3. Preserve error chain
	
	if err := processWithContext(); err != nil {
		fmt.Printf("  Exercise 4: Contextual error: %v\n", err)
	}
}

func processWithContext() error {
	context := map[string]interface{}{
		"operation": "data processing",
		"timestamp": time.Now(),
		"user_id":   "123",
	}
	
	if err := processData(); err != nil {
		return ErrorWrapper{
			Operation: "data processing",
			Context:   context,
			Err:       err,
		}
	}
	
	return nil
}

// Exercise 5: Implement Circuit Breaker
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Circuit Breaker")
	fmt.Println("====================================")
	
	// TODO: Implement circuit breaker pattern
	// 1. Create a circuit breaker
	// 2. Test with failing operations
	// 3. Verify circuit breaker behavior
	
	breaker := &CircuitBreaker{
		threshold: 3,
		timeout:   1 * time.Second,
	}
	
	// Test circuit breaker
	for i := 0; i < 5; i++ {
		err := breaker.Call(func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("  Exercise 5: Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("  Exercise 5: Call %d succeeded\n", i)
		}
	}
}

// Exercise 6: Implement Retry Mechanism
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Retry Mechanism")
	fmt.Println("====================================")
	
	// TODO: Implement retry mechanism with exponential backoff
	// 1. Create a function that can fail
	// 2. Implement retry with exponential backoff
	// 3. Test retry behavior
	
	if err := retryWithBackoff(func() error {
		return fmt.Errorf("operation failed")
	}, 3); err != nil {
		fmt.Printf("  Exercise 6: Retry failed: %v\n", err)
	}
}

func retryWithBackoff(fn func() error, maxRetries int) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			lastErr = err
			backoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
			log.Printf("  Exercise 6: Attempt %d failed, backing off for %v: %v", i+1, backoff, err)
			time.Sleep(backoff)
			continue
		}
		return nil
	}
	
	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Exercise 7: Implement Error Group
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Error Group")
	fmt.Println("=================================")
	
	// TODO: Implement error group for concurrent operations
	// 1. Create an error group
	// 2. Add multiple operations
	// 3. Wait for completion and collect errors
	
	group := &ErrorGroup{}
	
	for i := 0; i < 5; i++ {
		id := i
		group.Go(func() error {
			return processItem(fmt.Sprintf("item%d", id))
		})
	}
	
	time.Sleep(1 * time.Second) // Wait for completion
	
	if err := group.Wait(); err != nil {
		fmt.Printf("  Exercise 7: Group errors: %v\n", err)
	} else {
		fmt.Println("  Exercise 7: All operations completed successfully")
	}
}

// Exercise 8: Implement Timeout Pattern
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Timeout Pattern")
	fmt.Println("====================================")
	
	// TODO: Implement timeout pattern
	// 1. Create a context with timeout
	// 2. Run an operation with timeout
	// 3. Handle timeout errors
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	resultCh := make(chan error, 1)
	
	go func() {
		resultCh <- slowOperation()
	}()
	
	select {
	case err := <-resultCh:
		if err != nil {
			fmt.Printf("  Exercise 8: Operation failed: %v\n", err)
		} else {
			fmt.Println("  Exercise 8: Operation succeeded")
		}
	case <-ctx.Done():
		fmt.Printf("  Exercise 8: Operation timed out: %v\n", ctx.Err())
	}
}

func slowOperation() error {
	time.Sleep(2 * time.Second)
	return nil
}

// Exercise 9: Implement Fallback Pattern
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Fallback Pattern")
	fmt.Println("=====================================")
	
	// TODO: Implement fallback pattern
	// 1. Try primary operation
	// 2. If it fails, try fallback operation
	// 3. Handle both success and failure cases
	
	if err := fallbackPatternImpl(); err != nil {
		fmt.Printf("  Exercise 9: Fallback failed: %v\n", err)
	} else {
		fmt.Println("  Exercise 9: Fallback succeeded")
	}
}

func fallbackPatternImpl() error {
	// Try primary operation
	if err := primaryOperation(); err != nil {
		log.Printf("  Exercise 9: Primary operation failed: %v", err)
		
		// Try fallback operation
		if err := fallbackOperation(); err != nil {
			log.Printf("  Exercise 9: Fallback operation failed: %v", err)
			return fmt.Errorf("both operations failed")
		}
	}
	
	return nil
}

func primaryOperation() error {
	return fmt.Errorf("primary operation unavailable")
}

func fallbackOperation() error {
	return nil // Fallback succeeds
}

// Exercise 10: Implement Error Metrics
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Error Metrics")
	fmt.Println("===================================")
	
	// TODO: Implement error metrics collection
	// 1. Create error metrics collector
	// 2. Record different types of errors
	// 3. Display error statistics
	
	metrics := NewErrorMetrics()
	
	// Simulate some errors
	metrics.RecordError(fmt.Errorf("validation error"))
	metrics.RecordError(fmt.Errorf("network error"))
	metrics.RecordError(fmt.Errorf("validation error"))
	metrics.RecordError(fmt.Errorf("timeout error"))
	
	stats := metrics.GetStats()
	fmt.Printf("  Exercise 10: Error metrics: %v\n", stats)
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Error Handling in Concurrent Code Exercises")
	fmt.Println("=============================================")
	
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
