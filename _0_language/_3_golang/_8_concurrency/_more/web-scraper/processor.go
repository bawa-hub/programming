package main

import (
	"regexp"
	"strings"
	"time"
)

// DataProcessor processes scraped data
type DataProcessor struct{}

// NewDataProcessor creates a new data processor
func NewDataProcessor() *DataProcessor {
	return &DataProcessor{}
}

// Process processes scraped data and returns structured data
func (dp *DataProcessor) Process(result ScrapeResult) ScrapedData {
	data := ScrapedData{
		URL:         result.URL,
		Title:       result.Title,
		ProcessedAt: time.Now(),
	}
	
	if result.Status == "success" {
		// Count words
		words := strings.Fields(result.Content)
		data.WordCount = len(words)
		
		// Extract keywords (simple approach)
		data.Keywords = dp.extractKeywords(result.Content)
	}
	
	return data
}

// extractKeywords extracts keywords from content
func (dp *DataProcessor) extractKeywords(content string) []string {
	// Simple keyword extraction - find words that appear frequently
	words := strings.Fields(strings.ToLower(content))
	wordCount := make(map[string]int)
	
	// Count word frequency
	for _, word := range words {
		// Remove punctuation and short words
		cleanWord := regexp.MustCompile(`[^a-zA-Z]`).ReplaceAllString(word, "")
		if len(cleanWord) > 3 {
			wordCount[cleanWord]++
		}
	}
	
	// Get top 5 most frequent words
	var keywords []string
	for word, count := range wordCount {
		if count > 2 && len(keywords) < 5 {
			keywords = append(keywords, word)
		}
	}
	
	return keywords
}
