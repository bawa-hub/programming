// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/

class Solution {
public:
    int countGoodSubstrings(string s) {
        int i=0,j=0,cnt=0,n = s.size();
        unordered_map<char, int> mp;

        while(j<n) {
            mp[s[j]]++;
            if(j-i+1==3) {
               if(mp.size()==3) cnt++;
               mp[s[i]]--;
               if(mp[s[i]]==0) mp.erase(s[i]);
               i++;
            }
            j++;
        }
        return cnt;
    }
};