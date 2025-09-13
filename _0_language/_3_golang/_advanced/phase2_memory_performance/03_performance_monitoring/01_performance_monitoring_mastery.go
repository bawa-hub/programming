package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync/atomic"
	"time"
)

// 📊 PERFORMANCE MONITORING MASTERY
// Understanding performance monitoring and metrics collection

func main() {
	fmt.Println("📊 PERFORMANCE MONITORING MASTERY")
	fmt.Println("==================================")
	fmt.Println()

	// 1. Real-time Monitoring
	realTimeMonitoring()
	fmt.Println()

	// 2. Metrics Collection
	metricsCollection()
	fmt.Println()

	// 3. Alerting and Notifications
	alertingAndNotifications()
	fmt.Println()

	// 4. Performance Analysis
	performanceAnalysis()
	fmt.Println()

	// 5. Production Monitoring
	productionMonitoring()
	fmt.Println()

	// 6. Custom Metrics
	customMetrics()
	fmt.Println()

	// 7. Health Checks
	healthChecks()
	fmt.Println()

	// 8. Performance Dashboards
	performanceDashboards()
	fmt.Println()

	// 9. Monitoring Best Practices
	monitoringBestPractices()
	fmt.Println()

	// 10. Advanced Monitoring
	advancedMonitoring()
}

// 1. Real-time Monitoring
func realTimeMonitoring() {
	fmt.Println("1. Real-time Monitoring:")
	fmt.Println("Understanding real-time performance monitoring...")

	// Demonstrate basic monitoring
	basicMonitoring()
	
	// Show system metrics
	systemMetrics()
	
	// Demonstrate application metrics
	applicationMetrics()
}

func basicMonitoring() {
	fmt.Println("  📊 Basic monitoring example:")
	
	// Monitor function execution time
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	duration := time.Since(start)
	
	fmt.Printf("    Function execution time: %v\n", duration)
	fmt.Printf("    Memory usage: %d KB\n", getMemoryUsage())
	fmt.Printf("    Goroutine count: %d\n", runtime.NumGoroutine())
}

func getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / 1024
}

func systemMetrics() {
	fmt.Println("  📊 System metrics:")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("    Allocated memory: %d KB\n", m.Alloc/1024)
	fmt.Printf("    Total allocated: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("    System memory: %d KB\n", m.Sys/1024)
	fmt.Printf("    GC cycles: %d\n", m.NumGC)
	fmt.Printf("    GC pause time: %v\n", time.Duration(m.PauseTotalNs))
}

func applicationMetrics() {
	fmt.Println("  📊 Application metrics:")
	
	// Simulate application metrics
	requestCount := int64(1000)
	errorCount := int64(50)
	responseTime := 150 * time.Millisecond
	
	fmt.Printf("    Request count: %d\n", requestCount)
	fmt.Printf("    Error count: %d\n", errorCount)
	fmt.Printf("    Error rate: %.2f%%\n", float64(errorCount)/float64(requestCount)*100)
	fmt.Printf("    Average response time: %v\n", responseTime)
	fmt.Printf("    Requests per second: %.2f\n", float64(requestCount)/responseTime.Seconds())
}

// 2. Metrics Collection
func metricsCollection() {
	fmt.Println("2. Metrics Collection:")
	fmt.Println("Understanding metrics collection and storage...")

	// Demonstrate counter metrics
	counterMetrics()
	
	// Show gauge metrics
	gaugeMetrics()
	
	// Demonstrate histogram metrics
	histogramMetrics()
	
	// Show summary metrics
	summaryMetrics()
}

func counterMetrics() {
	fmt.Println("  📊 Counter metrics:")
	
	// Simulate request counter
	requestCounter := int64(0)
	
	// Increment counter
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&requestCounter, 1)
	}
	
	fmt.Printf("    Request counter: %d\n", atomic.LoadInt64(&requestCounter))
	
	// Simulate error counter
	errorCounter := int64(0)
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&errorCounter, 1)
	}
	
	fmt.Printf("    Error counter: %d\n", atomic.LoadInt64(&errorCounter))
	fmt.Printf("    Error rate: %.2f%%\n", float64(atomic.LoadInt64(&errorCounter))/float64(atomic.LoadInt64(&requestCounter))*100)
}

