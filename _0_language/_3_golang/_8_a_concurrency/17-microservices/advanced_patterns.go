package main

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 1: Service Mesh with Sidecar Proxy
type ServiceMeshWithSidecar struct {
	services    map[string]*ServiceInfo
	sidecars    map[string]*SidecarProxy
	mesh        *ServiceMesh
	mutex       sync.RWMutex
}

type ServiceInfo struct {
	Name    string
	Address string
	Port    int
	Zone    string
	Version string
}

type SidecarProxy struct {
	ServiceName string
	Address     string
	Port        int
	Metrics     *ProxyMetrics
}

type ProxyMetrics struct {
	RequestsProcessed int64
	RequestsFailed    int64
	AvgLatency        time.Duration
	LastRequest       time.Time
}

func NewServiceMeshWithSidecar() *ServiceMeshWithSidecar {
	return &ServiceMeshWithSidecar{
		services: make(map[string]*ServiceInfo),
		sidecars: make(map[string]*SidecarProxy),
		mesh:     NewServiceMesh(),
	}
}

func (sm *ServiceMeshWithSidecar) AddService(name, address string, port int, zone, version string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	service := &ServiceInfo{
		Name:    name,
		Address: address,
		Port:    port,
		Zone:    zone,
		Version: version,
	}
	
	sm.services[name] = service
	
	// Create sidecar proxy for the service
	sidecar := &SidecarProxy{
		ServiceName: name,
		Address:     address,
		Port:        port + 1000, // Sidecar runs on different port
		Metrics:     &ProxyMetrics{},
	}
	
	sm.sidecars[name] = sidecar
}

func (sm *ServiceMeshWithSidecar) SendRequest(from, to, method string, data map[string]interface{}) error {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	// Route through sidecar proxy
	sidecar, exists := sm.sidecars[to]
	if !exists {
		return fmt.Errorf("sidecar not found for service: %s", to)
	}
	
	// Update metrics
	sidecar.Metrics.RequestsProcessed++
	sidecar.Metrics.LastRequest = time.Now()
	
	// Simulate request processing
	fmt.Printf("    %s -> %s (via sidecar): %s\n", from, to, method)
	
	// Simulate occasional failures
	if time.Now().UnixNano()%10 == 0 {
		sidecar.Metrics.RequestsFailed++
		return fmt.Errorf("request failed")
	}
	
	return nil
}

func (sm *ServiceMeshWithSidecar) GetMetrics() map[string]*ProxyMetrics {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	metrics := make(map[string]*ProxyMetrics)
	for name, sidecar := range sm.sidecars {
		metrics[name] = sidecar.Metrics
	}
	return metrics
}

// Advanced Pattern 2: Distributed Tracing with OpenTelemetry
type DistributedTracer struct {
	traces    map[string]*Trace
	spans     map[string]*Span
	exporters []TraceExporter
	mutex     sync.RWMutex
}

type Trace struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Spans     []*Span
	Metadata  map[string]string
}

type Span struct {
	ID        string
	TraceID   string
	ParentID  string
	Operation string
	Service   string
	StartTime time.Time
	EndTime   time.Time
	Tags      map[string]string
	Logs      []LogEntry
}

type LogEntry struct {
	Timestamp time.Time
	Message   string
	Fields    map[string]interface{}
}

type TraceExporter interface {
	Export(trace *Trace) error
}

type ConsoleTraceExporter struct{}

func (cte *ConsoleTraceExporter) Export(trace *Trace) error {
	fmt.Printf("    Exporting trace %s with %d spans\n", trace.ID, len(trace.Spans))
	return nil
}

func NewDistributedTracer() *DistributedTracer {
	tracer := &DistributedTracer{
		traces:    make(map[string]*Trace),
		spans:     make(map[string]*Span),
		exporters: []TraceExporter{&ConsoleTraceExporter{}},
	}
	
	// Start background exporter
	go tracer.startExporter()
	
	return tracer
}

