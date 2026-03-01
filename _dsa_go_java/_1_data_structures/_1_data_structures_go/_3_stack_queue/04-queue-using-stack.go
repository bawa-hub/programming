// https://leetcode.com/problems/implement-queue-using-stacks/description/

package main

import "fmt"

type MyQueue struct {
	input  []int
	output []int
}

/**************** CONSTRUCTOR ****************/

func Constructor() MyQueue {
	return MyQueue{}
}

/**************** PUSH ****************/

func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

/**************** POP ****************/

func (q *MyQueue) Pop() int {

	if len(q.output) == 0 {
		for len(q.input) > 0 {
			n := q.input[len(q.input)-1]
			q.input = q.input[:len(q.input)-1]
			q.output = append(q.output, n)
		}
	}

	x := q.output[len(q.output)-1]
	q.output = q.output[:len(q.output)-1]

	return x
}

/**************** PEEK ****************/

func (q *MyQueue) Peek() int {

	if len(q.output) == 0 {
		for len(q.input) > 0 {
			n := q.input[len(q.input)-1]
			q.input = q.input[:len(q.input)-1]
			q.output = append(q.output, n)
		}
	}

	return q.output[len(q.output)-1]
}

/**************** EMPTY ****************/

func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

/**************** MAIN ****************/

func main() {

	q := Constructor()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println("Peek:", q.Peek())
	fmt.Println("Pop:", q.Pop())
	fmt.Println("Peek:", q.Peek())
	fmt.Println("Empty:", q.Empty())
}
