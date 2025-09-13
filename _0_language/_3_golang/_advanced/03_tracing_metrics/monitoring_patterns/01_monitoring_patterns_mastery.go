// 📊 MONITORING PATTERNS MASTERY
// Comprehensive monitoring and observability patterns for production systems
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"time"
)

// ============================================================================
// ALERTING SYSTEM
// ============================================================================

type AlertSeverity string

const (
	SeverityInfo     AlertSeverity = "info"
	SeverityWarning  AlertSeverity = "warning"
	SeverityCritical AlertSeverity = "critical"
	SeverityEmergency AlertSeverity = "emergency"
)

type Alert struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Severity    AlertSeverity `json:"severity"`
	Source      string        `json:"source"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     *time.Time    `json:"end_time,omitempty"`
	Status      string        `json:"status"` // firing, resolved
}

type AlertRule struct {
	Name        string            `json:"name"`
	Query       string            `json:"query"`
	Duration    time.Duration     `json:"duration"`
	Threshold   float64           `json:"threshold"`
	Severity    AlertSeverity     `json:"severity"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type AlertManager struct {
	rules    map[string]AlertRule
	alerts   map[string]Alert
	mu       sync.RWMutex
	notifier AlertNotifier
}

type AlertNotifier interface {
	SendAlert(alert Alert) error
}

type ConsoleNotifier struct{}

func (cn *ConsoleNotifier) SendAlert(alert Alert) error {
	fmt.Printf("🚨 ALERT: [%s] %s - %s\n", alert.Severity, alert.Title, alert.Description)
	return nil
}

func NewAlertManager() *AlertManager {
	return &AlertManager{
		rules:    make(map[string]AlertRule),
		alerts:   make(map[string]Alert),
		notifier: &ConsoleNotifier{},
	}
}

func (am *AlertManager) AddRule(rule AlertRule) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.rules[rule.Name] = rule
}

func (am *AlertManager) EvaluateRules(metrics map[string]float64) {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	for name, rule := range am.rules {
		if value, exists := metrics[rule.Query]; exists {
			if value > rule.Threshold {
				am.fireAlert(name, rule, value)
			} else {
				am.resolveAlert(name)
			}
		}
	}
}

func (am *AlertManager) fireAlert(ruleName string, rule AlertRule, value float64) {
	alertID := fmt.Sprintf("%s-%d", ruleName, time.Now().Unix())
	
	alert := Alert{
		ID:          alertID,
		Title:       rule.Name,
		Description: fmt.Sprintf("Value %.2f exceeds threshold %.2f", value, rule.Threshold),
		Severity:    rule.Severity,
		Source:      "monitoring-system",
		Labels:      rule.Labels,
		Annotations: rule.Annotations,
		StartTime:   time.Now(),
		Status:      "firing",
	}
	
	am.alerts[alertID] = alert
	am.notifier.SendAlert(alert)
}

func (am *AlertManager) resolveAlert(ruleName string) {
	for id, alert := range am.alerts {
		if alert.Labels["rule"] == ruleName && alert.Status == "firing" {
			now := time.Now()
			alert.EndTime = &now
			alert.Status = "resolved"
			am.alerts[id] = alert
			fmt.Printf("✅ RESOLVED: %s\n", alert.Title)
		}
	}
}

func (am *AlertManager) GetActiveAlerts() []Alert {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	var active []Alert
	for _, alert := range am.alerts {
		if alert.Status == "firing" {
			active = append(active, alert)
		}
	}
	return active
}

// ============================================================================
// DASHBOARD SYSTEM
// ============================================================================

type Dashboard struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Panels      []Panel   `json:"panels"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Panel struct {
	ID       string            `json:"id"`
	Title    string            `json:"title"`
	Type     string            `json:"type"` // graph, stat, table, log
	Query    string            `json:"query"`
	Options  map[string]interface{} `json:"options"`
	Position Position          `json:"position"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type DashboardManager struct {
	dashboards map[string]Dashboard
	mu         sync.RWMutex
}

func NewDashboardManager() *DashboardManager {
	return &DashboardManager{
		dashboards: make(map[string]Dashboard),
	}
}

