package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Example 1: Basic Atomic Operations
func basicAtomicOperations() {
	fmt.Println("\n1. Basic Atomic Operations")
	fmt.Println("=========================")
	
	var counter int64
	
	// Atomic add
	atomic.AddInt64(&counter, 1)
	fmt.Printf("  After add: %d\n", counter)
	
	// Atomic load
	value := atomic.LoadInt64(&counter)
	fmt.Printf("  Loaded value: %d\n", value)
	
	// Atomic store
	atomic.StoreInt64(&counter, 100)
	fmt.Printf("  After store: %d\n", counter)
	
	// Atomic swap
	old := atomic.SwapInt64(&counter, 200)
	fmt.Printf("  Swapped %d for %d\n", old, counter)
	
	// Atomic compare and swap
	swapped := atomic.CompareAndSwapInt64(&counter, 200, 300)
	fmt.Printf("  CAS result: %t, value: %d\n", swapped, counter)
	
	fmt.Println("  Basic atomic operations completed")
}

// Example 2: Atomic Value Operations
func atomicValueOperations() {
	fmt.Println("\n2. Atomic Value Operations")
	fmt.Println("=========================")
	
	var value atomic.Value
	
	// Store string
	value.Store("hello")
	loaded := value.Load().(string)
	fmt.Printf("  Stored and loaded: %s\n", loaded)
	
	// Store integer (need to create new atomic.Value for different type)
	var intValue atomic.Value
	intValue.Store(42)
	loadedInt := intValue.Load().(int)
	fmt.Printf("  Stored and loaded: %d\n", loadedInt)
	
	// Store struct
	type Data struct {
		ID   int
		Name string
	}
	
	var dataValue atomic.Value
	data := Data{ID: 1, Name: "test"}
	dataValue.Store(data)
	loadedData := dataValue.Load().(Data)
	fmt.Printf("  Stored and loaded: %+v\n", loadedData)
	
	fmt.Println("  Atomic value operations completed")
}

// Example 3: Atomic Boolean Operations
func atomicBooleanOperations() {
	fmt.Println("\n3. Atomic Boolean Operations")
	fmt.Println("============================")
	
	var flag int32
	
	// Set flag
	atomic.StoreInt32(&flag, 1)
	fmt.Printf("  Flag set: %d\n", atomic.LoadInt32(&flag))
	
	// Check flag
	if atomic.LoadInt32(&flag) == 1 {
		fmt.Println("  Flag is set")
	}
	
	// Toggle flag
	atomic.StoreInt32(&flag, 1-atomic.LoadInt32(&flag))
	fmt.Printf("  Flag toggled: %d\n", atomic.LoadInt32(&flag))
	
	// Toggle again
	atomic.StoreInt32(&flag, 1-atomic.LoadInt32(&flag))
	fmt.Printf("  Flag toggled again: %d\n", atomic.LoadInt32(&flag))
	
	fmt.Println("  Atomic boolean operations completed")
}

// Example 4: Compare and Swap Operations
func compareAndSwapOperations() {
	fmt.Println("\n4. Compare and Swap Operations")
	fmt.Println("=============================")
	
	var value int64 = 10
	
	// Try to change 10 to 20
	success := atomic.CompareAndSwapInt64(&value, 10, 20)
	fmt.Printf("  CAS 10->20: %t, value: %d\n", success, value)
	
	// Try to change 20 to 30 (will succeed)
	success = atomic.CompareAndSwapInt64(&value, 20, 30)
	fmt.Printf("  CAS 20->30: %t, value: %d\n", success, value)
	
	// Try to change 20 to 40 (will fail)
	success = atomic.CompareAndSwapInt64(&value, 20, 40)
	fmt.Printf("  CAS 20->40: %t, value: %d\n", success, value)
	
	fmt.Println("  Compare and swap operations completed")
}

