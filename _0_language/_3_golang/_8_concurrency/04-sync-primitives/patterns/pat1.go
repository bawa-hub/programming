package patterns

import "sync"

// Advanced Pattern 1: Thread-Safe Counter with Metrics
type SafeCounter struct {
	mu       sync.RWMutex
	counters map[string]int64
	metrics  map[string]int64
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counters: make(map[string]int64),
		metrics:  make(map[string]int64),
	}
}

func (sc *SafeCounter) Increment(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counters[key]++
	sc.metrics["increments"]++
}

func (sc *SafeCounter) Get(key string) int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counters[key]
}

func (sc *SafeCounter) GetAllCounters() map[string]int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sc.counters {
		result[k] = v
	}
	return result
}

func (sc *SafeCounter) GetMetrics() map[string]int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sc.metrics {
		result[k] = v
	}
	return result
}