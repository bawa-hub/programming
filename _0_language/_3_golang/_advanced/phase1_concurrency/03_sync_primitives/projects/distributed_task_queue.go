package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 🚀 DISTRIBUTED TASK QUEUE PROJECT
// A production-ready task queue with worker pools and synchronization

type Task struct {
	ID        string
	Data      interface{}
	Priority  int
	CreatedAt time.Time
	Status    TaskStatus
	Result    interface{}
	Error     error
}

type TaskStatus int

const (
	TaskPending TaskStatus = iota
	TaskProcessing
	TaskCompleted
	TaskFailed
)

type TaskQueue struct {
	tasks      []Task
	mu         sync.RWMutex
	cond       *sync.Cond
	shutdown   chan struct{}
	wg         sync.WaitGroup
	workers    int
	results    chan Task
}

type Worker struct {
	ID       int
	Queue    *TaskQueue
	Shutdown chan struct{}
}

func main() {
	fmt.Println("🚀 DISTRIBUTED TASK QUEUE")
	fmt.Println("==========================")

	// Create task queue with 3 workers
	queue := NewTaskQueue(3)
	defer queue.Close()

	// Start workers
	queue.StartWorkers()

	// Add some tasks
	fmt.Println("📝 Adding tasks...")
	
	// High priority tasks
	queue.AddTask(Task{
		ID:        "task-1",
		Data:      "High priority task 1",
		Priority:  10,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})
	
	queue.AddTask(Task{
		ID:        "task-2",
		Data:      "High priority task 2",
		Priority:  10,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})
	
	// Medium priority tasks
	queue.AddTask(Task{
		ID:        "task-3",
		Data:      "Medium priority task 1",
		Priority:  5,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})
	
	queue.AddTask(Task{
		ID:        "task-4",
		Data:      "Medium priority task 2",
		Priority:  5,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})
	
	// Low priority tasks
	queue.AddTask(Task{
		ID:        "task-5",
		Data:      "Low priority task 1",
		Priority:  1,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})
	
	queue.AddTask(Task{
		ID:        "task-6",
		Data:      "Low priority task 2",
		Priority:  1,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})

	// Add some tasks that will fail
	queue.AddTask(Task{
		ID:        "task-fail-1",
		Data:      "This task will fail",
		Priority:  5,
		CreatedAt: time.Now(),
		Status:    TaskPending,
	})

	// Collect results
	fmt.Println("📊 Collecting results...")
	results := queue.CollectResults()
	
	// Display results
	fmt.Println("\n📈 TASK EXECUTION RESULTS")
	fmt.Println("=========================")
	
	successCount := 0
	failCount := 0
	
	for _, task := range results {
		switch task.Status {
		case TaskCompleted:
			fmt.Printf("✅ %s: %v (Priority: %d)\n", task.ID, task.Result, task.Priority)
			successCount++
		case TaskFailed:
			fmt.Printf("❌ %s: %v (Priority: %d)\n", task.ID, task.Error, task.Priority)
			failCount++
		default:
			fmt.Printf("⏳ %s: %s (Priority: %d)\n", task.ID, task.Status, task.Priority)
		}
	}
	
	fmt.Printf("\n📊 SUMMARY:\n")
	fmt.Printf("  Total tasks: %d\n", len(results))
	fmt.Printf("  Completed: %d\n", successCount)
	fmt.Printf("  Failed: %d\n", failCount)
	fmt.Printf("  Success rate: %.2f%%\n", float64(successCount)/float64(len(results))*100)
}

// NewTaskQueue creates a new task queue
func NewTaskQueue(workers int) *TaskQueue {
	queue := &TaskQueue{
		workers:  workers,
		results:  make(chan Task, 100),
		shutdown: make(chan struct{}),
	}
	queue.cond = sync.NewCond(&queue.mu)
	return queue
}

// AddTask adds a task to the queue
func (q *TaskQueue) AddTask(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	// Insert task in priority order
	q.insertTaskByPriority(task)
	
	fmt.Printf("  📝 Added task %s (Priority: %d)\n", task.ID, task.Priority)
	
	// Notify workers that a new task is available
	q.cond.Signal()
}

// insertTaskByPriority inserts a task in priority order
func (q *TaskQueue) insertTaskByPriority(task Task) {
	// Find insertion point (higher priority first)
	insertIndex := len(q.tasks)
	for i, t := range q.tasks {
		if task.Priority > t.Priority {
			insertIndex = i
			break
		}
	}
	
	// Insert at the found position
	q.tasks = append(q.tasks[:insertIndex], append([]Task{task}, q.tasks[insertIndex:]...)...)
}

// GetNextTask gets the next task from the queue
func (q *TaskQueue) GetNextTask() *Task {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	// Wait for a task to be available
	for len(q.tasks) == 0 {
		select {
		case <-q.shutdown:
			return nil
		default:
			q.cond.Wait()
		}
	}
	
	// Get the highest priority task
	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	
	// Update task status
	task.Status = TaskProcessing
	return &task
}

// CompleteTask marks a task as completed
func (q *TaskQueue) CompleteTask(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	task.Status = TaskCompleted
	q.results <- task
}

// FailTask marks a task as failed
func (q *TaskQueue) FailTask(task Task, err error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	task.Status = TaskFailed
	task.Error = err
	q.results <- task
}

// StartWorkers starts the worker goroutines
func (q *TaskQueue) StartWorkers() {
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		worker := &Worker{
			ID:       i,
			Queue:    q,
			Shutdown: make(chan struct{}),
		}
		go worker.Start()
	}
}

// Start starts a worker
func (w *Worker) Start() {
	defer w.Queue.wg.Done()
	
	fmt.Printf("  🧵 Worker %d started\n", w.ID)
	
	for {
		select {
		case <-w.Shutdown:
			fmt.Printf("  🧵 Worker %d shutting down\n", w.ID)
			return
		default:
			// Get next task
			task := w.Queue.GetNextTask()
			if task == nil {
				// Shutdown signal received
				return
			}
			
			// Process task
			w.processTask(*task)
		}
	}
}

// processTask processes a single task
func (w *Worker) processTask(task Task) {
	fmt.Printf("  🧵 Worker %d processing task %s (Priority: %d)\n", 
		w.ID, task.ID, task.Priority)
	
	// Simulate work
	workDuration := time.Duration(rand.Intn(1000)+500) * time.Millisecond
	time.Sleep(workDuration)
	
	// Simulate task failure for specific tasks
	if task.ID == "task-fail-1" {
		w.Queue.FailTask(task, fmt.Errorf("simulated failure"))
		fmt.Printf("  ❌ Worker %d failed task %s\n", w.ID, task.ID)
		return
	}
	
	// Complete task
	task.Result = fmt.Sprintf("Processed by worker %d in %v", w.ID, workDuration)
	w.Queue.CompleteTask(task)
	fmt.Printf("  ✅ Worker %d completed task %s\n", w.ID, task.ID)
}

// CollectResults collects all task results
func (q *TaskQueue) CollectResults() []Task {
	var results []Task
	
	// Wait for all tasks to be processed
	time.Sleep(2 * time.Second)
	
	// Collect results from channel
	for {
		select {
		case task := <-q.results:
			results = append(results, task)
		default:
			// No more results available
			return results
		}
	}
}

// Close gracefully shuts down the task queue
func (q *TaskQueue) Close() {
	close(q.shutdown)
	q.cond.Broadcast() // Wake up all waiting workers
	q.wg.Wait()
	close(q.results)
}