func (dm *DashboardManager) CreateDashboard(dashboard Dashboard) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	
	dashboard.ID = fmt.Sprintf("dashboard-%d", time.Now().Unix())
	dashboard.CreatedAt = time.Now()
	dashboard.UpdatedAt = time.Now()
	
	dm.dashboards[dashboard.ID] = dashboard
}

func (dm *DashboardManager) GetDashboard(id string) (Dashboard, bool) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()
	
	dashboard, exists := dm.dashboards[id]
	return dashboard, exists
}

func (dm *DashboardManager) ListDashboards() []Dashboard {
	dm.mu.RLock()
	defer dm.mu.RUnlock()
	
	var dashboards []Dashboard
	for _, dashboard := range dm.dashboards {
		dashboards = append(dashboards, dashboard)
	}
	
	sort.Slice(dashboards, func(i, j int) bool {
		return dashboards[i].UpdatedAt.After(dashboards[j].UpdatedAt)
	})
	
	return dashboards
}

// ============================================================================
// SLA MONITORING
// ============================================================================

type SLA struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Target      float64       `json:"target"` // e.g., 99.9 for 99.9%
	Window      time.Duration `json:"window"`
	Metrics     []SLAMetric   `json:"metrics"`
}

type SLAMetric struct {
	Name     string  `json:"name"`
	Query    string  `json:"query"`
	Weight   float64 `json:"weight"`
	Required bool    `json:"required"`
}

type SLAMonitor struct {
	slas    map[string]SLA
	results map[string]SLAResult
	mu      sync.RWMutex
}

type SLAResult struct {
	SLA        string    `json:"sla"`
	Value      float64   `json:"value"`
	Target     float64   `json:"target"`
	Status     string    `json:"status"` // met, missed
	Timestamp  time.Time `json:"timestamp"`
	Details    map[string]float64 `json:"details"`
}

func NewSLAMonitor() *SLAMonitor {
	return &SLAMonitor{
		slas:    make(map[string]SLA),
		results: make(map[string]SLAResult),
	}
}

func (sm *SLAMonitor) AddSLA(sla SLA) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.slas[sla.Name] = sla
}

func (sm *SLAMonitor) EvaluateSLA(slaName string, metrics map[string]float64) SLAResult {
	sm.mu.RLock()
	sla, exists := sm.slas[slaName]
	sm.mu.RUnlock()
	
	if !exists {
		return SLAResult{Status: "error"}
	}
	
	// Calculate weighted average
	var totalWeight float64
	var weightedSum float64
	details := make(map[string]float64)
	
	for _, metric := range sla.Metrics {
		if value, exists := metrics[metric.Query]; exists {
			weightedSum += value * metric.Weight
			totalWeight += metric.Weight
			details[metric.Name] = value
		} else if metric.Required {
			// If required metric is missing, SLA is not met
			return SLAResult{
				SLA:       slaName,
				Value:     0,
				Target:    sla.Target,
				Status:    "missed",
				Timestamp: time.Now(),
				Details:   details,
			}
		}
	}
	
	var slaValue float64
	if totalWeight > 0 {
		slaValue = weightedSum / totalWeight
	}
	
	status := "met"
	if slaValue < sla.Target {
		status = "missed"
	}
	
	result := SLAResult{
		SLA:       slaName,
		Value:     slaValue,
		Target:    sla.Target,
		Status:    status,
		Timestamp: time.Now(),
		Details:   details,
	}
	
	sm.mu.Lock()
	sm.results[slaName] = result
	sm.mu.Unlock()
	
	return result
}

func (sm *SLAMonitor) GetSLAStatus(slaName string) (SLAResult, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	result, exists := sm.results[slaName]
	return result, exists
}

// ============================================================================
// PERFORMANCE MONITORING
// ============================================================================

type PerformanceMetrics struct {
	Timestamp     time.Time `json:"timestamp"`
	CPUUsage      float64   `json:"cpu_usage"`
	MemoryUsage   float64   `json:"memory_usage"`
	DiskUsage     float64   `json:"disk_usage"`
	NetworkIn     float64   `json:"network_in"`
	NetworkOut    float64   `json:"network_out"`
	RequestRate   float64   `json:"request_rate"`
	ResponseTime  float64   `json:"response_time"`
	ErrorRate     float64   `json:"error_rate"`
}

