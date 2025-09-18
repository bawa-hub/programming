package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Exercise 1: Basic Fan-Out/Fan-In Implementation
func Exercise1() {
	fmt.Println("\nExercise 1: Basic Fan-Out/Fan-In Implementation")
	fmt.Println("================================================")
	
	const numItems = 15
	const numWorkers = 3
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// TODO: Implement basic fan-out/fan-in
	// 1. Create workers
	// 2. Start workers
	// 3. Distribute work to workers
	// 4. Collect results
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise1 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d)\n", result.Processed, result.WorkerID)
	}
	
	fmt.Printf("Exercise 1 completed: %d items processed\n", len(results))
}

// Exercise 2: Buffered Fan-Out/Fan-In with Custom Buffer Sizes
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered Fan-Out/Fan-In with Custom Buffer Sizes")
	fmt.Println("============================================================")
	
	const numItems = 25
	const numWorkers = 4
	const inputBufferSize = 10
	const outputBufferSize = 20
	
	input := make(chan WorkItem, inputBufferSize)
	output := make(chan Result, outputBufferSize)
	
	// TODO: Implement buffered fan-out/fan-in
	// 1. Create workers with custom buffer sizes
	// 2. Implement buffered distribution
	// 3. Handle buffer overflow scenarios
	// 4. Optimize for throughput
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise2 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Buffer: %d)\n", 
			result.Processed, result.WorkerID, outputBufferSize)
	}
	
	fmt.Printf("Exercise 2 completed: %d items processed with buffering\n", len(results))
}

// Exercise 3: Fan-Out/Fan-In with Dynamic Workers
func Exercise3() {
	fmt.Println("\nExercise 3: Fan-Out/Fan-In with Dynamic Workers")
	fmt.Println("===============================================")
	
	const numItems = 30
	const minWorkers = 2
	const maxWorkers = 6
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, maxWorkers*2)
	
	// TODO: Implement dynamic worker scaling
	// 1. Start with minimum workers
	// 2. Scale up based on load
	// 3. Scale down when load decreases
	// 4. Handle worker lifecycle
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise3 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Dynamic)\n", 
			result.Processed, result.WorkerID)
	}
	
	fmt.Printf("Exercise 3 completed: %d items processed with dynamic workers\n", len(results))
}

// Exercise 4: Error Handling with Recovery
func Exercise4() {
	fmt.Println("\nExercise 4: Error Handling with Recovery")
	fmt.Println("========================================")
	
	const numItems = 20
	const numWorkers = 3
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	errors := make(chan error, numWorkers*2)
	
	// TODO: Implement error handling with recovery
	// 1. Handle worker panics
	// 2. Implement retry logic
	// 3. Graceful error recovery
	// 4. Error aggregation and reporting
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise4 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Error handler
	go func() {
		defer close(errors)
		time.Sleep(3 * time.Second)
	}()
	
	// Print errors
	go func() {
		for err := range errors {
			fmt.Printf("  ERROR: %v\n", err)
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  FAILED: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d)\n", 
				result.Processed, result.WorkerID)
		}
	}
	
	fmt.Printf("Exercise 4 completed: %d items processed with error handling\n", len(results))
}

// Exercise 5: Timeout and Cancellation
func Exercise5() {
	fmt.Println("\nExercise 5: Timeout and Cancellation")
	fmt.Println("====================================")
	
	const numItems = 25
	const numWorkers = 4
	const timeout = 3 * time.Second
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	// TODO: Implement timeout and cancellation
	// 1. Use context for cancellation
	// 2. Handle timeout scenarios
	// 3. Graceful shutdown of workers
	// 4. Resource cleanup
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			select {
			case input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise5 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}:
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Collect results with timeout
	var results []Result
	timeoutChan := time.After(timeout)
	
	done := make(chan bool)
	go func() {
		defer close(done)
		for result := range output {
			results = append(results, result)
			fmt.Printf("  Result: %s (Worker %d)\n", 
				result.Processed, result.WorkerID)
		}
	}()
	
	select {
	case <-done:
		fmt.Printf("Exercise 5 completed: %d items processed\n", len(results))
	case <-timeoutChan:
		fmt.Printf("Exercise 5 timeout: %d items processed before timeout\n", len(results))
	}
}

