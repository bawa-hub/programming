package main

import (
	"file-scanner/internal/cli"
	"file-scanner/internal/scanner"
	"file-scanner/pkg/models"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// Define command-line flags
	var (
		path        = flag.String("path", "", "Path to scan")
		depth       = flag.Int("depth", 10, "Maximum directory depth")
		symlinks    = flag.Bool("symlinks", false, "Follow symbolic links")
		hidden      = flag.Bool("hidden", false, "Include hidden files")
		system      = flag.Bool("system", false, "Include system files")
		concurrency = flag.Int("concurrency", runtime.NumCPU(), "Number of concurrent workers")
		verbose     = flag.Bool("verbose", false, "Verbose output")
		progress    = flag.Bool("progress", true, "Show progress")
		duplicates  = flag.Bool("duplicates", false, "Find duplicate files")
		hashes      = flag.Bool("hashes", false, "Calculate file hashes")
		export      = flag.String("export", "", "Export format (json, csv, txt)")
		output      = flag.String("output", "", "Output file path")
		interactive = flag.Bool("interactive", false, "Start in interactive mode")
		help        = flag.Bool("help", false, "Show help message")
	)
	
	flag.Parse()

	// Show help if requested
	if *help {
		showHelp()
		return
	}

	// Interactive mode
	if *interactive || *path == "" {
		startInteractiveMode()
		return
	}

	// Validate path
	if *path == "" {
		fmt.Fprintf(os.Stderr, "Error: path is required\n")
		os.Exit(1)
	}

	// Create scan options
	options := &models.ScanOptions{
		MaxDepth:        *depth,
		FollowSymlinks:  *symlinks,
		IncludeHidden:   *hidden,
		IncludeSystem:   *system,
		Concurrency:     *concurrency,
		BufferSize:      1024 * 1024, // 1MB
		Timeout:         30 * time.Minute,
		Progress:        *progress,
		Verbose:         *verbose,
		FindDuplicates:  *duplicates,
		CalculateHashes: *hashes,
	}

	// Create scanner
	fileScanner := scanner.NewScanner(options)

	// Start scanning
	fmt.Printf("🔍 Scanning %s...\n", *path)
	start := time.Now()

	result, err := fileScanner.Scan(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	duration := time.Since(start)

	// Display results
	displayResults(result, duration)

	// Export if requested
	if *export != "" {
		if err := exportResults(result, *export, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Export error: %v\n", err)
			os.Exit(1)
		}
	}
}

func startInteractiveMode() {
	fmt.Println("🚀 Starting File System Scanner in interactive mode...")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	cliApp := cli.New()
	cliApp.Run()
}

func displayResults(result *models.ScanResult, duration time.Duration) {
	fmt.Println()
	fmt.Println("📊 Scan Results:")
	fmt.Printf("  Root Path: %s\n", result.RootPath)
	fmt.Printf("  Total Files: %d\n", result.TotalFiles)
	fmt.Printf("  Total Directories: %d\n", result.TotalDirs)
	fmt.Printf("  Total Size: %s\n", formatBytes(result.TotalSize))
	fmt.Printf("  Scan Duration: %s\n", duration.Round(time.Millisecond))
	fmt.Printf("  Scan Rate: %.2f files/sec\n", result.GetScanRate())
	fmt.Printf("  Errors: %d\n", len(result.Errors))

	if len(result.Errors) > 0 {
		fmt.Println()
		fmt.Println("❌ Errors:")
		for i, err := range result.Errors {
			if i >= 10 { // Show only first 10 errors
				fmt.Printf("  ... and %d more errors\n", len(result.Errors)-10)
				break
			}
			fmt.Printf("  %s: %s\n", err.Path, err.Error)
		}
	}

	// Show file type distribution
	if result.Statistics != nil {
		fmt.Println()
		fmt.Println("📁 File Types:")
		for fileType, count := range result.Statistics.FileCountByType {
			fmt.Printf("  %s: %d files\n", fileType.String(), count)
		}

		// Show largest files
		if len(result.Statistics.LargestFiles) > 0 {
			fmt.Println()
			fmt.Println("📏 Largest Files:")
			for i, file := range result.Statistics.LargestFiles {
				if i >= 5 { // Show only top 5
					break
				}
				fmt.Printf("  %s (%s)\n", file.Path, formatBytes(file.Size))
			}
		}

		// Show duplicates if found
		if len(result.Statistics.DuplicateFiles) > 0 {
			fmt.Println()
			fmt.Println("🔄 Duplicate Files:")
			totalWasted := int64(0)
			for _, group := range result.Statistics.DuplicateFiles {
				totalWasted += int64(group.Count-1) * group.Size
			}
			fmt.Printf("  Total Duplicates: %d groups\n", len(result.Statistics.DuplicateFiles))
			fmt.Printf("  Wasted Space: %s\n", formatBytes(totalWasted))
		}
	}

	fmt.Println()
}

func exportResults(result *models.ScanResult, format, output string) error {
	if output == "" {
		output = fmt.Sprintf("scan_results.%s", format)
	}

	switch format {
	case "json":
		return exportJSON(result, output)
	case "csv":
		return exportCSV(result, output)
	case "txt":
		return exportTXT(result, output)
	default:
		return fmt.Errorf("unsupported export format: %s", format)
	}
}

func exportJSON(result *models.ScanResult, output string) error {
	fmt.Printf("📄 Exporting to JSON: %s\n", output)
	// Implementation would go here
	return nil
}

func exportCSV(result *models.ScanResult, output string) error {
	fmt.Printf("📊 Exporting to CSV: %s\n", output)
	// Implementation would go here
	return nil
}

func exportTXT(result *models.ScanResult, output string) error {
	fmt.Printf("📝 Exporting to TXT: %s\n", output)
	// Implementation would go here
	return nil
}

func showHelp() {
	fmt.Println("File System Scanner - Advanced file system analysis tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  file-scanner [flags]")
	fmt.Println("  file-scanner -interactive")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -path string")
	fmt.Println("        Path to scan (required for non-interactive mode)")
	fmt.Println("  -depth int")
	fmt.Println("        Maximum directory depth (default 10)")
	fmt.Println("  -symlinks")
	fmt.Println("        Follow symbolic links")
	fmt.Println("  -hidden")
	fmt.Println("        Include hidden files")
	fmt.Println("  -system")
	fmt.Println("        Include system files")
	fmt.Println("  -concurrency int")
	fmt.Println("        Number of concurrent workers (default: number of CPUs)")
	fmt.Println("  -verbose")
	fmt.Println("        Verbose output")
	fmt.Println("  -progress")
	fmt.Println("        Show progress (default true)")
	fmt.Println("  -duplicates")
	fmt.Println("        Find duplicate files")
	fmt.Println("  -hashes")
	fmt.Println("        Calculate file hashes")
	fmt.Println("  -export string")
	fmt.Println("        Export format (json, csv, txt)")
	fmt.Println("  -output string")
	fmt.Println("        Output file path")
	fmt.Println("  -interactive")
	fmt.Println("        Start in interactive mode")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  file-scanner -path /home/user")
	fmt.Println("  file-scanner -path /var/log -depth 5 -duplicates")
	fmt.Println("  file-scanner -path /tmp -export json -output scan.json")
	fmt.Println("  file-scanner -interactive")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  • Concurrent file system scanning")
	fmt.Println("  • Advanced file analysis and statistics")
	fmt.Println("  • Duplicate file detection")
	fmt.Println("  • File hash calculation")
	fmt.Println("  • Multiple export formats")
	fmt.Println("  • Interactive command-line interface")
	fmt.Println("  • Progress reporting and error handling")
	fmt.Println()
	fmt.Println("This tool demonstrates advanced Go concepts:")
	fmt.Println("  • Concurrency patterns (worker pools, goroutines)")
	fmt.Println("  • File system operations and I/O")
	fmt.Println("  • Data structures and algorithms")
	fmt.Println("  • Error handling and recovery")
	fmt.Println("  • Command-line interface development")
	fmt.Println("  • Performance optimization")
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
