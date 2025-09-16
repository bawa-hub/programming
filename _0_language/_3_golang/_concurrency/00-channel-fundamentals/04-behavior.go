package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. BLOCKING BEHAVIOR
// ============================================================================

func blockingBehavior() {

	// Channels block when:
	// - Sending on unbuffered channel (until someone receives)
	// - Sending on full buffered channel (until space available)
	// - Receiving on empty channel (until someone sends)
	// - Sending/receiving on nil channel (blocks forever)
		
	// Unbuffered channel - sender blocks until receiver
	ch1 := make(chan int)
	
	go func() {
		fmt.Println("  Sender: About to send on unbuffered channel...")
		ch1 <- 42  // This blocks until someone receives
		fmt.Println("  Sender: Data sent! (This prints after receiver gets data)")
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Receiver: About to receive from unbuffered channel...")
		data := <-ch1
		fmt.Printf("  Receiver: Got %d\n", data)
	}()
	
	time.Sleep(1 * time.Second)
	
	// Buffered channel - sender blocks when buffer is full
	ch2 := make(chan int, 2)  // Capacity of 2
	
	// Fill the buffer
	ch2 <- 1
	ch2 <- 2
	fmt.Println("  Buffer filled (2/2)")
	
	go func() {
		fmt.Println("  Sender: Trying to send 3rd value (buffer full)...")
		ch2 <- 3  // This blocks because buffer is full
		fmt.Println("  Sender: 3rd value sent! (This prints after receiver gets data)")
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Receiver: About to receive (making space)...")
		data := <-ch2
		fmt.Printf("  Receiver: Got %d (made space for sender)\n", data)
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 2. NON-BLOCKING BEHAVIOR
// ============================================================================

func nonBlockingBehavior() {

	// Channels don't block when:
	// - Sending on buffered channel with space
	// - Receiving from buffered channel with data
	// - Using select with default case
	
	ch := make(chan int, 3)
	
	// These sends won't block (buffer has space)
	ch <- 1  // No blocking
	ch <- 2  // No blocking
	ch <- 3  // No blocking
	fmt.Println("  All sends completed without blocking!")
	
	// These receives won't block (data available)
	fmt.Println("  Receiving from buffered channel...")
	fmt.Printf("  Received: %d\n", <-ch)  // No blocking
	fmt.Printf("  Received: %d\n", <-ch)  // No blocking
	fmt.Printf("  Received: %d\n", <-ch)  // No blocking
	fmt.Println("  All receives completed without blocking!")
}

// ============================================================================
// 3. CHANNEL STATES
// ============================================================================

func channelStates() {

	// Channels have different states:
	// - Open: Can send and receive
	// - Closed: Can only receive (gets zero value)
	// - Nil: Cannot send or receive (blocks forever)
	
	ch := make(chan int, 2)
	
	// Send data
	ch <- 1
	ch <- 2
	fmt.Println("  Channel is open - can send and receive")
	
	// Receive data
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
	
	close(ch)
	fmt.Println("  Channel is closed - can only receive")
	
	// Try to send (this would panic)
	// ch <- 3  // panic: send on closed channel
	
	// Can still receive (gets zero value)
	data, ok := <-ch
	fmt.Printf("  Received from closed channel: %d, ok: %t\n", data, ok)
	
	fmt.Println("\n3.3 Nil channel state")
	var nilCh chan int
	fmt.Printf("  Nil channel: %v\n", nilCh)
	fmt.Println("  Nil channel blocks forever on send/receive")
}

// ============================================================================
// 4. CHANNEL SYNCHRONIZATION PATTERNS
// ============================================================================

func channelSynchronizationPatterns() {

	// Use channel to signal completion
	done := make(chan bool)
	
	go func() {
		fmt.Println("  Worker: Starting work...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Worker: Work completed!")
		done <- true  // Signal completion
	}()
	
	fmt.Println("  Main: Waiting for worker...")
	<-done  // Wait for signal
	fmt.Println("  Main: Worker finished!")
	
	// Use channel to transfer data
	dataCh := make(chan string, 2)
	
	go func() {
		dataCh <- "Hello"
		dataCh <- "World"
		close(dataCh)
	}()
	
	fmt.Println("  Main: Receiving data...")
	for data := range dataCh {
		fmt.Printf("  Received: %s\n", data)
	}
}

// ============================================================================
// 5. CHANNEL DEADLOCKS
// ============================================================================

func channelDeadlocks() {

	// Deadlocks occur when:
	// - All goroutines are blocked waiting for each other
	// - Circular dependency in channel operations
	// - Sending to nil channel
	// - Receiving from nil channel
		
	// Scenario 1: Sending to nil channel
	// var nilCh chan int
	// nilCh <- 42  // This would block forever (commented out)
	
	// Scenario 2: Receiving from nil channel
	// data := <-nilCh  // This would block forever (commented out)
	
	// Scenario 3: Circular dependency
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	
	// This would create a deadlock (commented out)
	// go func() { ch1 <- <-ch2 }()  // Waits for ch2
	// go func() { ch2 <- <-ch1 }()  // Waits for ch1
	fmt.Println("  ⚠️  Circular dependency causes deadlock!")
}

// ============================================================================
// 7. CHANNEL MEMORY MODEL
// ============================================================================

func channelMemoryModel() {
	
	// Go's memory model guarantees:
	// - A send on a channel happens before the corresponding receive
	// - A receive on a channel happens before the corresponding send completes
	// - Closing a channel happens before a receive of the zero value
	
	ch := make(chan int)
	
	go func() {
		ch <- 42  // Send happens before receive
		fmt.Println("  Sender: Data sent")
	}()
	
	go func() {
		data := <-ch  // Receive happens after send
		fmt.Printf("  Receiver: Got %d\n", data)
	}()
	
	time.Sleep(500 * time.Millisecond)
}

// ============================================================================
// 8. CHANNEL LIFECYCLE MANAGEMENT
// ============================================================================

func channelLifecycleManagement() {

	// Stage 1: Creation
	ch := make(chan int, 2)
	fmt.Println("  Stage 1: Channel created")
	
	// Stage 2: Usage
	ch <- 1
	ch <- 2
	fmt.Println("  Stage 2: Channel used for communication")
	
	// Stage 3: Closing
	close(ch)
	fmt.Println("  Stage 3: Channel closed")
	
	// Stage 4: Cleanup (automatic)
	// Channel will be garbage collected
}
