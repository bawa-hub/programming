package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Exercise 1: Implement Atomic Counter
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Atomic Counter")
	fmt.Println("===================================")
	
	// TODO: Implement an atomic counter
	// 1. Create a counter using atomic operations
	// 2. Implement increment, decrement, and get methods
	// 3. Test with concurrent access
	
	counter := &AtomicCounter{}
	
	// Test increment
	counter.Increment()
	counter.Increment()
	fmt.Printf("  Exercise 1: After 2 increments: %d\n", counter.Get())
	
	// Test decrement
	counter.Decrement()
	fmt.Printf("  Exercise 1: After 1 decrement: %d\n", counter.Get())
	
	// Test concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("  Exercise 1: After concurrent increments: %d\n", counter.Get())
	
	fmt.Println("  Exercise 1: Atomic counter completed")
}

type AtomicCounter struct {
	value int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Decrement() {
	atomic.AddInt64(&ac.value, -1)
}

func (ac *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&ac.value)
}

// Exercise 2: Implement Lock-Free Stack
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Lock-Free Stack")
	fmt.Println("====================================")
	
	// TODO: Implement a lock-free stack
	// 1. Create a stack using atomic operations
	// 2. Implement push and pop methods
	// 3. Test with concurrent access
	
	stack := NewExerciseStack()
	
	// Test basic operations
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	
	fmt.Println("  Exercise 2: Pushed 3 items")
	
	// Test pop
	for i := 0; i < 3; i++ {
		value, ok := stack.Pop()
		if ok {
			fmt.Printf("  Exercise 2: Popped: %s\n", value)
		}
	}
	
	// Test concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				stack.Push(fmt.Sprintf("goroutine-%d-item-%d", id, j))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  Exercise 2: Pushed %d items concurrently\n", 5*10)
	
	fmt.Println("  Exercise 2: Lock-free stack completed")
}

type ExerciseNode struct {
	value interface{}
	next  unsafe.Pointer
}

type ExerciseStack struct {
	head unsafe.Pointer
}

func NewExerciseStack() *ExerciseStack {
	return &ExerciseStack{}
}

func (s *ExerciseStack) Push(value interface{}) {
	newNode := &ExerciseNode{value: value}
	
	for {
		current := atomic.LoadPointer(&s.head)
		newNode.next = unsafe.Pointer((*ExerciseNode)(current))
		
		if atomic.CompareAndSwapPointer(&s.head, current, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (s *ExerciseStack) Pop() (interface{}, bool) {
	for {
		current := atomic.LoadPointer(&s.head)
		if current == nil {
			return nil, false
		}
		
		node := (*ExerciseNode)(current)
		next := node.next
		
		if atomic.CompareAndSwapPointer(&s.head, current, unsafe.Pointer(next)) {
			return node.value, true
		}
	}
}

// Exercise 3: Implement Lock-Free Queue
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Lock-Free Queue")
	fmt.Println("====================================")
	
	// TODO: Implement a lock-free queue
	// 1. Create a queue using atomic operations
	// 2. Implement enqueue and dequeue methods
	// 3. Test with concurrent access
	
	queue := NewExerciseQueue()
	
	// Test basic operations
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	
	fmt.Println("  Exercise 3: Enqueued 3 items")
	
	// Test dequeue
	for i := 0; i < 3; i++ {
		value, ok := queue.Dequeue()
		if ok {
			fmt.Printf("  Exercise 3: Dequeued: %s\n", value)
		}
	}
	
	// Test concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				queue.Enqueue(fmt.Sprintf("goroutine-%d-item-%d", id, j))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  Exercise 3: Enqueued %d items concurrently\n", 5*10)
	
	fmt.Println("  Exercise 3: Lock-free queue completed")
}

type ExerciseQueueNode struct {
	value interface{}
	next  unsafe.Pointer
}

type ExerciseQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewExerciseQueue() *ExerciseQueue {
	dummy := &ExerciseQueueNode{}
	return &ExerciseQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *ExerciseQueue) Enqueue(value interface{}) {
	newNode := &ExerciseQueueNode{value: value}
	
	for {
		tail := atomic.LoadPointer(&q.tail)
		tailNode := (*ExerciseQueueNode)(tail)
		
		if atomic.CompareAndSwapPointer(&tailNode.next, nil, unsafe.Pointer(newNode)) {
			atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newNode))
			break
		}
	}
}

