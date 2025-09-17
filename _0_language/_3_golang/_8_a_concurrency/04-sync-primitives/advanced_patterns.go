package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Pattern 1: Thread-Safe Counter with Metrics
type SafeCounter struct {
	mu       sync.RWMutex
	counters map[string]int64
	metrics  map[string]int64
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counters: make(map[string]int64),
		metrics:  make(map[string]int64),
	}
}

func (sc *SafeCounter) Increment(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counters[key]++
	sc.metrics["increments"]++
}

func (sc *SafeCounter) Get(key string) int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counters[key]
}

func (sc *SafeCounter) GetAllCounters() map[string]int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sc.counters {
		result[k] = v
	}
	return result
}

func (sc *SafeCounter) GetMetrics() map[string]int64 {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sc.metrics {
		result[k] = v
	}
	return result
}

// Advanced Pattern 2: Read-Write Lock with Priority
type PriorityRWMutex struct {
	mu           sync.RWMutex
	readerCount  int64
	writerCount  int64
	readerWait   int64
	writerWait   int64
	readerSignal chan struct{}
	writerSignal chan struct{}
}

func NewPriorityRWMutex() *PriorityRWMutex {
	return &PriorityRWMutex{
		readerSignal: make(chan struct{}),
		writerSignal: make(chan struct{}),
	}
}

func (prw *PriorityRWMutex) RLock() {
	atomic.AddInt64(&prw.readerWait, 1)
	prw.mu.RLock()
	atomic.AddInt64(&prw.readerCount, 1)
	atomic.AddInt64(&prw.readerWait, -1)
}

func (prw *PriorityRWMutex) RUnlock() {
	atomic.AddInt64(&prw.readerCount, -1)
	prw.mu.RUnlock()
}

func (prw *PriorityRWMutex) Lock() {
	atomic.AddInt64(&prw.writerWait, 1)
	prw.mu.Lock()
	atomic.AddInt64(&prw.writerCount, 1)
	atomic.AddInt64(&prw.writerWait, -1)
}

func (prw *PriorityRWMutex) Unlock() {
	atomic.AddInt64(&prw.writerCount, -1)
	prw.mu.Unlock()
}

func (prw *PriorityRWMutex) GetStats() (readers, writers, readerWait, writerWait int64) {
	return atomic.LoadInt64(&prw.readerCount),
		atomic.LoadInt64(&prw.writerCount),
		atomic.LoadInt64(&prw.readerWait),
		atomic.LoadInt64(&prw.writerWait)
}

// Advanced Pattern 3: WaitGroup with Timeout
type TimeoutWaitGroup struct {
	wg      sync.WaitGroup
	timeout time.Duration
	done    chan struct{}
}

func NewTimeoutWaitGroup(timeout time.Duration) *TimeoutWaitGroup {
	return &TimeoutWaitGroup{
		timeout: timeout,
		done:    make(chan struct{}),
	}
}

func (twg *TimeoutWaitGroup) Add(delta int) {
	twg.wg.Add(delta)
}

func (twg *TimeoutWaitGroup) Done() {
	twg.wg.Done()
}

func (twg *TimeoutWaitGroup) Wait() error {
	go func() {
		twg.wg.Wait()
		close(twg.done)
	}()
	
	select {
	case <-twg.done:
		return nil
	case <-time.After(twg.timeout):
		return fmt.Errorf("waitgroup timeout after %v", twg.timeout)
	}
}

// Advanced Pattern 4: Once with Error Handling
type SafeOnce struct {
	once sync.Once
	err  error
}

func (so *SafeOnce) Do(fn func() error) error {
	so.once.Do(func() {
		so.err = fn()
	})
	return so.err
}

// Advanced Pattern 5: Condition Variable with Timeout
type TimeoutCond struct {
	mu    sync.Mutex
	cond  *sync.Cond
	ready bool
}

func NewTimeoutCond() *TimeoutCond {
	tc := &TimeoutCond{}
	tc.cond = sync.NewCond(&tc.mu)
	return tc
}

func (tc *TimeoutCond) Wait() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	for !tc.ready {
		tc.cond.Wait()
	}
}

func (tc *TimeoutCond) WaitWithTimeout(timeout time.Duration) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	
	done := make(chan struct{})
	go func() {
		for !tc.ready {
			tc.cond.Wait()
		}
		close(done)
	}()
	
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (tc *TimeoutCond) Signal() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.ready = true
	tc.cond.Signal()
}

