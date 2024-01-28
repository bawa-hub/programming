// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/solutions/685406/c-stack-next-smaller-element/

class Solution {
public:
    vector<int> finalPrices(vector<int>& prices) {
        int n = prices.size();
        stack<int> st;

        vector<int> res;

        for(int i=n-1;i>=0;i--) {
            while(!st.empty() && st.top()> prices[i]) st.pop();

            if(st.empty()) res.push_back(0);
            else res.push_back(st.top());

            st.push(prices[i]);
        }

        reverse(res.begin(), res.end());
        for(int i=0;i<n;i++) {
            res[i] = prices[i] -res[i];
        }

        return res;
    }
};