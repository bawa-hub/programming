package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// WebScraper handles web scraping operations
type WebScraper struct {
	client        *http.Client
	circuitBreaker *CircuitBreaker
	rateLimiter   *RateLimiter
}

// NewWebScraper creates a new web scraper
func NewWebScraper() *WebScraper {
	return &WebScraper{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		circuitBreaker: NewCircuitBreaker(3, 30*time.Second),
		rateLimiter:    NewRateLimiter(2, 1*time.Second), // 2 requests per second
	}
}

// Scrape scrapes a URL and returns the result
func (ws *WebScraper) Scrape(ctx context.Context, job ScrapeJob) ScrapeResult {
	start := time.Now()
	result := ScrapeResult{
		JobID:     job.ID,
		URL:       job.URL,
		Timestamp: time.Now(),
	}
	
	// Wait for rate limiter
	ws.rateLimiter.Wait()
	
	// Use circuit breaker
	err := ws.circuitBreaker.Call(func() error {
		// Check context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		// Create request with context
		req, err := http.NewRequestWithContext(ctx, "GET", job.URL, nil)
		if err != nil {
			return err
		}
		
		// Set user agent
		req.Header.Set("User-Agent", "Go-WebScraper/1.0")
		
		// Make request
		resp, err := ws.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		
		// Check status code
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
		}
		
		// Read body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		
		// Extract title
		title := ws.extractTitle(string(body))
		
		result.Title = title
		result.Content = string(body)
		result.Status = "success"
		
		return nil
	})
	
	result.Duration = time.Since(start)
	
	if err != nil {
		result.Error = err
		result.Status = "failed"
	}
	
	return result
}

// extractTitle extracts the title from HTML content
func (ws *WebScraper) extractTitle(html string) string {
	re := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return "No title found"
}

// Stop stops the scraper
func (ws *WebScraper) Stop() {
	ws.rateLimiter.Stop()
}
