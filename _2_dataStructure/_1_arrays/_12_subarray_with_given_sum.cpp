// https://leetcode.com/problems/subarray-sum-equals-k/
// https://practice.geeksforgeeks.org/problems/longest-sub-array-with-sum-k0809/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=longest-sub-array-with-sum-k

#include <bits/stdc++.h>
using namespace std;

// int longestSubArrWithSumK_BF(int arr[], int n, int k)
// {
//     int maxLength = 0;
//     for (int i = 0; i < n; i++)
//     {
//         int sum = 0;
//         for (int j = i; j < n; j++)
//         {
//             sum += arr[j];
//             if (sum == k)
//                 maxLength = max(maxLength, (j - i + 1));
//         }
//     }
//     return maxLength;
// }

// int main()
// {

//     int arr[] = {7, 1, 6, 0};
//     int n = sizeof(arr) / sizeof(arr[0]), k = 7;

//     cout << "Length of the longest subarray with sum K is " << longestSubArrWithSumK_BF(arr, n, k);

//     return 0;
// }

// Time Complexity: O(n^2) time to generate all possible subarrays.
// Space Complexity: O(1), we are not using any extra space.

// optimized
int longestSubArrWithSumK_Optimal(int arr[], int n, int k)
{
    int start = 0, end = -1, sum = 0, maxLength = 0;
    while (start < n)
    {
        while ((end + 1 < n) && (sum + arr[end + 1] <= k))
            sum += arr[++end];

        if (sum == k)
            maxLength = max(maxLength, (end - start + 1));

        sum -= arr[start];
        start++;
    }
    return maxLength;
}

int main()
{

    int arr[] = {7, 1, 6, 0};
    int n = sizeof(arr) / sizeof(arr[0]), k = 7;

    cout << "Length of the longest subarray with sum K is " << longestSubArrWithSumK_Optimal(arr, n, k);

    return 0;
}

// TC: O(n)
// SC: O(1)