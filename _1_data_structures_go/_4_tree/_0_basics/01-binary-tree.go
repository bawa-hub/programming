// https://leetcode.com/problems/binary-tree-preorder-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/

package main

import (
	"fmt"
)

/*********************** NODE *************************/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func createNode(data int) *Node {
	return &Node{data: data}
}

/*********************** INSERT NODE *************************/
// Level order insertion
func insertNode(root *Node, data int) *Node {

	if root == nil {
		return createNode(data)
	}

	queue := []*Node{root}

	for len(queue) > 0 {

		temp := queue[0]
		queue = queue[1:]

		if temp.left != nil {
			queue = append(queue, temp.left)
		} else {
			temp.left = createNode(data)
			return root
		}

		if temp.right != nil {
			queue = append(queue, temp.right)
		} else {
			temp.right = createNode(data)
			return root
		}
	}

	return root
}

/*********************** DFS *************************/

func preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	preorder(root.left)
	preorder(root.right)
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).


func inorder(root *Node) {
	if root == nil {
		return
	}

	inorder(root.left)
	fmt.Print(root.data, " ")
	inorder(root.right)
}
// Time Complexity: O(N)
// Space Complexity: O(N)

func postorder(root *Node) {
	if root == nil {
		return
	}

	postorder(root.left)
	postorder(root.right)
	fmt.Print(root.data, " ")
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).


/*********************** ITERATIVE PREORDER *************************/

func preorderIterative(root *Node) []int {
	if root == nil {
		return []int{}
	}

	stack := []*Node{root}
	result := []int{}

	for len(stack) > 0 {

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, n.data)

		if n.right != nil {
			stack = append(stack, n.right)
		}
		if n.left != nil {
			stack = append(stack, n.left)
		}
	}

	return result
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: In the worst case, (a tree with every node having a single right child and left-subtree), the space complexity can be considered as O(N).


/*********************** ITERATIVE INORDER *************************/

func inorderIterative(root *Node) []int {

	var result []int
	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, curr.data)

		curr = curr.right
	}

	return result
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: In the worst case (a tree with just left children), the space complexity will be O(N).


/*********************** ITERATIVE POSTORDER (2 STACK) *************************/

func postorderIterative(root *Node) []int {

	if root == nil {
		return []int{}
	}

	s1 := []*Node{root}
	s2 := []*Node{}
	result := []int{}

	for len(s1) > 0 {

		curr := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		s2 = append(s2, curr)

		if curr.left != nil {
			s1 = append(s1, curr.left)
		}
		if curr.right != nil {
			s1 = append(s1, curr.right)
		}
	}

	for len(s2) > 0 {
		n := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		result = append(result, n.data)
	}

	return result
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N+N)

func postorderIterativeSingleStack(root *Node) []int {

	var result []int
	if root == nil {
		return result
	}

	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {

		// go left
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		} else {

			temp := stack[len(stack)-1].right

			// if right child doesn't exist
			if temp == nil {

				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = append(result, temp.data)

				// check if node is right child
				for len(stack) > 0 &&
					temp == stack[len(stack)-1].right {

					temp = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					result = append(result, temp.data)
				}

			} else {
				curr = temp
			}
		}
	}

	return result
}
// Time Complexity: O(N).
// Space Complexity: O(N)


/*********************** BFS (LEVEL ORDER) *************************/

func levelOrder(root *Node) {

	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.data, " ")

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*********************** ZIGZAG TRAVERSAL *************************/

func zigzagLevelOrder(root *Node) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*Node{root}
	leftToRight := true

	for len(queue) > 0 {

		size := len(queue)
		row := make([]int, size)

		for i := 0; i < size; i++ {

			node := queue[0]
			queue = queue[1:]

			index := i
			if !leftToRight {
				index = size - 1 - i
			}

			row[index] = node.data

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		leftToRight = !leftToRight
		result = append(result, row)
	}

	return result
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*********************** MAIN *************************/

func main() {

	root := createNode(1)
	root.left = createNode(2)
	root.right = createNode(3)
	root.left.left = createNode(4)
	root.left.right = createNode(5)

	fmt.Println("Preorder:")
	preorder(root)

	fmt.Println("\nInorder:")
	inorder(root)

	fmt.Println("\nPostorder:")
	postorder(root)

	fmt.Println("\nLevel Order:")
	levelOrder(root)

	fmt.Println("\nIterative Preorder:", preorderIterative(root))
	fmt.Println("Iterative Inorder:", inorderIterative(root))
	fmt.Println("Iterative Postorder:", postorderIterative(root))

	fmt.Println("\nZigzag Traversal:")
	ans := zigzagLevelOrder(root)

	for _, row := range ans {
		fmt.Println(row)
	}
}