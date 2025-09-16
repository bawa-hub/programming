package main

import (
	"fmt"
)

// ============================================================================
// 1. DEADLOCK PITFALLS
// ============================================================================

func deadlockPitfalls() {
	fmt.Println("\n💀 DEADLOCK PITFALLS")
	fmt.Println("===================")

	// Deadlocks occur when all goroutines are blocked waiting for each other
	
	fmt.Println("\n1.1 Sending to nil channel")
	fmt.Println("  ⚠️  This would cause a deadlock:")
	fmt.Println("  var ch chan int")
	fmt.Println("  ch <- 42  // Blocks forever!")
	
	fmt.Println("\n1.2 Receiving from nil channel")
	fmt.Println("  ⚠️  This would cause a deadlock:")
	fmt.Println("  var ch chan int")
	fmt.Println("  data := <-ch  // Blocks forever!")
	
	fmt.Println("\n1.3 Circular dependency")
	fmt.Println("  ⚠️  This would cause a deadlock:")
	fmt.Println("  ch1 := make(chan int)")
	fmt.Println("  ch2 := make(chan int)")
	fmt.Println("  go func() { ch1 <- <-ch2 }()  // Waits for ch2")
	fmt.Println("  go func() { ch2 <- <-ch1 }()  // Waits for ch1")
	
	fmt.Println("\n1.4 Sending to unbuffered channel without receiver")
	fmt.Println("  ⚠️  This would cause a deadlock:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  ch <- 42  // Blocks forever if no receiver!")
}

// ============================================================================
// 2. CLOSED CHANNEL PITFALLS
// ============================================================================

func closedChannelPitfalls() {
	fmt.Println("\n🔒 CLOSED CHANNEL PITFALLS")
	fmt.Println("=========================")

	// Common mistakes when working with closed channels
	
	fmt.Println("\n2.1 Sending to closed channel")
	fmt.Println("  ⚠️  This causes a panic:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  close(ch)")
	fmt.Println("  ch <- 42  // panic: send on closed channel")
	
	fmt.Println("\n2.2 Closing already closed channel")
	fmt.Println("  ⚠️  This causes a panic:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  close(ch)")
	fmt.Println("  close(ch)  // panic: close of closed channel")
	
	fmt.Println("\n2.3 Not checking if channel is closed")
	fmt.Println("  ⚠️  This can cause issues:")
	fmt.Println("  for data := range ch {")
	fmt.Println("    // Process data")
	fmt.Println("  }")
	fmt.Println("  // If ch is never closed, this blocks forever!")
	
	fmt.Println("\n2.4 Receiving from closed channel")
	fmt.Println("  ✅ This is safe and returns zero value:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  close(ch)")
	fmt.Println("  data, ok := <-ch  // data = 0, ok = false")
}

// ============================================================================
// 3. BUFFER PITFALLS
// ============================================================================

func bufferPitfalls() {
	fmt.Println("\n📦 BUFFER PITFALLS")
	fmt.Println("=================")

	// Common mistakes with buffered channels
	
	fmt.Println("\n3.1 Assuming buffered channels are always non-blocking")
	fmt.Println("  ⚠️  Buffered channels still block when full:")
	fmt.Println("  ch := make(chan int, 2)")
	fmt.Println("  ch <- 1  // No blocking")
	fmt.Println("  ch <- 2  // No blocking")
	fmt.Println("  ch <- 3  // Blocks until space available!")
	
	fmt.Println("\n3.2 Not understanding buffer size")
	fmt.Println("  ⚠️  Buffer size affects behavior:")
	fmt.Println("  ch := make(chan int, 0)   // Unbuffered (synchronous)")
	fmt.Println("  ch := make(chan int, 1)   // Buffered (asynchronous)")
	fmt.Println("  ch := make(chan int, 10)  // Large buffer (decoupled)")
	
	fmt.Println("\n3.3 Buffer overflow")
	fmt.Println("  ⚠️  Sending more than buffer capacity blocks:")
	fmt.Println("  ch := make(chan int, 2)")
	fmt.Println("  for i := 0; i < 5; i++ {")
	fmt.Println("    ch <- i  // Last 3 sends will block!")
	fmt.Println("  }")
}

// ============================================================================
// 4. GOROUTINE LIFECYCLE PITFALLS
// ============================================================================

