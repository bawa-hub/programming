package main

import (
	"fmt"
	"os"
)

type Queue struct {
	arr      []int
	start    int
	end      int
	currSize int
	maxSize  int
}

/**************** CONSTRUCTOR ****************/

func NewQueue(maxSize int) *Queue {
	return &Queue{
		arr:      make([]int, maxSize),
		start:    -1,
		end:      -1,
		currSize: 0,
		maxSize:  maxSize,
	}
}

/**************** PUSH ****************/

func (q *Queue) Push(newElement int) {

	if q.currSize == q.maxSize {
		fmt.Println("Queue is full\nExiting...")
		os.Exit(1)
	}

	if q.end == -1 {
		q.start = 0
		q.end = 0
	} else {
		q.end = (q.end + 1) % q.maxSize
	}

	q.arr[q.end] = newElement
	q.currSize++

	fmt.Println("The element pushed is", newElement)
}

/**************** POP ****************/

func (q *Queue) Pop() int {

	if q.start == -1 {
		fmt.Println("Queue Empty")
		os.Exit(1)
	}

	popped := q.arr[q.start]

	if q.currSize == 1 {
		q.start = -1
		q.end = -1
	} else {
		q.start = (q.start + 1) % q.maxSize
	}

	q.currSize--
	return popped
}

/**************** TOP ****************/

func (q *Queue) Top() int {

	if q.start == -1 {
		fmt.Println("Queue is Empty")
		os.Exit(1)
	}

	return q.arr[q.start]
}

/**************** SIZE ****************/

func (q *Queue) Size() int {
	return q.currSize
}

/**************** MAIN ****************/

func main() {

	q := NewQueue(6)

	q.Push(4)
	q.Push(14)
	q.Push(24)
	q.Push(34)

	fmt.Println("Peek before deletion:", q.Top())
	fmt.Println("Size before deletion:", q.Size())

	fmt.Println("Deleted element:", q.Pop())

	fmt.Println("Peek after deletion:", q.Top())
	fmt.Println("Size after deletion:", q.Size())
}