package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// MonitorRealtime monitors processes in real-time
func (pm *ProcessManager) MonitorRealtime(opts *MonitorOptions) {
	fmt.Println("Starting real-time monitoring...")
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()
	
	// Create ticker for updates
	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()
	
	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Clear screen
	clearScreen()
	
	for {
		select {
		case <-ticker.C:
			pm.updateMonitorDisplay(opts)
		case <-sigChan:
			fmt.Println("\nMonitoring stopped.")
			return
		case <-pm.ctx.Done():
			return
		}
	}
}

// updateMonitorDisplay updates the monitoring display
func (pm *ProcessManager) updateMonitorDisplay(opts *MonitorOptions) {
	// Move cursor to top
	fmt.Print("\033[H")
	
	// Get processes
	processes, err := pm.GetProcesses(opts)
	if err != nil {
		fmt.Printf("Error getting processes: %v\n", err)
		return
	}
	
	// Print header
	fmt.Println("Process Monitor - Real-time")
	fmt.Println("==========================")
	fmt.Printf("Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Processes: %d\n", len(processes))
	fmt.Println()
	
	// Print process table
	pm.printMonitorTable(processes)
	
	// Print system stats
	stats, err := pm.GetSystemStats()
	if err == nil {
		pm.printSystemStats(stats)
	}
}

// printMonitorTable prints the monitoring table
func (pm *ProcessManager) printMonitorTable(processes []*Process) {
	// Table header
	fmt.Printf("%-8s %-20s %-12s %-8s %-8s %-8s %-12s\n",
		"PID", "NAME", "USER", "CPU%", "MEM(MB)", "THREADS", "STATUS")
	fmt.Println(strings.Repeat("-", 80))
	
	// Table rows
	for _, proc := range processes {
		fmt.Printf("%-8d %-20s %-12s %-8.2f %-8.2f %-8d %-12s\n",
			proc.PID,
			truncateString(proc.Name, 20),
			truncateString(proc.User, 12),
			proc.CPUPercent,
			proc.MemoryMB,
			proc.NumThreads,
			truncateString(proc.Status, 12))
	}
	fmt.Println()
}

// printSystemStats prints system statistics
func (pm *ProcessManager) printSystemStats(stats *SystemStats) {
	fmt.Println("System Statistics")
	fmt.Println("=================")
	fmt.Printf("Hostname: %s\n", stats.Hostname)
	fmt.Printf("OS: %s %s\n", stats.OS, stats.Architecture)
	fmt.Printf("Kernel: %s\n", stats.Kernel)
	fmt.Printf("Uptime: %d seconds\n", stats.Uptime)
	fmt.Printf("Total Processes: %d\n", stats.Processes)
	
	if stats.Memory != nil {
		fmt.Printf("Memory: %.2f%% used (%.2f GB / %.2f GB)\n",
			stats.Memory.Percent,
			float64(stats.Memory.Used)/1024/1024/1024,
			float64(stats.Memory.Total)/1024/1024/1024)
	}
	
	if stats.CPU != nil {
		fmt.Printf("CPU: %.2f%% usage (%d cores)\n",
			stats.CPU.Usage,
			stats.CPU.Count)
	}
	fmt.Println()
}

// clearScreen clears the terminal screen
func clearScreen() {
	fmt.Print("\033[2J")
}

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// PrintTable prints processes in table format
func (pm *ProcessManager) PrintTable(processes []*Process) {
	if len(processes) == 0 {
		fmt.Println("No processes found")
		return
	}
	
	// Table header
	fmt.Printf("%-8s %-20s %-12s %-8s %-8s %-8s %-12s %-20s\n",
		"PID", "NAME", "USER", "CPU%", "MEM(MB)", "THREADS", "STATUS", "COMMAND")
	fmt.Println(strings.Repeat("-", 100))
	
	// Table rows
	for _, proc := range processes {
		command := truncateString(proc.Command, 20)
		fmt.Printf("%-8d %-20s %-12s %-8.2f %-8.2f %-8d %-12s %-20s\n",
			proc.PID,
			truncateString(proc.Name, 20),
			truncateString(proc.User, 12),
			proc.CPUPercent,
			proc.MemoryMB,
			proc.NumThreads,
			truncateString(proc.Status, 12),
			command)
	}
}