// Example 5: CAS Loops
func casLoops() {
	fmt.Println("\n5. CAS Loops")
	fmt.Println("============")
	
	var counter int64
	
	// Atomic increment using CAS loop
	atomicIncrement(&counter)
	fmt.Printf("  After increment: %d\n", counter)
	
	// Atomic decrement using CAS loop
	atomicDecrement(&counter)
	fmt.Printf("  After decrement: %d\n", counter)
	
	// Multiple increments
	for i := 0; i < 5; i++ {
		atomicIncrement(&counter)
	}
	fmt.Printf("  After 5 increments: %d\n", counter)
	
	fmt.Println("  CAS loops completed")
}

func atomicIncrement(addr *int64) {
	for {
		current := atomic.LoadInt64(addr)
		newValue := current + 1
		if atomic.CompareAndSwapInt64(addr, current, newValue) {
			break
		}
	}
}

func atomicDecrement(addr *int64) {
	for {
		current := atomic.LoadInt64(addr)
		newValue := current - 1
		if atomic.CompareAndSwapInt64(addr, current, newValue) {
			break
		}
	}
}

// Example 6: Memory Ordering
func memoryOrdering() {
	fmt.Println("\n6. Memory Ordering")
	fmt.Println("==================")
	
	var data int64
	var ready int32
	
	// Writer thread
	go func() {
		data = 42
		atomic.StoreInt32(&ready, 1) // Release store
	}()
	
	// Reader thread
	go func() {
		for atomic.LoadInt32(&ready) == 0 {
			// Wait for data to be ready
		}
		// Acquire load ensures data is visible
		value := data
		fmt.Printf("  Read value: %d\n", value)
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Memory ordering completed")
}

// Example 7: Lock-Free Counter
func lockFreeCounter() {
	fmt.Println("\n7. Lock-Free Counter")
	fmt.Println("===================")
	
	counter := &LockFreeCounter{}
	
	// Concurrent increments
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("  Final count: %d\n", counter.Get())
	
	fmt.Println("  Lock-free counter completed")
}

type LockFreeCounter struct {
	count int64
}

func (lfc *LockFreeCounter) Increment() {
	atomic.AddInt64(&lfc.count, 1)
}

func (lfc *LockFreeCounter) Get() int64 {
	return atomic.LoadInt64(&lfc.count)
}

// Example 8: Lock-Free Stack
func lockFreeStack() {
	fmt.Println("\n8. Lock-Free Stack")
	fmt.Println("=================")
	
	stack := NewLockFreeStack()
	
	// Push some values
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	
	// Pop values
	for {
		value, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Printf("  Popped: %s\n", value)
	}
	
	fmt.Println("  Lock-free stack completed")
}

type Node struct {
	value interface{}
	next  *Node
}

type LockFreeStack struct {
	head unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value interface{}) {
	newNode := &Node{value: value}
	
	for {
		current := atomic.LoadPointer(&s.head)
		newNode.next = (*Node)(current)
		
		if atomic.CompareAndSwapPointer(&s.head, current, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() (interface{}, bool) {
	for {
		current := atomic.LoadPointer(&s.head)
		if current == nil {
			return nil, false
		}
		
		node := (*Node)(current)
		next := node.next
		
		if atomic.CompareAndSwapPointer(&s.head, current, unsafe.Pointer(next)) {
			return node.value, true
		}
	}
}

// Example 9: Lock-Free Queue
func lockFreeQueue() {
	fmt.Println("\n9. Lock-Free Queue")
	fmt.Println("=================")
	
	queue := NewLockFreeQueue()
	
	// Enqueue some values
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	
	// Dequeue values
	for {
		value, ok := queue.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("  Dequeued: %s\n", value)
	}
	
	fmt.Println("  Lock-free queue completed")
}

type QueueNode struct {
	value interface{}
	next  unsafe.Pointer
}

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	dummy := &QueueNode{}
	return &LockFreeQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *LockFreeQueue) Enqueue(value interface{}) {
	newNode := &QueueNode{value: value}
	
	for {
		tail := atomic.LoadPointer(&q.tail)
		tailNode := (*QueueNode)(tail)
		
		if atomic.CompareAndSwapPointer(&tailNode.next, nil, unsafe.Pointer(newNode)) {
			atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newNode))
			break
		}
	}
}

func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		headNode := (*QueueNode)(head)
		next := atomic.LoadPointer(&headNode.next)
		
		if next == nil {
			return nil, false
		}
		
		nextNode := (*QueueNode)(next)
		value := nextNode.value
		
		if atomic.CompareAndSwapPointer(&q.head, head, next) {
			return value, true
		}
	}
}

