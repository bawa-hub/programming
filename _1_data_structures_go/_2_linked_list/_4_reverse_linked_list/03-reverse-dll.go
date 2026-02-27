package main

import "fmt"

type Node struct {
	data int
	next *Node
	back *Node
}

func convertArr2DLL(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	head := &Node{data: arr[0]}
	prev := head

	for i := 1; i < len(arr); i++ {
		temp := &Node{
			data: arr[i],
			back: prev,
		}
		prev.next = temp
		prev = temp
	}

	return head
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func reverseDLLStack(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	stack := []int{}
	temp := head

	for temp != nil {
		stack = append(stack, temp.data)
		temp = temp.next
	}

	temp = head

	for temp != nil {
		n := len(stack)
		temp.data = stack[n-1]
		stack = stack[:n-1]
		temp = temp.next
	}

	return head
}
// Time  : O(N)
// Space : O(N)

// in place swap
func reverseDLL(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	var prev *Node
	current := head

	for current != nil {

		prev = current.back

		// swap links
		current.back = current.next
		current.next = prev

		current = current.back
	}

	return prev.back
}
// Time  : O(N)
// Space : O(1)


func main() {

	arr := []int{12, 5, 8, 7, 4}

	head := convertArr2DLL(arr)

	fmt.Println("DLL Initially:")
	printList(head)

	head = reverseDLL(head)

	fmt.Println("After Reversing:")
	printList(head)
}