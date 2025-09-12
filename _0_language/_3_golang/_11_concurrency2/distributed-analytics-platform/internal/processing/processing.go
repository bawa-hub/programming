package processing

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/storage"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// EventProcessingService represents the event processing service
type EventProcessingService struct {
	eventStore storage.EventStore
	metricStore storage.MetricStore
	logger     *zap.Logger
	mu         sync.RWMutex
	isRunning  bool
	ctx        context.Context
	cancel     context.CancelFunc
}

// NewEventProcessingService creates a new event processing service
func NewEventProcessingService(eventStore storage.EventStore) *EventProcessingService {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &EventProcessingService{
		eventStore: eventStore,
		logger:     zap.NewNop(), // Will be set by caller
		ctx:        ctx,
		cancel:     cancel,
		isRunning:  false,
	}
}

// SetLogger sets the logger for the service
func (eps *EventProcessingService) SetLogger(logger *zap.Logger) {
	eps.logger = logger
}

// SetMetricStore sets the metric store for the service
func (eps *EventProcessingService) SetMetricStore(metricStore storage.MetricStore) {
	eps.metricStore = metricStore
}

// Start starts the processing service
func (eps *EventProcessingService) Start() error {
	eps.mu.Lock()
	defer eps.mu.Unlock()
	
	if eps.isRunning {
		return ErrServiceAlreadyRunning
	}
	
	eps.isRunning = true
	eps.logger.Info("Event processing service started")
	
	return nil
}

// Stop stops the processing service
func (eps *EventProcessingService) Stop() error {
	eps.mu.Lock()
	defer eps.mu.Unlock()
	
	if !eps.isRunning {
		return nil
	}
	
	eps.isRunning = false
	eps.cancel()
	eps.logger.Info("Event processing service stopped")
	
	return nil
}

// ProcessEvent processes a single event
func (eps *EventProcessingService) ProcessEvent(ctx context.Context, event *models.Event) error {
	eps.mu.RLock()
	defer eps.mu.RUnlock()
	
	if !eps.isRunning {
		return ErrServiceNotRunning
	}
	
	// Process event based on type
	metrics, err := eps.processEventByType(event)
	if err != nil {
		eps.logger.Error("Failed to process event", 
			zap.String("event_id", event.ID),
			zap.Error(err),
		)
		return err
	}
	
	// Store metrics if metric store is available
	if eps.metricStore != nil && len(metrics) > 0 {
		for _, metric := range metrics {
			if err := eps.metricStore.Store(ctx, metric); err != nil {
				eps.logger.Error("Failed to store metric", 
					zap.String("metric_id", metric.ID),
					zap.Error(err),
				)
			}
		}
	}
	
	return nil
}

// processEventByType processes an event based on its type
func (eps *EventProcessingService) processEventByType(event *models.Event) ([]*models.Metric, error) {
	var metrics []*models.Metric
	
	switch event.Type {
	case models.EventTypePageView:
		metrics = eps.processPageView(event)
	case models.EventTypeClick:
		metrics = eps.processClick(event)
	case models.EventTypePurchase:
		metrics = eps.processPurchase(event)
	case models.EventTypeSignup:
		metrics = eps.processSignup(event)
	case models.EventTypeLogin:
		metrics = eps.processLogin(event)
	case models.EventTypeSearch:
		metrics = eps.processSearch(event)
	default:
		metrics = eps.processCustomEvent(event)
	}
	
	return metrics, nil
}

// processPageView processes a page view event
func (eps *EventProcessingService) processPageView(event *models.Event) []*models.Metric {
	// Extract page information
	page := getStringFromMap(event.Data, "page", "unknown")
	userID := event.UserID
	
	// Create metrics
	metrics := []*models.Metric{
		models.NewMetric("page_views", 1, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"page":    page,
				"user_id": userID,
			}),
		models.NewMetric("unique_page_views", 1, models.MetricTypeCounter).
			SetDimensions(map[string]interface{}{
				"page": page,
			}),
	}
	
	eps.logger.Debug("Processed page view event", 
		zap.String("page", page),
		zap.String("user_id", userID),
	)
	
	return metrics
}

