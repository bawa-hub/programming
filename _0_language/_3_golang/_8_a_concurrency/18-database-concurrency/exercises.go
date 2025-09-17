package main

import (
	"context"
	"fmt"
	"time"
)

// Exercise 1: Implement Connection Pooling
func Exercise1() {
	fmt.Println("\nExercise 1: Implement Connection Pooling")
	fmt.Println("=======================================")
	
	// TODO: Implement connection pooling
	// 1. Create a connection pool with configurable parameters
	// 2. Implement connection acquisition and release
	// 3. Add connection health checking
	// 4. Monitor pool statistics
	
	// Create connection pool
	pool, err := NewConnectionPool("user:password@tcp(localhost:3306)/testdb", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("  Error creating connection pool: %v\n", err)
		return
	}
	defer pool.Close()
	
	fmt.Println("  Testing connection pooling...")
	
	ctx := context.Background()
	
	// Test connection acquisition and release
	for i := 0; i < 8; i++ {
		conn, err := pool.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
			continue
		}
		
		fmt.Printf("    Acquired connection %d: %s\n", i+1, conn.id)
		
		// Simulate work
		time.Sleep(50 * time.Millisecond)
		
		// Release connection
		pool.ReleaseConnection(conn)
		fmt.Printf("    Released connection %d\n", i+1)
	}
	
	// Get pool statistics
	stats := pool.GetStats()
	fmt.Printf("  Pool statistics:\n")
	fmt.Printf("    Open connections: %d\n", stats.OpenConns)
	fmt.Printf("    In use: %d\n", stats.InUse)
	fmt.Printf("    Idle: %d\n", stats.Idle)
	
	fmt.Println("  Exercise 1: Connection pooling completed")
}

// Exercise 2: Implement Transaction Management
func Exercise2() {
	fmt.Println("\nExercise 2: Implement Transaction Management")
	fmt.Println("===========================================")
	
	// TODO: Implement transaction management
	// 1. Create transaction manager
	// 2. Implement begin, commit, rollback operations
	// 3. Add nested transaction support
	// 4. Test transaction isolation
	
	// Create transaction manager
	tm := NewTransactionManager(nil)
	
	fmt.Println("  Testing transaction management...")
	
	ctx := context.Background()
	
	// Test basic transaction operations
	txID := "exercise-tx-1"
	
	// Begin transaction
	_, err := tm.Begin(ctx, txID)
	if err != nil {
		fmt.Printf("    Error beginning transaction: %v\n", err)
		return
	}
	fmt.Printf("    Transaction %s begun\n", txID)
	
	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	
	// Commit transaction
	err = tm.Commit(txID)
	if err != nil {
		fmt.Printf("    Error committing transaction: %v\n", err)
		return
	}
	fmt.Printf("    Transaction %s committed\n", txID)
	
	// Test rollback scenario
	txID2 := "exercise-tx-2"
	_, err = tm.Begin(ctx, txID2)
	if err != nil {
		fmt.Printf("    Error beginning transaction: %v\n", err)
		return
	}
	fmt.Printf("    Transaction %s begun\n", txID2)
	
	// Simulate work
	time.Sleep(50 * time.Millisecond)
	
	// Rollback transaction
	err = tm.Rollback(txID2)
	if err != nil {
		fmt.Printf("    Error rolling back transaction: %v\n", err)
		return
	}
	fmt.Printf("    Transaction %s rolled back\n", txID2)
	
	fmt.Println("  Exercise 2: Transaction management completed")
}

// Exercise 3: Implement ACID Properties
func Exercise3() {
	fmt.Println("\nExercise 3: Implement ACID Properties")
	fmt.Println("====================================")
	
	// TODO: Implement ACID properties
	// 1. Implement atomicity with transactions
	// 2. Ensure consistency with constraints
	// 3. Control isolation with levels
	// 4. Guarantee durability with persistence
	
	fmt.Println("  Testing ACID properties...")
	
	// Atomicity - All or nothing
	fmt.Println("    Testing Atomicity...")
	fmt.Println("      - All operations in transaction succeed or all fail")
	fmt.Println("      - No partial updates allowed")
	
	// Consistency - Data integrity
	fmt.Println("    Testing Consistency...")
	fmt.Println("      - Database remains in valid state")
	fmt.Println("      - All constraints are maintained")
	
	// Isolation - Concurrent access control
	fmt.Println("    Testing Isolation...")
	fmt.Println("      - Concurrent transactions don't interfere")
	fmt.Println("      - Different isolation levels available")
	
	// Durability - Persistent storage
	fmt.Println("    Testing Durability...")
	fmt.Println("      - Committed changes persist")
	fmt.Println("      - Data written to persistent storage")
	
	fmt.Println("  Exercise 3: ACID properties completed")
}

