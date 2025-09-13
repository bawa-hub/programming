package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 🚀 PHASE 1 MASTERY DEMONSTRATION
// A comprehensive project showcasing all Phase 1 concepts:
// - Goroutines with advanced patterns
// - Channels for communication
// - Synchronization primitives
// - Context for cancellation and values

type Service struct {
	name        string
	url         string
	responseTime time.Duration
	errorRate   float64
}

type HealthCheckResult struct {
	Service     string
	Status      string
	ResponseTime time.Duration
	Error       error
	Timestamp   time.Time
}

type HealthChecker struct {
	services    []Service
	results     chan HealthCheckResult
	shutdown    chan struct{}
	wg          sync.WaitGroup
	mu          sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
}

func main() {
	fmt.Println("🚀 PHASE 1 MASTERY DEMONSTRATION")
	fmt.Println("=================================")
	fmt.Println("Showcasing: Goroutines + Channels + Sync + Context")
	fmt.Println()

	// Create services to monitor
	services := []Service{
		{name: "API Gateway", url: "https://httpbin.org/status/200", responseTime: 100 * time.Millisecond, errorRate: 0.1},
		{name: "User Service", url: "https://httpbin.org/status/200", responseTime: 200 * time.Millisecond, errorRate: 0.05},
		{name: "Payment Service", url: "https://httpbin.org/status/200", responseTime: 300 * time.Millisecond, errorRate: 0.15},
		{name: "Notification Service", url: "https://httpbin.org/status/200", responseTime: 150 * time.Millisecond, errorRate: 0.08},
		{name: "Database", url: "https://httpbin.org/status/200", responseTime: 50 * time.Millisecond, errorRate: 0.02},
	}

	// Create health checker with context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	checker := NewHealthChecker(services, ctx)
	defer checker.Close()

	// Start health checking
	fmt.Println("🏥 Starting health check system...")
	checker.Start()

	// Start result collector
	go checker.CollectResults()

	// Let it run for a while
	time.Sleep(25 * time.Second)

	// Graceful shutdown
	fmt.Println("\n🛑 Initiating graceful shutdown...")
	checker.Shutdown()

	// Wait for completion
	checker.Wait()
	fmt.Println("✅ Health check system stopped gracefully")
}

// NewHealthChecker creates a new health checker
func NewHealthChecker(services []Service, ctx context.Context) *HealthChecker {
	ctx, cancel := context.WithCancel(ctx)
	return &HealthChecker{
		services: services,
		results:  make(chan HealthCheckResult, 100),
		shutdown: make(chan struct{}),
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Start starts the health checking system
func (hc *HealthChecker) Start() {
	// Start health check workers
	for i := 0; i < 3; i++ {
		hc.wg.Add(1)
		go hc.healthCheckWorker(i)
	}

	// Start service monitor
	hc.wg.Add(1)
	go hc.serviceMonitor()

	// Start result processor
	hc.wg.Add(1)
	go hc.resultProcessor()
}

// healthCheckWorker performs health checks
func (hc *HealthChecker) healthCheckWorker(workerID int) {
	defer hc.wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check a random service
			service := hc.getRandomService()
			if service != nil {
				result := hc.checkService(service, workerID)
				hc.results <- result
			}
		case <-hc.ctx.Done():
			fmt.Printf("  🧵 Health check worker %d: Context cancelled\n", workerID)
			return
		case <-hc.shutdown:
			fmt.Printf("  🧵 Health check worker %d: Shutdown signal\n", workerID)
			return
		}
	}
}

// serviceMonitor monitors service health patterns
func (hc *HealthChecker) serviceMonitor() {
	defer hc.wg.Done()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hc.analyzeServiceHealth()
		case <-hc.ctx.Done():
			fmt.Println("  🧵 Service monitor: Context cancelled")
			return
		case <-hc.shutdown:
			fmt.Println("  🧵 Service monitor: Shutdown signal")
			return
		}
	}
}

// resultProcessor processes health check results
func (hc *HealthChecker) resultProcessor() {
	defer hc.wg.Done()

	for {
		select {
		case result := <-hc.results:
			hc.processResult(result)
		case <-hc.ctx.Done():
			fmt.Println("  🧵 Result processor: Context cancelled")
			return
		case <-hc.shutdown:
			fmt.Println("  🧵 Result processor: Shutdown signal")
			return
		}
	}
}

// getRandomService returns a random service
func (hc *HealthChecker) getRandomService() *Service {
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	if len(hc.services) == 0 {
		return nil
	}

	return &hc.services[rand.Intn(len(hc.services))]
}

// checkService performs a health check on a service
func (hc *HealthChecker) checkService(service *Service, workerID int) HealthCheckResult {
	start := time.Now()

	// Simulate network call with context
	_, cancel := context.WithTimeout(hc.ctx, 5*time.Second)
	defer cancel()

	// Simulate response time and error rate
	time.Sleep(service.responseTime)

	// Simulate error based on error rate
	var err error
	if rand.Float64() < service.errorRate {
		err = fmt.Errorf("service error")
	}

	responseTime := time.Since(start)

	status := "healthy"
	if err != nil {
		status = "unhealthy"
	}

	return HealthCheckResult{
		Service:      service.name,
		Status:       status,
		ResponseTime: responseTime,
		Error:        err,
		Timestamp:    time.Now(),
	}
}

// processResult processes a health check result
func (hc *HealthChecker) processResult(result HealthCheckResult) {
	statusIcon := "✅"
	if result.Error != nil {
		statusIcon = "❌"
	}

	fmt.Printf("  %s %s: %s (%.2fms) - %s\n",
		statusIcon,
		result.Service,
		result.Status,
		float64(result.ResponseTime.Nanoseconds())/1e6,
		result.Timestamp.Format("15:04:05"))
}

// analyzeServiceHealth analyzes overall service health
func (hc *HealthChecker) analyzeServiceHealth() {
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	fmt.Printf("  📊 Service Health Analysis (%s):\n", time.Now().Format("15:04:05"))
	for _, service := range hc.services {
		fmt.Printf("    - %s: Response time %.2fms, Error rate %.1f%%\n",
			service.name,
			float64(service.responseTime.Nanoseconds())/1e6,
			service.errorRate*100)
	}
}

// CollectResults collects results from the health checker
func (hc *HealthChecker) CollectResults() {
	// This is a placeholder - in a real implementation,
	// you might want to store results in a database or send them to a monitoring system
	fmt.Println("  📊 Result collector started")
}

// Shutdown gracefully shuts down the health checker
func (hc *HealthChecker) Shutdown() {
	select {
	case <-hc.shutdown:
		// Already closed
	default:
		close(hc.shutdown)
	}
	hc.cancel()
}

// Wait waits for all goroutines to complete
func (hc *HealthChecker) Wait() {
	hc.wg.Wait()
}

// Close closes the health checker
func (hc *HealthChecker) Close() {
	hc.Shutdown()
	close(hc.results)
}
