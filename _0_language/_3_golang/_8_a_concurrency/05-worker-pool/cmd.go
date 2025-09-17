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
	fmt.Println("🚀 Worker Pool Pattern - Usage")
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic worker pool examples")
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
	fmt.Println("🚀 Worker Pool Pattern Examples")
	fmt.Println("===============================")

	// Example 1: Basic Worker Pool
	basicWorkerPool()

	// Example 2: Buffered Worker Pool
	bufferedWorkerPool()

	// Example 3: Dynamic Worker Pool
	dynamicWorkerPool()

	// Example 4: Priority Worker Pool
	priorityWorkerPool()

	// Example 5: Worker Pool with Results
	workerPoolWithResults()

	// Example 6: Worker Pool with Error Handling
	workerPoolWithErrorHandling()

	// Example 7: Worker Pool with Timeout
	workerPoolWithTimeout()

	// Example 8: Worker Pool with Rate Limiting
	workerPoolWithRateLimiting()

	// Example 9: Worker Pool with Metrics
	workerPoolWithMetrics()

	// Example 10: Pipeline Worker Pool
	pipelineWorkerPool()

	// Example 11: Performance Comparison
	performanceComparison()

	// Example 12: Common Pitfalls
	commonPitfalls()
}
