package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Concurrency Learning - Project 3 ===")
	fmt.Println("Common Concurrency Patterns")
	fmt.Println()

	// Exercise 1: Fan-In/Fan-Out Pattern
	fmt.Println("Exercise 1: Fan-In/Fan-Out Pattern")
	demonstrateFanInFanOut()
	fmt.Println()

	// Exercise 2: Pipeline Pattern
	fmt.Println("Exercise 2: Pipeline Pattern")
	demonstratePipeline()
	fmt.Println()

	// Exercise 3: Generator Pattern
	fmt.Println("Exercise 3: Generator Pattern")
	demonstrateGenerator()
	fmt.Println()

	// Exercise 4: Worker Pool Pattern
	fmt.Println("Exercise 4: Worker Pool Pattern")
	demonstrateWorkerPool()
	fmt.Println()

	// Exercise 5: Graceful Shutdown
	fmt.Println("Exercise 5: Graceful Shutdown")
	demonstrateGracefulShutdown()
	fmt.Println()

	fmt.Println("=== All pattern exercises completed! ===")
	fmt.Println()
	fmt.Println("Run specific components:")
	fmt.Println("  go run main.go fan_in_out.go")
	fmt.Println("  go run main.go pipeline.go")
	fmt.Println("  go run main.go generator.go")
	fmt.Println("  go run main.go worker_pool.go")
	fmt.Println("  go run main.go graceful_shutdown.go")
}

func demonstrateFanInFanOut() {
	// Fan-out: Distribute work to multiple workers
	input := make(chan int)
	output := make(chan int)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for n := range input {
				result := n * n // Square the number
				fmt.Printf("Worker %d: processing %d -> %d\n", workerID, n, result)
				output <- result
			}
		}(i)
	}
	
	// Close output when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Send work
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Collect results (fan-in)
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("  Received: %d\n", result)
	}
}

func demonstratePipeline() {
	// Stage 1: Generate numbers
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 5; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers
	stage2 := make(chan int)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			result := n * n
			fmt.Printf("Stage 2: squaring %d -> %d\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Add 10
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
}

func demonstrateGenerator() {
	// Simple generator
	numbers := generateNumbers(1, 5)
	
	fmt.Println("Generated numbers:")
	for n := range numbers {
		fmt.Printf("  %d\n", n)
	}
	
	// Fibonacci generator
	fib := generateFibonacci(10)
	
	fmt.Println("Fibonacci sequence:")
	for n := range fib {
		fmt.Printf("  %d\n", n)
	}
}

func demonstrateWorkerPool() {
	// Create job queue
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				result := job * job
				fmt.Printf("Worker %d: processing job %d -> %d\n", workerID, job, result)
				results <- result
			}
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
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
	fmt.Println("Worker pool results:")
	for result := range results {
		fmt.Printf("  Result: %d\n", result)
	}
}

func demonstrateGracefulShutdown() {
	// Create a system that can be shut down gracefully
	system := NewGracefulSystem()
	
	// Start the system
	system.Start()
	
	// Let it run for a bit
	time.Sleep(2 * time.Second)
	
	// Shutdown gracefully
	fmt.Println("Initiating graceful shutdown...")
	system.Shutdown()
	
	// Wait for shutdown to complete
	<-system.Done()
	fmt.Println("System shutdown complete")
}

// Helper functions

func generateNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

func generateFibonacci(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		for i := 0; i < count; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

// GracefulSystem represents a system that can be shut down gracefully
type GracefulSystem struct {
	workers    int
	jobs       chan string
	results    chan string
	shutdown   chan struct{}
	done       chan struct{}
	wg         sync.WaitGroup
}

// NewGracefulSystem creates a new graceful system
func NewGracefulSystem() *GracefulSystem {
	return &GracefulSystem{
		workers:  3,
		jobs:    make(chan string, 100),
		results: make(chan string, 100),
		shutdown: make(chan struct{}),
		done:    make(chan struct{}),
	}
}

// Start starts the system
func (gs *GracefulSystem) Start() {
	// Start workers
	for i := 0; i < gs.workers; i++ {
		gs.wg.Add(1)
		go gs.worker(i)
	}
	
	// Start job producer
	gs.wg.Add(1)
	go gs.producer()
	
	// Start result collector
	gs.wg.Add(1)
	go gs.collector()
}

// worker processes jobs
func (gs *GracefulSystem) worker(id int) {
	defer gs.wg.Done()
	
	for {
		select {
		case job := <-gs.jobs:
			// Process job
			time.Sleep(500 * time.Millisecond)
			result := fmt.Sprintf("Worker %d processed: %s", id, job)
			gs.results <- result
		case <-gs.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// producer generates jobs
func (gs *GracefulSystem) producer() {
	defer gs.wg.Done()
	
	jobCount := 0
	for {
		select {
		case gs.jobs <- fmt.Sprintf("job_%d", jobCount):
			jobCount++
			time.Sleep(200 * time.Millisecond)
		case <-gs.shutdown:
			fmt.Println("Producer shutting down")
			return
		}
	}
}

// collector collects results
func (gs *GracefulSystem) collector() {
	defer gs.wg.Done()
	
	for {
		select {
		case result := <-gs.results:
			fmt.Printf("Result: %s\n", result)
		case <-gs.shutdown:
			fmt.Println("Collector shutting down")
			return
		}
	}
}

// Shutdown initiates graceful shutdown
func (gs *GracefulSystem) Shutdown() {
	close(gs.shutdown)
	
	// Wait for all workers to finish
	go func() {
		gs.wg.Wait()
		close(gs.done)
	}()
}

// Done returns a channel that signals when shutdown is complete
func (gs *GracefulSystem) Done() <-chan struct{} {
	return gs.done
}
