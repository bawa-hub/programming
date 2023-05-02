// https://leetcode.com/problems/subarray-sum-equals-k/

#include <bits/stdc++.h>
using namespace std;

// brute force
int findAllSubarraysWithGivenSum(vector<int> &arr, int k)
{
    int n = arr.size(); // size of the given array.
    int cnt = 0;        // Number of subarrays:

    for (int i = 0; i < n; i++)
    { // starting index i
        for (int j = i; j < n; j++)
        { // ending index j

            // calculate the sum of subarray [i...j]
            int sum = 0;
            for (int K = i; K <= j; K++)
                sum += arr[K];

            // Increase the count if sum == k:
            if (sum == k)
                cnt++;
        }
    }
    return cnt;
}
// Time Complexity: O(N3), where N = size of the array.
// Reason: We are using three nested loops here. Though all are not running for exactly N times, the time complexity will be approximately O(N3).

// Space Complexity: O(1) as we are not using any extra space.

// optimized
int findAllSubarraysWithGivenSum(vector<int> &arr, int k)
{
    int n = arr.size(); // size of the given array.
    int cnt = 0;        // Number of subarrays:

    for (int i = 0; i < n; i++)
    { // starting index i
        int sum = 0;
        for (int j = i; j < n; j++)
        { // ending index j
            // calculate the sum of subarray [i...j]
            // sum of [i..j-1] + arr[j]
            sum += arr[j];

            // Increase the count if sum == k:
            if (sum == k)
                cnt++;
        }
    }
    return cnt;
}
// Time Complexity: O(N2), where N = size of the array.
// Reason: We are using two nested loops here. As each of them is running for exactly N times, the time complexity will be approximately O(N2).

// Space Complexity: O(1) as we are not using any extra space.

// using hashing
int findAllSubarraysWithGivenSum(vector<int> &arr, int k)
{
    int n = arr.size(); // size of the given array.
    map<int, int> mpp;
    int preSum = 0, cnt = 0;

    mpp[0] = 1; // Setting 0 in the map.
    for (int i = 0; i < n; i++)
    {
        // add current element to prefix Sum:
        preSum += arr[i];

        // Calculate x-k:
        int remove = preSum - k;

        // Add the number of subarrays to be removed:
        cnt += mpp[remove];

        // Update the count of prefix sum
        // in the map.
        mpp[preSum] += 1;
    }
    return cnt;
}
// Time Complexity: O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N) but if we are using a map data structure, the time complexity will be O(N*logN). The least complexity will be O(N) as we are using a loop to traverse the array.

// Space Complexity: O(N) as we are using a map data structure.

int main()
{
    vector<int> arr = {3, 1, 2, 4};
    int k = 6;
    int cnt = findAllSubarraysWithGivenSum(arr, k);
    cout << "The number of subarrays is: " << cnt << "\n";
    return 0;
}