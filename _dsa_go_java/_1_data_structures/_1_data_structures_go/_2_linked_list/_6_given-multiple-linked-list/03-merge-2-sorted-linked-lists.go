// https://leetcode.com/problems/merge-two-sorted-lists/

// using external linked list
// TC: O(n+m)
// SC: O(n+m)

package main

import "fmt"

/************ LIST NODE ************/

type ListNode struct {
	Val  int
	Next *ListNode
}

/************ MERGE ************/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// if one list empty
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// ensure l1 starts with smaller value
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	res := l1

	for l1 != nil && l2 != nil {

		var temp *ListNode

		// move in sorted order
		for l1 != nil && l1.Val <= l2.Val {
			temp = l1
			l1 = l1.Next
		}

		// connect lists
		temp.Next = l2

		// swap pointers
		l1, l2 = l2, l1
	}

	return res
}
// TC: O(n+m)
// SC: O(1)


/************ HELPERS ************/

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {

	l1 := &ListNode{1,
		&ListNode{3,
			&ListNode{5, nil},
		},
	}

	l2 := &ListNode{2,
		&ListNode{4,
			&ListNode{6, nil},
		},
	}

	result := mergeTwoLists(l1, l2)
	printList(result)
}