// https://leetcode.com/problems/3sum/

#include <bits/stdc++.h>
using namespace std;

// brute force
vector<vector<int>> threeSum(vector<int> &nums)
{
    vector<vector<int>> ans;
    vector<int> temp;
    int i, j, k;
    for (i = 0; i < nums.size() - 2; i++)
    {
        for (j = i + 1; j < nums.size() - 1; j++)
        {
            for (k = j + 1; k < nums.size(); k++)
            {
                temp.clear();
                if (nums[i] + nums[j] + nums[k] == 0)
                {
                    temp.push_back(nums[i]);
                    temp.push_back(nums[j]);
                    temp.push_back(nums[k]);
                }
                if (temp.size() != 0)
                    ans.push_back(temp);
            }
        }
    }

    return ans;
}
// Time Complexity : O(n^3)   // 3 loops
// Space Complexity : O(3*k)  // k is the no.of triplets

// optimized approach
vector<vector<int>> threeSum(vector<int> &num)
{
    vector<vector<int>> res;
    sort(num.begin(), num.end());

    // moves for a
    for (int i = 0; i < (int)(num.size()) - 2; i++)
    {

        if (i == 0 || (i > 0 && num[i] != num[i - 1]))
        {

            int lo = i + 1, hi = (int)(num.size()) - 1, sum = 0 - num[i];

            while (lo < hi)
            {
                if (num[lo] + num[hi] == sum)
                {

                    vector<int> temp;
                    temp.push_back(num[i]);
                    temp.push_back(num[lo]);
                    temp.push_back(num[hi]);
                    res.push_back(temp);

                    while (lo < hi && num[lo] == num[lo + 1])
                        lo++;
                    while (lo < hi && num[hi] == num[hi - 1])
                        hi--;

                    lo++;
                    hi--;
                }
                else if (num[lo] + num[hi] < sum)
                    lo++;
                else
                    hi--;
            }
        }
    }
    return res;
}

// Time Complexity : O(n^2)
// Space Complexity : O(3*k)  // k is the no.of triplets.

int main()
{
    vector<int> arr{-1, 0, 1, 2, -1, -4};
    vector<vector<int>> ans;
    ans = threeSum(arr);
    cout << "The triplets are as follows: " << endl;
    for (int i = 0; i < ans.size(); i++)
    {
        for (int j = 0; j < ans[i].size(); j++)
        {
            cout << ans[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}