package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Exercise 1: Basic Goroutines
// Create a program that starts 5 goroutines, each printing a unique number.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Goroutines")
	fmt.Println("============================")
	
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Hello!\n", id)
		}(i)
	}
	
	// Wait for goroutines to complete
	time.Sleep(100 * time.Millisecond)
}

// Exercise 2: Goroutine Synchronization
// Use WaitGroup to wait for 3 goroutines to complete.
func Exercise2() {
	fmt.Println("\nExercise 2: Goroutine Synchronization")
	fmt.Println("=====================================")
	
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d: Starting work\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d: Work completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All workers completed!")
}

// Exercise 3: Goroutine Pool
// Implement a worker pool with 3 workers processing 10 jobs.
func Exercise3() {
	fmt.Println("\nExercise 3: Goroutine Pool")
	fmt.Println("==========================")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go exerciseWorker(i, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job %d processed, result: %d\n", r, result)
	}
}

func exerciseWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * job // Square the job number
	}
}

// Exercise 4: Goroutine Communication
// Create 2 goroutines that communicate through a channel.
func Exercise4() {
	fmt.Println("\nExercise 4: Goroutine Communication")
	fmt.Println("===================================")
	
	ch := make(chan string)
	
	// Producer goroutine
	go func() {
		messages := []string{"Hello", "from", "producer", "goroutine"}
		for _, msg := range messages {
			ch <- msg
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()
	
	// Consumer goroutine
	go func() {
		for msg := range ch {
			fmt.Printf("Consumer received: %s\n", msg)
		}
		fmt.Println("Consumer: Channel closed, exiting")
	}()
	
	// Wait for communication to complete
	time.Sleep(1 * time.Second)
}

// Exercise 5: Goroutine Lifecycle
// Implement a goroutine that can be started, paused, and stopped.
func Exercise5() {
	fmt.Println("\nExercise 5: Goroutine Lifecycle")
	fmt.Println("===============================")
	
	start := make(chan bool)
	pause := make(chan bool)
	stop := make(chan bool)
	
	// Controllable goroutine
	go func() {
		running := false
		for {
			select {
			case <-start:
				running = true
				fmt.Println("Goroutine: Started")
			case <-pause:
				if running {
					running = false
					fmt.Println("Goroutine: Paused")
				}
			case <-stop:
				fmt.Println("Goroutine: Stopped")
				return
			default:
				if running {
					fmt.Println("Goroutine: Working...")
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()
	
	// Control the goroutine
	start <- true
	time.Sleep(500 * time.Millisecond)
	
	pause <- true
	time.Sleep(500 * time.Millisecond)
	
	start <- true
	time.Sleep(500 * time.Millisecond)
	
	stop <- true
	time.Sleep(100 * time.Millisecond)
}

// Exercise 6: Advanced - Goroutine Pool with Rate Limiting
func Exercise6() {
	fmt.Println("\nExercise 6: Advanced - Rate Limited Pool")
	fmt.Println("=======================================")
	
	const numWorkers = 2
	const numJobs = 8
	const rateLimit = 2 // jobs per second
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	rateLimiter := time.Tick(time.Second / rateLimit)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go rateLimitedWorker(i, jobs, results, rateLimiter)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job %d completed with result: %d\n", r, result)
	}
}

func rateLimitedWorker(id int, jobs <-chan int, results chan<- int, rateLimiter <-chan time.Time) {
	for job := range jobs {
		<-rateLimiter // Wait for rate limit
		fmt.Printf("Worker %d: Processing job %d (rate limited)\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 3
	}
}

// Exercise 7: Goroutine Monitoring
func Exercise7() {
	fmt.Println("\nExercise 7: Goroutine Monitoring")
	fmt.Println("===============================")
	
	// Start some goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				fmt.Printf("Goroutine %d: iteration %d\n", id, j)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}
	
	// Monitor goroutines
	for i := 0; i < 3; i++ {
		fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(300 * time.Millisecond)
	}
}

// Exercise 8: Goroutine with Error Handling
func Exercise8() {
	fmt.Println("\nExercise 8: Goroutine with Error Handling")
	fmt.Println("=========================================")
	
	var wg sync.WaitGroup
	errors := make(chan error, 3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					errors <- fmt.Errorf("goroutine %d panicked: %v", id, r)
				}
			}()
			
			if id == 1 {
				panic("Simulated panic in goroutine 1")
			}
			
			fmt.Printf("Goroutine %d: Working normally\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	close(errors)
	
	// Check for errors
	for err := range errors {
		fmt.Printf("Error: %v\n", err)
	}
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Goroutine Exercises")
	fmt.Println("==================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	
	fmt.Println("\n✅ All exercises completed!")
}
