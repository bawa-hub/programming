package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 1: Multi-Database Connection Pool
type MultiDatabasePool struct {
	pools    map[string]*ConnectionPool
	balancer *DatabaseLoadBalancer
	mutex    sync.RWMutex
}

func NewMultiDatabasePool() *MultiDatabasePool {
	return &MultiDatabasePool{
		pools:    make(map[string]*ConnectionPool),
		balancer: NewDatabaseLoadBalancer(),
	}
}

func (mdp *MultiDatabasePool) AddDatabase(name, dsn string, maxConns, maxIdle int, maxLifetime time.Duration) error {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return err
	}
	
	mdp.mutex.Lock()
	defer mdp.mutex.Unlock()
	
	mdp.pools[name] = pool
	mdp.balancer.AddInstance(name, pool, 1.0)
	
	return nil
}

func (mdp *MultiDatabasePool) GetConnection(ctx context.Context, dbName string) (*MockConn, error) {
	mdp.mutex.RLock()
	pool, exists := mdp.pools[dbName]
	mdp.mutex.RUnlock()
	
	if !exists {
		return nil, fmt.Errorf("database %s not found", dbName)
	}
	
	return pool.GetConnection(ctx)
}

func (mdp *MultiDatabasePool) GetBalancedConnection(ctx context.Context) (*MockConn, string, error) {
	instance, err := mdp.balancer.GetInstance()
	if err != nil {
		return nil, "", err
	}
	
	conn, err := mdp.GetConnection(ctx, instance.Name)
	return conn, instance.Name, err
}

func (mdp *MultiDatabasePool) Close() error {
	mdp.mutex.Lock()
	defer mdp.mutex.Unlock()
	
	for _, pool := range mdp.pools {
		pool.Close()
	}
	
	return nil
}

// Advanced Pattern 2: Distributed Transaction Manager
type DistributedTransactionManager struct {
	databases map[string]*DatabaseManager
	coordinator *TransactionCoordinator
	mutex     sync.RWMutex
}

type TransactionCoordinator struct {
	transactions map[string]*DistributedTransaction
	mutex        sync.RWMutex
}

type DistributedTransaction struct {
	ID        string
	Databases []string
	Status    TransactionStatus
	StartTime time.Time
}

type TransactionStatus int

const (
	StatusPrepared TransactionStatus = iota
	StatusCommitted
	StatusAborted
)

func NewDistributedTransactionManager() *DistributedTransactionManager {
	return &DistributedTransactionManager{
		databases: make(map[string]*DatabaseManager),
		coordinator: &TransactionCoordinator{
			transactions: make(map[string]*DistributedTransaction),
		},
	}
}

func (dtm *DistributedTransactionManager) AddDatabase(name string, db *DatabaseManager) {
	dtm.mutex.Lock()
	defer dtm.mutex.Unlock()
	dtm.databases[name] = db
}

func (dtm *DistributedTransactionManager) BeginDistributedTransaction(ctx context.Context, txID string, databases []string) (*DistributedTransaction, error) {
	dtm.coordinator.mutex.Lock()
	defer dtm.coordinator.mutex.Unlock()
	
	// Check if all databases are available
	for _, dbName := range databases {
		if _, exists := dtm.databases[dbName]; !exists {
			return nil, fmt.Errorf("database %s not found", dbName)
		}
	}
	
	tx := &DistributedTransaction{
		ID:        txID,
		Databases: databases,
		Status:    StatusPrepared,
		StartTime: time.Now(),
	}
	
	dtm.coordinator.transactions[txID] = tx
	
	// Begin transaction on all databases
	for _, dbName := range databases {
		db := dtm.databases[dbName]
		err := db.ExecuteTransaction(ctx, func(tx interface{}) error {
			// Simulate transaction preparation
			return nil
		})
		if err != nil {
			// Rollback already prepared transactions
			dtm.rollbackDistributedTransaction(txID)
			return nil, err
		}
	}
	
	return tx, nil
}

