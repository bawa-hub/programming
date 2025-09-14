// 📊 PROMETHEUS METRICS MASTERY
// Advanced metrics collection and monitoring with Prometheus
package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// ============================================================================
// METRIC TYPES
// ============================================================================

type MetricType string

const (
	CounterType   MetricType = "counter"
	GaugeType     MetricType = "gauge"
	HistogramType MetricType = "histogram"
	SummaryType   MetricType = "summary"
)

type Metric struct {
	Name        string            `json:"name"`
	Help        string            `json:"help"`
	Type        MetricType        `json:"type"`
	Labels      map[string]string `json:"labels"`
	Value       float64           `json:"value"`
	Timestamp   time.Time         `json:"timestamp"`
	Buckets     []float64         `json:"buckets,omitempty"`     // For histogram
	Quantiles   []float64         `json:"quantiles,omitempty"`   // For summary
	Count       uint64            `json:"count,omitempty"`       // For histogram/summary
	Sum         float64           `json:"sum,omitempty"`         // For histogram/summary
}

type MetricFamily struct {
	Name    string   `json:"name"`
	Help    string   `json:"help"`
	Type    MetricType `json:"type"`
	Metrics []Metric `json:"metrics"`
}

// ============================================================================
// METRIC REGISTRY
// ============================================================================

type MetricRegistry struct {
	metrics map[string]*MetricFamily
	mu      sync.RWMutex
}

func NewMetricRegistry() *MetricRegistry {
	return &MetricRegistry{
		metrics: make(map[string]*MetricFamily),
	}
}

func (mr *MetricRegistry) RegisterCounter(name, help string) *Counter {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	counter := &Counter{
		name:   name,
		help:   help,
		value:  0,
		labels: make(map[string]string),
	}
	
	mr.metrics[name] = &MetricFamily{
		Name:    name,
		Help:    help,
		Type:    CounterType,
		Metrics: []Metric{},
	}
	
	return counter
}

func (mr *MetricRegistry) RegisterGauge(name, help string) *Gauge {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	gauge := &Gauge{
		name:   name,
		help:   help,
		value:  0,
		labels: make(map[string]string),
	}
	
	mr.metrics[name] = &MetricFamily{
		Name:    name,
		Help:    help,
		Type:    GaugeType,
		Metrics: []Metric{},
	}
	
	return gauge
}

func (mr *MetricRegistry) RegisterHistogram(name, help string, buckets []float64) *Histogram {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	if buckets == nil {
		buckets = []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}
	}
	
	histogram := &Histogram{
		name:    name,
		help:    help,
		buckets: buckets,
		counts:  make([]uint64, len(buckets)+1), // +1 for +Inf bucket
		sum:     0,
		labels:  make(map[string]string),
	}
	
	mr.metrics[name] = &MetricFamily{
		Name:    name,
		Help:    help,
		Type:    HistogramType,
		Metrics: []Metric{},
	}
	
	return histogram
}

func (mr *MetricRegistry) RegisterSummary(name, help string, quantiles []float64) *Summary {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	
	if quantiles == nil {
		quantiles = []float64{0.5, 0.9, 0.95, 0.99}
	}
	
	summary := &Summary{
		name:      name,
		help:      help,
		quantiles: quantiles,
		values:    make([]float64, 0),
		sum:       0,
		count:     0,
		labels:    make(map[string]string),
		mu:        sync.Mutex{},
	}
	
	mr.metrics[name] = &MetricFamily{
		Name:    name,
		Help:    help,
		Type:    SummaryType,
		Metrics: []Metric{},
	}
	
	return summary
}

func (mr *MetricRegistry) Collect() map[string]*MetricFamily {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	
	// Create a copy of the metrics
	result := make(map[string]*MetricFamily)
	for name, family := range mr.metrics {
		result[name] = &MetricFamily{
			Name:    family.Name,
			Help:    family.Help,
			Type:    family.Type,
			Metrics: make([]Metric, len(family.Metrics)),
		}
		copy(result[name].Metrics, family.Metrics)
	}
	
	return result
}

