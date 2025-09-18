package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// GOD-LEVEL CONCEPT 9: High-Performance Web Server
// Production-grade HTTP server with advanced concurrency patterns

func main() {
	fmt.Println("=== 🚀 GOD-LEVEL: High-Performance Web Server ===")
	
	// 1. Basic HTTP Server
	demonstrateBasicHTTPServer()
	
	// 2. Connection Pooling
	demonstrateConnectionPooling()
	
	// 3. Request Batching
	demonstrateRequestBatching()
	
	// 4. Circuit Breaker Integration
	demonstrateCircuitBreakerIntegration()
	
	// 5. Rate Limiting
	demonstrateRateLimiting()
	
	// 6. Load Balancing
	demonstrateLoadBalancing()
	
	// 7. Graceful Shutdown
	demonstrateGracefulShutdown()
	
	// 8. Metrics and Monitoring
	demonstrateMetricsAndMonitoring()
	
	// 9. Advanced Caching
	demonstrateAdvancedCaching()
	
	// 10. WebSocket Server
	demonstrateWebSocketServer()
}

// Basic HTTP Server
func demonstrateBasicHTTPServer() {
	fmt.Println("\n=== 1. BASIC HTTP SERVER ===")
	
	fmt.Println(`
🌐 High-Performance HTTP Server:
• Connection pooling and reuse
• Request/response batching
• Circuit breaker patterns
• Rate limiting and throttling
• Load balancing strategies
• Graceful shutdown handling
`)

	// Create server
	server := NewHighPerformanceServer(":8080")
	
	// Add routes
	server.AddRoute("GET", "/health", handleHealth)
	server.AddRoute("GET", "/metrics", handleMetrics)
	server.AddRoute("POST", "/api/data", handleData)
	server.AddRoute("GET", "/api/data/:id", handleGetData)
	
	// Start server in background
	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Wait a bit for server to start
	time.Sleep(100 * time.Millisecond)
	
	// Test server
	testServer("http://localhost:8080")
	
	// Stop server
	server.Stop()
}

func testServer(baseURL string) {
	// Test health endpoint
	resp, err := http.Get(baseURL + "/health")
	if err == nil {
		resp.Body.Close()
		fmt.Println("✅ Health endpoint working")
	}
	
	// Test metrics endpoint
	resp, err = http.Get(baseURL + "/metrics")
	if err == nil {
		resp.Body.Close()
		fmt.Println("✅ Metrics endpoint working")
	}
	
	// Test data endpoint
	data := map[string]interface{}{
		"name": "test",
		"value": 42,
	}
	jsonData, _ := json.Marshal(data)
	resp, err = http.Post(baseURL+"/api/data", "application/json", 
		bytes.NewReader(jsonData))
	if err == nil {
		resp.Body.Close()
		fmt.Println("✅ Data endpoint working")
	}
}

// Connection Pooling
func demonstrateConnectionPooling() {
	fmt.Println("\n=== 2. CONNECTION POOLING ===")
	
	fmt.Println(`
🔌 Connection Pooling:
• Reuse HTTP connections
• Limit concurrent connections
• Connection health checking
• Timeout management
`)

	// Create connection pool
	pool := NewConnectionPool(10, 30*time.Second)
	
	// Test connection pool
	for i := 0; i < 5; i++ {
		conn := pool.Get()
		if conn != nil {
			fmt.Printf("Got connection %d\n", i+1)
			time.Sleep(10 * time.Millisecond)
			pool.Put(conn)
		}
	}
	
	fmt.Println("💡 Connection pooling reduces connection overhead")
}

