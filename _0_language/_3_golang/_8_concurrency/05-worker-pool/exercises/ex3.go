package exercises

import (
	"fmt"
	"time"
)

// Exercise 3: Dynamic Worker Pool
// Create a worker pool that can adjust the number of workers based on workload.
func Exercise3() {
	fmt.Println("\nExercise 3: Dynamic Worker Pool")
	fmt.Println("==============================")
	
	const minWorkers = 2
	const maxWorkers = 4
	const numJobs = 12
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Dynamic worker pool
	pool := &ExerciseDynamicWorkerPool{
		workers:    minWorkers,
		minWorkers: minWorkers,
		maxWorkers: maxWorkers,
		jobs:       jobs,
		results:    results,
	}
	
	// Start initial workers
	pool.startWorkers()
	
	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Dynamic Job %d", i),
			}
			
			// Adjust workers based on queue size
			pool.adjustWorkers()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		pool.wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("Exercise 3 Results:")
	for result := range results {
		fmt.Printf("  Job %d: %s (took %v, worker %d)\n", 
			result.JobID, result.Data, result.Duration, result.WorkerID)
	}
}

type ExerciseDynamicWorkerPool struct {
	workers    int
	minWorkers int
	maxWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.RWMutex
}

func (dwp *ExerciseDynamicWorkerPool) startWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	for i := 0; i < dwp.workers; i++ {
		dwp.wg.Add(1)
		go dwp.worker(i)
	}
}

func (dwp *ExerciseDynamicWorkerPool) worker(id int) {
	defer dwp.wg.Done()
	
	for job := range dwp.jobs {
		start := time.Now()
		
		// Simulate work
		time.Sleep(time.Duration(job.ID*50) * time.Millisecond)
		
		result := Result{
			JobID:    job.ID,
			Data:     fmt.Sprintf("Dynamic: %s", job.Data),
			Duration: time.Since(start),
			WorkerID: id,
		}
		
		dwp.results <- result
		fmt.Printf("Worker %d processed dynamic job %d\n", id, job.ID)
	}
}

func (dwp *ExerciseDynamicWorkerPool) adjustWorkers() {
	dwp.mu.Lock()
	defer dwp.mu.Unlock()
	
	queueSize := len(dwp.jobs)
	
	if queueSize > 2 && dwp.workers < dwp.maxWorkers {
		// Add worker
		dwp.workers++
		dwp.wg.Add(1)
		go dwp.worker(dwp.workers - 1)
		fmt.Printf("Added worker, total: %d\n", dwp.workers)
	}
}