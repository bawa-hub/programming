#include <bits/stdc++.h>
using namespace std;

#define mod 1000000007

int main()
{
    long long n, x;
    cin >> n >> x;
    vector<long long> dp(x + 1, 0);
    dp[0] = 1; // base case
    vector<long long> coins(n);
    for (int i = 0; i <= n - 1; i++)
        cin >> coins[i];
    for (int i = 1; i <= x; i++)
    {
        for (int j = 0; j <= n - 1; j++)
        { // dice throw
            if (coins[j] > i)
                continue;
            dp[i] = (dp[i] + dp[i - coins[j]]) % mod;
        }
    }
    cout << dp[x] % mod;
}