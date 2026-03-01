package main

import "fmt"

/**************** NODE ****************/

type StackNode struct {
	data int
	next *StackNode
}

/**************** STACK ****************/

type Stack struct {
	top  *StackNode
	size int
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

/**************** PUSH ****************/

func (s *Stack) StackPush(x int) {

	element := &StackNode{data: x}
	element.next = s.top
	s.top = element

	fmt.Println("Element pushed")
	s.size++
}

/**************** POP ****************/

func (s *Stack) StackPop() int {

	if s.top == nil {
		return -1
	}

	topData := s.top.data
	s.top = s.top.next
	s.size--

	return topData
}

/**************** SIZE ****************/

func (s *Stack) StackSize() int {
	return s.size
}

/**************** EMPTY ****************/

func (s *Stack) StackIsEmpty() bool {
	return s.top == nil
}

/**************** PEEK ****************/

func (s *Stack) StackPeek() int {

	if s.top == nil {
		return -1
	}
	return s.top.data
}

/**************** PRINT ****************/

func (s *Stack) PrintStack() {

	current := s.top
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

/**************** MAIN ****************/

func main() {

	s := NewStack()

	s.StackPush(10)
	s.StackPush(20)
	s.StackPush(30)

	fmt.Println("Element popped:", s.StackPop())
	fmt.Println("Stack size:", s.StackSize())
	fmt.Println("Stack empty?", s.StackIsEmpty())
	fmt.Println("Top element:", s.StackPeek())

	fmt.Print("Stack elements: ")
	s.PrintStack()
}