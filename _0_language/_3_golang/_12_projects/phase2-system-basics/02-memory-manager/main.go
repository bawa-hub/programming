package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version = "1.0.0"
	build   = "dev"
)

func main() {
	// Command line flags
	var (
		monitorCmd    = flag.Bool("monitor", false, "Monitor memory usage")
		profileCmd    = flag.Bool("profile", false, "Generate memory profile")
		analyzeCmd    = flag.Bool("analyze", false, "Analyze memory profile")
		optimizeCmd   = flag.Bool("optimize", false, "Run memory optimization")
		poolCmd       = flag.Bool("pool", false, "Test memory pools")
		leakCmd       = flag.Bool("leak-detect", false, "Detect memory leaks")
		statsCmd      = flag.Bool("stats", false, "Show memory statistics")
		versionCmd    = flag.Bool("version", false, "Show version information")
		
		// Options
		watch         = flag.Bool("watch", false, "Watch mode (real-time updates)")
		interval      = flag.Duration("interval", 1*time.Second, "Update interval")
		output        = flag.String("output", "profile.prof", "Output file for profile")
		profile       = flag.String("profile", "", "Profile file to analyze")
		// profile1      = flag.String("profile1", "", "First profile for comparison")
		// profile2      = flag.String("profile2", "", "Second profile for comparison")
		format        = flag.String("format", "table", "Output format (table, json, csv)")
		pid           = flag.Int("pid", 0, "Process ID to monitor")
		threshold     = flag.Float64("threshold", 80.0, "Memory usage threshold percentage")
		duration      = flag.Duration("duration", 30*time.Second, "Monitoring duration")
	)
	
	flag.Parse()
	
	// Show version
	if *versionCmd {
		fmt.Printf("Memory Manager v%s (build %s)\n", version, build)
		return
	}
	
	// Create memory manager
	mm := NewMemoryManager()
	defer mm.Close()
	
	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nShutting down...")
		mm.Close()
		os.Exit(0)
	}()
	
	// Execute commands
	switch {
	case *monitorCmd:
		handleMonitor(mm, &MonitorOptions{
			Watch:     *watch,
			Interval:  *interval,
			Format:    *format,
			PID:       *pid,
			Threshold: *threshold,
			Duration:  *duration,
		})
		
	case *profileCmd:
		handleProfile(mm, &ProfileOptions{
			Output: *output,
			Format: *format,
		})
		
	case *analyzeCmd:
		handleAnalyze(mm, &AnalyzeOptions{
			Profile: *profile,
			Format:  *format,
		})
		
	case *optimizeCmd:
		handleOptimize(mm, &OptimizeOptions{
			Format: *format,
		})
		
	case *poolCmd:
		handlePool(mm, &PoolOptions{
			Format: *format,
		})
		
	case *leakCmd:
		handleLeakDetect(mm, &LeakDetectOptions{
			Format:    *format,
			Threshold: *threshold,
			Duration:  *duration,
		})
		
	case *statsCmd:
		handleStats(mm, &StatsOptions{
			Format: *format,
		})
		
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Memory Manager - Advanced Memory Management")
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("Usage: memory-manager [command] [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  monitor      Monitor memory usage")
	fmt.Println("  profile      Generate memory profile")
	fmt.Println("  analyze      Analyze memory profile")
	fmt.Println("  optimize     Run memory optimization")
	fmt.Println("  pool         Test memory pools")
	fmt.Println("  leak-detect  Detect memory leaks")
	fmt.Println("  stats        Show memory statistics")
	fmt.Println("  version      Show version information")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -watch              Watch mode (real-time updates)")
	fmt.Println("  -interval duration  Update interval (default 1s)")
	fmt.Println("  -output string      Output file for profile")
	fmt.Println("  -profile string     Profile file to analyze")
	fmt.Println("  -profile1 string    First profile for comparison")
	fmt.Println("  -profile2 string    Second profile for comparison")
	fmt.Println("  -format string      Output format (table, json, csv)")
	fmt.Println("  -pid int            Process ID to monitor")
	fmt.Println("  -threshold float    Memory usage threshold percentage")
	fmt.Println("  -duration duration  Monitoring duration")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  memory-manager -monitor")
	fmt.Println("  memory-manager -monitor -watch -interval=500ms")
	fmt.Println("  memory-manager -profile -output=mem.prof")
	fmt.Println("  memory-manager -analyze -profile=mem.prof")
	fmt.Println("  memory-manager -optimize")
	fmt.Println("  memory-manager -pool -test")
	fmt.Println("  memory-manager -leak-detect -threshold=90")
}

func handleMonitor(mm *MemoryManager, opts *MonitorOptions) {
	if opts.Watch {
		mm.MonitorRealtime(opts)
	} else {
		stats, err := mm.GetMemoryStats()
		if err != nil {
			fmt.Printf("Error getting memory stats: %v\n", err)
			return
		}
		
		switch opts.Format {
		case "json":
			mm.PrintJSON(stats)
		case "csv":
			mm.PrintCSV(stats)
		default:
			mm.PrintTable(stats)
		}
	}
}

func handleProfile(mm *MemoryManager, opts *ProfileOptions) {
	err := mm.GenerateProfile(opts)
	if err != nil {
		fmt.Printf("Error generating profile: %v\n", err)
		return
	}
	fmt.Printf("Memory profile generated: %s\n", opts.Output)
}

func handleAnalyze(mm *MemoryManager, opts *AnalyzeOptions) {
	if opts.Profile == "" {
		fmt.Println("Profile file is required")
		return
	}
	
	analysis, err := mm.AnalyzeProfile(opts)
	if err != nil {
		fmt.Printf("Error analyzing profile: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		mm.PrintJSON(analysis)
	case "csv":
		mm.PrintCSV(analysis)
	default:
		mm.PrintAnalysis(analysis)
	}
}

func handleOptimize(mm *MemoryManager, opts *OptimizeOptions) {
	results, err := mm.OptimizeMemory(opts)
	if err != nil {
		fmt.Printf("Error optimizing memory: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		mm.PrintJSON(results)
	case "csv":
		mm.PrintCSV(results)
	default:
		mm.PrintOptimizationResults(results)
	}
}

func handlePool(mm *MemoryManager, opts *PoolOptions) {
	results, err := mm.TestMemoryPools(opts)
	if err != nil {
		fmt.Printf("Error testing memory pools: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		mm.PrintJSON(results)
	case "csv":
		mm.PrintCSV(results)
	default:
		mm.PrintPoolResults(results)
	}
}

func handleLeakDetect(mm *MemoryManager, opts *LeakDetectOptions) {
	leaks, err := mm.DetectLeaks(opts)
	if err != nil {
		fmt.Printf("Error detecting leaks: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		mm.PrintJSON(leaks)
	case "csv":
		mm.PrintCSV(leaks)
	default:
		mm.PrintLeakReport(leaks)
	}
}

func handleStats(mm *MemoryManager, opts *StatsOptions) {
	stats, err := mm.GetDetailedStats()
	if err != nil {
		fmt.Printf("Error getting detailed stats: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		mm.PrintJSON(stats)
	case "csv":
		mm.PrintCSV(stats)
	default:
		mm.PrintDetailedStats(stats)
	}
}
