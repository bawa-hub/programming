package main

import (
	"fmt"
	"os"
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
	fmt.Println("🚀 Synchronization Primitives - Usage")
	fmt.Println("=====================================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic synchronization examples")
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
	fmt.Println("🚀 Synchronization Primitives Examples")
	fmt.Println("======================================")

	// Example 1: Basic Mutex
	basicMutex()

	// Example 2: RWMutex (Read-Write Mutex)
	rwMutex()

	// Example 3: WaitGroup
	waitGroup()

	// Example 4: Once (One-Time Execution)
	onceExample()

	// Example 5: Cond (Condition Variables)
	condExample()

	// Example 6: Atomic Operations
	atomicExample()

	// Example 7: Concurrent Map
	concurrentMap()

	// Example 8: Object Pool
	objectPool()

	// Example 9: Performance Comparison
	performanceComparison()

	// Example 10: Deadlock Prevention
	deadlockPrevention()

	// Example 11: Race Condition Detection
	raceConditionDetection()

	// Example 12: Common Pitfalls
	commonPitfalls()
}
