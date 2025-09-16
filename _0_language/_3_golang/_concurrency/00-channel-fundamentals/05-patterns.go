package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. BASIC SEND/RECEIVE PATTERN
// ============================================================================

func basicSendReceivePattern() {
	fmt.Println("\n📤📥 BASIC SEND/RECEIVE PATTERN")
	fmt.Println("=============================")

	// The most fundamental pattern: one goroutine sends, another receives
	
	fmt.Println("\n1.1 Simple send/receive")
	ch := make(chan string)
	
	// Sender goroutine
	go func() {
		fmt.Println("  Sender: Sending 'Hello'...")
		ch <- "Hello"
		fmt.Println("  Sender: 'Hello' sent!")
		
		fmt.Println("  Sender: Sending 'World'...")
		ch <- "World"
		fmt.Println("  Sender: 'World' sent!")
		
		close(ch)  // Close when done
		fmt.Println("  Sender: Channel closed!")
	}()
	
	// Receiver goroutine
	go func() {
		fmt.Println("  Receiver: Waiting for data...")
		for data := range ch {
			fmt.Printf("  Receiver: Got '%s'\n", data)
		}
		fmt.Println("  Receiver: Channel closed, done!")
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 2. WORKER PATTERN
// ============================================================================

func workerPattern() {
	fmt.Println("\n👷 WORKER PATTERN")
	fmt.Println("================")

	// One goroutine sends work, another processes it
	
	fmt.Println("\n2.1 Basic worker pattern")
	workCh := make(chan int, 3)
	doneCh := make(chan bool)
	
	// Worker goroutine
	go func() {
		fmt.Println("  Worker: Starting work...")
		for work := range workCh {
			fmt.Printf("  Worker: Processing work %d...\n", work)
			time.Sleep(200 * time.Millisecond)  // Simulate work
			fmt.Printf("  Worker: Work %d completed!\n", work)
		}
		fmt.Println("  Worker: All work done!")
		doneCh <- true  // Signal completion
	}()
	
	// Main goroutine sends work
	fmt.Println("  Main: Sending work...")
	workCh <- 1
	workCh <- 2
	workCh <- 3
	close(workCh)  // Signal no more work
	
	fmt.Println("  Main: Waiting for worker to finish...")
	<-doneCh  // Wait for worker
	fmt.Println("  Main: Worker finished!")
}

// ============================================================================
// 3. PIPELINE PATTERN
// ============================================================================

func pipelinePattern() {
	fmt.Println("\n🔗 PIPELINE PATTERN")
	fmt.Println("==================")

	// Data flows through multiple stages, each processing the data
	
	fmt.Println("\n3.1 Basic pipeline pattern")
	
	// Stage 1: Generate numbers
	numbers := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("  Stage 1: Generating %d\n", i)
			numbers <- i
		}
		close(numbers)
	}()
	
	// Stage 2: Square numbers
	squares := make(chan int, 3)
	go func() {
		for n := range numbers {
			square := n * n
			fmt.Printf("  Stage 2: Squaring %d = %d\n", n, square)
			squares <- square
		}
		close(squares)
	}()
	
	// Stage 3: Print results
	go func() {
		for square := range squares {
			fmt.Printf("  Stage 3: Result = %d\n", square)
		}
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 4. FAN-OUT PATTERN
// ============================================================================

func fanOutPattern() {
	fmt.Println("\n🌊 FAN-OUT PATTERN")
	fmt.Println("=================")

	// One goroutine sends data to multiple workers
	
	fmt.Println("\n4.1 Basic fan-out pattern")
	input := make(chan int, 5)
	
	// Send data
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("  Producer: Sending %d\n", i)
			input <- i
		}
		close(input)
	}()
	
	// Multiple workers
	worker1 := make(chan int, 2)
	worker2 := make(chan int, 2)
	worker3 := make(chan int, 2)
	
	// Distribute work to workers
	go func() {
		for data := range input {
			fmt.Printf("  Distributor: Sending %d to workers\n", data)
			worker1 <- data
			worker2 <- data
			worker3 <- data
		}
		close(worker1)
		close(worker2)
		close(worker3)
	}()
	
	// Worker 1
	go func() {
		for data := range worker1 {
			fmt.Printf("  Worker 1: Processing %d\n", data)
		}
	}()
	
	// Worker 2
	go func() {
		for data := range worker2 {
			fmt.Printf("  Worker 2: Processing %d\n", data)
		}
	}()
	
	// Worker 3
	go func() {
		for data := range worker3 {
			fmt.Printf("  Worker 3: Processing %d\n", data)
		}
	}()
	
	time.Sleep(1 * time.Second)
}

// ============================================================================
// 5. FAN-IN PATTERN
// ============================================================================

func fanInPattern() {
	fmt.Println("\n🌊 FAN-IN PATTERN")
	fmt.Println("================")

	// Multiple goroutines send data to one receiver
	
	fmt.Println("\n5.1 Basic fan-in pattern")
	output := make(chan string, 6)
	
	// Producer 1
	go func() {
		output <- "Producer 1: Hello"
		output <- "Producer 1: World"
	}()
	
	// Producer 2
	go func() {
		output <- "Producer 2: Go"
		output <- "Producer 2: Channels"
	}()
	
	// Producer 3
	go func() {
		output <- "Producer 3: Concurrency"
		output <- "Producer 3: Patterns"
	}()
	
	// Wait for all producers to finish
	go func() {
		time.Sleep(500 * time.Millisecond)
		close(output)
	}()
	
	// Consumer receives all data
	fmt.Println("  Consumer: Receiving data...")
	for data := range output {
		fmt.Printf("  %s\n", data)
	}
}

