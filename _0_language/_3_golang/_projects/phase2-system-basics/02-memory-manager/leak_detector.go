package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// DetectLeaks detects memory leaks
func (mm *MemoryManager) DetectLeaks(opts *LeakDetectOptions) (*LeakReport, error) {
	report := &LeakReport{
		Timestamp: time.Now(),
		Threshold: opts.Threshold,
		Duration:  opts.Duration,
		Leaks:     make([]LeakInfo, 0),
		Summary:   &LeakSummary{},
	}
	
	// Get initial memory stats
	initialStats, err := mm.GetMemoryStats()
	if err != nil {
		return nil, fmt.Errorf("failed to get initial stats: %w", err)
	}
	
	// Monitor memory over time
	monitorDuration := opts.Duration
	if monitorDuration == 0 {
		monitorDuration = 30 * time.Second
	}
	
	interval := 1 * time.Second
	iterations := int(monitorDuration / interval)
	
	measurements := make([]*MemoryStats, 0, iterations)
	
	for i := 0; i < iterations; i++ {
		stats, err := mm.GetMemoryStats()
		if err != nil {
			continue
		}
		measurements = append(measurements, stats)
		time.Sleep(interval)
	}
	
	// Analyze measurements for leaks
	report.Leaks = mm.analyzeLeaks(measurements, opts.Threshold)
	report.Summary = mm.calculateLeakSummary(report.Leaks, initialStats, measurements)
	
	return report, nil
}

// LeakReport represents a memory leak detection report
type LeakReport struct {
	Timestamp time.Time   `json:"timestamp"`
	Threshold float64     `json:"threshold"`
	Duration  time.Duration `json:"duration"`
	Leaks     []LeakInfo  `json:"leaks"`
	Summary   *LeakSummary `json:"summary"`
}

// LeakInfo represents information about a detected leak
type LeakInfo struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Size        uint64    `json:"size"`
	Rate        float64   `json:"rate"`
	Location    string    `json:"location"`
}

// LeakSummary represents a summary of leak detection
type LeakSummary struct {
	TotalLeaks     int     `json:"total_leaks"`
	CriticalLeaks  int     `json:"critical_leaks"`
	WarningLeaks   int     `json:"warning_leaks"`
	InfoLeaks      int     `json:"info_leaks"`
	TotalSize      uint64  `json:"total_size"`
	AverageRate    float64 `json:"average_rate"`
	MaxRate        float64 `json:"max_rate"`
	LeakScore      float64 `json:"leak_score"`
}

// analyzeLeaks analyzes memory measurements for leaks
func (mm *MemoryManager) analyzeLeaks(measurements []*MemoryStats, threshold float64) []LeakInfo {
	leaks := make([]LeakInfo, 0)
	
	if len(measurements) < 2 {
		return leaks
	}
	
	// Analyze heap growth
	heapLeaks := mm.analyzeHeapGrowth(measurements, threshold)
	leaks = append(leaks, heapLeaks...)
	
	// Analyze goroutine growth
	goroutineLeaks := mm.analyzeGoroutineGrowth(measurements, threshold)
	leaks = append(leaks, goroutineLeaks...)
	
	// Analyze object count growth
	objectLeaks := mm.analyzeObjectGrowth(measurements, threshold)
	leaks = append(leaks, objectLeaks...)
	
	return leaks
}

// analyzeHeapGrowth analyzes heap memory growth
func (mm *MemoryManager) analyzeHeapGrowth(measurements []*MemoryStats, threshold float64) []LeakInfo {
	leaks := make([]LeakInfo, 0)
	
	if len(measurements) < 2 {
		return leaks
	}
	
	initial := measurements[0]
	final := measurements[len(measurements)-1]
	
	// Calculate growth rate
	growth := int64(final.HeapAlloc) - int64(initial.HeapAlloc)
	rate := float64(growth) / float64(len(measurements)-1) // bytes per second
	
	// Check if growth exceeds threshold
	if rate > 0 && float64(growth)/float64(initial.HeapAlloc)*100 > threshold {
		severity := "info"
		if rate > 1024*1024 { // > 1MB/s
			severity = "critical"
		} else if rate > 1024*100 { // > 100KB/s
			severity = "warning"
		}
		
		leak := LeakInfo{
			Type:        "heap_growth",
			Description: fmt.Sprintf("Heap memory growing at %.2f bytes/second", rate),
			Severity:    severity,
			StartTime:   initial.Timestamp,
			EndTime:     final.Timestamp,
			Size:        uint64(growth),
			Rate:        rate,
			Location:    "heap",
		}
		leaks = append(leaks, leak)
	}
	
	return leaks
}

