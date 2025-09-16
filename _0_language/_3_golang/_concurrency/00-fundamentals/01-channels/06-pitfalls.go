package main

import (
	"fmt"
)

// ============================================================================
// 1. DEADLOCK PITFALLS
// ============================================================================

	// Deadlocks occur when all goroutines are blocked waiting for each other

	// Sending to nil channel
	// This would cause a deadlock:
	// var ch chan int
	// ch <- 42  // Blocks forever!

	// Receiving from nil channel
	// ⚠️  This would cause a deadlock:
	// var ch chan int
	// data := <-ch  // Blocks forever!

	// Circular dependency
	// ⚠️  This would cause a deadlock:
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go func() { ch1 <- <-ch2 }()  // Waits for ch2
	// go func() { ch2 <- <-ch1 }()  // Waits for ch1

	// Sending to unbuffered channel without receiver
	// ⚠️  This would cause a deadlock:
	// ch := make(chan int)
	// ch <- 42  // Blocks forever if no receiver!

// ============================================================================
// 2. CLOSED CHANNEL PITFALLS
// ============================================================================

	// Common mistakes when working with closed channels
	
	// Sending to closed channel
	// ⚠️  This causes a panic:
	// ch := make(chan int)
	// close(ch)
	// ch <- 42  // panic: send on closed channel
	
	// Closing already closed channel
	// ⚠️  This causes a panic:
	// ch := make(chan int)
	// close(ch)
	// close(ch)  // panic: close of closed channel
	
	// Not checking if channel is closed
	// ⚠️  This can cause issues:
	// for data := range ch {
	//  // Process data
	// }
	// // If ch is never closed, this blocks forever!
	
	// Receiving from closed channel
	//  ✅ This is safe and returns zero value:
	//  ch := make(chan int)
	//  close(ch)
	//  data, ok := <-ch  // data = 0, ok = false


// ============================================================================
// 3. BUFFER PITFALLS
// ============================================================================

	// Common mistakes with buffered channels
	
	// Assuming buffered channels are always non-blocking
	//  ⚠️  Buffered channels still block when full:
	//  ch := make(chan int, 2)
	//  ch <- 1  // No blocking
	//  ch <- 2  // No blocking
	//  ch <- 3  // Blocks until space available!
	
	// Not understanding buffer size
	//  ⚠️  Buffer size affects behavior:
	//  ch := make(chan int, 0)   // Unbuffered (synchronous)
	//  ch := make(chan int, 1)   // Buffered (asynchronous)
	//  ch := make(chan int, 10)  // Large buffer (decoupled)
	
	// Buffer overflow
	//  ⚠️  Sending more than buffer capacity blocks:
	//  ch := make(chan int, 2)
	//  for i := 0; i < 5; i++ {
	//    ch <- i  // Last 3 sends will block!
	//  }


// ============================================================================
// 4. GOROUTINE LIFECYCLE PITFALLS
// ============================================================================

	// Common mistakes with goroutine lifecycle
	
	// Not waiting for goroutines to finish
	//  ⚠️  This can cause program to exit early:
	//  go func() {
	//    time.Sleep(1 * time.Second)
	//    fmt.Println(\"Work done!\
	//  }()
	//  // Program exits before goroutine finishes!
	
	// Goroutine leaks
	//  ⚠️  This creates goroutine leaks:
	//  for i := 0; i < 1000; i++ {
	//    go func() {
	//      time.Sleep(1 * time.Hour)  // Never finishes!
	//    }()
	//  }
	
	// Not closing channels
	//  ⚠️  This can cause goroutines to block forever:
	//  ch := make(chan int)
	//  go func() {
	//    for data := range ch {  // Blocks forever if ch never closed
	//      // Process data
	//    }
	//  }()


// ============================================================================
// 5. SELECT PITFALLS
// ============================================================================

	// Common mistakes with select statements
	
	// Select without default case
	//  ⚠️  This can block forever:
	//  select {
	//    case data := <-ch1:  // If ch1 never sends, blocks forever
	//      // Process data
	//  }
	
	// Select with only nil channels
	//  ⚠️  This blocks forever:
	//  var ch1, ch2 chan int
	//  select {
	//    case <-ch1:  // nil channel, blocks forever
	//    case <-ch2:  // nil channel, blocks forever
	//  }
	
	//  Select with closed channels
	//  ⚠️  This can cause unexpected behavior:
	//  ch := make(chan int)
	//  close(ch)
	//  select {
	//    case data := <-ch:  // Always receives zero value
	//      // This case always executes!
	//  }


