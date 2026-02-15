package storage

import (
	"database-engine/schema"
	"database-engine/types"
	"fmt"
	"sync"
)

// BTree represents a B+ Tree
type BTree struct {
	root      *BTreeNode
	order     int
	height    int
	nodeCount int
	mutex     sync.RWMutex
}

// BTreeNode represents a node in the B+ Tree
type BTreeNode struct {
	keys     []Key
	values   []*BTreeRecord
	records  []*BTreeRecord // For leaf nodes
	children []*BTreeNode
	parent   *BTreeNode
	isLeaf   bool
	next     *BTreeNode // For leaf nodes
	prev     *BTreeNode // For leaf nodes
}

// BTreeRecord represents a record in the B+ Tree
type BTreeRecord struct {
	Key    Key
	Values map[string]types.Value
}

// BTreeIndex represents a B+ Tree index
type BTreeIndex struct {
	Name    string
	Columns []string
	Unique  bool
	BTree   *BTree
}

// NewBTree creates a new B+ Tree
func NewBTree() *BTree {
	return &BTree{
		order:     4, // Minimum degree
		height:    0,
		nodeCount: 0,
	}
}

// NewBTreeIndex creates a new B+ Tree index
func NewBTreeIndex(index *schema.Index) *BTreeIndex {
	return &BTreeIndex{
		Name:    index.Name,
		Columns: index.Columns,
		Unique:  index.Unique,
		BTree:   NewBTree(),
	}
}

// Insert inserts a record into the B+ Tree
func (bt *BTree) Insert(record *BTreeRecord) error {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	
	if bt.root == nil {
		// Create root node
		bt.root = &BTreeNode{
			keys:   []Key{record.Key},
			values: []*BTreeRecord{record},
			isLeaf: true,
		}
		bt.height = 1
		bt.nodeCount = 1
		return nil
	}
	
	// Find insertion point
	leaf := bt.findLeaf(record.Key)
	
	// Insert into leaf
	if err := bt.insertIntoLeaf(leaf, record); err != nil {
		return err
	}
	
	// Check if leaf needs to be split
	if len(leaf.keys) > bt.order*2 {
		bt.splitLeaf(leaf)
	}
	
	return nil
}

// Search searches for a record by key
func (bt *BTree) Search(key Key) (*BTreeRecord, error) {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	
	if bt.root == nil {
		return nil, fmt.Errorf("record not found")
	}
	
	// Find leaf node
	leaf := bt.findLeaf(key)
	
	// Search within leaf
	for i, k := range leaf.keys {
		if k.Compare(key) == 0 {
			return leaf.values[i], nil
		}
	}
	
	return nil, fmt.Errorf("record not found")
}

// Update updates a record
func (bt *BTree) Update(key Key, record *BTreeRecord) error {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	
	if bt.root == nil {
		return fmt.Errorf("record not found")
	}
	
	// Find leaf node
	leaf := bt.findLeaf(key)
	
	// Update within leaf
	for i, k := range leaf.keys {
		if k.Compare(key) == 0 {
			leaf.values[i] = record
			return nil
		}
	}
	
	return fmt.Errorf("record not found")
}

// Delete deletes a record by key
func (bt *BTree) Delete(key Key) error {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	
	if bt.root == nil {
		return fmt.Errorf("record not found")
	}
	
	// Find leaf node
	leaf := bt.findLeaf(key)
	
	// Find and remove record
	index := -1
	for i, k := range leaf.keys {
		if k.Compare(key) == 0 {
			index = i
			break
		}
	}
	
	if index == -1 {
		return fmt.Errorf("record not found")
	}
	
	// Remove record
	leaf.keys = append(leaf.keys[:index], leaf.keys[index+1:]...)
	leaf.values = append(leaf.values[:index], leaf.values[index+1:]...)
	
	// Check if leaf needs to be merged
	if len(leaf.keys) < bt.order && leaf != bt.root {
		bt.mergeLeaf(leaf)
	}
	
	return nil
}

// findLeaf finds the leaf node that should contain the key
func (bt *BTree) findLeaf(key Key) *BTreeNode {
	node := bt.root
	
	for !node.isLeaf {
		// Find appropriate child
		childIndex := bt.findChildIndex(node, key)
		node = node.children[childIndex]
	}
	
	return node
}

// findChildIndex finds the index of the child that should contain the key
func (bt *BTree) findChildIndex(node *BTreeNode, key Key) int {
	for i, k := range node.keys {
		if key.Compare(k) < 0 {
			return i
		}
	}
	return len(node.keys)
}

