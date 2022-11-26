// https://www.codingninjas.com/codestudio/problems/0-1-knapsack_920542?source=youtube&campaign=striver_dp_videos&utm_source=youtube&utm_medium=affiliate&utm_campaign=striver_dp_videos
// https: // www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/

#include <bits/stdc++.h>

using namespace std;

// recursive
// int knapsackUtil(vector<int>& wt, vector<int>& val, int ind, int W, vector<vector<int>>& dp){

//     if(ind == 0){
//         if(wt[0] <=W) return val[0];
//         else return 0;
//     }

//     if(dp[ind][W]!=-1)
//         return dp[ind][W];

//     int notTaken = 0 + knapsackUtil(wt,val,ind-1,W,dp);

//     int taken = INT_MIN;
//     if(wt[ind] <= W)
//         taken = val[ind] + knapsackUtil(wt,val,ind-1,W-wt[ind],dp);

//     return dp[ind][W] = max(notTaken,taken);
// }

// int knapsack(vector<int>& wt, vector<int>& val, int n, int W){

//     vector<vector<int>> dp(n,vector<int>(W+1,-1));
//     return knapsackUtil(wt, val, n-1, W, dp);
// }

// int main() {

//   vector<int> wt = {1,2,4,5};
//   vector<int> val = {5,4,8,6};
//   int W=5;

//   int n = wt.size();

//   cout<<"The Maximum value of items, thief can steal is " <<knapsack(wt,val,n,W);
// }

// tabulation
// int knapsack(vector<int> &wt, vector<int> &val, int n, int W)
// {

//     vector<vector<int>> dp(n, vector<int>(W + 1, 0));

//     // Base Condition

//     for (int i = wt[0]; i <= W; i++)
//     {
//         dp[0][i] = val[0];
//     }

//     for (int ind = 1; ind < n; ind++)
//     {
//         for (int cap = 0; cap <= W; cap++)
//         {

//             int notTaken = 0 + dp[ind - 1][cap];

//             int taken = INT_MIN;
//             if (wt[ind] <= cap)
//                 taken = val[ind] + dp[ind - 1][cap - wt[ind]];

//             dp[ind][cap] = max(notTaken, taken);
//         }
//     }

//     return dp[n - 1][W];
// }

// int main()
// {

//     vector<int> wt = {1, 2, 4, 5};
//     vector<int> val = {5, 4, 8, 6};
//     int W = 5;

//     int n = wt.size();

//     cout << "The Maximum value of items, thief can steal is " << knapsack(wt, val, n, W);
// }

// space optimized
int knapsack(vector<int> &wt, vector<int> &val, int n, int W)
{

    vector<int> prev(W + 1, 0);

    // Base Condition

    for (int i = wt[0]; i <= W; i++)
    {
        prev[i] = val[0];
    }

    for (int ind = 1; ind < n; ind++)
    {
        for (int cap = W; cap >= 0; cap--)
        {

            int notTaken = 0 + prev[cap];

            int taken = INT_MIN;
            if (wt[ind] <= cap)
                taken = val[ind] + prev[cap - wt[ind]];

            prev[cap] = max(notTaken, taken);
        }
    }

    return prev[W];
}

int main()
{

    vector<int> wt = {1, 2, 4, 5};
    vector<int> val = {5, 4, 8, 6};
    int W = 5;

    int n = wt.size();

    cout << "The Maximum value of items, thief can steal is " << knapsack(wt, val, n, W);
}