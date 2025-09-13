package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// ⚛️ ATOMIC OPERATIONS MASTERY
// Understanding atomic operations and memory ordering in Go

func main() {
	fmt.Println("⚛️ ATOMIC OPERATIONS MASTERY")
	fmt.Println("============================")

	// 1. Basic Atomic Operations
	fmt.Println("\n1. Basic Atomic Operations:")
	basicAtomicOperations()

	// 2. Atomic Counters
	fmt.Println("\n2. Atomic Counters:")
	atomicCounters()

	// 3. Atomic Boolean Operations
	fmt.Println("\n3. Atomic Boolean Operations:")
	atomicBooleanOperations()

	// 4. Atomic Pointer Operations
	fmt.Println("\n4. Atomic Pointer Operations:")
	atomicPointerOperations()

	// 5. Atomic Value Operations
	fmt.Println("\n5. Atomic Value Operations:")
	atomicValueOperations()

	// 6. Memory Ordering
	fmt.Println("\n6. Memory Ordering:")
	memoryOrdering()

	// 7. Lock-Free Data Structures
	fmt.Println("\n7. Lock-Free Data Structures:")
	lockFreeDataStructures()

	// 8. Performance Comparison
	fmt.Println("\n8. Performance Comparison:")
	performanceComparison()
}

// BASIC ATOMIC OPERATIONS: Understanding basic atomic operations
func basicAtomicOperations() {
	fmt.Println("Understanding basic atomic operations...")
	
	// Atomic integer operations
	var counter int64 = 0
	
	// Atomic add
	atomic.AddInt64(&counter, 10)
	fmt.Printf("  📊 After AddInt64(10): %d\n", counter)
	
	// Atomic increment
	atomic.AddInt64(&counter, 1)
	fmt.Printf("  📊 After increment: %d\n", counter)
	
	// Atomic decrement
	atomic.AddInt64(&counter, -1)
	fmt.Printf("  📊 After decrement: %d\n", counter)
	
	// Atomic load
	value := atomic.LoadInt64(&counter)
	fmt.Printf("  📊 Loaded value: %d\n", value)
	
	// Atomic store
	atomic.StoreInt64(&counter, 100)
	fmt.Printf("  📊 After StoreInt64(100): %d\n", counter)
	
	// Atomic compare and swap
	swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
	fmt.Printf("  📊 CompareAndSwap(100, 200): %t, value: %d\n", swapped, counter)
	
	// Atomic swap
	oldValue := atomic.SwapInt64(&counter, 300)
	fmt.Printf("  📊 Swap(300): old=%d, new=%d\n", oldValue, counter)
}

// ATOMIC COUNTERS: Understanding atomic counters
func atomicCounters() {
	fmt.Println("Understanding atomic counters...")
	
	// Create atomic counter
	counter := &AtomicCounter{}
	
	// Test concurrent access
	var wg sync.WaitGroup
	goroutines := 100
	operations := 1000
	
	start := time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("  📊 Atomic counter: %v (value: %d)\n", duration, counter.Value())
	fmt.Printf("  📊 Expected value: %d\n", goroutines*operations)
	
	// Test atomic operations on different types
	testAtomicTypes()
}

func testAtomicTypes() {
	// Test different atomic types
	var int32Val int32
	var int64Val int64
	var uint32Val uint32
	var uint64Val uint64
	
	// Int32 operations
	atomic.AddInt32(&int32Val, 10)
	atomic.StoreInt32(&int32Val, 100)
	fmt.Printf("  📊 Int32: %d\n", atomic.LoadInt32(&int32Val))
	
	// Int64 operations
	atomic.AddInt64(&int64Val, 20)
	atomic.StoreInt64(&int64Val, 200)
	fmt.Printf("  📊 Int64: %d\n", atomic.LoadInt64(&int64Val))
	
	// Uint32 operations
	atomic.AddUint32(&uint32Val, 30)
	atomic.StoreUint32(&uint32Val, 300)
	fmt.Printf("  📊 Uint32: %d\n", atomic.LoadUint32(&uint32Val))
	
	// Uint64 operations
	atomic.AddUint64(&uint64Val, 40)
	atomic.StoreUint64(&uint64Val, 400)
	fmt.Printf("  📊 Uint64: %d\n", atomic.LoadUint64(&uint64Val))
}

// ATOMIC BOOLEAN OPERATIONS: Understanding atomic boolean operations
func atomicBooleanOperations() {
	fmt.Println("Understanding atomic boolean operations...")
	
	// Atomic boolean operations
	var flag int32 = 0
	
	// Set flag to true
	atomic.StoreInt32(&flag, 1)
	fmt.Printf("  📊 Flag set to true: %t\n", atomic.LoadInt32(&flag) == 1)
	
	// Toggle flag
	oldValue := atomic.SwapInt32(&flag, 0)
	fmt.Printf("  📊 Flag toggled: old=%t, new=%t\n", oldValue == 1, atomic.LoadInt32(&flag) == 1)
	
	// Compare and swap
	swapped := atomic.CompareAndSwapInt32(&flag, 0, 1)
	fmt.Printf("  📊 CompareAndSwap(0, 1): %t, flag=%t\n", swapped, atomic.LoadInt32(&flag) == 1)
	
	// Test concurrent flag operations
	testConcurrentFlags()
}

