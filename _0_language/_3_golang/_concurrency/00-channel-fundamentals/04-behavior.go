package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. BLOCKING BEHAVIOR
// ============================================================================

func blockingBehavior() {
	fmt.Println("\n🚫 BLOCKING BEHAVIOR")
	fmt.Println("===================")

	// Channels block when:
	// - Sending on unbuffered channel (until someone receives)
	// - Sending on full buffered channel (until space available)
	// - Receiving on empty channel (until someone sends)
	// - Sending/receiving on nil channel (blocks forever)
	
	fmt.Println("\n1.1 When channels block")
	
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
	fmt.Println("\n1.2 Buffered channel blocking")
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
	fmt.Println("\n⚡ NON-BLOCKING BEHAVIOR")
	fmt.Println("======================")

	// Channels don't block when:
	// - Sending on buffered channel with space
	// - Receiving from buffered channel with data
	// - Using select with default case
	
	fmt.Println("\n2.1 Non-blocking with buffered channels")
	ch := make(chan int, 3)
	
	// These sends won't block (buffer has space)
	fmt.Println("  Sending to buffered channel...")
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
	fmt.Println("\n🔄 CHANNEL STATES")
	fmt.Println("================")

	// Channels have different states:
	// - Open: Can send and receive
	// - Closed: Can only receive (gets zero value)
	// - Nil: Cannot send or receive (blocks forever)
	
	fmt.Println("\n3.1 Open channel state")
	ch := make(chan int, 2)
	
	// Send data
	ch <- 1
	ch <- 2
	fmt.Println("  Channel is open - can send and receive")
	
	// Receive data
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
	
	fmt.Println("\n3.2 Closed channel state")
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
	fmt.Println("\n🔄 CHANNEL SYNCHRONIZATION PATTERNS")
	fmt.Println("==================================")

	fmt.Println("\n4.1 Signal pattern")
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
	
	fmt.Println("\n4.2 Data transfer pattern")
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
	fmt.Println("\n💀 CHANNEL DEADLOCKS")
	fmt.Println("===================")

	// Deadlocks occur when:
	// - All goroutines are blocked waiting for each other
	// - Circular dependency in channel operations
	// - Sending to nil channel
	// - Receiving from nil channel
	
	fmt.Println("\n5.1 Common deadlock scenarios")
	
	// Scenario 1: Sending to nil channel
	fmt.Println("  Scenario 1: Sending to nil channel")
	// var nilCh chan int
	// nilCh <- 42  // This would block forever (commented out)
	fmt.Println("  ⚠️  Sending to nil channel blocks forever!")
	
	// Scenario 2: Receiving from nil channel
	fmt.Println("\n  Scenario 2: Receiving from nil channel")
	// data := <-nilCh  // This would block forever (commented out)
	fmt.Println("  ⚠️  Receiving from nil channel blocks forever!")
	
	// Scenario 3: Circular dependency
	fmt.Println("\n  Scenario 3: Circular dependency")
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	
	// This would create a deadlock (commented out)
	// go func() { ch1 <- <-ch2 }()  // Waits for ch2
	// go func() { ch2 <- <-ch1 }()  // Waits for ch1
	fmt.Println("  ⚠️  Circular dependency causes deadlock!")
}

// ============================================================================
// 6. CHANNEL PERFORMANCE CHARACTERISTICS
// ============================================================================

func channelPerformanceCharacteristics() {
	fmt.Println("\n⚡ CHANNEL PERFORMANCE CHARACTERISTICS")
	fmt.Println("====================================")

	fmt.Println("\n6.1 Performance characteristics")
	
	// Unbuffered channels
	fmt.Println("  Unbuffered channels:")
	fmt.Println("    - Synchronous communication")
	fmt.Println("    - Higher overhead (context switching)")
	fmt.Println("    - Better for tight synchronization")
	
	// Buffered channels
	fmt.Println("\n  Buffered channels:")
	fmt.Println("    - Asynchronous communication")
	fmt.Println("    - Lower overhead (less context switching)")
	fmt.Println("    - Better for decoupling producer/consumer")
	
	// Channel operations
	fmt.Println("\n  Channel operations:")
	fmt.Println("    - Send/receive are O(1) operations")
	fmt.Println("    - Memory overhead is minimal")
	fmt.Println("    - Garbage collected automatically")
}

// ============================================================================
// 7. CHANNEL MEMORY MODEL
// ============================================================================

func channelMemoryModel() {
	fmt.Println("\n🧠 CHANNEL MEMORY MODEL")
	fmt.Println("======================")

	fmt.Println("\n7.1 Memory model guarantees")
	
	// Go's memory model guarantees:
	// - A send on a channel happens before the corresponding receive
	// - A receive on a channel happens before the corresponding send completes
	// - Closing a channel happens before a receive of the zero value
	
	fmt.Println("  Memory model guarantees:")
	fmt.Println("    - Send happens before receive")
	fmt.Println("    - Receive happens before send completes")
	fmt.Println("    - Close happens before zero value receive")
	
	// Example demonstrating happens-before relationship
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
	fmt.Println("\n♻️  CHANNEL LIFECYCLE MANAGEMENT")
	fmt.Println("==============================")

	fmt.Println("\n8.1 Channel lifecycle stages")
	
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
	fmt.Println("  Stage 4: Channel will be garbage collected")
	
	fmt.Println("\n8.2 Best practices for lifecycle management")
	fmt.Println("  ✅ Close channels when done")
	fmt.Println("  ✅ Use defer to ensure cleanup")
	fmt.Println("  ✅ Don't close channels multiple times")
	fmt.Println("  ✅ Let garbage collector handle cleanup")
}

// ============================================================================
// EXPORTED FUNCTIONS FOR MAIN
// ============================================================================

func runChannelBehavior() {
	fmt.Println("🎭 GO CHANNELS: BEHAVIOR")
	fmt.Println("========================")
	
	// Run all channel behavior examples
	blockingBehavior()
	nonBlockingBehavior()
	channelStates()
	channelSynchronizationPatterns()
	channelDeadlocks()
	channelPerformanceCharacteristics()
	channelMemoryModel()
	channelLifecycleManagement()
	
	fmt.Println("\n✅ Channel behavior completed!")
	fmt.Println("\nNext: Run 'go run . patterns' to learn about channel patterns")
}
