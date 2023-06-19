// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// https://practice.geeksforgeeks.org/problems/first-and-last-occurrences-of-x3116/1

// lower_bound and upper_bound
class Solution
{
public:
    vector<int> searchRange(vector<int> &nums, int target)
    {
        int n = nums.size();
        int lb = lowerBound(nums, n, target);
        if (lb == n || nums[lb] != target)
            return {-1, -1};
        return {lb, upperBound(nums, n, target) - 1};
    }

    int upperBound(vector<int> arr, int n, int x)
    {
        int l = 0, r = n - 1;
        int ans = n;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            if (arr[mid] > x)
            {
                ans = mid;
                r = mid - 1;
            }
            else
            {
                l = mid + 1;
            }
        }

        return ans;
    }

    int lowerBound(vector<int> arr, int n, int x)
    {
        int l = 0, r = n - 1;
        int ans = n;

        while (l <= r)
        {
            int mid = l + (r - l) / 2;

            if (arr[mid] >= x)
            {
                ans = mid;
                r = mid - 1;
            }
            else
            {
                l = mid + 1;
            }
        }

        return ans;
    }
};
// TC: O(2*log(n));
// SC: O(1)

// simple binary search
class Solution
{
public:
    vector<int> searchRange(vector<int> &nums, int target)
    {
        vector<int> res;
        int f = first(nums.size(), target, nums);
        res.push_back(f);
        int l = last(nums.size(), target, nums);
        res.push_back(l);
        return res;
    }

    int last(int n, int key, vector<int> &v)
    {
        int start = 0;
        int end = n - 1;
        int res = -1;

        while (start <= end)
        {
            int mid = start + (end - start) / 2;
            if (v[mid] == key)
            {
                res = mid;
                start = mid + 1;
            }
            else if (key < v[mid])
            {
                end = mid - 1;
            }
            else
            {
                start = mid + 1;
            }
        }
        return res;
    }

    int first(int n, int key, vector<int> &v)
    {
        int start = 0;
        int end = n - 1;
        int res = -1;

        while (start <= end)
        {
            int mid = start + (end - start) / 2;
            if (v[mid] == key)
            {
                res = mid;
                end = mid - 1;
            }
            else if (key < v[mid])
            {
                end = mid - 1;
            }
            else
            {
                start = mid + 1;
            }
        }
        return res;
    }
};