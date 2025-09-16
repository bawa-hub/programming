package main

import (
	"fmt"
	"os"
	"strconv"
)

// Import problem functions from solutions.go
// Note: In a real project, you would put these in a separate package

// ============================================================================
// PRACTICE PROBLEM RUNNER
// ============================================================================

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	arg := os.Args[1]
	
	if arg == "all" {
		runAllProblems()
		return
	}
	
	if arg == "easy" {
		runEasyProblems()
		return
	}
	
	if arg == "medium" {
		runMediumProblems()
		return
	}
	
	if arg == "advanced" {
		runAdvancedProblems()
		return
	}
	
	// Try to parse as problem number
	problemNum, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid argument: %s\n", arg)
		showUsage()
		return
	}
	
	if problemNum < 1 || problemNum > 50 {
		fmt.Printf("Problem number must be between 1 and 50, got: %d\n", problemNum)
		showUsage()
		return
	}
	
	runProblem(problemNum)
}

func showUsage() {
	fmt.Println("🔗 Go Channels Practice Problems")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run practice.go <problem_number>  # Run specific problem (1-50)")
	fmt.Println("  go run practice.go easy              # Run easy problems (1-15)")
	fmt.Println("  go run practice.go medium            # Run medium problems (16-35)")
	fmt.Println("  go run practice.go advanced          # Run advanced problems (36-50)")
	fmt.Println("  go run practice.go all               # Run all problems")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run practice.go 1                 # Run problem 1")
	fmt.Println("  go run practice.go 25                # Run problem 25")
	fmt.Println("  go run practice.go easy              # Run all easy problems")
	fmt.Println("  go run practice.go all               # Run all 50 problems")
	fmt.Println()
	fmt.Println("Problem Levels:")
	fmt.Println("  Easy (1-15):     Basic channel operations")
	fmt.Println("  Medium (16-35):  Buffered channels, select statements")
	fmt.Println("  Advanced (36-50): Complex channel interactions")
}

func runAllProblems() {
	fmt.Println("🚀 Running All 50 Problems")
	fmt.Println("==========================")
	
	problems := getAllProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n🎉 All 50 problems completed!")
}

func runEasyProblems() {
	fmt.Println("🟢 Running Easy Problems (1-15)")
	fmt.Println("===============================")
	
	problems := getEasyProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n✅ Easy problems completed!")
}

func runMediumProblems() {
	fmt.Println("🟡 Running Medium Problems (16-35)")
	fmt.Println("=================================")
	
	problems := getMediumProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+16)
		problem()
	}
	
	fmt.Println("\n✅ Medium problems completed!")
}

func runAdvancedProblems() {
	fmt.Println("🔴 Running Advanced Problems (36-50)")
	fmt.Println("===================================")
	
	problems := getAdvancedProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+36)
		problem()
	}
	
	fmt.Println("\n✅ Advanced problems completed!")
}

func runProblem(problemNum int) {
	fmt.Printf("🔗 Running Problem %d\n", problemNum)
	fmt.Println("==================")
	
	problems := getAllProblems()
	
	if problemNum < 1 || problemNum > len(problems) {
		fmt.Printf("Problem %d not found\n", problemNum)
		return
	}
	
	problems[problemNum-1]()
	fmt.Printf("\n✅ Problem %d completed!\n", problemNum)
}

func getAllProblems() []func() {
	return []func(){
		problem1, problem2, problem3, problem4, problem5,
		problem6, problem7, problem8, problem9, problem10,
		problem11, problem12, problem13, problem14, problem15,
		problem16, problem17, problem18, problem19, problem20,
		problem21, problem22, problem23, problem24, problem25,
		problem26, problem27, problem28, problem29, problem30,
		problem31, problem32, problem33, problem34, problem35,
		problem36, problem37, problem38, problem39, problem40,
		problem41, problem42, problem43, problem44, problem45,
		problem46, problem47, problem48, problem49, problem50,
	}
}

func getEasyProblems() []func() {
	return []func(){
		problem1, problem2, problem3, problem4, problem5,
		problem6, problem7, problem8, problem9, problem10,
		problem11, problem12, problem13, problem14, problem15,
	}
}

func getMediumProblems() []func() {
	return []func(){
		problem16, problem17, problem18, problem19, problem20,
		problem21, problem22, problem23, problem24, problem25,
		problem26, problem27, problem28, problem29, problem30,
		problem31, problem32, problem33, problem34, problem35,
	}
}

func getAdvancedProblems() []func() {
	return []func(){
		problem36, problem37, problem38, problem39, problem40,
		problem41, problem42, problem43, problem44, problem45,
		problem46, problem47, problem48, problem49, problem50,
	}
}
