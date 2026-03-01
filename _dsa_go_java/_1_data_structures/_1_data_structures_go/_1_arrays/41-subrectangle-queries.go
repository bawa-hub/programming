// https://leetcode.com/problems/subrectangle-queries/

package main

import "fmt"

type SubrectangleQueries struct {
	rectangle [][]int
}

// Constructor
func NewSubrectangleQueries(rectangle [][]int) *SubrectangleQueries {
	return &SubrectangleQueries{
		rectangle: rectangle,
	}
}

// updateSubrectangle updates all values in given sub-rectangle
func (s *SubrectangleQueries) updateSubrectangle(row1, col1, row2, col2, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			s.rectangle[i][j] = newValue
		}
	}
}

// getValue returns value at given position
func (s *SubrectangleQueries) getValue(row, col int) int {
	return s.rectangle[row][col]
}

func main() {
	rectangle := [][]int{
		{1, 2, 1},
		{4, 3, 4},
		{3, 2, 1},
		{1, 1, 1},
	}

	obj := NewSubrectangleQueries(rectangle)

	fmt.Println(obj.getValue(0, 2)) // 1

	obj.updateSubrectangle(0, 0, 3, 2, 5)

	fmt.Println(obj.getValue(0, 2)) // 5
	fmt.Println(obj.getValue(3, 1)) // 5
}