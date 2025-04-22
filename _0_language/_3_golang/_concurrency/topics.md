# Chapter 1: Concurrency Foundations
    What is Concurrency vs Parallelism?
    Goroutines: Go’s Lightweight Threads
    The Go Scheduler & How Goroutines Work Internally

# Chapter 2: Synchronization Basics
    Race Conditions & Data Safety
    sync.WaitGroup
    sync.Mutex & sync.RWMutex
    Atomic Operations (sync/atomic)
    Deadlocks

# Chapter 3: Communication with Channels
    Buffered and Unbuffered Channels
    Channel Directions
    Closing Channels & Range over Channels
    Select Statement
    Channel Deadlocks and Best Practices

🔹 Chapter 4: Advanced Patterns

    Fan-out / Fan-in Pattern

    Worker Pool Pattern

    Pipeline Pattern

    Publish-Subscribe Pattern with Channels

    Rate Limiting (Token Bucket / Leaky Bucket)

🔹 Chapter 5: Contexts and Goroutine Management

    context.Background, WithCancel, WithTimeout, WithDeadline

    Graceful Shutdown of Goroutines

    Parent-Child Goroutine Relationships

🔹 Chapter 6: Error Handling and Coordination

    Error Handling with Channels

    sync.Once and sync.Cond

    errgroup.Group Pattern (from golang.org/x/sync/errgroup)

🔹 Chapter 7: Testing & Debugging Concurrency

    The Go Race Detector (go test -race)

    Writing Concurrency-safe Unit Tests

    Benchmarking Concurrent Code

    Goroutine Leak Detection & Profiling

🔹 Chapter 8: Real-World Concurrency

    Concurrent Job Scheduler (like cron)

    Real-time Chat Server

    Concurrent Crawler

    Concurrent Map Implementation

    Thread-safe In-Memory Cache

🔹 Chapter 9: Interview Deep Dive

    20 Most Common Go Concurrency Interview Questions

    Writing & Explaining Lock-Free Structures

    How Go Compares to Java/Python in Concurrency

    Building a Concurrency Library from Scratch