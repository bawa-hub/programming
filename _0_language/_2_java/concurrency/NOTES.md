## Java Concurrency Mastery Roadmap

### 1. Big Picture: Why Concurrency?
- **Goal of concurrency**: Do more work in less *wall-clock* time by overlapping tasks (I/O, CPU) or improving responsiveness.
- **Concurrency vs parallelism**:
  - **Concurrency**: Structuring a program as multiple tasks that *can* run independently (even on a single core by time-slicing).
  - **Parallelism**: Actually running tasks at the same time on multiple cores.
- **Where it shows up in interviews**:
  - Correctness: race conditions, visibility, deadlocks, livelocks, starvation.
  - Tools: `synchronized`, `volatile`, `Lock`, `ReentrantLock`, `Semaphore`, `CountDownLatch`, `CyclicBarrier`, `ThreadLocal`, atomics, `ExecutorService`, `CompletableFuture`, parallel streams, fork/join.
  - Design: producer–consumer, async pipelines, bounded queues, throttling, rate limiting.

### 2. Java Memory & Thread Basics
- **Thread**:
  - Smallest unit of execution scheduled by the OS/JVM.
  - Each has its own stack; they share heap memory inside the same JVM process.
- **Creating threads**:
  - Extend `Thread` and override `run()`.
  - Implement `Runnable` or `Callable` and pass to a `Thread` or an `ExecutorService`.
  - See `concurrency.basics.ThreadCreationDemo` for hands-on examples.
- **Thread lifecycle**:
  - NEW → RUNNABLE → RUNNING → BLOCKED/WAITING/TIMED_WAITING → TERMINATED.

### 3. Fundamental Problems in Concurrency
- **Race condition**: Outcome depends on timing/ordering of threads due to shared mutable state.
- **Visibility problem**: One thread’s writes are not immediately visible to another (caches, reordering).
- **Atomicity**: Operations that must appear as a single, indivisible step (e.g., incrementing a counter).
-   - Try `concurrency.basics.SynchronizedCounterDemo` and run it multiple times.
- **Ordering / happens-before**:
  - Java Memory Model defines when writes by one thread are guaranteed visible to another.
  - Key tools: `synchronized`, `volatile`, `final` fields, `Lock` operations, thread start/join, executor task submission.

### 4. Low-Level Primitives
- **`synchronized`**:
  - Mutual exclusion: Only one thread can hold a given monitor at a time.
  - Provides *visibility* (entering/exiting a `synchronized` block establishes happens-before edges).
- **`volatile`**:
  - Guarantees visibility of reads/writes (no caching/reordering across volatile operations).
  - Does **not** provide atomicity for compound actions (like `count++`).
  - See `concurrency.basics.VisibilityVolatileDemo` and toggle `volatile` on/off to feel the difference.
- **`wait` / `notify` / `notifyAll`**:
  - Must be called inside `synchronized` on the same monitor object.
  - Pattern: while(conditionNotMet) { wait(); } then `notify`/`notifyAll` when condition becomes true.

### 5. Higher-Level Utilities (java.util.concurrent)
- **Locks & synchronizers**:
  - `Lock`, `ReentrantLock`, `ReadWriteLock`, `StampedLock`.
  - `Semaphore`, `CountDownLatch`, `CyclicBarrier`, `Phaser`, `Exchanger`.
  - See `concurrency.advanced.LocksDemo` for basic `ReentrantLock` usage.
- **Thread-safe collections**:
  - `ConcurrentHashMap`, `ConcurrentLinkedQueue`, `CopyOnWriteArrayList`, blocking queues (`ArrayBlockingQueue`, `LinkedBlockingQueue`, etc.).
  - See `concurrency.advanced.ConcurrentCollectionsDemo` for iteration and concurrent modifications.
- **Executors & thread pools**:
  - `Executor`, `ExecutorService`, `ScheduledExecutorService`, `ForkJoinPool`.
  - See `concurrency.advanced.ExecutorsDemo` for fixed/thread pool examples.
- **Async & reactive-style**:
  - `Future`, `CompletableFuture`, parallel streams, basic reactive patterns (even if outside core JDK).
  - See `concurrency.advanced.CompletableFutureDemo` for composing async tasks.

