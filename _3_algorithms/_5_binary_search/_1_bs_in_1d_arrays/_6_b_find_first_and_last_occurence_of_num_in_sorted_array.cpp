// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// https://practice.geeksforgeeks.org/problems/first-and-last-occurrences-of-x3116/1
// https://www.geeksforgeeks.org/find-first-and-last-positions-of-an-element-in-a-sorted-array/

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