func testConcurrentFlags() {
	var flag int32 = 0
	var wg sync.WaitGroup
	goroutines := 10
	
	// Multiple goroutines trying to set flag
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			swapped := atomic.CompareAndSwapInt32(&flag, 0, 1)
			if swapped {
				fmt.Printf("  📊 Goroutine %d set flag successfully\n", id)
			} else {
				fmt.Printf("  📊 Goroutine %d failed to set flag\n", id)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  📊 Final flag value: %t\n", atomic.LoadInt32(&flag) == 1)
}

// ATOMIC POINTER OPERATIONS: Understanding atomic pointer operations
func atomicPointerOperations() {
	fmt.Println("Understanding atomic pointer operations...")
	
	// Atomic pointer operations
	var ptr unsafe.Pointer
	
	// Create some data
	data1 := &Data{Value: 42, Name: "first"}
	data2 := &Data{Value: 100, Name: "second"}
	
	// Store pointer
	atomic.StorePointer(&ptr, unsafe.Pointer(data1))
	fmt.Printf("  📊 Stored pointer: %+v\n", (*Data)(atomic.LoadPointer(&ptr)))
	
	// Swap pointer
	oldPtr := atomic.SwapPointer(&ptr, unsafe.Pointer(data2))
	fmt.Printf("  📊 Swapped pointer: old=%+v, new=%+v\n", 
		(*Data)(oldPtr), (*Data)(atomic.LoadPointer(&ptr)))
	
	// Compare and swap pointer
	swapped := atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(data2), unsafe.Pointer(data1))
	fmt.Printf("  📊 CompareAndSwapPointer: %t, current=%+v\n", 
		swapped, (*Data)(atomic.LoadPointer(&ptr)))
}

// ATOMIC VALUE OPERATIONS: Understanding atomic value operations
func atomicValueOperations() {
	fmt.Println("Understanding atomic value operations...")
	
	// Atomic value operations
	var value atomic.Value
	
	// Store value
	value.Store(42)
	fmt.Printf("  📊 Stored value: %v\n", value.Load())
	
	// Store different types (atomic.Value requires consistent types)
	// Note: atomic.Value requires consistent types, so we'll use separate values
	valueString := atomic.Value{}
	valueString.Store("hello")
	fmt.Printf("  📊 Stored string: %v\n", valueString.Load())
	
	// Create new atomic.Value for different types
	value2 := atomic.Value{}
	value2.Store([]int{1, 2, 3})
	fmt.Printf("  📊 Stored slice: %v\n", value2.Load())
	
	// Store struct
	value3 := atomic.Value{}
	person := Person{Name: "John", Age: 30}
	value3.Store(person)
	fmt.Printf("  📊 Stored struct: %+v\n", value3.Load())
	
	// Test concurrent access
	testConcurrentAtomicValue()
}

func testConcurrentAtomicValue() {
	var value atomic.Value
	var wg sync.WaitGroup
	goroutines := 10
	
	// Multiple goroutines storing values
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value.Store(fmt.Sprintf("goroutine-%d", id))
			time.Sleep(time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("  📊 Final value: %v\n", value.Load())
}

// MEMORY ORDERING: Understanding memory ordering
func memoryOrdering() {
	fmt.Println("Understanding memory ordering...")
	
	// Memory ordering example
	var x, y int32
	var wg sync.WaitGroup
	
	// Goroutine 1: Write x, then y
	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&x, 1)
		atomic.StoreInt32(&y, 1)
	}()
	
	// Goroutine 2: Read y, then x
	wg.Add(1)
	go func() {
		defer wg.Done()
		for atomic.LoadInt32(&y) == 0 {
			runtime.Gosched()
		}
		if atomic.LoadInt32(&x) == 0 {
			fmt.Printf("  📊 Memory ordering violation detected!\n")
		} else {
			fmt.Printf("  📊 Memory ordering preserved\n")
		}
	}()
	
	wg.Wait()
	
	// Test memory barriers
	testMemoryBarriers()
}

func testMemoryBarriers() {
	// Memory barrier example
	var data int32
	var ready int32
	
	// Writer goroutine
	go func() {
		data = 42
		atomic.StoreInt32(&ready, 1)
	}()
	
	// Reader goroutine
	go func() {
		for atomic.LoadInt32(&ready) == 0 {
			runtime.Gosched()
		}
		fmt.Printf("  📊 Data read: %d\n", data)
	}()
	
	time.Sleep(time.Millisecond)
}

