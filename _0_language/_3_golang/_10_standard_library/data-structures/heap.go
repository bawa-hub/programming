package main

import (
	"container/heap"
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

type Process struct {
	PID        int
	Priority   int
	Memory     int
	StartTime  time.Time
}

func (p Process) String() string {
	return fmt.Sprintf("PID: %d, Priority: %d, Memory: %dMB", p.PID, p.Priority, p.Memory)
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

// Process heap for OS scheduling
type ProcessHeap []Process

func (h ProcessHeap) Len() int           { return len(h) }
func (h ProcessHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h ProcessHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ProcessHeap) Push(x interface{}) {
	*h = append(*h, x.(Process))
}

func (h *ProcessHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Updateable heap for dynamic priority changes
type UpdateableTask struct {
	ID       string
	Priority int
	Name     string
	Index    int // Index in heap for efficient updates
}

type UpdateableTaskHeap []*UpdateableTask

func (h UpdateableTaskHeap) Len() int           { return len(h) }
func (h UpdateableTaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h UpdateableTaskHeap) Swap(i, j int)      { 
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *UpdateableTaskHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*UpdateableTask)
	item.Index = n
	*h = append(*h, item)
}

func (h *UpdateableTaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*h = old[0 : n-1]
	return item
}

// Custom heap for absolute value comparison
type CustomHeap struct {
	items []int
	less  func(i, j int) bool
}

func (h CustomHeap) Len() int           { return len(h.items) }
func (h CustomHeap) Less(i, j int) bool { return h.less(i, j) }
func (h CustomHeap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *CustomHeap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

func (h *CustomHeap) Pop() interface{} {
	old := h.items
	n := len(old)
	x := old[n-1]
	h.items = old[0 : n-1]
	return x
}

// ListNode for merge K sorted lists
type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println("🚀 Go heap Package Mastery Examples")
	fmt.Println("====================================")

	// 1. Basic Min Heap Operations
	fmt.Println("\n1. Basic Min Heap Operations:")
	
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

	// 2. Task Priority Queue
	fmt.Println("\n2. Task Priority Queue:")
	
	// Create task heap
	taskHeap := &TaskHeap{
		{3, "Write report", 2 * time.Hour},
		{1, "Fix bug", 30 * time.Minute},
		{5, "Code review", 1 * time.Hour},
		{2, "Update docs", 45 * time.Minute},
		{4, "Deploy app", 15 * time.Minute},
	}
	
	heap.Init(taskHeap)
	fmt.Println("Processing tasks by priority:")
	for taskHeap.Len() > 0 {
		task := heap.Pop(taskHeap).(Task)
		fmt.Printf("  %s\n", task)
	}

	// 3. Max Heap Operations
	fmt.Println("\n3. Max Heap Operations:")
	
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

	// 4. Process Scheduling Simulation
	fmt.Println("\n4. Process Scheduling Simulation:")
	
	// Create process heap
	processHeap := &ProcessHeap{
		{1, 3, 512, time.Now()},
		{2, 1, 256, time.Now()},
		{3, 5, 1024, time.Now()},
		{4, 2, 128, time.Now()},
		{5, 4, 768, time.Now()},
	}
	
	heap.Init(processHeap)
	fmt.Println("Scheduling processes by priority:")
	for processHeap.Len() > 0 {
		process := heap.Pop(processHeap).(Process)
		fmt.Printf("  %s\n", process)
	}

	// 5. Dynamic Heap Operations
	fmt.Println("\n5. Dynamic Heap Operations:")
	
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

	// 6. Updateable Heap
	fmt.Println("\n6. Updateable Heap:")
	
	// Create updateable task heap
	updateableHeap := &UpdateableTaskHeap{
		{"task1", 3, "Write code", 0},
		{"task2", 1, "Fix bug", 1},
		{"task3", 5, "Review code", 2},
		{"task4", 2, "Update docs", 3},
	}
	
	heap.Init(updateableHeap)
	fmt.Println("Initial tasks:")
	for i, task := range *updateableHeap {
		fmt.Printf("  [%d] %s (priority: %d)\n", i, task.Name, task.Priority)
	}
	
	// Update task priority
	(*updateableHeap)[1].Priority = 4
	heap.Fix(updateableHeap, 1)
	fmt.Println("\nAfter updating task2 priority to 4:")
	for i, task := range *updateableHeap {
		fmt.Printf("  [%d] %s (priority: %d)\n", i, task.Name, task.Priority)
	}

	// 7. Heap with Custom Comparison
	fmt.Println("\n7. Heap with Custom Comparison:")
	
	// Create custom heap with absolute value comparison
	customHeap := &CustomHeap{
		items: []int{},
		less: func(i, j int) bool {
			// This will be set properly after creation
			return false
		},
	}
	
	// Set up the comparison function after creation
	customHeap.less = func(i, j int) bool {
		absI := customHeap.items[i]
		absJ := customHeap.items[j]
		if absI < 0 {
			absI = -absI
		}
		if absJ < 0 {
			absJ = -absJ
		}
		return absI < absJ
	}
	
	heap.Init(customHeap)
	
	// Add elements with negative values
	testValues := []int{3, -1, 4, -2, 5, -3}
	for _, val := range testValues {
		heap.Push(customHeap, val)
	}
	
	fmt.Println("Custom heap (sorted by absolute value):")
	for customHeap.Len() > 0 {
		val := heap.Pop(customHeap).(int)
		fmt.Printf("  %d\n", val)
	}

	// 8. Heap Performance Analysis
	fmt.Println("\n8. Heap Performance Analysis:")
	
	// Test heap operations with different sizes
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

	// 9. Heap Sort Implementation
	fmt.Println("\n9. Heap Sort Implementation:")
	
	// Implement heap sort
	heapSort := func(arr []int) {
		h := &IntHeap{}
		heap.Init(h)
		
		// Build heap
		for _, val := range arr {
			heap.Push(h, val)
		}
		
		// Extract elements in sorted order
		for i := 0; i < len(arr); i++ {
			arr[i] = heap.Pop(h).(int)
		}
	}
	
	// Test heap sort
	unsortedArray := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Printf("Original array: %v\n", unsortedArray)
	
	heapSort(unsortedArray)
	fmt.Printf("Heap sorted: %v\n", unsortedArray)

	// 10. Median Finding with Two Heaps
	fmt.Println("\n10. Median Finding with Two Heaps:")
	
	// Implement median finder using two heaps
	type MedianFinder struct {
		maxHeap *MaxHeap // Left side (smaller half)
		minHeap *IntHeap // Right side (larger half)
	}
	
	NewMedianFinder := func() *MedianFinder {
		return &MedianFinder{
			maxHeap: &MaxHeap{},
			minHeap: &IntHeap{},
		}
	}
	
	// Add number method
	AddNumber := func(mf *MedianFinder, num int) {
		if mf.maxHeap.Len() == 0 || num <= (*mf.maxHeap)[0] {
			heap.Push(mf.maxHeap, num)
		} else {
			heap.Push(mf.minHeap, num)
		}
		
		// Balance heaps
		if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
			val := heap.Pop(mf.maxHeap).(int)
			heap.Push(mf.minHeap, val)
		} else if mf.minHeap.Len() > mf.maxHeap.Len()+1 {
			val := heap.Pop(mf.minHeap).(int)
			heap.Push(mf.maxHeap, val)
		}
	}
	
	// Find median method
	FindMedian := func(mf *MedianFinder) float64 {
		if mf.maxHeap.Len() == mf.minHeap.Len() {
			return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
		} else if mf.maxHeap.Len() > mf.minHeap.Len() {
			return float64((*mf.maxHeap)[0])
		} else {
			return float64((*mf.minHeap)[0])
		}
	}
	
	// Test median finder
	medianFinder := NewMedianFinder()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	fmt.Println("Adding numbers and finding median:")
	for _, num := range numbers {
		AddNumber(medianFinder, num)
		median := FindMedian(medianFinder)
		fmt.Printf("  Added %d, median: %.1f\n", num, median)
	}

	// 11. Top K Elements
	fmt.Println("\n11. Top K Elements:")
	
	// Find top K elements using heap
	findTopK := func(arr []int, k int) []int {
		if k <= 0 || k > len(arr) {
			return nil
		}
		
		// Use min heap to keep only top K elements
		h := &IntHeap{}
		heap.Init(h)
		
		for _, num := range arr {
			if h.Len() < k {
				heap.Push(h, num)
			} else if num > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
		
		// Extract top K elements
		result := make([]int, k)
		for i := k - 1; i >= 0; i-- {
			result[i] = heap.Pop(h).(int)
		}
		
		return result
	}
	
	// Test top K elements
	testArray := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	k := 3
	topK := findTopK(testArray, k)
	fmt.Printf("Array: %v\n", testArray)
	fmt.Printf("Top %d elements: %v\n", k, topK)

	// 12. Merge K Sorted Lists
	fmt.Println("\n12. Merge K Sorted Lists:")
	
	// Merge K sorted lists using heap
	
	// Create sample sorted lists
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}
	
	lists := []*ListNode{list1, list2, list3}
	
	// Merge using heap
	listHeap := &ListNodeHeap{}
	heap.Init(listHeap)
	
	// Add first node from each list
	for _, list := range lists {
		if list != nil {
			heap.Push(listHeap, list)
		}
	}
	
	// Merge lists
	var result []int
	for listHeap.Len() > 0 {
		node := heap.Pop(listHeap).(*ListNode)
		result = append(result, node.Val)
		if node.Next != nil {
			heap.Push(listHeap, node.Next)
		}
	}
	
	fmt.Printf("Merged lists: %v\n", result)

	// 13. Heap with Duplicate Handling
	fmt.Println("\n13. Heap with Duplicate Handling:")
	
	// Count frequency and use heap for top frequent elements
	frequencyHeap := &IntHeap{}
	heap.Init(frequencyHeap)
	
	// Count frequencies
	freq := make(map[int]int)
	testNumbers := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}
	
	for _, num := range testNumbers {
		freq[num]++
	}
	
	// Add frequencies to heap
	for num, count := range freq {
		for i := 0; i < count; i++ {
			heap.Push(frequencyHeap, num)
		}
	}
	
	fmt.Printf("Numbers: %v\n", testNumbers)
	fmt.Printf("Frequency heap: %v\n", *frequencyHeap)
	
	// Find most frequent element
	mostFrequent := heap.Pop(frequencyHeap).(int)
	fmt.Printf("Most frequent element: %d\n", mostFrequent)

	// 14. Heap Memory Management
	fmt.Println("\n14. Heap Memory Management:")
	
	// Demonstrate heap memory usage
	largeHeap := &IntHeap{}
	heap.Init(largeHeap)
	
	// Add many elements
	for i := 0; i < 1000; i++ {
		heap.Push(largeHeap, i)
	}
	
	fmt.Printf("Large heap length: %d\n", largeHeap.Len())
	
	// Clear heap
	for largeHeap.Len() > 0 {
		heap.Pop(largeHeap)
	}
	
	fmt.Printf("After clearing, heap length: %d\n", largeHeap.Len())

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

	fmt.Println("\n🎉 heap Package Mastery Complete!")
}
