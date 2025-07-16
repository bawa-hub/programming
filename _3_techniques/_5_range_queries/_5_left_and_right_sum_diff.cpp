// https://leetcode.com/problems/left-and-right-sum-differences/description/

class Solution {
public:
    vector<int> leftRightDifference(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre(n), suff(n), res(n);

        int lsum=0;
        for(int i=0;i<nums.size();i++) {
            pre[i] = lsum;
            lsum+=nums[i];
        }

        int rsum=0;
        for(int i=nums.size()-1;i>=0;i--) {
            suff[i] = rsum;
            rsum+=nums[i];
        }

        for(int i=0;i<nums.size();i++) {
            res[i] = abs(pre[i]-suff[i]);
        }

        return res;
    }
};