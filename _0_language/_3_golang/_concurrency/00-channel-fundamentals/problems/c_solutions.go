package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ============================================================================
// BASIC SELECT (1-10)
// ============================================================================

// Problem 1: Simple Select
func problem1() {
	fmt.Println("\n=== Problem 1: Simple Select ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	}
}

// Problem 2: Select with Two Channels
func problem2() {
	fmt.Println("\n=== Problem 2: Select with Two Channels ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		ch1 <- "Message 1"
	}()
	
	go func() {
		ch2 <- "Message 2"
	}()
	
	// First message
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg2)
	}
	
	// Second message
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg2)
	}
}

// Problem 3: Select with Send
func problem3() {
	fmt.Println("\n=== Problem 3: Select with Send ===")
	
	ch := make(chan int)
	
	go func() {
		value := <-ch
		fmt.Printf("Received: %d\n", value)
	}()
	
	select {
	case ch <- 42:
		fmt.Println("Sent: 42")
	}
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 4: Select with Send and Receive
func problem4() {
	fmt.Println("\n=== Problem 4: Select with Send and Receive ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	// Start receiver for ch1
	go func() {
		msg := <-ch1
		fmt.Printf("Received: %d\n", msg)
	}()
	
	// Start receiver for ch2
	go func() {
		msg := <-ch2
		fmt.Printf("Received: %d\n", msg)
	}()
	
	// Send to ch1
	select {
	case ch1 <- 100:
		fmt.Println("Sent: 100")
	}
	
	// Send to ch2
	select {
	case ch2 <- 200:
		fmt.Println("Sent: 200")
	}
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 5: Select with Multiple Cases
func problem5() {
	fmt.Println("\n=== Problem 5: Select with Multiple Cases ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	go func() {
		ch1 <- 1
	}()
	
	go func() {
		ch2 <- 2
	}()
	
	go func() {
		ch3 <- 3
	}()
	
	// Handle all cases
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Case 1: %d\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Case 2: %d\n", msg2)
		case msg3 := <-ch3:
			fmt.Printf("Case 3: %d\n", msg3)
		}
	}
}

// Problem 6: Select with Channel Direction
func problem6() {
	fmt.Println("\n=== Problem 6: Select with Channel Direction ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Send-only function
	sendOnly := func(ch chan<- string) {
		select {
		case ch <- "Hello":
			fmt.Println("Sent: Hello")
		}
	}
	
	// Receive-only function
	receiveOnly := func(ch <-chan string) {
		select {
		case msg := <-ch:
			fmt.Printf("Received: %s\n", msg)
		}
	}
	
	go sendOnly(ch1)
	go receiveOnly(ch2)
	
	// Send to ch2
	ch2 <- "World"
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 7: Select with Channel Closing
func problem7() {
	fmt.Println("\n=== Problem 7: Select with Channel Closing ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 42
		close(ch)
	}()
	
	select {
	case msg, ok := <-ch:
		if ok {
			fmt.Printf("Received: %d\n", msg)
		} else {
			fmt.Println("Channel closed")
		}
	}
}

// Problem 8: Select with Range
func problem8() {
	fmt.Println("\n=== Problem 8: Select with Range ===")
	
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received: %d\n", msg)
		}
	}
}

// Problem 9: Select with Goroutine
func problem9() {
	fmt.Println("\n=== Problem 9: Select with Goroutine ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"
	}()
	
	go func() {
		select {
		case msg := <-ch:
			fmt.Printf("Goroutine: Received %s\n", msg)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 10: Select with Function Call
func problem10() {
	fmt.Println("\n=== Problem 10: Select with Function Call ===")
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	
	go func() {
		ch1 <- true
	}()
	
	go func() {
		ch2 <- true
	}()
	
	select {
	case <-ch1:
		fmt.Println("Function 1 called")
	case <-ch2:
		fmt.Println("Function 2 called")
	}
}

// ============================================================================
// MULTIPLE CHANNELS (11-20)
// ============================================================================

// Problem 11: Select with Three Channels
func problem11() {
	fmt.Println("\n=== Problem 11: Select with Three Channels ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	go func() {
		ch1 <- 1
	}()
	
	go func() {
		ch2 <- 2
	}()
	
	go func() {
		ch3 <- 3
	}()
	
	// Handle all three channels
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Channel 1: %d\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Channel 2: %d\n", msg2)
		case msg3 := <-ch3:
			fmt.Printf("Channel 3: %d\n", msg3)
		}
	}
}

// Problem 12: Select with Channel Priority
func problem12() {
	fmt.Println("\n=== Problem 12: Select with Channel Priority ===")
	
	highPriority := make(chan int)
	lowPriority := make(chan int)
	
	go func() {
		lowPriority <- 200
	}()
	
	go func() {
		highPriority <- 100
	}()
	
	// High priority first
	select {
	case msg := <-highPriority:
		fmt.Printf("High priority: %d\n", msg)
	case msg := <-lowPriority:
		fmt.Printf("Low priority: %d\n", msg)
	}
}

// Problem 13: Select with Channel Types
func problem13() {
	fmt.Println("\n=== Problem 13: Select with Channel Types ===")
	
	intCh := make(chan int)
	stringCh := make(chan string)
	boolCh := make(chan bool)
	
	go func() {
		intCh <- 42
	}()
	
	go func() {
		stringCh <- "Hello"
	}()
	
	go func() {
		boolCh <- true
	}()
	
	// Handle different types
	for i := 0; i < 3; i++ {
		select {
		case msg := <-intCh:
			fmt.Printf("Int: %d\n", msg)
		case msg := <-stringCh:
			fmt.Printf("String: %s\n", msg)
		case msg := <-boolCh:
			fmt.Printf("Bool: %t\n", msg)
		}
	}
}

// Problem 14: Select with Channel Arrays
func problem14() {
	fmt.Println("\n=== Problem 14: Select with Channel Arrays ===")
	
	channels := make([]chan int, 3)
	for i := range channels {
		channels[i] = make(chan int)
	}
	
	// Send to channels
	for i := range channels {
		go func(id int) {
			channels[id] <- id
		}(i)
	}
	
	// Receive from channels
	for i := 0; i < 3; i++ {
		select {
		case msg := <-channels[0]:
			fmt.Printf("Channel 0: %d\n", msg)
		case msg := <-channels[1]:
			fmt.Printf("Channel 1: %d\n", msg)
		case msg := <-channels[2]:
			fmt.Printf("Channel 2: %d\n", msg)
		}
	}
}

// Problem 15: Select with Channel Maps
func problem15() {
	fmt.Println("\n=== Problem 15: Select with Channel Maps ===")
	
	channels := make(map[string]chan string)
	channels["A"] = make(chan string)
	channels["B"] = make(chan string)
	channels["C"] = make(chan string)
	
	// Send to channels
	for key, ch := range channels {
		go func(k string, c chan string) {
			c <- k
		}(key, ch)
	}
	
	// Receive from channels
	for i := 0; i < 3; i++ {
		select {
		case msg := <-channels["A"]:
			fmt.Printf("Channel A: %s\n", msg)
		case msg := <-channels["B"]:
			fmt.Printf("Channel B: %s\n", msg)
		case msg := <-channels["C"]:
			fmt.Printf("Channel C: %s\n", msg)
		}
	}
}

// Problem 16: Select with Channel Slices
func problem16() {
	fmt.Println("\n=== Problem 16: Select with Channel Slices ===")
	
	channels := make([]chan int, 3)
	for i := range channels {
		channels[i] = make(chan int)
	}
	
	// Send to channels
	for i := range channels {
		go func(id int) {
			channels[id] <- id
		}(i)
	}
	
	// Receive from channels
	for i := 0; i < 3; i++ {
		select {
		case msg := <-channels[0]:
			fmt.Printf("Channel 0: %d\n", msg)
		case msg := <-channels[1]:
			fmt.Printf("Channel 1: %d\n", msg)
		case msg := <-channels[2]:
			fmt.Printf("Channel 2: %d\n", msg)
		}
	}
}

// Problem 17: Select with Channel Structs
func problem17() {
	fmt.Println("\n=== Problem 17: Select with Channel Structs ===")
	
	type Person struct {
		Name string
		Age  int
	}
	
	ch1 := make(chan Person)
	ch2 := make(chan Person)
	
	go func() {
		ch1 <- Person{Name: "John", Age: 30}
	}()
	
	go func() {
		ch2 <- Person{Name: "Jane", Age: 25}
	}()
	
	// Handle structs
	for i := 0; i < 2; i++ {
		select {
		case person := <-ch1:
			fmt.Printf("Person: %s, Age: %d\n", person.Name, person.Age)
		case person := <-ch2:
			fmt.Printf("Person: %s, Age: %d\n", person.Name, person.Age)
		}
	}
}

// Problem 18: Select with Channel Interfaces
func problem18() {
	fmt.Println("\n=== Problem 18: Select with Channel Interfaces ===")
	
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	
	go func() {
		ch1 <- 42
	}()
	
	go func() {
		ch2 <- "Hello"
	}()
	
	go func() {
		ch3 <- true
	}()
	
	// Handle interfaces
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Interface: %v\n", msg)
		case msg := <-ch2:
			fmt.Printf("Interface: %v\n", msg)
		case msg := <-ch3:
			fmt.Printf("Interface: %v\n", msg)
		}
	}
}

