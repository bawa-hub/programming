package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Analytics  AnalyticsConfig  `mapstructure:"analytics"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
	Log        LogConfig        `mapstructure:"log"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Database   DatabaseConfig   `mapstructure:"database"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Address      string        `mapstructure:"address"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// AnalyticsConfig represents analytics configuration
type AnalyticsConfig struct {
	Workers      int           `mapstructure:"workers"`
	BatchSize    int           `mapstructure:"batch_size"`
	BatchTimeout time.Duration `mapstructure:"batch_timeout"`
	CacheSize    int           `mapstructure:"cache_size"`
}

// MonitoringConfig represents monitoring configuration
type MonitoringConfig struct {
	MetricsPort         int               `mapstructure:"metrics_port"`
	HealthCheckInterval time.Duration     `mapstructure:"health_check_interval"`
	AlertThresholds     map[string]float64 `mapstructure:"alert_thresholds"`
}

// LogConfig represents logging configuration
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	Username           string        `mapstructure:"username"`
	Password           string        `mapstructure:"password"`
	Database           string        `mapstructure:"database"`
	SSLMode            string        `mapstructure:"ssl_mode"`
	MaxConnections     int           `mapstructure:"max_connections"`
	MaxIdleConnections int           `mapstructure:"max_idle_connections"`
	ConnectionMaxLifetime time.Duration `mapstructure:"connection_max_lifetime"`
}

// LoadConfig loads configuration from file
func LoadConfig(configFile string) (*Config, error) {
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
	// Server defaults
	viper.SetDefault("server.address", "0.0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", "30s")
	viper.SetDefault("server.write_timeout", "30s")
	viper.SetDefault("server.idle_timeout", "60s")

	// Analytics defaults
	viper.SetDefault("analytics.workers", 10)
	viper.SetDefault("analytics.batch_size", 1000)
	viper.SetDefault("analytics.batch_timeout", "1s")
	viper.SetDefault("analytics.cache_size", 10000)

	// Monitoring defaults
	viper.SetDefault("monitoring.metrics_port", 9090)
	viper.SetDefault("monitoring.health_check_interval", "10s")
	viper.SetDefault("monitoring.alert_thresholds", map[string]float64{
		"cpu_usage":    80.0,
		"memory_usage": 85.0,
		"error_rate":   5.0,
	})

	// Logging defaults
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")

	// Redis defaults
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.pool_size", 10)
	viper.SetDefault("redis.min_idle_conns", 5)

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.username", "analytics")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.database", "analytics")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_connections", 100)
	viper.SetDefault("database.max_idle_connections", 10)
	viper.SetDefault("database.connection_max_lifetime", "1h")
}
