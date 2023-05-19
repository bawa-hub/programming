// https://leetcode.com/problems/fruit-into-baskets/

// sliding window (by me)
class Solution {
public:
    int totalFruit(vector<int>& fruits) {
        int i=0,j=0,maxi = 0;
        int n = fruits.size();
        unordered_map<int, int> mp;

        while(j<n) {
          mp[fruits[j]]++;
          if(mp.size()<=2) {
              maxi = max(maxi, j-i+1);
          }

         if(mp.size()>2) {
             mp[fruits[i]]--;
             if(mp[fruits[i]]==0) mp.erase(fruits[i]);
             i++;
         }

          j++;
        }

        return maxi;
    }
};