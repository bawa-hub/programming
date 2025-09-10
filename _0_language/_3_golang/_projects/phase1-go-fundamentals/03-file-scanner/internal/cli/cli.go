package cli

import (
	"file-scanner/internal/analytics"
	"file-scanner/internal/scanner"
	"file-scanner/pkg/models"
	"file-scanner/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// CLI represents the command-line interface for the file scanner
type CLI struct {
	fileScanner *scanner.Scanner
	analyzer    *analytics.Analyzer
	inputScanner *bufio.Scanner
}

// New creates a new CLI instance
func New() *CLI {
	return &CLI{
		inputScanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the CLI application
func (cli *CLI) Run() {
	cli.printWelcome()
	cli.printHelp()
	
	for {
		fmt.Print(utils.Bold(utils.Cyan("scanner> ")))
		if !cli.inputScanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(cli.inputScanner.Text())
		if input == "" {
			continue
		}
		
		// Parse and execute command
		if err := cli.executeCommand(input); err != nil {
			fmt.Printf("%s %s\n", utils.Red("Error:"), err.Error())
		}
	}
}

// executeCommand parses and executes a command
func (cli *CLI) executeCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}
	
	command := strings.ToLower(parts[0])
	args := parts[1:]
	
	switch command {
	case "help", "h":
		cli.printHelp()
	case "scan", "s":
		return cli.handleScan(args)
	case "analyze", "a":
		return cli.handleAnalyze(args)
	case "filter", "f":
		return cli.handleFilter(args)
	case "search":
		return cli.handleSearch(args)
	case "stats", "st":
		cli.handleStats()
	case "export", "e":
		return cli.handleExport(args)
	case "config", "c":
		return cli.handleConfig(args)
	case "quit", "q", "exit":
		fmt.Println(utils.Green("Goodbye!"))
		os.Exit(0)
	default:
		return fmt.Errorf("unknown command: %s. Type 'help' for available commands", command)
	}
	
	return nil
}

// Command handlers

func (cli *CLI) handleScan(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: scan <path> [options]")
	}
	
	path := args[0]
	
	// Parse options
	options := &models.ScanOptions{
		MaxDepth:       10,
		FollowSymlinks: false,
		IncludeHidden:  false,
		IncludeSystem:  false,
		Concurrency:    4,
		BufferSize:     1024 * 1024,
		Timeout:        30 * time.Minute,
		Progress:       true,
		Verbose:        false,
	}
	
	// Parse additional options
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-d", "--depth":
			if i+1 < len(args) {
				depth, err := strconv.Atoi(args[i+1])
				if err != nil {
					return fmt.Errorf("invalid depth: %s", args[i+1])
				}
				options.MaxDepth = depth
				i++
			}
		case "-s", "--symlinks":
			options.FollowSymlinks = true
		case "-h", "--hidden":
			options.IncludeHidden = true
		case "-sys", "--system":
			options.IncludeSystem = true
		case "-c", "--concurrency":
			if i+1 < len(args) {
				conc, err := strconv.Atoi(args[i+1])
				if err != nil {
					return fmt.Errorf("invalid concurrency: %s", args[i+1])
				}
				options.Concurrency = conc
				i++
			}
		case "-v", "--verbose":
			options.Verbose = true
		case "--no-progress":
			options.Progress = false
		case "--duplicates":
			options.FindDuplicates = true
		case "--hashes":
			options.CalculateHashes = true
		}
	}
	
	// Create scanner
	cli.fileScanner = scanner.NewScanner(options)
	
	// Start scanning
	fmt.Printf("%s Scanning %s...\n", utils.Blue("🔍"), path)
	start := time.Now()
	
	result, err := cli.fileScanner.Scan(path)
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}
	
	duration := time.Since(start)
	
	// Display results
	cli.displayScanResults(result, duration)
	
	// Create analyzer
	cli.analyzer = analytics.NewAnalyzer(result)
	
	return nil
}

