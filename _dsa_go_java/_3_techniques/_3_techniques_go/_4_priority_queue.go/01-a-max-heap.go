package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	A []int
}

// ===== Helper functions =====

func (pq *PriorityQueue) parent(i int) int {
	return (i - 1) / 2
}

func (pq *PriorityQueue) left(i int) int {
	return 2*i + 1
}

func (pq *PriorityQueue) right(i int) int {
	return 2*i + 2
}

// ===== Heapify Down =====

func (pq *PriorityQueue) heapifyDown(i int) {
	left := pq.left(i)
	right := pq.right(i)

	largest := i

	if left < pq.size() && pq.A[left] > pq.A[largest] {
		largest = left
	}

	if right < pq.size() && pq.A[right] > pq.A[largest] {
		largest = right
	}

	if largest != i {
		pq.A[i], pq.A[largest] = pq.A[largest], pq.A[i]
		pq.heapifyDown(largest)
	}
}

// ===== Heapify Up =====

func (pq *PriorityQueue) heapifyUp(i int) {
	if i > 0 && pq.A[pq.parent(i)] < pq.A[i] {
		pq.A[i], pq.A[pq.parent(i)] = pq.A[pq.parent(i)], pq.A[i]
		pq.heapifyUp(pq.parent(i))
	}
}

// ===== Public Methods =====

func (pq *PriorityQueue) size() int {
	return len(pq.A)
}

func (pq *PriorityQueue) empty() bool {
	return pq.size() == 0
}

func (pq *PriorityQueue) push(key int) {
	pq.A = append(pq.A, key)
	pq.heapifyUp(pq.size() - 1)
}

func (pq *PriorityQueue) pop() error {
	if pq.empty() {
		return errors.New("heap underflow")
	}

	pq.A[0] = pq.A[pq.size()-1]
	pq.A = pq.A[:pq.size()-1]

	pq.heapifyDown(0)
	return nil
}

func (pq *PriorityQueue) top() (int, error) {
	if pq.empty() {
		return 0, errors.New("heap is empty")
	}
	return pq.A[0], nil
}

// ===== Main =====

func main() {
	pq := &PriorityQueue{}

	pq.push(3)
	pq.push(2)
	pq.push(15)

	fmt.Println("Size is", pq.size())

	val, _ := pq.top()
	fmt.Print(val, " ")
	pq.pop()

	val, _ = pq.top()
	fmt.Print(val, " ")
	pq.pop()

	pq.push(5)
	pq.push(4)
	pq.push(45)

	fmt.Println("\nSize is", pq.size())

	val, _ = pq.top()
	fmt.Print(val, " ")
	pq.pop()

	val, _ = pq.top()
	fmt.Print(val, " ")
	pq.pop()

	val, _ = pq.top()
	fmt.Print(val, " ")
	pq.pop()

	val, err := pq.top()
	if err != nil {
		fmt.Println("\nHeap is empty")
	} else {
		fmt.Println(val)
	}
}

// | Operation | Time     |
// | --------- | -------- |
// | push      | O(log n) |
// | pop       | O(log n) |
// | top       | O(1)     |