// Problem 19: Select with Channel Pointers
func problem19() {
	fmt.Println("\n=== Problem 19: Select with Channel Pointers ===")
	
	ch1 := make(chan *int)
	ch2 := make(chan *int)
	
	go func() {
		value := 100
		ch1 <- &value
	}()
	
	go func() {
		value := 200
		ch2 <- &value
	}()
	
	// Handle pointers
	for i := 0; i < 2; i++ {
		select {
		case ptr := <-ch1:
			fmt.Printf("Pointer: %d\n", *ptr)
		case ptr := <-ch2:
			fmt.Printf("Pointer: %d\n", *ptr)
		}
	}
}

// Problem 20: Select with Channel Functions
func problem20() {
	fmt.Println("\n=== Problem 20: Select with Channel Functions ===")
	
	ch1 := make(chan func() int)
	ch2 := make(chan func() int)
	
	go func() {
		ch1 <- func() int { return 42 }
	}()
	
	go func() {
		ch2 <- func() int { return 84 }
	}()
	
	// Handle functions
	for i := 0; i < 2; i++ {
		select {
		case fn := <-ch1:
			fmt.Printf("Function result: %d\n", fn())
		case fn := <-ch2:
			fmt.Printf("Function result: %d\n", fn())
		}
	}
}

// ============================================================================
// NON-BLOCKING OPERATIONS (21-30)
// ============================================================================

