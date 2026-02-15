package models

import (
	"fmt"
	"sync"
)

// LRU Cache implementation
type LRUCache struct {
	capacity int
	size     int
	cache    map[string]*Node
	head     *Node
	tail     *Node
	stats    *CacheStats
	mu       sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	
	return &LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[string]*Node),
		head:     nil,
		tail:     nil,
		stats:    NewCacheStats(),
	}
}

// Get retrieves a value by key
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if node, exists := c.cache[key]; exists {
		// Move to head (most recently used)
		c.moveToHead(node)
		c.stats.RecordHit()
		return node.GetValue(), true
	}
	
	c.stats.RecordMiss()
	return nil, false
}

// Put inserts or updates a key-value pair
func (c *LRUCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if node, exists := c.cache[key]; exists {
		// Update existing node
		node.SetValue(value)
		c.moveToHead(node)
		return
	}
	
	// Create new node
	newNode := NewNode(key, value)
	
	if c.size >= c.capacity {
		// Evict least recently used
		c.evictLRU()
	}
	
	// Add to head
	c.addToHead(newNode)
	c.cache[key] = newNode
	c.size++
}

// Delete removes a key-value pair
func (c *LRUCache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if node, exists := c.cache[key]; exists {
		c.removeNode(node)
		delete(c.cache, key)
		c.size--
		return true
	}
	
	return false
}

// Clear removes all items from cache
func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.cache = make(map[string]*Node)
	c.head = nil
	c.tail = nil
	c.size = 0
}

// Size returns current cache size
func (c *LRUCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.size
}

// Capacity returns cache capacity
func (c *LRUCache) Capacity() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.capacity
}

// IsFull checks if cache is at capacity
func (c *LRUCache) IsFull() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.size >= c.capacity
}

// IsEmpty checks if cache is empty
func (c *LRUCache) IsEmpty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.size == 0
}

// Contains checks if key exists in cache
func (c *LRUCache) Contains(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.cache[key]
	return exists
}

// Keys returns all keys in cache
func (c *LRUCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]string, 0, c.size)
	for key := range c.cache {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values in cache
func (c *LRUCache) Values() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	values := make([]interface{}, 0, c.size)
	for _, node := range c.cache {
		values = append(values, node.GetValue())
	}
	return values
}

// GetStats returns cache statistics
func (c *LRUCache) GetStats() *CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stats
}

// HitRate returns cache hit rate
func (c *LRUCache) HitRate() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stats.HitRate()
}

// MissRate returns cache miss rate
func (c *LRUCache) MissRate() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stats.MissRate()
}

// SetCapacity changes cache capacity
func (c *LRUCache) SetCapacity(newCapacity int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if newCapacity <= 0 {
		return fmt.Errorf("capacity must be positive")
	}
	
	oldCapacity := c.capacity
	c.capacity = newCapacity
	
	// If new capacity is smaller, evict excess items
	if newCapacity < oldCapacity && c.size > newCapacity {
		evictCount := c.size - newCapacity
		for i := 0; i < evictCount; i++ {
			c.evictLRU()
		}
	}
	
	return nil
}

// Resize changes cache capacity and evicts if necessary
func (c *LRUCache) Resize(newCapacity int) error {
	return c.SetCapacity(newCapacity)
}

// GetLRUOrder returns keys in LRU order (most recent first)
func (c *LRUCache) GetLRUOrder() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]string, 0, c.size)
	current := c.head
	
	for current != nil {
		keys = append(keys, current.GetKey())
		current = current.GetNext()
	}
	
	return keys
}

// Private helper methods

func (c *LRUCache) addToHead(node *Node) {
	node.SetPrev(nil)
	node.SetNext(c.head)
	
	if c.head != nil {
		c.head.SetPrev(node)
	}
	
	c.head = node
	
	if c.tail == nil {
		c.tail = node
	}
}

func (c *LRUCache) removeNode(node *Node) {
	if node.GetPrev() != nil {
		node.GetPrev().SetNext(node.GetNext())
	} else {
		c.head = node.GetNext()
	}
	
	if node.GetNext() != nil {
		node.GetNext().SetPrev(node.GetPrev())
	} else {
		c.tail = node.GetPrev()
	}
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) evictLRU() {
	if c.tail == nil {
		return
	}
	
	// Remove from tail
	lastNode := c.tail
	c.removeNode(lastNode)
	delete(c.cache, lastNode.GetKey())
	c.size--
	c.stats.RecordEviction()
}

// GetCacheInfo returns comprehensive cache information
func (c *LRUCache) GetCacheInfo() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return map[string]interface{}{
		"capacity":     c.capacity,
		"size":         c.size,
		"is_full":      c.size >= c.capacity,
		"is_empty":     c.size == 0,
		"hit_rate":     c.stats.HitRate(),
		"miss_rate":    c.stats.MissRate(),
		"total_hits":   c.stats.GetHits(),
		"total_misses": c.stats.GetMisses(),
		"evictions":    c.stats.GetEvictions(),
		"lru_order":    c.GetLRUOrder(),
	}
}