func (q *ExerciseQueue) Dequeue() (interface{}, bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		headNode := (*ExerciseQueueNode)(head)
		next := atomic.LoadPointer(&headNode.next)
		
		if next == nil {
			return nil, false
		}
		
		nextNode := (*ExerciseQueueNode)(next)
		value := nextNode.value
		
		if atomic.CompareAndSwapPointer(&q.head, head, next) {
			return value, true
		}
	}
}

// Exercise 4: Implement Lock-Free Ring Buffer
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Lock-Free Ring Buffer")
	fmt.Println("==========================================")
	
	// TODO: Implement a lock-free ring buffer
	// 1. Create a ring buffer using atomic operations
	// 2. Implement enqueue and dequeue methods
	// 3. Test with concurrent access
	
	rb := NewExerciseRingBuffer(5)
	
	// Test basic operations
	for i := 0; i < 7; i++ {
		success := rb.Enqueue(i)
		fmt.Printf("  Exercise 4: Enqueue %d: %t\n", i, success)
	}
	
	// Test dequeue
	for i := 0; i < 7; i++ {
		value, ok := rb.Dequeue()
		if ok {
			fmt.Printf("  Exercise 4: Dequeued: %v\n", value)
		} else {
			fmt.Println("  Exercise 4: Buffer empty")
		}
	}
	
	fmt.Println("  Exercise 4: Lock-free ring buffer completed")
}

type ExerciseRingBuffer struct {
	buffer []interface{}
	head   int64
	tail   int64
	size   int64
}

func NewExerciseRingBuffer(size int) *ExerciseRingBuffer {
	return &ExerciseRingBuffer{
		buffer: make([]interface{}, size),
		size:   int64(size),
	}
}