// Request Batching
func demonstrateRequestBatching() {
	fmt.Println("\n=== 3. REQUEST BATCHING ===")
	
	fmt.Println(`
📦 Request Batching:
• Batch multiple requests together
• Reduce network overhead
• Improve throughput
• Handle batch timeouts
`)

	// Create request batcher
	batcher := NewRequestBatcher(5, 100*time.Millisecond)
	
	// Start batcher
	batcher.Start()
	
	// Add requests
	for i := 0; i < 10; i++ {
		req := &Request{
			ID:   i,
			Data: fmt.Sprintf("request-%d", i),
		}
		batcher.AddRequest(req)
	}
	
	// Wait for batching
	time.Sleep(200 * time.Millisecond)
	
	// Stop batcher
	batcher.Stop()
	
	fmt.Println("💡 Request batching improves throughput")
}

// Circuit Breaker Integration
func demonstrateCircuitBreakerIntegration() {
	fmt.Println("\n=== 4. CIRCUIT BREAKER INTEGRATION ===")
	
	fmt.Println(`
⚡ Circuit Breaker Integration:
• Prevent cascade failures
• Fast failure detection
• Automatic recovery
• Health monitoring
`)

	// Create circuit breaker
	cb := NewCircuitBreaker(5, 10*time.Second, 30*time.Second)
	
	// Test circuit breaker
	for i := 0; i < 10; i++ {
		err := cb.Execute(func() error {
			// Simulate operation
			if rand.Float32() < 0.3 {
				return fmt.Errorf("operation failed")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("Operation %d failed: %v\n", i+1, err)
		} else {
			fmt.Printf("Operation %d succeeded\n", i+1)
		}
	}
	
	fmt.Println("💡 Circuit breaker prevents cascade failures")
}

// Rate Limiting
func demonstrateRateLimiting() {
	fmt.Println("\n=== 5. RATE LIMITING ===")
	
	fmt.Println(`
🚦 Rate Limiting:
• Token bucket algorithm
• Sliding window rate limiting
• Per-client rate limiting
• Burst handling
`)

	// Create rate limiter
	limiter := NewRateLimiter(10, time.Second)
	
	// Test rate limiting
	for i := 0; i < 15; i++ {
		allowed := limiter.Allow()
		if allowed {
			fmt.Printf("Request %d allowed\n", i+1)
		} else {
			fmt.Printf("Request %d rate limited\n", i+1)
		}
		time.Sleep(50 * time.Millisecond)
	}
	
	fmt.Println("💡 Rate limiting prevents abuse")
}

// Load Balancing
func demonstrateLoadBalancing() {
	fmt.Println("\n=== 6. LOAD BALANCING ===")
	
	fmt.Println(`
⚖️ Load Balancing:
• Round-robin balancing
• Weighted round-robin
• Least connections
• Health-based routing
`)

	// Create load balancer
	lb := NewLoadBalancer()
	
	// Add servers
	lb.AddServer("server1", 1.0)
	lb.AddServer("server2", 1.0)
	lb.AddServer("server3", 2.0) // Higher weight
	
	// Test load balancing
	for i := 0; i < 10; i++ {
		server := lb.GetServer()
		fmt.Printf("Request %d routed to %s\n", i+1, server)
	}
	
	fmt.Println("💡 Load balancing distributes load evenly")
}

// Graceful Shutdown
func demonstrateGracefulShutdown() {
	fmt.Println("\n=== 7. GRACEFUL SHUTDOWN ===")
	
	fmt.Println(`
🛑 Graceful Shutdown:
• Handle shutdown signals
• Wait for active requests
• Close connections gracefully
• Cleanup resources
`)

	// Create server with graceful shutdown
	server := NewGracefulServer(":8081")
	
	// Start server
	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Wait a bit
	time.Sleep(100 * time.Millisecond)
	
	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	} else {
		fmt.Println("✅ Graceful shutdown completed")
	}
	
	fmt.Println("💡 Graceful shutdown prevents data loss")
}

// Metrics and Monitoring
func demonstrateMetricsAndMonitoring() {
	fmt.Println("\n=== 8. METRICS AND MONITORING ===")
	
	fmt.Println(`
📊 Metrics and Monitoring:
• Request counters
• Response time histograms
• Error rates
• Resource utilization
`)

	// Create metrics collector
	metrics := NewMetricsCollector()
	
	// Simulate some requests
	for i := 0; i < 100; i++ {
		start := time.Now()
		
		// Simulate request processing
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		
		// Record metrics
		metrics.RecordRequest("GET", "/api/data", time.Since(start), nil)
	}
	
	// Print metrics
	metrics.PrintStats()
	
	fmt.Println("💡 Metrics help identify performance bottlenecks")
}

// Advanced Caching
func demonstrateAdvancedCaching() {
	fmt.Println("\n=== 9. ADVANCED CACHING ===")
	
	fmt.Println(`
💾 Advanced Caching:
• Multi-level caching
• Cache invalidation
• Cache warming
• Distributed caching
`)

	// Create multi-level cache
	cache := NewMultiLevelCache()
	
	// Test caching
	key := "user:123"
	value := "John Doe"
	
	// Set value
	cache.Set(key, value, 5*time.Minute)
	
	// Get value
	cached, found := cache.Get(key)
	if found {
		fmt.Printf("Cache hit: %s\n", cached)
	}
	
	// Test cache invalidation
	cache.Invalidate(key)
	
	cached, found = cache.Get(key)
	if !found {
		fmt.Println("Cache miss after invalidation")
	}
	
	fmt.Println("💡 Advanced caching improves performance")
}

// WebSocket Server
func demonstrateWebSocketServer() {
	fmt.Println("\n=== 10. WEBSOCKET SERVER ===")
	
	fmt.Println(`
🔌 WebSocket Server:
• Real-time communication
• Connection management
• Message broadcasting
• Heartbeat handling
`)

	// Create WebSocket server
	wsServer := NewWebSocketServer(":8082")
	
	// Start server
	go func() {
		if err := wsServer.Start(); err != nil {
			log.Printf("WebSocket server error: %v", err)
		}
	}()
	
	// Wait a bit
	time.Sleep(100 * time.Millisecond)
	
	// Stop server
	wsServer.Stop()
	
	fmt.Println("💡 WebSocket enables real-time communication")
}

// High-Performance Server Implementation
type HighPerformanceServer struct {
	addr        string
	routes      map[string]http.HandlerFunc
	server      *http.Server
	connPool    *ConnectionPool
	rateLimiter *RateLimiter
	circuitBreaker *CircuitBreaker
	metrics     *MetricsCollector
	cache       *MultiLevelCache
	mu          sync.RWMutex
}

func NewHighPerformanceServer(addr string) *HighPerformanceServer {
	return &HighPerformanceServer{
		addr:           addr,
		routes:         make(map[string]http.HandlerFunc),
		connPool:       NewConnectionPool(100, 30*time.Second),
		rateLimiter:    NewRateLimiter(1000, time.Second),
		circuitBreaker: NewCircuitBreaker(10, 5*time.Second, 30*time.Second),
		metrics:        NewMetricsCollector(),
		cache:          NewMultiLevelCache(),
	}
}

func (s *HighPerformanceServer) AddRoute(method, path string, handler http.HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	key := method + " " + path
	s.routes[key] = handler
}

func (s *HighPerformanceServer) Start() error {
	mux := http.NewServeMux()
	
	// Add all routes
	s.mu.RLock()
	for route, handler := range s.routes {
		mux.HandleFunc(route, s.wrapHandler(handler))
	}
	s.mu.RUnlock()
	
	s.server = &http.Server{
		Addr:         s.addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	fmt.Printf("Starting server on %s\n", s.addr)
	return s.server.ListenAndServe()
}

func (s *HighPerformanceServer) wrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Rate limiting
		if !s.rateLimiter.Allow() {
			http.Error(w, "Rate limited", http.StatusTooManyRequests)
			return
		}
		
		// Circuit breaker
		err := s.circuitBreaker.Execute(func() error {
			handler(w, r)
			return nil
		})
		
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}
		
		// Record metrics
		s.metrics.RecordRequest(r.Method, r.URL.Path, time.Since(start), nil)
	}
}

