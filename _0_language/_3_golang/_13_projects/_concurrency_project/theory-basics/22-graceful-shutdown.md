# 22 - Graceful Shutdown

## 🎯 Learning Objectives
- Understand what graceful shutdown is and why it's important
- Learn how to implement graceful shutdown in Go
- Master different shutdown patterns and strategies
- Practice with shutdown best practices
- Understand when to use different shutdown approaches

## 📚 Theory

### What is Graceful Shutdown?

**Graceful shutdown** is the process of shutting down an application in a controlled manner, allowing it to:
- Finish current operations
- Clean up resources
- Save state
- Notify other services
- Handle pending requests

**Why is it important?**
1. **Data integrity**: Ensure data is not lost
2. **Resource cleanup**: Free up resources properly
3. **User experience**: Handle requests gracefully
4. **System stability**: Prevent system crashes
5. **Monitoring**: Proper logging and metrics

### Shutdown Signals

Go applications can receive shutdown signals:
- **SIGINT**: Interrupt signal (Ctrl+C)
- **SIGTERM**: Termination signal
- **SIGQUIT**: Quit signal

### Shutdown Patterns

1. **Signal handling**: Listen for shutdown signals
2. **Context cancellation**: Use context for shutdown
3. **Resource cleanup**: Clean up resources properly
4. **Timeout handling**: Set shutdown timeouts
5. **State saving**: Save application state

## 💻 Code Examples

### Example 1: Basic Graceful Shutdown

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func basicGracefulShutdown() {
    fmt.Println("=== Basic Graceful Shutdown ===")
    
    // Create channel to receive shutdown signals
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
    
    // Start background work
    go func() {
        for i := 1; i <= 10; i++ {
            fmt.Printf("Working... %d\n", i)
            time.Sleep(500 * time.Millisecond)
        }
        fmt.Println("Work completed")
    }()
    
    // Wait for shutdown signal
    <-shutdown
    fmt.Println("Shutdown signal received, shutting down gracefully...")
    
    // Cleanup
    fmt.Println("Cleaning up resources...")
    time.Sleep(1 * time.Second)
    fmt.Println("Shutdown complete")
}

func main() {
    basicGracefulShutdown()
}
```

**Run this code:**
```bash
go run 22-graceful-shutdown.go
# Press Ctrl+C to test shutdown
```

### Example 2: Graceful Shutdown with Context

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func gracefulShutdownWithContext() {
    fmt.Println("=== Graceful Shutdown with Context ===")
    
    // Create context that can be cancelled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        cancel()
    }()
    
    // Start background work
    go func() {
        for i := 1; i <= 20; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("Work cancelled")
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(300 * time.Millisecond)
            }
        }
        fmt.Println("Work completed")
    }()
    
    // Wait for context to be cancelled
    <-ctx.Done()
    fmt.Println("Shutting down gracefully...")
    
    // Cleanup
    fmt.Println("Cleaning up resources...")
    time.Sleep(1 * time.Second)
    fmt.Println("Shutdown complete")
}

func main() {
    gracefulShutdownWithContext()
}
```

### Example 3: Graceful Shutdown with Timeout

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func gracefulShutdownWithTimeout() {
    fmt.Println("=== Graceful Shutdown with Timeout ===")
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        cancel()
    }()
    
    // Start background work
    go func() {
        for i := 1; i <= 20; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Work cancelled: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(400 * time.Millisecond)
            }
        }
        fmt.Println("Work completed")
    }()
    
    // Wait for context to be done
    <-ctx.Done()
    
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("Shutdown timeout exceeded, forcing shutdown")
    } else {
        fmt.Println("Shutting down gracefully...")
    }
    
    // Cleanup
    fmt.Println("Cleaning up resources...")
    time.Sleep(1 * time.Second)
    fmt.Println("Shutdown complete")
}

