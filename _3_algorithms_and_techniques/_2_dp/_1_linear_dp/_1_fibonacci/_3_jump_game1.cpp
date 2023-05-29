// https://leetcode.com/problems/jump-game/

class Solution {
public:
    bool canJump(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, -1);
        if(f(0, nums, dp)==true) return true;
        return false;
    }

    bool f(int idx, vector<int>& nums, vector<int> &dp) {
        if(idx==nums.size()-1) return true;
        if(idx>=nums.size()) return false;

        if(dp[idx] != -1) return dp[idx];

        for(int i=1;i<=nums[idx];i++) {
            if(i+idx<nums.size()) {
                if(f(idx+i, nums, dp)==true) return true;
            }
        }

        return dp[idx] = false;
    }
};