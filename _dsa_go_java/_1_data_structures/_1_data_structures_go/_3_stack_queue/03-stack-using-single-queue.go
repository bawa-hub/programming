package main

import "fmt"

type Stack struct {
	q []int
}

/**************** PUSH ****************/

func (s *Stack) Push(x int) {

	size := len(s.q)

	// push element
	s.q = append(s.q, x)

	// rotate previous elements
	for i := 0; i < size; i++ {
		front := s.q[0]
		s.q = s.q[1:]
		s.q = append(s.q, front)
	}
}

/**************** POP ****************/

func (s *Stack) Pop() int {

	if len(s.q) == 0 {
		return -1
	}

	top := s.q[0]
	s.q = s.q[1:]
	return top
}

/**************** TOP ****************/

func (s *Stack) Top() int {
	if len(s.q) == 0 {
		return -1
	}
	return s.q[0]
}

/**************** EMPTY ****************/

func (s *Stack) Empty() bool {
	return len(s.q) == 0
}

/**************** SIZE ****************/

func (s *Stack) Size() int {
	return len(s.q)
}

/**************** MAIN ****************/

func main() {

	s := Stack{}

	s.Push(3)
	s.Push(2)
	s.Push(4)
	s.Push(1)

	fmt.Println("Top of stack:", s.Top())
	fmt.Println("Size before pop:", s.Size())

	fmt.Println("Deleted element:", s.Pop())

	fmt.Println("Top after pop:", s.Top())
	fmt.Println("Size after pop:", s.Size())
}