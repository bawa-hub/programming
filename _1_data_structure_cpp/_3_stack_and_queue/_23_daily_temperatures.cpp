// https://leetcode.com/problems/daily-temperatures/

class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        stack<pair<int, int>> st;
        vector<int> res;

        for(int i=temperatures.size()-1;i>=0;i--) {
            while(!st.empty()&&st.top().first<=temperatures[i]) st.pop();

            if(!st.empty()) {
                 res.push_back(st.top().second-i);
            } else {
                res.push_back(0);
            }

            st.push({temperatures[i], i});
        }

        reverse(res.begin(), res.end());

        return res;
    }
};