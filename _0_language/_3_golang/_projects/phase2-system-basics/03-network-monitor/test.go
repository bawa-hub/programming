package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🌐 Network Monitor - Test")
	fmt.Println("========================")
	
	// Create network monitor
	nm := NewNetworkMonitor()
	defer nm.Close()
	
	// Test basic functionality
	fmt.Println("\n📊 Getting network statistics...")
	stats, err := nm.GetNetworkStats(&MonitorOptions{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Print basic stats
	fmt.Printf("Total Bytes Sent: %s\n", formatBytes(stats.TotalBytesSent))
	fmt.Printf("Total Bytes Recv: %s\n", formatBytes(stats.TotalBytesRecv))
	fmt.Printf("Total Packets Sent: %d\n", stats.TotalPacketsSent)
	fmt.Printf("Total Packets Recv: %d\n", stats.TotalPacketsRecv)
	fmt.Printf("Active Connections: %d\n", stats.ActiveConnections)
	fmt.Printf("Open Ports: %v\n", stats.OpenPorts)
	
	// Test traffic analysis
	fmt.Println("\n🚦 Testing traffic analysis...")
	traffic, err := nm.AnalyzeTraffic(&TrafficOptions{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Total Packets: %d\n", traffic.TotalPackets)
		fmt.Printf("Total Bytes: %s\n", formatBytes(traffic.TotalBytes))
		fmt.Printf("Protocols: %d\n", len(traffic.Protocols))
		fmt.Printf("Top Hosts: %d\n", len(traffic.TopHosts))
		fmt.Printf("Top Ports: %d\n", len(traffic.TopPorts))
		fmt.Printf("Connections: %d\n", len(traffic.Connections))
	}
	
	// Test diagnostics
	fmt.Println("\n🔍 Testing network diagnostics...")
	diagnostics, err := nm.RunDiagnostics(&DiagnoseOptions{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Overall Score: %d/100\n", diagnostics.Overall.Score)
		fmt.Printf("Overall Status: %s\n", diagnostics.Overall.Status)
		fmt.Printf("Internet: %t\n", diagnostics.Connectivity.Internet)
		fmt.Printf("Gateway: %t\n", diagnostics.Connectivity.Gateway)
		fmt.Printf("DNS: %t\n", diagnostics.DNS.Working)
		fmt.Printf("Interfaces Up: %d/%d\n", diagnostics.Interfaces.Up, diagnostics.Interfaces.Count)
	}
	
	// Test ping
	fmt.Println("\n🏓 Testing ping...")
	pingResults, err := nm.PingHost(&PingOptions{
		Host:    "google.com",
		Timeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Host: %s\n", pingResults.Host)
		fmt.Printf("Success: %t\n", pingResults.Success)
		if pingResults.Success {
			fmt.Printf("Latency: %v\n", pingResults.Latency)
			fmt.Printf("Packet Loss: %.2f%%\n", pingResults.PacketLoss)
		}
	}
	
	// Test port scan
	fmt.Println("\n🔍 Testing port scan...")
	scanResults, err := nm.ScanPorts(&ScanOptions{
		Host:    "localhost",
		Ports:   "22,80,443",
		Timeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Host: %s\n", scanResults.Host)
		fmt.Printf("Scan Time: %v\n", scanResults.ScanTime)
		fmt.Printf("Open Ports: %d\n", scanResults.OpenPorts)
		fmt.Printf("Closed Ports: %d\n", scanResults.ClosedPorts)
		fmt.Printf("Filtered Ports: %d\n", scanResults.FilteredPorts)
	}
	
	fmt.Println("\n✅ Network Monitor test completed!")
}
