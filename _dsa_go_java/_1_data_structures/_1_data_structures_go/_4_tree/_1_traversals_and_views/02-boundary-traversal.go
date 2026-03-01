// https://www.codingninjas.com/codestudio/problems/boundary-traversal_790725?utm_source=youtube&utm_medium=affiliate&utm_campaign=Striver_Tree_Videos
// https://practice.geeksforgeeks.org/problems/boundary-traversal-of-binary-tree/0
// https://leetcode.com/problems/boundary-of-binary-tree/description/

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

func isLeaf(root *Node) bool {
	return root.Left == nil && root.Right == nil
}

// ---------- Left Boundary ----------
func addLeftBoundary(root *Node, res *[]int) {
	cur := root.Left

	for cur != nil {
		if !isLeaf(cur) {
			*res = append(*res, cur.Data)
		}

		if cur.Left != nil {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
}

// ---------- Right Boundary ----------
func addRightBoundary(root *Node, res *[]int) {
	cur := root.Right
	tmp := []int{}

	for cur != nil {
		if !isLeaf(cur) {
			tmp = append(tmp, cur.Data)
		}

		if cur.Right != nil {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	// reverse order
	for i := len(tmp) - 1; i >= 0; i-- {
		*res = append(*res, tmp[i])
	}
}

// ---------- Leaf Nodes ----------
func addLeaves(root *Node, res *[]int) {
	if root == nil {
		return
	}

	if isLeaf(root) {
		*res = append(*res, root.Data)
		return
	}

	addLeaves(root.Left, res)
	addLeaves(root.Right, res)
}

// ---------- Boundary Traversal ----------
func printBoundary(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}

	if !isLeaf(root) {
		res = append(res, root.Data)
	}

	addLeftBoundary(root, &res)
	addLeaves(root, &res)
	addRightBoundary(root, &res)

	return res
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Left.Left = newNode(3)
	root.Left.Left.Right = newNode(4)
	root.Left.Left.Right.Left = newNode(5)
	root.Left.Left.Right.Right = newNode(6)
	root.Right = newNode(7)
	root.Right.Right = newNode(8)
	root.Right.Right.Left = newNode(9)
	root.Right.Right.Left.Left = newNode(10)
	root.Right.Right.Left.Right = newNode(11)

	boundary := printBoundary(root)

	fmt.Println("The Boundary Traversal is:", boundary)
}

// Time Complexity: O(N).
// Reason: The time complexity will be O(H) + O(H) + O(N) which is ≈ O(N)
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack while adding leaves. In the worst case (skewed tree), space complexity can be O(N).