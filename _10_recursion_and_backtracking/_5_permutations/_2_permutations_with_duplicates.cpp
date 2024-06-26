// https://leetcode.com/problems/permutations-ii/

class Solution
{
private:
    void dfs(vector<int> &nums, vector<bool> &used, vector<int> &list, vector<vector<int>> &res)
    {
        if (list.size() == nums.size())
        {
            res.push_back(list);
            return;
        }
        for (int i = 0; i < nums.size(); i++)
        {
            if (used[i])
                continue;
            if (i > 0 && nums[i - 1] == nums[i] && !used[i - 1])
                continue;
            used[i] = true;
            list.push_back(nums[i]);
            dfs(nums, used, list, res);
            used[i] = false;
            list.pop_back();
        }
    }

public:
    vector<vector<int>> permuteUnique(vector<int> &nums)
    {
        vector<vector<int>> res;
        if (nums.empty())
            return res;
        vector<bool> used(nums.size(), false);
        vector<int> list;
        sort(nums.begin(), nums.end());
        dfs(nums, used, list, res);
        return res;
    }
};