package main

import (
	"fmt"
	"sync"
	"time"
)

// 🔔 CONDITION VARIABLES MASTERY
// Understanding sync.Cond and condition variables in Go

func main() {
	fmt.Println("🔔 CONDITION VARIABLES MASTERY")
	fmt.Println("==============================")

	// 1. Basic Condition Variables
	fmt.Println("\n1. Basic Condition Variables:")
	basicConditionVariables()

	// 2. Producer-Consumer Pattern
	fmt.Println("\n2. Producer-Consumer Pattern:")
	producerConsumerPattern()

	// 3. Worker Pool with Condition Variables
	fmt.Println("\n3. Worker Pool with Condition Variables:")
	workerPoolWithConditionVariables()

	// 4. Barrier Pattern
	fmt.Println("\n4. Barrier Pattern:")
	barrierPattern()

	// 5. Semaphore Pattern
	fmt.Println("\n5. Semaphore Pattern:")
	semaphorePattern()

	// 6. Read-Write Lock with Condition Variables
	fmt.Println("\n6. Read-Write Lock with Condition Variables:")
	readWriteLockWithConditionVariables()

	// 7. Advanced Patterns
	fmt.Println("\n7. Advanced Patterns:")
	advancedPatterns()

	// 8. Best Practices
	fmt.Println("\n8. Best Practices:")
	bestPractices()
}

// BASIC CONDITION VARIABLES: Understanding basic condition variables
func basicConditionVariables() {
	fmt.Println("Understanding basic condition variables...")
	
	// Create condition variable
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	
	// Shared data
	sharedData := 0
	ready := false
	
	// Wait for condition
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		// Wait for condition
		for !ready {
			fmt.Println("  📊 Waiting for condition...")
			cond.Wait()
		}
		
		fmt.Printf("  📊 Condition met! Shared data: %d\n", sharedData)
	}()
	
	// Signal condition
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	sharedData = 42
	ready = true
	fmt.Println("  📊 Signaling condition...")
	cond.Signal()
	mu.Unlock()
	
	time.Sleep(100 * time.Millisecond)
}

// PRODUCER-CONSUMER PATTERN: Understanding producer-consumer with condition variables
func producerConsumerPattern() {
	fmt.Println("Understanding producer-consumer pattern...")
	
	// Create shared resources
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := make([]int, 0)
	maxSize := 5
	
	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			
			// Wait for space in queue
			for len(queue) >= maxSize {
				fmt.Printf("  📊 Producer waiting, queue full: %d\n", len(queue))
				cond.Wait()
			}
			
			// Add item to queue
			queue = append(queue, i)
			fmt.Printf("  📊 Produced: %d, queue size: %d\n", i, len(queue))
			
			// Notify consumers
			cond.Signal()
			mu.Unlock()
			
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Consumer
	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			
			// Wait for items in queue
			for len(queue) == 0 {
				fmt.Println("  📊 Consumer waiting, queue empty")
				cond.Wait()
			}
			
			// Remove item from queue
			item := queue[0]
			queue = queue[1:]
			fmt.Printf("  📊 Consumed: %d, queue size: %d\n", item, len(queue))
			
			// Notify producers
			cond.Signal()
			mu.Unlock()
			
			time.Sleep(75 * time.Millisecond)
		}
	}()
	
	time.Sleep(2 * time.Second)
}

// WORKER POOL WITH CONDITION VARIABLES: Understanding worker pool with condition variables
func workerPoolWithConditionVariables() {
	fmt.Println("Understanding worker pool with condition variables...")
	
	// Create worker pool
	pool := NewWorkerPool(3)
	
	// Add jobs
	for i := 0; i < 10; i++ {
		pool.AddJob(Job{
			ID:   i,
			Data: fmt.Sprintf("job-%d", i),
		})
	}
	
	// Start workers
	pool.Start()
	
	// Wait for completion
	pool.Wait()
	
	fmt.Printf("  📊 All jobs completed! Processed: %d\n", pool.ProcessedCount())
}

// BARRIER PATTERN: Understanding barrier pattern with condition variables
func barrierPattern() {
	fmt.Println("Understanding barrier pattern...")
	
	// Create barrier
	barrier := NewBarrier(3)
	
	// Create workers
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("  📊 Worker %d: Phase 1\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			
			// Wait at barrier
			barrier.Wait()
			
			fmt.Printf("  📊 Worker %d: Phase 2\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			
			// Wait at barrier again
			barrier.Wait()
			
			fmt.Printf("  📊 Worker %d: Phase 3\n", id)
		}(i)
	}
	
	time.Sleep(2 * time.Second)
}

