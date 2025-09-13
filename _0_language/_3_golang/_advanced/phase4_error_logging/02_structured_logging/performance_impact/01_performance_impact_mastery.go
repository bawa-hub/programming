// ⚡ PERFORMANCE IMPACT MASTERY
// Advanced logging performance optimization and monitoring
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// PERFORMANCE METRICS
// ============================================================================

type PerformanceMetrics struct {
	TotalLogs        int64         `json:"total_logs"`
	FilteredLogs     int64         `json:"filtered_logs"`
	TotalTime        time.Duration `json:"total_time"`
	AverageTime      time.Duration `json:"average_time"`
	MaxTime          time.Duration `json:"max_time"`
	MinTime          time.Duration `json:"min_time"`
	MemoryAllocated  int64         `json:"memory_allocated"`
	MemoryAllocs     int64         `json:"memory_allocs"`
	GCCollections    int64         `json:"gc_collections"`
	CPUUsage         float64       `json:"cpu_usage"`
}

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func (l LogLevel) Priority() int {
	return int(l)
}

// ============================================================================
// HIGH-PERFORMANCE LOGGER
// ============================================================================

type HighPerformanceLogger struct {
	level       LogLevel
	output      io.Writer
	metrics     *PerformanceMetrics
	mu          sync.RWMutex
	asyncChan   chan LogEntry
	asyncWg     sync.WaitGroup
	asyncStop   chan struct{}
	pool        *LogEntryPool
}

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

type LogEntryPool struct {
	pool sync.Pool
}

func NewLogEntryPool() *LogEntryPool {
	return &LogEntryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &LogEntry{
					Fields: make(map[string]interface{}),
				}
			},
		},
	}
}

func (p *LogEntryPool) Get() *LogEntry {
	entry := p.pool.Get().(*LogEntry)
	entry.Timestamp = time.Time{}
	entry.Level = 0
	entry.Message = ""
	// Clear fields map
	for k := range entry.Fields {
		delete(entry.Fields, k)
	}
	return entry
}

func (p *LogEntryPool) Put(entry *LogEntry) {
	p.pool.Put(entry)
}

func NewHighPerformanceLogger(level LogLevel, output io.Writer, async bool) *HighPerformanceLogger {
	logger := &HighPerformanceLogger{
		level:     level,
		output:    output,
		metrics:   &PerformanceMetrics{},
		pool:      NewLogEntryPool(),
	}
	
	if async {
		logger.asyncChan = make(chan LogEntry, 10000) // Large buffer
		logger.asyncStop = make(chan struct{})
		logger.startAsyncWorker()
	}
	
	return logger
}

func (hpl *HighPerformanceLogger) startAsyncWorker() {
	hpl.asyncWg.Add(1)
	go func() {
		defer hpl.asyncWg.Done()
		for {
			select {
			case entry := <-hpl.asyncChan:
				hpl.writeLog(entry)
			case <-hpl.asyncStop:
				return
			}
		}
	}()
}

func (hpl *HighPerformanceLogger) Stop() {
	if hpl.asyncChan != nil {
		close(hpl.asyncStop)
		hpl.asyncWg.Wait()
	}
}

func (hpl *HighPerformanceLogger) log(level LogLevel, message string, fields map[string]interface{}) {
	start := time.Now()
	
	// Quick level check
	if level.Priority() < hpl.level.Priority() {
		atomic.AddInt64(&hpl.metrics.FilteredLogs, 1)
		return
	}
	
	// Create log entry
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}
	
	// Update metrics
	atomic.AddInt64(&hpl.metrics.TotalLogs, 1)
	
	// Write log (sync or async)
	if hpl.asyncChan != nil {
		select {
		case hpl.asyncChan <- entry:
		default:
			// Channel full, fallback to sync
			hpl.writeLog(entry)
		}
	} else {
		hpl.writeLog(entry)
	}
	
	// Update timing metrics
	duration := time.Since(start)
	hpl.updateTiming(duration)
}

