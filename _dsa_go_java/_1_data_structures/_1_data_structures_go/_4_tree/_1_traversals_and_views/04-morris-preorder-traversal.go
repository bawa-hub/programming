// https://takeuforward.org/data-structure/morris-preorder-traversal-of-a-binary-tree/

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

func preorderTraversal(root *Node) []int {
	preorder := []int{}
	cur := root

	for cur != nil {

		// Case 1: No left child
		if cur.Left == nil {
			preorder = append(preorder, cur.Data)
			cur = cur.Right
		} else {

			// Find inorder predecessor
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			// Create thread
			if prev.Right == nil {
				preorder = append(preorder, cur.Data)
				prev.Right = cur
				cur = cur.Left
			} else {
				// Remove thread
				prev.Right = nil
				cur = cur.Right
			}
		}
	}

	return preorder
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Left.Right.Right = newNode(6)

	preorder := preorderTraversal(root)

	fmt.Println("The Preorder Traversal is:", preorder)
}
// Time Complexity: O(N).
// Space Complexity: O(1)