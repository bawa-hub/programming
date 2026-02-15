package cli

import (
	"bufio"
	"cli-calculator/internal/calculator"
	"cli-calculator/pkg/errors"
	"fmt"
	"os"
	"strings"
)

// CLI represents the command-line interface
type CLI struct {
	calculator *calculator.Calculator
	scanner    *bufio.Scanner
}

// New creates a new CLI instance
func New() *CLI {
	return &CLI{
		calculator: calculator.New(),
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

// Run starts the CLI application
func (cli *CLI) Run() {
	cli.printWelcome()
	cli.printHelp()

	for {
		fmt.Print("calc> ")
		if !cli.scanner.Scan() {
			break
		}

		input := strings.TrimSpace(cli.scanner.Text())
		if input == "" {
			continue
		}

		// Handle special commands
		if cli.handleSpecialCommands(input) {
			continue
		}

		// Process calculation
		result, err := cli.calculator.ParseAndCalculate(input)
		if err != nil {
			cli.printError(err)
			continue
		}

		fmt.Printf("Result: %.6f\n", result)
	}
}

// handleSpecialCommands handles special CLI commands
func (cli *CLI) handleSpecialCommands(input string) bool {
	switch input {
	case "help", "h":
		cli.printHelp()
		return true
	case "history":
		cli.printHistory()
		return true
	case "clear", "c":
		cli.calculator.ClearHistory()
		fmt.Println("History cleared.")
		return true
	case "quit", "q", "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)
	case "operations", "ops":
		cli.printOperations()
		return true
	}
	return false
}

// printWelcome prints the welcome message
func (cli *CLI) printWelcome() {
	fmt.Println("🧮 Welcome to the Go CLI Calculator!")
	fmt.Println("Type 'help' for available commands or enter a mathematical expression.")
	fmt.Println()
}

// printHelp prints the help message
func (cli *CLI) printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help, h        - Show this help message")
	fmt.Println("  history        - Show calculation history")
	fmt.Println("  clear, c       - Clear calculation history")
	fmt.Println("  operations, ops- Show available operations")
	fmt.Println("  quit, q, exit  - Exit the calculator")
	fmt.Println()
	fmt.Println("Available operations:")
	fmt.Println("  +  - Addition")
	fmt.Println("  -  - Subtraction")
	fmt.Println("  *  - Multiplication")
	fmt.Println("  /  - Division")
	fmt.Println("  ^  - Exponentiation")
	fmt.Println("  %  - Modulo")
	fmt.Println("  √  - Square root (prefix)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  2 + 3")
	fmt.Println("  10 * 5")
	fmt.Println("  √25")
	fmt.Println("  2 ^ 8")
	fmt.Println()
}

// printHistory prints the calculation history
func (cli *CLI) printHistory() {
	history := cli.calculator.GetHistory()
	if len(history) == 0 {
		fmt.Println("No calculations in history.")
		return
	}

	fmt.Println("Calculation History:")
	for i, entry := range history {
		fmt.Printf("%d. %s\n", i+1, entry)
	}
	fmt.Println()
}

// printOperations prints available operations
func (cli *CLI) printOperations() {
	operations := cli.calculator.GetAvailableOperations()
	fmt.Println("Available Operations:")
	for symbol, op := range operations {
		fmt.Printf("  %s - %s\n", symbol, op.Symbol())
	}
	fmt.Println()
}

// printError prints an error message
func (cli *CLI) printError(err error) {
	if calcErr, ok := err.(*errors.CalculatorError); ok {
		fmt.Printf("❌ Error (%s): %s\n", calcErr.Type, calcErr.Message)
	} else {
		fmt.Printf("❌ Error: %s\n", err.Error())
	}
}

// RunBatch runs the calculator in batch mode (non-interactive)
func (cli *CLI) RunBatch(expression string) {
	result, err := cli.calculator.ParseAndCalculate(expression)
	if err != nil {
		cli.printError(err)
		os.Exit(1)
	}
	fmt.Printf("%.6f\n", result)
}
