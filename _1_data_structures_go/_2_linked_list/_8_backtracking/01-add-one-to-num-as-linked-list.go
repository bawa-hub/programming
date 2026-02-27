// https://practice.geeksforgeeks.org/problems/add-1-to-a-number-represented-as-linked-list/1

package main

import "fmt"

type Node struct {
	data int
	next *Node
}


// recursive
func addOne(head *Node) *Node {

	carry := addHelper(head)

	if carry == 1 {
		newNode := &Node{data: 1}
		newNode.next = head
		head = newNode
	}

	return head
}
// TC: O(n)
// SC: O(n)

func addHelper(temp *Node) int {

	// add 1 at end
	if temp == nil {
		return 1
	}

	carry := addHelper(temp.next)

	temp.data += carry

	if temp.data < 10 {
		return 0
	}

	temp.data = 0
	return 1
}

// iterative
func addOneIterative(head *Node) *Node {

	// dummy handles cases like 999 -> 1000
	dummy := &Node{data: 0}
	dummy.next = head

	lastNonNine := dummy
	curr := head

	// find last non-9 node
	for curr != nil {
		if curr.data != 9 {
			lastNonNine = curr
		}
		curr = curr.next
	}

	// increment
	lastNonNine.data++

	// set remaining digits to 0
	curr = lastNonNine.next
	for curr != nil {
		curr.data = 0
		curr = curr.next
	}

	// if dummy changed → new head
	if dummy.data == 1 {
		return dummy
	}

	return dummy.next
}
// TC : O(3N)
// SC : O(1)

/************ PRINT ************/

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data)
		head = head.next
	}
	fmt.Println()
}

/************ MAIN ************/

func main() {

	head := &Node{1,
		&Node{9,
			&Node{9, nil},
		},
	}

	printList(head)

	head = addOne(head)

	printList(head)
}