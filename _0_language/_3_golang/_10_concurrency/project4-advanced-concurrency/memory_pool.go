package main

import (
	"fmt"
	"sync"
	"time"
)

// PooledObject represents an object that can be pooled
type PooledObject struct {
	ID        int
	Data      interface{}
	CreatedAt time.Time
	UsedAt    time.Time
	pool      *MemoryPool
}

// Reset resets the object for reuse
func (po *PooledObject) Reset() {
	po.Data = nil
	po.UsedAt = time.Now()
}

// MemoryPool represents a memory pool for objects
type MemoryPool struct {
	pool    chan *PooledObject
	factory func() *PooledObject
	maxSize int
	current int64
	mutex   sync.RWMutex
	closed  bool
}

// NewMemoryPool creates a new memory pool
func NewMemoryPool(maxSize int) *MemoryPool {
	mp := &MemoryPool{
		pool:    make(chan *PooledObject, maxSize),
		maxSize: maxSize,
		factory: func() *PooledObject {
			return &PooledObject{
				CreatedAt: time.Now(),
				UsedAt:    time.Now(),
			}
		},
	}
	
	// Pre-populate pool
	for i := 0; i < maxSize/2; i++ {
		obj := mp.factory()
		obj.ID = i
		obj.pool = mp
		mp.pool <- obj
		mp.current++
	}
	
	return mp
}

// Get gets an object from the pool
func (mp *MemoryPool) Get() *PooledObject {
	mp.mutex.RLock()
	if mp.closed {
		mp.mutex.RUnlock()
		return nil
	}
	mp.mutex.RUnlock()
	
	select {
	case obj := <-mp.pool:
		obj.Reset()
		return obj
	default:
		// Pool is empty, create new object
		mp.mutex.Lock()
		if mp.current < int64(mp.maxSize) {
			obj := mp.factory()
			obj.ID = int(mp.current)
			obj.pool = mp
			mp.current++
			mp.mutex.Unlock()
			obj.Reset()
			return obj
		}
		mp.mutex.Unlock()
		
		// Wait for an object to become available
		select {
		case obj := <-mp.pool:
			obj.Reset()
			return obj
		case <-time.After(5 * time.Second):
			// Timeout, create temporary object
			obj := mp.factory()
			obj.ID = -1 // Temporary object
			obj.pool = mp
			obj.Reset()
			return obj
		}
	}
}

// Put returns an object to the pool
func (mp *MemoryPool) Put(obj *PooledObject) {
	if obj == nil {
		return
	}
	
	mp.mutex.RLock()
	if mp.closed {
		mp.mutex.RUnlock()
		return
	}
	mp.mutex.RUnlock()
	
	// Only return objects that belong to this pool
	if obj.pool != mp {
		return
	}
	
	select {
	case mp.pool <- obj:
		// Object returned to pool
	default:
		// Pool is full, discard object
	}
}

// Size returns the current size of the pool
func (mp *MemoryPool) Size() int {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	return len(mp.pool)
}

// Close closes the memory pool
func (mp *MemoryPool) Close() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	if mp.closed {
		return
	}
	
	mp.closed = true
	close(mp.pool)
	
	// Clear all objects
	for obj := range mp.pool {
		obj.pool = nil
	}
}

// Stats returns pool statistics
func (mp *MemoryPool) Stats() PoolStats {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	
	return PoolStats{
		Current:    len(mp.pool),
		MaxSize:    mp.maxSize,
		TotalCreated: int(mp.current),
		Available:  len(mp.pool),
	}
}

// PoolStats represents pool statistics
type PoolStats struct {
	Current      int
	MaxSize      int
	TotalCreated int
	Available    int
}

// String returns a string representation of pool stats
func (ps PoolStats) String() string {
	return fmt.Sprintf("PoolStats{Current: %d, MaxSize: %d, TotalCreated: %d, Available: %d}",
		ps.Current, ps.MaxSize, ps.TotalCreated, ps.Available)
}

// AdvancedMemoryPool represents an advanced memory pool with metrics
type AdvancedMemoryPool struct {
	pool      chan *PooledObject
	factory   func() *PooledObject
	maxSize   int
	current   int64
	mutex     sync.RWMutex
	closed    bool
	metrics   *PoolMetrics
	cleanup   *CleanupManager
}

