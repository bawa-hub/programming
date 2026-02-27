// https://leetcode.com/problems/odd-even-linked-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/************ ODD EVEN LIST ************/

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head
	evenStart := head.Next
	even := head.Next

	for odd.Next != nil && even.Next != nil {

		// connect odd nodes
		odd.Next = odd.Next.Next
		odd = odd.Next

		// connect even nodes
		even.Next = even.Next.Next
		even = even.Next
	}

	// attach even list after odd list
	odd.Next = evenStart

	return head
}
// Time complexity: O(n)
// Space complexity: O(1)

/************ HELPER ************/

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

/************ MAIN ************/

func main() {

	head := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil},
				},
			},
		},
	}

	head = oddEvenList(head)
	printList(head)
}