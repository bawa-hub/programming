
// https://leetcode.com/problems/majority-element/

// moore voting algorithm
class Solution
{
public:
    int majorityElement(vector<int> &nums)
    {
        int count = 0;
        int candidate = 0;

        for (int num : nums)
        {
            if (count == 0)
            {
                candidate = num;
            }
            if (num == candidate)
                count += 1;
            else
                count -= 1;
        }

        return candidate;
    }
    // TC: O(n)
    // SC: O(1)
};
