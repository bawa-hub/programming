// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type Codec struct{}

/**************** SERIALIZE ****************/

func (c *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}

	var result []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "#")
			continue
		}

		result = append(result, strconv.Itoa(node.val))
		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}

	return strings.Join(result, ",")
}

/**************** DESERIALIZE ****************/

func (c *Codec) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	values := strings.Split(data, ",")

	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{val: val}

	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		// left child
		if values[i] != "#" {
			v, _ := strconv.Atoi(values[i])
			node.left = &TreeNode{val: v}
			queue = append(queue, node.left)
		}
		i++

		// right child
		if values[i] != "#" {
			v, _ := strconv.Atoi(values[i])
			node.right = &TreeNode{val: v}
			queue = append(queue, node.right)
		}
		i++
	}

	return root
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/**************** TEST ****************/

func main() {

	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
		},
		right: &TreeNode{
			val: 3,
			left:  &TreeNode{val: 4},
			right: &TreeNode{val: 5},
		},
	}

	codec := Codec{}

	data := codec.serialize(root)
	fmt.Println("Serialized:", data)

	newRoot := codec.deserialize(data)
	fmt.Println("Deserialize Root:", newRoot.val)
}