func (hpl *HighPerformanceLogger) writeLog(entry LogEntry) {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		// Fallback to simple format
		fmt.Fprintf(hpl.output, "[%s] %s: %s\n", 
			entry.Level.String(), 
			entry.Timestamp.Format(time.RFC3339), 
			entry.Message)
		return
	}
	
	fmt.Fprintln(hpl.output, string(jsonData))
}

func (hpl *HighPerformanceLogger) updateTiming(duration time.Duration) {
	// Update total time
	atomic.AddInt64((*int64)(&hpl.metrics.TotalTime), int64(duration))
	
	// Update average time (simple moving average)
	currentAvg := atomic.LoadInt64((*int64)(&hpl.metrics.AverageTime))
	totalLogs := atomic.LoadInt64(&hpl.metrics.TotalLogs)
	if totalLogs > 0 {
		newAvg := (currentAvg + int64(duration)) / 2
		atomic.StoreInt64((*int64)(&hpl.metrics.AverageTime), newAvg)
	}
	
	// Update max time
	for {
		currentMax := atomic.LoadInt64((*int64)(&hpl.metrics.MaxTime))
		if int64(duration) <= currentMax {
			break
		}
		if atomic.CompareAndSwapInt64((*int64)(&hpl.metrics.MaxTime), currentMax, int64(duration)) {
			break
		}
	}
	
	// Update min time
	for {
		currentMin := atomic.LoadInt64((*int64)(&hpl.metrics.MinTime))
		if currentMin == 0 || int64(duration) >= currentMin {
			break
		}
		if atomic.CompareAndSwapInt64((*int64)(&hpl.metrics.MinTime), currentMin, int64(duration)) {
			break
		}
	}
}

// Logging methods
func (hpl *HighPerformanceLogger) Trace(message string, fields ...map[string]interface{}) {
	hpl.log(TRACE, message, mergeFields(fields...))
}

func (hpl *HighPerformanceLogger) Debug(message string, fields ...map[string]interface{}) {
	hpl.log(DEBUG, message, mergeFields(fields...))
}

func (hpl *HighPerformanceLogger) Info(message string, fields ...map[string]interface{}) {
	hpl.log(INFO, message, mergeFields(fields...))
}

func (hpl *HighPerformanceLogger) Warn(message string, fields ...map[string]interface{}) {
	hpl.log(WARN, message, mergeFields(fields...))
}

func (hpl *HighPerformanceLogger) Error(message string, fields ...map[string]interface{}) {
	hpl.log(ERROR, message, mergeFields(fields...))
}

func (hpl *HighPerformanceLogger) Fatal(message string, fields ...map[string]interface{}) {
	hpl.log(FATAL, message, mergeFields(fields...))
	os.Exit(1)
}

func mergeFields(fields ...map[string]interface{}) map[string]interface{} {
	if len(fields) == 0 {
		return nil
	}
	if len(fields) == 1 {
		return fields[0]
	}
	
	result := make(map[string]interface{})
	for _, fieldMap := range fields {
		for k, v := range fieldMap {
			result[k] = v
		}
	}
	return result
}

func (hpl *HighPerformanceLogger) GetMetrics() *PerformanceMetrics {
	return &PerformanceMetrics{
		TotalLogs:       atomic.LoadInt64(&hpl.metrics.TotalLogs),
		FilteredLogs:    atomic.LoadInt64(&hpl.metrics.FilteredLogs),
		TotalTime:       time.Duration(atomic.LoadInt64((*int64)(&hpl.metrics.TotalTime))),
		AverageTime:     time.Duration(atomic.LoadInt64((*int64)(&hpl.metrics.AverageTime))),
		MaxTime:         time.Duration(atomic.LoadInt64((*int64)(&hpl.metrics.MaxTime))),
		MinTime:         time.Duration(atomic.LoadInt64((*int64)(&hpl.metrics.MinTime))),
		MemoryAllocated: atomic.LoadInt64(&hpl.metrics.MemoryAllocated),
		MemoryAllocs:    atomic.LoadInt64(&hpl.metrics.MemoryAllocs),
		GCCollections:   atomic.LoadInt64(&hpl.metrics.GCCollections),
		CPUUsage:        atomic.LoadFloat64(&hpl.metrics.CPUUsage),
	}
}

