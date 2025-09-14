package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WorkItem represents a unit of work to be processed
type WorkItem struct {
	ID       int
	Data     string
	Priority int
	Created  time.Time
}

// Result represents the result of processing a work item
type Result struct {
	ID        int
	Processed string
	WorkerID  int
	Duration  time.Duration
	Error     error
	Timestamp time.Time
}

// Simple Fan-Out/Fan-In implementation
func simpleFanOutFanIn(input <-chan WorkItem, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers*2)
	
	// Create worker input channels
	workerInputs := make([]chan WorkItem, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerInputs[i] = make(chan WorkItem, 2)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int, workerInput <-chan WorkItem) {
			defer wg.Done()
			for item := range workerInput {
				start := time.Now()
				
				// Simulate work processing
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				
				// Simulate occasional errors
				var err error
				if rand.Float32() < 0.1 { // 10% error rate
					err = fmt.Errorf("processing failed for item %d", item.ID)
				}
				
				// Send result
				result := Result{
					ID:        item.ID,
					Processed: fmt.Sprintf("Worker %d processed: %s", workerID, item.Data),
					WorkerID:  workerID,
					Duration:  time.Since(start),
					Error:     err,
					Timestamp: time.Now(),
				}
				
				output <- result
			}
		}(i, workerInputs[i])
	}
	
	// Fan-out: Distribute work to workers
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workerInputs[i])
			}
		}()
		
		workerIndex := 0
		for item := range input {
			workerInputs[workerIndex] <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	// Close output channel when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

// Example 1: Basic Fan-Out/Fan-In
func basicFanOutFanInExample() {
	fmt.Println("\n1. Basic Fan-Out/Fan-In")
	fmt.Println("========================")
	
	const numItems = 20
	const numWorkers = 4
	
	input := make(chan WorkItem, numItems)
	output := simpleFanOutFanIn(input, numWorkers)
	
	// Send work items
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Basic Fan-Out/Fan-In Results:")
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nProcessed %d items with %d workers\n", len(results), numWorkers)
}

// Example 2: Buffered Fan-Out/Fan-In
func bufferedFanOutFanInExample() {
	fmt.Println("\n2. Buffered Fan-Out/Fan-In")
	fmt.Println("===========================")
	
	const numItems = 30
	const numWorkers = 6
	const bufferSize = 5
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, bufferSize)
	
	// Create worker input channels
	workerInputs := make([]chan WorkItem, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerInputs[i] = make(chan WorkItem, bufferSize)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int, workerInput <-chan WorkItem) {
			defer wg.Done()
			for item := range workerInput {
				start := time.Now()
				
				// Simulate work processing
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				
				// Simulate occasional errors
				var err error
				if rand.Float32() < 0.1 { // 10% error rate
					err = fmt.Errorf("processing failed for item %d", item.ID)
				}
				
				// Send result
				result := Result{
					ID:        item.ID,
					Processed: fmt.Sprintf("Worker %d processed: %s", workerID, item.Data),
					WorkerID:  workerID,
					Duration:  time.Since(start),
					Error:     err,
					Timestamp: time.Now(),
				}
				
				output <- result
			}
		}(i, workerInputs[i])
	}
	
	// Fan-out with buffering
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workerInputs[i])
			}
		}()
		
		workerIndex := 0
		for i := 0; i < numItems; i++ {
			item := WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Buffered Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
			workerInputs[workerIndex] <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	// Close output channel when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Collect results
	fmt.Println("Buffered Fan-Out/Fan-In Results:")
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nProcessed %d items with %d workers (buffered)\n", len(results), numWorkers)
}

