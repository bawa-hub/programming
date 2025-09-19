package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// SINGLETON PATTERN IMPLEMENTATIONS
// =============================================================================

// 1. EAGER INITIALIZATION SINGLETON
// Instance is created when the package is loaded
type EagerSingleton struct {
	ID        string
	CreatedAt time.Time
}

var eagerInstance = &EagerSingleton{
	ID:        "eager-singleton",
	CreatedAt: time.Now(),
}

func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}

// 2. LAZY INITIALIZATION SINGLETON (Not Thread-Safe)
type LazySingleton struct {
	ID        string
	CreatedAt time.Time
}

var lazyInstance *LazySingleton

func GetLazyInstance() *LazySingleton {
	if lazyInstance == nil {
		lazyInstance = &LazySingleton{
			ID:        "lazy-singleton",
			CreatedAt: time.Now(),
		}
	}
	return lazyInstance
}

// 3. THREAD-SAFE LAZY INITIALIZATION SINGLETON
type ThreadSafeSingleton struct {
	ID        string
	CreatedAt time.Time
}

var (
	threadSafeInstance *ThreadSafeSingleton
	mu                 sync.Mutex
)

func GetThreadSafeInstance() *ThreadSafeSingleton {
	if threadSafeInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if threadSafeInstance == nil {
			threadSafeInstance = &ThreadSafeSingleton{
				ID:        "thread-safe-singleton",
				CreatedAt: time.Now(),
			}
		}
	}
	return threadSafeInstance
}

// 4. BILL PUGH SOLUTION (Recommended)
// Uses sync.Once for thread-safe lazy initialization
type BillPughSingleton struct {
	ID        string
	CreatedAt time.Time
}

var (
	instance *BillPughSingleton
	once     sync.Once
)

func GetBillPughInstance() *BillPughSingleton {
	once.Do(func() {
		instance = &BillPughSingleton{
			ID:        "bill-pugh-singleton",
			CreatedAt: time.Now(),
		}
	})
	return instance
}

// =============================================================================
// REAL-WORLD SINGLETON EXAMPLES
// =============================================================================

// 1. DATABASE CONNECTION MANAGER
type DatabaseConnectionManager struct {
	connectionString string
	isConnected      bool
	connectionCount  int
}

var (
	dbManager *DatabaseConnectionManager
	dbOnce    sync.Once
)

func GetDatabaseManager() *DatabaseConnectionManager {
	dbOnce.Do(func() {
		dbManager = &DatabaseConnectionManager{
			connectionString: "localhost:5432/mydb",
			isConnected:      false,
			connectionCount:  0,
		}
	})
	return dbManager
}

func (dbm *DatabaseConnectionManager) Connect() error {
	if !dbm.isConnected {
		fmt.Println("Connecting to database...")
		time.Sleep(100 * time.Millisecond) // Simulate connection time
		dbm.isConnected = true
		dbm.connectionCount++
		fmt.Printf("Connected to database. Connection count: %d\n", dbm.connectionCount)
	}
	return nil
}

func (dbm *DatabaseConnectionManager) Disconnect() {
	if dbm.isConnected {
		fmt.Println("Disconnecting from database...")
		dbm.isConnected = false
		fmt.Printf("Disconnected from database. Connection count: %d\n", dbm.connectionCount)
	}
}

func (dbm *DatabaseConnectionManager) GetConnectionString() string {
	return dbm.connectionString
}

// 2. LOGGER SINGLETON
type Logger struct {
	logLevel string
	logs     []string
}

var (
	logger *Logger
	logOnce sync.Once
)

func GetLogger() *Logger {
	logOnce.Do(func() {
		logger = &Logger{
			logLevel: "INFO",
			logs:     make([]string, 0),
		}
	})
	return logger
}

func (l *Logger) SetLogLevel(level string) {
	l.logLevel = level
}

func (l *Logger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s", timestamp, l.logLevel, message)
	l.logs = append(l.logs, logEntry)
	fmt.Println(logEntry)
}

func (l *Logger) GetLogs() []string {
	return l.logs
}

// 3. CONFIGURATION MANAGER
type ConfigManager struct {
	config map[string]interface{}
}

var (
	configManager *ConfigManager
	configOnce    sync.Once
)

func GetConfigManager() *ConfigManager {
	configOnce.Do(func() {
		configManager = &ConfigManager{
			config: make(map[string]interface{}),
		}
		// Load default configuration
		configManager.config["app_name"] = "MyApp"
		configManager.config["version"] = "1.0.0"
		configManager.config["debug"] = true
		configManager.config["max_connections"] = 100
	})
	return configManager
}

func (cm *ConfigManager) Get(key string) (interface{}, bool) {
	value, exists := cm.config[key]
	return value, exists
}

func (cm *ConfigManager) Set(key string, value interface{}) {
	cm.config[key] = value
}

func (cm *ConfigManager) GetAll() map[string]interface{} {
	return cm.config
}

// 4. CACHE MANAGER
type CacheManager struct {
	cache map[string]interface{}
	mu    sync.RWMutex
}

var (
	cacheManager *CacheManager
	cacheOnce    sync.Once
)

func GetCacheManager() *CacheManager {
	cacheOnce.Do(func() {
		cacheManager = &CacheManager{
			cache: make(map[string]interface{}),
		}
	})
	return cacheManager
}

func (cm *CacheManager) Set(key string, value interface{}) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.cache[key] = value
}