func (s *HighPerformanceServer) Stop() {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.server.Shutdown(ctx)
	}
}

// Handler functions
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Return metrics data
	json.NewEncoder(w).Encode(map[string]interface{}{
		"requests": 1000,
		"errors": 5,
		"avg_response_time": "50ms",
	})
}

func handleData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": rand.Intn(1000),
		"data": data,
		"created_at": time.Now().Format(time.RFC3339),
	})
}

func handleGetData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/data/"):]
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
		"name": "Sample Data",
		"value": 42,
	})
}

// Connection Pool Implementation
type ConnectionPool struct {
	pool    chan *http.Client
	maxSize int
	timeout time.Duration
}

func NewConnectionPool(maxSize int, timeout time.Duration) *ConnectionPool {
	pool := make(chan *http.Client, maxSize)
	
	// Pre-populate pool
	for i := 0; i < maxSize; i++ {
		client := &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     30 * time.Second,
			},
		}
		pool <- client
	}
	
	return &ConnectionPool{
		pool:    pool,
		maxSize: maxSize,
		timeout: timeout,
	}
}

func (cp *ConnectionPool) Get() *http.Client {
	select {
	case client := <-cp.pool:
		return client
	default:
		// Pool is empty, create new client
		return &http.Client{
			Timeout: cp.timeout,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     30 * time.Second,
			},
		}
	}
}

