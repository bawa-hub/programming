// https://leetcode.com/problems/longest-happy-prefix/description/

// Pattern
// Prefix hash comparison

// We want:
// prefix == suffix

// Example
// level

// prefix: le
// suffix: el
// Idea

// Compare hashes:
// hash(0,i) == hash(n-i-1,n-1)

// What You Learn:
// Prefix hashing
// Forward vs backward comparison
// String border detection

// Complexity
// O(n)

// Idea

// Compare
// prefix hash
// suffix hash

// We iterate length i.
// prefix = s[0..i]
// suffix = s[n-i-1..n-1]

class Solution {
public:
    string longestPrefix(string s) {

        const long long base = 131;
        const long long mod = 1e9+7;

        long long prefix=0;
        long long suffix=0;
        long long power=1;

        int n=s.size();
        int idx=0;

        for(int i=0;i<n-1;i++)
        {
            prefix=(prefix*base+s[i])%mod;

            suffix=(suffix+power*s[n-1-i])%mod;

            power=(power*base)%mod;

            if(prefix==suffix)
                idx=i+1;
        }

        return s.substr(0,idx);
    }
};