func main() {
    gracefulShutdownWithTimeout()
}
```

### Example 4: Graceful Shutdown with Resource Cleanup

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type Resource struct {
    ID   int
    Name string
}

type ResourceManager struct {
    resources map[int]*Resource
    mutex     sync.RWMutex
}

func NewResourceManager() *ResourceManager {
    return &ResourceManager{
        resources: make(map[int]*Resource),
    }
}

func (rm *ResourceManager) AddResource(id int, name string) {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()
    rm.resources[id] = &Resource{ID: id, Name: name}
    fmt.Printf("Added resource: %s (ID: %d)\n", name, id)
}

func (rm *ResourceManager) GetResource(id int) *Resource {
    rm.mutex.RLock()
    defer rm.mutex.RUnlock()
    return rm.resources[id]
}

func (rm *ResourceManager) Cleanup() {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()
    
    fmt.Println("Cleaning up resources...")
    for id, resource := range rm.resources {
        fmt.Printf("Cleaning up resource: %s (ID: %d)\n", resource.Name, id)
        time.Sleep(100 * time.Millisecond) // Simulate cleanup
    }
    rm.resources = make(map[int]*Resource)
    fmt.Println("Resource cleanup complete")
}

func gracefulShutdownWithResourceCleanup() {
    fmt.Println("=== Graceful Shutdown with Resource Cleanup ===")
    
    // Create resource manager
    rm := NewResourceManager()
    
    // Add some resources
    for i := 1; i <= 5; i++ {
        rm.AddResource(i, fmt.Sprintf("Resource-%d", i))
    }
    
    // Create context that can be cancelled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        cancel()
    }()
    
    // Start background work
    go func() {
        for i := 1; i <= 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("Work cancelled")
                return
            default:
                resource := rm.GetResource(i % 5)
                if resource != nil {
                    fmt.Printf("Using resource: %s\n", resource.Name)
                }
                time.Sleep(500 * time.Millisecond)
            }
        }
        fmt.Println("Work completed")
    }()
    
    // Wait for context to be cancelled
    <-ctx.Done()
    fmt.Println("Shutting down gracefully...")
    
    // Cleanup resources
    rm.Cleanup()
    
    fmt.Println("Shutdown complete")
}

func main() {
    gracefulShutdownWithResourceCleanup()
}
```

### Example 5: Graceful Shutdown with Multiple Services

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type Service struct {
    Name string
    ctx  context.Context
}

func (s *Service) Start() {
    go func() {
        for i := 1; i <= 20; i++ {
            select {
            case <-s.ctx.Done():
                fmt.Printf("Service %s stopped\n", s.Name)
                return
            default:
                fmt.Printf("Service %s working... %d\n", s.Name, i)
                time.Sleep(300 * time.Millisecond)
            }
        }
        fmt.Printf("Service %s completed\n", s.Name)
    }()
}

func gracefulShutdownWithMultipleServices() {
    fmt.Println("=== Graceful Shutdown with Multiple Services ===")
    
    // Create context that can be cancelled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Create services
    services := []*Service{
        {Name: "Database", ctx: ctx},
        {Name: "API", ctx: ctx},
        {Name: "Cache", ctx: ctx},
    }
    
    // Start all services
    for _, service := range services {
        service.Start()
    }
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        cancel()
    }()
    
    // Wait for context to be cancelled
    <-ctx.Done()
    fmt.Println("Shutting down all services...")
    
    // Wait for services to stop
    time.Sleep(1 * time.Second)
    fmt.Println("All services stopped")
}

func main() {
    gracefulShutdownWithMultipleServices()
}
```

### Example 6: Graceful Shutdown with WaitGroup

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func gracefulShutdownWithWaitGroup() {
    fmt.Println("=== Graceful Shutdown with WaitGroup ===")
    
    // Create context that can be cancelled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Create WaitGroup
    var wg sync.WaitGroup
    
    // Start multiple workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for j := 1; j <= 10; j++ {
                select {
                case <-ctx.Done():
                    fmt.Printf("Worker %d stopped\n", workerID)
                    return
                default:
                    fmt.Printf("Worker %d working... %d\n", workerID, j)
                    time.Sleep(400 * time.Millisecond)
                }
            }
            fmt.Printf("Worker %d completed\n", workerID)
        }(i)
    }
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        cancel()
    }()
    
    // Wait for all workers to complete
    go func() {
        wg.Wait()
        fmt.Println("All workers completed")
    }()
    
    // Wait for context to be cancelled
    <-ctx.Done()
    fmt.Println("Shutting down gracefully...")
    
    // Wait for workers to finish
    wg.Wait()
    fmt.Println("Shutdown complete")
}

func main() {
    gracefulShutdownWithWaitGroup()
}
```

### Example 7: Graceful Shutdown with HTTP Server

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func gracefulShutdownWithHTTPServer() {
    fmt.Println("=== Graceful Shutdown with HTTP Server ===")
    
    // Create HTTP server
    server := &http.Server{
        Addr: ":8080",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Printf("Handling request: %s %s\n", r.Method, r.URL.Path)
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Hello, World!"))
        }),
    }
    
    // Start server in goroutine
    go func() {
        fmt.Println("Starting HTTP server on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Server error: %v\n", err)
        }
    }()
    
    // Wait for shutdown signal
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
    <-shutdown
    
    fmt.Println("Shutdown signal received, shutting down server...")
    
    // Create context with timeout for shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Shutdown server
    if err := server.Shutdown(ctx); err != nil {
        fmt.Printf("Server shutdown error: %v\n", err)
    } else {
        fmt.Println("Server shutdown complete")
    }
}

