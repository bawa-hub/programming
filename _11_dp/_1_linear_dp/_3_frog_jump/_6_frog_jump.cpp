// https://atcoder.jp/contests/dp/tasks/dp_a
// https://practice.geeksforgeeks.org/problems/geek-jump/1

#include <bits/stdc++.h>
using namespace std;

// recursive/memoization
int solve(int ind, vector<int> &height, vector<int> &dp)
{
    if (ind == 0)
        return 0;
    if (dp[ind] != -1)
        return dp[ind];
    
    int jumpOne = solve(ind - 1, height, dp) + abs(height[ind] - height[ind - 1]);

    int jumpTwo = INT_MAX;
    if (ind > 1)
        jumpTwo = solve(ind - 2, height, dp) + abs(height[ind] - height[ind - 2]);

    return dp[ind] = min(jumpOne, jumpTwo);
}

int main()
{
    vector<int> height{30, 10, 60, 10, 60, 50};
    int n = height.size();
    vector<int> dp(n, -1);
    cout << solve(n - 1, height, dp);
}

// iterative/tabulation
int main()
{

    vector<int> height{30, 10, 60, 10, 60, 50};
    int n = height.size();
    vector<int> dp(n, -1);
    dp[0] = 0;
    for (int ind = 1; ind < n; ind++)
    {
        int jumpOne = dp[ind - 1] + abs(height[ind] - height[ind - 1]);
        int jumpTwo = INT_MAX;
        if (ind > 1)
            jumpTwo = dp[ind - 2] + abs(height[ind] - height[ind - 2]);

        dp[ind] = min(jumpOne, jumpTwo);
    }
    cout << dp[n - 1];
}

// space optimized
int main()
{
    vector<int> height{30, 10, 60, 10, 60, 50};
    int n = height.size();
    int prev = 0;
    int prev2 = 0;
    for (int i = 1; i < n; i++)
    {

        int jumpTwo = INT_MAX;
        int jumpOne = prev + abs(height[i] - height[i - 1]);
        if (i > 1)
            jumpTwo = prev2 + abs(height[i] - height[i - 2]);

        int cur_i = min(jumpOne, jumpTwo);
        prev2 = prev;
        prev = cur_i;
    }
    cout << prev;
}