func (dt *DistributedTracer) StartTrace(operation string) *Trace {
	dt.mutex.Lock()
	defer dt.mutex.Unlock()
	
	traceID := fmt.Sprintf("trace-%d", time.Now().UnixNano())
	trace := &Trace{
		ID:        traceID,
		StartTime: time.Now(),
		Spans:     []*Span{},
		Metadata:  make(map[string]string),
	}
	
	dt.traces[traceID] = trace
	return trace
}

func (dt *DistributedTracer) StartSpan(traceID, operation, service string) *Span {
	dt.mutex.Lock()
	defer dt.mutex.Unlock()
	
	spanID := fmt.Sprintf("span-%d", time.Now().UnixNano())
	span := &Span{
		ID:        spanID,
		TraceID:   traceID,
		Operation: operation,
		Service:   service,
		StartTime: time.Now(),
		Tags:      make(map[string]string),
		Logs:      []LogEntry{},
	}
	
	dt.spans[spanID] = span
	
	// Add to trace
	if trace, exists := dt.traces[traceID]; exists {
		trace.Spans = append(trace.Spans, span)
	}
	
	return span
}

func (dt *DistributedTracer) FinishSpan(span *Span) {
	dt.mutex.Lock()
	defer dt.mutex.Unlock()
	
	span.EndTime = time.Now()
}

func (dt *DistributedTracer) FinishTrace(trace *Trace) {
	dt.mutex.Lock()
	defer dt.mutex.Unlock()
	
	trace.EndTime = time.Now()
	
	// Export trace
	for _, exporter := range dt.exporters {
		exporter.Export(trace)
	}
}

func (dt *DistributedTracer) startExporter() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		dt.exportCompletedTraces()
	}
}

func (dt *DistributedTracer) exportCompletedTraces() {
	dt.mutex.RLock()
	defer dt.mutex.RUnlock()
	
	for _, trace := range dt.traces {
		if !trace.EndTime.IsZero() {
			// Trace is completed, export it
			for _, exporter := range dt.exporters {
				exporter.Export(trace)
			}
		}
	}
}

// Advanced Pattern 3: Service Mesh Security with mTLS
type SecureServiceMesh struct {
	services     map[string]*SecureServiceInfo
	certificates map[string]*Certificate
	policyEngine *SecurityPolicyEngine
	mutex        sync.RWMutex
}

type SecureServiceInfo struct {
	Name        string
	Address     string
	Port        int
	Zone        string
	Version     string
	Certificate *Certificate
	Policies    []SecurityPolicy
}

