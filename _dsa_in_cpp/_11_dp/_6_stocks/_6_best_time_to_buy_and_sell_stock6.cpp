// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

#include <bits/stdc++.h>

using namespace std;

// recursive
int getAns(vector<int> &Arr, int ind, int buy, int n, int fee, vector<vector<int>> &dp)
{

    if (ind == n)
        return 0; // base case

    if (dp[ind][buy] != -1)
        return dp[ind][buy];

    int profit;

    if (buy == 0)
    { // We can buy the stock
        profit = max(0 + getAns(Arr, ind + 1, 0, n, fee, dp), -Arr[ind] + getAns(Arr, ind + 1, 1, n, fee, dp));
    }

    if (buy == 1)
    { // We can sell the stock
        profit = max(0 + getAns(Arr, ind + 1, 1, n, fee, dp), Arr[ind] - fee + getAns(Arr, ind + 1, 0, n, fee, dp));
    }

    return dp[ind][buy] = profit;
}

int maximumProfit(int n, int fee, vector<int> &Arr)
{
    // Write your code here

    vector<vector<int>> dp(n, vector<int>(2, -1));

    if (n == 0)
        return 0;
    int ans = getAns(Arr, 0, 0, n, fee, dp);
    return ans;
}

// Time Complexity: O(N*2)
// Reason: There are N*2 states therefore at max ‘N*2’ new problems will be solved and we are running a for loop for ‘N’ times to calculate the total sum

// Space Complexity: O(N*2) + O(N)
// Reason: We are using a recursion stack space(O(N)) and a 2D array ( O(N*2)).

// tabulation

int maximumProfit(int n, int fee, vector<int> &Arr)
{
    // Write your code here
    if (n == 0)
        return 0;

    vector<vector<int>> dp(n + 1, vector<int>(2, 0));

    // base condition is handled by initializing dp array as 0

    for (int ind = n - 1; ind >= 0; ind--)
    {
        for (int buy = 0; buy <= 1; buy++)
        {
            int profit;

            if (buy == 0)
            { // We can buy the stock
                profit = max(0 + dp[ind + 1][0], -Arr[ind] + dp[ind + 1][1]);
            }

            if (buy == 1)
            { // We can sell the stock
                profit = max(0 + dp[ind + 1][1], Arr[ind] - fee + dp[ind + 1][0]);
            }

            dp[ind][buy] = profit;
        }
    }

    return dp[0][0];
}
// Time Complexity: O(N*2)
// Reason: There are two nested loops that account for O(N*2) complexity.

// Space Complexity: O(N*2)
// Reason: We are using an external array of size ‘N*2’. Stack Space is eliminated.

// space optimized
int maximumProfit(int n, int fee, vector<int> &Arr)
{
    if (n == 0)
        return 0;

    vector<long> ahead(2, 0);
    vector<long> cur(2, 0);

    // base condition
    ahead[0] = ahead[1] = 0;

    long profit;

    for (int ind = n - 1; ind >= 0; ind--)
    {
        for (int buy = 0; buy <= 1; buy++)
        {
            if (buy == 0)
            { // We can buy the stock
                profit = max(0 + ahead[0], -Arr[ind] + ahead[1]);
            }

            if (buy == 1)
            { // We can sell the stock
                profit = max(0 + ahead[1], Arr[ind] - fee + ahead[0]);
            }
            cur[buy] = profit;
        }

        ahead = cur;
    }
    return cur[0];
}
// Time Complexity: O(N*2)
// Reason: There are two nested loops that account for O(N*2) complexity

// Space Complexity: O(1)
// Reason: We are using an external array of size ‘2’.

int main()
{

    vector<int> prices = {1, 3, 2, 8, 4, 9};
    int n = prices.size();
    int fee = 2;

    cout << "The maximum profit that can be generated is " << maximumProfit(n, fee, prices);
}