func (cli *CLI) handleAnalyze(args []string) error {
	if cli.analyzer == nil {
		return fmt.Errorf("no scan results available. Run 'scan' first")
	}
	
	report := cli.analyzer.GenerateReport()
	cli.displayAnalysisReport(report)
	
	return nil
}

func (cli *CLI) handleFilter(args []string) error {
	if cli.analyzer == nil {
		return fmt.Errorf("no scan results available. Run 'scan' first")
	}
	
	if len(args) < 2 {
		return fmt.Errorf("usage: filter <field> <value>")
	}
	
	field := strings.ToLower(args[0])
	value := args[1]
	
	// This would implement filtering logic
	fmt.Printf("%s Filtering by %s = %s\n", utils.Yellow("🔍"), field, value)
	
	return nil
}

func (cli *CLI) handleSearch(args []string) error {
	if cli.analyzer == nil {
		return fmt.Errorf("no scan results available. Run 'scan' first")
	}
	
	if len(args) == 0 {
		return fmt.Errorf("usage: search <query>")
	}
	
	query := strings.Join(args, " ")
	fmt.Printf("%s Searching for: %s\n", utils.Yellow("🔍"), query)
	
	// This would implement search logic
	return nil
}

func (cli *CLI) handleStats() {
	if cli.analyzer == nil {
		fmt.Println(utils.Yellow("No scan results available. Run 'scan' first"))
		return
	}
	
	report := cli.analyzer.GenerateReport()
	cli.displayStats(report)
}

func (cli *CLI) handleExport(args []string) error {
	if cli.analyzer == nil {
		return fmt.Errorf("no scan results available. Run 'scan' first")
	}
	
	if len(args) < 2 {
		return fmt.Errorf("usage: export <format> <filepath>")
	}
	
	format := strings.ToLower(args[0])
	filepath := args[1]
	
	switch format {
	case "json":
		return cli.exportJSON(filepath)
	case "csv":
		return cli.exportCSV(filepath)
	case "txt":
		return cli.exportTXT(filepath)
	default:
		return fmt.Errorf("unsupported format: %s. Use: json, csv, txt", format)
	}
}

func (cli *CLI) handleConfig(args []string) error {
	if len(args) == 0 {
		cli.displayConfig()
		return nil
	}
	
	// This would implement configuration management
	fmt.Printf("%s Configuration updated\n", utils.Green("✓"))
	return nil
}

// Display functions

func (cli *CLI) printWelcome() {
	fmt.Println(utils.Bold(utils.Cyan("🔍 File System Scanner")))
	fmt.Println(utils.Yellow("Advanced file system analysis and scanning tool"))
	fmt.Println(utils.Yellow("Type 'help' for available commands"))
	fmt.Println()
}

func (cli *CLI) printHelp() {
	fmt.Println(utils.Bold("Available Commands:"))
	fmt.Println()
	
	commands := []struct {
		command string
		usage   string
		desc    string
	}{
		{"scan, s", "scan <path> [options]", "Scan a directory for files and directories"},
		{"analyze, a", "analyze", "Analyze scan results and generate report"},
		{"filter, f", "filter <field> <value>", "Filter scan results by field"},
		{"search", "search <query>", "Search for files by name or content"},
		{"stats, st", "stats", "Show scanning statistics"},
		{"export, e", "export <format> <filepath>", "Export results (json, csv, txt)"},
		{"config, c", "config [options]", "Show or modify configuration"},
		{"help, h", "help", "Show this help message"},
		{"quit, q, exit", "quit", "Exit the application"},
	}
	
	for _, cmd := range commands {
		fmt.Printf("  %-20s %-30s %s\n", 
			utils.Bold(cmd.command), 
			utils.Yellow(cmd.usage), 
			cmd.desc)
	}
	fmt.Println()
	
	fmt.Println(utils.Bold("Scan Options:"))
	fmt.Println("  -d, --depth <n>        Maximum directory depth")
	fmt.Println("  -s, --symlinks         Follow symbolic links")
	fmt.Println("  -h, --hidden           Include hidden files")
	fmt.Println("  -sys, --system         Include system files")
	fmt.Println("  -c, --concurrency <n>  Number of concurrent workers")
	fmt.Println("  -v, --verbose          Verbose output")
	fmt.Println("  --no-progress          Disable progress reporting")
	fmt.Println("  --duplicates           Find duplicate files")
	fmt.Println("  --hashes               Calculate file hashes")
	fmt.Println()
}

