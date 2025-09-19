package main

import (
	"fmt"
	"strings"
)

// =============================================================================
// BASIC ITERATOR PATTERN
// =============================================================================

// Iterator interface
type Iterator interface {
	Next() interface{}
	HasNext() bool
	Current() interface{}
	Reset()
}

// Iterable interface (Collection)
type Iterable interface {
	CreateIterator() Iterator
	GetCount() int
	GetItem(index int) interface{}
}

// Concrete Collection - Array
type ArrayCollection struct {
	items []interface{}
}

func NewArrayCollection(items []interface{}) *ArrayCollection {
	return &ArrayCollection{items: items}
}

func (ac *ArrayCollection) Add(item interface{}) {
	ac.items = append(ac.items, item)
}

func (ac *ArrayCollection) GetItem(index int) interface{} {
	if index >= 0 && index < len(ac.items) {
		return ac.items[index]
	}
	return nil
}

func (ac *ArrayCollection) GetCount() int {
	return len(ac.items)
}

func (ac *ArrayCollection) CreateIterator() Iterator {
	return &ArrayIterator{
		collection: ac,
		index:      0,
	}
}

// Array Iterator
type ArrayIterator struct {
	collection *ArrayCollection
	index      int
}

func (ai *ArrayIterator) Next() interface{} {
	if ai.HasNext() {
		item := ai.collection.GetItem(ai.index)
		ai.index++
		return item
	}
	return nil
}

func (ai *ArrayIterator) HasNext() bool {
	return ai.index < ai.collection.GetCount()
}

func (ai *ArrayIterator) Current() interface{} {
	if ai.index > 0 && ai.index <= ai.collection.GetCount() {
		return ai.collection.GetItem(ai.index - 1)
	}
	return nil
}

func (ai *ArrayIterator) Reset() {
	ai.index = 0
}

// =============================================================================
// ADVANCED ITERATOR TYPES
// =============================================================================

// Bidirectional Iterator
type BidirectionalIterator interface {
	Next() interface{}
	Previous() interface{}
	HasNext() bool
	HasPrevious() bool
	Current() interface{}
	Reset()
	SetPosition(index int)
}

// Bidirectional Array Iterator
type BidirectionalArrayIterator struct {
	collection *ArrayCollection
	index      int
}

func (bai *BidirectionalArrayIterator) Next() interface{} {
	if bai.HasNext() {
		item := bai.collection.GetItem(bai.index)
		bai.index++
		return item
	}
	return nil
}

func (bai *BidirectionalArrayIterator) Previous() interface{} {
	if bai.HasPrevious() {
		bai.index--
		return bai.collection.GetItem(bai.index)
	}
	return nil
}

func (bai *BidirectionalArrayIterator) HasNext() bool {
	return bai.index < bai.collection.GetCount()
}

func (bai *BidirectionalArrayIterator) HasPrevious() bool {
	return bai.index > 0
}

func (bai *BidirectionalArrayIterator) Current() interface{} {
	if bai.index >= 0 && bai.index < bai.collection.GetCount() {
		return bai.collection.GetItem(bai.index)
	}
	return nil
}

func (bai *BidirectionalArrayIterator) Reset() {
	bai.index = 0
}

func (bai *BidirectionalArrayIterator) SetPosition(index int) {
	if index >= 0 && index <= bai.collection.GetCount() {
		bai.index = index
	}
}

// Reverse Iterator
type ReverseIterator interface {
	Next() interface{}
	HasNext() bool
	Current() interface{}
	Reset()
}

// Reverse Array Iterator
type ReverseArrayIterator struct {
	collection *ArrayCollection
	index      int
}

func (rai *ReverseArrayIterator) Next() interface{} {
	if rai.HasNext() {
		item := rai.collection.GetItem(rai.index)
		rai.index--
		return item
	}
	return nil
}

func (rai *ReverseArrayIterator) HasNext() bool {
	return rai.index >= 0
}

func (rai *ReverseArrayIterator) Current() interface{} {
	if rai.index >= 0 && rai.index < rai.collection.GetCount() {
		return rai.collection.GetItem(rai.index)
	}
	return nil
}

