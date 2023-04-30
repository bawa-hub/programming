// https://leetcode.com/problems/spiral-matrix/

#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
    vector<int> spiralOrder(vector<vector<int>> &matrix)
    {
        int R = matrix.size();
        int C = matrix[0].size();
        vector<int> ans;

        int top = 0, left = 0, bottom = R - 1, right = C - 1;

        while (top <= bottom && left <= right)
        {
            for (int i = left; i <= right; i++)
                ans.push_back(matrix[top][i]);

            top++;

            for (int i = top; i <= bottom; i++)
                ans.push_back(matrix[i][right]);

            right--;

            if (top <= bottom)
            {
                for (int i = right; i >= left; i--)
                    ans.push_back(matrix[bottom][i]);

                bottom--;
            }

            if (left <= right)
            {
                for (int i = bottom; i >= top; i--)
                    ans.push_back(matrix[i][left]);

                left++;
            }
        }
        return ans;
    }
};
// Time Complexity: O(R x C)
// Reason: We are printing every element of the matrix so the time complexity is O(R x C) where R and C are rows and columns of the matrix.
// Space Complexity: O(1)
