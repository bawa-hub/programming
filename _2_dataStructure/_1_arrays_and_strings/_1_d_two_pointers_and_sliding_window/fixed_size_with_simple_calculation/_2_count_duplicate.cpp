// https://leetcode.com/problems/contains-duplicate-ii/

class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        int i=0,j=0,n=nums.size();
        unordered_map<int, int> mp;

        while(j<n) {
            mp[nums[j]]++;

            if(mp[nums[j]]>1) return true;

            if(j-i==k) {
                mp[nums[i]]--;
                i++;
            }

            j++;
        }
        return false;
    }
};