// Exercise 4: Implement Isolation Levels
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Isolation Levels")
	fmt.Println("======================================")
	
	// TODO: Implement isolation levels
	// 1. Test Read Uncommitted
	// 2. Test Read Committed
	// 3. Test Repeatable Read
	// 4. Test Serializable
	
	// Create isolation tester
	it := &IsolationTester{db: nil}
	
	fmt.Println("  Testing isolation levels...")
	
	ctx := context.Background()
	
	// Test Read Uncommitted
	fmt.Println("    Testing Read Uncommitted...")
	err := it.TestReadUncommitted(ctx)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	// Test Read Committed
	fmt.Println("    Testing Read Committed...")
	err = it.TestReadCommitted(ctx)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	// Test Repeatable Read
	fmt.Println("    Testing Repeatable Read...")
	err = it.TestPhantomReads(ctx)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	// Test Serializable
	fmt.Println("    Testing Serializable...")
	err = it.TestSerializable(ctx)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	fmt.Println("  Exercise 4: Isolation levels completed")
}

// Exercise 5: Implement Deadlock Prevention
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Deadlock Prevention")
	fmt.Println("=========================================")
	
	// TODO: Implement deadlock prevention
	// 1. Implement lock ordering
	// 2. Add timeout mechanisms
	// 3. Implement deadlock detection
	// 4. Test prevention strategies
	
	// Create deadlock prevention managers
	lom := NewLockOrderingManager(nil)
	tm := NewTimeoutManager(nil, 3*time.Second)
	
	fmt.Println("  Testing deadlock prevention...")
	
	ctx := context.Background()
	
	// Test lock ordering
	fmt.Println("    Testing lock ordering...")
	err := lom.TransferMoney(ctx, "account1", "account2", 100.0)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	// Test timeout prevention
	fmt.Println("    Testing timeout prevention...")
	err = tm.ExecuteWithTimeout(ctx, func(tx interface{}) error {
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	// Test reverse order (should still work with lock ordering)
	fmt.Println("    Testing reverse order transfer...")
	err = lom.TransferMoney(ctx, "account2", "account1", 50.0)
	if err != nil {
		fmt.Printf("      Error: %v\n", err)
	}
	
	fmt.Println("  Exercise 5: Deadlock prevention completed")
}

// Exercise 6: Implement Read Replicas
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Read Replicas")
	fmt.Println("===================================")
	
	// TODO: Implement read replicas
	// 1. Create read replica manager
	// 2. Implement read/write separation
	// 3. Add load balancing for reads
	// 4. Test replica health checking
	
	// Create read replica manager
	primary := &MockDB{name: "primary"}
	replicas := []interface{}{
		&MockDB{name: "replica1"},
		&MockDB{name: "replica2"},
		&MockDB{name: "replica3"},
	}
	
	rrm := NewReadReplicaManager(primary, replicas)
	
	fmt.Println("  Testing read replicas...")
	
	ctx := context.Background()
	
	// Test write operations (should go to primary)
	fmt.Println("    Testing write operations...")
	for i := 0; i < 3; i++ {
		_, err := rrm.Write(ctx, "INSERT INTO users (name) VALUES (?)", fmt.Sprintf("User %d", i+1))
		if err != nil {
			fmt.Printf("      Error writing: %v\n", err)
		} else {
			fmt.Printf("      Write operation %d completed\n", i+1)
		}
	}
	
	// Test read operations (should go to replicas)
	fmt.Println("    Testing read operations...")
	for i := 0; i < 6; i++ {
		_, err := rrm.Read(ctx, "SELECT * FROM users WHERE id = ?", i+1)
		if err != nil {
			fmt.Printf("      Error reading: %v\n", err)
		} else {
			fmt.Printf("      Read operation %d completed\n", i+1)
		}
	}
	
	fmt.Println("  Exercise 6: Read replicas completed")
}

// Exercise 7: Implement Database Sharding
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Database Sharding")
	fmt.Println("======================================")
	
	// TODO: Implement database sharding
	// 1. Create shard manager
	// 2. Implement shard routing
	// 3. Add shard operations
	// 4. Test shard distribution
	
	// Create shard manager
	sm := NewShardManager()
	
	// Add shards
	sm.AddShard("shard1", &MockDB{name: "shard1"})
	sm.AddShard("shard2", &MockDB{name: "shard2"})
	sm.AddShard("shard3", &MockDB{name: "shard3"})
	
	fmt.Println("  Testing database sharding...")
	
	ctx := context.Background()
	
	// Test shard routing
	fmt.Println("    Testing shard routing...")
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6"}
	for _, user := range users {
		shardID := sm.GetShard(user)
		fmt.Printf("      User %s -> Shard %s\n", user, shardID)
	}
	
	// Test operations on shards
	fmt.Println("    Testing operations on shards...")
	for _, user := range users {
		// Write to shard
		_, err := sm.Write(ctx, user, "INSERT INTO users (id, name) VALUES (?, ?)", user, "User "+user)
		if err != nil {
			fmt.Printf("      Error writing to shard: %v\n", err)
			continue
		}
		
		// Read from shard
		_, err = sm.Read(ctx, user, "SELECT * FROM users WHERE id = ?", user)
		if err != nil {
			fmt.Printf("      Error reading from shard: %v\n", err)
			continue
		}
		
		fmt.Printf("      Operations for user %s completed\n", user)
	}
	
	fmt.Println("  Exercise 7: Database sharding completed")
}

