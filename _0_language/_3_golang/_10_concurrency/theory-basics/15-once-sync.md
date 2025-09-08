# 15 - sync.Once

## 🎯 Learning Objectives
- Understand what sync.Once is and when to use it
- Learn how to use sync.Once for one-time initialization
- Master sync.Once patterns and best practices
- Practice with different sync.Once scenarios
- Understand sync.Once performance characteristics

## 📚 Theory

### What is sync.Once?

**sync.Once** is a synchronization primitive that ensures a function is executed only once, regardless of how many times it's called.

**Key characteristics:**
- **One-time execution**: Function runs only once
- **Thread-safe**: Can be called from multiple goroutines
- **Blocking**: Other goroutines wait until first call completes
- **No parameters**: Function must take no parameters

### Why Do We Need sync.Once?

**Problem**: Multiple goroutines might try to initialize the same resource simultaneously.

**Solution**: sync.Once ensures initialization happens only once, safely.

### Common Use Cases

1. **Singleton pattern**: Ensure only one instance exists
2. **Lazy initialization**: Initialize only when needed
3. **Resource setup**: Set up resources once
4. **Configuration loading**: Load configuration once

## 💻 Code Examples

### Example 1: Basic sync.Once Usage

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func basicSyncOnce() {
    fmt.Println("=== Basic sync.Once Usage ===")
    
    var once sync.Once
    var initialized bool
    
    // Function to initialize
    initFunc := func() {
        fmt.Println("Initializing...")
        time.Sleep(1 * time.Second)
        initialized = true
        fmt.Println("Initialization complete")
    }
    
    // Multiple goroutines try to initialize
    for i := 1; i <= 5; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d: calling Do\n", id)
            once.Do(initFunc)
            fmt.Printf("Goroutine %d: Do completed\n", id)
        }(i)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("Initialized: %v\n", initialized)
}

func main() {
    basicSyncOnce()
}
```

**Run this code:**
```bash
go run 15-once-sync.go
```

### Example 2: Singleton Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Database struct {
    connection string
}

var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        fmt.Println("Creating database instance...")
        time.Sleep(500 * time.Millisecond)
        instance = &Database{
            connection: "database_connection_string",
        }
        fmt.Println("Database instance created")
    })
    return instance
}

func singletonPattern() {
    fmt.Println("=== Singleton Pattern ===")
    
    // Multiple goroutines try to get database instance
    for i := 1; i <= 3; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d: getting database\n", id)
            db := GetDatabase()
            fmt.Printf("Goroutine %d: got database with connection: %s\n", id, db.connection)
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}

func main() {
    singletonPattern()
}
```

### Example 3: Lazy Initialization

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    data map[string]string
    once sync.Once
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) initialize() {
    fmt.Println("Initializing cache...")
    time.Sleep(500 * time.Millisecond)
    c.data["key1"] = "value1"
    c.data["key2"] = "value2"
    c.data["key3"] = "value3"
    fmt.Println("Cache initialized")
}

func (c *Cache) Get(key string) (string, bool) {
    c.once.Do(c.initialize)
    value, exists := c.data[key]
    return value, exists
}

func lazyInitialization() {
    fmt.Println("=== Lazy Initialization ===")
    
    cache := NewCache()
    
    // Multiple goroutines try to get values
    for i := 1; i <= 3; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d: getting key1\n", id)
            if value, exists := cache.Get("key1"); exists {
                fmt.Printf("Goroutine %d: got key1 = %s\n", id, value)
            }
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}

