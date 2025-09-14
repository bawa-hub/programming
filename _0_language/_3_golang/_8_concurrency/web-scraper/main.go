package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Concurrent Web Scraper ===")
	
	// Create components
	scraper := NewWebScraper()
	processor := NewDataProcessor()
	pubsub := NewPubSub()
	workerPool := NewWorkerPool(3, scraper, processor, pubsub)
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Start worker pool
	workerPool.Start(ctx)
	
	// Subscribe to events
	workerEvents := pubsub.Subscribe("worker")
	resultEvents := pubsub.Subscribe("results")
	dataEvents := pubsub.Subscribe("data")
	
	// Start event handlers
	go handleEvents("Worker", workerEvents)
	go handleEvents("Results", resultEvents)
	go handleEvents("Data", dataEvents)
	
	// Start data collector
	go collectData(workerPool.GetProcessedData())
	
	// Add some test URLs
	testURLs := []string{
		"https://httpbin.org/html",
		"https://httpbin.org/json",
		"https://httpbin.org/xml",
		"https://httpbin.org/robots.txt",
		"https://httpbin.org/user-agent",
	}
	
	// Add jobs
	for i, url := range testURLs {
		job := ScrapeJob{
			ID:        fmt.Sprintf("job-%d", i+1),
			URL:       url,
			Priority:  1,
			CreatedAt: time.Now(),
			Result:    make(chan ScrapeResult, 1),
		}
		
		workerPool.AddJob(job)
		fmt.Printf("Added job: %s\n", job.ID)
	}
	
	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	select {
	case <-sigChan:
		fmt.Println("\nReceived interrupt signal, shutting down...")
		cancel()
	case <-ctx.Done():
		fmt.Println("Context cancelled, shutting down...")
	}
	
	// Cleanup
	workerPool.Stop()
	scraper.Stop()
	
	fmt.Println("Web scraper stopped!")
}

// handleEvents handles events from the pub/sub system
func handleEvents(eventType string, events <-chan Event) {
	for event := range events {
		fmt.Printf("[%s] %s: %s\n", eventType, event.Timestamp.Format("15:04:05"), event.Message)
	}
}

// collectData collects and displays processed data
func collectData(data <-chan ScrapedData) {
	for data := range data {
		fmt.Printf("\n=== Processed Data ===\n")
		fmt.Printf("URL: %s\n", data.URL)
		fmt.Printf("Title: %s\n", data.Title)
		fmt.Printf("Word Count: %d\n", data.WordCount)
		fmt.Printf("Keywords: %v\n", data.Keywords)
		fmt.Printf("Processed At: %s\n", data.ProcessedAt.Format("15:04:05"))
		fmt.Printf("====================\n\n")
	}
}

// testCircuitBreaker demonstrates the circuit breaker functionality
func testCircuitBreaker() {
	fmt.Println("\n=== Testing Circuit Breaker ===")
	
	// Create a scraper with circuit breaker
	scraper := NewWebScraper()
	
	// Test with a URL that will fail
	failingURL := "https://httpbin.org/status/500"
	
	job := ScrapeJob{
		ID:        "test-job",
		URL:       failingURL,
		Priority:  1,
		CreatedAt: time.Now(),
		Result:    make(chan ScrapeResult, 1),
	}
	
	ctx := context.Background()
	
	// Try multiple times to trigger circuit breaker
	for i := 1; i <= 5; i++ {
		fmt.Printf("\nAttempt %d:\n", i)
		result := scraper.Scrape(ctx, job)
		fmt.Printf("Status: %s\n", result.Status)
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		}
		
		// Check circuit breaker state
		state := scraper.circuitBreaker.GetState()
		var stateStr string
		switch state {
		case StateClosed:
			stateStr = "CLOSED"
		case StateOpen:
			stateStr = "OPEN"
		case StateHalfOpen:
			stateStr = "HALF-OPEN"
		}
		fmt.Printf("Circuit State: %s\n", stateStr)
		
		time.Sleep(1 * time.Second)
	}
	
	scraper.Stop()
	fmt.Println("\nCircuit breaker test completed!")
}
