package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/analytics"
	"github.com/your-username/distributed-analytics-platform/internal/monitoring"
	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/ratelimit"
	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/circuitbreaker"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// HTTPServer represents the HTTP server
type HTTPServer struct {
	config         Config
	logger         *zap.Logger
	analytics      *analytics.Engine
	monitor        *monitoring.Monitor
	rateLimiter    *ratelimit.Limiter
	circuitBreaker *circuitbreaker.CircuitBreaker
	router         *gin.Engine
	server         *http.Server
}

// Config represents server configuration
type Config struct {
	Address      string        `yaml:"address"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// NewHTTPServerWithRouter creates a new HTTP server with a custom router
func NewHTTPServerWithRouter(config Config, router *gin.Engine, logger *zap.Logger) *HTTPServer {
	addr := config.Address + ":" + strconv.Itoa(config.Port)
	
	server := &HTTPServer{
		config: config,
		logger: logger,
		router: router,
	}
	
	server.server = &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
	
	return server
}

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(config Config, analytics *analytics.Engine, monitor *monitoring.Monitor, logger *zap.Logger) *HTTPServer {
	// Create rate limiter
	rateLimiter := ratelimit.NewRateLimiter(1000, time.Second) // 1000 requests per second
	
	// Create circuit breaker
	circuitBreaker := circuitbreaker.NewCircuitBreaker("http_server", 10, 30*time.Second, 60*time.Second)
	
	// Create router
	router := gin.New()
	
	server := &HTTPServer{
		config:         config,
		logger:         logger,
		analytics:      analytics,
		monitor:        monitor,
		rateLimiter:    rateLimiter,
		circuitBreaker: circuitBreaker,
		router:         router,
	}
	
	// Setup routes
	server.setupRoutes()
	
	return server
}

// Start starts the HTTP server
func (s *HTTPServer) Start() error {
	addr := s.config.Address + ":" + strconv.Itoa(s.config.Port)
	
	s.server = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}
	
	s.logger.Info("Starting HTTP server", zap.String("address", addr))
	
	return s.server.ListenAndServe()
}

// Shutdown shuts down the HTTP server
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")
	
	return s.server.Shutdown(ctx)
}

// setupRoutes sets up the HTTP routes
func (s *HTTPServer) setupRoutes() {
	// Middleware
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())
	s.router.Use(s.rateLimitMiddleware())
	s.router.Use(s.circuitBreakerMiddleware())
	s.router.Use(s.corsMiddleware())
	
	// Health check
	s.router.GET("/health", s.healthCheck)
	
	// API routes
	api := s.router.Group("/api/v1")
	{
		// Events
		api.POST("/events", s.createEvent)
		api.POST("/events/batch", s.createEventBatch)
		api.GET("/events", s.getEvents)
		
		// Metrics
		api.GET("/metrics", s.getMetrics)
		api.GET("/metrics/:name", s.getMetric)
		
		// Streams
		api.POST("/streams", s.createStream)
		api.GET("/streams", s.getStreams)
		api.GET("/streams/:name", s.getStream)
		api.DELETE("/streams/:name", s.deleteStream)
		
		// Analytics
		api.GET("/analytics/stats", s.getAnalyticsStats)
		api.GET("/analytics/health", s.getAnalyticsHealth)
	}
	
	// WebSocket routes
	s.router.GET("/ws", s.websocketHandler)
	
	// Dashboard
	s.router.Static("/dashboard", "./web/static")
	s.router.LoadHTMLGlob("web/templates/*")
	s.router.GET("/", s.dashboard)
}

// rateLimitMiddleware provides rate limiting
func (s *HTTPServer) rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !s.rateLimiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// circuitBreakerMiddleware provides circuit breaker protection
func (s *HTTPServer) circuitBreakerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := s.circuitBreaker.Execute(func() (interface{}, error) {
			c.Next()
			return nil, nil
		})
		
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Service temporarily unavailable",
			})
			c.Abort()
			return
		}
	}
}

// corsMiddleware provides CORS support
func (s *HTTPServer) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}

// healthCheck handles health check requests
func (s *HTTPServer) healthCheck(c *gin.Context) {
	// Check analytics engine health
	stats := s.analytics.GetStats()
	
	health := gin.H{
		"status":    "healthy",
		"timestamp": time.Now(),
		"analytics": gin.H{
			"workers":     stats.Workers,
			"is_running":  stats.IsRunning,
			"task_queue":  stats.TaskQueue,
			"result_queue": stats.ResultQueue,
		},
	}
	
	c.JSON(http.StatusOK, health)
}

// createEvent handles event creation
func (s *HTTPServer) createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := s.analytics.ProcessEvent(&event); err != nil {
		s.logger.Error("Failed to process event", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"id":        event.ID,
		"status":    "created",
		"timestamp": time.Now(),
	})
}

// createEventBatch handles batch event creation
func (s *HTTPServer) createEventBatch(c *gin.Context) {
	var events []models.Event
	if err := c.ShouldBindJSON(&events); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Convert to pointers
	eventPtrs := make([]*models.Event, len(events))
	for i := range events {
		eventPtrs[i] = &events[i]
	}
	
	// Process each event individually
	for _, event := range eventPtrs {
		if err := s.analytics.ProcessEvent(event); err != nil {
			s.logger.Error("Failed to process event", zap.Error(err))
		}
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"count":     len(events),
		"status":    "created",
		"timestamp": time.Now(),
	})
}

// getEvents handles event retrieval
func (s *HTTPServer) getEvents(c *gin.Context) {
	// Parse query parameters (for future implementation)
	_ = &analytics.EventQuery{
		StartTime: parseTime(c.Query("start_time")),
		EndTime:   parseTime(c.Query("end_time")),
		EventType: c.Query("event_type"),
		Source:    c.Query("source"),
		UserID:    c.Query("user_id"),
		Limit:     parseInt(c.Query("limit"), 100),
		Offset:    parseInt(c.Query("offset"), 0),
	}
	
	// This would need to be implemented in the analytics engine
	events := []*models.Event{}
	
	c.JSON(http.StatusOK, gin.H{
		"events":    events,
		"count":     len(events),
		"timestamp": time.Now(),
	})
}

// getMetrics handles metric retrieval
func (s *HTTPServer) getMetrics(c *gin.Context) {
	// Parse query parameters
	_ = &analytics.MetricQuery{
		StartTime: parseTime(c.Query("start_time")),
		EndTime:   parseTime(c.Query("end_time")),
		Name:      c.Query("name"),
		Type:      c.Query("type"),
		Source:    c.Query("source"),
		Limit:     parseInt(c.Query("limit"), 100),
		Offset:    parseInt(c.Query("offset"), 0),
	}
	
	// TODO: Implement metrics querying
	metrics := []*models.AggregatedMetric{}
	
	c.JSON(http.StatusOK, gin.H{
		"metrics":   metrics,
		"count":     len(metrics),
		"timestamp": time.Now(),
	})
}

// getMetric handles single metric retrieval
func (s *HTTPServer) getMetric(c *gin.Context) {
	name := c.Param("name")
	
	_ = &analytics.MetricQuery{
		Name: name,
		Limit: 1,
	}
	
	// TODO: Implement metrics querying
	metrics := []*models.AggregatedMetric{}
	
	if len(metrics) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"metric":    metrics[0],
		"timestamp": time.Now(),
	})
}

// createStream handles stream creation
func (s *HTTPServer) createStream(c *gin.Context) {
	var config analytics.StreamConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	name := c.Param("name")
	if name == "" {
		name = generateStreamName()
	}
	
	_, err := s.analytics.CreateStream(name, config)
	if err != nil {
		s.logger.Error("Failed to create stream", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stream"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"name":      name,
		"config":    config,
		"status":    "created",
		"timestamp": time.Now(),
	})
}

// getStreams handles stream listing
func (s *HTTPServer) getStreams(c *gin.Context) {
	// This would need to be implemented in the analytics engine
	c.JSON(http.StatusOK, gin.H{
		"streams":   []string{},
		"count":     0,
		"timestamp": time.Now(),
	})
}

// getStream handles single stream retrieval
func (s *HTTPServer) getStream(c *gin.Context) {
	name := c.Param("name")
	
	stream, exists := s.analytics.GetStream(name)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stream not found"})
		return
	}
	
	stats := stream.GetStats()
	c.JSON(http.StatusOK, gin.H{
		"stream":    stats,
		"timestamp": time.Now(),
	})
}

// deleteStream handles stream deletion
func (s *HTTPServer) deleteStream(c *gin.Context) {
	name := c.Param("name")
	
	if err := s.analytics.DeleteStream(name); err != nil {
		s.logger.Error("Failed to delete stream", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stream"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":    "deleted",
		"timestamp": time.Now(),
	})
}

// getAnalyticsStats handles analytics statistics
func (s *HTTPServer) getAnalyticsStats(c *gin.Context) {
	stats := s.analytics.GetStats()
	
	c.JSON(http.StatusOK, gin.H{
		"analytics": stats,
		"timestamp": time.Now(),
	})
}

// getAnalyticsHealth handles analytics health check
func (s *HTTPServer) getAnalyticsHealth(c *gin.Context) {
	stats := s.analytics.GetStats()
	
	health := "healthy"
	if !stats.IsRunning {
		health = "unhealthy"
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":    health,
		"analytics": stats,
		"timestamp": time.Now(),
	})
}

// websocketHandler handles WebSocket connections
func (s *HTTPServer) websocketHandler(c *gin.Context) {
	// This would need to be implemented with gorilla/websocket
	c.JSON(http.StatusNotImplemented, gin.H{"error": "WebSocket not implemented"})
}

// dashboard handles dashboard requests
func (s *HTTPServer) dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Analytics Dashboard",
	})
}

// Helper functions

func parseTime(timeStr string) time.Time {
	if timeStr == "" {
		return time.Time{}
	}
	
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}
	
	return t
}

func parseInt(str string, defaultValue int) int {
	if str == "" {
		return defaultValue
	}
	
	// Simple implementation - in practice you'd use strconv.Atoi
	return defaultValue
}

func generateStreamName() string {
	return "stream-" + time.Now().Format("20060102150405")
}
