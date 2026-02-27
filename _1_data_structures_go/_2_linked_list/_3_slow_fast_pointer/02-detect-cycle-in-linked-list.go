// https://leetcode.com/problems/linked-list-cycle/

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

func createCycle(head *Node, a, b int) {

	p1 := head
	p2 := head

	cnta := 0
	cntb := 0

	for cnta != a {
		p1 = p1.next
		cnta++
	}

	for cntb != b {
		p2 = p2.next
		cntb++
	}

	p2.next = p1
}

// hsashing
func cycleDetectHash(head *Node) bool {

	visited := make(map[*Node]bool)

	for head != nil {

		if visited[head] {
			return true
		}

		visited[head] = true
		head = head.next
	}

	return false
}
// Time Complexity: O(N)
// Reason: Entire list is iterated once.
// Space Complexity: O(N)
// Reason: All nodes present in the list are stored in a hash table.


// floyds cycle algo - fast slow pointer
func cycleDetect(head *Node) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}
// Time Complexity: O(N)
// Reason: In the worst case, all the nodes of the list are visited.
// Space Complexity: O(1)
// Reason: No extra data structure is used.

func main() {

	var head *Node

	insertNode(&head, 1)
	insertNode(&head, 2)
	insertNode(&head, 3)
	insertNode(&head, 4)

	createCycle(head, 1, 3)

	if cycleDetect(head) {
		fmt.Println("Cycle detected")
	} else {
		fmt.Println("Cycle not detected")
	}
}