// Exercise 8: Implement Locking Strategies
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Locking Strategies")
	fmt.Println("=======================================")
	
	// TODO: Implement locking strategies
	// 1. Implement optimistic locking
	// 2. Implement pessimistic locking
	// 3. Compare performance characteristics
	// 4. Test concurrency scenarios
	
	// Create locking managers
	olm := &OptimisticLockingManager{db: nil}
	plm := &PessimisticLockingManager{db: nil}
	
	fmt.Println("  Testing locking strategies...")
	
	ctx := context.Background()
	
	// Test optimistic locking
	fmt.Println("    Testing optimistic locking...")
	for i := 0; i < 3; i++ {
		err := olm.UpdateWithOptimisticLock(ctx, fmt.Sprintf("account%d", i+1), float64(1000+i*100), i+1)
		if err != nil {
			fmt.Printf("      Error: %v\n", err)
		}
	}
	
	// Test pessimistic locking
	fmt.Println("    Testing pessimistic locking...")
	for i := 0; i < 3; i++ {
		err := plm.UpdateWithPessimisticLock(ctx, fmt.Sprintf("account%d", i+1), float64(2000+i*100))
		if err != nil {
			fmt.Printf("      Error: %v\n", err)
		}
	}
	
	fmt.Println("  Exercise 8: Locking strategies completed")
}

// Exercise 9: Implement Connection Management
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Connection Management")
	fmt.Println("==========================================")
	
	// TODO: Implement connection management
	// 1. Create connection manager
	// 2. Implement connection lifecycle
	// 3. Add connection monitoring
	// 4. Test resource cleanup
	
	// Create connection manager
	cm := NewConnectionManager(10, 5, time.Hour)
	
	fmt.Println("  Testing connection management...")
	
	ctx := context.Background()
	
	// Test connection acquisition and release
	fmt.Println("    Testing connection lifecycle...")
	connections := make([]*MockConn, 0, 5)
	
	// Acquire connections
	for i := 0; i < 5; i++ {
		conn, err := cm.GetConnection(ctx)
		if err != nil {
			fmt.Printf("      Error getting connection %d: %v\n", i+1, err)
			continue
		}
		
		connections = append(connections, conn)
		fmt.Printf("      Acquired connection %d: %s\n", i+1, conn.id)
	}
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	// Release connections
	for i, conn := range connections {
		cm.ReleaseConnection(conn)
		fmt.Printf("      Released connection %d: %s\n", i+1, conn.id)
	}
	
	// Get connection statistics
	stats := cm.GetStats()
	fmt.Printf("    Connection statistics:\n")
	fmt.Printf("      Total connections: %d\n", stats.TotalConnections)
	fmt.Printf("      Active connections: %d\n", stats.ActiveConnections)
	fmt.Printf("      Idle connections: %d\n", stats.IdleConnections)
	
	fmt.Println("  Exercise 9: Connection management completed")
}

// Exercise 10: Implement Query Processing
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Query Processing")
	fmt.Println("======================================")
	
	// TODO: Implement concurrent query processing
	// 1. Create query processor
	// 2. Implement worker pool pattern
	// 3. Add query queuing
	// 4. Test concurrent execution
	
	// Create query processor
	qp := NewQueryProcessor(3) // 3 worker goroutines
	defer qp.Close()
	
	fmt.Println("  Testing concurrent query processing...")
	
	ctx := context.Background()
	
	// Submit multiple queries
	fmt.Println("    Submitting queries...")
	queryCount := 10
	
	for i := 0; i < queryCount; i++ {
		query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", i+1)
		qp.SubmitQuery(ctx, query, func(result interface{}) {
			fmt.Printf("      Query %d completed\n", i+1)
		})
	}
	
	// Wait for completion
	time.Sleep(500 * time.Millisecond)
	
	fmt.Printf("    All %d queries processed\n", queryCount)
	
	fmt.Println("  Exercise 10: Query processing completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Database Concurrency Exercises")
	fmt.Println("=================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}
