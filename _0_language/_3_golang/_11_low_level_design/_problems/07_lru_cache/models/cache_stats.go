package models

import (
	"sync"
	"time"
)

// Cache Statistics
type CacheStats struct {
	hits          int64
	misses        int64
	totalRequests int64
	evictions     int64
	lastAccess    time.Time
	mu            sync.RWMutex
}

func NewCacheStats() *CacheStats {
	return &CacheStats{
		hits:          0,
		misses:        0,
		totalRequests: 0,
		evictions:     0,
		lastAccess:    time.Now(),
	}
}

func (s *CacheStats) RecordHit() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hits++
	s.totalRequests++
	s.lastAccess = time.Now()
}

func (s *CacheStats) RecordMiss() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.misses++
	s.totalRequests++
	s.lastAccess = time.Now()
}

func (s *CacheStats) RecordEviction() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.evictions++
}

func (s *CacheStats) GetHits() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.hits
}

func (s *CacheStats) GetMisses() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.misses
}

func (s *CacheStats) GetTotalRequests() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.totalRequests
}

func (s *CacheStats) GetEvictions() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.evictions
}

func (s *CacheStats) GetLastAccess() time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastAccess
}

func (s *CacheStats) HitRate() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if s.totalRequests == 0 {
		return 0.0
	}
	return float64(s.hits) / float64(s.totalRequests)
}

func (s *CacheStats) MissRate() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if s.totalRequests == 0 {
		return 0.0
	}
	return float64(s.misses) / float64(s.totalRequests)
}

func (s *CacheStats) GetStats() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	return map[string]interface{}{
		"hits":           s.hits,
		"misses":         s.misses,
		"total_requests": s.totalRequests,
		"evictions":      s.evictions,
		"hit_rate":       s.HitRate(),
		"miss_rate":      s.MissRate(),
		"last_access":    s.lastAccess,
	}
}

func (s *CacheStats) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hits = 0
	s.misses = 0
	s.totalRequests = 0
	s.evictions = 0
	s.lastAccess = time.Now()
}
