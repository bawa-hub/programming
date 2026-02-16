package patterns

import (
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 6: Goroutine with Metrics
type MetricsGoroutine struct {
	processed int64
	errors    int64
	startTime time.Time
	mu        sync.RWMutex
}

func NewMetricsGoroutine() *MetricsGoroutine {
	return &MetricsGoroutine{
		startTime: time.Now(),
	}
}

func (m *MetricsGoroutine) RecordProcessed() {
	atomic.AddInt64(&m.processed, 1)
}

func (m *MetricsGoroutine) RecordError() {
	atomic.AddInt64(&m.errors, 1)
}

func (m *MetricsGoroutine) GetStats() (processed, errors int64, uptime time.Duration) {
	processed = atomic.LoadInt64(&m.processed)
	errors = atomic.LoadInt64(&m.errors)
	uptime = time.Since(m.startTime)
	return
}