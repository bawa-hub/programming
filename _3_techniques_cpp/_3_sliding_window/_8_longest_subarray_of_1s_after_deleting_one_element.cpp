// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int i=0,j=0,n=nums.size(),mxCnt=0,delCnt=0,maxi=INT_MIN;

        while(j<n) {
           if(nums[j]==1) {
               mxCnt++;
           } else if(nums[j]==0) {
             delCnt++;
           }

            if(delCnt>1) {
                while(delCnt>1) {
                    if(nums[i]==0) {
                        delCnt--;
                    } else mxCnt--;
                    i++;
                }
            }

             maxi = max(maxi, mxCnt);

            j++;
        }

        if(delCnt==0) return maxi-1;

        return maxi;
    }
};