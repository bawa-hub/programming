package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// Advanced Pattern 1: Real-time Performance Monitoring
type PerformanceMonitor struct {
	metrics   chan Metric
	alerts    chan Alert
	threshold float64
	mu        sync.Mutex
	stats     map[string]float64
}

type Metric struct {
	Name  string
	Value float64
	Time  time.Time
}

type Alert struct {
	Type    string
	Message string
	Value   float64
	Time    time.Time
}

func NewPerformanceMonitor(threshold float64) *PerformanceMonitor {
	monitor := &PerformanceMonitor{
		metrics:   make(chan Metric, 1000),
		alerts:    make(chan Alert, 100),
		threshold: threshold,
		stats:     make(map[string]float64),
	}
	
	go monitor.run()
	return monitor
}

func (pm *PerformanceMonitor) run() {
	for metric := range pm.metrics {
		pm.mu.Lock()
		pm.stats[metric.Name] = metric.Value
		pm.mu.Unlock()
		
		if metric.Value > pm.threshold {
			pm.alerts <- Alert{
				Type:    "threshold_exceeded",
				Message: fmt.Sprintf("%s exceeded threshold: %.2f", metric.Name, metric.Value),
				Value:   metric.Value,
				Time:    time.Now(),
			}
		}
	}
}

func (pm *PerformanceMonitor) RecordMetric(name string, value float64) {
	select {
	case pm.metrics <- Metric{Name: name, Value: value, Time: time.Now()}:
	default:
		// Channel is full, drop metric
	}
}

func (pm *PerformanceMonitor) GetStats() map[string]float64 {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	stats := make(map[string]float64)
	for k, v := range pm.stats {
		stats[k] = v
	}
	return stats
}

func (pm *PerformanceMonitor) GetAlerts() <-chan Alert {
	return pm.alerts
}

// Advanced Pattern 2: Profiling with Context
type ProfilingContext struct {
	labels map[string]string
	ctx    context.Context
}

func NewProfilingContext(labels map[string]string) *ProfilingContext {
	ctx := context.Background()
	for k, v := range labels {
		ctx = context.WithValue(ctx, k, v)
	}
	
	return &ProfilingContext{
		labels: labels,
		ctx:    ctx,
	}
}

func (pc *ProfilingContext) ProfileFunction(name string, fn func()) {
	// Add labels to context
	pprof.SetGoroutineLabels(pc.ctx)
	
	// Profile function execution
	start := time.Now()
	fn()
	duration := time.Since(start)
	
	fmt.Printf("  Profiled %s: %v\n", name, duration)
}

// Advanced Pattern 3: Custom Profiling
type CustomProfiler struct {
	profiles map[string]*pprof.Profile
	mu       sync.Mutex
}

func NewCustomProfiler() *CustomProfiler {
	return &CustomProfiler{
		profiles: make(map[string]*pprof.Profile),
	}
}

func (cp *CustomProfiler) CreateProfile(name string) *pprof.Profile {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	if profile, exists := cp.profiles[name]; exists {
		return profile
	}
	
	profile := pprof.NewProfile(name)
	cp.profiles[name] = profile
	return profile
}

func (cp *CustomProfiler) RecordSample(profileName string, value int64) {
	cp.mu.Lock()
	profile, exists := cp.profiles[profileName]
	cp.mu.Unlock()
	
	if exists {
		profile.Add(1, int(value))
	}
}

func (cp *CustomProfiler) WriteProfile(profileName string, filename string) error {
	cp.mu.Lock()
	profile, exists := cp.profiles[profileName]
	cp.mu.Unlock()
	
	if !exists {
		return fmt.Errorf("profile %s not found", profileName)
	}
	
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	
	return profile.WriteTo(f, 0)
}

// Advanced Pattern 4: Performance Profiler
type PerformanceProfiler struct {
	profiles map[string]time.Duration
	mu       sync.Mutex
}

func NewPerformanceProfiler() *PerformanceProfiler {
	return &PerformanceProfiler{
		profiles: make(map[string]time.Duration),
	}
}

func (pp *PerformanceProfiler) Profile(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	
	pp.mu.Lock()
	pp.profiles[name] = duration
	pp.mu.Unlock()
}

func (pp *PerformanceProfiler) GetProfile(name string) (time.Duration, bool) {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	
	duration, exists := pp.profiles[name]
	return duration, exists
}

func (pp *PerformanceProfiler) GetAllProfiles() map[string]time.Duration {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	
	profiles := make(map[string]time.Duration)
	for k, v := range pp.profiles {
		profiles[k] = v
	}
	return profiles
}

