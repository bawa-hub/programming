// https://leetcode.com/problems/generate-parentheses/
// generate paranthesis of n group ()

#include <vector>
using namespace std;

class Solution
{
public:
    void generate(vector<string> &strings, string &s, int open, int close)
    {
        if (open == 0 && close == 0)
        {
            strings.push_back(s);
            return;
        }

        if (open > 0)
        {
            s.push_back('(');
            generate(strings, s, open - 1, close);
            s.pop_back(); // backtracking step
        }

        if (close > 0)
        {
            if (open < close)
            {
                s.push_back(')');
                generate(strings, s, open, close - 1);
                s.pop_back();
            }
        }
    }
    vector<string> generateParenthesis(int n)
    {
        vector<string> strings;
        string s;
        generate(strings, s, n, n);
        return strings;
    }
};