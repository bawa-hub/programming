#include <bits/stdc++.h>

#define FAST_IO ios_base::sync_with_stdio(false), cin.tie(nullptr)
using namespace std;

signed main()
{
    FAST_IO;

    int n;
    cin >> n;
    int a[n], b[n], c[n];
    for (int i = 0; i < n; i++)
        cin >> a[i] >> b[i] >> c[i];

    int dp[n][3]; // dp[i][j] : i represents the day and j represents the activity
                  // j is 0 for activity A, 1 for activity B and 2 for activity C

    // Set the base values
    dp[0][0] = a[0]; // If you swim in the sea in sea on the first day, you cannot get more than a[0] points.
    dp[0][1] = b[0]; // If you catch bugs in the mountains on the first day, you cannot get more than b[0] points.
    dp[0][2] = c[0]; // If you do homework at home on the first day, you cannot get more than c[0] points.

    // i represents what day it is (0 based indexing)
    for (int i = 1; i < n; i++)
    {
        dp[i][0] = a[i] + max(dp[i - 1][1], dp[i - 1][2]); // If we do activity A today, we have to have done activities B or C on the previous day.
        dp[i][1] = b[i] + max(dp[i - 1][0], dp[i - 1][2]); // If we do activity B today, we have to have done activities A or C on the previous day.
        dp[i][2] = c[i] + max(dp[i - 1][1], dp[i - 1][0]); // If we do activity C today, we have to have done activities A or B on the previous day.
    }

    // Our answer is the maximum of all the points gained by doing activity A, B, or C on the last day.
    cout << max({dp[n - 1][0], dp[n - 1][1], dp[n - 1][2]});
    return 0;
}