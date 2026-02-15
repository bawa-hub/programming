package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	psnet "github.com/shirou/gopsutil/v3/net"
)

// NetworkMonitor manages network operations and monitoring
type NetworkMonitor struct {
	interfaces map[string]*NetworkInterface
	mutex      sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
	stats      *NetworkStats
}

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Name         string            `json:"name"`
	Index        int               `json:"index"`
	MTU          int               `json:"mtu"`
	HardwareAddr string            `json:"hardware_addr"`
	Flags        []string          `json:"flags"`
	Addresses    []string          `json:"addresses"`
	Stats        *InterfaceStats   `json:"stats"`
	IsUp         bool              `json:"is_up"`
	LastSeen     time.Time         `json:"last_seen"`
}

// InterfaceStats represents interface statistics
type InterfaceStats struct {
	BytesSent     uint64    `json:"bytes_sent"`
	BytesRecv     uint64    `json:"bytes_recv"`
	PacketsSent   uint64    `json:"packets_sent"`
	PacketsRecv   uint64    `json:"packets_recv"`
	Errin         uint64    `json:"errin"`
	Errout        uint64    `json:"errout"`
	Dropin        uint64    `json:"dropin"`
	Dropout       uint64    `json:"dropout"`
	Timestamp     time.Time `json:"timestamp"`
}

// NetworkStats represents overall network statistics
type NetworkStats struct {
	Timestamp     time.Time                    `json:"timestamp"`
	Interfaces    map[string]*NetworkInterface `json:"interfaces"`
	TotalBytesSent uint64                      `json:"total_bytes_sent"`
	TotalBytesRecv uint64                      `json:"total_bytes_recv"`
	TotalPacketsSent uint64                    `json:"total_packets_sent"`
	TotalPacketsRecv uint64                    `json:"total_packets_recv"`
	ActiveConnections int                      `json:"active_connections"`
	OpenPorts     []int                       `json:"open_ports"`
}

// MonitorOptions contains options for network monitoring
type MonitorOptions struct {
	Watch     bool
	Interval  time.Duration
	Interface string
	Format    string
}

// StatsOptions contains options for statistics
type StatsOptions struct {
	Format string
}

// NewNetworkMonitor creates a new network monitor
func NewNetworkMonitor() *NetworkMonitor {
	ctx, cancel := context.WithCancel(context.Background())
	return &NetworkMonitor{
		interfaces: make(map[string]*NetworkInterface),
		ctx:        ctx,
		cancel:     cancel,
		stats:      &NetworkStats{},
	}
}

// Close closes the network monitor
func (nm *NetworkMonitor) Close() {
	nm.cancel()
}

// GetNetworkStats gets current network statistics
func (nm *NetworkMonitor) GetNetworkStats(opts *MonitorOptions) (*NetworkStats, error) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()
	
	// Get network interfaces
	interfaces, err := psnet.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %w", err)
	}
	
	// Get interface statistics
	interfaceStats, err := psnet.IOCounters(true)
	if err != nil {
		return nil, fmt.Errorf("failed to get interface stats: %w", err)
	}
	
	// Get connections
	connections, err := psnet.Connections("all")
	if err != nil {
		return nil, fmt.Errorf("failed to get connections: %w", err)
	}
	
	// Update interfaces
	nm.interfaces = make(map[string]*NetworkInterface)
	totalBytesSent := uint64(0)
	totalBytesRecv := uint64(0)
	totalPacketsSent := uint64(0)
	totalPacketsRecv := uint64(0)
	
	for _, iface := range interfaces {
		// Find corresponding stats
		var stats *InterfaceStats
		for _, stat := range interfaceStats {
			if stat.Name == iface.Name {
				stats = &InterfaceStats{
					BytesSent:   stat.BytesSent,
					BytesRecv:   stat.BytesRecv,
					PacketsSent: stat.PacketsSent,
					PacketsRecv: stat.PacketsRecv,
					Errin:       stat.Errin,
					Errout:      stat.Errout,
					Dropin:      stat.Dropin,
					Dropout:     stat.Dropout,
					Timestamp:   time.Now(),
				}
				break
			}
		}
		
		// Get addresses
		addresses := make([]string, 0)
		for _, addr := range iface.Addrs {
			addresses = append(addresses, addr.Addr)
		}
		
		// Create interface
		netIface := &NetworkInterface{
			Name:         iface.Name,
			Index:        iface.Index,
			MTU:          iface.MTU,
			HardwareAddr: iface.HardwareAddr,
			Flags:        iface.Flags,
			Addresses:    addresses,
			Stats:        stats,
			IsUp:         len(iface.Flags) > 0 && iface.Flags[0] == "up",
			LastSeen:     time.Now(),
		}
		
		nm.interfaces[iface.Name] = netIface
		
		// Update totals
		if stats != nil {
			totalBytesSent += stats.BytesSent
			totalBytesRecv += stats.BytesRecv
			totalPacketsSent += stats.PacketsSent
			totalPacketsRecv += stats.PacketsRecv
		}
	}
	
	// Update stats
	nm.stats = &NetworkStats{
		Timestamp:        time.Now(),
		Interfaces:       nm.interfaces,
		TotalBytesSent:   totalBytesSent,
		TotalBytesRecv:   totalBytesRecv,
		TotalPacketsSent: totalPacketsSent,
		TotalPacketsRecv: totalPacketsRecv,
		ActiveConnections: len(connections),
		OpenPorts:        nm.getOpenPorts(),
	}
	
	return nm.stats, nil
}

