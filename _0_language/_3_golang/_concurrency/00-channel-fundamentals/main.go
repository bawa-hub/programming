package main

import (
	"fmt"
	"os"
)

// ============================================================================
// MAIN FUNCTION - CHANNEL FUNDAMENTALS DEMONSTRATION
// ============================================================================

func main() {
	// Check command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "basic":
			runBasicConcepts()
		case "types":
			runChannelTypes()
		case "operations":
			runChannelOperations()
		case "behavior":
			runChannelBehavior()
		case "patterns":
			runChannelPatterns()
		case "pitfalls":
			runChannelPitfalls()
		default:
			showUsage()
		}
	} else {
		// Run all examples by default
		runAllExamples()
	}
}

// ============================================================================
// RUN ALL EXAMPLES
// ============================================================================

func runAllExamples() {
	fmt.Println("🚀 GO CHANNELS: COMPLETE FUNDAMENTALS")
	fmt.Println("=====================================")
	fmt.Println("This comprehensive guide covers every aspect of Go channels")
	fmt.Println("from basic concepts to advanced patterns and common pitfalls.")
	fmt.Println()
	
	// Run all examples
	runBasicConcepts()
	runChannelTypes()
	runChannelOperations()
	runChannelBehavior()
	runChannelPatterns()
	runChannelPitfalls()
	
	fmt.Println("\n🎉 ALL CHANNEL FUNDAMENTALS COMPLETED!")
	fmt.Println("=====================================")
	fmt.Println("You now have a complete understanding of Go channels!")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Practice with the examples")
	fmt.Println("  2. Try building your own channel-based programs")
	fmt.Println("  3. Use 'go run -race' to test for race conditions")
	fmt.Println("  4. Use 'go vet' to check for common mistakes")
	fmt.Println("  5. Explore the main concurrency curriculum")
}

// ============================================================================
// RUN SPECIFIC EXAMPLES
// ============================================================================
// Note: The actual function implementations are in the individual files

// ============================================================================
// USAGE INFORMATION
// ============================================================================

func showUsage() {
	fmt.Println("🚀 GO CHANNELS: COMPLETE FUNDAMENTALS")
	fmt.Println("=====================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run .                    # Run all examples")
	fmt.Println("  go run . basic              # Run basic concepts")
	fmt.Println("  go run . types              # Run channel types")
	fmt.Println("  go run . operations         # Run channel operations")
	fmt.Println("  go run . behavior           # Run channel behavior")
	fmt.Println("  go run . patterns           # Run channel patterns")
	fmt.Println("  go run . pitfalls           # Run channel pitfalls")
	fmt.Println()
	fmt.Println("This comprehensive guide covers every aspect of Go channels")
	fmt.Println("from basic concepts to advanced patterns and common pitfalls.")
	fmt.Println()
	fmt.Println("Each file demonstrates specific channel concepts with")
	fmt.Println("detailed comments and explanations.")
}
