// https://leetcode.com/problems/reverse-linked-list/

package main

import "fmt"

/************ NODE ************/

type Node struct {
	data int
	next *Node
}

/************ INSERT ************/

func insertNode(head **Node, val int) {
	newNode := &Node{data: val}

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

// using stack
func reverseUsingStack(head *Node) *Node {

	if head == nil {
		return head
	}

	stack := []int{}
	temp := head

	// push values
	for temp != nil {
		stack = append(stack, temp.data)
		temp = temp.next
	}

	temp = head

	// pop values
	for temp != nil {
		n := len(stack)
		temp.data = stack[n-1]
		stack = stack[:n-1]
		temp = temp.next
	}

	return head
}
// Time Complexity: O(2N) This is because we traverse the linked list twice: once to push the values onto the stack, and once to pop the values and update the linked list. Both traversals take O(N) time, hence time complexity  O(2N) ~ O(N).
// Space Complexity: O(N) We use a stack to store the values of the linked list, and in the worst case, the stack will have all N values,  ie. storing the complete linked list. 

// 3 pointer method
func reverseIterative(head *Node) *Node {

	var prev *Node
	temp := head

	for temp != nil {

		front := temp.next
		temp.next = prev

		prev = temp
		temp = front
	}

	return prev
}
// Time Complexity: O(N) The code traverses the entire linked list once, where ‘n’ is the number of nodes in the list. This traversal has a linear time complexity, O(n).
// Space Complexity: O(1) The code uses only a constant amount of additional space, regardless of the linked list’s length. This is achieved by using three pointers (prev, temp and front) to reverse the list without any significant extra memory usage, resulting in constant space complexity, O(1).

// recursive
func reverseRecursive(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	newHead := reverseRecursive(head.next)

	front := head.next
	front.next = head
	head.next = nil

	return newHead
}
// Time Complexity: O(N) This is because we traverse the linked list twice: once to push the values onto the stack, and once to pop the values and update the linked list. Both traversals take O(N) time.
// Space Complexity : O(1) No additional space is used explicitly for data structures or allocations during the linked list reversal process. However, it’s important to note that there is an implicit use of stack space due to recursion. This recursive stack space stores function calls and associated variables during the recursive traversal and reversal of the linked list. Despite this, no extra memory beyond the program’s existing execution space is allocated, hence maintaining a space complexity of O(1).


func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
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

	fmt.Print("Original: ")
	printList(head)

	head = reverseIterative(head)

	fmt.Print("Reversed: ")
	printList(head)
}