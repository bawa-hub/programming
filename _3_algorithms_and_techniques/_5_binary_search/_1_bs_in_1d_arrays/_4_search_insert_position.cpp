// https://leetcode.com/problems/search-insert-position/

#include <bits/stdc++.h>
using namespace std;

// brute force
int find_index(int arr[], int n, int K)
{
    // Traverse the array
    for (int i = 0; i < n; i++)
        // If K is found
        if (arr[i] == K)
            return i;
        // If current array element exceeds K
        else if (arr[i] > K)
            return i;
    // If all elements are smaller than K
    return n;
}
// Time Complexity: O(N)
// Auxiliary Space: O(1)

// binary search
int find_index(int arr[], int n, int K)
{
    // Lower and upper bounds
    int start = 0;
    int end = n - 1;
    // Traverse the search space
    while (start <= end)
    {
        int mid = (start + end) / 2;
        // If K is found
        if (arr[mid] == K)
            return mid;
        else if (arr[mid] < K)
            start = mid + 1;
        else
            end = mid - 1;
    }
    // Return insert position
    return end + 1;
}
// Time Complexity: O(log N)
// Auxiliary Space: O(1)

// Driver Code
int main()
{
    int arr[] = {1, 3, 5, 6};
    int n = sizeof(arr) / sizeof(arr[0]);
    int K = 2;
    cout << find_index(arr, n, K) << endl;
    return 0;
}