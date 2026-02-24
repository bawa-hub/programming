// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tempi := head
	tempj := head

	for tempj != nil {
		if tempj.Val != tempi.Val {
			tempi = tempi.Next
			tempi.Val = tempj.Val
		}
		tempj = tempj.Next
	}

	// Cut off extra nodes
	tempi.Next = nil

	return head
}

// better
func deleteDuplicatesOptimized(head *ListNode) *ListNode {
	temp := head

	for temp != nil && temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
			continue
		}
		temp = temp.Next
	}

	return head
}