package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Example 1: Basic Channel Operations
func basicChannelOperations() {
	fmt.Println("1. Basic Channel Operations")
	fmt.Println("===========================")
	
	// Create a channel
	ch := make(chan string)
	
	// Send a value in a goroutine
	go func() {
		ch <- "Hello from goroutine!"
	}()
	
	// Receive the value
	message := <-ch
	fmt.Printf("Received: %s\n", message)
}

// Example 2: Buffered vs Unbuffered Channels
func bufferedVsUnbuffered() {
	fmt.Println("\n2. Buffered vs Unbuffered Channels")
	fmt.Println("===================================")
	
	// Unbuffered channel
	unbuffered := make(chan int)
	
	// Buffered channel
	buffered := make(chan int, 3)
	
	// Test unbuffered channel
	fmt.Println("Testing unbuffered channel:")
	go func() {
		unbuffered <- 1
		fmt.Println("Sent 1 to unbuffered channel")
	}()
	
	value := <-unbuffered
	fmt.Printf("Received %d from unbuffered channel\n", value)
	
	// Test buffered channel
	fmt.Println("\nTesting buffered channel:")
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("Sent 3 values to buffered channel")
	
	fmt.Printf("Received %d from buffered channel\n", <-buffered)
	fmt.Printf("Received %d from buffered channel\n", <-buffered)
	fmt.Printf("Received %d from buffered channel\n", <-buffered)
}

// Example 3: Channel Direction
func channelDirection() {
	fmt.Println("\n3. Channel Direction")
	fmt.Println("====================")
	
	// Create channels
	ch := make(chan int)
	
	// Send-only channel
	go sender(ch)
	
	// Receive-only channel
	receiver(ch)
}

func sender(ch chan<- int) {
	fmt.Println("Sender: Sending values...")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sender: Sent %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	fmt.Println("Receiver: Waiting for values...")
	for value := range ch {
		fmt.Printf("Receiver: Received %d\n", value)
	}
	fmt.Println("Receiver: Channel closed")
}

// Example 4: Channel Closing
func channelClosing() {
	fmt.Println("\n4. Channel Closing")
	fmt.Println("==================")
	
	ch := make(chan int)
	
	// Producer goroutine
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Producer: Sent %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Producer: Closing channel")
	}()
	
	// Consumer goroutine
	go func() {
		for {
			value, ok := <-ch
			if !ok {
				fmt.Println("Consumer: Channel closed")
				return
			}
			fmt.Printf("Consumer: Received %d\n", value)
		}
	}()
	
	time.Sleep(1 * time.Second)
}

