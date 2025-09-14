package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// RealContainerRuntime represents a real container runtime
type RealContainerRuntime struct {
	namespaceManager  *NamespaceManager
	cgroupManager     *CGroupManager
	filesystemManager *FilesystemManager
	networkManager    *NetworkManager
	containers        map[string]*RealContainer
}

// RealContainer represents a real container with full isolation
type RealContainer struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Image        string                 `json:"image"`
	Command      string                 `json:"command"`
	Status       string                 `json:"status"`
	Created      time.Time              `json:"created"`
	Started      time.Time              `json:"started"`
	Pid          int                    `json:"pid"`
	Namespaces   map[string]int         `json:"namespaces"`
	ResourceLimits *ResourceLimits      `json:"resource_limits"`
	Filesystem   *ContainerFilesystem   `json:"filesystem"`
	Network      *ContainerNetwork      `json:"network"`
	Process      *exec.Cmd              `json:"-"`
}

// NewRealContainerRuntime creates a new real container runtime
func NewRealContainerRuntime() *RealContainerRuntime {
	return &RealContainerRuntime{
		namespaceManager:  NewNamespaceManager(),
		containers:        make(map[string]*RealContainer),
	}
}

// CreateContainer creates a real container with full isolation
func (rtr *RealContainerRuntime) CreateContainer(name, image, command string, limits *ResourceLimits) (*RealContainer, error) {
	fmt.Printf("🐳 Creating real container: %s\n", name)
	
	// Check if running as root
	if os.Geteuid() != 0 {
		return nil, fmt.Errorf("root privileges required for real containerization")
	}
	
	// Check if running on Linux
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("real containerization only supported on Linux")
	}
	
	// Create container
	container := &RealContainer{
		ID:            generateContainerID(),
		Name:          name,
		Image:         image,
		Command:       command,
		Status:        "created",
		Created:       time.Now(),
		ResourceLimits: limits,
		Namespaces:    make(map[string]int),
	}
	
	// Create cgroup manager
	rtr.cgroupManager = NewCGroupManager(container.ID)
	
	// Create filesystem manager
	rtr.filesystemManager = NewFilesystemManager(container.ID)
	
	// Create network manager
	rtr.networkManager = NewNetworkManager(container.ID)
	
	// Create container filesystem
	filesystem, err := rtr.filesystemManager.CreateContainerFilesystem(image)
	if err != nil {
		return nil, fmt.Errorf("failed to create filesystem: %w", err)
	}
	container.Filesystem = filesystem
	
	// Create container network
	network, err := rtr.networkManager.CreateContainerNetwork()
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %w", err)
	}
	container.Network = network
	
	// Create cgroup
	if err := rtr.cgroupManager.CreateCGroup(); err != nil {
		return nil, fmt.Errorf("failed to create cgroup: %w", err)
	}
	
	// Set resource limits
	if limits != nil {
		if err := rtr.cgroupManager.SetResourceLimits(limits); err != nil {
			return nil, fmt.Errorf("failed to set resource limits: %w", err)
		}
	}
	
	// Store container
	rtr.containers[container.ID] = container
	
	fmt.Printf("✅ Real container created: %s (ID: %s)\n", name, container.ID)
	return container, nil
}

// StartContainer starts a real container
func (rtr *RealContainerRuntime) StartContainer(containerID string) error {
	container, exists := rtr.containers[containerID]
	if !exists {
		return fmt.Errorf("container not found: %s", containerID)
	}
	
	fmt.Printf("🚀 Starting real container: %s\n", container.Name)
	
	// Create isolated process with namespaces
	cmd, err := rtr.namespaceManager.CreateIsolatedProcess(container.Name, container.Command)
	if err != nil {
		return fmt.Errorf("failed to create isolated process: %w", err)
	}
	
	// Set up process attributes
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Set working directory to container root
	cmd.Dir = container.Filesystem.MergedPath
	
	// Set environment variables
	cmd.Env = append(os.Environ(), "HOSTNAME="+container.Network.Hostname)
	
	// Start the process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}
	
	// Add process to cgroup
	if err := rtr.cgroupManager.AddProcess(cmd.Process.Pid); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %w", err)
	}
	
	// Update container status
	container.Status = "running"
	container.Started = time.Now()
	container.Pid = cmd.Process.Pid
	container.Process = cmd
	
	// Start monitoring goroutine
	go rtr.monitorContainer(container)
	
	fmt.Printf("✅ Real container started: %s (PID: %d)\n", container.Name, container.Pid)
	return nil
}

