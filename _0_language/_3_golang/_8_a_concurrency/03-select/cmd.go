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
