// https://leetcode.com/problems/word-search/

#include <bits/stdc++.h>
using namespace std;

bool searchNext(vector<vector<char>> &board, string word, int row, int col,
                int index, int m, int n)
{

    // if index reaches at the end that means we have found the word
    if (index == word.length())
        return true;

    // Checking the boundaries if the character at which we are placed is not
    // the required character
    if (row < 0 || col < 0 || row == m || col == n || board[row][col] != word[index] or board[row][col] == '!')
        return false;

    // this is to prevent reusing of the same character
    char c = board[row][col];
    board[row][col] = '!';

    // top direction
    bool top = searchNext(board, word, row - 1, col, index + 1, m, n);
    // right direction
    bool right = searchNext(board, word, row, col + 1, index + 1, m, n);
    // bottom direction
    bool bottom = searchNext(board, word, row + 1, col, index + 1, m, n);
    // left direction
    bool left = searchNext(board, word, row, col - 1, index + 1, m, n);

    board[row][col] = c; // undo change

    return top || right || bottom || left;
}
bool exist(vector<vector<char>> board, string word)
{

    int m = board.size();
    int n = board[0].size();

    int index = 0;

    // First search the first character
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {

            if (board[i][j] == word[index])
            {
                if (searchNext(board, word, i, j, index, m, n))
                    return true;
            }
        }
    }

    return false;
}

int main()
{
    vector<vector<char>> board{{'A', 'B', 'C', 'E'},
                               {'S', 'F', 'C', 'S'},
                               {'A', 'D', 'E', 'E'}};

    string word = "ABCCED";

    bool res = exist(board, word);
    if (res == 1)
        cout << "True" << endl;
    else
        cout << "False" << endl;
}

//    Time Complexity:  O(m*n*4^k), where “K” is the length of the word. And we are searching for the letter m*n times in the worst case. Here 4 in 4^k is because at each level of our decision tree we are making 4 recursive calls which equal 4^k in the worst case.
// Space Complexity: O(K), Where k is the length of the given words.