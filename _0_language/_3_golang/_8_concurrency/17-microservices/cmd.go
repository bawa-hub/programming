package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [basic|exercises|advanced|all]")
		fmt.Println("  basic     - Run basic microservices examples")
		fmt.Println("  exercises - Run hands-on exercises")
		fmt.Println("  advanced  - Run advanced microservices patterns")
		fmt.Println("  all       - Run everything")
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
		fmt.Println("Available commands: basic, exercises, advanced, all")
	}
}

