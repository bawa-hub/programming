package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Data represents input data for the pipeline
type Data struct {
	ID    int
	Value string
	Key   string
}

// ProcessedData represents data after processing
type ProcessedData struct {
	ID    int
	Value string
	Key   string
	Stage string
}

// Result represents the final result
type Result struct {
	ID      int
	Value   string
	Key     string
	Stages  []string
	Duration time.Duration
}

// Example 1: Basic Pipeline
func basicPipeline() {
	fmt.Println("1. Basic Pipeline")
	fmt.Println("=================")
	
	const numItems = 10
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1: Process input
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			// Simulate processing
			time.Sleep(50 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Stage1: %s", data.Value),
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
			// Simulate processing
			time.Sleep(50 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Stage 3: Final processing
	go func() {
		defer close(output)
		for data := range stage2 {
			// Simulate processing
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Final: %s", data.Value),
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
				Value: fmt.Sprintf("Item %d", i),
				Key:   fmt.Sprintf("key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Basic Pipeline Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 2: Buffered Pipeline
func bufferedPipeline() {
	fmt.Println("\n2. Buffered Pipeline")
	fmt.Println("===================")
	
	const numItems = 10
	const bufferSize = 5
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
	fmt.Println("Buffered Pipeline Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 3: Fan-Out/Fan-In Pipeline
func fanOutFanInPipeline() {
	fmt.Println("\n3. Fan-Out/Fan-In Pipeline")
	fmt.Println("==========================")
	
	const numItems = 12
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
				// Simulate work
				time.Sleep(time.Duration(50+workerID*10) * time.Millisecond)
				result := Result{
					ID:      data.ID,
					Value:   fmt.Sprintf("Worker %d: %s", workerID, data.Value),
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
			// Round-robin distribution
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
	fmt.Println("Fan-Out/Fan-In Pipeline Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 4: Pipeline with Error Handling
func pipelineWithErrorHandling() {
	fmt.Println("\n4. Pipeline with Error Handling")
	fmt.Println("===============================")
	
	const numItems = 10
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	errors := make(chan error, numItems*2) // Make it larger to avoid blocking
	
	// Stage 1 with error handling
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			// Simulate occasional errors
			if data.ID%4 == 0 {
				select {
				case errors <- fmt.Errorf("stage1 failed for item %d", data.ID):
				default:
					// If error channel is full, skip this error
				}
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
			// Simulate occasional errors
			if data.ID%3 == 0 {
				select {
				case errors <- fmt.Errorf("stage2 failed for item %d", data.ID):
				default:
					// If error channel is full, skip this error
				}
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
	fmt.Println("Pipeline with Error Handling Results:")
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

// Example 5: Pipeline with Timeout
func pipelineWithTimeout() {
	fmt.Println("\n5. Pipeline with Timeout")
	fmt.Println("========================")
	
	const numItems = 8
	const timeout = 2 * time.Second
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
				
				// Simulate work with timeout
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
	fmt.Println("Pipeline with Timeout Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 6: Pipeline with Rate Limiting
func pipelineWithRateLimiting() {
	fmt.Println("\n6. Pipeline with Rate Limiting")
	fmt.Println("=============================")
	
	const numItems = 8
	const rateLimit = 2 // items per second
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
			<-rateLimiter.C // Wait for rate limiter
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
			<-rateLimiter.C // Wait for rate limiter
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
			<-rateLimiter.C // Wait for rate limiter
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
	fmt.Println("Pipeline with Rate Limiting Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 7: Pipeline with Metrics
func pipelineWithMetrics() {
	fmt.Println("\n7. Pipeline with Metrics")
	fmt.Println("=======================")
	
	const numItems = 10
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	metrics := &BasicPipelineMetrics{}
	
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
	fmt.Println("Pipeline with Metrics Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
	
	// Print metrics
	fmt.Printf("\nPipeline Metrics:\n")
	fmt.Printf("  Stage1 Total Time: %v\n", metrics.getStageTime("stage1"))
	fmt.Printf("  Stage2 Total Time: %v\n", metrics.getStageTime("stage2"))
	fmt.Printf("  Stage3 Total Time: %v\n", metrics.getStageTime("stage3"))
	fmt.Printf("  Total Items Processed: %d\n", metrics.getTotalItems())
}

type BasicPipelineMetrics struct {
	stage1Time    time.Duration
	stage2Time    time.Duration
	stage3Time    time.Duration
	totalItems    int64
	mu            sync.RWMutex
}

func (pm *BasicPipelineMetrics) recordStage(stage string, duration time.Duration) {
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

func (pm *BasicPipelineMetrics) getStageTime(stage string) time.Duration {
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

func (pm *BasicPipelineMetrics) getTotalItems() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.totalItems
}

// Example 8: Pipeline with Backpressure
func pipelineWithBackpressure() {
	fmt.Println("\n8. Pipeline with Backpressure")
	fmt.Println("=============================")
	
	const numItems = 8
	const backpressureTimeout = 100 * time.Millisecond
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1 with backpressure
	stage1 := make(chan ProcessedData, 2) // Small buffer to trigger backpressure
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
	stage2 := make(chan ProcessedData, 2) // Small buffer to trigger backpressure
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
	fmt.Println("Pipeline with Backpressure Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}

// Example 9: Performance Comparison
func performanceComparison() {
	fmt.Println("\n9. Performance Comparison")
	fmt.Println("=========================")
	
	const numItems = 1000
	
	// Sequential processing
	start := time.Now()
	for i := 0; i < numItems; i++ {
		// Simulate work
		time.Sleep(1 * time.Millisecond)
	}
	sequentialTime := time.Since(start)
	
	// Pipeline processing
	start = time.Now()
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Simple pipeline
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			time.Sleep(1 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: data.Value,
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	go func() {
		defer close(output)
		for data := range stage1 {
			time.Sleep(1 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   data.Value,
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2"},
				Duration: 2 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Item %d", i),
				Key:   fmt.Sprintf("key%d", i),
			}
		}
	}()
	
	// Collect results
	for range output {
		// Just consume results
	}
	
	pipelineTime := time.Since(start)
	
	fmt.Printf("Sequential processing: %v\n", sequentialTime)
	fmt.Printf("Pipeline processing: %v\n", pipelineTime)
	fmt.Printf("Pipeline is %.2fx faster\n", float64(sequentialTime)/float64(pipelineTime))
}

// Example 10: Common Pitfalls
func commonPitfalls() {
	fmt.Println("\n10. Common Pitfalls")
	fmt.Println("==================")
	
	fmt.Println("Pitfall 1: Channel Leaks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// stage1 := make(chan ProcessedData)")
	fmt.Println("// go func() {")
	fmt.Println("//     for data := range input {")
	fmt.Println("//         stage1 <- processStage1(data)")
	fmt.Println("//     }")
	fmt.Println("//     // Forgot to close stage1")
	fmt.Println("// }()")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// stage1 := make(chan ProcessedData)")
	fmt.Println("// go func() {")
	fmt.Println("//     defer close(stage1)")
	fmt.Println("//     for data := range input {")
	fmt.Println("//         stage1 <- processStage1(data)")
	fmt.Println("//     }")
	fmt.Println("// }()")
	
	fmt.Println("\nPitfall 2: Deadlocks")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// stage1 := make(chan ProcessedData) // Unbuffered")
	fmt.Println("// stage1 <- processStage1(data) // Can block")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// stage1 := make(chan ProcessedData, bufferSize) // Buffered")
	fmt.Println("// stage1 <- processStage1(data) // Won't block")
	
	fmt.Println("\nPitfall 3: Incorrect Buffer Sizing")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// stage1 := make(chan ProcessedData, 1) // Too small")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// stage1 := make(chan ProcessedData, bufferSize) // Appropriate size")
	
	fmt.Println("\nPitfall 4: Missing Error Handling")
	fmt.Println("// ❌ Wrong:")
	fmt.Println("// stage1 <- processStage1(data) // Can panic")
	fmt.Println("//")
	fmt.Println("// ✅ Correct:")
	fmt.Println("// defer func() {")
	fmt.Println("//     if r := recover(); r != nil {")
	fmt.Printf("//         errors <- fmt.Errorf(\"stage1 panicked: %%v\", r)\n")
	fmt.Println("//     }")
	fmt.Println("// }()")
	fmt.Println("// processed, err := processStage1WithError(data)")
	fmt.Println("// if err != nil { errors <- err } else { stage1 <- processed }")
}

// Utility function to show pipeline info
func showPipelineInfo() {
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
