package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// ConcurrentMap represents a thread-safe map
type ConcurrentMap struct {
	shards []*MapShard
	size   int
}

// MapShard represents a shard of the map
type MapShard struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// NewConcurrentMap creates a new concurrent map
func NewConcurrentMap() *ConcurrentMap {
	shards := make([]*MapShard, 16) // 16 shards for better concurrency
	for i := range shards {
		shards[i] = &MapShard{
			data: make(map[string]interface{}),
		}
	}
	return &ConcurrentMap{
		shards: shards,
		size:   16,
	}
}

// getShard returns the shard for a given key
func (cm *ConcurrentMap) getShard(key string) *MapShard {
	hash := fnv1a(key)
	return cm.shards[hash%uint32(len(cm.shards))]
}

// Set sets a value in the map
func (cm *ConcurrentMap) Set(key string, value interface{}) {
	shard := cm.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.data[key] = value
}

// Get gets a value from the map
func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := cm.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	value, exists := shard.data[key]
	return value, exists
}

// Delete deletes a value from the map
func (cm *ConcurrentMap) Delete(key string) {
	shard := cm.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	delete(shard.data, key)
}

// Size returns the size of the map
func (cm *ConcurrentMap) Size() int {
	total := 0
	for _, shard := range cm.shards {
		shard.mu.RLock()
		total += len(shard.data)
		shard.mu.RUnlock()
	}
	return total
}

// fnv1a implements FNV-1a hash function
func fnv1a(s string) uint32 {
	var hash uint32 = 2166136261
	for _, c := range s {
		hash ^= uint32(c)
		hash *= 16777619
	}
	return hash
}

// ConcurrentQueue represents a thread-safe queue
type ConcurrentQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// queueNode represents a node in the queue
type queueNode struct {
	value interface{}
	next  unsafe.Pointer
}

// NewConcurrentQueue creates a new concurrent queue
func NewConcurrentQueue() *ConcurrentQueue {
	dummy := &queueNode{}
	return &ConcurrentQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

// Enqueue adds an item to the queue
func (cq *ConcurrentQueue) Enqueue(value interface{}) {
	node := &queueNode{value: value}
	
	for {
		tail := (*queueNode)(atomic.LoadPointer(&cq.tail))
		next := (*queueNode)(atomic.LoadPointer(&tail.next))
		
		if tail == (*queueNode)(atomic.LoadPointer(&cq.tail)) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&tail.next, unsafe.Pointer(next), unsafe.Pointer(node)) {
					break
				}
			} else {
				atomic.CompareAndSwapPointer(&cq.tail, unsafe.Pointer(tail), unsafe.Pointer(next))
			}
		}
	}
	
	atomic.CompareAndSwapPointer(&cq.tail, unsafe.Pointer((*queueNode)(atomic.LoadPointer(&cq.tail))), unsafe.Pointer(node))
}

// Dequeue removes an item from the queue
func (cq *ConcurrentQueue) Dequeue() (interface{}, bool) {
	for {
		head := (*queueNode)(atomic.LoadPointer(&cq.head))
		tail := (*queueNode)(atomic.LoadPointer(&cq.tail))
		next := (*queueNode)(atomic.LoadPointer(&head.next))
		
		if head == (*queueNode)(atomic.LoadPointer(&cq.head)) {
			if head == tail {
				if next == nil {
					return nil, false
				}
				atomic.CompareAndSwapPointer(&cq.tail, unsafe.Pointer(tail), unsafe.Pointer(next))
			} else {
				if next == nil {
					continue
				}
				value := next.value
				if atomic.CompareAndSwapPointer(&cq.head, unsafe.Pointer(head), unsafe.Pointer(next)) {
					return value, true
				}
			}
		}
	}
}

// ConcurrentStack represents a thread-safe stack
type ConcurrentStack struct {
	top unsafe.Pointer
}

// stackNode represents a node in the stack
type stackNode struct {
	value interface{}
	next  unsafe.Pointer
}