type Certificate struct {
	ServiceName string
	PublicKey   string
	PrivateKey  string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

type SecurityPolicy struct {
	ServiceName string
	Rules       []PolicyRule
}

type PolicyRule struct {
	Action      string
	Source      string
	Destination string
	Method      string
	Path        string
}

type SecurityPolicyEngine struct {
	policies map[string][]SecurityPolicy
	mutex    sync.RWMutex
}

func NewSecureServiceMesh() *SecureServiceMesh {
	return &SecureServiceMesh{
		services:     make(map[string]*SecureServiceInfo),
		certificates: make(map[string]*Certificate),
		policyEngine: &SecurityPolicyEngine{
			policies: make(map[string][]SecurityPolicy),
		},
	}
}

func (ssm *SecureServiceMesh) AddService(name, address string, port int, zone, version string) {
	ssm.mutex.Lock()
	defer ssm.mutex.Unlock()
	
	// Generate certificate for service
	cert := &Certificate{
		ServiceName: name,
		PublicKey:   fmt.Sprintf("public-key-%s", name),
		PrivateKey:  fmt.Sprintf("private-key-%s", name),
		IssuedAt:    time.Now(),
		ExpiresAt:   time.Now().Add(365 * 24 * time.Hour),
	}
	
	ssm.certificates[name] = cert
	
	service := &SecureServiceInfo{
		Name:        name,
		Address:     address,
		Port:        port,
		Zone:        zone,
		Version:     version,
		Certificate: cert,
		Policies:    []SecurityPolicy{},
	}
	
	ssm.services[name] = service
}

func (ssm *SecureServiceMesh) SendSecureRequest(from, to, method string, data map[string]interface{}) error {
	ssm.mutex.RLock()
	defer ssm.mutex.RUnlock()
	
	// Check if service exists
	service, exists := ssm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Check certificate validity
	if time.Now().After(service.Certificate.ExpiresAt) {
		return fmt.Errorf("certificate expired for service: %s", to)
	}
	
	// Check security policies
	if !ssm.policyEngine.IsAllowed(from, to, method) {
		return fmt.Errorf("request blocked by security policy: %s -> %s", from, to)
	}
	
	// Simulate secure communication
	fmt.Printf("    Secure %s -> %s: %s (mTLS)\n", from, to, method)
	
	return nil
}

func (spe *SecurityPolicyEngine) IsAllowed(from, to, method string) bool {
	spe.mutex.RLock()
	defer spe.mutex.RUnlock()
	
	// Simple policy check - allow all for now
	return true
}

// Advanced Pattern 4: Service Mesh Observability
type ObservableServiceMesh struct {
	services    map[string]*ObservableService
	metrics     *MeshMetrics
	traces      *DistributedTracer
	logs        *LogAggregator
	mutex       sync.RWMutex
}

type ObservableService struct {
	Name        string
	Address     string
	Port        int
	Metrics     *ServiceMetrics
	Health      HealthStatus
	LastSeen    time.Time
}

type MeshMetrics struct {
	TotalRequests    int64
	SuccessfulRequests int64
	FailedRequests   int64
	AvgLatency       time.Duration
	P95Latency       time.Duration
	P99Latency       time.Duration
	ErrorRate        float64
}

type LogAggregator struct {
	logs  []LogEntry
	mutex sync.RWMutex
}

func NewObservableServiceMesh() *ObservableServiceMesh {
	return &ObservableServiceMesh{
		services: make(map[string]*ObservableService),
		metrics:  &MeshMetrics{},
		traces:   NewDistributedTracer(),
		logs:     &LogAggregator{logs: []LogEntry{}},
	}
}

func (osm *ObservableServiceMesh) AddService(name, address string, port int) {
	osm.mutex.Lock()
	defer osm.mutex.Unlock()
	
	service := &ObservableService{
		Name:    name,
		Address: address,
		Port:    port,
		Metrics: &ServiceMetrics{
			RequestCount:    0,
			AvgResponseTime: 0,
			ErrorCount:      0,
		},
		Health:   Healthy,
		LastSeen: time.Now(),
	}
	
	osm.services[name] = service
}

func (osm *ObservableServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	osm.mutex.Lock()
	defer osm.mutex.Unlock()
	
	// Update metrics
	osm.metrics.TotalRequests++
	
	// Start trace
	trace := osm.traces.StartTrace(fmt.Sprintf("%s -> %s", from, to))
	span := osm.traces.StartSpan(trace.ID, method, to)
	
	// Simulate request processing
	start := time.Now()
	
	// Simulate occasional failures
	if time.Now().UnixNano()%20 == 0 {
		osm.metrics.FailedRequests++
		osm.traces.FinishSpan(span)
		osm.traces.FinishTrace(trace)
		return fmt.Errorf("request failed")
	}
	
	latency := time.Since(start)
	osm.metrics.SuccessfulRequests++
	
	// Update service metrics
	if svc, exists := osm.services[to]; exists {
		svc.Metrics.RequestCount++
		svc.Metrics.AvgResponseTime = (svc.Metrics.AvgResponseTime + latency) / 2
		svc.LastSeen = time.Now()
	}
	
	// Update mesh metrics
	osm.metrics.AvgLatency = (osm.metrics.AvgLatency + latency) / 2
	
	// Finish trace
	osm.traces.FinishSpan(span)
	osm.traces.FinishTrace(trace)
	
	// Log request
	osm.logs.AddLog(LogEntry{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("%s -> %s: %s", from, to, method),
		Fields: map[string]interface{}{
			"from":   from,
			"to":     to,
			"method": method,
			"latency": latency,
		},
	})
	
	fmt.Printf("    %s -> %s: %s (%v)\n", from, to, method, latency)
	
	return nil
}

func (la *LogAggregator) AddLog(entry LogEntry) {
	la.mutex.Lock()
	defer la.mutex.Unlock()
	
	la.logs = append(la.logs, entry)
	
	// Keep only last 1000 logs
	if len(la.logs) > 1000 {
		la.logs = la.logs[len(la.logs)-1000:]
	}
}

func (osm *ObservableServiceMesh) GetMetrics() *MeshMetrics {
	osm.mutex.RLock()
	defer osm.mutex.RUnlock()
	
	// Calculate error rate
	if osm.metrics.TotalRequests > 0 {
		osm.metrics.ErrorRate = float64(osm.metrics.FailedRequests) / float64(osm.metrics.TotalRequests)
	}
	
	return osm.metrics
}

// Advanced Pattern 5: Service Mesh with Canary Deployment
type CanaryServiceMesh struct {
	services      map[string]*CanaryService
	canaryConfig  *CanaryConfig
	trafficRouter *TrafficRouter
	mutex         sync.RWMutex
}

type CanaryService struct {
	Name        string
	Instances   []*ServiceInstance
	CanaryInstances []*ServiceInstance
	TrafficSplit float64
	HealthCheck *HealthChecker
}

type CanaryConfig struct {
	CanaryPercentage float64
	PromotionThreshold float64
	RollbackThreshold  float64
	EvaluationPeriod   time.Duration
}

type TrafficRouter struct {
	routingRules map[string]*RoutingRule
	mutex        sync.RWMutex
}

type RoutingRule struct {
	ServiceName    string
	CanaryWeight   float64
	StableWeight   float64
	Conditions     []RoutingCondition
}

type RoutingCondition struct {
	Header string
	Value  string
	Weight float64
}

func NewCanaryServiceMesh() *CanaryServiceMesh {
	return &CanaryServiceMesh{
		services: make(map[string]*CanaryService),
		canaryConfig: &CanaryConfig{
			CanaryPercentage:     0.1,  // 10% traffic to canary
			PromotionThreshold:   0.95, // 95% success rate to promote
			RollbackThreshold:    0.05, // 5% error rate to rollback
			EvaluationPeriod:     5 * time.Minute,
		},
		trafficRouter: &TrafficRouter{
			routingRules: make(map[string]*RoutingRule),
		},
	}
}

func (csm *CanaryServiceMesh) AddService(name string, stableInstances, canaryInstances []*ServiceInstance) {
	csm.mutex.Lock()
	defer csm.mutex.Unlock()
	
	service := &CanaryService{
		Name:            name,
		Instances:       stableInstances,
		CanaryInstances: canaryInstances,
		TrafficSplit:    csm.canaryConfig.CanaryPercentage,
		HealthCheck:     NewHealthChecker(),
	}
	
	csm.services[name] = service
	
	// Create routing rule
	rule := &RoutingRule{
		ServiceName:  name,
		CanaryWeight: csm.canaryConfig.CanaryPercentage,
		StableWeight: 1.0 - csm.canaryConfig.CanaryPercentage,
		Conditions:   []RoutingCondition{},
	}
	
	csm.trafficRouter.routingRules[name] = rule
}

func (csm *CanaryServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	csm.mutex.RLock()
	defer csm.mutex.RUnlock()
	
	service, exists := csm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Determine which instance to use (canary or stable)
	useCanary := time.Now().UnixNano()%100 < int64(service.TrafficSplit*100)
	
	if useCanary && len(service.CanaryInstances) > 0 {
		// Use canary instance
		fmt.Printf("    %s -> %s (canary): %s\n", from, to, method)
	} else {
		// Use stable instance
		fmt.Printf("    %s -> %s (stable): %s\n", from, to, method)
	}
	
	// Simulate request processing
	time.Sleep(10 * time.Millisecond)
	
	// Simulate occasional failures
	if time.Now().UnixNano()%20 == 0 {
		return fmt.Errorf("request failed")
	}
	
	return nil
}

func (csm *CanaryServiceMesh) EvaluateCanary(serviceName string) {
	csm.mutex.Lock()
	defer csm.mutex.Unlock()
	
	_, exists := csm.services[serviceName]
	if !exists {
		return
	}
	
	// Simulate canary evaluation
	// In real implementation, this would check metrics, error rates, etc.
	
	// For now, just log the evaluation
	fmt.Printf("    Evaluating canary for %s\n", serviceName)
}

// Advanced Pattern 6: Service Mesh with Circuit Breaker
type CircuitBreakerServiceMesh struct {
	services        map[string]*CircuitBreakerService
	circuitBreakers map[string]*CircuitBreaker
	mutex           sync.RWMutex
}

type CircuitBreakerService struct {
	Name            string
	Instances       []*ServiceInstance
	CircuitBreaker  *CircuitBreaker
	HealthCheck     *HealthChecker
	LastHealthCheck time.Time
}

func NewCircuitBreakerServiceMesh() *CircuitBreakerServiceMesh {
	return &CircuitBreakerServiceMesh{
		services:        make(map[string]*CircuitBreakerService),
		circuitBreakers: make(map[string]*CircuitBreaker),
	}
}

func (cbsm *CircuitBreakerServiceMesh) AddService(name string, instances []*ServiceInstance) {
	cbsm.mutex.Lock()
	defer cbsm.mutex.Unlock()
	
	// Create circuit breaker for service
	circuitBreaker := NewCircuitBreaker(5, 30*time.Second)
	cbsm.circuitBreakers[name] = circuitBreaker
	
	service := &CircuitBreakerService{
		Name:           name,
		Instances:      instances,
		CircuitBreaker: circuitBreaker,
		HealthCheck:    NewHealthChecker(),
	}
	
	cbsm.services[name] = service
}

func (cbsm *CircuitBreakerServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	cbsm.mutex.RLock()
	defer cbsm.mutex.RUnlock()
	
	service, exists := cbsm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Use circuit breaker to protect against failures
	return service.CircuitBreaker.Execute(func() error {
		// Simulate request processing
		time.Sleep(10 * time.Millisecond)
		
		// Simulate occasional failures
		if time.Now().UnixNano()%20 == 0 {
			return fmt.Errorf("request failed")
		}
		
		fmt.Printf("    %s -> %s: %s\n", from, to, method)
		return nil
	})
}

// Advanced Pattern 7: Service Mesh with Load Balancing
type LoadBalancedServiceMesh struct {
	services     map[string]*LoadBalancedService
	loadBalancers map[string]LoadBalancer
	mutex        sync.RWMutex
}

type LoadBalancedService struct {
	Name          string
	Instances     []*ServiceInstance
	LoadBalancer  LoadBalancer
	HealthCheck   *HealthChecker
	LastHealthCheck time.Time
}

func NewLoadBalancedServiceMesh() *LoadBalancedServiceMesh {
	return &LoadBalancedServiceMesh{
		services:      make(map[string]*LoadBalancedService),
		loadBalancers: make(map[string]LoadBalancer),
	}
}

func (lbsm *LoadBalancedServiceMesh) AddService(name string, instances []*ServiceInstance) {
	lbsm.mutex.Lock()
	defer lbsm.mutex.Unlock()
	
	// Create load balancer for service
	loadBalancer := NewRoundRobinBalancer()
	lbsm.loadBalancers[name] = loadBalancer
	
	service := &LoadBalancedService{
		Name:         name,
		Instances:    instances,
		LoadBalancer: loadBalancer,
		HealthCheck:  NewHealthChecker(),
	}
	
	lbsm.services[name] = service
}

func (lbsm *LoadBalancedServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	lbsm.mutex.RLock()
	defer lbsm.mutex.RUnlock()
	
	service, exists := lbsm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Get healthy instances
	healthyInstances := []ServiceInstance{}
	for _, inst := range service.Instances {
		if inst.Health == Healthy {
			healthyInstances = append(healthyInstances, *inst)
		}
	}
	
	if len(healthyInstances) == 0 {
		return fmt.Errorf("no healthy instances available for service: %s", to)
	}
	
	// Select instance using load balancer
	selectedInstance := service.LoadBalancer.SelectInstance(healthyInstances)
	
	// Simulate request processing
	time.Sleep(10 * time.Millisecond)
	
	fmt.Printf("    %s -> %s (%s): %s\n", from, to, selectedInstance.ID, method)
	
	return nil
}

// Advanced Pattern 8: Service Mesh with Rate Limiting
type RateLimitedServiceMesh struct {
	services      map[string]*RateLimitedService
	rateLimiters  map[string]*RateLimiter
	mutex         sync.RWMutex
}

type RateLimitedService struct {
	Name         string
	Instances    []*ServiceInstance
	RateLimiter  *RateLimiter
	HealthCheck  *HealthChecker
	LastHealthCheck time.Time
}

func NewRateLimitedServiceMesh() *RateLimitedServiceMesh {
	return &RateLimitedServiceMesh{
		services:     make(map[string]*RateLimitedService),
		rateLimiters: make(map[string]*RateLimiter),
	}
}

func (rlsm *RateLimitedServiceMesh) AddService(name string, instances []*ServiceInstance, rate int, interval time.Duration) {
	rlsm.mutex.Lock()
	defer rlsm.mutex.Unlock()
	
	// Create rate limiter for service
	rateLimiter := NewRateLimiter(rate, interval)
	rlsm.rateLimiters[name] = rateLimiter
	
	service := &RateLimitedService{
		Name:        name,
		Instances:   instances,
		RateLimiter: rateLimiter,
		HealthCheck: NewHealthChecker(),
	}
	
	rlsm.services[name] = service
}

func (rlsm *RateLimitedServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	rlsm.mutex.RLock()
	defer rlsm.mutex.RUnlock()
	
	service, exists := rlsm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Check rate limit
	if !service.RateLimiter.Allow() {
		return fmt.Errorf("rate limit exceeded for service: %s", to)
	}
	
	// Simulate request processing
	time.Sleep(10 * time.Millisecond)
	
	fmt.Printf("    %s -> %s: %s\n", from, to, method)
	
	return nil
}

// Advanced Pattern 9: Service Mesh with Monitoring
type MonitoredServiceMesh struct {
	services    map[string]*MonitoredService
	monitor     *ServiceMonitor
	metrics     *MeshMetrics
	mutex       sync.RWMutex
}

type MonitoredService struct {
	Name        string
	Instances   []*ServiceInstance
	Metrics     *ServiceMetrics
	HealthCheck *HealthChecker
	LastHealthCheck time.Time
}

func NewMonitoredServiceMesh() *MonitoredServiceMesh {
	return &MonitoredServiceMesh{
		services: make(map[string]*MonitoredService),
		monitor:  NewServiceMonitor(),
		metrics:  &MeshMetrics{},
	}
}

func (msm *MonitoredServiceMesh) AddService(name string, instances []*ServiceInstance) {
	msm.mutex.Lock()
	defer msm.mutex.Unlock()
	
	service := &MonitoredService{
		Name:        name,
		Instances:   instances,
		Metrics:     &ServiceMetrics{},
		HealthCheck: NewHealthChecker(),
	}
	
	msm.services[name] = service
	msm.monitor.AddService(name, "http://localhost:8080/health")
}

func (msm *MonitoredServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	msm.mutex.Lock()
	defer msm.mutex.Unlock()
	
	service, exists := msm.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Update metrics
	msm.metrics.TotalRequests++
	service.Metrics.RequestCount++
	
	// Simulate request processing
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	latency := time.Since(start)
	
	// Update latency metrics
	service.Metrics.AvgResponseTime = (service.Metrics.AvgResponseTime + latency) / 2
	msm.metrics.AvgLatency = (msm.metrics.AvgLatency + latency) / 2
	
	// Simulate occasional failures
	if time.Now().UnixNano()%20 == 0 {
		msm.metrics.FailedRequests++
		service.Metrics.ErrorCount++
		return fmt.Errorf("request failed")
	}
	
	msm.metrics.SuccessfulRequests++
	
	fmt.Printf("    %s -> %s: %s (%v)\n", from, to, method, latency)
	
	return nil
}

func (msm *MonitoredServiceMesh) GetMetrics() *MeshMetrics {
	msm.mutex.RLock()
	defer msm.mutex.RUnlock()
	
	// Calculate error rate
	if msm.metrics.TotalRequests > 0 {
		msm.metrics.ErrorRate = float64(msm.metrics.FailedRequests) / float64(msm.metrics.TotalRequests)
	}
	
	return msm.metrics
}

// Advanced Pattern 10: Service Mesh with Fault Injection
type FaultInjectionServiceMesh struct {
	services        map[string]*FaultInjectionService
	faultInjectors  map[string]*FaultInjector
	mutex           sync.RWMutex
}

type FaultInjectionService struct {
	Name           string
	Instances      []*ServiceInstance
	FaultInjector  *FaultInjector
	HealthCheck    *HealthChecker
	LastHealthCheck time.Time
}

func NewFaultInjectionServiceMesh() *FaultInjectionServiceMesh {
	return &FaultInjectionServiceMesh{
		services:       make(map[string]*FaultInjectionService),
		faultInjectors: make(map[string]*FaultInjector),
	}
}

func (fism *FaultInjectionServiceMesh) AddService(name string, instances []*ServiceInstance) {
	fism.mutex.Lock()
	defer fism.mutex.Unlock()
	
	// Create fault injector for service
	faultInjector := NewFaultInjector()
	faultInjector.AddFault(name, "timeout", 0.1)  // 10% timeout
	faultInjector.AddFault(name, "error", 0.05)   // 5% error
	fism.faultInjectors[name] = faultInjector
	
	service := &FaultInjectionService{
		Name:          name,
		Instances:     instances,
		FaultInjector: faultInjector,
		HealthCheck:   NewHealthChecker(),
	}
	
	fism.services[name] = service
}

func (fism *FaultInjectionServiceMesh) SendRequest(from, to, method string, data map[string]interface{}) error {
	fism.mutex.RLock()
	defer fism.mutex.RUnlock()
	
	service, exists := fism.services[to]
	if !exists {
		return fmt.Errorf("service not found: %s", to)
	}
	
	// Inject faults
	return service.FaultInjector.InjectFault(to, func() error {
		// Simulate request processing
		time.Sleep(10 * time.Millisecond)
		
		fmt.Printf("    %s -> %s: %s\n", from, to, method)
		return nil
	})
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Microservices Patterns")
	fmt.Println("==================================")
	
	// Pattern 1: Service Mesh with Sidecar Proxy
	fmt.Println("\n1. Service Mesh with Sidecar Proxy")
	sidecarMesh := NewServiceMeshWithSidecar()
	sidecarMesh.AddService("user-service", "localhost", 8081, "us-east-1", "1.0.0")
	sidecarMesh.AddService("order-service", "localhost", 8082, "us-east-1", "1.0.0")
	
	for i := 0; i < 5; i++ {
		sidecarMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	// Pattern 2: Distributed Tracing
	fmt.Println("\n2. Distributed Tracing")
	tracer := NewDistributedTracer()
	trace := tracer.StartTrace("microservices-operation")
	
	userSpan := tracer.StartSpan(trace.ID, "get-user", "user-service")
	time.Sleep(100 * time.Millisecond)
	tracer.FinishSpan(userSpan)
	
	orderSpan := tracer.StartSpan(trace.ID, "create-order", "order-service")
	time.Sleep(150 * time.Millisecond)
	tracer.FinishSpan(orderSpan)
	
	tracer.FinishTrace(trace)
	
	// Pattern 3: Service Mesh Security
	fmt.Println("\n3. Service Mesh Security")
	secureMesh := NewSecureServiceMesh()
	secureMesh.AddService("user-service", "localhost", 8081, "us-east-1", "1.0.0")
	secureMesh.AddService("order-service", "localhost", 8082, "us-east-1", "1.0.0")
	
	secureMesh.SendSecureRequest("client", "user-service", "get-user", map[string]interface{}{
		"user_id": "123",
	})
	
	// Pattern 4: Service Mesh Observability
	fmt.Println("\n4. Service Mesh Observability")
	observableMesh := NewObservableServiceMesh()
	observableMesh.AddService("user-service", "localhost", 8081)
	observableMesh.AddService("order-service", "localhost", 8082)
	
	for i := 0; i < 5; i++ {
		observableMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	metrics := observableMesh.GetMetrics()
	fmt.Printf("    Total requests: %d\n", metrics.TotalRequests)
	fmt.Printf("    Error rate: %.2f%%\n", metrics.ErrorRate*100)
	
	// Pattern 5: Canary Deployment
	fmt.Println("\n5. Canary Deployment")
	canaryMesh := NewCanaryServiceMesh()
	
	stableInstances := []*ServiceInstance{
		{ID: "stable-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "stable-2", Address: "localhost", Port: 8082, Health: Healthy},
	}
	
	canaryInstances := []*ServiceInstance{
		{ID: "canary-1", Address: "localhost", Port: 8083, Health: Healthy},
	}
	
	canaryMesh.AddService("user-service", stableInstances, canaryInstances)
	
	for i := 0; i < 10; i++ {
		canaryMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	// Pattern 6: Circuit Breaker Service Mesh
	fmt.Println("\n6. Circuit Breaker Service Mesh")
	circuitBreakerMesh := NewCircuitBreakerServiceMesh()
	
	instances := []*ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "instance-2", Address: "localhost", Port: 8082, Health: Healthy},
	}
	
	circuitBreakerMesh.AddService("user-service", instances)
	
	for i := 0; i < 10; i++ {
		circuitBreakerMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	// Pattern 7: Load Balanced Service Mesh
	fmt.Println("\n7. Load Balanced Service Mesh")
	loadBalancedMesh := NewLoadBalancedServiceMesh()
	
	loadBalancedInstances := []*ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "instance-2", Address: "localhost", Port: 8082, Health: Healthy},
		{ID: "instance-3", Address: "localhost", Port: 8083, Health: Healthy},
	}
	
	loadBalancedMesh.AddService("user-service", loadBalancedInstances)
	
	for i := 0; i < 10; i++ {
		loadBalancedMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	// Pattern 8: Rate Limited Service Mesh
	fmt.Println("\n8. Rate Limited Service Mesh")
	rateLimitedMesh := NewRateLimitedServiceMesh()
	
	rateLimitedInstances := []*ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
	}
	
	rateLimitedMesh.AddService("user-service", rateLimitedInstances, 5, time.Second)
	
	for i := 0; i < 10; i++ {
		err := rateLimitedMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
		if err != nil {
			fmt.Printf("    Request %d failed: %v\n", i+1, err)
		}
	}
	
	// Pattern 9: Monitored Service Mesh
	fmt.Println("\n9. Monitored Service Mesh")
	monitoredMesh := NewMonitoredServiceMesh()
	
	monitoredInstances := []*ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
		{ID: "instance-2", Address: "localhost", Port: 8082, Health: Healthy},
	}
	
	monitoredMesh.AddService("user-service", monitoredInstances)
	
	for i := 0; i < 5; i++ {
		monitoredMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
	}
	
	metrics = monitoredMesh.GetMetrics()
	fmt.Printf("    Total requests: %d\n", metrics.TotalRequests)
	fmt.Printf("    Error rate: %.2f%%\n", metrics.ErrorRate*100)
	
	// Pattern 10: Fault Injection Service Mesh
	fmt.Println("\n10. Fault Injection Service Mesh")
	faultInjectionMesh := NewFaultInjectionServiceMesh()
	
	faultInjectionInstances := []*ServiceInstance{
		{ID: "instance-1", Address: "localhost", Port: 8081, Health: Healthy},
	}
	
	faultInjectionMesh.AddService("user-service", faultInjectionInstances)
	
	for i := 0; i < 10; i++ {
		err := faultInjectionMesh.SendRequest("client", "user-service", "get-user", map[string]interface{}{
			"user_id": fmt.Sprintf("user-%d", i+1),
		})
		if err != nil {
			fmt.Printf("    Request %d failed: %v\n", i+1, err)
		}
	}
	
	fmt.Println("\n🎉 All advanced patterns completed!")
}
