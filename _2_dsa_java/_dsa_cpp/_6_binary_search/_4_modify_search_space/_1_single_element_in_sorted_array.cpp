// https://leetcode.com/problems/single-element-in-a-sorted-array/

// brute force 1
int singleNonDuplicate(vector<int>& arr) {
    int n = arr.size(); //size of the array.
    if (n == 1) return arr[0];

    for (int i = 0; i < n; i++) {

        //Check for first index:
        if (i == 0) {
            if (arr[i] != arr[i + 1])
                return arr[i];
        }
        //Check for last index:
        else if (i == n - 1) {
            if (arr[i] != arr[i - 1])
                return arr[i];
        }
        else {
            if (arr[i] != arr[i - 1] && arr[i] != arr[i + 1])
                return arr[i];
        }
    }

    // dummy return statement:
    return -1;
}
// Time Complexity: O(N), N = size of the given array.
// Reason: We are traversing the entire array.
// Space Complexity: O(1) as we are not using any extra space.

// using XOR
int singleNonDuplicate(vector<int>& arr) {
    int n = arr.size(); //size of the array.
    int ans = 0;
    // XOR all the elements:
    for (int i = 0; i < n; i++) {
        ans = ans ^ arr[i];
    }
    return ans;
}
// Time Complexity: O(N), N = size of the given array.
// Reason: We are traversing the entire array.
// Space Complexity: O(1) as we are not using any extra space.

// binary search
class Solution
{
public:
    int singleNonDuplicate(vector<int> &nums)
    {

        int n = nums.size();

        // conditions to avoid conditions in while loop
        if (n == 1)
            return nums[0];
        if (nums[0] != nums[1])
            return nums[0];
        if (nums[n - 1] != nums[n - 2])
            return nums[n - 1];

        int low = 1, high = n - 2;
        while (low <= high)
        {
            int mid = low + (high - low) / 2;

            if (nums[mid] != nums[mid - 1] && nums[mid] != nums[mid + 1])
                return nums[mid];

            if (mid % 2 == 1 && nums[mid] == nums[mid - 1] || (mid % 2 == 0 && nums[mid] == nums[mid + 1]))
            {
                low = mid + 1;
            }
            else
            {
                high = mid - 1;
            }
        }

        return -1;
    }
};
// Time Complexity: O(logN), N = size of the given array.
// Reason: We are basically using the Binary Search algorithm.
// Space Complexity: O(1) as we are not using any extra space.