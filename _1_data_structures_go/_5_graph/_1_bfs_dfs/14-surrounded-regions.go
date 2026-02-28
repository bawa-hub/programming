// https://leetcode.com/problems/surrounded-regions/

type Solution struct {
	rows int
	cols int
}

func (s *Solution) dfs(board [][]byte, i, j int) {

	if i < 0 || j < 0 || i >= s.rows || j >= s.cols {
		return
	}

	if board[i][j] != 'O' {
		return
	}

	// mark safe region
	board[i][j] = '#'

	s.dfs(board, i+1, j)
	s.dfs(board, i-1, j)
	s.dfs(board, i, j+1)
	s.dfs(board, i, j-1)
}

func (s *Solution) solve(board [][]byte) {

	if len(board) == 0 {
		return
	}

	s.rows = len(board)
	s.cols = len(board[0])

	// DFS from boundary cells
	for i := 0; i < s.rows; i++ {
		for j := 0; j < s.cols; j++ {
			if i == 0 || j == 0 || i == s.rows-1 || j == s.cols-1 {
				s.dfs(board, i, j)
			}
		}
	}

	// convert regions
	for i := 0; i < s.rows; i++ {
		for j := 0; j < s.cols; j++ {

			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}