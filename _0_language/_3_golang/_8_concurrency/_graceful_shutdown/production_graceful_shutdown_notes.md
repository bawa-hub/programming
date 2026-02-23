# Production-Level Graceful Shutdown in Go (Complete Notes)

------------------------------------------------------------------------

# 1. What is Graceful Shutdown?

Graceful shutdown is the controlled termination of an application where:

-   New requests are stopped
-   In-flight requests are allowed to finish
-   Background workers are stopped safely
-   External resources are closed properly
-   The process exits without data loss

It prevents abrupt termination and ensures system consistency.

------------------------------------------------------------------------

# 2. Why Graceful Shutdown is Critical in Production

In real systems:

-   Kubernetes sends SIGTERM during pod termination
-   Docker stop sends SIGTERM
-   Load balancers drain connections during deployments
-   Auto-scaling removes instances

Without graceful shutdown:

-   Requests may be dropped
-   DB transactions may remain incomplete
-   Kafka offsets may not commit
-   Goroutines may leak
-   System may become inconsistent

------------------------------------------------------------------------

# 3. Graceful Shutdown Lifecycle

1.  Receive termination signal (SIGINT / SIGTERM)
2.  Stop accepting new traffic
3.  Cancel running operations
4.  Wait for in-flight work to complete
5.  Close external dependencies
6.  Exit process

------------------------------------------------------------------------

# 4. Signal Handling (Modern Go Pattern)

Recommended approach (Go 1.16+):

``` go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()

<-ctx.Done()
```

Why this is better:

-   Automatically cancels context on signal
-   Cleaner lifecycle management
-   Avoids manual signal channels

------------------------------------------------------------------------

# 5. Production HTTP Server Setup

Always configure timeouts:

``` go
srv := &http.Server{
    Addr:              ":8080",
    Handler:           handler,
    ReadTimeout:       10 * time.Second,
    WriteTimeout:      15 * time.Second,
    IdleTimeout:       60 * time.Second,
    ReadHeaderTimeout: 5 * time.Second,
}
```

Why timeouts matter:

-   Prevent Slowloris attacks
-   Prevent hanging connections
-   Avoid resource exhaustion

------------------------------------------------------------------------

# 6. Starting Server Safely

``` go
go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server failed: %v", err)
    }
}()
```

Why check `http.ErrServerClosed`?

Because Shutdown() returns that error normally.

------------------------------------------------------------------------

# 7. Graceful HTTP Shutdown

``` go
shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()

if err := srv.Shutdown(shutdownCtx); err != nil {
    log.Printf("shutdown error: %v", err)
}
```

What `Shutdown()` does:

-   Stops accepting new connections
-   Waits for active requests to finish
-   Closes idle connections

------------------------------------------------------------------------

# 8. Closing External Resources

Correct order:

1.  Stop traffic
2.  Stop background workers
3.  Close message producers/consumers
4.  Close Redis/cache
5.  Close database connection pool

Example for database (GORM):

``` go
sqlDB, _ := db.DB()
sqlDB.Close()
```

Important:

GORM wraps `*sql.DB`, so close the underlying pool.

------------------------------------------------------------------------

# 9. WaitGroup for Background Workers

``` go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    worker(ctx)
}()

wg.Wait()
```

Always wait for background goroutines to exit.

------------------------------------------------------------------------

# 10. Kubernetes Shutdown Lifecycle

When pod is terminated:

1.  SIGTERM sent
2.  Pod enters Terminating state
3.  readinessProbe fails (removed from LB)
4.  terminationGracePeriodSeconds starts (default 30s)
5.  If not exited → SIGKILL

Your shutdown must complete within grace period.

------------------------------------------------------------------------

# 11. Advanced Production Practices

-   Make shutdown idempotent
-   Handle second signal for force exit
-   Add shutdown timeout
-   Fail readiness probe during shutdown
-   Flush metrics before exit
-   Structured logging
-   Drain queues before exit

Example of second-signal force exit:

``` go
go func() {
    <-ctx.Done()
    <-ctx.Done()
    os.Exit(1)
}()
```

------------------------------------------------------------------------

# 12. Common Mistakes

❌ Not handling SIGTERM\
❌ No timeout on shutdown\
❌ Not cancelling goroutines\
❌ Closing channels from receiver\
❌ Forgetting DB close\
❌ Not waiting for workers

------------------------------------------------------------------------

# 13. Interview Answer (Short Version)

"Graceful shutdown means handling termination signals like SIGTERM,
stopping new requests, cancelling in-flight operations using context,
waiting for goroutines using WaitGroup, closing external resources like
database connections, and exiting safely. In Go, this is implemented
using signal.NotifyContext, server.Shutdown with timeout, and structured
cleanup order."

------------------------------------------------------------------------

# 14. Key Takeaways

-   Always use context for cancellation
-   Always define HTTP timeouts
-   Always use Shutdown() instead of Close()
-   Always close external resources
-   Always wait for goroutines
-   Always use timeout for shutdown

------------------------------------------------------------------------

END OF DOCUMENT