// Example 5: Select Statement
func selectStatement() {
	fmt.Println("\n5. Select Statement")
	fmt.Println("===================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Start goroutines
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	// Use select to receive from either channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

// Example 6: Pipeline Pattern
func pipelinePattern() {
	fmt.Println("\n6. Pipeline Pattern")
	fmt.Println("===================")
	
	// Create pipeline stages
	numbers := make(chan int)
	squares := make(chan int)
	results := make(chan int)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()
	
	// Stage 2: Square numbers
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()
	
	// Stage 3: Add 10 to squares
	go func() {
		defer close(results)
		for s := range squares {
			results <- s + 10
		}
	}()
	
	// Collect results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Example 7: Fan-Out Pattern
func fanOutPattern() {
	fmt.Println("\n7. Fan-Out Pattern")
	fmt.Println("==================")
	
	input := make(chan int)
	outputs := make([]chan int, 3)
	
	// Create output channels
	for i := range outputs {
		outputs[i] = make(chan int)
	}
	
	// Fan-out goroutine
	go func() {
		defer func() {
			for _, out := range outputs {
				close(out)
			}
		}()
		
		for n := range input {
			for _, out := range outputs {
				out <- n
			}
		}
	}()
	
	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()
	
	// Collect from all outputs
	var wg sync.WaitGroup
	for i, out := range outputs {
		wg.Add(1)
		go func(id int, ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				fmt.Printf("Worker %d: Received %d\n", id, value)
			}
		}(i, out)
	}
	
	wg.Wait()
}

// Example 8: Fan-In Pattern
func fanInPattern() {
	fmt.Println("\n8. Fan-In Pattern")
	fmt.Println("=================")
	
	// Create input channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// Fan-in goroutine
	output := fanInMain(ch1, ch2, ch3)
	
	// Send data to inputs
	go func() {
		defer close(ch1)
		for j := 1; j <= 3; j++ {
			ch1 <- 10 + j
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch2)
		for j := 1; j <= 3; j++ {
			ch2 <- 20 + j
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch3)
		for j := 1; j <= 3; j++ {
			ch3 <- 30 + j
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Collect from output
	fmt.Println("Fan-in results:")
	for value := range output {
		fmt.Printf("Received: %d\n", value)
	}
}

func fanInMain(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		var wg sync.WaitGroup
		
		for _, input := range inputs {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for value := range ch {
					output <- value
				}
			}(input)
		}
		
		wg.Wait()
	}()
	
	return output
}

// Example 9: Channel Timeout
func channelTimeout() {
	fmt.Println("\n9. Channel Timeout")
	fmt.Println("==================")
	
	ch := make(chan string)
	
	// Send a message after delay
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Delayed message"
	}()
	
	// Wait for message with timeout
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
}

// Example 10: Channel Performance
func channelPerformance() {
	fmt.Println("\n10. Channel Performance")
	fmt.Println("=======================")
	
	// Test unbuffered channel performance
	start := time.Now()
	ch := make(chan int)
	
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			ch <- i
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			<-ch
		}
	}()
	
	wg.Wait()
	unbufferedTime := time.Since(start)
	
	// Test buffered channel performance
	start = time.Now()
	bufferedCh := make(chan int, 1000)
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			bufferedCh <- i
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			<-bufferedCh
		}
	}()
	
	wg.Wait()
	bufferedTime := time.Since(start)
	
	fmt.Printf("Unbuffered channel: %v\n", unbufferedTime)
	fmt.Printf("Buffered channel: %v\n", bufferedTime)
	fmt.Printf("Buffered is %.2fx faster\n", float64(unbufferedTime)/float64(bufferedTime))
}

// Example 11: Common Pitfalls
func commonPitfalls() {
	fmt.Println("\n11. Common Pitfalls")
	fmt.Println("===================")
	
	// Pitfall 1: Sending to closed channel (commented out to avoid panic)
	fmt.Println("Pitfall 1: Sending to closed channel")
	fmt.Println("// ch := make(chan int)")
	fmt.Println("// close(ch)")
	fmt.Println("// ch <- 1  // This would panic!")
	fmt.Println("// Use select to avoid panic")
	
	// Pitfall 2: Receiving from closed channel
	fmt.Println("\nPitfall 2: Receiving from closed channel")
	ch := make(chan int)
	close(ch)
	value, ok := <-ch
	fmt.Printf("Value: %d, OK: %t (channel is closed)\n", value, ok)
	
	// Pitfall 3: Nil channel operations (commented out to avoid blocking)
	fmt.Println("\nPitfall 3: Nil channel operations")
	fmt.Println("// var ch chan int")
	fmt.Println("// ch <- 1  // This would block forever!")
	fmt.Println("// Always initialize channels with make()")
	
	// Pitfall 4: Deadlock with unbuffered channels
	fmt.Println("\nPitfall 4: Deadlock with unbuffered channels")
	fmt.Println("// ch := make(chan int)")
	fmt.Println("// ch <- 1  // This would block waiting for receiver")
	fmt.Println("// value := <-ch  // Never reached")
	fmt.Println("// Use goroutines or buffered channels")
}

// Utility function to show channel info
func showChannelInfo() {
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
