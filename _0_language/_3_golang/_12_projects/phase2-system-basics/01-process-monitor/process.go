package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// Process represents a system process
type Process struct {
	PID         int32     `json:"pid"`
	PPID        int32     `json:"ppid"`
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	User        string    `json:"user"`
	Status      string    `json:"status"`
	CPUPercent  float64   `json:"cpu_percent"`
	MemoryMB    float64   `json:"memory_mb"`
	CreateTime  time.Time `json:"create_time"`
	StartTime   time.Time `json:"start_time"`
	NumThreads  int32     `json:"num_threads"`
	NumFDs      int32     `json:"num_fds"`
	WorkingDir  string    `json:"working_dir"`
	Executable  string    `json:"executable"`
	Args        []string  `json:"args"`
	Environment map[string]string `json:"environment"`
}

// ProcessManager manages system processes
type ProcessManager struct {
	processes map[int32]*Process
	mutex     sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewProcessManager creates a new process manager
func NewProcessManager() *ProcessManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &ProcessManager{
		processes: make(map[int32]*Process),
		ctx:       ctx,
		cancel:    cancel,
	}
}

// Close closes the process manager
func (pm *ProcessManager) Close() {
	pm.cancel()
}

// ListOptions contains options for listing processes
type ListOptions struct {
	Name     string
	User     string
	Detailed bool
	Output   string
	Limit    int
}

// MonitorOptions contains options for monitoring processes
type MonitorOptions struct {
	PID      int
	Name     string
	Watch    bool
	Interval time.Duration
	Output   string
}

// StartOptions contains options for starting a process
type StartOptions struct {
	Command string
	Args    string
}

// TreeOptions contains options for process tree
type TreeOptions struct {
	PID    int
	Output string
}

// StatsOptions contains options for system statistics
type StatsOptions struct {
	Output string
}

// ListProcesses lists all processes with given options
func (pm *ProcessManager) ListProcesses(opts *ListOptions) ([]*Process, error) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	
	// Get all processes
	processes, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("failed to get processes: %w", err)
	}
	
	var result []*Process
	count := 0
	
	for _, p := range processes {
		if opts.Limit > 0 && count >= opts.Limit {
			break
		}
		
		proc, err := pm.getProcessInfo(p)
		if err != nil {
			continue // Skip processes we can't access
		}
		
		// Apply filters
		if opts.Name != "" && !strings.Contains(strings.ToLower(proc.Name), strings.ToLower(opts.Name)) {
			continue
		}
		
		if opts.User != "" && proc.User != opts.User {
			continue
		}
		
		result = append(result, proc)
		count++
	}
	
	return result, nil
}

// GetProcesses gets processes for monitoring
func (pm *ProcessManager) GetProcesses(opts *MonitorOptions) ([]*Process, error) {
	if opts.PID > 0 {
		// Get specific process
		p, err := process.NewProcess(int32(opts.PID))
		if err != nil {
			return nil, fmt.Errorf("failed to get process %d: %w", opts.PID, err)
		}
		
		proc, err := pm.getProcessInfo(p)
		if err != nil {
			return nil, err
		}
		
		return []*Process{proc}, nil
	}
	
	// Get all processes with name filter
	return pm.ListProcesses(&ListOptions{
		Name:   opts.Name,
		Limit:  100,
	})
}