// Example 10: Lock-Free Ring Buffer
func lockFreeRingBuffer() {
	fmt.Println("\n10. Lock-Free Ring Buffer")
	fmt.Println("========================")
	
	rb := NewLockFreeRingBuffer(5)
	
	// Enqueue values
	for i := 0; i < 7; i++ {
		success := rb.Enqueue(i)
		fmt.Printf("  Enqueue %d: %t\n", i, success)
	}
	
	// Dequeue values
	for i := 0; i < 7; i++ {
		value, ok := rb.Dequeue()
		if ok {
			fmt.Printf("  Dequeued: %v\n", value)
		} else {
			fmt.Println("  Queue empty")
		}
	}
	
	fmt.Println("  Lock-free ring buffer completed")
}

type LockFreeRingBuffer struct {
	buffer []interface{}
	head   int64
	tail   int64
	size   int64
}

func NewLockFreeRingBuffer(size int) *LockFreeRingBuffer {
	return &LockFreeRingBuffer{
		buffer: make([]interface{}, size),
		size:   int64(size),
	}
}

func (rb *LockFreeRingBuffer) Enqueue(value interface{}) bool {
	for {
		currentTail := atomic.LoadInt64(&rb.tail)
		nextTail := (currentTail + 1) % rb.size
		
		if nextTail == atomic.LoadInt64(&rb.head) {
			return false // Buffer full
		}
		
		if atomic.CompareAndSwapInt64(&rb.tail, currentTail, nextTail) {
			rb.buffer[currentTail] = value
			return true
		}
	}
}

func (rb *LockFreeRingBuffer) Dequeue() (interface{}, bool) {
	for {
		currentHead := atomic.LoadInt64(&rb.head)
		
		if currentHead == atomic.LoadInt64(&rb.tail) {
			return nil, false // Buffer empty
		}
		
		value := rb.buffer[currentHead]
		nextHead := (currentHead + 1) % rb.size
		
		if atomic.CompareAndSwapInt64(&rb.head, currentHead, nextHead) {
			return value, true
		}
	}
}

// Example 11: Performance Comparison
func performanceComparison() {
	fmt.Println("\n11. Performance Comparison")
	fmt.Println("==========================")
	
	// Lock-based counter
	lockCounter := &LockCounter{}
	
	// Lock-free counter
	lockFreeCounter := &LockFreeCounter{}
	
	// Test lock-based counter
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				lockCounter.Increment()
			}
		}()
	}
	wg.Wait()
	lockTime := time.Since(start)
	
	// Test lock-free counter
	start = time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				lockFreeCounter.Increment()
			}
		}()
	}
	wg.Wait()
	lockFreeTime := time.Since(start)
	
	fmt.Printf("  Lock-based time: %v\n", lockTime)
	fmt.Printf("  Lock-free time: %v\n", lockFreeTime)
	fmt.Printf("  Lock-free speedup: %.2fx\n", float64(lockTime)/float64(lockFreeTime))
	
	fmt.Println("  Performance comparison completed")
}

type LockCounter struct {
	mu    sync.Mutex
	count int64
}

func (lc *LockCounter) Increment() {
	lc.mu.Lock()
	lc.count++
	lc.mu.Unlock()
}

func (lc *LockCounter) Get() int64 {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.count
}

