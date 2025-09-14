# container Package - Container Data Structures 📦

The `container` package provides implementations of heap, list, and ring data structures. These are essential for advanced data manipulation and algorithm implementation.

## 🎯 Key Concepts

### 1. **Heap Package**
- `Interface` - Interface for heap operations
- `Init()` - Initialize heap
- `Push()` - Add element to heap
- `Pop()` - Remove and return minimum element
- `Remove()` - Remove element at index
- `Fix()` - Fix heap after element change

### 2. **List Package**
- `List` - Doubly linked list
- `Element` - List element
- `New()` - Create new list
- `PushFront()` - Add to front
- `PushBack()` - Add to back
- `InsertBefore()` - Insert before element
- `InsertAfter()` - Insert after element
- `Remove()` - Remove element
- `MoveToFront()` - Move to front
- `MoveToBack()` - Move to back

### 3. **Ring Package**
- `Ring` - Circular list
- `New()` - Create new ring
- `Next()` - Next element
- `Prev()` - Previous element
- `Move()` - Move n steps
- `Link()` - Link rings
- `Unlink()` - Unlink elements
- `Do()` - Execute function for each element

### 4. **Heap Interface Methods**
- `Len()` - Length of heap
- `Less(i, j int) bool` - Compare elements
- `Swap(i, j int)` - Swap elements
- `Push(x interface{})` - Add element
- `Pop() interface{}` - Remove minimum

### 5. **List Operations**
- `Front()` - First element
- `Back()` - Last element
- `Len()` - Length of list
- `InsertBefore()` - Insert before
- `InsertAfter()` - Insert after
- `Remove()` - Remove element

### 6. **Ring Operations**
- `Len()` - Length of ring
- `Next()` - Next element
- `Prev()` - Previous element
- `Move()` - Move n steps
- `Link()` - Link rings
- `Unlink()` - Unlink elements

## 🚀 Common Patterns

### Heap Usage
```go
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
min := heap.Pop(h).(int)
```

### List Usage
```go
list := list.New()
list.PushBack(1)
list.PushFront(2)
for e := list.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

### Ring Usage
```go
r := ring.New(5)
for i := 0; i < r.Len(); i++ {
    r.Value = i
    r = r.Next()
}
```

## ⚠️ Common Pitfalls

1. **Heap not initialized** - Always call heap.Init()
2. **List element reuse** - Don't reuse removed elements
3. **Ring empty check** - Check if ring is empty before operations
4. **Type assertions** - Always check type assertions
5. **Memory leaks** - Remove elements when no longer needed

## 🎯 Best Practices

1. **Initialize heaps** - Always call heap.Init()
2. **Check list length** - Verify list is not empty
3. **Use appropriate container** - Choose right data structure
4. **Handle errors** - Check for nil pointers
5. **Clean up resources** - Remove unused elements

## 🔍 Advanced Features

### Custom Heap
```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

### List with Custom Types
```go
type Person struct {
    Name string
    Age  int
}

list := list.New()
list.PushBack(Person{"Alice", 30})
```

### Ring with Operations
```go
r := ring.New(3)
r.Do(func(x interface{}) {
    fmt.Println(x)
})
```

## 📚 Real-world Applications

1. **Priority Queues** - Task scheduling, event processing
2. **LRU Cache** - Least recently used cache implementation
3. **Circular Buffers** - Data streaming, ring buffers
4. **Graph Algorithms** - BFS, DFS implementations
5. **Data Processing** - Efficient data manipulation

## 🧠 Memory Tips

- **container** = **C**ontainer **O**perations **N**etwork **T**oolkit **A**lgorithm **I**nterface **N**etwork **E**ngine **R**esource
- **heap** = **H**eap operations
- **list** = **L**inked list
- **ring** = **R**ing buffer
- **Push** = **P**ush element
- **Pop** = **P**op element
- **Next** = **N**ext element
- **Prev** = **P**revious element

Remember: The container package is your gateway to advanced data structures in Go! 🎯
