

// Rabin–Karp Algorithm
// The Rabin–Karp algorithm is a string searching algorithm used to find a pattern inside a text using hashing.
// Instead of comparing strings character-by-character, we compare hash values.
// If hashes match → then verify characters.

// 1. Problem Rabin–Karp Solves

// Find all occurrences of pattern P in text T.

// Example
// Text:    ABCCDDAEFG
// Pattern: CDD

// Expected match:
// ABCCDDAEFG
//    CDD

// Index = 3

// Rabin–Karp Key Idea
// Instead of comparing characters:

// compare hash(pattern)
// with
// hash(substring)

// Example
// hash("CDD") == hash("CDD")
// If equal → then verify characters.

// Sliding Window Concept

// Example
// Text = ABCDEFG
// Pattern length = 3

// Windows

// ABC
// BCD
// CDE
// DEF
// EFG

// Instead of recomputing hash for every substring:
// hash(ABC)
// hash(BCD)
// hash(CDE)

// we update the hash using rolling hash.


// Rabin–Karp Algorithm Steps
// 1 compute hash(pattern)
// 2 compute hash(first window of text)
// 3 for each window

//     if hashes match
//         verify characters

//     slide window
//     update hash

#include <bits/stdc++.h>
using namespace std;

vector<int> rabinKarp(string text,string pattern)
{
    const long long base = 256;
    const long long mod = 1e9+7;

    int n=text.size();
    int m=pattern.size();

    long long patternHash=0;
    long long windowHash=0;
    long long power=1;

    vector<int> ans;

    for(int i=0;i<m-1;i++)
        power=(power*base)%mod;

    for(int i=0;i<m;i++)
    {
        patternHash=(patternHash*base+pattern[i])%mod;
        windowHash=(windowHash*base+text[i])%mod;
    }

    for(int i=0;i<=n-m;i++)
    {
        if(patternHash==windowHash)
        {
            if(text.substr(i,m)==pattern)
                ans.push_back(i);
        }

        if(i<n-m)
        {
            windowHash=
            (windowHash-text[i]*power%mod+mod)%mod;

            windowHash=
            (windowHash*base+text[i+m])%mod;
        }
    }

    return ans;
}

// Average case - O(n + m)
// Worst case - O(n * m)



// Problem 1 — Find the Index of the First Occurrence in a String
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string
class Solution {
public:
    int strStr(string text, string pattern) {

        int n = text.size();
        int m = pattern.size();

        const long long base = 256;
        const long long mod = 1e9+7;

        long long patternHash = 0;
        long long windowHash = 0;
        long long power = 1;

        for(int i=0;i<m-1;i++)
            power = (power*base)%mod;

        for(int i=0;i<m;i++)
        {
            patternHash = (patternHash*base + pattern[i])%mod;
            windowHash = (windowHash*base + text[i])%mod;
        }

        for(int i=0;i<=n-m;i++)
        {
            if(patternHash==windowHash)
            {
                if(text.substr(i,m)==pattern)
                    return i;
            }

            if(i<n-m)
            {
                windowHash =
                (windowHash - text[i]*power%mod + mod)%mod;

                windowHash =
                (windowHash*base + text[i+m])%mod;
            }
        }

        return -1;
    }
};

// Problem 2 — All Pattern Matches
// Example
// Text    = AABAACAADAABAABA
// Pattern = AABA
// Answer
// 0 9 12

vector<int> rabinKarp(string text,string pattern)
{
    vector<int> ans;

    int n=text.size();
    int m=pattern.size();

    const long long base=256;
    const long long mod=1e9+7;

    long long pHash=0;
    long long wHash=0;
    long long power=1;

    for(int i=0;i<m-1;i++)
        power=(power*base)%mod;

    for(int i=0;i<m;i++)
    {
        pHash=(pHash*base+pattern[i])%mod;
        wHash=(wHash*base+text[i])%mod;
    }

    for(int i=0;i<=n-m;i++)
    {
        if(pHash==wHash)
        {
            if(text.substr(i,m)==pattern)
                ans.push_back(i);
        }

        if(i<n-m)
        {
            wHash=(wHash-text[i]*power%mod+mod)%mod;
            wHash=(wHash*base+text[i+m])%mod;
        }
    }

    return ans;
}