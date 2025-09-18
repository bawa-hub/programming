package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Example 1: Basic Channel Ownership Pattern
func basicChannelOwnership() {
	fmt.Println("\n1. Basic Channel Ownership Pattern")
	fmt.Println("==================================")
	
	// Channel owner creates and manages the channel
	ch := basicChannelOwner()
	
	// Consumer uses the channel
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}
	
	fmt.Println("  Channel ownership pattern completed")
}

func basicChannelOwner() <-chan int {
	ch := make(chan int, 5)
	
	go func() {
		defer close(ch) // Owner closes the channel
		
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	
	return ch // Return receive-only channel
}

// Example 2: Channel Factory Pattern
func channelFactoryPattern() {
	fmt.Println("\n2. Channel Factory Pattern")
	fmt.Println("==========================")
	
	// Create channel using factory
	ch := basicCreateChannel(3)
	
	// Send data
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	
	// Receive data
	for value := range ch {
		fmt.Printf("  Factory channel: %d\n", value)
	}
	
	fmt.Println("  Channel factory pattern completed")
}

func basicCreateChannel(capacity int) chan int {
	return make(chan int, capacity)
}

// Example 3: Channel Wrapper Pattern
func channelWrapperPattern() {
	fmt.Println("\n3. Channel Wrapper Pattern")
	fmt.Println("==========================")
	
	// Create safe channel wrapper
	safeCh := NewSafeChannel(3)
	
	// Send data
	go func() {
		defer safeCh.Close()
		for i := 0; i < 5; i++ {
			if success := safeCh.Send(i); success {
				fmt.Printf("  Sent: %d\n", i)
			} else {
				fmt.Printf("  Failed to send: %d\n", i)
			}
		}
	}()
	
	// Receive data
	for i := 0; i < 5; i++ {
		if value, ok := safeCh.Receive(); ok {
			fmt.Printf("  Received: %d\n", value)
		}
	}
	
	fmt.Println("  Channel wrapper pattern completed")
}

type SafeChannel struct {
	ch     chan int
	closed bool
	mu     sync.Mutex
}

func NewSafeChannel(capacity int) *SafeChannel {
	return &SafeChannel{
		ch: make(chan int, capacity),
	}
}

func (sc *SafeChannel) Send(value int) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	if sc.closed {
		return false
	}
	
	select {
	case sc.ch <- value:
		return true
	default:
		return false
	}
}

func (sc *SafeChannel) Receive() (int, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	if sc.closed {
		return 0, false
	}
	
	select {
	case value := <-sc.ch:
		return value, true
	default:
		return 0, false
	}
}

func (sc *SafeChannel) Close() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	if !sc.closed {
		close(sc.ch)
		sc.closed = true
	}
}

// Example 4: Graceful Shutdown Pattern
func gracefulShutdownPattern() {
	fmt.Println("\n4. Graceful Shutdown Pattern")
	fmt.Println("============================")
	
	ch := make(chan int)
	done := make(chan struct{})
	
	// Producer
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case ch <- i:
				fmt.Printf("  Produced: %d\n", i)
			case <-done:
				fmt.Println("  Producer shutting down")
				return
			}
		}
	}()
	
	// Consumer
	go func() {
		defer close(done)
		for value := range ch {
			fmt.Printf("  Consumed: %d\n", value)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Wait for completion
	<-done
	fmt.Println("  Graceful shutdown completed")
}

// Example 5: Nil Channel Tricks
func nilChannelTricks() {
	fmt.Println("\n5. Nil Channel Tricks")
	fmt.Println("====================")
	
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
			fmt.Printf("  Received from ch1: %d\n", value)
			ch2 = nil // Disable ch2
		case value := <-ch2:
			fmt.Printf("  Received from ch2: %s\n", value)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("  Timeout reached")
			break
		}
	}
	
	fmt.Println("  Nil channel tricks completed")
}

// Example 6: Channel Switching Pattern
func channelSwitchingPattern() {
	fmt.Println("\n6. Channel Switching Pattern")
	fmt.Println("============================")
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	
	go func() {
		for i := 10; i < 13; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	
	// Switch between channels
	for {
		select {
		case value, ok := <-ch1:
			if !ok {
				ch1 = nil // Disable ch1
				continue
			}
			fmt.Printf("  From ch1: %d\n", value)
		case value, ok := <-ch2:
			if !ok {
				ch2 = nil // Disable ch2
				continue
			}
			fmt.Printf("  From ch2: %d\n", value)
		case <-time.After(1 * time.Second):
			fmt.Println("  Timeout reached")
			break
		}
		
		// Exit when both channels are disabled
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	
	fmt.Println("  Channel switching pattern completed")
}

// Example 7: Channel Pipeline Pattern
func channelPipelinePattern() {
	fmt.Println("\n7. Channel Pipeline Pattern")
	fmt.Println("===========================")
	
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
		fmt.Printf("  Square: %d\n", square)
	}
	
	fmt.Println("  Channel pipeline pattern completed")
}

// Example 8: Channel Fan-Out Pattern
func channelFanOutPattern() {
	fmt.Println("\n8. Channel Fan-Out Pattern")
	fmt.Println("=========================")
	
	input := make(chan int)
	
	// Start multiple workers
	workers := make([]<-chan int, 3)
	for i := 0; i < 3; i++ {
		workers[i] = basicWorker(input, i)
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
				fmt.Printf("  Worker %d result: %d\n", id, result)
			}
		}(i, worker)
	}
	
	wg.Wait()
	fmt.Println("  Channel fan-out pattern completed")
}

