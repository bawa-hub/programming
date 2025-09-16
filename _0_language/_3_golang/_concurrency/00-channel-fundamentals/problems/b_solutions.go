package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// ============================================================================
// ALTERNATING PATTERNS (1-5)
// ============================================================================

// Problem 1: Even-Odd Printing
func problem1() {
	fmt.Println("\n=== Problem 1: Even-Odd Printing ===")
	
	// Create channels for coordination
	evenCh := make(chan bool)
	oddCh := make(chan bool)
	
	// Even goroutine
	go func() {
		for i := 0; i < 6; i += 2 {
			<-evenCh // Wait for permission to print
			fmt.Printf("Even: %d\n", i)
			oddCh <- true // Signal odd goroutine to print
		}
	}()
	
	// Odd goroutine
	go func() {
		for i := 1; i < 6; i += 2 {
			<-oddCh // Wait for permission to print
			fmt.Printf("Odd: %d\n", i)
			evenCh <- true // Signal even goroutine to print
		}
	}()
	
	// Start the sequence
	evenCh <- true
	
	// Wait for completion
	time.Sleep(100 * time.Millisecond)
}

// Problem 2: A-B Alternating
func problem2() {
	fmt.Println("\n=== Problem 2: A-B Alternating ===")
	
	aCh := make(chan bool)
	bCh := make(chan bool)
	
	// A goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-aCh // Wait for permission
			fmt.Println("A")
			bCh <- true // Signal B
		}
	}()
	
	// B goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-bCh // Wait for permission
			fmt.Println("B")
			aCh <- true // Signal A
		}
	}()
	
	// Start the sequence
	aCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 3: 1-2-3 Sequence
func problem3() {
	fmt.Println("\n=== Problem 3: 1-2-3 Sequence ===")
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	
	// Goroutine 1
	go func() {
		for i := 0; i < 2; i++ {
			<-ch1
			fmt.Println("1")
			ch2 <- true
		}
	}()
	
	// Goroutine 2
	go func() {
		for i := 0; i < 2; i++ {
			<-ch2
			fmt.Println("2")
			ch3 <- true
		}
	}()
	
	// Goroutine 3
	go func() {
		for i := 0; i < 2; i++ {
			<-ch3
			fmt.Println("3")
			ch1 <- true
		}
	}()
	
	// Start the sequence
	ch1 <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 4: Color Alternating
func problem4() {
	fmt.Println("\n=== Problem 4: Color Alternating ===")
	
	redCh := make(chan bool)
	blueCh := make(chan bool)
	
	// Red goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-redCh
			fmt.Println("Red")
			blueCh <- true
		}
	}()
	
	// Blue goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-blueCh
			fmt.Println("Blue")
			redCh <- true
		}
	}()
	
	// Start the sequence
	redCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 5: Number-Letter Alternating
func problem5() {
	fmt.Println("\n=== Problem 5: Number-Letter Alternating ===")
	
	numberCh := make(chan bool)
	letterCh := make(chan bool)
	
	// Number goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			<-numberCh
			fmt.Printf("%d\n", i)
			letterCh <- true
		}
	}()
	
	// Letter goroutine
	go func() {
		for i := 0; i < 3; i++ {
			<-letterCh
			fmt.Printf("%c\n", 'A'+i)
			numberCh <- true
		}
	}()
	
	// Start the sequence
	numberCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// TURN-BASED OPERATIONS (6-10)
// ============================================================================

