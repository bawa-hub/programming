// https://leetcode.com/problems/crawler-log-folder/

class Solution {
public:
    int minOperations(vector<string>& logs) {
        stack<string> st;

        for(int i=0;i<logs.size();i++) {
            if(st.empty() && (logs[i] == "../" || logs[i] == "./")) continue;
            else if(st.empty()) st.push(logs[i]); 
            else {
                if(logs[i] == "../") st.pop();
                else if(logs[i] == "./") continue;
                else st.push(logs[i]);
            }
        }

        return st.size();
    }
};