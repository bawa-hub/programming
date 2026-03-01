// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

class Solution {
public:
    string removeDuplicates(string s) {
        stack<char> st;

        for(int i=0;i<s.size();i++) {
            if(st.empty()) st.push(s[i]);
            else {
                if(!(st.top()==s[i])) {
                    st.push(s[i]);
                } else {
                    st.pop();
                }
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