
// https://leetcode.com/problems/reverse-nodes-in-k-group/

package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

/******** INSERT NODE ********/

func insertNode(head **Node, val int) {
	newNode := &Node{num: val}

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

/******** LENGTH ********/

func lengthOfLinkedList(head *Node) int {
	length := 0
	for head != nil {
		length++
		head = head.next
	}
	return length
}

/******** REVERSE K GROUP ********/

func reverseKNodes(head *Node, k int) *Node {

	if head == nil || head.next == nil {
		return head
	}

	length := lengthOfLinkedList(head)

	dummy := &Node{}
	dummy.next = head

	pre := dummy

	for length >= k {

		cur := pre.next
		nex := cur.next

		for i := 1; i < k; i++ {
			cur.next = nex.next
			nex.next = pre.next
			pre.next = nex
			nex = cur.next
		}

		pre = cur
		length -= k
	}

	return dummy.next
}

/******** PRINT ********/

func printList(head *Node) {
	for head != nil {
		if head.next != nil {
			fmt.Print(head.num, "->")
		} else {
			fmt.Print(head.num)
		}
		head = head.next
	}
	fmt.Println()
}

/******** MAIN ********/

func main() {

	var head *Node
	k := 3

	insertNode(&head, 1)
	insertNode(&head, 2)
	insertNode(&head, 3)
	insertNode(&head, 4)
	insertNode(&head, 5)
	insertNode(&head, 6)
	insertNode(&head, 7)
	insertNode(&head, 8)

	fmt.Print("Original: ")
	printList(head)

	head = reverseKNodes(head, k)

	fmt.Print("After Reverse K: ")
	printList(head)
}


// Time Complexity: O(N)
// Reason: Nested iteration with O((N/k)*k) which makes it equal to O(N).
// Space Complexity: O(1)
// Reason: No extra data structures are used for computation