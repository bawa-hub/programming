package ingestion

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/storage"
	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/workerpool"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// EventIngestionService represents the event ingestion service
type EventIngestionService struct {
	eventStore    storage.EventStore
	workerPool    *workerpool.Pool
	logger        *zap.Logger
	mu            sync.RWMutex
	isRunning     bool
	ctx           context.Context
	cancel        context.CancelFunc
}

// NewEventIngestionService creates a new event ingestion service
func NewEventIngestionService(eventStore storage.EventStore, workerPool *workerpool.Pool) *EventIngestionService {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &EventIngestionService{
		eventStore: eventStore,
		workerPool: workerPool,
		logger:     zap.NewNop(), // Will be set by caller
		ctx:        ctx,
		cancel:     cancel,
		isRunning:  false,
	}
}

// SetLogger sets the logger for the service
func (eis *EventIngestionService) SetLogger(logger *zap.Logger) {
	eis.logger = logger
}

// Start starts the ingestion service
func (eis *EventIngestionService) Start() error {
	eis.mu.Lock()
	defer eis.mu.Unlock()
	
	if eis.isRunning {
		return ErrServiceAlreadyRunning
	}
	
	eis.isRunning = true
	eis.logger.Info("Event ingestion service started")
	
	return nil
}

// Stop stops the ingestion service
func (eis *EventIngestionService) Stop() error {
	eis.mu.Lock()
	defer eis.mu.Unlock()
	
	if !eis.isRunning {
		return nil
	}
	
	eis.isRunning = false
	eis.cancel()
	eis.logger.Info("Event ingestion service stopped")
	
	return nil
}

// IngestEvent ingests a single event
func (eis *EventIngestionService) IngestEvent(ctx context.Context, event *models.Event) error {
	eis.mu.RLock()
	defer eis.mu.RUnlock()
	
	if !eis.isRunning {
		return ErrServiceNotRunning
	}
	
	// Validate event
	if err := event.Validate(); err != nil {
		return err
	}
	
	// Store event
	if err := eis.eventStore.Store(ctx, event); err != nil {
		eis.logger.Error("Failed to store event", zap.Error(err))
		return err
	}
	
	// Submit to worker pool for processing
	task := workerpool.Task{
		ID:   generateID(),
		Data: event,
		Handler: func(data interface{}) (interface{}, error) {
			event := data.(*models.Event)
			return eis.processEvent(event)
		},
	}
	
	if err := eis.workerPool.Submit(task); err != nil {
		eis.logger.Error("Failed to submit event to worker pool", zap.Error(err))
		return err
	}
	
	return nil
}

// IngestBatch ingests a batch of events
func (eis *EventIngestionService) IngestBatch(ctx context.Context, events []*models.Event) error {
	eis.mu.RLock()
	defer eis.mu.RUnlock()
	
	if !eis.isRunning {
		return ErrServiceNotRunning
	}
	
	if len(events) == 0 {
		return nil
	}
	
	// Validate events
	for _, event := range events {
		if err := event.Validate(); err != nil {
			return err
		}
	}
	
	// Store events
	if err := eis.eventStore.StoreBatch(ctx, events); err != nil {
		eis.logger.Error("Failed to store event batch", zap.Error(err))
		return err
	}
	
	// Submit to worker pool for processing
	task := workerpool.Task{
		ID:   generateID(),
		Data: events,
		Handler: func(data interface{}) (interface{}, error) {
			events := data.([]*models.Event)
			return eis.processEventBatch(events)
		},
	}
	
	if err := eis.workerPool.Submit(task); err != nil {
		eis.logger.Error("Failed to submit event batch to worker pool", zap.Error(err))
		return err
	}
	
	return nil
}

// processEvent processes a single event
func (eis *EventIngestionService) processEvent(event *models.Event) (interface{}, error) {
	// This is where you would add event processing logic
	// For now, we'll just log the event
	eis.logger.Debug("Processing event", 
		zap.String("id", event.ID),
		zap.String("type", string(event.Type)),
		zap.String("source", event.Source),
	)
	
	return event, nil
}

// processEventBatch processes a batch of events
func (eis *EventIngestionService) processEventBatch(events []*models.Event) (interface{}, error) {
	processed := 0
	
	for _, event := range events {
		if _, err := eis.processEvent(event); err != nil {
			eis.logger.Error("Failed to process event in batch", 
				zap.String("event_id", event.ID),
				zap.Error(err),
			)
		} else {
			processed++
		}
	}
	
	eis.logger.Info("Processed event batch",
		zap.Int("total", len(events)),
		zap.Int("processed", processed),
	)
	
	return processed, nil
}

// GetStats returns ingestion service statistics
func (eis *EventIngestionService) GetStats() IngestionStats {
	eis.mu.RLock()
	defer eis.mu.RUnlock()
	
	poolStats := eis.workerPool.GetStats()
	
	return IngestionStats{
		IsRunning:   eis.isRunning,
		Workers:     poolStats.Workers,
		TaskQueue:   poolStats.TaskQueue,
		ResultQueue: poolStats.ResultQueue,
	}
}

// IngestionStats represents ingestion service statistics
type IngestionStats struct {
	IsRunning   bool
	Workers     int
	TaskQueue   int
	ResultQueue int
}

// Ingestion errors
var (
	ErrServiceAlreadyRunning = &IngestionError{msg: "service is already running"}
	ErrServiceNotRunning     = &IngestionError{msg: "service is not running"}
)

// IngestionError represents an ingestion-related error
type IngestionError struct {
	msg string
}

func (e *IngestionError) Error() string {
	return e.msg
}

// generateID generates a unique ID
func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random string
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
