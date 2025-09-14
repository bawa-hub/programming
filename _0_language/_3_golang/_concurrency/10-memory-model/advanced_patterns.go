package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// LockFreeStack represents a lock-free stack using atomic operations
type LockFreeStack struct {
	head unsafe.Pointer
}

type node struct {
	value int
	next  unsafe.Pointer
}

// NewLockFreeStack creates a new lock-free stack
func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

// Push pushes a value onto the stack
func (s *LockFreeStack) Push(value int) {
	n := &node{value: value}
	
	for {
		head := atomic.LoadPointer(&s.head)
		n.next = head
		
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(n)) {
			break
		}
	}
}

// Pop pops a value from the stack
func (s *LockFreeStack) Pop() (int, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}
		
		n := (*node)(head)
		next := atomic.LoadPointer(&n.next)
		
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return n.value, true
		}
	}
}

// MemoryPool represents a memory pool to avoid allocations
type MemoryPool struct {
	pool sync.Pool
}

// NewMemoryPool creates a new memory pool
func NewMemoryPool() *MemoryPool {
	return &MemoryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 1024)
			},
		},
	}
}

// Get gets a buffer from the pool
func (p *MemoryPool) Get() []byte {
	return p.pool.Get().([]byte)
}

// Put puts a buffer back into the pool
func (p *MemoryPool) Put(buf []byte) {
	p.pool.Put(buf)
}

// DoubleCheckedLocking represents a singleton with double-checked locking
type DoubleCheckedLocking struct {
	data string
}

