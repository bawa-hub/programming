package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// SystemProfiler manages system performance profiling
type SystemProfiler struct {
	ctx        context.Context
	cancel     context.CancelFunc
	mutex      sync.RWMutex
	profiles   map[string]*ProfileInfo
	metrics    *SystemMetrics
}

// ProfileInfo represents profile information
type ProfileInfo struct {
	Type      string    `json:"type"`
	Filename  string    `json:"filename"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  time.Duration `json:"duration"`
	Size      int64     `json:"size"`
	Active    bool      `json:"active"`
}

// SystemMetrics represents system performance metrics
type SystemMetrics struct {
	Timestamp     time.Time     `json:"timestamp"`
	CPU           *CPUMetrics   `json:"cpu"`
	Memory        *MemoryMetrics `json:"memory"`
	Goroutines    *GoroutineMetrics `json:"goroutines"`
	GC            *GCMetrics    `json:"gc"`
	Runtime       *RuntimeMetrics `json:"runtime"`
}

// CPUMetrics represents CPU performance metrics
type CPUMetrics struct {
	Usage       float64   `json:"usage_percent"`
	Count       int       `json:"count"`
	Frequency   float64   `json:"frequency_mhz"`
	LoadAvg     []float64 `json:"load_avg"`
	Temperature float64   `json:"temperature_c"`
}

// MemoryMetrics represents memory performance metrics
type MemoryMetrics struct {
	Total       uint64  `json:"total_bytes"`
	Used        uint64  `json:"used_bytes"`
	Available   uint64  `json:"available_bytes"`
	Free        uint64  `json:"free_bytes"`
	UsedPercent float64 `json:"used_percent"`
	SwapTotal   uint64  `json:"swap_total_bytes"`
	SwapUsed    uint64  `json:"swap_used_bytes"`
	SwapFree    uint64  `json:"swap_free_bytes"`
}

// GoroutineMetrics represents goroutine metrics
type GoroutineMetrics struct {
	Count       int           `json:"count"`
	MaxCount    int           `json:"max_count"`
	AvgCount    float64       `json:"avg_count"`
	GrowthRate  float64       `json:"growth_rate"`
	Blocked     int           `json:"blocked"`
	Running     int           `json:"running"`
	Waiting     int           `json:"waiting"`
	Created     int64         `json:"created"`
	Destroyed   int64         `json:"destroyed"`
}

// GCMetrics represents garbage collection metrics
type GCMetrics struct {
	Cycles       uint32        `json:"cycles"`
	PauseTotal   time.Duration `json:"pause_total_ns"`
	PauseAvg     time.Duration `json:"pause_avg_ns"`
	PauseMax     time.Duration `json:"pause_max_ns"`
	PauseMin     time.Duration `json:"pause_min_ns"`
	CPUPercent   float64       `json:"cpu_percent"`
	HeapSize     uint64        `json:"heap_size_bytes"`
	HeapObjects  uint64        `json:"heap_objects"`
	HeapAlloc    uint64        `json:"heap_alloc_bytes"`
	HeapSys      uint64        `json:"heap_sys_bytes"`
	HeapIdle     uint64        `json:"heap_idle_bytes"`
	HeapInuse    uint64        `json:"heap_inuse_bytes"`
	HeapReleased uint64        `json:"heap_released_bytes"`
}

// RuntimeMetrics represents Go runtime metrics
type RuntimeMetrics struct {
	Version         string        `json:"version"`
	GoOS            string        `json:"goos"`
	GoArch          string        `json:"goarch"`
	NumCPU          int           `json:"num_cpu"`
	NumGoroutine    int           `json:"num_goroutine"`
	NumCgoCall      int64         `json:"num_cgo_call"`
	MemStats        *runtime.MemStats `json:"mem_stats"`
	GCStats         interface{} `json:"gc_stats"`
}

// MonitorOptions contains options for monitoring
type MonitorOptions struct {
	Watch    bool
	Interval time.Duration
	Format   string
	Export   string
}

// NewSystemProfiler creates a new system profiler
func NewSystemProfiler() *SystemProfiler {
	ctx, cancel := context.WithCancel(context.Background())
	return &SystemProfiler{
		ctx:      ctx,
		cancel:   cancel,
		profiles: make(map[string]*ProfileInfo),
		metrics:  &SystemMetrics{},
	}
}

// Close closes the system profiler
func (sp *SystemProfiler) Close() {
	sp.cancel()
}

// StartCPUProfile starts CPU profiling
func (sp *SystemProfiler) StartCPUProfile(filename string) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CPU profile file: %w", err)
	}
	
	err = pprof.StartCPUProfile(file)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to start CPU profile: %w", err)
	}
	
	sp.profiles["cpu"] = &ProfileInfo{
		Type:      "cpu",
		Filename:  filename,
		StartTime: time.Now(),
		Active:    true,
	}
	
	return nil
}

// StopCPUProfile stops CPU profiling
func (sp *SystemProfiler) StopCPUProfile() error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	pprof.StopCPUProfile()
	
	if profile, exists := sp.profiles["cpu"]; exists {
		profile.EndTime = time.Now()
		profile.Duration = profile.EndTime.Sub(profile.StartTime)
		profile.Active = false
		
		// Get file size
		if stat, err := os.Stat(profile.Filename); err == nil {
			profile.Size = stat.Size()
		}
	}
	
	return nil
}

// StartMemoryProfile starts memory profiling
func (sp *SystemProfiler) StartMemoryProfile(filename string) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create memory profile file: %w", err)
	}
	
	// Force garbage collection before profiling
	runtime.GC()
	
	err = pprof.WriteHeapProfile(file)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to write memory profile: %w", err)
	}
	
	file.Close()
	
	sp.profiles["memory"] = &ProfileInfo{
		Type:      "memory",
		Filename:  filename,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Duration:  0,
		Active:    false,
	}
	
	// Get file size
	if stat, err := os.Stat(filename); err == nil {
		sp.profiles["memory"].Size = stat.Size()
	}
	
	return nil
}

// StopMemoryProfile stops memory profiling
func (sp *SystemProfiler) StopMemoryProfile() error {
	// Memory profiling is already stopped after WriteHeapProfile
	return nil
}

// StartGoroutineProfile starts goroutine profiling
func (sp *SystemProfiler) StartGoroutineProfile(filename string) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create goroutine profile file: %w", err)
	}
	
	err = pprof.Lookup("goroutine").WriteTo(file, 0)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to write goroutine profile: %w", err)
	}
	
	file.Close()
	
	sp.profiles["goroutine"] = &ProfileInfo{
		Type:      "goroutine",
		Filename:  filename,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Duration:  0,
		Active:    false,
	}
	
	// Get file size
	if stat, err := os.Stat(filename); err == nil {
		sp.profiles["goroutine"].Size = stat.Size()
	}
	
	return nil
}

// StopGoroutineProfile stops goroutine profiling
func (sp *SystemProfiler) StopGoroutineProfile() error {
	// Goroutine profiling is already stopped after WriteTo
	return nil
}

// StartBlockProfile starts block profiling
func (sp *SystemProfiler) StartBlockProfile(filename string) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create block profile file: %w", err)
	}
	
	err = pprof.Lookup("block").WriteTo(file, 0)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to write block profile: %w", err)
	}
	
	file.Close()
	
	sp.profiles["block"] = &ProfileInfo{
		Type:      "block",
		Filename:  filename,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Duration:  0,
		Active:    false,
	}
	
	// Get file size
	if stat, err := os.Stat(filename); err == nil {
		sp.profiles["block"].Size = stat.Size()
	}
	
	return nil
}

// StopBlockProfile stops block profiling
func (sp *SystemProfiler) StopBlockProfile() error {
	// Block profiling is already stopped after WriteTo
	return nil
}

// StartMutexProfile starts mutex profiling
func (sp *SystemProfiler) StartMutexProfile(filename string) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create mutex profile file: %w", err)
	}
	
	err = pprof.Lookup("mutex").WriteTo(file, 0)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to write mutex profile: %w", err)
	}
	
	file.Close()
	
	sp.profiles["mutex"] = &ProfileInfo{
		Type:      "mutex",
		Filename:  filename,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Duration:  0,
		Active:    false,
	}
	
	// Get file size
	if stat, err := os.Stat(filename); err == nil {
		sp.profiles["mutex"].Size = stat.Size()
	}
	
	return nil
}

// StopMutexProfile stops mutex profiling
func (sp *SystemProfiler) StopMutexProfile() error {
	// Mutex profiling is already stopped after WriteTo
	return nil
}

// GetSystemMetrics gets current system metrics
func (sp *SystemProfiler) GetSystemMetrics() (*SystemMetrics, error) {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	metrics := &SystemMetrics{
		Timestamp: time.Now(),
	}
	
	// Get CPU metrics
	cpuMetrics, err := sp.getCPUMetrics()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU metrics: %w", err)
	}
	metrics.CPU = cpuMetrics
	
	// Get memory metrics
	memMetrics, err := sp.getMemoryMetrics()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory metrics: %w", err)
	}
	metrics.Memory = memMetrics
	
	// Get goroutine metrics
	goroutineMetrics := sp.getGoroutineMetrics()
	metrics.Goroutines = goroutineMetrics
	
	// Get GC metrics
	gcMetrics := sp.getGCMetrics()
	metrics.GC = gcMetrics
	
	// Get runtime metrics
	runtimeMetrics := sp.getRuntimeMetrics()
	metrics.Runtime = runtimeMetrics
	
	sp.metrics = metrics
	return metrics, nil
}

// getCPUMetrics gets CPU performance metrics
func (sp *SystemProfiler) getCPUMetrics() (*CPUMetrics, error) {
	// Get CPU usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	
	// Get CPU info
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	
	// Get load average (simplified)
	loadAvg := []float64{0, 0, 0}
	
	usage := 0.0
	if len(cpuPercent) > 0 {
		usage = cpuPercent[0]
	}
	
	frequency := 0.0
	if len(cpuInfo) > 0 {
		frequency = cpuInfo[0].Mhz
	}
	
	return &CPUMetrics{
		Usage:       usage,
		Count:       len(cpuInfo),
		Frequency:   frequency,
		LoadAvg:     loadAvg,
		Temperature: 0, // Not available on all systems
	}, nil
}

// getMemoryMetrics gets memory performance metrics
func (sp *SystemProfiler) getMemoryMetrics() (*MemoryMetrics, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		// Swap not available on all systems
		swapInfo = &mem.SwapMemoryStat{}
	}
	
	return &MemoryMetrics{
		Total:       memInfo.Total,
		Used:        memInfo.Used,
		Available:   memInfo.Available,
		Free:        memInfo.Free,
		UsedPercent: memInfo.UsedPercent,
		SwapTotal:   swapInfo.Total,
		SwapUsed:    swapInfo.Used,
		SwapFree:    swapInfo.Free,
	}, nil
}

// getGoroutineMetrics gets goroutine metrics
func (sp *SystemProfiler) getGoroutineMetrics() *GoroutineMetrics {
	count := runtime.NumGoroutine()
	
	return &GoroutineMetrics{
		Count:      count,
		MaxCount:   count, // Simplified - would need historical tracking
		AvgCount:   float64(count), // Simplified
		GrowthRate: 0, // Simplified
		Blocked:    0, // Simplified
		Running:    count, // Simplified
		Waiting:    0, // Simplified
		Created:    0, // Simplified
		Destroyed:  0, // Simplified
	}
}

// getGCMetrics gets garbage collection metrics
func (sp *SystemProfiler) getGCMetrics() *GCMetrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	
	return &GCMetrics{
		Cycles:       memStats.NumGC,
		PauseTotal:   time.Duration(memStats.PauseTotalNs),
		PauseAvg:     time.Duration(memStats.PauseTotalNs) / time.Duration(max(1, memStats.NumGC)),
		PauseMax:     time.Duration(memStats.PauseNs[(memStats.NumGC+255)%256]),
		PauseMin:     time.Duration(memStats.PauseNs[(memStats.NumGC+255)%256]),
		CPUPercent:   float64(memStats.GCSys) / float64(time.Second) * 100,
		HeapSize:     memStats.HeapSys,
		HeapObjects:  memStats.HeapObjects,
		HeapAlloc:    memStats.HeapAlloc,
		HeapSys:      memStats.HeapSys,
		HeapIdle:     memStats.HeapIdle,
		HeapInuse:    memStats.HeapInuse,
		HeapReleased: memStats.HeapReleased,
	}
}

// getRuntimeMetrics gets Go runtime metrics
func (sp *SystemProfiler) getRuntimeMetrics() *RuntimeMetrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	
	return &RuntimeMetrics{
		Version:      runtime.Version(),
		GoOS:         runtime.GOOS,
		GoArch:       runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
		NumCgoCall:   runtime.NumCgoCall(),
		MemStats:     &memStats,
		GCStats:      nil, // Simplified
	}
}

// MonitorRealtime monitors system performance in real-time
func (sp *SystemProfiler) MonitorRealtime(opts *MonitorOptions) {
	fmt.Println("Starting real-time system monitoring...")
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
			sp.updateMonitorDisplay(opts, startTime)
		case <-sp.ctx.Done():
			return
		}
	}
}

// updateMonitorDisplay updates the monitoring display
func (sp *SystemProfiler) updateMonitorDisplay(opts *MonitorOptions, startTime time.Time) {
	// Move cursor to top
	fmt.Print("\033[H")
	
	// Get system metrics
	metrics, err := sp.GetSystemMetrics()
	if err != nil {
		fmt.Printf("Error getting system metrics: %v\n", err)
		return
	}
	
	// Print header
	fmt.Println("System Profiler - Real-time")
	fmt.Println("===========================")
	fmt.Printf("Time: %s (Running: %v)\n", 
		time.Now().Format("2006-01-02 15:04:05"), 
		time.Since(startTime).Round(time.Second))
	fmt.Println()
	
	// Print system metrics
	sp.printSystemMetrics(metrics)
}

// printSystemMetrics prints system metrics
func (sp *SystemProfiler) printSystemMetrics(metrics *SystemMetrics) {
	// CPU metrics
	fmt.Printf("CPU Usage: %.2f%% (Load: %.2f, %.2f, %.2f)\n", 
		metrics.CPU.Usage, 
		metrics.CPU.LoadAvg[0], 
		metrics.CPU.LoadAvg[1], 
		metrics.CPU.LoadAvg[2])
	
	// Memory metrics
	fmt.Printf("Memory: %s / %s (%.2f%%)\n", 
		formatBytes(metrics.Memory.Used), 
		formatBytes(metrics.Memory.Total), 
		metrics.Memory.UsedPercent)
	
	// Goroutine metrics
	fmt.Printf("Goroutines: %d\n", metrics.Goroutines.Count)
	
	// GC metrics
	fmt.Printf("GC Cycles: %d (Pause: %v avg)\n", 
		metrics.GC.Cycles, 
		metrics.GC.PauseAvg)
	
	// Runtime metrics
	fmt.Printf("Go Version: %s (%s/%s)\n", 
		metrics.Runtime.Version, 
		metrics.Runtime.GoOS, 
		metrics.Runtime.GoArch)
	
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

// max returns the maximum of two uint32 values
func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

// PrintSystemMetrics prints system metrics
func (sp *SystemProfiler) PrintSystemMetrics(metrics *SystemMetrics) {
	fmt.Println("System Performance Metrics")
	fmt.Println("=========================")
	fmt.Printf("Timestamp: %s\n", metrics.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println()
	
	// CPU metrics
	fmt.Println("CPU:")
	fmt.Printf("  Usage: %.2f%%\n", metrics.CPU.Usage)
	fmt.Printf("  Count: %d\n", metrics.CPU.Count)
	fmt.Printf("  Frequency: %.2f MHz\n", metrics.CPU.Frequency)
	fmt.Printf("  Load Average: %.2f, %.2f, %.2f\n", 
		metrics.CPU.LoadAvg[0], 
		metrics.CPU.LoadAvg[1], 
		metrics.CPU.LoadAvg[2])
	fmt.Println()
	
	// Memory metrics
	fmt.Println("Memory:")
	fmt.Printf("  Total: %s\n", formatBytes(metrics.Memory.Total))
	fmt.Printf("  Used: %s (%.2f%%)\n", 
		formatBytes(metrics.Memory.Used), 
		metrics.Memory.UsedPercent)
	fmt.Printf("  Available: %s\n", formatBytes(metrics.Memory.Available))
	fmt.Printf("  Free: %s\n", formatBytes(metrics.Memory.Free))
	fmt.Printf("  Swap Total: %s\n", formatBytes(metrics.Memory.SwapTotal))
	fmt.Printf("  Swap Used: %s\n", formatBytes(metrics.Memory.SwapUsed))
	fmt.Println()
	
	// Goroutine metrics
	fmt.Println("Goroutines:")
	fmt.Printf("  Count: %d\n", metrics.Goroutines.Count)
	fmt.Printf("  Max Count: %d\n", metrics.Goroutines.MaxCount)
	fmt.Printf("  Average Count: %.2f\n", metrics.Goroutines.AvgCount)
	fmt.Printf("  Growth Rate: %.2f\n", metrics.Goroutines.GrowthRate)
	fmt.Println()
	
	// GC metrics
	fmt.Println("Garbage Collection:")
	fmt.Printf("  Cycles: %d\n", metrics.GC.Cycles)
	fmt.Printf("  Pause Total: %v\n", metrics.GC.PauseTotal)
	fmt.Printf("  Pause Average: %v\n", metrics.GC.PauseAvg)
	fmt.Printf("  Pause Max: %v\n", metrics.GC.PauseMax)
	fmt.Printf("  CPU Percent: %.4f%%\n", metrics.GC.CPUPercent)
	fmt.Printf("  Heap Size: %s\n", formatBytes(metrics.GC.HeapSize))
	fmt.Printf("  Heap Objects: %d\n", metrics.GC.HeapObjects)
	fmt.Println()
	
	// Runtime metrics
	fmt.Println("Runtime:")
	fmt.Printf("  Version: %s\n", metrics.Runtime.Version)
	fmt.Printf("  OS: %s\n", metrics.Runtime.GoOS)
	fmt.Printf("  Architecture: %s\n", metrics.Runtime.GoArch)
	fmt.Printf("  CPU Count: %d\n", metrics.Runtime.NumCPU)
	fmt.Printf("  Goroutine Count: %d\n", metrics.Runtime.NumGoroutine)
	fmt.Printf("  CGO Calls: %d\n", metrics.Runtime.NumCgoCall)
}

// PrintJSON prints data in JSON format
func (sp *SystemProfiler) PrintJSON(data interface{}) {
	// This is a simplified version - in a real implementation,
	// you would use encoding/json
	fmt.Printf("JSON output: %+v\n", data)
}

// PrintCSV prints data in CSV format
func (sp *SystemProfiler) PrintCSV(data interface{}) {
	// This is a simplified version - in a real implementation,
	// you would use encoding/csv
	fmt.Printf("CSV output: %+v\n", data)
}
