package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
	ex "gselect/exercises"
	pat "gselect/patterns"
)

// Example 1: Basic Select Statement
func basicSelect() {
	fmt.Println("1. Basic Select Statement")
	fmt.Println("=========================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Send messages after delays
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
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

// Example 2: Non-blocking Operations
func nonBlockingOperations() {
	fmt.Println("\n2. Non-blocking Operations")
	fmt.Println("==========================")
	
	ch := make(chan string, 1)
	
	// Non-blocking send
	select {
	case ch <- "Hello":
		fmt.Println("Message sent successfully")
	default:
		fmt.Println("Channel is full")
	}
	
	// Non-blocking receive
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available")
	}
	
	// Try to receive from empty channel
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available (channel is empty)")
	}
}

// Example 3: Default Cases
func defaultCases() {
	fmt.Println("\n3. Default Cases")
	fmt.Println("================")
	
	ch := make(chan string)
	
	// Select with default
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message, doing other work...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Other work completed")
	}
	
	// Send a message and try again
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "Delayed message"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Still no message")
	}
}

// Example 4: Timeout Patterns
func timeoutPatterns() {
	fmt.Println("\n4. Timeout Patterns")
	fmt.Println("===================")
	
	ch := make(chan string)
	
	// Timeout with time.After
	fmt.Println("Testing timeout with time.After:")
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
	
	// Send a message after timeout
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Late message"
	}()
	
	// Multiple timeouts
	fmt.Println("\nTesting multiple timeouts:")
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Short timeout")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Long timeout")
	}
}

// Example 5: Priority Handling
func priorityHandling() {
	fmt.Println("\n5. Priority Handling")
	fmt.Println("===================")
	
	highPriority := make(chan string, 10)
	normalPriority := make(chan string, 10)
	lowPriority := make(chan string, 10)
	
	// Send messages to different priority channels
	go func() {
		time.Sleep(50 * time.Millisecond)
		normalPriority <- "Normal priority message"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		lowPriority <- "Low priority message"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		highPriority <- "High priority message"
	}()
	
	// Handle messages with priority
	for i := 0; i < 3; i++ {
		select {
		case msg := <-highPriority:
			fmt.Printf("HIGH PRIORITY: %s\n", msg)
		case msg := <-normalPriority:
			fmt.Printf("Normal priority: %s\n", msg)
		case msg := <-lowPriority:
			fmt.Printf("Low priority: %s\n", msg)
		}
	}
}

// Example 6: Channel Multiplexing
func channelMultiplexing() {
	fmt.Println("\n6. Channel Multiplexing")
	fmt.Println("=======================")
	
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
					output <- "From input1: " + msg
				}
			case msg, ok := <-input2:
				if !ok {
					input2 = nil // Disable this case
				} else {
					output <- "From input2: " + msg
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
	for msg := range output {
		fmt.Printf("Multiplexed: %s\n", msg)
	}
}

// Example 7: Select with Loops
func selectWithLoops() {
	fmt.Println("\n7. Select with Loops")
	fmt.Println("====================")
	
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
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return
			}
			fmt.Printf("Received: %s\n", msg)
		case <-quit:
			fmt.Println("Quit signal received, exiting loop")
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout in loop")
		}
	}
}

// Example 8: Select with Ticker
func selectWithTicker() {
	fmt.Println("\n8. Select with Ticker")
	fmt.Println("====================")
	
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
	for i := 0; i < 10; i++ {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received: %s\n", msg)
		case <-ticker.C:
			fmt.Println("Tick!")
		}
	}
}

// Example 9: Select with Context
func selectWithContext() {
	fmt.Println("\n9. Select with Context")
	fmt.Println("======================")
	
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	// Send a message after delay
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Context message"
	}()
	
	// Select with context
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	}
}

