package main

import (
	"context"
	"fmt"
	"time"
)

// Exercise 1: Basic Select
// Create a select statement that handles two channels.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Select")
	fmt.Println("========================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Send messages
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Hello from channel 1"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "Hello from channel 2"
	}()
	
	// Select from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
}

// Exercise 2: Non-blocking Select
// Implement non-blocking send and receive operations.
func Exercise2() {
	fmt.Println("\nExercise 2: Non-blocking Select")
	fmt.Println("===============================")
	
	ch := make(chan string, 2)
	
	// Non-blocking send
	fmt.Println("Testing non-blocking send:")
	select {
	case ch <- "Message 1":
		fmt.Println("✓ Message 1 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	select {
	case ch <- "Message 2":
		fmt.Println("✓ Message 2 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	select {
	case ch <- "Message 3":
		fmt.Println("✓ Message 3 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	// Non-blocking receive
	fmt.Println("\nTesting non-blocking receive:")
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("✓ Received: %s\n", msg)
		default:
			fmt.Println("✗ No message available")
		}
	}
}

// Exercise 3: Timeout Select
// Add timeout handling to channel operations.
func Exercise3() {
	fmt.Println("\nExercise 3: Timeout Select")
	fmt.Println("==========================")
	
	ch := make(chan string)
	
	// Test with timeout
	fmt.Println("Testing timeout (no message sent):")
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
	
	// Test with message sent after timeout
	fmt.Println("\nTesting timeout (message sent after timeout):")
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Late message"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
}

// Exercise 4: Priority Select
// Implement priority handling for multiple channels.
func Exercise4() {
	fmt.Println("\nExercise 4: Priority Select")
	fmt.Println("===========================")
	
	highPriority := make(chan string, 5)
	normalPriority := make(chan string, 5)
	lowPriority := make(chan string, 5)
	
	// Send messages to different priority channels
	go func() {
		time.Sleep(50 * time.Millisecond)
		normalPriority <- "Normal message 1"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		lowPriority <- "Low message 1"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		highPriority <- "High message 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		normalPriority <- "Normal message 2"
	}()
	
	// Handle messages with priority
	for i := 0; i < 4; i++ {
		select {
		case msg := <-highPriority:
			fmt.Printf("🔥 HIGH PRIORITY: %s\n", msg)
		case msg := <-normalPriority:
			fmt.Printf("📝 Normal priority: %s\n", msg)
		case msg := <-lowPriority:
			fmt.Printf("📄 Low priority: %s\n", msg)
		}
	}
}

// Exercise 5: Multiplexing Select
// Create a fan-in pattern using select.
func Exercise5() {
	fmt.Println("\nExercise 5: Multiplexing Select")
	fmt.Println("===============================")
	
	input1 := make(chan string)
	input2 := make(chan string)
	output := make(chan string)
	
	// Fan-in goroutine
	go func() {
		defer close(output)
		for {
			select {
			case msg, ok := <-input1:
				if !ok {
					input1 = nil // Disable this case
				} else {
					output <- "Input1: " + msg
				}
			case msg, ok := <-input2:
				if !ok {
					input2 = nil // Disable this case
				} else {
					output <- "Input2: " + msg
				}
			}
			
			// Exit when both channels are closed
			if input1 == nil && input2 == nil {
				return
			}
		}
	}()
	
	// Send messages to inputs
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 1; i <= 3; i++ {
			input2 <- fmt.Sprintf("Data %d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	// Collect multiplexed output
	fmt.Println("Multiplexed output:")
	for msg := range output {
		fmt.Printf("  %s\n", msg)
	}
}

// Exercise 6: Loop Select
// Implement a select statement in a loop with proper exit conditions.
func Exercise6() {
	fmt.Println("\nExercise 6: Loop Select")
	fmt.Println("=======================")
	
	ch := make(chan string)
	quit := make(chan bool)
	
	// Send messages
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Send quit signal
	go func() {
		time.Sleep(600 * time.Millisecond)
		quit <- true
	}()
	
	// Loop with select
	messageCount := 0
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return
			}
			messageCount++
			fmt.Printf("Received: %s (count: %d)\n", msg, messageCount)
		case <-quit:
			fmt.Println("Quit signal received, exiting loop")
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout in loop")
		}
	}
}

// Exercise 7: Error Handling Select
// Add error handling to select operations.
func Exercise7() {
	fmt.Println("\nExercise 7: Error Handling Select")
	fmt.Println("=================================")
	
	ch := make(chan string)
	errCh := make(chan error)
	
	// Send messages and errors
	go func() {
		defer close(ch)
		defer close(errCh)
		
		for i := 1; i <= 5; i++ {
			if i == 3 {
				errCh <- fmt.Errorf("Error at message %d", i)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Select with error handling
	messageCount := 0
	errorCount := 0
	
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				fmt.Printf("Total messages: %d, Total errors: %d\n", messageCount, errorCount)
				return
			}
			messageCount++
			fmt.Printf("✓ Received: %s\n", msg)
		case err, ok := <-errCh:
			if !ok {
				fmt.Println("Error channel closed")
				return
			}
			errorCount++
			fmt.Printf("✗ Error: %v\n", err)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout")
			return
		}
	}
}

// Exercise 8: Dynamic Select
// Create a select statement that handles a variable number of channels.
func Exercise8() {
	fmt.Println("\nExercise 8: Dynamic Select")
	fmt.Println("==========================")
	
	// Create multiple channels
	channels := make([]chan string, 3)
	for i := range channels {
		channels[i] = make(chan string)
	}
	
	// Send messages to different channels
	for i, ch := range channels {
		go func(id int, ch chan<- string) {
			defer close(ch)
			for j := 1; j <= 2; j++ {
				ch <- fmt.Sprintf("Channel %d, Message %d", id, j)
				time.Sleep(time.Duration(100+id*50) * time.Millisecond)
			}
		}(i, ch)
	}
	
	// Dynamic select (simulated with fixed number of channels)
	fmt.Println("Dynamic select results:")
	for i := 0; i < 6; i++ {
		select {
		case msg := <-channels[0]:
			fmt.Printf("  From channel 0: %s\n", msg)
		case msg := <-channels[1]:
			fmt.Printf("  From channel 1: %s\n", msg)
		case msg := <-channels[2]:
			fmt.Printf("  From channel 2: %s\n", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("  Timeout")
			return
		}
	}
}

// Exercise 9: Select with Ticker
func Exercise9() {
	fmt.Println("\nExercise 9: Select with Ticker")
	fmt.Println("=============================")
	
	ch := make(chan string)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	
	// Send messages
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	
	// Select with ticker
	tickCount := 0
	messageCount := 0
	
	for i := 0; i < 10; i++ {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				fmt.Printf("Total ticks: %d, Total messages: %d\n", tickCount, messageCount)
				return
			}
			messageCount++
			fmt.Printf("📨 Received: %s\n", msg)
		case <-ticker.C:
			tickCount++
			fmt.Printf("⏰ Tick %d\n", tickCount)
		}
	}
}

// Exercise 10: Select with Context
func Exercise10() {
	fmt.Println("\nExercise 10: Select with Context")
	fmt.Println("===============================")
	
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	// Send a message after delay
	go func() {
		time.Sleep(300 * time.Millisecond)
		select {
		case ch <- "Context message":
			fmt.Println("Message sent successfully")
		case <-ctx.Done():
			fmt.Println("Context cancelled before message could be sent")
		}
		close(ch)
	}()
	
	// Select with context
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-ctx.Done():
		fmt.Printf("Context cancelled: %v\n", ctx.Err())
	}
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Select Exercises")
	fmt.Println("===============================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}
