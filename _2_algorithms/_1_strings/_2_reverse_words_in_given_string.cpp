// https://leetcode.com/problems/reverse-words-in-a-string/

#include<iostream>
using namespace std;

class Solution
{
public:
    string reverseWords(string s)
    {
        string res = ""; // to store the final result
        int i = s.length() - 1;

        while (i >= 0)
        {
            if (s[i] == ' ')
                i--; // ignore the spaces
            else
            {
                int j = i;
                string temp = " "; // whenever adding a word start with a space
                while (j >= 0 && s[j] != ' ')
                {
                    temp += s[j--];
                }
                reverse(temp.begin(), temp.end());
                // since we added the word in reverse order so we need to reverse it
                res += temp;
                i = j;
            }
        }
        res.pop_back();
        /* last character will be a space since we are adding a space each time we
           add a word */
        return res;
    }
};
// TC:O(N) for traversing the string + O(N) for reversing
// SC:O(N)

class Solution1 {
public:

    void reverseString(string &s, int l, int r) {
        while(l < r) {
            swap(s[l], s[r]);
            l++;
            r--;
        }
    }

    string reverseWords(string s) {

        int n = s.size();
        int i = 0, j = 0;

        // remove extra spaces
        while(j < n) {

            while(j < n && s[j] == ' ') j++;

            while(j < n && s[j] != ' ')
                s[i++] = s[j++];

            while(j < n && s[j] == ' ') j++;

            if(j < n)
                s[i++] = ' ';
        }

        s.resize(i);

        // reverse entire string
        reverseString(s, 0, s.size() - 1);

        // reverse each word
        int start = 0;

        for(int end = 0; end <= s.size(); end++) {

            if(end == s.size() || s[end] == ' ') {
                reverseString(s, start, end - 1);
                start = end + 1;
            }
        }

        return s;
    }
};
// TC: O(n)
// SC: O(1)