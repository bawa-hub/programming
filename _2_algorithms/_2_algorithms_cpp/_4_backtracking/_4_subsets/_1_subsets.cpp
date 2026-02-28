// https://leetcode.com/problems/subsets/

class Solution
{
public:
    void recursion(int i, vector<int> &ds, vector<vector<int>> &res, vector<int> &nums, int N)
    {
        if (i == N)
        {
            res.push_back(ds);
            return;
        }

        // take
        ds.push_back(nums[i]);
        recursion(i + 1, ds, res, nums, N);
        ds.pop_back();

        // not take
        recursion(i + 1, ds, res, nums, N);
    }

    vector<vector<int>> subsets(vector<int> &nums)
    {
        vector<vector<int>> res;
        vector<int> ds;
        recursion(0, ds, res, nums, nums.size());
        return res;
    }
};