// NewAdvancedMemoryPool creates a new advanced memory pool
func NewAdvancedMemoryPool(maxSize int) *AdvancedMemoryPool {
	amp := &AdvancedMemoryPool{
		pool:    make(chan *PooledObject, maxSize),
		maxSize: maxSize,
		factory: func() *PooledObject {
			return &PooledObject{
				CreatedAt: time.Now(),
				UsedAt:    time.Now(),
			}
		},
		metrics: NewPoolMetrics(),
		cleanup: NewCleanupManager(),
	}
	
	// Pre-populate pool
	for i := 0; i < maxSize/2; i++ {
		obj := amp.factory()
		obj.ID = i
		obj.pool = &MemoryPool{} // Dummy pool reference
		amp.pool <- obj
		amp.current++
	}
	
	// Start cleanup manager
	amp.cleanup.Start()
	
	return amp
}

// Get gets an object from the advanced pool
func (amp *AdvancedMemoryPool) Get() *PooledObject {
	amp.mutex.RLock()
	if amp.closed {
		amp.mutex.RUnlock()
		return nil
	}
	amp.mutex.RUnlock()
	
	start := time.Now()
	
	select {
	case obj := <-amp.pool:
		obj.Reset()
		amp.metrics.RecordGet(time.Since(start))
		return obj
	default:
		// Pool is empty, create new object
		amp.mutex.Lock()
		if amp.current < int64(amp.maxSize) {
			obj := amp.factory()
			obj.ID = int(amp.current)
			obj.pool = &MemoryPool{} // Dummy pool reference
			amp.current++
			amp.mutex.Unlock()
			obj.Reset()
			amp.metrics.RecordGet(time.Since(start))
			return obj
		}
		amp.mutex.Unlock()
		
		// Wait for an object to become available
		select {
		case obj := <-amp.pool:
			obj.Reset()
			amp.metrics.RecordGet(time.Since(start))
			return obj
		case <-time.After(5 * time.Second):
			// Timeout, create temporary object
			obj := amp.factory()
			obj.ID = -1 // Temporary object
			obj.pool = &MemoryPool{} // Dummy pool reference
			obj.Reset()
			amp.metrics.RecordGet(time.Since(start))
			return obj
		}
	}
}

// Put returns an object to the advanced pool
func (amp *AdvancedMemoryPool) Put(obj *PooledObject) {
	if obj == nil {
		return
	}
	
	amp.mutex.RLock()
	if amp.closed {
		amp.mutex.RUnlock()
		return
	}
	amp.mutex.RUnlock()
	
	start := time.Now()
	
	select {
	case amp.pool <- obj:
		amp.metrics.RecordPut(time.Since(start))
	default:
		// Pool is full, discard object
		amp.metrics.RecordDiscard()
	}
}

// Close closes the advanced memory pool
func (amp *AdvancedMemoryPool) Close() {
	amp.mutex.Lock()
	defer amp.mutex.Unlock()
	
	if amp.closed {
		return
	}
	
	amp.closed = true
	close(amp.pool)
	
	// Stop cleanup manager
	amp.cleanup.Stop()
	
	// Clear all objects
	for obj := range amp.pool {
		obj.pool = nil
	}
}

// Stats returns advanced pool statistics
func (amp *AdvancedMemoryPool) Stats() AdvancedPoolStats {
	amp.mutex.RLock()
	defer amp.mutex.RUnlock()
	
	return AdvancedPoolStats{
		PoolStats: PoolStats{
			Current:      len(amp.pool),
			MaxSize:      amp.maxSize,
			TotalCreated: int(amp.current),
			Available:    len(amp.pool),
		},
		Metrics: amp.metrics.GetStats(),
	}
}

// AdvancedPoolStats represents advanced pool statistics
type AdvancedPoolStats struct {
	PoolStats
	Metrics PoolMetricsStats
}

// String returns a string representation of advanced pool stats
func (aps AdvancedPoolStats) String() string {
	return fmt.Sprintf("AdvancedPoolStats{%s, Metrics: %s}",
		aps.PoolStats.String(), aps.Metrics.String())
}

// PoolMetrics tracks pool metrics
type PoolMetrics struct {
	gets      int64
	puts      int64
	discards  int64
	getTime   time.Duration
	putTime   time.Duration
	mutex     sync.RWMutex
}

// NewPoolMetrics creates a new pool metrics
func NewPoolMetrics() *PoolMetrics {
	return &PoolMetrics{}
}

// RecordGet records a get operation
func (pm *PoolMetrics) RecordGet(duration time.Duration) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.gets++
	pm.getTime += duration
}

// RecordPut records a put operation
func (pm *PoolMetrics) RecordPut(duration time.Duration) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.puts++
	pm.putTime += duration
}

// RecordDiscard records a discard operation
func (pm *PoolMetrics) RecordDiscard() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.discards++
}

