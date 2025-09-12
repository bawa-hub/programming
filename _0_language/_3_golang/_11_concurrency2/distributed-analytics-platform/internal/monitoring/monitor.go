package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Monitor represents the monitoring system
type Monitor struct {
	config      Config
	logger      *zap.Logger
	metrics     map[string]float64
	mu          sync.RWMutex
	isRunning   bool
	ctx         context.Context
	cancel      context.CancelFunc
}

// Config represents the monitoring configuration
type Config struct {
	MetricsPort         int               `json:"metrics_port"`
	HealthCheckInterval time.Duration     `json:"health_check_interval"`
	AlertThresholds     map[string]float64 `json:"alert_thresholds"`
}

// NewMonitor creates a new monitor
func NewMonitor(config Config, logger *zap.Logger) (*Monitor, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	monitor := &Monitor{
		config:  config,
		logger:  logger,
		metrics: make(map[string]float64),
		ctx:     ctx,
		cancel:  cancel,
	}
	
	return monitor, nil
}

// Start starts the monitoring system
func (m *Monitor) Start() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.isRunning {
		return ErrMonitorAlreadyRunning
	}
	
	m.isRunning = true
	
	// Start metrics collection
	go m.collectMetrics()
	
	m.logger.Info("Monitoring system started",
		zap.Int("metrics_port", m.config.MetricsPort),
	)
	
	return nil
}

// Stop stops the monitoring system
func (m *Monitor) Stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if !m.isRunning {
		return nil
	}
	
	m.isRunning = false
	m.cancel()
	
	m.logger.Info("Monitoring system stopped")
	return nil
}

// Shutdown gracefully shuts down the monitor
func (m *Monitor) Shutdown(ctx context.Context) error {
	return m.Stop()
}

// collectMetrics collects system metrics
func (m *Monitor) collectMetrics() {
	ticker := time.NewTicker(m.config.HealthCheckInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			m.updateMetrics()
		case <-m.ctx.Done():
			return
		}
	}
}

// updateMetrics updates the metrics
func (m *Monitor) updateMetrics() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// Update basic metrics
	m.metrics["timestamp"] = float64(time.Now().Unix())
	m.metrics["uptime"] = float64(time.Since(time.Now()).Seconds())
	
	// Check alert thresholds
	m.checkAlerts()
}

// checkAlerts checks if any metrics exceed alert thresholds
func (m *Monitor) checkAlerts() {
	for metric, threshold := range m.config.AlertThresholds {
		if value, exists := m.metrics[metric]; exists {
			if value > threshold {
				m.logger.Warn("Alert threshold exceeded",
					zap.String("metric", metric),
					zap.Float64("value", value),
					zap.Float64("threshold", threshold),
				)
			}
		}
	}
}

// IncEventsIngested increments the events ingested counter
func (m *Monitor) IncEventsIngested(eventType string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	key := "events_ingested_" + eventType
	m.metrics[key]++
}

// IncRateLimitBlocked increments the rate limit blocked counter
func (m *Monitor) IncRateLimitBlocked() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.metrics["rate_limit_blocked"]++
}

// SetMetric sets a metric value
func (m *Monitor) SetMetric(name string, value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.metrics[name] = value
}

// GetMetric gets a metric value
func (m *Monitor) GetMetric(name string) (float64, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	value, exists := m.metrics[name]
	return value, exists
}

// GetMetrics returns all metrics
func (m *Monitor) GetMetrics() map[string]float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	// Return a copy of the metrics
	metrics := make(map[string]float64)
	for k, v := range m.metrics {
		metrics[k] = v
	}
	return metrics
}

// CheckHealth checks the health of the system
func (m *Monitor) CheckHealth() HealthStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	status := HealthStatus{
		OverallStatus: "healthy",
		Timestamp:     time.Now(),
		Metrics:       make(map[string]interface{}),
	}
	
	// Check if monitoring is running
	if !m.isRunning {
		status.OverallStatus = "unhealthy"
		status.Metrics["monitoring"] = "stopped"
	} else {
		status.Metrics["monitoring"] = "running"
	}
	
	// Add basic metrics
	for name, value := range m.metrics {
		status.Metrics[name] = value
	}
	
	return status
}

// HealthStatus represents the health status of the system
type HealthStatus struct {
	OverallStatus string                 `json:"overall_status"`
	Timestamp     time.Time              `json:"timestamp"`
	Metrics       map[string]interface{} `json:"metrics"`
}

// MetricsHandler returns an HTTP handler for metrics
func (m *Monitor) MetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics := m.GetMetrics()
		
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		
		for name, value := range metrics {
			w.Write([]byte(name + " " + fmt.Sprintf("%.2f", value) + "\n"))
		}
	})
}

// Monitor errors
var (
	ErrMonitorAlreadyRunning = &MonitorError{msg: "monitor is already running"}
	ErrMonitorNotRunning     = &MonitorError{msg: "monitor is not running"}
)

// MonitorError represents a monitor-related error
type MonitorError struct {
	msg string
}

func (e *MonitorError) Error() string {
	return e.msg
}