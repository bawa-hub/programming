// https://leetcode.com/problems/split-a-string-in-balanced-strings/description/

class Solution {
public:
    int balancedStringSplit(string s) {
        int n=s.size();
        int cnt=0,ans=0;

        for(int i=0;i<n;i++) {
            if(s[i]=='R') cnt++;
            else if(s[i]=='L') cnt--;
            if(cnt==0) ans++;
        }

        return ans;
    }
};