package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// ContainerRuntime manages container operations
type ContainerRuntime struct {
	containers map[string]*Container
	mutex      sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
}

// Container represents a container instance
type Container struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Command     string            `json:"command"`
	Status      string            `json:"status"`
	Created     time.Time         `json:"created"`
	Started     time.Time         `json:"started"`
	Stopped     time.Time         `json:"stopped"`
	Environment map[string]string `json:"environment"`
	Ports       []PortMapping     `json:"ports"`
	Volumes     []VolumeMount     `json:"volumes"`
	Network     string            `json:"network"`
	Process     *exec.Cmd         `json:"-"`
	Logs        []string          `json:"logs"`
	Stats       *ContainerStats   `json:"stats"`
}

// PortMapping represents port mapping configuration
type PortMapping struct {
	HostPort      string `json:"host_port"`
	ContainerPort string `json:"container_port"`
	Protocol      string `json:"protocol"`
}

// VolumeMount represents volume mount configuration
type VolumeMount struct {
	HostPath      string `json:"host_path"`
	ContainerPath string `json:"container_path"`
	Mode          string `json:"mode"`
}

// ContainerStats represents container performance statistics
type ContainerStats struct {
	CPUUsage    float64 `json:"cpu_usage_percent"`
	MemoryUsage uint64  `json:"memory_usage_bytes"`
	MemoryLimit uint64  `json:"memory_limit_bytes"`
	NetworkRx   uint64  `json:"network_rx_bytes"`
	NetworkTx   uint64  `json:"network_tx_bytes"`
	BlockRead   uint64  `json:"block_read_bytes"`
	BlockWrite  uint64  `json:"block_write_bytes"`
	PIDs        int     `json:"pids"`
	Timestamp   time.Time `json:"timestamp"`
}

// CreateOptions contains options for container creation
type CreateOptions struct {
	Name    string
	Image   string
	Command string
	Env     string
	Ports   string
	Volumes string
	Network string
}

// NewContainerRuntime creates a new container runtime
func NewContainerRuntime() *ContainerRuntime {
	ctx, cancel := context.WithCancel(context.Background())
	return &ContainerRuntime{
		containers: make(map[string]*Container),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Close closes the container runtime
func (cr *ContainerRuntime) Close() {
	cr.cancel()
	
	// Stop all running containers
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	for _, container := range cr.containers {
		if container.Status == "running" {
			cr.stopContainerInternal(container)
		}
	}
}

// CreateContainer creates a new container
func (cr *ContainerRuntime) CreateContainer(opts *CreateOptions) (*Container, error) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	// Generate container ID
	containerID := cr.generateContainerID()
	
	// Parse environment variables
	env := make(map[string]string)
	if opts.Env != "" {
		envVars := strings.Split(opts.Env, ",")
		for _, envVar := range envVars {
			parts := strings.SplitN(envVar, "=", 2)
			if len(parts) == 2 {
				env[parts[0]] = parts[1]
			}
		}
	}
	
	// Parse port mappings
	ports := cr.parsePortMappings(opts.Ports)
	
	// Parse volume mounts
	volumes := cr.parseVolumeMounts(opts.Volumes)
	
	// Create container
	container := &Container{
		ID:          containerID,
		Name:        opts.Name,
		Image:       opts.Image,
		Command:     opts.Command,
		Status:      "created",
		Created:     time.Now(),
		Environment: env,
		Ports:       ports,
		Volumes:     volumes,
		Network:     opts.Network,
		Logs:        make([]string, 0),
		Stats:       &ContainerStats{},
	}
	
	cr.containers[opts.Name] = container
	
	return container, nil
}

// StartContainer starts a container
func (cr *ContainerRuntime) StartContainer(name string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status == "running" {
		return fmt.Errorf("container is already running: %s", name)
	}
	
	return cr.startContainerInternal(container)
}

// startContainerInternal starts a container internally
func (cr *ContainerRuntime) startContainerInternal(container *Container) error {
	// Create command
	var cmd *exec.Cmd
	if container.Command != "" {
		cmd = exec.CommandContext(cr.ctx, "sh", "-c", container.Command)
	} else {
		cmd = exec.CommandContext(cr.ctx, "sh", "-c", "echo 'Container started' && sleep 3600")
	}
	
	// Set environment variables
	env := os.Environ()
	for key, value := range container.Environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env
	
	// Set up logging
	cmd.Stdout = &ContainerWriter{container: container}
	cmd.Stderr = &ContainerWriter{container: container}
	
	// Start the process
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}
	
	container.Process = cmd
	container.Status = "running"
	container.Started = time.Now()
	
	// Start monitoring goroutine
	go cr.monitorContainer(container)
	
	return nil
}