// PrintJSON prints processes in JSON format
func (pm *ProcessManager) PrintJSON(data interface{}) {
	// This is a simplified version - in a real implementation,
	// you would use encoding/json
	fmt.Printf("JSON output: %+v\n", data)
}

// PrintCSV prints processes in CSV format
func (pm *ProcessManager) PrintCSV(processes []*Process) {
	if len(processes) == 0 {
		return
	}
	
	// CSV header
	fmt.Println("PID,NAME,USER,CPU_PERCENT,MEMORY_MB,THREADS,STATUS,COMMAND")
	
	// CSV rows
	for _, proc := range processes {
		command := strings.ReplaceAll(proc.Command, ",", ";")
		fmt.Printf("%d,%s,%s,%.2f,%.2f,%d,%s,%s\n",
			proc.PID,
			proc.Name,
			proc.User,
			proc.CPUPercent,
			proc.MemoryMB,
			proc.NumThreads,
			proc.Status,
			command)
	}
}

// PrintTree prints the process tree
func (pm *ProcessManager) PrintTree(tree *ProcessNode) {
	pm.printTreeNode(tree, 0)
}

// printTreeNode prints a process tree node recursively
func (pm *ProcessManager) printTreeNode(node *ProcessNode, depth int) {
	indent := strings.Repeat("  ", depth)
	
	// Print current process
	fmt.Printf("%s├─ PID: %d, Name: %s, User: %s, CPU: %.2f%%, Memory: %.2f MB\n",
		indent,
		node.Process.PID,
		node.Process.Name,
		node.Process.User,
		node.Process.CPUPercent,
		node.Process.MemoryMB)
	
	// Print children
	for _, child := range node.Children {
		pm.printTreeNode(child, depth+1)
	}
}

// PrintStats prints system statistics
func (pm *ProcessManager) PrintStats(stats *SystemStats) {
	fmt.Println("System Statistics")
	fmt.Println("=================")
	fmt.Printf("Hostname: %s\n", stats.Hostname)
	fmt.Printf("OS: %s %s\n", stats.OS, stats.Architecture)
	fmt.Printf("Kernel: %s\n", stats.Kernel)
	fmt.Printf("Uptime: %d seconds\n", stats.Uptime)
	fmt.Printf("Total Processes: %d\n", stats.Processes)
	
	if stats.Memory != nil {
		fmt.Printf("\nMemory Information:\n")
		fmt.Printf("  Total: %.2f GB\n", float64(stats.Memory.Total)/1024/1024/1024)
		fmt.Printf("  Available: %.2f GB\n", float64(stats.Memory.Available)/1024/1024/1024)
		fmt.Printf("  Used: %.2f GB\n", float64(stats.Memory.Used)/1024/1024/1024)
		fmt.Printf("  Free: %.2f GB\n", float64(stats.Memory.Free)/1024/1024/1024)
		fmt.Printf("  Usage: %.2f%%\n", stats.Memory.Percent)
	}
	
	if stats.CPU != nil {
		fmt.Printf("\nCPU Information:\n")
		fmt.Printf("  Model: %s\n", stats.CPU.Model)
		fmt.Printf("  Cores: %d\n", stats.CPU.Count)
		fmt.Printf("  Frequency: %.2f GHz\n", float64(stats.CPU.Frequency)/1000000000)
		fmt.Printf("  Usage: %.2f%%\n", stats.CPU.Usage)
	}
	
	fmt.Printf("\nTimestamp: %s\n", stats.Timestamp.Format("2006-01-02 15:04:05"))
}