// Problem 21: Select with Default
func problem21() {
	fmt.Println("\n=== Problem 21: Select with Default ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No data available")
	}
}

// Problem 22: Non-blocking Send
func problem22() {
	fmt.Println("\n=== Problem 22: Non-blocking Send ===")
	
	ch := make(chan int)
	
	select {
	case ch <- 42:
		fmt.Println("Sent: 42")
	default:
		fmt.Println("Send would block")
	}
}

// Problem 23: Non-blocking Receive
func problem23() {
	fmt.Println("\n=== Problem 23: Non-blocking Receive ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No data available")
	}
}

// Problem 24: Non-blocking Multiple Channels
func problem24() {
	fmt.Println("\n=== Problem 24: Non-blocking Multiple Channels ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	select {
	case msg := <-ch1:
		fmt.Printf("Channel 1: %s\n", msg)
	default:
		fmt.Println("Channel 1: No data")
	}
	
	select {
	case msg := <-ch2:
		fmt.Printf("Channel 2: %s\n", msg)
	default:
		fmt.Println("Channel 2: No data")
	}
}

// Problem 25: Non-blocking with Timeout
func problem25() {
	fmt.Println("\n=== Problem 25: Non-blocking with Timeout ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("No data available")
	}
}

// Problem 26: Non-blocking with Error Handling
func problem26() {
	fmt.Println("\n=== Problem 26: Non-blocking with Error Handling ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Error: No data available")
	}
}