// processClick processes a click event
func (eps *EventProcessingService) processClick(event *models.Event) []*models.Metric {
	// Extract click information
	element := getStringFromMap(event.Data, "element", "unknown")
	page := getStringFromMap(event.Data, "page", "unknown")
	
	// Create metric
	metric := models.NewMetric("clicks", 1, models.MetricTypeCounter).
		SetDimensions(map[string]interface{}{
			"element": element,
			"page":    page,
		})
	
	eps.logger.Debug("Processed click event", 
		zap.String("element", element),
		zap.String("page", page),
	)
	
	return []*models.Metric{metric}
}

// processPurchase processes a purchase event
func (eps *EventProcessingService) processPurchase(event *models.Event) []*models.Metric {
	// Extract purchase information
	amount := getFloat64FromMap(event.Data, "amount", 0.0)
	product := getStringFromMap(event.Data, "product", "unknown")
	
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
	
	eps.logger.Debug("Processed purchase event", 
		zap.String("product", product),
		zap.Float64("amount", amount),
	)
	
	return metrics
}

// processSignup processes a signup event
func (eps *EventProcessingService) processSignup(event *models.Event) []*models.Metric {
	// Create metric
	metric := models.NewMetric("signups", 1, models.MetricTypeCounter)
	
	eps.logger.Debug("Processed signup event")
	
	return []*models.Metric{metric}
}

// processLogin processes a login event
func (eps *EventProcessingService) processLogin(event *models.Event) []*models.Metric {
	// Create metric
	metric := models.NewMetric("logins", 1, models.MetricTypeCounter)
	
	eps.logger.Debug("Processed login event")
	
	return []*models.Metric{metric}
}

// processSearch processes a search event
func (eps *EventProcessingService) processSearch(event *models.Event) []*models.Metric {
	// Extract search information
	query := getStringFromMap(event.Data, "query", "unknown")
	results := getIntFromMap(event.Data, "results", 0)
	
	// Create metrics
	metrics := []*models.Metric{
		models.NewMetric("searches", 1, models.MetricTypeCounter),
		models.NewMetric("search_results", float64(results), models.MetricTypeGauge).
			SetDimensions(map[string]interface{}{
				"query": query,
			}),
	}
	
	eps.logger.Debug("Processed search event", 
		zap.String("query", query),
		zap.Int("results", results),
	)
	
	return metrics
}

// processCustomEvent processes a custom event
func (eps *EventProcessingService) processCustomEvent(event *models.Event) []*models.Metric {
	// Create metric
	metric := models.NewMetric("custom_events", 1, models.MetricTypeCounter).
		SetDimensions(map[string]interface{}{
			"event_type": event.Type,
		})
	
	eps.logger.Debug("Processed custom event", 
		zap.String("event_type", string(event.Type)),
	)
	
	return []*models.Metric{metric}
}

// GetStats returns processing service statistics
func (eps *EventProcessingService) GetStats() ProcessingStats {
	eps.mu.RLock()
	defer eps.mu.RUnlock()
	
	return ProcessingStats{
		IsRunning: eps.isRunning,
	}
}

// ProcessingStats represents processing service statistics
type ProcessingStats struct {
	IsRunning bool
}

// Processing errors
var (
	ErrServiceAlreadyRunning = &ProcessingError{msg: "service is already running"}
	ErrServiceNotRunning     = &ProcessingError{msg: "service is not running"}
)

// ProcessingError represents a processing-related error
type ProcessingError struct {
	msg string
}

func (e *ProcessingError) Error() string {
	return e.msg
}

// Helper functions

func getStringFromMap(data map[string]interface{}, key, defaultValue string) string {
	if value, exists := data[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

func getFloat64FromMap(data map[string]interface{}, key string, defaultValue float64) float64 {
	if value, exists := data[key]; exists {
		if f, ok := value.(float64); ok {
			return f
		}
	}
	return defaultValue
}

func getIntFromMap(data map[string]interface{}, key string, defaultValue int) int {
	if value, exists := data[key]; exists {
		if i, ok := value.(int); ok {
			return i
		}
	}
	return defaultValue
}
