package main

import "fmt"

type Node struct {
	data int
	next *Node
	back *Node
}

// Convert array to DLL
func convertArr2DLL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{data: arr[0]}
	prev := head

	for i := 1; i < len(arr); i++ {
		temp := &Node{data: arr[i]}
		prev.next = temp
		temp.back = prev
		prev = temp
	}

	return head
}

// Traverse forward
func traverse(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

//
// DELETE OPERATIONS
//

func deleteHead(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	newHead := head.next
	newHead.back = nil
	head.next = nil

	return newHead
}

func deleteTail(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	tail := head
	for tail.next != nil {
		tail = tail.next
	}

	prev := tail.back
	prev.next = nil
	tail.back = nil

	return head
}

func deleteKthElement(head *Node, k int) *Node {
	if head == nil {
		return nil
	}

	count := 1
	curr := head

	for curr != nil && count < k {
		curr = curr.next
		count++
	}

	if curr == nil {
		return head
	}

	prev := curr.back
	next := curr.next

	// only node
	if prev == nil && next == nil {
		return nil
	}

	// head
	if prev == nil {
		return deleteHead(head)
	}

	// tail
	if next == nil {
		return deleteTail(head)
	}

	prev.next = next
	next.back = prev

	curr.next = nil
	curr.back = nil

	return head
}

func deleteGivenNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	prev := node.back
	next := node.next

	// Only node in list
	if prev == nil && next == nil {
		return nil
	}

	// Node is head
	if prev == nil {
		next.back = nil
		node.next = nil
		return next
	}

	// Node is tail
	if next == nil {
		prev.next = nil
		node.back = nil
		return nil
	}

	// Node is in middle
	prev.next = next
	next.back = prev

	node.next = nil
	node.back = nil

	return nil
}

//
// INSERT OPERATIONS
//

func insertBeforeHead(head *Node, val int) *Node {
	newHead := &Node{data: val, next: head}
	if head != nil {
		head.back = newHead
	}
	return newHead
}

func insertBeforeTail(head *Node, val int) *Node {
	if head == nil || head.next == nil {
		return insertBeforeHead(head, val)
	}

	tail := head
	for tail.next != nil {
		tail = tail.next
	}

	prev := tail.back

	newNode := &Node{data: val, next: tail, back: prev}

	prev.next = newNode
	tail.back = newNode

	return head
}

func insertBeforeKthElement(head *Node, k, val int) *Node {
	if k == 1 {
		return insertBeforeHead(head, val)
	}

	count := 1
	curr := head

	for curr != nil && count < k {
		curr = curr.next
		count++
	}

	if curr == nil {
		return head
	}

	prev := curr.back
	newNode := &Node{data: val, next: curr, back: prev}

	prev.next = newNode
	curr.back = newNode

	return head
}

func insertBeforeNode(node *Node, val int) {
	if node == nil || node.back == nil {
		return
	}

	prev := node.back
	newNode := &Node{data: val, next: node, back: prev}

	prev.next = newNode
	node.back = newNode
}

func insertAtTail(head *Node, val int) *Node {
	newNode := &Node{data: val}

	if head == nil {
		return newNode
	}

	tail := head
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = newNode
	newNode.back = tail

	return head
}

func main() {
	arr := []int{3, 5, 8, 7, 6}
	head := convertArr2DLL(arr)

	fmt.Print("Initial DLL: ")
	traverse(head)

	head = insertBeforeHead(head, 1)
	fmt.Print("After insert before head: ")
	traverse(head)

	head = deleteKthElement(head, 3)
	fmt.Print("After deleting 3rd element: ")
	traverse(head)
}