func (dtm *DistributedTransactionManager) CommitDistributedTransaction(ctx context.Context, txID string) error {
	dtm.coordinator.mutex.Lock()
	defer dtm.coordinator.mutex.Unlock()
	
	tx, exists := dtm.coordinator.transactions[txID]
	if !exists {
		return fmt.Errorf("transaction %s not found", txID)
	}
	
	// Commit on all databases
	for _, dbName := range tx.Databases {
		db := dtm.databases[dbName]
		err := db.ExecuteTransaction(ctx, func(tx interface{}) error {
			// Simulate transaction commit
			return nil
		})
		if err != nil {
			// Rollback on failure
			dtm.rollbackDistributedTransaction(txID)
			return err
		}
	}
	
	tx.Status = StatusCommitted
	delete(dtm.coordinator.transactions, txID)
	
	return nil
}

func (dtm *DistributedTransactionManager) rollbackDistributedTransaction(txID string) {
	tx, exists := dtm.coordinator.transactions[txID]
	if !exists {
		return
	}
	
	// Rollback on all databases
	for _, dbName := range tx.Databases {
		if db, exists := dtm.databases[dbName]; exists {
			db.ExecuteTransaction(context.Background(), func(tx interface{}) error {
				// Simulate transaction rollback
				return nil
			})
		}
	}
	
	tx.Status = StatusAborted
	delete(dtm.coordinator.transactions, txID)
}

// Advanced Pattern 3: Database Circuit Breaker
type DatabaseCircuitBreaker struct {
	db           interface{}
	state        CircuitState
	failureCount int
	threshold    int
	timeout      time.Duration
	lastFailure  time.Time
	mutex        sync.RWMutex
}

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

func NewDatabaseCircuitBreaker(db interface{}, threshold int, timeout time.Duration) *DatabaseCircuitBreaker {
	return &DatabaseCircuitBreaker{
		db:        db,
		state:     StateClosed,
		threshold: threshold,
		timeout:   timeout,
	}
}

func (dcb *DatabaseCircuitBreaker) Execute(ctx context.Context, operation func() error) error {
	dcb.mutex.Lock()
	defer dcb.mutex.Unlock()
	
	// Check if circuit should be opened
	if dcb.state == StateClosed && dcb.failureCount >= dcb.threshold {
		dcb.state = StateOpen
		dcb.lastFailure = time.Now()
	}
	
	// Check if circuit should be half-opened
	if dcb.state == StateOpen && time.Since(dcb.lastFailure) > dcb.timeout {
		dcb.state = StateHalfOpen
	}
	
	// Execute operation based on state
	switch dcb.state {
	case StateClosed, StateHalfOpen:
		err := operation()
		if err != nil {
			dcb.failureCount++
			dcb.lastFailure = time.Now()
			if dcb.state == StateHalfOpen {
				dcb.state = StateOpen
			}
			return err
		}
		
		dcb.failureCount = 0
		if dcb.state == StateHalfOpen {
			dcb.state = StateClosed
		}
		return nil
		
	case StateOpen:
		return fmt.Errorf("circuit breaker is open")
		
	default:
		return fmt.Errorf("unknown circuit breaker state")
	}
}

// Advanced Pattern 4: Database Connection Pool with Metrics
type MetricsConnectionPool struct {
	*ConnectionPool
	metrics *PoolMetrics
	mutex   sync.RWMutex
}

type PoolMetrics struct {
	TotalRequests    int64
	SuccessfulRequests int64
	FailedRequests   int64
	AvgResponseTime  time.Duration
	MaxResponseTime  time.Duration
	MinResponseTime  time.Duration
}

func NewMetricsConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*MetricsConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	mcp := &MetricsConnectionPool{
		ConnectionPool: pool,
		metrics:        &PoolMetrics{},
	}
	
	// Start metrics collection
	go mcp.collectMetrics()
	
	return mcp, nil
}

func (mcp *MetricsConnectionPool) GetConnectionWithMetrics(ctx context.Context) (*MockConn, error) {
	start := time.Now()
	
	conn, err := mcp.GetConnection(ctx)
	
	mcp.mutex.Lock()
	defer mcp.mutex.Unlock()
	
	mcp.metrics.TotalRequests++
	
	if err != nil {
		mcp.metrics.FailedRequests++
	} else {
		mcp.metrics.SuccessfulRequests++
	}
	
	responseTime := time.Since(start)
	mcp.updateResponseTime(responseTime)
	
	return conn, err
}

