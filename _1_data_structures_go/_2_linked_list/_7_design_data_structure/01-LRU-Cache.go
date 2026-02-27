//  https://leetcode.com/problems/lru-cache/


package main

import "fmt"

/************ NODE ************/

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

/************ LRU CACHE ************/

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

/************ CONSTRUCTOR ************/

func Constructor(capacity int) LRUCache {

	head := &Node{-1, -1, nil, nil}
	tail := &Node{-1, -1, nil, nil}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

/************ ADD NODE (FRONT) ************/

func (l *LRUCache) addNode(newNode *Node) {

	temp := l.head.next

	newNode.next = temp
	newNode.prev = l.head

	l.head.next = newNode
	temp.prev = newNode
}

/************ DELETE NODE ************/

func (l *LRUCache) deleteNode(delNode *Node) {

	prev := delNode.prev
	next := delNode.next

	prev.next = next
	next.prev = prev
}

/************ GET ************/

func (l *LRUCache) Get(key int) int {

	if node, ok := l.cache[key]; ok {

		l.deleteNode(node)
		l.addNode(node)

		return node.val
	}

	return -1
}

/************ PUT ************/

func (l *LRUCache) Put(key int, value int) {

	if node, ok := l.cache[key]; ok {
		l.deleteNode(node)
		delete(l.cache, key)
	}

	if len(l.cache) == l.capacity {
		lru := l.tail.prev
		l.deleteNode(lru)
		delete(l.cache, lru.key)
	}

	newNode := &Node{key, value, nil, nil}
	l.addNode(newNode)
	l.cache[key] = newNode
}
// Time Complexity:O(1)
// Space Complexity:O(1)

/************ TEST ************/

func main() {

	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	fmt.Println(lru.Get(1)) // 1

	lru.Put(3, 3) // evicts key 2

	fmt.Println(lru.Get(2)) // -1

	lru.Put(4, 4) // evicts key 1

	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}