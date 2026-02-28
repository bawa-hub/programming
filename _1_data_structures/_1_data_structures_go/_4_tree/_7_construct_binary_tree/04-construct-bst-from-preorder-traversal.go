// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** MAIN FUNCTION ****************/

func bstFromPreorder(preorder []int) *TreeNode {
	i := 0
	return build(preorder, &i, math.MaxInt)
}

/**************** BUILD FUNCTION ****************/

func build(A []int, i *int, bound int) *TreeNode {

	if *i == len(A) || A[*i] > bound {
		return nil
	}

	root := &TreeNode{val: A[*i]}
	*i++

	root.left = build(A, i, root.val)
	root.right = build(A, i, bound)

	return root
}

/**************** TEST ****************/

func main() {

	preorder := []int{8, 5, 1, 7, 10, 12}

	root := bstFromPreorder(preorder)

	fmt.Println("Root:", root.val)
	fmt.Println("Left:", root.left.val)
	fmt.Println("Right:", root.right.val)
}