// https://leetcode.com/problems/linked-list-cycle-ii/

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

func createCycle(head *Node, pos int) {

	if head == nil {
		return
	}

	ptr := head
	temp := head
	cnt := 0

	for temp.next != nil {

		if cnt < pos {
			ptr = ptr.next
			cnt++
		}

		temp = temp.next
	}

	temp.next = ptr
}

// hashing method
func detectCycleHash(head *Node) *Node {

	visited := make(map[*Node]bool)

	for head != nil {

		if visited[head] {
			return head
		}

		visited[head] = true
		head = head.next
	}

	return nil
}

// Time Complexity: O(N)
// Reason: Iterating the entire list once.
// Space Complexity: O(N)
// Reason: We store all nodes in a hash table.

// optimal - floyd cycles algo
func detectCycle(head *Node) *Node {

	if head == nil || head.next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next

		if slow == fast {

			entry := head

			for entry != slow {
				entry = entry.next
				slow = slow.next
			}

			return entry
		}
	}

	return nil
}
// Time Complexity: O(N)
// Reason: We can take overall iterations and club them to O(N)
// Space Complexity: O(1)
// Reason: No extra data structure is used.

func main() {

	var head *Node

	insertNode(&head, 1)
	insertNode(&head, 2)
	insertNode(&head, 3)
	insertNode(&head, 4)
	insertNode(&head, 5)
	insertNode(&head, 6)

	createCycle(head, 2)

	nodeReceive := detectCycle(head)

	if nodeReceive == nil {
		fmt.Println("No cycle")
	} else {

		temp := head
		pos := 0

		for temp != nodeReceive {
			temp = temp.next
			pos++
		}

		fmt.Println("Tail connects at pos", pos)
	}
}