// Example 10: Select Performance
func selectPerformance() {
	fmt.Println("\n10. Select Performance")
	fmt.Println("======================")
	
	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)
	
	// Fill channels
	for i := 0; i < 1000; i++ {
		ch1 <- i
		ch2 <- i
	}
	
	// Test select performance
	start := time.Now()
	count := 0
	
	for i := 0; i < 1000; i++ {
		select {
		case <-ch1:
			count++
		case <-ch2:
			count++
		}
	}
	
	duration := time.Since(start)
	fmt.Printf("Processed %d messages in %v\n", count, duration)
	fmt.Printf("Average time per select: %v\n", duration/1000)
}

// Example 11: Select with Error Handling
func selectWithErrorHandling() {
	fmt.Println("\n11. Select with Error Handling")
	fmt.Println("==============================")
	
	ch := make(chan string)
	errCh := make(chan error)
	
	// Send messages and errors
	go func() {
		defer close(ch)
		defer close(errCh)
		
		for i := 1; i <= 3; i++ {
			if i == 2 {
				errCh <- fmt.Errorf("Error at message %d", i)
				continue
			}
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Select with error handling
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received: %s\n", msg)
		case err, ok := <-errCh:
			if !ok {
				fmt.Println("Error channel closed")
				return
			}
			fmt.Printf("Error: %v\n", err)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout")
			return
		}
	}
}

// Example 12: Common Pitfalls
func commonPitfalls() {
	fmt.Println("\n12. Common Pitfalls")
	fmt.Println("===================")
	
	// Pitfall 1: Forgetting default case
	fmt.Println("Pitfall 1: Forgetting default case")
	fmt.Println("// This can block forever if no channels are ready")
	fmt.Println("// select { case msg := <-ch: ... }")
	fmt.Println("// Add default case or timeout")
	
	// Pitfall 2: Not handling channel closing
	fmt.Println("\nPitfall 2: Not handling channel closing")
	ch := make(chan string)
	close(ch)
	
	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Println("Channel is closed (handled properly)")
		} else {
			fmt.Printf("Received: %s\n", msg)
		}
	default:
		fmt.Println("No message available")
	}
	
	// Pitfall 3: Race conditions
	fmt.Println("\nPitfall 3: Race conditions")
	var counter int64
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		ch1 <- 1
	}()
	
	go func() {
		ch2 <- 2
	}()
	
	select {
	case <-ch1:
		atomic.AddInt64(&counter, 1)
	case <-ch2:
		atomic.AddInt64(&counter, 1)
	}
	
	fmt.Printf("Counter: %d (using atomic operations)\n", atomic.LoadInt64(&counter))
}

