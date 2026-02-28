// https://practice.geeksforgeeks.org/problems/k-distance-from-root/1


package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

type Pair struct {
	Node  *Node
	Level int
}

func Kdistance(root *Node, k int) {
	if root == nil {
		return
	}

	queue := []Pair{{root, 0}}

	for len(queue) > 0 {

		size := len(queue)

		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			node := p.Node
			level := p.Level

			if level == k {
				fmt.Println(node.Data)
			} else {
				if node.Left != nil {
					queue = append(queue, Pair{node.Left, level + 1})
				}
				if node.Right != nil {
					queue = append(queue, Pair{node.Right, level + 1})
				}
			}
		}
	}
}

func main() {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Right.Left = newNode(8)

	Kdistance(root, 2)
}