func (cp *ConnectionPool) Put(client *http.Client) {
	select {
	case cp.pool <- client:
		// Successfully returned to pool
	default:
		// Pool is full, discard client
	}
}

// Request Batcher Implementation
type RequestBatcher struct {
	batchSize    int
	batchTimeout time.Duration
	requests     chan *Request
	stopCh       chan struct{}
	wg           sync.WaitGroup
}

type Request struct {
	ID   int
	Data string
}

func NewRequestBatcher(batchSize int, batchTimeout time.Duration) *RequestBatcher {
	return &RequestBatcher{
		batchSize:    batchSize,
		batchTimeout: batchTimeout,
		requests:     make(chan *Request, 1000),
		stopCh:       make(chan struct{}),
	}
}

func (rb *RequestBatcher) Start() {
	rb.wg.Add(1)
	go rb.processBatches()
}

func (rb *RequestBatcher) processBatches() {
	defer rb.wg.Done()
	
	var batch []*Request
	timer := time.NewTimer(rb.batchTimeout)
	defer timer.Stop()
	
	for {
		select {
		case req := <-rb.requests:
			batch = append(batch, req)
			
			if len(batch) >= rb.batchSize {
				rb.processBatch(batch)
				batch = nil
				timer.Reset(rb.batchTimeout)
			}
			
		case <-timer.C:
			if len(batch) > 0 {
				rb.processBatch(batch)
				batch = nil
			}
			timer.Reset(rb.batchTimeout)
			
		case <-rb.stopCh:
			if len(batch) > 0 {
				rb.processBatch(batch)
			}
			return
		}
	}
}

func (rb *RequestBatcher) processBatch(batch []*Request) {
	fmt.Printf("Processing batch of %d requests\n", len(batch))
	
	// Simulate batch processing
	time.Sleep(10 * time.Millisecond)
	
	// Process each request
	for _, req := range batch {
		fmt.Printf("Processed request %d: %s\n", req.ID, req.Data)
	}
}

func (rb *RequestBatcher) AddRequest(req *Request) {
	select {
	case rb.requests <- req:
		// Successfully added
	default:
		// Channel is full, drop request
		fmt.Printf("Dropped request %d: %s\n", req.ID, req.Data)
	}
}

func (rb *RequestBatcher) Stop() {
	close(rb.stopCh)
	rb.wg.Wait()
}

