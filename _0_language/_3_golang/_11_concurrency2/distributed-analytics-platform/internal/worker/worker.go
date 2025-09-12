package worker

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/workerpool"
	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/circuitbreaker"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// Worker represents a data processing worker
type Worker struct {
	config         Config
	logger         *zap.Logger
	workerPool     *workerpool.Pool
	circuitBreaker *circuitbreaker.CircuitBreaker
	eventQueue     chan *models.Event
	mu             sync.RWMutex
	ctx            context.Context
	cancel         context.CancelFunc
	isRunning      bool
}

// Config represents worker configuration
type Config struct {
	Workers      int           `yaml:"workers"`
	BatchSize    int           `yaml:"batch_size"`
	BatchTimeout time.Duration `yaml:"batch_timeout"`
	PollInterval time.Duration `yaml:"poll_interval"`
}

// NewWorker creates a new worker
func NewWorker(config Config, logger *zap.Logger) (*Worker, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Create worker pool
	workerPool := workerpool.NewPool(config.Workers)
	
	// Create circuit breaker
	circuitBreaker := circuitbreaker.NewCircuitBreaker("worker", 5, 30*time.Second, 60*time.Second)
	
	worker := &Worker{
		config:         config,
		logger:         logger,
		workerPool:     workerPool,
		circuitBreaker: circuitBreaker,
		eventQueue:     make(chan *models.Event, config.BatchSize*2),
		ctx:            ctx,
		cancel:         cancel,
		isRunning:      false,
	}
	
	return worker, nil
}

// Start starts the worker
func (w *Worker) Start() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	
	if w.isRunning {
		return ErrWorkerAlreadyRunning
	}
	
	w.isRunning = true
	
	// Start worker pool
	w.workerPool.Start()
	
	// Start event processing
	go w.processEvents()
	
	// Start batch processing
	go w.processBatches()
	
	w.logger.Info("Worker started",
		zap.Int("workers", w.config.Workers),
		zap.Int("batch_size", w.config.BatchSize),
	)
	
	return nil
}

// Stop stops the worker
func (w *Worker) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()
	
	if !w.isRunning {
		return
	}
	
	w.isRunning = false
	w.cancel()
	
	// Stop worker pool
	w.workerPool.Stop()
	
	// Close event queue
	close(w.eventQueue)
	
	w.logger.Info("Worker stopped")
}

// Shutdown gracefully shuts down the worker
func (w *Worker) Shutdown(ctx context.Context) error {
	w.logger.Info("Shutting down worker")
	
	w.Stop()
	
	// Wait for worker pool to finish
	w.workerPool.Wait()
	
	w.logger.Info("Worker shutdown complete")
	return nil
}

// ProcessEvent processes a single event
func (w *Worker) ProcessEvent(event *models.Event) error {
	w.mu.RLock()
	defer w.mu.RUnlock()
	
	if !w.isRunning {
		return ErrWorkerNotRunning
	}
	
	select {
	case w.eventQueue <- event:
		return nil
	case <-w.ctx.Done():
		return ErrWorkerStopped
	default:
		return ErrQueueFull
	}
}

// ProcessBatch processes a batch of events
func (w *Worker) ProcessBatch(events []*models.Event) error {
	w.mu.RLock()
	defer w.mu.RUnlock()
	
	if !w.isRunning {
		return ErrWorkerNotRunning
	}
	
	// Submit batch to worker pool
	task := workerpool.Task{
		ID:   generateID(),
		Data: events,
		Handler: func(data interface{}) (interface{}, error) {
			events := data.([]*models.Event)
			return w.processEventBatch(events)
		},
	}
	
	if err := w.workerPool.Submit(task); err != nil {
		return err
	}
	
	return nil
}

// processEvents processes events from the queue
func (w *Worker) processEvents() {
	for {
		select {
		case event, ok := <-w.eventQueue:
			if !ok {
				return
			}
			
			// Process event
			if err := w.processEvent(event); err != nil {
				w.logger.Error("Failed to process event", zap.Error(err))
			}
			
		case <-w.ctx.Done():
			return
		}
	}
}

// processBatches processes events in batches
func (w *Worker) processBatches() {
	ticker := time.NewTicker(w.config.BatchTimeout)
	defer ticker.Stop()
	
	batch := make([]*models.Event, 0, w.config.BatchSize)
	
	for {
		select {
		case event, ok := <-w.eventQueue:
			if !ok {
				// Process remaining batch
				if len(batch) > 0 {
					w.processEventBatch(batch)
				}
				return
			}
			
			batch = append(batch, event)
			
			// Process batch if it's full
			if len(batch) >= w.config.BatchSize {
				w.processEventBatch(batch)
				batch = batch[:0]
			}
			
		case <-ticker.C:
			// Process batch on timeout
			if len(batch) > 0 {
				w.processEventBatch(batch)
				batch = batch[:0]
			}
			
		case <-w.ctx.Done():
			return
		}
	}
}

// processEvent processes a single event
func (w *Worker) processEvent(event *models.Event) error {
	// Use circuit breaker for processing
	_, err := w.circuitBreaker.Execute(func() (interface{}, error) {
		return w.processEventInternal(event)
	})
	
	return err
}

// processEventInternal processes an event internally
func (w *Worker) processEventInternal(event *models.Event) (interface{}, error) {
	// Validate event
	if err := event.Validate(); err != nil {
		return nil, err
	}
	
	// Process based on event type
	switch event.Type {
	case models.EventTypePageView:
		return w.processPageView(event)
	case models.EventTypeClick:
		return w.processClick(event)
	case models.EventTypePurchase:
		return w.processPurchase(event)
	case models.EventTypeSignup:
		return w.processSignup(event)
	case models.EventTypeLogin:
		return w.processLogin(event)
	case models.EventTypeSearch:
		return w.processSearch(event)
	default:
		return w.processCustomEvent(event)
	}
}

