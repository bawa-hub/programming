// https://leetcode.com/problems/missing-number/

class Solution
{
public:
    int missingNumber(vector<int> &nums)
    {
        int size = nums.size();

        int total = size * (size + 1) / 2;
        for (int i = 0; i < size; i++)
            total -= nums[i];

        return total;
    }
};