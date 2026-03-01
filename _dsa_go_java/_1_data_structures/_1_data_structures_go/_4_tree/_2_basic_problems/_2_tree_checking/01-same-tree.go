// https://leetcode.com/problems/same-tree/

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func isIdentical(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Data == node2.Data &&
		isIdentical(node1.Left, node2.Left) &&
		isIdentical(node1.Right, node2.Right)
}

func newNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func main() {

	root1 := newNode(1)
	root1.Left = newNode(2)
	root1.Right = newNode(3)
	root1.Right.Left = newNode(4)
	root1.Right.Right = newNode(5)

	root2 := newNode(1)
	root2.Left = newNode(2)
	root2.Right = newNode(3)
	root2.Right.Left = newNode(4)

	if isIdentical(root1, root2) {
		fmt.Println("Two Trees are identical")
	} else {
		fmt.Println("Two trees are non-identical")
	}
}

// Time Complexity: O(N).
// Reason: We are doing a tree traversal.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).