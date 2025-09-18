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
	fmt.Println("🚀 Goroutines Deep Dive - Usage")
	fmt.Println("===============================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic goroutine examples")
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
	fmt.Println("🚀 Goroutines Deep Dive Examples")
	fmt.Println("================================")

	// Example 1: Basic Goroutine
	fmt.Println("\n1. Basic Goroutine:")
	basicGoroutine()

	// Example 2: Multiple Goroutines
	fmt.Println("\n2. Multiple Goroutines:")
	multipleGoroutines()

	// Example 3: Goroutine with WaitGroup
	fmt.Println("\n3. Goroutine with WaitGroup:")
	goroutineWithWaitGroup()

	// Example 4: Goroutine Pool
	fmt.Println("\n4. Goroutine Pool:")
	goroutinePool()

	// Example 5: Goroutine Communication
	fmt.Println("\n5. Goroutine Communication:")
	goroutineCommunication()

	// Example 6: Goroutine Lifecycle
	fmt.Println("\n6. Goroutine Lifecycle:")
	goroutineLifecycle()

	// Example 7: Performance Comparison
	fmt.Println("\n7. Performance Comparison:")
	performanceComparison()

	// Example 8: Common Pitfalls
	fmt.Println("\n8. Common Pitfalls:")
	commonPitfalls()
}
