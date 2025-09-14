package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// TestProcessManager tests the process manager functionality
func TestProcessManager() {
	fmt.Println("🧪 Testing Process Manager...")
	fmt.Println("=============================")
	
	// Create process manager
	pm := NewProcessManager()
	defer pm.Close()
	
	// Test 1: List processes
	fmt.Println("\n1. Testing process listing...")
	processes, err := pm.ListProcesses(&ListOptions{Limit: 5})
	if err != nil {
		fmt.Printf("❌ Error listing processes: %v\n", err)
	} else {
		fmt.Printf("✅ Listed %d processes\n", len(processes))
		for i, proc := range processes {
			if i >= 3 { // Show only first 3
				break
			}
			fmt.Printf("   PID: %d, Name: %s, User: %s\n", 
				proc.PID, proc.Name, proc.User)
		}
	}
	
	// Test 2: Start a test process
	fmt.Println("\n2. Testing process start...")
	testCmd := "sleep"
	testArgs := "5"
	
	pid, err := pm.StartProcess(&StartOptions{
		Command: testCmd,
		Args:    testArgs,
	})
	if err != nil {
		fmt.Printf("❌ Error starting process: %v\n", err)
	} else {
		fmt.Printf("✅ Started process with PID: %d\n", pid)
		
		// Wait a bit
		time.Sleep(1 * time.Second)
		
		// Test 3: Monitor specific process
		fmt.Println("\n3. Testing process monitoring...")
		proc, err := pm.GetProcesses(&MonitorOptions{PID: pid})
		if err != nil {
			fmt.Printf("❌ Error monitoring process: %v\n", err)
		} else {
			fmt.Printf("✅ Monitored process: %+v\n", proc[0].Name)
		}
		
		// Test 4: Stop the process
		fmt.Println("\n4. Testing process stop...")
		err = pm.StopProcess(pid)
		if err != nil {
			fmt.Printf("❌ Error stopping process: %v\n", err)
		} else {
			fmt.Printf("✅ Stopped process %d\n", pid)
		}
	}
	
	// Test 5: System statistics
	fmt.Println("\n5. Testing system statistics...")
	stats, err := pm.GetSystemStats()
	if err != nil {
		fmt.Printf("❌ Error getting system stats: %v\n", err)
	} else {
		fmt.Printf("✅ System stats: %s, %d processes\n", 
			stats.Hostname, stats.Processes)
	}
	
	// Test 6: Process tree
	fmt.Println("\n6. Testing process tree...")
	tree, err := pm.GetProcessTree(&TreeOptions{})
	if err != nil {
		fmt.Printf("❌ Error getting process tree: %v\n", err)
	} else {
		fmt.Printf("✅ Process tree root: PID %d, Name: %s\n", 
			tree.Process.PID, tree.Process.Name)
		fmt.Printf("   Children: %d\n", len(tree.Children))
	}
	
	fmt.Println("\n🎉 Process Manager tests completed!")
}

// RunBasicTest runs a basic functionality test
func RunBasicTest() {
	fmt.Println("🚀 Process Manager - Basic Test")
	fmt.Println("===============================")
	
	// Create process manager
	pm := NewProcessManager()
	defer pm.Close()
	
	// List some processes
	fmt.Println("\n📋 Listing processes...")
	processes, err := pm.ListProcesses(&ListOptions{Limit: 10})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Print table
	pm.PrintTable(processes)
	
	// Show system stats
	fmt.Println("\n📊 System Statistics:")
	stats, err := pm.GetSystemStats()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	pm.PrintStats(stats)
}

// RunInteractiveTest runs an interactive test
func RunInteractiveTest() {
	fmt.Println("🚀 Process Manager - Interactive Test")
	fmt.Println("=====================================")
	
	// Create process manager
	pm := NewProcessManager()
	defer pm.Close()
	
	// Start a background process
	fmt.Println("\n🔄 Starting background process...")
	pid, err := pm.StartProcess(&StartOptions{
		Command: "sleep",
		Args:    "30",
	})
	if err != nil {
		fmt.Printf("Error starting process: %v\n", err)
		return
	}
	
	fmt.Printf("Started process with PID: %d\n", pid)
	
	// Monitor for a few seconds
	fmt.Println("\n👀 Monitoring process...")
	for i := 0; i < 5; i++ {
		processes, err := pm.GetProcesses(&MonitorOptions{PID: pid})
		if err != nil {
			fmt.Printf("Error monitoring: %v\n", err)
			break
		}
		
		if len(processes) > 0 {
			proc := processes[0]
			fmt.Printf("PID: %d, Name: %s, CPU: %.2f%%, Memory: %.2f MB\n",
				proc.PID, proc.Name, proc.CPUPercent, proc.MemoryMB)
		}
		
		time.Sleep(1 * time.Second)
	}
	
	// Stop the process
	fmt.Println("\n🛑 Stopping process...")
	err = pm.StopProcess(pid)
	if err != nil {
		fmt.Printf("Error stopping process: %v\n", err)
	} else {
		fmt.Printf("Process %d stopped successfully\n", pid)
	}
	
	fmt.Println("\n✅ Interactive test completed!")
}

// RunPerformanceTest runs a performance test
func RunPerformanceTest() {
	fmt.Println("🚀 Process Manager - Performance Test")
	fmt.Println("====================================")
	
	// Create process manager
	pm := NewProcessManager()
	defer pm.Close()
	
	// Test process listing performance
	fmt.Println("\n⏱️  Testing process listing performance...")
	start := time.Now()
	
	processes, err := pm.ListProcesses(&ListOptions{Limit: 100})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	duration := time.Since(start)
	fmt.Printf("Listed %d processes in %v\n", len(processes), duration)
	
	// Test monitoring performance
	fmt.Println("\n⏱️  Testing monitoring performance...")
	start = time.Now()
	
	for i := 0; i < 10; i++ {
		_, err := pm.GetProcesses(&MonitorOptions{})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
	
	duration = time.Since(start)
	fmt.Printf("10 monitoring cycles completed in %v (avg: %v per cycle)\n", 
		duration, duration/10)
	
	// Test system stats performance
	fmt.Println("\n⏱️  Testing system stats performance...")
	start = time.Now()
	
	for i := 0; i < 10; i++ {
		_, err := pm.GetSystemStats()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
	
	duration = time.Since(start)
	fmt.Printf("10 system stats calls completed in %v (avg: %v per call)\n", 
		duration, duration/10)
	
	fmt.Println("\n✅ Performance test completed!")
}
