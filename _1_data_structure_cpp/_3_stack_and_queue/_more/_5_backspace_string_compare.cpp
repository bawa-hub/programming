// https://leetcode.com/problems/backspace-string-compare/

class Solution {
public:
    bool backspaceCompare(string s, string t) {
        s = result(s);
        t = result(t);
        return s==t;
    }

    string result(string s) {
        stack<char> st;

        for(int i=0;i<s.size();i++) {
            if(st.empty() && s[i]!='#') {
                st.push(s[i]);
            } else if(st.empty() && s[i]=='#') continue;
            else {
                if(s[i]=='#') st.pop();
                else st.push(s[i]);
            }
        }

        string res = "";
        while(!st.empty()) {
            res+=st.top();
            st.pop();
        }
        reverse(res.begin(), res.end());
        return res;
    }
};