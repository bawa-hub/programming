// https://practice.geeksforgeeks.org/problems/find-length-of-loop/1
// https://www.geeksforgeeks.org/find-length-of-loop-in-linked-list/

type Node struct {
	data int
	next *Node
}

// using hashing
func countNodesInLoopHash(head *Node) int {

	visited := make(map[*Node]int)

	temp := head
	timer := 1

	for temp != nil {

		if val, ok := visited[temp]; ok {
			return timer - val
		}

		visited[temp] = timer
		temp = temp.next
		timer++
	}

	return 0
}
// TC: O(n*2*x)
// SC: O(n)

// floyd cycled algo - slow fast pointer
func countNodesInLoop(head *Node) int {

	if head == nil {
		return 0
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return findLength(slow)
		}
	}

	return 0
}
func findLength(meet *Node) int {

	count := 1
	temp := meet.next

	for temp != meet {
		count++
		temp = temp.next
	}

	return count
}
// Time Complexity: O(N + lenOfLoop)
// Reason: In the worst case, all the nodes of the list are visited.
// Space Complexity: O(1)
// Reason: No extra data structure is used.