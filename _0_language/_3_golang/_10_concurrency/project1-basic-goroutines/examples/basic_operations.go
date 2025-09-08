package main

import (
	"fmt"
	"time"
)

// Basic goroutine examples
func basicGoroutineExamples() {
	fmt.Println("=== Basic Goroutine Examples ===")

	// Example 1: Simple goroutine
	fmt.Println("1. Simple goroutine:")
	go func() {
		fmt.Println("Hello from goroutine!")
	}()
	time.Sleep(100 * time.Millisecond)

	// Example 2: Multiple goroutines
	fmt.Println("\n2. Multiple goroutines:")
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}
	time.Sleep(200 * time.Millisecond)

	// Example 3: Goroutine with return value (using channel)
	fmt.Println("\n3. Goroutine with return value:")
	result := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		result <- 42
	}()
	value := <-result
	fmt.Printf("Received value: %d\n", value)
}

// Channel examples
func channelExamples() {
	fmt.Println("\n=== Channel Examples ===")

	// Example 1: Unbuffered channel
	fmt.Println("1. Unbuffered channel:")
	ch := make(chan string)
	go func() {
		ch <- "Hello from channel!"
	}()
	msg := <-ch
	fmt.Printf("Received: %s\n", msg)

	// Example 2: Buffered channel
	fmt.Println("\n2. Buffered channel:")
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	fmt.Printf("Channel length: %d\n", len(bufferedCh))
	
	for i := 0; i < 3; i++ {
		fmt.Printf("Received: %d\n", <-bufferedCh)
	}

	// Example 3: Channel direction
	fmt.Println("\n3. Channel direction (send-only, receive-only):")
	sendOnly := make(chan<- int)
	receiveOnly := make(<-chan int)
	
	// These would be used in function parameters
	_ = sendOnly
	_ = receiveOnly
	fmt.Println("Channel directions defined (used in function parameters)")
}

// Select statement examples
func selectExamples() {
	fmt.Println("\n=== Select Statement Examples ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start goroutines
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()

	// Example 1: Basic select
	fmt.Println("1. Basic select (first available):")
	select {
	case msg := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg)
	}

	// Example 2: Select with timeout
	fmt.Println("\n2. Select with timeout:")
	timeout := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s\n", msg)
	case <-timeout:
		fmt.Println("Timeout occurred")
	}

	// Example 3: Select with default
	fmt.Println("\n3. Select with default:")
	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available, using default")
	}
}

// Producer-Consumer pattern
func producerConsumerExample() {
	fmt.Println("\n=== Producer-Consumer Pattern ===")

	// Create channels
	numbers := make(chan int, 5)
	done := make(chan bool)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Producing: %d\n", i)
			numbers <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(numbers)
		done <- true
	}()

	// Consumer
	go func() {
		for num := range numbers {
			fmt.Printf("Consuming: %d\n", num)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Wait for producer to finish
	<-done
	fmt.Println("Producer finished")
}

func main() {
	basicGoroutineExamples()
	channelExamples()
	selectExamples()
	producerConsumerExample()
}
