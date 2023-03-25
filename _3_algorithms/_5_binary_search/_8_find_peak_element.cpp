// https://leetcode.com/problems/find-peak-element/

#include <iostream>

using namespace std;

// brute force
int peakEleBruteForce(int arr[], int n)
{
    if (arr[0] >= arr[1])
        return arr[0];

    for (int i = 1; i < n - 1; i++)
        if (arr[i] >= arr[i - 1] && arr[i] >= arr[i + 1])
            return arr[i];

    return arr[n - 1];
}
// Time Complexity: O(n), we traverse the whole array once.
// Space Complexity: O(1), we are not using any extra space.

// binary search
int peakEleOptimal(int arr[], int n)
{
    int start = 0, end = n - 1;

    while (start < end)
    {
        int mid = (start + end) / 2;

        if (mid == 0)
            return arr[0] >= arr[1] ? arr[0] : arr[1];

        if (mid == n - 1)
            return arr[n - 1] >= arr[n - 2] ? arr[n - 1] : arr[n - 2];

        // Cheking whether peak ele is in mid position
        if (arr[mid] >= arr[mid - 1] && arr[mid] >= arr[mid + 1])
            return arr[mid];

        // If left ele is greater then ignore 2nd half of the elements
        if (arr[mid] < arr[mid - 1])
            end = mid - 1;

        // Else ignore first half of the elements
        else
            start = mid + 1;
    }

    return arr[start];
}
// Time Complexity: O(log(n)), at every time we shrink the search space to half resulting in log(n) time complexity.
// Space Complexity: O(1), we are not using any extra space.
int main()
{

    int arr[] = {3, 5, 4, 1, 1};
    int n = sizeof(arr) / sizeof(arr[0]);

    cout << "Peak Element is " << peakEleBruteForce(arr, n);

    return 0;
}