func (tc *TimeoutCond) Broadcast() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.ready = true
	tc.cond.Broadcast()
}

// Advanced Pattern 6: Atomic Counter with Statistics
type AtomicCounter struct {
	value      int64
	increments int64
	decrements int64
	resets     int64
}

func (ac *AtomicCounter) Increment() int64 {
	atomic.AddInt64(&ac.increments, 1)
	return atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Decrement() int64 {
	atomic.AddInt64(&ac.decrements, 1)
	return atomic.AddInt64(&ac.value, -1)
}

func (ac *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&ac.value)
}

func (ac *AtomicCounter) Reset() int64 {
	atomic.AddInt64(&ac.resets, 1)
	return atomic.SwapInt64(&ac.value, 0)
}

func (ac *AtomicCounter) GetStats() (value, increments, decrements, resets int64) {
	return atomic.LoadInt64(&ac.value),
		atomic.LoadInt64(&ac.increments),
		atomic.LoadInt64(&ac.decrements),
		atomic.LoadInt64(&ac.resets)
}

// Advanced Pattern 7: Concurrent Map with Statistics
type StatsMap struct {
	mu     sync.RWMutex
	data   map[string]interface{}
	stats  map[string]int64
}

func NewStatsMap() *StatsMap {
	return &StatsMap{
		data:  make(map[string]interface{}),
		stats: make(map[string]int64),
	}
}

func (sm *StatsMap) Store(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
	sm.stats["stores"]++
}

func (sm *StatsMap) Load(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	if ok {
		sm.stats["loads"]++
	} else {
		sm.stats["misses"]++
	}
	return value, ok
}

func (sm *StatsMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, exists := sm.data[key]; exists {
		delete(sm.data, key)
		sm.stats["deletes"]++
	}
}

func (sm *StatsMap) GetStats() map[string]int64 {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	result := make(map[string]int64)
	for k, v := range sm.stats {
		result[k] = v
	}
	return result
}

// Advanced Pattern 8: Object Pool with Statistics
type StatsPool struct {
	pool     sync.Pool
	created  int64
	reused   int64
	returned int64
}

func NewStatsPool(newFunc func() interface{}) *StatsPool {
	return &StatsPool{
		pool: sync.Pool{New: newFunc},
	}
}

func (sp *StatsPool) Get() interface{} {
	obj := sp.pool.Get()
	if obj == nil {
		atomic.AddInt64(&sp.created, 1)
	} else {
		atomic.AddInt64(&sp.reused, 1)
	}
	return obj
}

func (sp *StatsPool) Put(obj interface{}) {
	atomic.AddInt64(&sp.returned, 1)
	sp.pool.Put(obj)
}

func (sp *StatsPool) GetStats() (created, reused, returned int64) {
	return atomic.LoadInt64(&sp.created),
		atomic.LoadInt64(&sp.reused),
		atomic.LoadInt64(&sp.returned)
}

// Advanced Pattern 9: Barrier Synchronization
type Barrier struct {
	mu      sync.Mutex
	cond    *sync.Cond
	count   int
	target  int
	phase   int
}

func NewBarrier(target int) *Barrier {
	b := &Barrier{target: target}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Barrier) Wait() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	phase := b.phase
	b.count++
	
	if b.count == b.target {
		b.count = 0
		b.phase++
		b.cond.Broadcast()
	} else {
		for phase == b.phase {
			b.cond.Wait()
		}
	}
	
	return phase
}

// Advanced Pattern 10: Semaphore
type Semaphore struct {
	permits chan struct{}
}

func NewSemaphore(permits int) *Semaphore {
	s := &Semaphore{
		permits: make(chan struct{}, permits),
	}
	
	// Fill with permits
	for i := 0; i < permits; i++ {
		s.permits <- struct{}{}
	}
	
	return s
}

func (s *Semaphore) Acquire() {
	<-s.permits
}

