// https://www.codingninjas.com/codestudio/problems/981269?topList=striver-sde-sheet-problems&utm_source=striver&utm_medium=website

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Pair struct {
	Node  *Node
	State int
}

func allTraversal(root *Node, pre, in, post *[]int) {
	if root == nil {
		return
	}

	stack := []Pair{{root, 1}}

	for len(stack) > 0 {

		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Preorder
		if it.State == 1 {
			*pre = append(*pre, it.Node.Data)

			it.State++
			stack = append(stack, it)

			if it.Node.Left != nil {
				stack = append(stack, Pair{it.Node.Left, 1})
			}

		} else if it.State == 2 { // Inorder
			*in = append(*in, it.Node.Data)

			it.State++
			stack = append(stack, it)

			if it.Node.Right != nil {
				stack = append(stack, Pair{it.Node.Right, 1})
			}

		} else { // Postorder
			*post = append(*post, it.Node.Data)
		}
	}
}

func newNode(data int) *Node {
	return &Node{Data: data}
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Right = newNode(3)
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(7)

	pre := []int{}
	in := []int{}
	post := []int{}

	allTraversal(root, &pre, &in, &post)

	fmt.Println("The preorder Traversal is :", pre)
	fmt.Println("The inorder Traversal is :", in)
	fmt.Println("The postorder Traversal is :", post)
}