// StopContainer stops a real container
func (rtr *RealContainerRuntime) StopContainer(containerID string) error {
	container, exists := rtr.containers[containerID]
	if !exists {
		return fmt.Errorf("container not found: %s", containerID)
	}
	
	fmt.Printf("🛑 Stopping real container: %s\n", container.Name)
	
	// Send SIGTERM to container process
	if container.Process != nil && container.Process.Process != nil {
		if err := container.Process.Process.Signal(syscall.SIGTERM); err != nil {
			return fmt.Errorf("failed to send SIGTERM: %w", err)
		}
		
		// Wait for process to terminate
		done := make(chan error, 1)
		go func() {
			done <- container.Process.Wait()
		}()
		
		select {
		case <-time.After(10 * time.Second):
			// Force kill if not terminated
			container.Process.Process.Kill()
			<-done
		case <-done:
			// Process terminated gracefully
		}
	}
	
	// Update container status
	container.Status = "stopped"
	
	fmt.Printf("✅ Real container stopped: %s\n", container.Name)
	return nil
}

// RemoveContainer removes a real container
func (rtr *RealContainerRuntime) RemoveContainer(containerID string) error {
	container, exists := rtr.containers[containerID]
	if !exists {
		return fmt.Errorf("container not found: %s", containerID)
	}
	
	fmt.Printf("🗑️ Removing real container: %s\n", container.Name)
	
	// Stop container if running
	if container.Status == "running" {
		if err := rtr.StopContainer(containerID); err != nil {
			return fmt.Errorf("failed to stop container: %w", err)
		}
	}
	
	// Cleanup cgroup
	if err := rtr.cgroupManager.CleanupCGroup(); err != nil {
		fmt.Printf("Warning: failed to cleanup cgroup: %v\n", err)
	}
	
	// Cleanup filesystem
	if err := rtr.filesystemManager.CleanupFilesystem(container.Filesystem); err != nil {
		fmt.Printf("Warning: failed to cleanup filesystem: %v\n", err)
	}
	
	// Cleanup network
	if err := rtr.networkManager.CleanupNetwork(container.Network); err != nil {
		fmt.Printf("Warning: failed to cleanup network: %v\n", err)
	}
	
	// Remove from containers map
	delete(rtr.containers, containerID)
	
	fmt.Printf("✅ Real container removed: %s\n", container.Name)
	return nil
}

// ListContainers lists all containers
func (rtr *RealContainerRuntime) ListContainers() {
	fmt.Println("🐳 Real Containers")
	fmt.Println("==================")
	
	if len(rtr.containers) == 0 {
		fmt.Println("No containers found")
		return
	}
	
	for _, container := range rtr.containers {
		fmt.Printf("ID: %s\n", container.ID)
		fmt.Printf("Name: %s\n", container.Name)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Printf("Created: %s\n", container.Created.Format("2006-01-02 15:04:05"))
		if container.Status == "running" {
			fmt.Printf("Started: %s\n", container.Started.Format("2006-01-02 15:04:05"))
			fmt.Printf("PID: %d\n", container.Pid)
		}
		fmt.Printf("IP: %s\n", container.Network.ContainerIP)
		fmt.Println("---")
	}
}

// GetContainerInfo returns detailed container information
func (rtr *RealContainerRuntime) GetContainerInfo(containerID string) (*ContainerInfo, error) {
	container, exists := rtr.containers[containerID]
	if !exists {
		return nil, fmt.Errorf("container not found: %s", containerID)
	}
	
	info := &ContainerInfo{
		Container: container,
	}
	
	// Get resource usage
	if usage, err := rtr.cgroupManager.GetResourceUsage(); err == nil {
		info.ResourceUsage = usage
	}
	
	// Get filesystem info
	if fsInfo, err := rtr.filesystemManager.GetFilesystemInfo(container.Filesystem); err == nil {
		info.FilesystemInfo = fsInfo
	}
	
	// Get network info
	if netInfo, err := rtr.networkManager.GetNetworkInfo(container.Network); err == nil {
		info.NetworkInfo = netInfo
	}
	
	return info, nil
}

