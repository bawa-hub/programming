package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// WorkerPool manages a pool of workers for scraping
type WorkerPool struct {
	scraper   *WebScraper
	processor *DataProcessor
	pubsub    *PubSub
	jobs      chan ScrapeJob
	processed chan ScrapedData
	wg        sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int, scraper *WebScraper, processor *DataProcessor, pubsub *PubSub) *WorkerPool {
	return &WorkerPool{
		scraper:   scraper,
		processor: processor,
		pubsub:    pubsub,
		jobs:      make(chan ScrapeJob, 100),
		processed: make(chan ScrapedData, 100),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start(ctx context.Context) {
	// Start workers
	for i := 0; i < 3; i++ {
		wp.wg.Add(1)
		go wp.worker(ctx, i+1)
	}
}

// worker processes jobs from the job queue
func (wp *WorkerPool) worker(ctx context.Context, workerID int) {
	defer wp.wg.Done()
	
	for {
		select {
		case job := <-wp.jobs:
			wp.pubsub.Publish("worker", Event{
				Type:      "job_started",
				Message:   fmt.Sprintf("Worker %d started job %s", workerID, job.ID),
				Timestamp: time.Now(),
				Data:      job,
			})
			
			result := wp.scraper.Scrape(ctx, job)
			
			wp.pubsub.Publish("results", Event{
				Type:      "result_processed",
				Message:   fmt.Sprintf("Result processed for %s", result.URL),
				Timestamp: time.Now(),
				Data:      result,
			})
			
			// Process data if successful
			if result.Status == "success" {
				data := wp.processor.Process(result)
				wp.processed <- data
				
				wp.pubsub.Publish("data", Event{
					Type:      "data_processed",
					Message:   fmt.Sprintf("Data processed for %s", data.URL),
					Timestamp: time.Now(),
					Data:      data,
				})
			}
			
			wp.pubsub.Publish("worker", Event{
				Type:      "job_completed",
				Message:   fmt.Sprintf("Worker %d completed job %s", workerID, job.ID),
				Timestamp: time.Now(),
				Data:      result,
			})
			
		case <-ctx.Done():
			fmt.Printf("Worker %d: Shutting down\n", workerID)
			return
		}
	}
}

// AddJob adds a job to the queue
func (wp *WorkerPool) AddJob(job ScrapeJob) {
	wp.jobs <- job
}

// GetProcessedData returns a channel for processed data
func (wp *WorkerPool) GetProcessedData() <-chan ScrapedData {
	return wp.processed
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.processed)
}
