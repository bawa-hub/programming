// https://leetcode.com/contest/biweekly-contest-104/problems/maximum-or/
class Solution {
public:
    long long maximumOr(vector<int>& nums, int k) {
        int n = nums.size();
        vector<long long> preSum(n+1, 0);
         vector<long long> suffSum(n+1, 0);
        long long res, pow = 1;
        for (int i = 0; i < k; i++) pow *= 2;
                 
        for (int i = 0; i < n; i++)
            preSum[i + 1] = preSum[i] | nums[i];
    
        for (int i = n - 1; i >= 0; i--) suffSum[i] = suffSum[i + 1] | nums[i];
    
        res = 0;
        for (int i = 0; i < n; i++) res = max(res, preSum[i] | (nums[i] * pow) | suffSum[i + 1]);
 
        return res;
    }
};