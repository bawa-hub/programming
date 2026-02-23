# SIGINT and SIGTERM -- Complete Notes (Backend & Interview Ready)

------------------------------------------------------------------------

# 1. What is a Signal?

A signal is a software interrupt sent by the operating system to a
process to notify it that an event has occurred.

Signals are part of the Unix/Linux process control mechanism.

They are asynchronous notifications delivered by the kernel.

------------------------------------------------------------------------

# 2. What is SIGINT?

SIGINT = Signal Interrupt

-   Triggered by: Ctrl + C in terminal
-   Sent by: User
-   Purpose: Interrupt a running process
-   Default behavior: Terminate the process

Example:

When you press:

Ctrl + C

The OS sends SIGINT to the running program.

------------------------------------------------------------------------

# 3. What is SIGTERM?

SIGTERM = Signal Terminate

-   Triggered by: kill `<pid>`{=html}
-   Sent by: System, Docker, Kubernetes, Process Manager
-   Purpose: Politely request process termination
-   Default behavior: Terminate the process

Unlike SIGINT, SIGTERM is typically used in production environments.

------------------------------------------------------------------------

# 4. Key Difference Between SIGINT and SIGTERM

  Feature           SIGINT          SIGTERM
  ----------------- --------------- -----------------------
  Source            User (Ctrl+C)   System / Orchestrator
  Usage             Development     Production
  Intent            Interrupt       Graceful termination
  Can be handled?   Yes             Yes

------------------------------------------------------------------------

# 5. Default Signal Behavior

If a process does not handle SIGINT or SIGTERM:

→ The process exits immediately.

If the process handles them:

→ It can perform graceful shutdown (cleanup, close DB, finish requests).

------------------------------------------------------------------------

# 6. What is SIGKILL?

SIGKILL = Force Kill

-   Triggered by: kill -9 `<pid>`{=html}
-   Cannot be caught
-   Cannot be ignored
-   Immediately terminates the process

Important:

SIGKILL does NOT allow graceful shutdown.

------------------------------------------------------------------------

# 7. How Signals Work Internally

1.  OS event occurs (Ctrl+C, kill command, container stop)
2.  Kernel sends signal to target process
3.  Process receives signal
4.  If handler exists → execute handler
5.  Else → default behavior

Signals are asynchronous and interrupt normal execution flow.

------------------------------------------------------------------------

# 8. Signal Handling in Go

Modern Go approach:

``` go
ctx, stop := signal.NotifyContext(
    context.Background(),
    syscall.SIGINT,
    syscall.SIGTERM,
)
defer stop()

<-ctx.Done()
```

This allows graceful shutdown when either signal is received.

------------------------------------------------------------------------

# 9. Kubernetes Shutdown Flow

When a Pod is terminated:

1.  Kubernetes sends SIGTERM
2.  Pod enters "Terminating" state
3.  terminationGracePeriodSeconds starts (default 30s)
4.  If process still running → SIGKILL is sent

Therefore:

Graceful shutdown must complete before grace period ends.

------------------------------------------------------------------------

# 10. Why SIGTERM is Important in Production

-   Used by Docker stop
-   Used by Kubernetes
-   Used by systemd
-   Used by cloud orchestration systems

Production systems rely on SIGTERM for controlled shutdown.

------------------------------------------------------------------------

# 11. Interview-Ready Answer (Short Version)

"SIGINT and SIGTERM are Unix signals used to terminate processes. SIGINT
is typically sent when a user presses Ctrl+C, while SIGTERM is sent by
systems like Docker or Kubernetes to request graceful termination.
Applications can intercept these signals to perform cleanup before
exiting."

------------------------------------------------------------------------

# 12. Advanced Notes (Senior Level)

-   Signals are delivered by the kernel.
-   They interrupt normal process execution.
-   They can be handled using signal handlers.
-   SIGKILL cannot be intercepted.
-   Graceful shutdown logic must handle SIGTERM.
-   Signal handling must be idempotent.

------------------------------------------------------------------------

# 13. Summary

-   SIGINT → User interrupt (Ctrl+C)
-   SIGTERM → System termination request
-   SIGKILL → Force kill (no cleanup possible)
-   Always handle SIGTERM in production systems
-   Always implement graceful shutdown

------------------------------------------------------------------------

END OF DOCUMENT
