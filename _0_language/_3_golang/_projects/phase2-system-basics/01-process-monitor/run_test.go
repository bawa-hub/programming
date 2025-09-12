package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		testType = flag.String("test", "basic", "Test type: basic, interactive, performance, full")
	)
	flag.Parse()
	
	fmt.Println("🚀 Process Manager - Test Runner")
	fmt.Println("================================")
	
	switch *testType {
	case "basic":
		RunBasicTest()
	case "interactive":
		RunInteractiveTest()
	case "performance":
		RunPerformanceTest()
	case "full":
		TestProcessManager()
	default:
		fmt.Println("Available test types:")
		fmt.Println("  basic       - Basic functionality test")
		fmt.Println("  interactive - Interactive process test")
		fmt.Println("  performance - Performance benchmark test")
		fmt.Println("  full        - Complete test suite")
		fmt.Println()
		fmt.Println("Usage: go run run_test.go -test=basic")
		os.Exit(1)
	}
}
