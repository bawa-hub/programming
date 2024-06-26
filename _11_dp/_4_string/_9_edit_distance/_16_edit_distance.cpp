// https://leetcode.com/problems/edit-distance/

#include <bits/stdc++.h>
using namespace std;

int editDistanceUtil(string &S1, string &S2, int i, int j, vector<vector<int>> &dp)
{

    if (i < 0)
        return j + 1;
    if (j < 0)
        return i + 1;

    if (dp[i][j] != -1)
        return dp[i][j];

    if (S1[i] == S2[j])
        return dp[i][j] = 0 + editDistanceUtil(S1, S2, i - 1, j - 1, dp);

    // Minimum of three choices
    else
        return dp[i][j] = 1 + min(editDistanceUtil(S1, S2, i - 1, j - 1, dp),
                                  min(editDistanceUtil(S1, S2, i - 1, j, dp), editDistanceUtil(S1, S2, i, j - 1, dp)));
}

int editDistance(string &S1, string &S2)
{

    int n = S1.size();
    int m = S2.size();

    vector<vector<int>> dp(n, vector<int>(m, -1));
    return editDistanceUtil(S1, S2, n - 1, m - 1, dp);
}

// Time Complexity: O(N*M)
// Reason: There are N*M states therefore at max ‘N*M’ new problems will be solved.

// Space Complexity: O(N*M) + O(N+M)
// Reason: We are using a recursion stack space(O(N+M)) and a 2D array ( O(N*M)).

// tabulation
int editDistance(string &S1, string &S2)
{

    int n = S1.size();
    int m = S2.size();

    vector<vector<int>> dp(n + 1, vector<int>(m + 1, 0));

    for (int i = 0; i <= n; i++)
    {
        dp[i][0] = i;
    }
    for (int j = 0; j <= m; j++)
    {
        dp[0][j] = j;
    }

    for (int i = 1; i < n + 1; i++)
    {
        for (int j = 1; j < m + 1; j++)
        {
            if (S1[i - 1] == S2[j - 1])
                dp[i][j] = 0 + dp[i - 1][j - 1];

            else
                dp[i][j] = 1 + min(dp[i - 1][j - 1], min(dp[i - 1][j], dp[i][j - 1]));
        }
    }

    return dp[n][m];
}
// Time Complexity: O(N*M)
// Reason: There are two nested loops

// Space Complexity: O(N*M)
// Reason: We are using an external array of size ‘N*M’. Stack Space is eliminated.

// space optimization
int editDistance(string &S1, string &S2)
{

    int n = S1.size();
    int m = S2.size();

    vector<int> prev(m + 1, 0);
    vector<int> cur(m + 1, 0);

    for (int j = 0; j <= m; j++)
    {
        prev[j] = j;
    }

    for (int i = 1; i < n + 1; i++)
    {
        cur[0] = i;
        for (int j = 1; j < m + 1; j++)
        {
            if (S1[i - 1] == S2[j - 1])
                cur[j] = 0 + prev[j - 1];

            else
                cur[j] = 1 + min(prev[j - 1], min(prev[j], cur[j - 1]));
        }
        prev = cur;
    }

    return prev[m];
}
// Time Complexity: O(N*M)
// Reason: There are two nested loops.

// Space Complexity: O(M)
// Reason: We are using an external array of size ‘M+1’ to store two rows.

int main()
{

    string s1 = "horse";
    string s2 = "ros";

    cout << "The minimum number of operations required is: " << editDistance(s1, s2);
}
