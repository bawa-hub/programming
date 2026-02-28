// https://leetcode.com/problems/symmetric-tree/

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func isSymmetricUtil(root1, root2 *Node) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}

	return root1.Data == root2.Data &&
		isSymmetricUtil(root1.Left, root2.Right) &&
		isSymmetricUtil(root1.Right, root2.Left)
}

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isSymmetricUtil(root.Left, root.Right)
}

func newNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Left.Left = newNode(3)
	root.Left.Right = newNode(4)
	root.Right = newNode(2)
	root.Right.Left = newNode(4)
	root.Right.Right = newNode(3)

	res := isSymmetric(root)

	if res {
		fmt.Println("The tree is symmetrical")
	} else {
		fmt.Println("The tree is NOT symmetrical")
	}
}

// Time Complexity: O(N)
// Reason: We are doing simple tree traversal and changing both root1 and root2 simultaneously.
// Space Complexity: O(N)
// Reason: In the worst case (skewed tree), space complexity can be O(N)