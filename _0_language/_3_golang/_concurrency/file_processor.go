package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ============================================================================
// 1. BASIC CHANNEL CONCEPTS - File Processing
// ============================================================================

// FileInfo represents information about a processed file
type FileInfo struct {
	Path      string
	WordCount int
	LineCount int
	CharCount int
	Error     error
}

// Job represents a file processing job
type Job struct {
	FilePath string
	ID       int
}

// Result represents the result of processing a file
type Result struct {
	Job      Job
	FileInfo FileInfo
}

// ============================================================================
// 2. CHANNEL TYPES AND PATTERNS
// ============================================================================

// Global channels for the file processor
var (
	// Unbuffered channels for synchronization
	jobQueue    = make(chan Job)    // Jobs to be processed
	resultQueue = make(chan Result) // Results from processing
	
	// Buffered channels for performance
	workerPool = make(chan struct{}, 3) // Limit to 3 concurrent workers
	
	// Channel for system shutdown
	shutdown = make(chan bool)
)

// ============================================================================
// 3. FILE PROCESSOR - Demonstrates Channel Communication
// ============================================================================

// FileProcessor manages file processing using channels
type FileProcessor struct {
	workerCount int
	results     []Result
}

func NewFileProcessor(workerCount int) *FileProcessor {
	return &FileProcessor{
		workerCount: workerCount,
		results:     make([]Result, 0),
	}
}

// ProcessFiles demonstrates channel-based file processing
func (fp *FileProcessor) ProcessFiles(filePaths []string) {
	fmt.Println("🚀 Starting File Processor")
	fmt.Printf("📁 Processing %d files with %d workers\n", len(filePaths), fp.workerCount)
	fmt.Println()
	
	// Start workers
	for i := 0; i < fp.workerCount; i++ {
		go fp.worker(i)
	}
	
	// Start result collector
	go fp.collectResults()
	
	// Send jobs to queue
	go func() {
		for i, filePath := range filePaths {
			job := Job{
				FilePath: filePath,
				ID:       i + 1,
			}
			jobQueue <- job
			fmt.Printf("📤 Sent job %d: %s\n", job.ID, filepath.Base(job.FilePath))
		}
		close(jobQueue) // Close channel when all jobs sent
	}()
	
	// Wait for all results
	time.Sleep(2 * time.Second)
	
	// Print results
	fp.printResults()
}

// ============================================================================
// 4. WORKER GOROUTINE - Demonstrates Channel Synchronization
// ============================================================================

// worker demonstrates channel-based worker pattern
func (fp *FileProcessor) worker(workerID int) {
	fmt.Printf("👷 Worker %d started\n", workerID)
	
	for job := range jobQueue {
		// Acquire worker slot (buffered channel)
		workerPool <- struct{}{}
		
		fmt.Printf("🔄 Worker %d processing: %s\n", workerID, filepath.Base(job.FilePath))
		
		// Process file
		fileInfo := fp.processFile(job.FilePath)
		
		// Send result
		result := Result{
			Job:      job,
			FileInfo: fileInfo,
		}
		resultQueue <- result
		
		// Release worker slot
		<-workerPool
		
		fmt.Printf("✅ Worker %d completed: %s\n", workerID, filepath.Base(job.FilePath))
	}
	
	fmt.Printf("🛑 Worker %d finished\n", workerID)
}

// ============================================================================
// 5. FILE PROCESSING - Demonstrates Channel Communication
// ============================================================================

// processFile demonstrates channel-based file processing
func (fp *FileProcessor) processFile(filePath string) FileInfo {
	fileInfo := FileInfo{Path: filePath}
	
	// Read file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fileInfo.Error = err
		return fileInfo
	}
	
	// Count words, lines, characters
	fileInfo.WordCount = len(strings.Fields(string(content)))
	fileInfo.LineCount = strings.Count(string(content), "\n") + 1
	fileInfo.CharCount = len(content)
	
	return fileInfo
}

// ============================================================================
// 6. RESULT COLLECTOR - Demonstrates Channel Synchronization
// ============================================================================