func (cli *CLI) displayScanResults(result *models.ScanResult, duration time.Duration) {
	fmt.Println()
	fmt.Println(utils.Bold(utils.Green("📊 Scan Results:")))
	fmt.Printf("  Root Path: %s\n", result.RootPath)
	fmt.Printf("  Total Files: %d\n", result.TotalFiles)
	fmt.Printf("  Total Directories: %d\n", result.TotalDirs)
	fmt.Printf("  Total Size: %s\n", formatBytes(result.TotalSize))
	fmt.Printf("  Scan Duration: %s\n", duration.Round(time.Millisecond))
	fmt.Printf("  Scan Rate: %.2f files/sec\n", result.GetScanRate())
	fmt.Printf("  Errors: %d\n", len(result.Errors))
	
	if len(result.Errors) > 0 {
		fmt.Println()
		fmt.Println(utils.Bold(utils.Red("❌ Errors:")))
		for i, err := range result.Errors {
			if i >= 5 { // Show only first 5 errors
				fmt.Printf("  ... and %d more errors\n", len(result.Errors)-5)
				break
			}
			fmt.Printf("  %s: %s\n", err.Path, err.Error)
		}
	}
	
	fmt.Println()
}

func (cli *CLI) displayAnalysisReport(report *analytics.Report) {
	fmt.Println()
	fmt.Println(utils.Bold(utils.Green("📈 Analysis Report:")))
	
	// Summary
	cli.displaySummary(report.Summary)
	
	// File Types
	cli.displayFileTypes(report.FileTypes)
	
	// Extensions
	cli.displayExtensions(report.Extensions)
	
	// Size Analysis
	cli.displaySizeAnalysis(report.SizeAnalysis)
	
	// Duplicates
	if len(report.Duplicates.Groups) > 0 {
		cli.displayDuplicates(report.Duplicates)
	}
	
	// Recommendations
	if len(report.Recommendations) > 0 {
		cli.displayRecommendations(report.Recommendations)
	}
	
	fmt.Println()
}

func (cli *CLI) displaySummary(summary *analytics.SummaryReport) {
	fmt.Println(utils.Bold("📋 Summary:"))
	fmt.Printf("  Total Files: %d\n", summary.TotalFiles)
	fmt.Printf("  Total Directories: %d\n", summary.TotalDirs)
	fmt.Printf("  Total Size: %s\n", formatBytes(summary.TotalSize))
	fmt.Printf("  Average File Size: %s\n", formatBytes(summary.AverageFileSize))
	fmt.Printf("  Scan Duration: %s\n", summary.ScanDuration.Round(time.Millisecond))
	fmt.Printf("  Scan Rate: %.2f files/sec\n", summary.ScanRate)
	fmt.Printf("  Errors: %d\n", summary.ErrorCount)
	fmt.Println()
}

func (cli *CLI) displayFileTypes(fileTypes *analytics.FileTypeReport) {
	fmt.Println(utils.Bold("📁 File Types:"))
	for i, stat := range fileTypes.Top {
		if i >= 10 { // Show only top 10
			break
		}
		fmt.Printf("  %s %s: %d files (%s)\n", 
			stat.Icon, stat.Type, stat.Count, formatBytes(stat.Size))
	}
	fmt.Println()
}

