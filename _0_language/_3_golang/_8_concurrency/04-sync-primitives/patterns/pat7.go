package patterns

import "sync"

// Advanced Pattern 7: Concurrent Map with Statistics
type StatsMap struct {
	mu     sync.RWMutex
	data   map[string]interface{}
	stats  map[string]int64
}

func NewStatsMap() *StatsMap {
	return &StatsMap{
		data:  make(map[string]interface{}),
		stats: make(map[string]int64),
	}
}

func (sm *StatsMap) Store(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
	sm.stats["stores"]++
}

func (sm *StatsMap) Load(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	if ok {
		sm.stats["loads"]++
	} else {
		sm.stats["misses"]++
	}
	return value, ok
}

func (sm *StatsMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, exists := sm.data[key]; exists {
		delete(sm.data, key)
		sm.stats["deletes"]++
	}
}

func (sm *StatsMap) GetStats() map[string]int64 {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sm.stats {
		result[k] = v
	}
	return result
}