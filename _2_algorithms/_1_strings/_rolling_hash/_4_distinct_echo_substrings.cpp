#include<vector>
#include<unordered_set>
using namespace std;

// Distinct Echo Substrings
// https://leetcode.com/problems/distinct-echo-substrings/description/

// Problem
// Find number of substrings that appear twice consecutively.

// Example:

// abcabcabc

// Echo substrings:

// abcabc
// bcabca
// cabcab

// Idea

// Substring length = 2k
// first half == second half

// We compare hashes:
// hash(l, l+k-1) == hash(l+k, l+2k-1)

class Solution {
public:
    int distinctEchoSubstrings(string text) {

        int n=text.size();
        const long long base=31;
        const long long mod=1e9+7;

        vector<long long> power(n+1);
        vector<long long> prefix(n+1);

        power[0]=1;

        for(int i=1;i<=n;i++)
            power[i]=(power[i-1]*base)%mod;

        for(int i=0;i<n;i++)
            prefix[i+1]=(prefix[i]*base+(text[i]-'a'+1))%mod;

        auto getHash=[&](int l,int r){
            return (prefix[r+1]-prefix[l]*power[r-l+1]%mod+mod)%mod;
        };

        unordered_set<long long> seen;

        for(int len=1;len*2<=n;len++)
        {
            for(int i=0;i+2*len<=n;i++)
            {
                if(getHash(i,i+len-1)==getHash(i+len,i+2*len-1))
                {
                    seen.insert(getHash(i,i+2*len-1));
                }
            }
        }

        return seen.size();
    }
};