func (cli *CLI) displayExtensions(extensions *analytics.ExtensionReport) {
	fmt.Println(utils.Bold("📄 File Extensions:"))
	for i, stat := range extensions.Top {
		if i >= 10 { // Show only top 10
			break
		}
		fmt.Printf("  .%s: %d files (%s) - %s\n", 
			stat.Extension, stat.Count, formatBytes(stat.Size), stat.MimeType)
	}
	fmt.Println()
}

func (cli *CLI) displaySizeAnalysis(sizeAnalysis *analytics.SizeReport) {
	fmt.Println(utils.Bold("📏 Size Analysis:"))
	fmt.Printf("  Total Size: %s\n", formatBytes(sizeAnalysis.TotalSize))
	fmt.Printf("  Average Size: %s\n", formatBytes(sizeAnalysis.AverageSize))
	fmt.Printf("  Median Size: %s\n", formatBytes(sizeAnalysis.MedianSize))
	fmt.Printf("  Largest File: %s\n", formatBytes(sizeAnalysis.LargestFile))
	fmt.Printf("  Smallest File: %s\n", formatBytes(sizeAnalysis.SmallestFile))
	
	fmt.Println("  Size Distribution:")
	for _, category := range sizeAnalysis.SizeCategories {
		fmt.Printf("    %s: %d files (%s)\n", 
			category.Range, category.Count, formatBytes(category.Size))
	}
	fmt.Println()
}

func (cli *CLI) displayDuplicates(duplicates *analytics.DuplicateReport) {
	fmt.Println(utils.Bold("🔄 Duplicate Files:"))
	fmt.Printf("  Total Duplicates: %d\n", duplicates.TotalDuplicates)
	fmt.Printf("  Wasted Space: %s\n", formatBytes(duplicates.WastedSpace))
	
	if len(duplicates.Groups) > 0 {
		fmt.Println("  Duplicate Groups:")
		for i, group := range duplicates.Groups {
			if i >= 5 { // Show only first 5 groups
				fmt.Printf("    ... and %d more groups\n", len(duplicates.Groups)-5)
				break
			}
			fmt.Printf("    %s (%d files):\n", formatBytes(group.Size), group.Count)
			for j, file := range group.Files {
				if j >= 3 { // Show only first 3 files per group
					fmt.Printf("      ... and %d more files\n", group.Count-3)
					break
				}
				fmt.Printf("      %s\n", file.Path)
			}
		}
	}
	fmt.Println()
}

func (cli *CLI) displayRecommendations(recommendations []string) {
	fmt.Println(utils.Bold("💡 Recommendations:"))
	for i, rec := range recommendations {
		fmt.Printf("  %d. %s\n", i+1, rec)
	}
	fmt.Println()
}

func (cli *CLI) displayStats(report *analytics.Report) {
	cli.displaySummary(report.Summary)
}

func (cli *CLI) displayConfig() {
	fmt.Println(utils.Bold("⚙️ Configuration:"))
	fmt.Println("  Default Options:")
	fmt.Println("    Max Depth: 10")
	fmt.Println("    Follow Symlinks: false")
	fmt.Println("    Include Hidden: false")
	fmt.Println("    Include System: false")
	fmt.Println("    Concurrency: 4")
	fmt.Println("    Buffer Size: 1MB")
	fmt.Println("    Timeout: 30 minutes")
	fmt.Println("    Progress: true")
	fmt.Println("    Verbose: false")
	fmt.Println()
}

// Export functions

func (cli *CLI) exportJSON(filepath string) error {
	fmt.Printf("%s Exporting to JSON: %s\n", utils.Green("✓"), filepath)
	// Implementation would go here
	return nil
}

func (cli *CLI) exportCSV(filepath string) error {
	fmt.Printf("%s Exporting to CSV: %s\n", utils.Green("✓"), filepath)
	// Implementation would go here
	return nil
}

func (cli *CLI) exportTXT(filepath string) error {
	fmt.Printf("%s Exporting to TXT: %s\n", utils.Green("✓"), filepath)
	// Implementation would go here
	return nil
}

// Utility functions

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