// NewConcurrentStack creates a new concurrent stack
func NewConcurrentStack() *ConcurrentStack {
	return &ConcurrentStack{}
}

// Push adds an item to the stack
func (cs *ConcurrentStack) Push(value interface{}) {
	node := &stackNode{value: value}
	
	for {
		top := (*stackNode)(atomic.LoadPointer(&cs.top))
		node.next = unsafe.Pointer(top)
		if atomic.CompareAndSwapPointer(&cs.top, unsafe.Pointer(top), unsafe.Pointer(node)) {
			break
		}
	}
}

// Pop removes an item from the stack
func (cs *ConcurrentStack) Pop() (interface{}, bool) {
	for {
		top := (*stackNode)(atomic.LoadPointer(&cs.top))
		if top == nil {
			return nil, false
		}
		
		next := (*stackNode)(atomic.LoadPointer(&top.next))
		if atomic.CompareAndSwapPointer(&cs.top, unsafe.Pointer(top), unsafe.Pointer(next)) {
			return top.value, true
		}
	}
}

// ConcurrentRingBuffer represents a thread-safe ring buffer
type ConcurrentRingBuffer struct {
	buffer []interface{}
	size   int
	head   int64
	tail   int64
	mask   int64
}

// NewConcurrentRingBuffer creates a new concurrent ring buffer
func NewConcurrentRingBuffer(size int) *ConcurrentRingBuffer {
	// Ensure size is a power of 2
	actualSize := 1
	for actualSize < size {
		actualSize <<= 1
	}
	
	return &ConcurrentRingBuffer{
		buffer: make([]interface{}, actualSize),
		size:   actualSize,
		mask:   int64(actualSize - 1),
	}
}

// Put adds an item to the ring buffer
func (crb *ConcurrentRingBuffer) Put(value interface{}) bool {
	tail := atomic.LoadInt64(&crb.tail)
	head := atomic.LoadInt64(&crb.head)
	
	if tail-head >= int64(crb.size) {
		return false // Buffer is full
	}
	
	crb.buffer[tail&crb.mask] = value
	atomic.StoreInt64(&crb.tail, tail+1)
	return true
}

// Get gets an item from the ring buffer
func (crb *ConcurrentRingBuffer) Get() (interface{}, bool) {
	head := atomic.LoadInt64(&crb.head)
	tail := atomic.LoadInt64(&crb.tail)
	
	if head >= tail {
		return nil, false // Buffer is empty
	}
	
	value := crb.buffer[head&crb.mask]
	atomic.StoreInt64(&crb.head, head+1)
	return value, true
}

// Size returns the current size of the ring buffer
func (crb *ConcurrentRingBuffer) Size() int {
	tail := atomic.LoadInt64(&crb.tail)
	head := atomic.LoadInt64(&crb.head)
	return int(tail - head)
}

// ConcurrentCounter represents a thread-safe counter
type ConcurrentCounter struct {
	value int64
}

// NewConcurrentCounter creates a new concurrent counter
func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{}
}

// Increment increments the counter
func (cc *ConcurrentCounter) Increment() {
	atomic.AddInt64(&cc.value, 1)
}

// Decrement decrements the counter
func (cc *ConcurrentCounter) Decrement() {
	atomic.AddInt64(&cc.value, -1)
}

// Add adds a value to the counter
func (cc *ConcurrentCounter) Add(delta int64) {
	atomic.AddInt64(&cc.value, delta)
}

// Get returns the current value
func (cc *ConcurrentCounter) Get() int64 {
	return atomic.LoadInt64(&cc.value)
}

// Set sets the counter value
func (cc *ConcurrentCounter) Set(value int64) {
	atomic.StoreInt64(&cc.value, value)
}

// CompareAndSwap compares and swaps the value
func (cc *ConcurrentCounter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&cc.value, old, new)
}

// ConcurrentSet represents a thread-safe set
type ConcurrentSet struct {
	mu   sync.RWMutex
	data map[interface{}]bool
}

// NewConcurrentSet creates a new concurrent set
func NewConcurrentSet() *ConcurrentSet {
	return &ConcurrentSet{
		data: make(map[interface{}]bool),
	}
}