// getProcessInfo gets detailed information about a process
func (pm *ProcessManager) getProcessInfo(p *process.Process) (*Process, error) {
	// Basic info
	pid := p.Pid
	name, _ := p.Name()
	status, _ := p.Status()
	createTime, _ := p.CreateTime()
	startTime, _ := p.CreateTime()
	numThreads, _ := p.NumThreads()
	numFDs, _ := p.NumFDs()
	
	// CPU and memory
	cpuPercent, _ := p.CPUPercent()
	memInfo, _ := p.MemoryInfo()
	memoryMB := float64(0)
	if memInfo != nil {
		memoryMB = float64(memInfo.RSS) / 1024 / 1024
	}
	
	// Command and args
	cmdline, _ := p.Cmdline()
	exe, _ := p.Exe()
	cwd, _ := p.Cwd()
	
	// Parse args from cmdline
	args := strings.Fields(cmdline)
	if len(args) > 0 {
		args = args[1:] // Remove command name
	}
	
	// Parent process
	ppid, _ := p.Ppid()
	
	// User
	username := "unknown"
	if u, err := p.Username(); err == nil {
		username = u
	}
	
	// Environment
	env, _ := p.Environ()
	environment := make(map[string]string)
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			environment[parts[0]] = parts[1]
		}
	}
	
	return &Process{
		PID:         pid,
		PPID:        ppid,
		Name:        name,
		Command:     cmdline,
		User:        username,
		Status:      strings.Join(status, ","),
		CPUPercent:  cpuPercent,
		MemoryMB:    memoryMB,
		CreateTime:  time.Unix(createTime/1000, 0),
		StartTime:   time.Unix(startTime/1000, 0),
		NumThreads:  numThreads,
		NumFDs:      numFDs,
		WorkingDir:  cwd,
		Executable:  exe,
		Args:        args,
		Environment: environment,
	}, nil
}

// StartProcess starts a new process
func (pm *ProcessManager) StartProcess(opts *StartOptions) (int, error) {
	// Parse command and arguments
	cmd := opts.Command
	args := []string{}
	
	if opts.Args != "" {
		args = strings.Fields(opts.Args)
	}
	
	// Create command
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	
	// Start process
	err := command.Start()
	if err != nil {
		return 0, fmt.Errorf("failed to start process: %w", err)
	}
	
	// Store process info
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	
	proc := &Process{
		PID:        int32(command.Process.Pid),
		Name:       cmd,
		Command:    cmd + " " + strings.Join(args, " "),
		Status:     "running",
		CreateTime: time.Now(),
		StartTime:  time.Now(),
	}
	
	pm.processes[proc.PID] = proc
	
	return command.Process.Pid, nil
}

// StopProcess stops a process gracefully
func (pm *ProcessManager) StopProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process %d: %w", pid, err)
	}
	
	// Send SIGTERM for graceful shutdown
	err = proc.Signal(syscall.SIGTERM)
	if err != nil {
		return fmt.Errorf("failed to send SIGTERM to process %d: %w", pid, err)
	}
	
	// Wait for process to exit
	done := make(chan error, 1)
	go func() {
		_, err := proc.Wait()
		done <- err
	}()
	
	// Wait with timeout
	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("process %d did not exit gracefully: %w", pid, err)
		}
	case <-time.After(10 * time.Second):
		// Force kill if it doesn't exit
		proc.Signal(syscall.SIGKILL)
		return fmt.Errorf("process %d did not exit within timeout", pid)
	}
	
	// Remove from our tracking
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	delete(pm.processes, int32(pid))
	
	return nil
}

// KillProcess kills a process immediately
func (pm *ProcessManager) KillProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process %d: %w", pid, err)
	}
	
	// Send SIGKILL for immediate termination
	err = proc.Signal(syscall.SIGKILL)
	if err != nil {
		return fmt.Errorf("failed to send SIGKILL to process %d: %w", pid, err)
	}
	
	// Remove from our tracking
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	delete(pm.processes, int32(pid))
	
	return nil
}

// GetProcessTree gets the process tree
func (pm *ProcessManager) GetProcessTree(opts *TreeOptions) (*ProcessNode, error) {
	// Get all processes
	processes, err := pm.ListProcesses(&ListOptions{Limit: 1000})
	if err != nil {
		return nil, err
	}
	
	// Build process map
	processMap := make(map[int32]*Process)
	for _, proc := range processes {
		processMap[proc.PID] = proc
	}
	
	// Find root process
	var root *Process
	if opts.PID > 0 {
		root = processMap[int32(opts.PID)]
	} else {
		// Find init process (PID 1)
		root = processMap[1]
	}
	
	if root == nil {
		return nil, fmt.Errorf("root process not found")
	}
	
	// Build tree
	return pm.buildProcessTree(root, processMap), nil
}

