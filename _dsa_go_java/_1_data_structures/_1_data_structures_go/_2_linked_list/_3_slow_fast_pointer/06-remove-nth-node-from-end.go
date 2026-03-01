// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func insertNode(head **Node, val int) {
	newNode := &Node{val: val}

	if *head == nil {
		*head = newNode
		return
	}

	temp := *head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// brute force - length method
func deleteNthFromEnd(head *Node, n int) *Node {

	if head == nil {
		return nil
	}

	length := 0
	temp := head

	for temp != nil {
		length++
		temp = temp.next
	}

	// delete head
	if length == n {
		return head.next
	}

	pos := length - n
	temp = head

	for i := 1; i < pos; i++ {
		temp = temp.next
	}

	temp.next = temp.next.next

	return head
}
// Time Complexity: O(L)+O(L-N), We are calculating the length of the linked list and then iterating up to the (L-N)th node of the linked list, where L is the total length of the list.
// Space Complexity:  O(1), as we have not used any extra space.

// fast and slow pointer
func deleteNthFromEndTwoPointer(head *Node, n int) *Node {

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.next
	}

	// remove head
	if fast == nil {
		return head.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next

	return head
}
// Time Complexity: O(N) since the fast pointer will traverse the entire linked list, where N is the length of the linked list.
// Space Complexity: O(1), as we have not used any extra space.

// best approach - using dummy node
func removeNthFromEnd(head *Node, n int) *Node {

	dummy := &Node{next: head}

	fast := dummy
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next

	return dummy.next
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.val, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {

	var head *Node

	insertNode(&head, 1)
	insertNode(&head, 2)
	insertNode(&head, 3)
	insertNode(&head, 4)
	insertNode(&head, 5)

	head = removeNthFromEnd(head, 2)

	printList(head)
}