type PerformanceMonitor struct {
	metrics    []PerformanceMetrics
	maxMetrics int
	mu         sync.RWMutex
}

func NewPerformanceMonitor(maxMetrics int) *PerformanceMonitor {
	return &PerformanceMonitor{
		metrics:    make([]PerformanceMetrics, 0, maxMetrics),
		maxMetrics: maxMetrics,
	}
}

func (pm *PerformanceMonitor) RecordMetrics(metrics PerformanceMetrics) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.metrics = append(pm.metrics, metrics)
	
	// Keep only the last maxMetrics entries
	if len(pm.metrics) > pm.maxMetrics {
		pm.metrics = pm.metrics[len(pm.metrics)-pm.maxMetrics:]
	}
}

func (pm *PerformanceMonitor) GetMetrics(limit int) []PerformanceMetrics {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	if limit <= 0 || limit > len(pm.metrics) {
		limit = len(pm.metrics)
	}
	
	start := len(pm.metrics) - limit
	if start < 0 {
		start = 0
	}
	
	return pm.metrics[start:]
}

func (pm *PerformanceMonitor) GetAverageMetrics(window time.Duration) PerformanceMetrics {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	cutoff := time.Now().Add(-window)
	var validMetrics []PerformanceMetrics
	
	for _, metric := range pm.metrics {
		if metric.Timestamp.After(cutoff) {
			validMetrics = append(validMetrics, metric)
		}
	}
	
	if len(validMetrics) == 0 {
		return PerformanceMetrics{}
	}
	
	var avg PerformanceMetrics
	avg.Timestamp = time.Now()
	
	for _, metric := range validMetrics {
		avg.CPUUsage += metric.CPUUsage
		avg.MemoryUsage += metric.MemoryUsage
		avg.DiskUsage += metric.DiskUsage
		avg.NetworkIn += metric.NetworkIn
		avg.NetworkOut += metric.NetworkOut
		avg.RequestRate += metric.RequestRate
		avg.ResponseTime += metric.ResponseTime
		avg.ErrorRate += metric.ErrorRate
	}
	
	count := float64(len(validMetrics))
	avg.CPUUsage /= count
	avg.MemoryUsage /= count
	avg.DiskUsage /= count
	avg.NetworkIn /= count
	avg.NetworkOut /= count
	avg.RequestRate /= count
	avg.ResponseTime /= count
	avg.ErrorRate /= count
	
	return avg
}

// ============================================================================
// OBSERVABILITY AGGREGATOR
// ============================================================================

type ObservabilityAggregator struct {
	alertManager      *AlertManager
	dashboardManager  *DashboardManager
	slaMonitor        *SLAMonitor
	performanceMonitor *PerformanceMonitor
	mu                sync.RWMutex
}

func NewObservabilityAggregator() *ObservabilityAggregator {
	return &ObservabilityAggregator{
		alertManager:       NewAlertManager(),
		dashboardManager:   NewDashboardManager(),
		slaMonitor:         NewSLAMonitor(),
		performanceMonitor: NewPerformanceMonitor(1000),
	}
}

func (oa *ObservabilityAggregator) ProcessMetrics(metrics map[string]float64) {
	// Evaluate alert rules
	oa.alertManager.EvaluateRules(metrics)
	
	// Record performance metrics
	perfMetrics := PerformanceMetrics{
		Timestamp:    time.Now(),
		CPUUsage:     metrics["cpu_usage"],
		MemoryUsage:  metrics["memory_usage"],
		RequestRate:  metrics["request_rate"],
		ResponseTime: metrics["response_time"],
		ErrorRate:    metrics["error_rate"],
	}
	oa.performanceMonitor.RecordMetrics(perfMetrics)
}