// collectResults demonstrates channel-based result collection
func (fp *FileProcessor) collectResults() {
	fmt.Println("📊 Result collector started")
	
	for result := range resultQueue {
		fp.results = append(fp.results, result)
		fmt.Printf("📈 Collected result for: %s\n", filepath.Base(result.Job.FilePath))
	}
	
	fmt.Println("📊 Result collector finished")
}

// ============================================================================
// 7. RESULTS DISPLAY - Demonstrates Channel Data Processing
// ============================================================================

// printResults demonstrates processing channel data
func (fp *FileProcessor) printResults() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 FILE PROCESSING RESULTS")
	fmt.Println(strings.Repeat("=", 60))
	
	totalWords := 0
	totalLines := 0
	totalChars := 0
	successCount := 0
	
	for _, result := range fp.results {
		if result.FileInfo.Error != nil {
			fmt.Printf("❌ %s: Error - %v\n", filepath.Base(result.FileInfo.Path), result.FileInfo.Error)
		} else {
			fmt.Printf("✅ %s: %d words, %d lines, %d chars\n", 
				filepath.Base(result.FileInfo.Path),
				result.FileInfo.WordCount,
				result.FileInfo.LineCount,
				result.FileInfo.CharCount)
			
			totalWords += result.FileInfo.WordCount
			totalLines += result.FileInfo.LineCount
			totalChars += result.FileInfo.CharCount
			successCount++
		}
	}
	
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("📈 SUMMARY: %d files processed successfully\n", successCount)
	fmt.Printf("📝 Total: %d words, %d lines, %d characters\n", totalWords, totalLines, totalChars)
	fmt.Println(strings.Repeat("=", 60))
}

// ============================================================================
// 8. MAIN FUNCTION - Demonstrates Goroutine Coordination
// ============================================================================

func main() {
	fmt.Println("🚀 Go Channel Learning Project")
	fmt.Println("📁 Simple File Processor")
	fmt.Println("🎯 Learning: Channels, Goroutines, Synchronization")
	fmt.Println()
	
	// Create sample files for testing
	createSampleFiles()
	
	// File paths to process
	filePaths := []string{
		"sample1.txt",
		"sample2.txt", 
		"sample3.txt",
		"sample4.txt",
		"sample5.txt",
	}
	
	// Create file processor with 3 workers
	processor := NewFileProcessor(3)
	
	// Process files
	processor.ProcessFiles(filePaths)
	
	// Clean up sample files
	cleanupSampleFiles()
	
	fmt.Println("\n🎉 Project completed! You've learned:")
	fmt.Println("   ✅ Unbuffered channels for synchronization")
	fmt.Println("   ✅ Buffered channels for performance")
	fmt.Println("   ✅ Goroutines for concurrency")
	fmt.Println("   ✅ Channel communication between goroutines")
	fmt.Println("   ✅ Error handling with channels")
}

// ============================================================================
// 9. HELPER FUNCTIONS - Demonstrates Channel Patterns
// ============================================================================

// createSampleFiles creates test files
func createSampleFiles() {
	fmt.Println("📝 Creating sample files...")
	
	samples := []struct {
		filename string
		content  string
	}{
		{"sample1.txt", "Hello world! This is sample file 1.\nIt has multiple lines.\nPerfect for testing."},
		{"sample2.txt", "Go channels are amazing!\nThey make concurrency simple.\nLearn them well."},
		{"sample3.txt", "This is a longer file with more content.\nIt will help us test our file processor.\nChannels make everything better!"},
		{"sample4.txt", "Short file."},
		{"sample5.txt", "Another test file\nwith some content\nfor processing."},
	}
	
	for _, sample := range samples {
		err := ioutil.WriteFile(sample.filename, []byte(sample.content), 0644)
		if err != nil {
			fmt.Printf("❌ Error creating %s: %v\n", sample.filename, err)
		} else {
			fmt.Printf("✅ Created %s\n", sample.filename)
		}
	}
	fmt.Println()
}

// cleanupSampleFiles removes test files
func cleanupSampleFiles() {
	fmt.Println("\n🧹 Cleaning up sample files...")
	
	files := []string{"sample1.txt", "sample2.txt", "sample3.txt", "sample4.txt", "sample5.txt"}
	
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("❌ Error removing %s: %v\n", file, err)
		} else {
			fmt.Printf("✅ Removed %s\n", file)
		}
	}
}