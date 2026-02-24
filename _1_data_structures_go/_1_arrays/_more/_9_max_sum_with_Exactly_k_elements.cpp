// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/description/

class Solution {
public:
    int maximizeSum(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());
        
        int m = nums[nums.size()-1];
        int sum = 0;
          sum+=m;
        
        for(int i=0;i<k-1;i++) {
            m++;
            sum+=m;
        }
        
        return sum;
    }
};