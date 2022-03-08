#include <bits/stdc++.h>
using namespace std;

int main()
{
    int n;
    cin >> n;
    vector<int> dp(n + 1);
    dp[0] = 0;
    for (int i = 1; i <= n; i++)
    {
        int temp = i;
        int minSteps = 1e6 + 2;
        while (temp != 0)
        {
            int di = temp % 10;
            temp /= 10;
            if (di == 0)
                continue;
            minSteps = min(minSteps, 1 + dp[i - di]);
        }
        dp[i] = minSteps;
    }

    cout << dp[n];
}