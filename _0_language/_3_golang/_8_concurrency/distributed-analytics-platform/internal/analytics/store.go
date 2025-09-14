package analytics

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// EventStore represents an event store
type EventStore struct {
	events   []*models.Event
	mu       sync.RWMutex
	cache    map[string]*models.Event
	cacheMu  sync.RWMutex
	maxSize  int
	logger   *zap.Logger
}

// NewEventStore creates a new event store
func NewEventStore(maxSize int, logger *zap.Logger) *EventStore {
	return &EventStore{
		events:  make([]*models.Event, 0),
		cache:   make(map[string]*models.Event),
		maxSize: maxSize,
		logger:  logger,
	}
}

// Store stores an event
func (es *EventStore) Store(event *models.Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	// Add to events slice
	es.events = append(es.events, event)
	
	// Add to cache
	es.cacheMu.Lock()
	es.cache[event.ID] = event
	es.cacheMu.Unlock()
	
	// Trim if over max size
	if len(es.events) > es.maxSize {
		es.events = es.events[1:]
	}
	
	return nil
}

// StoreBatch stores a batch of events
func (es *EventStore) StoreBatch(events []*models.Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	// Add to events slice
	es.events = append(es.events, events...)
	
	// Add to cache
	es.cacheMu.Lock()
	for _, event := range events {
		es.cache[event.ID] = event
	}
	es.cacheMu.Unlock()
	
	// Trim if over max size
	if len(es.events) > es.maxSize {
		es.events = es.events[len(es.events)-es.maxSize:]
	}
	
	return nil
}

// Get retrieves an event by ID
func (es *EventStore) Get(id string) (*models.Event, bool) {
	es.cacheMu.RLock()
	defer es.cacheMu.RUnlock()
	
	event, exists := es.cache[id]
	return event, exists
}

// Query queries events with filters
func (es *EventStore) Query(query *EventQuery) ([]*models.Event, error) {
	es.mu.RLock()
	defer es.mu.RUnlock()
	
	var results []*models.Event
	
	for _, event := range es.events {
		if es.matchesQuery(event, query) {
			results = append(results, event)
		}
	}
	
	return results, nil
}

// Count returns the number of events
func (es *EventStore) Count() int {
	es.mu.RLock()
	defer es.mu.RUnlock()
	return len(es.events)
}

// Close closes the event store
func (es *EventStore) Close() error {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	es.events = nil
	es.cache = nil
	
	return nil
}

// matchesQuery checks if an event matches the query
func (es *EventStore) matchesQuery(event *models.Event, query *EventQuery) bool {
	// Check time range
	if !query.StartTime.IsZero() && event.Timestamp.Before(query.StartTime) {
		return false
	}
	if !query.EndTime.IsZero() && event.Timestamp.After(query.EndTime) {
		return false
	}
	
	// Check event type
	if query.EventType != "" && string(event.Type) != query.EventType {
		return false
	}
	
	// Check source
	if query.Source != "" && event.Source != query.Source {
		return false
	}
	
	// Check user ID
	if query.UserID != "" && event.UserID != query.UserID {
		return false
	}
	
	// Check session ID
	if query.SessionID != "" && event.SessionID != query.SessionID {
		return false
	}
	
	return true
}

// EventQuery represents an event query
type EventQuery struct {
	StartTime  time.Time
	EndTime    time.Time
	EventType  string
	Source     string
	UserID     string
	SessionID  string
	Limit      int
	Offset     int
}

// MetricStore represents a metric store
type MetricStore struct {
	metrics  []*models.Metric
	mu       sync.RWMutex
	cache    map[string]*models.Metric
	cacheMu  sync.RWMutex
	maxSize  int
	logger   *zap.Logger
}

// NewMetricStore creates a new metric store
func NewMetricStore(maxSize int, logger *zap.Logger) *MetricStore {
	return &MetricStore{
		metrics: make([]*models.Metric, 0),
		cache:   make(map[string]*models.Metric),
		maxSize: maxSize,
		logger:  logger,
	}
}