func basicWorker(input <-chan int, id int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		for value := range input {
			// Process value
			result := value * value
			output <- result
		}
	}()
	
	return output
}

// Example 9: Channel Fan-In Pattern
func channelFanInPattern() {
	fmt.Println("\n9. Channel Fan-In Pattern")
	fmt.Println("=========================")
	
	// Create multiple input channels
	inputs := make([]<-chan int, 3)
	for i := 0; i < 3; i++ {
		inputs[i] = basicGenerateNumbers(i, 3)
	}
	
	// Fan-in to single output
	output := basicFanIn(inputs...)
	
	// Process results
	for value := range output {
		fmt.Printf("  Fan-in result: %d\n", value)
	}
	
	fmt.Println("  Channel fan-in pattern completed")
}

func basicGenerateNumbers(id, count int) <-chan int {
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- id*100 + i
		}
	}()
	
	return ch
}

func basicFanIn(inputs ...<-chan int) <-chan int {
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

// Example 10: Error Channel Pattern
func errorChannelPattern() {
	fmt.Println("\n10. Error Channel Pattern")
	fmt.Println("=========================")
	
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
			fmt.Printf("  Received: %d\n", value)
		case err := <-errorCh:
			fmt.Printf("  Error: %v\n", err)
			return
		}
	}
}

// Example 11: Channel Batching Pattern
func channelBatchingPattern() {
	fmt.Println("\n11. Channel Batching Pattern")
	fmt.Println("============================")
	
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
		fmt.Printf("  Batch: %v\n", batch)
	}
	
	fmt.Println("  Channel batching pattern completed")
}

// Example 12: Channel Rate Limiting Pattern
func channelRateLimitingPattern() {
	fmt.Println("\n12. Channel Rate Limiting Pattern")
	fmt.Println("=================================")
	
	input := make(chan int)
	output := make(chan int)
	
	// Rate limiter: 5 items per second
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()
	
	go func() {
		defer close(output)
		
		for value := range input {
			<-limiter.C // Wait for rate limit
			output <- value
		}
	}()
	
	// Send data
	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()
	
	// Process output
	for value := range output {
		fmt.Printf("  Rate limited: %d\n", value)
	}
	
	fmt.Println("  Channel rate limiting pattern completed")
}

// Example 13: Channel Generator Pattern
func channelGeneratorPattern() {
	fmt.Println("\n13. Channel Generator Pattern")
	fmt.Println("=============================")
	
	// Using the generator
	for value := range channelGenerator() {
		fmt.Printf("  Generated: %d\n", value)
	}
	
	fmt.Println("  Channel generator pattern completed")
}

func channelGenerator() <-chan int {
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	
	return ch
}

// Example 14: Channel Transformer Pattern
func channelTransformerPattern() {
	fmt.Println("\n14. Channel Transformer Pattern")
	fmt.Println("===============================")
	
	input := channelGenerator()
	squared := channelTransformer(input, func(x int) int { return x * x })
	
	for value := range squared {
		fmt.Printf("  Squared: %d\n", value)
	}
	
	fmt.Println("  Channel transformer pattern completed")
}

func channelTransformer(input <-chan int, transform func(int) int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		for value := range input {
			output <- transform(value)
		}
	}()
	
	return output
}

// Example 15: Channel Filter Pattern
func channelFilterPattern() {
	fmt.Println("\n15. Channel Filter Pattern")
	fmt.Println("==========================")
	
	input := channelGenerator()
	evens := channelFilter(input, func(x int) bool { return x%2 == 0 })
	
	for value := range evens {
		fmt.Printf("  Even: %d\n", value)
	}
	
	fmt.Println("  Channel filter pattern completed")
}

func channelFilter(input <-chan int, predicate func(int) bool) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		for value := range input {
			if predicate(value) {
				output <- value
			}
		}
	}()
	
	return output
}

// Example 16: Channel Accumulator Pattern
func channelAccumulatorPattern() {
	fmt.Println("\n16. Channel Accumulator Pattern")
	fmt.Println("===============================")
	
	input := channelGenerator()
	sums := channelAccumulator(input)
	
	for sum := range sums {
		fmt.Printf("  Running sum: %d\n", sum)
	}
	
	fmt.Println("  Channel accumulator pattern completed")
}

func channelAccumulator(input <-chan int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		
		sum := 0
		for value := range input {
			sum += value
			output <- sum
		}
	}()
	
	return output
}

