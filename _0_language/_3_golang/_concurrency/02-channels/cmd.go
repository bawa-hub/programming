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