func main() {
    gracefulShutdownWithHTTPServer()
}
```

### Example 8: Graceful Shutdown Best Practices

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

type Application struct {
    services []Service
    ctx      context.Context
    cancel   context.CancelFunc
    wg       sync.WaitGroup
}

type Service interface {
    Start(ctx context.Context) error
    Stop() error
    Name() string
}

type DatabaseService struct {
    name string
}

func (d *DatabaseService) Start(ctx context.Context) error {
    fmt.Printf("Starting %s service\n", d.name)
    return nil
}

func (d *DatabaseService) Stop() error {
    fmt.Printf("Stopping %s service\n", d.name)
    time.Sleep(500 * time.Millisecond) // Simulate cleanup
    return nil
}

func (d *DatabaseService) Name() string {
    return d.name
}

type APIService struct {
    name string
}

func (a *APIService) Start(ctx context.Context) error {
    fmt.Printf("Starting %s service\n", a.name)
    return nil
}

func (a *APIService) Stop() error {
    fmt.Printf("Stopping %s service\n", a.name)
    time.Sleep(300 * time.Millisecond) // Simulate cleanup
    return nil
}

func (a *APIService) Name() string {
    return a.name
}

func NewApplication() *Application {
    ctx, cancel := context.WithCancel(context.Background())
    return &Application{
        services: []Service{
            &DatabaseService{name: "Database"},
            &APIService{name: "API"},
        },
        ctx:    ctx,
        cancel: cancel,
    }
}

func (app *Application) Start() {
    fmt.Println("Starting application...")
    
    // Start all services
    for _, service := range app.services {
        app.wg.Add(1)
        go func(s Service) {
            defer app.wg.Done()
            if err := s.Start(app.ctx); err != nil {
                fmt.Printf("Error starting %s: %v\n", s.Name(), err)
            }
        }(service)
    }
    
    fmt.Println("Application started")
}

func (app *Application) Stop() {
    fmt.Println("Stopping application...")
    
    // Cancel context to signal shutdown
    app.cancel()
    
    // Wait for services to stop
    app.wg.Wait()
    
    // Stop all services
    for _, service := range app.services {
        if err := service.Stop(); err != nil {
            fmt.Printf("Error stopping %s: %v\n", service.Name(), err)
        }
    }
    
    fmt.Println("Application stopped")
}

func gracefulShutdownBestPractices() {
    fmt.Println("=== Graceful Shutdown Best Practices ===")
    
    // Create application
    app := NewApplication()
    
    // Start application
    app.Start()
    
    // Handle shutdown signals
    go func() {
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
        <-shutdown
        fmt.Println("Shutdown signal received")
        app.Stop()
    }()
    
    // Keep application running
    select {}
}

func main() {
    gracefulShutdownBestPractices()
}
```

## 🧪 Key Concepts to Remember

1. **Signal handling**: Listen for shutdown signals
2. **Context cancellation**: Use context for shutdown
3. **Resource cleanup**: Clean up resources properly
4. **Timeout handling**: Set shutdown timeouts
5. **State saving**: Save application state
6. **WaitGroup**: Wait for goroutines to finish

## 🎯 When to Use Different Shutdown Approaches

1. **Signal handling**: When you need to handle OS signals
2. **Context cancellation**: When you need to cancel operations
3. **Resource cleanup**: When you need to clean up resources
4. **Timeout handling**: When you need to set shutdown timeouts
5. **WaitGroup**: When you need to wait for goroutines

## 🎯 Best Practices

1. **Always handle shutdown signals**:
   ```go
   // GOOD - handle shutdown signals
   shutdown := make(chan os.Signal, 1)
   signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
   <-shutdown
   ```

2. **Use context for cancellation**:
   ```go
   // GOOD - use context
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   ```

3. **Clean up resources**:
   ```go
   // GOOD - clean up resources
   defer func() {
       // Cleanup code
   }()
   ```

4. **Set shutdown timeouts**:
   ```go
   // GOOD - set timeout
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   ```

## 🎯 Common Pitfalls

1. **Not handling shutdown signals**:
   ```go
   // BAD - no shutdown handling
   for {
       // Long running operation
   }
   
   // GOOD - handle shutdown
   go func() {
       shutdown := make(chan os.Signal, 1)
       signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
       <-shutdown
       // Handle shutdown
   }()
   ```

2. **Not cleaning up resources**:
   ```go
   // BAD - no cleanup
   func main() {
       // Start services
       // No cleanup
   }
   
   // GOOD - cleanup resources
   func main() {
       defer cleanup()
       // Start services
   }
   ```

3. **Not using timeouts**:
   ```go
   // BAD - no timeout
   ctx, cancel := context.WithCancel(context.Background())
   
   // GOOD - use timeout
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   ```

## 🏋️ Exercise

**Problem**: Create a program that simulates a microservice with graceful shutdown:
- Multiple services (database, API, cache)
- Each service can be started and stopped
- Handle shutdown signals properly
- Clean up resources on shutdown
- Use context for cancellation
- Show how to implement proper graceful shutdown

**Hint**: Use a struct to represent the microservice and implement proper startup and shutdown methods.

## 🚀 Next Steps

Now that you understand graceful shutdown, let's learn about **monitoring concurrency** in the next file: `23-monitoring-concurrency.md`

---

**Remember**: Always run your concurrent code with `go run -race` to check for race conditions!
