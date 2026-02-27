// https://leetcode.com/problems/binary-tree-paths/

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** HELPER ****************/

func helper(root *TreeNode, answer *[]string, curr string) {

	if root == nil {
		return
	}

	// if not leaf
	if root.left != nil || root.right != nil {
		curr += strconv.Itoa(root.val) + "->"
	} else {
		// leaf node
		curr += strconv.Itoa(root.val)
		*answer = append(*answer, curr)
		return
	}

	helper(root.left, answer, curr)
	helper(root.right, answer, curr)
}

/**************** MAIN FUNCTION ****************/

func binaryTreePaths(root *TreeNode) []string {

	answer := []string{}
	helper(root, &answer, "")
	return answer
}

/**************** TEST ****************/

func main() {

root := &TreeNode{val: 1}
root.left = &TreeNode{val: 2}
root.right = &TreeNode{val: 3}
root.left.right = &TreeNode{val: 5}

	res := binaryTreePaths(root)

	fmt.Println(res)
}