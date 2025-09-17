package main

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 1: Implement Channel Ownership Pattern
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Channel Ownership Pattern")
	fmt.Println("==============================================")
	
	// TODO: Implement a channel owner that creates, manages, and closes a channel
	// 1. Create a function that returns a receive-only channel
	// 2. The function should be responsible for closing the channel
	// 3. Send some data through the channel
	// 4. Consume the data from the channel
	
	ch := channelOwner()
	
	for value := range ch {
		fmt.Printf("  Exercise 1: Received %d\n", value)
	}
	
	fmt.Println("Exercise 1 completed")
}

func channelOwner() <-chan int {
	ch := make(chan int, 5)
	
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	
	return ch
}

// Exercise 2: Implement Channel Factory Pattern
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Channel Factory Pattern")
	fmt.Println("===========================================")
	
	// TODO: Implement a channel factory function
	// 1. Create a function that creates channels with specified capacity
	// 2. Create a function that creates initialized channels
	// 3. Test both factory functions
	
	// Basic factory
	ch1 := createChannel(3)
	
	// Initialized factory
	ch2 := createInitializedChannel([]int{10, 20, 30})
	
	// Test basic factory
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()
	
	for value := range ch1 {
		fmt.Printf("  Exercise 2: Basic factory %d\n", value)
	}
	
	// Test initialized factory
	for value := range ch2 {
		fmt.Printf("  Exercise 2: Initialized factory %d\n", value)
	}
	
	fmt.Println("Exercise 2 completed")
}

func createChannel(capacity int) chan int {
	return make(chan int, capacity)
}

func createInitializedChannel(values []int) <-chan int {
	ch := make(chan int, len(values))
	
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	
	return ch
}

// Exercise 3: Implement Channel Wrapper Pattern
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Channel Wrapper Pattern")
	fmt.Println("===========================================")
	
	// TODO: Implement a channel wrapper with additional functionality
	// 1. Create a wrapper struct that adds safety features
	// 2. Implement Send, Receive, and Close methods
	// 3. Add thread safety and error handling
	// 4. Test the wrapper
	
	wrapper := NewChannelWrapper(3)
	
	// Test sending
	for i := 0; i < 5; i++ {
		if success := wrapper.Send(i); success {
			fmt.Printf("  Exercise 3: Sent %d\n", i)
		} else {
			fmt.Printf("  Exercise 3: Failed to send %d\n", i)
		}
	}
	
	// Test receiving
	for i := 0; i < 5; i++ {
		if value, ok := wrapper.Receive(); ok {
			fmt.Printf("  Exercise 3: Received %d\n", value)
		}
	}
	
	wrapper.Close()
	fmt.Println("Exercise 3 completed")
}

type ChannelWrapper struct {
	ch     chan int
	closed bool
	mu     sync.Mutex
}

func NewChannelWrapper(capacity int) *ChannelWrapper {
	return &ChannelWrapper{
		ch: make(chan int, capacity),
	}
}

func (cw *ChannelWrapper) Send(value int) bool {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	
	if cw.closed {
		return false
	}
	
	select {
	case cw.ch <- value:
		return true
	default:
		return false
	}
}

func (cw *ChannelWrapper) Receive() (int, bool) {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	
	if cw.closed {
		return 0, false
	}
	
	select {
	case value := <-cw.ch:
		return value, true
	default:
		return 0, false
	}
}

func (cw *ChannelWrapper) Close() {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	
	if !cw.closed {
		close(cw.ch)
		cw.closed = true
	}
}

// Exercise 4: Implement Graceful Shutdown Pattern
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Graceful Shutdown Pattern")
	fmt.Println("==============================================")
	
	// TODO: Implement graceful shutdown with channels
	// 1. Create a producer that can be shut down gracefully
	// 2. Create a consumer that can be shut down gracefully
	// 3. Use a done channel to signal shutdown
	// 4. Test the graceful shutdown
	
	ch := make(chan int)
	done := make(chan struct{})
	
	// Producer
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Printf("  Exercise 4: Produced %d\n", i)
			case <-done:
				fmt.Println("  Exercise 4: Producer shutting down")
				return
			}
		}
	}()
	
	// Consumer
	go func() {
		defer close(done)
		for value := range ch {
			fmt.Printf("  Exercise 4: Consumed %d\n", value)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Wait for completion
	<-done
	fmt.Println("Exercise 4 completed")
}

// Exercise 5: Implement Nil Channel Tricks
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Nil Channel Tricks")
	fmt.Println("======================================")
	
	// TODO: Implement nil channel tricks
	// 1. Create two channels
	// 2. Use nil channels to disable select cases
	// 3. Implement dynamic channel management
	// 4. Test the nil channel behavior
	
	ch1 := make(chan int)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 42
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "hello"
	}()
	
	// Disable ch2 after receiving from ch1
	for i := 0; i < 2; i++ {
		select {
		case value := <-ch1:
			fmt.Printf("  Exercise 5: Received from ch1: %d\n", value)
			ch2 = nil // Disable ch2
		case value := <-ch2:
			fmt.Printf("  Exercise 5: Received from ch2: %s\n", value)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("  Exercise 5: Timeout reached")
			break
		}
	}
	
	fmt.Println("Exercise 5 completed")
}