// GetStats returns the current metrics
func (pm *PoolMetrics) GetStats() PoolMetricsStats {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	
	avgGetTime := time.Duration(0)
	if pm.gets > 0 {
		avgGetTime = pm.getTime / time.Duration(pm.gets)
	}
	
	avgPutTime := time.Duration(0)
	if pm.puts > 0 {
		avgPutTime = pm.putTime / time.Duration(pm.puts)
	}
	
	return PoolMetricsStats{
		Gets:        pm.gets,
		Puts:        pm.puts,
		Discards:    pm.discards,
		AvgGetTime:  avgGetTime,
		AvgPutTime:  avgPutTime,
	}
}

// PoolMetricsStats represents pool metrics statistics
type PoolMetricsStats struct {
	Gets       int64
	Puts       int64
	Discards   int64
	AvgGetTime time.Duration
	AvgPutTime time.Duration
}

// String returns a string representation of pool metrics stats
func (pms PoolMetricsStats) String() string {
	return fmt.Sprintf("PoolMetricsStats{Gets: %d, Puts: %d, Discards: %d, AvgGetTime: %v, AvgPutTime: %v}",
		pms.Gets, pms.Puts, pms.Discards, pms.AvgGetTime, pms.AvgPutTime)
}

// CleanupManager manages cleanup of pooled objects
type CleanupManager struct {
	ticker *time.Ticker
	done   chan struct{}
	wg     sync.WaitGroup
}

// NewCleanupManager creates a new cleanup manager
func NewCleanupManager() *CleanupManager {
	return &CleanupManager{
		ticker: time.NewTicker(1 * time.Minute),
		done:   make(chan struct{}),
	}
}

// Start starts the cleanup manager
func (cm *CleanupManager) Start() {
	cm.wg.Add(1)
	go cm.run()
}

// run runs the cleanup manager
func (cm *CleanupManager) run() {
	defer cm.wg.Done()
	
	for {
		select {
		case <-cm.ticker.C:
			// Perform cleanup
			cm.cleanup()
		case <-cm.done:
			return
		}
	}
}

// cleanup performs cleanup operations
func (cm *CleanupManager) cleanup() {
	// This is where you would implement cleanup logic
	// For example, removing old objects from the pool
	fmt.Println("Performing cleanup...")
}

// Stop stops the cleanup manager
func (cm *CleanupManager) Stop() {
	cm.ticker.Stop()
	close(cm.done)
	cm.wg.Wait()
}

// ObjectPool represents a generic object pool
type ObjectPool struct {
	pool    chan interface{}
	factory func() interface{}
	reset   func(interface{})
	maxSize int
	current int64
	mutex   sync.RWMutex
	closed  bool
}

// NewObjectPool creates a new object pool
func NewObjectPool(maxSize int, factory func() interface{}, reset func(interface{})) *ObjectPool {
	return &ObjectPool{
		pool:    make(chan interface{}, maxSize),
		factory: factory,
		reset:   reset,
		maxSize: maxSize,
	}
}

// Get gets an object from the pool
func (op *ObjectPool) Get() interface{} {
	op.mutex.RLock()
	if op.closed {
		op.mutex.RUnlock()
		return nil
	}
	op.mutex.RUnlock()
	
	select {
	case obj := <-op.pool:
		if op.reset != nil {
			op.reset(obj)
		}
		return obj
	default:
		// Pool is empty, create new object
		op.mutex.Lock()
		if op.current < int64(op.maxSize) {
			op.current++
			op.mutex.Unlock()
			return op.factory()
		}
		op.mutex.Unlock()
		
		// Wait for an object to become available
		select {
		case obj := <-op.pool:
			if op.reset != nil {
				op.reset(obj)
			}
			return obj
		case <-time.After(5 * time.Second):
			// Timeout, create temporary object
			return op.factory()
		}
	}
}

// Put returns an object to the pool
func (op *ObjectPool) Put(obj interface{}) {
	if obj == nil {
		return
	}
	
	op.mutex.RLock()
	if op.closed {
		op.mutex.RUnlock()
		return
	}
	op.mutex.RUnlock()
	
	select {
	case op.pool <- obj:
		// Object returned to pool
	default:
		// Pool is full, discard object
	}
}

// Close closes the object pool
func (op *ObjectPool) Close() {
	op.mutex.Lock()
	defer op.mutex.Unlock()
	
	if op.closed {
		return
	}
	
	op.closed = true
	close(op.pool)
	
	// Clear all objects
	for obj := range op.pool {
		_ = obj // Discard
	}
}

// Size returns the current size of the pool
func (op *ObjectPool) Size() int {
	op.mutex.RLock()
	defer op.mutex.RUnlock()
	return len(op.pool)
}
