// https://practice.geeksforgeeks.org/problems/matrix-chain-multiplication0303/1

#include <bits/stdc++.h>
using namespace std;

// top-down
int f(vector<int> &arr, int i, int j, vector<vector<int>> &dp)
{

    // base condition
    if (i == j)
        return 0;

    if (dp[i][j] != -1)
        return dp[i][j];

    int mini = INT_MAX;

    // partioning loop
    for (int k = i; k <= j - 1; k++)
    {

        int ans = f(arr, i, k, dp) + f(arr, k + 1, j, dp) + arr[i - 1] * arr[k] * arr[j];

        mini = min(mini, ans);
    }

    return dp[i][j] = mini;
}

int matrixMultiplication(vector<int> &arr, int N)
{

    vector<vector<int>> dp(N, vector<int>(N, -1));

    int i = 1;
    int j = N - 1;

    return f(arr, i, j, dp);
}
// Time Complexity: O(N*N*N)
// Reason: There are N*N states and we explicitly run a loop inside the function which will run for N times, therefore at max ‘N*N*N’ new problems will be solved.

// Space Complexity: O(N*N) + O(N)
// Reason: We are using an auxiliary recursion stack space(O(N))and a 2D array ( O(N*N)).


// tabulation
int matrixMultiplication(vector<int>& arr, int N) {
    // Create a DP table to store the minimum number of operations
    vector<vector<int>> dp(N, vector<int>(N, -1));

    // Initialize the diagonal elements of the DP table to 0
    for (int i = 0; i < N; i++) {
        dp[i][i] = 0;
    }

    // Loop for the length of the chain
    for (int len = 2; len < N; len++) {
        for (int i = 1; i < N - len + 1; i++) {
            int j = i + len - 1;
            dp[i][j] = INT_MAX;

            // Try different partition points to find the minimum
            for (int k = i; k < j; k++) {
                int cost = dp[i][k] + dp[k + 1][j] + arr[i - 1] * arr[k] * arr[j];
                dp[i][j] = min(dp[i][j], cost);
            }
        }
    }

    // The result is stored in dp[1][N-1]
    return dp[1][N - 1];
}

// Time Complexity: O(N*N*N)
// Reason: There are N*N states and we explicitly run a loop inside the function which will run for N times, therefore at max ‘N*N*N’ new problems will be solved.

// Space Complexity: O(N*N)
// Reason: We are using a 2D array ( O(N*N)) space.

int main()
{

    vector<int> arr = {10, 20, 30, 40, 50};

    int n = arr.size();

    cout << "The minimum number of operations is " << matrixMultiplication(arr, n);

    return 0;
}