// monitorContainer monitors container resource usage
func (rtr *RealContainerRuntime) monitorContainer(container *RealContainer) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			// Check if container is still running
			if container.Process != nil && container.Process.Process != nil {
				if err := container.Process.Process.Signal(syscall.Signal(0)); err != nil {
					// Process is not running
					container.Status = "stopped"
					return
				}
			}
			
			// Update resource usage
			if usage, err := rtr.cgroupManager.GetResourceUsage(); err == nil {
				fmt.Printf("📊 Container %s - Memory: %s, CPU: %.2fs\n", 
					container.Name, 
					formatBytes(usage.MemoryUsage), 
					usage.CPUUsage)
			}
		}
	}
}

// ContainerInfo represents detailed container information
type ContainerInfo struct {
	Container      *RealContainer      `json:"container"`
	ResourceUsage  *ResourceUsage      `json:"resource_usage"`
	FilesystemInfo *FilesystemInfo     `json:"filesystem_info"`
	NetworkInfo    *NetworkInfo        `json:"network_info"`
}

// generateContainerID generates a unique container ID
func generateContainerID() string {
	return fmt.Sprintf("real-%d", time.Now().UnixNano())
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

func main() {
	// Check if running as root
	if os.Geteuid() != 0 {
		fmt.Println("❌ Root privileges required for real containerization")
		fmt.Println("Please run with: sudo go run main.go")
		os.Exit(1)
	}
	
	// Check if running on Linux
	if runtime.GOOS != "linux" {
		fmt.Println("❌ Real containerization only supported on Linux")
		os.Exit(1)
	}
	
	// Parse command line arguments
	var (
		action      = flag.String("action", "", "Action to perform (create, start, stop, remove, list, info)")
		name        = flag.String("name", "", "Container name")
		image       = flag.String("image", "ubuntu:latest", "Container image")
		command     = flag.String("command", "sh", "Command to run")
		memory      = flag.String("memory", "512m", "Memory limit")
		cpus        = flag.String("cpus", "1.0", "CPU limit")
		ports       = flag.String("ports", "", "Port mappings (e.g., 8080:80,9090:9090)")
		containerID = flag.String("id", "", "Container ID")
	)
	flag.Parse()
	
	// Create real container runtime
	rtr := NewRealContainerRuntime()
	
	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n🛑 Shutting down real container runtime...")
		os.Exit(0)
	}()
	
	// Handle actions
	switch *action {
	case "create":
		if *name == "" {
			fmt.Println("❌ Container name required")
			os.Exit(1)
		}
		
		// Parse resource limits
		limits := &ResourceLimits{
			MemoryLimit: parseMemoryLimit(*memory),
			CPULimit:    parseCPULimit(*cpus),
		}
		
		// Create container
		container, err := rtr.CreateContainer(*name, *image, *command, limits)
		if err != nil {
			fmt.Printf("❌ Failed to create container: %v\n", err)
			os.Exit(1)
		}
		
		// Add port mappings
		if *ports != "" {
			portMappings := parsePortMappings(*ports)
			for _, mapping := range portMappings {
				if err := rtr.networkManager.AddPortMapping(container.Network, mapping.HostPort, mapping.ContainerPort, mapping.Protocol); err != nil {
					fmt.Printf("Warning: failed to add port mapping: %v\n", err)
				}
			}
		}
		
		fmt.Printf("✅ Container created successfully: %s\n", container.ID)
		
	case "start":
		if *containerID == "" {
			fmt.Println("❌ Container ID required")
			os.Exit(1)
		}
		
		if err := rtr.StartContainer(*containerID); err != nil {
			fmt.Printf("❌ Failed to start container: %v\n", err)
			os.Exit(1)
		}
		
	case "stop":
		if *containerID == "" {
			fmt.Println("❌ Container ID required")
			os.Exit(1)
		}
		
		if err := rtr.StopContainer(*containerID); err != nil {
			fmt.Printf("❌ Failed to stop container: %v\n", err)
			os.Exit(1)
		}
		
	case "remove":
		if *containerID == "" {
			fmt.Println("❌ Container ID required")
			os.Exit(1)
		}
		
		if err := rtr.RemoveContainer(*containerID); err != nil {
			fmt.Printf("❌ Failed to remove container: %v\n", err)
			os.Exit(1)
		}
		
	case "list":
		rtr.ListContainers()
		
	case "info":
		if *containerID == "" {
			fmt.Println("❌ Container ID required")
			os.Exit(1)
		}
		
		info, err := rtr.GetContainerInfo(*containerID)
		if err != nil {
			fmt.Printf("❌ Failed to get container info: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("Container: %s (%s)\n", info.Container.Name, info.Container.ID)
		fmt.Printf("Status: %s\n", info.Container.Status)
		fmt.Printf("Image: %s\n", info.Container.Image)
		fmt.Printf("Command: %s\n", info.Container.Command)
		fmt.Printf("Created: %s\n", info.Container.Created.Format("2006-01-02 15:04:05"))
		
		if info.ResourceUsage != nil {
			fmt.Printf("Memory Usage: %s\n", formatBytes(info.ResourceUsage.MemoryUsage))
			fmt.Printf("CPU Usage: %.2fs\n", info.ResourceUsage.CPUUsage)
		}
		
		if info.NetworkInfo != nil {
			fmt.Printf("IP Address: %s\n", info.NetworkInfo.ContainerIP)
			fmt.Printf("Gateway: %s\n", info.NetworkInfo.GatewayIP)
		}
		
	default:
		fmt.Println("🐳 Real Container Runtime")
		fmt.Println("=========================")
		fmt.Println("A real container runtime with namespace isolation, cgroups, and networking")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  sudo go run main.go -action=create -name=my-container -image=ubuntu:latest")
		fmt.Println("  sudo go run main.go -action=start -id=container-id")
		fmt.Println("  sudo go run main.go -action=stop -id=container-id")
		fmt.Println("  sudo go run main.go -action=remove -id=container-id")
		fmt.Println("  sudo go run main.go -action=list")
		fmt.Println("  sudo go run main.go -action=info -id=container-id")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -name     Container name")
		fmt.Println("  -image    Container image (default: ubuntu:latest)")
		fmt.Println("  -command  Command to run (default: sh)")
		fmt.Println("  -memory   Memory limit (e.g., 512m, 1g)")
		fmt.Println("  -cpus     CPU limit (e.g., 1.0, 2.5)")
		fmt.Println("  -ports    Port mappings (e.g., 8080:80,9090:9090)")
		fmt.Println("  -id       Container ID")
	}
}

// parseMemoryLimit parses memory limit string
func parseMemoryLimit(memory string) int64 {
	if memory == "" {
		return 0
	}
	
	// Simple parsing - in real implementation, this would be more sophisticated
	switch memory[len(memory)-1] {
	case 'm', 'M':
		return int64(parseInt(memory[:len(memory)-1])) * 1024 * 1024
	case 'g', 'G':
		return int64(parseInt(memory[:len(memory)-1])) * 1024 * 1024 * 1024
	default:
		return int64(parseInt(memory)) * 1024 * 1024
	}
}

// parseCPULimit parses CPU limit string
func parseCPULimit(cpus string) int64 {
	if cpus == "" {
		return 0
	}
	
	// Simple parsing - in real implementation, this would be more sophisticated
	return int64(parseFloat(cpus) * 100)
}

// parsePortMappings parses port mapping string
func parsePortMappings(ports string) []PortMapping {
	mappings := make([]PortMapping, 0)
	
	if ports == "" {
		return mappings
	}
	
	portList := strings.Split(ports, ",")
	for _, port := range portList {
		parts := strings.Split(port, ":")
		if len(parts) == 2 {
			mapping := PortMapping{
				HostPort:      parts[0],
				ContainerPort: parts[1],
				Protocol:      "tcp",
				HostIP:        "0.0.0.0",
			}
			mappings = append(mappings, mapping)
		}
	}
	
	return mappings
}

// parseInt parses integer string
func parseInt(s string) int {
	if s == "" {
		return 0
	}
	
	result := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		}
	}
	
	return result
}

// parseFloat parses float string
func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	
	// Simple parsing - in real implementation, this would be more sophisticated
	return float64(parseInt(s))
}
