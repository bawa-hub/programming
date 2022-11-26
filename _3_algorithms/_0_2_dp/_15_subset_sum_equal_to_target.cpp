// https://www.codingninjas.com/codestudio/problems/subset-sum-equal-to-k_1550954?leftPanelTab=0

#include <bits/stdc++.h>
using namespace std;

// recursive/memoization
// bool subsetSumUtil(int ind, int target, vector<int> &arr, vector<vector<int>> &dp)
// {
//     if (target == 0)
//         return true;

//     if (ind == 0)
//         return arr[0] == target;

//     if (dp[ind][target] != -1)
//         return dp[ind][target];

//     bool notTaken = subsetSumUtil(ind - 1, target, arr, dp);

//     bool taken = false;
//     if (arr[ind] <= target)
//         taken = subsetSumUtil(ind - 1, target - arr[ind], arr, dp);

//     return dp[ind][target] = notTaken || taken;
// }

// bool subsetSumToK(int n, int k, vector<int> &arr)
// {
//     vector<vector<int>> dp(n, vector<int>(k + 1, -1));

//     return subsetSumUtil(n - 1, k, arr, dp);
// }

// int main()
// {

//     vector<int> arr = {1, 2, 3, 4};
//     int k = 4;
//     int n = arr.size();

//     if (subsetSumToK(n, k, arr))
//         cout << "Subset with given target found";
//     else
//         cout << "Subset with given target not found";
// }

// iterative/tabulation
// bool subsetSumToK(int n, int k, vector<int> &arr)
// {
//     vector<vector<bool>> dp(n, vector<bool>(k + 1, false));

//     for (int i = 0; i < n; i++)
//     {
//         dp[i][0] = true;
//     }

//     if (arr[0] <= k)
//         dp[0][arr[0]] = true;

//     for (int ind = 1; ind < n; ind++)
//     {
//         for (int target = 1; target <= k; target++)
//         {

//             bool notTaken = dp[ind - 1][target];

//             bool taken = false;
//             if (arr[ind] <= target)
//                 taken = dp[ind - 1][target - arr[ind]];

//             dp[ind][target] = notTaken || taken;
//         }
//     }

//     return dp[n - 1][k];
// }

// int main()
// {

//     vector<int> arr = {1, 2, 3, 4};
//     int k = 4;
//     int n = arr.size();

//     if (subsetSumToK(n, k, arr))
//         cout << "Subset with given target found";
//     else
//         cout << "Subset with given target not found";
// }

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