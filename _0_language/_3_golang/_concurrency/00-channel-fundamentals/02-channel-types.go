package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. UNBUFFERED CHANNELS
// ============================================================================

func unbufferedChannels() {
	fmt.Println("\n📤 UNBUFFERED CHANNELS")
	fmt.Println("======================")

	// Unbuffered channels have a capacity of 0
	// They provide synchronous communication
	// Sender blocks until receiver is ready
	// Receiver blocks until sender is ready
	
	fmt.Println("\n1.1 What are unbuffered channels?")
	fmt.Println("Capacity: 0")
	fmt.Println("Behavior: Synchronous (blocking)")
	fmt.Println("Use case: When you need guaranteed synchronization")
	
	// Create an unbuffered channel
	ch := make(chan int)  // No capacity specified = unbuffered
	fmt.Printf("Unbuffered channel: %v\n", ch)
	
	fmt.Println("\n1.2 Unbuffered channel behavior")
	
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
	fmt.Println("\n📦 BUFFERED CHANNELS")
	fmt.Println("===================")

	// Buffered channels have a capacity > 0
	// They provide asynchronous communication
	// Sender only blocks when buffer is full
	// Receiver only blocks when buffer is empty
	
	fmt.Println("\n2.1 What are buffered channels?")
	fmt.Println("Capacity: > 0 (specified when creating)")
	fmt.Println("Behavior: Asynchronous (non-blocking until buffer full/empty)")
	fmt.Println("Use case: When you want to decouple sender and receiver")
	
	// Create a buffered channel with capacity 3
	ch := make(chan int, 3)  // Capacity of 3
	fmt.Printf("Buffered channel (capacity 3): %v\n", ch)
	
	fmt.Println("\n2.2 Buffered channel behavior")
	
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
	fmt.Println("\n⚖️  CHANNEL CAPACITY COMPARISON")
	fmt.Println("=============================")

	fmt.Println("\n3.1 Unbuffered vs Buffered")
	
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
	
	fmt.Println("\n3.2 When to use which?")
	fmt.Println("Unbuffered: When you need tight synchronization")
	fmt.Println("Buffered: When you want to decouple producer and consumer")
	fmt.Println("Large buffer: When you have bursty traffic")
}

// ============================================================================
// 4. CHANNEL LENGTH AND CAPACITY
// ============================================================================

func channelLengthAndCapacity() {
	fmt.Println("\n📏 CHANNEL LENGTH AND CAPACITY")
	fmt.Println("=============================")

	// len(ch) - number of elements currently in the channel
	// cap(ch) - maximum number of elements the channel can hold
	
	ch := make(chan int, 5)  // Capacity of 5
	
	fmt.Println("\n4.1 Initial state")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Send some data
	ch <- 1
	ch <- 2
	ch <- 3
	
	fmt.Println("\n4.2 After sending 3 values")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Send more data
	ch <- 4
	ch <- 5
	
	fmt.Println("\n4.3 After sending 5 values (buffer full)")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Receive some data
	<-ch
	<-ch
	
	fmt.Println("\n4.4 After receiving 2 values")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Clear the channel
	for len(ch) > 0 {
		<-ch
	}
	
	fmt.Println("\n4.5 After clearing channel")
	fmt.Printf("Length: %d, Capacity: %d\n", len(ch), cap(ch))
}

// ============================================================================
// 5. NIL CHANNELS
// ============================================================================

func nilChannels() {
	fmt.Println("\n🚫 NIL CHANNELS")
	fmt.Println("===============")

	// A nil channel is a channel that hasn't been initialized
	// Operations on nil channels block forever
	
	fmt.Println("\n5.1 What is a nil channel?")
	var ch chan int  // This is nil
	fmt.Printf("Nil channel: %v\n", ch)
	fmt.Printf("Is nil: %t\n", ch == nil)
	
	fmt.Println("\n5.2 Nil channel behavior")
	fmt.Println("⚠️  Sending to nil channel blocks forever")
	fmt.Println("⚠️  Receiving from nil channel blocks forever")
	fmt.Println("⚠️  Closing nil channel panics")
	
	// This would block forever (commented out to avoid hanging)
	// ch <- 42  // This would block forever
	// data := <-ch  // This would block forever
}

// ============================================================================
// 6. CHANNEL TYPES DEMONSTRATION
// ============================================================================

func channelTypesDemonstration() {
	fmt.Println("\n🎭 CHANNEL TYPES DEMONSTRATION")
	fmt.Println("=============================")

	fmt.Println("\n6.1 Different channel types")
	
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
	
	fmt.Println("\n6.2 Type checking")
	fmt.Printf("ch2 is unbuffered: %t\n", cap(ch2) == 0)
	fmt.Printf("ch4 is buffered: %t\n", cap(ch4) > 0)
	fmt.Printf("ch5 is buffered: %t\n", cap(ch5) > 0)
}

// ============================================================================
// EXPORTED FUNCTIONS FOR MAIN
// ============================================================================

func runChannelTypes() {
	fmt.Println("🔗 GO CHANNELS: TYPES AND CAPACITY")
	fmt.Println("==================================")
	
	// Run all channel type examples
	unbufferedChannels()
	bufferedChannels()
	channelCapacityComparison()
	channelLengthAndCapacity()
	nilChannels()
	channelTypesDemonstration()
	
	fmt.Println("\n✅ Channel types completed!")
	fmt.Println("\nNext: Run 'go run . operations' to learn about channel operations")
}
