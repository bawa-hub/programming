// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

class Solution {
public:
    long long countSubarrays(vector<int>& nums, long long k) {
        long long i=0,j=0,n=nums.size(),sum=0,cnt=0,res=0;
        long long score;

        while(j<n) {
          cnt++;
          sum+=nums[j];
          score=sum*cnt;

          while(score>=k) {
              sum-=nums[i];
              cnt--;
              score=sum*cnt;
              i++;
          }

          res+=(j-i+1);

          j++;
        }

        return res;
    }
};