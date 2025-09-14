# sync/atomic Package - Atomic Operations ⚛️

The `sync/atomic` package provides low-level atomic memory primitives useful for implementing synchronization algorithms. These operations are lock-free and provide atomic operations on basic types.

## 🎯 Key Concepts

### 1. **Atomic Operations**
- `AddInt64()` - Atomic addition
- `LoadInt64()` - Atomic load
- `StoreInt64()` - Atomic store
- `SwapInt64()` - Atomic swap
- `CompareAndSwapInt64()` - Atomic compare and swap
- `AddUint64()` - Atomic addition for unsigned
- `LoadUint64()` - Atomic load for unsigned
- `StoreUint64()` - Atomic store for unsigned

### 2. **Pointer Operations**
- `LoadPointer()` - Atomic load pointer
- `StorePointer()` - Atomic store pointer
- `SwapPointer()` - Atomic swap pointer
- `CompareAndSwapPointer()` - Atomic compare and swap pointer

### 3. **Value Operations**
- `Value` - Generic atomic value
- `Load()` - Load value
- `Store()` - Store value
- `Swap()` - Swap value
- `CompareAndSwap()` - Compare and swap value

### 4. **Atomic Types**
- `Int64` - Atomic int64
- `Uint64` - Atomic uint64
- `Uint32` - Atomic uint32
- `Int32` - Atomic int32
- `Bool` - Atomic bool
- `Pointer` - Atomic pointer

### 5. **Memory Ordering**
- Sequential consistency
- Acquire-release semantics
- Relaxed ordering
- Memory barriers

### 6. **Use Cases**
- Counters
- Flags
- Pointers
- State management
- Lock-free data structures

## 🚀 Common Patterns

### Basic Atomic Operations
```go
var counter int64

// Atomic increment
atomic.AddInt64(&counter, 1)

// Atomic load
value := atomic.LoadInt64(&counter)

// Atomic store
atomic.StoreInt64(&counter, 100)

// Atomic compare and swap
swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
```

### Atomic Counter
```go
type AtomicCounter struct {
    value int64
}

func (c *AtomicCounter) Increment() int64 {
    return atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.value)
}
```

### Atomic Flag
```go
type AtomicFlag struct {
    value int32
}

func (f *AtomicFlag) Set() {
    atomic.StoreInt32(&f.value, 1)
}

func (f *AtomicFlag) IsSet() bool {
    return atomic.LoadInt32(&f.value) == 1
}
```

### Atomic Pointer
```go
type AtomicPointer struct {
    ptr unsafe.Pointer
}

func (p *AtomicPointer) Store(ptr unsafe.Pointer) {
    atomic.StorePointer(&p.ptr, ptr)
}

func (p *AtomicPointer) Load() unsafe.Pointer {
    return atomic.LoadPointer(&p.ptr)
}
```

## ⚠️ Common Pitfalls

1. **Race conditions** - Still possible with atomic operations
2. **Memory ordering** - Understand memory ordering guarantees
3. **Pointer arithmetic** - Be careful with pointer operations
4. **Type safety** - Use proper types for atomic operations
5. **Performance** - Atomic operations have overhead

## 🎯 Best Practices

1. **Use atomic types** - Prefer atomic types over raw operations
2. **Minimize atomic operations** - Use them only when necessary
3. **Understand ordering** - Know memory ordering guarantees
4. **Test thoroughly** - Atomic operations are hard to test
5. **Document assumptions** - Document memory ordering assumptions

## 🔍 Advanced Features

### Custom Atomic Operations
```go
func atomicIncrement(addr *int64) int64 {
    for {
        old := atomic.LoadInt64(addr)
        new := old + 1
        if atomic.CompareAndSwapInt64(addr, old, new) {
            return new
        }
    }
}
```

### Atomic State Machine
```go
type State int32

const (
    Idle State = iota
    Running
    Stopped
)

type AtomicState struct {
    state int32
}

func (s *AtomicState) SetState(newState State) {
    atomic.StoreInt32(&s.state, int32(newState))
}

func (s *AtomicState) GetState() State {
    return State(atomic.LoadInt32(&s.state))
}
```

### Lock-Free Stack
```go
type Node struct {
    value int
    next  *Node
}

type LockFreeStack struct {
    head unsafe.Pointer
}

func (s *LockFreeStack) Push(value int) {
    node := &Node{value: value}
    for {
        head := atomic.LoadPointer(&s.head)
        node.next = (*Node)(head)
        if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
            return
        }
    }
}
```

## 📚 Real-world Applications

1. **Counters** - Request counters, metrics
2. **Flags** - Feature flags, status flags
3. **Pointers** - Lock-free data structures
4. **State** - State machines
5. **Caches** - Atomic cache operations

## 🧠 Memory Tips

- **atomic** = **A**tomic **T**ype **O**perations **M**emory **I**nterface **C**ontrol
- **Add** = **A**dd value
- **Load** = **L**oad value
- **Store** = **S**tore value
- **Swap** = **S**wap value
- **CompareAndSwap** = **C**ompare **A**nd **S**wap
- **Pointer** = **P**ointer operations
- **Value** = **V**alue operations

Remember: The atomic package is your gateway to lock-free programming in Go! 🎯
