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
		monitorCmd    = flag.Bool("monitor", false, "Monitor network interfaces")
		trafficCmd    = flag.Bool("traffic", false, "Analyze network traffic")
		diagnoseCmd   = flag.Bool("diagnose", false, "Run network diagnostics")
		pingCmd       = flag.Bool("ping", false, "Test connectivity")
		scanCmd       = flag.Bool("scan", false, "Scan network ports")
		statsCmd      = flag.Bool("stats", false, "Show network statistics")
		versionCmd    = flag.Bool("version", false, "Show version information")
		
		// Options
		watch         = flag.Bool("watch", false, "Watch mode (real-time updates)")
		interval      = flag.Duration("interval", 1*time.Second, "Update interval")
		interfaceName = flag.String("interface", "", "Network interface to monitor")
		protocol      = flag.String("protocol", "", "Protocol to filter (tcp, udp, http)")
		host          = flag.String("host", "", "Host to test/scan")
		ports         = flag.String("ports", "1-1000", "Port range to scan")
		timeout       = flag.Duration("timeout", 5*time.Second, "Operation timeout")
		format        = flag.String("format", "table", "Output format (table, json, csv)")
		export        = flag.String("export", "", "Export data to file")
	)
	
	flag.Parse()
	
	// Show version
	if *versionCmd {
		fmt.Printf("Network Monitor v%s (build %s)\n", version, build)
		return
	}
	
	// Create network monitor
	nm := NewNetworkMonitor()
	defer nm.Close()
	
	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nShutting down...")
		nm.Close()
		os.Exit(0)
	}()
	
	// Execute commands
	switch {
	case *monitorCmd:
		handleMonitor(nm, &MonitorOptions{
			Watch:     *watch,
			Interval:  *interval,
			Interface: *interfaceName,
			Format:    *format,
		})
		
	case *trafficCmd:
		handleTraffic(nm, &TrafficOptions{
			Protocol: *protocol,
			Format:   *format,
			Export:   *export,
		})
		
	case *diagnoseCmd:
		handleDiagnose(nm, &DiagnoseOptions{
			Format: *format,
		})
		
	case *pingCmd:
		handlePing(nm, &PingOptions{
			Host:    *host,
			Timeout: *timeout,
			Format:  *format,
		})
		
	case *scanCmd:
		handleScan(nm, &ScanOptions{
			Host:    *host,
			Ports:   *ports,
			Timeout: *timeout,
			Format:  *format,
		})
		
	case *statsCmd:
		handleStats(nm, &StatsOptions{
			Format: *format,
		})
		
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Network Monitor - Advanced Network Analysis")
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("Usage: network-monitor [command] [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  monitor      Monitor network interfaces")
	fmt.Println("  traffic      Analyze network traffic")
	fmt.Println("  diagnose     Run network diagnostics")
	fmt.Println("  ping         Test connectivity")
	fmt.Println("  scan         Scan network ports")
	fmt.Println("  stats        Show network statistics")
	fmt.Println("  version      Show version information")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -watch              Watch mode (real-time updates)")
	fmt.Println("  -interval duration  Update interval (default 1s)")
	fmt.Println("  -interface string   Network interface to monitor")
	fmt.Println("  -protocol string    Protocol to filter (tcp, udp, http)")
	fmt.Println("  -host string        Host to test/scan")
	fmt.Println("  -ports string       Port range to scan (default 1-1000)")
	fmt.Println("  -timeout duration   Operation timeout (default 5s)")
	fmt.Println("  -format string      Output format (table, json, csv)")
	fmt.Println("  -export string      Export data to file")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  network-monitor -monitor")
	fmt.Println("  network-monitor -monitor -watch -interface=eth0")
	fmt.Println("  network-monitor -traffic -protocol=tcp")
	fmt.Println("  network-monitor -ping -host=google.com")
	fmt.Println("  network-monitor -scan -host=192.168.1.1 -ports=1-1000")
}

func handleMonitor(nm *NetworkMonitor, opts *MonitorOptions) {
	if opts.Watch {
		nm.MonitorRealtime(opts)
	} else {
		stats, err := nm.GetNetworkStats(opts)
		if err != nil {
			fmt.Printf("Error getting network stats: %v\n", err)
			return
		}
		
		switch opts.Format {
		case "json":
			nm.PrintJSON(stats)
		case "csv":
			nm.PrintCSV(stats)
		default:
			nm.PrintTable(stats)
		}
	}
}

func handleTraffic(nm *NetworkMonitor, opts *TrafficOptions) {
	traffic, err := nm.AnalyzeTraffic(opts)
	if err != nil {
		fmt.Printf("Error analyzing traffic: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		nm.PrintJSON(traffic)
	case "csv":
		nm.PrintCSV(traffic)
	default:
		nm.PrintTraffic(traffic)
	}
	
	if opts.Export != "" {
		err := nm.ExportTraffic(traffic, opts.Export)
		if err != nil {
			fmt.Printf("Error exporting traffic: %v\n", err)
		} else {
			fmt.Printf("Traffic data exported to: %s\n", opts.Export)
		}
	}
}

func handleDiagnose(nm *NetworkMonitor, opts *DiagnoseOptions) {
	results, err := nm.RunDiagnostics(opts)
	if err != nil {
		fmt.Printf("Error running diagnostics: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		nm.PrintJSON(results)
	case "csv":
		nm.PrintCSV(results)
	default:
		nm.PrintDiagnostics(results)
	}
}

func handlePing(nm *NetworkMonitor, opts *PingOptions) {
	if opts.Host == "" {
		fmt.Println("Host is required for ping")
		return
	}
	
	results, err := nm.PingHost(opts)
	if err != nil {
		fmt.Printf("Error pinging host: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		nm.PrintJSON(results)
	case "csv":
		nm.PrintCSV(results)
	default:
		nm.PrintPingResults(results)
	}
}

func handleScan(nm *NetworkMonitor, opts *ScanOptions) {
	if opts.Host == "" {
		fmt.Println("Host is required for scan")
		return
	}
	
	results, err := nm.ScanPorts(opts)
	if err != nil {
		fmt.Printf("Error scanning ports: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		nm.PrintJSON(results)
	case "csv":
		nm.PrintCSV(results)
	default:
		nm.PrintScanResults(results)
	}
}

func handleStats(nm *NetworkMonitor, opts *StatsOptions) {
	stats, err := nm.GetDetailedStats()
	if err != nil {
		fmt.Printf("Error getting detailed stats: %v\n", err)
		return
	}
	
	switch opts.Format {
	case "json":
		nm.PrintJSON(stats)
	case "csv":
		nm.PrintCSV(stats)
	default:
		nm.PrintDetailedStats(stats)
	}
}
