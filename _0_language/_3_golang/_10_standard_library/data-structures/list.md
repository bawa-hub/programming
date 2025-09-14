# list Package - Doubly Linked Lists 🔗

The `list` package implements a doubly linked list. It's essential for efficient insertion, deletion, and traversal operations.

## 🎯 Key Concepts

### 1. **List Structure**
- `List` - Doubly linked list type
- `Element` - List element type
- `New()` - Create new list
- `Len()` - Get list length
- `Front()` - Get first element
- `Back()` - Get last element

### 2. **Element Operations**
- `Next()` - Next element
- `Prev()` - Previous element
- `Value` - Element value
- `List()` - Get containing list

### 3. **Insertion Operations**
- `PushFront()` - Add to front
- `PushBack()` - Add to back
- `InsertBefore()` - Insert before element
- `InsertAfter()` - Insert after element

### 4. **Removal Operations**
- `Remove()` - Remove element
- `MoveToFront()` - Move to front
- `MoveToBack()` - Move to back

### 5. **Traversal Operations**
- `Front()` - First element
- `Back()` - Last element
- `Next()` - Next element
- `Prev()` - Previous element

## 🚀 Common Patterns

### Basic List Operations
```go
list := list.New()
list.PushBack(1)
list.PushFront(0)
list.PushBack(2)

for e := list.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
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
list.PushBack(Person{"Bob", 25})
```

### List Manipulation
```go
// Insert before
newElement := list.InsertBefore(Person{"Charlie", 35}, existingElement)

// Move element
list.MoveToFront(existingElement)
list.MoveToBack(existingElement)

// Remove element
list.Remove(existingElement)
```

## ⚠️ Common Pitfalls

1. **Nil pointer dereference** - Check for nil before operations
2. **Element reuse** - Don't reuse removed elements
3. **List modification during traversal** - Be careful when modifying during iteration
4. **Memory leaks** - Remove elements when no longer needed
5. **Type assertions** - Always check type assertions

## 🎯 Best Practices

1. **Check for nil** - Always check for nil before operations
2. **Use appropriate operations** - Choose right insertion/removal method
3. **Handle errors** - Check for nil pointers
4. **Clean up resources** - Remove unused elements
5. **Consider performance** - List operations are O(1) for insertion/deletion

## 🔍 Advanced Features

### List with Indexing
```go
type IndexedList struct {
    list  *list.List
    index map[int]*list.Element
}

func (il *IndexedList) Get(index int) interface{} {
    if elem, exists := il.index[index]; exists {
        return elem.Value
    }
    return nil
}
```

### List with Custom Operations
```go
type CustomList struct {
    list *list.List
}

func (cl *CustomList) InsertSorted(value int) {
    for e := cl.list.Front(); e != nil; e = e.Next() {
        if e.Value.(int) > value {
            cl.list.InsertBefore(value, e)
            return
        }
    }
    cl.list.PushBack(value)
}
```

## 📚 Real-world Applications

1. **LRU Cache** - Least recently used cache
2. **Undo/Redo** - Command pattern implementation
3. **Music Playlist** - Playlist management
4. **Browser History** - Navigation history
5. **Task Queue** - Task management system

## 🧠 Memory Tips

- **list** = **L**inked **I**nterface **S**tructure **T**oolkit
- **New** = **N**ew list
- **Push** = **P**ush element
- **Insert** = **I**nsert element
- **Remove** = **R**emove element
- **Move** = **M**ove element
- **Front** = **F**ront element
- **Back** = **B**ack element

Remember: The list package is your gateway to efficient linked list operations in Go! 🎯
