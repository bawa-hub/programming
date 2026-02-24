// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/

class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
        // support variables
        int len = nums.size(), freqs[101] = {};
        vector<int> res(len);
        // populating freqs
        for (int n: nums) freqs[n]++;
        for (int i = 0; i < len; i++) {
            res[i] = accumulate(freqs, freqs + nums[i], 0);
        }
        return res;
    }
};