// ============================================================================
// LAZY LOGGER
// ============================================================================

type LazyLogger struct {
	level  LogLevel
	output io.Writer
}

func NewLazyLogger(level LogLevel, output io.Writer) *LazyLogger {
	return &LazyLogger{
		level:  level,
		output: output,
	}
}

func (ll *LazyLogger) log(level LogLevel, message string, lazyFields func() map[string]interface{}) {
	// Quick level check first
	if level.Priority() < ll.level.Priority() {
		return
	}
	
	// Only call lazy function if level check passes
	fields := lazyFields()
	
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}
	
	jsonData, _ := json.Marshal(entry)
	fmt.Fprintln(ll.output, string(jsonData))
}

func (ll *LazyLogger) Trace(message string, lazyFields func() map[string]interface{}) {
	ll.log(TRACE, message, lazyFields)
}

func (ll *LazyLogger) Debug(message string, lazyFields func() map[string]interface{}) {
	ll.log(DEBUG, message, lazyFields)
}

func (ll *LazyLogger) Info(message string, lazyFields func() map[string]interface{}) {
	ll.log(INFO, message, lazyFields)
}

func (ll *LazyLogger) Warn(message string, lazyFields func() map[string]interface{}) {
	ll.log(WARN, message, lazyFields)
}

func (ll *LazyLogger) Error(message string, lazyFields func() map[string]interface{}) {
	ll.log(ERROR, message, lazyFields)
}

// ============================================================================
// PERFORMANCE PROFILER
// ============================================================================

type PerformanceProfiler struct {
	startTime    time.Time
	startMem     runtime.MemStats
	endTime      time.Time
	endMem       runtime.MemStats
	measurements []Measurement
	mu           sync.Mutex
}

type Measurement struct {
	Name      string        `json:"name"`
	Duration  time.Duration `json:"duration"`
	Memory    int64         `json:"memory"`
	Allocs    int64         `json:"allocs"`
	Timestamp time.Time     `json:"timestamp"`
}

func NewPerformanceProfiler() *PerformanceProfiler {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	
	return &PerformanceProfiler{
		startTime:    time.Now(),
		startMem:     memStats,
		measurements: make([]Measurement, 0),
	}
}

func (pp *PerformanceProfiler) Start() {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.startTime = time.Now()
	runtime.ReadMemStats(&pp.startMem)
}

func (pp *PerformanceProfiler) End() {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.endTime = time.Now()
	runtime.ReadMemStats(&pp.endMem)
}

func (pp *PerformanceProfiler) Measure(name string, fn func()) {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	startTime := time.Now()
	
	fn()
	
	var endMem runtime.MemStats
	runtime.ReadMemStats(&endMem)
	endTime := time.Now()
	
	measurement := Measurement{
		Name:      name,
		Duration:  endTime.Sub(startTime),
		Memory:    int64(endMem.Alloc - startMem.Alloc),
		Allocs:    int64(endMem.Mallocs - startMem.Mallocs),
		Timestamp: startTime,
	}
	
	pp.measurements = append(pp.measurements, measurement)
}

func (pp *PerformanceProfiler) GetResults() []Measurement {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	return pp.measurements
}

func (pp *PerformanceProfiler) GetTotalDuration() time.Duration {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	return pp.endTime.Sub(pp.startTime)
}

