// https://leetcode.com/problems/coin-change-2/
#include <bits/stdc++.h>

using namespace std;

// recursive
// long countWaysToMakeChangeUtil(vector<int>& arr,int ind, int T, vector<vector<long
// >>& dp){

//     if(ind == 0){
//         return (T%arr[0]==0);
//     }

//     if(dp[ind][T]!=-1)
//         return dp[ind][T];

//     long notTaken = countWaysToMakeChangeUtil(arr,ind-1,T,dp);

//     long taken = 0;
//     if(arr[ind] <= T)
//         taken = countWaysToMakeChangeUtil(arr,ind,T-arr[ind],dp);

//     return dp[ind][T] = notTaken + taken;
// }

// long countWaysToMakeChange(vector<int>& arr, int n, int T){

//     vector<vector<long>> dp(n,vector<long>(T+1,-1));
//     return countWaysToMakeChangeUtil(arr,n-1, T, dp);
// }

// int main() {

//   vector<int> arr ={1,2,3};
//   int target=4;

//   int n =arr.size();

//   cout<<"The total number of ways is " <<countWaysToMakeChange(arr,n,target);
// }

// tabulation
// long countWaysToMakeChange(vector<int>& arr, int n, int T){

//     vector<vector<long>> dp(n,vector<long>(T+1,0));

//     //Initializing base condition
//     for(int i=0;i<=T;i++){
//         if(i%arr[0]==0)
//             dp[0][i]=1;
//         // Else condition is automatically fulfilled,
//         // as dp array is initialized to zero
//     }

//     for(int ind=1; ind<n;ind++){
//         for(int target=0;target<=T;target++){
//             long notTaken = dp[ind-1][target];

//             long taken = 0;
//             if(arr[ind]<=target)
//                 taken = dp[ind][target-arr[ind]];

//             dp[ind][target] = notTaken + taken;
//         }
//     }

//     return dp[n-1][T];
// }

// int main() {

//   vector<int> arr ={1,2,3};
//   int target=4;

//   int n =arr.size();

//   cout<<"The total number of ways is " <<countWaysToMakeChange(arr,n,target);
// }

// space optimized
long countWaysToMakeChange(vector<int> &arr, int n, int T)
{

    vector<long> prev(T + 1, 0);

    // Initializing base condition
    for (int i = 0; i <= T; i++)
    {
        if (i % arr[0] == 0)
            prev[i] = 1;
        // Else condition is automatically fulfilled,
        // as prev array is initialized to zero
    }

    for (int ind = 1; ind < n; ind++)
    {
        vector<long> cur(T + 1, 0);
        for (int target = 0; target <= T; target++)
        {
            long notTaken = prev[target];

            long taken = 0;
            if (arr[ind] <= target)
                taken = cur[target - arr[ind]];

            cur[target] = notTaken + taken;
        }
        prev = cur;
    }

    return prev[T];
}

int main()
{

    vector<int> arr = {1, 2, 3};
    int target = 4;

    int n = arr.size();

    cout << "The total number of ways is " << countWaysToMakeChange(arr, n, target);
}
