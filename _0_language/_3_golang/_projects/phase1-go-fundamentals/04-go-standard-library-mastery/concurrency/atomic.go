package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Custom types for demonstration
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() int64 {
	return atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Decrement() int64 {
	return atomic.AddInt64(&c.value, -1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func (c *AtomicCounter) Set(value int64) {
	atomic.StoreInt64(&c.value, value)
}

func (c *AtomicCounter) Add(delta int64) int64 {
	return atomic.AddInt64(&c.value, delta)
}

func (c *AtomicCounter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&c.value, old, new)
}

func (c *AtomicCounter) Swap(new int64) int64 {
	return atomic.SwapInt64(&c.value, new)
}

type AtomicFlag struct {
	value int32
}

func (f *AtomicFlag) Set() {
	atomic.StoreInt32(&f.value, 1)
}

func (f *AtomicFlag) Clear() {
	atomic.StoreInt32(&f.value, 0)
}

func (f *AtomicFlag) IsSet() bool {
	return atomic.LoadInt32(&f.value) == 1
}

func (f *AtomicFlag) Toggle() bool {
	for {
		old := atomic.LoadInt32(&f.value)
		new := 1 - old
		if atomic.CompareAndSwapInt32(&f.value, old, new) {
			return new == 1
		}
	}
}

type AtomicState struct {
	state int32
}

const (
	Idle State = iota
	Running
	Paused
	Stopped
)

type State int32

func (s *AtomicState) SetState(newState State) {
	atomic.StoreInt32(&s.state, int32(newState))
}

func (s *AtomicState) GetState() State {
	return State(atomic.LoadInt32(&s.state))
}

func (s *AtomicState) CompareAndSwapState(oldState, newState State) bool {
	return atomic.CompareAndSwapInt32(&s.state, int32(oldState), int32(newState))
}

func (s *AtomicState) String() string {
	switch s.GetState() {
	case Idle:
		return "Idle"
	case Running:
		return "Running"
	case Paused:
		return "Paused"
	case Stopped:
		return "Stopped"
	default:
		return "Unknown"
	}
}

type AtomicPointer struct {
	ptr unsafe.Pointer
}

func (p *AtomicPointer) Store(ptr unsafe.Pointer) {
	atomic.StorePointer(&p.ptr, ptr)
}

func (p *AtomicPointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&p.ptr)
}

func (p *AtomicPointer) Swap(ptr unsafe.Pointer) unsafe.Pointer {
	return atomic.SwapPointer(&p.ptr, ptr)
}

func (p *AtomicPointer) CompareAndSwap(old, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&p.ptr, old, new)
}

type AtomicValue struct {
	value atomic.Value
}

func (v *AtomicValue) Store(value interface{}) {
	v.value.Store(value)
}

func (v *AtomicValue) Load() interface{} {
	return v.value.Load()
}

func (v *AtomicValue) Swap(new interface{}) interface{} {
	return v.value.Swap(new)
}

func (v *AtomicValue) CompareAndSwap(old, new interface{}) bool {
	return v.value.CompareAndSwap(old, new)
}

// Lock-free stack implementation
type Node struct {
	value int
	next  *Node
}

type LockFreeStack struct {
	head unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	node := &Node{value: value}
	for {
		head := atomic.LoadPointer(&s.head)
		node.next = (*Node)(head)
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			return
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}
		node := (*Node)(head)
		next := node.next
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(next)) {
			return node.value, true
		}
	}
}

func (s *LockFreeStack) IsEmpty() bool {
	return atomic.LoadPointer(&s.head) == nil
}

// Atomic ring buffer
type AtomicRingBuffer struct {
	buffer []int
	head   int64
	tail   int64
	size   int64
}

func NewAtomicRingBuffer(size int) *AtomicRingBuffer {
	return &AtomicRingBuffer{
		buffer: make([]int, size),
		size:   int64(size),
	}
}

func (rb *AtomicRingBuffer) Push(value int) bool {
	tail := atomic.LoadInt64(&rb.tail)
	head := atomic.LoadInt64(&rb.head)
	
	if (tail+1)%rb.size == head {
		return false // Buffer full
	}
	
	rb.buffer[tail] = value
	atomic.StoreInt64(&rb.tail, (tail+1)%rb.size)
	return true
}

