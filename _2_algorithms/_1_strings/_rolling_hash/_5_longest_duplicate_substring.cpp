#include <unordered_set>

using namespace std;


// Longest Duplicate Substring
// https://leetcode.com/problems/longest-duplicate-substring/description/

// Problem
// Given string s, return longest substring that appears at least twice.

// Example

// Input
// banana

// Output
// ana

// Key Idea

// We combine
// Binary Search + Rolling Hash


class Solution {
public:
    string longestDupSubstring(string s) {

        int n = s.size();

        const long long base = 31;
        const long long mod1 = 1e9+7;
        const long long mod2 = 1e9+9;

        vector<long long> p1(n+1), p2(n+1);
        vector<long long> pow1(n+1), pow2(n+1);

        pow1[0]=pow2[0]=1;

        for(int i=1;i<=n;i++){
            pow1[i]=(pow1[i-1]*base)%mod1;
            pow2[i]=(pow2[i-1]*base)%mod2;
        }

        for(int i=0;i<n;i++){
            p1[i+1]=(p1[i]*base+(s[i]-'a'+1))%mod1;
            p2[i+1]=(p2[i]*base+(s[i]-'a'+1))%mod2;
        }

        auto getHash=[&](int l,int r){

            long long h1 =
            (p1[r+1]-p1[l]*pow1[r-l+1]%mod1+mod1)%mod1;

            long long h2 =
            (p2[r+1]-p2[l]*pow2[r-l+1]%mod2+mod2)%mod2;

            return ((unsigned long long)h1<<32)|h2;
        };

        int left=1,right=n;
        string ans="";

        while(left<=right){

            int mid=(left+right)/2;

            unordered_set<unsigned long long> seen;

            bool found=false;

            for(int i=0;i+mid<=n;i++){

                auto h=getHash(i,i+mid-1);

                if(seen.count(h)){
                    ans=s.substr(i,mid);
                    found=true;
                    break;
                }

                seen.insert(h);
            }

            if(found) left=mid+1;
            else right=mid-1;
        }

        return ans;
    }
};
// TC: O(n log n)
