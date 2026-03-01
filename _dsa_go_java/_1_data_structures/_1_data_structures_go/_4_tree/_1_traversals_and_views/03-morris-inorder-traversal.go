// https://takeuforward.org/data-structure/morris-inorder-traversal-of-a-binary-tree/

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(data int) *Node {
	return &Node{Data: data}
}

func inorderTraversal(root *Node) []int {
	inorder := []int{}
	cur := root

	for cur != nil {

		// Case 1: No left child
		if cur.Left == nil {
			inorder = append(inorder, cur.Data)
			cur = cur.Right
		} else {

			// Find inorder predecessor
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			// Create thread
			if prev.Right == nil {
				prev.Right = cur
				cur = cur.Left
			} else {
				// Remove thread
				prev.Right = nil
				inorder = append(inorder, cur.Data)
				cur = cur.Right
			}
		}
	}

	return inorder
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Left.Right.Right = newNode(6)

	inorder := inorderTraversal(root)

	fmt.Println("The Inorder Traversal is:", inorder)
}
// Time Complexity: O(N).
// Space Complexity: O(1)