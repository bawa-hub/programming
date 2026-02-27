// https://leetcode.com/problems/palindrome-linked-list/

package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

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

// brute force using array
func isPalindromeArray(head *Node) bool {

	arr := []int{}

	for head != nil {
		arr = append(arr, head.num)
		head = head.next
	}

	n := len(arr)

	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}

	return true
}
// Time Complexity: O(2N)
// Reason: Iterating through the list to store elements in the array.
// Space Complexity: O(N)
// Reason: Using an array to store list elements for further computations.


// using stack
func isPalindromeStack(head *Node) bool {

	stack := []int{}
	temp := head

	for temp != nil {
		stack = append(stack, temp.num)
		temp = temp.next
	}

	temp = head

	for temp != nil {
		n := len(stack)

		if temp.num != stack[n-1] {
			return false
		}

		stack = stack[:n-1]
		temp = temp.next
	}

	return true
}
// Time Complexity: O(2 * N) This is because we traverse the linked list twice: once to push the values onto the stack, and once to pop the values and compare with the linked list. Both traversals take O(2*N) ~ O(N) time.
// Space Complexity: O(N) We use a stack to store the values of the linked list, and in the worst case, the stack will have all N values,  ie. storing the complete linked list. 


// optimal
func isPalindrome(head *Node) bool {

	if head == nil || head.next == nil {
		return true
	}

	slow := head
	fast := head

	// find middle
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// reverse second half
	secondHalf := reverse(slow.next)

	first := head
	second := secondHalf

	for second != nil {
		if first.num != second.num {
			return false
		}
		first = first.next
		second = second.next
	}

	return true
}
// Time Complexity: O(N/2)+O(N/2)+O(N/2)
// Reason: O(N/2) for finding the middle element, reversing the list from the middle element, and traversing again to find palindrome respectively.
// Space Complexity: O(1)
// Reason: No extra data structures are used.

func reverse(head *Node) *Node {

	var prev *Node
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}