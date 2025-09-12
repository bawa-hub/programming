package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// ConcurrencyManager manages database concurrency
type ConcurrencyManager struct {
	locks      map[string]*Lock
	waitQueue  map[string][]*LockRequest
	mutex      sync.RWMutex
	deadlockDetector *DeadlockDetector
}

// Lock represents a database lock
type Lock struct {
	Resource  string
	Type      LockType
	Holder    string
	Waiters   []*LockRequest
	CreatedAt time.Time
}

// LockType represents the type of lock
type LockType int

const (
	SharedLock LockType = iota
	ExclusiveLock
	IntentSharedLock
	IntentExclusiveLock
)

// LockRequest represents a lock request
type LockRequest struct {
	Requester string
	Type      LockType
	Timestamp time.Time
	Granted   bool
}

// DeadlockDetector detects deadlocks
type DeadlockDetector struct {
	waitForGraph map[string][]string
	mutex        sync.RWMutex
}

// NewConcurrencyManager creates a new concurrency manager
func NewConcurrencyManager() *ConcurrencyManager {
	return &ConcurrencyManager{
		locks:      make(map[string]*Lock),
		waitQueue:  make(map[string][]*LockRequest),
		deadlockDetector: NewDeadlockDetector(),
	}
}

// Initialize initializes the concurrency manager
func (cm *ConcurrencyManager) Initialize() error {
	// Initialize concurrency manager
	return nil
}

// AcquireLock acquires a lock
func (cm *ConcurrencyManager) AcquireLock(resource, requester string, lockType LockType) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	
	// Check if lock already exists
	if lock, exists := cm.locks[resource]; exists {
		// Check if requester already holds the lock
		if lock.Holder == requester {
			// Upgrade lock if needed
			if lockType > lock.Type {
				lock.Type = lockType
			}
			return nil
		}
		
		// Check if lock can be granted
		if cm.canGrantLock(lock, lockType) {
			// Grant lock
			lock.Holder = requester
			lock.Type = lockType
			lock.CreatedAt = time.Now()
			return nil
		}
		
		// Add to wait queue
		request := &LockRequest{
			Requester: requester,
			Type:      lockType,
			Timestamp: time.Now(),
			Granted:   false,
		}
		lock.Waiters = append(lock.Waiters, request)
		cm.waitQueue[resource] = append(cm.waitQueue[resource], request)
		
		// Check for deadlock
		if cm.deadlockDetector.DetectDeadlock(requester, lock.Holder) {
			// Resolve deadlock
			cm.resolveDeadlock(resource, requester)
		}
		
		return fmt.Errorf("lock not available, added to wait queue")
	}
	
	// Create new lock
	lock := &Lock{
		Resource:  resource,
		Type:      lockType,
		Holder:    requester,
		Waiters:   make([]*LockRequest, 0),
		CreatedAt: time.Now(),
	}
	cm.locks[resource] = lock
	
	return nil
}

// ReleaseLock releases a lock
func (cm *ConcurrencyManager) ReleaseLock(resource, holder string) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	
	// Check if lock exists
	lock, exists := cm.locks[resource]
	if !exists {
		return fmt.Errorf("lock for resource %s does not exist", resource)
	}
	
	// Check if holder owns the lock
	if lock.Holder != holder {
		return fmt.Errorf("lock holder %s does not match requester %s", lock.Holder, holder)
	}
	
	// Grant lock to next waiter
	if len(lock.Waiters) > 0 {
		nextWaiter := lock.Waiters[0]
		lock.Waiters = lock.Waiters[1:]
		lock.Holder = nextWaiter.Requester
		lock.Type = nextWaiter.Type
		lock.CreatedAt = time.Now()
		nextWaiter.Granted = true
		
		// Remove from wait queue
		cm.removeFromWaitQueue(resource, nextWaiter)
	} else {
		// No waiters, remove lock
		delete(cm.locks, resource)
	}
	
	return nil
}

