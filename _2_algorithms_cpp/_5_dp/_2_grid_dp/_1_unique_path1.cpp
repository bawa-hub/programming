// https://leetcode.com/problems/unique-paths/

#include <bits/stdc++.h>
using namespace std;

// recursive
int countWaysUtil(int i, int j, vector<vector<int>> &dp)
{
    if (i == 0 && j == 0)
        return 1;
    if (i < 0 || j < 0)
        return 0;
    if (dp[i][j] != -1)
        return dp[i][j];

    int up = countWaysUtil(i - 1, j, dp);
    int left = countWaysUtil(i, j - 1, dp);

    return dp[i][j] = up + left;
}

int countWays(int m, int n)
{
    vector<vector<int>> dp(m, vector<int>(n, -1));
    return countWaysUtil(m - 1, n - 1, dp);
}
// Time Complexity: O(M*N)
// Reason: At max, there will be M*N calls of recursion.
// Space Complexity: O((N-1)+(M-1)) + O(M*N)
// Reason: We are using a recursion stack space:O((N-1)+(M-1)), here (N-1)+(M-1) is the path length and an external DP Array of size ‘M*N’.

// tabulation
int countWaysUtil(int m, int n, vector<vector<int>> &dp)
{
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {

            // base condition
            if (i == 0 && j == 0)
            {
                dp[i][j] = 1;
                continue;
            }

            int up = 0;
            int left = 0;

            if (i > 0)
                up = dp[i - 1][j];
            if (j > 0)
                left = dp[i][j - 1];

            dp[i][j] = up + left;
        }
    }

    return dp[m - 1][n - 1];
}

int countWays(int m, int n)
{
    vector<vector<int>> dp(m, vector<int>(n, -1));
    return countWaysUtil(m, n, dp);
}
// Time Complexity: O(M*N)
// Reason: There are two nested loops
// Space Complexity: O(M*N)
// Reason: We are using an external array of size ‘M*N’’.

// space optimiztion
int countWays(int m, int n)
{
    vector<int> prev(n, 0);
    for (int i = 0; i < m; i++)
    {
        vector<int> temp(n, 0);
        for (int j = 0; j < n; j++)
        {
            if (i == 0 && j == 0)
            {
                temp[j] = 1;
                continue;
            }

            int up = 0;
            int left = 0;

            if (i > 0)
                up = prev[j];
            if (j > 0)
                left = temp[j - 1];

            temp[j] = up + left;
        }
        prev = temp;
    }

    return prev[n - 1];
}
// Time Complexity: O(M*N)
// Reason: There are two nested loops
// Space Complexity: O(N)
// Reason: We are using an external array of size ‘N’ to store only one row.

int main()
{

    int m = 3;
    int n = 2;

    cout << countWays(m, n);
}