func (rb *AtomicRingBuffer) Pop() (int, bool) {
	head := atomic.LoadInt64(&rb.head)
	tail := atomic.LoadInt64(&rb.tail)
	
	if head == tail {
		return 0, false // Buffer empty
	}
	
	value := rb.buffer[head]
	atomic.StoreInt64(&rb.head, (head+1)%rb.size)
	return value, true
}

func (rb *AtomicRingBuffer) IsEmpty() bool {
	head := atomic.LoadInt64(&rb.head)
	tail := atomic.LoadInt64(&rb.tail)
	return head == tail
}

func (rb *AtomicRingBuffer) IsFull() bool {
	head := atomic.LoadInt64(&rb.head)
	tail := atomic.LoadInt64(&rb.tail)
	return (tail+1)%rb.size == head
}

// Atomic semaphore
type AtomicSemaphore struct {
	count int64
	max   int64
}

func NewAtomicSemaphore(max int64) *AtomicSemaphore {
	return &AtomicSemaphore{
		count: max,
		max:   max,
	}
}

func (s *AtomicSemaphore) Acquire() bool {
	for {
		count := atomic.LoadInt64(&s.count)
		if count <= 0 {
			return false
		}
		if atomic.CompareAndSwapInt64(&s.count, count, count-1) {
			return true
		}
	}
}

func (s *AtomicSemaphore) Release() {
	atomic.AddInt64(&s.count, 1)
}

func (s *AtomicSemaphore) Count() int64 {
	return atomic.LoadInt64(&s.count)
}

