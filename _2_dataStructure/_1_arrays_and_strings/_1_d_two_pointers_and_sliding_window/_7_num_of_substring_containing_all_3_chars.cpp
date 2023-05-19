// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

// by me
class Solution {
public:
    int numberOfSubstrings(string s) {
        
        int i=0,j=0,n=s.size(),cnt=0;
        unordered_map<char, int> mp;

        while(j<n) {
         mp[s[j]]++;

         if(mp.size()==3) {
             while(mp.size()==3) {
             cnt+=(s.size()-1-j+1);
             mp[s[i]]--;
             if(mp[s[i]]==0) mp.erase(s[i]);
             i++;
             }
         }

         j++;

        }

        return cnt;
    }
};