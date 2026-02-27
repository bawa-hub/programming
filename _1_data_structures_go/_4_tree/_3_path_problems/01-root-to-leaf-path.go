// https://practice.geeksforgeeks.org/problems/root-to-leaf-paths/1
// https://www.geeksforgeeks.org/given-a-binary-tree-print-all-root-to-leaf-paths/

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

/**************** HELPER ****************/

func helper(root *Node, path []int, ans *[][]int) {

	if root == nil {
		return
	}

	// add current node
	path = append(path, root.data)

	// leaf node
	if root.left == nil && root.right == nil {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	helper(root.left, path, ans)
	helper(root.right, path, ans)
}

/**************** MAIN FUNCTION ****************/

func Paths(root *Node) [][]int {

	ans := [][]int{}
	if root == nil {
		return ans
	}

	helper(root, []int{}, &ans)
	return ans
}

/**************** TEST ****************/

func main() {

	root := &Node{1}
	root.left = &Node{2}
	root.right = &Node{3}
	root.left.left = &Node{4}
	root.left.right = &Node{5}

	result := Paths(root)

	for _, p := range result {
		fmt.Println(p)
	}
}