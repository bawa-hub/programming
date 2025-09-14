package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 🧵 ADVANCED GOROUTINE PATTERNS
// Master the patterns that make Go concurrency powerful

func main() {
	fmt.Println("🚀 ADVANCED GOROUTINE PATTERNS")
	fmt.Println("===============================")

	// 1. Fan-In Pattern
	fmt.Println("\n1. Fan-In Pattern:")
	fanInPattern()

	// 2. Fan-Out Pattern
	fmt.Println("\n2. Fan-Out Pattern:")
	fanOutPattern()

	// 3. Pipeline Pattern
	fmt.Println("\n3. Pipeline Pattern:")
	pipelinePattern()

	// 4. Worker Pool Pattern
	fmt.Println("\n4. Worker Pool Pattern:")
	workerPoolPattern()

	// 5. Graceful Shutdown Pattern
	fmt.Println("\n5. Graceful Shutdown Pattern:")
	gracefulShutdownPattern()

	// 6. Goroutine Pool with Dynamic Scaling
	fmt.Println("\n6. Dynamic Scaling Pool:")
	dynamicScalingPool()
}

// FAN-IN PATTERN: Collecting results from multiple goroutines
func fanInPattern() {
	fmt.Println("Collecting results from multiple sources...")
	
	// Create multiple input channels
	input1 := make(chan int)
	input2 := make(chan int)
	input3 := make(chan int)
	
	// Start goroutines that send data
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- i * 10
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 1; i <= 3; i++ {
			input2 <- i * 20
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input3)
		for i := 1; i <= 3; i++ {
			input3 <- i * 30
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Fan-in: collect from all channels
	output := fanIn(input1, input2, input3)
	
	fmt.Println("  📊 Fan-in results:")
	for result := range output {
		fmt.Printf("    Received: %d\n", result)
	}
}

// Fan-in function that merges multiple channels
func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				output <- value
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

// FAN-OUT PATTERN: Distributing work to multiple workers
func fanOutPattern() {
	fmt.Println("Distributing work to multiple workers...")
	
	// Create work channel
	work := make(chan int, 10)
	
	// Fill work channel
	go func() {
		defer close(work)
		for i := 1; i <= 10; i++ {
			work <- i
		}
	}()
	
	// Create multiple workers
	numWorkers := 3
	results := make(chan int, 10)
	var wg sync.WaitGroup
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range work {
				fmt.Printf("  🧵 Worker %d processing job %d\n", workerID, job)
				time.Sleep(100 * time.Millisecond) // Simulate work
				results <- job * 2
			}
		}(i)
	}
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("  📊 Fan-out results:")
	for result := range results {
		fmt.Printf("    Result: %d\n", result)
	}
}

// PIPELINE PATTERN: Chaining processing stages
func pipelinePattern() {
	fmt.Println("Creating a processing pipeline...")
	
	// Stage 1: Generate numbers
	numbers := generateNumbers(5)
	
	// Stage 2: Square numbers
	squared := squareNumbers(numbers)
	
	// Stage 3: Filter even numbers
	even := filterEven(squared)
	
	// Stage 4: Collect results
	fmt.Println("  📊 Pipeline results:")
	for result := range even {
		fmt.Printf("    Pipeline result: %d\n", result)
	}
}

// Pipeline stage 1: Generate numbers
func generateNumbers(count int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i := 1; i <= count; i++ {
			output <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return output
}

// Pipeline stage 2: Square numbers
func squareNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for value := range input {
			fmt.Printf("  🧵 Squaring: %d\n", value)
			output <- value * value
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return output
}

// Pipeline stage 3: Filter even numbers
func filterEven(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for value := range input {
			if value%2 == 0 {
				fmt.Printf("  🧵 Filtering even: %d\n", value)
				output <- value
			}
		}
	}()
	return output
}

// WORKER POOL PATTERN: Fixed number of workers processing jobs
func workerPoolPattern() {
	fmt.Println("Creating a worker pool...")
	
	// Configuration
	numWorkers := 3
	numJobs := 10
	
	// Create channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	
	// Send jobs
	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("  📊 Worker pool results:")
	for result := range results {
		fmt.Printf("    Result: %d\n", result)
	}
}

// Worker function for the worker pool
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("  🧵 Worker %d processing job %d\n", id, job)
		time.Sleep(200 * time.Millisecond) // Simulate work
		results <- job * job
	}
}

// GRACEFUL SHUTDOWN PATTERN: Properly shutting down goroutines
func gracefulShutdownPattern() {
	fmt.Println("Implementing graceful shutdown...")
	
	// Create shutdown signal
	shutdown := make(chan struct{})
	
	// Start a service
	service := startService(shutdown)
	
	// Let it run for a bit
	time.Sleep(3 * time.Second)
	
	// Gracefully shutdown
	fmt.Println("  🛑 Initiating graceful shutdown...")
	close(shutdown)
	
	// Wait for service to stop
	<-service
	fmt.Println("  ✅ Service stopped gracefully")
}

// Service that can be gracefully shut down
func startService(shutdown <-chan struct{}) <-chan struct{} {
	stopped := make(chan struct{})
	
	go func() {
		defer close(stopped)
		
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Println("  🧵 Service working...")
			case <-shutdown:
				fmt.Println("  🧵 Service shutting down...")
				time.Sleep(500 * time.Millisecond) // Cleanup time
				return
			}
		}
	}()
	
	return stopped
}

// DYNAMIC SCALING POOL: Pool that adjusts size based on load
func dynamicScalingPool() {
	fmt.Println("Creating a dynamically scaling pool...")
	
	// Configuration
	minWorkers := 2
	maxWorkers := 5
	jobQueue := make(chan int, 20)
	results := make(chan int, 20)
	
	// Start with minimum workers
	activeWorkers := minWorkers
	var wg sync.WaitGroup
	
	// Start initial workers
	for i := 0; i < activeWorkers; i++ {
		wg.Add(1)
		go dynamicWorker(i, jobQueue, results, &wg)
	}
	
	// Load balancer goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		
		for range ticker.C {
			queueLen := len(jobQueue)
			currentWorkers := runtime.NumGoroutine() - 2 // Subtract main and this goroutine
			
			fmt.Printf("  📊 Queue length: %d, Active workers: %d\n", queueLen, currentWorkers)
			
			// Scale up if queue is full and we can add more workers
			if queueLen > 10 && currentWorkers < maxWorkers {
				fmt.Printf("  ⬆️  Scaling up to %d workers\n", currentWorkers+1)
				wg.Add(1)
				go dynamicWorker(currentWorkers, jobQueue, results, &wg)
			}
		}
	}()
	
	// Send jobs
	go func() {
		defer close(jobQueue)
		for i := 1; i <= 15; i++ {
			jobQueue <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("  📊 Dynamic pool results:")
	for result := range results {
		fmt.Printf("    Result: %d\n", result)
	}
}

// Dynamic worker that can be added/removed
func dynamicWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("  🧵 Dynamic Worker %d processing job %d\n", id, job)
		time.Sleep(300 * time.Millisecond) // Simulate work
		results <- job * 3
	}
}
