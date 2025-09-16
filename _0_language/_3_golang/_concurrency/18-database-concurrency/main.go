package main

import (
	"context"
	"fmt"
	"time"
)

// Example 1: Basic Database Connection Management
func basicDatabaseConnection() {
	fmt.Println("\n1. Basic Database Connection Management")
	fmt.Println("======================================")
	
	// Simulate database connection
	config := &DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		Database: "testdb",
		Username: "user",
		Password: "password",
		MaxConns: 10,
		MaxIdle:  5,
	}
	
	// Create database manager
	dm, err := NewDatabaseManager(config)
	if err != nil {
		fmt.Printf("  Error creating database manager: %v\n", err)
		return
	}
	defer dm.Close()
	
	fmt.Println("  Database manager created successfully")
	fmt.Printf("  Max connections: %d\n", config.MaxConns)
	fmt.Printf("  Max idle connections: %d\n", config.MaxIdle)
	
	fmt.Println("  Basic database connection management completed")
}

// Example 2: Connection Pooling
func connectionPooling() {
	fmt.Println("\n2. Connection Pooling")
	fmt.Println("====================")
	
	// Create connection pool
	pool, err := NewConnectionPool("user:password@tcp(localhost:3306)/testdb", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("  Error creating connection pool: %v\n", err)
		return
	}
	defer pool.Close()
	
	// Test connection pool
	fmt.Println("  Testing connection pool...")
	
	ctx := context.Background()
	
	// Get multiple connections
	for i := 0; i < 5; i++ {
		conn, err := pool.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
			continue
		}
		
		fmt.Printf("    Got connection %d\n", i+1)
		
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		
		// Release connection
		pool.ReleaseConnection(conn)
		fmt.Printf("    Released connection %d\n", i+1)
	}
	
	// Get pool statistics
	stats := pool.GetStats()
	fmt.Printf("  Pool stats - Open: %d, InUse: %d, Idle: %d\n", 
		stats.OpenConns, stats.InUse, stats.Idle)
	
	fmt.Println("  Connection pooling completed")
}

// Example 3: Transaction Management
func transactionManagement() {
	fmt.Println("\n3. Transaction Management")
	fmt.Println("========================")
	
	// Create transaction manager
	tm := NewTransactionManager(nil) // Simulated database
	
	fmt.Println("  Testing transaction management...")
	
	ctx := context.Background()
	
	// Simulate transaction operations
	txID := "tx-123"
	
	// Begin transaction
	_, err := tm.Begin(ctx, txID)
	if err != nil {
		fmt.Printf("  Error beginning transaction: %v\n", err)
		return
	}
	
	fmt.Printf("  Transaction %s begun\n", txID)
	
	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	
	// Commit transaction
	err = tm.Commit(txID)
	if err != nil {
		fmt.Printf("  Error committing transaction: %v\n", err)
		return
	}
	
	fmt.Printf("  Transaction %s committed\n", txID)
	
	fmt.Println("  Transaction management completed")
}

// Example 4: ACID Properties
func acidProperties() {
	fmt.Println("\n4. ACID Properties")
	fmt.Println("==================")
	
	// Simulate ACID operations
	fmt.Println("  Testing ACID properties...")
	
	// Atomicity - All or nothing
	fmt.Println("    Atomicity: Ensuring all operations succeed or all fail")
	
	// Consistency - Data integrity
	fmt.Println("    Consistency: Maintaining data integrity constraints")
	
	// Isolation - Concurrent access control
	fmt.Println("    Isolation: Preventing concurrent operations from interfering")
	
	// Durability - Persistent storage
	fmt.Println("    Durability: Ensuring committed changes persist")
	
	fmt.Println("  ACID properties demonstration completed")
}

