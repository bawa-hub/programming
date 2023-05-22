// https://leetcode.com/problems/minimum-size-subarray-sum/

class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int i=0,j=0,n=nums.size(),sum=0,mini=INT_MAX;

        while(j<n) {
           sum+=nums[j];

           if(sum>=target) {
               mini = min(mini, j-i+1);
               while(sum>=target) {
                   sum-=nums[i];
                   i++;
                   if(sum>=target) mini = min(mini, j-i+1);
               }
           }

           j++;
        }

        if(mini == INT_MAX) return 0;

        return mini;
    }
};