// analyzeGoroutineGrowth analyzes goroutine count growth
func (mm *MemoryManager) analyzeGoroutineGrowth(measurements []*MemoryStats, threshold float64) []LeakInfo {
	leaks := make([]LeakInfo, 0)
	
	if len(measurements) < 2 {
		return leaks
	}
	
	initial := measurements[0]
	final := measurements[len(measurements)-1]
	
	// Calculate growth rate
	growth := final.NumGoroutines - initial.NumGoroutines
	rate := float64(growth) / float64(len(measurements)-1) // goroutines per second
	
	// Check if growth exceeds threshold
	if rate > 0 && float64(growth)/float64(initial.NumGoroutines)*100 > threshold {
		severity := "info"
		if rate > 10 { // > 10 goroutines/second
			severity = "critical"
		} else if rate > 1 { // > 1 goroutine/second
			severity = "warning"
		}
		
		leak := LeakInfo{
			Type:        "goroutine_growth",
			Description: fmt.Sprintf("Goroutine count growing at %.2f goroutines/second", rate),
			Severity:    severity,
			StartTime:   initial.Timestamp,
			EndTime:     final.Timestamp,
			Size:        uint64(growth),
			Rate:        rate,
			Location:    "goroutines",
		}
		leaks = append(leaks, leak)
	}
	
	return leaks
}

// analyzeObjectGrowth analyzes heap object count growth
func (mm *MemoryManager) analyzeObjectGrowth(measurements []*MemoryStats, threshold float64) []LeakInfo {
	leaks := make([]LeakInfo, 0)
	
	if len(measurements) < 2 {
		return leaks
	}
	
	initial := measurements[0]
	final := measurements[len(measurements)-1]
	
	// Calculate growth rate
	growth := int64(final.HeapObjects) - int64(initial.HeapObjects)
	rate := float64(growth) / float64(len(measurements)-1) // objects per second
	
	// Check if growth exceeds threshold
	if rate > 0 && float64(growth)/float64(initial.HeapObjects)*100 > threshold {
		severity := "info"
		if rate > 1000 { // > 1000 objects/second
			severity = "critical"
		} else if rate > 100 { // > 100 objects/second
			severity = "warning"
		}
		
		leak := LeakInfo{
			Type:        "object_growth",
			Description: fmt.Sprintf("Heap objects growing at %.2f objects/second", rate),
			Severity:    severity,
			StartTime:   initial.Timestamp,
			EndTime:     final.Timestamp,
			Size:        uint64(growth),
			Rate:        rate,
			Location:    "heap_objects",
		}
		leaks = append(leaks, leak)
	}
	
	return leaks
}

// calculateLeakSummary calculates leak detection summary
func (mm *MemoryManager) calculateLeakSummary(leaks []LeakInfo, initial *MemoryStats, measurements []*MemoryStats) *LeakSummary {
	summary := &LeakSummary{
		TotalLeaks: len(leaks),
	}
	
	// Count leaks by severity
	for _, leak := range leaks {
		switch leak.Severity {
		case "critical":
			summary.CriticalLeaks++
		case "warning":
			summary.WarningLeaks++
		case "info":
			summary.InfoLeaks++
		}
		summary.TotalSize += leak.Size
	}
	
	// Calculate rates
	if len(leaks) > 0 {
		totalRate := 0.0
		maxRate := 0.0
		for _, leak := range leaks {
			totalRate += leak.Rate
			if leak.Rate > maxRate {
				maxRate = leak.Rate
			}
		}
		summary.AverageRate = totalRate / float64(len(leaks))
		summary.MaxRate = maxRate
	}
	
	// Calculate leak score (0-100, higher is worse)
	summary.LeakScore = float64(summary.CriticalLeaks)*50 + 
		float64(summary.WarningLeaks)*25 + 
		float64(summary.InfoLeaks)*10
	
	if summary.LeakScore > 100 {
		summary.LeakScore = 100
	}
	
	return summary
}

