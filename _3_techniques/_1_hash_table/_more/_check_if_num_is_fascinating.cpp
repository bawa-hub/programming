// https://leetcode.com/problems/check-if-the-number-is-fascinating/description/

class Solution {
public:
    bool isFascinating(int n) {
        int f = 2*n;
        int l = 3*n;
        string s = to_string(n)+to_string(f)+to_string(l);
        return check(s);
    }
    
    bool check(string s) {
        if(s.size()<0) return false;
        unordered_map<char,int> mp;
        for(int i=0;i<s.size();i++) {
            mp[s[i]]++;
        }
        
        for(auto i : mp) {
            if(i.first=='0') return false;
            if(i.second>1) return false;
        }
        return true;
    }
};