package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/config"
	"github.com/your-username/distributed-analytics-platform/internal/server"
	"github.com/your-username/distributed-analytics-platform/internal/analytics"
	"github.com/your-username/distributed-analytics-platform/internal/monitoring"
	"github.com/your-username/distributed-analytics-platform/internal/storage"
	"github.com/your-username/distributed-analytics-platform/internal/ingestion"
	"github.com/your-username/distributed-analytics-platform/internal/processing"
	"github.com/your-username/distributed-analytics-platform/internal/api"
	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/workerpool"
	"github.com/your-username/distributed-analytics-platform/pkg/utils"
)

// main is the entry point for the analytics server
func main() {
	// Parse command line flags
	configFile := flag.String("config", "configs/server.yaml", "Path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger, err := utils.NewLogger(cfg.Log.Level, cfg.Log.Format)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Starting Distributed Analytics Platform Server",
		zap.String("version", "1.0.0"),
		zap.String("config", *configFile),
	)

	// Initialize storage
	eventStore := storage.NewInMemoryEventStore(cfg.Analytics.CacheSize, logger)
	metricStore := storage.NewInMemoryMetricStore(cfg.Analytics.CacheSize, logger)

	// Initialize worker pool
	workerPool := workerpool.NewPool(cfg.Analytics.Workers)
	workerPool.Start()
	defer workerPool.Stop()

	// Initialize ingestion service
	ingestionService := ingestion.NewEventIngestionService(eventStore, workerPool)
	ingestionService.SetLogger(logger)
	ingestionService.Start()
	defer ingestionService.Stop()

	// Initialize processing service
	processingService := processing.NewEventProcessingService(eventStore)
	processingService.SetLogger(logger)
	processingService.SetMetricStore(metricStore)
	processingService.Start()
	defer processingService.Stop()

	// Initialize analytics engine
	analyticsConfig := analytics.Config{
		Workers:      cfg.Analytics.Workers,
		BatchSize:    cfg.Analytics.BatchSize,
		BatchTimeout: cfg.Analytics.BatchTimeout,
		CacheSize:    cfg.Analytics.CacheSize,
	}
	analyticsEngine, err := analytics.NewEngine(analyticsConfig, logger)
	if err != nil {
		logger.Fatal("Failed to initialize analytics engine", zap.Error(err))
	}
	
	// Start analytics engine
	if err := analyticsEngine.Start(); err != nil {
		logger.Fatal("Failed to start analytics engine", zap.Error(err))
	}

	// Initialize monitoring
	monitoringConfig := monitoring.Config{
		MetricsPort:         cfg.Monitoring.MetricsPort,
		HealthCheckInterval: cfg.Monitoring.HealthCheckInterval,
		AlertThresholds:     cfg.Monitoring.AlertThresholds,
	}
	monitor, err := monitoring.NewMonitor(monitoringConfig, logger)
	if err != nil {
		logger.Fatal("Failed to initialize monitoring", zap.Error(err))
	}
	
	// Start monitoring
	if err := monitor.Start(); err != nil {
		logger.Fatal("Failed to start monitoring", zap.Error(err))
	}

	// Initialize API router
	router := api.NewRouter(analyticsEngine, processingService)
	router.SetLogger(logger)
	router.SetIngestionService(ingestionService)
	ginRouter := router.SetupRoutes()

	// Initialize HTTP server
	serverConfig := server.Config{
		Address:      cfg.Server.Address,
		Port:         cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}
	httpServer := server.NewHTTPServerWithRouter(serverConfig, ginRouter, logger)

	// Start server in goroutine
	go func() {
		logger.Info("Starting HTTP server",
			zap.String("address", cfg.Server.Address),
			zap.Int("port", cfg.Server.Port),
		)
		
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start HTTP server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	// Shutdown analytics engine
	if err := analyticsEngine.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown analytics engine", zap.Error(err))
	}

	// Shutdown monitoring
	if err := monitor.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown monitoring", zap.Error(err))
	}

	logger.Info("Server exited")
}