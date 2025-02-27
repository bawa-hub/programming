// https://leetcode.com/problems/longest-increasing-subsequence/

#include <bits/stdc++.h>
using namespace std;

int getAns(int arr[], int n, int ind, int prev_index, vector<vector<int>> &dp)
{

    // base condition
    if (ind == n)
        return 0;

    if (dp[ind][prev_index + 1] != -1)
        return dp[ind][prev_index + 1];

    int notTake = 0 + getAns(arr, n, ind + 1, prev_index, dp);

    int take = 0;

    if (prev_index == -1 || arr[ind] > arr[prev_index])
    {
        take = 1 + getAns(arr, n, ind + 1, ind, dp);
    }

    // coordinate shift for tacking -1 index
    return dp[ind][prev_index + 1] = max(notTake, take);
}

int longestIncreasingSubsequence(int arr[], int n)
{

    vector<vector<int>> dp(n, vector<int>(n + 1, -1));

    return getAns(arr, n, 0, -1, dp);
}

// Time Complexity: O(N*N)
// Reason: There are N*N states therefore at max ‘N*N’ new problems will be solved.

// Space Complexity: O(N*N) + O(N)
// Reason: We are using an auxiliary recursion stack space(O(N)) (see the recursive tree, in the worst case we will go till N calls at a time) and a 2D array ( O(N*N+1)).

int main()
{

    int arr[] = {10, 9, 2, 5, 3, 7, 101, 18};

    int n = sizeof(arr) / sizeof(arr[0]);

    cout << "The length of the longest increasing subsequence is "
         << longestIncreasingSubsequence(arr, n);

    return 0;
}