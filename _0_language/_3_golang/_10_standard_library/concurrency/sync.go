package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Custom types for demonstration
type Counter struct {
	mu    sync.RWMutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func (c *Counter) Set(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count = value
}

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Keys() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	keys := make([]string, 0, len(sm.data))
	for k := range sm.data {
		keys = append(keys, k)
	}
	return keys
}

type WorkerPool struct {
	workers    int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.Mutex
	activeJobs int
}

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID int
	Data  string
	Error error
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		jobs:    make(chan Job, 100),
		results: make(chan Result, 100),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobs {
		wp.mu.Lock()
		wp.activeJobs++
		wp.mu.Unlock()
		
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		result := Result{
			JobID: job.ID,
			Data:  fmt.Sprintf("Processed by worker %d: %s", id, job.Data),
		}
		
		wp.results <- result
		
		wp.mu.Lock()
		wp.activeJobs--
		wp.mu.Unlock()
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) ActiveJobs() int {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	return wp.activeJobs
}

type Singleton struct {
	ID   int
	Name string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			ID:   1,
			Name: "Singleton Instance",
		}
	})
	return instance
}

type Barrier struct {
	mu      sync.Mutex
	count   int
	target  int
	cond    *sync.Cond
	waiting bool
}

func NewBarrier(target int) *Barrier {
	b := &Barrier{target: target}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Barrier) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.count++
	if b.count >= b.target {
		b.waiting = true
		b.cond.Broadcast()
	} else {
		for !b.waiting {
			b.cond.Wait()
		}
	}
}

type RateLimiter struct {
	mu       sync.Mutex
	requests int
	lastTime time.Time
	rate     int
	window   time.Duration
}

func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		window: window,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	if now.Sub(rl.lastTime) > rl.window {
		rl.requests = 0
		rl.lastTime = now
	}
	
	if rl.requests >= rl.rate {
		return false
	}
	
	rl.requests++
	return true
}