// getOpenPorts gets currently open ports
func (nm *NetworkMonitor) getOpenPorts() []int {
	ports := make([]int, 0)
	
	// Common ports to check
	commonPorts := []int{22, 23, 25, 53, 80, 110, 143, 443, 993, 995, 3389, 5432, 3306, 6379, 27017}
	
	for _, port := range commonPorts {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), 100*time.Millisecond)
		if err == nil {
			ports = append(ports, port)
			conn.Close()
		}
	}
	
	return ports
}

// MonitorRealtime monitors network in real-time
func (nm *NetworkMonitor) MonitorRealtime(opts *MonitorOptions) {
	fmt.Println("Starting real-time network monitoring...")
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()
	
	// Create ticker for updates
	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()
	
	// Clear screen
	clearScreen()
	
	startTime := time.Now()
	
	for {
		select {
		case <-ticker.C:
			nm.updateMonitorDisplay(opts, startTime)
		case <-nm.ctx.Done():
			return
		}
	}
}

// updateMonitorDisplay updates the monitoring display
func (nm *NetworkMonitor) updateMonitorDisplay(opts *MonitorOptions, startTime time.Time) {
	// Move cursor to top
	fmt.Print("\033[H")
	
	// Get network stats
	stats, err := nm.GetNetworkStats(opts)
	if err != nil {
		fmt.Printf("Error getting network stats: %v\n", err)
		return
	}
	
	// Print header
	fmt.Println("Network Monitor - Real-time")
	fmt.Println("===========================")
	fmt.Printf("Time: %s (Running: %v)\n", 
		time.Now().Format("2006-01-02 15:04:05"), 
		time.Since(startTime).Round(time.Second))
	fmt.Println()
	
	// Print network stats
	nm.printNetworkStats(stats, opts.Interface)
}

