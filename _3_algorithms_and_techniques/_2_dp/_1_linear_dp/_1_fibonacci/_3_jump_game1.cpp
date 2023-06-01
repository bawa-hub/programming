// https://leetcode.com/problems/jump-game/

// memoization
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

// tabulation
bool canJump(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, -1);
        dp[n-1] = 1; //base case;
        
        for(int idx = n-2; idx >= 0; idx--) {
            if(nums[idx] == 0) {
                dp[idx] = false;
                continue;   
            }
            
            int flag = 0;
            int reach = idx + nums[idx];
            for(int jump=idx + 1; jump <= reach; jump++) {
                if(jump < nums.size() && dp[jump]) {
                    dp[idx] = true;
                    flag = 1;  
                    break;
                }
            }
            if(flag == 1) 
                continue;
           
            dp[idx] = false;
			
        }
        return dp[0]; 
    }