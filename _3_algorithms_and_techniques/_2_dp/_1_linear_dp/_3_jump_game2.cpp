// https://leetcode.com/problems/jump-game-ii/
// https://www.geeksforgeeks.org/minimum-number-of-jumps-to-reach-end-of-a-given-array/
// https://leetcode.com/problems/jump-game-ii/solutions/1192401/easy-solutions-w-explanation-optimizations-from-brute-force-to-dp-to-greedy-bfs/
// recursive
class Solution {
public:
    int jump(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, -1);
        return f(0, nums, dp);
    }

    int f(int idx, vector<int>& nums,vector<int> &dp) {
        if(idx>=nums.size()-1) return 0;
        if(dp[idx]!=-1) return dp[idx];
        
        int mini = 1e6;
        for(int i=1;i<=nums[idx];i++) {
            if(i+idx<nums.size()) {
             mini = min(mini, 1+f(idx+i, nums, dp));
            }
        }

        return dp[idx] = mini;
    }
};

