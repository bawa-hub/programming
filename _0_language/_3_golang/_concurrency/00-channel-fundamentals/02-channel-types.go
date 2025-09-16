package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. UNBUFFERED CHANNELS
// ============================================================================

func unbufferedChannels() {

	// Unbuffered channels have a capacity of 0
	// They provide synchronous communication
	// Sender blocks until receiver is ready
	// Receiver blocks until sender is ready
	// Use case: When you need guaranteed synchronization
	
	// Create an unbuffered channel
	ch := make(chan int)  // No capacity specified = unbuffered
	fmt.Printf("Unbuffered channel: %v\n", ch)
		
	// This will demonstrate the blocking behavior
	go func() {
		fmt.Println("  Sender: About to send data...")
		ch <- 42  // This blocks until someone receives
		fmt.Println("  Sender: Data sent! (This prints after receiver gets data)")
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)  // Wait a bit
		fmt.Println("  Receiver: About to receive data...")
		data := <-ch  // This blocks until someone sends
		fmt.Printf("  Receiver: Got data: %d\n", data)
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 2. BUFFERED CHANNELS
// ============================================================================

func bufferedChannels() {

	// Buffered channels have a capacity > 0
	// They provide asynchronous communication
	// Sender only blocks when buffer is full
	// Receiver only blocks when buffer is empty
	// Use case: When you want to decouple sender and receiver
	
	// Create a buffered channel with capacity 3
	ch := make(chan int, 3)  // Capacity of 3
	fmt.Printf("Buffered channel (capacity 3): %v\n", ch)
		
	// Send multiple values without blocking (up to capacity)
	fmt.Println("  Sending 3 values to buffered channel...")
	ch <- 1  // Doesn't block (buffer has space)
	ch <- 2  // Doesn't block (buffer has space)
	ch <- 3  // Doesn't block (buffer has space)
	fmt.Println("  All 3 values sent without blocking!")
	
	// This would block because buffer is full
	go func() {
		fmt.Println("  Sender: Trying to send 4th value...")
		ch <- 4  // This will block because buffer is full
		fmt.Println("  Sender: 4th value sent! (This prints after receiver gets data)")
	}()
	
	// Receive values
	fmt.Println("  Receiving values...")
	for i := 0; i < 4; i++ {
		data := <-ch
		fmt.Printf("  Received: %d\n", data)
		time.Sleep(200 * time.Millisecond)  // Small delay to see the behavior
	}
}

// ============================================================================
// 3. CHANNEL CAPACITY COMPARISON
// ============================================================================

func channelCapacityComparison() {
	
	// Unbuffered channel
	unbuffered := make(chan int)
	fmt.Printf("Unbuffered channel capacity: %d\n", cap(unbuffered))
	
	// Buffered channels with different capacities
	buffered1 := make(chan int, 1)
	buffered5 := make(chan int, 5)
	buffered10 := make(chan int, 10)
	
	fmt.Printf("Buffered channel (cap 1): %d\n", cap(buffered1))
	fmt.Printf("Buffered channel (cap 5): %d\n", cap(buffered5))
	fmt.Printf("Buffered channel (cap 10): %d\n", cap(buffered10))
}

// ============================================================================
// 4. CHANNEL LENGTH AND CAPACITY
// ============================================================================

func channelLengthAndCapacity() {
	// len(ch) - number of elements currently in the channel
	// cap(ch) - maximum number of elements the channel can hold
	
	ch := make(chan int, 5)  // Capacity of 5

	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Send some data
	ch <- 1
	ch <- 2
	ch <- 3
	
	fmt.Println("\nAfter sending 3 values")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Send more data
	ch <- 4
	ch <- 5
	
	fmt.Println("\n After sending 5 values (buffer full)")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Receive some data
	<-ch
	<-ch
	
	fmt.Println("\n After receiving 2 values")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Clear the channel
	for len(ch) > 0 {
		<-ch
	}
	
	fmt.Println("\n After clearing channel")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
}

// ============================================================================
// 5. NIL CHANNELS
// ============================================================================

func nilChannels() {
	// A nil channel is a channel that hasn't been initialized
	// Operations on nil channels block forever
	// Closing nil channel panics
	
	var ch chan int  // This is nil
	fmt.Printf("Nil channel: %v\n", ch)
	fmt.Printf("Is nil: %t\n", ch == nil)

	// This would block forever (commented out to avoid hanging)
	// ch <- 42  // This would block forever
	// data := <-ch  // This would block forever
}

// ============================================================================
// 6. CHANNEL TYPES DEMONSTRATION
// ============================================================================

func channelTypesDemonstration() {
	fmt.Println("\n Different channel types")
	
	// Different ways to create channels
	var ch1 chan int                    // Nil channel
	ch2 := make(chan int)               // Unbuffered channel
	ch3 := make(chan int, 0)            // Unbuffered channel (explicit)
	ch4 := make(chan int, 1)            // Buffered channel (capacity 1)
	ch5 := make(chan int, 10)           // Buffered channel (capacity 10)
	
	fmt.Printf("Nil channel: %v (nil: %t)\n", ch1, ch1 == nil)
	fmt.Printf("Unbuffered: %v (cap: %d)\n", ch2, cap(ch2))
	fmt.Printf("Unbuffered explicit: %v (cap: %d)\n", ch3, cap(ch3))
	fmt.Printf("Buffered (1): %v (cap: %d)\n", ch4, cap(ch4))
	fmt.Printf("Buffered (10): %v (cap: %d)\n", ch5, cap(ch5))
	
	fmt.Println("\n Type checking")
	fmt.Printf("ch2 is unbuffered: %t\n", cap(ch2) == 0)
	fmt.Printf("ch4 is buffered: %t\n", cap(ch4) > 0)
	fmt.Printf("ch5 is buffered: %t\n", cap(ch5) > 0)
}
