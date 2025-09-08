package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID     int
	Data   string
	Result chan string
}

// Worker processes jobs from the job queue
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job %d (%s)\n", id, job.ID, job.Data)
		
		// Simulate work
		time.Sleep(time.Duration(job.ID) * 200 * time.Millisecond)
		
		// Send result back
		result := fmt.Sprintf("Job %d completed by worker %d", job.ID, id)
		job.Result <- result
		close(job.Result)
	}
	
	fmt.Printf("Worker %d: Finished\n", id)
}

func main() {
	fmt.Println("=== Worker Pool Pattern ===")
	
	// Configuration
	numWorkers := 3
	numJobs := 5
	
	// Create channels
	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup
	
	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	
	// Create jobs and send them
	jobResults := make([]Job, numJobs)
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:     i,
			Data:   fmt.Sprintf("data-%d", i),
			Result: make(chan string, 1),
		}
		jobResults[i-1] = job
		jobs <- job
		fmt.Printf("Sent job %d\n", i)
	}
	close(jobs)
	
	// Collect results
	fmt.Println("\nCollecting results:")
	for i, job := range jobResults {
		result := <-job.Result
		fmt.Printf("Result %d: %s\n", i+1, result)
	}
	
	// Wait for all workers to finish
	wg.Wait()
	
	fmt.Println("All jobs completed!")
}
