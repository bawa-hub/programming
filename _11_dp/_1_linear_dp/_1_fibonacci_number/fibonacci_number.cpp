// https://leetcode.com/problems/fibonacci-number/description/

#include <bits/stdc++.h>
using namespace std;

// Top-Down (recursion+memoization) approach
int f(int n, vector<int> &dp)
{
    if (n <= 1)
        return n;

    if (dp[n] != -1)
        return dp[n];

    return dp[n] = f(n - 1, dp) + f(n - 2, dp);
}
// Time complexity - O(N)
// Space complexity - O(N) + O(N) - stack space + array

// Bottom-up (tabulation) approach
int f(int n, vector<int> &dp)
{

    dp[0] = 0;
    dp[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        dp[i] = dp[i - 1] + dp[i - 2];
    }
    return dp[n];
}
// Time complexity - O(N)
// Space complexity -O(N)

int main()
{
    int n;
    cin >> n;
    vector<int> dp(n + 1, -1);
    cout << f(n, dp);
    return 0;
}

// space optimized
int main()
{
    int n;
    cin >> n;
    int prev2 = 0;
    int prev = 1;
    for (int i = 2; i <= n; i++)
    {
        int current = prev + prev2;
        prev2 = prev;
        prev = current;
    }
    cout << prev;
    return 0;
}