// LOCK-FREE DATA STRUCTURES: Understanding lock-free programming
func lockFreeDataStructures() {
	fmt.Println("Understanding lock-free data structures...")
	
	// Lock-free stack
	stack := NewLockFreeStack()
	
	// Test concurrent operations
	var wg sync.WaitGroup
	goroutines := 10
	operations := 100
	
	start := time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				// Push
				stack.Push(id*1000 + j)
				// Pop
				stack.Pop()
			}
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("  📊 Lock-free stack: %v\n", duration)
	
	// Lock-free queue
	queue := NewLockFreeQueue()
	
	start = time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				// Enqueue
				queue.Enqueue(id*1000 + j)
				// Dequeue
				queue.Dequeue()
			}
		}(i)
	}
	
	wg.Wait()
	duration = time.Since(start)
	fmt.Printf("  📊 Lock-free queue: %v\n", duration)
}

// PERFORMANCE COMPARISON: Comparing atomic vs mutex performance
func performanceComparison() {
	fmt.Println("Understanding performance comparison...")
	
	// Test atomic vs mutex performance
	testAtomicVsMutex()
	
	// Test different atomic operations
	testAtomicOperationPerformance()
}

func testAtomicVsMutex() {
	// Atomic counter
	atomicCounter := &AtomicCounter{}
	
	// Mutex counter
	mutexCounter := &MutexCounter{}
	
	// Test atomic performance
	start := time.Now()
	testCounterPerformance(atomicCounter, "Atomic")
	atomicTime := time.Since(start)
	
	// Test mutex performance
	start = time.Now()
	testCounterPerformance(mutexCounter, "Mutex")
	mutexTime := time.Since(start)
	
	fmt.Printf("  📊 Atomic time: %v\n", atomicTime)
	fmt.Printf("  📊 Mutex time: %v\n", mutexTime)
	fmt.Printf("  📊 Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}

func testCounterPerformance(counter Counter, name string) {
	var wg sync.WaitGroup
	goroutines := 10
	operations := 10000
	
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
}

func testAtomicOperationPerformance() {
	var value int64
	
	// Test different atomic operations
	operations := 1000000
	
	// Load performance
	start := time.Now()
	for i := 0; i < operations; i++ {
		atomic.LoadInt64(&value)
	}
	loadTime := time.Since(start)
	
	// Store performance
	start = time.Now()
	for i := 0; i < operations; i++ {
		atomic.StoreInt64(&value, int64(i))
	}
	storeTime := time.Since(start)
	
	// Add performance
	start = time.Now()
	for i := 0; i < operations; i++ {
		atomic.AddInt64(&value, 1)
	}
	addTime := time.Since(start)
	
	// CompareAndSwap performance
	start = time.Now()
	for i := 0; i < operations; i++ {
		atomic.CompareAndSwapInt64(&value, int64(i), int64(i+1))
	}
	casTime := time.Since(start)
	
	fmt.Printf("  📊 Load time: %v\n", loadTime)
	fmt.Printf("  📊 Store time: %v\n", storeTime)
	fmt.Printf("  📊 Add time: %v\n", addTime)
	fmt.Printf("  📊 CompareAndSwap time: %v\n", casTime)
}

// IMPLEMENTATIONS

// Atomic Counter
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Mutex Counter for comparison
type MutexCounter struct {
	value int64
	mu    sync.Mutex
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Counter interface
type Counter interface {
	Increment()
	Value() int64
}

// Lock-Free Stack
type LockFreeStack struct {
	head unsafe.Pointer
}

type stackNode struct {
	value int
	next  unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	node := &stackNode{value: value}
	for {
		head := atomic.LoadPointer(&s.head)
		node.next = head
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}
		node := (*stackNode)(head)
		next := atomic.LoadPointer(&node.next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return node.value, true
		}
	}
}

// Lock-Free Queue
type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type queueNode struct {
	value int
	next  unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	dummy := &queueNode{}
	return &LockFreeQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *LockFreeQueue) Enqueue(value int) {
	node := &queueNode{value: value}
	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*queueNode)(tail).next)
		if next == nil {
			if atomic.CompareAndSwapPointer(&(*queueNode)(tail).next, nil, unsafe.Pointer(node)) {
				break
			}
		} else {
			atomic.CompareAndSwapPointer(&q.tail, tail, next)
		}
	}
	atomic.CompareAndSwapPointer(&q.tail, atomic.LoadPointer(&q.tail), unsafe.Pointer(node))
}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*queueNode)(head).next)
		
		if head == tail {
			if next == nil {
				return 0, false
			}
			atomic.CompareAndSwapPointer(&q.tail, tail, next)
		} else {
			if next == nil {
				continue
			}
			value := (*queueNode)(next).value
			if atomic.CompareAndSwapPointer(&q.head, head, next) {
				return value, true
			}
		}
	}
}

// Data structures
type Data struct {
	Value int
	Name  string
}

type Person struct {
	Name string
	Age  int
}