// Example 5: Isolation Levels
func isolationLevels() {
	fmt.Println("\n5. Isolation Levels")
	fmt.Println("===================")
	
	// Create isolation tester
	it := &IsolationTester{db: nil} // Simulated database
	
	fmt.Println("  Testing isolation levels...")
	
	ctx := context.Background()
	
	// Test different isolation levels
	fmt.Println("    Testing Read Uncommitted...")
	err := it.TestReadUncommitted(ctx)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("    Testing Read Committed...")
	err = it.TestReadCommitted(ctx)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("    Testing Phantom Reads...")
	err = it.TestPhantomReads(ctx)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("    Testing Serializable...")
	err = it.TestSerializable(ctx)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("  Isolation levels testing completed")
}

// Example 6: Deadlock Prevention
func deadlockPrevention() {
	fmt.Println("\n6. Deadlock Prevention")
	fmt.Println("======================")
	
	// Create deadlock prevention managers
	lom := NewLockOrderingManager(nil) // Simulated database
	tm := NewTimeoutManager(nil, 5*time.Second) // Simulated database
	
	fmt.Println("  Testing deadlock prevention...")
	
	ctx := context.Background()
	
	// Test lock ordering
	fmt.Println("    Testing lock ordering...")
	err := lom.TransferMoney(ctx, "account1", "account2", 100.0)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	// Test timeout
	fmt.Println("    Testing timeout prevention...")
	err = tm.ExecuteWithTimeout(ctx, func(tx interface{}) error {
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("  Deadlock prevention testing completed")
}

// Example 7: Read Replicas
func readReplicas() {
	fmt.Println("\n7. Read Replicas")
	fmt.Println("================")
	
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
	
	// Test write operations (go to primary)
	fmt.Println("    Testing write operations...")
	_, err := rrm.Write(ctx, "INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	// Test read operations (go to replicas)
	fmt.Println("    Testing read operations...")
	for i := 0; i < 5; i++ {
		_, err := rrm.Read(ctx, "SELECT * FROM users WHERE id = ?", i+1)
		if err != nil {
			fmt.Printf("    Error: %v\n", err)
			continue
		}
		fmt.Printf("    Read operation %d completed\n", i+1)
	}
	
	fmt.Println("  Read replicas testing completed")
}

// Example 8: Database Sharding
func databaseSharding() {
	fmt.Println("\n8. Database Sharding")
	fmt.Println("===================")
	
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
	for i := 0; i < 10; i++ {
		shardID := sm.GetShard(fmt.Sprintf("user%d", i))
		fmt.Printf("    User %d -> Shard %s\n", i, shardID)
	}
	
	// Test operations on shards
	fmt.Println("    Testing operations on shards...")
	for i := 0; i < 5; i++ {
		userID := fmt.Sprintf("user%d", i)
		shardID := sm.GetShard(userID)
		
		// Write to shard
		_, err := sm.Write(ctx, userID, "INSERT INTO users (id, name) VALUES (?, ?)", userID, "User "+userID)
		if err != nil {
			fmt.Printf("    Error writing to shard %s: %v\n", shardID, err)
			continue
		}
		
		// Read from shard
		_, err = sm.Read(ctx, userID, "SELECT * FROM users WHERE id = ?", userID)
		if err != nil {
			fmt.Printf("    Error reading from shard %s: %v\n", shardID, err)
			continue
		}
		
		fmt.Printf("    Operations on shard %s completed\n", shardID)
	}
	
	fmt.Println("  Database sharding testing completed")
}

// Example 9: Optimistic vs Pessimistic Locking
func lockingStrategies() {
	fmt.Println("\n9. Optimistic vs Pessimistic Locking")
	fmt.Println("====================================")
	
	// Create locking managers
	olm := &OptimisticLockingManager{db: nil} // Simulated database
	plm := &PessimisticLockingManager{db: nil} // Simulated database
	
	fmt.Println("  Testing locking strategies...")
	
	ctx := context.Background()
	
	// Test optimistic locking
	fmt.Println("    Testing optimistic locking...")
	err := olm.UpdateWithOptimisticLock(ctx, "account1", 1000.0, 1)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	// Test pessimistic locking
	fmt.Println("    Testing pessimistic locking...")
	err = plm.UpdateWithPessimisticLock(ctx, "account2", 2000.0)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("  Locking strategies testing completed")
}

// Example 10: Database Connection Management
func databaseConnectionManagement() {
	fmt.Println("\n10. Database Connection Management")
	fmt.Println("==================================")
	
	// Create connection manager
	cm := NewConnectionManager(10, 5, time.Hour)
	
	fmt.Println("  Testing database connection management...")
	
	ctx := context.Background()
	
	// Test connection acquisition and release
	fmt.Println("    Testing connection lifecycle...")
	for i := 0; i < 5; i++ {
		conn, err := cm.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
			continue
		}
		
		fmt.Printf("    Got connection %d\n", i+1)
		
		// Simulate work
		time.Sleep(50 * time.Millisecond)
		
		// Release connection
		cm.ReleaseConnection(conn)
		fmt.Printf("    Released connection %d\n", i+1)
	}
	
	// Test connection pool statistics
	stats := cm.GetStats()
	fmt.Printf("    Pool stats - Total: %d, Active: %d, Idle: %d\n", 
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)
	
	fmt.Println("  Database connection management completed")
}

// Example 11: Concurrent Query Processing
func concurrentQueryProcessing() {
	fmt.Println("\n11. Concurrent Query Processing")
	fmt.Println("===============================")
	
	// Create query processor
	qp := NewQueryProcessor(5) // 5 worker goroutines
	
	fmt.Println("  Testing concurrent query processing...")
	
	ctx := context.Background()
	
	// Submit multiple queries
	fmt.Println("    Submitting queries...")
	for i := 0; i < 10; i++ {
		query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", i+1)
		qp.SubmitQuery(ctx, query, func(result interface{}) {
			fmt.Printf("    Query %d completed\n", i+1)
		})
	}
	
	// Wait for completion
	time.Sleep(500 * time.Millisecond)
	
	fmt.Println("  Concurrent query processing completed")
}

// Example 12: Database Monitoring
func databaseMonitoring() {
	fmt.Println("\n12. Database Monitoring")
	fmt.Println("=======================")
	
	// Create database monitor
	dm := NewDatabaseMonitor()
	
	// Add databases to monitor
	dm.AddDatabase("primary", &MockDB{name: "primary"})
	dm.AddDatabase("replica1", &MockDB{name: "replica1"})
	dm.AddDatabase("replica2", &MockDB{name: "replica2"})
	
	fmt.Println("  Testing database monitoring...")
	
	// Start monitoring
	go dm.StartMonitoring()
	
	// Simulate some activity
	time.Sleep(200 * time.Millisecond)
	
	// Get monitoring data
	metrics := dm.GetMetrics()
	fmt.Printf("    Databases monitored: %d\n", len(metrics))
	
	for dbName, metric := range metrics {
		fmt.Printf("    %s: %d queries, %v avg response time\n", 
			dbName, metric.QueryCount, metric.AvgResponseTime)
	}
	
	fmt.Println("  Database monitoring completed")
}

// Example 13: Performance Optimization
func performanceOptimization() {
	fmt.Println("\n13. Performance Optimization")
	fmt.Println("============================")
	
	// Create performance optimizer
	po := NewPerformanceOptimizer()
	
	fmt.Println("  Testing performance optimization...")
	
	// Test query optimization
	fmt.Println("    Testing query optimization...")
	optimizedQuery := po.OptimizeQuery("SELECT * FROM users WHERE name = 'John' AND age > 25")
	fmt.Printf("    Optimized query: %s\n", optimizedQuery)
	
	// Test index recommendations
	fmt.Println("    Testing index recommendations...")
	indexes := po.RecommendIndexes("users", []string{"name", "age", "email"})
	for _, index := range indexes {
		fmt.Printf("    Recommended index: %s\n", index)
	}
	
	// Test connection pool optimization
	fmt.Println("    Testing connection pool optimization...")
	config := po.OptimizeConnectionPool(100, 50, time.Hour)
	fmt.Printf("    Optimized config: MaxConns=%d, MaxIdle=%d\n", 
		config.MaxConns, config.MaxIdle)
	
	fmt.Println("  Performance optimization completed")
}

// Example 14: Error Handling
func errorHandling() {
	fmt.Println("\n14. Error Handling")
	fmt.Println("==================")
	
	// Create error handler
	eh := NewDatabaseErrorHandler()
	
	fmt.Println("  Testing error handling...")
	
	ctx := context.Background()
	
	// Test different types of errors
	fmt.Println("    Testing connection errors...")
	err := eh.HandleConnectionError(ctx, fmt.Errorf("connection refused"))
	if err != nil {
		fmt.Printf("    Error handled: %v\n", err)
	}
	
	fmt.Println("    Testing timeout errors...")
	err = eh.HandleTimeoutError(ctx, fmt.Errorf("query timeout"))
	if err != nil {
		fmt.Printf("    Error handled: %v\n", err)
	}
	
	fmt.Println("    Testing deadlock errors...")
	err = eh.HandleDeadlockError(ctx, fmt.Errorf("deadlock detected"))
	if err != nil {
		fmt.Printf("    Error handled: %v\n", err)
	}
	
	fmt.Println("  Error handling completed")
}

// Example 15: Testing Strategies
func testingStrategies() {
	fmt.Println("\n15. Testing Strategies")
	fmt.Println("======================")
	
	// Create test suite
	ts := NewDatabaseTestSuite()
	
	fmt.Println("  Testing database testing strategies...")
	
	// Add test cases
	ts.AddTest("connection_test", func() error {
		fmt.Println("    Running connection test...")
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	
	ts.AddTest("transaction_test", func() error {
		fmt.Println("    Running transaction test...")
		time.Sleep(75 * time.Millisecond)
		return nil
	})
	
	ts.AddTest("concurrency_test", func() error {
		fmt.Println("    Running concurrency test...")
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	
	// Run tests
	fmt.Println("    Running test suite...")
	results := ts.RunTests()
	
	for testName, result := range results {
		if result.Error != nil {
			fmt.Printf("    %s: FAILED - %v\n", testName, result.Error)
		} else {
			fmt.Printf("    %s: PASSED (%v)\n", testName, result.Duration)
		}
	}
	
	fmt.Println("  Testing strategies completed")
}

// Example 16: Best Practices
func bestPractices() {
	fmt.Println("\n16. Best Practices")
	fmt.Println("==================")
	
	fmt.Println("  Database concurrency best practices:")
	fmt.Println("    1. Use connection pooling to manage database connections")
	fmt.Println("    2. Implement proper transaction management with rollback")
	fmt.Println("    3. Choose appropriate isolation levels for your use case")
	fmt.Println("    4. Implement deadlock prevention strategies")
	fmt.Println("    5. Use read replicas for read-heavy workloads")
	fmt.Println("    6. Implement proper error handling and retry logic")
	fmt.Println("    7. Monitor database performance and connection usage")
	fmt.Println("    8. Use prepared statements to prevent SQL injection")
	fmt.Println("    9. Implement proper logging and monitoring")
	fmt.Println("    10. Test concurrency scenarios thoroughly")
	
	fmt.Println("  Best practices demonstration completed")
}

// Example 17: Connection Pool Metrics
func connectionPoolMetrics() {
	fmt.Println("\n17. Connection Pool Metrics")
	fmt.Println("===========================")
	
	// Create connection pool with metrics
	pool, err := NewConnectionPool("user:password@tcp(localhost:3306)/testdb", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("  Error creating connection pool: %v\n", err)
		return
	}
	defer pool.Close()
	
	fmt.Println("  Testing connection pool metrics...")
	
	ctx := context.Background()
	
	// Simulate some activity
	for i := 0; i < 20; i++ {
		conn, err := pool.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
			continue
		}
		
		// Simulate work
		time.Sleep(10 * time.Millisecond)
		
		// Release connection
		pool.ReleaseConnection(conn)
	}
	
	// Get and display metrics
	stats := pool.GetStats()
	fmt.Printf("    Open connections: %d\n", stats.OpenConns)
	fmt.Printf("    In use: %d\n", stats.InUse)
	fmt.Printf("    Idle: %d\n", stats.Idle)
	fmt.Printf("    Wait count: %d\n", stats.WaitCount)
	fmt.Printf("    Wait duration: %v\n", stats.WaitDuration)
	fmt.Printf("    Max idle closed: %d\n", stats.MaxIdleClosed)
	fmt.Printf("    Max lifetime closed: %d\n", stats.MaxLifetimeClosed)
	
	fmt.Println("  Connection pool metrics completed")
}

// Example 18: Database Health Checks
func databaseHealthChecks() {
	fmt.Println("\n18. Database Health Checks")
	fmt.Println("==========================")
	
	// Create health checker
	hc := NewDatabaseHealthChecker()
	
	// Add databases to check
	hc.AddDatabase("primary", "localhost:3306")
	hc.AddDatabase("replica1", "localhost:3307")
	hc.AddDatabase("replica2", "localhost:3308")
	
	fmt.Println("  Testing database health checks...")
	
	// Start health checking
	go hc.StartHealthChecks()
	
	// Wait for some checks
	time.Sleep(200 * time.Millisecond)
	
	// Get health status
	status := hc.GetHealthStatus()
	for dbName, health := range status {
		fmt.Printf("    %s: %s\n", dbName, health)
	}
	
	fmt.Println("  Database health checks completed")
}

// Example 19: Database Load Balancing
func databaseLoadBalancing() {
	fmt.Println("\n19. Database Load Balancing")
	fmt.Println("===========================")
	
	// Create load balancer
	lb := NewDatabaseLoadBalancer()
	
	// Add database instances
	lb.AddInstance("db1", &MockDB{name: "db1"}, 1.0)
	lb.AddInstance("db2", &MockDB{name: "db2"}, 2.0)
	lb.AddInstance("db3", &MockDB{name: "db3"}, 1.5)
	
	fmt.Println("  Testing database load balancing...")
	
	_ = context.Background()
	
	// Test load balancing
	for i := 0; i < 10; i++ {
		instance, err := lb.GetInstance()
		if err != nil {
			fmt.Printf("    Error getting instance: %v\n", err)
			continue
		}
		
		fmt.Printf("    Request %d routed to: %s\n", i+1, instance.Name)
		
		// Simulate work
		time.Sleep(10 * time.Millisecond)
	}
	
	fmt.Println("  Database load balancing completed")
}

// Example 20: Database Caching
func databaseCaching() {
	fmt.Println("\n20. Database Caching")
	fmt.Println("===================")
	
	// Create database cache
	cache := NewDatabaseCache(100, 5*time.Minute)
	
	fmt.Println("  Testing database caching...")
	
	_ = context.Background()
	
	// Test cache operations
	fmt.Println("    Testing cache operations...")
	
	// Set some values
	cache.Set("user:1", "John Doe", 1*time.Minute)
	cache.Set("user:2", "Jane Smith", 1*time.Minute)
	cache.Set("user:3", "Bob Johnson", 1*time.Minute)
	
	// Get values
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("user:%d", i)
		value, found := cache.Get(key)
		if found {
			fmt.Printf("    Cache hit for %s: %s\n", key, value)
		} else {
			fmt.Printf("    Cache miss for %s\n", key)
		}
	}
	
	// Test cache statistics
	stats := cache.GetStats()
	fmt.Printf("    Cache stats - Hits: %d, Misses: %d, Size: %d\n", 
		stats.Hits, stats.Misses, stats.Size)
	
	fmt.Println("  Database caching completed")
}

// Run all basic examples
func runBasicExamples() {
	fmt.Println("🗄️ Database Concurrency Examples")
	fmt.Println("=================================")
	
	basicDatabaseConnection()
	connectionPooling()
	transactionManagement()
	acidProperties()
	isolationLevels()
	deadlockPrevention()
	readReplicas()
	databaseSharding()
	lockingStrategies()
	databaseConnectionManagement()
	concurrentQueryProcessing()
	databaseMonitoring()
	performanceOptimization()
	errorHandling()
	testingStrategies()
	bestPractices()
	connectionPoolMetrics()
	databaseHealthChecks()
	databaseLoadBalancing()
	databaseCaching()
}