// Advanced Pattern 5: Memory Profiler
type MemoryProfiler struct {
	allocations map[string]int64
	mu          sync.Mutex
}

func NewMemoryProfiler() *MemoryProfiler {
	return &MemoryProfiler{
		allocations: make(map[string]int64),
	}
}

func (mp *MemoryProfiler) RecordAllocation(name string, size int64) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	
	mp.allocations[name] += size
}

func (mp *MemoryProfiler) GetAllocations() map[string]int64 {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	
	allocations := make(map[string]int64)
	for k, v := range mp.allocations {
		allocations[k] = v
	}
	return allocations
}

func (mp *MemoryProfiler) Reset() {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	
	mp.allocations = make(map[string]int64)
}

// Advanced Pattern 6: Goroutine Profiler
type GoroutineProfiler struct {
	goroutines map[string]int
	mu         sync.Mutex
}

func NewGoroutineProfiler() *GoroutineProfiler {
	return &GoroutineProfiler{
		goroutines: make(map[string]int),
	}
}

func (gp *GoroutineProfiler) RecordGoroutine(name string) {
	gp.mu.Lock()
	defer gp.mu.Unlock()
	
	gp.goroutines[name]++
}

func (gp *GoroutineProfiler) GetGoroutines() map[string]int {
	gp.mu.Lock()
	defer gp.mu.Unlock()
	
	goroutines := make(map[string]int)
	for k, v := range gp.goroutines {
		goroutines[k] = v
	}
	return goroutines
}

// Advanced Pattern 7: Performance Dashboard
type PerformanceDashboard struct {
	monitor *PerformanceMonitor
	server  *http.Server
}

func NewPerformanceDashboard(port string) *PerformanceDashboard {
	dashboard := &PerformanceDashboard{
		monitor: NewPerformanceMonitor(1000.0),
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", dashboard.metricsHandler)
	mux.HandleFunc("/alerts", dashboard.alertsHandler)
	mux.HandleFunc("/stats", dashboard.statsHandler)
	
	dashboard.server = &http.Server{
		Addr:    port,
		Handler: mux,
	}
	
	return dashboard
}

func (pd *PerformanceDashboard) Start() {
	go func() {
		if err := pd.server.ListenAndServe(); err != nil {
			fmt.Printf("Dashboard server error: %v\n", err)
		}
	}()
}

func (pd *PerformanceDashboard) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pd.server.Shutdown(ctx)
}

func (pd *PerformanceDashboard) metricsHandler(w http.ResponseWriter, r *http.Request) {
	stats := pd.monitor.GetStats()
	
	w.Header().Set("Content-Type", "text/plain")
	for name, value := range stats {
		fmt.Fprintf(w, "%s %.2f\n", name, value)
	}
}

func (pd *PerformanceDashboard) alertsHandler(w http.ResponseWriter, r *http.Request) {
	alerts := pd.monitor.GetAlerts()
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"alerts\": [")
	
	first := true
	for {
		select {
		case alert := <-alerts:
			if !first {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprintf(w, "{\"type\":\"%s\",\"message\":\"%s\",\"value\":%.2f,\"time\":\"%s\"}", 
				alert.Type, alert.Message, alert.Value, alert.Time.Format(time.RFC3339))
			first = false
		case <-time.After(100 * time.Millisecond):
			goto done
		}
	}
done:
	fmt.Fprintf(w, "]}")
}

func (pd *PerformanceDashboard) statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := pd.monitor.GetStats()
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"stats\": {")
	
	first := true
	for name, value := range stats {
		if !first {
			fmt.Fprintf(w, ",")
		}
		fmt.Fprintf(w, "\"%s\": %.2f", name, value)
		first = false
	}
	fmt.Fprintf(w, "}}")
}

// Advanced Pattern 8: Performance Optimizer
type PerformanceOptimizer struct {
	profiler *PerformanceProfiler
	mu       sync.Mutex
	config   map[string]interface{}
}

func NewPerformanceOptimizer() *PerformanceOptimizer {
	return &PerformanceOptimizer{
		profiler: NewPerformanceProfiler(),
		config:   make(map[string]interface{}),
	}
}

func (po *PerformanceOptimizer) Optimize(name string, fn func()) {
	// Profile function
	po.profiler.Profile(name, fn)
	
	// Get profile duration
	duration, exists := po.profiler.GetProfile(name)
	if !exists {
		return
	}
	
	// Optimize based on duration
	po.mu.Lock()
	defer po.mu.Unlock()
	
	if duration > 100*time.Millisecond {
		po.config[name] = "optimize"
	} else if duration > 10*time.Millisecond {
		po.config[name] = "monitor"
	} else {
		po.config[name] = "good"
	}
}