func (mcp *MetricsConnectionPool) updateResponseTime(responseTime time.Duration) {
	if mcp.metrics.AvgResponseTime == 0 {
		mcp.metrics.AvgResponseTime = responseTime
	} else {
		mcp.metrics.AvgResponseTime = (mcp.metrics.AvgResponseTime + responseTime) / 2
	}
	
	if responseTime > mcp.metrics.MaxResponseTime {
		mcp.metrics.MaxResponseTime = responseTime
	}
	
	if mcp.metrics.MinResponseTime == 0 || responseTime < mcp.metrics.MinResponseTime {
		mcp.metrics.MinResponseTime = responseTime
	}
}

func (mcp *MetricsConnectionPool) collectMetrics() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		mcp.mutex.RLock()
		metrics := *mcp.metrics
		mcp.mutex.RUnlock()
		
		fmt.Printf("    Pool metrics - Total: %d, Success: %d, Failed: %d, Avg: %v\n",
			metrics.TotalRequests, metrics.SuccessfulRequests, metrics.FailedRequests, metrics.AvgResponseTime)
	}
}

func (mcp *MetricsConnectionPool) GetMetrics() *PoolMetrics {
	mcp.mutex.RLock()
	defer mcp.mutex.RUnlock()
	
	metrics := *mcp.metrics
	return &metrics
}

// Advanced Pattern 5: Database Connection Pool with Health Checks
type HealthCheckConnectionPool struct {
	*ConnectionPool
	healthChecker *PoolHealthChecker
	healthy       bool
	mutex         sync.RWMutex
}

type PoolHealthChecker struct {
	interval time.Duration
	timeout  time.Duration
	stopCh   chan struct{}
}

func NewHealthCheckConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*HealthCheckConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	hcp := &HealthCheckConnectionPool{
		ConnectionPool: pool,
		healthChecker: &PoolHealthChecker{
			interval: 30 * time.Second,
			timeout:  5 * time.Second,
			stopCh:   make(chan struct{}),
		},
		healthy: true,
	}
	
	// Start health checking
	go hcp.healthChecker.Start(hcp)
	
	return hcp, nil
}

func (hc *PoolHealthChecker) Start(pool *HealthCheckConnectionPool) {
	ticker := time.NewTicker(hc.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			pool.checkHealth()
		case <-hc.stopCh:
			return
		}
	}
}

func (hcp *HealthCheckConnectionPool) checkHealth() {
	_, cancel := context.WithTimeout(context.Background(), hcp.healthChecker.timeout)
	defer cancel()
	
	// Simulate health check
	time.Sleep(10 * time.Millisecond)
	
	hcp.mutex.Lock()
	hcp.healthy = true
	hcp.mutex.Unlock()
}

func (hcp *HealthCheckConnectionPool) GetConnection(ctx context.Context) (*MockConn, error) {
	hcp.mutex.RLock()
	healthy := hcp.healthy
	hcp.mutex.RUnlock()
	
	if !healthy {
		return nil, fmt.Errorf("connection pool is unhealthy")
	}
	
	return hcp.ConnectionPool.GetConnection(ctx)
}

func (hcp *HealthCheckConnectionPool) IsHealthy() bool {
	hcp.mutex.RLock()
	defer hcp.mutex.RUnlock()
	return hcp.healthy
}

// Advanced Pattern 6: Database Connection Pool with Load Balancing
type LoadBalancedConnectionPool struct {
	pools    []*ConnectionPool
	balancer *RoundRobinBalancer
	mutex    sync.RWMutex
}

type RoundRobinBalancer struct {
	current int
	mutex   sync.Mutex
}

func NewLoadBalancedConnectionPool(pools []*ConnectionPool) *LoadBalancedConnectionPool {
	return &LoadBalancedConnectionPool{
		pools:    pools,
		balancer: &RoundRobinBalancer{},
	}
}

func (lbcp *LoadBalancedConnectionPool) GetConnection(ctx context.Context) (*MockConn, error) {
	lbcp.mutex.RLock()
	defer lbcp.mutex.RUnlock()
	
	if len(lbcp.pools) == 0 {
		return nil, fmt.Errorf("no connection pools available")
	}
	
	pool := lbcp.balancer.SelectPool(lbcp.pools)
	return pool.GetConnection(ctx)
}

