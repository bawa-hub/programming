# heap Package - Heap Operations 🏗️

The `heap` package provides heap operations for any type that implements heap.Interface. It's essential for priority queues, scheduling, and efficient data management.

## 🎯 Key Concepts

### 1. **Heap Interface**
- `Interface` - Interface for heap operations
- `Len()` - Length of heap
- `Less(i, j int) bool` - Compare elements at indices i and j
- `Swap(i, j int)` - Swap elements at indices i and j
- `Push(x interface{})` - Add element to heap
- `Pop() interface{}` - Remove and return minimum element

### 2. **Heap Functions**
- `Init()` - Initialize heap
- `Push()` - Add element to heap
- `Pop()` - Remove and return minimum element
- `Remove()` - Remove element at index
- `Fix()` - Fix heap after element change

### 3. **Heap Properties**
- **Min Heap**: Parent is always smaller than children
- **Max Heap**: Parent is always larger than children
- **Complete Binary Tree**: All levels filled except possibly last
- **Heap Property**: Maintained after each operation

### 4. **Common Operations**
- **Insert**: Add new element (O(log n))
- **Extract Min/Max**: Remove minimum/maximum (O(log n))
- **Peek**: View minimum/maximum without removing (O(1))
- **Update**: Change element priority (O(log n))
- **Delete**: Remove arbitrary element (O(log n))

### 5. **Heap Types**
- **Min Heap**: Smallest element at root
- **Max Heap**: Largest element at root
- **Binary Heap**: Most common implementation
- **Fibonacci Heap**: Advanced heap with better amortized performance

## 🚀 Common Patterns

### Basic Min Heap
```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

### Priority Queue
```go
type Task struct {
    Priority int
    Name     string
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
    *h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

## ⚠️ Common Pitfalls

1. **Not calling Init()** - Always initialize heap before use
2. **Inconsistent comparison** - Less function must be consistent
3. **Index out of bounds** - Check indices in Less and Swap
4. **Type assertions** - Always check type assertions
5. **Memory leaks** - Remove elements when no longer needed

## 🎯 Best Practices

1. **Initialize heap** - Always call heap.Init()
2. **Consistent comparison** - Ensure Less function is consistent
3. **Handle errors** - Check for nil pointers and type assertions
4. **Use appropriate heap type** - Choose min or max heap based on needs
5. **Consider performance** - Heap operations are O(log n)

## 🔍 Advanced Features

### Custom Heap with Update
```go
type UpdateableHeap struct {
    items []Item
    index map[string]int
}

func (h *UpdateableHeap) Update(item Item) {
    if idx, exists := h.index[item.Key]; exists {
        h.items[idx] = item
        heap.Fix(h, idx)
    } else {
        heap.Push(h, item)
    }
}
```

### Heap with Custom Comparison
```go
type CustomHeap struct {
    items []Item
    less  func(i, j int) bool
}

func (h CustomHeap) Less(i, j int) bool {
    return h.less(i, j)
}
```

## 📚 Real-world Applications

1. **Priority Queues** - Task scheduling, event processing
2. **Dijkstra's Algorithm** - Shortest path finding
3. **A* Search** - Pathfinding with heuristics
4. **Merge K Sorted Lists** - Efficient merging
5. **Median Finding** - Running median calculation

## 🧠 Memory Tips

- **heap** = **H**eap **E**ngine **A**lgorithm **P**rocessor
- **Init** = **I**nitialize heap
- **Push** = **P**ush element
- **Pop** = **P**op minimum
- **Fix** = **F**ix heap
- **Remove** = **R**emove element
- **Less** = **L**ess comparison
- **Swap** = **S**wap elements

Remember: The heap package is your gateway to efficient priority queues in Go! 🎯
