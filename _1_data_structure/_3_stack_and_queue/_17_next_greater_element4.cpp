// https://leetcode.com/problems/next-greater-element-iv/
// https://leetcode.com/problems/next-greater-element-iv/solutions/2799346/monotonic-stack-min-heap/


class Solution {
public:
    vector<int> secondGreaterElement(vector<int>& nums) {
        int n = nums.size();
        vector<int> gE(n, -1);
        stack<int> st;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;
        
        for(int i = 0; i < n; i++) {
            while(!pq.empty() && pq.top().first < nums[i]) {
                gE[pq.top().second] = nums[i];
                pq.pop();
            }
            while(!st.empty() && nums[st.top()] < nums[i]) {
                pq.push(make_pair(nums[st.top()], st.top()));
                st.pop();
            }
            st.push(i);
        }
        
        return gE;
    }
};