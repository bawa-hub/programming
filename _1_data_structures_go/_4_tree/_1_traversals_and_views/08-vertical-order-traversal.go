// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/


package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Item struct {
	Node *Node
	X    int // vertical
	Y    int // level
}

func findVertical(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	// x -> y -> values
	nodes := make(map[int]map[int][]int)

	queue := []Item{{root, 0, 0}}

	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]

		node := it.Node
		x, y := it.X, it.Y

		if _, ok := nodes[x]; !ok {
			nodes[x] = make(map[int][]int)
		}

		nodes[x][y] = append(nodes[x][y], node.Data)

		if node.Left != nil {
			queue = append(queue, Item{node.Left, x - 1, y + 1})
		}

		if node.Right != nil {
			queue = append(queue, Item{node.Right, x + 1, y + 1})
		}
	}

	// sort vertical keys
	xKeys := []int{}
	for k := range nodes {
		xKeys = append(xKeys, k)
	}
	sort.Ints(xKeys)

	ans := [][]int{}

	for _, x := range xKeys {
		col := []int{}

		// sort levels
		yKeys := []int{}
		for k := range nodes[x] {
			yKeys = append(yKeys, k)
		}
		sort.Ints(yKeys)

		for _, y := range yKeys {
			vals := nodes[x][y]
			sort.Ints(vals) // simulate multiset behavior
			col = append(col, vals...)
		}

		ans = append(ans, col)
	}

	return ans
}

func newNode(data int) *Node {
	return &Node{Data: data}
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(10)
	root.Left.Left.Right = newNode(5)
	root.Left.Left.Right.Right = newNode(6)
	root.Right = newNode(3)
	root.Right.Left = newNode(9)
	root.Right.Right = newNode(10)

	result := findVertical(root)

	fmt.Println("The Vertical Traversal is:")
	for _, col := range result {
		fmt.Println(col)
	}
}


// Time Complexity: O(N*logN*logN*logN)
// Space Complexity: O(N)