// StopContainer stops a container
func (cr *ContainerRuntime) StopContainer(name string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status != "running" {
		return fmt.Errorf("container is not running: %s", name)
	}
	
	return cr.stopContainerInternal(container)
}

// stopContainerInternal stops a container internally
func (cr *ContainerRuntime) stopContainerInternal(container *Container) error {
	if container.Process != nil {
		container.Process.Process.Kill()
		container.Process.Wait()
	}
	
	container.Status = "stopped"
	container.Stopped = time.Now()
	
	return nil
}

// RemoveContainer removes a container
func (cr *ContainerRuntime) RemoveContainer(name string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status == "running" {
		return fmt.Errorf("cannot remove running container: %s", name)
	}
	
	delete(cr.containers, name)
	return nil
}

// ListContainers lists all containers
func (cr *ContainerRuntime) ListContainers() ([]*Container, error) {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()
	
	containers := make([]*Container, 0, len(cr.containers))
	for _, container := range cr.containers {
		containers = append(containers, container)
	}
	
	return containers, nil
}

// ExecContainer executes a command in a container
func (cr *ContainerRuntime) ExecContainer(name, command string) (string, error) {
	cr.mutex.RLock()
	container, exists := cr.containers[name]
	cr.mutex.RUnlock()
	
	if !exists {
		return "", fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status != "running" {
		return "", fmt.Errorf("container is not running: %s", name)
	}
	
	// Simulate command execution
	output := fmt.Sprintf("Executing command '%s' in container %s\n", command, name)
	output += fmt.Sprintf("Command output: %s\n", command)
	output += "Command completed successfully\n"
	
	// Add to logs
	cr.mutex.Lock()
	container.Logs = append(container.Logs, fmt.Sprintf("[%s] EXEC: %s", time.Now().Format("2006-01-02 15:04:05"), command))
	cr.mutex.Unlock()
	
	return output, nil
}

// GetContainerLogs gets container logs
func (cr *ContainerRuntime) GetContainerLogs(name string) ([]string, error) {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return nil, fmt.Errorf("container not found: %s", name)
	}
	
	return container.Logs, nil
}

// GetContainerStats gets container statistics
func (cr *ContainerRuntime) GetContainerStats(name string) (*ContainerStats, error) {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return nil, fmt.Errorf("container not found: %s", name)
	}
	
	return container.Stats, nil
}

// MonitorContainerStats monitors container statistics in real-time
func (cr *ContainerRuntime) MonitorContainerStats(name string, interval time.Duration) {
	fmt.Printf("Monitoring container stats: %s (interval: %v)\n", name, interval)
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()
	
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			stats, err := cr.GetContainerStats(name)
			if err != nil {
				fmt.Printf("Error getting stats: %v\n", err)
				continue
			}
			
			cr.printStatsRealtime(stats)
			
		case <-cr.ctx.Done():
			return
		}
	}
}

// PauseContainer pauses a container
func (cr *ContainerRuntime) PauseContainer(name string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status != "running" {
		return fmt.Errorf("container is not running: %s", name)
	}
	
	container.Status = "paused"
	return nil
}