func (rrb *RoundRobinBalancer) SelectPool(pools []*ConnectionPool) *ConnectionPool {
	rrb.mutex.Lock()
	defer rrb.mutex.Unlock()
	
	pool := pools[rrb.current%len(pools)]
	rrb.current++
	return pool
}

// Advanced Pattern 7: Database Connection Pool with Caching
type CachedConnectionPool struct {
	*ConnectionPool
	cache   *DatabaseCache
	mutex   sync.RWMutex
}

func NewCachedConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*CachedConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	ccp := &CachedConnectionPool{
		ConnectionPool: pool,
		cache:          NewDatabaseCache(1000, 5*time.Minute),
	}
	
	return ccp, nil
}

func (ccp *CachedConnectionPool) ExecuteCachedQuery(ctx context.Context, query string, args ...interface{}) (string, error) {
	// Create cache key
	cacheKey := fmt.Sprintf("%s:%v", query, args)
	
	// Try to get from cache
	ccp.mutex.RLock()
	cached, found := ccp.cache.Get(cacheKey)
	ccp.mutex.RUnlock()
	
	if found {
		return cached, nil
	}
	
	// Execute query
	conn, err := ccp.GetConnection(ctx)
	if err != nil {
		return "", err
	}
	defer ccp.ReleaseConnection(conn)
	
	// Simulate query execution
	time.Sleep(50 * time.Millisecond)
	result := fmt.Sprintf("Result for query: %s", query)
	
	// Cache result
	ccp.mutex.Lock()
	ccp.cache.Set(cacheKey, result, 5*time.Minute)
	ccp.mutex.Unlock()
	
	return result, nil
}

// Advanced Pattern 8: Database Connection Pool with Retry Logic
type RetryConnectionPool struct {
	*ConnectionPool
	maxRetries int
	retryDelay time.Duration
}

func NewRetryConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration, maxRetries int, retryDelay time.Duration) (*RetryConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	rcp := &RetryConnectionPool{
		ConnectionPool: pool,
		maxRetries:     maxRetries,
		retryDelay:     retryDelay,
	}
	
	return rcp, nil
}

func (rcp *RetryConnectionPool) GetConnectionWithRetry(ctx context.Context) (*MockConn, error) {
	var lastErr error
	
	for i := 0; i <= rcp.maxRetries; i++ {
		conn, err := rcp.GetConnection(ctx)
		if err == nil {
			return conn, nil
		}
		
		lastErr = err
		
		if i < rcp.maxRetries {
			time.Sleep(rcp.retryDelay)
		}
	}
	
	return nil, fmt.Errorf("failed to get connection after %d retries: %v", rcp.maxRetries, lastErr)
}

// Advanced Pattern 9: Database Connection Pool with Monitoring
type MonitoredConnectionPool struct {
	*ConnectionPool
	monitor *PoolMonitor
	mutex   sync.RWMutex
}

type PoolMonitor struct {
	alerts    []Alert
	threshold int
	mutex     sync.RWMutex
}

type Alert struct {
	Type      string
	Message   string
	Timestamp time.Time
}

func NewMonitoredConnectionPool(dsn string, maxConns, maxIdle int, maxLifetime time.Duration) (*MonitoredConnectionPool, error) {
	pool, err := NewConnectionPool(dsn, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	mcp := &MonitoredConnectionPool{
		ConnectionPool: pool,
		monitor: &PoolMonitor{
			alerts:    make([]Alert, 0),
			threshold: 80, // 80% utilization threshold
		},
	}
	
	// Start monitoring
	go mcp.startMonitoring()
	
	return mcp, nil
}

func (mcp *MonitoredConnectionPool) startMonitoring() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		mcp.checkPoolHealth()
	}
}