// Exercise 6: Rate Limiting with Different Strategies
func Exercise6() {
	fmt.Println("\nExercise 6: Rate Limiting with Different Strategies")
	fmt.Println("==================================================")
	
	const numItems = 30
	const numWorkers = 3
	const rateLimit = 5 // items per second
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// TODO: Implement rate limiting strategies
	// 1. Token bucket rate limiting
	// 2. Sliding window rate limiting
	// 3. Fixed window rate limiting
	// 4. Adaptive rate limiting
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise6 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Rate Limited)\n", 
			result.Processed, result.WorkerID)
	}
	
	fmt.Printf("Exercise 6 completed: %d items processed with rate limiting\n", len(results))
}

// Exercise 7: Metrics Collection and Analysis
func Exercise7() {
	fmt.Println("\nExercise 7: Metrics Collection and Analysis")
	fmt.Println("===========================================")
	
	const numItems = 40
	const numWorkers = 5
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// TODO: Implement comprehensive metrics
	// 1. Throughput metrics
	// 2. Latency metrics
	// 3. Error rate metrics
	// 4. Worker utilization metrics
	// 5. Resource usage metrics
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise7 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results and calculate metrics
	var results []Result
	startTime := time.Now()
	
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Duration: %v)\n", 
			result.Processed, result.WorkerID, result.Duration)
	}
	
	totalDuration := time.Since(startTime)
	
	// Calculate and display metrics
	fmt.Printf("\nExercise 7 Metrics:\n")
	fmt.Printf("  Total Items: %d\n", len(results))
	fmt.Printf("  Total Duration: %v\n", totalDuration)
	fmt.Printf("  Throughput: %.2f items/sec\n", float64(len(results))/totalDuration.Seconds())
	fmt.Printf("  Average Latency: %v\n", calculateAverageLatency(results))
	fmt.Printf("  Error Rate: %.2f%%\n", calculateErrorRate(results))
}

// Exercise 8: Backpressure with Adaptive Buffering
func Exercise8() {
	fmt.Println("\nExercise 8: Backpressure with Adaptive Buffering")
	fmt.Println("===============================================")
	
	const numItems = 35
	const numWorkers = 4
	const initialBufferSize = 5
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// TODO: Implement adaptive buffering
	// 1. Monitor queue lengths
	// 2. Adjust buffer sizes dynamically
	// 3. Implement backpressure signals
	// 4. Handle overflow scenarios
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise8 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Adaptive Buffer)\n", 
			result.Processed, result.WorkerID)
	}
	
	fmt.Printf("Exercise 8 completed: %d items processed with adaptive buffering\n", len(results))
}

// Exercise 9: Circuit Breaker Integration
func Exercise9() {
	fmt.Println("\nExercise 9: Circuit Breaker Integration")
	fmt.Println("=======================================")
	
	const numItems = 25
	const numWorkers = 3
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	errors := make(chan error, numWorkers*2)
	
	// TODO: Implement circuit breaker
	// 1. Monitor error rates
	// 2. Open circuit when threshold exceeded
	// 3. Half-open state for testing
	// 4. Automatic recovery
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise9 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Error handler
	go func() {
		defer close(errors)
		time.Sleep(3 * time.Second)
	}()
	
	// Print errors
	go func() {
		for err := range errors {
			fmt.Printf("  ERROR: %v\n", err)
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  FAILED: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d)\n", 
				result.Processed, result.WorkerID)
		}
	}
	
	fmt.Printf("Exercise 9 completed: %d items processed with circuit breaker\n", len(results))
}

// Exercise 10: Caching Layer Integration
func Exercise10() {
	fmt.Println("\nExercise 10: Caching Layer Integration")
	fmt.Println("======================================")
	
	const numItems = 30
	const numWorkers = 4
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// TODO: Implement caching layer
	// 1. Check cache before processing
	// 2. Store results in cache
	// 3. Cache invalidation strategies
	// 4. Cache hit/miss metrics
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Exercise10 Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
		fmt.Printf("  Result: %s (Worker %d, Cached)\n", 
			result.Processed, result.WorkerID)
	}
	
	fmt.Printf("Exercise 10 completed: %d items processed with caching\n", len(results))
}

// Helper functions for exercises

func calculateAverageLatency(results []Result) time.Duration {
	if len(results) == 0 {
		return 0
	}
	
	var total time.Duration
	for _, result := range results {
		total += result.Duration
	}
	
	return total / time.Duration(len(results))
}

func calculateErrorRate(results []Result) float64 {
	if len(results) == 0 {
		return 0
	}
	
	var errors int
	for _, result := range results {
		if result.Error != nil {
			errors++
		}
	}
	
	return float64(errors) / float64(len(results)) * 100
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Fan-Out/Fan-In Pattern Exercises")
	fmt.Println("====================================")
	
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
