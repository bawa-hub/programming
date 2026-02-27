package main

import "fmt"

/**************** NODE ****************/

type Node struct {
	data  int
	left  *Node
	right *Node
}

/**************** BST ****************/

type BST struct {
	root *Node
}

/**************** INSERT ****************/

func insertNode(t *Node, x int) *Node {

	if t == nil {
		return &Node{data: x}
	}

	if x < t.data {
		t.left = insertNode(t.left, x)
	} else if x > t.data {
		t.right = insertNode(t.right, x)
	}

	return t
}

func (bst *BST) Insert(x int) {
	bst.root = insertNode(bst.root, x)
}

/**************** INORDER ****************/

func inorder(t *Node) {
	if t == nil {
		return
	}

	inorder(t.left)
	fmt.Print(t.data, " ")
	inorder(t.right)
}

func (bst *BST) Display() {
	inorder(bst.root)
	fmt.Println()
}

/**************** SEARCH ****************/

func find(t *Node, x int) *Node {

	if t == nil {
		return nil
	}

	if x < t.data {
		return find(t.left, x)
	} else if x > t.data {
		return find(t.right, x)
	}

	return t
}

func (bst *BST) Search(x int) *Node {
	return find(bst.root, x)
}

/**************** FIND MIN ****************/

func findMin(t *Node) *Node {

	if t == nil {
		return nil
	}

	for t.left != nil {
		t = t.left
	}
	return t
}

/**************** FIND MAX ****************/

func findMax(t *Node) *Node {

	if t == nil {
		return nil
	}

	for t.right != nil {
		t = t.right
	}
	return t
}

/**************** REMOVE ****************/

func removeNode(t *Node, x int) *Node {

	if t == nil {
		return nil
	}

	if x < t.data {
		t.left = removeNode(t.left, x)
	} else if x > t.data {
		t.right = removeNode(t.right, x)
	} else {

		// node with two children
		if t.left != nil && t.right != nil {
			temp := findMin(t.right)
			t.data = temp.data
			t.right = removeNode(t.right, temp.data)
		} else {
			// one child or no child
			if t.left == nil {
				return t.right
			}
			return t.left
		}
	}

	return t
}

func (bst *BST) Remove(x int) {
	bst.root = removeNode(bst.root, x)
}

/**************** MAIN ****************/

func main() {

	tree := BST{}

	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(30)

	tree.Display()

	tree.Remove(20)
	tree.Display()

	tree.Remove(25)
	tree.Display()

	tree.Remove(30)
	tree.Display()
}