### 7. Classic Bugs & Patterns
- **Deadlock**:
  - Two or more threads wait forever for each other’s locks.
  - Usually caused by inconsistent lock acquisition order.
  - See `concurrency.advanced.DeadlockDemo` and then refactor it to avoid deadlock.
- **Livelock**:
  - Threads keep changing state in response to each other but no real progress is made.
- **Starvation**:
  - Some threads never get CPU or lock access due to scheduling or unfair locks.
- **Best practices**:
  - Prefer higher-level constructs (`ExecutorService`, `BlockingQueue`, `CompletableFuture`) over manual `Thread` and `wait/notify` in production code.
  - Avoid sharing mutable state; where needed, encapsulate it and guard with clear synchronization strategy.
  - Document which thread owns which data; design for immutability where possible.

### 6. Strategy for Mastery
- **Step 1**: Understand low-level primitives deeply (`synchronized`, `volatile`, `wait/notify`).
- **Step 2**: Learn and implement classic problems (producer–consumer, bounded buffer, readers–writers, dining philosophers).
- **Step 3**: Master `java.util.concurrent` tools and when to choose which.
- **Step 4**: Learn patterns for designing concurrent systems and spotting bugs in code.
- **Step 5**: Performance tuning, contention reduction, and avoiding over-engineering.

---

> As we go, we’ll link each section of these notes to specific Java classes under `src/main/java/concurrency` so you can open them and run experiments.

﻿

### 8. Thread Safety & Object Sharing

- **What does “thread-safe” mean?**
  - A class is **thread-safe** if it behaves correctly when accessed from multiple threads without additional external synchronization.
  - Correctness here means: invariants are preserved, no data races, no unexpected exceptions, and results match a valid sequential specification.
- **Common strategies to achieve thread-safety**:
  - **Immutability**:
    - All fields are `final` (or effectively final) and state never changes after construction.
    - No `this` reference escapes during construction (no publishing `this` from constructor).
    - Examples: `String`, wrapper classes like `Integer`, many `java.time` classes.
  - **Thread confinement**:
    - An object is used only from a single thread.
    - Example: using local variables inside methods, or objects kept strictly within a single thread.
  - **Confinement by ownership**:
    - A structure is accessed only under a single lock or via a particular component.
  - **Using thread-safe building blocks**:
    - `ConcurrentHashMap`, `BlockingQueue`, atomics, `CopyOnWriteArrayList`, etc.
  - **Explicit synchronization protocol**:
    - Well-defined which locks protect which data, documented in code comments.

- **Safe publication of objects**:
  - Publishing means making an object reference available to other threads.
  - **Unsafe publication** can lead to other threads seeing a partially constructed object.
  - **Safe publication mechanisms**:
    - Static initializers: objects stored in `static final` fields initialized in a static block.
    - Final fields: properly constructed objects with `final` fields have stronger visibility guarantees.
    - Storing into a `volatile` field or `AtomicReference`.
    - Storing into a thread-safe collection or under a synchronized block that all readers also use.

- **Levels of thread safety** (useful terminology in interviews):
  - **Immutable**: cannot be changed after construction.
  - **Thread-safe**: can be used from multiple threads without external synchronization.
  - **Conditionally thread-safe**: some methods are thread-safe only when certain conditions are met.
  - **Not thread-safe**: client must provide external synchronization.
  - **Thread-compatible**: safe as long as each instance is used only from a single thread.

### 9. Java Memory Model (JMM) – In Depth

- **Why JMM matters**:
  - Modern CPUs reorder instructions, and compilers/JITs also reorder and optimize.
  - The JMM defines **legal behaviors** of concurrent programs and **happens-before** rules that guarantee visibility and ordering.

- **Key concepts**:
  - **Happens-before**:
    - If A *happens-before* B, then:
      - All effects of A are visible to B.
      - A is ordered before B.
  - **Data race**:
    - Two threads access the same variable, at least one is a write, and there is no happens-before ordering between them.
    - Data races lead to **undefined behavior** within the limits of the JMM (non-intuitive results).

