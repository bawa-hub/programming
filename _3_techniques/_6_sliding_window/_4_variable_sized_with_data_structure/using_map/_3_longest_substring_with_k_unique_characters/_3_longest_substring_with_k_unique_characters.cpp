// https://practice.geeksforgeeks.org/problems/longest-k-unique-characters-substring0853/1

int longestKSubstr(string s, int k) {
     int i=0,j=0,maxi=-1;
     int n = s.size();
     unordered_map<char, int> mp;
     
     while(j<n) {
         mp[s[j]]++;
         
         if(mp.size() == k) {
             maxi = max(maxi, j-i+1);
         }
         
         if(mp.size() > k) {
             while(mp.size()>k) {
                 mp[s[i]]--;
                 if(mp[s[i]]==0) mp.erase(s[i]);
                 i++;
             }
         }
         
        j++;
     }
     
     return maxi;
    }