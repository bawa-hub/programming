// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func insertBST(root *Node, val int) *Node {
	if root == nil {
		return &Node{Data: val}
	}

	if val < root.Data {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}

	return root
}

// ---------- Kth Largest ----------
func kthLargest(root *Node, k *int) *Node {
	if root == nil {
		return nil
	}

	right := kthLargest(root.Right, k)
	if right != nil {
		return right
	}

	*k--
	if *k == 0 {
		return root
	}

	return kthLargest(root.Left, k)
}

// ---------- Kth Smallest ----------
func kthSmallest(root *Node, k *int) *Node {
	if root == nil {
		return nil
	}

	left := kthSmallest(root.Left, k)
	if left != nil {
		return left
	}

	*k--
	if *k == 0 {
		return root
	}

	return kthSmallest(root.Right, k)
}

func main() {

	arr := []int{10, 40, 45, 20, 25, 30, 50}
	var root *Node

	for _, v := range arr {
		root = insertBST(root, v)
	}

	k := 3
	kCopy := k

	large := kthLargest(root, &k)
	k = kCopy
	small := kthSmallest(root, &k)

	if large == nil {
		fmt.Println("Invalid input")
	} else {
		fmt.Println("kth largest element is", large.Data)
	}

	if small == nil {
		fmt.Println("Invalid input")
	} else {
		fmt.Println("kth smallest element is", small.Data)
	}
}
// Time Complexity: O(min(K,N))
// Space Complexity: O(min(K,N))

// for leetcode

// recursive

func kthSmallest(root *TreeNode, k int) int {
	var result int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		k--
		if k == 0 {
			result = node.Val
			return
		}

		inorder(node.Right)
	}

	inorder(root)
	return result
}

// iterative
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

	return -1
}