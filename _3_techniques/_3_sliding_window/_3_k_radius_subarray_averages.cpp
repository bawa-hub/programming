// https://leetcode.com/problems/k-radius-subarray-averages

class Solution
{
public:
    vector<int> getAverages(vector<int> &nums, int k)
    {
        int i = 0, j = 0, z = k, n = nums.size();
        long long sum = 0;
        vector<int> res(n, -1);
        while (j < n)
        {
            sum += nums[j];
            if (j - i == 2 * k)
            {
                long long avg = sum / (j - i + 1);
                res[z++] = avg;
                sum -= nums[i];
                i++;
            }
            j++;
        }
        return res;
    }
};