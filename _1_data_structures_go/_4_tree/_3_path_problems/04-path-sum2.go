// https://leetcode.com/problems/path-sum-ii/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** MAIN FUNCTION ****************/

func pathSum(root *TreeNode, targetSum int) [][]int {

	res := [][]int{}
	helper(root, 0, []int{}, &res, targetSum)
	return res
}

/**************** HELPER ****************/

func helper(root *TreeNode, sum int, curr []int,
	res *[][]int, target int) {

	if root == nil {
		return
	}

	sum += root.val
	curr = append(curr, root.val)

	// leaf node
	if root.left == nil && root.right == nil {
		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			*res = append(*res, temp)
		}
		return
	}

	helper(root.left, sum, curr, res, target)
	helper(root.right, sum, curr, res, target)
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
			right: &TreeNode{
				val: 4,
				left:  &TreeNode{val: 5},
				right: &TreeNode{val: 1},
			},
		},
	}

	fmt.Println(pathSum(root, 22))
}