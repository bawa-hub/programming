// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

/******** DELETE MIDDLE ********/

func deleteMiddle(head *ListNode) *ListNode {

	if head == nil || head.next == nil {
		return nil
	}

	dummy := &ListNode{next: head}

	slow := dummy
	fast := dummy

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// delete middle
	slow.next = slow.next.next

	return dummy.next
}

/******** HELPERS ********/

func insert(head **ListNode, val int) {
	newNode := &ListNode{val: val}

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

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.val, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {

	var head *ListNode

	insert(&head, 1)
	insert(&head, 3)
	insert(&head, 4)
	insert(&head, 7)
	insert(&head, 1)
	insert(&head, 2)
	insert(&head, 6)

	head = deleteMiddle(head)

	printList(head)
}