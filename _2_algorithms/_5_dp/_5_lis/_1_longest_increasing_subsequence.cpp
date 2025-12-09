// https://leetcode.com/problems/longest-increasing-subsequence/

#include <bits/stdc++.h>
using namespace std;

// memoization
int getAns(int arr[], int n, int ind, int prev_index, vector<vector<int>> &dp)
{

    // base condition
    if (ind == n)
        return 0;

    if (dp[ind][prev_index + 1] != -1)
        return dp[ind][prev_index + 1];

    int notTake = 0 + getAns(arr, n, ind + 1, prev_index, dp);

    int take = 0;

    if (prev_index == -1 || arr[ind] > arr[prev_index])
    {
        take = 1 + getAns(arr, n, ind + 1, ind, dp);
    }

    // coordinate shift for tacking -1 index
    return dp[ind][prev_index + 1] = max(notTake, take);
}

int longestIncreasingSubsequence(int arr[], int n)
{

    vector<vector<int>> dp(n, vector<int>(n + 1, -1));

    return getAns(arr, n, 0, -1, dp);
}

// Time Complexity: O(N*N)
// Reason: There are N*N states therefore at max ‘N*N’ new problems will be solved.

// Space Complexity: O(N*N) + O(N)
// Reason: We are using an auxiliary recursion stack space(O(N)) (see the recursive tree, in the worst case we will go till N calls at a time) and a 2D array ( O(N*N+1)).

// tabulation
int longestIncreasingSubsequence(int arr[], int n){
    
    vector<vector<int>> dp(n+1,vector<int>(n+1,0));
    
    for(int ind = n-1; ind>=0; ind --){
        for (int prev_index = ind-1; prev_index >=-1; prev_index --){
            
            int notTake = 0 + dp[ind+1][prev_index +1];
    
            int take = 0;
    
            if(prev_index == -1 || arr[ind] > arr[prev_index]){
                
                take = 1 + dp[ind+1][ind+1];
            }
    
            dp[ind][prev_index+1] = max(notTake,take);
            
        }
    }
    
    return dp[0][0];
}

// Time Complexity: O(N*N)
// Reason: There are two nested loops
// Space Complexity: O(N*N)
// Reason: We are using an external array of size ‘(N+1)*(N+1)’. Stack Space is eliminated.

// space optimized
int longestIncreasingSubsequence(int arr[], int n){
    
    vector<int> next(n+1,0);
    
    vector<int> cur(n+1,0);
    
    for(int ind = n-1; ind>=0; ind --){
        for (int prev_index = ind-1; prev_index >=-1; prev_index --){
            
            int notTake = 0 + next[prev_index +1];
    
            int take = 0;
    
            if(prev_index == -1 || arr[ind] > arr[prev_index]){
                
                take = 1 + next[ind+1];
            }
    
            cur[prev_index+1] = max(notTake,take);
        }
        next = cur;
    }
    
    return cur[0];
}

// Time Complexity: O(N*N)
// Reason: There are two nested loops.
// Space Complexity: 2*O(N)
// Reason: We are only using two rows of size n.

// tabulation algorithm
int longestIncreasingSubsequence(int arr[], int n){
    
    vector<int> dp(n,1);
    
    for(int i=0; i<=n-1; i++){
        for(int prev_index = 0; prev_index <=i-1; prev_index ++){
            
            if(arr[prev_index]<arr[i]){
                dp[i] = max(dp[i], 1 + dp[prev_index]);
            }
        }
    }
    
    int ans = -1;
    
    for(int i=0; i<=n-1; i++){
        ans = max(ans, dp[i]);
    }
    
    return ans;
}
// Time Complexity: O(N*N)
// Reason: There are two nested loops.
// Space Complexity: O(N)
// Reason: We are only using two rows of size ‘N’.

int main()
{

    int arr[] = {10, 9, 2, 5, 3, 7, 101, 18};

    int n = sizeof(arr) / sizeof(arr[0]);

    cout << "The length of the longest increasing subsequence is "
         << longestIncreasingSubsequence(arr, n);

    return 0;
}


// binary search
class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int n = nums.size();
        vector<int> temp;
        temp.push_back(nums[0]);

        int len = 1;

        for(int i=0;i<n;i++) {
            if(nums[i]>temp.back()) {
                temp.push_back(nums[i]);
                len++;
            } else {
                int ind = lower_bound(temp.begin(), temp.end(), nums[i]) - temp.begin();
                temp[ind] = nums[i];
            }
        }

        return len;
    }
};

// TC - O(nlogn)
// SC - O(n)