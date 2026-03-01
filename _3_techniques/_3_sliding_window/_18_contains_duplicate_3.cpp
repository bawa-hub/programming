// https://leetcode.com/problems/contains-duplicate-iii
// https://leetcode.com/problems/contains-duplicate-iii/solutions/2891742/c-solution-unsing-map-and-lower-bound-sliding-window/

class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int indexDiff, int valueDiff) {
        int i=0;
        map<int,int> mp;
        int n=nums.size();
        for(int j=0;j<n;j++){
            auto val=mp.lower_bound(nums[j]);
            if(val!=mp.end() and (val->first-nums[j])<=valueDiff){
                return true;
            }
            if(val!=mp.begin()){
                val--;
                if(abs(val->first-nums[j])<=valueDiff){
                    return true;
                }
            }
            mp[nums[j]]++;
            if((j-i)==indexDiff){
                mp[nums[i]]--;
                if(mp[nums[i]]==0){
                    mp.erase(nums[i]);
                }
                i++;
            }
        }
        return false;
    }
};

// using multiset
class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) 
    {
        int n = nums.size();
        multiset<int> ms;     //to store window elements in sorted order
        
		int i=0, j=0;
        while(j<n)
        {
            auto up = ms.upper_bound(nums[j]);
            if((up != ms.end() and *up-nums[j] <= t) || (up != ms.begin() and nums[j] - *(--up) <= t))
                return true;
            ms.insert(nums[j]);
            
            if(ms.size() == k+1)
            {
                ms.erase(nums[i]);
                i++;
            }
            j++;
        }
        return false;
    }
};