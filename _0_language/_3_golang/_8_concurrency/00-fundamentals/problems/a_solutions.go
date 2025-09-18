package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// EASY LEVEL SOLUTIONS (1-15)
// ============================================================================

// Problem 1: Basic Send and Receive
func problem1() {
	fmt.Println("\n=== Problem 1: Basic Send and Receive ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 42
		fmt.Println("Sent: 42")
	}()
	
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

// Problem 2: String Channel
func problem2() {
	fmt.Println("\n=== Problem 2: String Channel ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "Hello World"
	}()
	
	message := <-ch
	fmt.Printf("Message: %s\n", message)
}

// Problem 3: Multiple Values
func problem3() {
	fmt.Println("\n=== Problem 3: Multiple Values ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

// Problem 4: Channel Direction
func problem4() {
	fmt.Println("\n=== Problem 4: Channel Direction ===")
	
	ch := make(chan int)
	
	// Send-only function
	sendOnly := func(ch chan<- int) {
		ch <- 100
		fmt.Println("Sender sent: 100")
	}
	
	// Receive-only function
	receiveOnly := func(ch <-chan int) {
		value := <-ch
		fmt.Printf("Receiver got: %d\n", value)
	}
	
	go sendOnly(ch)
	go receiveOnly(ch)
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 5: Channel as Signal
func problem5() {
	fmt.Println("\n=== Problem 5: Channel as Signal ===")
	
	done := make(chan bool)
	
	go func() {
		fmt.Println("Work started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Work completed")
		done <- true
	}()
	
	<-done
	fmt.Println("Signal received")
}

// Problem 6: Close Channel
func problem6() {
	fmt.Println("\n=== Problem 6: Close Channel ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 42
		close(ch)
		fmt.Println("Channel closed")
	}()
	
	value := <-ch
	fmt.Printf("Received: %d\n", value)
	
	value, ok := <-ch
	fmt.Printf("Received from closed: %d, ok: %t\n", value, ok)
}

// Problem 7: Range Over Channel
func problem7() {
	fmt.Println("\n=== Problem 7: Range Over Channel ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Done receiving")
}

// Problem 8: Buffered Channel
func problem8() {
	fmt.Println("\n=== Problem 8: Buffered Channel ===")
	
	ch := make(chan int, 3)
	
	ch <- 1
	fmt.Println("Sent: 1")
	ch <- 2
	fmt.Println("Sent: 2")
	ch <- 3
	fmt.Println("Sent: 3")
	fmt.Println("All sent without blocking")
}

// Problem 9: Channel Length
func problem9() {
	fmt.Println("\n=== Problem 9: Channel Length ===")
	
	ch := make(chan int, 5)
	fmt.Printf("Channel length: %d\n", len(ch))
	
	ch <- 1
	ch <- 2
	fmt.Printf("After sending 2 items: %d\n", len(ch))
	
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("After sending 5 items: %d\n", len(ch))
}

// Problem 10: Channel Capacity
func problem10() {
	fmt.Println("\n=== Problem 10: Channel Capacity ===")
	
	unbuffered := make(chan int)
	buffered := make(chan int, 5)
	
	fmt.Printf("Unbuffered capacity: %d\n", cap(unbuffered))
	fmt.Printf("Buffered capacity: %d\n", cap(buffered))
}

// Problem 11: Nil Channel Check
func problem11() {
	fmt.Println("\n=== Problem 11: Nil Channel Check ===")
	
	var nilCh chan int
	realCh := make(chan int)
	
	fmt.Printf("Channel is nil: %t\n", nilCh == nil)
	fmt.Printf("Channel is nil: %t\n", realCh == nil)
}

// Problem 12: Simple Select
func problem12() {
	fmt.Println("\n=== Problem 12: Simple Select ===")
	
	ch := make(chan int)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 42
	}()
	
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

// Problem 13: Select with Default
func problem13() {
	fmt.Println("\n=== Problem 13: Select with Default ===")
	
	ch := make(chan int)
	
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("No data available")
	}
}

// Problem 14: Multiple Channels
func problem14() {
	fmt.Println("\n=== Problem 14: Multiple Channels ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		ch1 <- 100
	}()
	
	go func() {
		ch2 <- 200
	}()
	
	select {
	case value := <-ch1:
		fmt.Printf("Received from ch1: %d\n", value)
	case value := <-ch2:
		fmt.Printf("Received from ch2: %d\n", value)
	}
}

// Problem 15: Channel State Check
func problem15() {
	fmt.Println("\n=== Problem 15: Channel State Check ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 42
		close(ch)
	}()
	
	value, ok := <-ch
	fmt.Printf("Channel open: %t\n", ok)
	fmt.Printf("Value: %d\n", value)
	
	value, ok = <-ch
	fmt.Printf("Channel closed: %t\n", ok)
}

// ============================================================================
// MEDIUM LEVEL SOLUTIONS (16-35)
// ============================================================================

// Problem 16: Buffered Channel Blocking
func problem16() {
	fmt.Println("\n=== Problem 16: Buffered Channel Blocking ===")
	
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	
	select {
	case ch <- 3:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Buffer full, send would block")
	}
}

// Problem 17: Select with Timeout
func problem17() {
	fmt.Println("\n=== Problem 17: Select with Timeout ===")
	
	ch := make(chan int)
	
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timed out")
	}
}

// Problem 18: Non-blocking Send
func problem18() {
	fmt.Println("\n=== Problem 18: Non-blocking Send ===")
	
	ch := make(chan int)
	
	select {
	case ch <- 42:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Send would block")
	}
}

// Problem 19: Channel Comparison
func problem19() {
	fmt.Println("\n=== Problem 19: Channel Comparison ===")
	
	unbuffered := make(chan int)
	buffered := make(chan int, 1)
	
	fmt.Println("Unbuffered: Synchronous")
	fmt.Println("Buffered: Asynchronous")
	
	// Demonstrate buffered channel
	buffered <- 1
	fmt.Println("Buffered channel accepted value immediately")
}

// Problem 20: Multiple Senders
func problem20() {
	fmt.Println("\n=== Problem 20: Multiple Senders ===")
	
	ch := make(chan int)
	
	go func() {
		ch <- 1
		fmt.Println("Sender 1: 1")
	}()
	
	go func() {
		ch <- 2
		fmt.Println("Sender 2: 2")
	}()
	
	go func() {
		ch <- 3
		fmt.Println("Sender 3: 3")
	}()
	
	time.Sleep(200 * time.Millisecond)
}

// Problem 21: Multiple Receivers
func problem21() {
	fmt.Println("\n=== Problem 21: Multiple Receivers ===")
	
	ch := make(chan int)
	
	go func() {
		value := <-ch
		fmt.Printf("Receiver 1: %d\n", value)
	}()
	
	go func() {
		value := <-ch
		fmt.Printf("Receiver 2: %d\n", value)
	}()
	
	ch <- 42
	ch <- 42
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 22: Channel as Counter
func problem22() {
	fmt.Println("\n=== Problem 22: Channel as Counter ===")
	
	ch := make(chan int)
	
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	for value := range ch {
		fmt.Printf("Count: %d\n", value)
	}
}

// Problem 23: Select with Multiple Cases
func problem23() {
	fmt.Println("\n=== Problem 23: Select with Multiple Cases ===")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	go func() {
		ch1 <- 100
	}()
	
	go func() {
		ch2 <- 200
	}()
	
	go func() {
		ch3 <- 300
	}()
	
	for i := 0; i < 3; i++ {
		select {
		case value := <-ch1:
			fmt.Printf("Case 1: %d\n", value)
		case value := <-ch2:
			fmt.Printf("Case 2: %d\n", value)
		case value := <-ch3:
			fmt.Printf("Case 3: %d\n", value)
		}
	}
}

// Problem 24: Channel with Struct
func problem24() {
	fmt.Println("\n=== Problem 24: Channel with Struct ===")
	
	type Person struct {
		Name string
		Age  int
	}
	
	ch := make(chan Person)
	
	go func() {
		ch <- Person{Name: "John", Age: 30}
	}()
	
	person := <-ch
	fmt.Printf("Person: %s, Age: %d\n", person.Name, person.Age)
}

// Problem 25: Channel with Slice
func problem25() {
	fmt.Println("\n=== Problem 25: Channel with Slice ===")
	
	ch := make(chan []int)
	
	go func() {
		ch <- []int{1, 2, 3, 4, 5}
	}()
	
	slice := <-ch
	fmt.Printf("Received slice: %v\n", slice)
}

// Problem 26: Channel with Map
func problem26() {
	fmt.Println("\n=== Problem 26: Channel with Map ===")
	
	ch := make(chan map[string]string)
	
	go func() {
		ch <- map[string]string{"hello": "world"}
	}()
	
	m := <-ch
	fmt.Printf("Received map: %v\n", m)
}

// Problem 27: Channel with Interface
func problem27() {
	fmt.Println("\n=== Problem 27: Channel with Interface ===")
	
	ch := make(chan interface{})
	
	go func() {
		ch <- 42
		ch <- "hello"
		ch <- true
		close(ch)
	}()
	
	for value := range ch {
		fmt.Printf("Received: %v\n", value)
	}
}

// Problem 28: Channel with Pointer
func problem28() {
	fmt.Println("\n=== Problem 28: Channel with Pointer ===")
	
	ch := make(chan *int)
	
	go func() {
		value := 100
		ch <- &value
	}()
	
	ptr := <-ch
	fmt.Printf("Value: %d\n", *ptr)
}

// Problem 29: Channel with Function
func problem29() {
	fmt.Println("\n=== Problem 29: Channel with Function ===")
	
	ch := make(chan func() int)
	
	go func() {
		ch <- func() int { return 42 }
	}()
	
	fn := <-ch
	result := fn()
	fmt.Printf("Function result: %d\n", result)
}

// Problem 30: Channel with Channel
func problem30() {
	fmt.Println("\n=== Problem 30: Channel with Channel ===")
	
	ch := make(chan chan int)
	
	go func() {
		innerCh := make(chan int)
		ch <- innerCh
		innerCh <- 42
	}()
	
	innerCh := <-ch
	fmt.Println("Received channel, sending data")
	value := <-innerCh
	fmt.Printf("Data sent through received channel: %d\n", value)
}

// Problem 31: Select with Nil Channels
func problem31() {
	fmt.Println("\n=== Problem 31: Select with Nil Channels ===")
	
	var nilCh chan int
	realCh := make(chan int)
	
	go func() {
		realCh <- 42
	}()
	
	select {
	case value := <-nilCh:
		fmt.Printf("Received from nil: %d\n", value)
	case value := <-realCh:
		fmt.Printf("Received from real: %d\n", value)
	default:
		fmt.Println("Only non-nil channels are considered")
	}
}

// Problem 32: Channel with Context
func problem32() {
	fmt.Println("\n=== Problem 32: Channel with Context ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	ch := make(chan int)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- 42
	}()
	
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
	}
}

// Problem 33: Channel with Error
func problem33() {
	fmt.Println("\n=== Problem 33: Channel with Error ===")
	
	ch := make(chan error)
	
	go func() {
		ch <- fmt.Errorf("something went wrong")
	}()
	
	err := <-ch
	fmt.Printf("Error: %v\n", err)
}

// Problem 34: Channel with Result
func problem34() {
	fmt.Println("\n=== Problem 34: Channel with Result ===")
	
	type Result struct {
		Value int
		Error error
	}
	
	ch := make(chan Result)
	
	go func() {
		ch <- Result{Value: 42, Error: nil}
	}()
	
	result := <-ch
	fmt.Printf("Result: %d, Error: %v\n", result.Value, result.Error)
}

// Problem 35: Channel with Status
func problem35() {
	fmt.Println("\n=== Problem 35: Channel with Status ===")
	
	ch := make(chan string)
	
	go func() {
		ch <- "started"
		time.Sleep(100 * time.Millisecond)
		ch <- "processing"
		time.Sleep(100 * time.Millisecond)
		ch <- "completed"
		close(ch)
	}()
	
	for status := range ch {
		fmt.Printf("Status: %s\n", status)
	}
}

// ============================================================================
// ADVANCED LEVEL SOLUTIONS (36-50)
// ============================================================================

// Problem 36: Channel with Mutex
func problem36() {
	fmt.Println("\n=== Problem 36: Channel with Mutex ===")
	
	var mu sync.Mutex
	counter := 0
	ch := make(chan int, 1000)
	
	// Start workers
	for i := 0; i < 10; i++ {
		go func() {
			for range ch {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	
	// Send work
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Counter: %d\n", counter)
}

// Problem 37: Channel with WaitGroup
func problem37() {
	fmt.Println("\n=== Problem 37: Channel with WaitGroup ===")
	
	var wg sync.WaitGroup
	ch := make(chan int)
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ch <- id
		}(i)
	}
	
	go func() {
		wg.Wait()
		close(ch)
	}()
	
	for value := range ch {
		fmt.Printf("Worker %d completed\n", value)
	}
	fmt.Println("All goroutines completed")
}

// Problem 38: Channel with Atomic
func problem38() {
	fmt.Println("\n=== Problem 38: Channel with Atomic ===")
	
	var counter int64
	ch := make(chan int, 1000)
	
	// Start workers
	for i := 0; i < 10; i++ {
		go func() {
			for range ch {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	
	// Send work
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Atomic counter: %d\n", atomic.LoadInt64(&counter))
}

// Problem 39: Channel with Timer
func problem39() {
	fmt.Println("\n=== Problem 39: Channel with Timer ===")
	
	timer := time.NewTimer(100 * time.Millisecond)
	defer timer.Stop()
	
	count := 0
	for {
		select {
		case <-timer.C:
			count++
			fmt.Printf("Tick: %d\n", count)
			if count >= 3 {
				return
			}
			timer.Reset(100 * time.Millisecond)
		}
	}
}

// Problem 40: Channel with Ticker
func problem40() {
	fmt.Println("\n=== Problem 40: Channel with Ticker ===")
	
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Ticker: %d\n", count)
			if count >= 3 {
				return
			}
		}
	}
}

// Problem 41: Channel with Rate Limiting
func problem41() {
	fmt.Println("\n=== Problem 41: Channel with Rate Limiting ===")
	
	rateLimiter := make(chan time.Time, 2)
	
	// Allow 2 requests per second
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rateLimiter <- time.Now():
			default:
			}
		}
	}()
	
	// Simulate requests
	for i := 1; i <= 4; i++ {
		select {
		case <-rateLimiter:
			fmt.Printf("Request %d: allowed\n", i)
		default:
			fmt.Printf("Request %d: rate limited\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// Problem 42: Channel with Circuit Breaker
func problem42() {
	fmt.Println("\n=== Problem 42: Channel with Circuit Breaker ===")
	
	type CircuitState int
	const (
		Closed CircuitState = iota
		Open
		HalfOpen
	)
	
	state := Closed
	failures := 0
	lastFailTime := time.Time{}
	
	ch := make(chan string)
	
	go func() {
		for i := 0; i < 5; i++ {
			now := time.Now()
			
			if state == Open && now.Sub(lastFailTime) > 1*time.Second {
				state = HalfOpen
				fmt.Println("Circuit: half-open")
			}
			
			if state == Closed || state == HalfOpen {
				// Simulate failure
				if i%2 == 0 {
					failures++
					lastFailTime = now
					if failures >= 2 {
						state = Open
						fmt.Println("Circuit: open")
					}
				} else {
					failures = 0
					state = Closed
					fmt.Println("Circuit: closed")
				}
			} else {
				fmt.Println("Circuit: open")
			}
			
			time.Sleep(500 * time.Millisecond)
		}
	}()
	
	time.Sleep(3 * time.Second)
}

// Problem 43: Channel with Load Balancer
func problem43() {
	fmt.Println("\n=== Problem 43: Channel with Load Balancer ===")
	
	servers := []string{"Server 1", "Server 2", "Server 3"}
	serverCh := make(chan string, len(servers))
	
	// Initialize servers
	for _, server := range servers {
		serverCh <- server
	}
	
	// Load balancer
	go func() {
		for i := 1; i <= 5; i++ {
			server := <-serverCh
			fmt.Printf("Request %d: %s\n", i, server)
			serverCh <- server // Put server back
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	time.Sleep(1 * time.Second)
}

// Problem 44: Channel with Priority Queue
func problem44() {
	fmt.Println("\n=== Problem 44: Channel with Priority Queue ===")
	
	type Task struct {
		Priority int
		Data     string
	}
	
	highPriority := make(chan Task, 10)
	lowPriority := make(chan Task, 10)
	
	// Add tasks
	go func() {
		highPriority <- Task{Priority: 1, Data: "High priority task 1"}
		highPriority <- Task{Priority: 1, Data: "High priority task 2"}
		lowPriority <- Task{Priority: 2, Data: "Low priority task 1"}
		close(highPriority)
		close(lowPriority)
	}()
	
	// Process tasks with priority
	for {
		select {
		case task, ok := <-highPriority:
			if !ok {
				highPriority = nil
			} else {
				fmt.Printf("High priority: %s\n", task.Data)
			}
		case task, ok := <-lowPriority:
			if !ok {
				lowPriority = nil
			} else {
				fmt.Printf("Low priority: %s\n", task.Data)
			}
		}
		
		if highPriority == nil && lowPriority == nil {
			break
		}
	}
}

// Problem 45: Channel with Batch Processing
func problem45() {
	fmt.Println("\n=== Problem 45: Channel with Batch Processing ===")
	
	input := make(chan int, 10)
	batchSize := 3
	
	// Send data
	go func() {
		for i := 1; i <= 7; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Process in batches
	batch := make([]int, 0, batchSize)
	for value := range input {
		batch = append(batch, value)
		if len(batch) == batchSize {
			fmt.Printf("Batch: %v\n", batch)
			batch = batch[:0] // Reset batch
		}
	}
	
	if len(batch) > 0 {
		fmt.Printf("Batch: %v\n", batch)
	}
}

// Problem 46: Channel with Retry Logic
func problem46() {
	fmt.Println("\n=== Problem 46: Channel with Retry Logic ===")
	
	maxRetries := 3
	ch := make(chan bool)
	
	go func() {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Attempt %d: ", attempt)
			
			// Simulate operation
			success := attempt == 3 // Succeed on 3rd attempt
			
			if success {
				fmt.Println("success")
				ch <- true
				return
			} else {
				fmt.Println("failed")
				time.Sleep(100 * time.Millisecond)
			}
		}
		ch <- false
	}()
	
	result := <-ch
	if result {
		fmt.Println("Operation succeeded")
	} else {
		fmt.Println("Operation failed after all retries")
	}
}

// Problem 47: Channel with Exponential Backoff
func problem47() {
	fmt.Println("\n=== Problem 47: Channel with Exponential Backoff ===")
	
	baseDelay := 1 * time.Second
	maxDelay := 8 * time.Second
	ch := make(chan bool)
	
	go func() {
		delay := baseDelay
		for attempt := 1; attempt <= 4; attempt++ {
			fmt.Printf("Retry after: %v\n", delay)
			time.Sleep(delay)
			
			// Double the delay for next attempt
			delay *= 2
			if delay > maxDelay {
				delay = maxDelay
			}
		}
		ch <- true
	}()
	
	<-ch
}

// Problem 48: Channel with Health Check
func problem48() {
	fmt.Println("\n=== Problem 48: Channel with Health Check ===")
	
	healthCh := make(chan bool)
	
	go func() {
		for i := 0; i < 5; i++ {
			// Simulate health check
			healthy := i%3 != 0 // Unhealthy every 3rd check
			healthCh <- healthy
			
			if healthy {
				fmt.Println("Service: healthy")
			} else {
				fmt.Println("Service: unhealthy")
			}
			
			time.Sleep(500 * time.Millisecond)
		}
		close(healthCh)
	}()
	
	for range healthCh {
		// Health check results processed
	}
}

// Problem 49: Channel with Metrics
func problem49() {
	fmt.Println("\n=== Problem 49: Channel with Metrics ===")
	
	type Metrics struct {
		Requests int
		Errors   int
	}
	
	metricsCh := make(chan Metrics)
	
	go func() {
		// Simulate metrics collection
		for i := 0; i < 10; i++ {
			metrics := Metrics{
				Requests: 100 + i*10,
				Errors:   i,
			}
			metricsCh <- metrics
			time.Sleep(200 * time.Millisecond)
		}
		close(metricsCh)
	}()
	
	for metrics := range metricsCh {
		fmt.Printf("Metrics: requests=%d, errors=%d\n", metrics.Requests, metrics.Errors)
	}
}

// Problem 50: Channel with Configuration
func problem50() {
	fmt.Println("\n=== Problem 50: Channel with Configuration ===")
	
	type Config struct {
		Timeout  time.Duration
		Retries  int
	}
	
	configCh := make(chan Config)
	
	go func() {
		configs := []Config{
			{Timeout: 5 * time.Second, Retries: 3},
			{Timeout: 10 * time.Second, Retries: 5},
			{Timeout: 3 * time.Second, Retries: 2},
		}
		
		for _, config := range configs {
			configCh <- config
			time.Sleep(500 * time.Millisecond)
		}
		close(configCh)
	}()
	
	for config := range configCh {
		fmt.Printf("Config updated: timeout=%v, retries=%d\n", config.Timeout, config.Retries)
	}
}