- **Important happens-before rules**:
  - **Program order**:
    - Within a single thread, actions appear to happen in program order.
  - **Monitor lock**:
    - An unlock on a monitor (exit of a `synchronized` block/method) happens-before every subsequent lock on that same monitor.
  - **Volatile variables**:
    - A write to a `volatile` field happens-before every subsequent read of that field.
  - **Thread start/join**:
    - A call to `Thread.start()` on a thread happens-before any actions in the started thread.
    - All actions in a thread happen-before a successful return from `Thread.join()` on that thread.
  - **Executor tasks**:
    - Submitting a task to an `ExecutorService` happens-before any actions in that task.
    - Completion of a task (e.g., `Future.get()`) happens-before returning from `get()`.
  - **Final fields**:
    - Properly constructed objects with `final` fields guarantee that other threads see the initialized values when they first use the reference that was safely published.

- **Reordering and visibility pitfalls**:
  - Without `volatile` or synchronization, the JVM/CPU is allowed to:
    - Cache values in registers or CPU caches.
    - Reorder independent reads/writes.
  - This is why naive double-checked locking without `volatile` is broken.

### 10. Intrinsic Locks & `wait/notify` Details

- **Intrinsic locks (monitor locks)**:
  - Every Java object has an internal **monitor** used for `synchronized`.
  - `synchronized (lock) { ... }` acquires the monitor associated with `lock`.
  - Intrinsic locks are:
    - **Reentrant**: the same thread can acquire the same lock multiple times.
    - **Mutually exclusive**: at most one thread holds a given lock at a time.

- **Choosing lock granularity**:
  - **Coarse-grained lock**: one lock for many pieces of data.
    - Simpler, but can create contention and reduce parallelism.
  - **Fine-grained locks**: separate locks for independent data.
    - More parallelism, but more complex and more potential for deadlocks.

- **`wait`, `notify`, `notifyAll`**:
  - Must be called by a thread that **owns the monitor** (`synchronized` on the same object).
  - Pattern:
    ```java
    synchronized (lock) {
        while (!condition) { // use while, not if, to handle spurious wakeups
            lock.wait();     // releases the lock and waits
        }
        // condition is true here, proceed
    }
    ```
  - **Spurious wakeups**:
    - A thread may wake up from `wait()` even if `notify`/`notifyAll` was not called.
    - Therefore, always use a `while` loop to re-check the condition.
  - **`notify` vs `notifyAll`**:
    - `notify` wakes a single waiting thread; `notifyAll` wakes all.
    - For many scenarios with multiple conditions or complex invariants, `notifyAll` is safer (though potentially less efficient).

### 11. Explicit Locks & Conditions

- **Why use `Lock` instead of `synchronized`?**
  - More flexible:
    - `lockInterruptibly()` – respond to interrupts while waiting for the lock.
    - `tryLock()` – attempt to acquire lock without blocking forever.
    - `tryLock(timeout, unit)` – bounded wait to avoid deadlock.
    - Multiple `Condition` objects per lock for different wait-sets.
  - Optional fairness policies (e.g., `new ReentrantLock(true)`).

- **`ReentrantLock` basics**:
  - Usage pattern:
    ```java
    lock.lock();
    try {
        // critical section
    } finally {
        lock.unlock();
    }
    ```
  - Similar semantics to intrinsic locks, but supports more operations as above.

- **`ReadWriteLock`**:
  - Typically `ReentrantReadWriteLock`.
  - Allows:
    - Multiple readers concurrently **when no writer holds the lock**.
    - A single writer with exclusive access.
  - Best for read-heavy workloads where writes are rare and well-separated.
  - Beware of writer starvation if many readers; fair mode can mitigate but may reduce throughput.

- **`StampedLock`**:
  - Provides:
    - Write locks.
    - Read locks.
    - **Optimistic reads**: non-blocking read that must later be validated.
  - Pattern:
    ```java
    long stamp = lock.tryOptimisticRead();
    // read state into locals
    if (!lock.validate(stamp)) {
        stamp = lock.readLock();
        try {
            // re-read state safely
        } finally {
            lock.unlockRead(stamp);
        }
    }
    ```
  - Can give better scalability for certain read-mostly workloads.
  - Not reentrant and does not support `Condition` objects; be careful with usage.

### 12. Synchronizers in Depth

- **`Semaphore`**:
  - Maintains a set of permits.
  - `acquire()` blocks until a permit is available; `release()` returns a permit.
  - Use cases:
    - Limit the number of concurrent accesses to a resource (e.g., 10 concurrent DB connections).
    - Implement simple resource pools.

