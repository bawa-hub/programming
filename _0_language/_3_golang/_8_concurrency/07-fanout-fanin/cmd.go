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
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("🚀 Fan-Out/Fan-In Pattern Commands")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  basic     - Run basic fan-out/fan-in examples")
	fmt.Println("  exercises - Run hands-on exercises")
	fmt.Println("  advanced  - Run advanced patterns")
	fmt.Println("  help      - Show this help message")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . basic")
	fmt.Println("  go run . exercises")
	fmt.Println("  go run . advanced")
}
