package transaction

import (
	"database-engine/storage"
	"fmt"
	"sync"
	"time"
)

// TransactionManager manages database transactions
type TransactionManager struct {
	transactions map[TransactionID]*Transaction
	nextTxID     TransactionID
	mutex        sync.RWMutex
}

// Transaction represents a database transaction
type Transaction struct {
	ID        TransactionID
	State     TransactionState
	StartTime time.Time
	EndTime   time.Time
	Changes   []*Change
}

// TransactionID represents a transaction ID
type TransactionID int64

// TransactionState represents the state of a transaction
type TransactionState int

const (
	Active TransactionState = iota
	Committed
	Aborted
)

// Change represents a change in a transaction
type Change struct {
	Type      ChangeType
	TableName string
	Key       storage.Key
	OldRow    *storage.Row
	NewRow    *storage.Row
}

// ChangeType represents the type of change
type ChangeType int

const (
	Insert ChangeType = iota
	Update
	Delete
)

// NewTransactionManager creates a new transaction manager
func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		transactions: make(map[TransactionID]*Transaction),
		nextTxID:     1,
	}
}

// Initialize initializes the transaction manager
func (tm *TransactionManager) Initialize(se storage.StorageEngine) error {
	// Initialize transaction manager
	return nil
}

// BeginTransaction begins a new transaction
func (tm *TransactionManager) BeginTransaction() *Transaction {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	tx := &Transaction{
		ID:        tm.nextTxID,
		State:     Active,
		StartTime: time.Now(),
		Changes:   make([]*Change, 0),
	}
	
	tm.nextTxID++
	tm.transactions[tx.ID] = tx
	
	return tx
}

// CommitTransaction commits a transaction
func (tm *TransactionManager) CommitTransaction(tx *Transaction) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	// Check transaction state
	if tx.State != Active {
		return fmt.Errorf("transaction %d is not active", tx.ID)
	}
	
	// Mark transaction as committed
	tx.State = Committed
	tx.EndTime = time.Now()
	
	// Remove from active transactions
	delete(tm.transactions, tx.ID)
	
	return nil
}

// RollbackTransaction rolls back a transaction
func (tm *TransactionManager) RollbackTransaction(tx *Transaction) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	// Check transaction state
	if tx.State != Active {
		return fmt.Errorf("transaction %d is not active", tx.ID)
	}
	
	// Mark transaction as aborted
	tx.State = Aborted
	tx.EndTime = time.Now()
	
	// Remove from active transactions
	delete(tm.transactions, tx.ID)
	
	return nil
}

// GetTransaction gets a transaction by ID
func (tm *TransactionManager) GetTransaction(txID TransactionID) (*Transaction, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	tx, exists := tm.transactions[txID]
	if !exists {
		return nil, fmt.Errorf("transaction %d not found", txID)
	}
	
	return tx, nil
}

// ListTransactions lists all active transactions
func (tm *TransactionManager) ListTransactions() []*Transaction {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	transactions := make([]*Transaction, 0, len(tm.transactions))
	for _, tx := range tm.transactions {
		transactions = append(transactions, tx)
	}
	
	return transactions
}

// AddChange adds a change to a transaction
func (tm *TransactionManager) AddChange(tx *Transaction, change *Change) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	if tx.State == Active {
		tx.Changes = append(tx.Changes, change)
	}
}

// GetChanges gets all changes for a transaction
func (tm *TransactionManager) GetChanges(tx *Transaction) []*Change {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	return tx.Changes
}

// IsActive checks if a transaction is active
func (tm *TransactionManager) IsActive(tx *Transaction) bool {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	return tx.State == Active
}

// GetTransactionCount gets the number of active transactions
func (tm *TransactionManager) GetTransactionCount() int {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	return len(tm.transactions)
}

// CleanupInactiveTransactions cleans up inactive transactions
func (tm *TransactionManager) CleanupInactiveTransactions() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	now := time.Now()
	for txID, tx := range tm.transactions {
		// Clean up transactions that have been inactive for more than 1 hour
		if now.Sub(tx.StartTime) > time.Hour {
			delete(tm.transactions, txID)
		}
	}
}
