// https://leetcode.com/problems/fruit-into-baskets/

// sliding window (by me)
class Solution {
public:
    int totalFruit(vector<int>& fruits) {
        int i=0,j=0,maxi=0;
        unordered_map<int, int> mp;

        while(j<fruits.size()) {
            mp[fruits[j]]++;
            while(mp.size()>2) {
               if(--mp[fruits[i]] == 0) mp.erase(fruits[i]);
               i++;
            }
            maxi = max(maxi, j-i+1);
            j++;
        }

        return maxi;
    }
};