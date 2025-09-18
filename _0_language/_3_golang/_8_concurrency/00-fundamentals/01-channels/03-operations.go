package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. SENDING DATA
// ============================================================================

func sendingData() {
	// Sending data: ch <- data
	// This operation blocks until someone receives the data (unbuffered)
	// or until there's space in the buffer (buffered)
	
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
	// Receiving data: data := <-ch
	// This operation blocks until someone sends data
	
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
	// data, ok := <-ch
	// ok is true if data was received
	// ok is false if channel is closed
	
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
	// close(ch) - closes the channel
	// After closing, you cannot send data
	// You can still receive data that was already sent
	
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

	// for data := range ch
	// This automatically receives data until channel is closed
	
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

	// Using select with default case for non-blocking operations
	
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

// ✅ Always check if channel is closed when receiving
// ✅ Use buffered channels when you want to decouple sender/receiver
// ✅ Use unbuffered channels when you need tight synchronization
// ✅ Close channels to signal completion
// ✅ Use range to receive data until channel is closed
// ✅ Use select with default for non-blocking operations
// ❌ Don't send to closed channels
// ❌ Don't close channels multiple times
// ❌ Don't send to nil channels
// ❌ Don't receive from nil channels
// ❌ Don't forget to close channels when done