func goroutineLifecyclePitfalls() {
	fmt.Println("\n🔄 GOROUTINE LIFECYCLE PITFALLS")
	fmt.Println("=============================")

	// Common mistakes with goroutine lifecycle
	
	fmt.Println("\n4.1 Not waiting for goroutines to finish")
	fmt.Println("  ⚠️  This can cause program to exit early:")
	fmt.Println("  go func() {")
	fmt.Println("    time.Sleep(1 * time.Second)")
	fmt.Println("    fmt.Println(\"Work done!\")")
	fmt.Println("  }()")
	fmt.Println("  // Program exits before goroutine finishes!")
	
	fmt.Println("\n4.2 Goroutine leaks")
	fmt.Println("  ⚠️  This creates goroutine leaks:")
	fmt.Println("  for i := 0; i < 1000; i++ {")
	fmt.Println("    go func() {")
	fmt.Println("      time.Sleep(1 * time.Hour)  // Never finishes!")
	fmt.Println("    }()")
	fmt.Println("  }")
	
	fmt.Println("\n4.3 Not closing channels")
	fmt.Println("  ⚠️  This can cause goroutines to block forever:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  go func() {")
	fmt.Println("    for data := range ch {  // Blocks forever if ch never closed")
	fmt.Println("      // Process data")
	fmt.Println("    }")
	fmt.Println("  }()")
}

// ============================================================================
// 5. SELECT PITFALLS
// ============================================================================

func selectPitfalls() {
	fmt.Println("\n🎯 SELECT PITFALLS")
	fmt.Println("=================")

	// Common mistakes with select statements
	
	fmt.Println("\n5.1 Select without default case")
	fmt.Println("  ⚠️  This can block forever:")
	fmt.Println("  select {")
	fmt.Println("    case data := <-ch1:  // If ch1 never sends, blocks forever")
	fmt.Println("      // Process data")
	fmt.Println("  }")
	
	fmt.Println("\n5.2 Select with only nil channels")
	fmt.Println("  ⚠️  This blocks forever:")
	fmt.Println("  var ch1, ch2 chan int")
	fmt.Println("  select {")
	fmt.Println("    case <-ch1:  // nil channel, blocks forever")
	fmt.Println("    case <-ch2:  // nil channel, blocks forever")
	fmt.Println("  }")
	
	fmt.Println("\n5.3 Select with closed channels")
	fmt.Println("  ⚠️  This can cause unexpected behavior:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  close(ch)")
	fmt.Println("  select {")
	fmt.Println("    case data := <-ch:  // Always receives zero value")
	fmt.Println("      // This case always executes!")
	fmt.Println("  }")
}

// ============================================================================
// 6. MEMORY LEAK PITFALLS
// ============================================================================

func memoryLeakPitfalls() {
	fmt.Println("\n💧 MEMORY LEAK PITFALLS")
	fmt.Println("======================")

	// Common mistakes that cause memory leaks
	
	fmt.Println("\n6.1 Keeping references to large data")
	fmt.Println("  ⚠️  This can cause memory leaks:")
	fmt.Println("  ch := make(chan []byte, 1000)")
	fmt.Println("  for i := 0; i < 1000; i++ {")
	fmt.Println("    data := make([]byte, 1024*1024)  // 1MB")
	fmt.Println("    ch <- data  // Keeps reference in channel")
	fmt.Println("  }")
	
	fmt.Println("\n6.2 Not clearing channel buffers")
	fmt.Println("  ⚠️  This can cause memory leaks:")
	fmt.Println("  ch := make(chan int, 1000)")
	fmt.Println("  // Fill channel with data")
	fmt.Println("  // Never clear it - data stays in memory!")
	
	fmt.Println("\n6.3 Goroutine leaks")
	fmt.Println("  ⚠️  This can cause memory leaks:")
	fmt.Println("  for i := 0; i < 1000; i++ {")
	fmt.Println("    go func() {")
	fmt.Println("      time.Sleep(1 * time.Hour)  // Never finishes!")
	fmt.Println("    }()")
	fmt.Println("  }")
}

// ============================================================================
// 7. RACE CONDITION PITFALLS
// ============================================================================

