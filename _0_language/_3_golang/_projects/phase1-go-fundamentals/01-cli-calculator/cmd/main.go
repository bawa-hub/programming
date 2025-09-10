package main

import (
	"cli-calculator/internal/cli"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	var (
		expression = flag.String("expr", "", "Mathematical expression to evaluate")
		interactive = flag.Bool("i", false, "Run in interactive mode")
		help = flag.Bool("help", false, "Show help message")
	)
	
	flag.Parse()

	// Show help if requested
	if *help {
		showHelp()
		return
	}

	// Create CLI instance
	calcCLI := cli.New()

	// Check if expression is provided
	if *expression != "" {
		// Run in batch mode
		calcCLI.RunBatch(*expression)
		return
	}

	// Check if interactive mode is explicitly requested or no flags provided
	if *interactive || len(os.Args) == 1 {
		// Run in interactive mode
		calcCLI.Run()
		return
	}

	// If no expression provided and not in interactive mode, show help
	showHelp()
}

func showHelp() {
	fmt.Println("Go CLI Calculator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cli-calculator [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -expr string")
	fmt.Println("        Mathematical expression to evaluate")
	fmt.Println("  -i    Run in interactive mode")
	fmt.Println("  -help Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cli-calculator -expr '2 + 3'")
	fmt.Println("  cli-calculator -i")
	fmt.Println("  cli-calculator")
	fmt.Println()
	fmt.Println("Interactive mode commands:")
	fmt.Println("  help, h        - Show help message")
	fmt.Println("  history        - Show calculation history")
	fmt.Println("  clear, c       - Clear calculation history")
	fmt.Println("  operations, ops- Show available operations")
	fmt.Println("  quit, q, exit  - Exit the calculator")
}
