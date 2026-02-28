// https://leetcode.com/problems/permutations-ii/

#include <vector>
using namespace std;

class Solution
{
public:
    void backtrack(vector<int> &nums, vector<bool> &used, vector<int> &path, vector<vector<int>> &result)
    {
        if (path.size() == nums.size())
        {
            result.push_back(path);
            return;
        }

        for (int i = 0; i < nums.size(); ++i)
        {
            if (used[i])
                continue;

            // 🚨 Skip duplicate numbers at the same level
            if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])
                continue;

            used[i] = true;
            path.push_back(nums[i]);

            backtrack(nums, used, path, result);

            path.pop_back();
            used[i] = false;
        }
    }

    vector<vector<int>> permuteUnique(vector<int> &nums)
    {
        sort(nums.begin(), nums.end()); // 🔑 sort to detect duplicates
        vector<vector<int>> result;
        vector<int> path;
        vector<bool> used(nums.size(), false);
        backtrack(nums, used, path, result);
        return result;
    }
};