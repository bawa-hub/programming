package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Exercise 1: Basic Pipeline
// Create a basic pipeline with three stages.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Pipeline")
	fmt.Println("==========================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1: Process input
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			time.Sleep(50 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Exercise1 Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2: Process stage1 output
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			time.Sleep(50 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Exercise1 Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Stage 3: Final processing
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Exercise1 Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Exercise Item %d", i),
				Key:   fmt.Sprintf("exercise_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 1 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 2: Buffered Pipeline
// Implement a pipeline with buffered channels.
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered Pipeline")
	fmt.Println("=============================")
	
	const numItems = 8
	const bufferSize = 3
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1 with buffering
	stage1 := make(chan ProcessedData, bufferSize)
	go func() {
		defer close(stage1)
		for data := range input {
			time.Sleep(30 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Buffered Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with buffering
	stage2 := make(chan ProcessedData, bufferSize)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			time.Sleep(30 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Buffered Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with buffering
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(30 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Buffered Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 90 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Buffered Item %d", i),
				Key:   fmt.Sprintf("buffered_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 2 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 3: Fan-Out/Fan-In Pipeline
// Create a pipeline that distributes work to multiple workers.
func Exercise3() {
	fmt.Println("\nExercise 3: Fan-Out/Fan-In Pipeline")
	fmt.Println("===================================")
	
	const numItems = 10
	const numWorkers = 3
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Fan-out: Distribute work to multiple workers
	workers := make([]chan Data, numWorkers)
	for i := range workers {
		workers[i] = make(chan Data)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i, worker := range workers {
		wg.Add(1)
		go func(workerID int, jobs <-chan Data) {
			defer wg.Done()
			for data := range jobs {
				time.Sleep(time.Duration(50+workerID*10) * time.Millisecond)
				result := Result{
					ID:      data.ID,
					Value:   fmt.Sprintf("FanOut Worker %d: %s", workerID, data.Value),
					Key:     data.Key,
					Stages:  []string{"fanout", "worker", "fanin"},
					Duration: time.Duration(50+workerID*10) * time.Millisecond,
				}
				output <- result
			}
		}(i, worker)
	}
	
	// Distribute work
	go func() {
		defer func() {
			for _, worker := range workers {
				close(worker)
			}
		}()
		
		for data := range input {
			workers[data.ID%numWorkers] <- data
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("FanOut Item %d", i),
				Key:   fmt.Sprintf("fanout_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 3 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 4: Pipeline with Error Handling
// Implement a pipeline that handles errors from any stage.
func Exercise4() {
	fmt.Println("\nExercise 4: Pipeline with Error Handling")
	fmt.Println("=======================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	errors := make(chan error, numItems)
	
	// Stage 1 with error handling
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			if data.ID%3 == 0 {
				errors <- fmt.Errorf("stage1 failed for item %d", data.ID)
				continue
			}
			
			time.Sleep(50 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Error Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with error handling
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			if data.ID%2 == 0 {
				errors <- fmt.Errorf("stage2 failed for item %d", data.ID)
				continue
			}
			
			time.Sleep(50 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Error Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with error handling
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Error Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Close errors channel when done
	go func() {
		time.Sleep(2 * time.Second)
		close(errors)
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Error Item %d", i),
				Key:   fmt.Sprintf("error_key%d", i),
			}
		}
	}()
	
	// Collect results and errors
	fmt.Println("Exercise 4 Results:")
	for {
		select {
		case result, ok := <-output:
			if !ok {
				output = nil
			} else {
				fmt.Printf("  SUCCESS: Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
			}
		case err, ok := <-errors:
			if !ok {
				errors = nil
			} else {
				fmt.Printf("  ERROR: %v\n", err)
			}
		}
		
		if output == nil && errors == nil {
			break
		}
	}
}

// Exercise 5: Pipeline with Timeout
// Create a pipeline that handles timeouts for individual stages.
func Exercise5() {
	fmt.Println("\nExercise 5: Pipeline with Timeout")
	fmt.Println("=================================")
	
	const numItems = 6
	const timeout = 1 * time.Second
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	// Stage 1 with timeout
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for {
			select {
			case data, ok := <-input:
				if !ok {
					return
				}
				
				select {
				case <-time.After(time.Duration(data.ID*100) * time.Millisecond):
					stage1 <- ProcessedData{
						ID:    data.ID,
						Value: fmt.Sprintf("Timeout Stage1: %s", data.Value),
						Key:   data.Key,
						Stage: "stage1",
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Stage 2 with timeout
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for {
			select {
			case data, ok := <-stage1:
				if !ok {
					return
				}
				
				select {
				case <-time.After(50 * time.Millisecond):
					stage2 <- ProcessedData{
						ID:    data.ID,
						Value: fmt.Sprintf("Timeout Stage2: %s", data.Value),
						Key:   data.Key,
						Stage: "stage2",
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Final stage with timeout
	go func() {
		defer close(output)
		for {
			select {
			case data, ok := <-stage2:
				if !ok {
					return
				}
				
				select {
				case <-time.After(50 * time.Millisecond):
					output <- Result{
						ID:      data.ID,
						Value:   fmt.Sprintf("Timeout Final: %s", data.Value),
						Key:     data.Key,
						Stages:  []string{"stage1", "stage2", "stage3"},
						Duration: time.Duration(data.ID*100+100) * time.Millisecond,
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			select {
			case input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Timeout Item %d", i),
				Key:   fmt.Sprintf("timeout_key%d", i),
			}:
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 5 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 6: Pipeline with Rate Limiting
// Implement a pipeline that limits the rate of data processing.
func Exercise6() {
	fmt.Println("\nExercise 6: Pipeline with Rate Limiting")
	fmt.Println("=======================================")
	
	const numItems = 6
	const rateLimit = 3 // items per second
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Rate limiter
	rateLimiter := time.NewTicker(time.Second / rateLimit)
	defer rateLimiter.Stop()
	
	// Stage 1 with rate limiting
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Rate Limited Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with rate limiting
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Rate Limited Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with rate limiting
	go func() {
		defer close(output)
		for data := range stage2 {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Rate Limited Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 300 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Rate Limited Item %d", i),
				Key:   fmt.Sprintf("rate_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 6 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 7: Pipeline with Metrics
// Create a pipeline that collects and reports performance metrics.
func Exercise7() {
	fmt.Println("\nExercise 7: Pipeline with Metrics")
	fmt.Println("=================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	metrics := &ExercisePipelineMetrics{}
	
	// Stage 1 with metrics
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage1", time.Since(start))
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Metrics Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with metrics
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage2", time.Since(start))
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Metrics Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with metrics
	go func() {
		defer close(output)
		for data := range stage2 {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage3", time.Since(start))
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Metrics Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Metrics Item %d", i),
				Key:   fmt.Sprintf("metrics_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 7 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
	
	// Print metrics
	fmt.Printf("\nExercise 7 Metrics:\n")
	fmt.Printf("  Stage1 Total Time: %v\n", metrics.getStageTime("stage1"))
	fmt.Printf("  Stage2 Total Time: %v\n", metrics.getStageTime("stage2"))
	fmt.Printf("  Stage3 Total Time: %v\n", metrics.getStageTime("stage3"))
	fmt.Printf("  Total Items Processed: %d\n", metrics.getTotalItems())
}

type ExercisePipelineMetrics struct {
	stage1Time    time.Duration
	stage2Time    time.Duration
	stage3Time    time.Duration
	totalItems    int64
	mu            sync.RWMutex
}

func (pm *ExercisePipelineMetrics) recordStage(stage string, duration time.Duration) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.totalItems++
	switch stage {
	case "stage1":
		pm.stage1Time += duration
	case "stage2":
		pm.stage2Time += duration
	case "stage3":
		pm.stage3Time += duration
	}
}

func (pm *ExercisePipelineMetrics) getStageTime(stage string) time.Duration {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	switch stage {
	case "stage1":
		return pm.stage1Time
	case "stage2":
		return pm.stage2Time
	case "stage3":
		return pm.stage3Time
	default:
		return 0
	}
}

func (pm *ExercisePipelineMetrics) getTotalItems() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.totalItems
}

// Exercise 8: Pipeline with Backpressure
// Implement a pipeline that handles backpressure.
func Exercise8() {
	fmt.Println("\nExercise 8: Pipeline with Backpressure")
	fmt.Println("======================================")
	
	const numItems = 6
	const backpressureTimeout = 50 * time.Millisecond
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1 with backpressure
	stage1 := make(chan ProcessedData, 1) // Small buffer to trigger backpressure
	go func() {
		defer close(stage1)
		for data := range input {
			select {
			case stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Backpressure Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage1 for item %d\n", data.ID)
			}
		}
	}()
	
	// Stage 2 with backpressure
	stage2 := make(chan ProcessedData, 1) // Small buffer to trigger backpressure
	go func() {
		defer close(stage2)
		for data := range stage1 {
			select {
			case stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Backpressure Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage2 for item %d\n", data.ID)
			}
		}
	}()
	
	// Final stage with backpressure
	go func() {
		defer close(output)
		for data := range stage2 {
			select {
			case output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Backpressure Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage3 for item %d\n", data.ID)
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Backpressure Item %d", i),
				Key:   fmt.Sprintf("backpressure_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 8 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Exercise 9: Pipeline with Circuit Breaker
// Create a pipeline that uses circuit breakers.
func Exercise9() {
	fmt.Println("\nExercise 9: Pipeline with Circuit Breaker")
	fmt.Println("=========================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	errors := make(chan error, numItems)
	
	// Simple circuit breaker simulation
	breaker1 := &SimpleCircuitBreaker{threshold: 3, timeout: 500 * time.Millisecond}
	breaker2 := &SimpleCircuitBreaker{threshold: 3, timeout: 500 * time.Millisecond}
	
	// Stage 1 with circuit breaker
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			if breaker1.Allow() {
				// Simulate occasional failures
				if data.ID%4 == 0 {
					breaker1.RecordFailure()
					errors <- fmt.Errorf("stage1 failed for item %d", data.ID)
					continue
				}
				
				time.Sleep(50 * time.Millisecond)
				breaker1.RecordSuccess()
				stage1 <- ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Circuit Breaker Stage1: %s", data.Value),
					Key:   data.Key,
					Stage: "stage1",
				}
			} else {
				errors <- fmt.Errorf("circuit breaker open at stage1 for item %d", data.ID)
			}
		}
	}()
	
	// Stage 2 with circuit breaker
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			if breaker2.Allow() {
				// Simulate occasional failures
				if data.ID%3 == 0 {
					breaker2.RecordFailure()
					errors <- fmt.Errorf("stage2 failed for item %d", data.ID)
					continue
				}
				
				time.Sleep(50 * time.Millisecond)
				breaker2.RecordSuccess()
				stage2 <- ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Circuit Breaker Stage2: %s", data.Value),
					Key:   data.Key,
					Stage: "stage2",
				}
			} else {
				errors <- fmt.Errorf("circuit breaker open at stage2 for item %d", data.ID)
			}
		}
	}()
	
	// Final stage
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Circuit Breaker Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Close errors channel when done
	go func() {
		time.Sleep(2 * time.Second)
		close(errors)
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Circuit Breaker Item %d", i),
				Key:   fmt.Sprintf("circuit_key%d", i),
			}
		}
	}()
	
	// Collect results and errors
	fmt.Println("Exercise 9 Results:")
	for {
		select {
		case result, ok := <-output:
			if !ok {
				output = nil
			} else {
				fmt.Printf("  SUCCESS: Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
			}
		case err, ok := <-errors:
			if !ok {
				errors = nil
			} else {
				fmt.Printf("  ERROR: %v\n", err)
			}
		}
		
		if output == nil && errors == nil {
			break
		}
	}
}

type SimpleCircuitBreaker struct {
	failures    int
	threshold   int
	timeout     time.Duration
	lastFailure time.Time
	state       int // 0: closed, 1: open, 2: half-open
}

func (cb *SimpleCircuitBreaker) Allow() bool {
	if cb.state == 0 { // closed
		return true
	} else if cb.state == 1 { // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = 2 // half-open
			return true
		}
		return false
	} else { // half-open
		return true
	}
}

func (cb *SimpleCircuitBreaker) RecordSuccess() {
	if cb.state == 2 { // half-open
		cb.state = 0 // closed
		cb.failures = 0
	}
}

func (cb *SimpleCircuitBreaker) RecordFailure() {
	cb.failures++
	cb.lastFailure = time.Now()
	
	if cb.failures >= cb.threshold {
		cb.state = 1 // open
	}
}

// Exercise 10: Pipeline with Caching
// Implement a pipeline that caches results.
func Exercise10() {
	fmt.Println("\nExercise 10: Pipeline with Caching")
	fmt.Println("==================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	cache := NewSimpleCache()
	
	// Stage 1 with caching
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			// Check cache first
			if cached, found := cache.Get(data.Key); found {
				stage1 <- cached
			} else {
				time.Sleep(50 * time.Millisecond)
				processed := ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Cached Stage1: %s", data.Value),
					Key:   data.Key,
					Stage: "stage1",
				}
				cache.Set(data.Key, processed)
				stage1 <- processed
			}
		}
	}()
	
	// Stage 2 with caching
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			// Check cache first
			cacheKey := fmt.Sprintf("stage2_%s", data.Key)
			if cached, found := cache.Get(cacheKey); found {
				stage2 <- cached
			} else {
				time.Sleep(50 * time.Millisecond)
				processed := ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Cached Stage2: %s", data.Value),
					Key:   data.Key,
					Stage: "stage2",
				}
				cache.Set(cacheKey, processed)
				stage2 <- processed
			}
		}
	}()
	
	// Final stage
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Cached Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Cached Item %d", i),
				Key:   fmt.Sprintf("cached_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 10 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
	
	// Print cache stats
	fmt.Printf("\nExercise 10 Cache Stats:\n")
	fmt.Printf("  Cache Hits: %d\n", cache.GetHits())
	fmt.Printf("  Cache Misses: %d\n", cache.GetMisses())
}

type SimpleCache struct {
	data   map[string]ProcessedData
	hits   int64
	misses int64
	mu     sync.RWMutex
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		data: make(map[string]ProcessedData),
	}
}

func (c *SimpleCache) Get(key string) (ProcessedData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if value, found := c.data[key]; found {
		c.hits++
		return value, true
	}
	c.misses++
	return ProcessedData{}, false
}

func (c *SimpleCache) Set(key string, value ProcessedData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *SimpleCache) GetHits() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hits
}

func (c *SimpleCache) GetMisses() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.misses
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Pipeline Exercises")
	fmt.Println("=================================")
	
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
	
	fmt.Println("\n✅ All exercises completed!")
}