func main() {
	fmt.Println("🚀 Go sync/atomic Package Mastery Examples")
	fmt.Println("==========================================")

	// 1. Basic Atomic Operations
	fmt.Println("\n1. Basic Atomic Operations:")
	
	var counter int64
	
	// Atomic increment
	atomic.AddInt64(&counter, 1)
	fmt.Printf("Counter after increment: %d\n", counter)
	
	// Atomic load
	value := atomic.LoadInt64(&counter)
	fmt.Printf("Counter value: %d\n", value)
	
	// Atomic store
	atomic.StoreInt64(&counter, 100)
	fmt.Printf("Counter after store: %d\n", counter)
	
	// Atomic compare and swap
	swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
	fmt.Printf("CAS operation: %t, new value: %d\n", swapped, atomic.LoadInt64(&counter))
	
	// Atomic swap
	oldValue := atomic.SwapInt64(&counter, 300)
	fmt.Printf("Swap operation: old=%d, new=%d\n", oldValue, atomic.LoadInt64(&counter))

	// 2. Atomic Counter
	fmt.Println("\n2. Atomic Counter:")
	
	atomicCounter := &AtomicCounter{}
	
	// Increment counter concurrently
	var wg sync.WaitGroup
	numGoroutines := 100
	numIncrements := 1000
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				atomicCounter.Increment()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d (expected: %d)\n", atomicCounter.Value(), numGoroutines*numIncrements)

	// 3. Atomic Flag
	fmt.Println("\n3. Atomic Flag:")
	
	flag := &AtomicFlag{}
	
	// Test flag operations
	fmt.Printf("Flag initially set: %t\n", flag.IsSet())
	
	flag.Set()
	fmt.Printf("Flag after set: %t\n", flag.IsSet())
	
	flag.Clear()
	fmt.Printf("Flag after clear: %t\n", flag.IsSet())
	
	// Toggle flag
	toggled := flag.Toggle()
	fmt.Printf("Flag after toggle: %t (was: %t)\n", flag.IsSet(), toggled)
	
	toggled = flag.Toggle()
	fmt.Printf("Flag after toggle: %t (was: %t)\n", flag.IsSet(), toggled)

	// 4. Atomic State
	fmt.Println("\n4. Atomic State:")
	
	state := &AtomicState{}
	
	// Test state transitions
	fmt.Printf("Initial state: %s\n", state.String())
	
	state.SetState(Running)
	fmt.Printf("State after set to Running: %s\n", state.String())
	
	state.SetState(Paused)
	fmt.Printf("State after set to Paused: %s\n", state.String())
	
	// Test compare and swap
	swapped = state.CompareAndSwapState(Paused, Stopped)
	fmt.Printf("CAS Paused->Stopped: %t, new state: %s\n", swapped, state.String())
	
	swapped = state.CompareAndSwapState(Running, Idle)
	fmt.Printf("CAS Running->Idle: %t, new state: %s\n", swapped, state.String())

	// 5. Atomic Pointer
	fmt.Println("\n5. Atomic Pointer:")
	
	atomicPtr := &AtomicPointer{}
	
	// Test pointer operations
	var data1, data2 int = 42, 84
	
	atomicPtr.Store(unsafe.Pointer(&data1))
	ptr1 := atomicPtr.Load()
	fmt.Printf("Stored pointer to data1: %p, loaded: %p\n", &data1, ptr1)
	
	ptr2 := atomicPtr.Swap(unsafe.Pointer(&data2))
	fmt.Printf("Swapped pointer: old=%p, new=%p\n", ptr2, atomicPtr.Load())
	
	swapped = atomicPtr.CompareAndSwap(unsafe.Pointer(&data2), unsafe.Pointer(&data1))
	fmt.Printf("CAS pointer: %t, current: %p\n", swapped, atomicPtr.Load())

	// 6. Atomic Value
	fmt.Println("\n6. Atomic Value:")
	
	atomicValue := &AtomicValue{}
	
	// Test value operations
	atomicValue.Store("Hello")
	value2 := atomicValue.Load()
	fmt.Printf("Stored 'Hello', loaded: %v\n", value2)
	
	oldValue2 := atomicValue.Swap("World")
	fmt.Printf("Swapped value: old=%v, new=%v\n", oldValue2, atomicValue.Load())
	
	swapped = atomicValue.CompareAndSwap("World", "Go")
	fmt.Printf("CAS value: %t, current: %v\n", swapped, atomicValue.Load())

	// 7. Lock-Free Stack
	fmt.Println("\n7. Lock-Free Stack:")
	
	stack := NewLockFreeStack()
	
	// Push values
	for i := 1; i <= 5; i++ {
		stack.Push(i)
		fmt.Printf("Pushed: %d\n", i)
	}
	
	// Pop values
	for !stack.IsEmpty() {
		value, ok := stack.Pop()
		if ok {
			fmt.Printf("Popped: %d\n", value)
		}
	}

	// 8. Atomic Ring Buffer
	fmt.Println("\n8. Atomic Ring Buffer:")
	
	ringBuffer := NewAtomicRingBuffer(5)
	
	// Push values
	for i := 1; i <= 7; i++ {
		success := ringBuffer.Push(i)
		fmt.Printf("Push %d: %t\n", i, success)
	}
	
	// Pop values
	for !ringBuffer.IsEmpty() {
		value, ok := ringBuffer.Pop()
		if ok {
			fmt.Printf("Pop: %d\n", value)
		}
	}

	// 9. Atomic Semaphore
	fmt.Println("\n9. Atomic Semaphore:")
	
	semaphore := NewAtomicSemaphore(3)
	
	// Test semaphore
	fmt.Printf("Initial semaphore count: %d\n", semaphore.Count())
	
	// Acquire permits
	for i := 0; i < 5; i++ {
		acquired := semaphore.Acquire()
		fmt.Printf("Acquire %d: %t, count: %d\n", i+1, acquired, semaphore.Count())
	}
	
	// Release permits
	for i := 0; i < 3; i++ {
		semaphore.Release()
		fmt.Printf("Release %d, count: %d\n", i+1, semaphore.Count())
	}

	// 10. Performance Comparison
	fmt.Println("\n10. Performance Comparison:")
	
	// Test mutex vs atomic performance
	const iterations = 1000000
	
	// Mutex version
	var mutexCounter int64
	var mu sync.Mutex
	var mutexWg sync.WaitGroup
	
	start := time.Now()
	for i := 0; i < 100; i++ {
		mutexWg.Add(1)
		go func() {
			defer mutexWg.Done()
			for j := 0; j < iterations/100; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	mutexWg.Wait()
	mutexTime := time.Since(start)
	
	// Atomic version
	var atomicCounter2 int64
	var atomicWg sync.WaitGroup
	
	start = time.Now()
	for i := 0; i < 100; i++ {
		atomicWg.Add(1)
		go func() {
			defer atomicWg.Done()
			for j := 0; j < iterations/100; j++ {
				atomic.AddInt64(&atomicCounter2, 1)
			}
		}()
	}
	atomicWg.Wait()
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v\n", mutexTime)
	fmt.Printf("Atomic time: %v\n", atomicTime)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))

	// 11. Atomic Operations with Different Types
	fmt.Println("\n11. Atomic Operations with Different Types:")
	
	// Int32 operations
	var int32Val int32
	atomic.StoreInt32(&int32Val, 42)
	fmt.Printf("Int32 value: %d\n", atomic.LoadInt32(&int32Val))
	
	// Uint64 operations
	var uint64Val uint64
	atomic.StoreUint64(&uint64Val, 100)
	fmt.Printf("Uint64 value: %d\n", atomic.LoadUint64(&uint64Val))
	
	// Uint32 operations
	var uint32Val uint32
	atomic.StoreUint32(&uint32Val, 200)
	fmt.Printf("Uint32 value: %d\n", atomic.LoadUint32(&uint32Val))

	// 12. Atomic Operations with Pointers
	fmt.Println("\n12. Atomic Operations with Pointers:")
	
	var ptr unsafe.Pointer
	var data int = 42
	
	// Store pointer
	atomic.StorePointer(&ptr, unsafe.Pointer(&data))
	fmt.Printf("Stored pointer: %p\n", atomic.LoadPointer(&ptr))
	
	// Swap pointer
	var newData int = 84
	oldPtr := atomic.SwapPointer(&ptr, unsafe.Pointer(&newData))
	fmt.Printf("Swapped pointer: old=%p, new=%p\n", oldPtr, atomic.LoadPointer(&ptr))
	
	// Compare and swap pointer
	swapped = atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(&newData), unsafe.Pointer(&data))
	fmt.Printf("CAS pointer: %t, current: %p\n", swapped, atomic.LoadPointer(&ptr))

	// 13. Atomic Operations with Values
	fmt.Println("\n13. Atomic Operations with Values:")
	
	// Store different types in separate atomic.Value instances
	var stringValue atomic.Value
	var intValue atomic.Value
	var sliceValue atomic.Value
	
	stringValue.Store("string")
	fmt.Printf("Stored string: %v\n", stringValue.Load())
	
	intValue.Store(42)
	fmt.Printf("Stored int: %v\n", intValue.Load())
	
	sliceValue.Store([]int{1, 2, 3})
	fmt.Printf("Stored slice: %v\n", sliceValue.Load())

	// 14. Memory Ordering
	fmt.Println("\n14. Memory Ordering:")
	
	// Demonstrate memory ordering with atomic operations
	var x, y int64
	var wg2 sync.WaitGroup
	
	// Writer
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 1)
	}()
	
	// Reader
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for atomic.LoadInt64(&y) == 0 {
			// Busy wait
		}
		fmt.Printf("x = %d, y = %d\n", atomic.LoadInt64(&x), atomic.LoadInt64(&y))
	}()
	
	wg2.Wait()

	// 15. Atomic Operations in Loops
	fmt.Println("\n15. Atomic Operations in Loops:")
	
	var loopCounter int64
	
	// Increment in loop
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&loopCounter, 1)
	}
	fmt.Printf("Loop counter: %d\n", atomic.LoadInt64(&loopCounter))
	
	// Compare and swap in loop
	target := int64(5)
	for {
		old := atomic.LoadInt64(&loopCounter)
		if old >= target {
			break
		}
		if atomic.CompareAndSwapInt64(&loopCounter, old, old+1) {
			fmt.Printf("CAS succeeded: %d\n", atomic.LoadInt64(&loopCounter))
		}
	}

	// 16. Atomic Operations with Timeouts
	fmt.Println("\n16. Atomic Operations with Timeouts:")
	
	var timeoutCounter int64
	timeout := time.After(100 * time.Millisecond)
	
	go func() {
		for {
			select {
			case <-timeout:
				return
			default:
				atomic.AddInt64(&timeoutCounter, 1)
			}
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Timeout counter: %d\n", atomic.LoadInt64(&timeoutCounter))

	// 17. Atomic Operations with Channels
	fmt.Println("\n17. Atomic Operations with Channels:")
	
	var channelCounter int64
	ch := make(chan int, 10)
	
	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			atomic.AddInt64(&channelCounter, 1)
			ch <- i
		}
		close(ch)
	}()
	
	// Consumer
	for range ch {
		// Process values
	}
	
	fmt.Printf("Channel counter: %d\n", atomic.LoadInt64(&channelCounter))

	// 18. Runtime Statistics
	fmt.Println("\n18. Runtime Statistics:")
	
	// Get goroutine count
	goroutineCount := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines: %d\n", goroutineCount)
	
	// Get memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory allocated: %d bytes\n", m.Alloc)
	fmt.Printf("Number of GC cycles: %d\n", m.NumGC)

	fmt.Println("\n🎉 sync/atomic Package Mastery Complete!")
}