func (s *Semaphore) Release() {
	select {
	case s.permits <- struct{}{}:
	default:
		// Semaphore is full
	}
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case <-s.permits:
		return true
	default:
		return false
	}
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Synchronization Patterns")
	fmt.Println("====================================")
	
	// Pattern 1: Thread-Safe Counter with Metrics
	fmt.Println("\n1. Thread-Safe Counter with Metrics:")
	counter := NewSafeCounter()
	
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				counter.Increment(fmt.Sprintf("key%d", id))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Counters: %v\n", counter.GetAllCounters())
	fmt.Printf("Metrics: %v\n", counter.GetMetrics())
	
	// Pattern 2: Priority RWMutex
	fmt.Println("\n2. Priority RWMutex:")
	prw := NewPriorityRWMutex()
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			prw.RLock()
			fmt.Printf("Reader %d: reading\n", id)
			time.Sleep(100 * time.Millisecond)
			prw.RUnlock()
		}(i)
	}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		prw.Lock()
		fmt.Println("Writer: writing")
		time.Sleep(100 * time.Millisecond)
		prw.Unlock()
	}()
	
	wg.Wait()
	readers, writers, readerWait, writerWait := prw.GetStats()
	fmt.Printf("Stats: readers=%d, writers=%d, readerWait=%d, writerWait=%d\n", readers, writers, readerWait, writerWait)
	
	// Pattern 3: WaitGroup with Timeout
	fmt.Println("\n3. WaitGroup with Timeout:")
	twg := NewTimeoutWaitGroup(500 * time.Millisecond)
	
	for i := 0; i < 3; i++ {
		twg.Add(1)
		go func(id int) {
			defer twg.Done()
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Worker %d completed\n", id)
		}(i)
	}
	
	if err := twg.Wait(); err != nil {
		fmt.Printf("WaitGroup error: %v\n", err)
	} else {
		fmt.Println("All workers completed")
	}
	
	// Pattern 4: Once with Error Handling
	fmt.Println("\n4. Once with Error Handling:")
	so := &SafeOnce{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := so.Do(func() error {
				fmt.Printf("Initializing from goroutine %d\n", id)
				if id == 1 {
					return fmt.Errorf("initialization failed")
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Goroutine %d: error: %v\n", id, err)
			} else {
				fmt.Printf("Goroutine %d: success\n", id)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Pattern 5: Condition Variable with Timeout
	fmt.Println("\n5. Condition Variable with Timeout:")
	tc := NewTimeoutCond()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		if tc.WaitWithTimeout(1 * time.Second) {
			fmt.Println("Condition met within timeout")
		} else {
			fmt.Println("Timeout waiting for condition")
		}
	}()
	
	time.Sleep(500 * time.Millisecond)
	tc.Signal()
	wg.Wait()
	
	// Pattern 6: Atomic Counter with Statistics
	fmt.Println("\n6. Atomic Counter with Statistics:")
	ac := &AtomicCounter{}
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ac.Increment()
			}
		}(i)
	}
	
	wg.Wait()
	value, increments, decrements, resets := ac.GetStats()
	fmt.Printf("Counter stats: value=%d, increments=%d, decrements=%d, resets=%d\n", value, increments, decrements, resets)
	
	// Pattern 7: Concurrent Map with Statistics
	fmt.Println("\n7. Concurrent Map with Statistics:")
	sm := NewStatsMap()
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			sm.Store(key, fmt.Sprintf("value%d", id))
			sm.Load(key)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Map stats: %v\n", sm.GetStats())
	
	// Pattern 8: Object Pool with Statistics
	fmt.Println("\n8. Object Pool with Statistics:")
	sp := NewStatsPool(func() interface{} {
		return &Buffer{ID: time.Now().UnixNano()}
	})
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			buf := sp.Get().(*Buffer)
			time.Sleep(10 * time.Millisecond)
			sp.Put(buf)
		}(i)
	}
	
	wg.Wait()
	created, reused, returned := sp.GetStats()
	fmt.Printf("Pool stats: created=%d, reused=%d, returned=%d\n", created, reused, returned)
	
	// Pattern 9: Barrier Synchronization
	fmt.Println("\n9. Barrier Synchronization:")
	barrier := NewBarrier(3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: before barrier\n", id)
			phase := barrier.Wait()
			fmt.Printf("Goroutine %d: after barrier phase %d\n", id, phase)
		}(i)
	}
	
	wg.Wait()
	
	// Pattern 10: Semaphore
	fmt.Println("\n10. Semaphore:")
	sem := NewSemaphore(2)
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			fmt.Printf("Goroutine %d: acquired semaphore\n", id)
			time.Sleep(100 * time.Millisecond)
			sem.Release()
			fmt.Printf("Goroutine %d: released semaphore\n", id)
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}
