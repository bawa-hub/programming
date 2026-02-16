package main

import (
	"fmt"
	"time"
	pat "goroutines/patterns"
)


// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Goroutine Patterns")
	fmt.Println("==============================")
	
	// Pattern 1: Dynamic Pool
	fmt.Println("\n1. Dynamic Goroutine Pool:")
	dynamicPool := pat.NewDynamicPool(3)
	dynamicPool.Start()
	
	for i := 0; i < 10; i++ {
		dynamicPool.Submit(func() {
			fmt.Printf("Dynamic pool job %d\n", i)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	time.Sleep(2 * time.Second)
	dynamicPool.Stop()
	
	// Pattern 2: Context Cancellation
	pat.ContextGoroutine()
	
	// Pattern 3: Heartbeat
	fmt.Println("\n3. Heartbeat Goroutine:")
	heartbeat := pat.NewHeartbeatGoroutine()
	heartbeat.Start()
	
	go func() {
		for t := range heartbeat.GetHeartbeat() {
			fmt.Printf("Heartbeat: %v\n", t)
		}
	}()
	
	time.Sleep(3 * time.Second)
	heartbeat.Stop()
	
	// Pattern 4: Circuit Breaker
	fmt.Println("\n4. Circuit Breaker:")
	cb := pat.NewCircuitBreaker(3, 1*time.Second)
	
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
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
	
	// Pattern 5: Backpressure
	fmt.Println("\n5. Backpressure Goroutine:")
	bp := pat.NewBackpressureGoroutine(2)
	bp.Start()
	
	for i := 0; i < 5; i++ {
		success := bp.Send(i)
		if success {
			fmt.Printf("Sent item %d\n", i)
		} else {
			fmt.Printf("Failed to send item %d (backpressure)\n", i)
		}
	}
	
	// Pattern 6: Metrics
	fmt.Println("\n6. Metrics Goroutine:")
	metrics := pat.NewMetricsGoroutine()
	
	for i := 0; i < 5; i++ {
		metrics.RecordProcessed()
		if i%2 == 0 {
			metrics.RecordError()
		}
	}
	
	processed, errors, uptime := metrics.GetStats()
	fmt.Printf("Processed: %d, Errors: %d, Uptime: %v\n", processed, errors, uptime)
	
	// Pattern 7: Graceful Shutdown
	fmt.Println("\n7. Graceful Shutdown:")
	graceful := pat.NewGracefulGoroutine()
	graceful.Start()
	
	time.Sleep(2 * time.Second)
	graceful.Shutdown()
	graceful.Wait()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
