package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CGroupManager manages cgroups for container resource limits
type CGroupManager struct {
	cgroupPath string
	containerID string
}

// ResourceLimits represents container resource limits
type ResourceLimits struct {
	MemoryLimit    int64  // Memory limit in bytes
	CPULimit       int64  // CPU limit in percentage (0-100)
	CPUQuota       int64  // CPU quota in microseconds
	CPUPeriod      int64  // CPU period in microseconds
	PIDsLimit      int64  // Maximum number of processes
	BlockIOLimit   int64  // Block I/O limit
	NetworkLimit   int64  // Network bandwidth limit
}

// NewCGroupManager creates a new cgroup manager
func NewCGroupManager(containerID string) *CGroupManager {
	return &CGroupManager{
		cgroupPath:  "/sys/fs/cgroup",
		containerID: containerID,
	}
}

// CreateCGroup creates a cgroup for the container
func (cgm *CGroupManager) CreateCGroup() error {
	fmt.Printf("⚡ Creating cgroup for container: %s\n", cgm.containerID)
	
	// Check if cgroups v2 is available
	if cgm.isCGroupV2Available() {
		return cgm.createCGroupV2()
	}
	
	// Fall back to cgroups v1
	return cgm.createCGroupV1()
}

// createCGroupV2 creates a cgroup using cgroups v2
func (cgm *CGroupManager) createCGroupV2() error {
	fmt.Println("  Using cgroups v2")
	
	// Create container cgroup directory
	containerCGroupPath := filepath.Join(cgm.cgroupPath, cgm.containerID)
	if err := os.MkdirAll(containerCGroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup directory: %w", err)
	}
	
	fmt.Printf("  CGroup created at: %s\n", containerCGroupPath)
	return nil
}

// createCGroupV1 creates a cgroup using cgroups v1
func (cgm *CGroupManager) createCGroupV1() error {
	fmt.Println("  Using cgroups v1")
	
	// Create cgroup directories for different controllers
	controllers := []string{"memory", "cpu", "pids", "blkio"}
	
	for _, controller := range controllers {
		cgroupPath := filepath.Join(cgm.cgroupPath, controller, cgm.containerID)
		if err := os.MkdirAll(cgroupPath, 0755); err != nil {
			return fmt.Errorf("failed to create %s cgroup: %w", controller, err)
		}
	}
	
	return nil
}

// SetResourceLimits sets resource limits for the container
func (cgm *CGroupManager) SetResourceLimits(limits *ResourceLimits) error {
	fmt.Println("⚡ Setting container resource limits...")
	
	if cgm.isCGroupV2Available() {
		return cgm.setCGroupV2Limits(limits)
	}
	
	return cgm.setCGroupV1Limits(limits)
}

// setCGroupV2Limits sets limits using cgroups v2
func (cgm *CGroupManager) setCGroupV2Limits(limits *ResourceLimits) error {
	containerCGroupPath := filepath.Join(cgm.cgroupPath, cgm.containerID)
	
	// Set memory limit
	if limits.MemoryLimit > 0 {
		memoryFile := filepath.Join(containerCGroupPath, "memory.max")
		if err := os.WriteFile(memoryFile, []byte(fmt.Sprintf("%d", limits.MemoryLimit)), 0644); err != nil {
			return fmt.Errorf("failed to set memory limit: %w", err)
		}
		fmt.Printf("  Memory limit set to: %s\n", formatBytes(uint64(limits.MemoryLimit)))
	}
	
	// Set CPU limit
	if limits.CPULimit > 0 {
		cpuFile := filepath.Join(containerCGroupPath, "cpu.max")
		cpuLimit := fmt.Sprintf("%d %d", limits.CPUQuota, limits.CPUPeriod)
		if err := os.WriteFile(cpuFile, []byte(cpuLimit), 0644); err != nil {
			return fmt.Errorf("failed to set CPU limit: %w", err)
		}
		fmt.Printf("  CPU limit set to: %d%%\n", limits.CPULimit)
	}
	
	// Set PIDs limit
	if limits.PIDsLimit > 0 {
		pidsFile := filepath.Join(containerCGroupPath, "pids.max")
		if err := os.WriteFile(pidsFile, []byte(fmt.Sprintf("%d", limits.PIDsLimit)), 0644); err != nil {
			return fmt.Errorf("failed to set PIDs limit: %w", err)
		}
		fmt.Printf("  PIDs limit set to: %d\n", limits.PIDsLimit)
	}
	
	return nil
}

