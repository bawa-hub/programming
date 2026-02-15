# LRU Cache System Design

## Problem Statement
Design a Least Recently Used (LRU) cache system that supports get and put operations in O(1) time complexity. The cache should have a maximum capacity and evict the least recently used item when the capacity is exceeded.

## Requirements Analysis

### Functional Requirements
1. **Cache Operations**
   - `get(key)`: Retrieve value by key, return -1 if not found
   - `put(key, value)`: Insert or update key-value pair
   - `delete(key)`: Remove key-value pair from cache
   - `clear()`: Remove all items from cache

2. **Capacity Management**
   - Fixed maximum capacity
   - Automatic eviction when capacity exceeded
   - LRU eviction policy (least recently used item removed first)

3. **Performance Requirements**
   - O(1) time complexity for get and put operations
   - O(1) space complexity for cache operations
   - Thread-safe operations for concurrent access

4. **Cache Statistics**
   - Hit rate and miss rate tracking
   - Cache size and capacity monitoring
   - Access frequency statistics

### Non-Functional Requirements
1. **Performance**: O(1) operations for get and put
2. **Scalability**: Support for large cache sizes
3. **Reliability**: Thread-safe concurrent access
4. **Memory Efficiency**: Optimal memory usage

## Core Entities

### 1. Cache Node
- **Attributes**: key, value, prev, next pointers
- **Behavior**: Link/unlink operations for doubly linked list

### 2. LRU Cache
- **Attributes**: capacity, size, head, tail, cache map
- **Behavior**: get, put, delete, evict operations

### 3. Cache Statistics
- **Attributes**: hits, misses, total_requests
- **Behavior**: Calculate hit rate, miss rate

## Design Patterns Used

### 1. Doubly Linked List Pattern
- Efficient O(1) insertion and deletion
- Maintains access order for LRU policy
- Head represents most recently used, tail represents least recently used

### 2. HashMap Pattern
- O(1) key lookup for cache operations
- Maps keys to cache nodes for fast access

### 3. Strategy Pattern
- Different eviction policies (LRU, LFU, FIFO)
- Configurable cache behavior

### 4. Observer Pattern
- Cache hit/miss notifications
- Statistics tracking and reporting

### 5. Template Method Pattern
- Standardized cache operation flow
- Consistent error handling and logging

## Data Structure Design

### Doubly Linked List + HashMap
```
HashMap: key -> Node
Doubly Linked List: head <-> node1 <-> node2 <-> ... <-> tail

Most Recently Used (head) -> Least Recently Used (tail)
```

### Node Structure
```go
type Node struct {
    key   string
    value interface{}
    prev  *Node
    next  *Node
}
```

### Cache Structure
```go
type LRUCache struct {
    capacity int
    size     int
    cache    map[string]*Node
    head     *Node
    tail     *Node
    stats    *CacheStats
    mu       sync.RWMutex
}
```

## Key Design Decisions

### 1. Data Structure Choice
- **Doubly Linked List**: O(1) insertion/deletion at any position
- **HashMap**: O(1) key lookup
- **Combination**: Achieves O(1) for both get and put operations

### 2. LRU Implementation
- **Move to Head**: When accessed, move node to head of list
- **Evict from Tail**: When capacity exceeded, remove tail node
- **Update Pointers**: Maintain doubly linked list integrity

### 3. Thread Safety
- **Read-Write Mutex**: Allow concurrent reads, exclusive writes
- **Fine-grained Locking**: Minimize lock contention
- **Atomic Operations**: For statistics and counters

### 4. Memory Management
- **Node Pool**: Reuse nodes to reduce GC pressure
- **Lazy Initialization**: Create nodes only when needed
- **Memory Monitoring**: Track memory usage and growth

## API Design

### Core Operations
```go
// Basic cache operations
func (c *LRUCache) Get(key string) (interface{}, bool)
func (c *LRUCache) Put(key string, value interface{})
func (c *LRUCache) Delete(key string) bool
func (c *LRUCache) Clear()

// Cache management
func (c *LRUCache) Size() int
func (c *LRUCache) Capacity() int
func (c *LRUCache) IsFull() bool
func (c *LRUCache) IsEmpty() bool

// Statistics
func (c *LRUCache) GetStats() *CacheStats
func (c *LRUCache) HitRate() float64
func (c *LRUCache) MissRate() float64
```

