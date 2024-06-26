// https://leetcode.com/problems/remove-outermost-parentheses/

class Solution
{
public:
    string removeOuterParentheses(string s)
    {

        stack<char> st;
        string ans = "";

        for (auto c : s)
        {

            if (!st.empty())
                ans += c;
            
            if (c == '(')
                st.push(c);
            else
            {
                st.pop();
                if (st.empty())
                    ans.pop_back();
            }
        }

        return ans;
    }
};

// Time complexity:
// O(N)
// Space complexity:
// O(N)

class Solution {
public:
    string removeOuterParentheses(string S) {
        int count = 0;
        string str;
        for (char c : S) {
            if (c == '(') {
                if (count++) {
                    str += '(';
                }
            } else {
                if (--count) {
                    str += ')';
                }
            }
        }
        return str;
    }
};
// TC: O(n)
// SC: O(1)
