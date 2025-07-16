// https://leetcode.com/problems/climbing-stairs/

#include <bits/stdc++.h>
using namespace std;

// recursive
class Solution
{
public:
    int climbStairs(int n)
    {
        vector<int> dp(n + 1, -1);
        return f(n, dp);
    }

    int f(int n, vector<int> &dp)
    {
        if (n == 0 || n == 1)
            return 1;

        if (dp[n] != -1)
            return dp[n];

        return dp[n] = f(n - 1, dp) + f(n - 2, dp);
    }
};
// TC: O(N)
// SC: O(N)

// iterative
int main()
{

    int n = 3;
    vector<int> dp(n + 1, -1);

    dp[0] = 1;
    dp[1] = 1;

    for (int i = 2; i <= n; i++)
    {
        dp[i] = dp[i - 1] + dp[i - 2];
    }
    cout << dp[n];
    return 0;
}
// Time Complexity: O(N)
// Reason: We are running a simple iterative loop
// Space Complexity: O(N)
// Reason: We are using an external array of size ‘n+1’.

// space optimized
int main()
{

    int n = 3;

    int prev2 = 1;
    int prev = 1;

    for (int i = 2; i <= n; i++)
    {
        int cur_i = prev2 + prev;
        prev2 = prev;
        prev = cur_i;
    }
    cout << prev;
    return 0;
}

// Time Complexity: O(N)
// Reason: We are running a simple iterative loop
// Space Complexity: O(1)
// Reason: We are not using any extra space.