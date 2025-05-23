// https://leetcode.com/problems/make-the-string-great/

class Solution {
public:
    string makeGood(string s) {
        stack<char> st;

        for(int i=0;i<s.size();i++) {
            if(st.empty()) st.push(s[i]);
            else {
                if(abs(s[i]-st.top())==32) st.pop();
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