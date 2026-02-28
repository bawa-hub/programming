// https://leetcode.com/problems/house-robber/
// https://practice.geeksforgeeks.org/problems/max-sum-without-adjacents2430/1
// https://www.geeksforgeeks.org/maximum-sum-such-that-no-two-elements-are-adjacent/


#include <iostream>
#include <vector>

using namespace std;

// recursive/memoization
int solveUtil(int ind, vector<int> &arr, vector<int> &dp)
{

    if (dp[ind] != -1)
        return dp[ind];

    if (ind == 0)
        return arr[ind];
    if (ind < 0)
        return 0;

    int pick = arr[ind] + solveUtil(ind - 2, arr, dp);
    int nonPick = 0 + solveUtil(ind - 1, arr, dp);

    return dp[ind] = max(pick, nonPick);
}

// tabulation/iterative
int solveUtil(int n, vector<int> &arr, vector<int> &dp)
{

    dp[0] = arr[0];

    for (int i = 1; i < n; i++)
    {
        int pick = arr[i];
        if (i > 1)
            pick += dp[i - 2];
        int nonPick = 0 + dp[i - 1];

        dp[i] = max(pick, nonPick);
    }

    return dp[n - 1];
}

int solve(int n, vector<int> &arr)
{
    vector<int> dp(n, -1);
    return solveUtil(n, arr, dp);
}

// space optimized
int solve(int n, vector<int> &arr)
{
    int prev = arr[0];
    int prev2 = 0;

    for (int i = 1; i < n; i++)
    {
        int pick = arr[i];
        if (i > 1)
            pick += prev2;
        int nonPick = 0 + prev;

        int cur_i = max(pick, nonPick);
        prev2 = prev;
        prev = cur_i;
    }
    return prev;
}

int main()
{

    vector<int> arr{2, 1, 4, 9};
    int n = arr.size();
    cout << solve(n, arr);
}

// Time Complexity: O(N)
// Reason: The overlapping subproblems will return the answer in constant time O(1). Therefore the total number of new subproblems we solve is ‘n’. Hence total time complexity is O(N).
// Space Complexity: O(N)
// Reason: We are using a recursion stack space(O(N)) and an array (again O(N)). Therefore total space complexity will be O(N) + O(N) ≈ O(N)