func (pp *PerformanceProfiler) GetTotalMemory() int64 {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	return int64(pp.endMem.Alloc - pp.startMem.Alloc)
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstratePerformanceMeasurement() {
	fmt.Println("\n=== Performance Measurement ===")
	
	// Create profiler
	profiler := NewPerformanceProfiler()
	profiler.Start()
	
	// Test different logging approaches
	profiler.Measure("sync_logging", func() {
		logger := NewHighPerformanceLogger(INFO, os.Stdout, false)
		for i := 0; i < 1000; i++ {
			logger.Info("Sync log message", map[string]interface{}{
				"iteration": i,
				"timestamp": time.Now().UnixNano(),
			})
		}
	})
	
	profiler.Measure("async_logging", func() {
		logger := NewHighPerformanceLogger(INFO, os.Stdout, true)
		for i := 0; i < 1000; i++ {
			logger.Info("Async log message", map[string]interface{}{
				"iteration": i,
				"timestamp": time.Now().UnixNano(),
			})
		}
		logger.Stop()
	})
	
	profiler.End()
	
	// Display results
	results := profiler.GetResults()
	for _, measurement := range results {
		fmt.Printf("   📊 %s: %v, Memory: %d bytes, Allocs: %d\n", 
			measurement.Name, measurement.Duration, measurement.Memory, measurement.Allocs)
	}
	
	fmt.Printf("   📊 Total Duration: %v\n", profiler.GetTotalDuration())
	fmt.Printf("   📊 Total Memory: %d bytes\n", profiler.GetTotalMemory())
}

func demonstrateLazyEvaluation() {
	fmt.Println("\n=== Lazy Evaluation ===")
	
	logger := NewLazyLogger(INFO, os.Stdout)
	
	// Expensive operation that should only run if INFO level is enabled
	expensiveOperation := func() map[string]interface{} {
		fmt.Println("   🔍 Expensive operation executed!")
		time.Sleep(10 * time.Millisecond) // Simulate expensive operation
		return map[string]interface{}{
			"expensive_data": "This took a long time to compute",
			"computation_time": "10ms",
		}
	}
	
	// This will execute expensive operation
	logger.Info("Info with expensive data", expensiveOperation)
	
	// Change level to ERROR - expensive operation won't run
	logger.level = ERROR
	logger.Info("This won't execute expensive operation", expensiveOperation)
	
	fmt.Println("   📊 Lazy evaluation prevents unnecessary computation")
}

func demonstrateMemoryOptimization() {
	fmt.Println("\n=== Memory Optimization ===")
	
	// Test with object pooling
	pool := NewLogEntryPool()
	
	// Measure memory usage
	var memStats1, memStats2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStats1)
	
	// Create many log entries without pooling
	entries1 := make([]LogEntry, 1000)
	for i := 0; i < 1000; i++ {
		entries1[i] = LogEntry{
			Timestamp: time.Now(),
			Level:     INFO,
			Message:   "Test message",
			Fields:    map[string]interface{}{"id": i},
		}
	}
	
	runtime.GC()
	runtime.ReadMemStats(&memStats2)
	
	fmt.Printf("   📊 Memory without pooling: %d bytes\n", memStats2.Alloc-memStats1.Alloc)
	
	// Test with object pooling
	runtime.GC()
	runtime.ReadMemStats(&memStats1)
	
	entries2 := make([]*LogEntry, 1000)
	for i := 0; i < 1000; i++ {
		entry := pool.Get()
		entry.Timestamp = time.Now()
		entry.Level = INFO
		entry.Message = "Test message"
		entry.Fields["id"] = i
		entries2[i] = entry
	}
	
	// Return entries to pool
	for _, entry := range entries2 {
		pool.Put(entry)
	}
	
	runtime.GC()
	runtime.ReadMemStats(&memStats2)
	
	fmt.Printf("   📊 Memory with pooling: %d bytes\n", memStats2.Alloc-memStats1.Alloc)
}