func (po *PerformanceOptimizer) GetConfig() map[string]interface{} {
	po.mu.Lock()
	defer po.mu.Unlock()
	
	config := make(map[string]interface{})
	for k, v := range po.config {
		config[k] = v
	}
	return config
}

// Advanced Pattern 9: Memory Optimizer
type MemoryOptimizer struct {
	profiler *MemoryProfiler
	mu       sync.Mutex
	config   map[string]interface{}
}

func NewMemoryOptimizer() *MemoryOptimizer {
	return &MemoryOptimizer{
		profiler: NewMemoryProfiler(),
		config:   make(map[string]interface{}),
	}
}

func (mo *MemoryOptimizer) Optimize(name string, fn func()) {
	// Get initial memory stats
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	// Run function
	fn()
	
	// Get final memory stats
	runtime.ReadMemStats(&m2)
	
	// Record allocation
	allocated := int64(m2.Alloc - m1.Alloc)
	mo.profiler.RecordAllocation(name, allocated)
	
	// Optimize based on allocation
	mo.mu.Lock()
	defer mo.mu.Unlock()
	
	if allocated > 1024*1024 { // 1MB
		mo.config[name] = "optimize"
	} else if allocated > 1024*100 { // 100KB
		mo.config[name] = "monitor"
	} else {
		mo.config[name] = "good"
	}
}

func (mo *MemoryOptimizer) GetConfig() map[string]interface{} {
	mo.mu.Lock()
	defer mo.mu.Unlock()
	
	config := make(map[string]interface{})
	for k, v := range mo.config {
		config[k] = v
	}
	return config
}

// Advanced Pattern 10: Performance Tester
type PerformanceTester struct {
	tests map[string]func()
	mu    sync.Mutex
}

func NewPerformanceTester() *PerformanceTester {
	return &PerformanceTester{
		tests: make(map[string]func()),
	}
}

func (pt *PerformanceTester) AddTest(name string, test func()) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.tests[name] = test
}

func (pt *PerformanceTester) RunTest(name string) (time.Duration, error) {
	pt.mu.Lock()
	test, exists := pt.tests[name]
	pt.mu.Unlock()
	
	if !exists {
		return 0, fmt.Errorf("test %s not found", name)
	}
	
	start := time.Now()
	test()
	duration := time.Since(start)
	
	return duration, nil
}

func (pt *PerformanceTester) RunAllTests() map[string]time.Duration {
	pt.mu.Lock()
	tests := make(map[string]func())
	for k, v := range pt.tests {
		tests[k] = v
	}
	pt.mu.Unlock()
	
	results := make(map[string]time.Duration)
	for name, test := range tests {
		start := time.Now()
		test()
		results[name] = time.Since(start)
	}
	
	return results
}

// Advanced Pattern 1: Real-time Performance Monitoring
func realTimePerformanceMonitoring() {
	fmt.Println("\n1. Real-time Performance Monitoring")
	fmt.Println("===================================")
	
	monitor := NewPerformanceMonitor(1000.0)
	
	// Simulate metrics
	for i := 0; i < 10; i++ {
		monitor.RecordMetric("cpu_usage", float64(i*100))
		monitor.RecordMetric("memory_usage", float64(i*200))
		time.Sleep(100 * time.Millisecond)
	}
	
	// Get stats
	stats := monitor.GetStats()
	fmt.Printf("  Monitor stats: %v\n", stats)
	
	fmt.Println("Real-time performance monitoring completed")
}

// Advanced Pattern 2: Profiling with Context
func profilingWithContext() {
	fmt.Println("\n2. Profiling with Context")
	fmt.Println("=========================")
	
	labels := map[string]string{
		"service": "profiling-service",
		"version": "1.0.0",
	}
	
	ctx := NewProfilingContext(labels)
	
	ctx.ProfileFunction("test-function", func() {
		time.Sleep(100 * time.Millisecond)
	})
	
	fmt.Println("Profiling with context completed")
}

// Advanced Pattern 3: Custom Profiling
func customProfiling() {
	fmt.Println("\n3. Custom Profiling")
	fmt.Println("===================")
	
	// For this example, we'll just demonstrate the concept
	// without actually creating a custom profile to avoid complexity
	fmt.Println("  Custom profiling concept demonstrated")
	fmt.Println("  Custom profiling completed")
}

