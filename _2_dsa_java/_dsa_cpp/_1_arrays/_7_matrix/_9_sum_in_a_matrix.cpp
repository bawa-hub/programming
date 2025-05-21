// https://leetcode.com/problems/sum-in-a-matrix/description/

class Solution {
public:
    int matrixSum(vector<vector<int>>& nums) {
        for(int i=0;i<nums.size();i++) {
            sort(nums[i].begin(), nums[i].end());
        }
        
        long sum = 0;
        
        vector<int> me;
        int s = nums[0].size()-1;
        for(int j=s;j>=0;j--) {
            for(int i=0;i<nums.size();i++){
            me.push_back(nums[i][j]);
        }
            int ele = maxi(me);
            sum+=ele;
            me.clear();
            
        }
        
        return sum;
        
        
        
    }
    
    int maxi(vector<int> &v) {
        int m = INT_MIN;
        for(int i=0;i<v.size();i++) {
            m = max(m, v[i]);
        }
        if(m==INT_MIN) return 0;
        return m;
    }
};