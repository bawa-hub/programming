// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
// https://www.geeksforgeeks.org/find-possible-words-phone-digits/

class Solution
{
private:
    void solve(string digits, string output, int index, vector<string> &ans, string *mapping)
    {
        // base case
        if (index >= digits.length())
        {
            ans.push_back(output);
            return;
        }

        int number = digits[index] - '0'; // covert string to number .
        string value = mapping[number];   // map string value related to number.

        for (int i = 0; i < value.length(); i++)
        {
            output.push_back(value[i]);                     //
            solve(digits, output, index + 1, ans, mapping); // recursive call .
            output.pop_back();                              // back track beacuse of reach to the empty string .
        }
    }

public:
    vector<string> letterCombinations(string digits)
    {
        vector<string> ans;

        if (digits.length() == 0)
        {
            return ans;
        }

        string output = "";
        int index = 0;

        string mapping[10] = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

        solve(digits, output, index, ans, mapping);
        return ans;
    }
};