// Example 12: Atomic Pointer Operations
func atomicPointerOperations() {
	fmt.Println("\n12. Atomic Pointer Operations")
	fmt.Println("=============================")
	
	var ptr unsafe.Pointer
	
	// Create some data
	data1 := &struct{ value int }{value: 1}
	data2 := &struct{ value int }{value: 2}
	
	// Atomic load pointer
	loaded := atomic.LoadPointer(&ptr)
	fmt.Printf("  Initial pointer: %v\n", loaded)
	
	// Atomic store pointer
	atomic.StorePointer(&ptr, unsafe.Pointer(data1))
	loaded = atomic.LoadPointer(&ptr)
	fmt.Printf("  After store data1: %v\n", loaded)
	
	// Atomic swap pointer
	old := atomic.SwapPointer(&ptr, unsafe.Pointer(data2))
	fmt.Printf("  Swapped %v for %v\n", old, atomic.LoadPointer(&ptr))
	
	// Atomic compare and swap pointer
	swapped := atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(data2), unsafe.Pointer(data1))
	fmt.Printf("  CAS result: %t\n", swapped)
	
	fmt.Println("  Atomic pointer operations completed")
}

// Example 13: Memory Pool
func memoryPool() {
	fmt.Println("\n13. Memory Pool")
	fmt.Println("===============")
	
	pool := NewMemoryPool(1024)
	
	// Allocate memory
	ptr1 := pool.Get()
	ptr2 := pool.Get()
	
	fmt.Printf("  Allocated pointers: %v, %v\n", ptr1, ptr2)
	
	// Return memory
	pool.Put(ptr1)
	pool.Put(ptr2)
	
	// Allocate again (should reuse)
	ptr3 := pool.Get()
	fmt.Printf("  Reused pointer: %v\n", ptr3)
	
	fmt.Println("  Memory pool completed")
}

type MemoryPool struct {
	freeList unsafe.Pointer
	size     int
}

func NewMemoryPool(size int) *MemoryPool {
	return &MemoryPool{size: size}
}

func (mp *MemoryPool) Get() unsafe.Pointer {
	for {
		current := atomic.LoadPointer(&mp.freeList)
		if current == nil {
			return unsafe.Pointer(&make([]byte, mp.size)[0])
		}
		
		next := *(*unsafe.Pointer)(current)
		if atomic.CompareAndSwapPointer(&mp.freeList, current, next) {
			return current
		}
	}
}

func (mp *MemoryPool) Put(ptr unsafe.Pointer) {
	for {
		current := atomic.LoadPointer(&mp.freeList)
		*(*unsafe.Pointer)(ptr) = current
		
		if atomic.CompareAndSwapPointer(&mp.freeList, current, ptr) {
			break
		}
	}
}

// Example 14: Lock-Free Hash Table
func basicLockFreeHashTable() {
	fmt.Println("\n14. Lock-Free Hash Table")
	fmt.Println("=======================")
	
	ht := NewBasicLockFreeHashTable(10)
	
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
	
	fmt.Println("  Lock-free hash table completed")
}

type BasicHashNode struct {
	key   string
	value interface{}
	next  unsafe.Pointer
}

type BasicLockFreeHashTable struct {
	buckets []unsafe.Pointer
	size    int
}

func NewBasicLockFreeHashTable(size int) *BasicLockFreeHashTable {
	return &BasicLockFreeHashTable{
		buckets: make([]unsafe.Pointer, size),
		size:    size,
	}
}

func (ht *BasicLockFreeHashTable) hash(key string) int {
	hash := 0
	for _, c := range key {
		hash = hash*31 + int(c)
	}
	return hash % ht.size
}