// setCGroupV1Limits sets limits using cgroups v1
func (cgm *CGroupManager) setCGroupV1Limits(limits *ResourceLimits) error {
	// Set memory limit
	if limits.MemoryLimit > 0 {
		memoryFile := filepath.Join(cgm.cgroupPath, "memory", cgm.containerID, "memory.limit_in_bytes")
		if err := os.WriteFile(memoryFile, []byte(fmt.Sprintf("%d", limits.MemoryLimit)), 0644); err != nil {
			return fmt.Errorf("failed to set memory limit: %w", err)
		}
		fmt.Printf("  Memory limit set to: %s\n", formatBytes(uint64(limits.MemoryLimit)))
	}
	
	// Set CPU limit
	if limits.CPULimit > 0 {
		cpuQuotaFile := filepath.Join(cgm.cgroupPath, "cpu", cgm.containerID, "cpu.cfs_quota_us")
		cpuPeriodFile := filepath.Join(cgm.cgroupPath, "cpu", cgm.containerID, "cpu.cfs_period_us")
		
		// Set CPU period (default 100ms)
		if limits.CPUPeriod == 0 {
			limits.CPUPeriod = 100000 // 100ms in microseconds
		}
		
		if err := os.WriteFile(cpuPeriodFile, []byte(fmt.Sprintf("%d", limits.CPUPeriod)), 0644); err != nil {
			return fmt.Errorf("failed to set CPU period: %w", err)
		}
		
		// Set CPU quota
		if limits.CPUQuota == 0 {
			limits.CPUQuota = (limits.CPULimit * limits.CPUPeriod) / 100
		}
		
		if err := os.WriteFile(cpuQuotaFile, []byte(fmt.Sprintf("%d", limits.CPUQuota)), 0644); err != nil {
			return fmt.Errorf("failed to set CPU quota: %w", err)
		}
		
		fmt.Printf("  CPU limit set to: %d%%\n", limits.CPULimit)
	}
	
	// Set PIDs limit
	if limits.PIDsLimit > 0 {
		pidsFile := filepath.Join(cgm.cgroupPath, "pids", cgm.containerID, "pids.max")
		if err := os.WriteFile(pidsFile, []byte(fmt.Sprintf("%d", limits.PIDsLimit)), 0644); err != nil {
			return fmt.Errorf("failed to set PIDs limit: %w", err)
		}
		fmt.Printf("  PIDs limit set to: %d\n", limits.PIDsLimit)
	}
	
	return nil
}

// AddProcess adds a process to the cgroup
func (cgm *CGroupManager) AddProcess(pid int) error {
	fmt.Printf("⚡ Adding process %d to cgroup\n", pid)
	
	if cgm.isCGroupV2Available() {
		return cgm.addProcessToCGroupV2(pid)
	}
	
	return cgm.addProcessToCGroupV1(pid)
}

// addProcessToCGroupV2 adds process to cgroup v2
func (cgm *CGroupManager) addProcessToCGroupV2(pid int) error {
	cgroupProcsFile := filepath.Join(cgm.cgroupPath, cgm.containerID, "cgroup.procs")
	
	// Add process to cgroup
	if err := os.WriteFile(cgroupProcsFile, []byte(fmt.Sprintf("%d", pid)), 0644); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %w", err)
	}
	
	return nil
}