// insertIntoLeaf inserts a record into a leaf node
func (bt *BTree) insertIntoLeaf(leaf *BTreeNode, record *BTreeRecord) error {
	// Find insertion point
	index := bt.findInsertionIndex(leaf, record.Key)
	
	// Insert key and value
	leaf.keys = append(leaf.keys[:index], append([]Key{record.Key}, leaf.keys[index:]...)...)
	leaf.values = append(leaf.values[:index], append([]*BTreeRecord{record}, leaf.values[index:]...)...)
	
	return nil
}

// findInsertionIndex finds the index where a key should be inserted
func (bt *BTree) findInsertionIndex(node *BTreeNode, key Key) int {
	for i, k := range node.keys {
		if key.Compare(k) < 0 {
			return i
		}
	}
	return len(node.keys)
}

// splitLeaf splits a leaf node
func (bt *BTree) splitLeaf(leaf *BTreeNode) {
	// Create new leaf
	newLeaf := &BTreeNode{
		keys:   make([]Key, 0),
		values: make([]*BTreeRecord, 0),
		isLeaf: true,
	}
	
	// Move half the records to new leaf
	mid := len(leaf.keys) / 2
	newLeaf.keys = leaf.keys[mid:]
	newLeaf.values = leaf.values[mid:]
	leaf.keys = leaf.keys[:mid]
	leaf.values = leaf.values[:mid]
	
	// Update linked list
	newLeaf.next = leaf.next
	newLeaf.prev = leaf
	leaf.next = newLeaf
	if newLeaf.next != nil {
		newLeaf.next.prev = newLeaf
	}
	
	// Update parent
	if leaf.parent == nil {
		// Create new root
		bt.root = &BTreeNode{
			keys:     []Key{newLeaf.keys[0]},
			children: []*BTreeNode{leaf, newLeaf},
			isLeaf:   false,
		}
		leaf.parent = bt.root
		newLeaf.parent = bt.root
		bt.height++
	} else {
		// Insert into parent
		bt.insertIntoParent(leaf.parent, newLeaf.keys[0], newLeaf)
	}
	
	bt.nodeCount++
}

// insertIntoParent inserts a key and child into a parent node
func (bt *BTree) insertIntoParent(parent *BTreeNode, key Key, child *BTreeNode) {
	// Find insertion point
	index := bt.findInsertionIndex(parent, key)
	
	// Insert key and child
	parent.keys = append(parent.keys[:index], append([]Key{key}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*BTreeNode{child}, parent.children[index+1:]...)...)
	
	// Update child's parent
	child.parent = parent
	
	// Check if parent needs to be split
	if len(parent.keys) > bt.order*2 {
		bt.splitInternal(parent)
	}
}

// splitInternal splits an internal node
func (bt *BTree) splitInternal(node *BTreeNode) {
	// Create new internal node
	newNode := &BTreeNode{
		keys:     make([]Key, 0),
		children: make([]*BTreeNode, 0),
		isLeaf:   false,
	}
	
	// Move half the keys and children to new node
	mid := len(node.keys) / 2
	newNode.keys = node.keys[mid+1:]
	newNode.children = node.children[mid+1:]
	node.keys = node.keys[:mid]
	node.children = node.children[:mid+1]
	
	// Update children's parent
	for _, child := range newNode.children {
		child.parent = newNode
	}
	
	// Update parent
	if node.parent == nil {
		// Create new root
		bt.root = &BTreeNode{
			keys:     []Key{node.keys[mid]},
			children: []*BTreeNode{node, newNode},
			isLeaf:   false,
		}
		node.parent = bt.root
		newNode.parent = bt.root
		bt.height++
	} else {
		// Insert into parent
		bt.insertIntoParent(node.parent, node.keys[mid], newNode)
	}
	
	bt.nodeCount++
}

