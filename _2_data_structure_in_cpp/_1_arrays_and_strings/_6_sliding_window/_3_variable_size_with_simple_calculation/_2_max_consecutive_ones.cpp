// https://leetcode.com/problems/max-consecutive-ones-iii/

// by me
class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int i=0,j=0,n=nums.size(),maxi=INT_MIN,cnt=0, zeroToOne=0;

        while(j<n) {
            if(nums[j]==1) cnt++;
            else {
                if(zeroToOne<k) {
                  cnt++;
                  zeroToOne++;
                } else zeroToOne++;
            }

            // if(zeroToOne == k) {
                maxi = max(maxi, cnt);


            if(zeroToOne>k) {
              while(zeroToOne>k) {
                  if(nums[i]==1) cnt--;
                  else {
                      zeroToOne--;
                  }
                  i++;
              }
            }
            j++;
        }

        if(maxi==INT_MIN) return 0;

        return maxi;
    }
};