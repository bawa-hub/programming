package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var packageName string
	flag.StringVar(&packageName, "package", "", "Package to run (fmt, os, io, time, math, reflect, errors, log)")
	flag.Parse()

	if packageName == "" {
		showHelp()
		return
	}

	switch strings.ToLower(packageName) {
	case "fmt":
		runFmtExamples()
	case "os":
		runOsExamples()
	case "io":
		runIoExamples()
	case "time":
		runTimeExamples()
	case "math":
		runMathExamples()
	case "reflect":
		runReflectExamples()
	case "errors":
		runErrorsExamples()
	case "log":
		runLogExamples()
	case "all":
		runAllExamples()
	default:
		fmt.Printf("Unknown package: %s\n", packageName)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("🚀 Go Standard Library Mastery")
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("Usage: go run main.go -package <package_name>")
	fmt.Println()
	fmt.Println("Available packages:")
	fmt.Println("  fmt      - Formatting and printing")
	fmt.Println("  os       - Operating system interface")
	fmt.Println("  io       - I/O primitives")
	fmt.Println("  time     - Time and date operations")
	fmt.Println("  math     - Mathematical functions")
	fmt.Println("  reflect  - Runtime reflection")
	fmt.Println("  errors   - Error handling")
	fmt.Println("  log      - Logging")
	fmt.Println("  all      - Run all examples")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go -package fmt")
	fmt.Println("  go run main.go -package all")
}

func runFmtExamples() {
	fmt.Println("Running fmt package examples...")
	// This would call the fmt examples
	fmt.Println("fmt examples would run here")
}

func runOsExamples() {
	fmt.Println("Running os package examples...")
	// This would call the os examples
	fmt.Println("os examples would run here")
}

func runIoExamples() {
	fmt.Println("Running io package examples...")
	// This would call the io examples
	fmt.Println("io examples would run here")
}

func runTimeExamples() {
	fmt.Println("Running time package examples...")
	// This would call the time examples
	fmt.Println("time examples would run here")
}

func runMathExamples() {
	fmt.Println("Running math package examples...")
	// This would call the math examples
	fmt.Println("math examples would run here")
}

func runReflectExamples() {
	fmt.Println("Running reflect package examples...")
	// This would call the reflect examples
	fmt.Println("reflect examples would run here")
}

func runErrorsExamples() {
	fmt.Println("Running errors package examples...")
	// This would call the errors examples
	fmt.Println("errors examples would run here")
}

func runLogExamples() {
	fmt.Println("Running log package examples...")
	// This would call the log examples
	fmt.Println("log examples would run here")
}

func runAllExamples() {
	fmt.Println("Running all package examples...")
	
	packages := []string{"fmt", "os", "io", "time", "math", "reflect", "errors", "log"}
	
	for _, pkg := range packages {
		fmt.Printf("\n" + strings.Repeat("=", 50) + "\n")
		fmt.Printf("Running %s package examples\n", pkg)
		fmt.Printf(strings.Repeat("=", 50) + "\n")
		
		switch pkg {
		case "fmt":
			runFmtExamples()
		case "os":
			runOsExamples()
		case "io":
			runIoExamples()
		case "time":
			runTimeExamples()
		case "math":
			runMathExamples()
		case "reflect":
			runReflectExamples()
		case "errors":
			runErrorsExamples()
		case "log":
			runLogExamples()
		}
	}
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("All examples completed!")
	fmt.Println(strings.Repeat("=", 50))
}
