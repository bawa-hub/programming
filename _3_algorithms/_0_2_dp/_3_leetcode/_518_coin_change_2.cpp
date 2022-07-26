// https://leetcode.com/problems/coin-change-2/

#include <bits/stdc++.h>
using namespace std;

const int N = 2510;

int dp[310][10010];

int func(int idx, int amount, vector<int> &coins)
{

    if (amount == 0)
        return 1;

    if (idx < 0)
        return 0;

    if (dp[idx][amount] != -1)
        return dp[idx][amount];
    int ways = 0;

    for (int coin_amount = 0; coin_amount <= amount; coin_amount += coins[idx])
    {
        ways += func(idx - 1, amount - coin_amount, coins);
    }

    return dp[idx][amount] = ways;
}

int coinChange(vector<int> &coins, int amount)
{
    memset(dp, -1, sizeof(dp));
    return func(coins.size() - 1, amount, coins);
}

int main()
{
    vector<int> coins = {2};
    cout << coinChange(coins, 11);
}