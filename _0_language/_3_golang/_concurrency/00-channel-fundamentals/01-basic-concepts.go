package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. BASIC CHANNEL CONCEPTS
// ============================================================================

func basicChannelConcepts() {

	// 1.1 What is a channel?
	// A channel is a communication mechanism that allows goroutines to send and receive data
	// Think of it as a pipe where data flows from one goroutine to another

	// 1.2 Creating a channel
	// Syntax: make(chan Type)
	// This creates an unbuffered channel of the specified type
	
	fmt.Println("\n1.2 Creating a channel")
	var ch chan int                    // Declare a channel (nil by default)
	ch = make(chan int)                // Initialize the channel
	fmt.Printf("Channel created: %v\n", ch)
	fmt.Printf("Channel type: %T\n", ch)
	fmt.Printf("Channel is nil: %t\n", ch == nil)

	// 1.3 Channel zero value
	// The zero value of a channel is nil
	// You cannot send or receive on a nil channel (it will block forever)
	
	fmt.Println("\n1.3 Channel zero value")
	var nilCh chan int
	fmt.Printf("Nil channel: %v\n", nilCh)
	fmt.Printf("Nil channel is nil: %t\n", nilCh == nil)

	// 1.4 Basic send and receive
	// Sending: ch <- data
	// Receiving: data := <-ch
	
	fmt.Println("\n1.4 Basic send and receive")
	
	// Start a goroutine to receive data
	go func() {
		fmt.Println("  Goroutine: Waiting to receive data...")
		data := <-ch  // This will block until data is sent
		fmt.Printf("  Goroutine: Received data: %d\n", data)
	}()
	
	// Send data from main goroutine
	fmt.Println("  Main: Sending data...")
	ch <- 42  // This will block until someone receives
	fmt.Println("  Main: Data sent successfully!")
	
	// Wait a bit to see the output
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// 2. CHANNEL SYNCHRONIZATION
// ============================================================================

func channelSynchronization() {

	// Channels naturally synchronize goroutines
	// When you send data, the sender blocks until someone receives
	// When you receive data, the receiver blocks until someone sends
		
	ch := make(chan string)
	
	// Goroutine 1: Sender
	go func() {
		fmt.Println("  Sender: About to send 'Hello'")
		ch <- "Hello"  // Blocks until someone receives
		fmt.Println("  Sender: 'Hello' sent successfully")
		
		ch <- "World"  // Blocks until someone receives
		fmt.Println("  Sender: 'World' sent successfully")
	}()
	
	// Goroutine 2: Receiver
	go func() {
		time.Sleep(500 * time.Millisecond)  // Wait a bit
		fmt.Println("  Receiver: About to receive data")
		
		msg1 := <-ch  // Blocks until someone sends
		fmt.Printf("  Receiver: Received: %s\n", msg1)
		
		msg2 := <-ch  // Blocks until someone sends
		fmt.Printf("  Receiver: Received: %s\n", msg2)
	}()
	
	// Wait for both goroutines to complete
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 3. CHANNEL AS A SIGNAL
// ============================================================================

func channelAsSignal() {

	// Channels can be used to signal between goroutines
	// You don't always need to send data - just the act of sending/receiving can be a signal
	
	fmt.Println("\n3.1 Using channel as a signal")
	
	done := make(chan bool)  // Channel to signal completion
	
	// Worker goroutine
	go func() {
		fmt.Println("  Worker: Starting work...")
		time.Sleep(1 * time.Second)  // Simulate work
		fmt.Println("  Worker: Work completed!")
		
		done <- true  // Signal that work is done
	}()
	
	// Main goroutine waits for signal
	fmt.Println("  Main: Waiting for worker to finish...")
	<-done  // Block until we receive the signal
	fmt.Println("  Main: Worker finished!")
}

// ============================================================================
// 4. CHANNEL DIRECTIONALITY
// ============================================================================

func channelDirectionality() {

	// Channels can be:
	// - Bidirectional: chan Type (can send and receive)
	// - Send-only: chan<- Type (can only send)
	// - Receive-only: <-chan Type (can only receive)
		
	// Bidirectional channel
	ch := make(chan int)
	fmt.Printf("Bidirectional channel: %T\n", ch)
	
	// Send-only channel (chan<- int)
	var sendCh chan<- int = ch
	fmt.Printf("Send-only channel: %T\n", sendCh)
	
	// Receive-only channel (<-chan int)
	var recvCh <-chan int = ch
	fmt.Printf("Receive-only channel: %T\n", recvCh)
		
	// Function that only sends data
	go func() {
		sendCh <- 100  // Can only send
		fmt.Println("  Sent: 100")
	}()
	
	// Function that only receives data
	go func() {
		data := <-recvCh  // Can only receive
		fmt.Printf("  Received: %d\n", data)
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// 5. CHANNEL LIFECYCLE
// ============================================================================

func channelLifecycle() {

	// Channels have a lifecycle:
	// 1. Created with make()
	// 2. Used for communication
	// 3. Closed with close()
	// 4. Garbage collected when no longer referenced
	
	fmt.Println("\nChannel lifecycle stages")
	
	ch := make(chan int)
	fmt.Println("  Stage 1: Channel created")
	
	// Send some data
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)  // Close the channel
		fmt.Println("  Stage 2: Channel closed")
	}()
	
	// Receive data until channel is closed
	fmt.Println("  Stage 3: Receiving data...")
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("  Stage 4: Channel is closed, no more data")
			break
		}
		fmt.Printf("  Received: %d\n", data)
	}
}
