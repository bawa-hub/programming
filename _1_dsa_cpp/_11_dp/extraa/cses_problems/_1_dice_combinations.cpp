#include <bits/stdc++.h>
using namespace std;

#define mod 1000000007

int main()
{
    long long n;
    cin >> n;
    vector<long long> dp(n + 1, 0);
    dp[0] = dp[1] = 1; // base case
    for (int i = 2; i <= n; i++)
    {
        for (int j = 1; j <= 6; j++)
        { // dice throw
            if (j > i)
                continue;
            dp[i] = (dp[i] % mod + dp[i - j] % mod) % mod;
        }
    }
    cout << dp[n] % mod;
}