// Problem 27: Non-blocking with Retry
func problem27() {
	fmt.Println("\n=== Problem 27: Non-blocking with Retry ===")
	
	ch := make(chan string)
	
	for i := 1; i <= 3; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("Received: %s\n", msg)
			return
		default:
			fmt.Printf("Retry: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Problem 28: Non-blocking with Fallback
func problem28() {
	fmt.Println("\n=== Problem 28: Non-blocking with Fallback ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Fallback: Using default value")
	}
}

// Problem 29: Non-blocking with Circuit Breaker
func problem29() {
	fmt.Println("\n=== Problem 29: Non-blocking with Circuit Breaker ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Circuit: Open")
	}
}

// Problem 30: Non-blocking with Load Balancer
func problem30() {
	fmt.Println("\n=== Problem 30: Non-blocking with Load Balancer ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	select {
	case msg := <-ch1:
		fmt.Printf("Server 1: %s\n", msg)
	default:
		fmt.Println("Server 1: Available")
	}
	
	select {
	case msg := <-ch2:
		fmt.Printf("Server 2: %s\n", msg)
	default:
		fmt.Println("Server 2: Available")
	}
}

// ============================================================================
// TIMEOUT PATTERNS (31-40)
// ============================================================================

// Problem 31: Select with Timeout
func problem31() {
	fmt.Println("\n=== Problem 31: Select with Timeout ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timed out")
	}
}

// Problem 32: Select with Multiple Timeouts
func problem32() {
	fmt.Println("\n=== Problem 32: Select with Multiple Timeouts ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout 1: 100ms")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout 2: 200ms")
	}
}

// Problem 33: Select with Ticker
func problem33() {
	fmt.Println("\n=== Problem 33: Select with Ticker ===")
	
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Tick: %d\n", count)
			if count >= 3 {
				return
			}
		}
	}
}

// Problem 34: Select with Timer
func problem34() {
	fmt.Println("\n=== Problem 34: Select with Timer ===")
	
	timer := time.NewTimer(100 * time.Millisecond)
	defer timer.Stop()
	
	count := 0
	for {
		select {
		case <-timer.C:
			count++
			fmt.Printf("Timer: %d\n", count)
			if count >= 3 {
				return
			}
			timer.Reset(100 * time.Millisecond)
		}
	}
}

// Problem 35: Select with Context Timeout
func problem35() {
	fmt.Println("\n=== Problem 35: Select with Context Timeout ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	}
}

// Problem 36: Select with Deadline
func problem36() {
	fmt.Println("\n=== Problem 36: Select with Deadline ===")
	
	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-ctx.Done():
		fmt.Println("Deadline exceeded")
	}
}

// Problem 37: Select with Timeout and Default
func problem37() {
	fmt.Println("\n=== Problem 37: Select with Timeout and Default ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("No data available")
	default:
		fmt.Println("No data available")
	}
}

// Problem 38: Select with Timeout and Retry
func problem38() {
	fmt.Println("\n=== Problem 38: Select with Timeout and Retry ===")
	
	ch := make(chan string)
	
	for i := 1; i <= 3; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("Received: %s\n", msg)
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("Retry: %d\n", i)
		}
	}
}

// Problem 39: Select with Timeout and Fallback
func problem39() {
	fmt.Println("\n=== Problem 39: Select with Timeout and Fallback ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Fallback: Using default value")
	}
}

// Problem 40: Select with Timeout and Error
func problem40() {
	fmt.Println("\n=== Problem 40: Select with Timeout and Error ===")
	
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Error: Operation timed out")
	}
}

// ============================================================================
// ADVANCED PATTERNS (41-50)
// ============================================================================

// Problem 41: Select with Nil Channels
func problem41() {
	fmt.Println("\n=== Problem 41: Select with Nil Channels ===")
	
	var nilCh chan string
	realCh := make(chan string)
	
	go func() {
		realCh <- "Hello"
	}()
	
	select {
	case msg := <-nilCh:
		fmt.Printf("Received from nil: %s\n", msg)
	case msg := <-realCh:
		fmt.Printf("Received from real: %s\n", msg)
	default:
		fmt.Println("Only non-nil channels are considered")
	}
}

// Problem 42: Select with Channel Closing
func problem42() {
	fmt.Println("\n=== Problem 42: Select with Channel Closing ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"
		close(ch)
	}()
	
	select {
	case msg, ok := <-ch:
		if ok {
			fmt.Printf("Received: %s\n", msg)
		} else {
			fmt.Println("Channel closed")
		}
	}
}