// ProcessNode represents a node in the process tree
type ProcessNode struct {
	Process  *Process       `json:"process"`
	Children []*ProcessNode `json:"children"`
}

// buildProcessTree builds the process tree recursively
func (pm *ProcessManager) buildProcessTree(proc *Process, processMap map[int32]*Process) *ProcessNode {
	node := &ProcessNode{
		Process:  proc,
		Children: []*ProcessNode{},
	}
	
	// Find children
	for _, p := range processMap {
		if p.PPID == proc.PID {
			child := pm.buildProcessTree(p, processMap)
			node.Children = append(node.Children, child)
		}
	}
	
	return node
}

// GetSystemStats gets system statistics
func (pm *ProcessManager) GetSystemStats() (*SystemStats, error) {
	// Get system info
	hostInfo, err := getHostInfo()
	if err != nil {
		return nil, err
	}
	
	// Get memory info
	memInfo, err := getMemoryInfo()
	if err != nil {
		return nil, err
	}
	
	// Get CPU info
	cpuInfo, err := getCPUInfo()
	if err != nil {
		return nil, err
	}
	
	// Get process count
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	
	return &SystemStats{
		Hostname:     hostInfo.Hostname,
		OS:           hostInfo.OS,
		Architecture: hostInfo.Architecture,
		Kernel:       hostInfo.KernelVersion,
		Uptime:       hostInfo.Uptime,
		Processes:    len(processes),
		Memory:       memInfo,
		CPU:          cpuInfo,
		Timestamp:    time.Now(),
	}, nil
}

// SystemStats represents system statistics
type SystemStats struct {
	Hostname     string      `json:"hostname"`
	OS           string      `json:"os"`
	Architecture string      `json:"architecture"`
	Kernel       string      `json:"kernel"`
	Uptime       uint64      `json:"uptime"`
	Processes    int         `json:"processes"`
	Memory       *MemoryInfo `json:"memory"`
	CPU          *CPUInfo    `json:"cpu"`
	Timestamp    time.Time   `json:"timestamp"`
}

// MemoryInfo represents memory information
type MemoryInfo struct {
	Total     uint64  `json:"total"`
	Available uint64  `json:"available"`
	Used      uint64  `json:"used"`
	Free      uint64  `json:"free"`
	Percent   float64 `json:"percent"`
}

// CPUInfo represents CPU information
type CPUInfo struct {
	Count     int     `json:"count"`
	Usage     float64 `json:"usage"`
	Model     string  `json:"model"`
	Frequency uint64  `json:"frequency"`
}

// Helper functions for system information
func getHostInfo() (*struct {
	Hostname      string
	OS            string
	Architecture  string
	KernelVersion string
	Uptime        uint64
}, error) {
	// This is a simplified version - in a real implementation,
	// you would use gopsutil or similar libraries
	return &struct {
		Hostname      string
		OS            string
		Architecture  string
		KernelVersion string
		Uptime        uint64
	}{
		Hostname:      "localhost",
		OS:            "darwin",
		Architecture:  "arm64",
		KernelVersion: "24.5.0",
		Uptime:        12345,
	}, nil
}

func getMemoryInfo() (*MemoryInfo, error) {
	// This is a simplified version - in a real implementation,
	// you would use gopsutil or similar libraries
	return &MemoryInfo{
		Total:     8589934592, // 8GB
		Available: 4294967296, // 4GB
		Used:      4294967296, // 4GB
		Free:      4294967296, // 4GB
		Percent:   50.0,
	}, nil
}

func getCPUInfo() (*CPUInfo, error) {
	// This is a simplified version - in a real implementation,
	// you would use gopsutil or similar libraries
	return &CPUInfo{
		Count:     8,
		Usage:     25.5,
		Model:     "Apple M2",
		Frequency: 3200000000, // 3.2GHz
	}, nil
}