// ResumeContainer resumes a container
func (cr *ContainerRuntime) ResumeContainer(name string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	if container.Status != "paused" {
		return fmt.Errorf("container is not paused: %s", name)
	}
	
	container.Status = "running"
	return nil
}

// DeployContainers deploys containers from a configuration file
func (cr *ContainerRuntime) DeployContainers(file string) error {
	// Simulate deployment from file
	fmt.Printf("Deploying containers from file: %s\n", file)
	
	// Create sample containers
	containers := []*CreateOptions{
		{Name: "web-server", Image: "nginx:latest", Command: "nginx -g 'daemon off;'", Ports: "80:80"},
		{Name: "database", Image: "postgres:latest", Command: "postgres", Ports: "5432:5432"},
		{Name: "cache", Image: "redis:latest", Command: "redis-server", Ports: "6379:6379"},
	}
	
	for _, opts := range containers {
		_, err := cr.CreateContainer(opts)
		if err != nil {
			return fmt.Errorf("failed to create container %s: %w", opts.Name, err)
		}
		
		err = cr.StartContainer(opts.Name)
		if err != nil {
			return fmt.Errorf("failed to start container %s: %w", opts.Name, err)
		}
	}
	
	return nil
}

// ScaleContainer scales a container to the specified number of replicas
func (cr *ContainerRuntime) ScaleContainer(name string, replicas int) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	// Simulate scaling
	fmt.Printf("Scaling container %s to %d replicas\n", name, replicas)
	
	// Update container stats to reflect scaling
	container.Stats.PIDs = replicas
	
	return nil
}

// CheckContainerHealth checks container health
func (cr *ContainerRuntime) CheckContainerHealth(name string) (*HealthStatus, error) {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return nil, fmt.Errorf("container not found: %s", name)
	}
	
	health := &HealthStatus{
		Container: name,
		Status:    "healthy",
		Timestamp: time.Now(),
	}
	
	if container.Status == "running" {
		health.Status = "healthy"
	} else if container.Status == "stopped" {
		health.Status = "unhealthy"
	} else if container.Status == "paused" {
		health.Status = "paused"
	}
	
	return health, nil
}

// RestartContainer restarts a container
func (cr *ContainerRuntime) RestartContainer(name, policy string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	container, exists := cr.containers[name]
	if !exists {
		return fmt.Errorf("container not found: %s", name)
	}
	
	// Stop container if running
	if container.Status == "running" {
		cr.stopContainerInternal(container)
	}
	
	// Start container
	err := cr.startContainerInternal(container)
	if err != nil {
		return fmt.Errorf("failed to restart container: %w", err)
	}
	
	return nil
}

// Helper functions

func (cr *ContainerRuntime) generateContainerID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (cr *ContainerRuntime) parsePortMappings(ports string) []PortMapping {
	if ports == "" {
		return nil
	}
	
	mappings := make([]PortMapping, 0)
	portList := strings.Split(ports, ",")
	
	for _, port := range portList {
		parts := strings.Split(port, ":")
		if len(parts) == 2 {
			mappings = append(mappings, PortMapping{
				HostPort:      parts[0],
				ContainerPort: parts[1],
				Protocol:      "tcp",
			})
		}
	}
	
	return mappings
}

func (cr *ContainerRuntime) parseVolumeMounts(volumes string) []VolumeMount {
	if volumes == "" {
		return nil
	}
	
	mounts := make([]VolumeMount, 0)
	volumeList := strings.Split(volumes, ",")
	
	for _, volume := range volumeList {
		parts := strings.Split(volume, ":")
		if len(parts) >= 2 {
			mode := "rw"
			if len(parts) == 3 {
				mode = parts[2]
			}
			
			mounts = append(mounts, VolumeMount{
				HostPath:      parts[0],
				ContainerPath: parts[1],
				Mode:          mode,
			})
		}
	}
	
	return mounts
}

