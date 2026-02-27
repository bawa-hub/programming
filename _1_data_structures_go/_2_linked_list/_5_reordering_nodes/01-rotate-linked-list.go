// https://leetcode.com/problems/rotate-list/

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

/******** BRUTE FORCE ROTATE ********/

func rotateRight(head *Node, k int) *Node {

	if head == nil || head.next == nil {
		return head
	}

	for i := 0; i < k; i++ {

		temp := head

		for temp.next.next != nil {
			temp = temp.next
		}

		end := temp.next
		temp.next = nil
		end.next = head
		head = end
	}

	return head
}
// Time Complexity: O(Number of nodes present in the list*k)
// Reason: For k times, we are iterating through the entire list to get the last element and move it to first.
// Space Complexity: O(1)
// Reason: No extra data structures is used for computations

// optimal
func rotateRightOptimal(head *Node, k int) *Node {

	if head == nil || head.next == nil || k == 0 {
		return head
	}

	// find length
	temp := head
	length := 1

	for temp.next != nil {
		temp = temp.next
		length++
	}

	// make circular list
	temp.next = head

	// reduce rotations
	k = k % length

	end := length - k

	for end > 0 {
		temp = temp.next
		end--
	}

	// new head
	head = temp.next

	// break circle
	temp.next = nil

	return head
}
// Time Complexity: O(length of list) + O(length of list – (length of list%k))
// Reason: O(length of the list) for calculating the length of the list. O(length of the list – (length of list%k)) for breaking link.
// Space Complexity: O(1)
// Reason: No extra data structure is used for computation.

/******** PRINT ********/

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.num, " ")
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

	head = rotateRight(head, 2)
	printList(head)
}