func main() {
    lazyInitialization()
}
```

### Example 4: Resource Setup

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type ResourceManager struct {
    resources map[string]string
    once      sync.Once
}

func NewResourceManager() *ResourceManager {
    return &ResourceManager{
        resources: make(map[string]string),
    }
}

func (rm *ResourceManager) setup() {
    fmt.Println("Setting up resources...")
    time.Sleep(1 * time.Second)
    rm.resources["resource1"] = "ready"
    rm.resources["resource2"] = "ready"
    rm.resources["resource3"] = "ready"
    fmt.Println("Resources setup complete")
}

func (rm *ResourceManager) UseResource(name string) bool {
    rm.once.Do(rm.setup)
    
    if status, exists := rm.resources[name]; exists {
        fmt.Printf("Using resource %s (status: %s)\n", name, status)
        return true
    }
    fmt.Printf("Resource %s not found\n", name)
    return false
}

func resourceSetup() {
    fmt.Println("=== Resource Setup ===")
    
    rm := NewResourceManager()
    
    // Multiple goroutines try to use resources
    for i := 1; i <= 3; i++ {
        go func(id int) {
            resourceName := fmt.Sprintf("resource%d", id)
            rm.UseResource(resourceName)
        }(i)
    }
    
    time.Sleep(2 * time.Second)
}

func main() {
    resourceSetup()
}
```

### Example 5: Configuration Loading

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Config struct {
    DatabaseURL string
    APIKey      string
    DebugMode   bool
}

var (
    config *Config
    once   sync.Once
)

func LoadConfig() *Config {
    once.Do(func() {
        fmt.Println("Loading configuration...")
        time.Sleep(800 * time.Millisecond)
        config = &Config{
            DatabaseURL: "postgres://localhost:5432/mydb",
            APIKey:      "secret-api-key",
            DebugMode:   true,
        }
        fmt.Println("Configuration loaded")
    })
    return config
}

func configurationLoading() {
    fmt.Println("=== Configuration Loading ===")
    
    // Multiple goroutines try to load config
    for i := 1; i <= 3; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d: loading config\n", id)
            cfg := LoadConfig()
            fmt.Printf("Goroutine %d: got config - DB: %s, Debug: %v\n", 
                id, cfg.DatabaseURL, cfg.DebugMode)
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}

func main() {
    configurationLoading()
}
```

### Example 6: sync.Once with Error Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Service struct {
    initialized bool
    error       error
    once        sync.Once
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) initialize() {
    fmt.Println("Initializing service...")
    time.Sleep(500 * time.Millisecond)
    
    // Simulate initialization that might fail
    if time.Now().Unix()%2 == 0 {
        s.error = fmt.Errorf("initialization failed")
        fmt.Println("Service initialization failed")
        return
    }
    
    s.initialized = true
    fmt.Println("Service initialized successfully")
}

func (s *Service) GetStatus() (bool, error) {
    s.once.Do(s.initialize)
    return s.initialized, s.error
}

func syncOnceWithErrorHandling() {
    fmt.Println("=== sync.Once with Error Handling ===")
    
    service := NewService()
    
    // Multiple goroutines try to get status
    for i := 1; i <= 3; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d: getting status\n", id)
            initialized, err := service.GetStatus()
            if err != nil {
                fmt.Printf("Goroutine %d: error - %v\n", id, err)
            } else {
                fmt.Printf("Goroutine %d: initialized = %v\n", id, initialized)
            }
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}

func main() {
    syncOnceWithErrorHandling()
}
```

### Example 7: sync.Once Performance

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func syncOncePerformance() {
    fmt.Println("=== sync.Once Performance ===")
    
    const iterations = 1000000
    
    // Test with sync.Once
    start := time.Now()
    var once sync.Once
    var counter int
    
    for i := 0; i < iterations; i++ {
        once.Do(func() {
            counter++
        })
    }
    
    syncOnceTime := time.Since(start)
    
    // Test with mutex
    start = time.Now()
    var mutex sync.Mutex
    var initialized bool
    counter = 0
    
    for i := 0; i < iterations; i++ {
        mutex.Lock()
        if !initialized {
            counter++
            initialized = true
        }
        mutex.Unlock()
    }
    
    mutexTime := time.Since(start)
    
    fmt.Printf("sync.Once time: %v\n", syncOnceTime)
    fmt.Printf("Mutex time: %v\n", mutexTime)
    fmt.Printf("sync.Once speedup: %.2fx\n", float64(mutexTime)/float64(syncOnceTime))
}

