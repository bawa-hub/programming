package main

import (
	"fmt"
	"time"
)

// Worker function that processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(time.Duration(job) * 200 * time.Millisecond)
		results <- job * job // Return square of the job
	}
}

func main() {
	fmt.Println("=== Fan-out Pattern with Select ===")
	
	// Create channels
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	
	// Send jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
			fmt.Printf("Sent job %d\n", j)
			time.Sleep(100 * time.Millisecond)
		}
		close(jobs)
	}()
	
	// Collect results as they come in
	resultsReceived := 0
	for resultsReceived < 5 {
		select {
		case result := <-results:
			fmt.Printf("Result received: %d\n", result)
			resultsReceived++
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout waiting for results!")
			break
		}
	}
	
	fmt.Println("All results collected!")
}