func gaugeMetrics() {
	fmt.Println("  📊 Gauge metrics:")
	
	// Simulate active connections
	activeConnections := int64(42)
	fmt.Printf("    Active connections: %d\n", activeConnections)
	
	// Simulate memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("    Memory usage: %d KB\n", m.Alloc/1024)
	
	// Simulate CPU usage
	fmt.Printf("    CPU usage: %.2f%%\n", 45.67)
	
	// Simulate queue length
	queueLength := int64(15)
	fmt.Printf("    Queue length: %d\n", queueLength)
}

func histogramMetrics() {
	fmt.Println("  📊 Histogram metrics:")
	
	// Simulate response time histogram
	responseTimes := []time.Duration{
		50 * time.Millisecond,
		100 * time.Millisecond,
		150 * time.Millisecond,
		200 * time.Millisecond,
		250 * time.Millisecond,
	}
	
	fmt.Println("    Response time distribution:")
	for i, rt := range responseTimes {
		fmt.Printf("      Bucket %d: %v\n", i, rt)
	}
	
	// Calculate percentiles
	sort.Slice(responseTimes, func(i, j int) bool {
		return responseTimes[i] < responseTimes[j]
	})
	
	p50 := responseTimes[len(responseTimes)/2]
	p95 := responseTimes[int(float64(len(responseTimes))*0.95)]
	p99 := responseTimes[int(float64(len(responseTimes))*0.99)]
	
	fmt.Printf("    P50: %v\n", p50)
	fmt.Printf("    P95: %v\n", p95)
	fmt.Printf("    P99: %v\n", p99)
}

func summaryMetrics() {
	fmt.Println("  📊 Summary metrics:")
	
	// Simulate request duration summary
	durations := []time.Duration{
		10 * time.Millisecond,
		20 * time.Millisecond,
		30 * time.Millisecond,
		40 * time.Millisecond,
		50 * time.Millisecond,
	}
	
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	
	avg := total / time.Duration(len(durations))
	fmt.Printf("    Average duration: %v\n", avg)
	fmt.Printf("    Total requests: %d\n", len(durations))
	fmt.Printf("    Total duration: %v\n", total)
}

// 3. Alerting and Notifications
func alertingAndNotifications() {
	fmt.Println("3. Alerting and Notifications:")
	fmt.Println("Understanding alerting systems and notifications...")

	// Demonstrate alert rules
	alertRules()
	
	// Show notification channels
	notificationChannels()
	
	// Demonstrate alert escalation
	alertEscalation()
}

func alertRules() {
	fmt.Println("  📊 Alert rules:")
	
	// Simulate alert conditions
	errorRate := 5.5
	responseTime := 200 * time.Millisecond
	memoryUsage := 85.0
	
	fmt.Printf("    Error rate: %.2f%% (threshold: 5.0%%)\n", errorRate)
	if errorRate > 5.0 {
		fmt.Println("    🚨 ALERT: High error rate!")
	}
	
	fmt.Printf("    Response time: %v (threshold: 150ms)\n", responseTime)
	if responseTime > 150*time.Millisecond {
		fmt.Println("    🚨 ALERT: High response time!")
	}
	
	fmt.Printf("    Memory usage: %.2f%% (threshold: 80%%)\n", memoryUsage)
	if memoryUsage > 80.0 {
		fmt.Println("    🚨 ALERT: High memory usage!")
	}
}

func notificationChannels() {
	fmt.Println("  📊 Notification channels:")
	
	// Simulate different notification channels
	channels := []string{"email", "slack", "pagerduty", "webhook"}
	
	for _, channel := range channels {
		fmt.Printf("    Sending alert to %s...\n", channel)
		// Simulate sending notification
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("    ✅ Alert sent to %s\n", channel)
	}
}

