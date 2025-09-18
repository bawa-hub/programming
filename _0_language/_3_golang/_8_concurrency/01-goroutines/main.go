package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// This file contains the basic goroutine examples
// Use cmd.go to run the examples with command line arguments

// Example 1: Basic Goroutine
func basicGoroutine() {
	fmt.Println("Starting basic goroutine...")
	
	go func() {
		fmt.Println("Hello from goroutine!")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Goroutine finished!")
	}()

	// Give goroutine time to run
	time.Sleep(200 * time.Millisecond)
}

// Example 2: Multiple Goroutines
func multipleGoroutines() {
	fmt.Println("Starting 5 goroutines...")
	
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Starting\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Goroutine %d: Finished\n", id)
		}(i)
	}

	// Wait for all goroutines
	time.Sleep(1 * time.Second)
}

// Example 3: Goroutine with WaitGroup
func goroutineWithWaitGroup() {
	var wg sync.WaitGroup
	
	fmt.Println("Starting goroutines with WaitGroup...")
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d: Starting\n", id)
			time.Sleep(time.Duration(id*200) * time.Millisecond)
			fmt.Printf("Worker %d: Finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed!")
}

// Example 4: Goroutine Pool
func goroutinePool() {
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 2
	}
}

// Example 5: Goroutine Communication
func goroutineCommunication() {
	ch := make(chan string)
	
	// Sender goroutine
	go func() {
		ch <- "Hello"
		ch <- "from"
		ch <- "goroutine"
		close(ch)
	}()
	
	// Receiver (main goroutine)
	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
	}
}

// Example 6: Goroutine Lifecycle
func goroutineLifecycle() {
	fmt.Println("Starting goroutine lifecycle demo...")
	
	// Start a goroutine that can be controlled
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Goroutine: Received stop signal")
				return
			default:
				fmt.Println("Goroutine: Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	
	// Let it run for a bit
	time.Sleep(2 * time.Second)
	
	// Stop the goroutine
	done <- true
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Goroutine lifecycle demo completed")
}

// Example 7: Performance Comparison
func performanceComparison() {
	fmt.Println("Comparing goroutine vs function call performance...")
	
	// Function calls
	start := time.Now()
	for i := 0; i < 1000; i++ {
		func() {
			// Simulate work
			_ = i * i
		}()
	}
	funcTime := time.Since(start)
	
	// Goroutines
	start = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate work
			_ = i * i
		}()
	}
	wg.Wait()
	goroutineTime := time.Since(start)
	
	fmt.Printf("Function calls: %v\n", funcTime)
	fmt.Printf("Goroutines: %v\n", goroutineTime)
	fmt.Printf("Overhead: %.2fx\n", float64(goroutineTime)/float64(funcTime))
}

// Example 8: Common Pitfalls
func commonPitfalls() {
	fmt.Println("Demonstrating common pitfalls...")
	
	// Pitfall 1: Variable capture in loops
	fmt.Println("\nPitfall 1: Variable capture in loops")
	fmt.Println("Wrong way:")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("Wrong: i = %d\n", i) // Always prints 3
		}()
	}
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("Correct way:")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Correct: i = %d\n", i) // Prints 0, 1, 2
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	
	// Pitfall 2: Goroutine leak (commented out to avoid actual leak)
	fmt.Println("\nPitfall 2: Goroutine leak (demonstration)")
	fmt.Println("// This would create a leak:")
	fmt.Println("// go func() { for { } }()")
	fmt.Println("// Use context or channels to prevent leaks")
}

// Utility function to show goroutine info
func showGoroutineInfo() {
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