// Example 3: Priority-Based Fan-Out/Fan-In
func priorityFanOutFanInExample() {
	fmt.Println("\n3. Priority-Based Fan-Out/Fan-In")
	fmt.Println("=================================")
	
	const numItems = 25
	const numWorkers = 3
	
	input := make(chan WorkItem, numItems)
	output := make(chan Result, numWorkers*2)
	
	// Create priority channels
	highPriority := make(chan WorkItem, 10)
	mediumPriority := make(chan WorkItem, 10)
	lowPriority := make(chan WorkItem, 10)
	
	// Create worker input channels
	workerInputs := make([]chan WorkItem, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerInputs[i] = make(chan WorkItem, 5)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int, workerInput <-chan WorkItem) {
			defer wg.Done()
			for item := range workerInput {
				start := time.Now()
				
				// Simulate work processing
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				
				// Simulate occasional errors
				var err error
				if rand.Float32() < 0.1 { // 10% error rate
					err = fmt.Errorf("processing failed for item %d", item.ID)
				}
				
				// Send result
				result := Result{
					ID:        item.ID,
					Processed: fmt.Sprintf("Worker %d processed: %s", workerID, item.Data),
					WorkerID:  workerID,
					Duration:  time.Since(start),
					Error:     err,
					Timestamp: time.Now(),
				}
				
				output <- result
			}
		}(i, workerInputs[i])
	}
	
	// Priority distributor
	go func() {
		defer func() {
			close(highPriority)
			close(mediumPriority)
			close(lowPriority)
		}()
		
		for i := 0; i < numItems; i++ {
			item := WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Priority Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
			
			switch item.Priority {
			case 0:
				highPriority <- item
			case 1:
				mediumPriority <- item
			default:
				lowPriority <- item
			}
		}
	}()
	
	// Fan-out to workers based on priority
	go func() {
		defer func() {
			for i := 0; i < numWorkers; i++ {
				close(workerInputs[i])
			}
		}()
		
		workerIndex := 0
		
		// Process high priority first
		for item := range highPriority {
			workerInputs[workerIndex] <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
		
		// Process medium priority
		for item := range mediumPriority {
			workerInputs[workerIndex] <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
		
		// Process low priority
		for item := range lowPriority {
			workerInputs[workerIndex] <- item
			workerIndex = (workerIndex + 1) % numWorkers
		}
	}()
	
	// Close output channel when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Collect results
	fmt.Println("Priority-Based Fan-Out/Fan-In Results:")
	var results []Result
	for result := range output {
		results = append(results, result)
		if result.Error != nil {
			fmt.Printf("  ERROR: Item %d failed: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("  SUCCESS: %s (Worker %d, Duration: %v)\n", 
				result.Processed, result.WorkerID, result.Duration)
		}
	}
	
	fmt.Printf("\nProcessed %d items with priority-based distribution\n", len(results))
}

// Example 4: Performance Comparison
func performanceComparisonExample() {
	fmt.Println("\n4. Performance Comparison")
	fmt.Println("=========================")
	
	const numItems = 1000
	const numWorkers = 8
	
	// Sequential processing
	fmt.Println("Sequential Processing:")
	start := time.Now()
	for i := 0; i < numItems; i++ {
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	sequentialTime := time.Since(start)
	fmt.Printf("  Sequential time: %v\n", sequentialTime)
	
	// Fan-out/Fan-in processing
	fmt.Println("\nFan-Out/Fan-In Processing:")
	input := make(chan WorkItem, numItems)
	output := simpleFanOutFanIn(input, numWorkers)
	
	start = time.Now()
	
	// Send work
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- WorkItem{
				ID:       i,
				Data:     fmt.Sprintf("Perf Item %d", i),
				Priority: rand.Intn(3),
				Created:  time.Now(),
			}
		}
	}()
	
	// Collect results
	var results []Result
	for result := range output {
		results = append(results, result)
	}
	
	fanOutFanInTime := time.Since(start)
	fmt.Printf("  Fan-Out/Fan-In time: %v\n", fanOutFanInTime)
	fmt.Printf("  Processed %d items\n", len(results))
	
	// Performance comparison
	speedup := float64(sequentialTime) / float64(fanOutFanInTime)
	fmt.Printf("\nPerformance Comparison:\n")
	fmt.Printf("  Sequential: %v\n", sequentialTime)
	fmt.Printf("  Fan-Out/Fan-In: %v\n", fanOutFanInTime)
	fmt.Printf("  Speedup: %.2fx\n", speedup)
	fmt.Printf("  Efficiency: %.2f%%\n", speedup/float64(numWorkers)*100)
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🚀 Fan-Out/Fan-In Pattern Examples")
	fmt.Println("===================================")
	
	basicFanOutFanInExample()
	bufferedFanOutFanInExample()
	priorityFanOutFanInExample()
	performanceComparisonExample()
}
