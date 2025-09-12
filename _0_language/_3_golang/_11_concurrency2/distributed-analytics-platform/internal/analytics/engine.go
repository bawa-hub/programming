package analytics

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/workerpool"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// Engine represents the analytics engine
type Engine struct {
	config      Config
	logger      *zap.Logger
	workerPool  *workerpool.Pool
	streams     map[string]*Stream
	mu          sync.RWMutex
	isRunning   bool
	ctx         context.Context
	cancel      context.CancelFunc
}

// Config represents the analytics engine configuration
type Config struct {
	Workers      int           `json:"workers"`
	BatchSize    int           `json:"batch_size"`
	BatchTimeout time.Duration `json:"batch_timeout"`
	CacheSize    int           `json:"cache_size"`
}

// NewEngine creates a new analytics engine
func NewEngine(config Config, logger *zap.Logger) (*Engine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Create worker pool
	workerPool := workerpool.NewPool(config.Workers)
	
	engine := &Engine{
		config:     config,
		logger:     logger,
		workerPool: workerPool,
		streams:    make(map[string]*Stream),
		ctx:        ctx,
		cancel:     cancel,
	}
	
	return engine, nil
}

// Start starts the analytics engine
func (e *Engine) Start() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	
	if e.isRunning {
		return ErrEngineAlreadyRunning
	}
	
	e.isRunning = true
	e.workerPool.Start()
	
	e.logger.Info("Analytics engine started",
		zap.Int("workers", e.config.Workers),
		zap.Int("batch_size", e.config.BatchSize),
	)
	
	return nil
}

// Stop stops the analytics engine
func (e *Engine) Stop() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	
	if !e.isRunning {
		return nil
	}
	
	e.isRunning = false
	e.cancel()
	e.workerPool.Stop()
	
	// Stop all streams
	for name, stream := range e.streams {
		stream.Stop()
		delete(e.streams, name)
	}
	
	e.logger.Info("Analytics engine stopped")
	return nil
}

// Shutdown gracefully shuts down the engine
func (e *Engine) Shutdown(ctx context.Context) error {
	return e.Stop()
}

// ProcessEvent processes an event
func (e *Engine) ProcessEvent(event *models.Event) error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	if !e.isRunning {
		return ErrEngineNotRunning
	}
	
	// Create task
	task := workerpool.Task{
		ID:   event.ID,
		Data: event,
		Handler: func(data interface{}) (interface{}, error) {
			event := data.(*models.Event)
			return e.processEvent(event)
		},
	}
	
	// Submit to worker pool
	if err := e.workerPool.Submit(task); err != nil {
		e.logger.Error("Failed to submit event to worker pool", zap.Error(err))
		return err
	}
	
	return nil
}

// processEvent processes a single event
func (e *Engine) processEvent(event *models.Event) (interface{}, error) {
	// This is where you would add event processing logic
	// For now, we'll just log the event
	e.logger.Debug("Processing event", 
		zap.String("id", event.ID),
		zap.String("type", string(event.Type)),
		zap.String("source", event.Source),
	)
	
	return event, nil
}

// CreateStream creates a new analytics stream
func (e *Engine) CreateStream(name string, config StreamConfig) (*Stream, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	
	if _, exists := e.streams[name]; exists {
		return nil, ErrStreamExists
	}
	
	stream := NewStream(name, config, e.logger)
	e.streams[name] = stream
	
	// Start stream
	stream.Start(e.ctx)
	
	e.logger.Info("Analytics stream created", zap.String("name", name))
	return stream, nil
}

// GetStream retrieves a stream by name
func (e *Engine) GetStream(name string) (*Stream, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	stream, exists := e.streams[name]
	return stream, exists
}

// DeleteStream deletes a stream
func (e *Engine) DeleteStream(name string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	
	stream, exists := e.streams[name]
	if !exists {
		return ErrStreamNotFound
	}
	
	stream.Stop()
	delete(e.streams, name)
	
	e.logger.Info("Analytics stream deleted", zap.String("name", name))
	return nil
}

// ListStreams returns a list of all active stream names
func (e *Engine) ListStreams() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	names := make([]string, 0, len(e.streams))
	for name := range e.streams {
		names = append(names, name)
	}
	return names
}

// GetStats returns engine statistics
func (e *Engine) GetStats() EngineStats {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	poolStats := e.workerPool.GetStats()
	
	return EngineStats{
		Workers:     poolStats.Workers,
		IsRunning:   e.isRunning,
		TaskQueue:   poolStats.TaskQueue,
		ResultQueue: poolStats.ResultQueue,
		Streams:     len(e.streams),
	}
}

// EngineStats represents engine statistics
type EngineStats struct {
	Workers     int
	IsRunning   bool
	TaskQueue   int
	ResultQueue int
	Streams     int
}

// Engine errors
var (
	ErrEngineAlreadyRunning = &EngineError{msg: "engine is already running"}
	ErrEngineNotRunning     = &EngineError{msg: "engine is not running"}
	ErrStreamExists         = &EngineError{msg: "stream already exists"}
	ErrStreamNotFound       = &EngineError{msg: "stream not found"}
)

// EngineError represents an engine-related error
type EngineError struct {
	msg string
}

func (e *EngineError) Error() string {
	return e.msg
}