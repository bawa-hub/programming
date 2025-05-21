#include <bits/stdc++.h>
using namespace std;

#define mod 1000000007

int main()
{
    int n;
    cin >> n;
    bool grid[n + 1][n + 1];
    for (int i = 1; i <= n; i++)
    {
        for (int j = 1; j <= n; j++)
        {
            char ch;
            cin >> ch;
            if (ch == '.')
                grid[i][j] = 0;
            else
                grid[i][j] = 1;
        }
    }

    int dp[n + 1][n + 1];
    for (int i = n; i >= 1; i--)
    {
        for (int j = n; j >= 1; j--)
        {
            if (i == n && j == n)
            {
                dp[i][j] = 1;
            }
            else
            {
                int op1 = (j == n) ? 0 : dp[i][j + 1];
                int op2 = (i == n) ? 0 : dp[i + 1][j];
                dp[i][j] = (op1 + op2) % mod;
                if (grid[i][j])
                    dp[i][j] = 0;
            }
        }
    }
    if (grid[n][n])
        cout << 0;
    else
        cout << dp[1][1];
}