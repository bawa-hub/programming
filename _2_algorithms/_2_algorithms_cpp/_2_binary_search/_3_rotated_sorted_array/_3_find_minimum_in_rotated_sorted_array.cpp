// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

#include <bits/stdc++.h>

using namespace std;
int main()
{
    int nums[5] = {4, 5, 1, 2, 3};
    int i;
    int left = 0, right = 4, minVal = INT_MAX;

    while (left <= right)
    {
        // search space is already sorted , return first element as answ
        if (nums[left] < nums[right])
        {
            minVal = min(minVal, nums[left]);
            break;
        }
        int mid = left + (right - left) / 2;

        if (nums[left] <= nums[mid])
        {
            minVal = min(minVal, nums[left]);
            left = mid + 1;
        }
        else
        {
            minVal = min(minVal, nums[mid]);
            right = mid - 1;
        }
    }
    cout << "Minimum Value is " << minVal;
    return 0;
}

// Time Complexity : O(logn)
// Space Complexity: O(1)