// Circuit Breaker Implementation
type CircuitBreaker struct {
	failureThreshold int
	resetTimeout     time.Duration
	timeout          time.Duration
	failures         int64
	lastFailure      time.Time
	state            string // closed, open, half-open
	mu               sync.RWMutex
}

func NewCircuitBreaker(failureThreshold int, resetTimeout, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: failureThreshold,
		resetTimeout:     resetTimeout,
		timeout:          timeout,
		state:            "closed",
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if circuit is open
	if cb.state == "open" {
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	// Execute function with timeout
	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()
	
	select {
	case err := <-done:
		if err != nil {
			cb.failures++
			cb.lastFailure = time.Now()
			
			if cb.failures >= int64(cb.failureThreshold) {
				cb.state = "open"
			}
			return err
		}
		
		// Success - reset failures
		cb.failures = 0
		cb.state = "closed"
		return nil
		
	case <-time.After(cb.timeout):
		cb.failures++
		cb.lastFailure = time.Now()
		cb.state = "open"
		return fmt.Errorf("operation timeout")
	}
}

// Rate Limiter Implementation
type RateLimiter struct {
	tokens   int64
	capacity int64
	rate     time.Duration
	lastRefill time.Time
	mu       sync.Mutex
}

func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     int64(rate),
		capacity:   int64(rate),
		rate:       window,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	// Refill tokens
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	tokensToAdd := int64(elapsed / rl.rate)
	
	if tokensToAdd > 0 {
		rl.tokens = min(rl.capacity, rl.tokens+tokensToAdd)
		rl.lastRefill = now
	}
	
	// Check if we have tokens
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	
	return false
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// Load Balancer Implementation
type LoadBalancer struct {
	servers []Server
	index   int64
	mu      sync.RWMutex
}

type Server struct {
	Name   string
	Weight float64
	Active bool
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers: make([]Server, 0),
	}
}

func (lb *LoadBalancer) AddServer(name string, weight float64) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	lb.servers = append(lb.servers, Server{
		Name:   name,
		Weight: weight,
		Active: true,
	})
}

func (lb *LoadBalancer) GetServer() string {
	lb.mu.RLock()
	defer lb.mu.RUnlock()
	
	if len(lb.servers) == 0 {
		return "none"
	}
	
	// Round-robin with weights
	totalWeight := 0.0
	for _, server := range lb.servers {
		if server.Active {
			totalWeight += server.Weight
		}
	}
	
	if totalWeight == 0 {
		return "none"
	}
	
	// Find server based on weight
	random := rand.Float64() * totalWeight
	current := 0.0
	
	for _, server := range lb.servers {
		if server.Active {
			current += server.Weight
			if random <= current {
				return server.Name
			}
		}
	}
	
	// Fallback to first active server
	for _, server := range lb.servers {
		if server.Active {
			return server.Name
		}
	}
	
	return "none"
}

// Graceful Server Implementation
type GracefulServer struct {
	server *http.Server
	mu     sync.RWMutex
}

func NewGracefulServer(addr string) *GracefulServer {
	return &GracefulServer{
		server: &http.Server{
			Addr: addr,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello from graceful server"))
			}),
		},
	}
}

func (gs *GracefulServer) Start() error {
	return gs.server.ListenAndServe()
}

func (gs *GracefulServer) Shutdown(ctx context.Context) error {
	return gs.server.Shutdown(ctx)
}

// Metrics Collector Implementation
type MetricsCollector struct {
	requests     int64
	errors       int64
	totalTime    int64
	mu           sync.RWMutex
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{}
}

func (mc *MetricsCollector) RecordRequest(method, path string, duration time.Duration, err error) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	
	atomic.AddInt64(&mc.requests, 1)
	atomic.AddInt64(&mc.totalTime, int64(duration))
	
	if err != nil {
		atomic.AddInt64(&mc.errors, 1)
	}
}

