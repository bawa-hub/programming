// https://leetcode.com/problems/flood-fill/

package main

import "fmt"

type Solution struct{}

func (s *Solution) dfs(
	row, col int,
	ans [][]int,
	image [][]int,
	newColor int,
	delRow []int,
	delCol []int,
	iniColor int,
) {

	ans[row][col] = newColor

	n := len(image)
	m := len(image[0])

	for i := 0; i < 4; i++ {
		nrow := row + delRow[i]
		ncol := col + delCol[i]

		if nrow >= 0 && nrow < n &&
			ncol >= 0 && ncol < m &&
			image[nrow][ncol] == iniColor &&
			ans[nrow][ncol] != newColor {

			s.dfs(nrow, ncol, ans, image, newColor, delRow, delCol, iniColor)
		}
	}
}

func (s *Solution) floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	iniColor := image[sr][sc]
	ans := make([][]int, len(image))

	for i := range image {
		ans[i] = make([]int, len(image[0]))
		copy(ans[i], image[i])
	}

	delRow := []int{-1, 0, 1, 0}
	delCol := []int{0, 1, 0, -1}

	s.dfs(sr, sc, ans, image, newColor, delRow, delCol, iniColor)

	return ans
}

func main() {

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	obj := Solution{}
	ans := obj.floodFill(image, 1, 1, 2)

	for _, row := range ans {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}