// SEMAPHORE PATTERN: Understanding semaphore pattern with condition variables
func semaphorePattern() {
	fmt.Println("Understanding semaphore pattern...")
	
	// Create semaphore
	semaphore := NewSemaphore(2)
	
	// Create workers
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("  📊 Worker %d: Waiting for semaphore\n", id)
			semaphore.Acquire()
			
			fmt.Printf("  📊 Worker %d: Acquired semaphore\n", id)
			time.Sleep(500 * time.Millisecond)
			
			fmt.Printf("  📊 Worker %d: Releasing semaphore\n", id)
			semaphore.Release()
		}(i)
	}
	
	time.Sleep(3 * time.Second)
}

// READ-WRITE LOCK WITH CONDITION VARIABLES: Understanding read-write lock with condition variables
func readWriteLockWithConditionVariables() {
	fmt.Println("Understanding read-write lock with condition variables...")
	
	// Create read-write lock
	rwLock := NewReadWriteLock()
	
	// Create readers
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				rwLock.RLock()
				fmt.Printf("  📊 Reader %d: Reading data\n", id)
				time.Sleep(100 * time.Millisecond)
				rwLock.RUnlock()
			}
		}(i)
	}
	
	// Create writers
	for i := 0; i < 2; i++ {
		go func(id int) {
			for j := 0; j < 2; j++ {
				rwLock.Lock()
				fmt.Printf("  📊 Writer %d: Writing data\n", id)
				time.Sleep(200 * time.Millisecond)
				rwLock.Unlock()
			}
		}(i)
	}
	
	time.Sleep(2 * time.Second)
}

// ADVANCED PATTERNS: Understanding advanced condition variable patterns
func advancedPatterns() {
	fmt.Println("Understanding advanced patterns...")
	
	// Pattern 1: Timeout with condition variables
	fmt.Println("  📊 Pattern 1: Timeout with condition variables")
	timeoutPattern()
	
	// Pattern 2: Multiple condition variables
	fmt.Println("  📊 Pattern 2: Multiple condition variables")
	multipleConditionVariables()
	
	// Pattern 3: Condition variable with priority
	fmt.Println("  📊 Pattern 3: Condition variable with priority")
	priorityPattern()
}

func timeoutPattern() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	// Wait with timeout
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		// Wait for condition with timeout
		timeout := time.After(500 * time.Millisecond)
		for !ready {
			select {
			case <-timeout:
				fmt.Println("    ⏰ Timeout waiting for condition")
				return
			default:
				cond.Wait()
			}
		}
		
		fmt.Println("    ✅ Condition met before timeout")
	}()
	
	// Signal after timeout
	time.Sleep(600 * time.Millisecond)
	
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}

func multipleConditionVariables() {
	var mu sync.Mutex
	cond1 := sync.NewCond(&mu)
	cond2 := sync.NewCond(&mu)
	
	ready1 := false
	ready2 := false
	
	// Wait for both conditions
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		for !ready1 || !ready2 {
			if !ready1 {
				cond1.Wait()
			}
			if !ready2 {
				cond2.Wait()
			}
		}
		
		fmt.Println("    ✅ Both conditions met")
	}()
	
	// Signal conditions
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready1 = true
	cond1.Signal()
	mu.Unlock()
	
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready2 = true
	cond2.Signal()
	mu.Unlock()
}

func priorityPattern() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := make([]int, 0)
	processing := false
	
	// High priority worker
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		for {
			for len(queue) == 0 || processing {
				cond.Wait()
			}
			
			processing = true
			item := queue[0]
			queue = queue[1:]
			fmt.Printf("    📊 High priority processing: %d\n", item)
			time.Sleep(100 * time.Millisecond)
			processing = false
			cond.Signal()
		}
	}()
	
	// Low priority worker
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		for {
			for len(queue) == 0 || processing {
				cond.Wait()
			}
			
			processing = true
			item := queue[0]
			queue = queue[1:]
			fmt.Printf("    📊 Low priority processing: %d\n", item)
			time.Sleep(100 * time.Millisecond)
			processing = false
			cond.Signal()
		}
	}()
	
	// Add items
	mu.Lock()
	queue = append(queue, 1, 2, 3, 4, 5)
	cond.Broadcast()
	mu.Unlock()
	
	time.Sleep(1 * time.Second)
}

