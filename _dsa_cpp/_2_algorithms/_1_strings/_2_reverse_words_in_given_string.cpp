// https://leetcode.com/problems/reverse-words-in-a-string/

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