// Example 17: Channel Pool Pattern
func channelPoolPattern() {
	fmt.Println("\n17. Channel Pool Pattern")
	fmt.Println("========================")
	
	pool := NewBasicChannelPool(3)
	
	// Use channels from pool
	for i := 0; i < 5; i++ {
		ch := pool.GetChannel()
		go func(id int, channel chan int) {
			channel <- id
		}(i, ch)
	}
	
	fmt.Println("  Channel pool pattern completed")
}

type BasicChannelPool struct {
	channels []chan int
	current  int
	mu       sync.Mutex
}

func NewBasicChannelPool(size int) *BasicChannelPool {
	pool := &BasicChannelPool{
		channels: make([]chan int, size),
	}
	
	for i := 0; i < size; i++ {
		pool.channels[i] = make(chan int, 10)
	}
	
	return pool
}

func (cp *BasicChannelPool) GetChannel() chan int {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	ch := cp.channels[cp.current]
	cp.current = (cp.current + 1) % len(cp.channels)
	return ch
}

// Example 18: Channel Mock Pattern
func channelMockPattern() {
	fmt.Println("\n18. Channel Mock Pattern")
	fmt.Println("========================")
	
	mock := NewChannelMock()
	
	// Test sending
	for i := 0; i < 3; i++ {
		if success := mock.Send(i); success {
			fmt.Printf("  Mock sent: %d\n", i)
		}
	}
	
	// Test receiving
	for i := 0; i < 3; i++ {
		if value, ok := mock.Receive(); ok {
			fmt.Printf("  Mock received: %d\n", value)
		}
	}
	
	mock.Close()
	fmt.Println("  Channel mock pattern completed")
}

type ChannelMock struct {
	sendCh    chan int
	receiveCh chan int
	closed    bool
	mu        sync.Mutex
}

func NewChannelMock() *ChannelMock {
	return &ChannelMock{
		sendCh:    make(chan int),
		receiveCh: make(chan int),
	}
}

func (cm *ChannelMock) Send(value int) bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	if cm.closed {
		return false
	}
	
	select {
	case cm.sendCh <- value:
		return true
	default:
		return false
	}
}

func (cm *ChannelMock) Receive() (int, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	if cm.closed {
		return 0, false
	}
	
	select {
	case value := <-cm.receiveCh:
		return value, true
	default:
		return 0, false
	}
}

func (cm *ChannelMock) Close() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	if !cm.closed {
		close(cm.sendCh)
		close(cm.receiveCh)
		cm.closed = true
	}
}

// Example 19: Channel Test Helper
func channelTestHelper() {
	fmt.Println("\n19. Channel Test Helper")
	fmt.Println("=======================")
	
	ch := make(chan int)
	
	// Test helper function
	testChannel := func(ch chan int) {
		go func() {
			defer close(ch)
			for i := 0; i < 3; i++ {
				ch <- i
			}
		}()
		
		var results []int
		for value := range ch {
			results = append(results, value)
		}
		
		expected := []int{0, 1, 2}
		if !reflect.DeepEqual(results, expected) {
			fmt.Printf("  Expected %v, got %v\n", expected, results)
		} else {
			fmt.Printf("  Test passed: %v\n", results)
		}
	}
	
	testChannel(ch)
	fmt.Println("  Channel test helper completed")
}

// Example 20: Channel Anti-Patterns
func channelAntiPatterns() {
	fmt.Println("\n20. Channel Anti-Patterns")
	fmt.Println("=========================")
	
	fmt.Println("  ❌ Anti-patterns to avoid:")
	fmt.Println("    - Channel leaks (not closing channels)")
	fmt.Println("    - Goroutine leaks (goroutines that never exit)")
	fmt.Println("    - Deadlocks (circular channel dependencies)")
	fmt.Println("    - Race conditions (unsafe channel access)")
	fmt.Println("    - Blocking operations (unbuffered channels)")
	
	fmt.Println("  ✅ Best practices:")
	fmt.Println("    - Always close channels when done")
	fmt.Println("    - Use buffered channels when appropriate")
	fmt.Println("    - Handle channel errors gracefully")
	fmt.Println("    - Use select with default for non-blocking operations")
	fmt.Println("    - Test channel patterns thoroughly")
	
	fmt.Println("  Channel anti-patterns completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🔗 Channel Patterns & Idioms Examples")
	fmt.Println("=====================================")
	
	basicChannelOwnership()
	channelFactoryPattern()
	channelWrapperPattern()
	gracefulShutdownPattern()
	nilChannelTricks()
	channelSwitchingPattern()
	channelPipelinePattern()
	channelFanOutPattern()
	channelFanInPattern()
	errorChannelPattern()
	channelBatchingPattern()
	channelRateLimitingPattern()
	channelGeneratorPattern()
	channelTransformerPattern()
	channelFilterPattern()
	channelAccumulatorPattern()
	channelPoolPattern()
	channelMockPattern()
	channelTestHelper()
	channelAntiPatterns()
}
