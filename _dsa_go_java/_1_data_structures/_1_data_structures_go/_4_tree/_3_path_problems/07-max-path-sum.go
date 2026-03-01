// https://leetcode.com/problems/binary-tree-maximum-path-sum/

package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

/**************** CORE LOGIC ****************/

func findMaxPathSum(root *Node, maxi *int) int {

	if root == nil {
		return 0
	}

	leftMaxPath := max(0, findMaxPathSum(root.left, maxi))
	rightMaxPath := max(0, findMaxPathSum(root.right, maxi))

	val := root.data

	// path passing through current node
	*maxi = max(*maxi, leftMaxPath+rightMaxPath+val)

	// return best single path upward
	return max(leftMaxPath, rightMaxPath) + val
}

func maxPathSum(root *Node) int {

	maxi := math.MinInt
	findMaxPathSum(root, &maxi)
	return maxi
}

/**************** UTIL ****************/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**************** TEST ****************/

func main() {

	root := &Node{data: -10}
	root.left = &Node{data: 9}
	root.right = &Node{data: 20}
	root.right.left = &Node{data: 15}
	root.right.right = &Node{data: 7}

	answer := maxPathSum(root)

	fmt.Println("The Max Path Sum for this tree is", answer)
}