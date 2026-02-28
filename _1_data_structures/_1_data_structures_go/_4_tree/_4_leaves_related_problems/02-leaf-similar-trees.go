// https://leetcode.com/problems/leaf-similar-trees/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** MAIN FUNCTION ****************/

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	nodes1 := []int{}
	nodes2 := []int{}

	dfs(root1, &nodes1)
	dfs(root2, &nodes2)

	if len(nodes1) != len(nodes2) {
		return false
	}

	for i := 0; i < len(nodes1); i++ {
		if nodes1[i] != nodes2[i] {
			return false
		}
	}

	return true
}

/**************** DFS ****************/

func dfs(root *TreeNode, nodes *[]int) {

	if root == nil {
		return
	}

	// leaf node
	if root.left == nil && root.right == nil {
		*nodes = append(*nodes, root.val)
	}

	dfs(root.left, nodes)
	dfs(root.right, nodes)
}

/**************** TEST ****************/

func main() {

	root1 := &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 5,
			left: &TreeNode{val: 6},
			right: &TreeNode{
				val: 2,
				left:  &TreeNode{val: 7},
				right: &TreeNode{val: 4},
			},
		},
		right: &TreeNode{
			val: 1,
			left:  &TreeNode{val: 9},
			right: &TreeNode{val: 8},
		},
	}

	root2 := &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 5,
			left:  &TreeNode{val: 6},
			right: &TreeNode{val: 7},
		},
		right: &TreeNode{
			val: 1,
			left:  &TreeNode{val: 4},
			right: &TreeNode{
				val: 2,
				left:  &TreeNode{val: 9},
				right: &TreeNode{val: 8},
			},
		},
	}

	fmt.Println(leafSimilar(root1, root2)) // true
}