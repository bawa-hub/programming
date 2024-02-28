// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

class Solution
{
public:
    bool search(vector<int> &nums, int target)
    {
        int l = 0;
        int r = nums.size() - 1;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            if (nums[mid] == target)
                return true;
            // with duplicates we can have this contdition, just update left & right
            if ((nums[l] == nums[mid]) && (nums[r] == nums[mid]))
            {
                l++;
                r--;
            }
            // first half
            // first half is in order
            else if (nums[l] <= nums[mid])
            {
                // target is in first  half
                if ((nums[l] <= target) && (nums[mid] > target))
                    r = mid - 1;
                else
                    l = mid + 1;
            }
            // second half
            // second half is order
            // target is in second half
            else
            {
                if ((nums[mid] < target) && (nums[r] >= target))
                    l = mid + 1;
                else
                    r = mid - 1;
            }
        }
        return false;
    }
};

// Time Complexity: O(logN) for the best and average case. O(N/2) for the worst case. Here, N = size of the given array.
// Reason: In the best and average scenarios, the binary search algorithm is primarily utilized and hence the time complexity is O(logN). However, in the worst-case scenario, where all array elements are the same but not the target (e.g., given array = {3, 3, 3, 3, 3, 3, 3}), we continue to reduce the search space by adjusting the low and high pointers until they intersect. This worst-case situation incurs a time complexity of O(N/2).

// Space Complexity: O(1)
// Reason: We have not used any extra data structures, this makes space complexity, even in the worst case as O(1).