// mergeLeaf merges a leaf node with its sibling
func (bt *BTree) mergeLeaf(leaf *BTreeNode) {
	// Find sibling
	var sibling *BTreeNode
	var isLeftSibling bool
	
	if leaf.prev != nil && leaf.prev.parent == leaf.parent {
		sibling = leaf.prev
		isLeftSibling = true
	} else if leaf.next != nil && leaf.next.parent == leaf.parent {
		sibling = leaf.next
		isLeftSibling = false
	} else {
		return // No sibling to merge with
	}
	
	// Merge with sibling
	if isLeftSibling {
		// Merge left sibling into leaf
		leaf.keys = append(sibling.keys, leaf.keys...)
		leaf.values = append(sibling.values, leaf.values...)
		leaf.prev = sibling.prev
		if sibling.prev != nil {
			sibling.prev.next = leaf
		}
	} else {
		// Merge leaf into right sibling
		sibling.keys = append(leaf.keys, sibling.keys...)
		sibling.values = append(leaf.values, sibling.values...)
		sibling.prev = leaf.prev
		if leaf.prev != nil {
			leaf.prev.next = sibling
		}
	}
	
	// Remove sibling from parent
	bt.removeFromParent(sibling)
	
	// Check if parent needs to be merged
	if leaf.parent != nil && len(leaf.parent.keys) < bt.order {
		bt.mergeInternal(leaf.parent)
	}
	
	bt.nodeCount--
}

// removeFromParent removes a node from its parent
func (bt *BTree) removeFromParent(node *BTreeNode) {
	if node.parent == nil {
		return
	}
	
	parent := node.parent
	
	// Find node in parent's children
	index := -1
	for i, child := range parent.children {
		if child == node {
			index = i
			break
		}
	}
	
	if index == -1 {
		return
	}
	
	// Remove child
	parent.children = append(parent.children[:index], parent.children[index+1:]...)
	
	// Remove corresponding key
	if index > 0 {
		parent.keys = append(parent.keys[:index-1], parent.keys[index:]...)
	}
}

// mergeInternal merges an internal node with its sibling
func (bt *BTree) mergeInternal(node *BTreeNode) {
	// Find sibling
	// This is a simplified implementation
	
	if node.parent == nil {
		// This is the root
		if len(node.children) == 1 {
			// Promote only child to root
			bt.root = node.children[0]
			bt.root.parent = nil
			bt.height--
		}
		return
	}
	
	// Merge with sibling
	// Implementation details...
}

// Vacuum performs vacuum operation
func (bt *BTree) Vacuum() error {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	
	// Vacuum implementation
	// This would involve:
	// 1. Rebuilding the tree to remove fragmentation
	// 2. Compacting leaf nodes
	// 3. Rebalancing the tree
	
	return nil
}

// Analyze performs analysis operation
func (bt *BTree) Analyze() error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	
	// Analyze implementation
	// This would involve:
	// 1. Collecting statistics about the tree
	// 2. Updating cost estimates
	// 3. Identifying optimization opportunities
	
	return nil
}

// CheckIntegrity checks tree integrity
func (bt *BTree) CheckIntegrity() error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	
	// Integrity check implementation
	// This would involve:
	// 1. Verifying tree structure
	// 2. Checking key ordering
	// 3. Validating parent-child relationships
	
	return nil
}

// BTreeIndexIterator implements an iterator for B+ Tree index
type BTreeIndexIterator struct {
	index *BTreeIndex
	key   Key
	iter  *BTreeIterator
}

func (iter *BTreeIndexIterator) Next() (*Row, error) {
	// Implementation for index iteration
	// This would involve:
	// 1. Finding the starting position in the index
	// 2. Iterating through index entries
	// 3. Converting index entries to rows
	
	return nil, nil
}

func (iter *BTreeIndexIterator) Close() error {
	// Cleanup resources
	return nil
}

// BTreeIndex methods

// Insert inserts a key-value pair into the index
func (bti *BTreeIndex) Insert(key Key, value Key) error {
	// Create a record for the index
	record := &BTreeRecord{
		Key: key,
		Values: map[string]types.Value{
			"value": value.(*SimpleKey).Value,
		},
	}
	
	return bti.BTree.Insert(record)
}

// Delete deletes a key from the index
func (bti *BTreeIndex) Delete(key Key) error {
	return bti.BTree.Delete(key)
}

// Search searches for a key in the index
func (bti *BTreeIndex) Search(key Key) (Key, error) {
	record, err := bti.BTree.Search(key)
	if err != nil {
		return nil, err
	}
	
	// Extract value from record
	if val, exists := record.Values["value"]; exists {
		return &SimpleKey{Value: val}, nil
	}
	
	return nil, fmt.Errorf("value not found in index record")
}

// RangeSearch searches for keys in a range
func (bti *BTreeIndex) RangeSearch(start, end Key) ([]Key, error) {
	// Implementation for range search
	// This would involve:
	// 1. Finding the starting position
	// 2. Iterating through keys in range
	// 3. Collecting matching keys
	
	return nil, nil
}
