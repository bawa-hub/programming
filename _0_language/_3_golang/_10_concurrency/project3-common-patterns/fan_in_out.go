package main

import (
	"fmt"
	"sync"
	"time"
)

// FanOutFanIn demonstrates the fan-out/fan-in pattern
func FanOutFanIn() {
	fmt.Println("=== Fan-Out/Fan-In Pattern ===")
	
	// Input channel
	input := make(chan int)
	
	// Fan-out: Distribute work to multiple workers
	worker1 := fanOutWorker(input, 1)
	worker2 := fanOutWorker(input, 2)
	worker3 := fanOutWorker(input, 3)
	
	// Fan-in: Collect results from all workers
	output := fanIn(worker1, worker2, worker3)
	
	// Send work
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("  %s\n", result)
	}
}

// fanOutWorker creates a worker that processes input
func fanOutWorker(input <-chan int, workerID int) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		for n := range input {
			// Simulate work
			time.Sleep(100 * time.Millisecond)
			result := fmt.Sprintf("Worker %d: %d -> %d", workerID, n, n*n)
			output <- result
		}
	}()
	
	return output
}

// fanIn collects results from multiple channels
func fanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for result := range ch {
				output <- result
			}
		}(input)
	}
	
	// Close output when all inputs are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

// AdvancedFanOutFanIn demonstrates advanced fan-out/fan-in with error handling
func AdvancedFanOutFanIn() {
	fmt.Println("\n=== Advanced Fan-Out/Fan-In with Error Handling ===")
	
	// Input channel
	input := make(chan int)
	
	// Fan-out: Distribute work to multiple workers
	worker1 := advancedFanOutWorker(input, 1)
	worker2 := advancedFanOutWorker(input, 2)
	worker3 := advancedFanOutWorker(input, 3)
	
	// Fan-in: Collect results from all workers
	output := advancedFanIn(worker1, worker2, worker3)
	
	// Send work
	go func() {
		for i := 1; i <= 15; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("  %s\n", result)
	}
}

// WorkerResult represents the result of a worker operation
type WorkerResult struct {
	WorkerID int
	Input    int
	Output   int
	Error    error
}

// advancedFanOutWorker creates a worker that processes input with error handling
func advancedFanOutWorker(input <-chan int, workerID int) <-chan WorkerResult {
	output := make(chan WorkerResult)
	
	go func() {
		defer close(output)
		for n := range input {
			// Simulate work with occasional errors
			time.Sleep(50 * time.Millisecond)
			
			var result WorkerResult
			result.WorkerID = workerID
			result.Input = n
			
			// Simulate error for certain inputs
			if n%7 == 0 {
				result.Error = fmt.Errorf("simulated error for input %d", n)
			} else {
				result.Output = n * n
			}
			
			output <- result
		}
	}()
	
	return output
}

// advancedFanIn collects results from multiple channels with error handling
func advancedFanIn(inputs ...<-chan WorkerResult) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan WorkerResult) {
			defer wg.Done()
			for result := range ch {
				if result.Error != nil {
					output <- fmt.Sprintf("Worker %d: ERROR processing %d - %v", 
						result.WorkerID, result.Input, result.Error)
				} else {
					output <- fmt.Sprintf("Worker %d: %d -> %d", 
						result.WorkerID, result.Input, result.Output)
				}
			}
		}(input)
	}
	
	// Close output when all inputs are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

// DynamicFanOutFanIn demonstrates dynamic fan-out/fan-in
func DynamicFanOutFanIn() {
	fmt.Println("\n=== Dynamic Fan-Out/Fan-In ===")
	
	// Input channel
	input := make(chan int)
	
	// Create workers dynamically based on workload
	numWorkers := 5
	workers := make([]<-chan string, numWorkers)
	
	for i := 0; i < numWorkers; i++ {
		workers[i] = dynamicFanOutWorker(input, i+1)
	}
	
	// Fan-in: Collect results from all workers
	output := fanIn(workers...)
	
	// Send work
	go func() {
		for i := 1; i <= 20; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("  %s\n", result)
	}
}

// dynamicFanOutWorker creates a worker that processes input with dynamic load balancing
func dynamicFanOutWorker(input <-chan int, workerID int) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		for n := range input {
			// Simulate work with variable duration
			duration := time.Duration(n%3+1) * 50 * time.Millisecond
			time.Sleep(duration)
			
			result := fmt.Sprintf("Worker %d: %d -> %d (took %v)", 
				workerID, n, n*n, duration)
			output <- result
		}
	}()
	
	return output
}

// FanOutFanInWithTimeout demonstrates fan-out/fan-in with timeout
func FanOutFanInWithTimeout() {
	fmt.Println("\n=== Fan-Out/Fan-In with Timeout ===")
	
	// Input channel
	input := make(chan int)
	
	// Fan-out: Distribute work to multiple workers
	worker1 := timeoutFanOutWorker(input, 1)
	worker2 := timeoutFanOutWorker(input, 2)
	worker3 := timeoutFanOutWorker(input, 3)
	
	// Fan-in: Collect results from all workers
	output := timeoutFanIn(worker1, worker2, worker3)
	
	// Send work
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results with timeout
	timeout := time.After(3 * time.Second)
	
	fmt.Println("Results:")
	for {
		select {
		case result, ok := <-output:
			if !ok {
				fmt.Println("All results received")
				return
			}
			fmt.Printf("  %s\n", result)
		case <-timeout:
			fmt.Println("Timeout reached, stopping collection")
			return
		}
	}
}

// timeoutFanOutWorker creates a worker that processes input with timeout
func timeoutFanOutWorker(input <-chan int, workerID int) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		for n := range input {
			// Simulate work with timeout
			done := make(chan string, 1)
			go func() {
				time.Sleep(200 * time.Millisecond)
				done <- fmt.Sprintf("Worker %d: %d -> %d", workerID, n, n*n)
			}()
			
			select {
			case result := <-done:
				output <- result
			case <-time.After(150 * time.Millisecond):
				output <- fmt.Sprintf("Worker %d: %d -> TIMEOUT", workerID, n)
			}
		}
	}()
	
	return output
}

// timeoutFanIn collects results from multiple channels with timeout
func timeoutFanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for result := range ch {
				output <- result
			}
		}(input)
	}
	
	// Close output when all inputs are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

// FanOutFanInWithBatching demonstrates fan-out/fan-in with batching
func FanOutFanInWithBatching() {
	fmt.Println("\n=== Fan-Out/Fan-In with Batching ===")
	
	// Input channel
	input := make(chan int)
	
	// Fan-out: Distribute work to multiple workers
	worker1 := batchFanOutWorker(input, 1, 3)
	worker2 := batchFanOutWorker(input, 2, 3)
	worker3 := batchFanOutWorker(input, 3, 3)
	
	// Fan-in: Collect results from all workers
	output := fanIn(worker1, worker2, worker3)
	
	// Send work
	go func() {
		for i := 1; i <= 15; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("  %s\n", result)
	}
}

// batchFanOutWorker creates a worker that processes input in batches
func batchFanOutWorker(input <-chan int, workerID int, batchSize int) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		
		batch := make([]int, 0, batchSize)
		
		for n := range input {
			batch = append(batch, n)
			
			if len(batch) >= batchSize {
				// Process batch
				result := fmt.Sprintf("Worker %d: processed batch %v", workerID, batch)
				output <- result
				batch = batch[:0] // Reset batch
			}
		}
		
		// Process remaining items
		if len(batch) > 0 {
			result := fmt.Sprintf("Worker %d: processed final batch %v", workerID, batch)
			output <- result
		}
	}()
	
	return output
}
