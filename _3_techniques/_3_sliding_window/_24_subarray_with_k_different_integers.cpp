// https://leetcode.com/problems/subarrays-with-k-different-integers/
// https://www.youtube.com/watch?v=akwRFY2eyXs

#include <vector>
#include <unordered_map>
using namespace std;


class Solution {
public:
    int subarraysWithKDistinct(vector<int>& A, int K) {
        return subarraysWithAtMostKDistinct(A, K) - subarraysWithAtMostKDistinct(A, K - 1);
    }

private:
    int subarraysWithAtMostKDistinct(vector<int>& s, int k) {
        unordered_map<int, int> mp;
        int i=0,j=0,cnt=0,res=0,n=s.size();
        
        while (j < n) {
            mp[s[j]]++;
            if (mp[s[j]] == 1) cnt++;

            while (cnt > k) {
                mp[s[i]]--;
                if (mp[s[i]] == 0) cnt--;
                i++;
            }
            
            res += (j-i+1);

            j++;
        }
        return res;
    }
};