func (cr *ContainerRuntime) monitorContainer(container *Container) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			cr.updateContainerStats(container)
		case <-cr.ctx.Done():
			return
		}
	}
}

func (cr *ContainerRuntime) updateContainerStats(container *Container) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	
	// Simulate stats update
	container.Stats = &ContainerStats{
		CPUUsage:    float64(time.Now().Unix() % 100),
		MemoryUsage: uint64(1024 * 1024 * (1 + time.Now().Unix()%10)),
		MemoryLimit: uint64(1024 * 1024 * 100),
		NetworkRx:   uint64(time.Now().Unix() % 1000000),
		NetworkTx:   uint64(time.Now().Unix() % 1000000),
		BlockRead:   uint64(time.Now().Unix() % 100000),
		BlockWrite:  uint64(time.Now().Unix() % 100000),
		PIDs:        1,
		Timestamp:   time.Now(),
	}
}

func (cr *ContainerRuntime) printStatsRealtime(stats *ContainerStats) {
	fmt.Printf("\rCPU: %.2f%% | Memory: %s/%s | Network: %s↓ %s↑ | PIDs: %d",
		stats.CPUUsage,
		formatBytes(stats.MemoryUsage),
		formatBytes(stats.MemoryLimit),
		formatBytes(stats.NetworkRx),
		formatBytes(stats.NetworkTx),
		stats.PIDs)
}

// ContainerWriter implements io.Writer for container logs
type ContainerWriter struct {
	container *Container
}

func (cw *ContainerWriter) Write(p []byte) (n int, err error) {
	cw.container.Logs = append(cw.container.Logs, string(p))
	return len(p), nil
}

// HealthStatus represents container health status
type HealthStatus struct {
	Container string    `json:"container"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// Print functions

func (cr *ContainerRuntime) PrintContainers(containers []*Container) {
	fmt.Println("Container List")
	fmt.Println("==============")
	fmt.Printf("%-20s %-20s %-15s %-10s %-20s\n",
		"Name", "Image", "Status", "Ports", "Created")
	fmt.Println(strings.Repeat("-", 90))
	
	for _, container := range containers {
		ports := ""
		if len(container.Ports) > 0 {
			ports = fmt.Sprintf("%s:%s", container.Ports[0].HostPort, container.Ports[0].ContainerPort)
		}
		
		fmt.Printf("%-20s %-20s %-15s %-10s %-20s\n",
			container.Name,
			container.Image,
			container.Status,
			ports,
			container.Created.Format("2006-01-02 15:04:05"))
	}
}

func (cr *ContainerRuntime) PrintLogs(logs []string) {
	fmt.Println("Container Logs")
	fmt.Println("==============")
	for _, log := range logs {
		fmt.Print(log)
	}
}

func (cr *ContainerRuntime) PrintStats(stats *ContainerStats) {
	fmt.Println("Container Statistics")
	fmt.Println("===================")
	fmt.Printf("CPU Usage: %.2f%%\n", stats.CPUUsage)
	fmt.Printf("Memory Usage: %s / %s\n", formatBytes(stats.MemoryUsage), formatBytes(stats.MemoryLimit))
	fmt.Printf("Network: %s↓ %s↑\n", formatBytes(stats.NetworkRx), formatBytes(stats.NetworkTx))
	fmt.Printf("Block I/O: %s↓ %s↑\n", formatBytes(stats.BlockRead), formatBytes(stats.BlockWrite))
	fmt.Printf("PIDs: %d\n", stats.PIDs)
	fmt.Printf("Timestamp: %s\n", stats.Timestamp.Format("2006-01-02 15:04:05"))
}

func (cr *ContainerRuntime) PrintJSON(data interface{}) {
	fmt.Printf("JSON output: %+v\n", data)
}

func (cr *ContainerRuntime) PrintCSV(data interface{}) {
	fmt.Printf("CSV output: %+v\n", data)
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
