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
	fmt.Println("🚀 Pipeline Pattern - Usage")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic pipeline examples")
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
	fmt.Println("🚀 Pipeline Pattern Examples")
	fmt.Println("============================")

	// Example 1: Basic Pipeline
	basicPipeline()

	// Example 2: Buffered Pipeline
	bufferedPipeline()

	// Example 3: Fan-Out/Fan-In Pipeline
	fanOutFanInPipeline()

	// Example 4: Pipeline with Error Handling
	pipelineWithErrorHandling()

	// Example 5: Pipeline with Timeout
	pipelineWithTimeout()

	// Example 6: Pipeline with Rate Limiting
	pipelineWithRateLimiting()

	// Example 7: Pipeline with Metrics
	pipelineWithMetrics()

	// Example 8: Pipeline with Backpressure
	pipelineWithBackpressure()

	// Example 9: Performance Comparison
	performanceComparison()

	// Example 10: Common Pitfalls
	commonPitfalls()
}