// Store stores a metric
func (ms *MetricStore) Store(metric *models.Metric) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	
	// Add to metrics slice
	ms.metrics = append(ms.metrics, metric)
	
	// Add to cache
	ms.cacheMu.Lock()
	ms.cache[metric.ID] = metric
	ms.cacheMu.Unlock()
	
	// Trim if over max size
	if len(ms.metrics) > ms.maxSize {
		ms.metrics = ms.metrics[1:]
	}
	
	return nil
}

// StoreBatch stores a batch of metrics
func (ms *MetricStore) StoreBatch(metrics []*models.Metric) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	
	// Add to metrics slice
	ms.metrics = append(ms.metrics, metrics...)
	
	// Add to cache
	ms.cacheMu.Lock()
	for _, metric := range metrics {
		ms.cache[metric.ID] = metric
	}
	ms.cacheMu.Unlock()
	
	// Trim if over max size
	if len(ms.metrics) > ms.maxSize {
		ms.metrics = ms.metrics[len(ms.metrics)-ms.maxSize:]
	}
	
	return nil
}

// Get retrieves a metric by ID
func (ms *MetricStore) Get(id string) (*models.Metric, bool) {
	ms.cacheMu.RLock()
	defer ms.cacheMu.RUnlock()
	
	metric, exists := ms.cache[id]
	return metric, exists
}

// Query queries metrics with filters
func (ms *MetricStore) Query(query *MetricQuery) ([]*models.AggregatedMetric, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	
	// Group metrics by name and dimensions
	groups := make(map[string]*models.AggregatedMetric)
	
	for _, metric := range ms.metrics {
		if ms.matchesQuery(metric, query) {
			key := ms.getGroupKey(metric)
			
			if group, exists := groups[key]; exists {
				group.AddValue(metric.Value)
			} else {
				group = models.NewAggregatedMetric(metric.Name, metric.Type, query.StartTime, query.EndTime).
					SetDimensions(metric.Dimensions)
				group.AddValue(metric.Value)
				groups[key] = group
			}
		}
	}
	
	// Convert to slice
	var results []*models.AggregatedMetric
	for _, group := range groups {
		results = append(results, group)
	}
	
	return results, nil
}

// Count returns the number of metrics
func (ms *MetricStore) Count() int {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return len(ms.metrics)
}

// Close closes the metric store
func (ms *MetricStore) Close() error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	
	ms.metrics = nil
	ms.cache = nil
	
	return nil
}

// matchesQuery checks if a metric matches the query
func (ms *MetricStore) matchesQuery(metric *models.Metric, query *MetricQuery) bool {
	// Check time range
	if !query.StartTime.IsZero() && metric.Timestamp.Before(query.StartTime) {
		return false
	}
	if !query.EndTime.IsZero() && metric.Timestamp.After(query.EndTime) {
		return false
	}
	
	// Check metric name
	if query.Name != "" && metric.Name != query.Name {
		return false
	}
	
	// Check metric type
	if query.Type != "" && string(metric.Type) != query.Type {
		return false
	}
	
	// Check source
	if query.Source != "" && metric.Source != query.Source {
		return false
	}
	
	return true
}

// getGroupKey generates a group key for metrics
func (ms *MetricStore) getGroupKey(metric *models.Metric) string {
	key := metric.Name
	for k, v := range metric.Dimensions {
		key += ":" + k + "=" + toString(v)
	}
	return key
}

// toString converts a value to string
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return string(rune(val))
	case float64:
		return string(rune(val))
	default:
		return "unknown"
	}
}

// MetricQuery represents a metric query
type MetricQuery struct {
	StartTime time.Time
	EndTime   time.Time
	Name      string
	Type      string
	Source    string
	Limit     int
	Offset    int
}