// Problem 43: Select with Channel State
func problem43() {
	fmt.Println("\n=== Problem 43: Select with Channel State ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"
		close(ch)
	}()
	
	select {
	case msg, ok := <-ch:
		if ok {
			fmt.Printf("Channel open: %t\n", ok)
			fmt.Printf("Value: %s\n", msg)
		} else {
			fmt.Printf("Channel closed: %t\n", ok)
		}
	}
}

// Problem 44: Select with Channel Capacity
func problem44() {
	fmt.Println("\n=== Problem 44: Select with Channel Capacity ===")
	
	ch := make(chan int, 5)
	
	select {
	case ch <- 1:
		fmt.Printf("Channel capacity: %d\n", cap(ch))
	}
}

// Problem 45: Select with Channel Length
func problem45() {
	fmt.Println("\n=== Problem 45: Select with Channel Length ===")
	
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	
	select {
	case ch <- 4:
		fmt.Printf("Channel length: %d\n", len(ch))
	}
}

// Problem 46: Select with Channel Comparison
func problem46() {
	fmt.Println("\n=== Problem 46: Select with Channel Comparison ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	select {
	case ch1 <- 1:
		fmt.Printf("Channels are equal: %t\n", ch1 == ch2)
	}
}

// Problem 47: Select with Channel Assignment
func problem47() {
	fmt.Println("\n=== Problem 47: Select with Channel Assignment ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	select {
	case ch1 <- 1:
		ch2 = ch1
		fmt.Println("Channel assigned: true")
	}
	
	// Use ch2 to avoid unused variable warning
	_ = ch2
}

// Problem 48: Select with Channel Range
func problem48() {
	fmt.Println("\n=== Problem 48: Select with Channel Range ===")
	
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	
	select {
	case ch <- 5:
		fmt.Printf("Channel range: 0 to %d\n", cap(ch))
	}
}

// Problem 49: Select with Channel Iteration
func problem49() {
	fmt.Println("\n=== Problem 49: Select with Channel Iteration ===")
	
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	
	count := 0
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			count++
			fmt.Printf("Iteration: %d\n", msg)
		}
	}
}

// Problem 50: Select with Channel Composition
func problem50() {
	fmt.Println("\n=== Problem 50: Select with Channel Composition ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		ch1 <- 42
	}()
	
	go func() {
		ch2 <- 84
	}()
	
	select {
	case msg1 := <-ch1:
		select {
		case msg2 := <-ch2:
			fmt.Printf("Composition: %d\n", msg1+msg2)
		}
	}
}

// ============================================================================
// MAIN FUNCTION
// ============================================================================

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	arg := os.Args[1]
	
	if arg == "all" {
		runAllProblems()
		return
	}
	
	if arg == "basic" {
		runBasicProblems()
		return
	}
	
	if arg == "multiple" {
		runMultipleProblems()
		return
	}
	
	if arg == "non-blocking" {
		runNonBlockingProblems()
		return
	}
	
	if arg == "timeout" {
		runTimeoutProblems()
		return
	}
	
	if arg == "advanced" {
		runAdvancedProblems()
		return
	}
	
	// Try to parse as problem number
	problemNum, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid argument: %s\n", arg)
		showUsage()
		return
	}
	
	if problemNum < 1 || problemNum > 50 {
		fmt.Printf("Problem number must be between 1 and 50, got: %d\n", problemNum)
		showUsage()
		return
	}
	
	runProblem(problemNum)
}

func showUsage() {
	fmt.Println("🔀 Go Select Statement Problems")
	fmt.Println("===============================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run select_problems.go <problem_number>  # Run specific problem (1-50)")
	fmt.Println("  go run select_problems.go basic            # Run basic problems (1-10)")
	fmt.Println("  go run select_problems.go multiple         # Run multiple channel problems (11-20)")
	fmt.Println("  go run select_problems.go non-blocking     # Run non-blocking problems (21-30)")
	fmt.Println("  go run select_problems.go timeout          # Run timeout problems (31-40)")
	fmt.Println("  go run select_problems.go advanced         # Run advanced problems (41-50)")
	fmt.Println("  go run select_problems.go all              # Run all problems")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run select_problems.go 1                # Run problem 1")
	fmt.Println("  go run select_problems.go 25                # Run problem 25")
	fmt.Println("  go run select_problems.go basic             # Run basic problems")
	fmt.Println("  go run select_problems.go all               # Run all 50 problems")
	fmt.Println()
	fmt.Println("Problem Categories:")
	fmt.Println("  Basic (1-10):        Simple select statements")
	fmt.Println("  Multiple (11-20):    Multiple channels, multiplexing")
	fmt.Println("  Non-blocking (21-30): Default cases, non-blocking operations")
	fmt.Println("  Timeout (31-40):     Timeout handling, time.After, time.Tick")
	fmt.Println("  Advanced (41-50):    Complex patterns, real-world scenarios")
}

