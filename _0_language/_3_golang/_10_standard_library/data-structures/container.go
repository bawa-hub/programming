package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"math/rand"
	"time"
)

// Custom types for demonstration
type Task struct {
	Priority int
	Name     string
	Duration time.Duration
}

func (t Task) String() string {
	return fmt.Sprintf("%s (priority: %d, duration: %v)", t.Name, t.Priority, t.Duration)
}

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d, %s)", p.Name, p.Age, p.City)
}

// Heap implementations
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

// Max heap (reverse of min heap)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println("🚀 Go container Package Mastery Examples")
	fmt.Println("=========================================")

	// 1. Heap - Basic Operations
	fmt.Println("\n1. Heap - Basic Operations:")
	
	// Create and initialize heap
	h := &IntHeap{2, 1, 5, 3, 4}
	heap.Init(h)
	fmt.Printf("Initial heap: %v\n", *h)
	
	// Push new element
	heap.Push(h, 0)
	fmt.Printf("After pushing 0: %v\n", *h)
	
	// Pop minimum element
	min := heap.Pop(h).(int)
	fmt.Printf("Popped minimum: %d\n", min)
	fmt.Printf("Heap after pop: %v\n", *h)
	
	// Push more elements
	heap.Push(h, 6)
	heap.Push(h, 1)
	fmt.Printf("After pushing 6 and 1: %v\n", *h)

	// 2. Heap - Task Priority Queue
	fmt.Println("\n2. Heap - Task Priority Queue:")
	
	// Create task heap
	taskHeap := &TaskHeap{
		{3, "Write report", 2 * time.Hour},
		{1, "Fix bug", 30 * time.Minute},
		{5, "Code review", 1 * time.Hour},
		{2, "Update docs", 45 * time.Minute},
		{4, "Deploy app", 15 * time.Minute},
	}
	
	heap.Init(taskHeap)
	fmt.Println("Initial tasks:")
	for taskHeap.Len() > 0 {
		task := heap.Pop(taskHeap).(Task)
		fmt.Printf("  %s\n", task)
	}

	// 3. Heap - Max Heap
	fmt.Println("\n3. Heap - Max Heap:")
	
	// Create max heap
	maxHeap := &MaxHeap{3, 1, 4, 1, 5, 9, 2, 6}
	heap.Init(maxHeap)
	fmt.Printf("Max heap: %v\n", *maxHeap)
	
	// Pop maximum elements
	fmt.Println("Popping maximum elements:")
	for maxHeap.Len() > 0 {
		max := heap.Pop(maxHeap).(int)
		fmt.Printf("  %d\n", max)
	}

	// 4. Heap - Dynamic Operations
	fmt.Println("\n4. Heap - Dynamic Operations:")
	
	// Create heap and add elements dynamically
	dynamicHeap := &IntHeap{}
	heap.Init(dynamicHeap)
	
	// Add elements
	elements := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	for _, elem := range elements {
		heap.Push(dynamicHeap, elem)
		fmt.Printf("Added %d, heap: %v\n", elem, *dynamicHeap)
	}
	
	// Remove elements
	fmt.Println("Removing elements:")
	for dynamicHeap.Len() > 0 {
		min := heap.Pop(dynamicHeap).(int)
		fmt.Printf("  Removed %d, remaining: %v\n", min, *dynamicHeap)
	}

	// 5. List - Basic Operations
	fmt.Println("\n5. List - Basic Operations:")
	
	// Create new list
	l := list.New()
	
	// Add elements to front and back
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(0)
	l.PushBack(3)
	
	fmt.Println("List elements (forward):")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %v\n", e.Value)
	}
	
	fmt.Println("List elements (backward):")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("  %v\n", e.Value)
	}
	
	fmt.Printf("List length: %d\n", l.Len())

	// 6. List - Insertion and Removal
	fmt.Println("\n6. List - Insertion and Removal:")
	
	// Create list with people
	peopleList := list.New()
	
	// Add people
	peopleList.PushBack(Person{"Alice", 30, "NYC"})
	bob := peopleList.PushBack(Person{"Bob", 25, "LA"})
	peopleList.PushBack(Person{"Charlie", 35, "Chicago"})
	
	fmt.Println("Initial list:")
	for e := peopleList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}
	
	// Insert before Bob
	peopleList.InsertBefore(Person{"David", 28, "Boston"}, bob)
	fmt.Println("\nAfter inserting David before Bob:")
	for e := peopleList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}
	
	// Insert after Charlie
	peopleList.InsertAfter(Person{"Eve", 32, "Seattle"}, peopleList.Back())
	fmt.Println("\nAfter inserting Eve after Charlie:")
	for e := peopleList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}
	
	// Remove Bob
	peopleList.Remove(bob)
	fmt.Println("\nAfter removing Bob:")
	for e := peopleList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}

	// 7. List - Moving Elements
	fmt.Println("\n7. List - Moving Elements:")
	
	// Create list
	moveList := list.New()
	
	// Add elements
	moveList.PushBack("First")
	moveList.PushBack("Second")
	e3 := moveList.PushBack("Third")
	moveList.PushBack("Fourth")
	
	fmt.Println("Initial list:")
	for e := moveList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}
	
	// Move Third to front
	moveList.MoveToFront(e3)
	fmt.Println("\nAfter moving Third to front:")
	for e := moveList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}
	
	// Move First to back
	moveList.MoveToBack(moveList.Front())
	fmt.Println("\nAfter moving First to back:")
	for e := moveList.Front(); e != nil; e = e.Next() {
		fmt.Printf("  %s\n", e.Value)
	}

	// 8. List - LRU Cache Implementation
	fmt.Println("\n8. List - LRU Cache Implementation:")
	
	// Simple LRU cache using list
	type LRUCache struct {
		capacity int
		list     *list.List
		elements map[string]*list.Element
	}
	
	NewLRUCache := func(capacity int) *LRUCache {
		return &LRUCache{
			capacity: capacity,
			list:     list.New(),
			elements: make(map[string]*list.Element),
		}
	}
	
	// Add method
	Add := func(cache *LRUCache, key string, value interface{}) {
		if elem, exists := cache.elements[key]; exists {
			// Update existing element
			elem.Value = value
			cache.list.MoveToFront(elem)
		} else {
			// Add new element
			if cache.list.Len() >= cache.capacity {
				// Remove least recently used
				back := cache.list.Back()
				cache.list.Remove(back)
				delete(cache.elements, back.Value.(string))
			}
			elem := cache.list.PushFront(value)
			cache.elements[key] = elem
		}
	}
	
	// Get method
	Get := func(cache *LRUCache, key string) (interface{}, bool) {
		if elem, exists := cache.elements[key]; exists {
			cache.list.MoveToFront(elem)
			return elem.Value, true
		}
		return nil, false
	}
	
	// Test LRU cache
	cache := NewLRUCache(3)
	Add(cache, "key1", "value1")
	Add(cache, "key2", "value2")
	Add(cache, "key3", "value3")
	Add(cache, "key4", "value4") // This should remove key1
	
	fmt.Println("LRU Cache operations:")
	if val, ok := Get(cache, "key1"); ok {
		fmt.Printf("  key1: %v\n", val)
	} else {
		fmt.Println("  key1: not found (evicted)")
	}
	
	if val, ok := Get(cache, "key2"); ok {
		fmt.Printf("  key2: %v\n", val)
	} else {
		fmt.Println("  key2: not found")
	}

	// 9. Ring - Basic Operations
	fmt.Println("\n9. Ring - Basic Operations:")
	
	// Create ring with 5 elements
	r := ring.New(5)
	
	// Initialize ring values
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	
	fmt.Println("Ring elements:")
	r.Do(func(x interface{}) {
		fmt.Printf("  %v\n", x)
	})
	
	// Move around the ring
	fmt.Printf("Current position: %v\n", r.Value)
	r = r.Next()
	fmt.Printf("After Next(): %v\n", r.Value)
	r = r.Prev()
	fmt.Printf("After Prev(): %v\n", r.Value)
	r = r.Move(3)
	fmt.Printf("After Move(3): %v\n", r.Value)

	// 10. Ring - Ring Operations
	fmt.Println("\n10. Ring - Ring Operations:")
	
	// Create two rings
	r1 := ring.New(3)
	r2 := ring.New(2)
	
	// Initialize rings
	for i := 0; i < r1.Len(); i++ {
		r1.Value = fmt.Sprintf("R1-%d", i)
		r1 = r1.Next()
	}
	
	for i := 0; i < r2.Len(); i++ {
		r2.Value = fmt.Sprintf("R2-%d", i)
		r2 = r2.Next()
	}
	
	fmt.Println("Ring 1:")
	r1.Do(func(x interface{}) {
		fmt.Printf("  %v\n", x)
	})
	
	fmt.Println("Ring 2:")
	r2.Do(func(x interface{}) {
		fmt.Printf("  %v\n", x)
	})
	
	// Link rings
	r1.Link(r2)
	fmt.Println("After linking Ring 2 to Ring 1:")
	r1.Do(func(x interface{}) {
		fmt.Printf("  %v\n", x)
	})

	// 11. Ring - Circular Buffer
	fmt.Println("\n11. Ring - Circular Buffer:")
	
	// Implement circular buffer using ring
	type CircularBuffer struct {
		ring *ring.Ring
		size int
	}
	
	NewCircularBuffer := func(size int) *CircularBuffer {
		return &CircularBuffer{
			ring: ring.New(size),
			size: size,
		}
	}
	
	// Add method
	AddToBuffer := func(cb *CircularBuffer, value interface{}) {
		cb.ring.Value = value
		cb.ring = cb.ring.Next()
	}
	
	// Get all values
	GetAllValues := func(cb *CircularBuffer) []interface{} {
		values := make([]interface{}, 0, cb.size)
		cb.ring.Do(func(x interface{}) {
			if x != nil {
				values = append(values, x)
			}
		})
		return values
	}
	
	// Test circular buffer
	buffer := NewCircularBuffer(4)
	
	fmt.Println("Adding values to circular buffer:")
	for i := 1; i <= 6; i++ {
		AddToBuffer(buffer, i)
		fmt.Printf("  Added %d, buffer: %v\n", i, GetAllValues(buffer))
	}

	// 12. Ring - Josephus Problem
	fmt.Println("\n12. Ring - Josephus Problem:")
	
	// Josephus problem: eliminate every k-th person in a circle
	Josephus := func(n, k int) int {
		r := ring.New(n)
		
		// Initialize with people numbered 1 to n
		for i := 1; i <= n; i++ {
			r.Value = i
			r = r.Next()
		}
		
		// Eliminate every k-th person
		for r.Len() > 1 {
			// Move to the person before the one to eliminate
			for i := 0; i < k-1; i++ {
				r = r.Next()
			}
			// Remove the k-th person
			r = r.Prev()
			r.Unlink(1)
			r = r.Next()
		}
		
		return r.Value.(int)
	}
	
	// Test Josephus problem
	n, k := 7, 3
	survivor := Josephus(n, k)
	fmt.Printf("With %d people, eliminating every %d-th person, survivor is: %d\n", n, k, survivor)

	// 13. Heap - Performance Comparison
	fmt.Println("\n13. Heap - Performance Comparison:")
	
	// Compare heap operations with different sizes
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		// Test heap operations
		h := &IntHeap{}
		heap.Init(h)
		
		// Generate random numbers
		rand.Seed(time.Now().UnixNano())
		numbers := make([]int, size)
		for i := range numbers {
			numbers[i] = rand.Intn(1000)
		}
		
		// Measure push time
		start := time.Now()
		for _, num := range numbers {
			heap.Push(h, num)
		}
		pushTime := time.Since(start)
		
		// Measure pop time
		start = time.Now()
		for h.Len() > 0 {
			heap.Pop(h)
		}
		popTime := time.Since(start)
		
		fmt.Printf("Size %d: Push time: %v, Pop time: %v\n", size, pushTime, popTime)
	}

	// 14. List - Performance Comparison
	fmt.Println("\n14. List - Performance Comparison:")
	
	// Compare list operations
	operations := []int{100, 1000, 10000}
	
	for _, ops := range operations {
		// Test list operations
		l := list.New()
		
		// Measure push back time
		start := time.Now()
		for i := 0; i < ops; i++ {
			l.PushBack(i)
		}
		pushBackTime := time.Since(start)
		
		// Measure push front time
		start = time.Now()
		for i := 0; i < ops; i++ {
			l.PushFront(i)
		}
		pushFrontTime := time.Since(start)
		
		fmt.Printf("Operations %d: PushBack: %v, PushFront: %v\n", ops, pushBackTime, pushFrontTime)
	}

	// 15. Advanced Heap Operations
	fmt.Println("\n15. Advanced Heap Operations:")
	
	// Create heap with tasks
	advancedTaskHeap := &TaskHeap{
		{3, "Task A", 1 * time.Hour},
		{1, "Task B", 30 * time.Minute},
		{5, "Task C", 2 * time.Hour},
		{2, "Task D", 45 * time.Minute},
	}
	
	heap.Init(advancedTaskHeap)
	fmt.Println("Initial tasks:")
	for i, task := range *advancedTaskHeap {
		fmt.Printf("  [%d] %s\n", i, task)
	}
	
	// Update task priority
	(*advancedTaskHeap)[1].Priority = 4
	heap.Fix(advancedTaskHeap, 1)
	fmt.Println("\nAfter updating task at index 1:")
	for i, task := range *advancedTaskHeap {
		fmt.Printf("  [%d] %s\n", i, task)
	}
	
	// Remove task at index 2
	removed := heap.Remove(advancedTaskHeap, 2)
	fmt.Printf("\nRemoved task: %s\n", removed)
	fmt.Println("Remaining tasks:")
	for i, task := range *advancedTaskHeap {
		fmt.Printf("  [%d] %s\n", i, task)
	}

	fmt.Println("\n🎉 container Package Mastery Complete!")
}
