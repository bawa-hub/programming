package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// Convert array to linked list
func convertArr2LL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{data: arr[0]}
	curr := head

	for i := 1; i < len(arr); i++ {
		curr.next = &Node{data: arr[i]}
		curr = curr.next
	}
	return head
}

// Traverse
func traverse(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

// Length
func lengthOfLinkedList(head *Node) int {
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

// Search
func searchNode(head *Node, val int) bool {
	curr := head
	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

// Insert at start
func insertAtStart(head *Node, val int) *Node {
	return &Node{data: val, next: head}
}

// Insert at end
func insertAtLast(head *Node, val int) *Node {
	newNode := &Node{data: val}
	if head == nil {
		return newNode
	}

	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	return head
}

// Insert at position
func insertAtPosition(head *Node, val, pos int) *Node {
	if pos == 1 {
		return &Node{data: val, next: head}
	}

	curr := head
	count := 1

	for curr != nil && count < pos-1 {
		curr = curr.next
		count++
	}

	if curr == nil {
		return head
	}

	newNode := &Node{data: val, next: curr.next}
	curr.next = newNode
	return head
}

func insertBeforeNode(head *Node, data, val int) *Node {
	if head == nil {
		return nil
	}

	// If head itself has target value
	if head.data == val {
		return &Node{data: data, next: head}
	}

	curr := head

	for curr.next != nil {
		if curr.next.data == val {
			newNode := &Node{data: data, next: curr.next}
			curr.next = newNode
			return head
		}
		curr = curr.next
	}

	// If value not found, return original list
	return head
}

// Delete head
func deleteHeadNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.next
}

// Delete last
func deleteLastNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	curr := head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
	return head
}

// Delete kth node
func deleteKthNode(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head.next
	}

	curr := head
	prev := (*Node)(nil)
	count := 0

	for curr != nil {
		count++
		if count == k {
			prev.next = curr.next
			return head
		}
		prev = curr
		curr = curr.next
	}
	return head
}

// Delete by value
func deleteNodeWithValue(head *Node, val int) *Node {
	if head == nil {
		return nil
	}

	if head.data == val {
		return head.next
	}

	curr := head
	for curr.next != nil {
		if curr.next.data == val {
			curr.next = curr.next.next
			break
		}
		curr = curr.next
	}
	return head
}

func main() {
	arr := []int{2, 5, 8, 7, 6}
	head := convertArr2LL(arr)

	fmt.Print("Initial: ")
	traverse(head)

	head = insertAtStart(head, 1)
	fmt.Print("After insert at start: ")
	traverse(head)

	head = insertAtLast(head, 10)
	fmt.Print("After insert at last: ")
	traverse(head)

	head = deleteKthNode(head, 3)
	fmt.Print("After deleting 3rd node: ")
	traverse(head)
}