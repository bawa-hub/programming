// https://leetcode.com/problems/maximum-width-of-binary-tree/

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

type Pair struct {
	Node *Node
	ID   int
}

func widthOfBinaryTree(root *Node) int {
	if root == nil {
		return 0
	}

	ans := 0
	queue := []Pair{{root, 0}}

	for len(queue) > 0 {
		size := len(queue)
		curMin := queue[0].ID

		var leftMost, rightMost int

		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			curID := p.ID - curMin // prevent overflow
			node := p.Node

			if i == 0 {
				leftMost = curID
			}
			if i == size-1 {
				rightMost = curID
			}

			if node.Left != nil {
				queue = append(queue, Pair{
					node.Left,
					curID*2 + 1,
				})
			}

			if node.Right != nil {
				queue = append(queue, Pair{
					node.Right,
					curID*2 + 2,
				})
			}
		}

		width := rightMost - leftMost + 1
		if width > ans {
			ans = width
		}
	}

	return ans
}

func main() {

	root := newNode(1)
	root.Left = newNode(3)
	root.Left.Left = newNode(5)
	root.Left.Left.Left = newNode(7)
	root.Right = newNode(2)
	root.Right.Right = newNode(4)
	root.Right.Right.Right = newNode(6)

	maxWidth := widthOfBinaryTree(root)
	fmt.Println("The maximum width of the Binary Tree is", maxWidth)
}


// Time Complexity: O(N)
// Space Complexity: O(N)