### Advanced Operations
```go
// Cache inspection
func (c *LRUCache) Keys() []string
func (c *LRUCache) Values() []interface{}
func (c *LRUCache) Contains(key string) bool

// Batch operations
func (c *LRUCache) PutAll(items map[string]interface{})
func (c *LRUCache) GetAll(keys []string) map[string]interface{}

// Cache configuration
func (c *LRUCache) SetCapacity(capacity int)
func (c *LRUCache) Resize(newCapacity int)
```

## Implementation Details

### 1. Get Operation
```go
func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    if node, exists := c.cache[key]; exists {
        // Move to head (most recently used)
        c.moveToHead(node)
        c.stats.RecordHit()
        return node.value, true
    }
    
    c.stats.RecordMiss()
    return nil, false
}
```

### 2. Put Operation
```go
func (c *LRUCache) Put(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    if node, exists := c.cache[key]; exists {
        // Update existing node
        node.value = value
        c.moveToHead(node)
        return
    }
    
    // Create new node
    newNode := &Node{key: key, value: value}
    
    if c.size >= c.capacity {
        // Evict least recently used
        c.evictLRU()
    }
    
    // Add to head
    c.addToHead(newNode)
    c.cache[key] = newNode
    c.size++
}
```

### 3. Eviction Strategy
```go
func (c *LRUCache) evictLRU() {
    if c.tail == nil {
        return
    }
    
    // Remove from tail
    lastNode := c.tail
    c.removeNode(lastNode)
    delete(c.cache, lastNode.key)
    c.size--
}
```

### 4. Node Management
```go
func (c *LRUCache) addToHead(node *Node) {
    node.prev = nil
    node.next = c.head
    
    if c.head != nil {
        c.head.prev = node
    }
    
    c.head = node
    
    if c.tail == nil {
        c.tail = node
    }
}

func (c *LRUCache) removeNode(node *Node) {
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        c.head = node.next
    }
    
    if node.next != nil {
        node.next.prev = node.prev
    } else {
        c.tail = node.prev
    }
}
```

## Advanced Features

### 1. Cache Statistics
```go
type CacheStats struct {
    hits          int64
    misses        int64
    totalRequests int64
    evictions     int64
    mu            sync.RWMutex
}

func (s *CacheStats) HitRate() float64 {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    if s.totalRequests == 0 {
        return 0.0
    }
    return float64(s.hits) / float64(s.totalRequests)
}
```

### 2. Cache Warming
```go
func (c *LRUCache) WarmUp(items map[string]interface{}) {
    for key, value := range items {
        if c.size < c.capacity {
            c.Put(key, value)
        }
    }
}
```

### 3. Cache Persistence
```go
func (c *LRUCache) SaveToFile(filename string) error {
    // Serialize cache to file
    data := c.serialize()
    return ioutil.WriteFile(filename, data, 0644)
}

func (c *LRUCache) LoadFromFile(filename string) error {
    // Deserialize cache from file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    return c.deserialize(data)
}
```

### 4. Cache Monitoring
```go
type CacheMonitor struct {
    cache     *LRUCache
    metrics   chan CacheMetric
    stopChan  chan bool
}

type CacheMetric struct {
    Timestamp time.Time
    HitRate   float64
    Size      int
    Capacity  int
}

func (m *CacheMonitor) Start() {
    go m.collectMetrics()
}

func (m *CacheMonitor) collectMetrics() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            metric := CacheMetric{
                Timestamp: time.Now(),
                HitRate:   m.cache.HitRate(),
                Size:      m.cache.Size(),
                Capacity:  m.cache.Capacity(),
            }
            m.metrics <- metric
        case <-m.stopChan:
            return
        }
    }
}
```

## Performance Optimization