// Utility function to show select info
func showSelectInfo() {
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

func runBasicExamples() {
	fmt.Println("🚀 Select Statement Mastery Examples")
	fmt.Println("====================================")

	// Example 1: Basic Select Statement
	basicSelect()

	// Example 2: Non-blocking Operations
	nonBlockingOperations()

	// Example 3: Default Cases
	defaultCases()

	// Example 4: Timeout Patterns
	timeoutPatterns()

	// Example 5: Priority Handling
	priorityHandling()

	// Example 6: Channel Multiplexing
	channelMultiplexing()

	// Example 7: Select with Loops
	selectWithLoops()

	// Example 8: Select with Ticker
	selectWithTicker()

	// Example 9: Select with Context
	selectWithContext()

	// Example 10: Select Performance
	selectPerformance()

	// Example 11: Select with Error Handling
	selectWithErrorHandling()

	// Example 12: Common Pitfalls
	commonPitfalls()
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Select Exercises")
	fmt.Println("===============================")
	
	ex.Exercise1()
	ex.Exercise2()
	ex.Exercise3()
	ex.Exercise4()
	ex.Exercise5()
	ex.Exercise6()
	ex.Exercise7()
	ex.Exercise8()
	ex.Exercise9()
	ex.Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "basic":
		runBasicExamples()
	case "exercises":
		RunAllExercises()
	case "advanced":
		RunAdvancedPatterns()
	case "all":
		runBasicExamples()
		fmt.Println("\n" + "==================================================")
		RunAllExercises()
		fmt.Println("\n" + "==================================================")
		RunAdvancedPatterns()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func showUsage() {
	fmt.Println("🚀 Select Statement Mastery - Usage")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic select examples")
	fmt.Println("  exercises - Run all exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  all       - Run everything")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println("  go run . all")
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Select Patterns")
	fmt.Println("===========================")
	
	// Pattern 1: Event Loop
	fmt.Println("\n1. Select-based Event Loop:")
	eventLoop := pat.NewEventLoop()
	eventLoop.RegisterHandler("test", func(e pat.Event) {
		fmt.Printf("Event received: %s - %v\n", e.Type, e.Data)
	})
	eventLoop.Start()
	
	eventLoop.EmitEvent(pat.Event{Type: "test", Data: "Hello World"})
	eventLoop.SendCommand(pat.Command{Type: "quit"})
	time.Sleep(100 * time.Millisecond)
	
	// Pattern 2: Rate Limiter
	fmt.Println("\n2. Select-based Rate Limiter:")
	rateLimiter := pat.NewSelectRateLimiter(100*time.Millisecond, 3)
	
	for i := 0; i < 10; i++ {
		if rateLimiter.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i)
		}
	}
	
	requests, allowed := rateLimiter.Stats()
	fmt.Printf("Stats: %d requests, %d allowed\n", requests, allowed)
	rateLimiter.Stop()
	
	// Pattern 3: Load Balancer
	fmt.Println("\n3. Select-based Load Balancer:")
	loadBalancer := pat.NewSelectLoadBalancer(3)
	
	for i := 0; i < 10; i++ {
		loadBalancer.Submit(pat.Job{ID: i, Data: fmt.Sprintf("Data %d", i)})
	}
	
	time.Sleep(1 * time.Second)
	loadBalancer.Stop()
	
	// Pattern 4: Circuit Breaker
	fmt.Println("\n4. Select-based Circuit Breaker:")
	circuitBreaker := pat.NewSelectCircuitBreaker(3, 1*time.Second)
	
	for i := 0; i < 5; i++ {
		err := circuitBreaker.Call(func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i)
		}
	}
	circuitBreaker.Close()
	
	// Pattern 5: Message Router
	fmt.Println("\n5. Select-based Message Router:")
	router := pat.NewSelectMessageRouter()
	router.Start()
	
	sub1 := router.Subscribe("topic1")
	sub2 := router.Subscribe("topic1")
	
	go func() {
		for msg := range sub1 {
			fmt.Printf("Subscriber 1: %v\n", msg)
		}
	}()
	
	go func() {
		for msg := range sub2 {
			fmt.Printf("Subscriber 2: %v\n", msg)
		}
	}()
	
	router.Publish("topic1", "Message 1")
	router.Publish("topic1", "Message 2")
	
	time.Sleep(100 * time.Millisecond)
	router.Stop()
	
	// Pattern 6: Priority Pool
	fmt.Println("\n6. Select-based Priority Pool:")
	priorityPool := pat.NewSelectPriorityPool(2)
	
	for i := 0; i < 5; i++ {
		priorities := []string{"high", "normal", "low"}
		priority := priorities[i%3]
		priorityPool.Submit(pat.Job{ID: i, Data: fmt.Sprintf("Data %d", i)}, priority)
	}
	
	time.Sleep(1 * time.Second)
	priorityPool.Stop()
	
	// Pattern 7: Context Manager
	fmt.Println("\n7. Select-based Context Manager:")
	contextManager := pat.NewSelectContextManager()
	
	ctx := contextManager.CreateContext("test")
	if ctx != nil {
		fmt.Println("Context created successfully")
	}
	
	// Test context cancellation
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Context still active")
	}
	
	contextManager.CancelContext("test")
	
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled after explicit cancel")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Context still active")
	}
	
	fmt.Println("\n✅ All advanced patterns completed!")
}