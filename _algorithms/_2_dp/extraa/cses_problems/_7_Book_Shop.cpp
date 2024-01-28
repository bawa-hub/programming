#include <bits/stdc++.h>
using namespace std;

int dp[1005][100005];
int main()
{
    long long n, x;
    cin >> n >> x;
    vector<int> price(n + 1);
    vector<int> pages(n + 1);

    for (int i = 1; i <= n; i++)
        cin >> price[i];

    for (int i = 1; i <= n; i++)
        cin >> pages[i];

    vector<vector<int>> dp(n + 1, vector<int>(x + 1, 0));

    for (int i = 1; i <= n; i++)
    {
        for (int j = 1; j <= x; j++)
        {
            if (price[i] > j)
            {
                dp[i][j] = dp[i - 1][j];
            }
            else
            {
                dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - price[i]] + pages[i]);
            }
        }
    }

    cout << dp[n][x];
}