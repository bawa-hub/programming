// https://www.codingninjas.com/codestudio/problems/longest-subarray-with-sum-k_6682399
// https://www.geeksforgeeks.org/problems/longest-sub-array-with-sum-k0809

#include <bits/stdc++.h>
using namespace std;

// brute force
int getLongestSubarray(vector<int> &a, long long k)
{
    int n = a.size(); // size of the array.

    int len = 0;
    for (int i = 0; i < n; i++)
    { // starting index
        for (int j = i; j < n; j++)
        { // ending index
            // add all the elements of
            // subarray = a[i...j]:
            long long s = 0;
            for (int K = i; K <= j; K++)
            {
                s += a[K];
            }

            if (s == k)
                len = max(len, j - i + 1);
        }
    }
    return len;
}
// Time Complexity: O(n^3)
// Space Complexity: O(1)

// optimized brute force
int getLongestSubarray(int arr[], int n, int k)
{
    int maxLength = 0;
    for (int i = 0; i < n; i++)
    {
        int sum = 0;
        for (int j = i; j < n; j++)
        {
            sum += arr[j];
            if (sum == k)
                maxLength = max(maxLength, (j - i + 1));
        }
    }
    return maxLength;
}
// Time Complexity: O(n^2) time to generate all possible subarrays.
// Space Complexity: O(1), we are not using any extra space.

// using hashing
int getLongestSubarray(vector<int> &a, long long k)
{
    int n = a.size(); // size of the array.

    map<long long, int> preSumMap;
    long long sum = 0;
    int maxLen = 0;
    for (int i = 0; i < n; i++)
    {
        // calculate the prefix sum till index i:
        sum += a[i];

        // if the sum = k, update the maxLen:
        if (sum == k)
        {
            maxLen = max(maxLen, i + 1);
        }

        // calculate the sum of remaining part i.e. (prefixSum - k):
        long long rem = sum - k;

    
        // if (prefixSum-k) is present that means we have subarray with sum k.
        if (preSumMap.find(rem) != preSumMap.end())
        {
            int len = i - preSumMap[rem];
            maxLen = max(maxLen, len);
        }

        // Finally, update the map if not already present, because you will not get the longest subarray
        if (preSumMap.find(sum) == preSumMap.end())
        {
            preSumMap[sum] = i;
        }
    }

    return maxLen;
}
// Time Complexity: O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
// Space Complexity: O(N) as we are using a map data structure.

// using two pointer (sliding window)
int getLongestSubarray(vector<int> &a, long long k)
{
    int n = a.size(); // size of the array.

    int left = 0, right = 0; // 2 pointers
    long long sum = a[0];
    int maxLen = 0;
    while (right < n)
    {
        // if sum > k, reduce the subarray from left
        // until sum becomes less or equal to k:
        while (left <= right && sum > k)
        {
            sum -= a[left];
            left++;
        }

        // if sum = k, update the maxLen i.e. answer:
        if (sum == k)
        {
            maxLen = max(maxLen, right - left + 1);
        }

        // Move forward thw right pointer:
        right++;
        if (right < n)
            sum += a[right];
    }

    return maxLen;
}
// Time Complexity: O(2*N)
// Space Complexity: O(1)