func runAllProblems() {
	fmt.Println("🚀 Running All 50 Select Problems")
	fmt.Println("=================================")
	
	problems := getAllProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n🎉 All 50 problems completed!")
}

func runBasicProblems() {
	fmt.Println("🟢 Running Basic Select Problems (1-10)")
	fmt.Println("=======================================")
	
	problems := getBasicProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n✅ Basic problems completed!")
}

func runMultipleProblems() {
	fmt.Println("🟡 Running Multiple Channel Problems (11-20)")
	fmt.Println("===========================================")
	
	problems := getMultipleProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+11)
		problem()
	}
	
	fmt.Println("\n✅ Multiple channel problems completed!")
}

func runNonBlockingProblems() {
	fmt.Println("🟠 Running Non-blocking Problems (21-30)")
	fmt.Println("=======================================")
	
	problems := getNonBlockingProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+21)
		problem()
	}
	
	fmt.Println("\n✅ Non-blocking problems completed!")
}

func runTimeoutProblems() {
	fmt.Println("🔴 Running Timeout Problems (31-40)")
	fmt.Println("==================================")
	
	problems := getTimeoutProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+31)
		problem()
	}
	
	fmt.Println("\n✅ Timeout problems completed!")
}

func runAdvancedProblems() {
	fmt.Println("🟣 Running Advanced Problems (41-50)")
	fmt.Println("===================================")
	
	problems := getAdvancedProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+41)
		problem()
	}
	
	fmt.Println("\n✅ Advanced problems completed!")
}

func runProblem(problemNum int) {
	fmt.Printf("🔀 Running Problem %d\n", problemNum)
	fmt.Println("==================")
	
	problems := getAllProblems()
	
	if problemNum < 1 || problemNum > len(problems) {
		fmt.Printf("Problem %d not found\n", problemNum)
		return
	}
	
	problems[problemNum-1]()
	fmt.Printf("\n✅ Problem %d completed!\n", problemNum)
}

func getAllProblems() []func() {
	return []func(){
		problem1, problem2, problem3, problem4, problem5,
		problem6, problem7, problem8, problem9, problem10,
		problem11, problem12, problem13, problem14, problem15,
		problem16, problem17, problem18, problem19, problem20,
		problem21, problem22, problem23, problem24, problem25,
		problem26, problem27, problem28, problem29, problem30,
		problem31, problem32, problem33, problem34, problem35,
		problem36, problem37, problem38, problem39, problem40,
		problem41, problem42, problem43, problem44, problem45,
		problem46, problem47, problem48, problem49, problem50,
	}
}

func getBasicProblems() []func() {
	return []func(){
		problem1, problem2, problem3, problem4, problem5,
		problem6, problem7, problem8, problem9, problem10,
	}
}

func getMultipleProblems() []func() {
	return []func(){
		problem11, problem12, problem13, problem14, problem15,
		problem16, problem17, problem18, problem19, problem20,
	}
}

func getNonBlockingProblems() []func() {
	return []func(){
		problem21, problem22, problem23, problem24, problem25,
		problem26, problem27, problem28, problem29, problem30,
	}
}

func getTimeoutProblems() []func() {
	return []func(){
		problem31, problem32, problem33, problem34, problem35,
		problem36, problem37, problem38, problem39, problem40,
	}
}

func getAdvancedProblems() []func() {
	return []func(){
		problem41, problem42, problem43, problem44, problem45,
		problem46, problem47, problem48, problem49, problem50,
	}
}
