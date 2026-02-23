# Graceful Shutdown -- Complete Interview Notes

------------------------------------------------------------------------

# 1. What is Graceful Shutdown?

Graceful shutdown is the process of terminating an application in a
controlled manner such that:

-   No in-flight requests are dropped
-   No data is lost
-   No resources are leaked
-   The system remains consistent
-   The service exits cleanly

It ensures reliability and stability in production systems.

------------------------------------------------------------------------

# 2. Why Do We Need Graceful Shutdown?

In real-world systems:

-   Kubernetes sends SIGTERM when stopping a pod
-   Load balancers remove instances during deployments
-   Cloud auto-scaling terminates instances
-   Manual service restarts happen

If shutdown is not graceful:

-   Requests may be dropped
-   Transactions may be partially committed
-   DB connections may leak
-   Kafka consumers may reprocess messages incorrectly

------------------------------------------------------------------------

# 3. Difference: Graceful vs Forceful Shutdown

  Graceful Shutdown         Forceful Shutdown
  ------------------------- ---------------------------
  Finishes in-flight work   Immediately kills process
  Releases resources        May leak resources
  Prevents data loss        Can corrupt data
  Controlled exit           Abrupt termination

------------------------------------------------------------------------

# 4. Core Steps in Graceful Shutdown

1.  Receive termination signal (SIGTERM / SIGINT)
2.  Stop accepting new work
3.  Cancel in-flight operations
4.  Wait for goroutines/workers to finish
5.  Close resources (DB, Kafka, Redis, etc.)
6.  Exit process

------------------------------------------------------------------------

# 5. Signal Handling in Go

Common signals:

-   SIGINT (Ctrl+C)
-   SIGTERM (Kubernetes, Docker stop)
-   SIGKILL (Cannot be caught)

Example:

``` go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
<-sigCh
```

------------------------------------------------------------------------

# 6. Context-Based Cancellation Pattern

The standard production pattern in Go:

``` go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go worker(ctx)

<-sigCh
cancel()
```

Workers must listen:

``` go
select {
case <-ctx.Done():
    return
}
```

------------------------------------------------------------------------

# 7. Using WaitGroup to Wait for Workers

``` go
var wg sync.WaitGroup
wg.Add(1)

go func() {
    defer wg.Done()
    worker(ctx)
}()

wg.Wait()
```

------------------------------------------------------------------------

# 8. Full Production Example

``` go
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

func worker(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker shutting down...")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    var wg sync.WaitGroup

    wg.Add(1)
    go worker(ctx, &wg)

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

    <-sigCh
    fmt.Println("Shutdown signal received")

    cancel()
    wg.Wait()

    fmt.Println("Cleanup complete. Exiting.")
}
```

------------------------------------------------------------------------

# 9. Graceful Shutdown for HTTP Servers

``` go
srv := &http.Server{Addr: ":8080"}

go srv.ListenAndServe()

<-sigCh

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

srv.Shutdown(ctx)
```

What Shutdown() does:

-   Stops accepting new connections
-   Waits for active requests
-   Closes idle connections

------------------------------------------------------------------------

# 10. Kubernetes Shutdown Lifecycle

When a Pod is terminated:

1.  SIGTERM is sent
2.  Pod enters Terminating state
3.  readinessProbe fails (removed from load balancer)
4.  terminationGracePeriodSeconds starts (default 30s)
5.  If still running → SIGKILL

Shutdown must complete within grace period.

------------------------------------------------------------------------

# 11. Shutdown in Message Consumers (Kafka Example)

Steps:

1.  Stop polling new messages
2.  Finish processing current message
3.  Commit offset
4.  Close consumer

Failure to do this may cause duplicate processing.

------------------------------------------------------------------------

# 12. Common Mistakes

❌ Not listening for SIGTERM\
❌ Not cancelling goroutines\
❌ Forgetting WaitGroup\
❌ Not setting timeout for shutdown\
❌ Closing channels from receiver side\
❌ Ignoring DB/Redis/Kafka close

------------------------------------------------------------------------

# 13. Advanced Considerations

-   Idempotent shutdown (handle multiple signals safely)
-   Shutdown timeouts
-   Health check failing before shutdown
-   Metrics flushing
-   Structured logging
-   Draining queues
-   Circuit breaker states

------------------------------------------------------------------------

# 14. Interview Answer Template (30 Seconds)

"Graceful shutdown means handling termination signals like SIGTERM,
stopping new requests, cancelling in-flight operations using context,
waiting for goroutines using WaitGroup, cleaning up resources like
database connections, and exiting safely. In Go, this is typically
implemented using signal.Notify, context cancellation, and
server.Shutdown for HTTP services."

------------------------------------------------------------------------

# 15. Key Takeaways

-   Sender handles closing
-   Always use context for cancellation
-   Always wait for workers
-   Always use timeout for shutdown
-   Production systems require graceful shutdown

------------------------------------------------------------------------

END OF NOTES
