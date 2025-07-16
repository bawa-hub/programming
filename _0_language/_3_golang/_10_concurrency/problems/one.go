// 🔥 Problem Statement

// You’re building a simulation where 100 workers are doing a task (e.g., "Processing task #x").
// Your job:
//     Launch 100 goroutines, each doing a small task.
//     Ensure they all finish before your program exits.
//     Print the task number each goroutine is handling.

// 🛠️ Your Tasks

//     Create a function doWork(id int) that prints:

//     Goroutine <id> is processing

//     Use a loop in main() to launch 100 goroutines, each with a unique ID.

//     Use sync.WaitGroup to wait for all 100 goroutines to finish.

package main

import (
	"fmt"
	"sync"
)

func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signal this goroutine is done
	fmt.Printf("Goroutine %d is processing\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		// Avoid closure bug by passing i as argument to goroutine
		go doWork(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines done!")
}

// ⚠️ Why This Works

//     wg.Add(1) adds one goroutine to wait for.

//     defer wg.Done() ensures even if the function panics or returns early, the wait group gets signaled.

//     We pass i as an argument to doWork() — this avoids the classic trap where goroutines close over the same loop variable (i).

// 📌 Sample Output (order may vary):

// Goroutine 3 is processing
// Goroutine 1 is processing
// Goroutine 7 is processing
// ...
// All goroutines done!



// 🧠 Interview Insight:
//     "Why pass the loop variable as a function argument?" → Because goroutines share the loop variable and its memory location, you’d otherwise get all goroutines printing the same value (usually 100). Passing i by value avoids this trap.