func alertEscalation() {
	fmt.Println("  📊 Alert escalation:")
	
	// Simulate alert escalation levels
	alertLevels := []string{"INFO", "WARNING", "CRITICAL", "EMERGENCY"}
	
	for i, level := range alertLevels {
		fmt.Printf("    Level %d: %s\n", i+1, level)
		
		// Simulate escalation based on severity
		if level == "CRITICAL" {
			fmt.Println("      🚨 Escalating to on-call engineer")
		} else if level == "EMERGENCY" {
			fmt.Println("      🚨 Escalating to management")
		}
	}
}

// 4. Performance Analysis
func performanceAnalysis() {
	fmt.Println("4. Performance Analysis:")
	fmt.Println("Understanding performance analysis techniques...")

	// Demonstrate bottleneck identification
	bottleneckIdentification()
	
	// Show performance regression detection
	performanceRegressionDetection()
	
	// Demonstrate capacity planning
	capacityPlanning()
}

func bottleneckIdentification() {
	fmt.Println("  📊 Bottleneck identification:")
	
	// Simulate different system components
	components := map[string]time.Duration{
		"Database":     100 * time.Millisecond,
		"Cache":        10 * time.Millisecond,
		"API Gateway":  50 * time.Millisecond,
		"Business Logic": 200 * time.Millisecond,
		"External API": 150 * time.Millisecond,
	}
	
	var totalTime time.Duration
	for component, duration := range components {
		fmt.Printf("    %s: %v\n", component, duration)
		totalTime += duration
	}
	
	fmt.Printf("    Total time: %v\n", totalTime)
	
	// Identify bottleneck
	var bottleneck string
	var maxDuration time.Duration
	for component, duration := range components {
		if duration > maxDuration {
			maxDuration = duration
			bottleneck = component
		}
	}
	
	fmt.Printf("    🎯 Bottleneck: %s (%v, %.1f%% of total)\n", 
		bottleneck, maxDuration, float64(maxDuration)/float64(totalTime)*100)
}

func performanceRegressionDetection() {
	fmt.Println("  📊 Performance regression detection:")
	
	// Simulate performance over time
	baseline := 100 * time.Millisecond
	current := 150 * time.Millisecond
	
	fmt.Printf("    Baseline performance: %v\n", baseline)
	fmt.Printf("    Current performance: %v\n", current)
	
	regression := float64(current-baseline) / float64(baseline) * 100
	fmt.Printf("    Performance regression: %.1f%%\n", regression)
	
	if regression > 10 {
		fmt.Println("    🚨 ALERT: Significant performance regression detected!")
	} else if regression > 5 {
		fmt.Println("    ⚠️  WARNING: Minor performance regression detected")
	} else {
		fmt.Println("    ✅ Performance within acceptable range")
	}
}

func capacityPlanning() {
	fmt.Println("  📊 Capacity planning:")
	
	// Simulate capacity analysis
	currentLoad := 70.0
	growthRate := 20.0 // 20% per month
	months := 6
	
	fmt.Printf("    Current load: %.1f%%\n", currentLoad)
	fmt.Printf("    Growth rate: %.1f%% per month\n", growthRate)
	
	for month := 1; month <= months; month++ {
		futureLoad := currentLoad * (1 + growthRate/100*float64(month))
		fmt.Printf("    Month %d: %.1f%%\n", month, futureLoad)
		
		if futureLoad > 90 {
			fmt.Printf("      🚨 Capacity limit reached in month %d!\n", month)
			break
		}
	}
}

// 5. Production Monitoring
func productionMonitoring() {
	fmt.Println("5. Production Monitoring:")
	fmt.Println("Understanding production monitoring systems...")

	// Demonstrate health checks
	healthCheckMonitoring()
	
	// Show graceful degradation
	gracefulDegradation()
	
	// Demonstrate circuit breakers
	circuitBreakerMonitoring()
}

func healthCheckMonitoring() {
	fmt.Println("  📊 Health check monitoring:")
	
	// Simulate health check endpoints
	endpoints := []string{
		"/health",
		"/ready",
		"/live",
		"/metrics",
	}
	
	for _, endpoint := range endpoints {
		// Simulate health check
		healthy := rand.Float64() > 0.1 // 90% success rate
		status := "✅ Healthy"
		if !healthy {
			status = "❌ Unhealthy"
		}
		
		fmt.Printf("    %s: %s\n", endpoint, status)
	}
}

