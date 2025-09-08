package main

import (
	"fmt"
	"time"
)

// Advanced channel patterns and techniques
func channelPatterns() {
	fmt.Println("=== Advanced Channel Patterns ===")

	// Pattern 1: Signal channels
	fmt.Println("1. Signal Channels:")
	signalChannelExample()

	// Pattern 2: Done channels
	fmt.Println("\n2. Done Channels:")
	doneChannelExample()

	// Pattern 3: Rate limiting
	fmt.Println("\n3. Rate Limiting:")
	rateLimitingExample()

	// Pattern 4: Channel multiplexing
	fmt.Println("\n4. Channel Multiplexing:")
	multiplexingExample()

	// Pattern 5: Channel closing patterns
	fmt.Println("\n5. Channel Closing Patterns:")
	channelClosingExample()
}

// Signal channels: Use channels to signal events
func signalChannelExample() {
	// Create a signal channel
	ready := make(chan bool)

	// Worker that signals when ready
	go func() {
		fmt.Println("Worker: Starting work...")
		time.Sleep(1 * time.Second)
		fmt.Println("Worker: Work completed!")
		ready <- true // Signal completion
	}()

	// Main goroutine waits for signal
	fmt.Println("Main: Waiting for worker to complete...")
	<-ready
	fmt.Println("Main: Worker is ready!")
}

// Done channels: Use channels to signal cancellation
func doneChannelExample() {
	done := make(chan bool)

	// Worker that can be cancelled
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				fmt.Println("Worker: Cancelled!")
				return
			default:
				fmt.Printf("Worker: Processing item %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
		fmt.Println("Worker: Completed all items")
	}()

	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Cancelling worker...")
	close(done) // Signal cancellation
	time.Sleep(100 * time.Millisecond)
}

// Rate limiting: Control the rate of operations
func rateLimitingExample() {
	// Create a rate limiter (ticker)
	rateLimiter := time.Tick(200 * time.Millisecond)

	// Process items with rate limiting
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	for _, item := range items {
		<-rateLimiter // Wait for rate limiter
		fmt.Printf("Processing: %s\n", item)
	}
}

// Channel multiplexing: Combine multiple channels
func multiplexingExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Start goroutines that send to different channels
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	go func() {
		time.Sleep(400 * time.Millisecond)
		ch3 <- "Message from channel 3"
	}()

	// Multiplex: listen to all channels
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg)
		case msg := <-ch3:
			fmt.Printf("Received from ch3: %s\n", msg)
		}
	}
}

// Channel closing patterns
func channelClosingExample() {
	// Pattern 1: Close from sender
	fmt.Println("Pattern 1: Close from sender")
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // Sender closes
	}()

	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}

	// Pattern 2: Close from receiver (using done channel)
	fmt.Println("\nPattern 2: Close from receiver")
	ch2 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 5; i++ {
			select {
			case ch2 <- i:
				fmt.Printf("Sent: %d\n", i)
			case <-done:
				fmt.Println("Sender: Stopping...")
				close(ch2)
				return
			}
		}
		close(ch2)
	}()

	// Receive a few items, then signal done
	for i := 0; i < 3; i++ {
		if num, ok := <-ch2; ok {
			fmt.Printf("Received: %d\n", num)
		}
	}
	done <- true
	time.Sleep(100 * time.Millisecond)
}

// Advanced: Channel with timeout and retry
func advancedChannelPatterns() {
	fmt.Println("\n=== Advanced Channel Patterns ===")

	// Channel with timeout
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Delayed message"
	}()

	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received")
	}

	// Channel with retry mechanism
	retryChannelExample()
}

func retryChannelExample() {
	fmt.Println("\nRetry Channel Example:")
	
	maxRetries := 3
	ch := make(chan string)
	
	// Simulate unreliable sender
	go func() {
		for i := 0; i < maxRetries; i++ {
			time.Sleep(500 * time.Millisecond)
			if i == 1 { // Simulate success on second try
				ch <- "Success message"
				return
			}
		}
		close(ch)
	}()

	// Retry logic
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case msg, ok := <-ch:
			if ok {
				fmt.Printf("Success on attempt %d: %s\n", attempt, msg)
				return
			} else {
				fmt.Printf("Attempt %d: Channel closed\n", attempt)
				return
			}
		case <-time.After(1 * time.Second):
			fmt.Printf("Attempt %d: Timeout\n", attempt)
		}
	}
	fmt.Println("All retry attempts failed")
}

// Channel composition patterns
func channelCompositionPatterns() {
	fmt.Println("\n=== Channel Composition Patterns ===")

	// Pattern: Chain of channels
	input := make(chan int)
	output := make(chan int)

	// Stage 1: Double the input
	go func() {
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	// Send input
	go func() {
		for i := 1; i <= 3; i++ {
			input <- i
		}
		close(input)
	}()

	// Collect output
	fmt.Println("Chain pattern results:")
	for result := range output {
		fmt.Printf("Result: %d\n", result)
	}
}

func main() {
	channelPatterns()
	advancedChannelPatterns()
	channelCompositionPatterns()
}
