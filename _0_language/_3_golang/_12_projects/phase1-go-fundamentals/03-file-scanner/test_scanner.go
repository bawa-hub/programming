package main

import (
	"file-scanner/internal/scanner"
	"file-scanner/pkg/models"
	"fmt"
	"time"
)

func main() {
	// Create simple scan options
	options := &models.ScanOptions{
		MaxDepth:       2,
		FollowSymlinks: false,
		IncludeHidden:  false,
		IncludeSystem:  false,
		Concurrency:    1, // Use single worker to avoid complexity
		BufferSize:     1024,
		Timeout:        5 * time.Second, // Short timeout
		Progress:       false,
		Verbose:        false,
	}

	// Create scanner
	fileScanner := scanner.NewScanner(options)

	// Test scan
	fmt.Println("Testing file scanner...")
	result, err := fileScanner.Scan("./testdata")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Display results
	fmt.Printf("Scan completed successfully!\n")
	fmt.Printf("Files: %d, Dirs: %d, Size: %d bytes\n", 
		result.TotalFiles, result.TotalDirs, result.TotalSize)
	fmt.Printf("Duration: %v\n", result.ScanDuration)
}
