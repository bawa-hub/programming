// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/

class Solution
{
public:
    int sumByD(vector<int> &nums, int div)
    {
        int n = nums.size(), sum = 0;
        for (int i = 0; i < n; i++)
        {
            sum += (ceil((double)nums[i] / (double)div));
        }
        return sum;
    }
    int smallestDivisor(vector<int> &nums, int threshold)
    {
        int l = 1, r = *max_element(nums.begin(), nums.end());
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            if (sumByD(nums, mid) <= threshold)
            {
                r = mid - 1;
            }
            else
            {
                l = mid + 1;
            }
        }
        return l;
    }
};