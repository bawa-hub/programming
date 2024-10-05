// https://leetcode.com/problems/longest-substring-without-repeating-characters/

#include <bits/stdc++.h>

using namespace std;

// brute force
int solve(string str)
{

    if (str.size() == 0)
        return 0;
    int maxans = INT_MIN;
    for (int i = 0; i < str.length(); i++) // outer loop for traversing the string
    {
        unordered_set<int> set;
        for (int j = i; j < str.length(); j++) // nested loop for getting different string starting with str[i]
        {
            if (set.find(str[j]) != set.end()) // if element if found so mark it as ans and break from the loop
            {
                maxans = max(maxans, j - i);
                break;
            }
            set.insert(str[j]);
        }
    }
    return maxans;
}
// Time Complexity: O( N2 )
// Space Complexity: O(N) where N is the size of HashSet taken for storing the elements

// better approach
int solve(string str)
{

    if (str.size() == 0)
        return 0;
    int maxans = INT_MIN;
    unordered_set<int> set;
    int l = 0;
    for (int r = 0; r < str.length(); r++) // outer loop for traversing the string
    {
        if (set.find(str[r]) != set.end()) // if duplicate element is found
        {
            while (l < r && set.find(str[r]) != set.end())
            {
                set.erase(str[l]);
                l++;
            }
        }
        set.insert(str[r]);
        maxans = max(maxans, r - l + 1);
    }
    return maxans;
}
// Time Complexity: O( 2*N ) (sometimes left and right both have to travel a complete array)
// Space Complexity: O(N) where N is the size of HashSet taken for storing the elements

int main()
{
    string str = "takeUforward";
    cout << "The length of the longest substring without repeating characters is " << solve(str);
    return 0;
}

// best approach
class Solution
{
public:
    int lengthofLongestSubstring(string s)
    {
        vector<int> mpp(256, -1);

        int left = 0, right = 0;
        int n = s.size();
        int len = 0;
        while (right < n)
        {
            if (mpp[s[right]] != -1)
                left = max(mpp[s[right]] + 1, left);

            mpp[s[right]] = right;

            len = max(len, right - left + 1);
            right++;
        }
        return len;
    }
};
// Time Complexity: O( N )
// Space Complexity: O(N) where N represents the size of HashSet where we are storing our elements

int main()
{
    string str = "takeUforward";
    Solution obj;
    cout << "The length of the longest substring without repeating characters is " << obj.lengthofLongestSubstring(str);
    return 0;
}

// sliding window (by me)
class Solution {
public:
        int lengthOfLongestSubstring(string s) {
         int i=0,j=0,n=s.size(),maxi=0;
        unordered_map<int, int> mp;

        while(j<n) {
            mp[s[j]]++;
            if(mp[s[j]]>1) {
                 while(mp[s[j]]>1) {
                     mp[s[i++]]--;
                    //  if(mp[s[i]]==0) mp.erase(s[i]);
                 }
            }
            maxi=max(maxi, j-i+1);
            j++;
        }

        return maxi;
    }
};