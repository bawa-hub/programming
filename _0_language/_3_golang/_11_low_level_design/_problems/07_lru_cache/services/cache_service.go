package services

import (
	"fmt"
	"sync"
	"time"

	"lru_cache/models"
)

// Cache Service for advanced cache operations
type CacheService struct {
	caches map[string]*models.LRUCache
	mu     sync.RWMutex
}

func NewCacheService() *CacheService {
	return &CacheService{
		caches: make(map[string]*models.LRUCache),
	}
}

// CreateCache creates a new LRU cache
func (cs *CacheService) CreateCache(name string, capacity int) (*models.LRUCache, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive")
	}
	
	if _, exists := cs.caches[name]; exists {
		return nil, fmt.Errorf("cache with name '%s' already exists", name)
	}
	
	cache := models.NewLRUCache(capacity)
	cs.caches[name] = cache
	return cache, nil
}

// GetCache retrieves a cache by name
func (cs *CacheService) GetCache(name string) *models.LRUCache {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.caches[name]
}

// DeleteCache removes a cache
func (cs *CacheService) DeleteCache(name string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	if cache, exists := cs.caches[name]; exists {
		cache.Clear()
		delete(cs.caches, name)
		return true
	}
	
	return false
}

// ListCaches returns all cache names
func (cs *CacheService) ListCaches() []string {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	
	names := make([]string, 0, len(cs.caches))
	for name := range cs.caches {
		names = append(names, name)
	}
	return names
}

// GetCacheStats returns statistics for all caches
func (cs *CacheService) GetCacheStats() map[string]interface{} {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	
	stats := make(map[string]interface{})
	for name, cache := range cs.caches {
		stats[name] = cache.GetCacheInfo()
	}
	return stats
}

// WarmUpCache preloads cache with data
func (cs *CacheService) WarmUpCache(name string, data map[string]interface{}) error {
	cs.mu.RLock()
	cache := cs.caches[name]
	cs.mu.RUnlock()
	
	if cache == nil {
		return fmt.Errorf("cache '%s' not found", name)
	}
	
	for key, value := range data {
		if cache.Size() < cache.Capacity() {
			cache.Put(key, value)
		}
	}
	
	return nil
}

// BatchPut performs batch put operations
func (cs *CacheService) BatchPut(cacheName string, items map[string]interface{}) error {
	cs.mu.RLock()
	cache := cs.caches[cacheName]
	cs.mu.RUnlock()
	
	if cache == nil {
		return fmt.Errorf("cache '%s' not found", cacheName)
	}
	
	for key, value := range items {
		cache.Put(key, value)
	}
	
	return nil
}

// BatchGet performs batch get operations
func (cs *CacheService) BatchGet(cacheName string, keys []string) map[string]interface{} {
	cs.mu.RLock()
	cache := cs.caches[cacheName]
	cs.mu.RUnlock()
	
	if cache == nil {
		return nil
	}
	
	result := make(map[string]interface{})
	for _, key := range keys {
		if value, exists := cache.Get(key); exists {
			result[key] = value
		}
	}
	
	return result
}

// Cache Monitor for real-time monitoring
type CacheMonitor struct {
	service   *CacheService
	metrics   chan CacheMetric
	stopChan  chan bool
	interval  time.Duration
}

type CacheMetric struct {
	Timestamp time.Time
	CacheName string
	HitRate   float64
	Size      int
	Capacity  int
	Evictions int64
}

func NewCacheMonitor(service *CacheService, interval time.Duration) *CacheMonitor {
	return &CacheMonitor{
		service:  service,
		metrics:  make(chan CacheMetric, 100),
		stopChan: make(chan bool),
		interval: interval,
	}
}

func (cm *CacheMonitor) Start() {
	go cm.collectMetrics()
}

func (cm *CacheMonitor) Stop() {
	cm.stopChan <- true
}

func (cm *CacheMonitor) GetMetrics() <-chan CacheMetric {
	return cm.metrics
}

func (cm *CacheMonitor) collectMetrics() {
	ticker := time.NewTicker(cm.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			cm.collectCacheMetrics()
		case <-cm.stopChan:
			return
		}
	}
}

func (cm *CacheMonitor) collectCacheMetrics() {
	cacheNames := cm.service.ListCaches()
	
	for _, name := range cacheNames {
		cache := cm.service.GetCache(name)
		if cache != nil {
			metric := CacheMetric{
				Timestamp: time.Now(),
				CacheName: name,
				HitRate:   cache.HitRate(),
				Size:      cache.Size(),
				Capacity:  cache.Capacity(),
				Evictions: cache.GetStats().GetEvictions(),
			}
			
			select {
			case cm.metrics <- metric:
			default:
				// Channel full, skip this metric
			}
		}
	}
}

// Cache Load Balancer for distributed caching
type CacheLoadBalancer struct {
	caches []*models.LRUCache
	index  int
	mu     sync.Mutex
}

func NewCacheLoadBalancer(caches []*models.LRUCache) *CacheLoadBalancer {
	return &CacheLoadBalancer{
		caches: caches,
		index:  0,
	}
}

func (clb *CacheLoadBalancer) Get(key string) (interface{}, bool) {
	for _, cache := range clb.caches {
		if value, exists := cache.Get(key); exists {
			return value, true
		}
	}
	return nil, false
}

func (clb *CacheLoadBalancer) Put(key string, value interface{}) {
	clb.mu.Lock()
	defer clb.mu.Unlock()
	
	// Round-robin distribution
	cache := clb.caches[clb.index]
	cache.Put(key, value)
	clb.index = (clb.index + 1) % len(clb.caches)
}

func (clb *CacheLoadBalancer) Delete(key string) bool {
	for _, cache := range clb.caches {
		if cache.Delete(key) {
			return true
		}
	}
	return false
}

func (clb *CacheLoadBalancer) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	
	for i, cache := range clb.caches {
		stats[fmt.Sprintf("cache_%d", i)] = cache.GetCacheInfo()
	}
	
	return stats
}
