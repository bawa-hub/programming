// https://leetcode.com/problems/max-consecutive-ones-iii/

class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int i=0,j=0,cnt=0,maxi=0,flip=0;

        while(j<nums.size()) {
             if(nums[j]==1) cnt++;
             else if (nums[j]==0 && flip++ <= k) cnt++;

             while(flip > k) {
                 if(nums[i]==0) flip--;
                 i++;
                 cnt--;
             }

             maxi = max(maxi, j-i+1);
             j++;
        }

        return maxi;
    }
};