// https://leetcode.com/problems/n-queens-ii/

// same as n_queens, only diff is it return count
class Solution
{
public:
    bool isSafe1(int row, int col, vector<string> board, int n)
    {
        // check upper element
        int duprow = row;
        int dupcol = col;

        while (row >= 0 && col >= 0)
        {
            if (board[row][col] == 'Q')
                return false;
            row--;
            col--;
        }

        col = dupcol;
        row = duprow;
        while (col >= 0)
        {
            if (board[row][col] == 'Q')
                return false;
            col--;
        }

        row = duprow;
        col = dupcol;
        while (row < n && col >= 0)
        {
            if (board[row][col] == 'Q')
                return false;
            row++;
            col--;
        }
        return true;
    }

public:
    void solve(int col, vector<string> &board, int &ans, int n)
    {
        if (col == n)
        {
            ans++;
            return;
        }
        for (int row = 0; row < n; row++)
        {
            if (isSafe1(row, col, board, n))
            {
                board[row][col] = 'Q';
                solve(col + 1, board, ans, n);
                board[row][col] = '.';
            }
        }
    }

public:
    int totalNQueens(int n)
    {
        int ans = 0;
        vector<string> board(n);
        string s(n, '.');
        for (int i = 0; i < n; i++)
        {
            board[i] = s;
        }
        solve(0, board, ans, n);
        return ans;
    }
};