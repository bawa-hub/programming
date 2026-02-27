// https://leetcode.com/problems/min-stack/

package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int64
	mini  int64
}

/************ CONSTRUCTOR ************/

func Constructor() MinStack {
	return MinStack{
		stack: []int64{},
		mini:  math.MaxInt64,
	}
}

/************ PUSH ************/

func (m *MinStack) Push(val int) {

	value := int64(val)

	if len(m.stack) == 0 {
		m.mini = value
		m.stack = append(m.stack, value)
		return
	}

	if value < m.mini {
		encoded := 2*value - m.mini
		m.stack = append(m.stack, encoded)
		m.mini = value
	} else {
		m.stack = append(m.stack, value)
	}
}

/************ POP ************/

func (m *MinStack) Pop() {

	if len(m.stack) == 0 {
		return
	}

	top := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]

	if top < m.mini {
		m.mini = 2*m.mini - top
	}
}

/************ TOP ************/

func (m *MinStack) Top() int {

	if len(m.stack) == 0 {
		return -1
	}

	top := m.stack[len(m.stack)-1]

	if top < m.mini {
		return int(m.mini)
	}

	return int(top)
}

/************ GET MIN ************/

func (m *MinStack) GetMin() int {
	return int(m.mini)
}

// Time Complexity: O(1)
// Space Complexity: O(N)

/************ TEST ************/

func main() {

	ms := Constructor()

	ms.Push(5)
	ms.Push(3)
	ms.Push(7)
	ms.Push(2)

	fmt.Println("Min:", ms.GetMin()) // 2
	fmt.Println("Top:", ms.Top())    // 2

	ms.Pop()

	fmt.Println("Min:", ms.GetMin()) // 3
}