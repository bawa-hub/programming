// https://leetcode.com/problems/intersection-of-two-linked-lists/

package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

/************ INSERT NODE ************/

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

/************ INTERSECTION CHECK ************/

// brute force
func intersectionPresent(head1, head2 *Node) *Node {

	for head2 != nil {

		temp := head1

		for temp != nil {
			if temp == head2 {
				return head2
			}
			temp = temp.next
		}

		head2 = head2.next
	}

	return nil
}
// Time Complexity: O(m*n)
// Reason: For each node in list 2 entire lists 1 are iterated.
// Space Complexity: O(1)
// Reason: No extra space is used.

// hashing

func intersectionPresentHashing(head1, head2 *Node) *Node {

	visited := make(map[*Node]bool)

	for head1 != nil {
		visited[head1] = true
		head1 = head1.next
	}

	for head2 != nil {
		if visited[head2] {
			return head2
		}
		head2 = head2.next
	}

	return nil
}
// Time Complexity: O(n+m)
// Reason: Iterating through list 1 first takes O(n), then iterating through list 2 takes O(m).
// Space Complexity: O(n)
// Reason: Storing list 1 node addresses in unordered_set.

// using length difference
func getDifference(head1, head2 *Node) int {

	len1, len2 := 0, 0

	for head1 != nil || head2 != nil {

		if head1 != nil {
			len1++
			head1 = head1.next
		}

		if head2 != nil {
			len2++
			head2 = head2.next
		}
	}

	return len1 - len2
}

func intersectionPresentUsingDiff(head1, head2 *Node) *Node {

	diff := getDifference(head1, head2)

	if diff < 0 {
		for diff != 0 {
			head2 = head2.next
			diff++
		}
	} else {
		for diff != 0 {
			head1 = head1.next
			diff--
		}
	}

	for head1 != nil && head2 != nil {

		if head1 == head2 {
			return head1
		}

		head1 = head1.next
		head2 = head2.next
	}

	return nil
}
// Time Complexity:
// O(2max(length of list1,length of list2))+O(abs(length of list1-length of list2))+O(min(length of list1,length of list2))
// Reason: Finding the length of both lists takes max(length of list1, length of list2) because it is found simultaneously for both of them. Moving the head pointer ahead by a difference of them. The next one is for searching.
// Space Complexity: O(1)
// Reason: No extra space is used.

// optimal using two pointer
func intersectionPresentOptimal(head1, head2 *Node) *Node {

	d1 := head1
	d2 := head2

	for d1 != d2 {

		if d1 == nil {
			d1 = head2
		} else {
			d1 = d1.next
		}

		if d2 == nil {
			d2 = head1
		} else {
			d2 = d2.next
		}
	}

	return d1
}
// Time Complexity: O(2*max(length of list1,length of list2))
// Reason: Uses the same concept of the difference of lengths of two lists.
// Space Complexity: O(1)
// Reason: No extra data structure is used

/************ PRINT LIST ************/

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

/************ MAIN ************/

func main() {

	var head *Node
	insertNode(&head, 1)
	insertNode(&head, 3)
	insertNode(&head, 1)
	insertNode(&head, 2)
	insertNode(&head, 4)

	head1 := head

	// intersection node
	intersect := head.next.next.next

	var headSec *Node
	insertNode(&headSec, 3)

	head2 := headSec
	headSec.next = intersect

	fmt.Print("List1: ")
	printList(head1)

	fmt.Print("List2: ")
	printList(head2)

	ans := intersectionPresent(head1, head2)

	if ans == nil {
		fmt.Println("No intersection")
	} else {
		fmt.Println("Intersection at:", ans.num)
	}
}