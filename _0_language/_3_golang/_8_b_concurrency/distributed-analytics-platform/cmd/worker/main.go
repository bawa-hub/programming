package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/worker"
	"github.com/your-username/distributed-analytics-platform/pkg/utils"
)

// main is the entry point for the analytics worker
func main() {
	// Parse command line flags
	configFile := flag.String("config", "configs/worker.yaml", "Path to configuration file")
	flag.Parse()

	// Load configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger, err := utils.NewLogger(config.Log.Level, config.Log.Format)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Starting Distributed Analytics Platform Worker",
		zap.String("version", "1.0.0"),
		zap.String("config", *configFile),
	)

	// Initialize worker
	workerConfig := worker.Config{
		Workers:      config.Worker.Workers,
		BatchSize:    config.Worker.BatchSize,
		BatchTimeout: config.Worker.BatchTimeout,
		PollInterval: config.Worker.PollInterval,
	}
	worker, err := worker.NewWorker(workerConfig, logger)
	if err != nil {
		logger.Fatal("Failed to initialize worker", zap.Error(err))
	}

	// Start worker
	if err := worker.Start(); err != nil {
		logger.Fatal("Failed to start worker", zap.Error(err))
	}

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down worker...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := worker.Shutdown(ctx); err != nil {
		logger.Fatal("Worker forced to shutdown", zap.Error(err))
	}

	logger.Info("Worker exited")
}

// loadConfig loads configuration from file
func loadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// Set default values
	setDefaults()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	// Worker defaults
	viper.SetDefault("worker.workers", 5)
	viper.SetDefault("worker.batch_size", 100)
	viper.SetDefault("worker.batch_timeout", "5s")
	viper.SetDefault("worker.poll_interval", "1s")

	// Logging defaults
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
}

// Config represents the worker configuration
type Config struct {
	Worker WorkerConfig `mapstructure:"worker"`
	Log    LogConfig    `mapstructure:"log"`
}

type WorkerConfig struct {
	Workers      int           `mapstructure:"workers"`
	BatchSize    int           `mapstructure:"batch_size"`
	BatchTimeout time.Duration `mapstructure:"batch_timeout"`
	PollInterval time.Duration `mapstructure:"poll_interval"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}
