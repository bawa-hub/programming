package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. SENDING DATA
// ============================================================================

func sendingData() {
	fmt.Println("\n📤 SENDING DATA")
	fmt.Println("===============")

	// Sending data: ch <- data
	// This operation blocks until someone receives the data (unbuffered)
	// or until there's space in the buffer (buffered)
	
	fmt.Println("\n1.1 Basic sending")
	ch := make(chan int)
	
	go func() {
		fmt.Println("  Sender: About to send 42...")
		ch <- 42  // Send data
		fmt.Println("  Sender: 42 sent successfully!")
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Receiver: About to receive...")
		data := <-ch
		fmt.Printf("  Receiver: Got %d\n", data)
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 2. RECEIVING DATA
// ============================================================================

func receivingData() {
	fmt.Println("\n📥 RECEIVING DATA")
	fmt.Println("================")

	// Receiving data: data := <-ch
	// This operation blocks until someone sends data
	
	fmt.Println("\n2.1 Basic receiving")
	ch := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Sender: Sending 'Hello'...")
		ch <- "Hello"
	}()
	
	go func() {
		fmt.Println("  Receiver: Waiting for data...")
		data := <-ch  // Receive data
		fmt.Printf("  Receiver: Got '%s'\n", data)
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 3. RECEIVING WITH OK CHECK
// ============================================================================

func receivingWithOkCheck() {
	fmt.Println("\n✅ RECEIVING WITH OK CHECK")
	fmt.Println("=========================")

	// data, ok := <-ch
	// ok is true if data was received
	// ok is false if channel is closed
	
	fmt.Println("\n3.1 Using ok to check channel state")
	ch := make(chan int, 3)
	
	// Send some data
	ch <- 1
	ch <- 2
	ch <- 3
	
	// Close the channel
	close(ch)
	
	// Receive data and check ok
	fmt.Println("  Receiving data with ok check...")
	for i := 0; i < 5; i++ {  // Try to receive more than we sent
		data, ok := <-ch
		if ok {
			fmt.Printf("  Received: %d (channel open)\n", data)
		} else {
			fmt.Println("  Channel is closed!")
			break
		}
	}
}

// ============================================================================
// 4. CLOSING CHANNELS
// ============================================================================

func closingChannels() {
	fmt.Println("\n🔒 CLOSING CHANNELS")
	fmt.Println("==================")

	// close(ch) - closes the channel
	// After closing, you cannot send data
	// You can still receive data that was already sent
	
	fmt.Println("\n4.1 Basic channel closing")
	ch := make(chan int, 3)
	
	// Send some data
	ch <- 1
	ch <- 2
	ch <- 3
	
	fmt.Println("  Data sent, closing channel...")
	close(ch)
	fmt.Println("  Channel closed!")
	
	// Try to send after closing (this would panic)
	// ch <- 4  // This would panic: "send on closed channel"
	
	// But we can still receive data
	fmt.Println("  Receiving data after closing...")
	for data := range ch {
		fmt.Printf("  Received: %d\n", data)
	}
}

// ============================================================================
// 5. RANGE OVER CHANNELS
// ============================================================================

func rangeOverChannels() {
	fmt.Println("\n🔄 RANGE OVER CHANNELS")
	fmt.Println("=====================")

	// for data := range ch
	// This automatically receives data until channel is closed
	
	fmt.Println("\n5.1 Using range to receive data")
	ch := make(chan string, 3)
	
	// Send data in a goroutine
	go func() {
		ch <- "First"
		ch <- "Second"
		ch <- "Third"
		close(ch)  // Must close for range to work
	}()
	
	// Receive data using range
	fmt.Println("  Receiving data with range...")
	for data := range ch {
		fmt.Printf("  Received: %s\n", data)
	}
	fmt.Println("  Range loop ended (channel closed)")
}

// ============================================================================
// 6. CHANNEL OPERATIONS IN DIFFERENT CONTEXTS
// ============================================================================

func channelOperationsInContexts() {
	fmt.Println("\n🎭 CHANNEL OPERATIONS IN CONTEXTS")
	fmt.Println("=================================")

	fmt.Println("\n6.1 Sending in different goroutines")
	ch := make(chan int, 2)
	
	// Multiple senders
	go func() {
		ch <- 1
		fmt.Println("  Sender 1: Sent 1")
	}()
	
	go func() {
		ch <- 2
		fmt.Println("  Sender 2: Sent 2")
	}()
	
	// Multiple receivers
	go func() {
		data := <-ch
		fmt.Printf("  Receiver 1: Got %d\n", data)
	}()
	
	go func() {
		data := <-ch
		fmt.Printf("  Receiver 2: Got %d\n", data)
	}()
	
	time.Sleep(500 * time.Millisecond)
}

// ============================================================================
// 7. NON-BLOCKING OPERATIONS
// ============================================================================

func nonBlockingOperations() {
	fmt.Println("\n⚡ NON-BLOCKING OPERATIONS")
	fmt.Println("=========================")

	// Using select with default case for non-blocking operations
	
	fmt.Println("\n7.1 Non-blocking send")
	ch := make(chan int, 1)
	
	// This send won't block because channel has space
	select {
	case ch <- 42:
		fmt.Println("  Non-blocking send: Success!")
	default:
		fmt.Println("  Non-blocking send: Would block")
	}
	
	// This send would block, so default case executes
	select {
	case ch <- 43:
		fmt.Println("  Non-blocking send: Success!")
	default:
		fmt.Println("  Non-blocking send: Would block (buffer full)")
	}
	
	fmt.Println("\n7.2 Non-blocking receive")
	
	// This receive won't block because data is available
	select {
	case data := <-ch:
		fmt.Printf("  Non-blocking receive: Got %d\n", data)
	default:
		fmt.Println("  Non-blocking receive: No data available")
	}
	
	// This receive would block, so default case executes
	select {
	case data := <-ch:
		fmt.Printf("  Non-blocking receive: Got %d\n", data)
	default:
		fmt.Println("  Non-blocking receive: No data available")
	}
}

// ============================================================================
// 8. CHANNEL OPERATION ERRORS
// ============================================================================

func channelOperationErrors() {
	fmt.Println("\n❌ CHANNEL OPERATION ERRORS")
	fmt.Println("==========================")

	fmt.Println("\n8.1 Common channel operation errors")
	
	// Error 1: Sending to closed channel
	fmt.Println("  Error 1: Sending to closed channel")
	ch := make(chan int)
	close(ch)
	
	// This would panic (commented out to avoid crash)
	// ch <- 42  // panic: send on closed channel
	fmt.Println("  ⚠️  Sending to closed channel causes panic!")
	
	// Error 2: Closing already closed channel
	fmt.Println("\n  Error 2: Closing already closed channel")
	// close(ch)  // panic: close of closed channel
	fmt.Println("  ⚠️  Closing already closed channel causes panic!")
	
	// Error 3: Sending to nil channel
	fmt.Println("\n  Error 3: Sending to nil channel")
	// var nilCh chan int
	// nilCh <- 42  // This would block forever
	fmt.Println("  ⚠️  Sending to nil channel blocks forever!")
	
	// Error 4: Receiving from nil channel
	fmt.Println("\n  Error 4: Receiving from nil channel")
	// data := <-nilCh  // This would block forever
	fmt.Println("  ⚠️  Receiving from nil channel blocks forever!")
}

// ============================================================================
// 9. CHANNEL OPERATION BEST PRACTICES
// ============================================================================

func channelOperationBestPractices() {
	fmt.Println("\n💡 CHANNEL OPERATION BEST PRACTICES")
	fmt.Println("==================================")

	fmt.Println("\n9.1 Best practices for channel operations")
	
	fmt.Println("  ✅ Always check if channel is closed when receiving")
	fmt.Println("  ✅ Use buffered channels when you want to decouple sender/receiver")
	fmt.Println("  ✅ Use unbuffered channels when you need tight synchronization")
	fmt.Println("  ✅ Close channels to signal completion")
	fmt.Println("  ✅ Use range to receive data until channel is closed")
	fmt.Println("  ✅ Use select with default for non-blocking operations")
	
	fmt.Println("\n9.2 What to avoid")
	fmt.Println("  ❌ Don't send to closed channels")
	fmt.Println("  ❌ Don't close channels multiple times")
	fmt.Println("  ❌ Don't send to nil channels")
	fmt.Println("  ❌ Don't receive from nil channels")
	fmt.Println("  ❌ Don't forget to close channels when done")
}

// ============================================================================
// EXPORTED FUNCTIONS FOR MAIN
// ============================================================================

func runChannelOperations() {
	fmt.Println("🔧 GO CHANNELS: OPERATIONS")
	fmt.Println("==========================")
	
	// Run all channel operation examples
	sendingData()
	receivingData()
	receivingWithOkCheck()
	closingChannels()
	rangeOverChannels()
	channelOperationsInContexts()
	nonBlockingOperations()
	channelOperationErrors()
	channelOperationBestPractices()
	
	fmt.Println("\n✅ Channel operations completed!")
	fmt.Println("\nNext: Run 'go run . behavior' to learn about channel behavior")
}
