
// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
// https://www.lintcode.com/problem/386/
// https://www.codingninjas.com/codestudio/problems/distinct-characters_2221410

int atmostK(string & s, int k) {
        unordered_map<char, int> mp;
        int i=0,j=0,cnt=0,maxi=INT_MIN,n=s.size();
        
        while (j < n) {
            mp[s[j]]++;
            if (mp[s[j]] == 1) cnt++;

            while (cnt > k) {
                mp[s[i]]--;
                if (mp[s[i]] == 0) cnt--;
                i++;
            }
            
            maxi = max(maxi, j-i+1);

            j++;
        }
        return maxi;
    }