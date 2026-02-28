// https://leetcode.com/problems/path-sum-iii/
// https://leetcode.com/problems/path-sum-iii/solutions/779575/c-3-dfs-based-solutions-explained-and-compared-up-to-100-time-75-space/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** SOLUTION ****************/

func pathSum(root *TreeNode, targetSum int) int {

	count := 0
	mp := make(map[int64]int)

	calculate(root, int64(targetSum), &count, 0, mp)

	return count
}

func calculate(root *TreeNode,
	target int64,
	count *int,
	curr int64,
	mp map[int64]int) {

	if root == nil {
		return
	}

	curr += int64(root.val)

	// case 1: path from root
	if curr == target {
		*count++
	}

	// case 2: prefix sum exists
	if val, ok := mp[curr-target]; ok {
		*count += val
	}

	// store prefix sum
	mp[curr]++

	calculate(root.left, target, count, curr, mp)
	calculate(root.right, target, count, curr, mp)

	// backtrack
	mp[curr]--
}

/**************** TEST ****************/

func main() {

	root := &TreeNode{
		val: 10,
		left: &TreeNode{
			val: 5,
			left: &TreeNode{
				val: 3,
				left:  &TreeNode{val: 3},
				right: &TreeNode{val: -2},
			},
			right: &TreeNode{
				val: 2,
				right: &TreeNode{val: 1},
			},
		},
		right: &TreeNode{
			val: -3,
			right: &TreeNode{val: 11},
		},
	}

	fmt.Println(pathSum(root, 8)) // Output: 3
}