func gracefulDegradation() {
	fmt.Println("  📊 Graceful degradation:")
	
	// Simulate service degradation
	services := map[string]bool{
		"Database":      true,
		"Cache":         false, // Degraded
		"External API":  true,
		"File Storage":  false, // Degraded
	}
	
	fmt.Println("    Service status:")
	for service, available := range services {
		status := "✅ Available"
		if !available {
			status = "⚠️  Degraded"
		}
		fmt.Printf("      %s: %s\n", service, status)
	}
	
	// Simulate fallback behavior
	fmt.Println("    Fallback behavior:")
	fmt.Println("      Cache: Using database directly")
	fmt.Println("      File Storage: Using temporary storage")
}

func circuitBreakerMonitoring() {
	fmt.Println("  📊 Circuit breaker monitoring:")
	
	// Simulate circuit breaker states
	states := []string{"CLOSED", "OPEN", "HALF_OPEN"}
	
	for _, state := range states {
		fmt.Printf("    Circuit breaker state: %s\n", state)
		
		switch state {
		case "CLOSED":
			fmt.Println("      ✅ Normal operation")
		case "OPEN":
			fmt.Println("      🚨 Circuit open - requests blocked")
		case "HALF_OPEN":
			fmt.Println("      ⚠️  Testing if service recovered")
		}
	}
}

// 6. Custom Metrics
func customMetrics() {
	fmt.Println("6. Custom Metrics:")
	fmt.Println("Understanding custom metrics implementation...")

	// Demonstrate custom counter
	customCounter()
	
	// Show custom gauge
	customGauge()
	
	// Demonstrate custom histogram
	customHistogram()
}

func customCounter() {
	fmt.Println("  📊 Custom counter metrics:")
	
	// Simulate custom business metrics
	ordersProcessed := int64(0)
	revenue := int64(0)
	
	// Simulate order processing
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&ordersProcessed, 1)
		atomic.AddInt64(&revenue, int64(rand.Intn(100)+10))
	}
	
	fmt.Printf("    Orders processed: %d\n", atomic.LoadInt64(&ordersProcessed))
	fmt.Printf("    Total revenue: $%d\n", atomic.LoadInt64(&revenue))
	fmt.Printf("    Average order value: $%.2f\n", 
		float64(atomic.LoadInt64(&revenue))/float64(atomic.LoadInt64(&ordersProcessed)))
}

func customGauge() {
	fmt.Println("  📊 Custom gauge metrics:")
	
	// Simulate custom application gauges
	activeUsers := int64(42)
	queueSize := int64(15)
	cacheHitRate := 85.5
	
	fmt.Printf("    Active users: %d\n", atomic.LoadInt64(&activeUsers))
	fmt.Printf("    Queue size: %d\n", atomic.LoadInt64(&queueSize))
	fmt.Printf("    Cache hit rate: %.1f%%\n", cacheHitRate)
}

func customHistogram() {
	fmt.Println("  📊 Custom histogram metrics:")
	
	// Simulate custom business histograms
	orderValues := []int{10, 25, 50, 75, 100, 150, 200, 300, 500, 1000}
	
	fmt.Println("    Order value distribution:")
	for _, value := range orderValues {
		fmt.Printf("      $%d: %d orders\n", value, rand.Intn(100))
	}
}

// 7. Health Checks
func healthChecks() {
	fmt.Println("7. Health Checks:")
	fmt.Println("Understanding health check implementation...")

	// Demonstrate basic health check
	basicHealthCheck()
	
	// Show readiness check
	readinessCheck()
	
	// Demonstrate liveness check
	livenessCheck()
}

func basicHealthCheck() {
	fmt.Println("  📊 Basic health check:")
	
	// Simulate health check
	healthy := true
	message := "OK"
	
	// Check system resources
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	if m.Alloc > 100*1024*1024 { // 100MB
		healthy = false
		message = "High memory usage"
	}
	
	if runtime.NumGoroutine() > 1000 {
		healthy = false
		message = "Too many goroutines"
	}
	
	status := "✅ Healthy"
	if !healthy {
		status = "❌ Unhealthy"
	}
	
	fmt.Printf("    Status: %s\n", status)
	fmt.Printf("    Message: %s\n", message)
	fmt.Printf("    Memory: %d KB\n", m.Alloc/1024)
	fmt.Printf("    Goroutines: %d\n", runtime.NumGoroutine())
}

