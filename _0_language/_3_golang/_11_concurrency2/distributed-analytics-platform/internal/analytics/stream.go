package analytics

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/pkg/concurrency/reactive"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

// Stream represents an analytics stream
type Stream struct {
	name        string
	config      StreamConfig
	logger      *zap.Logger
	subscribers map[string]*Subscriber
	mu          sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
	reactive    *reactive.Stream[*models.Event]
}

// StreamConfig represents stream configuration
type StreamConfig struct {
	BufferSize    int           `yaml:"buffer_size"`
	FlushInterval time.Duration `yaml:"flush_interval"`
	MaxSubscribers int          `yaml:"max_subscribers"`
}

// Subscriber represents a stream subscriber
type Subscriber struct {
	ID       string
	Channel  chan *models.Event
	Filter   func(*models.Event) bool
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewStream creates a new analytics stream
func NewStream(name string, config StreamConfig, logger *zap.Logger) *Stream {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Set defaults
	if config.BufferSize == 0 {
		config.BufferSize = 1000
	}
	if config.FlushInterval == 0 {
		config.FlushInterval = 1 * time.Second
	}
	if config.MaxSubscribers == 0 {
		config.MaxSubscribers = 100
	}
	
	stream := &Stream{
		name:        name,
		config:      config,
		logger:      logger,
		subscribers: make(map[string]*Subscriber),
		ctx:         ctx,
		cancel:      cancel,
	}
	
	// Create reactive stream
	stream.reactive = reactive.NewStream[*models.Event]()
	
	return stream
}

// Start starts the stream
func (s *Stream) Start(ctx context.Context) {
	s.logger.Info("Starting analytics stream", zap.String("name", s.name))
	
	// Start flush ticker
	ticker := time.NewTicker(s.config.FlushInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Stopping analytics stream", zap.String("name", s.name))
			return
		case <-ticker.C:
			s.flush()
		}
	}
}

// Stop stops the stream
func (s *Stream) Stop() {
	s.cancel()
	
	// Close all subscribers
	s.mu.Lock()
	for _, subscriber := range s.subscribers {
		subscriber.cancel()
		close(subscriber.Channel)
	}
	s.subscribers = make(map[string]*Subscriber)
	s.mu.Unlock()
}

// Publish publishes an event to the stream
func (s *Stream) Publish(event *models.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// Emit to reactive stream
	s.reactive.Emit(event)
	
	// Send to subscribers
	for _, subscriber := range s.subscribers {
		select {
		case subscriber.Channel <- event:
		default:
			// Subscriber channel is full, skip
			s.logger.Warn("Subscriber channel full, skipping event",
				zap.String("subscriber", subscriber.ID),
				zap.String("stream", s.name))
		}
	}
	
	return nil
}

// Subscribe subscribes to the stream
func (s *Stream) Subscribe(id string, filter func(*models.Event) bool) (*Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if len(s.subscribers) >= s.config.MaxSubscribers {
		return nil, ErrMaxSubscribersReached
	}
	
	ctx, cancel := context.WithCancel(s.ctx)
	subscriber := &Subscriber{
		ID:      id,
		Channel: make(chan *models.Event, s.config.BufferSize),
		Filter:  filter,
		ctx:     ctx,
		cancel:  cancel,
	}
	
	s.subscribers[id] = subscriber
	
	s.logger.Info("Subscriber added to stream",
		zap.String("subscriber", id),
		zap.String("stream", s.name))
	
	return subscriber, nil
}

// Unsubscribe unsubscribes from the stream
func (s *Stream) Unsubscribe(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	subscriber, exists := s.subscribers[id]
	if !exists {
		return ErrSubscriberNotFound
	}
	
	subscriber.cancel()
	close(subscriber.Channel)
	delete(s.subscribers, id)
	
	s.logger.Info("Subscriber removed from stream",
		zap.String("subscriber", id),
		zap.String("stream", s.name))
	
	return nil
}

// GetSubscribers returns the list of subscribers
func (s *Stream) GetSubscribers() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var ids []string
	for id := range s.subscribers {
		ids = append(ids, id)
	}
	
	return ids
}

// GetStats returns stream statistics
func (s *Stream) GetStats() StreamStats {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	return StreamStats{
		Name:         s.name,
		Subscribers:  len(s.subscribers),
		BufferSize:   s.config.BufferSize,
		FlushInterval: s.config.FlushInterval,
	}
}

// flush flushes the stream
func (s *Stream) flush() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// Flush reactive stream
	s.reactive.Flush()
	
	s.logger.Debug("Stream flushed", zap.String("name", s.name))
}

// StreamStats represents stream statistics
type StreamStats struct {
	Name          string
	Subscribers   int
	BufferSize    int
	FlushInterval time.Duration
}

// Stream errors
var (
	ErrMaxSubscribersReached = &StreamError{msg: "max subscribers reached"}
	ErrSubscriberNotFound    = &StreamError{msg: "subscriber not found"}
)

// StreamError represents a stream-related error
type StreamError struct {
	msg string
}

func (e *StreamError) Error() string {
	return e.msg
}

// StreamManager manages multiple streams
type StreamManager struct {
	streams map[string]*Stream
	mu      sync.RWMutex
	logger  *zap.Logger
}

// NewStreamManager creates a new stream manager
func NewStreamManager(logger *zap.Logger) *StreamManager {
	return &StreamManager{
		streams: make(map[string]*Stream),
		logger:  logger,
	}
}

// CreateStream creates a new stream
func (sm *StreamManager) CreateStream(name string, config StreamConfig) (*Stream, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if _, exists := sm.streams[name]; exists {
		return nil, ErrStreamManagerExists
	}
	
	stream := NewStream(name, config, sm.logger)
	sm.streams[name] = stream
	
	// Start stream
	go stream.Start(context.Background())
	
	return stream, nil
}

// GetStream returns a stream by name
func (sm *StreamManager) GetStream(name string) (*Stream, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	stream, exists := sm.streams[name]
	return stream, exists
}

// DeleteStream deletes a stream
func (sm *StreamManager) DeleteStream(name string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	stream, exists := sm.streams[name]
	if !exists {
		return ErrStreamManagerNotFound
	}
	
	stream.Stop()
	delete(sm.streams, name)
	
	return nil
}

// PublishToStream publishes an event to a specific stream
func (sm *StreamManager) PublishToStream(streamName string, event *models.Event) error {
	stream, exists := sm.GetStream(streamName)
	if !exists {
		return ErrStreamNotFound
	}
	
	return stream.Publish(event)
}

// PublishToAll publishes an event to all streams
func (sm *StreamManager) PublishToAll(event *models.Event) error {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	var errors []error
	for _, stream := range sm.streams {
		if err := stream.Publish(event); err != nil {
			errors = append(errors, err)
		}
	}
	
	if len(errors) > 0 {
		return ErrPublishFailed
	}
	
	return nil
}

// GetStats returns manager statistics
func (sm *StreamManager) GetStats() StreamManagerStats {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	var totalSubscribers int
	for _, stream := range sm.streams {
		totalSubscribers += len(stream.subscribers)
	}
	
	return StreamManagerStats{
		Streams:          len(sm.streams),
		TotalSubscribers: totalSubscribers,
	}
}

// StreamManagerStats represents stream manager statistics
type StreamManagerStats struct {
	Streams          int
	TotalSubscribers int
}

// Stream manager errors
var (
	ErrStreamManagerExists    = &StreamError{msg: "stream already exists"}
	ErrStreamManagerNotFound  = &StreamError{msg: "stream not found"}
	ErrPublishFailed   = &StreamError{msg: "publish failed"}
)
