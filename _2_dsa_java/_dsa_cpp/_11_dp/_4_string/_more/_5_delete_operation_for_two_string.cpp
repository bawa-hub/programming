// https://leetcode.com/problems/delete-operation-for-two-strings/

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

int canYouMake(string str1, string str2)
{

    int n = str1.size();
    int m = str2.size();

    int k = lcs(str1, str2);

    return (n - k) + (m - k);
}

// Time Complexity: O(N*M)
// Reason: There are two nested loops

// Space Complexity: O(N*M)
// Reason: We are using an external array of size (N*M). Stack Space is eliminated.

// space optimization
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

int canYouMake(string str1, string str2)
{

    int n = str1.size();
    int m = str2.size();

    int k = lcs(str1, str2);

    return (n - k) + (m - k);
}
// Time Complexity: O(N*M)
// Reason: There are two nested loops.

// Space Complexity: O(M)
// Reason: We are using an external array of size ‘M+1’ to store only two rows.

int main()
{

    string str1 = "abcd";
    string str2 = "anc";

    cout << "The Minimum operations required to convert str1 to str2: " << canYouMake(str1, str2);
}