func readinessCheck() {
	fmt.Println("  📊 Readiness check:")
	
	// Simulate readiness checks
	checks := map[string]bool{
		"Database connection": true,
		"Cache connection":    true,
		"External API":        false, // Not ready
		"File system":         true,
	}
	
	ready := true
	for check, status := range checks {
		statusStr := "✅ Ready"
		if !status {
			statusStr = "❌ Not ready"
			ready = false
		}
		fmt.Printf("    %s: %s\n", check, statusStr)
	}
	
	if ready {
		fmt.Println("    Overall: ✅ Ready to serve traffic")
	} else {
		fmt.Println("    Overall: ❌ Not ready to serve traffic")
	}
}

func livenessCheck() {
	fmt.Println("  📊 Liveness check:")
	
	// Simulate liveness check
	alive := true
	uptime := time.Since(time.Now().Add(-time.Hour)) // Simulate 1 hour uptime
	
	fmt.Printf("    Uptime: %v\n", uptime)
	fmt.Printf("    Memory: %d KB\n", getMemoryUsage())
	fmt.Printf("    Goroutines: %d\n", runtime.NumGoroutine())
	
	if alive {
		fmt.Println("    Status: ✅ Alive")
	} else {
		fmt.Println("    Status: ❌ Dead")
	}
}

// 8. Performance Dashboards
func performanceDashboards() {
	fmt.Println("8. Performance Dashboards:")
	fmt.Println("Understanding performance dashboard implementation...")

	// Demonstrate dashboard metrics
	dashboardMetrics()
	
	// Show real-time updates
	realTimeUpdates()
	
	// Demonstrate alerting integration
	alertingIntegration()
}

func dashboardMetrics() {
	fmt.Println("  📊 Dashboard metrics:")
	
	// Simulate dashboard data
	metrics := map[string]interface{}{
		"Requests per second": 150.5,
		"Average response time": "120ms",
		"Error rate": "2.5%",
		"Active connections": 42,
		"Memory usage": "65%",
		"CPU usage": "45%",
	}
	
	for metric, value := range metrics {
		fmt.Printf("    %s: %v\n", metric, value)
	}
}

func realTimeUpdates() {
	fmt.Println("  📊 Real-time updates:")
	
	// Simulate real-time metric updates
	for i := 0; i < 5; i++ {
		rps := 100 + rand.Intn(50)
		responseTime := 100 + rand.Intn(100)
		errorRate := rand.Float64() * 5
		
		fmt.Printf("    Update %d: RPS=%d, ResponseTime=%dms, ErrorRate=%.1f%%\n", 
			i+1, rps, responseTime, errorRate)
		
		time.Sleep(100 * time.Millisecond)
	}
}

func alertingIntegration() {
	fmt.Println("  📊 Alerting integration:")
	
	// Simulate alert integration
	alerts := []string{
		"High error rate detected",
		"Response time exceeded threshold",
		"Memory usage critical",
	}
	
	for i, alert := range alerts {
		fmt.Printf("    Alert %d: %s\n", i+1, alert)
	}
}

