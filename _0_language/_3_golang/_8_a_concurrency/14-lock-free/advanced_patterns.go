package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
	"unsafe"
)

// Advanced Pattern 1: Lock-Free Skip List
type SkipListNode struct {
	key     int
	value   interface{}
	next    []unsafe.Pointer
	marked  int32
	deleted int32
}

type LockFreeSkipList struct {
	head   *SkipListNode
	levels int
}

func NewLockFreeSkipList() *LockFreeSkipList {
	head := &SkipListNode{
		key:  -1,
		next: make([]unsafe.Pointer, 32),
	}
	
	return &LockFreeSkipList{
		head:   head,
		levels: 32,
	}
}

func (sl *LockFreeSkipList) randomLevel() int {
	level := 1
	for rand.Float32() < 0.5 && level < sl.levels {
		level++
	}
	return level
}

func (sl *LockFreeSkipList) find(key int) (*SkipListNode, []*SkipListNode) {
	preds := make([]*SkipListNode, sl.levels)
	current := sl.head
	
	for level := sl.levels - 1; level >= 0; level-- {
		for {
			next := atomic.LoadPointer(&current.next[level])
			if next == nil {
				break
			}
			
			nextNode := (*SkipListNode)(next)
			if nextNode.key < key {
				current = nextNode
			} else {
				break
			}
		}
		preds[level] = current
	}
	
	return current, preds
}

func (sl *LockFreeSkipList) Insert(key int, value interface{}) bool {
	level := sl.randomLevel()
	newNode := &SkipListNode{
		key:   key,
		value: value,
		next:  make([]unsafe.Pointer, level),
	}
	
	for {
		pred, preds := sl.find(key)
		
		if pred.key == key {
			return false // Key already exists
		}
		
		// Insert at each level
		for i := 0; i < level; i++ {
			newNode.next[i] = preds[i].next[i]
		}
		
		// Update predecessors
		for i := 0; i < level; i++ {
			if !atomic.CompareAndSwapPointer(&preds[i].next[i], newNode.next[i], unsafe.Pointer(newNode)) {
				// Retry
				break
			}
		}
		
		return true
	}
}

func (sl *LockFreeSkipList) Delete(key int) bool {
	for {
		pred, preds := sl.find(key)
		
		if pred.key != key {
			return false // Key not found
		}
		
		// Mark as deleted
		atomic.StoreInt32(&pred.deleted, 1)
		
		// Update predecessors
		for i := 0; i < len(pred.next); i++ {
			atomic.CompareAndSwapPointer(&preds[i].next[i], unsafe.Pointer(pred), pred.next[i])
		}
		
		return true
	}
}

func (sl *LockFreeSkipList) Get(key int) (interface{}, bool) {
	pred, _ := sl.find(key)
	
	if pred.key == key && atomic.LoadInt32(&pred.deleted) == 0 {
		return pred.value, true
	}
	
	return nil, false
}

// Advanced Pattern 2: Lock-Free Hash Table with Chaining
type LockFreeHashTable struct {
	buckets []unsafe.Pointer
	size    int
}

type HashNode struct {
	key   string
	value interface{}
	next  unsafe.Pointer
}

func NewLockFreeHashTable(size int) *LockFreeHashTable {
	return &LockFreeHashTable{
		buckets: make([]unsafe.Pointer, size),
		size:    size,
	}
}

func (ht *LockFreeHashTable) hash(key string) int {
	hash := 0
	for _, c := range key {
		hash = hash*31 + int(c)
	}
	return hash % ht.size
}

