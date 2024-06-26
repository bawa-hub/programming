// https://leetcode.com/problems/longest-palindromic-subsequence/

#include <bits/stdc++.h>

using namespace std;

int lcs(string s1, string s2)
{

    int n = s1.size();
    int m = s2.size();

    vector<vector<int>> dp(n + 1, vector<int>(m + 1, -1));
    for (int i = 0; i <= n; i++)
    {
        dp[i][0] = 0;
    }
    for (int i = 0; i <= m; i++)
    {
        dp[0][i] = 0;
    }

    for (int ind1 = 1; ind1 <= n; ind1++)
    {
        for (int ind2 = 1; ind2 <= m; ind2++)
        {
            if (s1[ind1 - 1] == s2[ind2 - 1])
                dp[ind1][ind2] = 1 + dp[ind1 - 1][ind2 - 1];
            else
                dp[ind1][ind2] = 0 + max(dp[ind1 - 1][ind2], dp[ind1][ind2 - 1]);
        }
    }

    return dp[n][m];
}

int longestPalindromeSubsequence(string s)
{
    string t = s;
    reverse(s.begin(), s.end());
    return lcs(s, t);
}
// Time Complexity: O(N*N)
// Reason: There are two nested loops

// Space Complexity: O(N*N)
// Reason: We are using an external array of size ‘(N*N)’. Stack Space is eliminated.

// space optimized
int lcs(string s1, string s2)
{

    int n = s1.size();
    int m = s2.size();

    vector<int> prev(m + 1, 0), cur(m + 1, 0);

    // Base Case is covered as we have initialized the prev and cur to 0.

    for (int ind1 = 1; ind1 <= n; ind1++)
    {
        for (int ind2 = 1; ind2 <= m; ind2++)
        {
            if (s1[ind1 - 1] == s2[ind2 - 1])
                cur[ind2] = 1 + prev[ind2 - 1];
            else
                cur[ind2] = 0 + max(prev[ind2], cur[ind2 - 1]);
        }
        prev = cur;
    }

    return prev[m];
}

int longestPalindromeSubsequence(string s)
{
    string t = s;
    reverse(s.begin(), s.end());
    return lcs(s, t);
}
// Time Complexity: O(N*N)
// Reason: There are two nested loops.

// Space Complexity: O(N)
// Reason: We are using an external array of size ‘N+1’ to store only two rows.

int main()
{

    string s = "bbabcbcab";

    cout << "The Length of Longest Palindromic Subsequence is " << longestPalindromeSubsequence(s);
}