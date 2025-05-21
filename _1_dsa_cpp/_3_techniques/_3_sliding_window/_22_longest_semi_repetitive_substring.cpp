// https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/description/

class Solution {
public:
    int longestSemiRepetitiveSubstring(string s) {
        int i=0,cnt=0,n=s.size(),maxi=INT_MIN;
        
        unordered_map<char,int> mp;
        
        for(int j=1;j<n;j++) {
            if(s[j]==s[j-1]) cnt++;
            if(cnt>1) {
                while(cnt>1&&i<n-1) {
                    if(s[i]==s[i+1]) cnt--;
                    i++;
                }
            }
            maxi=max(maxi, j-i+1);
        }
        
        if(maxi==INT_MIN) return 1;
        return maxi;
        
        
    }
};