package main

import (
	"fmt"
	"time"
)

// 📡 CHANNEL TYPES MASTERY
// Understanding the different types and behaviors of channels

func main() {
	fmt.Println("📡 CHANNEL TYPES MASTERY")
	fmt.Println("========================")

	// 1. Unbuffered Channels
	fmt.Println("\n1. Unbuffered Channels:")
	unbufferedChannels()

	// 2. Buffered Channels
	fmt.Println("\n2. Buffered Channels:")
	bufferedChannels()

	// 3. Directional Channels
	fmt.Println("\n3. Directional Channels:")
	directionalChannels()

	// 4. Channel Zero Values
	fmt.Println("\n4. Channel Zero Values:")
	channelZeroValues()

	// 5. Channel Capacity and Length
	fmt.Println("\n5. Channel Capacity and Length:")
	channelCapacityAndLength()

	// 6. Channel Closing
	fmt.Println("\n6. Channel Closing:")
	channelClosing()
}

// UNBUFFERED CHANNELS: Synchronous communication
func unbufferedChannels() {
	fmt.Println("Understanding unbuffered channels...")
	
	// Create unbuffered channel
	ch := make(chan int)
	
	// Send and receive must happen simultaneously
	go func() {
		fmt.Println("  🧵 Goroutine: Sending value 42")
		ch <- 42
		fmt.Println("  🧵 Goroutine: Value sent")
	}()
	
	// Receive the value
	value := <-ch
	fmt.Printf("  📡 Received: %d\n", value)
	
	// Demonstrate blocking behavior
	fmt.Println("  🔄 Demonstrating blocking behavior...")
	
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("  🧵 Goroutine: Sending value 100")
		ch <- 100
	}()
	
	fmt.Println("  ⏳ Main: Waiting for value...")
	value = <-ch
	fmt.Printf("  📡 Received: %d\n", value)
}

// BUFFERED CHANNELS: Asynchronous communication
func bufferedChannels() {
	fmt.Println("Understanding buffered channels...")
	
	// Create buffered channel with capacity 3
	ch := make(chan int, 3)
	
	// Can send multiple values without blocking
	fmt.Println("  📤 Sending values to buffered channel...")
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("  📊 Channel length: %d, capacity: %d\n", len(ch), cap(ch))
	
	// Can receive values
	fmt.Println("  📥 Receiving values...")
	fmt.Printf("  📡 Received: %d\n", <-ch)
	fmt.Printf("  📡 Received: %d\n", <-ch)
	fmt.Printf("  📡 Received: %d\n", <-ch)
	fmt.Printf("  📊 Channel length: %d, capacity: %d\n", len(ch), cap(ch))
	
	// Demonstrate blocking when buffer is full
	fmt.Println("  🔄 Demonstrating buffer overflow...")
	
	// Fill the buffer
	ch <- 10
	ch <- 20
	ch <- 30
	
	// This will block until space is available
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("  🧵 Goroutine: Receiving value to make space")
		<-ch
	}()
	
	fmt.Println("  ⏳ Main: Trying to send to full buffer...")
	ch <- 40
	fmt.Println("  ✅ Main: Value sent after space became available")
}

// DIRECTIONAL CHANNELS: Send-only and receive-only channels
func directionalChannels() {
	fmt.Println("Understanding directional channels...")
	
	// Create bidirectional channel
	ch := make(chan int)
	
	// Pass as send-only channel
	go sender(ch)
	
	// Pass as receive-only channel
	receiver(ch)
}

// Function that takes a send-only channel
func sender(ch chan<- int) {
	fmt.Println("  📤 Sender: Sending values...")
	for i := 1; i <= 3; i++ {
		ch <- i * 10
		fmt.Printf("  📤 Sender: Sent %d\n", i*10)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	fmt.Println("  📤 Sender: Channel closed")
}

// Function that takes a receive-only channel
func receiver(ch <-chan int) {
	fmt.Println("  📥 Receiver: Receiving values...")
	for value := range ch {
		fmt.Printf("  📥 Receiver: Received %d\n", value)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println("  📥 Receiver: Channel closed, done receiving")
}

// CHANNEL ZERO VALUES: Understanding nil channels
func channelZeroValues() {
	fmt.Println("Understanding channel zero values...")
	
	// Zero value of channel is nil
	var ch chan int
	fmt.Printf("  📊 Zero value channel: %v\n", ch)
	fmt.Printf("  📊 Is nil: %t\n", ch == nil)
	
	// Sending to nil channel blocks forever
	fmt.Println("  ⚠️  Sending to nil channel will block forever!")
	
	// Receiving from nil channel blocks forever
	fmt.Println("  ⚠️  Receiving from nil channel will block forever!")
	
	// Use select to handle nil channels
	fmt.Println("  🔄 Demonstrating nil channel handling with select...")
	
	var nilCh chan int
	timeout := time.After(1 * time.Second)
	
	select {
	case <-nilCh:
		fmt.Println("  📡 This will never execute")
	case <-timeout:
		fmt.Println("  ⏰ Timeout: nil channel blocks forever")
	}
}

// CHANNEL CAPACITY AND LENGTH: Understanding channel state
func channelCapacityAndLength() {
	fmt.Println("Understanding channel capacity and length...")
	
	// Create buffered channel
	ch := make(chan string, 5)
	
	fmt.Printf("  📊 Initial - Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Add some values
	ch <- "Hello"
	ch <- "World"
	ch <- "Go"
	
	fmt.Printf("  📊 After adding 3 values - Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Remove a value
	value := <-ch
	fmt.Printf("  📡 Removed: %s\n", value)
	fmt.Printf("  📊 After removing 1 value - Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Demonstrate length vs capacity
	fmt.Println("  🔄 Demonstrating length vs capacity...")
	
	// Fill to capacity
	ch <- "A"
	ch <- "B"
	fmt.Printf("  📊 Full buffer - Length: %d, Capacity: %d\n", len(ch), cap(ch))
	
	// Empty the channel
	for len(ch) > 0 {
		value := <-ch
		fmt.Printf("  📡 Removed: %s\n", value)
	}
	fmt.Printf("  📊 Empty buffer - Length: %d, Capacity: %d\n", len(ch), cap(ch))
}

// CHANNEL CLOSING: Proper channel lifecycle management
func channelClosing() {
	fmt.Println("Understanding channel closing...")
	
	// Create channel
	ch := make(chan int, 3)
	
	// Send some values
	ch <- 1
	ch <- 2
	ch <- 3
	
	// Close the channel
	close(ch)
	fmt.Println("  🔒 Channel closed")
	
	// Demonstrate closed channel behavior
	fmt.Println("  📥 Receiving from closed channel...")
	
	// Can still receive values that were sent before closing
	for i := 0; i < 3; i++ {
		value, ok := <-ch
		if ok {
			fmt.Printf("  📡 Received: %d\n", value)
		} else {
			fmt.Println("  📡 Channel is closed")
		}
	}
	
	// Sending to closed channel panics
	fmt.Println("  ⚠️  Sending to closed channel will panic!")
	
	// Use defer and recover to handle panic gracefully
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  🚨 Panic recovered: %v\n", r)
		}
	}()
	
	// This will panic
	ch <- 4
}
