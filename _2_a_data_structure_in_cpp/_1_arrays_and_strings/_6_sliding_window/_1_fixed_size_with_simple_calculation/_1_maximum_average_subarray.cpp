// https://leetcode.com/problems/maximum-average-subarray-i/

class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        int i=0,j=0,n=nums.size();

       double sum=0;
       double maxi = INT_MIN;
        while(j<n) {
         sum+=nums[j];

         if(j-i+1==k) {
             double avg = sum/k;
             maxi = max(maxi, avg);
             sum-=nums[i];
             i++;
         }

         j++;
        }

        return maxi;
    }
};