func (rb *ExerciseRingBuffer) Enqueue(value interface{}) bool {
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

func (rb *ExerciseRingBuffer) Dequeue() (interface{}, bool) {
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

// Exercise 5: Implement Lock-Free Hash Table
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Lock-Free Hash Table")
	fmt.Println("==========================================")
	
	// TODO: Implement a lock-free hash table
	// 1. Create a hash table using atomic operations
	// 2. Implement put and get methods
	// 3. Test with concurrent access
	
	ht := NewExerciseHashTable(10)
	
	// Test basic operations
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")
	
	value1, ok1 := ht.Get("key1")
	value2, ok2 := ht.Get("key2")
	value3, ok3 := ht.Get("key3")
	value4, ok4 := ht.Get("key4")
	
	fmt.Printf("  Exercise 5: key1: %v, %t\n", value1, ok1)
	fmt.Printf("  Exercise 5: key2: %v, %t\n", value2, ok2)
	fmt.Printf("  Exercise 5: key3: %v, %t\n", value3, ok3)
	fmt.Printf("  Exercise 5: key4: %v, %t\n", value4, ok4)
	
	fmt.Println("  Exercise 5: Lock-free hash table completed")
}

type ExerciseHashNode struct {
	key   string
	value interface{}
	next  unsafe.Pointer
}

type ExerciseHashTable struct {
	buckets []unsafe.Pointer
	size    int
}

func NewExerciseHashTable(size int) *ExerciseHashTable {
	return &ExerciseHashTable{
		buckets: make([]unsafe.Pointer, size),
		size:    size,
	}
}

func (ht *ExerciseHashTable) hash(key string) int {
	hash := 0
	for _, c := range key {
		hash = hash*31 + int(c)
	}
	return hash % ht.size
}

func (ht *ExerciseHashTable) Put(key string, value interface{}) {
	hash := ht.hash(key)
	newNode := &ExerciseHashNode{key: key, value: value}
	
	for {
		current := atomic.LoadPointer(&ht.buckets[hash])
		newNode.next = current
		
		if atomic.CompareAndSwapPointer(&ht.buckets[hash], current, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (ht *ExerciseHashTable) Get(key string) (interface{}, bool) {
	hash := ht.hash(key)
	current := atomic.LoadPointer(&ht.buckets[hash])
	
	for current != nil {
		node := (*ExerciseHashNode)(current)
		if node.key == key {
			return node.value, true
		}
		current = node.next
	}
	
	return nil, false
}

// Exercise 6: Implement Memory Pool
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Memory Pool")
	fmt.Println("=================================")
	
	// TODO: Implement a memory pool using atomic operations
	// 1. Create a memory pool for reusing memory
	// 2. Implement get and put methods
	// 3. Test with concurrent access
	
	pool := NewExerciseMemoryPool(1024)
	
	// Test basic operations
	ptr1 := pool.Get()
	ptr2 := pool.Get()
	
	fmt.Printf("  Exercise 6: Allocated: %v, %v\n", ptr1, ptr2)
	
	// Return memory
	pool.Put(ptr1)
	pool.Put(ptr2)
	
	// Allocate again (should reuse)
	ptr3 := pool.Get()
	fmt.Printf("  Exercise 6: Reused: %v\n", ptr3)
	
	fmt.Println("  Exercise 6: Memory pool completed")
}

type ExerciseMemoryPool struct {
	freeList unsafe.Pointer
	size     int
}

func NewExerciseMemoryPool(size int) *ExerciseMemoryPool {
	return &ExerciseMemoryPool{size: size}
}

func (mp *ExerciseMemoryPool) Get() unsafe.Pointer {
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

func (mp *ExerciseMemoryPool) Put(ptr unsafe.Pointer) {
	for {
		current := atomic.LoadPointer(&mp.freeList)
		*(*unsafe.Pointer)(ptr) = current
		
		if atomic.CompareAndSwapPointer(&mp.freeList, current, ptr) {
			break
		}
	}
}

// Exercise 7: Implement Reference Counting
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Reference Counting")
	fmt.Println("=======================================")
	
	// TODO: Implement reference counting using atomic operations
	// 1. Create a reference counted object
	// 2. Implement acquire and release methods
	// 3. Test with concurrent access
	
	rc := &ExerciseRefCounted{data: "test data"}
	
	// Test basic operations
	rc.Acquire()
	rc.Acquire()
	fmt.Printf("  Exercise 7: Reference count: %d\n", rc.GetCount())
	
	// Test release
	released := rc.Release()
	fmt.Printf("  Exercise 7: Released: %t, count: %d\n", released, rc.GetCount())
	
	released = rc.Release()
	fmt.Printf("  Exercise 7: Released: %t, count: %d\n", released, rc.GetCount())
	
	fmt.Println("  Exercise 7: Reference counting completed")
}

type ExerciseRefCounted struct {
	count int64
	data  interface{}
}

func (rc *ExerciseRefCounted) Acquire() {
	atomic.AddInt64(&rc.count, 1)
}

func (rc *ExerciseRefCounted) Release() bool {
	count := atomic.AddInt64(&rc.count, -1)
	return count == 0
}

func (rc *ExerciseRefCounted) GetCount() int64 {
	return atomic.LoadInt64(&rc.count)
}

// Exercise 8: Implement Work Stealing Queue
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Work Stealing Queue")
	fmt.Println("=======================================")
	
	// TODO: Implement a work stealing queue using atomic operations
	// 1. Create a work stealing queue
	// 2. Implement push, pop, and steal methods
	// 3. Test with concurrent access
	
	wsq := NewExerciseWorkStealingQueue(10)
	
	// Test basic operations
	for i := 0; i < 5; i++ {
		success := wsq.Push(fmt.Sprintf("work-%d", i))
		fmt.Printf("  Exercise 8: Pushed work-%d: %t\n", i, success)
	}
	
	// Test pop
	for i := 0; i < 5; i++ {
		work, ok := wsq.Pop()
		if ok {
			fmt.Printf("  Exercise 8: Popped: %s\n", work)
		} else {
			fmt.Println("  Exercise 8: No work available")
		}
	}
	
	fmt.Println("  Exercise 8: Work stealing queue completed")
}

type ExerciseWorkStealingQueue struct {
	tasks    []interface{}
	head     int64
	tail     int64
	capacity int64
}

func NewExerciseWorkStealingQueue(capacity int) *ExerciseWorkStealingQueue {
	return &ExerciseWorkStealingQueue{
		tasks:    make([]interface{}, capacity),
		capacity: int64(capacity),
	}
}

func (wsq *ExerciseWorkStealingQueue) Push(task interface{}) bool {
	currentTail := atomic.LoadInt64(&wsq.tail)
	nextTail := (currentTail + 1) % wsq.capacity
	
	if nextTail == atomic.LoadInt64(&wsq.head) {
		return false // Queue full
	}
	
	wsq.tasks[currentTail] = task
	atomic.StoreInt64(&wsq.tail, nextTail)
	return true
}

func (wsq *ExerciseWorkStealingQueue) Pop() (interface{}, bool) {
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

func (wsq *ExerciseWorkStealingQueue) Steal() (interface{}, bool) {
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

// Exercise 9: Implement Performance Comparison
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Performance Comparison")
	fmt.Println("===========================================")
	
	// TODO: Compare lock-based vs lock-free performance
	// 1. Create both lock-based and lock-free implementations
	// 2. Benchmark both with concurrent access
	// 3. Compare performance results
	
	// Lock-based counter
	lockCounter := &ExerciseLockCounter{}
	
	// Lock-free counter
	lockFreeCounter := &ExerciseLockFreeCounter{}
	
	// Test lock-based counter
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
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
			for j := 0; j < 1000; j++ {
				lockFreeCounter.Increment()
			}
		}()
	}
	wg.Wait()
	lockFreeTime := time.Since(start)
	
	fmt.Printf("  Exercise 9: Lock-based time: %v\n", lockTime)
	fmt.Printf("  Exercise 9: Lock-free time: %v\n", lockFreeTime)
	fmt.Printf("  Exercise 9: Lock-free speedup: %.2fx\n", float64(lockTime)/float64(lockFreeTime))
	
	fmt.Println("  Exercise 9: Performance comparison completed")
}

type ExerciseLockCounter struct {
	mu    sync.Mutex
	count int64
}

func (lc *ExerciseLockCounter) Increment() {
	lc.mu.Lock()
	lc.count++
	lc.mu.Unlock()
}

func (lc *ExerciseLockCounter) Get() int64 {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.count
}

type ExerciseLockFreeCounter struct {
	count int64
}

func (lfc *ExerciseLockFreeCounter) Increment() {
	atomic.AddInt64(&lfc.count, 1)
}

func (lfc *ExerciseLockFreeCounter) Get() int64 {
	return atomic.LoadInt64(&lfc.count)
}

// Exercise 10: Implement Stress Testing
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Stress Testing")
	fmt.Println("====================================")
	
	// TODO: Implement stress testing for lock-free data structures
	// 1. Create a lock-free data structure
	// 2. Test with many concurrent goroutines
	// 3. Verify correctness under stress
	
	stack := NewExerciseStack()
	
	// Stress test with many goroutines
	var wg sync.WaitGroup
	numGoroutines := 50
	numOperations := 100
	
	start := time.Now()
	
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
	
	// Pop operations
	popped := 0
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				_, ok := stack.Pop()
				if ok {
					popped++
				}
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("  Exercise 10: Pushed %d operations\n", numGoroutines*numOperations)
	fmt.Printf("  Exercise 10: Popped %d operations\n", popped)
	fmt.Printf("  Exercise 10: Duration: %v\n", duration)
	fmt.Printf("  Exercise 10: Operations per second: %.0f\n", float64(numGoroutines*numOperations*2)/duration.Seconds())
	
	fmt.Println("  Exercise 10: Stress testing completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Lock-Free Programming Exercises")
	fmt.Println("==================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
