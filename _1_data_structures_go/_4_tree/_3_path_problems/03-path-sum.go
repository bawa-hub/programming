// https://leetcode.com/problems/path-sum/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** MAIN FUNCTION ****************/

func hasPathSum(root *TreeNode, targetSum int) bool {
	return helper(root, 0, targetSum)
}

/**************** HELPER ****************/

func helper(root *TreeNode, sum int, target int) bool {

	if root == nil {
		return false
	}

	sum += root.val

	// leaf node
	if root.left == nil && root.right == nil {
		return sum == target
	}

	return helper(root.left, sum, target) ||
		helper(root.right, sum, target)
}

/**************** TEST ****************/

func main() {

	root := &TreeNode{
		val: 5,
		left: &TreeNode{
			val: 4,
			left: &TreeNode{
				val: 11,
				left:  &TreeNode{val: 7},
				right: &TreeNode{val: 2},
			},
		},
		right: &TreeNode{
			val: 8,
			left:  &TreeNode{val: 13},
			right: &TreeNode{val: 4, right: &TreeNode{val: 1}},
		},
	}

	fmt.Println(hasPathSum(root, 22)) // true
}