### 1. Memory Pool
```go
type NodePool struct {
    pool sync.Pool
}

func NewNodePool() *NodePool {
    return &NodePool{
        pool: sync.Pool{
            New: func() interface{} {
                return &Node{}
            },
        },
    }
}

func (p *NodePool) Get() *Node {
    return p.pool.Get().(*Node)
}

func (p *NodePool) Put(node *Node) {
    node.key = ""
    node.value = nil
    node.prev = nil
    node.next = nil
    p.pool.Put(node)
}
```

### 2. Batch Operations
```go
func (c *LRUCache) PutBatch(items map[string]interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    for key, value := range items {
        if c.size >= c.capacity {
            c.evictLRU()
        }
        
        if node, exists := c.cache[key]; exists {
            node.value = value
            c.moveToHead(node)
        } else {
            newNode := &Node{key: key, value: value}
            c.addToHead(newNode)
            c.cache[key] = newNode
            c.size++
        }
    }
}
```

### 3. Lazy Loading
```go
type LazyLRUCache struct {
    *LRUCache
    loader func(string) (interface{}, error)
}

func (c *LazyLRUCache) Get(key string) (interface{}, bool) {
    if value, exists := c.LRUCache.Get(key); exists {
        return value, true
    }
    
    // Load from external source
    if c.loader != nil {
        if value, err := c.loader(key); err == nil {
            c.Put(key, value)
            return value, true
        }
    }
    
    return nil, false
}
```

## Testing Strategy

### 1. Unit Tests
- Test basic get/put operations
- Test eviction when capacity exceeded
- Test LRU ordering maintenance
- Test edge cases (empty cache, single item)

### 2. Performance Tests
- Benchmark get/put operations
- Memory usage profiling
- Concurrent access testing
- Load testing with large datasets

### 3. Integration Tests
- Test with real-world data patterns
- Test cache persistence and recovery
- Test monitoring and statistics

### 4. Stress Tests
- High-frequency operations
- Memory pressure testing
- Long-running stability tests

## Use Cases

### 1. Web Caching
- HTTP response caching
- Database query result caching
- Session data caching

### 2. Application Caching
- Object caching in applications
- Configuration data caching
- User session caching

### 3. System Caching
- File system caching
- Network request caching
- API response caching

### 4. Database Caching
- Query result caching
- Connection pooling
- Metadata caching

## Interview Tips

### 1. Start Simple
- Begin with basic get/put operations
- Add capacity management
- Implement LRU eviction policy
- Add thread safety

### 2. Ask Clarifying Questions
- What is the expected capacity?
- What data types will be stored?
- Any specific performance requirements?
- Thread safety requirements?

### 3. Consider Edge Cases
- What happens with capacity 0?
- How to handle null values?
- What if key doesn't exist?
- Memory constraints?

### 4. Discuss Trade-offs
- Memory vs. performance
- Complexity vs. efficiency
- Synchronization vs. concurrency
- Persistence vs. speed

### 5. Show System Thinking
- Discuss scalability considerations
- Consider monitoring and metrics
- Think about error handling
- Plan for future enhancements

## Common Interview Questions

### 1. Basic Implementation
- "Implement an LRU cache with get and put operations"
- "How would you ensure O(1) time complexity?"
- "What data structures would you use?"

### 2. Advanced Features
- "How would you add statistics tracking?"
- "How would you make it thread-safe?"
- "How would you add persistence?"

### 3. Optimization
- "How would you optimize memory usage?"
- "How would you handle high-frequency operations?"
- "How would you add monitoring?"

### 4. System Design
- "How would you scale this to multiple servers?"
- "How would you handle cache invalidation?"
- "How would you add different eviction policies?"

## Conclusion

The LRU Cache System is an excellent example of a data structure problem that tests your understanding of:
- Hash maps and linked lists
- Time and space complexity analysis
- Thread safety and concurrency
- Performance optimization
- System design principles

The key is to start with a simple implementation and gradually add complexity while maintaining the core O(1) performance requirements. Focus on the data structure design first, then add thread safety, statistics, and advanced features.
