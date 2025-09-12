package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/analytics"
	"github.com/your-username/distributed-analytics-platform/internal/ingestion"
	"github.com/your-username/distributed-analytics-platform/internal/processing"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// Router represents the API router
type Router struct {
	analyticsEngine *analytics.Engine
	ingestionService *ingestion.EventIngestionService
	processingService *processing.EventProcessingService
	logger          *zap.Logger
}

// NewRouter creates a new API router
func NewRouter(analyticsEngine *analytics.Engine, processingService *processing.EventProcessingService) *Router {
	return &Router{
		analyticsEngine:  analyticsEngine,
		processingService: processingService,
		logger:          zap.NewNop(), // Will be set by caller
	}
}

// SetLogger sets the logger for the router
func (r *Router) SetLogger(logger *zap.Logger) {
	r.logger = logger
}

// SetIngestionService sets the ingestion service for the router
func (r *Router) SetIngestionService(ingestionService *ingestion.EventIngestionService) {
	r.ingestionService = ingestionService
}

// SetupRoutes sets up all the API routes
func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.New()
	
	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(r.corsMiddleware())
	
	// Health check
	router.GET("/health", r.healthCheck)
	
	// API routes
	api := router.Group("/api/v1")
	{
		// Events
		api.POST("/events", r.createEvent)
		api.POST("/events/batch", r.createEventBatch)
		api.GET("/events", r.getEvents)
		
		// Metrics
		api.GET("/metrics", r.getMetrics)
		api.GET("/metrics/:name", r.getMetric)
		
		// Streams
		api.POST("/streams", r.createStream)
		api.GET("/streams", r.getStreams)
		api.GET("/streams/:name", r.getStream)
		api.DELETE("/streams/:name", r.deleteStream)
		
		// Analytics
		api.GET("/analytics/stats", r.getAnalyticsStats)
		api.GET("/analytics/health", r.getAnalyticsHealth)
	}
	
	// WebSocket routes
	router.GET("/ws", r.websocketHandler)
	
	// Dashboard
	router.Static("/dashboard", "./web/static")
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", r.dashboard)
	
	return router
}

// corsMiddleware provides CORS support
func (r *Router) corsMiddleware() gin.HandlerFunc {
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
func (r *Router) healthCheck(c *gin.Context) {
	// Check analytics engine health
	stats := r.analyticsEngine.GetStats()
	
	health := gin.H{
		"status":    "healthy",
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
		"analytics": gin.H{
			"workers":      stats.Workers,
			"is_running":   stats.IsRunning,
			"task_queue":   stats.TaskQueue,
			"result_queue": stats.ResultQueue,
		},
	}
	
	c.JSON(http.StatusOK, health)
}

// createEvent handles event creation
func (r *Router) createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Generate ID if not provided
	if event.ID == "" {
		event.ID = generateID()
	}
	
	// Set timestamp if not provided
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	
	if r.ingestionService != nil {
		if err := r.ingestionService.IngestEvent(c.Request.Context(), &event); err != nil {
			r.logger.Error("Failed to ingest event", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
			return
		}
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"id":        event.ID,
		"status":    "created",
		"timestamp": time.Now(),
	})
}

// createEventBatch handles batch event creation
func (r *Router) createEventBatch(c *gin.Context) {
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
	
	if r.ingestionService != nil {
		if err := r.ingestionService.IngestBatch(c.Request.Context(), eventPtrs); err != nil {
			r.logger.Error("Failed to ingest event batch", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event batch"})
			return
		}
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"count":     len(events),
		"status":    "created",
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getEvents handles event retrieval
func (r *Router) getEvents(c *gin.Context) {
	// This would need to be implemented with actual querying
	events := []*models.Event{}
	
	c.JSON(http.StatusOK, gin.H{
		"events":    events,
		"count":     len(events),
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getMetrics handles metric retrieval
func (r *Router) getMetrics(c *gin.Context) {
	// This would need to be implemented with actual querying
	metrics := []*models.AggregatedMetric{}
	
	c.JSON(http.StatusOK, gin.H{
		"metrics":   metrics,
		"count":     len(metrics),
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getMetric handles single metric retrieval
func (r *Router) getMetric(c *gin.Context) {
	name := c.Param("name")
	
	// This would need to be implemented with actual querying
	c.JSON(http.StatusOK, gin.H{
		"metric":    gin.H{"name": name, "value": 0},
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// createStream handles stream creation
func (r *Router) createStream(c *gin.Context) {
	var config analytics.StreamConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	name := c.Param("name")
	if name == "" {
		name = "stream-" + "2025-09-11T00:00:00Z" // This would be time.Now().Format("20060102150405")
	}
	
	_, err := r.analyticsEngine.CreateStream(name, config)
	if err != nil {
		r.logger.Error("Failed to create stream", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stream"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"name":      name,
		"config":    config,
		"status":    "created",
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getStreams handles stream listing
func (r *Router) getStreams(c *gin.Context) {
	// This would need to be implemented
	c.JSON(http.StatusOK, gin.H{
		"streams":   []string{},
		"count":     0,
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getStream handles single stream retrieval
func (r *Router) getStream(c *gin.Context) {
	name := c.Param("name")
	
	stream, exists := r.analyticsEngine.GetStream(name)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stream not found"})
		return
	}
	
	stats := stream.GetStats()
	c.JSON(http.StatusOK, gin.H{
		"stream":    stats,
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// deleteStream handles stream deletion
func (r *Router) deleteStream(c *gin.Context) {
	name := c.Param("name")
	
	if err := r.analyticsEngine.DeleteStream(name); err != nil {
		r.logger.Error("Failed to delete stream", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stream"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":    "deleted",
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getAnalyticsStats handles analytics statistics
func (r *Router) getAnalyticsStats(c *gin.Context) {
	stats := r.analyticsEngine.GetStats()
	
	c.JSON(http.StatusOK, gin.H{
		"analytics": stats,
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// getAnalyticsHealth handles analytics health check
func (r *Router) getAnalyticsHealth(c *gin.Context) {
	stats := r.analyticsEngine.GetStats()
	
	health := "healthy"
	if !stats.IsRunning {
		health = "unhealthy"
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":    health,
		"analytics": stats,
		"timestamp": "2025-09-11T00:00:00Z", // This would be time.Now()
	})
}

// websocketHandler handles WebSocket connections
func (r *Router) websocketHandler(c *gin.Context) {
	// This would need to be implemented with gorilla/websocket
	c.JSON(http.StatusNotImplemented, gin.H{"error": "WebSocket not implemented"})
}

// dashboard handles dashboard requests
func (r *Router) dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Analytics Dashboard",
	})
}