func main() {
	fmt.Println("🚀 Go sync Package Mastery Examples")
	fmt.Println("===================================")

	// 1. Basic Mutex Operations
	fmt.Println("\n1. Basic Mutex Operations:")
	
	counter := &Counter{}
	
	// Increment counter concurrently
	var wg sync.WaitGroup
	numGoroutines := 100
	numIncrements := 1000
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter value: %d (expected: %d)\n", counter.Value(), numGoroutines*numIncrements)

	// 2. RWMutex Operations
	fmt.Println("\n2. RWMutex Operations:")
	
	safeMap := NewSafeMap()
	
	// Write operations
	var writeWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		writeWg.Add(1)
		go func(i int) {
			defer writeWg.Done()
			key := fmt.Sprintf("key%d", i)
			value := i * 10
			safeMap.Set(key, value)
			fmt.Printf("Set %s = %d\n", key, value)
		}(i)
	}
	writeWg.Wait()
	
	// Read operations
	var readWg sync.WaitGroup
	for i := 0; i < 20; i++ {
		readWg.Add(1)
		go func(i int) {
			defer readWg.Done()
			key := fmt.Sprintf("key%d", i%10)
			if value, exists := safeMap.Get(key); exists {
				fmt.Printf("Read %s = %d\n", key, value)
			}
		}(i)
	}
	readWg.Wait()
	
	fmt.Printf("Map keys: %v\n", safeMap.Keys())

	// 3. WaitGroup Patterns
	fmt.Println("\n3. WaitGroup Patterns:")
	
	// Basic WaitGroup
	var basicWg sync.WaitGroup
	results := make([]int, 10)
	
	for i := 0; i < 10; i++ {
		basicWg.Add(1)
		go func(i int) {
			defer basicWg.Done()
			// Simulate work
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			results[i] = i * i
		}(i)
	}
	
	basicWg.Wait()
	fmt.Printf("Results: %v\n", results)

	// 4. Once Pattern
	fmt.Println("\n4. Once Pattern:")
	
	// Test singleton pattern
	var onceWg sync.WaitGroup
	instances := make([]*Singleton, 10)
	
	for i := 0; i < 10; i++ {
		onceWg.Add(1)
		go func(i int) {
			defer onceWg.Done()
			instances[i] = GetInstance()
		}(i)
	}
	
	onceWg.Wait()
	
	// Check if all instances are the same
	allSame := true
	for i := 1; i < len(instances); i++ {
		if instances[i] != instances[0] {
			allSame = false
			break
		}
	}
	
	fmt.Printf("All instances same: %t\n", allSame)
	fmt.Printf("Instance: %+v\n", instances[0])

	// 5. Condition Variables
	fmt.Println("\n5. Condition Variables:")
	
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	// Producer
	go func() {
		time.Sleep(100 * time.Millisecond)
		mu.Lock()
		ready = true
		cond.Signal()
		mu.Unlock()
		fmt.Println("Producer: Data ready")
	}()
	
	// Consumer
	mu.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Consumer: Data received")
	mu.Unlock()

	// 6. Atomic Operations
	fmt.Println("\n6. Atomic Operations:")
	
	var atomicCounter int64
	var atomicWg sync.WaitGroup
	
	// Atomic increment
	for i := 0; i < 1000; i++ {
		atomicWg.Add(1)
		go func() {
			defer atomicWg.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	
	atomicWg.Wait()
	fmt.Printf("Atomic counter: %d\n", atomic.LoadInt64(&atomicCounter))
	
	// Atomic compare and swap
	oldValue := atomic.LoadInt64(&atomicCounter)
	newValue := oldValue + 100
	swapped := atomic.CompareAndSwapInt64(&atomicCounter, oldValue, newValue)
	fmt.Printf("CAS operation: %t, new value: %d\n", swapped, atomic.LoadInt64(&atomicCounter))

	// 7. Worker Pool Pattern
	fmt.Println("\n7. Worker Pool Pattern:")
	
	pool := NewWorkerPool(5)
	pool.Start()
	
	// Submit jobs
	go func() {
		for i := 0; i < 20; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Job %d", i),
			}
			pool.Submit(job)
		}
		pool.Close()
	}()
	
	// Collect results
	var resultWg sync.WaitGroup
	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		for result := range pool.results {
			fmt.Printf("Result: %s\n", result.Data)
		}
	}()
	
	resultWg.Wait()

	// 8. Rate Limiting
	fmt.Println("\n8. Rate Limiting:")
	
	limiter := NewRateLimiter(5, time.Second)
	
	// Test rate limiting
	for i := 0; i < 10; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request %d: %t\n", i+1, allowed)
		time.Sleep(100 * time.Millisecond)
	}

	// 9. Deadlock Prevention
	fmt.Println("\n9. Deadlock Prevention:")
	
	// Example of proper lock ordering
	type Account struct {
		mu      sync.Mutex
		balance int
		id      int
	}
	
	transfer := func(from, to *Account, amount int) {
		// Always lock in the same order (by ID)
		if from.id < to.id {
			from.mu.Lock()
			to.mu.Lock()
		} else {
			to.mu.Lock()
			from.mu.Lock()
		}
		defer from.mu.Unlock()
		defer to.mu.Unlock()
		
		if from.balance >= amount {
			from.balance -= amount
			to.balance += amount
			fmt.Printf("Transferred %d from account %d to account %d\n", amount, from.id, to.id)
		}
	}
	
	account1 := &Account{balance: 1000, id: 1}
	account2 := &Account{balance: 500, id: 2}
	
	transfer(account1, account2, 200)
	fmt.Printf("Account 1 balance: %d\n", account1.balance)
	fmt.Printf("Account 2 balance: %d\n", account2.balance)

	// 10. Performance Comparison
	fmt.Println("\n10. Performance Comparison:")
	
	// Test mutex vs atomic performance
	const iterations = 1000000
	
	// Mutex version
	start := time.Now()
	var mutexCounter int64
	var mutexWg sync.WaitGroup
	for i := 0; i < 100; i++ {
		mutexWg.Add(1)
		go func() {
			defer mutexWg.Done()
			for j := 0; j < iterations/100; j++ {
				atomic.AddInt64(&mutexCounter, 1)
			}
		}()
	}
	mutexWg.Wait()
	mutexTime := time.Since(start)
	
	// Atomic version
	start = time.Now()
	var atomicCounter2 int64
	var atomicWg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		atomicWg2.Add(1)
		go func() {
			defer atomicWg2.Done()
			for j := 0; j < iterations/100; j++ {
				atomic.AddInt64(&atomicCounter2, 1)
			}
		}()
	}
	atomicWg2.Wait()
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v\n", mutexTime)
	fmt.Printf("Atomic time: %v\n", atomicTime)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))

	// 11. Channel-based Synchronization
	fmt.Println("\n11. Channel-based Synchronization:")
	
	// Using channels for synchronization
	done := make(chan bool)
	
	go func() {
		// Do some work
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Work completed")
		done <- true
	}()
	
	// Wait for completion
	<-done
	fmt.Println("All work done")

	// 12. Semaphore Pattern
	fmt.Println("\n12. Semaphore Pattern:")
	
	// Implement semaphore using buffered channel
	semaphore := make(chan struct{}, 3) // Allow 3 concurrent operations
	
	// Worker function
	worker := func(id int) {
		semaphore <- struct{}{} // Acquire
		defer func() { <-semaphore }() // Release
		
		fmt.Printf("Worker %d started\n", id)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Worker %d finished\n", id)
	}
	
	// Start workers
	var semWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		semWg.Add(1)
		go func(i int) {
			defer semWg.Done()
			worker(i)
		}(i)
	}
	
	semWg.Wait()

	// 13. Pipeline Pattern
	fmt.Println("\n13. Pipeline Pattern:")
	
	// Stage 1: Generate numbers
	numbers := make(chan int, 10)
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
	}()
	
	// Stage 2: Square numbers
	squares := make(chan int, 10)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()
	
	// Stage 3: Print results
	for square := range squares {
		fmt.Printf("Square: %d\n", square)
	}

	// 14. Fan-out/Fan-in Pattern
	fmt.Println("\n14. Fan-out/Fan-in Pattern:")
	
	// Fan-out: Distribute work to multiple workers
	input := make(chan int, 10)
	output := make(chan int, 10)
	
	// Start workers
	var fanWg sync.WaitGroup
	for i := 0; i < 3; i++ {
		fanWg.Add(1)
		go func(workerID int) {
			defer fanWg.Done()
			for n := range input {
				result := n * n
				fmt.Printf("Worker %d processed %d -> %d\n", workerID, n, result)
				output <- result
			}
		}(i)
	}
	
	// Close output when all workers are done
	go func() {
		fanWg.Wait()
		close(output)
	}()
	
	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()
	
	// Collect results
	for result := range output {
		fmt.Printf("Final result: %d\n", result)
	}

	// 15. Advanced Synchronization
	fmt.Println("\n15. Advanced Synchronization:")
	
	// Test barrier
	barrier := NewBarrier(3)
	var barrierWg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		barrierWg.Add(1)
		go func(id int) {
			defer barrierWg.Done()
			fmt.Printf("Goroutine %d waiting at barrier\n", id)
			barrier.Wait()
			fmt.Printf("Goroutine %d passed barrier\n", id)
		}(i)
	}
	
	barrierWg.Wait()
	fmt.Println("All goroutines passed the barrier")

	// 16. Memory Model and Synchronization
	fmt.Println("\n16. Memory Model and Synchronization:")
	
	// Demonstrate memory visibility
	var (
		flag    int32
		message string
		mu2     sync.Mutex
	)
	
	// Writer
	go func() {
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		message = "Hello, World!"
		mu2.Unlock()
		atomic.StoreInt32(&flag, 1)
	}()
	
	// Reader
	for atomic.LoadInt32(&flag) == 0 {
		// Busy wait
	}
	
	mu2.Lock()
	fmt.Printf("Message: %s\n", message)
	mu2.Unlock()

	// 17. Goroutine Leak Prevention
	fmt.Println("\n17. Goroutine Leak Prevention:")
	
	// Use context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Worker with context
	workerFunc := func(ctx context.Context, id int) {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d cancelled\n", id)
				return
			case <-ticker.C:
				fmt.Printf("Worker %d working\n", id)
			}
		}
	}
	
	// Start workers
	var leakWg sync.WaitGroup
	for i := 0; i < 3; i++ {
		leakWg.Add(1)
		go func(i int) {
			defer leakWg.Done()
			workerFunc(ctx, i)
		}(i)
	}
	
	leakWg.Wait()

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

	fmt.Println("\n🎉 sync Package Mastery Complete!")
}
