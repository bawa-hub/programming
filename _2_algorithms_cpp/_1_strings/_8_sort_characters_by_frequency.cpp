// https://leetcode.com/problems/sort-characters-by-frequency/
// https://practice.geeksforgeeks.org/problems/sorting-elements-of-an-array-by-frequency/0

class Solution
{
public:
    string frequencySort(string s)
    {
        map<char, int> mp;
        multimap<int, char, greater<int>> mul; // multimap sort in desending order for this we dont have to take the loop fort this.
        string s1 = "";
        for (int i = 0; i < s.length(); i++)
        {
            mp[s[i]]++;
        }

        for (auto it : mp)
        {
            mul.insert({it.second, it.first});
        }

        for (auto it : mul)
        {
            for (int i = 0; i < it.first; i++)
            {
                s1 += it.second;
            }
        }
        return s1;
    }
};