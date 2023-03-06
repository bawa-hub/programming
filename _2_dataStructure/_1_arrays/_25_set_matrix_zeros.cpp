// https://leetcode.com/problems/set-matrix-zeroes/

#include <bits/stdc++.h>

using namespace std;

// brute force
void setZeroes(vector<vector<int>> &matrix)
{
    int rows = matrix.size(), cols = matrix[0].size();
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < cols; j++)
        {
            if (matrix[i][j] == 0)
            {

                int ind = i - 1;
                while (ind >= 0)
                {
                    if (matrix[ind][j] != 0)
                    {
                        matrix[ind][j] = -1;
                    }
                    ind--;
                }
                ind = i + 1;
                while (ind < rows)
                {
                    if (matrix[ind][j] != 0)
                    {
                        matrix[ind][j] = -1;
                    }
                    ind++;
                }
                ind = j - 1;
                while (ind >= 0)
                {
                    if (matrix[i][ind] != 0)
                    {
                        matrix[i][ind] = -1;
                    }
                    ind--;
                }
                ind = j + 1;
                while (ind < cols)
                {
                    if (matrix[i][ind] != 0)
                    {
                        matrix[i][ind] = -1;
                    }
                    ind++;
                }
            }
        }
    }
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < cols; j++)
        {
            if (matrix[i][j] <= 0)
            {
                matrix[i][j] = 0;
            }
        }
    }
}
// Time Complexity:O((N*M)*(N + M)). O(N*M) for traversing through each element and (N+M)for traversing to row and column of elements having value 0.
// Space Complexity:O(1)

// better approach
void setZeroes(vector<vector<int>> &matrix)
{
    int rows = matrix.size(), cols = matrix[0].size();
    vector<int> dummy1(rows, -1), dummy2(cols, -1);
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < cols; j++)
        {
            if (matrix[i][j] == 0)
            {
                dummy1[i] = 0;
                dummy2[j] = 0;
            }
        }
    }
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < cols; j++)
        {
            if (dummy1[i] == 0 || dummy2[j] == 0)
            {
                matrix[i][j] = 0;
            }
        }
    }
}
// Time Complexity: O(N*M + N*M)
// Space Complexity: O(N)

// optimized better approach
void setZeroes(vector<vector<int>> &matrix)
{
    int col0 = 1, rows = matrix.size(), cols = matrix[0].size();
    for (int i = 0; i < rows; i++)
    {
        // checking if 0 is present in the 0th column or not
        if (matrix[i][0] == 0)
            col0 = 0;
        for (int j = 1; j < cols; j++)
        {
            if (matrix[i][j] == 0)
            {
                matrix[i][0] = 0;
                matrix[0][j] = 0;
            }
        }
    }
    // traversing in the reverse direction and
    // checking if the row or col has 0 or not
    // and setting values of matrix accordingly.
    for (int i = rows - 1; i >= 0; i--)
    {
        for (int j = cols - 1; j >= 1; j--)
        {
            if (matrix[i][0] == 0 || matrix[0][j] == 0)
            {
                matrix[i][j] = 0;
            }
        }
        if (col0 == 0)
        {
            matrix[i][0] = 0;
        }
    }
}
// Time Complexity: O(2*(N*M)), as we are traversing two times in a matrix,
// Space Complexity: O(1)

int main()
{
    vector<vector<int>> arr;
    arr = {{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}};
    setZeroes(arr);
    cout << "The Final Matrix is " << endl;
    for (int i = 0; i < arr.size(); i++)
    {
        for (int j = 0; j < arr[0].size(); j++)
        {
            cout << arr[i][j] << " ";
        }
        cout << "\n";
    }
}