// https://leetcode.com/problems/search-in-rotated-sorted-array/

using namespace std;

// linear search
int search(vector<int> &nums, int target)
{
    for (int i = 0; i < nums.size(); i++)
    {
        if (nums[i] == target)
            return i;
    }
    return -1;
}

// Time Complexity : O(N)
// Reason: We have to iterate through the entire array to check if the target is present in the array.
// Space Complexity: O(1)
// Reason: We have not used any extra data structures, this makes space complexity, even in the worst case as O(1).

// binary search
int search(vector<int> &nums, int target)
{
    int low = 0, high = nums.size() - 1;

    while (low <= high)
    {
        int mid = (low + high) >> 1;
        if (nums[mid] == target)
            return mid;

        // if left side is sorted
        if (nums[low] <= nums[mid])
        {
            if (nums[low] <= target && nums[mid] >= target)
                high = mid - 1;
            else
                low = mid + 1;
        }
        // right half is sorted
        else
        {
            if (nums[mid] <= target && target <= nums[high])
                low = mid + 1;
            else
                high = mid - 1;
        }
    }
    return -1;
}

// Time Complexity: O(log(N))
// Reason: We are performing a binary search, this turns time complexity to O(log(N)) where N is the size of the array.
// Space Complexity: O(1)
// Reason: We do not use any extra data structure, this turns space complexity to O(1).
