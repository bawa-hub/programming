package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "basic":
		runBasicExamples()
	case "exercises":
		runExercises()
	case "advanced":
		runAdvancedPatterns()
	case "all":
		runBasicExamples()
		runExercises()
		runAdvancedPatterns()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("🔗 Channel Patterns & Idioms Commands")
	fmt.Println("====================================")
	fmt.Println("")
	fmt.Println("Usage: go run . <command>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  basic     - Run basic channel patterns and idioms")
	fmt.Println("  exercises - Run hands-on exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  all       - Run all examples and exercises")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println("  go run . all")
}