// ============================================================================
// 6. TIMEOUT PATTERN
// ============================================================================

func timeoutPattern() {
	fmt.Println("\n⏰ TIMEOUT PATTERN")
	fmt.Println("=================")

	// Use select with time.After for timeouts
	
	fmt.Println("\n6.1 Basic timeout pattern")
	ch := make(chan string)
	
	// Simulate slow operation
	go func() {
		time.Sleep(2 * time.Second)  // This takes 2 seconds
		ch <- "Slow operation completed"
	}()
	
	// Wait with timeout
	fmt.Println("  Main: Waiting for operation with timeout...")
	select {
	case result := <-ch:
		fmt.Printf("  Main: Got result: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("  Main: Operation timed out!")
	}
}

// ============================================================================
// 7. SELECT PATTERN
// ============================================================================

func selectPattern() {
	fmt.Println("\n🎯 SELECT PATTERN")
	fmt.Println("================")

	// Use select to handle multiple channels
	
	fmt.Println("\n7.1 Basic select pattern")
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	// Send data to both channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Data from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Data from channel 2"
	}()
	
	// Select from both channels
	fmt.Println("  Main: Waiting for data from either channel...")
	for i := 0; i < 2; i++ {
		select {
		case data1 := <-ch1:
			fmt.Printf("  Main: Got from ch1: %s\n", data1)
		case data2 := <-ch2:
			fmt.Printf("  Main: Got from ch2: %s\n", data2)
		}
	}
}

// ============================================================================
// 8. SIGNAL PATTERN
// ============================================================================

func signalPattern() {
	fmt.Println("\n📡 SIGNAL PATTERN")
	fmt.Println("================")

	// Use channels to signal between goroutines
	
	fmt.Println("\n8.1 Basic signal pattern")
	start := make(chan bool)
	done := make(chan bool)
	
	// Worker goroutine
	go func() {
		fmt.Println("  Worker: Waiting for start signal...")
		<-start  // Wait for start signal
		fmt.Println("  Worker: Starting work...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Worker: Work completed!")
		done <- true  // Signal completion
	}()
	
	// Main goroutine
	fmt.Println("  Main: Sending start signal...")
	start <- true  // Send start signal
	fmt.Println("  Main: Waiting for completion...")
	<-done  // Wait for completion
	fmt.Println("  Main: Worker finished!")
}

// ============================================================================
// 9. QUIT PATTERN
// ============================================================================

func quitPattern() {
	fmt.Println("\n🚪 QUIT PATTERN")
	fmt.Println("===============")

	// Use a quit channel to signal goroutines to stop
	
	fmt.Println("\n9.1 Basic quit pattern")
	work := make(chan int, 3)
	quit := make(chan bool)
	
	// Worker goroutine
	go func() {
		for {
			select {
			case data := <-work:
				fmt.Printf("  Worker: Processing %d\n", data)
			case <-quit:
				fmt.Println("  Worker: Quit signal received, stopping...")
				return
			}
		}
	}()
	
	// Send some work
	work <- 1
	work <- 2
	work <- 3
	
	// Wait a bit
	time.Sleep(500 * time.Millisecond)
	
	// Send quit signal
	fmt.Println("  Main: Sending quit signal...")
	quit <- true
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// 10. CHANNEL COMBINATIONS
// ============================================================================

func channelCombinations() {
	fmt.Println("\n🔗 CHANNEL COMBINATIONS")
	fmt.Println("======================")

	// Combine multiple patterns for complex behavior
	
	fmt.Println("\n10.1 Pipeline with fan-out and fan-in")
	
	// Stage 1: Generate data
	numbers := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	// Stage 2: Process data (fan-out to multiple workers)
	worker1 := make(chan int, 2)
	worker2 := make(chan int, 2)
	
	go func() {
		for n := range numbers {
			worker1 <- n
			worker2 <- n
		}
		close(worker1)
		close(worker2)
	}()
	
	// Stage 3: Collect results (fan-in)
	results := make(chan int, 10)
	
	go func() {
		for data := range worker1 {
			results <- data * 2
		}
	}()
	
	go func() {
		for data := range worker2 {
			results <- data * 3
		}
	}()
	
	// Wait for all workers to finish
	go func() {
		time.Sleep(1 * time.Second)
		close(results)
	}()
	
	// Stage 4: Print results
	fmt.Println("  Results:")
	for result := range results {
		fmt.Printf("    %d\n", result)
	}
}

// ============================================================================
// EXPORTED FUNCTIONS FOR MAIN
// ============================================================================

func runChannelPatterns() {
	fmt.Println("🎨 GO CHANNELS: PATTERNS")
	fmt.Println("========================")
	
	// Run all channel pattern examples
	basicSendReceivePattern()
	workerPattern()
	pipelinePattern()
	fanOutPattern()
	fanInPattern()
	timeoutPattern()
	selectPattern()
	signalPattern()
	quitPattern()
	channelCombinations()
	
	fmt.Println("\n✅ Channel patterns completed!")
	fmt.Println("\nNext: Run 'go run . pitfalls' to learn about common pitfalls")
}
