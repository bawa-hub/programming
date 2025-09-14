package main

import (
	"fmt"
	"time"
)

// 📡 SELECT STATEMENT PATTERNS
// Mastering the select statement for channel multiplexing

func main() {
	fmt.Println("📡 SELECT STATEMENT PATTERNS")
	fmt.Println("============================")

	// 1. Basic Select
	fmt.Println("\n1. Basic Select:")
	basicSelect()

	// 2. Select with Default
	fmt.Println("\n2. Select with Default:")
	selectWithDefault()

	// 3. Select with Timeout
	fmt.Println("\n3. Select with Timeout:")
	selectWithTimeout()

	// 4. Select with Multiple Channels
	fmt.Println("\n4. Select with Multiple Channels:")
	selectWithMultipleChannels()

	// 5. Select for Non-blocking Operations
	fmt.Println("\n5. Non-blocking Operations:")
	nonBlockingOperations()

	// 6. Select for Channel Multiplexing
	fmt.Println("\n6. Channel Multiplexing:")
	channelMultiplexing()

	// 7. Select for Graceful Shutdown
	fmt.Println("\n7. Graceful Shutdown:")
	gracefulShutdown()
}

// BASIC SELECT: Choosing between channels
func basicSelect() {
	fmt.Println("Understanding basic select...")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Start goroutines that send to different channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case ch1 <- "Message from channel 1":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		select {
		case ch2 <- "Message from channel 2":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	// Select will choose the first available channel
	select {
	case msg1 := <-ch1:
		fmt.Printf("  📡 Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("  📡 Received from ch2: %s\n", msg2)
	}
	
	// Wait for goroutines to complete
	time.Sleep(300 * time.Millisecond)
	
	// Clean up
	close(ch1)
	close(ch2)
}

// SELECT WITH DEFAULT: Non-blocking select
func selectWithDefault() {
	fmt.Println("Understanding select with default...")
	
	ch := make(chan string)
	
	// Try to receive with default case
	select {
	case msg := <-ch:
		fmt.Printf("  📡 Received: %s\n", msg)
	default:
		fmt.Println("  📡 No message available, using default")
	}
	
	// Send a message and try again
	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case ch <- "Hello from goroutine":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	// Wait a bit and try again
	time.Sleep(150 * time.Millisecond)
	select {
	case msg := <-ch:
		fmt.Printf("  📡 Received: %s\n", msg)
	default:
		fmt.Println("  📡 No message available, using default")
	}
	
	close(ch)
}

// SELECT WITH TIMEOUT: Adding timeouts to channel operations
func selectWithTimeout() {
	fmt.Println("Understanding select with timeout...")
	
	ch := make(chan string)
	
	// Create timeout channel
	timeout := time.After(1 * time.Second)
	
	// Start a goroutine that might take longer than timeout
	go func() {
		time.Sleep(2 * time.Second)
		select {
		case ch <- "This message will arrive too late":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	// Select with timeout
	select {
	case msg := <-ch:
		fmt.Printf("  📡 Received: %s\n", msg)
	case <-timeout:
		fmt.Println("  ⏰ Timeout: No message received within 1 second")
	}
	
	// Wait for goroutine to complete
	time.Sleep(1 * time.Second)
	close(ch)
}

// SELECT WITH MULTIPLE CHANNELS: Handling multiple input sources
func selectWithMultipleChannels() {
	fmt.Println("Understanding select with multiple channels...")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	// Start goroutines that send to different channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case ch1 <- "Fast message":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	go func() {
		time.Sleep(300 * time.Millisecond)
		select {
		case ch2 <- "Medium message":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		select {
		case ch3 <- "Slow message":
		default:
			// Channel might be closed, ignore
		}
	}()
	
	// Select from multiple channels
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("  📡 Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("  📡 Received from ch2: %s\n", msg2)
		case msg3 := <-ch3:
			fmt.Printf("  📡 Received from ch3: %s\n", msg3)
		}
	}
	
	// Wait for goroutines to complete
	time.Sleep(600 * time.Millisecond)
	
	// Clean up
	close(ch1)
	close(ch2)
	close(ch3)
}

// NON-BLOCKING OPERATIONS: Using select for non-blocking sends/receives
func nonBlockingOperations() {
	fmt.Println("Understanding non-blocking operations...")
	
	ch := make(chan int, 2)
	
	// Non-blocking send
	fmt.Println("  📤 Non-blocking send:")
	select {
	case ch <- 1:
		fmt.Println("  ✅ Successfully sent 1")
	default:
		fmt.Println("  ❌ Channel full, send failed")
	}
	
	select {
	case ch <- 2:
		fmt.Println("  ✅ Successfully sent 2")
	default:
		fmt.Println("  ❌ Channel full, send failed")
	}
	
	select {
	case ch <- 3:
		fmt.Println("  ✅ Successfully sent 3")
	default:
		fmt.Println("  ❌ Channel full, send failed")
	}
	
	// Non-blocking receive
	fmt.Println("  📥 Non-blocking receive:")
	for i := 0; i < 3; i++ {
		select {
		case value := <-ch:
			fmt.Printf("  ✅ Received: %d\n", value)
		default:
			fmt.Println("  ❌ No value available")
		}
	}
	
	close(ch)
}

// CHANNEL MULTIPLEXING: Combining multiple channels into one
func channelMultiplexing() {
	fmt.Println("Understanding channel multiplexing...")
	
	// Create input channels
	input1 := make(chan string)
	input2 := make(chan string)
	input3 := make(chan string)
	
	// Start goroutines that send data
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- fmt.Sprintf("Input1-%d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 1; i <= 3; i++ {
			input2 <- fmt.Sprintf("Input2-%d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input3)
		for i := 1; i <= 3; i++ {
			input3 <- fmt.Sprintf("Input3-%d", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Multiplex all inputs into one output
	output := multiplex(input1, input2, input3)
	
	// Collect all results
	fmt.Println("  📊 Multiplexed results:")
	for result := range output {
		fmt.Printf("  📡 Received: %s\n", result)
	}
}

// Multiplex function that combines multiple channels
func multiplex(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		
		// Use select to read from any available input
		for {
			select {
			case msg, ok := <-inputs[0]:
				if !ok {
					inputs[0] = nil
				} else {
					output <- msg
				}
			case msg, ok := <-inputs[1]:
				if !ok {
					inputs[1] = nil
				} else {
					output <- msg
				}
			case msg, ok := <-inputs[2]:
				if !ok {
					inputs[2] = nil
				} else {
					output <- msg
				}
			}
			
			// Check if all inputs are closed
			if inputs[0] == nil && inputs[1] == nil && inputs[2] == nil {
				break
			}
		}
	}()
	
	return output
}

// GRACEFUL SHUTDOWN: Using select for graceful shutdown
func gracefulShutdown() {
	fmt.Println("Understanding graceful shutdown...")
	
	// Create shutdown signal
	shutdown := make(chan struct{})
	
	// Start a service
	service := startServiceWithSelect(shutdown)
	
	// Let it run for a bit
	time.Sleep(3 * time.Second)
	
	// Gracefully shutdown
	fmt.Println("  🛑 Initiating graceful shutdown...")
	close(shutdown)
	
	// Wait for service to stop
	<-service
	fmt.Println("  ✅ Service stopped gracefully")
}

// Service that uses select for graceful shutdown
func startServiceWithSelect(shutdown <-chan struct{}) <-chan struct{} {
	stopped := make(chan struct{})
	
	go func() {
		defer close(stopped)
		
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Println("  🧵 Service working...")
			case <-shutdown:
				fmt.Println("  🧵 Service shutting down...")
				time.Sleep(500 * time.Millisecond) // Cleanup time
				return
			}
		}
	}()
	
	return stopped
}
