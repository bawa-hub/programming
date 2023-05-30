// https://leetcode.com/problems/find-peak-element/
// https://leetcode.com/problems/find-peak-element/solutions/1290642/intuition-behind-conditions-complete-explanation-diagram-binary-search/

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
int findPeakElement(vector<int> &nums)
{

    if (nums.size() == 1)
        return 0;
    int n = nums.size();
    if (nums[0] > nums[1])
        return 0;
    if (nums[n - 1] > nums[n - 2])
        return n - 1;
    int start = 1;
    int end = n - 2;
    while (start <= end)
    {
        int mid = start + (end - start) / 2;
        if (nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1])
            return mid;
        else if (nums[mid] < nums[mid - 1])
            end = mid - 1;
        else if (nums[mid] < nums[mid + 1])
            start = mid + 1;
    }
    return -1;
}