func (oa *ObservabilityAggregator) GetSystemStatus() map[string]interface{} {
	oa.mu.RLock()
	defer oa.mu.RUnlock()
	
	activeAlerts := oa.alertManager.GetActiveAlerts()
	dashboards := oa.dashboardManager.ListDashboards()
	
	status := map[string]interface{}{
		"timestamp":    time.Now(),
		"active_alerts": len(activeAlerts),
		"dashboards":   len(dashboards),
		"alerts":       activeAlerts,
	}
	
	return status
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateAlertingSystem() {
	fmt.Println("\n=== Alerting System ===")
	
	alertManager := NewAlertManager()
	
	// Add alert rules
	alertManager.AddRule(AlertRule{
		Name:        "High CPU Usage",
		Query:       "cpu_usage",
		Duration:    5 * time.Minute,
		Threshold:   80.0,
		Severity:    SeverityWarning,
		Labels:      map[string]string{"service": "api", "rule": "cpu"},
		Annotations: map[string]string{"summary": "CPU usage is high"},
	})
	
	alertManager.AddRule(AlertRule{
		Name:        "High Error Rate",
		Query:       "error_rate",
		Duration:    2 * time.Minute,
		Threshold:   5.0,
		Severity:    SeverityCritical,
		Labels:      map[string]string{"service": "api", "rule": "errors"},
		Annotations: map[string]string{"summary": "Error rate is too high"},
	})
	
	// Simulate metrics
	metrics := map[string]float64{
		"cpu_usage":    85.0, // Should trigger alert
		"memory_usage": 60.0,
		"error_rate":   2.0,  // Should not trigger alert
		"request_rate": 100.0,
	}
	
	alertManager.EvaluateRules(metrics)
	
	activeAlerts := alertManager.GetActiveAlerts()
	fmt.Printf("   📊 Active Alerts: %d\n", len(activeAlerts))
	for _, alert := range activeAlerts {
		fmt.Printf("   🚨 %s: %s\n", alert.Severity, alert.Title)
	}
}

func demonstrateDashboardSystem() {
	fmt.Println("\n=== Dashboard System ===")
	
	dashboardManager := NewDashboardManager()
	
	// Create a system overview dashboard
	dashboard := Dashboard{
		Name:        "System Overview",
		Description: "Main system monitoring dashboard",
		Panels: []Panel{
			{
				Title: "CPU Usage",
				Type:  "graph",
				Query: "cpu_usage",
				Position: Position{X: 0, Y: 0, W: 6, H: 4},
			},
			{
				Title: "Memory Usage",
				Type:  "graph",
				Query: "memory_usage",
				Position: Position{X: 6, Y: 0, W: 6, H: 4},
			},
			{
				Title: "Request Rate",
				Type:  "stat",
				Query: "request_rate",
				Position: Position{X: 0, Y: 4, W: 4, H: 2},
			},
			{
				Title: "Error Rate",
				Type:  "stat",
				Query: "error_rate",
				Position: Position{X: 4, Y: 4, W: 4, H: 2},
			},
		},
	}
	
	dashboardManager.CreateDashboard(dashboard)
	
	dashboards := dashboardManager.ListDashboards()
	fmt.Printf("   📊 Created %d dashboards\n", len(dashboards))
	for _, dash := range dashboards {
		fmt.Printf("   📊 %s: %d panels\n", dash.Name, len(dash.Panels))
	}
}

func demonstrateSLAMonitoring() {
	fmt.Println("\n=== SLA Monitoring ===")
	
	slaMonitor := NewSLAMonitor()
	
	// Define SLA
	sla := SLA{
		Name:        "API Availability",
		Description: "API service availability SLA",
		Target:      99.9, // 99.9%
		Window:      24 * time.Hour,
		Metrics: []SLAMetric{
			{
				Name:     "Uptime",
				Query:    "uptime_percentage",
				Weight:   1.0,
				Required: true,
			},
		},
	}
	
	slaMonitor.AddSLA(sla)
	
	// Simulate metrics
	metrics := map[string]float64{
		"uptime_percentage": 99.95, // Above target
	}
	
	result := slaMonitor.EvaluateSLA("API Availability", metrics)
	fmt.Printf("   📊 SLA: %s\n", result.SLA)
	fmt.Printf("   📊 Value: %.2f%% (Target: %.2f%%)\n", result.Value, result.Target)
	fmt.Printf("   📊 Status: %s\n", result.Status)
}

func demonstratePerformanceMonitoring() {
	fmt.Println("\n=== Performance Monitoring ===")
	
	perfMonitor := NewPerformanceMonitor(100)
	
	// Simulate performance metrics over time
	for i := 0; i < 10; i++ {
		metrics := PerformanceMetrics{
			Timestamp:    time.Now().Add(time.Duration(i) * time.Minute),
			CPUUsage:     50 + rand.Float64()*30,
			MemoryUsage:  60 + rand.Float64()*20,
			RequestRate:  100 + rand.Float64()*50,
			ResponseTime: 100 + rand.Float64()*200,
			ErrorRate:    rand.Float64() * 2,
		}
		
		perfMonitor.RecordMetrics(metrics)
	}
	
	// Get average metrics for last hour
	avgMetrics := perfMonitor.GetAverageMetrics(time.Hour)
	fmt.Printf("   📊 Average CPU Usage: %.2f%%\n", avgMetrics.CPUUsage)
	fmt.Printf("   📊 Average Memory Usage: %.2f%%\n", avgMetrics.MemoryUsage)
	fmt.Printf("   📊 Average Request Rate: %.2f req/s\n", avgMetrics.RequestRate)
	fmt.Printf("   📊 Average Response Time: %.2f ms\n", avgMetrics.ResponseTime)
	fmt.Printf("   📊 Average Error Rate: %.2f%%\n", avgMetrics.ErrorRate)
}

func demonstrateObservabilityAggregation() {
	fmt.Println("\n=== Observability Aggregation ===")
	
	aggregator := NewObservabilityAggregator()
	
	// Set up SLA
	aggregator.slaMonitor.AddSLA(SLA{
		Name:    "Service Health",
		Target:  99.0,
		Window:  time.Hour,
		Metrics: []SLAMetric{
			{Name: "Availability", Query: "availability", Weight: 1.0, Required: true},
		},
	})
	
	// Process metrics
	metrics := map[string]float64{
		"cpu_usage":    75.0,
		"memory_usage": 65.0,
		"request_rate": 150.0,
		"response_time": 120.0,
		"error_rate":   1.5,
		"availability": 99.5,
	}
	
	aggregator.ProcessMetrics(metrics)
	
	status := aggregator.GetSystemStatus()
	fmt.Printf("   📊 System Status:\n")
	fmt.Printf("   📊 Active Alerts: %v\n", status["active_alerts"])
	fmt.Printf("   📊 Dashboards: %v\n", status["dashboards"])
}

func demonstrateMonitoringAPI() {
	fmt.Println("\n=== Monitoring API ===")
	
	aggregator := NewObservabilityAggregator()
	
	// Set up some initial data
	aggregator.alertManager.AddRule(AlertRule{
		Name:      "Test Alert",
		Query:     "test_metric",
		Threshold: 10.0,
		Severity:  SeverityWarning,
	})
	
	// Create HTTP handlers
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics := map[string]float64{
			"test_metric": 15.0, // Will trigger alert
		}
		aggregator.ProcessMetrics(metrics)
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	})
	
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := aggregator.GetSystemStatus()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	})
	
	fmt.Println("   📊 Monitoring API endpoints available:")
	fmt.Println("   📊 GET /metrics - Current metrics")
	fmt.Println("   📊 GET /status - System status")
}

func main() {
	fmt.Println("📊 MONITORING PATTERNS MASTERY")
	fmt.Println("==============================")
	
	demonstrateAlertingSystem()
	demonstrateDashboardSystem()
	demonstrateSLAMonitoring()
	demonstratePerformanceMonitoring()
	demonstrateObservabilityAggregation()
	demonstrateMonitoringAPI()
	
	fmt.Println("\n🎉 MONITORING PATTERNS MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Alerting systems and rules")
	fmt.Println("✅ Dashboard creation and management")
	fmt.Println("✅ SLA monitoring and tracking")
	fmt.Println("✅ Performance monitoring patterns")
	fmt.Println("✅ Observability aggregation")
	fmt.Println("✅ Monitoring API endpoints")
	
	fmt.Println("\n🚀 You are now ready for Phase 5: Package Design & Architecture Mastery!")
}