func (mr *MetricRegistry) ExportPrometheusFormat() string {
	families := mr.Collect()
	
	var result string
	for _, family := range families {
		result += fmt.Sprintf("# HELP %s %s\n", family.Name, family.Help)
		result += fmt.Sprintf("# TYPE %s %s\n", family.Name, family.Type)
		
		for _, metric := range family.Metrics {
			labels := ""
			if len(metric.Labels) > 0 {
				var labelPairs []string
				for k, v := range metric.Labels {
					labelPairs = append(labelPairs, fmt.Sprintf("%s=\"%s\"", k, v))
				}
				sort.Strings(labelPairs)
				labels = "{" + fmt.Sprintf("%s", labelPairs) + "}"
			}
			
			if family.Type == HistogramType {
				// Export histogram buckets
			for _, bucket := range metric.Buckets {
				result += fmt.Sprintf("%s_bucket%s %d\n", family.Name, labels, bucket, metric.Count)
			}
				result += fmt.Sprintf("%s_bucket%s +Inf %d\n", family.Name, labels, metric.Count)
				result += fmt.Sprintf("%s_sum%s %f\n", family.Name, labels, metric.Sum)
				result += fmt.Sprintf("%s_count%s %d\n", family.Name, labels, metric.Count)
			} else if family.Type == SummaryType {
				// Export summary quantiles
			for _, quantile := range metric.Quantiles {
				result += fmt.Sprintf("%s%s %f\n", family.Name, labels, quantile, metric.Value)
			}
				result += fmt.Sprintf("%s_sum%s %f\n", family.Name, labels, metric.Sum)
				result += fmt.Sprintf("%s_count%s %d\n", family.Name, labels, metric.Count)
			} else {
				result += fmt.Sprintf("%s%s %f\n", family.Name, labels, metric.Value)
			}
		}
		result += "\n"
	}
	
	return result
}

// ============================================================================
// COUNTER IMPLEMENTATION
// ============================================================================

type Counter struct {
	name   string
	help   string
	value  uint64
	labels map[string]string
	mu     sync.RWMutex
}

func (c *Counter) Inc() {
	atomic.AddUint64(&c.value, 1)
}

func (c *Counter) Add(delta float64) {
	atomic.AddUint64(&c.value, uint64(delta))
}

func (c *Counter) Get() float64 {
	return float64(atomic.LoadUint64(&c.value))
}

func (c *Counter) WithLabels(labels map[string]string) *Counter {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	newCounter := &Counter{
		name:   c.name,
		help:   c.help,
		value:  c.value,
		labels: make(map[string]string),
	}
	
	// Copy existing labels
	for k, v := range c.labels {
		newCounter.labels[k] = v
	}
	
	// Add new labels
	for k, v := range labels {
		newCounter.labels[k] = v
	}
	
	return newCounter
}

func (c *Counter) ToMetric() Metric {
	return Metric{
		Name:      c.name,
		Type:      CounterType,
		Labels:    c.labels,
		Value:     c.Get(),
		Timestamp: time.Now(),
	}
}

// ============================================================================
// GAUGE IMPLEMENTATION
// ============================================================================

type Gauge struct {
	name   string
	help   string
	value  float64
	labels map[string]string
	mu     sync.RWMutex
}

func (g *Gauge) Set(value float64) {
	atomic.StoreUint64((*uint64)(unsafe.Pointer(&g.value)), uint64(value))
}

func (g *Gauge) Inc() {
	atomic.AddUint64((*uint64)(unsafe.Pointer(&g.value)), 1)
}

func (g *Gauge) Dec() {
	atomic.AddUint64((*uint64)(unsafe.Pointer(&g.value)), ^uint64(0))
}

func (g *Gauge) Add(delta float64) {
	atomic.AddUint64((*uint64)(unsafe.Pointer(&g.value)), uint64(delta))
}

func (g *Gauge) Get() float64 {
	return float64(atomic.LoadUint64((*uint64)(unsafe.Pointer(&g.value))))
}