func main() {
    syncOncePerformance()
}
```

### Example 8: sync.Once Best Practices

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Logger struct {
    level string
    once  sync.Once
}

func NewLogger() *Logger {
    return &Logger{}
}

func (l *Logger) setup() {
    fmt.Println("Setting up logger...")
    time.Sleep(300 * time.Millisecond)
    l.level = "INFO"
    fmt.Println("Logger setup complete")
}

func (l *Logger) Log(message string) {
    l.once.Do(l.setup)
    fmt.Printf("[%s] %s\n", l.level, message)
}

func syncOnceBestPractices() {
    fmt.Println("=== sync.Once Best Practices ===")
    
    logger := NewLogger()
    
    // Multiple goroutines use logger
    for i := 1; i <= 3; i++ {
        go func(id int) {
            logger.Log(fmt.Sprintf("Message from goroutine %d", id))
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}

func main() {
    syncOnceBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **One-time execution**: Function runs only once
2. **Thread-safe**: Can be called from multiple goroutines
3. **Blocking**: Other goroutines wait until first call completes
4. **No parameters**: Function must take no parameters
5. **Performance**: Very efficient for one-time initialization

## 🎯 When to Use sync.Once

1. **Singleton pattern**: Ensure only one instance exists
2. **Lazy initialization**: Initialize only when needed
3. **Resource setup**: Set up resources once
4. **Configuration loading**: Load configuration once
5. **One-time operations**: Any operation that should happen only once

## 🎯 Best Practices

1. **Use for initialization**:
   ```go
   var once sync.Once
   var instance *MyType
   
   func GetInstance() *MyType {
       once.Do(func() {
           instance = &MyType{}
       })
       return instance
   }
   ```

2. **Keep initialization function simple**:
   ```go
   // BAD - complex initialization
   once.Do(func() {
       // Complex logic here
   })
   
   // GOOD - simple initialization
   once.Do(func() {
       instance = &MyType{}
   })
   ```

3. **Handle errors properly**:
   ```go
   var once sync.Once
   var err error
   
   func Initialize() error {
       once.Do(func() {
           err = doInitialization()
       })
       return err
   }
   ```

## 🎯 Common Pitfalls

1. **Using sync.Once for multiple operations**:
   ```go
   // BAD - sync.Once is for one operation
   var once sync.Once
   once.Do(func() { /* operation 1 */ })
   once.Do(func() { /* operation 2 */ }) // This won't run
   
   // GOOD - use separate sync.Once for each operation
   var once1, once2 sync.Once
   once1.Do(func() { /* operation 1 */ })
   once2.Do(func() { /* operation 2 */ })
   ```

2. **Not handling errors**:
   ```go
   // BAD - error is lost
   once.Do(func() {
       if err := doSomething(); err != nil {
           // Error is lost
       }
   })
   
   // GOOD - handle errors
   var err error
   once.Do(func() {
       err = doSomething()
   })
   if err != nil {
       // Handle error
   }
   ```

3. **Using sync.Once for repeated operations**:
   ```go
   // BAD - sync.Once is for one-time operations
   for i := 0; i < 10; i++ {
       once.Do(func() {
           // This only runs once
       })
   }
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a connection pool:
- Multiple goroutines try to get a database connection
- Use sync.Once to initialize the connection pool only once
- Show how sync.Once ensures the pool is created only once
- Handle the case where initialization might fail

**Hint**: Use a struct to represent the connection pool and sync.Once to ensure it's initialized only once.

## 🚀 Next Steps

Now that you understand sync.Once, let's learn about **atomic operations** in the next file: `16-atomic-operations.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
