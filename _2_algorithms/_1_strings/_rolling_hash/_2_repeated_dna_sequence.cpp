


// Pattern
// Sliding Window Hashing (Rabin–Karp)

// Problem
// https://leetcode.com/problems/repeated-dna-sequences/description/
// Find all substrings of length 10 that appear more than once.

// Example
// AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT

// Output
// AAAAACCCCC
// CCCCCAAAAA

// What You Learn
// Rolling hash sliding window
// Hash set usage
// Avoid recomputation

// Complexity
// O(n)

// This problem builds the foundation of Rabin-Karp thinking.

// Key Observation

// Length is fixed = 10
// So we can use rolling hash sliding window.
// window size = 10
// Each window:
// s[i..i+9]

// Idea
// Use two sets:
// seen
// duplicate

// Steps:
// 1 compute hash
// 2 if already seen → duplicate
// 3 otherwise add to seen

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {

        int n = s.size();
        int k = 10;

        if(n < k) return {};

        const long long base = 131;
        const long long mod1 = 1e9+7;
        const long long mod2 = 1e9+9;

        long long hash1 = 0, hash2 = 0;
        long long power1 = 1, power2 = 1;

        for(int i=0;i<k;i++)
        {
            hash1 = (hash1*base + s[i]) % mod1;
            hash2 = (hash2*base + s[i]) % mod2;

            if(i < k-1)
            {
                power1 = (power1*base) % mod1;
                power2 = (power2*base) % mod2;
            }
        }

        unordered_set<unsigned long long> seen;
        unordered_set<unsigned long long> added;

        auto combine=[&](long long a,long long b){
            return ((unsigned long long)a<<32) | b;
        };

        seen.insert(combine(hash1,hash2));

        vector<string> ans;

        for(int i=k;i<n;i++)
        {
            hash1 = (hash1 - s[i-k]*power1 % mod1 + mod1) % mod1;
            hash1 = (hash1*base + s[i]) % mod1;

            hash2 = (hash2 - s[i-k]*power2 % mod2 + mod2) % mod2;
            hash2 = (hash2*base + s[i]) % mod2;

            unsigned long long h = combine(hash1,hash2);

            if(seen.count(h))
            {
                if(!added.count(h))
                {
                    ans.push_back(s.substr(i-k+1,k));
                    added.insert(h);
                }
            }
            else
            {
                seen.insert(h);
            }
        }

        return ans;
    }
};