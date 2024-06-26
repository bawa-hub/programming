// https://leetcode.com/problems/sum-of-unique-elements/description/

class Solution {
public:
    int sumOfUnique(vector<int>& nums) {
        unordered_map<int, int> mp;
        for(int i=0;i<nums.size();i++) {
            mp[nums[i]]++;
        }
        int sum = 0;
        for(auto i : mp) {
            if(i.second==1) sum+=i.first;
        }

        return sum;
    }
};