func (g *Gauge) WithLabels(labels map[string]string) *Gauge {
	g.mu.Lock()
	defer g.mu.Unlock()
	
	newGauge := &Gauge{
		name:   g.name,
		help:   g.help,
		value:  g.value,
		labels: make(map[string]string),
	}
	
	// Copy existing labels
	for k, v := range g.labels {
		newGauge.labels[k] = v
	}
	
	// Add new labels
	for k, v := range labels {
		newGauge.labels[k] = v
	}
	
	return newGauge
}

func (g *Gauge) ToMetric() Metric {
	return Metric{
		Name:      g.name,
		Type:      GaugeType,
		Labels:    g.labels,
		Value:     g.Get(),
		Timestamp: time.Now(),
	}
}

// ============================================================================
// HISTOGRAM IMPLEMENTATION
// ============================================================================

type Histogram struct {
	name    string
	help    string
	buckets []float64
	counts  []uint64
	sum     float64
	labels  map[string]string
	mu      sync.RWMutex
}

func (h *Histogram) Observe(value float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	// Find the appropriate bucket
	bucketIndex := len(h.buckets) // Default to +Inf bucket
	for i, bucket := range h.buckets {
		if value <= bucket {
			bucketIndex = i
			break
		}
	}
	
	atomic.AddUint64(&h.counts[bucketIndex], 1)
	atomic.AddUint64((*uint64)(unsafe.Pointer(&h.sum)), uint64(value))
}

func (h *Histogram) GetCount() uint64 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	var total uint64
	for _, count := range h.counts {
		total += atomic.LoadUint64(&count)
	}
	return total
}

func (h *Histogram) GetSum() float64 {
	return float64(atomic.LoadUint64((*uint64)(unsafe.Pointer(&h.sum))))
}

func (h *Histogram) WithLabels(labels map[string]string) *Histogram {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	newHistogram := &Histogram{
		name:    h.name,
		help:    h.help,
		buckets: h.buckets,
		counts:  make([]uint64, len(h.counts)),
		sum:     h.sum,
		labels:  make(map[string]string),
	}
	
	// Copy existing labels
	for k, v := range h.labels {
		newHistogram.labels[k] = v
	}
	
	// Add new labels
	for k, v := range labels {
		newHistogram.labels[k] = v
	}
	
	return newHistogram
}

func (h *Histogram) ToMetric() Metric {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	// Copy current counts
	counts := make([]uint64, len(h.counts))
	for i, count := range h.counts {
		counts[i] = atomic.LoadUint64(&count)
	}
	
	return Metric{
		Name:      h.name,
		Type:      HistogramType,
		Labels:    h.labels,
		Buckets:   h.buckets,
		Count:     h.GetCount(),
		Sum:       h.GetSum(),
		Timestamp: time.Now(),
	}
}

// ============================================================================
// SUMMARY IMPLEMENTATION
// ============================================================================

type Summary struct {
	name      string
	help      string
	quantiles []float64
	values    []float64
	sum       float64
	count     uint64
	labels    map[string]string
	mu        sync.Mutex
}

func (s *Summary) Observe(value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.values = append(s.values, value)
	s.sum += value
	s.count++
}

func (s *Summary) GetQuantile(quantile float64) float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if len(s.values) == 0 {
		return 0
	}
	
	// Sort values
	sortedValues := make([]float64, len(s.values))
	copy(sortedValues, s.values)
	sort.Float64s(sortedValues)
	
	// Calculate quantile
	index := quantile * float64(len(sortedValues)-1)
	if index < 0 {
		index = 0
	}
	if index >= float64(len(sortedValues)) {
		index = float64(len(sortedValues) - 1)
	}
	
	lower := int(index)
	upper := lower + 1
	
	if upper >= len(sortedValues) {
		return sortedValues[lower]
	}
	
	weight := index - float64(lower)
	return sortedValues[lower]*(1-weight) + sortedValues[upper]*weight
}

func (s *Summary) GetSum() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sum
}