func (rai *ReverseArrayIterator) Reset() {
	rai.index = rai.collection.GetCount() - 1
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. LINKED LIST ITERATOR
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func (ll *LinkedList) Add(value interface{}) {
	newNode := &ListNode{Value: value, Next: nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.size++
}

func (ll *LinkedList) GetCount() int {
	return ll.size
}

func (ll *LinkedList) CreateIterator() Iterator {
	return &LinkedListIterator{
		list:  ll,
		current: ll.head,
	}
}

type LinkedListIterator struct {
	list    *LinkedList
	current *ListNode
}

func (lli *LinkedListIterator) Next() interface{} {
	if lli.HasNext() {
		value := lli.current.Value
		lli.current = lli.current.Next
		return value
	}
	return nil
}

func (lli *LinkedListIterator) HasNext() bool {
	return lli.current != nil
}

func (lli *LinkedListIterator) Current() interface{} {
	if lli.current != nil {
		return lli.current.Value
	}
	return nil
}

func (lli *LinkedListIterator) Reset() {
	lli.current = lli.list.head
}

// 2. TREE ITERATOR
type TreeNode struct {
	Value    interface{}
	Children []*TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTree(rootValue interface{}) *Tree {
	return &Tree{
		root: &TreeNode{Value: rootValue, Children: make([]*TreeNode, 0)},
	}
}

func (t *Tree) AddChild(parent *TreeNode, value interface{}) *TreeNode {
	child := &TreeNode{Value: value, Children: make([]*TreeNode, 0)}
	parent.Children = append(parent.Children, child)
	return child
}

func (t *Tree) CreateIterator() Iterator {
	return &TreeIterator{
		tree:  t,
		stack: []*TreeNode{t.root},
	}
}

type TreeIterator struct {
	tree  *Tree
	stack []*TreeNode
}

func (ti *TreeIterator) Next() interface{} {
	if ti.HasNext() {
		current := ti.stack[len(ti.stack)-1]
		ti.stack = ti.stack[:len(ti.stack)-1]
		
		// Add children to stack (reverse order for pre-order traversal)
		for i := len(current.Children) - 1; i >= 0; i-- {
			ti.stack = append(ti.stack, current.Children[i])
		}
		
		return current.Value
	}
	return nil
}

func (ti *TreeIterator) HasNext() bool {
	return len(ti.stack) > 0
}

func (ti *TreeIterator) Current() interface{} {
	if len(ti.stack) > 0 {
		return ti.stack[len(ti.stack)-1].Value
	}
	return nil
}

func (ti *TreeIterator) Reset() {
	ti.stack = []*TreeNode{ti.tree.root}
}

// 3. FILTERED ITERATOR
type FilteredIterator struct {
	iterator Iterator
	filter   func(interface{}) bool
	nextItem interface{}
	hasNext  bool
}

func NewFilteredIterator(iterator Iterator, filter func(interface{}) bool) *FilteredIterator {
	fi := &FilteredIterator{
		iterator: iterator,
		filter:   filter,
		nextItem: nil,
		hasNext:  false,
	}
	fi.findNext()
	return fi
}

func (fi *FilteredIterator) findNext() {
	fi.hasNext = false
	fi.nextItem = nil
	
	for fi.iterator.HasNext() {
		item := fi.iterator.Next()
		if fi.filter(item) {
			fi.nextItem = item
			fi.hasNext = true
			break
		}
	}
}

func (fi *FilteredIterator) Next() interface{} {
	if fi.HasNext() {
		item := fi.nextItem
		fi.findNext()
		return item
	}
	return nil
}

func (fi *FilteredIterator) HasNext() bool {
	return fi.hasNext
}

func (fi *FilteredIterator) Current() interface{} {
	return fi.nextItem
}

func (fi *FilteredIterator) Reset() {
	fi.iterator.Reset()
	fi.findNext()
}

// 4. RANGE ITERATOR
type RangeIterator struct {
	start, end, step int
	current          int
}

func NewRangeIterator(start, end, step int) *RangeIterator {
	return &RangeIterator{
		start:   start,
		end:     end,
		step:    step,
		current: start,
	}
}

func (ri *RangeIterator) Next() interface{} {
	if ri.HasNext() {
		value := ri.current
		ri.current += ri.step
		return value
	}
	return nil
}

func (ri *RangeIterator) HasNext() bool {
	if ri.step > 0 {
		return ri.current < ri.end
	} else {
		return ri.current > ri.end
	}
}

func (ri *RangeIterator) Current() interface{} {
	return ri.current
}

func (ri *RangeIterator) Reset() {
	ri.current = ri.start
}

// 5. INFINITE ITERATOR
type InfiniteIterator struct {
	generator func() interface{}
	current   interface{}
}

func NewInfiniteIterator(generator func() interface{}) *InfiniteIterator {
	return &InfiniteIterator{
		generator: generator,
		current:   nil,
	}
}

func (ii *InfiniteIterator) Next() interface{} {
	ii.current = ii.generator()
	return ii.current
}

func (ii *InfiniteIterator) HasNext() bool {
	return true // Always has next
}

func (ii *InfiniteIterator) Current() interface{} {
	return ii.current
}

func (ii *InfiniteIterator) Reset() {
	ii.current = nil
}

// =============================================================================
// ITERATOR UTILITIES
// =============================================================================

// Iterator Factory
type IteratorFactory struct{}

func NewIteratorFactory() *IteratorFactory {
	return &IteratorFactory{}
}

func (ifactory *IteratorFactory) CreateForwardIterator(collection Iterable) Iterator {
	return collection.CreateIterator()
}

func (ifactory *IteratorFactory) CreateReverseIterator(collection *ArrayCollection) ReverseIterator {
	return &ReverseArrayIterator{
		collection: collection,
		index:      collection.GetCount() - 1,
	}
}

func (ifactory *IteratorFactory) CreateBidirectionalIterator(collection *ArrayCollection) BidirectionalIterator {
	return &BidirectionalArrayIterator{
		collection: collection,
		index:      0,
	}
}

func (ifactory *IteratorFactory) CreateFilteredIterator(iterator Iterator, filter func(interface{}) bool) Iterator {
	return NewFilteredIterator(iterator, filter)
}

// Iterator Utilities
type IteratorUtils struct{}

func NewIteratorUtils() *IteratorUtils {
	return &IteratorUtils{}
}

func (iu *IteratorUtils) ToSlice(iterator Iterator) []interface{} {
	var result []interface{}
	for iterator.HasNext() {
		result = append(result, iterator.Next())
	}
	return result
}

func (iu *IteratorUtils) ForEach(iterator Iterator, action func(interface{})) {
	for iterator.HasNext() {
		action(iterator.Next())
	}
}

func (iu *IteratorUtils) Count(iterator Iterator) int {
	count := 0
	for iterator.HasNext() {
		iterator.Next()
		count++
	}
	return count
}

func (iu *IteratorUtils) Any(iterator Iterator, predicate func(interface{}) bool) bool {
	for iterator.HasNext() {
		if predicate(iterator.Next()) {
			return true
		}
	}
	return false
}

func (iu *IteratorUtils) All(iterator Iterator, predicate func(interface{}) bool) bool {
	for iterator.HasNext() {
		if !predicate(iterator.Next()) {
			return false
		}
	}
	return true
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== ITERATOR PATTERN DEMONSTRATION ===\n")

	// 1. BASIC ITERATOR
	fmt.Println("1. BASIC ITERATOR:")
	items := []interface{}{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	arrayCollection := NewArrayCollection(items)
	iterator := arrayCollection.CreateIterator()
	
	fmt.Println("Forward iteration:")
	for iterator.HasNext() {
		fmt.Printf("  %s\n", iterator.Next())
	}
	fmt.Println()

	// 2. BIDIRECTIONAL ITERATOR
	fmt.Println("2. BIDIRECTIONAL ITERATOR:")
	bidirectionalIterator := &BidirectionalArrayIterator{
		collection: arrayCollection,
		index:      0,
	}
	
	fmt.Println("Forward iteration:")
	for bidirectionalIterator.HasNext() {
		fmt.Printf("  %s\n", bidirectionalIterator.Next())
	}
	
	fmt.Println("Backward iteration:")
	for bidirectionalIterator.HasPrevious() {
		fmt.Printf("  %s\n", bidirectionalIterator.Previous())
	}
	fmt.Println()

	// 3. REVERSE ITERATOR
	fmt.Println("3. REVERSE ITERATOR:")
	reverseIterator := &ReverseArrayIterator{
		collection: arrayCollection,
		index:      arrayCollection.GetCount() - 1,
	}
	
	fmt.Println("Reverse iteration:")
	for reverseIterator.HasNext() {
		fmt.Printf("  %s\n", reverseIterator.Next())
	}
	fmt.Println()

	// 4. LINKED LIST ITERATOR
	fmt.Println("4. LINKED LIST ITERATOR:")
	linkedList := NewLinkedList()
	linkedList.Add("First")
	linkedList.Add("Second")
	linkedList.Add("Third")
	linkedList.Add("Fourth")
	
	linkedListIterator := linkedList.CreateIterator()
	fmt.Println("Linked list iteration:")
	for linkedListIterator.HasNext() {
		fmt.Printf("  %s\n", linkedListIterator.Next())
	}
	fmt.Println()

	// 5. TREE ITERATOR
	fmt.Println("5. TREE ITERATOR:")
	tree := NewTree("Root")
	child1 := tree.AddChild(tree.root, "Child 1")
	child2 := tree.AddChild(tree.root, "Child 2")
	tree.AddChild(child1, "Grandchild 1")
	tree.AddChild(child1, "Grandchild 2")
	tree.AddChild(child2, "Grandchild 3")
	
	treeIterator := tree.CreateIterator()
	fmt.Println("Tree pre-order traversal:")
	for treeIterator.HasNext() {
		fmt.Printf("  %s\n", treeIterator.Next())
	}
	fmt.Println()

	// 6. FILTERED ITERATOR
	fmt.Println("6. FILTERED ITERATOR:")
	numbers := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numberCollection := NewArrayCollection(numbers)
	numberIterator := numberCollection.CreateIterator()
	
	// Filter even numbers
	evenFilter := func(item interface{}) bool {
		return item.(int)%2 == 0
	}
	filteredIterator := NewFilteredIterator(numberIterator, evenFilter)
	
	fmt.Println("Even numbers only:")
	for filteredIterator.HasNext() {
		fmt.Printf("  %d\n", filteredIterator.Next())
	}
	fmt.Println()

	// 7. RANGE ITERATOR
	fmt.Println("7. RANGE ITERATOR:")
	rangeIterator := NewRangeIterator(0, 10, 2)
	fmt.Println("Range 0 to 10 step 2:")
	for rangeIterator.HasNext() {
		fmt.Printf("  %d\n", rangeIterator.Next())
	}
	fmt.Println()

	// 8. INFINITE ITERATOR
	fmt.Println("8. INFINITE ITERATOR (limited to 5 items):")
	counter := 0
	infiniteIterator := NewInfiniteIterator(func() interface{} {
		counter++
		return fmt.Sprintf("Item %d", counter)
	})
	
	fmt.Println("Infinite sequence (first 5 items):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %s\n", infiniteIterator.Next())
	}
	fmt.Println()

	// 9. ITERATOR UTILITIES
	fmt.Println("9. ITERATOR UTILITIES:")
	utils := NewIteratorUtils()
	testCollection := NewArrayCollection([]interface{}{1, 2, 3, 4, 5})
	testIterator := testCollection.CreateIterator()
	
	// Convert to slice
	slice := utils.ToSlice(testIterator)
	fmt.Printf("Converted to slice: %v\n", slice)
	
	// ForEach
	testIterator.Reset()
	fmt.Println("ForEach demonstration:")
	utils.ForEach(testIterator, func(item interface{}) {
		fmt.Printf("  Processing: %d\n", item)
	})
	
	// Count
	testIterator.Reset()
	count := utils.Count(testIterator)
	fmt.Printf("Count: %d\n", count)
	
	// Any
	testIterator.Reset()
	hasEven := utils.Any(testIterator, func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	fmt.Printf("Has even number: %t\n", hasEven)
	
	// All
	testIterator.Reset()
	allPositive := utils.All(testIterator, func(item interface{}) bool {
		return item.(int) > 0
	})
	fmt.Printf("All positive: %t\n", allPositive)
	fmt.Println()

	// 10. ITERATOR FACTORY
	fmt.Println("10. ITERATOR FACTORY:")
	factory := NewIteratorFactory()
	
	// Create different types of iterators
	forwardIter := factory.CreateForwardIterator(arrayCollection)
	reverseIter := factory.CreateReverseIterator(arrayCollection)
	bidirectionalIter := factory.CreateBidirectionalIterator(arrayCollection)
	
	fmt.Println("Forward iterator:")
	for forwardIter.HasNext() {
		fmt.Printf("  %s\n", forwardIter.Next())
	}
	
	fmt.Println("Reverse iterator:")
	for reverseIter.HasNext() {
		fmt.Printf("  %s\n", reverseIter.Next())
	}
	
	fmt.Println("Bidirectional iterator (forward then backward):")
	for bidirectionalIter.HasNext() {
		fmt.Printf("  %s\n", bidirectionalIter.Next())
	}
	for bidirectionalIter.HasPrevious() {
		fmt.Printf("  %s\n", bidirectionalIter.Previous())
	}
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
