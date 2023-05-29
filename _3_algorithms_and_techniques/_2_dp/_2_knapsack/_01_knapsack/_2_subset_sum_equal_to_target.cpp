// https://www.codingninjas.com/codestudio/problems/subset-sum-equal-to-k_1550954?leftPanelTab=0
// https://practice.geeksforgeeks.org/problems/subset-sum-problem-1611555638/1
// https://www.geeksforgeeks.org/subset-sum-problem-dp-25/

#include <bits/stdc++.h>
using namespace std;

// recursive/memoization
bool subsetSumUtil(int ind, int target, vector<int> &arr, vector<vector<int>> &dp)
{
    if (target == 0)
        return true;

    if (ind == 0)
        return arr[0] == target;

    if (dp[ind][target] != -1)
        return dp[ind][target];

    bool notTaken = subsetSumUtil(ind - 1, target, arr, dp);

    bool taken = false;
    if (arr[ind] <= target)
        taken = subsetSumUtil(ind - 1, target - arr[ind], arr, dp);

    return dp[ind][target] = notTaken || taken;
}

bool subsetSumToK(int n, int k, vector<int> &arr)
{
    vector<vector<int>> dp(n, vector<int>(k + 1, -1));

    return subsetSumUtil(n - 1, k, arr, dp);
}
// Time Complexity: O(N*K)
// Reason: There are N*K states therefore at max ‘N*K’ new problems will be solved.
// Space Complexity: O(N*K) + O(N)
// Reason: We are using a recursion stack space(O(N)) and a 2D array ( O(N*K)).

// iterative/tabulation
bool subsetSumToK(int n, int k, vector<int> &arr)
{
    vector<vector<bool>> dp(n, vector<bool>(k + 1, false));

    for (int i = 0; i < n; i++)
    {
        dp[i][0] = true;
    }

    if (arr[0] <= k)
        dp[0][arr[0]] = true;

    for (int ind = 1; ind < n; ind++)
    {
        for (int target = 1; target <= k; target++)
        {

            bool notTaken = dp[ind - 1][target];

            bool taken = false;
            if (arr[ind] <= target)
                taken = dp[ind - 1][target - arr[ind]];

            dp[ind][target] = notTaken || taken;
        }
    }

    return dp[n - 1][k];
}
// Time Complexity: O(N*K)
// Reason: There are two nested loops
// Space Complexity: O(N*K)
// Reason: We are using an external array of size ‘N*K’. Stack Space is eliminated.

// space optimized
bool subsetSumToK(int n, int k, vector<int> &arr)
{
    vector<bool> prev(k + 1, false);

    prev[0] = true;

    if (arr[0] <= k)
        prev[arr[0]] = true;

    for (int ind = 1; ind < n; ind++)
    {
        vector<bool> cur(k + 1, false);
        cur[0] = true;
        for (int target = 1; target <= k; target++)
        {
            bool notTaken = prev[target];

            bool taken = false;
            if (arr[ind] <= target)
                taken = prev[target - arr[ind]];

            cur[target] = notTaken || taken;
        }
        prev = cur;
    }

    return prev[k];
}
// Time Complexity: O(N*K)
// Reason: There are three nested loops
// Space Complexity: O(K)
// Reason: We are using an external array of size ‘K+1’ to store only one row.

int main()
{

    vector<int> arr = {1, 2, 3, 4};
    int k = 4;
    int n = arr.size();

    if (subsetSumToK(n, k, arr))
        cout << "Subset with given target found";
    else
        cout << "Subset with given target not found";
}