// 9. Monitoring Best Practices
func monitoringBestPractices() {
	fmt.Println("9. Monitoring Best Practices:")
	fmt.Println("Best practices for effective monitoring...")

	fmt.Println("  📝 Best Practice 1: Monitor the right metrics")
	fmt.Println("    - Focus on business-critical metrics")
	fmt.Println("    - Monitor user-facing performance")
	fmt.Println("    - Track error rates and availability")
	
	fmt.Println("  📝 Best Practice 2: Set appropriate thresholds")
	fmt.Println("    - Use historical data to set baselines")
	fmt.Println("    - Avoid alert fatigue with proper thresholds")
	fmt.Println("    - Implement alert escalation policies")
	
	fmt.Println("  📝 Best Practice 3: Use multiple monitoring layers")
	fmt.Println("    - Application-level monitoring")
	fmt.Println("    - Infrastructure monitoring")
	fmt.Println("    - Business metrics monitoring")
	
	fmt.Println("  📝 Best Practice 4: Implement proper alerting")
	fmt.Println("    - Use different alert levels")
	fmt.Println("    - Implement alert suppression")
	fmt.Println("    - Test alerting systems regularly")
	
	fmt.Println("  📝 Best Practice 5: Monitor performance trends")
	fmt.Println("    - Track performance over time")
	fmt.Println("    - Identify performance regressions")
	fmt.Println("    - Plan for capacity growth")
	
	fmt.Println("  📝 Best Practice 6: Use distributed tracing")
	fmt.Println("    - Track requests across services")
	fmt.Println("    - Identify performance bottlenecks")
	fmt.Println("    - Debug complex issues")
	
	fmt.Println("  📝 Best Practice 7: Monitor security metrics")
	fmt.Println("    - Track authentication failures")
	fmt.Println("    - Monitor suspicious activity")
	fmt.Println("    - Alert on security events")
}

// 10. Advanced Monitoring
func advancedMonitoring() {
	fmt.Println("10. Advanced Monitoring:")
	fmt.Println("Advanced monitoring techniques and patterns...")

	// Demonstrate distributed tracing
	distributedTracing()
	
	// Show anomaly detection
	anomalyDetection()
	
	// Demonstrate predictive monitoring
	predictiveMonitoring()
}

func distributedTracing() {
	fmt.Println("  📊 Distributed tracing:")
	
	// Simulate distributed trace
	traceID := "trace-12345"
	spanID := "span-67890"
	
	fmt.Printf("    Trace ID: %s\n", traceID)
	fmt.Printf("    Span ID: %s\n", spanID)
	
	// Simulate service calls
	services := []string{"API Gateway", "User Service", "Database", "Cache"}
	
	for i, service := range services {
		startTime := time.Now().Add(-time.Duration(100-i*20) * time.Millisecond)
		duration := time.Duration(20+i*10) * time.Millisecond
		
		fmt.Printf("    %s: %v - %v (%v)\n", 
			service, startTime.Format("15:04:05.000"), 
			startTime.Add(duration).Format("15:04:05.000"), duration)
	}
}

func anomalyDetection() {
	fmt.Println("  📊 Anomaly detection:")
	
	// Simulate anomaly detection
	metrics := []float64{100, 105, 98, 102, 99, 101, 150, 103, 97, 99} // 150 is anomaly
	
	fmt.Println("    Metric values:", metrics)
	
	// Simple anomaly detection (z-score)
	mean := 0.0
	for _, v := range metrics {
		mean += v
	}
	mean /= float64(len(metrics))
	
	variance := 0.0
	for _, v := range metrics {
		variance += (v - mean) * (v - mean)
	}
	variance /= float64(len(metrics))
	stdDev := math.Sqrt(variance)
	
	fmt.Printf("    Mean: %.2f\n", mean)
	fmt.Printf("    Standard deviation: %.2f\n", stdDev)
	
	for i, v := range metrics {
		zScore := (v - mean) / stdDev
		if math.Abs(zScore) > 2 {
			fmt.Printf("    🚨 Anomaly detected at index %d: value=%.2f, z-score=%.2f\n", 
				i, v, zScore)
		}
	}
}

func predictiveMonitoring() {
	fmt.Println("  📊 Predictive monitoring:")
	
	// Simulate predictive analysis
	currentLoad := 70.0
	trend := 5.0 // 5% increase per hour
	
	fmt.Printf("    Current load: %.1f%%\n", currentLoad)
	fmt.Printf("    Trend: +%.1f%% per hour\n", trend)
	
	// Predict future load
	for hour := 1; hour <= 6; hour++ {
		predictedLoad := currentLoad + trend*float64(hour)
		fmt.Printf("    Hour %d: %.1f%%\n", hour, predictedLoad)
		
		if predictedLoad > 90 {
			fmt.Printf("      🚨 Predicted capacity limit reached in hour %d!\n", hour)
			break
		}
	}
}