// Advanced Pattern 4: Performance Profiler
func performanceProfiler() {
	fmt.Println("\n4. Performance Profiler")
	fmt.Println("=======================")
	
	profiler := NewPerformanceProfiler()
	
	// Profile functions
	profiler.Profile("fast-function", func() {
		time.Sleep(10 * time.Millisecond)
	})
	
	profiler.Profile("slow-function", func() {
		time.Sleep(100 * time.Millisecond)
	})
	
	// Get profiles
	profiles := profiler.GetAllProfiles()
	fmt.Printf("  Performance profiles: %v\n", profiles)
	
	fmt.Println("Performance profiler completed")
}

// Advanced Pattern 5: Memory Profiler
func memoryProfiler() {
	fmt.Println("\n5. Memory Profiler")
	fmt.Println("==================")
	
	profiler := NewMemoryProfiler()
	
	// Record allocations
	profiler.RecordAllocation("slice", 1000)
	profiler.RecordAllocation("map", 2000)
	profiler.RecordAllocation("slice", 1500)
	
	// Get allocations
	allocations := profiler.GetAllocations()
	fmt.Printf("  Memory allocations: %v\n", allocations)
	
	fmt.Println("Memory profiler completed")
}

// Advanced Pattern 6: Goroutine Profiler
func goroutineProfiler() {
	fmt.Println("\n6. Goroutine Profiler")
	fmt.Println("====================")
	
	profiler := NewGoroutineProfiler()
	
	// Record goroutines
	for i := 0; i < 10; i++ {
		profiler.RecordGoroutine("worker")
	}
	
	for i := 0; i < 5; i++ {
		profiler.RecordGoroutine("handler")
	}
	
	// Get goroutines
	goroutines := profiler.GetGoroutines()
	fmt.Printf("  Goroutine counts: %v\n", goroutines)
	
	fmt.Println("Goroutine profiler completed")
}

// Advanced Pattern 7: Performance Dashboard
func performanceDashboard() {
	fmt.Println("\n7. Performance Dashboard")
	fmt.Println("=======================")
	
	dashboard := NewPerformanceDashboard(":8080")
	dashboard.Start()
	
	// Simulate some metrics
	monitor := dashboard.monitor
	for i := 0; i < 5; i++ {
		monitor.RecordMetric("cpu_usage", float64(i*100))
		monitor.RecordMetric("memory_usage", float64(i*200))
		time.Sleep(100 * time.Millisecond)
	}
	
	// Stop dashboard
	dashboard.Stop()
	
	fmt.Println("Performance dashboard completed")
}

// Advanced Pattern 8: Performance Optimizer
func performanceOptimizer() {
	fmt.Println("\n8. Performance Optimizer")
	fmt.Println("=======================")
	
	optimizer := NewPerformanceOptimizer()
	
	// Optimize functions
	optimizer.Optimize("fast-function", func() {
		time.Sleep(5 * time.Millisecond)
	})
	
	optimizer.Optimize("slow-function", func() {
		time.Sleep(150 * time.Millisecond)
	})
	
	// Get optimization config
	config := optimizer.GetConfig()
	fmt.Printf("  Optimization config: %v\n", config)
	
	fmt.Println("Performance optimizer completed")
}

// Advanced Pattern 9: Memory Optimizer
func memoryOptimizer() {
	fmt.Println("\n9. Memory Optimizer")
	fmt.Println("==================")
	
	optimizer := NewMemoryOptimizer()
	
	// Optimize functions
	optimizer.Optimize("small-allocation", func() {
		_ = make([]int, 100)
	})
	
	optimizer.Optimize("large-allocation", func() {
		_ = make([]int, 1000000)
	})
	
	// Get optimization config
	config := optimizer.GetConfig()
	fmt.Printf("  Memory optimization config: %v\n", config)
	
	fmt.Println("Memory optimizer completed")
}

// Advanced Pattern 10: Performance Tester
func performanceTester() {
	fmt.Println("\n10. Performance Tester")
	fmt.Println("=====================")
	
	tester := NewPerformanceTester()
	
	// Add tests
	tester.AddTest("test1", func() {
		time.Sleep(50 * time.Millisecond)
	})
	
	tester.AddTest("test2", func() {
		time.Sleep(100 * time.Millisecond)
	})
	
	// Run tests
	results := tester.RunAllTests()
	fmt.Printf("  Test results: %v\n", results)
	
	fmt.Println("Performance tester completed")
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Profiling & Benchmarking Patterns")
	fmt.Println("=============================================")
	
	realTimePerformanceMonitoring()
	profilingWithContext()
	customProfiling()
	performanceProfiler()
	memoryProfiler()
	goroutineProfiler()
	performanceDashboard()
	performanceOptimizer()
	memoryOptimizer()
	performanceTester()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