// addProcessToCGroupV1 adds process to cgroup v1
func (cgm *CGroupManager) addProcessToCGroupV1(pid int) error {
	controllers := []string{"memory", "cpu", "pids", "blkio"}
	
	for _, controller := range controllers {
		cgroupProcsFile := filepath.Join(cgm.cgroupPath, controller, cgm.containerID, "cgroup.procs")
		
		// Add process to cgroup
		if err := os.WriteFile(cgroupProcsFile, []byte(fmt.Sprintf("%d", pid)), 0644); err != nil {
			return fmt.Errorf("failed to add process to %s cgroup: %w", controller, err)
		}
	}
	
	return nil
}

// GetResourceUsage gets current resource usage
func (cgm *CGroupManager) GetResourceUsage() (*ResourceUsage, error) {
	usage := &ResourceUsage{
		Timestamp: time.Now(),
	}
	
	if cgm.isCGroupV2Available() {
		return cgm.getCGroupV2Usage(usage)
	}
	
	return cgm.getCGroupV1Usage(usage)
}

// getCGroupV2Usage gets usage from cgroup v2
func (cgm *CGroupManager) getCGroupV2Usage(usage *ResourceUsage) error {
	containerCGroupPath := filepath.Join(cgm.cgroupPath, cgm.containerID)
	
	// Get memory usage
	memoryFile := filepath.Join(containerCGroupPath, "memory.current")
	if data, err := os.ReadFile(memoryFile); err == nil {
		if memory, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64); err == nil {
			usage.MemoryUsage = uint64(memory)
		}
	}
	
	// Get CPU usage
	cpuFile := filepath.Join(containerCGroupPath, "cpu.stat")
	if data, err := os.ReadFile(cpuFile); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "usage_usec") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					if cpu, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
						usage.CPUUsage = float64(cpu) / 1000000.0 // Convert to seconds
					}
				}
			}
		}
	}
	
	return nil
}

// getCGroupV1Usage gets usage from cgroup v1
func (cgm *CGroupManager) getCGroupV1Usage(usage *ResourceUsage) error {
	// Get memory usage
	memoryFile := filepath.Join(cgm.cgroupPath, "memory", cgm.containerID, "memory.usage_in_bytes")
	if data, err := os.ReadFile(memoryFile); err == nil {
		if memory, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64); err == nil {
			usage.MemoryUsage = uint64(memory)
		}
	}
	
	// Get CPU usage
	cpuFile := filepath.Join(cgm.cgroupPath, "cpu", cgm.containerID, "cpuacct.usage")
	if data, err := os.ReadFile(cpuFile); err == nil {
		if cpu, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64); err == nil {
			usage.CPUUsage = float64(cpu) / 1000000000.0 // Convert to seconds
		}
	}
	
	return nil
}

// CleanupCGroup removes the cgroup
func (cgm *CGroupManager) CleanupCGroup() error {
	fmt.Printf("🧹 Cleaning up cgroup for container: %s\n", cgm.containerID)
	
	if cgm.isCGroupV2Available() {
		return cgm.cleanupCGroupV2()
	}
	
	return cgm.cleanupCGroupV1()
}

// cleanupCGroupV2 removes cgroup v2
func (cgm *CGroupManager) cleanupCGroupV2() error {
	containerCGroupPath := filepath.Join(cgm.cgroupPath, cgm.containerID)
	
	// Remove all processes from cgroup first
	cgroupProcsFile := filepath.Join(containerCGroupPath, "cgroup.procs")
	if data, err := os.ReadFile(cgroupProcsFile); err == nil {
		pids := strings.Split(strings.TrimSpace(string(data)), "\n")
		for _, pid := range pids {
			if pid != "" {
				// Move process to parent cgroup
				parentCGroupPath := filepath.Join(cgm.cgroupPath, "cgroup.procs")
				os.WriteFile(parentCGroupPath, []byte(pid), 0644)
			}
		}
	}
	
	// Remove cgroup directory
	if err := os.RemoveAll(containerCGroupPath); err != nil {
		return fmt.Errorf("failed to remove cgroup: %w", err)
	}
	
	return nil
}

