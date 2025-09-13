package main

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 1: Basic Channel Operations
// Create a program that sends and receives values through a channel.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Channel Operations")
	fmt.Println("====================================")
	
	ch := make(chan string)
	
	// Send values
	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "from"
		ch <- "Go"
		close(ch)
	}()
	
	// Receive values
	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
	}
}

// Exercise 2: Buffered vs Unbuffered
// Compare the behavior of buffered and unbuffered channels.
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered vs Unbuffered")
	fmt.Println("===================================")
	
	// Unbuffered channel
	unbuffered := make(chan int)
	
	// Buffered channel
	buffered := make(chan int, 3)
	
	// Test unbuffered
	fmt.Println("Unbuffered channel test:")
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Sending %d to unbuffered channel\n", i)
			unbuffered <- i
		}
		close(unbuffered)
	}()
	
	for value := range unbuffered {
		fmt.Printf("Received %d from unbuffered channel\n", value)
	}
	
	// Test buffered
	fmt.Println("\nBuffered channel test:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sending %d to buffered channel\n", i)
		buffered <- i
	}
	close(buffered)
	
	for value := range buffered {
		fmt.Printf("Received %d from buffered channel\n", value)
	}
}

// Exercise 3: Channel Direction
// Create functions that use send-only and receive-only channels.
func Exercise3() {
	fmt.Println("\nExercise 3: Channel Direction")
	fmt.Println("=============================")
	
	ch := make(chan int)
	
	// Send-only function
	go sendOnly(ch)
	
	// Receive-only function
	receiveOnly(ch)
}

func sendOnly(ch chan<- int) {
	fmt.Println("Send-only function: Sending values...")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func receiveOnly(ch <-chan int) {
	fmt.Println("Receive-only function: Receiving values...")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Channel closed")
}

// Exercise 4: Channel Closing
// Implement proper channel closing and cleanup.
func Exercise4() {
	fmt.Println("\nExercise 4: Channel Closing")
	fmt.Println("===========================")
	
	ch := make(chan int)
	
	// Producer
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Producer: Sent %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Producer: Finished sending")
	}()
	
	// Consumer
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Consumer: Channel closed, exiting")
			break
		}
		fmt.Printf("Consumer: Received %d\n", value)
	}
}

// Exercise 5: Select Statement
// Use select to handle multiple channels.
func Exercise5() {
	fmt.Println("\nExercise 5: Select Statement")
	fmt.Println("============================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	// Start goroutines
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch3 <- "Message from channel 3"
	}()
	
	// Use select to receive from any channel
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case msg3 := <-ch3:
			fmt.Printf("Received from ch3: %s\n", msg3)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

// Exercise 6: Pipeline Pattern
// Create a pipeline that processes data through multiple stages.
func Exercise6() {
	fmt.Println("\nExercise 6: Pipeline Pattern")
	fmt.Println("============================")
	
	// Create pipeline stages
	numbers := make(chan int)
	squares := make(chan int)
	results := make(chan int)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("Stage 1: Generated %d\n", i)
		}
	}()
	
	// Stage 2: Square numbers
	go func() {
		defer close(squares)
		for n := range numbers {
			squared := n * n
			squares <- squared
			fmt.Printf("Stage 2: Squared %d = %d\n", n, squared)
		}
	}()
	
	// Stage 3: Add 10 to squares
	go func() {
		defer close(results)
		for s := range squares {
			result := s + 10
			results <- result
			fmt.Printf("Stage 3: Added 10 to %d = %d\n", s, result)
		}
	}()
	
	// Collect results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
}

// Exercise 7: Fan-Out/Fan-In
// Implement fan-out and fan-in patterns.
func Exercise7() {
	fmt.Println("\nExercise 7: Fan-Out/Fan-In")
	fmt.Println("==========================")
	
	// Create input channel
	input := make(chan int)
	
	// Create output channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// Fan-out goroutine
	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()
		
		for n := range input {
			ch1 <- n
			ch2 <- n
			ch3 <- n
		}
	}()
	
	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()
	
	// Fan-in goroutine
	merged := fanInExercise(ch1, ch2, ch3)
	
	// Collect results
	fmt.Println("Fan-out/Fan-in results:")
	for value := range merged {
		fmt.Printf("Received: %d\n", value)
	}
}

func fanInExercise(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		var wg sync.WaitGroup
		
		for _, input := range inputs {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for value := range ch {
					output <- value
				}
			}(input)
		}
		
		wg.Wait()
	}()
	
	return output
}

// Exercise 8: Channel Timeout
// Add timeout handling to channel operations.
func Exercise8() {
	fmt.Println("\nExercise 8: Channel Timeout")
	fmt.Println("===========================")
	
	ch := make(chan string)
	
	// Send a message after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "Delayed message"
	}()
	
	// Wait for message with timeout
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
}

// Exercise 9: Channel with Quit Signal
func Exercise9() {
	fmt.Println("\nExercise 9: Channel with Quit Signal")
	fmt.Println("====================================")
	
	workCh := make(chan int)
	quitCh := make(chan bool)
	
	// Worker goroutine
	go func() {
		for {
			select {
			case work := <-workCh:
				fmt.Printf("Processing work: %d\n", work)
				time.Sleep(100 * time.Millisecond)
			case <-quitCh:
				fmt.Println("Worker: Received quit signal, exiting")
				return
			}
		}
	}()
	
	// Send some work
	for i := 1; i <= 3; i++ {
		workCh <- i
	}
	
	// Send quit signal
	quitCh <- true
	time.Sleep(100 * time.Millisecond)
}

// Exercise 10: Channel Pool
func Exercise10() {
	fmt.Println("\nExercise 10: Channel Pool")
	fmt.Println("=========================")
	
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
		fmt.Printf("Job %d completed with result: %d\n", r, result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * job
	}
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Channel Exercises")
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
