package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// main is the entry point for the analytics client
func main() {
	// Parse command line flags
	serverURL := flag.String("server", "http://localhost:8080", "Analytics server URL")
	eventType := flag.String("type", "page_view", "Event type")
	count := flag.Int("count", 10, "Number of events to send")
	rate := flag.Int("rate", 1, "Events per second")
	flag.Parse()

	fmt.Printf("Sending %d %s events to %s at %d events/second\n", 
		*count, *eventType, *serverURL, *rate)

	// Create events
	events := createEvents(*eventType, *count)

	// Send events
	if err := sendEvents(*serverURL, events, *rate); err != nil {
		fmt.Printf("Error sending events: %v\n", err)
		return
	}

	fmt.Println("Events sent successfully!")
}

// createEvents creates a list of events
func createEvents(eventType string, count int) []models.Event {
	events := make([]models.Event, count)
	
	for i := 0; i < count; i++ {
		event := models.NewEvent(eventType, "client", map[string]interface{}{
			"page": fmt.Sprintf("/page-%d", i%5),
			"user_id": fmt.Sprintf("user-%d", i%10),
		})
		
		events[i] = *event
	}
	
	return events
}

// sendEvents sends events to the server
func sendEvents(serverURL string, events []models.Event, rate int) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	// Calculate delay between requests
	delay := time.Second / time.Duration(rate)
	
	for i, event := range events {
		// Send single event
		if err := sendEvent(client, serverURL, event); err != nil {
			return fmt.Errorf("failed to send event %d: %w", i, err)
		}
		
		// Wait between requests
		if i < len(events)-1 {
			time.Sleep(delay)
		}
	}
	
	return nil
}

// sendEvent sends a single event to the server
func sendEvent(client *http.Client, serverURL string, event models.Event) error {
	// Marshal event to JSON
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	
	// Create request
	req, err := http.NewRequest("POST", serverURL+"/api/v1/events", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	
	// Check response
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}
	
	return nil
}

// sendBatch sends events in batches
func sendBatch(client *http.Client, serverURL string, events []models.Event) error {
	// Marshal events to JSON
	jsonData, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("failed to marshal events: %w", err)
	}
	
	// Create request
	req, err := http.NewRequest("POST", serverURL+"/api/v1/events/batch", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	
	// Check response
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}
	
	return nil
}
