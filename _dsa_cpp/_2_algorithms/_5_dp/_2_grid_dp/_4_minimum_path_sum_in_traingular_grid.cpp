// https://leetcode.com/problems/triangle/

#include <bits/stdc++.h>

using namespace std;

// recursive
int minimumPathSumUtil(int i, int j, vector<vector<int>> &triangle, int n,
                       vector<vector<int>> &dp)
{

    if (dp[i][j] != -1)
        return dp[i][j];

    if (i == n - 1)
        return triangle[i][j];

    int down = triangle[i][j] + minimumPathSumUtil(i + 1, j, triangle, n, dp);
    int diagonal = triangle[i][j] + minimumPathSumUtil(i + 1, j + 1, triangle, n, dp);

    return dp[i][j] = min(down, diagonal);
}

int minimumPathSum(vector<vector<int>> &triangle, int n)
{
    vector<vector<int>> dp(n, vector<int>(n, -1));
    return minimumPathSumUtil(0, 0, triangle, n, dp);
}
// Time Complexity: O(N*N)
// Reason: At max, there will be (half of, due to triangle) N*N calls of recursion.
// Space Complexity: O(N) + O(N*N)
// Reason: We are using a recursion stack space: O((N), where N is the path length and an external DP Array of size ‘N*N’.

// iterative
int minimumPathSum(vector<vector<int>> &triangle, int n)
{
    vector<vector<int>> dp(n, vector<int>(n, 0));

    for (int j = 0; j < n; j++)
    {
        dp[n - 1][j] = triangle[n - 1][j];
    }

    for (int i = n - 2; i >= 0; i--)
    {
        for (int j = i; j >= 0; j--)
        {

            int down = triangle[i][j] + dp[i + 1][j];
            int diagonal = triangle[i][j] + dp[i + 1][j + 1];

            dp[i][j] = min(down, diagonal);
        }
    }

    return dp[0][0];
}
// Time Complexity: O(N*N)
// Reason: There are two nested loops
// Space Complexity: O(N*N)
// Reason: We are using an external array of size ‘N*N’. The stack space will be eliminated

// space optimized
int minimumPathSum(vector<vector<int>> &triangle, int n)
{
    vector<int> front(n, 0), cur(n, 0);

    for (int j = 0; j < n; j++)
    {
        front[j] = triangle[n - 1][j];
    }

    for (int i = n - 2; i >= 0; i--)
    {
        for (int j = i; j >= 0; j--)
        {

            int down = triangle[i][j] + front[j];
            int diagonal = triangle[i][j] + front[j + 1];

            cur[j] = min(down, diagonal);
        }
        front = cur;
    }

    return front[0];
}
// Time Complexity: O(N*N)
// Reason: There are two nested loops
// Space Complexity: O(N)
// Reason: We are using an external array of size ‘N’ to store only one row.

int main()
{

    vector<vector<int>> triangle{{1},
                                 {2, 3},
                                 {3, 6, 7},
                                 {8, 9, 6, 10}};

    int n = triangle.size();

    cout << minimumPathSum(triangle, n);
}