package main

import (
	"time"
)

// ScrapeJob represents a URL to be scraped
type ScrapeJob struct {
	ID        string
	URL       string
	Priority  int
	CreatedAt time.Time
	Result    chan ScrapeResult
}

// ScrapeResult contains the result of a scraping operation
type ScrapeResult struct {
	JobID     string
	URL       string
	Title     string
	Content   string
	Status    string
	Error     error
	Duration  time.Duration
	Timestamp time.Time
}

// ScrapedData represents processed scraped data
type ScrapedData struct {
	URL       string
	Title     string
	WordCount int
	Keywords  []string
	ProcessedAt time.Time
}

// Event represents a system event
type Event struct {
	Type      string
	Message   string
	Timestamp time.Time
	Data      interface{}
}