// canGrantLock checks if a lock can be granted
func (cm *ConcurrencyManager) canGrantLock(lock *Lock, requestedType LockType) bool {
	// Shared locks can be granted if no exclusive lock is held
	if requestedType == SharedLock {
		return lock.Type != ExclusiveLock
	}
	
	// Exclusive locks can only be granted if no other locks are held
	if requestedType == ExclusiveLock {
		return len(lock.Waiters) == 0
	}
	
	return false
}

// removeFromWaitQueue removes a request from the wait queue
func (cm *ConcurrencyManager) removeFromWaitQueue(resource string, request *LockRequest) {
	if queue, exists := cm.waitQueue[resource]; exists {
		for i, req := range queue {
			if req == request {
				cm.waitQueue[resource] = append(queue[:i], queue[i+1:]...)
				break
			}
		}
	}
}

// resolveDeadlock resolves a deadlock
func (cm *ConcurrencyManager) resolveDeadlock(resource, requester string) {
	// Simple deadlock resolution - abort the requester
	// In a real implementation, this would be more sophisticated
	
	// Find and abort the requester's transaction
	// This is a placeholder implementation
	fmt.Printf("Deadlock detected for resource %s, requester %s\n", resource, requester)
}

// GetLockInfo gets information about a lock
func (cm *ConcurrencyManager) GetLockInfo(resource string) (*LockInfo, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	
	lock, exists := cm.locks[resource]
	if !exists {
		return nil, fmt.Errorf("lock for resource %s does not exist", resource)
	}
	
	return &LockInfo{
		Resource:  lock.Resource,
		Type:      lock.Type.String(),
		Holder:    lock.Holder,
		Waiters:   len(lock.Waiters),
		CreatedAt: lock.CreatedAt,
	}, nil
}

// ListLocks lists all locks
func (cm *ConcurrencyManager) ListLocks() []*LockInfo {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	
	locks := make([]*LockInfo, 0, len(cm.locks))
	for _, lock := range cm.locks {
		locks = append(locks, &LockInfo{
			Resource:  lock.Resource,
			Type:      lock.Type.String(),
			Holder:    lock.Holder,
			Waiters:   len(lock.Waiters),
			CreatedAt: lock.CreatedAt,
		})
	}
	
	return locks
}

// GetLockCount gets the number of active locks
func (cm *ConcurrencyManager) GetLockCount() int {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	
	return len(cm.locks)
}

// LockInfo represents lock information
type LockInfo struct {
	Resource  string    `json:"resource"`
	Type      string    `json:"type"`
	Holder    string    `json:"holder"`
	Waiters   int       `json:"waiters"`
	CreatedAt time.Time `json:"created_at"`
}

// String method for LockType
func (lt LockType) String() string {
	switch lt {
	case SharedLock:
		return "SHARED"
	case ExclusiveLock:
		return "EXCLUSIVE"
	case IntentSharedLock:
		return "INTENT_SHARED"
	case IntentExclusiveLock:
		return "INTENT_EXCLUSIVE"
	default:
		return "UNKNOWN"
	}
}

// NewDeadlockDetector creates a new deadlock detector
func NewDeadlockDetector() *DeadlockDetector {
	return &DeadlockDetector{
		waitForGraph: make(map[string][]string),
	}
}

// DetectDeadlock detects if there's a deadlock
func (dd *DeadlockDetector) DetectDeadlock(requester, holder string) bool {
	dd.mutex.Lock()
	defer dd.mutex.Unlock()
	
	// Add edge to wait-for graph
	dd.waitForGraph[requester] = append(dd.waitForGraph[requester], holder)
	
	// Check for cycle using DFS
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	
	return dd.hasCycle(requester, visited, recStack)
}

// hasCycle checks if there's a cycle in the wait-for graph
func (dd *DeadlockDetector) hasCycle(node string, visited, recStack map[string]bool) bool {
	visited[node] = true
	recStack[node] = true
	
	for _, neighbor := range dd.waitForGraph[node] {
		if !visited[neighbor] {
			if dd.hasCycle(neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true
		}
	}
	
	recStack[node] = false
	return false
}

// ClearWaitForGraph clears the wait-for graph
func (dd *DeadlockDetector) ClearWaitForGraph() {
	dd.mutex.Lock()
	defer dd.mutex.Unlock()
	
	dd.waitForGraph = make(map[string][]string)
}