var (
	instance *DoubleCheckedLocking
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *DoubleCheckedLocking {
	if instance == nil {
		once.Do(func() {
			instance = &DoubleCheckedLocking{data: "initialized"}
		})
	}
	return instance
}

// LockFreeCounter represents a lock-free counter
type LockFreeCounter struct {
	value int64
}

// Increment increments the counter
func (c *LockFreeCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Decrement decrements the counter
func (c *LockFreeCounter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

// Value returns the current value
func (c *LockFreeCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// LockFreeRingBuffer represents a lock-free ring buffer
type LockFreeRingBuffer struct {
	buffer []int
	size   int
	head   int64
	tail   int64
}

// NewLockFreeRingBuffer creates a new lock-free ring buffer
func NewLockFreeRingBuffer(size int) *LockFreeRingBuffer {
	return &LockFreeRingBuffer{
		buffer: make([]int, size),
		size:   size,
	}
}

// Push pushes a value into the buffer
func (rb *LockFreeRingBuffer) Push(value int) bool {
	tail := atomic.LoadInt64(&rb.tail)
	head := atomic.LoadInt64(&rb.head)
	
	nextTail := (tail + 1) % int64(rb.size)
	if nextTail == head {
		return false // Buffer is full
	}
	
	rb.buffer[tail] = value
	atomic.StoreInt64(&rb.tail, nextTail)
	return true
}

// Pop pops a value from the buffer
func (rb *LockFreeRingBuffer) Pop() (int, bool) {
	head := atomic.LoadInt64(&rb.head)
	tail := atomic.LoadInt64(&rb.tail)
	
	if head == tail {
		return 0, false // Buffer is empty
	}
	
	value := rb.buffer[head]
	atomic.StoreInt64(&rb.head, (head+1)%int64(rb.size))
	return value, true
}

// Advanced Pattern 1: Lock-Free Stack
func lockFreeStackExample() {
	fmt.Println("\n1. Lock-Free Stack")
	fmt.Println("==================")
	
	stack := NewLockFreeStack()
	
	// Push values
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	
	// Pop values
	for i := 0; i < 10; i++ {
		if value, ok := stack.Pop(); ok {
			fmt.Printf("  Popped: %d\n", value)
		}
	}
	
	fmt.Println("Lock-free stack example completed")
}

// Advanced Pattern 2: Memory Pool
func memoryPoolExample() {
	fmt.Println("\n2. Memory Pool")
	fmt.Println("==============")
	
	pool := NewMemoryPool()
	
	// Get buffers from pool
	buffers := make([][]byte, 10)
	for i := 0; i < 10; i++ {
		buffers[i] = pool.Get()
		fmt.Printf("  Got buffer %d: len=%d\n", i, len(buffers[i]))
	}
	
	// Put buffers back to pool
	for i := 0; i < 10; i++ {
		pool.Put(buffers[i])
		fmt.Printf("  Put buffer %d back\n", i)
	}
	
	fmt.Println("Memory pool example completed")
}

// Advanced Pattern 3: Double-Checked Locking
func doubleCheckedLockingExample() {
	fmt.Println("\n3. Double-Checked Locking")
	fmt.Println("========================")
	
	// Get singleton instances
	instances := make([]*DoubleCheckedLocking, 10)
	for i := 0; i < 10; i++ {
		instances[i] = GetInstance()
	}
	
	// Verify all instances are the same
	allSame := true
	for i := 1; i < 10; i++ {
		if instances[i] != instances[0] {
			allSame = false
			break
		}
	}
	
	fmt.Printf("  All instances same: %v\n", allSame)
	fmt.Printf("  Data: %s\n", instances[0].data)
	fmt.Println("Double-checked locking example completed")
}

// Advanced Pattern 4: Lock-Free Counter
func lockFreeCounterExample() {
	fmt.Println("\n4. Lock-Free Counter")
	fmt.Println("===================")
	
	counter := &LockFreeCounter{}
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Counter value: %d (expected: 10000)\n", counter.Value())
	fmt.Println("Lock-free counter example completed")
}

// Advanced Pattern 5: Lock-Free Ring Buffer
func lockFreeRingBufferExample() {
	fmt.Println("\n5. Lock-Free Ring Buffer")
	fmt.Println("=======================")
	
	rb := NewLockFreeRingBuffer(5)
	
	// Push values
	for i := 0; i < 7; i++ {
		success := rb.Push(i)
		fmt.Printf("  Push %d: %v\n", i, success)
	}
	
	// Pop values
	for i := 0; i < 7; i++ {
		if value, ok := rb.Pop(); ok {
			fmt.Printf("  Pop: %d\n", value)
		} else {
			fmt.Printf("  Pop: failed (buffer empty)\n")
		}
	}
	
	fmt.Println("Lock-free ring buffer example completed")
}

// Advanced Pattern 6: False Sharing Prevention
func falseSharingPreventionExample() {
	fmt.Println("\n6. False Sharing Prevention")
	fmt.Println("===========================")
	
	// Bad: False sharing
	type BadCounter struct {
		counter1 int64
		counter2 int64 // Same cache line as counter1
	}
	
	badCounter := &BadCounter{}
	
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(&badCounter.counter1, 1)
		atomic.AddInt64(&badCounter.counter2, 1)
	}
	badDuration := time.Since(start)
	
	// Good: Avoid false sharing with padding
	type GoodCounter struct {
		counter1 int64
		_        [7]int64 // Padding to avoid false sharing
		counter2 int64
	}
	
	goodCounter := &GoodCounter{}
	
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(&goodCounter.counter1, 1)
		atomic.AddInt64(&goodCounter.counter2, 1)
	}
	goodDuration := time.Since(start)
	
	fmt.Printf("  Bad (false sharing): %v\n", badDuration)
	fmt.Printf("  Good (no false sharing): %v\n", goodDuration)
	fmt.Printf("  Improvement: %.2fx faster\n", float64(badDuration)/float64(goodDuration))
	fmt.Println("False sharing prevention example completed")
}

// Advanced Pattern 7: Memory Barrier with Atomic Operations
func memoryBarrierExample() {
	fmt.Println("\n7. Memory Barrier with Atomic Operations")
	fmt.Println("=======================================")
	
	var x, y int64
	
	// Goroutine 1: writes x, then y
	go func() {
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 1) // This acts as a memory barrier
		fmt.Println("  Set x=1, y=1")
	}()
	
	// Goroutine 2: reads y, then x
	go func() {
		for atomic.LoadInt64(&y) != 1 {
			// Wait for y to be set
		}
		fmt.Printf("  x=%d, y=%d (x guaranteed to be 1)\n", 
			atomic.LoadInt64(&x), atomic.LoadInt64(&y))
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Memory barrier example completed")
}

// Advanced Pattern 8: Lock-Free Hash Table
func lockFreeHashTableExample() {
	fmt.Println("\n8. Lock-Free Hash Table")
	fmt.Println("======================")
	
	// Simple lock-free hash table using atomic operations
	type LockFreeHashTable struct {
		buckets []*LockFreeCounter
		size    int
	}
	
	ht := &LockFreeHashTable{
		buckets: make([]*LockFreeCounter, 16),
		size:    16,
	}
	
	// Initialize buckets
	for i := 0; i < ht.size; i++ {
		ht.buckets[i] = &LockFreeCounter{}
	}
	
	// Hash function
	hash := func(key string) int {
		h := 0
		for _, c := range key {
			h = h*31 + int(c)
		}
		return h % ht.size
	}
	
	// Increment counters
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	for i := 0; i < 1000; i++ {
		key := keys[i%len(keys)]
		bucket := hash(key)
		ht.buckets[bucket].Increment()
	}
	
	// Print results
	for i, bucket := range ht.buckets {
		if bucket.Value() > 0 {
			fmt.Printf("  Bucket %d: %d\n", i, bucket.Value())
		}
	}
	
	fmt.Println("Lock-free hash table example completed")
}

// Advanced Pattern 9: Lock-Free Queue
func lockFreeQueueExample() {
	fmt.Println("\n9. Lock-Free Queue")
	fmt.Println("=================")
	
	// Simple lock-free queue using atomic operations
	type LockFreeQueue struct {
		head unsafe.Pointer
		tail unsafe.Pointer
	}
	
	type queueNode struct {
		value int
		next  unsafe.Pointer
	}
	
	queue := &LockFreeQueue{}
	
	// Enqueue values
	for i := 0; i < 10; i++ {
		node := &queueNode{value: i}
		
		for {
			tail := atomic.LoadPointer(&queue.tail)
			if tail == nil {
				// First node
				if atomic.CompareAndSwapPointer(&queue.head, nil, unsafe.Pointer(node)) {
					atomic.StorePointer(&queue.tail, unsafe.Pointer(node))
					break
				}
			} else {
				// Add to tail
				next := atomic.LoadPointer(&(*queueNode)(tail).next)
				if next == nil {
					if atomic.CompareAndSwapPointer(&(*queueNode)(tail).next, nil, unsafe.Pointer(node)) {
						atomic.StorePointer(&queue.tail, unsafe.Pointer(node))
						break
					}
				} else {
					atomic.StorePointer(&queue.tail, next)
				}
			}
		}
	}
	
	// Dequeue values
	for i := 0; i < 10; i++ {
		for {
			head := atomic.LoadPointer(&queue.head)
			if head == nil {
				break
			}
			
			node := (*queueNode)(head)
			next := atomic.LoadPointer(&node.next)
			
			if atomic.CompareAndSwapPointer(&queue.head, head, next) {
				fmt.Printf("  Dequeued: %d\n", node.value)
				break
			}
		}
	}
	
	fmt.Println("Lock-free queue example completed")
}

// Advanced Pattern 10: Performance Comparison
func performanceComparisonExample() {
	fmt.Println("\n10. Performance Comparison")
	fmt.Println("=========================")
	
	iterations := 1000000
	
	// Mutex performance
	start := time.Now()
	var mu sync.Mutex
	var counter1 int
	for i := 0; i < iterations; i++ {
		mu.Lock()
		counter1++
		mu.Unlock()
	}
	mutexDuration := time.Since(start)
	
	// Atomic performance
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicDuration := time.Since(start)
	
	// Lock-free counter performance
	start = time.Now()
	lfCounter := &LockFreeCounter{}
	for i := 0; i < iterations; i++ {
		lfCounter.Increment()
	}
	lockFreeDuration := time.Since(start)
	
	fmt.Printf("  Mutex: %v\n", mutexDuration)
	fmt.Printf("  Atomic: %v\n", atomicDuration)
	fmt.Printf("  Lock-free: %v\n", lockFreeDuration)
	fmt.Printf("  Atomic is %.2fx faster than mutex\n", float64(mutexDuration)/float64(atomicDuration))
	fmt.Printf("  Lock-free is %.2fx faster than mutex\n", float64(mutexDuration)/float64(lockFreeDuration))
	fmt.Println("Performance comparison example completed")
}

// Advanced Pattern 11: Memory Model Validation
func memoryModelValidationExample() {
	fmt.Println("\n11. Memory Model Validation")
	fmt.Println("===========================")
	
	// Test sequential consistency
	var x, y int64
	
	// Goroutine 1: writes x, then y
	go func() {
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 1)
	}()
	
	// Goroutine 2: reads y, then x
	go func() {
		for atomic.LoadInt64(&y) != 1 {
			// Wait for y to be set
		}
		xVal := atomic.LoadInt64(&x)
		if xVal != 1 {
			fmt.Printf("  ERROR: x=%d, expected 1\n", xVal)
		} else {
			fmt.Println("  Sequential consistency validated")
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Memory model validation example completed")
}

// Advanced Pattern 12: Lock-Free Producer-Consumer
func lockFreeProducerConsumerExample() {
	fmt.Println("\n12. Lock-Free Producer-Consumer")
	fmt.Println("==============================")
	
	// Use ring buffer for lock-free producer-consumer
	rb := NewLockFreeRingBuffer(10)
	
	// Producer
	go func() {
		for i := 0; i < 20; i++ {
			for !rb.Push(i) {
				// Buffer full, wait
				time.Sleep(1 * time.Millisecond)
			}
			fmt.Printf("  Produced: %d\n", i)
		}
	}()
	
	// Consumer
	go func() {
		for i := 0; i < 20; i++ {
			for {
				if value, ok := rb.Pop(); ok {
					fmt.Printf("  Consumed: %d\n", value)
					break
				}
				// Buffer empty, wait
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	fmt.Println("Lock-free producer-consumer example completed")
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Memory Model Patterns")
	fmt.Println("=================================")
	
	lockFreeStackExample()
	memoryPoolExample()
	doubleCheckedLockingExample()
	lockFreeCounterExample()
	lockFreeRingBufferExample()
	falseSharingPreventionExample()
	memoryBarrierExample()
	lockFreeHashTableExample()
	lockFreeQueueExample()
	performanceComparisonExample()
	memoryModelValidationExample()
	lockFreeProducerConsumerExample()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