// ============================================================================
// 6. MEMORY LEAK PITFALLS
// ============================================================================

	// Common mistakes that cause memory leaks
	
	// Keeping references to large data
	//  ⚠️  This can cause memory leaks:
	//  ch := make(chan []byte, 1000)
	//  for i := 0; i < 1000; i++ {
	//    data := make([]byte, 1024*1024)  // 1MB
	//    ch <- data  // Keeps reference in channel
	//  }
	
	// Not clearing channel buffers
	//  ⚠️  This can cause memory leaks:
	//  ch := make(chan int, 1000)
	//  // Fill channel with data
	//  // Never clear it - data stays in memory!
	
	// Goroutine leaks
	//  ⚠️  This can cause memory leaks:
	//  for i := 0; i < 1000; i++ {
	//    go func() {
	//      time.Sleep(1 * time.Hour)  // Never finishes!
	//    }()
	//  }


// ============================================================================
// 7. RACE CONDITION PITFALLS
// ============================================================================

	// Common mistakes that cause race conditions
	
	// Accessing shared data without synchronization
	//  ⚠️  This can cause race conditions:
	//  var counter int
	//  go func() { counter++ }()  // Race condition!
	//  go func() { counter++ }()  // Race condition!
	
	// Using channels incorrectly for synchronization
	//  ⚠️  This can cause race conditions:
	//  var data int
	//  ch := make(chan int)
	//  go func() {
	//    data = 42  // Race condition!
	//    ch <- 1
	//  }()
	//  <-ch
	//  fmt.Println(data)  // May not be 42!

// ============================================================================
// 8. PERFORMANCE PITFALLS
// ============================================================================

	// Common mistakes that hurt performance
	
	// Using unbuffered channels when buffered would be better
	//  ⚠️  This can hurt performance:
	//  ch := make(chan int)  // Unbuffered
	//  // Sender and receiver must synchronize for each value
	
	// Creating too many goroutines
	//  ⚠️  This can hurt performance:
	//  for i := 0; i < 1000000; i++ {
	//    go func() { /* work */ }()  // Too many goroutines!
	//  }
	
	// Using channels for simple data sharing
	//  ⚠️  This can hurt performance:
	//  ch := make(chan int)
	//  ch <- 42  // Overhead of channel communication
	//  data := <-ch
	//  // Use direct variable access when possible


// ============================================================================
// 9. HOW TO AVOID PITFALLS
// ============================================================================


	// Best practices to avoid pitfalls
	
	//  ✅ Always check if channel is closed when receiving
	//  ✅ Use buffered channels when you want to decouple sender/receiver
	//  ✅ Use unbuffered channels when you need tight synchronization
	//  ✅ Close channels to signal completion
	//  ✅ Use select with default for non-blocking operations
	//  ✅ Use sync.WaitGroup to wait for goroutines
	//  ✅ Use context.Context for cancellation
	//  ✅ Don't send to nil channels
	//  ✅ Don't receive from nil channels
	//  ✅ Don't close channels multiple times
	
	// Debugging tips
	//  🔍 Use 'go run -race' to detect race conditions
	//  🔍 Use 'go vet' to detect common mistakes
	//  🔍 Use 'go tool trace' to analyze performance
	//  🔍 Use 'go tool pprof' to profile memory usage
	//  🔍 Use 'go test -race' to test for race conditions

// ============================================================================
// 10. COMMON MISTAKES SUMMARY
// ============================================================================

	// Most common channel mistakes
	//  1. Sending to nil channel (deadlock)
	//  2. Receiving from nil channel (deadlock)
	//  3. Sending to closed channel (panic)
	//  4. Closing already closed channel (panic)
	//  5. Not closing channels (goroutine leaks)
	//  6. Not waiting for goroutines (early exit)
	//  7. Using unbuffered channels when buffered would be better
	//  8. Not checking if channel is closed
	//  9. Creating too many goroutines
	//  10. Using channels for simple data sharing
	
	//  How to fix them
	//  🔧 Always initialize channels with make()
	//  🔧 Use close() to signal completion
	//  🔧 Use sync.WaitGroup to wait for goroutines
	//  🔧 Use select with default for non-blocking operations
	//  🔧 Use buffered channels when appropriate
	//  🔧 Check if channel is closed when receiving
	//  🔧 Use context.Context for cancellation
	//  🔧 Use direct variable access when possible
	//  🔧 Use go run -race to detect race conditions
	//  🔧 Use go vet to detect common mistakes