// printNetworkStats prints network statistics
func (nm *NetworkMonitor) printNetworkStats(stats *NetworkStats, filterInterface string) {
	fmt.Printf("Network Overview:\n")
	fmt.Printf("  Total Bytes Sent: %s\n", formatBytes(stats.TotalBytesSent))
	fmt.Printf("  Total Bytes Recv: %s\n", formatBytes(stats.TotalBytesRecv))
	fmt.Printf("  Total Packets Sent: %d\n", stats.TotalPacketsSent)
	fmt.Printf("  Total Packets Recv: %d\n", stats.TotalPacketsRecv)
	fmt.Printf("  Active Connections: %d\n", stats.ActiveConnections)
	fmt.Printf("  Open Ports: %v\n", stats.OpenPorts)
	fmt.Println()
	
	// Print interfaces
	fmt.Println("Network Interfaces:")
	fmt.Printf("%-15s %-8s %-15s %-15s %-10s %-10s %-8s\n",
		"Interface", "Status", "Bytes Sent", "Bytes Recv", "Packets", "Errors", "Addresses")
	fmt.Println(strings.Repeat("-", 90))
	
	for name, iface := range stats.Interfaces {
		if filterInterface != "" && name != filterInterface {
			continue
		}
		
		status := "DOWN"
		if iface.IsUp {
			status = "UP"
		}
		
		packets := uint64(0)
		errors := uint64(0)
		if iface.Stats != nil {
			packets = iface.Stats.PacketsSent + iface.Stats.PacketsRecv
			errors = iface.Stats.Errin + iface.Stats.Errout + iface.Stats.Dropin + iface.Stats.Dropout
		}
		
		addresses := strings.Join(iface.Addresses, ", ")
		if len(addresses) > 15 {
			addresses = addresses[:12] + "..."
		}
		
		fmt.Printf("%-15s %-8s %-15s %-15s %-10d %-10d %-8s\n",
			name,
			status,
			formatBytes(iface.Stats.BytesSent),
			formatBytes(iface.Stats.BytesRecv),
			packets,
			errors,
			addresses)
	}
	fmt.Println()
}

// clearScreen clears the terminal screen
func clearScreen() {
	fmt.Print("\033[2J")
}

// formatBytes formats bytes into human readable format
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// PrintTable prints network stats in table format
func (nm *NetworkMonitor) PrintTable(stats *NetworkStats) {
	fmt.Printf("%-15s %-8s %-15s %-15s %-10s %-10s %-8s\n",
		"Interface", "Status", "Bytes Sent", "Bytes Recv", "Packets", "Errors", "Addresses")
	fmt.Println(strings.Repeat("-", 90))
	
	for name, iface := range stats.Interfaces {
		status := "DOWN"
		if iface.IsUp {
			status = "UP"
		}
		
		packets := uint64(0)
		errors := uint64(0)
		if iface.Stats != nil {
			packets = iface.Stats.PacketsSent + iface.Stats.PacketsRecv
			errors = iface.Stats.Errin + iface.Stats.Errout + iface.Stats.Dropin + iface.Stats.Dropout
		}
		
		addresses := strings.Join(iface.Addresses, ", ")
		if len(addresses) > 15 {
			addresses = addresses[:12] + "..."
		}
		
		fmt.Printf("%-15s %-8s %-15s %-15s %-10d %-10d %-8s\n",
			name,
			status,
			formatBytes(iface.Stats.BytesSent),
			formatBytes(iface.Stats.BytesRecv),
			packets,
			errors,
			addresses)
	}
}

// PrintJSON prints data in JSON format
func (nm *NetworkMonitor) PrintJSON(data interface{}) {
	// This is a simplified version - in a real implementation,
	// you would use encoding/json
	fmt.Printf("JSON output: %+v\n", data)
}

// PrintCSV prints data in CSV format
func (nm *NetworkMonitor) PrintCSV(data interface{}) {
	// This is a simplified version - in a real implementation,
	// you would use encoding/csv
	fmt.Printf("CSV output: %+v\n", data)
}

// GetDetailedStats gets detailed network statistics
func (nm *NetworkMonitor) GetDetailedStats() (*DetailedStats, error) {
	stats, err := nm.GetNetworkStats(&MonitorOptions{})
	if err != nil {
		return nil, err
	}
	
	return &DetailedStats{
		Network:    stats,
		Timestamp:  time.Now(),
	}, nil
}

// DetailedStats represents detailed network statistics
type DetailedStats struct {
	Network   *NetworkStats `json:"network"`
	Timestamp time.Time     `json:"timestamp"`
}

// PrintDetailedStats prints detailed network statistics
func (nm *NetworkMonitor) PrintDetailedStats(stats *DetailedStats) {
	fmt.Println("Detailed Network Statistics")
	fmt.Println("===========================")
	fmt.Printf("Timestamp: %s\n", stats.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println()
	
	// Print network stats
	nm.printNetworkStats(stats.Network, "")
}