// Problem 6: Three-Player Game
func problem6() {
	fmt.Println("\n=== Problem 6: Three-Player Game ===")
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	
	// Player 1
	go func() {
		for i := 0; i < 2; i++ {
			<-ch1
			fmt.Println("Player 1's turn")
			ch2 <- true
		}
	}()
	
	// Player 2
	go func() {
		for i := 0; i < 2; i++ {
			<-ch2
			fmt.Println("Player 2's turn")
			ch3 <- true
		}
	}()
	
	// Player 3
	go func() {
		for i := 0; i < 2; i++ {
			<-ch3
			fmt.Println("Player 3's turn")
			ch1 <- true
		}
	}()
	
	// Start the game
	ch1 <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 7: Round-Robin Processing
func problem7() {
	fmt.Println("\n=== Problem 7: Round-Robin Processing ===")
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	
	// Worker 1
	go func() {
		for i := 1; i <= 6; i += 3 {
			<-ch1
			fmt.Printf("Worker 1: Item %d\n", i)
			ch2 <- true
		}
	}()
	
	// Worker 2
	go func() {
		for i := 2; i <= 6; i += 3 {
			<-ch2
			fmt.Printf("Worker 2: Item %d\n", i)
			ch3 <- true
		}
	}()
	
	// Worker 3
	go func() {
		for i := 3; i <= 6; i += 3 {
			<-ch3
			fmt.Printf("Worker 3: Item %d\n", i)
			ch1 <- true
		}
	}()
	
	// Start processing
	ch1 <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 8: Ordered Execution
func problem8() {
	fmt.Println("\n=== Problem 8: Ordered Execution ===")
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	
	// Function 1
	go func() {
		<-ch1
		fmt.Println("Function 1 executed")
		ch2 <- true
	}()
	
	// Function 2
	go func() {
		<-ch2
		fmt.Println("Function 2 executed")
		ch3 <- true
	}()
	
	// Function 3
	go func() {
		<-ch3
		fmt.Println("Function 3 executed")
	}()
	
	// Start execution
	ch1 <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 9: Sequential Steps
func problem9() {
	fmt.Println("\n=== Problem 9: Sequential Steps ===")
	
	stepA := make(chan bool)
	stepB := make(chan bool)
	stepC := make(chan bool)
	
	// Step A
	go func() {
		<-stepA
		fmt.Println("Step A")
		stepB <- true
	}()
	
	// Step B
	go func() {
		<-stepB
		fmt.Println("Step B")
		stepC <- true
	}()
	
	// Step C
	go func() {
		<-stepC
		fmt.Println("Step C")
	}()
	
	// Start steps
	stepA <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 10: Turn-Based Counter
func problem10() {
	fmt.Println("\n=== Problem 10: Turn-Based Counter ===")
	
	counter1Ch := make(chan bool)
	counter2Ch := make(chan bool)
	
	// Counter 1
	go func() {
		for i := 1; i <= 6; i += 2 {
			<-counter1Ch
			fmt.Printf("Counter 1: %d\n", i)
			counter2Ch <- true
		}
	}()
	
	// Counter 2
	go func() {
		for i := 2; i <= 6; i += 2 {
			<-counter2Ch
			fmt.Printf("Counter 2: %d\n", i)
			counter1Ch <- true
		}
	}()
	
	// Start counting
	counter1Ch <- true
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// SIGNAL COORDINATION (11-15)
// ============================================================================

// Problem 11: Start Signal
func problem11() {
	fmt.Println("\n=== Problem 11: Start Signal ===")
	
	startCh := make(chan bool)
	doneCh := make(chan bool)
	
	// Worker goroutine
	go func() {
		fmt.Println("Waiting for start signal...")
		<-startCh
		fmt.Println("Start signal received!")
		fmt.Println("Work started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Work completed")
		doneCh <- true
	}()
	
	// Send start signal after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		startCh <- true
	}()
	
	// Wait for completion
	<-doneCh
}

// Problem 12: Handshake Protocol
func problem12() {
	fmt.Println("\n=== Problem 12: Handshake Protocol ===")
	
	handshakeCh := make(chan bool)
	responseCh := make(chan bool)
	
	// Goroutine 1
	go func() {
		fmt.Println("Goroutine 1: Sending handshake")
		handshakeCh <- true
		<-responseCh
		fmt.Println("Goroutine 1: Received response")
		fmt.Println("Handshake complete")
	}()
	
	// Goroutine 2
	go func() {
		<-handshakeCh
		fmt.Println("Goroutine 2: Received handshake")
		fmt.Println("Goroutine 2: Sending response")
		responseCh <- true
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 13: Barrier Synchronization
func problem13() {
	fmt.Println("\n=== Problem 13: Barrier Synchronization ===")
	
	barrierCh := make(chan bool)
	proceedCh := make(chan bool)
	
	// Goroutine 1
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine 1: Reached barrier")
		barrierCh <- true
		<-proceedCh
		fmt.Println("Goroutine 1: Proceeding...")
	}()
	
	// Goroutine 2
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Goroutine 2: Reached barrier")
		barrierCh <- true
		<-proceedCh
		fmt.Println("Goroutine 2: Proceeding...")
	}()
	
	// Goroutine 3
	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Goroutine 3: Reached barrier")
		barrierCh <- true
		<-proceedCh
		fmt.Println("Goroutine 3: Proceeding...")
	}()
	
	// Wait for all to reach barrier
	<-barrierCh
	<-barrierCh
	<-barrierCh
	fmt.Println("All goroutines reached barrier")
	
	// Signal all to proceed
	proceedCh <- true
	proceedCh <- true
	proceedCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 14: Stop Signal
func problem14() {
	fmt.Println("\n=== Problem 14: Stop Signal ===")
	
	stopCh := make(chan bool)
	doneCh := make(chan bool)
	
	// Worker 1
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("Goroutine 1: Stopping")
				doneCh <- true
				return
			default:
				fmt.Println("Goroutine 1: Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	
	// Worker 2
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("Goroutine 2: Stopping")
				doneCh <- true
				return
			default:
				fmt.Println("Goroutine 2: Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	
	// Send stop signal after delay
	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Stop signal received")
		stopCh <- true
		stopCh <- true
	}()
	
	// Wait for both to stop
	<-doneCh
	<-doneCh
}

// Problem 15: Ready Signal
func problem15() {
	fmt.Println("\n=== Problem 15: Ready Signal ===")
	
	readyCh := make(chan bool)
	startCh := make(chan bool)
	
	// Goroutine 1
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine 1: Ready")
		readyCh <- true
		<-startCh
		fmt.Println("Goroutine 1: Starting...")
	}()
	
	// Goroutine 2
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Goroutine 2: Ready")
		readyCh <- true
		<-startCh
		fmt.Println("Goroutine 2: Starting...")
	}()
	
	// Goroutine 3
	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Goroutine 3: Ready")
		readyCh <- true
		<-startCh
		fmt.Println("Goroutine 3: Starting...")
	}()
	
	// Wait for all to be ready
	<-readyCh
	<-readyCh
	<-readyCh
	fmt.Println("All ready! Starting...")
	
	// Signal all to start
	startCh <- true
	startCh <- true
	startCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// DATA EXCHANGE (16-20)
// ============================================================================

// Problem 16: Ping-Pong
func problem16() {
	fmt.Println("\n=== Problem 16: Ping-Pong ===")
	
	pingCh := make(chan bool)
	pongCh := make(chan bool)
	
	// Ping goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-pingCh
			fmt.Println("Ping")
			pongCh <- true
		}
	}()
	
	// Pong goroutine
	go func() {
		for i := 0; i < 6; i++ {
			<-pongCh
			fmt.Println("Pong")
			pingCh <- true
		}
	}()
	
	// Start the game
	pingCh <- true
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 17: Data Relay
func problem17() {
	fmt.Println("\n=== Problem 17: Data Relay ===")
	
	stage1Ch := make(chan string)
	stage2Ch := make(chan string)
	stage3Ch := make(chan string)
	
	// Stage 1
	go func() {
		data := <-stage1Ch
		fmt.Printf("Stage 1: Processing %s\n", data)
		stage2Ch <- data
	}()
	
	// Stage 2
	go func() {
		data := <-stage2Ch
		fmt.Printf("Stage 2: Processing %s\n", data)
		stage3Ch <- data
	}()
	
	// Stage 3
	go func() {
		data := <-stage3Ch
		fmt.Printf("Stage 3: Processing %s\n", data)
		fmt.Println("Data relay complete")
	}()
	
	// Start the relay
	stage1Ch <- "data"
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 18: Request-Response
func problem18() {
	fmt.Println("\n=== Problem 18: Request-Response ===")
	
	requestCh := make(chan string)
	responseCh := make(chan string)
	
	// Server goroutine
	go func() {
		for i := 0; i < 2; i++ {
			request := <-requestCh
			fmt.Printf("Response: %s World\n", request)
			responseCh <- fmt.Sprintf("%s World", request)
		}
	}()
	
	// Client goroutine
	go func() {
		requests := []string{"Hello", "How are you?"}
		for _, req := range requests {
			fmt.Printf("Request: %s\n", req)
			requestCh <- req
			response := <-responseCh
			fmt.Printf("Response: %s\n", response)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 19: Data Pipeline
func problem19() {
	fmt.Println("\n=== Problem 19: Data Pipeline ===")
	
	inputCh := make(chan string)
	processCh := make(chan string)
	outputCh := make(chan string)
	
	// Input stage
	go func() {
		data := <-inputCh
		fmt.Printf("Stage 1: Input %s\n", data)
		processCh <- data
	}()
	
	// Process stage
	go func() {
		data := <-processCh
		fmt.Printf("Stage 2: Processing %s\n", data)
		outputCh <- data
	}()
	
	// Output stage
	go func() {
		data := <-outputCh
		fmt.Printf("Stage 3: Output %s\n", data)
		fmt.Println("Pipeline complete")
	}()
	
	// Start the pipeline
	inputCh <- "data"
	
	time.Sleep(100 * time.Millisecond)
}

// Problem 20: Message Passing
func problem20() {
	fmt.Println("\n=== Problem 20: Message Passing ===")
	
	messageCh := make(chan string)
	
	// Sender goroutine
	go func() {
		messages := []string{"Hello", "World", "Go", "Channels"}
		for i, msg := range messages {
			fmt.Printf("Message %d: %s\n", i+1, msg)
			messageCh <- msg
		}
		close(messageCh)
	}()
	
	// Receiver goroutine
	go func() {
		for msg := range messageCh {
			fmt.Printf("Received: %s\n", msg)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// ============================================================================
// MAIN FUNCTION
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
	
	if arg == "alternating" {
		runAlternatingProblems()
		return
	}
	
	if arg == "turn-based" {
		runTurnBasedProblems()
		return
	}
	
	if arg == "signals" {
		runSignalProblems()
		return
	}
	
	if arg == "data-exchange" {
		runDataExchangeProblems()
		return
	}
	
	// Try to parse as problem number
	problemNum, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid argument: %s\n", arg)
		showUsage()
		return
	}
	
	if problemNum < 1 || problemNum > 20 {
		fmt.Printf("Problem number must be between 1 and 20, got: %d\n", problemNum)
		showUsage()
		return
	}
	
	runProblem(problemNum)
}

func showUsage() {
	fmt.Println("🔗 Basic Channel Synchronization Problems")
	fmt.Println("=========================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run basic_sync_problems.go <problem_number>  # Run specific problem (1-20)")
	fmt.Println("  go run basic_sync_problems.go alternating       # Run alternating problems (1-5)")
	fmt.Println("  go run basic_sync_problems.go turn-based        # Run turn-based problems (6-10)")
	fmt.Println("  go run basic_sync_problems.go signals           # Run signal problems (11-15)")
	fmt.Println("  go run basic_sync_problems.go data-exchange     # Run data exchange problems (16-20)")
	fmt.Println("  go run basic_sync_problems.go all               # Run all problems")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run basic_sync_problems.go 1                 # Run problem 1 (Even-Odd)")
	fmt.Println("  go run basic_sync_problems.go 16                # Run problem 16 (Ping-Pong)")
	fmt.Println("  go run basic_sync_problems.go alternating       # Run alternating problems")
	fmt.Println("  go run basic_sync_problems.go all               # Run all 20 problems")
	fmt.Println()
	fmt.Println("Problem Categories:")
	fmt.Println("  Alternating (1-5):     Even/odd, A/B, sequences")
	fmt.Println("  Turn-Based (6-10):     Taking turns, round-robin")
	fmt.Println("  Signals (11-15):       Start/stop, handshakes")
	fmt.Println("  Data Exchange (16-20): Ping-pong, pipelines")
}

func runAllProblems() {
	fmt.Println("🚀 Running All 20 Basic Sync Problems")
	fmt.Println("=====================================")
	
	problems := getAllProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n🎉 All 20 problems completed!")
}

func runAlternatingProblems() {
	fmt.Println("🔄 Running Alternating Problems (1-5)")
	fmt.Println("====================================")
	
	problems := getAlternatingProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+1)
		problem()
	}
	
	fmt.Println("\n✅ Alternating problems completed!")
}

func runTurnBasedProblems() {
	fmt.Println("🔄 Running Turn-Based Problems (6-10)")
	fmt.Println("====================================")
	
	problems := getTurnBasedProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+6)
		problem()
	}
	
	fmt.Println("\n✅ Turn-based problems completed!")
}

func runSignalProblems() {
	fmt.Println("📡 Running Signal Problems (11-15)")
	fmt.Println("=================================")
	
	problems := getSignalProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+11)
		problem()
	}
	
	fmt.Println("\n✅ Signal problems completed!")
}

func runDataExchangeProblems() {
	fmt.Println("🔄 Running Data Exchange Problems (16-20)")
	fmt.Println("=======================================")
	
	problems := getDataExchangeProblems()
	
	for i, problem := range problems {
		fmt.Printf("\n--- Problem %d ---\n", i+16)
		problem()
	}
	
	fmt.Println("\n✅ Data exchange problems completed!")
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
	}
}

func getAlternatingProblems() []func() {
	return []func(){
		problem1, problem2, problem3, problem4, problem5,
	}
}

func getTurnBasedProblems() []func() {
	return []func(){
		problem6, problem7, problem8, problem9, problem10,
	}
}

func getSignalProblems() []func() {
	return []func(){
		problem11, problem12, problem13, problem14, problem15,
	}
}

func getDataExchangeProblems() []func() {
	return []func(){
		problem16, problem17, problem18, problem19, problem20,
	}
}
