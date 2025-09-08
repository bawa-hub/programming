package main

import (
	"fmt"
	"sync"
	"time"
)

// Connection represents a database connection
type Connection struct {
	ID        int
	CreatedAt time.Time
	InUse     bool
}

// String returns a string representation of the connection
func (c *Connection) String() string {
	return fmt.Sprintf("Connection{ID: %d, InUse: %v, Age: %v}", 
		c.ID, c.InUse, time.Since(c.CreatedAt))
}

// ConnectionPool manages a pool of database connections
type ConnectionPool struct {
	connections chan *Connection
	maxSize     int
	created     int
	mutex       sync.RWMutex
	closed      bool
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(maxSize int) *ConnectionPool {
	pool := &ConnectionPool{
		connections: make(chan *Connection, maxSize),
		maxSize:     maxSize,
	}
	
	// Pre-create some connections
	for i := 0; i < maxSize/2; i++ {
		conn := &Connection{
			ID:        i + 1,
			CreatedAt: time.Now(),
			InUse:     false,
		}
		pool.connections <- conn
		pool.created++
	}
	
	return pool
}

// GetConnection gets a connection from the pool
func (p *ConnectionPool) GetConnection() (*Connection, error) {
	p.mutex.RLock()
	if p.closed {
		p.mutex.RUnlock()
		return nil, fmt.Errorf("connection pool is closed")
	}
	p.mutex.RUnlock()
	
	select {
	case conn := <-p.connections:
		conn.InUse = true
		fmt.Printf("Got connection: %s\n", conn)
		return conn, nil
	default:
		// No available connections, create new one if under limit
		p.mutex.Lock()
		if p.created < p.maxSize {
			p.created++
			conn := &Connection{
				ID:        p.created,
				CreatedAt: time.Now(),
				InUse:     true,
			}
			p.mutex.Unlock()
			fmt.Printf("Created new connection: %s\n", conn)
			return conn, nil
		}
		p.mutex.Unlock()
		
		// Wait for a connection to become available
		select {
		case conn := <-p.connections:
			conn.InUse = true
			fmt.Printf("Got connection after wait: %s\n", conn)
			return conn, nil
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("timeout waiting for connection")
		}
	}
}

// ReturnConnection returns a connection to the pool
func (p *ConnectionPool) ReturnConnection(conn *Connection) error {
	p.mutex.RLock()
	if p.closed {
		p.mutex.RUnlock()
		return fmt.Errorf("connection pool is closed")
	}
	p.mutex.RUnlock()
	
	conn.InUse = false
	fmt.Printf("Returned connection: %s\n", conn)
	
	select {
	case p.connections <- conn:
		return nil
	default:
		// Pool is full, discard connection
		fmt.Printf("Pool full, discarding connection: %s\n", conn)
		return nil
	}
}

// Close closes the connection pool
func (p *ConnectionPool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	
	if p.closed {
		return
	}
	
	p.closed = true
	close(p.connections)
	
	// Close all connections
	for conn := range p.connections {
		fmt.Printf("Closing connection: %s\n", conn)
	}
}

// GetStats returns pool statistics
func (p *ConnectionPool) GetStats() (int, int, int) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	
	available := len(p.connections)
	return available, p.created, p.maxSize
}

// DemonstrateConnectionPool demonstrates the connection pool
func DemonstrateConnectionPool() {
	fmt.Println("=== Connection Pool Demonstration ===")
	
	// Create a pool with max 5 connections
	pool := NewConnectionPool(5)
	
	// Get some connections
	conn1, err := pool.GetConnection()
	if err != nil {
		fmt.Printf("Error getting connection: %v\n", err)
		return
	}
	
	conn2, err := pool.GetConnection()
	if err != nil {
		fmt.Printf("Error getting connection: %v\n", err)
		return
	}
	
	// Show stats
	available, created, max := pool.GetStats()
	fmt.Printf("Pool stats: available=%d, created=%d, max=%d\n", available, created, max)
	
	// Simulate work
	time.Sleep(1 * time.Second)
	
	// Return connections
	pool.ReturnConnection(conn1)
	pool.ReturnConnection(conn2)
	
	// Show stats again
	available, created, max = pool.GetStats()
	fmt.Printf("Pool stats after return: available=%d, created=%d, max=%d\n", available, created, max)
	
	// Close pool
	pool.Close()
}

// WorkerPool represents a pool of workers that use connections
type WorkerPool struct {
	pool       *ConnectionPool
	workers    int
	jobs       chan string
	results    chan string
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(pool *ConnectionPool, workers int) *WorkerPool {
	return &WorkerPool{
		pool:    pool,
		workers: workers,
		jobs:    make(chan string, 100),
		results: make(chan string, 100),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker processes jobs using connections from the pool
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for job := range wp.jobs {
		// Get connection from pool
		conn, err := wp.pool.GetConnection()
		if err != nil {
			wp.results <- fmt.Sprintf("Worker %d: Error getting connection for job %s: %v", id, job, err)
			continue
		}
		
		// Simulate work
		time.Sleep(500 * time.Millisecond)
		
		// Return connection to pool
		wp.pool.ReturnConnection(conn)
		
		// Send result
		wp.results <- fmt.Sprintf("Worker %d: Completed job %s using %s", id, job, conn)
	}
}

// SubmitJob submits a job to the worker pool
func (wp *WorkerPool) SubmitJob(job string) {
	wp.jobs <- job
}

// Close closes the worker pool
func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

// GetResults returns the results channel
func (wp *WorkerPool) GetResults() <-chan string {
	return wp.results
}

// DemonstrateWorkerPool demonstrates the worker pool with connection pool
func DemonstrateWorkerPool() {
	fmt.Println("\n=== Worker Pool with Connection Pool Demonstration ===")
	
	// Create connection pool
	pool := NewConnectionPool(3)
	defer pool.Close()
	
	// Create worker pool
	workerPool := NewWorkerPool(pool, 5)
	workerPool.Start()
	
	// Submit jobs
	jobs := []string{"job1", "job2", "job3", "job4", "job5", "job6", "job7", "job8"}
	for _, job := range jobs {
		workerPool.SubmitJob(job)
	}
	
	// Close worker pool
	workerPool.Close()
	
	// Collect results
	fmt.Println("Results:")
	for result := range workerPool.GetResults() {
		fmt.Printf("  %s\n", result)
	}
}
