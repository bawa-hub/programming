// https://leetcode.com/problems/middle-of-the-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// brute force
func middleNode(head *ListNode) *ListNode {
	n := 0
	temp := head

	for temp != nil {
		n++
		temp = temp.Next
	}

	temp = head
	for i := 0; i < n/2; i++ {
		temp = temp.Next
	}

	return temp
}

// Time Complexity: O(N) + O(N/2)
// Space Complexity: O(1)

// fast and slow pointer
func middleNodeOptimized(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Time Complexity: O(N/2)
// Space Complexity: O(1)
