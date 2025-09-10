package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	
	"go-basics-mastery/arrays-slices"
	"go-basics-mastery/examples"
	"go-basics-mastery/interfaces"
	"go-basics-mastery/pointers"
	"go-basics-mastery/primitives"
	"go-basics-mastery/structs"
)

func main() {
	// Define command-line flags
	var (
		module = flag.String("module", "all", "Module to run (primitives, arrays, structs, interfaces, pointers, examples, all)")
		help   = flag.Bool("help", false, "Show help message")
	)
	
	flag.Parse()

	// Show help if requested
	if *help {
		showHelp()
		return
	}

	// Run the specified module
	switch *module {
	case "primitives":
		runPrimitives()
	case "arrays":
		runArraysSlices()
	case "structs":
		runStructs()
	case "interfaces":
		runInterfaces()
	case "pointers":
		runPointers()
	case "examples":
		runExamples()
	case "all":
		runAll()
	default:
		fmt.Printf("Unknown module: %s\n", *module)
		showHelp()
		os.Exit(1)
	}
}

func runPrimitives() {
	fmt.Println("🧮 RUNNING PRIMITIVE DATA TYPES MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	primitives.RunAllPrimitiveExamples()
}

func runArraysSlices() {
	fmt.Println("📊 RUNNING ARRAYS AND SLICES MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	arrays_slices.RunAllArraySliceExamples()
}

func runStructs() {
	fmt.Println("🏗️ RUNNING STRUCTS MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	structs.RunAllStructExamples()
}

func runInterfaces() {
	fmt.Println("🔌 RUNNING INTERFACES MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	interfaces.RunAllInterfaceExamples()
}

func runPointers() {
	fmt.Println("📍 RUNNING POINTERS MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	pointers.RunAllPointerExamples()
}

func runExamples() {
	fmt.Println("🚀 RUNNING PRACTICAL EXAMPLES MODULE")
	fmt.Println("=" + strings.Repeat("=", 50))
	examples.RunAllExamples()
}

func runAll() {
	fmt.Println("🎯 GO BASICS MASTERY - COMPLETE DEMONSTRATION")
	fmt.Println("=" + strings.Repeat("=", 60))
	
	runPrimitives()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	
	runArraysSlices()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	
	runStructs()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	
	runInterfaces()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	
	runPointers()
	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
	
	runExamples()
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎉 CONGRATULATIONS! You've completed the Go Basics Mastery course!")
	fmt.Println("You now have a solid foundation in Go fundamentals.")
	fmt.Println("Ready to move on to Project 2: File System Scanner!")
}

func showHelp() {
	fmt.Println("Go Basics Mastery - Complete Go Fundamentals Course")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run main.go [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -module string")
	fmt.Println("        Module to run (primitives, arrays, structs, interfaces, pointers, examples, all)")
	fmt.Println("        (default \"all\")")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Modules:")
	fmt.Println("  primitives  - Primitive data types (int, float, string, bool, etc.)")
	fmt.Println("  arrays      - Arrays and slices with operations")
	fmt.Println("  structs     - Structs, methods, and embedding")
	fmt.Println("  interfaces  - Interfaces and polymorphism")
	fmt.Println("  pointers    - Pointers and memory management")
	fmt.Println("  examples    - Practical examples combining all concepts")
	fmt.Println("  all         - Run all modules (default)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go -module primitives")
	fmt.Println("  go run main.go -module arrays")
	fmt.Println("  go run main.go -module all")
	fmt.Println()
	fmt.Println("Learning Path:")
	fmt.Println("  1. Start with primitives to understand basic data types")
	fmt.Println("  2. Move to arrays/slices to learn collections")
	fmt.Println("  3. Study structs for object-oriented concepts")
	fmt.Println("  4. Explore interfaces for polymorphism")
	fmt.Println("  5. Practice pointers for memory management")
	fmt.Println("  6. See everything combined in examples")
}