func (mcp *MonitoredConnectionPool) checkPoolHealth() {
	stats := mcp.GetStats()
	
	// Check utilization
	utilization := float64(stats.InUse) / float64(stats.OpenConns) * 100
	
	if utilization > float64(mcp.monitor.threshold) {
		alert := Alert{
			Type:      "HIGH_UTILIZATION",
			Message:   fmt.Sprintf("Pool utilization is %.2f%%", utilization),
			Timestamp: time.Now(),
		}
		
		mcp.monitor.mutex.Lock()
		mcp.monitor.alerts = append(mcp.monitor.alerts, alert)
		mcp.monitor.mutex.Unlock()
		
		fmt.Printf("    Alert: %s - %s\n", alert.Type, alert.Message)
	}
}

func (mcp *MonitoredConnectionPool) GetAlerts() []Alert {
	mcp.monitor.mutex.RLock()
	defer mcp.monitor.mutex.RUnlock()
	
	alerts := make([]Alert, len(mcp.monitor.alerts))
	copy(alerts, mcp.monitor.alerts)
	return alerts
}

// Advanced Pattern 10: Database Connection Pool with Failover
type FailoverConnectionPool struct {
	primary   *ConnectionPool
	secondary *ConnectionPool
	current   *ConnectionPool
	mutex     sync.RWMutex
}

func NewFailoverConnectionPool(primaryDSN, secondaryDSN string, maxConns, maxIdle int, maxLifetime time.Duration) (*FailoverConnectionPool, error) {
	primary, err := NewConnectionPool(primaryDSN, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	secondary, err := NewConnectionPool(secondaryDSN, maxConns, maxIdle, maxLifetime)
	if err != nil {
		return nil, err
	}
	
	fcp := &FailoverConnectionPool{
		primary:   primary,
		secondary: secondary,
		current:   primary,
	}
	
	// Start failover monitoring
	go fcp.monitorFailover()
	
	return fcp, nil
}

func (fcp *FailoverConnectionPool) GetConnection(ctx context.Context) (*MockConn, error) {
	fcp.mutex.RLock()
	current := fcp.current
	fcp.mutex.RUnlock()
	
	conn, err := current.GetConnection(ctx)
	if err != nil {
		// Try failover
		fcp.mutex.Lock()
		if fcp.current == fcp.primary {
			fcp.current = fcp.secondary
		} else {
			fcp.current = fcp.primary
		}
		fcp.mutex.Unlock()
		
		// Retry with new current
		fcp.mutex.RLock()
		current = fcp.current
		fcp.mutex.RUnlock()
		
		return current.GetConnection(ctx)
	}
	
	return conn, nil
}

func (fcp *FailoverConnectionPool) monitorFailover() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		fcp.checkPrimaryHealth()
	}
}