func (s *Summary) GetCount() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func (s *Summary) WithLabels(labels map[string]string) *Summary {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	newSummary := &Summary{
		name:      s.name,
		help:      s.help,
		quantiles: s.quantiles,
		values:    make([]float64, len(s.values)),
		sum:       s.sum,
		count:     s.count,
		labels:    make(map[string]string),
	}
	
	copy(newSummary.values, s.values)
	
	// Copy existing labels
	for k, v := range s.labels {
		newSummary.labels[k] = v
	}
	
	// Add new labels
	for k, v := range labels {
		newSummary.labels[k] = v
	}
	
	return newSummary
}

func (s *Summary) ToMetric() Metric {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	return Metric{
		Name:      s.name,
		Type:      SummaryType,
		Labels:    s.labels,
		Quantiles: s.quantiles,
		Count:     s.count,
		Sum:       s.sum,
		Timestamp: time.Now(),
	}
}

// ============================================================================
// METRICS SERVER
// ============================================================================

type MetricsServer struct {
	registry *MetricRegistry
	server   *http.Server
}

func NewMetricsServer(addr string, registry *MetricRegistry) *MetricsServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, registry.ExportPrometheusFormat())
	})
	
	return &MetricsServer{
		registry: registry,
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (ms *MetricsServer) Start() error {
	return ms.server.ListenAndServe()
}

func (ms *MetricsServer) Stop() error {
	return ms.server.Close()
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateCounters() {
	fmt.Println("\n=== Counters ===")
	
	registry := NewMetricRegistry()
	
	// Create counters
	requestCounter := registry.RegisterCounter("http_requests_total", "Total number of HTTP requests")
	errorCounter := registry.RegisterCounter("http_errors_total", "Total number of HTTP errors")
	
	// Simulate some requests
	for i := 0; i < 100; i++ {
		requestCounter.Inc()
		if i%10 == 0 {
			errorCounter.Inc()
		}
	}
	
	// Create labeled counters
	statusCounter := requestCounter.WithLabels(map[string]string{"status": "200"})
	statusCounter.Add(50)
	
	statusCounter404 := requestCounter.WithLabels(map[string]string{"status": "404"})
	statusCounter404.Add(10)
	
	fmt.Printf("   📊 Total requests: %.0f\n", requestCounter.Get())
	fmt.Printf("   📊 Total errors: %.0f\n", errorCounter.Get())
	fmt.Printf("   📊 200 responses: %.0f\n", statusCounter.Get())
	fmt.Printf("   📊 404 responses: %.0f\n", statusCounter404.Get())
}

func demonstrateGauges() {
	fmt.Println("\n=== Gauges ===")
	
	registry := NewMetricRegistry()
	
	// Create gauges
	memoryGauge := registry.RegisterGauge("memory_usage_bytes", "Current memory usage in bytes")
	cpuGauge := registry.RegisterGauge("cpu_usage_percent", "Current CPU usage percentage")
	
	// Simulate memory usage
	memoryGauge.Set(1024 * 1024 * 100) // 100MB
	cpuGauge.Set(75.5)
	
	// Simulate changes
	memoryGauge.Add(1024 * 1024 * 10) // Add 10MB
	cpuGauge.Dec() // Decrease CPU usage
	
	fmt.Printf("   📊 Memory usage: %.0f bytes\n", memoryGauge.Get())
	fmt.Printf("   📊 CPU usage: %.1f%%\n", cpuGauge.Get())
}

func demonstrateHistograms() {
	fmt.Println("\n=== Histograms ===")
	
	registry := NewMetricRegistry()
	
	// Create histogram
	responseTimeHist := registry.RegisterHistogram("http_request_duration_seconds", "HTTP request duration in seconds", []float64{0.1, 0.5, 1.0, 2.0, 5.0})
	
	// Simulate response times
	for i := 0; i < 1000; i++ {
		// Generate random response time (0-2 seconds)
		responseTime := rand.Float64() * 2
		responseTimeHist.Observe(responseTime)
	}
	
	metric := responseTimeHist.ToMetric()
	fmt.Printf("   📊 Total requests: %d\n", metric.Count)
	fmt.Printf("   📊 Total duration: %.2f seconds\n", metric.Sum)
	fmt.Printf("   📊 Average duration: %.3f seconds\n", metric.Sum/float64(metric.Count))
}

func demonstrateSummaries() {
	fmt.Println("\n=== Summaries ===")
	
	registry := NewMetricRegistry()
	
	// Create summary
	responseTimeSummary := registry.RegisterSummary("http_request_duration_seconds", "HTTP request duration in seconds", []float64{0.5, 0.9, 0.99})
	
	// Simulate response times
	for i := 0; i < 1000; i++ {
		// Generate random response time (0-2 seconds)
		responseTime := rand.Float64() * 2
		responseTimeSummary.Observe(responseTime)
	}
	
	metric := responseTimeSummary.ToMetric()
	fmt.Printf("   📊 Total requests: %d\n", metric.Count)
	fmt.Printf("   📊 Total duration: %.2f seconds\n", metric.Sum)
	fmt.Printf("   📊 50th percentile: %.3f seconds\n", responseTimeSummary.GetQuantile(0.5))
	fmt.Printf("   📊 95th percentile: %.3f seconds\n", responseTimeSummary.GetQuantile(0.95))
	fmt.Printf("   📊 99th percentile: %.3f seconds\n", responseTimeSummary.GetQuantile(0.99))
}

func demonstratePrometheusExport() {
	fmt.Println("\n=== Prometheus Export ===")
	
	registry := NewMetricRegistry()
	
	// Create various metrics
	requestCounter := registry.RegisterCounter("http_requests_total", "Total HTTP requests")
	requestCounter.Inc()
	requestCounter.Inc()
	
	memoryGauge := registry.RegisterGauge("memory_usage_bytes", "Memory usage")
	memoryGauge.Set(1024 * 1024 * 50) // 50MB
	
	responseTimeHist := registry.RegisterHistogram("http_request_duration_seconds", "Request duration", []float64{0.1, 0.5, 1.0, 2.0, 5.0})
	responseTimeHist.Observe(0.1)
	responseTimeHist.Observe(0.2)
	responseTimeHist.Observe(0.15)
	
	// Export in Prometheus format
	prometheusFormat := registry.ExportPrometheusFormat()
	fmt.Println("   📊 Prometheus format:")
	fmt.Println(prometheusFormat)
}

func demonstrateMetricsServer() {
	fmt.Println("\n=== Metrics Server ===")
	
	registry := NewMetricRegistry()
	
	// Create some metrics
	requestCounter := registry.RegisterCounter("http_requests_total", "Total HTTP requests")
	requestCounter.Inc()
	requestCounter.Inc()
	
	memoryGauge := registry.RegisterGauge("memory_usage_bytes", "Memory usage")
	memoryGauge.Set(1024 * 1024 * 100) // 100MB
	
	// Create metrics server
	server := NewMetricsServer(":8080", registry)
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Metrics server error: %v", err)
		}
	}()
	
	// Wait a bit for server to start
	time.Sleep(100 * time.Millisecond)
	
	// Make request to metrics endpoint
	resp, err := http.Get("http://localhost:8080/metrics")
	if err != nil {
		fmt.Printf("   ❌ Error fetching metrics: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("   ❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("   📊 Metrics server response:\n%s\n", string(body))
	
	// Stop server
	server.Stop()
}

func main() {
	fmt.Println("📊 PROMETHEUS METRICS MASTERY")
	fmt.Println("=============================")
	
	demonstrateCounters()
	demonstrateGauges()
	demonstrateHistograms()
	demonstrateSummaries()
	demonstratePrometheusExport()
	demonstrateMetricsServer()
	
	fmt.Println("\n🎉 PROMETHEUS METRICS MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Counter metrics for counting events")
	fmt.Println("✅ Gauge metrics for current values")
	fmt.Println("✅ Histogram metrics for distributions")
	fmt.Println("✅ Summary metrics for quantiles")
	fmt.Println("✅ Prometheus format export")
	fmt.Println("✅ HTTP metrics server")
	
	fmt.Println("\n🚀 You are now ready for Health Checks Mastery!")
}