func raceConditionPitfalls() {
	fmt.Println("\n🏃 RACE CONDITION PITFALLS")
	fmt.Println("=========================")

	// Common mistakes that cause race conditions
	
	fmt.Println("\n7.1 Accessing shared data without synchronization")
	fmt.Println("  ⚠️  This can cause race conditions:")
	fmt.Println("  var counter int")
	fmt.Println("  go func() { counter++ }()  // Race condition!")
	fmt.Println("  go func() { counter++ }()  // Race condition!")
	
	fmt.Println("\n7.2 Using channels incorrectly for synchronization")
	fmt.Println("  ⚠️  This can cause race conditions:")
	fmt.Println("  var data int")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  go func() {")
	fmt.Println("    data = 42  // Race condition!")
	fmt.Println("    ch <- 1")
	fmt.Println("  }()")
	fmt.Println("  <-ch")
	fmt.Println("  fmt.Println(data)  // May not be 42!")
}

// ============================================================================
// 8. PERFORMANCE PITFALLS
// ============================================================================

func performancePitfalls() {
	fmt.Println("\n⚡ PERFORMANCE PITFALLS")
	fmt.Println("======================")

	// Common mistakes that hurt performance
	
	fmt.Println("\n8.1 Using unbuffered channels when buffered would be better")
	fmt.Println("  ⚠️  This can hurt performance:")
	fmt.Println("  ch := make(chan int)  // Unbuffered")
	fmt.Println("  // Sender and receiver must synchronize for each value")
	
	fmt.Println("\n8.2 Creating too many goroutines")
	fmt.Println("  ⚠️  This can hurt performance:")
	fmt.Println("  for i := 0; i < 1000000; i++ {")
	fmt.Println("    go func() { /* work */ }()  // Too many goroutines!")
	fmt.Println("  }")
	
	fmt.Println("\n8.3 Using channels for simple data sharing")
	fmt.Println("  ⚠️  This can hurt performance:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  ch <- 42  // Overhead of channel communication")
	fmt.Println("  data := <-ch")
	fmt.Println("  // Use direct variable access when possible")
}

// ============================================================================
// 9. HOW TO AVOID PITFALLS
// ============================================================================

func howToAvoidPitfalls() {
	fmt.Println("\n✅ HOW TO AVOID PITFALLS")
	fmt.Println("========================")

	fmt.Println("\n9.1 Best practices to avoid pitfalls")
	
	fmt.Println("  ✅ Always check if channel is closed when receiving")
	fmt.Println("  ✅ Use buffered channels when you want to decouple sender/receiver")
	fmt.Println("  ✅ Use unbuffered channels when you need tight synchronization")
	fmt.Println("  ✅ Close channels to signal completion")
	fmt.Println("  ✅ Use select with default for non-blocking operations")
	fmt.Println("  ✅ Use sync.WaitGroup to wait for goroutines")
	fmt.Println("  ✅ Use context.Context for cancellation")
	fmt.Println("  ✅ Don't send to nil channels")
	fmt.Println("  ✅ Don't receive from nil channels")
	fmt.Println("  ✅ Don't close channels multiple times")
	
	fmt.Println("\n9.2 Debugging tips")
	fmt.Println("  🔍 Use 'go run -race' to detect race conditions")
	fmt.Println("  🔍 Use 'go vet' to detect common mistakes")
	fmt.Println("  🔍 Use 'go tool trace' to analyze performance")
	fmt.Println("  🔍 Use 'go tool pprof' to profile memory usage")
	fmt.Println("  🔍 Use 'go test -race' to test for race conditions")
}

// ============================================================================
// 10. COMMON MISTAKES SUMMARY
// ============================================================================

func commonMistakesSummary() {
	fmt.Println("\n📋 COMMON MISTAKES SUMMARY")
	fmt.Println("=========================")

	fmt.Println("\n10.1 Most common channel mistakes")
	
	fmt.Println("  1. Sending to nil channel (deadlock)")
	fmt.Println("  2. Receiving from nil channel (deadlock)")
	fmt.Println("  3. Sending to closed channel (panic)")
	fmt.Println("  4. Closing already closed channel (panic)")
	fmt.Println("  5. Not closing channels (goroutine leaks)")
	fmt.Println("  6. Not waiting for goroutines (early exit)")
	fmt.Println("  7. Using unbuffered channels when buffered would be better")
	fmt.Println("  8. Not checking if channel is closed")
	fmt.Println("  9. Creating too many goroutines")
	fmt.Println("  10. Using channels for simple data sharing")
	
	fmt.Println("\n10.2 How to fix them")
	
	fmt.Println("  🔧 Always initialize channels with make()")
	fmt.Println("  🔧 Use close() to signal completion")
	fmt.Println("  🔧 Use sync.WaitGroup to wait for goroutines")
	fmt.Println("  🔧 Use select with default for non-blocking operations")
	fmt.Println("  🔧 Use buffered channels when appropriate")
	fmt.Println("  🔧 Check if channel is closed when receiving")
	fmt.Println("  🔧 Use context.Context for cancellation")
	fmt.Println("  🔧 Use direct variable access when possible")
	fmt.Println("  🔧 Use go run -race to detect race conditions")
	fmt.Println("  🔧 Use go vet to detect common mistakes")
}

// ============================================================================
// EXPORTED FUNCTIONS FOR MAIN
// ============================================================================

func runChannelPitfalls() {
	fmt.Println("⚠️  GO CHANNELS: PITFALLS")
	fmt.Println("========================")
	
	// Run all channel pitfall examples
	deadlockPitfalls()
	closedChannelPitfalls()
	bufferPitfalls()
	goroutineLifecyclePitfalls()
	selectPitfalls()
	memoryLeakPitfalls()
	raceConditionPitfalls()
	performancePitfalls()
	howToAvoidPitfalls()
	commonMistakesSummary()
	
	fmt.Println("\n✅ Channel pitfalls completed!")
	fmt.Println("\n🎉 You now understand Go channels completely!")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Practice with the examples")
	fmt.Println("  2. Try building your own channel-based programs")
	fmt.Println("  3. Use 'go run -race' to test for race conditions")
	fmt.Println("  4. Use 'go vet' to check for common mistakes")
}