// cleanupCGroupV1 removes cgroup v1
func (cgm *CGroupManager) cleanupCGroupV1() error {
	controllers := []string{"memory", "cpu", "pids", "blkio"}
	
	for _, controller := range controllers {
		cgroupPath := filepath.Join(cgm.cgroupPath, controller, cgm.containerID)
		
		// Remove all processes from cgroup first
		cgroupProcsFile := filepath.Join(cgroupPath, "cgroup.procs")
		if data, err := os.ReadFile(cgroupProcsFile); err == nil {
			pids := strings.Split(strings.TrimSpace(string(data)), "\n")
			for _, pid := range pids {
				if pid != "" {
					// Move process to parent cgroup
					parentCGroupPath := filepath.Join(cgm.cgroupPath, controller, "cgroup.procs")
					os.WriteFile(parentCGroupPath, []byte(pid), 0644)
				}
			}
		}
		
		// Remove cgroup directory
		if err := os.RemoveAll(cgroupPath); err != nil {
			return fmt.Errorf("failed to remove %s cgroup: %w", controller, err)
		}
	}
	
	return nil
}

// isCGroupV2Available checks if cgroups v2 is available
func (cgm *CGroupManager) isCGroupV2Available() bool {
	// Check if cgroups v2 is mounted
	if _, err := os.Stat("/sys/fs/cgroup/cgroup.controllers"); err == nil {
		return true
	}
	
	// Check if cgroups v2 is available in kernel
	if _, err := os.Stat("/sys/fs/cgroup/cgroup.procs"); err == nil {
		return true
	}
	
	return false
}

// CheckCGroupSupport checks if cgroups are supported
func (cgm *CGroupManager) CheckCGroupSupport() error {
	fmt.Println("🔍 Checking cgroup support...")
	
	// Check if cgroup filesystem is mounted
	if _, err := os.Stat(cgm.cgroupPath); err != nil {
		return fmt.Errorf("cgroup filesystem not mounted: %w", err)
	}
	
	// Check if we have write permissions
	if err := os.MkdirAll(filepath.Join(cgm.cgroupPath, "test"), 0755); err != nil {
		return fmt.Errorf("no write permission to cgroup filesystem: %w", err)
	}
	
	// Clean up test directory
	os.RemoveAll(filepath.Join(cgm.cgroupPath, "test"))
	
	fmt.Println("✅ CGroup support available")
	return nil
}

// PrintCGroupInfo prints cgroup information
func (cgm *CGroupManager) PrintCGroupInfo() {
	fmt.Println("⚡ CGroup Information")
	fmt.Println("====================")
	
	// Check cgroup support
	if err := cgm.CheckCGroupSupport(); err != nil {
		fmt.Printf("❌ CGroup support check failed: %v\n", err)
		return
	}
	
	// Check cgroup version
	if cgm.isCGroupV2Available() {
		fmt.Println("✅ CGroup v2 available")
	} else {
		fmt.Println("✅ CGroup v1 available")
	}
	
	// Get current cgroup
	if currentCGroup, err := cgm.getCurrentCGroup(); err == nil {
		fmt.Printf("Current cgroup: %s\n", currentCGroup)
	}
}

// getCurrentCGroup gets the current cgroup
func (cgm *CGroupManager) getCurrentCGroup() (string, error) {
	if cgm.isCGroupV2Available() {
		data, err := os.ReadFile("/proc/self/cgroup")
		if err != nil {
			return "", err
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 {
			parts := strings.Split(lines[0], ":")
			if len(parts) >= 3 {
				return parts[2], nil
			}
		}
	}
	
	return "", fmt.Errorf("unable to determine current cgroup")
}

// ResourceUsage represents current resource usage
type ResourceUsage struct {
	Timestamp   time.Time `json:"timestamp"`
	MemoryUsage uint64    `json:"memory_usage_bytes"`
	CPUUsage    float64   `json:"cpu_usage_seconds"`
	PIDsCount   int       `json:"pids_count"`
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