func demonstrateHighThroughputLogging() {
	fmt.Println("\n=== High Throughput Logging ===")
	
	// Create async logger
	logger := NewHighPerformanceLogger(INFO, os.Stdout, true)
	defer logger.Stop()
	
	// Measure throughput
	start := time.Now()
	numLogs := 10000
	
	// Send logs concurrently
	var wg sync.WaitGroup
	numGoroutines := 10
	logsPerGoroutine := numLogs / numGoroutines
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < logsPerGoroutine; j++ {
				logger.Info("High throughput log", map[string]interface{}{
					"goroutine": goroutineID,
					"iteration": j,
					"timestamp": time.Now().UnixNano(),
				})
			}
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	// Get metrics
	metrics := logger.GetMetrics()
	
	fmt.Printf("   📊 Logs per second: %.0f\n", float64(numLogs)/duration.Seconds())
	fmt.Printf("   📊 Total logs: %d\n", metrics.TotalLogs)
	fmt.Printf("   📊 Filtered logs: %d\n", metrics.FilteredLogs)
	fmt.Printf("   📊 Average time per log: %v\n", metrics.AverageTime)
}

func demonstrateProfiling() {
	fmt.Println("\n=== Profiling ===")
	
	// Create CPU profile
	cpuFile, _ := os.Create("cpu.prof")
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()
	
	// Create memory profile
	memFile, _ := os.Create("mem.prof")
	defer memFile.Close()
	defer pprof.WriteHeapProfile(memFile)
	
	// Run intensive logging
	logger := NewHighPerformanceLogger(INFO, os.Stdout, true)
	defer logger.Stop()
	
	for i := 0; i < 5000; i++ {
		logger.Info("Profiling log", map[string]interface{}{
			"iteration": i,
			"data":      "This is some data for profiling",
		})
	}
	
	fmt.Println("   📊 CPU and memory profiles created")
	fmt.Println("   📊 Run 'go tool pprof cpu.prof' to analyze CPU profile")
	fmt.Println("   📊 Run 'go tool pprof mem.prof' to analyze memory profile")
}

func demonstrateOptimizationTechniques() {
	fmt.Println("\n=== Optimization Techniques ===")
	
	// Test different optimization strategies
	strategies := []struct {
		name string
		fn   func()
	}{
		{
			name: "Level Check First",
			fn: func() {
				logger := NewHighPerformanceLogger(ERROR, os.Stdout, false)
				for i := 0; i < 1000; i++ {
					// This will be filtered out quickly
					logger.Debug("Debug message", map[string]interface{}{"id": i})
				}
			},
		},
		{
			name: "Minimal Field Creation",
			fn: func() {
				logger := NewHighPerformanceLogger(INFO, os.Stdout, false)
				for i := 0; i < 1000; i++ {
					// Minimal field creation
					logger.Info("Info message", map[string]interface{}{"id": i})
				}
			},
		},
		{
			name: "Async Processing",
			fn: func() {
				logger := NewHighPerformanceLogger(INFO, os.Stdout, true)
				for i := 0; i < 1000; i++ {
					logger.Info("Async message", map[string]interface{}{"id": i})
				}
				logger.Stop()
			},
		},
	}
	
	profiler := NewPerformanceProfiler()
	profiler.Start()
	
	for _, strategy := range strategies {
		profiler.Measure(strategy.name, strategy.fn)
	}
	
	profiler.End()
	
	results := profiler.GetResults()
	for _, measurement := range results {
		fmt.Printf("   📊 %s: %v\n", measurement.Name, measurement.Duration)
	}
}

func main() {
	fmt.Println("⚡ PERFORMANCE IMPACT MASTERY")
	fmt.Println("=============================")
	
	demonstratePerformanceMeasurement()
	demonstrateLazyEvaluation()
	demonstrateMemoryOptimization()
	demonstrateHighThroughputLogging()
	demonstrateProfiling()
	demonstrateOptimizationTechniques()
	
	fmt.Println("\n🎉 PERFORMANCE IMPACT MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Performance measurement and profiling")
	fmt.Println("✅ Lazy evaluation patterns")
	fmt.Println("✅ Memory optimization techniques")
	fmt.Println("✅ High-throughput logging strategies")
	fmt.Println("✅ CPU and memory profiling")
	fmt.Println("✅ Various optimization techniques")
	
	fmt.Println("\n🚀 You are now ready for Tracing and Metrics Mastery!")
}
