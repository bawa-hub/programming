// https://leetcode.com/problems/set-matrix-zeroes/
// https://takeuforward.org/data-structure/set-matrix-zero/

package main

import "fmt"

func setZeroesBrute(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	// First pass: mark affected cells as -1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if matrix[i][j] == 0 {

				// Up
				ind := i - 1
				for ind >= 0 {
					if matrix[ind][j] != 0 {
						matrix[ind][j] = -1
					}
					ind--
				}

				// Down
				ind = i + 1
				for ind < rows {
					if matrix[ind][j] != 0 {
						matrix[ind][j] = -1
					}
					ind++
				}

				// Left
				ind = j - 1
				for ind >= 0 {
					if matrix[i][ind] != 0 {
						matrix[i][ind] = -1
					}
					ind--
				}

				// Right
				ind = j + 1
				for ind < cols {
					if matrix[i][ind] != 0 {
						matrix[i][ind] = -1
					}
					ind++
				}
			}
		}
	}

	// Second pass: convert all <= 0 to 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] <= 0 {
				matrix[i][j] = 0
			}
		}
	}
}
// Time Complexity:O((N*M)*(N + M)). O(N*M) for traversing through each element and (N+M)for traversing to row and column of elements having value 0.
// Space Complexity:O(1)

func setZeroesBetter(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	dummyRow := make([]int, rows)
	dummyCol := make([]int, cols)

	// First pass: mark rows & columns
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dummyRow[i] = 1
				dummyCol[j] = 1
			}
		}
	}

	// Second pass: update matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dummyRow[i] == 1 || dummyCol[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
// Time Complexity: O(N*M + N*M)
// Space Complexity: O(N)

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	col0 := 1

	// Step 1: Mark rows & columns
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {

				// mark row
				matrix[i][0] = 0

				// mark column
				if j != 0 {
					matrix[0][j] = 0
				} else {
					col0 = 0
				}
			}
		}
	}

	// Step 2: Update inner matrix
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Step 3: Handle first row
	if matrix[0][0] == 0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	// Step 4: Handle first column
	if col0 == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
}
// Time Complexity: O(2*(N*M)), as we are traversing two times in a matrix,
// Space Complexity: O(1)

func main() {
	arr := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(arr)

	fmt.Println("The Final Matrix is:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}