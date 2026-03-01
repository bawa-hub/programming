// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/

class Solution {
public:
    int minimumDifference(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());

        int i=0,j=0;
        int mini = INT_MAX;
        while(j<nums.size()) {
            if(j-i+1==k) {
             int diff = nums[j]-nums[i];
             mini = min(mini, diff);
             i++;
            }
            j++;
        }
        return mini;
    }
};