func (ht *LockFreeHashTable) Put(key string, value interface{}) {
	hash := ht.hash(key)
	newNode := &HashNode{key: key, value: value}
	
	for {
		current := atomic.LoadPointer(&ht.buckets[hash])
		newNode.next = current
		
		if atomic.CompareAndSwapPointer(&ht.buckets[hash], current, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (ht *LockFreeHashTable) Get(key string) (interface{}, bool) {
	hash := ht.hash(key)
	current := atomic.LoadPointer(&ht.buckets[hash])
	
	for current != nil {
		node := (*HashNode)(current)
		if node.key == key {
			return node.value, true
		}
		current = node.next
	}
	
	return nil, false
}

// Advanced Pattern 3: Lock-Free Work Stealing Deque
type WorkStealingDeque struct {
	tasks    []interface{}
	head     int64
	tail     int64
	capacity int64
}

func NewWorkStealingDeque(capacity int) *WorkStealingDeque {
	return &WorkStealingDeque{
		tasks:    make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (wsd *WorkStealingDeque) PushBottom(task interface{}) bool {
	currentTail := atomic.LoadInt64(&wsd.tail)
	nextTail := (currentTail + 1) % wsd.capacity
	
	if nextTail == atomic.LoadInt64(&wsd.head) {
		return false // Deque full
	}
	
	wsd.tasks[currentTail] = task
	atomic.StoreInt64(&wsd.tail, nextTail)
	return true
}

func (wsd *WorkStealingDeque) PopBottom() (interface{}, bool) {
	currentTail := atomic.LoadInt64(&wsd.tail)
	currentHead := atomic.LoadInt64(&wsd.head)
	
	if currentHead == currentTail {
		return nil, false // Deque empty
	}
	
	// Try to pop from tail
	newTail := (currentTail - 1 + wsd.capacity) % wsd.capacity
	if atomic.CompareAndSwapInt64(&wsd.tail, currentTail, newTail) {
		task := wsd.tasks[newTail]
		return task, true
	}
	
	return nil, false
}

func (wsd *WorkStealingDeque) Steal() (interface{}, bool) {
	currentHead := atomic.LoadInt64(&wsd.head)
	currentTail := atomic.LoadInt64(&wsd.tail)
	
	if currentHead == currentTail {
		return nil, false // Deque empty
	}
	
	// Try to steal from head
	task := wsd.tasks[currentHead]
	newHead := (currentHead + 1) % wsd.capacity
	
	if atomic.CompareAndSwapInt64(&wsd.head, currentHead, newHead) {
		return task, true
	}
	
	return nil, false
}

// Advanced Pattern 4: Lock-Free Memory Allocator
type LockFreeAllocator struct {
	freeList unsafe.Pointer
	size     int
}

func NewLockFreeAllocator(size int) *LockFreeAllocator {
	return &LockFreeAllocator{size: size}
}

func (lfa *LockFreeAllocator) Allocate() unsafe.Pointer {
	for {
		current := atomic.LoadPointer(&lfa.freeList)
		if current == nil {
			return unsafe.Pointer(&make([]byte, lfa.size)[0])
		}
		
		next := *(*unsafe.Pointer)(current)
		if atomic.CompareAndSwapPointer(&lfa.freeList, current, next) {
			return current
		}
	}
}

func (lfa *LockFreeAllocator) Deallocate(ptr unsafe.Pointer) {
	for {
		current := atomic.LoadPointer(&lfa.freeList)
		*(*unsafe.Pointer)(ptr) = current
		
		if atomic.CompareAndSwapPointer(&lfa.freeList, current, ptr) {
			break
		}
	}
}

// Advanced Pattern 5: Lock-Free Reference Counting
type LockFreeRefCounted struct {
	count int64
	data  interface{}
}

func NewLockFreeRefCounted(data interface{}) *LockFreeRefCounted {
	return &LockFreeRefCounted{
		count: 1,
		data:  data,
	}
}

func (rc *LockFreeRefCounted) Acquire() {
	atomic.AddInt64(&rc.count, 1)
}

func (rc *LockFreeRefCounted) Release() bool {
	count := atomic.AddInt64(&rc.count, -1)
	return count == 0
}

func (rc *LockFreeRefCounted) GetCount() int64 {
	return atomic.LoadInt64(&rc.count)
}

// Advanced Pattern 6: Lock-Free Cache
type LockFreeCache struct {
	buckets []unsafe.Pointer
	size    int
}

type CacheEntry struct {
	key      string
	value    interface{}
	expiry   time.Time
	next     unsafe.Pointer
	refCount int64
}

func NewLockFreeCache(size int) *LockFreeCache {
	return &LockFreeCache{
		buckets: make([]unsafe.Pointer, size),
		size:    size,
	}
}

func (c *LockFreeCache) hash(key string) int {
	hash := 0
	for _, c := range key {
		hash = hash*31 + int(c)
	}
	return hash % c.size
}

func (c *LockFreeCache) Get(key string) (interface{}, bool) {
	hash := c.hash(key)
	current := atomic.LoadPointer(&c.buckets[hash])
	
	for current != nil {
		entry := (*CacheEntry)(current)
		if entry.key == key {
			if time.Now().Before(entry.expiry) {
				atomic.AddInt64(&entry.refCount, 1)
				return entry.value, true
			}
			// Entry expired, remove it
			c.remove(key)
			return nil, false
		}
		current = entry.next
	}
	
	return nil, false
}

func (c *LockFreeCache) Set(key string, value interface{}, ttl time.Duration) {
	hash := c.hash(key)
	newEntry := &CacheEntry{
		key:      key,
		value:    value,
		expiry:   time.Now().Add(ttl),
		refCount: 1,
	}
	
	for {
		current := atomic.LoadPointer(&c.buckets[hash])
		newEntry.next = current
		
		if atomic.CompareAndSwapPointer(&c.buckets[hash], current, unsafe.Pointer(newEntry)) {
			break
		}
	}
}

func (c *LockFreeCache) remove(key string) {
	// Implementation for removing expired entries
}

// Advanced Pattern 7: Lock-Free Priority Queue
type LockFreePriorityQueue struct {
	head unsafe.Pointer
}

type PriorityNode struct {
	value    interface{}
	priority int
	next     unsafe.Pointer
}

func NewLockFreePriorityQueue() *LockFreePriorityQueue {
	return &LockFreePriorityQueue{}
}

func (pq *LockFreePriorityQueue) Insert(value interface{}, priority int) {
	newNode := &PriorityNode{
		value:    value,
		priority: priority,
	}
	
	for {
		current := atomic.LoadPointer(&pq.head)
		if current == nil {
			// Empty queue
			if atomic.CompareAndSwapPointer(&pq.head, nil, unsafe.Pointer(newNode)) {
				break
			}
		} else {
			node := (*PriorityNode)(current)
			if priority > node.priority {
				// Insert at head
				newNode.next = current
				if atomic.CompareAndSwapPointer(&pq.head, current, unsafe.Pointer(newNode)) {
					break
				}
			} else {
				// Find insertion point
				for {
					next := atomic.LoadPointer(&node.next)
					if next == nil {
						// Insert at end
						if atomic.CompareAndSwapPointer(&node.next, nil, unsafe.Pointer(newNode)) {
							break
						}
					} else {
						nextNode := (*PriorityNode)(next)
						if priority > nextNode.priority {
							// Insert here
							newNode.next = next
							if atomic.CompareAndSwapPointer(&node.next, next, unsafe.Pointer(newNode)) {
								break
							}
						} else {
							node = nextNode
						}
					}
				}
				break
			}
		}
	}
}

func (pq *LockFreePriorityQueue) ExtractMax() (interface{}, bool) {
	for {
		current := atomic.LoadPointer(&pq.head)
		if current == nil {
			return nil, false // Queue empty
		}
		
		node := (*PriorityNode)(current)
		next := node.next
		
		if atomic.CompareAndSwapPointer(&pq.head, current, next) {
			return node.value, true
		}
	}
}

// Advanced Pattern 8: Lock-Free Bounded Queue
type LockFreeBoundedQueue struct {
	buffer   []interface{}
	head     int64
	tail     int64
	capacity int64
}

func NewLockFreeBoundedQueue(capacity int) *LockFreeBoundedQueue {
	return &LockFreeBoundedQueue{
		buffer:   make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (q *LockFreeBoundedQueue) Enqueue(value interface{}) bool {
	for {
		currentTail := atomic.LoadInt64(&q.tail)
		nextTail := (currentTail + 1) % q.capacity
		
		if nextTail == atomic.LoadInt64(&q.head) {
			return false // Queue full
		}
		
		if atomic.CompareAndSwapInt64(&q.tail, currentTail, nextTail) {
			q.buffer[currentTail] = value
			return true
		}
	}
}

func (q *LockFreeBoundedQueue) Dequeue() (interface{}, bool) {
	for {
		currentHead := atomic.LoadInt64(&q.head)
		
		if currentHead == atomic.LoadInt64(&q.tail) {
			return nil, false // Queue empty
		}
		
		value := q.buffer[currentHead]
		nextHead := (currentHead + 1) % q.capacity
		
		if atomic.CompareAndSwapInt64(&q.head, currentHead, nextHead) {
			return value, true
		}
	}
}

// Advanced Pattern 9: Lock-Free Set
type LockFreeSet struct {
	buckets []unsafe.Pointer
	size    int
}

type SetNode struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLockFreeSet(size int) *LockFreeSet {
	return &LockFreeSet{
		buckets: make([]unsafe.Pointer, size),
		size:    size,
	}
}

func (s *LockFreeSet) hash(value interface{}) int {
	hash := 0
	switch v := value.(type) {
	case string:
		for _, c := range v {
			hash = hash*31 + int(c)
		}
	case int:
		hash = v
	}
	return hash % s.size
}

func (s *LockFreeSet) Add(value interface{}) bool {
	hash := s.hash(value)
	newNode := &SetNode{value: value}
	
	for {
		current := atomic.LoadPointer(&s.buckets[hash])
		
		// Check if value already exists
		temp := current
		for temp != nil {
			node := (*SetNode)(temp)
			if node.value == value {
				return false // Value already exists
			}
			temp = node.next
		}
		
		newNode.next = current
		if atomic.CompareAndSwapPointer(&s.buckets[hash], current, unsafe.Pointer(newNode)) {
			return true
		}
	}
}

func (s *LockFreeSet) Contains(value interface{}) bool {
	hash := s.hash(value)
	current := atomic.LoadPointer(&s.buckets[hash])
	
	for current != nil {
		node := (*SetNode)(current)
		if node.value == value {
			return true
		}
		current = node.next
	}
	
	return false
}

// Advanced Pattern 10: Lock-Free Thread Pool
type LockFreeThreadPool struct {
	workers []*Worker
	jobs    chan Job
	quit    chan bool
}

type Worker struct {
	id    int
	jobs  chan Job
	quit  chan bool
	pool  *LockFreeThreadPool
}

type Job struct {
	ID   int
	Data interface{}
	Fn   func(interface{}) interface{}
}

func NewLockFreeThreadPool(numWorkers int) *LockFreeThreadPool {
	pool := &LockFreeThreadPool{
		workers: make([]*Worker, numWorkers),
		jobs:    make(chan Job, 1000),
		quit:    make(chan bool),
	}
	
	for i := 0; i < numWorkers; i++ {
		worker := &Worker{
			id:   i,
			jobs: pool.jobs,
			quit: pool.quit,
			pool: pool,
		}
		pool.workers[i] = worker
		go worker.start()
	}
	
	return pool
}

func (w *Worker) start() {
	for {
		select {
		case job := <-w.jobs:
			// Process job
			result := job.Fn(job.Data)
			fmt.Printf("  Worker %d processed job %d: %v\n", w.id, job.ID, result)
		case <-w.quit:
			return
		}
	}
}

func (tp *LockFreeThreadPool) Submit(job Job) {
	tp.jobs <- job
}

func (tp *LockFreeThreadPool) Shutdown() {
	close(tp.quit)
}

// Advanced Pattern 1: Lock-Free Skip List
func lockFreeSkipList() {
	fmt.Println("\n1. Lock-Free Skip List")
	fmt.Println("=====================")
	
	sl := NewLockFreeSkipList()
	
	// Insert some values
	sl.Insert(1, "one")
	sl.Insert(2, "two")
	sl.Insert(3, "three")
	
	// Get values
	value1, ok1 := sl.Get(1)
	value2, ok2 := sl.Get(2)
	value3, ok3 := sl.Get(3)
	value4, ok4 := sl.Get(4)
	
	fmt.Printf("  Key 1: %v, %t\n", value1, ok1)
	fmt.Printf("  Key 2: %v, %t\n", value2, ok2)
	fmt.Printf("  Key 3: %v, %t\n", value3, ok3)
	fmt.Printf("  Key 4: %v, %t\n", value4, ok4)
	
	// Delete a value
	sl.Delete(2)
	value2, ok2 = sl.Get(2)
	fmt.Printf("  Key 2 after delete: %v, %t\n", value2, ok2)
	
	fmt.Println("Lock-free skip list completed")
}

// Advanced Pattern 2: Lock-Free Hash Table
func lockFreeHashTable() {
	fmt.Println("\n2. Lock-Free Hash Table")
	fmt.Println("======================")
	
	ht := NewLockFreeHashTable(10)
	
	// Put some values
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")
	
	// Get values
	value1, ok1 := ht.Get("key1")
	value2, ok2 := ht.Get("key2")
	value3, ok3 := ht.Get("key3")
	value4, ok4 := ht.Get("key4")
	
	fmt.Printf("  key1: %v, %t\n", value1, ok1)
	fmt.Printf("  key2: %v, %t\n", value2, ok2)
	fmt.Printf("  key3: %v, %t\n", value3, ok3)
	fmt.Printf("  key4: %v, %t\n", value4, ok4)
	
	fmt.Println("Lock-free hash table completed")
}

// Advanced Pattern 3: Work Stealing Deque
func workStealingDeque() {
	fmt.Println("\n3. Work Stealing Deque")
	fmt.Println("=====================")
	
	wsd := NewWorkStealingDeque(10)
	
	// Push some work
	for i := 0; i < 5; i++ {
		success := wsd.PushBottom(fmt.Sprintf("work-%d", i))
		fmt.Printf("  Pushed work-%d: %t\n", i, success)
	}
	
	// Pop work
	for i := 0; i < 5; i++ {
		work, ok := wsd.PopBottom()
		if ok {
			fmt.Printf("  Popped: %s\n", work)
		} else {
			fmt.Println("  No work available")
		}
	}
	
	fmt.Println("Work stealing deque completed")
}

// Advanced Pattern 4: Lock-Free Memory Allocator
func lockFreeMemoryAllocator() {
	fmt.Println("\n4. Lock-Free Memory Allocator")
	fmt.Println("============================")
	
	allocator := NewLockFreeAllocator(1024)
	
	// Allocate memory
	ptr1 := allocator.Allocate()
	ptr2 := allocator.Allocate()
	
	fmt.Printf("  Allocated: %v, %v\n", ptr1, ptr2)
	
	// Deallocate memory
	allocator.Deallocate(ptr1)
	allocator.Deallocate(ptr2)
	
	// Allocate again (should reuse)
	ptr3 := allocator.Allocate()
	fmt.Printf("  Reused: %v\n", ptr3)
	
	fmt.Println("Lock-free memory allocator completed")
}

// Advanced Pattern 5: Lock-Free Reference Counting
func lockFreeReferenceCounting() {
	fmt.Println("\n5. Lock-Free Reference Counting")
	fmt.Println("===============================")
	
	rc := NewLockFreeRefCounted("test data")
	
	// Acquire references
	rc.Acquire()
	rc.Acquire()
	fmt.Printf("  Reference count: %d\n", rc.GetCount())
	
	// Release references
	released := rc.Release()
	fmt.Printf("  Released: %t, count: %d\n", released, rc.GetCount())
	
	released = rc.Release()
	fmt.Printf("  Released: %t, count: %d\n", released, rc.GetCount())
	
	fmt.Println("Lock-free reference counting completed")
}

// Advanced Pattern 6: Lock-Free Cache
func lockFreeCache() {
	fmt.Println("\n6. Lock-Free Cache")
	fmt.Println("=================")
	
	cache := NewLockFreeCache(10)
	
	// Set some values
	cache.Set("key1", "value1", 1*time.Second)
	cache.Set("key2", "value2", 1*time.Second)
	cache.Set("key3", "value3", 1*time.Second)
	
	// Get values
	value1, ok1 := cache.Get("key1")
	value2, ok2 := cache.Get("key2")
	value3, ok3 := cache.Get("key3")
	value4, ok4 := cache.Get("key4")
	
	fmt.Printf("  key1: %v, %t\n", value1, ok1)
	fmt.Printf("  key2: %v, %t\n", value2, ok2)
	fmt.Printf("  key3: %v, %t\n", value3, ok3)
	fmt.Printf("  key4: %v, %t\n", value4, ok4)
	
	fmt.Println("Lock-free cache completed")
}

// Advanced Pattern 7: Lock-Free Priority Queue
func lockFreePriorityQueue() {
	fmt.Println("\n7. Lock-Free Priority Queue")
	fmt.Println("===========================")
	
	pq := NewLockFreePriorityQueue()
	
	// Insert some values with priorities
	pq.Insert("low", 1)
	pq.Insert("high", 10)
	pq.Insert("medium", 5)
	pq.Insert("highest", 15)
	
	// Extract values (should be in priority order)
	for i := 0; i < 4; i++ {
		value, ok := pq.ExtractMax()
		if ok {
			fmt.Printf("  Extracted: %s\n", value)
		} else {
			fmt.Println("  Queue empty")
		}
	}
	
	fmt.Println("Lock-free priority queue completed")
}

// Advanced Pattern 8: Lock-Free Bounded Queue
func lockFreeBoundedQueue() {
	fmt.Println("\n8. Lock-Free Bounded Queue")
	fmt.Println("==========================")
	
	q := NewLockFreeBoundedQueue(5)
	
	// Enqueue some values
	for i := 0; i < 7; i++ {
		success := q.Enqueue(i)
		fmt.Printf("  Enqueued %d: %t\n", i, success)
	}
	
	// Dequeue values
	for i := 0; i < 7; i++ {
		value, ok := q.Dequeue()
		if ok {
			fmt.Printf("  Dequeued: %v\n", value)
		} else {
			fmt.Println("  Queue empty")
		}
	}
	
	fmt.Println("Lock-free bounded queue completed")
}

// Advanced Pattern 9: Lock-Free Set
func lockFreeSet() {
	fmt.Println("\n9. Lock-Free Set")
	fmt.Println("===============")
	
	s := NewLockFreeSet(10)
	
	// Add some values
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")
	s.Add("apple") // Duplicate
	
	// Check if values exist
	fmt.Printf("  Contains apple: %t\n", s.Contains("apple"))
	fmt.Printf("  Contains banana: %t\n", s.Contains("banana"))
	fmt.Printf("  Contains cherry: %t\n", s.Contains("cherry"))
	fmt.Printf("  Contains orange: %t\n", s.Contains("orange"))
	
	fmt.Println("Lock-free set completed")
}

// Advanced Pattern 10: Lock-Free Thread Pool
func lockFreeThreadPool() {
	fmt.Println("\n10. Lock-Free Thread Pool")
	fmt.Println("=========================")
	
	pool := NewLockFreeThreadPool(3)
	
	// Submit some jobs
	for i := 0; i < 5; i++ {
		job := Job{
			ID:   i,
			Data: i * 2,
			Fn: func(data interface{}) interface{} {
				return data.(int) * 2
			},
		}
		pool.Submit(job)
	}
	
	// Wait for jobs to complete
	time.Sleep(100 * time.Millisecond)
	
	// Shutdown pool
	pool.Shutdown()
	
	fmt.Println("Lock-free thread pool completed")
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Lock-Free Programming Patterns")
	fmt.Println("==========================================")
	
	lockFreeSkipList()
	lockFreeHashTable()
	workStealingDeque()
	lockFreeMemoryAllocator()
	lockFreeReferenceCounting()
	lockFreeCache()
	lockFreePriorityQueue()
	lockFreeBoundedQueue()
	lockFreeSet()
	lockFreeThreadPool()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