// processEventBatch processes a batch of events
func (w *Worker) processEventBatch(events []*models.Event) (interface{}, error) {
	processed := 0
	errors := 0
	
	for _, event := range events {
		if err := w.processEvent(event); err != nil {
			errors++
			w.logger.Error("Failed to process event in batch", 
				zap.String("event_id", event.ID),
				zap.Error(err))
		} else {
			processed++
		}
	}
	
	w.logger.Info("Processed event batch",
		zap.Int("total", len(events)),
		zap.Int("processed", processed),
		zap.Int("errors", errors),
	)
	
	return processed, nil
}

// processPageView processes a page view event
func (w *Worker) processPageView(event *models.Event) (interface{}, error) {
	// Extract page information
	page := event.Data["page"].(string)
	userID := event.UserID
	
	// Create metrics
	metrics := []*models.Metric{
		models.NewMetric("page_views", 1, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"page": page,
				"user_id": userID,
			}),
		models.NewMetric("unique_page_views", 1, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"page": page,
			}),
	}
	
	// Process metrics
	for _, metric := range metrics {
		w.logger.Debug("Created metric", 
			zap.String("name", metric.Name),
			zap.Float64("value", metric.Value),
		)
	}
	
	return len(metrics), nil
}

// processClick processes a click event
func (w *Worker) processClick(event *models.Event) (interface{}, error) {
	// Extract click information
	element := event.Data["element"].(string)
	page := event.Data["page"].(string)
	
	// Create metric
	metric := models.NewMetric("clicks", 1, models.MetricTypeCounter).
		SetDimensions(map[string]interface{}{
			"element": element,
			"page": page,
		})
	
	w.logger.Debug("Created click metric", 
		zap.String("element", element),
		zap.String("page", page),
	)
	
	return metric, nil
}

// processPurchase processes a purchase event
func (w *Worker) processPurchase(event *models.Event) (interface{}, error) {
	// Extract purchase information
	amount := event.Data["amount"].(float64)
	product := event.Data["product"].(string)
	
	// Create metrics
	metrics := []*models.Metric{
		models.NewMetric("purchases", 1, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"product": product,
			}),
		models.NewMetric("revenue", amount, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"product": product,
			}),
	}
	
	w.logger.Debug("Created purchase metrics", 
		zap.String("product", product),
		zap.Float64("amount", amount),
	)
	
	return len(metrics), nil
}

// processSignup processes a signup event
func (w *Worker) processSignup(event *models.Event) (interface{}, error) {
	// Create metric
	metric := models.NewMetric("signups", 1, models.MetricTypeCounter)
	
	w.logger.Debug("Created signup metric")
	
	return metric, nil
}

// processLogin processes a login event
func (w *Worker) processLogin(event *models.Event) (interface{}, error) {
	// Create metric
	metric := models.NewMetric("logins", 1, models.MetricTypeCounter)
	
	w.logger.Debug("Created login metric")
	
	return metric, nil
}

// processSearch processes a search event
func (w *Worker) processSearch(event *models.Event) (interface{}, error) {
	// Extract search information
	query := event.Data["query"].(string)
	results := event.Data["results"].(int)
	
	// Create metrics
	metrics := []*models.Metric{
		models.NewMetric("searches", 1, models.MetricTypeCounter),
		models.NewMetric("search_results", float64(results), models.MetricTypeGauge).
			SetDimensions(map[string]interface{}{
				"query": query,
			}),
	}
	
	w.logger.Debug("Created search metrics", 
		zap.String("query", query),
		zap.Int("results", results),
	)
	
	return len(metrics), nil
}

// processCustomEvent processes a custom event
func (w *Worker) processCustomEvent(event *models.Event) (interface{}, error) {
	// Create metric
	metric := models.NewMetric("custom_events", 1, models.MetricTypeCounter).
		SetDimensions(map[string]interface{}{
			"event_type": event.Type,
		})
	
	w.logger.Debug("Created custom event metric", 
		zap.String("event_type", event.Type),
	)
	
	return metric, nil
}

// GetStats returns worker statistics
func (w *Worker) GetStats() WorkerStats {
	w.mu.RLock()
	defer w.mu.RUnlock()
	
	poolStats := w.workerPool.GetStats()
	cbStats := w.circuitBreaker.GetStats()
	
	return WorkerStats{
		IsRunning:      w.isRunning,
		Workers:        poolStats.Workers,
		TaskQueue:      poolStats.TaskQueue,
		ResultQueue:    poolStats.ResultQueue,
		EventQueue:     len(w.eventQueue),
		CircuitBreaker: cbStats,
	}
}

// WorkerStats represents worker statistics
type WorkerStats struct {
	IsRunning      bool
	Workers        int
	TaskQueue      int
	ResultQueue    int
	EventQueue     int
	CircuitBreaker circuitbreaker.CircuitBreakerStats
}

// Worker errors
var (
	ErrWorkerAlreadyRunning = &WorkerError{msg: "worker is already running"}
	ErrWorkerNotRunning     = &WorkerError{msg: "worker is not running"}
	ErrWorkerStopped        = &WorkerError{msg: "worker is stopped"}
	ErrQueueFull            = &WorkerError{msg: "event queue is full"}
)

// WorkerError represents a worker-related error
type WorkerError struct {
	msg string
}

func (e *WorkerError) Error() string {
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
