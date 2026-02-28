package main

import "fmt"

/**************** NODE ****************/

type QueueNode struct {
	val  int
	next *QueueNode
}

/**************** QUEUE ****************/

type Queue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

/**************** EMPTY ****************/

func (q *Queue) Empty() bool {
	return q.front == nil
}

/**************** PEEK ****************/

func (q *Queue) Peek() int {

	if q.Empty() {
		fmt.Println("Queue is Empty")
		return -1
	}

	return q.front.val
}

/**************** ENQUEUE ****************/

func (q *Queue) Enqueue(value int) {

	temp := &QueueNode{val: value}

	if q.front == nil {
		q.front = temp
		q.rear = temp
	} else {
		q.rear.next = temp
		q.rear = temp
	}

	fmt.Println(value, "Inserted into Queue")
	q.size++
}

/**************** DEQUEUE ****************/

func (q *Queue) Dequeue() {

	if q.front == nil {
		fmt.Println("Queue is Empty")
		return
	}

	fmt.Println(q.front.val, "Removed From Queue")

	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
}

/**************** MAIN ****************/

func main() {

	Q := Queue{}

	Q.Enqueue(10)
	Q.Enqueue(20)
	Q.Enqueue(30)
	Q.Enqueue(40)
	Q.Enqueue(50)

	Q.Dequeue()

	fmt.Println("Queue size:", Q.size)
	fmt.Println("Peek element:", Q.Peek())
}