func (cm *CacheManager) Get(key string) (interface{}, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, exists := cm.cache[key]
	return value, exists
}

func (cm *CacheManager) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.cache, key)
}

func (cm *CacheManager) Clear() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.cache = make(map[string]interface{})
}

func (cm *CacheManager) Size() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.cache)
}

// =============================================================================
// THREAD SAFETY TESTING
// =============================================================================

func testThreadSafety() {
	fmt.Println("\n=== THREAD SAFETY TEST ===")
	
	var wg sync.WaitGroup
	instances := make([]*BillPughSingleton, 100)
	
	// Create 100 goroutines to test thread safety
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = GetBillPughInstance()
		}(i)
	}
	
	wg.Wait()
	
	// Check if all instances are the same
	firstInstance := instances[0]
	allSame := true
	for i := 1; i < len(instances); i++ {
		if instances[i] != firstInstance {
			allSame = false
			break
		}
	}
	
	if allSame {
		fmt.Println("✅ Thread safety test passed - all instances are the same")
	} else {
		fmt.Println("❌ Thread safety test failed - instances are different")
	}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== SINGLETON PATTERN DEMONSTRATION ===\n")

	// 1. EAGER INITIALIZATION
	fmt.Println("1. EAGER INITIALIZATION:")
	instance1 := GetEagerInstance()
	instance2 := GetEagerInstance()
	fmt.Printf("Instance 1: %p, Created at: %v\n", instance1, instance1.CreatedAt)
	fmt.Printf("Instance 2: %p, Created at: %v\n", instance2, instance2.CreatedAt)
	fmt.Printf("Same instance: %t\n", instance1 == instance2)
	fmt.Println()

	// 2. LAZY INITIALIZATION
	fmt.Println("2. LAZY INITIALIZATION:")
	lazy1 := GetLazyInstance()
	lazy2 := GetLazyInstance()
	fmt.Printf("Lazy Instance 1: %p, Created at: %v\n", lazy1, lazy1.CreatedAt)
	fmt.Printf("Lazy Instance 2: %p, Created at: %v\n", lazy2, lazy2.CreatedAt)
	fmt.Printf("Same instance: %t\n", lazy1 == lazy2)
	fmt.Println()

	// 3. THREAD-SAFE LAZY INITIALIZATION
	fmt.Println("3. THREAD-SAFE LAZY INITIALIZATION:")
	threadSafe1 := GetThreadSafeInstance()
	threadSafe2 := GetThreadSafeInstance()
	fmt.Printf("Thread-Safe Instance 1: %p, Created at: %v\n", threadSafe1, threadSafe1.CreatedAt)
	fmt.Printf("Thread-Safe Instance 2: %p, Created at: %v\n", threadSafe2, threadSafe2.CreatedAt)
	fmt.Printf("Same instance: %t\n", threadSafe1 == threadSafe2)
	fmt.Println()

	// 4. BILL PUGH SOLUTION
	fmt.Println("4. BILL PUGH SOLUTION:")
	billPugh1 := GetBillPughInstance()
	billPugh2 := GetBillPughInstance()
	fmt.Printf("Bill Pugh Instance 1: %p, Created at: %v\n", billPugh1, billPugh1.CreatedAt)
	fmt.Printf("Bill Pugh Instance 2: %p, Created at: %v\n", billPugh2, billPugh2.CreatedAt)
	fmt.Printf("Same instance: %t\n", billPugh1 == billPugh2)
	fmt.Println()

	// 5. REAL-WORLD EXAMPLES
	fmt.Println("5. REAL-WORLD EXAMPLES:")

	// Database Manager
	fmt.Println("Database Manager:")
	db1 := GetDatabaseManager()
	db2 := GetDatabaseManager()
	fmt.Printf("Same DB manager: %t\n", db1 == db2)
	db1.Connect()
	db2.Connect() // Should not create new connection
	db1.Disconnect()
	fmt.Println()

	// Logger
	fmt.Println("Logger:")
	logger1 := GetLogger()
	logger2 := GetLogger()
	fmt.Printf("Same logger: %t\n", logger1 == logger2)
	logger1.SetLogLevel("DEBUG")
	logger1.Log("This is a debug message")
	logger2.Log("This is an info message")
	fmt.Printf("Total logs: %d\n", len(logger1.GetLogs()))
	fmt.Println()

	// Configuration Manager
	fmt.Println("Configuration Manager:")
	config1 := GetConfigManager()
	config2 := GetConfigManager()
	fmt.Printf("Same config manager: %t\n", config1 == config2)
	appName, _ := config1.Get("app_name")
	fmt.Printf("App name: %v\n", appName)
	config1.Set("new_setting", "new_value")
	newSetting, _ := config2.Get("new_setting")
	fmt.Printf("New setting: %v\n", newSetting)
	fmt.Println()

	// Cache Manager
	fmt.Println("Cache Manager:")
	cache1 := GetCacheManager()
	cache2 := GetCacheManager()
	fmt.Printf("Same cache manager: %t\n", cache1 == cache2)
	cache1.Set("key1", "value1")
	cache1.Set("key2", "value2")
	value, exists := cache2.Get("key1")
	fmt.Printf("Key1 exists: %t, Value: %v\n", exists, value)
	fmt.Printf("Cache size: %d\n", cache1.Size())
	fmt.Println()

	// 6. THREAD SAFETY TEST
	testThreadSafety()

	fmt.Println("\n=== END OF DEMONSTRATION ===")
}
