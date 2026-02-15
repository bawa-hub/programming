package main

import (
	"flag"
	"fmt"
	"log"
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
		listCmd     = flag.Bool("list", false, "List all processes")
		monitorCmd  = flag.Bool("monitor", false, "Monitor processes")
		startCmd    = flag.Bool("start", false, "Start a new process")
		stopCmd     = flag.Bool("stop", false, "Stop a process")
		killCmd     = flag.Bool("kill", false, "Kill a process")
		treeCmd     = flag.Bool("tree", false, "Show process tree")
		statsCmd    = flag.Bool("stats", false, "Show system statistics")
		versionCmd  = flag.Bool("version", false, "Show version information")
		
		// Options
		pid         = flag.Int("pid", 0, "Process ID")
		name        = flag.String("name", "", "Process name filter")
		user        = flag.String("user", "", "User filter")
		detailed    = flag.Bool("detailed", false, "Show detailed information")
		watch       = flag.Bool("watch", false, "Watch mode (real-time updates)")
		interval    = flag.Duration("interval", 1*time.Second, "Update interval")
		command     = flag.String("cmd", "", "Command to execute")
		args        = flag.String("args", "", "Command arguments")
		output      = flag.String("output", "table", "Output format (table, json, csv)")
		limit       = flag.Int("limit", 100, "Maximum number of processes to show")
	)
	
	flag.Parse()
	
	// Show version
	if *versionCmd {
		fmt.Printf("Process Manager v%s (build %s)\n", version, build)
		return
	}
	
	// Create process manager
	pm := NewProcessManager()
	defer pm.Close()
	
	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nShutting down...")
		pm.Close()
		os.Exit(0)
	}()
	
	// Execute commands
	switch {
	case *listCmd:
		handleList(pm, &ListOptions{
			Name:     *name,
			User:     *user,
			Detailed: *detailed,
			Output:   *output,
			Limit:    *limit,
		})
		
	case *monitorCmd:
		handleMonitor(pm, &MonitorOptions{
			PID:      *pid,
			Name:     *name,
			Watch:    *watch,
			Interval: *interval,
			Output:   *output,
		})
		
	case *startCmd:
		handleStart(pm, &StartOptions{
			Command: *command,
			Args:    *args,
		})
		
	case *stopCmd:
		handleStop(pm, *pid)
		
	case *killCmd:
		handleKill(pm, *pid)
		
	case *treeCmd:
		handleTree(pm, &TreeOptions{
			PID:    *pid,
			Output: *output,
		})
		
	case *statsCmd:
		handleStats(pm, &StatsOptions{
			Output: *output,
		})
		
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Process Manager - Advanced System Programming")
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println("Usage: process-manager [command] [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list      List all processes")
	fmt.Println("  monitor   Monitor processes")
	fmt.Println("  start     Start a new process")
	fmt.Println("  stop      Stop a process")
	fmt.Println("  kill      Kill a process")
	fmt.Println("  tree      Show process tree")
	fmt.Println("  stats     Show system statistics")
	fmt.Println("  version   Show version information")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -pid int           Process ID")
	fmt.Println("  -name string       Process name filter")
	fmt.Println("  -user string       User filter")
	fmt.Println("  -detailed          Show detailed information")
	fmt.Println("  -watch             Watch mode (real-time updates)")
	fmt.Println("  -interval duration Update interval (default 1s)")
	fmt.Println("  -cmd string        Command to execute")
	fmt.Println("  -args string       Command arguments")
	fmt.Println("  -output string     Output format (table, json, csv)")
	fmt.Println("  -limit int         Maximum number of processes to show")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  process-manager -list")
	fmt.Println("  process-manager -list -name=chrome -detailed")
	fmt.Println("  process-manager -monitor -pid=1234 -watch")
	fmt.Println("  process-manager -start -cmd=ls -args='-la'")
	fmt.Println("  process-manager -tree -pid=1")
	fmt.Println("  process-manager -stats -output=json")
}

func handleList(pm *ProcessManager, opts *ListOptions) {
	processes, err := pm.ListProcesses(opts)
	if err != nil {
		log.Fatalf("Failed to list processes: %v", err)
	}
	
	switch opts.Output {
	case "json":
		pm.PrintJSON(processes)
	case "csv":
		pm.PrintCSV(processes)
	default:
		pm.PrintTable(processes)
	}
}

func handleMonitor(pm *ProcessManager, opts *MonitorOptions) {
	if opts.Watch {
		pm.MonitorRealtime(opts)
	} else {
		processes, err := pm.GetProcesses(opts)
		if err != nil {
			log.Fatalf("Failed to get processes: %v", err)
		}
		
		switch opts.Output {
		case "json":
			pm.PrintJSON(processes)
		case "csv":
			pm.PrintCSV(processes)
		default:
			pm.PrintTable(processes)
		}
	}
}

func handleStart(pm *ProcessManager, opts *StartOptions) {
	if opts.Command == "" {
		log.Fatal("Command is required")
	}
	
	pid, err := pm.StartProcess(opts)
	if err != nil {
		log.Fatalf("Failed to start process: %v", err)
	}
	
	fmt.Printf("Process started with PID: %d\n", pid)
}

func handleStop(pm *ProcessManager, pid int) {
	if pid == 0 {
		log.Fatal("PID is required")
	}
	
	err := pm.StopProcess(pid)
	if err != nil {
		log.Fatalf("Failed to stop process: %v", err)
	}
	
	fmt.Printf("Process %d stopped\n", pid)
}

func handleKill(pm *ProcessManager, pid int) {
	if pid == 0 {
		log.Fatal("PID is required")
	}
	
	err := pm.KillProcess(pid)
	if err != nil {
		log.Fatalf("Failed to kill process: %v", err)
	}
	
	fmt.Printf("Process %d killed\n", pid)
}

func handleTree(pm *ProcessManager, opts *TreeOptions) {
	tree, err := pm.GetProcessTree(opts)
	if err != nil {
		log.Fatalf("Failed to get process tree: %v", err)
	}
	
	switch opts.Output {
	case "json":
		pm.PrintJSON(tree)
	case "csv":
		fmt.Println("CSV output not supported for process tree")
	default:
		pm.PrintTree(tree)
	}
}

func handleStats(pm *ProcessManager, opts *StatsOptions) {
	stats, err := pm.GetSystemStats()
	if err != nil {
		log.Fatalf("Failed to get system stats: %v", err)
	}
	
	switch opts.Output {
	case "json":
		pm.PrintJSON(stats)
	case "csv":
		fmt.Println("CSV output not supported for system stats")
	default:
		pm.PrintStats(stats)
	}
}
