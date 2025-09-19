package models

import (
	"sync"
)

// Cache Node for doubly linked list
type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
	mu    sync.RWMutex
}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (n *Node) GetKey() string {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.key
}

func (n *Node) GetValue() interface{} {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.value
}

func (n *Node) SetValue(value interface{}) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.value = value
}

func (n *Node) GetPrev() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.prev
}

func (n *Node) SetPrev(prev *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.prev = prev
}

func (n *Node) GetNext() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.next = next
}

func (n *Node) Reset() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.key = ""
	n.value = nil
	n.prev = nil
	n.next = nil
}

func (n *Node) GetDetails() map[string]interface{} {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	return map[string]interface{}{
		"key":   n.key,
		"value": n.value,
		"has_prev": n.prev != nil,
		"has_next": n.next != nil,
	}
}
