package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// 🕷️ CONCURRENT WEB SCRAPER PROJECT
// A real-world application demonstrating advanced goroutine patterns

type Scraper struct {
	client      *http.Client
	concurrency int
	results     chan ScrapeResult
	errors      chan error
	shutdown    chan struct{}
	wg          sync.WaitGroup
}

type ScrapeResult struct {
	URL        string
	StatusCode int
	Size       int64
	Duration   time.Duration
	Error      error
}

func main() {
	fmt.Println("🕷️ CONCURRENT WEB SCRAPER")
	fmt.Println("==========================")

	// URLs to scrape
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/3",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/4",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
	}

	// Create scraper with 3 concurrent workers
	scraper := NewScraper(3)
	defer scraper.Close()

	// Start scraping
	start := time.Now()
	scraper.ScrapeURLs(urls)

	// Collect results
	results := scraper.CollectResults()
	duration := time.Since(start)

	// Display results
	fmt.Printf("\n📊 SCRAPING RESULTS (Completed in %v)\n", duration)
	fmt.Println("=====================================")
	
	successCount := 0
	totalSize := int64(0)
	
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("❌ %s - Error: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("✅ %s - Status: %d, Size: %d bytes, Duration: %v\n", 
				result.URL, result.StatusCode, result.Size, result.Duration)
			successCount++
			totalSize += result.Size
		}
	}
	
	fmt.Printf("\n📈 SUMMARY:\n")
	fmt.Printf("  Total URLs: %d\n", len(urls))
	fmt.Printf("  Successful: %d\n", successCount)
	fmt.Printf("  Failed: %d\n", len(urls)-successCount)
	fmt.Printf("  Total Size: %d bytes\n", totalSize)
	fmt.Printf("  Average Time per URL: %v\n", duration/time.Duration(len(urls)))
	fmt.Printf("  Goroutines used: %d\n", runtime.NumGoroutine())
}

// NewScraper creates a new web scraper with specified concurrency
func NewScraper(concurrency int) *Scraper {
	return &Scraper{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		concurrency: concurrency,
		results:     make(chan ScrapeResult, concurrency*2),
		errors:      make(chan error, concurrency*2),
		shutdown:    make(chan struct{}),
	}
}

// ScrapeURLs starts scraping the provided URLs concurrently
func (s *Scraper) ScrapeURLs(urls []string) {
	// Create work channel
	work := make(chan string, len(urls))
	
	// Fill work channel
	go func() {
		defer close(work)
		for _, url := range urls {
			work <- url
		}
	}()
	
	// Start worker goroutines
	for i := 0; i < s.concurrency; i++ {
		s.wg.Add(1)
		go s.worker(i, work)
	}
	
	// Close results when all workers are done
	go func() {
		s.wg.Wait()
		close(s.results)
	}()
}

// worker processes URLs from the work channel
func (s *Scraper) worker(id int, work <-chan string) {
	defer s.wg.Done()
	
	for url := range work {
		select {
		case <-s.shutdown:
			return
		default:
			result := s.scrapeURL(url)
			s.results <- result
		}
	}
}

// scrapeURL scrapes a single URL
func (s *Scraper) scrapeURL(url string) ScrapeResult {
	start := time.Now()
	
	resp, err := s.client.Get(url)
	if err != nil {
		return ScrapeResult{
			URL:      url,
			Error:    err,
			Duration: time.Since(start),
		}
	}
	defer resp.Body.Close()
	
	// Read response body to get size
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ScrapeResult{
			URL:        url,
			StatusCode: resp.StatusCode,
			Error:      err,
			Duration:   time.Since(start),
		}
	}
	
	return ScrapeResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		Size:       int64(len(body)),
		Duration:   time.Since(start),
	}
}

// CollectResults collects all results from the scraper
func (s *Scraper) CollectResults() []ScrapeResult {
	var results []ScrapeResult
	
	for result := range s.results {
		results = append(results, result)
	}
	
	return results
}

// Close gracefully shuts down the scraper
func (s *Scraper) Close() {
	close(s.shutdown)
	s.wg.Wait()
}