// Add adds an item to the set
func (cs *ConcurrentSet) Add(item interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.data[item] = true
}

// Remove removes an item from the set
func (cs *ConcurrentSet) Remove(item interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.data, item)
}

// Contains checks if an item is in the set
func (cs *ConcurrentSet) Contains(item interface{}) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.data[item]
}

// Size returns the size of the set
func (cs *ConcurrentSet) Size() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return len(cs.data)
}

// Clear clears the set
func (cs *ConcurrentSet) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.data = make(map[interface{}]bool)
}

// ToSlice returns a slice of all items in the set
func (cs *ConcurrentSet) ToSlice() []interface{} {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	
	slice := make([]interface{}, 0, len(cs.data))
	for item := range cs.data {
		slice = append(slice, item)
	}
	return slice
}

// ConcurrentList represents a thread-safe list
type ConcurrentList struct {
	mu   sync.RWMutex
	data []interface{}
}

// NewConcurrentList creates a new concurrent list
func NewConcurrentList() *ConcurrentList {
	return &ConcurrentList{
		data: make([]interface{}, 0),
	}
}

// Add adds an item to the list
func (cl *ConcurrentList) Add(item interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.data = append(cl.data, item)
}

// Get gets an item at the specified index
func (cl *ConcurrentList) Get(index int) (interface{}, bool) {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	
	if index < 0 || index >= len(cl.data) {
		return nil, false
	}
	return cl.data[index], true
}

// Remove removes an item at the specified index
func (cl *ConcurrentList) Remove(index int) bool {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	
	if index < 0 || index >= len(cl.data) {
		return false
	}
	
	cl.data = append(cl.data[:index], cl.data[index+1:]...)
	return true
}

// Size returns the size of the list
func (cl *ConcurrentList) Size() int {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	return len(cl.data)
}

// Clear clears the list
func (cl *ConcurrentList) Clear() {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.data = cl.data[:0]
}

// ToSlice returns a slice of all items in the list
func (cl *ConcurrentList) ToSlice() []interface{} {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	
	slice := make([]interface{}, len(cl.data))
	copy(slice, cl.data)
	return slice
}

// ConcurrentHashMap represents a thread-safe hash map with better performance
type ConcurrentHashMap struct {
	shards []*HashMapShard
	size   int
}

// HashMapShard represents a shard of the hash map
type HashMapShard struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// NewConcurrentHashMap creates a new concurrent hash map
func NewConcurrentHashMap() *ConcurrentHashMap {
	shards := make([]*HashMapShard, 32) // 32 shards for better concurrency
	for i := range shards {
		shards[i] = &HashMapShard{
			data: make(map[string]interface{}),
		}
	}
	return &ConcurrentHashMap{
		shards: shards,
		size:   32,
	}
}

// getShard returns the shard for a given key
func (chm *ConcurrentHashMap) getShard(key string) *HashMapShard {
	hash := fnv1a(key)
	return chm.shards[hash%uint32(len(chm.shards))]
}

// Set sets a value in the hash map
func (chm *ConcurrentHashMap) Set(key string, value interface{}) {
	shard := chm.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.data[key] = value
}

// Get gets a value from the hash map
func (chm *ConcurrentHashMap) Get(key string) (interface{}, bool) {
	shard := chm.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	value, exists := shard.data[key]
	return value, exists
}

// Delete deletes a value from the hash map
func (chm *ConcurrentHashMap) Delete(key string) {
	shard := chm.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	delete(shard.data, key)
}

// Size returns the size of the hash map
func (chm *ConcurrentHashMap) Size() int {
	total := 0
	for _, shard := range chm.shards {
		shard.mu.RLock()
		total += len(shard.data)
		shard.mu.RUnlock()
	}
	return total
}

// Clear clears the hash map
func (chm *ConcurrentHashMap) Clear() {
	for _, shard := range chm.shards {
		shard.mu.Lock()
		shard.data = make(map[string]interface{})
		shard.mu.Unlock()
	}
}
