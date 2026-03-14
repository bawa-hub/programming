


// Find Substring With Given Hash Value
// https://leetcode.com/problems/find-substring-with-given-hash-value/description/

// Pattern
// Reverse rolling hash

// Instead of sliding left → right
// We slide right → left

// This problem tests deep understanding of rolling hash math.

// What You Learn
// Reverse rolling hash
// Mod arithmetic tricks
// Polynomial hashing control


// Key Idea

// Compute hash from right → left.
// Because formula:
// hash =
// s[i]*power^0
// + s[i+1]*power^1

// So sliding window works better backward.

class Solution {
public:
    string subStrHash(string s,int power,int mod,
                      int k,int hashValue) {

        int n=s.size();

        long long hash=0;
        long long p=1;

        int start=0;

        for(int i=0;i<k;i++)
            p=(p*power)%mod;

        for(int i=n-1;i>=0;i--)
        {
            hash=(hash*power+(s[i]-'a'+1))%mod;

            if(i+k<n)
            {
                hash=(hash-(s[i+k]-'a'+1)*p%mod+mod)%mod;
            }

            if(i+k<=n && hash==hashValue)
                start=i;
        }

        return s.substr(start,k);
    }
};