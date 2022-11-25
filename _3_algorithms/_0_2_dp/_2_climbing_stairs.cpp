// https://leetcode.com/problems/climbing-stairs/

#include <bits/stdc++.h>
using namespace std;

// recursive
// class Solution {
// public:
//     int climbStairs(int n) {
//         if(n==0 || n==1) return 1;
//         return climbStairs(n-1)+climbStairs(n-2);
//     }
// };

// iterative
// int main()
// {

//     int n = 3;
//     vector<int> dp(n + 1, -1);

//     dp[0] = 1;
//     dp[1] = 1;

//     for (int i = 2; i <= n; i++)
//     {
//         dp[i] = dp[i - 1] + dp[i - 2];
//     }
//     cout << dp[n];
//     return 0;
// }
// TC: O(N)
// SC: O(N)

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

// TC: O(N)
// SC: O(1)