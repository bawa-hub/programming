package main

import (
	"fmt"
	"sort"
	"strings"
)

// 🎯 GENERIC TYPES MASTERY
// This file demonstrates comprehensive generic usage and patterns

// Generic constraints
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 | ~string
}

type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 | ~string
}

// Generic structs
type GenericContainer[T any] struct {
	Items []T
	Count int
}

type GenericNode[T any] struct {
	Value    T
	Children []*GenericNode[T]
	Parent   *GenericNode[T]
}

type GenericPair[K, V any] struct {
	Key   K
	Value V
}

type GenericStack[T any] struct {
	items []T
}

type GenericQueue[T any] struct {
	items []T
}

type GenericSet[T comparable] struct {
	items map[T]bool
}

type GenericMapType[K comparable, V any] struct {
	items map[K]V
}

// Generic interfaces
type GenericComparable[T any] interface {
	Compare(other T) int
}

type GenericStringer[T any] interface {
	String() string
}

type GenericCloner[T any] interface {
	Clone() T
}

// Generic functions
func GenericMax[T Comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func GenericMin[T Comparable](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func GenericSum[T Numeric](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func GenericAverage[T Numeric](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}
	sum := GenericSum(slice)
	return float64(sum) / float64(len(slice))
}

func GenericContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GenericIndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func GenericFilter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func GenericMap[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

func GenericReduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func GenericSort[T Comparable](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func GenericReverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

func GenericUnique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func GenericChunk[T any](slice []T, size int) [][]T {
	var result [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

func GenericFlatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Generic type methods
func (gc *GenericContainer[T]) Add(item T) {
	gc.Items = append(gc.Items, item)
	gc.Count++
}

func (gc *GenericContainer[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(gc.Items) {
		var zero T
		return zero, false
	}
	return gc.Items[index], true
}

func (gc *GenericContainer[T]) Remove(index int) bool {
	if index < 0 || index >= len(gc.Items) {
		return false
	}
	gc.Items = append(gc.Items[:index], gc.Items[index+1:]...)
	gc.Count--
	return true
}

func (gc *GenericContainer[T]) Clear() {
	gc.Items = nil
	gc.Count = 0
}

func (gc *GenericContainer[T]) Size() int {
	return gc.Count
}

func (gc *GenericContainer[T]) IsEmpty() bool {
	return gc.Count == 0
}

// Generic Stack methods
func (gs *GenericStack[T]) Push(item T) {
	gs.items = append(gs.items, item)
}

func (gs *GenericStack[T]) Pop() (T, bool) {
	if len(gs.items) == 0 {
		var zero T
		return zero, false
	}
	index := len(gs.items) - 1
	item := gs.items[index]
	gs.items = gs.items[:index]
	return item, true
}

func (gs *GenericStack[T]) Peek() (T, bool) {
	if len(gs.items) == 0 {
		var zero T
		return zero, false
	}
	return gs.items[len(gs.items)-1], true
}

func (gs *GenericStack[T]) Size() int {
	return len(gs.items)
}

func (gs *GenericStack[T]) IsEmpty() bool {
	return len(gs.items) == 0
}

// Generic Queue methods
func (gq *GenericQueue[T]) Enqueue(item T) {
	gq.items = append(gq.items, item)
}

func (gq *GenericQueue[T]) Dequeue() (T, bool) {
	if len(gq.items) == 0 {
		var zero T
		return zero, false
	}
	item := gq.items[0]
	gq.items = gq.items[1:]
	return item, true
}

func (gq *GenericQueue[T]) Front() (T, bool) {
	if len(gq.items) == 0 {
		var zero T
		return zero, false
	}
	return gq.items[0], true
}

func (gq *GenericQueue[T]) Size() int {
	return len(gq.items)
}

func (gq *GenericQueue[T]) IsEmpty() bool {
	return len(gq.items) == 0
}

// Generic Set methods
func NewGenericSet[T comparable]() *GenericSet[T] {
	return &GenericSet[T]{
		items: make(map[T]bool),
	}
}

func (gs *GenericSet[T]) Add(item T) {
	gs.items[item] = true
}

func (gs *GenericSet[T]) Remove(item T) {
	delete(gs.items, item)
}

func (gs *GenericSet[T]) Contains(item T) bool {
	return gs.items[item]
}

func (gs *GenericSet[T]) Size() int {
	return len(gs.items)
}

func (gs *GenericSet[T]) IsEmpty() bool {
	return len(gs.items) == 0
}

func (gs *GenericSet[T]) ToSlice() []T {
	var result []T
	for item := range gs.items {
		result = append(result, item)
	}
	return result
}

func (gs *GenericSet[T]) Union(other *GenericSet[T]) *GenericSet[T] {
	result := NewGenericSet[T]()
	for item := range gs.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

func (gs *GenericSet[T]) Intersection(other *GenericSet[T]) *GenericSet[T] {
	result := NewGenericSet[T]()
	for item := range gs.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (gs *GenericSet[T]) Difference(other *GenericSet[T]) *GenericSet[T] {
	result := NewGenericSet[T]()
	for item := range gs.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Generic Map methods
func NewGenericMap[K comparable, V any]() *GenericMapType[K, V] {
	return &GenericMapType[K, V]{
		items: make(map[K]V),
	}
}

func (gm *GenericMapType[K, V]) Set(key K, value V) {
	gm.items[key] = value
}

func (gm *GenericMapType[K, V]) Get(key K) (V, bool) {
	value, exists := gm.items[key]
	return value, exists
}

func (gm *GenericMapType[K, V]) Delete(key K) {
	delete(gm.items, key)
}

func (gm *GenericMapType[K, V]) Contains(key K) bool {
	_, exists := gm.items[key]
	return exists
}

func (gm *GenericMapType[K, V]) Size() int {
	return len(gm.items)
}

func (gm *GenericMapType[K, V]) IsEmpty() bool {
	return len(gm.items) == 0
}

func (gm *GenericMapType[K, V]) Keys() []K {
	keys := make([]K, 0, len(gm.items))
	for key := range gm.items {
		keys = append(keys, key)
	}
	return keys
}

func (gm *GenericMapType[K, V]) Values() []V {
	values := make([]V, 0, len(gm.items))
	for _, value := range gm.items {
		values = append(values, value)
	}
	return values
}

func (gm *GenericMapType[K, V]) Clear() {
	gm.items = make(map[K]V)
}

// Generic Node methods
func NewGenericNode[T any](value T) *GenericNode[T] {
	return &GenericNode[T]{
		Value:    value,
		Children: make([]*GenericNode[T], 0),
	}
}

func (gn *GenericNode[T]) AddChild(child *GenericNode[T]) {
	child.Parent = gn
	gn.Children = append(gn.Children, child)
}

func (gn *GenericNode[T]) RemoveChild(child *GenericNode[T]) bool {
	for i, c := range gn.Children {
		if c == child {
			gn.Children = append(gn.Children[:i], gn.Children[i+1:]...)
			child.Parent = nil
			return true
		}
	}
	return false
}

func (gn *GenericNode[T]) IsLeaf() bool {
	return len(gn.Children) == 0
}

func (gn *GenericNode[T]) IsRoot() bool {
	return gn.Parent == nil
}

func (gn *GenericNode[T]) Depth() int {
	if gn.IsRoot() {
		return 0
	}
	return gn.Parent.Depth() + 1
}

func (gn *GenericNode[T]) Height() int {
	if gn.IsLeaf() {
		return 0
	}
	maxHeight := 0
	for _, child := range gn.Children {
		childHeight := child.Height()
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}
	return maxHeight + 1
}

// Generic Manager for CRUD operations
type GenericManager struct {
	Containers map[string]interface{}
	Stacks     map[string]interface{}
	Queues     map[string]interface{}
	Sets       map[string]interface{}
	Maps       map[string]interface{}
	Trees      map[string]interface{}
}

// NewGenericManager creates a new generic manager
func NewGenericManager() *GenericManager {
	return &GenericManager{
		Containers: make(map[string]interface{}),
		Stacks:     make(map[string]interface{}),
		Queues:     make(map[string]interface{}),
		Sets:       make(map[string]interface{}),
		Maps:       make(map[string]interface{}),
		Trees:      make(map[string]interface{}),
	}
}

// CRUD Operations for Generics

// Create - Initialize generic instances
func (gm *GenericManager) Create() {
	fmt.Println("🔧 Creating generic instances...")
	
	// Create containers
	gm.createContainers()
	
	// Create stacks
	gm.createStacks()
	
	// Create queues
	gm.createQueues()
	
	// Create sets
	gm.createSets()
	
	// Create maps
	gm.createMaps()
	
	// Create trees
	gm.createTrees()
	
	fmt.Println("✅ Generic instances created successfully")
}

// createContainers creates various generic containers
func (gm *GenericManager) createContainers() {
	// Int container
	intContainer := &GenericContainer[int]{}
	intContainer.Add(1)
	intContainer.Add(2)
	intContainer.Add(3)
	gm.Containers["int"] = intContainer
	
	// String container
	stringContainer := &GenericContainer[string]{}
	stringContainer.Add("apple")
	stringContainer.Add("banana")
	stringContainer.Add("cherry")
	gm.Containers["string"] = stringContainer
	
	// Person container
	personContainer := &GenericContainer[Person]{}
	personContainer.Add(Person{ID: 1, Name: "Alice", Age: 30})
	personContainer.Add(Person{ID: 2, Name: "Bob", Age: 25})
	gm.Containers["person"] = personContainer
	
	fmt.Println("Created generic containers")
}

// createStacks creates various generic stacks
func (gm *GenericManager) createStacks() {
	// Int stack
	intStack := &GenericStack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	gm.Stacks["int"] = intStack
	
	// String stack
	stringStack := &GenericStack[string]{}
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")
	gm.Stacks["string"] = stringStack
	
	fmt.Println("Created generic stacks")
}

// createQueues creates various generic queues
func (gm *GenericManager) createQueues() {
	// Int queue
	intQueue := &GenericQueue[int]{}
	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)
	gm.Queues["int"] = intQueue
	
	// String queue
	stringQueue := &GenericQueue[string]{}
	stringQueue.Enqueue("first")
	stringQueue.Enqueue("second")
	stringQueue.Enqueue("third")
	gm.Queues["string"] = stringQueue
	
	fmt.Println("Created generic queues")
}

// createSets creates various generic sets
func (gm *GenericManager) createSets() {
	// Int set
	intSet := NewGenericSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	intSet.Add(1) // Duplicate
	gm.Sets["int"] = intSet
	
	// String set
	stringSet := NewGenericSet[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")
	stringSet.Add("cherry")
	stringSet.Add("apple") // Duplicate
	gm.Sets["string"] = stringSet
	
	fmt.Println("Created generic sets")
}

// createMaps creates various generic maps
func (gm *GenericManager) createMaps() {
	// String to int map
	stringIntMap := NewGenericMap[string, int]()
	stringIntMap.Set("apple", 5)
	stringIntMap.Set("banana", 3)
	stringIntMap.Set("cherry", 8)
	gm.Maps["string_int"] = stringIntMap
	
	// Int to string map
	intStringMap := NewGenericMap[int, string]()
	intStringMap.Set(1, "one")
	intStringMap.Set(2, "two")
	intStringMap.Set(3, "three")
	gm.Maps["int_string"] = intStringMap
	
	// String to Person map
	stringPersonMap := NewGenericMap[string, Person]()
	stringPersonMap.Set("alice", Person{ID: 1, Name: "Alice", Age: 30})
	stringPersonMap.Set("bob", Person{ID: 2, Name: "Bob", Age: 25})
	gm.Maps["string_person"] = stringPersonMap
	
	fmt.Println("Created generic maps")
}

// createTrees creates various generic trees
func (gm *GenericManager) createTrees() {
	// Int tree
	intRoot := NewGenericNode(1)
	intChild1 := NewGenericNode(2)
	intChild2 := NewGenericNode(3)
	intRoot.AddChild(intChild1)
	intRoot.AddChild(intChild2)
	gm.Trees["int"] = intRoot
	
	// String tree
	stringRoot := NewGenericNode("root")
	stringChild1 := NewGenericNode("child1")
	stringChild2 := NewGenericNode("child2")
	stringRoot.AddChild(stringChild1)
	stringRoot.AddChild(stringChild2)
	gm.Trees["string"] = stringRoot
	
	fmt.Println("Created generic trees")
}

// Read - Display generic information
func (gm *GenericManager) Read() {
	fmt.Println("\n📖 READING GENERIC INSTANCES:")
	fmt.Println("=============================")
	
	// Read containers
	gm.readContainers()
	
	// Read stacks
	gm.readStacks()
	
	// Read queues
	gm.readQueues()
	
	// Read sets
	gm.readSets()
	
	// Read maps
	gm.readMaps()
	
	// Read trees
	gm.readTrees()
}

// readContainers displays container information
func (gm *GenericManager) readContainers() {
	fmt.Println("Containers:")
	
	if intContainer, ok := gm.Containers["int"].(*GenericContainer[int]); ok {
		fmt.Printf("  Int Container: %v (size: %d)\n", intContainer.Items, intContainer.Size())
	}
	
	if stringContainer, ok := gm.Containers["string"].(*GenericContainer[string]); ok {
		fmt.Printf("  String Container: %v (size: %d)\n", stringContainer.Items, stringContainer.Size())
	}
	
	if personContainer, ok := gm.Containers["person"].(*GenericContainer[Person]); ok {
		fmt.Printf("  Person Container: %v (size: %d)\n", personContainer.Items, personContainer.Size())
	}
}

// readStacks displays stack information
func (gm *GenericManager) readStacks() {
	fmt.Println("\nStacks:")
	
	if intStack, ok := gm.Stacks["int"].(*GenericStack[int]); ok {
		fmt.Printf("  Int Stack: size %d\n", intStack.Size())
		if top, ok := intStack.Peek(); ok {
			fmt.Printf("    Top: %d\n", top)
		}
	}
	
	if stringStack, ok := gm.Stacks["string"].(*GenericStack[string]); ok {
		fmt.Printf("  String Stack: size %d\n", stringStack.Size())
		if top, ok := stringStack.Peek(); ok {
			fmt.Printf("    Top: %s\n", top)
		}
	}
}

// readQueues displays queue information
func (gm *GenericManager) readQueues() {
	fmt.Println("\nQueues:")
	
	if intQueue, ok := gm.Queues["int"].(*GenericQueue[int]); ok {
		fmt.Printf("  Int Queue: size %d\n", intQueue.Size())
		if front, ok := intQueue.Front(); ok {
			fmt.Printf("    Front: %d\n", front)
		}
	}
	
	if stringQueue, ok := gm.Queues["string"].(*GenericQueue[string]); ok {
		fmt.Printf("  String Queue: size %d\n", stringQueue.Size())
		if front, ok := stringQueue.Front(); ok {
			fmt.Printf("    Front: %s\n", front)
		}
	}
}

// readSets displays set information
func (gm *GenericManager) readSets() {
	fmt.Println("\nSets:")
	
	if intSet, ok := gm.Sets["int"].(*GenericSet[int]); ok {
		fmt.Printf("  Int Set: %v (size: %d)\n", intSet.ToSlice(), intSet.Size())
	}
	
	if stringSet, ok := gm.Sets["string"].(*GenericSet[string]); ok {
		fmt.Printf("  String Set: %v (size: %d)\n", stringSet.ToSlice(), stringSet.Size())
	}
}

// readMaps displays map information
func (gm *GenericManager) readMaps() {
	fmt.Println("\nMaps:")
	
	if stringIntMap, ok := gm.Maps["string_int"].(*GenericMapType[string, int]); ok {
		fmt.Printf("  String-Int Map: %v (size: %d)\n", stringIntMap.items, stringIntMap.Size())
	}
	
	if intStringMap, ok := gm.Maps["int_string"].(*GenericMapType[int, string]); ok {
		fmt.Printf("  Int-String Map: %v (size: %d)\n", intStringMap.items, intStringMap.Size())
	}
	
	if stringPersonMap, ok := gm.Maps["string_person"].(*GenericMapType[string, Person]); ok {
		fmt.Printf("  String-Person Map: %v (size: %d)\n", stringPersonMap.items, stringPersonMap.Size())
	}
}

// readTrees displays tree information
func (gm *GenericManager) readTrees() {
	fmt.Println("\nTrees:")
	
	if intRoot, ok := gm.Trees["int"].(*GenericNode[int]); ok {
		fmt.Printf("  Int Tree: root=%d, height=%d, depth=%d\n", 
			intRoot.Value, intRoot.Height(), intRoot.Depth())
		fmt.Printf("    Children: %d\n", len(intRoot.Children))
	}
	
	if stringRoot, ok := gm.Trees["string"].(*GenericNode[string]); ok {
		fmt.Printf("  String Tree: root=%s, height=%d, depth=%d\n", 
			stringRoot.Value, stringRoot.Height(), stringRoot.Depth())
		fmt.Printf("    Children: %d\n", len(stringRoot.Children))
	}
}

// Update - Modify generic instances
func (gm *GenericManager) Update() {
	fmt.Println("\n🔄 UPDATING GENERIC INSTANCES:")
	fmt.Println("==============================")
	
	// Update containers
	gm.updateContainers()
	
	// Update stacks
	gm.updateStacks()
	
	// Update queues
	gm.updateQueues()
	
	// Update sets
	gm.updateSets()
	
	// Update maps
	gm.updateMaps()
	
	// Update trees
	gm.updateTrees()
	
	fmt.Println("✅ Generic instances updated successfully")
}

// updateContainers updates container instances
func (gm *GenericManager) updateContainers() {
	if intContainer, ok := gm.Containers["int"].(*GenericContainer[int]); ok {
		intContainer.Add(4)
		intContainer.Add(5)
		fmt.Println("Updated int container")
	}
	
	if stringContainer, ok := gm.Containers["string"].(*GenericContainer[string]); ok {
		stringContainer.Add("date")
		stringContainer.Add("elderberry")
		fmt.Println("Updated string container")
	}
}

// updateStacks updates stack instances
func (gm *GenericManager) updateStacks() {
	if intStack, ok := gm.Stacks["int"].(*GenericStack[int]); ok {
		intStack.Push(4)
		intStack.Push(5)
		fmt.Println("Updated int stack")
	}
	
	if stringStack, ok := gm.Stacks["string"].(*GenericStack[string]); ok {
		stringStack.Push("fourth")
		stringStack.Push("fifth")
		fmt.Println("Updated string stack")
	}
}

// updateQueues updates queue instances
func (gm *GenericManager) updateQueues() {
	if intQueue, ok := gm.Queues["int"].(*GenericQueue[int]); ok {
		intQueue.Enqueue(4)
		intQueue.Enqueue(5)
		fmt.Println("Updated int queue")
	}
	
	if stringQueue, ok := gm.Queues["string"].(*GenericQueue[string]); ok {
		stringQueue.Enqueue("fourth")
		stringQueue.Enqueue("fifth")
		fmt.Println("Updated string queue")
	}
}

// updateSets updates set instances
func (gm *GenericManager) updateSets() {
	if intSet, ok := gm.Sets["int"].(*GenericSet[int]); ok {
		intSet.Add(4)
		intSet.Add(5)
		fmt.Println("Updated int set")
	}
	
	if stringSet, ok := gm.Sets["string"].(*GenericSet[string]); ok {
		stringSet.Add("date")
		stringSet.Add("elderberry")
		fmt.Println("Updated string set")
	}
}

// updateMaps updates map instances
func (gm *GenericManager) updateMaps() {
	if stringIntMap, ok := gm.Maps["string_int"].(*GenericMapType[string, int]); ok {
		stringIntMap.Set("grape", 4)
		stringIntMap.Set("kiwi", 2)
		fmt.Println("Updated string-int map")
	}
	
	if intStringMap, ok := gm.Maps["int_string"].(*GenericMapType[int, string]); ok {
		intStringMap.Set(4, "four")
		intStringMap.Set(5, "five")
		fmt.Println("Updated int-string map")
	}
}

// updateTrees updates tree instances
func (gm *GenericManager) updateTrees() {
	if intRoot, ok := gm.Trees["int"].(*GenericNode[int]); ok {
		newChild := NewGenericNode(4)
		intRoot.AddChild(newChild)
		fmt.Println("Updated int tree")
	}
	
	if stringRoot, ok := gm.Trees["string"].(*GenericNode[string]); ok {
		newChild := NewGenericNode("child3")
		stringRoot.AddChild(newChild)
		fmt.Println("Updated string tree")
	}
}

// Delete - Remove generic instances
func (gm *GenericManager) Delete() {
	fmt.Println("\n🗑️  DELETING GENERIC INSTANCES:")
	fmt.Println("===============================")
	
	// Delete containers
	gm.deleteContainers()
	
	// Delete stacks
	gm.deleteStacks()
	
	// Delete queues
	gm.deleteQueues()
	
	// Delete sets
	gm.deleteSets()
	
	// Delete maps
	gm.deleteMaps()
	
	// Delete trees
	gm.deleteTrees()
	
	fmt.Println("✅ Generic instances deleted successfully")
}

// deleteContainers deletes container instances
func (gm *GenericManager) deleteContainers() {
	if intContainer, ok := gm.Containers["int"].(*GenericContainer[int]); ok {
		intContainer.Remove(0) // Remove first item
		fmt.Println("Deleted from int container")
	}
	
	if stringContainer, ok := gm.Containers["string"].(*GenericContainer[string]); ok {
		stringContainer.Clear()
		fmt.Println("Cleared string container")
	}
}

// deleteStacks deletes stack instances
func (gm *GenericManager) deleteStacks() {
	if intStack, ok := gm.Stacks["int"].(*GenericStack[int]); ok {
		intStack.Pop()
		fmt.Println("Popped from int stack")
	}
	
	if stringStack, ok := gm.Stacks["string"].(*GenericStack[string]); ok {
		stringStack.Pop()
		fmt.Println("Popped from string stack")
	}
}

// deleteQueues deletes queue instances
func (gm *GenericManager) deleteQueues() {
	if intQueue, ok := gm.Queues["int"].(*GenericQueue[int]); ok {
		intQueue.Dequeue()
		fmt.Println("Dequeued from int queue")
	}
	
	if stringQueue, ok := gm.Queues["string"].(*GenericQueue[string]); ok {
		stringQueue.Dequeue()
		fmt.Println("Dequeued from string queue")
	}
}

// deleteSets deletes set instances
func (gm *GenericManager) deleteSets() {
	if intSet, ok := gm.Sets["int"].(*GenericSet[int]); ok {
		intSet.Remove(1)
		fmt.Println("Removed from int set")
	}
	
	if stringSet, ok := gm.Sets["string"].(*GenericSet[string]); ok {
		stringSet.Remove("apple")
		fmt.Println("Removed from string set")
	}
}

// deleteMaps deletes map instances
func (gm *GenericManager) deleteMaps() {
	if stringIntMap, ok := gm.Maps["string_int"].(*GenericMapType[string, int]); ok {
		stringIntMap.Delete("apple")
		fmt.Println("Deleted from string-int map")
	}
	
	if intStringMap, ok := gm.Maps["int_string"].(*GenericMapType[int, string]); ok {
		intStringMap.Delete(1)
		fmt.Println("Deleted from int-string map")
	}
}

// deleteTrees deletes tree instances
func (gm *GenericManager) deleteTrees() {
	if intRoot, ok := gm.Trees["int"].(*GenericNode[int]); ok {
		if len(intRoot.Children) > 0 {
			intRoot.RemoveChild(intRoot.Children[0])
			fmt.Println("Removed child from int tree")
		}
	}
	
	if stringRoot, ok := gm.Trees["string"].(*GenericNode[string]); ok {
		if len(stringRoot.Children) > 0 {
			stringRoot.RemoveChild(stringRoot.Children[0])
			fmt.Println("Removed child from string tree")
		}
	}
}

// Advanced Generic Operations

// DemonstrateGenericFunctions shows generic function usage
func (gm *GenericManager) DemonstrateGenericFunctions() {
	fmt.Println("\n🔧 GENERIC FUNCTIONS DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Demonstrate basic generic functions
	numbers := []int{1, 2, 3, 4, 5}
	strings := []string{"apple", "banana", "cherry"}
	
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Max: %d\n", GenericMax(numbers[0], numbers[1]))
	fmt.Printf("Min: %d\n", GenericMin(numbers[0], numbers[1]))
	fmt.Printf("Sum: %d\n", GenericSum(numbers))
	fmt.Printf("Average: %.2f\n", GenericAverage(numbers))
	
	fmt.Printf("\nStrings: %v\n", strings)
	fmt.Printf("Max: %s\n", GenericMax(strings[0], strings[1]))
	fmt.Printf("Min: %s\n", GenericMin(strings[0], strings[1]))
	
	// Demonstrate filtering
	evenNumbers := GenericFilter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evenNumbers)
	
	// Demonstrate mapping
	doubledNumbers := GenericMap(numbers, func(x int) int {
		return x * 2
	})
	fmt.Printf("Doubled numbers: %v\n", doubledNumbers)
	
	// Demonstrate reducing
	sum := GenericReduce(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Printf("Sum via reduce: %d\n", sum)
	
	// Demonstrate sorting
	sortedNumbers := GenericSort(numbers)
	fmt.Printf("Sorted numbers: %v\n", sortedNumbers)
	
	// Demonstrate reversing
	reversedNumbers := GenericReverse(numbers)
	fmt.Printf("Reversed numbers: %v\n", reversedNumbers)
	
	// Demonstrate unique
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5}
	uniqueNumbers := GenericUnique(duplicates)
	fmt.Printf("Unique numbers: %v\n", uniqueNumbers)
	
	// Demonstrate chunking
	chunks := GenericChunk(numbers, 2)
	fmt.Printf("Chunks: %v\n", chunks)
	
	// Demonstrate flattening
	flattened := GenericFlatten(chunks)
	fmt.Printf("Flattened: %v\n", flattened)
}

// DemonstrateGenericConstraints shows generic constraint usage
func (gm *GenericManager) DemonstrateGenericConstraints() {
	fmt.Println("\n🔒 GENERIC CONSTRAINTS DEMONSTRATION:")
	fmt.Println("====================================")
	
	// Demonstrate numeric constraints
	intNumbers := []int{1, 2, 3, 4, 5}
	floatNumbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	
	fmt.Printf("Int numbers: %v\n", intNumbers)
	fmt.Printf("Int sum: %d\n", GenericSum(intNumbers))
	fmt.Printf("Int average: %.2f\n", GenericAverage(intNumbers))
	
	fmt.Printf("\nFloat numbers: %v\n", floatNumbers)
	fmt.Printf("Float sum: %.2f\n", GenericSum(floatNumbers))
	fmt.Printf("Float average: %.2f\n", GenericAverage(floatNumbers))
	
	// Demonstrate comparable constraints
	comparableItems := []string{"apple", "banana", "cherry"}
	fmt.Printf("\nComparable items: %v\n", comparableItems)
	fmt.Printf("Contains 'banana': %t\n", GenericContains(comparableItems, "banana"))
	fmt.Printf("Index of 'cherry': %d\n", GenericIndexOf(comparableItems, "cherry"))
	fmt.Printf("Max: %s\n", GenericMax(comparableItems[0], comparableItems[1]))
	fmt.Printf("Min: %s\n", GenericMin(comparableItems[0], comparableItems[1]))
}

// ComparableInt is a custom type that implements generic interfaces
type ComparableInt int

func (ci ComparableInt) Compare(other ComparableInt) int {
	if ci < other {
		return -1
	} else if ci > other {
		return 1
	}
	return 0
}

func (ci ComparableInt) String() string {
	return fmt.Sprintf("ComparableInt(%d)", int(ci))
}

func (ci ComparableInt) Clone() ComparableInt {
	return ci
}

// DemonstrateGenericInterfaces shows generic interface usage
func (gm *GenericManager) DemonstrateGenericInterfaces() {
	fmt.Println("\n🔌 GENERIC INTERFACES DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Demonstrate generic interface usage
	items := []ComparableInt{3, 1, 4, 1, 5}
	fmt.Printf("Items: %v\n", items)
	
	// Sort using generic interface
	sortedItems := GenericSort(items)
	fmt.Printf("Sorted: %v\n", sortedItems)
	
	// Demonstrate interface methods
	item1 := ComparableInt(5)
	item2 := ComparableInt(3)
	fmt.Printf("Compare %s and %s: %d\n", item1.String(), item2.String(), item1.Compare(item2))
	
	cloned := item1.Clone()
	fmt.Printf("Cloned: %s\n", cloned.String())
}

// DemonstrateGenericTypeInference shows type inference
func (gm *GenericManager) DemonstrateGenericTypeInference() {
	fmt.Println("\n🔍 GENERIC TYPE INFERENCE DEMONSTRATION:")
	fmt.Println("=======================================")
	
	// Type inference with function calls
	numbers := []int{1, 2, 3, 4, 5}
	
	// Type inference works automatically
	sum := GenericSum(numbers)
	fmt.Printf("Sum of %v: %d\n", numbers, sum)
	
	// Type inference with different types
	floats := []float64{1.1, 2.2, 3.3}
	floatSum := GenericSum(floats)
	fmt.Printf("Sum of %v: %.2f\n", floats, floatSum)
	
	// Type inference with strings
	words := []string{"hello", "world", "golang"}
	// Note: strings don't implement Numeric, so this won't work
	// stringSum := GenericSum(words) // This would cause a compile error
	
	// Use words to avoid unused variable warning
	_ = words
	
	// Type inference with maps
	stringIntMap := NewGenericMap[string, int]()
	stringIntMap.Set("apple", 5)
	stringIntMap.Set("banana", 3)
	
	// Type inference with sets
	intSet := NewGenericSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	
	fmt.Printf("Map: %v\n", stringIntMap.items)
	fmt.Printf("Set: %v\n", intSet.ToSlice())
}

// DemonstrateGenericPerformance shows generic performance considerations
func (gm *GenericManager) DemonstrateGenericPerformance() {
	fmt.Println("\n⚡ GENERIC PERFORMANCE DEMONSTRATION:")
	fmt.Println("====================================")
	
	// Create large slices for performance testing
	largeIntSlice := make([]int, 1000)
	for i := range largeIntSlice {
		largeIntSlice[i] = i
	}
	
	largeStringSlice := make([]string, 1000)
	for i := range largeStringSlice {
		largeStringSlice[i] = fmt.Sprintf("item_%d", i)
	}
	
	// Test generic operations
	fmt.Printf("Large int slice size: %d\n", len(largeIntSlice))
	
	// Test sum operation
	sum := GenericSum(largeIntSlice)
	fmt.Printf("Sum of large slice: %d\n", sum)
	
	// Test filtering
	evenNumbers := GenericFilter(largeIntSlice, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("Even numbers count: %d\n", len(evenNumbers))
	
	// Test mapping
	doubledNumbers := GenericMap(largeIntSlice, func(x int) int {
		return x * 2
	})
	fmt.Printf("Doubled numbers count: %d\n", len(doubledNumbers))
	
	// Test sorting
	sortedNumbers := GenericSort(largeIntSlice)
	fmt.Printf("First 5 sorted numbers: %v\n", sortedNumbers[:5])
	
	// Test with strings
	fmt.Printf("\nLarge string slice size: %d\n", len(largeStringSlice))
	
	// Test string operations
	filteredStrings := GenericFilter(largeStringSlice, func(s string) bool {
		return strings.Contains(s, "5")
	})
	fmt.Printf("Strings containing '5': %d\n", len(filteredStrings))
	
	// Test string mapping
	upperStrings := GenericMap(largeStringSlice, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("First 3 upper strings: %v\n", upperStrings[:3])
	
	// Use the variables to avoid unused variable warnings
	_ = filteredStrings
	_ = upperStrings
}
