# 🗄️ Level 5, Topic 2: Database Concurrency

## 🚀 Overview
Welcome to **Database Concurrency**! This topic covers the critical aspects of building concurrent systems that interact with databases safely and efficiently. We'll explore connection pooling, transaction management, read replicas, ACID properties, isolation levels, and deadlock prevention.

---

## 📚 Table of Contents

1. [Database Concurrency Fundamentals](#database-concurrency-fundamentals)
2. [Connection Pooling](#connection-pooling)
3. [Transaction Management](#transaction-management)
4. [ACID Properties](#acid-properties)
5. [Isolation Levels](#isolation-levels)
6. [Deadlock Prevention](#deadlock-prevention)
7. [Read Replicas](#read-replicas)
8. [Database Sharding](#database-sharding)
9. [Optimistic vs Pessimistic Locking](#optimistic-vs-pessimistic-locking)
10. [Database Connection Management](#database-connection-management)
11. [Concurrent Query Processing](#concurrent-query-processing)
12. [Database Monitoring](#database-monitoring)
13. [Performance Optimization](#performance-optimization)
14. [Error Handling](#error-handling)
15. [Testing Strategies](#testing-strategies)
16. [Best Practices](#best-practices)

---

## 🗄️ Database Concurrency Fundamentals

### Understanding Database Concurrency

Database concurrency refers to the ability of a database system to handle multiple simultaneous operations while maintaining data consistency and integrity. In concurrent systems, multiple goroutines may access the same database simultaneously, requiring careful coordination.

### Key Challenges

#### 1. Data Consistency
- Ensuring data remains consistent across concurrent operations
- Preventing partial updates and inconsistent states
- Maintaining referential integrity

#### 2. Performance
- Minimizing lock contention
- Optimizing query execution
- Efficient connection management

#### 3. Scalability
- Handling increasing concurrent load
- Distributing database operations
- Managing connection pools

### Database Concurrency Patterns

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Database Connection
type DatabaseManager struct {
	db     *sql.DB
	config *DatabaseConfig
	mutex  sync.RWMutex
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	MaxConns int
	MaxIdle  int
}

func NewDatabaseManager(config *DatabaseConfig) (*DatabaseManager, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	// Configure connection pool
	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.MaxIdle)
	db.SetConnMaxLifetime(time.Hour)
	
	return &DatabaseManager{
		db:     db,
		config: config,
	}, nil
}

func (dm *DatabaseManager) Close() error {
	return dm.db.Close()
}

// Example 2: Concurrent Database Operations
func (dm *DatabaseManager) ExecuteQuery(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	return dm.db.QueryContext(ctx, query, args...)
}

func (dm *DatabaseManager) ExecuteTransaction(ctx context.Context, fn func(*sql.Tx) error) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	tx, err := dm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()
	
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit()
}
```

---

## 🏊 Connection Pooling

### Understanding Connection Pooling

Connection pooling is a technique used to maintain a cache of database connections that can be reused across multiple requests. This reduces the overhead of creating and destroying connections.

### Connection Pool Benefits

#### 1. Performance
- Reuse existing connections
- Reduce connection establishment overhead
- Better resource utilization

#### 2. Scalability
- Handle more concurrent requests
- Control resource usage
- Prevent connection exhaustion

#### 3. Reliability
- Connection health monitoring
- Automatic reconnection
- Graceful degradation

### Connection Pool Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Connection Pool
type ConnectionPool struct {
	db          *sql.DB
	maxConns    int
	maxIdle     int
	maxLifetime time.Duration
	stats       *PoolStats
	mutex       sync.RWMutex
}

type PoolStats struct {
	OpenConns     int
	InUse         int
	Idle          int
	WaitCount     int64
	WaitDuration  time.Duration
	MaxIdleClosed int64
	MaxLifetimeClosed int64
}

func NewConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*ConnectionPool, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	// Configure connection pool
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(maxLifetime)
	
	pool := &ConnectionPool{
		db:          db,
		maxConns:    maxConns,
		maxIdle:     maxIdle,
		maxLifetime: maxLifetime,
		stats:       &PoolStats{},
	}
	
	// Start stats monitoring
	go pool.monitorStats()
	
	return pool, nil
}

func (cp *ConnectionPool) GetConnection(ctx context.Context) (*sql.Conn, error) {
	conn, err := cp.db.Conn(ctx)
	if err != nil {
		cp.mutex.Lock()
		cp.stats.WaitCount++
		cp.mutex.Unlock()
		return nil, err
	}
	
	cp.mutex.Lock()
	cp.stats.OpenConns++
	cp.stats.InUse++
	cp.mutex.Unlock()
	
	return conn, nil
}

func (cp *ConnectionPool) ReleaseConnection(conn *sql.Conn) {
	cp.mutex.Lock()
	cp.stats.InUse--
	cp.stats.Idle++
	cp.mutex.Unlock()
	
	conn.Close()
}

func (cp *ConnectionPool) GetStats() *PoolStats {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()
	
	stats := *cp.stats
	stats.OpenConns = cp.db.Stats().OpenConnections
	stats.Idle = cp.db.Stats().Idle
	stats.InUse = stats.OpenConns - stats.Idle
	
	return &stats
}

func (cp *ConnectionPool) monitorStats() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		cp.mutex.Lock()
		cp.stats.MaxIdleClosed = cp.db.Stats().MaxIdleClosed
		cp.stats.MaxLifetimeClosed = cp.db.Stats().MaxLifetimeClosed
		cp.mutex.Unlock()
	}
}

// Example 2: Advanced Connection Pool with Health Checks
type AdvancedConnectionPool struct {
	*ConnectionPool
	healthChecker *HealthChecker
	reconnectCh   chan struct{}
}

type HealthChecker struct {
	interval time.Duration
	timeout  time.Duration
	stopCh   chan struct{}
}

func NewAdvancedConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*AdvancedConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	acp := &AdvancedConnectionPool{
		ConnectionPool: pool,
		healthChecker: &HealthChecker{
			interval: 30 * time.Second,
			timeout:  5 * time.Second,
			stopCh:   make(chan struct{}),
		},
		reconnectCh: make(chan struct{}, 1),
	}
	
	// Start health checking
	go acp.healthChecker.Start(acp.db)
	
	return acp, nil
}

func (hc *HealthChecker) Start(db *sql.DB) {
	ticker := time.NewTicker(hc.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), hc.timeout)
			if err := db.PingContext(ctx); err != nil {
				fmt.Printf("Database health check failed: %v\n", err)
			}
			cancel()
		case <-hc.stopCh:
			return
		}
	}
}

func (acp *AdvancedConnectionPool) GetConnectionWithRetry(ctx context.Context, maxRetries int) (*sql.Conn, error) {
	for i := 0; i < maxRetries; i++ {
		conn, err := acp.GetConnection(ctx)
		if err == nil {
			return conn, nil
		}
		
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}
	
	return nil, fmt.Errorf("failed to get connection after %d retries", maxRetries)
}
```

---

## 🔄 Transaction Management

### Understanding Transactions

A transaction is a sequence of database operations that are treated as a single unit of work. Transactions ensure that either all operations succeed (commit) or all operations fail (rollback).

### Transaction Properties

#### 1. Atomicity
- All operations in a transaction succeed or all fail
- No partial updates

#### 2. Consistency
- Database remains in a valid state
- All constraints are maintained

#### 3. Isolation
- Concurrent transactions don't interfere with each other
- Different isolation levels provide different guarantees

#### 4. Durability
- Committed changes persist even after system failure
- Data is written to persistent storage

### Transaction Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Transaction Management
type TransactionManager struct {
	db     *sql.DB
	active map[string]*sql.Tx
	mutex  sync.RWMutex
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{
		db:     db,
		active: make(map[string]*sql.Tx),
	}
}

func (tm *TransactionManager) Begin(ctx context.Context, txID string) (*sql.Tx, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	
	tm.active[txID] = tx
	return tx, nil
}

func (tm *TransactionManager) Commit(txID string) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	tx, exists := tm.active[txID]
	if !exists {
		return fmt.Errorf("transaction %s not found", txID)
	}
	
	err := tx.Commit()
	delete(tm.active, txID)
	return err
}

func (tm *TransactionManager) Rollback(txID string) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	tx, exists := tm.active[txID]
	if !exists {
		return fmt.Errorf("transaction %s not found", txID)
	}
	
	err := tx.Rollback()
	delete(tm.active, txID)
	return err
}

// Example 2: Nested Transaction Support
type NestedTransactionManager struct {
	*TransactionManager
	nested map[string][]string
}

func NewNestedTransactionManager(db *sql.DB) *NestedTransactionManager {
	return &NestedTransactionManager{
		TransactionManager: NewTransactionManager(db),
		nested:             make(map[string][]string),
	}
}

func (ntm *NestedTransactionManager) BeginNested(ctx context.Context, parentTxID, childTxID string) (*sql.Tx, error) {
	ntm.mutex.Lock()
	defer ntm.mutex.Unlock()
	
	// Check if parent transaction exists
	if _, exists := ntm.active[parentTxID]; !exists {
		return nil, fmt.Errorf("parent transaction %s not found", parentTxID)
	}
	
	// Create savepoint
	tx := ntm.active[parentTxID]
	_, err := tx.ExecContext(ctx, "SAVEPOINT "+childTxID)
	if err != nil {
		return nil, err
	}
	
	// Track nested transaction
	ntm.nested[parentTxID] = append(ntm.nested[parentTxID], childTxID)
	
	return tx, nil
}

func (ntm *NestedTransactionManager) RollbackNested(parentTxID, childTxID string) error {
	ntm.mutex.Lock()
	defer ntm.mutex.Unlock()
	
	tx, exists := ntm.active[parentTxID]
	if !exists {
		return fmt.Errorf("parent transaction %s not found", parentTxID)
	}
	
	// Rollback to savepoint
	_, err := tx.Exec("ROLLBACK TO SAVEPOINT " + childTxID)
	if err != nil {
		return err
	}
	
	// Remove from nested list
	nested := ntm.nested[parentTxID]
	for i, id := range nested {
		if id == childTxID {
			ntm.nested[parentTxID] = append(nested[:i], nested[i+1:]...)
			break
		}
	}
	
	return nil
}

// Example 3: Transaction with Retry Logic
type RetryableTransactionManager struct {
	*TransactionManager
	maxRetries int
	retryDelay time.Duration
}

func NewRetryableTransactionManager(db *sql.DB, maxRetries int, retryDelay time.Duration) *RetryableTransactionManager {
	return &RetryableTransactionManager{
		TransactionManager: NewTransactionManager(db),
		maxRetries:         maxRetries,
		retryDelay:         retryDelay,
	}
}

func (rtm *RetryableTransactionManager) ExecuteWithRetry(ctx context.Context, fn func(*sql.Tx) error) error {
	var lastErr error
	
	for i := 0; i < rtm.maxRetries; i++ {
		tx, err := rtm.db.BeginTx(ctx, nil)
		if err != nil {
			lastErr = err
			continue
		}
		
		if err := fn(tx); err != nil {
			tx.Rollback()
			lastErr = err
			
			// Check if error is retryable
			if !rtm.isRetryableError(err) {
				break
			}
			
			if i < rtm.maxRetries-1 {
				time.Sleep(rtm.retryDelay)
			}
			continue
		}
		
		if err := tx.Commit(); err != nil {
			lastErr = err
			if !rtm.isRetryableError(err) {
				break
			}
			
			if i < rtm.maxRetries-1 {
				time.Sleep(rtm.retryDelay)
			}
			continue
		}
		
		return nil
	}
	
	return fmt.Errorf("transaction failed after %d retries: %v", rtm.maxRetries, lastErr)
}

func (rtm *RetryableTransactionManager) isRetryableError(err error) bool {
	// Check for retryable database errors
	// This is a simplified implementation
	return err != nil
}
```

---

## 🔒 ACID Properties

### Understanding ACID

ACID (Atomicity, Consistency, Isolation, Durability) is a set of properties that guarantee reliable processing of database transactions.

### ACID Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
)

// Example 1: Atomicity - All or Nothing
type AtomicOperation struct {
	db *sql.DB
}

func (ao *AtomicOperation) TransferMoney(ctx context.Context, fromAccount, toAccount string, amount float64) error {
	tx, err := ao.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Check if source account has sufficient funds
	var balance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = ?", fromAccount).Scan(&balance)
	if err != nil {
		return err
	}
	
	if balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	
	// Debit from source account
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccount)
	if err != nil {
		return err
	}
	
	// Credit to destination account
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccount)
	if err != nil {
		return err
	}
	
	// Commit transaction (atomicity)
	return tx.Commit()
}

// Example 2: Consistency - Data Integrity
type ConsistencyManager struct {
	db     *sql.DB
	checks []ConsistencyCheck
	mutex  sync.RWMutex
}

type ConsistencyCheck struct {
	Name        string
	Query       string
	Description string
}

func (cm *ConsistencyManager) AddCheck(check ConsistencyCheck) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.checks = append(cm.checks, check)
}

func (cm *ConsistencyManager) ValidateConsistency(ctx context.Context) error {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	
	for _, check := range cm.checks {
		var count int
		err := cm.db.QueryRowContext(ctx, check.Query).Scan(&count)
		if err != nil {
			return fmt.Errorf("consistency check '%s' failed: %v", check.Name, err)
		}
		
		if count > 0 {
			return fmt.Errorf("consistency check '%s' failed: %s", check.Name, check.Description)
		}
	}
	
	return nil
}

// Example 3: Isolation - Concurrent Access Control
type IsolationManager struct {
	db *sql.DB
}

func (im *IsolationManager) ReadWithIsolation(ctx context.Context, query string, isolationLevel sql.IsolationLevel) (*sql.Rows, error) {
	tx, err := im.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: isolationLevel,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}
	
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	
	// Note: In real implementation, you'd need to handle the transaction lifecycle
	// This is simplified for demonstration
	return rows, nil
}

// Example 4: Durability - Persistent Storage
type DurabilityManager struct {
	db *sql.DB
}

func (dm *DurabilityManager) WriteWithDurability(ctx context.Context, query string, args ...interface{}) error {
	// Use WAL (Write-Ahead Logging) for durability
	tx, err := dm.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	
	// Force flush to disk
	err = tx.Commit()
	if err != nil {
		return err
	}
	
	// Additional durability measures
	_, err = dm.db.Exec("FLUSH LOGS")
	return err
}
```

---

## 🔐 Isolation Levels

### Understanding Isolation Levels

Isolation levels control the degree to which one transaction must be isolated from resource or data modifications made by other transactions.

### Isolation Level Types

#### 1. Read Uncommitted
- Lowest isolation level
- Allows dirty reads
- No locks on reads

#### 2. Read Committed
- Prevents dirty reads
- Allows non-repeatable reads
- Default in most databases

#### 3. Repeatable Read
- Prevents dirty and non-repeatable reads
- May allow phantom reads
- Locks read data

#### 4. Serializable
- Highest isolation level
- Prevents all concurrency issues
- May cause deadlocks

### Isolation Level Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Isolation Level Testing
type IsolationTester struct {
	db *sql.DB
}

func (it *IsolationTester) TestReadUncommitted(ctx context.Context) error {
	// Start transaction 1
	tx1, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		return err
	}
	defer tx1.Rollback()
	
	// Start transaction 2
	tx2, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		return err
	}
	defer tx2.Rollback()
	
	// Transaction 1: Update data
	_, err = tx1.ExecContext(ctx, "UPDATE accounts SET balance = 1000 WHERE id = 'account1'")
	if err != nil {
		return err
	}
	
	// Transaction 2: Read uncommitted data (dirty read)
	var balance float64
	err = tx2.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = 'account1'").Scan(&balance)
	if err != nil {
		return err
	}
	
	fmt.Printf("Read uncommitted balance: %.2f\n", balance)
	
	// Transaction 1: Rollback
	tx1.Rollback()
	
	// Transaction 2: Read again (should be different)
	err = tx2.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = 'account1'").Scan(&balance)
	if err != nil {
		return err
	}
	
	fmt.Printf("After rollback balance: %.2f\n", balance)
	
	return tx2.Commit()
}

func (it *IsolationTester) TestReadCommitted(ctx context.Context) error {
	// Start transaction 1
	tx1, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}
	defer tx1.Rollback()
	
	// Start transaction 2
	tx2, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}
	defer tx2.Rollback()
	
	// Transaction 1: Read data
	var balance1 float64
	err = tx1.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = 'account1'").Scan(&balance1)
	if err != nil {
		return err
	}
	
	// Transaction 2: Update data
	_, err = tx2.ExecContext(ctx, "UPDATE accounts SET balance = 2000 WHERE id = 'account1'")
	if err != nil {
		return err
	}
	tx2.Commit()
	
	// Transaction 1: Read again (non-repeatable read)
	var balance2 float64
	err = tx1.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = 'account1'").Scan(&balance2)
	if err != nil {
		return err
	}
	
	fmt.Printf("First read: %.2f, Second read: %.2f\n", balance1, balance2)
	
	return tx1.Commit()
}

// Example 2: Phantom Read Testing
func (it *IsolationTester) TestPhantomReads(ctx context.Context) error {
	// Start transaction 1
	tx1, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return err
	}
	defer tx1.Rollback()
	
	// Start transaction 2
	tx2, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return err
	}
	defer tx2.Rollback()
	
	// Transaction 1: Count records
	var count1 int
	err = tx1.QueryRowContext(ctx, "SELECT COUNT(*) FROM accounts WHERE balance > 1000").Scan(&count1)
	if err != nil {
		return err
	}
	
	// Transaction 2: Insert new record
	_, err = tx2.ExecContext(ctx, "INSERT INTO accounts (id, balance) VALUES ('account3', 1500)")
	if err != nil {
		return err
	}
	tx2.Commit()
	
	// Transaction 1: Count again (phantom read)
	var count2 int
	err = tx1.QueryRowContext(ctx, "SELECT COUNT(*) FROM accounts WHERE balance > 1000").Scan(&count2)
	if err != nil {
		return err
	}
	
	fmt.Printf("First count: %d, Second count: %d\n", count1, count2)
	
	return tx1.Commit()
}

// Example 3: Serializable Isolation
func (it *IsolationTester) TestSerializable(ctx context.Context) error {
	// Start transaction 1
	tx1, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx1.Rollback()
	
	// Start transaction 2
	tx2, err := it.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx2.Rollback()
	
	// Both transactions try to update the same data
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		_, err := tx1.ExecContext(ctx, "UPDATE accounts SET balance = balance + 100 WHERE id = 'account1'")
		if err != nil {
			fmt.Printf("Transaction 1 error: %v\n", err)
		}
	}()
	
	go func() {
		defer wg.Done()
		_, err := tx2.ExecContext(ctx, "UPDATE accounts SET balance = balance + 200 WHERE id = 'account1'")
		if err != nil {
			fmt.Printf("Transaction 2 error: %v\n", err)
		}
	}()
	
	wg.Wait()
	
	// One transaction will succeed, one will fail
	err1 := tx1.Commit()
	err2 := tx2.Commit()
	
	if err1 != nil && err2 != nil {
		return fmt.Errorf("both transactions failed: %v, %v", err1, err2)
	}
	
	return nil
}
```

---

## ⚰️ Deadlock Prevention

### Understanding Deadlocks

A deadlock occurs when two or more transactions are waiting for each other to release locks, creating a circular dependency that prevents any of them from proceeding.

### Deadlock Prevention Strategies

#### 1. Lock Ordering
- Always acquire locks in the same order
- Prevents circular dependencies

#### 2. Timeout
- Set maximum wait time for locks
- Automatically abort long-running transactions

#### 3. Deadlock Detection
- Monitor for circular dependencies
- Abort one transaction to break the cycle

### Deadlock Prevention Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Lock Ordering
type LockOrderingManager struct {
	db     *sql.DB
	lockOrder map[string]int
	mutex  sync.RWMutex
}

func NewLockOrderingManager(db *sql.DB) *LockOrderingManager {
	return &LockOrderingManager{
		db: db,
		lockOrder: map[string]int{
			"account1": 1,
			"account2": 2,
			"account3": 3,
		},
	}
}

func (lom *LockOrderingManager) TransferMoney(ctx context.Context, fromAccount, toAccount string, amount float64) error {
	// Determine lock order
	fromOrder := lom.getLockOrder(fromAccount)
	toOrder := lom.getLockOrder(toAccount)
	
	// Ensure consistent lock ordering
	if fromOrder > toOrder {
		fromAccount, toAccount = toAccount, fromAccount
		fromOrder, toOrder = toOrder, fromOrder
	}
	
	tx, err := lom.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Lock accounts in order
	_, err = tx.ExecContext(ctx, "SELECT * FROM accounts WHERE id = ? FOR UPDATE", fromAccount)
	if err != nil {
		return err
	}
	
	_, err = tx.ExecContext(ctx, "SELECT * FROM accounts WHERE id = ? FOR UPDATE", toAccount)
	if err != nil {
		return err
	}
	
	// Perform transfer
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccount)
	if err != nil {
		return err
	}
	
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccount)
	if err != nil {
		return err
	}
	
	return tx.Commit()
}

func (lom *LockOrderingManager) getLockOrder(accountID string) int {
	lom.mutex.RLock()
	defer lom.mutex.RUnlock()
	
	if order, exists := lom.lockOrder[accountID]; exists {
		return order
	}
	return 999 // Default order for unknown accounts
}

// Example 2: Timeout-based Deadlock Prevention
type TimeoutManager struct {
	db      *sql.DB
	timeout time.Duration
}

func NewTimeoutManager(db *sql.DB, timeout time.Duration) *TimeoutManager {
	return &TimeoutManager{
		db:      db,
		timeout: timeout,
	}
}

func (tm *TimeoutManager) ExecuteWithTimeout(ctx context.Context, fn func(*sql.Tx) error) error {
	// Create context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, tm.timeout)
	defer cancel()
	
	tx, err := tm.db.BeginTx(timeoutCtx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Execute function with timeout
	done := make(chan error, 1)
	go func() {
		done <- fn(tx)
	}()
	
	select {
	case err := <-done:
		if err != nil {
			return err
		}
		return tx.Commit()
	case <-timeoutCtx.Done():
		return fmt.Errorf("transaction timed out after %v", tm.timeout)
	}
}

// Example 3: Deadlock Detection
type DeadlockDetector struct {
	db           *sql.DB
	activeLocks  map[string][]string
	mutex        sync.RWMutex
	checkInterval time.Duration
	stopCh       chan struct{}
}

func NewDeadlockDetector(db *sql.DB, checkInterval time.Duration) *DeadlockDetector {
	return &DeadlockDetector{
		db:            db,
		activeLocks:  make(map[string][]string),
		checkInterval: checkInterval,
		stopCh:       make(chan struct{}),
	}
}

func (dd *DeadlockDetector) Start() {
	ticker := time.NewTicker(dd.checkInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			dd.detectDeadlocks()
		case <-dd.stopCh:
			return
		}
	}
}

func (dd *DeadlockDetector) Stop() {
	close(dd.stopCh)
}

func (dd *DeadlockDetector) detectDeadlocks() {
	dd.mutex.RLock()
	defer dd.mutex.RUnlock()
	
	// Simplified deadlock detection
	// In real implementation, you'd query the database for lock information
	for txID, locks := range dd.activeLocks {
		if len(locks) > 1 {
			// Check for circular dependencies
			if dd.hasCircularDependency(txID, locks) {
				fmt.Printf("Deadlock detected for transaction %s\n", txID)
				dd.abortTransaction(txID)
			}
		}
	}
}

func (dd *DeadlockDetector) hasCircularDependency(txID string, locks []string) bool {
	// Simplified circular dependency check
	// In real implementation, you'd build a graph and check for cycles
	return len(locks) > 2
}

func (dd *DeadlockDetector) abortTransaction(txID string) {
	// In real implementation, you'd abort the transaction
	fmt.Printf("Aborting transaction %s\n", txID)
}

// Example 4: Optimistic Locking
type OptimisticLockingManager struct {
	db *sql.DB
}

func (olm *OptimisticLockingManager) UpdateWithOptimisticLock(ctx context.Context, id string, newBalance float64, expectedVersion int) error {
	tx, err := olm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Check version before update
	var currentVersion int
	err = tx.QueryRowContext(ctx, "SELECT version FROM accounts WHERE id = ?", id).Scan(&currentVersion)
	if err != nil {
		return err
	}
	
	if currentVersion != expectedVersion {
		return fmt.Errorf("version mismatch: expected %d, got %d", expectedVersion, currentVersion)
	}
	
	// Update with version increment
	result, err := tx.ExecContext(ctx, 
		"UPDATE accounts SET balance = ?, version = version + 1 WHERE id = ? AND version = ?",
		newBalance, id, expectedVersion)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("update failed: version mismatch")
	}
	
	return tx.Commit()
}
```

---

## 📖 Read Replicas

### Understanding Read Replicas

Read replicas are copies of the primary database that are used for read operations, reducing load on the primary database and improving performance.

### Read Replica Benefits

#### 1. Performance
- Distribute read load
- Reduce primary database pressure
- Improve response times

#### 2. Scalability
- Handle more concurrent reads
- Scale horizontally
- Better resource utilization

#### 3. Availability
- Backup for read operations
- Disaster recovery
- Geographic distribution

### Read Replica Implementation

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Read Replica Manager
type ReadReplicaManager struct {
	primary    *sql.DB
	replicas   []*sql.DB
	replicaIndex int
	mutex      sync.RWMutex
}

func NewReadReplicaManager(primary *sql.DB, replicas []*sql.DB) *ReadReplicaManager {
	return &ReadReplicaManager{
		primary:    primary,
		replicas:   replicas,
		replicaIndex: 0,
	}
}

func (rrm *ReadReplicaManager) Write(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Always write to primary
	return rrm.primary.ExecContext(ctx, query, args...)
}

func (rrm *ReadReplicaManager) Read(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Read from replica with round-robin
	rrm.mutex.Lock()
	replica := rrm.replicas[rrm.replicaIndex]
	rrm.replicaIndex = (rrm.replicaIndex + 1) % len(rrm.replicas)
	rrm.mutex.Unlock()
	
	return replica.QueryContext(ctx, query, args...)
}

func (rrm *ReadReplicaManager) ReadOne(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// Read from replica with round-robin
	rrm.mutex.Lock()
	replica := rrm.replicas[rrm.replicaIndex]
	rrm.replicaIndex = (rrm.replicaIndex + 1) % len(rrm.replicas)
	rrm.mutex.Unlock()
	
	return replica.QueryRowContext(ctx, query, args...)
}

// Example 2: Advanced Read Replica with Health Checks
type AdvancedReadReplicaManager struct {
	*ReadReplicaManager
	healthChecker *ReplicaHealthChecker
	healthyReplicas []*sql.DB
}

type ReplicaHealthChecker struct {
	interval time.Duration
	timeout  time.Duration
	stopCh   chan struct{}
}

func NewAdvancedReadReplicaManager(primary *sql.DB, replicas []*sql.DB) *AdvancedReadReplicaManager {
	rrm := NewReadReplicaManager(primary, replicas)
	
	acp := &AdvancedReadReplicaManager{
		ReadReplicaManager: rrm,
		healthChecker: &ReplicaHealthChecker{
			interval: 30 * time.Second,
			timeout:  5 * time.Second,
			stopCh:   make(chan struct{}),
		},
		healthyReplicas: make([]*sql.DB, 0, len(replicas)),
	}
	
	// Start health checking
	go acp.healthChecker.Start(acp)
	
	return acp
}

func (rhc *ReplicaHealthChecker) Start(manager *AdvancedReadReplicaManager) {
	ticker := time.NewTicker(rhc.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			manager.checkReplicaHealth()
		case <-rhc.stopCh:
			return
		}
	}
}

func (acp *AdvancedReadReplicaManager) checkReplicaHealth() {
	acp.mutex.Lock()
	defer acp.mutex.Unlock()
	
	acp.healthyReplicas = acp.healthyReplicas[:0]
	
	for _, replica := range acp.replicas {
		ctx, cancel := context.WithTimeout(context.Background(), acp.healthChecker.timeout)
		if err := replica.PingContext(ctx); err == nil {
			acp.healthyReplicas = append(acp.healthyReplicas, replica)
		}
		cancel()
	}
}

func (acp *AdvancedReadReplicaManager) Read(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	acp.mutex.RLock()
	defer acp.mutex.RUnlock()
	
	if len(acp.healthyReplicas) == 0 {
		// Fallback to primary if no healthy replicas
		return acp.primary.QueryContext(ctx, query, args...)
	}
	
	// Use round-robin on healthy replicas
	replica := acp.healthyReplicas[acp.replicaIndex%len(acp.healthyReplicas)]
	acp.replicaIndex++
	
	return replica.QueryContext(ctx, query, args...)
}

// Example 3: Read Replica with Consistency Levels
type ConsistencyLevel int

const (
	EventualConsistency ConsistencyLevel = iota
	StrongConsistency
	SessionConsistency
)

type ConsistencyAwareReadReplicaManager struct {
	*AdvancedReadReplicaManager
	consistencyLevel ConsistencyLevel
}

func NewConsistencyAwareReadReplicaManager(primary *sql.DB, replicas []*sql.DB, level ConsistencyLevel) *ConsistencyAwareReadReplicaManager {
	return &ConsistencyAwareReadReplicaManager{
		AdvancedReadReplicaManager: NewAdvancedReadReplicaManager(primary, replicas),
		consistencyLevel:           level,
	}
}

func (carm *ConsistencyAwareReadReplicaManager) Read(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch carm.consistencyLevel {
	case StrongConsistency:
		// Always read from primary for strong consistency
		return carm.primary.QueryContext(ctx, query, args...)
	case SessionConsistency:
		// Use session-based consistency
		return carm.readWithSessionConsistency(ctx, query, args...)
	case EventualConsistency:
		// Use replica with eventual consistency
		return carm.AdvancedReadReplicaManager.Read(ctx, query, args...)
	default:
		return carm.AdvancedReadReplicaManager.Read(ctx, query, args...)
	}
}

func (carm *ConsistencyAwareReadReplicaManager) readWithSessionConsistency(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Simplified session consistency implementation
	// In real implementation, you'd track session state
	return carm.AdvancedReadReplicaManager.Read(ctx, query, args...)
}
```

---

## 🎯 Summary

Database Concurrency is essential for building production-ready concurrent systems. Key takeaways:

1. **Connection Pooling** improves performance and resource utilization
2. **Transaction Management** ensures data consistency and integrity
3. **ACID Properties** guarantee reliable database operations
4. **Isolation Levels** control concurrency and consistency trade-offs
5. **Deadlock Prevention** avoids circular dependencies and timeouts
6. **Read Replicas** improve performance and scalability

This topic provides the foundation for building robust, concurrent database systems! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different concurrency patterns
3. **Apply** techniques to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** database concurrency patterns

Ready to become a Database Concurrency expert? Let's dive into the implementation! 💪

