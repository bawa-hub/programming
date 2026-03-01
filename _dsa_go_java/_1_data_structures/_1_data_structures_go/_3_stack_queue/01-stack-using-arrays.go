package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(x int) {
	s.arr = append(s.arr, x)
}

func (s *Stack) Pop() int {
	n := len(s.arr)
	x := s.arr[n-1]
	s.arr = s.arr[:n-1]
	return x
}

func (s *Stack) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func main() {

	s := Stack{}

	s.Push(6)
	s.Push(3)
	s.Push(7)

	fmt.Println("Top:", s.Top())
	fmt.Println("Size:", s.Size())

	fmt.Println("Deleted:", s.Pop())

	fmt.Println("Size:", s.Size())
	fmt.Println("Top:", s.Top())
}