package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🐳 Container Runtime - Test")
	fmt.Println("===========================")
	
	// Create container runtime
	cr := NewContainerRuntime()
	defer cr.Close()
	
	// Test basic functionality
	fmt.Println("\n📦 Testing container creation...")
	container, err := cr.CreateContainer(&CreateOptions{
		Name:    "test-container",
		Image:   "ubuntu:latest",
		Command: "echo 'Hello from container!' && sleep 10",
		Env:     "ENV1=value1,ENV2=value2",
		Ports:   "8080:80,3306:3306",
		Volumes: "/host/path:/container/path",
		Network: "bridge",
	})
	if err != nil {
		fmt.Printf("Error creating container: %v\n", err)
		return
	}
	
	fmt.Printf("Container created: %s (ID: %s)\n", container.Name, container.ID)
	fmt.Printf("Image: %s\n", container.Image)
	fmt.Printf("Status: %s\n", container.Status)
	fmt.Printf("Environment: %v\n", container.Environment)
	fmt.Printf("Ports: %v\n", container.Ports)
	fmt.Printf("Volumes: %v\n", container.Volumes)
	
	// Test starting container
	fmt.Println("\n🚀 Testing container start...")
	err = cr.StartContainer("test-container")
	if err != nil {
		fmt.Printf("Error starting container: %v\n", err)
	} else {
		fmt.Println("Container started successfully")
	}
	
	// Wait a bit
	time.Sleep(2 * time.Second)
	
	// Test container stats
	fmt.Println("\n📊 Testing container stats...")
	stats, err := cr.GetContainerStats("test-container")
	if err != nil {
		fmt.Printf("Error getting container stats: %v\n", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", stats.CPUUsage)
		fmt.Printf("Memory Usage: %s / %s\n", formatBytes(stats.MemoryUsage), formatBytes(stats.MemoryLimit))
		fmt.Printf("Network: %s↓ %s↑\n", formatBytes(stats.NetworkRx), formatBytes(stats.NetworkTx))
		fmt.Printf("PIDs: %d\n", stats.PIDs)
	}
	
	// Test container execution
	fmt.Println("\n⚡ Testing container execution...")
	output, err := cr.ExecContainer("test-container", "ls -la")
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	} else {
		fmt.Printf("Command output:\n%s\n", output)
	}
	
	// Test container logs
	fmt.Println("\n📝 Testing container logs...")
	logs, err := cr.GetContainerLogs("test-container")
	if err != nil {
		fmt.Printf("Error getting container logs: %v\n", err)
	} else {
		fmt.Printf("Container logs (%d entries):\n", len(logs))
		for i, log := range logs {
			if i < 5 { // Show first 5 logs
				fmt.Printf("  %s", log)
			}
		}
		if len(logs) > 5 {
			fmt.Printf("  ... and %d more logs\n", len(logs)-5)
		}
	}
	
	// Test container health
	fmt.Println("\n🏥 Testing container health...")
	health, err := cr.CheckContainerHealth("test-container")
	if err != nil {
		fmt.Printf("Error checking container health: %v\n", err)
	} else {
		fmt.Printf("Container health: %s\n", health.Status)
	}
	
	// Test container pause/resume
	fmt.Println("\n⏸️  Testing container pause...")
	err = cr.PauseContainer("test-container")
	if err != nil {
		fmt.Printf("Error pausing container: %v\n", err)
	} else {
		fmt.Println("Container paused successfully")
	}
	
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n▶️  Testing container resume...")
	err = cr.ResumeContainer("test-container")
	if err != nil {
		fmt.Printf("Error resuming container: %v\n", err)
	} else {
		fmt.Println("Container resumed successfully")
	}
	
	// Test container scaling
	fmt.Println("\n📈 Testing container scaling...")
	err = cr.ScaleContainer("test-container", 3)
	if err != nil {
		fmt.Printf("Error scaling container: %v\n", err)
	} else {
		fmt.Println("Container scaled to 3 replicas")
	}
	
	// Test container restart
	fmt.Println("\n🔄 Testing container restart...")
	err = cr.RestartContainer("test-container", "always")
	if err != nil {
		fmt.Printf("Error restarting container: %v\n", err)
	} else {
		fmt.Println("Container restarted successfully")
	}
	
	// Test container listing
	fmt.Println("\n📋 Testing container listing...")
	containers, err := cr.ListContainers()
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
	} else {
		fmt.Printf("Found %d containers:\n", len(containers))
		for _, container := range containers {
			fmt.Printf("  - %s (%s) - %s\n", container.Name, container.Image, container.Status)
		}
	}
	
	// Test container deployment
	fmt.Println("\n🚀 Testing container deployment...")
	err = cr.DeployContainers("compose.yaml")
	if err != nil {
		fmt.Printf("Error deploying containers: %v\n", err)
	} else {
		fmt.Println("Containers deployed successfully")
	}
	
	// List all containers after deployment
	fmt.Println("\n📋 Final container list:")
	containers, err = cr.ListContainers()
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
	} else {
		fmt.Printf("Found %d containers:\n", len(containers))
		for _, container := range containers {
			fmt.Printf("  - %s (%s) - %s\n", container.Name, container.Image, container.Status)
		}
	}
	
	// Test stopping container
	fmt.Println("\n🛑 Testing container stop...")
	err = cr.StopContainer("test-container")
	if err != nil {
		fmt.Printf("Error stopping container: %v\n", err)
	} else {
		fmt.Println("Container stopped successfully")
	}
	
	// Test removing container
	fmt.Println("\n🗑️  Testing container removal...")
	err = cr.RemoveContainer("test-container")
	if err != nil {
		fmt.Printf("Error removing container: %v\n", err)
	} else {
		fmt.Println("Container removed successfully")
	}
	
	fmt.Println("\n✅ Container Runtime test completed!")
}
