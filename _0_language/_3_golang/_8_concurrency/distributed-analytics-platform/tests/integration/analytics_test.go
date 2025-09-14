package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/your-username/distributed-analytics-platform/internal/analytics"
	"github.com/your-username/distributed-analytics-platform/pkg/models"
)

func TestAnalyticsEngine(t *testing.T) {
	// Create logger
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	// Create analytics engine
	config := analytics.Config{
		Workers:      5,
		BatchSize:    100,
		BatchTimeout: 1 * time.Second,
		CacheSize:    1000,
	}

	engine, err := analytics.NewEngine(config, logger)
	require.NoError(t, err)
	defer engine.Shutdown(context.Background())

	// Test single event processing
	t.Run("ProcessSingleEvent", func(t *testing.T) {
		event := models.NewEvent("page_view", "test", map[string]interface{}{
			"page": "/test-page",
			"user_id": "user-123",
		})

		err := engine.ProcessEvent(event)
		assert.NoError(t, err)
	})

	// Test batch event processing
	t.Run("ProcessBatchEvents", func(t *testing.T) {
		events := make([]*models.Event, 10)
		for i := 0; i < 10; i++ {
			events[i] = models.NewEvent("click", "test", map[string]interface{}{
				"element": "button",
				"page": "/test-page",
			})
		}

		err := engine.ProcessBatch(events)
		assert.NoError(t, err)
	})

	// Test metrics query
	t.Run("QueryMetrics", func(t *testing.T) {
		query := &analytics.MetricQuery{
			StartTime: time.Now().Add(-1 * time.Hour),
			EndTime:   time.Now(),
			Limit:     100,
		}

		metrics, err := engine.GetMetrics(query)
		assert.NoError(t, err)
		assert.NotNil(t, metrics)
	})

	// Test stream creation
	t.Run("CreateStream", func(t *testing.T) {
		streamConfig := analytics.StreamConfig{
			BufferSize:     1000,
			FlushInterval:  5 * time.Second,
			MaxSubscribers: 10,
		}

		stream, err := engine.CreateStream("test-stream", streamConfig)
		assert.NoError(t, err)
		assert.NotNil(t, stream)

		// Clean up
		err = engine.DeleteStream("test-stream")
		assert.NoError(t, err)
	})

	// Test engine statistics
	t.Run("GetStats", func(t *testing.T) {
		stats := engine.GetStats()
		assert.NotNil(t, stats)
		assert.GreaterOrEqual(t, stats.Workers, 0)
	})
}

func TestConcurrentEventProcessing(t *testing.T) {
	// Create logger
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	// Create analytics engine
	config := analytics.Config{
		Workers:      10,
		BatchSize:    50,
		BatchTimeout: 100 * time.Millisecond,
		CacheSize:    1000,
	}

	engine, err := analytics.NewEngine(config, logger)
	require.NoError(t, err)
	defer engine.Shutdown(context.Background())

	// Test concurrent event processing
	t.Run("ConcurrentEvents", func(t *testing.T) {
		const numGoroutines = 10
		const eventsPerGoroutine = 100

		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(goroutineID int) {
				defer func() { done <- true }()

				for j := 0; j < eventsPerGoroutine; j++ {
					event := models.NewEvent("page_view", "test", map[string]interface{}{
						"page":     "/test-page",
						"user_id":  "user-123",
						"goroutine": goroutineID,
						"event":    j,
					})

					err := engine.ProcessEvent(event)
					assert.NoError(t, err)
				}
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		// Verify all events were processed
		time.Sleep(2 * time.Second) // Allow time for processing

		stats := engine.GetStats()
		assert.GreaterOrEqual(t, stats.Events, numGoroutines*eventsPerGoroutine)
	})
}

func TestEventValidation(t *testing.T) {
	// Create logger
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	// Create analytics engine
	config := analytics.Config{
		Workers:      5,
		BatchSize:    100,
		BatchTimeout: 1 * time.Second,
		CacheSize:    1000,
	}

	engine, err := analytics.NewEngine(config, logger)
	require.NoError(t, err)
	defer engine.Shutdown(context.Background())

	// Test invalid event
	t.Run("InvalidEvent", func(t *testing.T) {
		event := &models.Event{
			// Missing required fields
		}

		err := engine.ProcessEvent(event)
		assert.Error(t, err)
	})

	// Test valid event
	t.Run("ValidEvent", func(t *testing.T) {
		event := models.NewEvent("page_view", "test", map[string]interface{}{
			"page": "/test-page",
		})

		err := engine.ProcessEvent(event)
		assert.NoError(t, err)
	})
}

func TestStreamProcessing(t *testing.T) {
	// Create logger
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	// Create analytics engine
	config := analytics.Config{
		Workers:      5,
		BatchSize:    100,
		BatchTimeout: 1 * time.Second,
		CacheSize:    1000,
	}

	engine, err := analytics.NewEngine(config, logger)
	require.NoError(t, err)
	defer engine.Shutdown(context.Background())

	// Create stream
	streamConfig := analytics.StreamConfig{
		BufferSize:     100,
		FlushInterval:  100 * time.Millisecond,
		MaxSubscribers: 5,
	}

	stream, err := engine.CreateStream("test-stream", streamConfig)
	require.NoError(t, err)

	// Test stream publishing
	t.Run("PublishToStream", func(t *testing.T) {
		event := models.NewEvent("page_view", "test", map[string]interface{}{
			"page": "/test-page",
		})

		err := stream.Publish(event)
		assert.NoError(t, err)
	})

	// Test stream subscription
	t.Run("SubscribeToStream", func(t *testing.T) {
		subscriber, err := stream.Subscribe("test-subscriber", func(event *models.Event) bool {
			return event.Type == "page_view"
		})
		require.NoError(t, err)
		assert.NotNil(t, subscriber)

		// Clean up
		err = stream.Unsubscribe("test-subscriber")
		assert.NoError(t, err)
	})

	// Clean up stream
	err = engine.DeleteStream("test-stream")
	assert.NoError(t, err)
}

func TestPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	// Create logger
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	// Create analytics engine
	config := analytics.Config{
		Workers:      20,
		BatchSize:    1000,
		BatchTimeout: 10 * time.Millisecond,
		CacheSize:    10000,
	}

	engine, err := analytics.NewEngine(config, logger)
	require.NoError(t, err)
	defer engine.Shutdown(context.Background())

	// Test high-throughput event processing
	t.Run("HighThroughput", func(t *testing.T) {
		const numEvents = 10000
		start := time.Now()

		for i := 0; i < numEvents; i++ {
			event := models.NewEvent("page_view", "test", map[string]interface{}{
				"page":     "/test-page",
				"user_id":  "user-123",
				"event_id": i,
			})

			err := engine.ProcessEvent(event)
			assert.NoError(t, err)
		}

		duration := time.Since(start)
		throughput := float64(numEvents) / duration.Seconds()

		t.Logf("Processed %d events in %v (%.2f events/sec)", 
			numEvents, duration, throughput)

		// Verify reasonable throughput (should be > 1000 events/sec)
		assert.Greater(t, throughput, 1000.0)
	})
}