// BEST PRACTICES: Understanding best practices
func bestPractices() {
	fmt.Println("Understanding best practices...")
	
	// 1. Always use condition variables with mutex
	fmt.Println("  📝 Best Practice 1: Always use condition variables with mutex")
	useWithMutex()
	
	// 2. Use Broadcast() for multiple waiters
	fmt.Println("  📝 Best Practice 2: Use Broadcast() for multiple waiters")
	useBroadcast()
	
	// 3. Check condition in loop
	fmt.Println("  📝 Best Practice 3: Check condition in loop")
	checkConditionInLoop()
	
	// 4. Avoid spurious wakeups
	fmt.Println("  📝 Best Practice 4: Avoid spurious wakeups")
	avoidSpuriousWakeups()
}

func useWithMutex() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		for !ready {
			cond.Wait()
		}
		
		fmt.Println("    ✅ Condition met with mutex")
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}

func useBroadcast() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	// Multiple waiters
	for i := 0; i < 3; i++ {
		go func(id int) {
			mu.Lock()
			defer mu.Unlock()
			
			for !ready {
				cond.Wait()
			}
			
			fmt.Printf("    ✅ Waiter %d notified\n", id)
		}(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready = true
	cond.Broadcast() // Notify all waiters
	mu.Unlock()
	
	time.Sleep(100 * time.Millisecond)
}

func checkConditionInLoop() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		// Always check condition in loop
		for !ready {
			cond.Wait()
		}
		
		fmt.Println("    ✅ Condition checked in loop")
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}

func avoidSpuriousWakeups() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		// Check condition in loop to avoid spurious wakeups
		for !ready {
			cond.Wait()
		}
		
		fmt.Println("    ✅ Avoided spurious wakeup")
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}

// IMPLEMENTATIONS

// Worker Pool
type WorkerPool struct {
	workers    int
	jobs       chan Job
	done       chan bool
	mu         sync.Mutex
	cond       *sync.Cond
	processed  int
	completed  bool
}

type Job struct {
	ID   int
	Data string
}

func NewWorkerPool(workers int) *WorkerPool {
	pool := &WorkerPool{
		workers: workers,
		jobs:    make(chan Job, 100),
		done:    make(chan bool, workers),
	}
	pool.cond = sync.NewCond(&pool.mu)
	return pool
}

func (p *WorkerPool) AddJob(job Job) {
	p.jobs <- job
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.workers; i++ {
		go p.worker(i)
	}
}

func (p *WorkerPool) worker(id int) {
	for job := range p.jobs {
		fmt.Printf("  📊 Worker %d processing job %d: %s\n", id, job.ID, job.Data)
		time.Sleep(100 * time.Millisecond)
		
		p.mu.Lock()
		p.processed++
		p.cond.Signal()
		p.mu.Unlock()
	}
	p.done <- true
}

func (p *WorkerPool) Wait() {
	close(p.jobs)
	
	// Wait for all workers to finish
	for i := 0; i < p.workers; i++ {
		<-p.done
	}
}

func (p *WorkerPool) ProcessedCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.processed
}

// Barrier
type Barrier struct {
	count    int
	waiting  int
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewBarrier(count int) *Barrier {
	barrier := &Barrier{count: count}
	barrier.cond = sync.NewCond(&barrier.mu)
	return barrier
}

func (b *Barrier) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.waiting++
	
	if b.waiting == b.count {
		b.waiting = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
}

// Semaphore
type Semaphore struct {
	permits int
	mu      sync.Mutex
	cond    *sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
	sem := &Semaphore{permits: permits}
	sem.cond = sync.NewCond(&sem.mu)
	return sem
}

func (s *Semaphore) Acquire() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	for s.permits == 0 {
		s.cond.Wait()
	}
	
	s.permits--
}

func (s *Semaphore) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.permits++
	s.cond.Signal()
}

// Read-Write Lock
type ReadWriteLock struct {
	readers  int
	writing  bool
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewReadWriteLock() *ReadWriteLock {
	rw := &ReadWriteLock{}
	rw.cond = sync.NewCond(&rw.mu)
	return rw
}

func (rw *ReadWriteLock) RLock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	
	for rw.writing {
		rw.cond.Wait()
	}
	
	rw.readers++
}

func (rw *ReadWriteLock) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	
	rw.readers--
	if rw.readers == 0 {
		rw.cond.Signal()
	}
}

func (rw *ReadWriteLock) Lock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	
	for rw.readers > 0 || rw.writing {
		rw.cond.Wait()
	}
	
	rw.writing = true
}

func (rw *ReadWriteLock) Unlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	
	rw.writing = false
	rw.cond.Broadcast()
}