// PrintLeakReport prints leak detection report
func (mm *MemoryManager) PrintLeakReport(report *LeakReport) {
	fmt.Println("Memory Leak Detection Report")
	fmt.Println("============================")
	fmt.Printf("Timestamp: %s\n", report.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %v\n", report.Duration)
	fmt.Printf("Threshold: %.2f%%\n", report.Threshold)
	fmt.Println()
	
	// Print summary
	fmt.Println("Summary:")
	fmt.Printf("  Total Leaks: %d\n", report.Summary.TotalLeaks)
	fmt.Printf("  Critical: %d\n", report.Summary.CriticalLeaks)
	fmt.Printf("  Warning: %d\n", report.Summary.WarningLeaks)
	fmt.Printf("  Info: %d\n", report.Summary.InfoLeaks)
	fmt.Printf("  Total Size: %s\n", formatBytes(report.Summary.TotalSize))
	fmt.Printf("  Average Rate: %.2f bytes/second\n", report.Summary.AverageRate)
	fmt.Printf("  Max Rate: %.2f bytes/second\n", report.Summary.MaxRate)
	fmt.Printf("  Leak Score: %.2f/100\n", report.Summary.LeakScore)
	fmt.Println()
	
	// Print leaks
	if len(report.Leaks) > 0 {
		fmt.Println("Detected Leaks:")
		fmt.Printf("%-20s %-10s %-30s %-15s %-15s %-10s\n",
			"Type", "Severity", "Description", "Size", "Rate", "Location")
		fmt.Println(strings.Repeat("-", 100))
		
		for _, leak := range report.Leaks {
			fmt.Printf("%-20s %-10s %-30s %-15s %-15.2f %-10s\n",
				leak.Type,
				leak.Severity,
				truncateString(leak.Description, 30),
				formatBytes(leak.Size),
				leak.Rate,
				leak.Location)
		}
		fmt.Println()
	} else {
		fmt.Println("No memory leaks detected! 🎉")
		fmt.Println()
	}
	
	// Print recommendations
	if report.Summary.LeakScore > 50 {
		fmt.Println("Recommendations:")
		if report.Summary.CriticalLeaks > 0 {
			fmt.Println("  - Investigate critical leaks immediately")
		}
		if report.Summary.WarningLeaks > 0 {
			fmt.Println("  - Monitor warning leaks closely")
		}
		fmt.Println("  - Consider running memory profiling")
		fmt.Println("  - Review code for potential memory leaks")
		fmt.Println("  - Implement proper resource cleanup")
	}
}

// MemoryLeakDetector represents a memory leak detector
type MemoryLeakDetector struct {
	measurements []*MemoryStats
	mutex        sync.RWMutex
	threshold    float64
}

// NewMemoryLeakDetector creates a new memory leak detector
func NewMemoryLeakDetector(threshold float64) *MemoryLeakDetector {
	return &MemoryLeakDetector{
		measurements: make([]*MemoryStats, 0),
		threshold:    threshold,
	}
}

// AddMeasurement adds a memory measurement
func (mld *MemoryLeakDetector) AddMeasurement(stats *MemoryStats) {
	mld.mutex.Lock()
	defer mld.mutex.Unlock()
	
	mld.measurements = append(mld.measurements, stats)
	
	// Keep only last 100 measurements
	if len(mld.measurements) > 100 {
		mld.measurements = mld.measurements[1:]
	}
}

// DetectLeaks detects leaks from measurements
func (mld *MemoryLeakDetector) DetectLeaks() []LeakInfo {
	mld.mutex.RLock()
	defer mld.mutex.RUnlock()
	
	if len(mld.measurements) < 2 {
		return []LeakInfo{}
	}
	
	// Analyze for leaks
	leaks := make([]LeakInfo, 0)
	
	// Check heap growth
	initial := mld.measurements[0]
	final := mld.measurements[len(mld.measurements)-1]
	
	heapGrowth := int64(final.HeapAlloc) - int64(initial.HeapAlloc)
	if heapGrowth > 0 {
		rate := float64(heapGrowth) / float64(len(mld.measurements)-1)
		if rate > 0 && float64(heapGrowth)/float64(initial.HeapAlloc)*100 > mld.threshold {
			leak := LeakInfo{
				Type:        "heap_growth",
				Description: fmt.Sprintf("Heap growing at %.2f bytes/second", rate),
				Severity:    "warning",
				StartTime:   initial.Timestamp,
				EndTime:     final.Timestamp,
				Size:        uint64(heapGrowth),
				Rate:        rate,
				Location:    "heap",
			}
			leaks = append(leaks, leak)
		}
	}
	
	return leaks
}
