// https://leetcode.com/problems/surrounded-regions/

#include <bits/stdc++.h>
using namespace std;

class Solution
{
    int rows, cols;
    void dfs(vector<vector<char>> &board, int i, int j)
    {
        if (i < 0 || j < 0 || i >= rows || j >= cols)
            return;

        if (board[i][j] != 'O')
            return;

        board[i][j] = '#';
        dfs(board, i + 1, j);
        dfs(board, i - 1, j);
        dfs(board, i, j + 1);
        dfs(board, i, j - 1);
    }

public:
    void solve(vector<vector<char>> &board)
    {
        rows = board.size();
        cols = board[0].size();
        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                if (i == 0 || j == 0 || i == rows - 1 || j == cols - 1)
                    dfs(board, i, j);
            }
        }

        for (int i = 0; i < rows; i++)
        {
            for (int j = 0; j < cols; j++)
            {
                if (board[i][j] == '#')
                    board[i][j] = 'O';
                else if (board[i][j] == 'O')
                    board[i][j] = 'X';
            }
        }
    }
};