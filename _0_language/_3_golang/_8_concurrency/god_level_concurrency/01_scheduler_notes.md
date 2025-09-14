# 🚀 GOD-LEVEL: Go Runtime Scheduler Deep Dive

## 📚 Theory Notes

### **G-M-P Model (Goroutines, Machine, Processors)**

The Go runtime scheduler is built around the G-M-P model:

```
┌─────────────────────────────────────────────────────────────┐
│                    MACHINE (M)                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ Processor 1 │  │ Processor 2 │  │ Processor 3 │  ...   │
│  │     (P)     │  │     (P)     │  │     (P)     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│       │                │                │                 │
│   ┌───▼───┐        ┌───▼───┐        ┌───▼───┐             │
│   │ G1,G2 │        │ G3,G4 │        │ G5,G6 │             │
│   │ G7,G8 │        │ G9,G10│        │G11,G12│             │
│   └───────┘        └───────┘        └───────┘             │
└─────────────────────────────────────────────────────────────┘
```

#### **Components:**

1. **G (Goroutine)**
   - Lightweight thread managed by Go runtime
   - 2KB initial stack (grows/shrinks as needed)
   - Contains: stack, program counter, scheduling info
   - Very cheap to create (thousands can run on single OS thread)

2. **M (Machine)**
   - OS thread managed by Go runtime
   - Executes goroutines on a processor
   - Blocked M's don't consume CPU
   - Number limited by `GOMAXPROCS`

3. **P (Processor)**
   - Logical processor that runs goroutines
   - Has local run queue
   - Number equals `GOMAXPROCS`
   - Manages work distribution

### **Work Stealing Algorithm**

#### **How it Works:**
1. Each P has a local run queue (FIFO)
2. When P's queue is empty, it steals from other P's queues
3. Steals from the middle of the queue (random access)
4. Global run queue as fallback
5. Network poller handles I/O completion

#### **Benefits:**
- Load balancing across processors
- No central dispatcher (scalable)
- Efficient cache utilization
- Reduces contention

### **Scheduling Decisions**

#### **Triggers:**
1. **Blocking system calls** (I/O operations)
2. **Channel operations** (send/receive)
3. **Time slice expiration** (10ms)
4. **Function calls** (stack growth)
5. **Garbage collection**
6. **Explicit yields** (`runtime.Gosched()`)

#### **Preemption (Go 1.14+):**
- **Cooperative preemption** (function calls)
- **Non-cooperative preemption** (async preemption)
- Ensures fair scheduling
- Prevents goroutine starvation

### **Memory Management**

#### **Stack Management:**
- Goroutine stacks start at 2KB
- Stacks grow/shrink as needed
- No stack overflow (unlike C)
- Efficient memory usage

#### **GC Interaction:**
- GC is concurrent and low-latency
- Work stealing reduces GC pressure
- Scheduler coordinates with GC
- Minimal impact on application performance

### **NUMA Awareness**

#### **NUMA (Non-Uniform Memory Access):**
- Modern systems have multiple CPU sockets
- Memory access speed varies by distance
- Go runtime is NUMA-aware
- Work stealing considers NUMA topology

#### **CPU Affinity Considerations:**
- Keep related goroutines on same CPU
- Minimize cross-socket memory access
- Use `sync.Pool` for CPU-local caching
- Consider NUMA topology for large systems

## 🔧 Practical Applications

### **When to Use Different Patterns:**

1. **CPU-Intensive Work:**
   - Use worker pools
   - Consider `GOMAXPROCS` tuning
   - Watch for false sharing

2. **I/O-Intensive Work:**
   - Many goroutines are fine
   - Use channels for coordination
   - Consider connection pooling

3. **Mixed Workloads:**
   - Separate CPU and I/O goroutines
   - Use different patterns for each
   - Monitor scheduling behavior

### **Performance Tuning:**

1. **GOMAXPROCS:**
   - Default: number of CPUs
   - Can be set at runtime
   - Affects parallelism, not concurrency

2. **Goroutine Count:**
   - Monitor with `runtime.NumGoroutine()`
   - Watch for goroutine leaks
   - Use profiling to identify issues

3. **Scheduling Profiling:**
   - Use `GODEBUG=schedtrace=1000`
   - Use `runtime/trace` package
   - Monitor work distribution

## 🎯 Key Takeaways

1. **G-M-P Model** is the foundation of Go concurrency
2. **Work stealing** provides efficient load balancing
3. **Preemption** ensures fair scheduling
4. **Memory management** is automatic and efficient
5. **NUMA awareness** optimizes for modern hardware
6. **Scheduling decisions** happen at specific triggers
7. **Performance tuning** requires understanding the model

## 🚨 Common Pitfalls

1. **Goroutine Leaks:**
   - Always ensure goroutines can exit
   - Use context for cancellation
   - Monitor goroutine count

2. **False Sharing:**
   - Variables on same cache line
   - Use padding to separate variables
   - Consider `sync.Pool` for local data

3. **Scheduling Issues:**
   - Long-running CPU work blocks others
   - Use `runtime.Gosched()` for fairness
   - Consider breaking up work

4. **Memory Issues:**
   - Large goroutine stacks
   - Inefficient memory usage
   - GC pressure from many goroutines

## 📖 Further Reading

- Go Runtime Source Code
- Go Memory Model Specification
- CPU Architecture and Cache Systems
- Work Stealing Algorithms
- NUMA Programming

---

*This is GOD-LEVEL knowledge that separates good developers from concurrency masters!*
