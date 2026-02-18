package main

import (
	ch "channels/exercises"
	pat "channels/patterns"
	"fmt"
	"os"
	"time"
)

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
	fmt.Println("🚀 Channels Fundamentals - Usage")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic channel examples")
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

func runBasicExamples() {
	fmt.Println("🚀 Channels Fundamentals Examples")
	fmt.Println("==================================")

	// Example 1: Basic Channel Operations
	basicChannelOperations()

	// Example 2: Buffered vs Unbuffered Channels
	bufferedVsUnbuffered()

	// Example 3: Channel Direction
	channelDirection()

	// Example 4: Channel Closing
	channelClosing()

	// Example 5: Select Statement
	selectStatement()

	// Example 6: Pipeline Pattern
	pipelinePattern()

	// Example 7: Fan-Out Pattern
	fanOutPattern()

	// Example 8: Fan-In Pattern
	fanInPattern()

	// Example 9: Channel Timeout
	channelTimeout()

	// Example 10: Channel Performance
	channelPerformance()

	// Example 11: Common Pitfalls
	commonPitfalls()
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Channel Exercises")
	fmt.Println("=================================")
	
	ch.Exercise1()
	ch.Exercise2()
	ch.Exercise3()
	ch.Exercise4()
	ch.Exercise5()
	ch.Exercise6()
	ch.Exercise7()
	ch.Exercise8()
	ch.Exercise9()
	ch.Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Channel Patterns")
	fmt.Println("============================")
	
	// Pattern 1: State Machine
	fmt.Println("\n1. Channel-based State Machine:")
	sm := pat.NewStateMachine()
	sm.DoAction(func() {
		fmt.Printf("Current state: %d\n", sm.GetState())
	})
	sm.SetState(1)
	sm.DoAction(func() {
		fmt.Printf("Current state: %d\n", sm.GetState())
	})
	sm.Stop()
	
	// Pattern 2: Rate Limiter
	fmt.Println("\n2. Channel-based Rate Limiter:")
	rl := pat.NewRateLimiter(100*time.Millisecond, 3)
	for i := 0; i < 5; i++ {
		if rl.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i)
		}
	}
	rl.Stop()
	
	// Pattern 3: Circuit Breaker
	fmt.Println("\n3. Channel-based Circuit Breaker:")
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
	cb.Close()
	
	// Pattern 4: Event Bus
	fmt.Println("\n4. Channel-based Event Bus:")
	eb := pat.NewEventBus()
	
	// Subscribe to events
	ch1 := eb.Subscribe("events")
	ch2 := eb.Subscribe("events")
	
	// Listen for events
	go func() {
		for event := range ch1 {
			fmt.Printf("Subscriber 1 received: %v\n", event)
		}
	}()
	
	go func() {
		for event := range ch2 {
			fmt.Printf("Subscriber 2 received: %v\n", event)
		}
	}()
	
	// Publish events
	eb.Publish("events", "Event 1")
	eb.Publish("events", "Event 2")
	
	time.Sleep(100 * time.Millisecond)
	
	// Pattern 5: Priority Worker Pool
	fmt.Println("\n5. Channel-based Priority Worker Pool:")
	pwp := pat.NewPriorityWorkerPool(2)
	pwp.Start()
	
	// Submit jobs with different priorities
	for i := 0; i < 5; i++ {
		job := pat.Job{
			ID:       i,
			Priority: (i % 3) + 1, // 1, 2, 3
			Data:     fmt.Sprintf("Data %d", i),
			Result:   make(chan interface{}, 1),
		}
		pwp.Submit(job)
	}
	
	time.Sleep(1 * time.Second)
	pwp.Stop()
	
	// Pattern 6: Load Balancer
	fmt.Println("\n6. Channel-based Load Balancer:")
	lb := pat.NewLoadBalancer(3)
	
	for i := 0; i < 5; i++ {
		job := Job{
			ID:     i,
			Data:   fmt.Sprintf("Data %d", i),
			Result: make(chan interface{}, 1),
		}
		lb.Submit(job)
	}
	
	time.Sleep(1 * time.Second)
	lb.Stop()
	
	// Pattern 7: Channel-based Context
	fmt.Println("\n7. Channel-based Context:")
	ctx := pat.NewChannelContext()
	
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
		}
	}()
	
	time.Sleep(1 * time.Second)
	ctx.Cancel()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