- **`CountDownLatch`**:
  - Initialized with a **count**.
  - Threads call `await()` to block until the count reaches zero.
  - Other threads call `countDown()` to decrement the count.
  - One-time use: once the count reaches zero, it cannot be reset.
  - Use cases:
    - Wait for a set of worker threads to finish.
    - Wait for a set of services to be initialized before accepting traffic.

- **`CyclicBarrier`**:
  - Similar to `CountDownLatch`, but **reusable**.
  - A fixed number of parties call `await()`, and all are released when the last arrives.
  - Optionally runs a barrier action (a `Runnable`) when the barrier trips.
  - Use cases:
    - Iterative algorithms where threads must align at specific phases/steps.

- **`Phaser`**:
  - More flexible barrier for dynamic number of parties and multiple phases.
  - Threads can register/deregister dynamically.
  - Use when participants can join/leave over time.

- **`Exchanger`**:
  - Two threads exchange objects at a synchronization point.
  - Use cases:
    - Producer–consumer style handoff where the producer and consumer exchange buffers.

### 13. Concurrent Collections Deep Dive

- **Why not just use `Collections.synchronizedList` / `synchronizedMap`?**
  - Those wrappers synchronize each individual method but:
    - Compound operations like `if (!map.containsKey(k)) map.put(k, v);` are not atomic.
    - Iteration still needs external synchronization.
  - `java.util.concurrent` collections provide:
    - Better scalability via finer-grained locking or lock-free algorithms.
    - Well-defined concurrent behaviors.

- **`ConcurrentHashMap`**:
  - Thread-safe map with high concurrency.
  - Iterators are **weakly consistent**:
    - They reflect some state of the map at or after creation, but not necessarily a snapshot.
    - They do not throw `ConcurrentModificationException`.
  - Provides atomic methods:
    - `putIfAbsent`, `computeIfAbsent`, `compute`, `merge`.
  - Typical pattern in interviews:
    ```java
    map.computeIfAbsent(key, k -> new Value());
    ```
    which avoids race conditions of check-then-act.

- **`CopyOnWriteArrayList` / `CopyOnWriteArraySet`**:
  - On mutation (add/remove), they create a new underlying array.
  - Great for read-mostly workloads where iteration is far more common than modification.
  - Iterators see a **snapshot** of the list at creation time.

- **Blocking queues**:
  - `ArrayBlockingQueue`, `LinkedBlockingQueue`, `PriorityBlockingQueue`, `SynchronousQueue`, etc.
  - Fundamental to classic producer–consumer patterns.
  - Operations like `put` and `take` block, enabling backpressure and bounded resource usage.

### 14. Executors & Thread Pools – Deep Dive

- **Why use thread pools?**
  - Reuse threads instead of frequently creating/destroying them.
  - Centralized management of concurrency level.
  - Separation of task submission (producers) from execution policy.

- **Common executor factory methods** (`Executors`):
  - `newFixedThreadPool(n)`: pool with `n` threads, unbounded queue by default.
  - `newCachedThreadPool()`: creates new threads as needed, reusing idle ones; good for many short-lived async tasks, but be careful with unbounded thread growth.
  - `newSingleThreadExecutor()`: single worker thread; serializes tasks.
  - `newScheduledThreadPool(n)`: supports delayed and periodic execution.

- **Tuning considerations**:
  - For **CPU-bound** work:
    - Target pool size around `#cores` (or `cores ± 1`), depending on context.
  - For **I/O-bound** work:
    - Often can have more threads than cores, because many are blocked on I/O.
  - Avoid blocking operations inside a heavily used shared pool (especially `ForkJoinPool.commonPool()`).
  - Understand **rejection policies**:
    - `AbortPolicy` (default), `CallerRunsPolicy`, `DiscardPolicy`, `DiscardOldestPolicy`.

- **Custom thread pools via `ThreadPoolExecutor`**:
  - Parameters:
    - `corePoolSize`, `maximumPoolSize`, `keepAliveTime`, `workQueue`, `threadFactory`, `handler` (rejection).
  - Allows fine-grained control over how tasks are queued and executed.

### 15. Fork/Join & Parallel Streams