func (mc *MetricsCollector) PrintStats() {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	
	requests := atomic.LoadInt64(&mc.requests)
	errors := atomic.LoadInt64(&mc.errors)
	totalTime := atomic.LoadInt64(&mc.totalTime)
	
	fmt.Printf("Metrics: %d requests, %d errors, avg time: %dms\n", 
		requests, errors, totalTime/max(requests, 1))
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Multi-Level Cache Implementation
type MultiLevelCache struct {
	L1 *sync.Map // Fast in-memory cache
	L2 *sync.Map // Slower but larger cache
	mu sync.RWMutex
}

func NewMultiLevelCache() *MultiLevelCache {
	return &MultiLevelCache{
		L1: &sync.Map{},
		L2: &sync.Map{},
	}
}

func (mlc *MultiLevelCache) Set(key string, value interface{}, ttl time.Duration) {
	// Set in L1 cache
	mlc.L1.Store(key, value)
	
	// Set in L2 cache with TTL
	mlc.L2.Store(key, map[string]interface{}{
		"value": value,
		"expiry": time.Now().Add(ttl),
	})
}

func (mlc *MultiLevelCache) Get(key string) (interface{}, bool) {
	// Try L1 first
	if value, ok := mlc.L1.Load(key); ok {
		return value, true
	}
	
	// Try L2
	if item, ok := mlc.L2.Load(key); ok {
		data := item.(map[string]interface{})
		expiry := data["expiry"].(time.Time)
		
		if time.Now().Before(expiry) {
			value := data["value"]
			// Promote to L1
			mlc.L1.Store(key, value)
			return value, true
		}
		
		// Expired, remove from L2
		mlc.L2.Delete(key)
	}
	
	return nil, false
}

func (mlc *MultiLevelCache) Invalidate(key string) {
	mlc.L1.Delete(key)
	mlc.L2.Delete(key)
}

// WebSocket Server Implementation
type WebSocketServer struct {
	addr     string
	server   *http.Server
	clients  map[*WebSocketClient]bool
	register chan *WebSocketClient
	unregister chan *WebSocketClient
	broadcast chan []byte
	mu       sync.RWMutex
}

type WebSocketClient struct {
	conn   interface{} // Simplified for demo
	send   chan []byte
	server *WebSocketServer
}

func NewWebSocketServer(addr string) *WebSocketServer {
	return &WebSocketServer{
		addr:        addr,
		clients:     make(map[*WebSocketClient]bool),
		register:    make(chan *WebSocketClient),
		unregister:  make(chan *WebSocketClient),
		broadcast:   make(chan []byte),
	}
}

func (wss *WebSocketServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wss.handleWebSocket)
	
	wss.server = &http.Server{
		Addr:    wss.addr,
		Handler: mux,
	}
	
	// Start hub
	go wss.run()
	
	return wss.server.ListenAndServe()
}

func (wss *WebSocketServer) run() {
	for {
		select {
		case client := <-wss.register:
			wss.mu.Lock()
			wss.clients[client] = true
			wss.mu.Unlock()
			
		case client := <-wss.unregister:
			wss.mu.Lock()
			if _, ok := wss.clients[client]; ok {
				delete(wss.clients, client)
				close(client.send)
			}
			wss.mu.Unlock()
			
		case message := <-wss.broadcast:
			wss.mu.RLock()
			for client := range wss.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(wss.clients, client)
				}
			}
			wss.mu.RUnlock()
		}
	}
}

func (wss *WebSocketServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Simple WebSocket implementation
	client := &WebSocketClient{
		conn:   nil, // Simplified
		send:   make(chan []byte, 256),
		server: wss,
	}
	
	wss.register <- client
	
	// Send welcome message
	client.send <- []byte("Welcome to WebSocket server")
	
	// Close after demo
	time.Sleep(100 * time.Millisecond)
	wss.unregister <- client
}

func (wss *WebSocketServer) Stop() {
	if wss.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		wss.server.Shutdown(ctx)
	}
}
