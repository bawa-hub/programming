package exercises

import (
	"fmt"
	"time"
)

// Exercise 3: Goroutine Pool
// Implement a worker pool with 3 workers processing 10 jobs.
func Exercise3() {
	fmt.Println("\nExercise 3: Goroutine Pool")
	fmt.Println("==========================")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start workers
	for i := 0; i < numWorkers; i++ {
		go exerciseWorker(i, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job %d processed, result: %d\n", r, result)
	}
}

func exerciseWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * job // Square the job number
	}
}