- **`ForkJoinPool` basics**:
  - Designed for tasks that can be recursively split into smaller subtasks (divide-and-conquer).
  - Uses **work stealing**: idle threads “steal” work from busier threads’ queues.
  - Tasks implement `RecursiveTask<V>` or `RecursiveAction`.

- **Parallel streams**:
  - Created by calling `.parallel()` on a stream.
  - Uses the **common ForkJoinPool** by default.
  - Best use cases:
    - CPU-bound operations.
    - Large data sets.
    - Stateless, side-effect-free operations.
  - Pitfalls:
    - Shared mutable state inside parallel operations → data races.
    - Blocking I/O in parallel stream operations → thread starvation.
    - Small collections → overhead outweighs benefits.

### 16. `CompletableFuture` – Advanced Usage

- **Core ideas**:
  - Represents a future result of an async computation with powerful composition APIs.
  - Can be manually completed or completed by async tasks.

- **Creation**:
  - `CompletableFuture.supplyAsync(Supplier)` – compute a value asynchronously.
  - `runAsync(Runnable)` – run a task without a result.
  - Both can take an `Executor` to control the thread pool.

- **Composition patterns**:
  - **Transforming results**:
    - `thenApply(fn)` – synchronous transform.
    - `thenApplyAsync(fn)` – transform possibly on another thread.
  - **Chaining dependent tasks**:
    - `thenCompose(fn)` – flat-map style; for async operations that depend on previous result.
  - **Combining independent tasks**:
    - `thenCombine(other, combiner)` – combine results of two futures.
    - `allOf(f1, f2, ...)` – wait for all to complete.
    - `anyOf(f1, f2, ...)` – completes when any completes.

- **Error handling**:
  - `exceptionally(fn)` – recover from an exception and provide fallback.
  - `handle((result, ex) -> ...)` – handle both success and failure in one place.
  - `whenComplete((result, ex) -> ...)` – side-effect after completion, does not change final result.

- **Timeouts and cancellation**:
  - `orTimeout(timeout, unit)` – complete exceptionally if not done in time.
  - `completeOnTimeout(value, timeout, unit)` – substitute default value if timeout occurs.
  - `cancel(true)` – attempt to cancel; may interrupt running task if supported.

### 17. Performance & Contention

- **Contention & scalability**:
  - Heavy lock contention reduces parallel speedup.
  - Techniques:
    - Reduce critical-section size.
    - Use more granular locks or lock-free structures.
    - Use read-write locks when reads dominate and writes are rare.

- **False sharing**:
  - Multiple threads frequently updating different variables that share the same cache line.
  - Can cause performance degradation even though they are logically independent.
  - Mitigation:
    - Padding (e.g., manual padding fields, or using specialized classes like `LongAdder` for hot counters).

- **Choosing the right primitive**:
  - Simple shared counter:
    - Low contention: `AtomicLong` is OK.
    - High contention: `LongAdder` or `LongAccumulator` can scale better.
  - Queue between producer and consumer:
    - Use a `BlockingQueue`.
  - Shared map with heavy concurrent access:
    - Use `ConcurrentHashMap`.

### 18. Interview-Focused Patterns & Tips

- **Classic coding problems** (you should be able to implement and reason about):
  - Producer–consumer with `BlockingQueue` and with `wait/notify`.
  - Bounded buffer with backpressure.
  - Readers–writers lock pattern.
  - Dining philosophers (demonstrate deadlock and then fix via ordering or other strategies).
  - Rate limiter / throttling using `Semaphore` or scheduled tasks.

- **What interviewers look for**:
  - Clear mental model of memory visibility and race conditions.
  - Ability to **explain** why `volatile` is not enough for compound operations.
  - Correct use of higher-level abstractions instead of reinventing low-level mechanisms.
  - Ability to reason about deadlock, livelock, and starvation, and propose fixes.

- **How to practice effectively**:
  - Take each section above and:
    - Write a small example or modify the existing ones in `concurrency.basics` and `concurrency.advanced`.
    - Add logging, sleeps, and artificial contention to explore behavior.
  - Re-explain concepts (JMM, happens-before, deadlock avoidance) **out loud** as if teaching.
  - Implement the same pattern using:
    - Raw `Thread` + `synchronized`.
    - `ExecutorService` + blocking queues.
    - `CompletableFuture` where appropriate.