func (ht *BasicLockFreeHashTable) Put(key string, value interface{}) {
	hash := ht.hash(key)
	newNode := &BasicHashNode{key: key, value: value}
	
	for {
		current := atomic.LoadPointer(&ht.buckets[hash])
		newNode.next = current
		
		if atomic.CompareAndSwapPointer(&ht.buckets[hash], current, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (ht *BasicLockFreeHashTable) Get(key string) (interface{}, bool) {
	hash := ht.hash(key)
	current := atomic.LoadPointer(&ht.buckets[hash])
	
	for current != nil {
		node := (*BasicHashNode)(current)
		if node.key == key {
			return node.value, true
		}
		current = node.next
	}
	
	return nil, false
}

// Example 15: Work Stealing Queue
func workStealingQueue() {
	fmt.Println("\n15. Work Stealing Queue")
	fmt.Println("=======================")
	
	wsq := NewWorkStealingQueue(10)
	
	// Push some work
	for i := 0; i < 5; i++ {
		success := wsq.Push(fmt.Sprintf("work-%d", i))
		fmt.Printf("  Pushed work-%d: %t\n", i, success)
	}
	
	// Pop work
	for i := 0; i < 5; i++ {
		work, ok := wsq.Pop()
		if ok {
			fmt.Printf("  Popped: %s\n", work)
		} else {
			fmt.Println("  No work available")
		}
	}
	
	fmt.Println("  Work stealing queue completed")
}

type WorkStealingQueue struct {
	tasks    []interface{}
	head     int64
	tail     int64
	capacity int64
}

func NewWorkStealingQueue(capacity int) *WorkStealingQueue {
	return &WorkStealingQueue{
		tasks:    make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (wsq *WorkStealingQueue) Push(task interface{}) bool {
	currentTail := atomic.LoadInt64(&wsq.tail)
	nextTail := (currentTail + 1) % wsq.capacity
	
	if nextTail == atomic.LoadInt64(&wsq.head) {
		return false // Queue full
	}
	
	wsq.tasks[currentTail] = task
	atomic.StoreInt64(&wsq.tail, nextTail)
	return true
}

func (wsq *WorkStealingQueue) Pop() (interface{}, bool) {
	currentTail := atomic.LoadInt64(&wsq.tail)
	currentHead := atomic.LoadInt64(&wsq.head)
	
	if currentHead == currentTail {
		return nil, false // Queue empty
	}
	
	// Try to pop from tail
	newTail := (currentTail - 1 + wsq.capacity) % wsq.capacity
	if atomic.CompareAndSwapInt64(&wsq.tail, currentTail, newTail) {
		task := wsq.tasks[newTail]
		return task, true
	}
	
	return nil, false
}

func (wsq *WorkStealingQueue) Steal() (interface{}, bool) {
	currentHead := atomic.LoadInt64(&wsq.head)
	currentTail := atomic.LoadInt64(&wsq.tail)
	
	if currentHead == currentTail {
		return nil, false // Queue empty
	}
	
	// Try to steal from head
	task := wsq.tasks[currentHead]
	newHead := (currentHead + 1) % wsq.capacity
	
	if atomic.CompareAndSwapInt64(&wsq.head, currentHead, newHead) {
		return task, true
	}
	
	return nil, false
}

// Example 16: Reference Counting
func referenceCounting() {
	fmt.Println("\n16. Reference Counting")
	fmt.Println("=====================")
	
	rc := &RefCounted{data: "test data"}
	
	// Acquire references
	rc.Acquire()
	rc.Acquire()
	fmt.Printf("  Reference count: %d\n", rc.GetCount())
	
	// Release references
	released := rc.Release()
	fmt.Printf("  Released: %t, count: %d\n", released, rc.GetCount())
	
	released = rc.Release()
	fmt.Printf("  Released: %t, count: %d\n", released, rc.GetCount())
	
	fmt.Println("  Reference counting completed")
}

type RefCounted struct {
	count int64
	data  interface{}
}

func (rc *RefCounted) Acquire() {
	atomic.AddInt64(&rc.count, 1)
}

func (rc *RefCounted) Release() bool {
	count := atomic.AddInt64(&rc.count, -1)
	return count == 0
}

func (rc *RefCounted) GetCount() int64 {
	return atomic.LoadInt64(&rc.count)
}

// Example 17: ABA Problem Demonstration
func abaProblem() {
	fmt.Println("\n17. ABA Problem Demonstration")
	fmt.Println("=============================")
	
	var ptr unsafe.Pointer
	var counter int64
	
	// Thread 1: Reads and processes
	go func() {
		for i := 0; i < 5; i++ {
			current := atomic.LoadPointer(&ptr)
			if current != nil {
				fmt.Printf("  Thread 1: Processing %v\n", current)
				time.Sleep(10 * time.Millisecond)
				// Try to update (might fail due to ABA)
				if atomic.CompareAndSwapPointer(&ptr, current, current) {
					fmt.Printf("  Thread 1: Successfully updated\n")
				} else {
					fmt.Printf("  Thread 1: Update failed (ABA problem)\n")
				}
			}
		}
	}()
	
	// Thread 2: Changes value back and forth
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Millisecond)
			atomic.StorePointer(&ptr, unsafe.Pointer(&counter))
			atomic.AddInt64(&counter, 1)
			time.Sleep(5 * time.Millisecond)
			atomic.StorePointer(&ptr, unsafe.Pointer(&counter))
			atomic.AddInt64(&counter, 1)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  ABA problem demonstration completed")
}

// Example 18: Lock-Free Allocator
func basicLockFreeAllocator() {
	fmt.Println("\n18. Lock-Free Allocator")
	fmt.Println("======================")
	
	allocator := NewBasicLockFreeAllocator(1024)
	
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
	
	fmt.Println("  Lock-free allocator completed")
}

type BasicLockFreeAllocator struct {
	freeList unsafe.Pointer
	size     int
}

func NewBasicLockFreeAllocator(size int) *BasicLockFreeAllocator {
	return &BasicLockFreeAllocator{size: size}
}

func (lfa *BasicLockFreeAllocator) Allocate() unsafe.Pointer {
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

func (lfa *BasicLockFreeAllocator) Deallocate(ptr unsafe.Pointer) {
	for {
		current := atomic.LoadPointer(&lfa.freeList)
		*(*unsafe.Pointer)(ptr) = current
		
		if atomic.CompareAndSwapPointer(&lfa.freeList, current, ptr) {
			break
		}
	}
}

// Example 19: Concurrent Testing
func concurrentTesting() {
	fmt.Println("\n19. Concurrent Testing")
	fmt.Println("=====================")
	
	stack := NewLockFreeStack()
	
	// Concurrent push/pop
	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100
	
	// Push operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				stack.Push(fmt.Sprintf("goroutine-%d-operation-%d", id, j))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  Pushed %d operations\n", numGoroutines*numOperations)
	
	// Pop operations
	var popped int64
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				_, ok := stack.Pop()
				if ok {
					atomic.AddInt64(&popped, 1)
				}
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("  Popped %d operations\n", atomic.LoadInt64(&popped))
	
	fmt.Println("  Concurrent testing completed")
}

// Example 20: Stress Testing
func stressTesting() {
	fmt.Println("\n20. Stress Testing")
	fmt.Println("=================")
	
	counter := &LockFreeCounter{}
	
	// Stress test with many goroutines
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 1000
	
	start := time.Now()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("  Final count: %d\n", counter.Get())
	fmt.Printf("  Duration: %v\n", duration)
	fmt.Printf("  Operations per second: %.0f\n", float64(numGoroutines*numOperations)/duration.Seconds())
	
	fmt.Println("  Stress testing completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("⚛️ Lock-Free Programming Examples")
	fmt.Println("=================================")
	
	basicAtomicOperations()
	atomicValueOperations()
	atomicBooleanOperations()
	compareAndSwapOperations()
	casLoops()
	memoryOrdering()
	lockFreeCounter()
	lockFreeStack()
	lockFreeQueue()
	lockFreeRingBuffer()
	performanceComparison()
	atomicPointerOperations()
	memoryPool()
	basicLockFreeHashTable()
	workStealingQueue()
	referenceCounting()
	abaProblem()
	basicLockFreeAllocator()
	concurrentTesting()
	stressTesting()
}
