// https://leetcode.com/problems/permutations/
// https://takeuforward.org/data-structure/print-all-permutations-of-a-string-array/

// https://www.geeksforgeeks.org/permutations-and-combinations/
// Selecting the data or objects from a certain group is said to be permutation
// number of ways to pick r things out of n different things in a specific order and replacement is not allowed 
// total permutations == nPr = n!/(n-r)! (n factorial)


#include <bits/stdc++.h>
using namespace std;

// extra space
class Solution
{
private:
    void recurPermute(vector<int> &ds, vector<int> &nums, vector<vector<int>> &ans, int freq[])
    {
        if (ds.size() == nums.size())
        {
            ans.push_back(ds);
            return;
        }
        for (int i = 0; i < nums.size(); i++)
        {
            if (!freq[i])
            {
                ds.push_back(nums[i]);
                freq[i] = 1;
                recurPermute(ds, nums, ans, freq);
                freq[i] = 0;
                ds.pop_back();
            }
        }
    }

public:
    vector<vector<int>> permute(vector<int> &nums)
    {
        vector<vector<int>> ans;
        vector<int> ds;
        int freq[nums.size()];
        for (int i = 0; i < nums.size(); i++)
            freq[i] = 0;
        recurPermute(ds, nums, ans, freq);
        return ans;
    }
};
// TC: O(n*n!)
// SC: O(n + n)

// space optimized
class Solution
{
private:
    void recurPermute(int index, vector<int> &nums, vector<vector<int>> &ans)
    {
        if (index == nums.size())
        {
            ans.push_back(nums);
            return;
        }
        for (int i = index; i < nums.size(); i++)
        {
            swap(nums[index], nums[i]);
            recurPermute(index + 1, nums, ans);
            swap(nums[index], nums[i]);
        }
    }

public:
    vector<vector<int>> permute(vector<int> &nums)
    {
        vector<vector<int>> ans;
        recurPermute(0, nums, ans);
        return ans;
    }
};
// TC: O(n*n!)
// SC: O(n)


int main()
{
    Solution obj;
    vector<int> v{1, 2, 3};
    vector<vector<int>> sum = obj.permute(v);
    cout << "All Permutations are " << endl;
    for (int i = 0; i < sum.size(); i++)
    {
        for (int j = 0; j < sum[i].size(); j++)
            cout << sum[i][j] << " ";
        cout << endl;
    }
}