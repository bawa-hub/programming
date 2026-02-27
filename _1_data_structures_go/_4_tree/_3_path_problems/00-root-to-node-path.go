// https://www.interviewbit.com/problems/path-to-given-node/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func getPath(root *Node, path *[]int, x int) bool {

	if root == nil {
		return false
	}

	// add current node
	*path = append(*path, root.data)

	// node found
	if root.data == x {
		return true
	}

	// search left or right
	if getPath(root.left, path, x) ||
		getPath(root.right, path, x) {
		return true
	}

	// backtrack
	*path = (*path)[:len(*path)-1]

	return false
}

func main() {

	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	path := []int{}

	if getPath(root, &path, 5) {
		fmt.Println("Path:", path)
	} else {
		fmt.Println("Node not found")
	}
}