// https://practice.geeksforgeeks.org/problems/top-view-of-binary-tree/1
// https://takeuforward.org/data-structure/top-view-of-a-binary-tree/


type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Pair struct {
	Node *Node
	Line int
}

func topView(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	mpp := make(map[int]int)
	queue := []Pair{{root, 0}}

	minLine, maxLine := 0, 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		node := p.Node
		line := p.Line

		// store only first node at this vertical line
		if _, exists := mpp[line]; !exists {
			mpp[line] = node.Data
		}

		if line < minLine {
			minLine = line
		}
		if line > maxLine {
			maxLine = line
		}

		if node.Left != nil {
			queue = append(queue, Pair{node.Left, line - 1})
		}

		if node.Right != nil {
			queue = append(queue, Pair{node.Right, line + 1})
		}
	}

	for i := minLine; i <= maxLine; i++ {
		ans = append(ans, mpp[i])
	}

	return ans
}
// Time Complexity: O(N)
// Space Complexity: O(N)