// https://leetcode.com/problems/baseball-game/

class Solution {
public:
    int calPoints(vector<string>& operations) {
        stack<int> st;

        for(int i=0;i<operations.size();i++) {
            if(operations[i] == "+") {
                int top = st.top();
                st.pop();
                int sum = top + st.top();
                st.push(top);
                st.push(sum);
            } else if(operations[i]=="D") {
               st.push(2*st.top());
            } else if(operations[i]=="C") {
               st.pop();
            } else {
                st.push(stoi(operations[i]));
            }
        }

        int res = 0;
        while(!st.empty()) {
           res+=st.top();
           st.pop();
        }

        return res;
    }
};