// Exercise 6: Implement Channel Pipeline Pattern
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Channel Pipeline Pattern")
	fmt.Println("=============================================")
	
	// TODO: Implement a channel pipeline
	// 1. Create a 3-stage pipeline: generate -> square -> print
	// 2. Use channels to connect the stages
	// 3. Ensure proper channel closing
	// 4. Test the pipeline
	
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 0; i < 5; i++ {
			numbers <- i
		}
	}()
	
	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()
	
	// Stage 3: Print results
	for square := range squares {
		fmt.Printf("  Exercise 6: Square %d\n", square)
	}
	
	fmt.Println("Exercise 6 completed")
}

// Exercise 7: Implement Channel Fan-Out Pattern
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Channel Fan-Out Pattern")
	fmt.Println("============================================")
	
	// TODO: Implement channel fan-out
	// 1. Create an input channel
	// 2. Create multiple worker goroutines
	// 3. Distribute work among workers
	// 4. Collect results from all workers
	
	input := make(chan int)
	
	// Start multiple workers
	workers := make([]<-chan int, 3)
	for i := 0; i < 3; i++ {
		workers[i] = worker(input, i)
	}
	
	// Send data
	go func() {
		defer close(input)
		for i := 0; i < 6; i++ {
			input <- i
		}
	}()
	
	// Collect results
	var wg sync.WaitGroup
	wg.Add(len(workers))
	
	for i, worker := range workers {
		go func(id int, w <-chan int) {
			defer wg.Done()
			for result := range w {
				fmt.Printf("  Exercise 7: Worker %d result: %d\n", id, result)
			}
		}(i, worker)
	}
	
	wg.Wait()
	fmt.Println("Exercise 7 completed")
}

func worker(input <-chan int, id int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		for value := range input {
			result := value * value
			output <- result
		}
	}()
	
	return output
}

// Exercise 8: Implement Channel Fan-In Pattern
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Channel Fan-In Pattern")
	fmt.Println("===========================================")
	
	// TODO: Implement channel fan-in
	// 1. Create multiple input channels
	// 2. Create a fan-in function that combines them
	// 3. Use WaitGroup to wait for all inputs
	// 4. Test the fan-in pattern
	
	// Create multiple input channels
	inputs := make([]<-chan int, 3)
	for i := 0; i < 3; i++ {
		inputs[i] = generateNumbers(i, 3)
	}
	
	// Fan-in to single output
	output := fanIn(inputs...)
	
	// Process results
	for value := range output {
		fmt.Printf("  Exercise 8: Fan-in result: %d\n", value)
	}
	
	fmt.Println("Exercise 8 completed")
}

func generateNumbers(id, count int) <-chan int {
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- id*100 + i
		}
	}()
	
	return ch
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		
		var wg sync.WaitGroup
		wg.Add(len(inputs))
		
		for _, input := range inputs {
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

// Exercise 9: Implement Error Channel Pattern
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Error Channel Pattern")
	fmt.Println("=========================================")
	
	// TODO: Implement error handling with channels
	// 1. Create data and error channels
	// 2. Send data and errors through appropriate channels
	// 3. Handle both data and errors in the main loop
	// 4. Test error handling
	
	dataCh := make(chan int)
	errorCh := make(chan error)
	
	go func() {
		defer close(dataCh)
		defer close(errorCh)
		
		for i := 0; i < 5; i++ {
			if i == 3 {
				errorCh <- fmt.Errorf("error at value %d", i)
				return
			}
			dataCh <- i
		}
	}()
	
	for {
		select {
		case value, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("  Exercise 9: Received: %d\n", value)
		case err := <-errorCh:
			fmt.Printf("  Exercise 9: Error: %v\n", err)
			return
		}
	}
}

// Exercise 10: Implement Channel Batching Pattern
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Channel Batching Pattern")
	fmt.Println("=============================================")
	
	// TODO: Implement channel batching
	// 1. Create input and output channels
	// 2. Implement batching logic with size and time limits
	// 3. Use select to handle both data and timeouts
	// 4. Test the batching pattern
	
	input := make(chan int)
	output := make(chan []int)
	
	go func() {
		defer close(output)
		
		batch := make([]int, 0, 3)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case value, ok := <-input:
				if !ok {
					if len(batch) > 0 {
						output <- batch
					}
					return
				}
				batch = append(batch, value)
				if len(batch) >= 3 {
					output <- batch
					batch = make([]int, 0, 3)
				}
			case <-ticker.C:
				if len(batch) > 0 {
					output <- batch
					batch = make([]int, 0, 3)
				}
			}
		}
	}()
	
	// Send data
	go func() {
		defer close(input)
		for i := 0; i < 7; i++ {
			input <- i
		}
	}()
	
	// Process batches
	for batch := range output {
		fmt.Printf("  Exercise 10: Batch: %v\n", batch)
	}
	
	fmt.Println("Exercise 10 completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Channel Patterns & Idioms Exercises")
	fmt.Println("=====================================")
	
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
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
