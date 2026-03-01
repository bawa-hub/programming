// https://leetcode.com/problems/sum-of-left-leaves/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** MAIN FUNCTION ****************/

func sumOfLeftLeaves(root *TreeNode) int {

	sum := 0
	helper(root, &sum, -1)
	return sum
}

/**************** HELPER ****************/

func helper(root *TreeNode, sum *int, lr int) {

	if root == nil {
		return
	}

	// left leaf
	if root.left == nil && root.right == nil && lr == 0 {
		*sum += root.val
	}

	helper(root.left, sum, 0)
	helper(root.right, sum, 1)
}

/**************** TEST ****************/

func main() {

	root := &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 9,
		},
		right: &TreeNode{
			val: 20,
			left:  &TreeNode{val: 15},
			right: &TreeNode{val: 7},
		},
	}

	fmt.Println(sumOfLeftLeaves(root)) // 24
}