func (fcp *FailoverConnectionPool) checkPrimaryHealth() {
	// Simulate health check
	time.Sleep(10 * time.Millisecond)
	
	// If primary is healthy, switch back to it
	fcp.mutex.Lock()
	if fcp.current == fcp.secondary {
		fcp.current = fcp.primary
		fmt.Println("    Switched back to primary database")
	}
	fcp.mutex.Unlock()
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Database Concurrency Patterns")
	fmt.Println("========================================")
	
	// Pattern 1: Multi-Database Connection Pool
	fmt.Println("\n1. Multi-Database Connection Pool")
	mdp := NewMultiDatabasePool()
	mdp.AddDatabase("db1", "user:pass@tcp(localhost:3306)/db1", 10, 5, time.Hour)
	mdp.AddDatabase("db2", "user:pass@tcp(localhost:3307)/db2", 10, 5, time.Hour)
	
	ctx := context.Background()
	conn, dbName, err := mdp.GetBalancedConnection(ctx)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    Got connection from %s: %s\n", dbName, conn.id)
	}
	
	// Pattern 2: Distributed Transaction Manager
	fmt.Println("\n2. Distributed Transaction Manager")
	dtm := NewDistributedTransactionManager()
	dtm.AddDatabase("db1", &DatabaseManager{})
	dtm.AddDatabase("db2", &DatabaseManager{})
	
	tx, err := dtm.BeginDistributedTransaction(ctx, "dist-tx-1", []string{"db1", "db2"})
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    Distributed transaction %s begun\n", tx.ID)
		err = dtm.CommitDistributedTransaction(ctx, tx.ID)
		if err != nil {
			fmt.Printf("    Error: %v\n", err)
		} else {
			fmt.Printf("    Distributed transaction %s committed\n", tx.ID)
		}
	}
	
	// Pattern 3: Database Circuit Breaker
	fmt.Println("\n3. Database Circuit Breaker")
	dcb := NewDatabaseCircuitBreaker(nil, 3, 5*time.Second)
	
	for i := 0; i < 5; i++ {
		err := dcb.Execute(ctx, func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		if err != nil {
			fmt.Printf("    Operation %d failed: %v\n", i+1, err)
		} else {
			fmt.Printf("    Operation %d succeeded\n", i+1)
		}
	}
	
	// Pattern 4: Database Connection Pool with Metrics
	fmt.Println("\n4. Database Connection Pool with Metrics")
	mcp, err := NewMetricsConnectionPool("user:pass@tcp(localhost:3306)/db", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		for i := 0; i < 5; i++ {
			_, err := mcp.GetConnectionWithMetrics(ctx)
			if err != nil {
				fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
			}
		}
		
		metrics := mcp.GetMetrics()
		fmt.Printf("    Metrics - Total: %d, Success: %d, Failed: %d\n",
			metrics.TotalRequests, metrics.SuccessfulRequests, metrics.FailedRequests)
	}
	
	// Pattern 5: Database Connection Pool with Health Checks
	fmt.Println("\n5. Database Connection Pool with Health Checks")
	hcp, err := NewHealthCheckConnectionPool("user:pass@tcp(localhost:3306)/db", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		healthy := hcp.IsHealthy()
		fmt.Printf("    Pool is healthy: %t\n", healthy)
	}
	
	// Pattern 6: Database Connection Pool with Load Balancing
	fmt.Println("\n6. Database Connection Pool with Load Balancing")
	pools := []*ConnectionPool{
		&ConnectionPool{db: &MockDB{name: "pool1"}},
		&ConnectionPool{db: &MockDB{name: "pool2"}},
		&ConnectionPool{db: &MockDB{name: "pool3"}},
	}
	
	lbcp := NewLoadBalancedConnectionPool(pools)
	for i := 0; i < 5; i++ {
		conn, err := lbcp.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection %d: %v\n", i+1, err)
		} else {
			fmt.Printf("    Got connection %d: %s\n", i+1, conn.id)
		}
	}
	
	// Pattern 7: Database Connection Pool with Caching
	fmt.Println("\n7. Database Connection Pool with Caching")
	ccp, err := NewCachedConnectionPool("user:pass@tcp(localhost:3306)/db", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		result, err := ccp.ExecuteCachedQuery(ctx, "SELECT * FROM users WHERE id = ?", 1)
		if err != nil {
			fmt.Printf("    Error executing query: %v\n", err)
		} else {
			fmt.Printf("    Query result: %s\n", result)
		}
	}
	
	// Pattern 8: Database Connection Pool with Retry Logic
	fmt.Println("\n8. Database Connection Pool with Retry Logic")
	rcp, err := NewRetryConnectionPool("user:pass@tcp(localhost:3306)/db", 10, 5, time.Hour, 3, time.Second)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		conn, err := rcp.GetConnectionWithRetry(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection: %v\n", err)
		} else {
			fmt.Printf("    Got connection with retry: %s\n", conn.id)
		}
	}
	
	// Pattern 9: Database Connection Pool with Monitoring
	fmt.Println("\n9. Database Connection Pool with Monitoring")
	mcp2, err := NewMonitoredConnectionPool("user:pass@tcp(localhost:3306)/db", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		time.Sleep(100 * time.Millisecond) // Let monitoring run
		alerts := mcp2.GetAlerts()
		fmt.Printf("    Generated %d alerts\n", len(alerts))
	}
	
	// Pattern 10: Database Connection Pool with Failover
	fmt.Println("\n10. Database Connection Pool with Failover")
	fcp, err := NewFailoverConnectionPool("user:pass@tcp(localhost:3306)/primary", "user:pass@tcp(localhost:3307)/secondary", 10, 5, time.Hour)
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		conn, err := fcp.GetConnection(ctx)
		if err != nil {
			fmt.Printf("    Error getting connection: %v\n", err)
		} else {
			fmt.Printf("    Got connection with failover: %s\n", conn.id)
		}
	}
	
	fmt.Println("\n🎉 All advanced patterns completed!")
}
