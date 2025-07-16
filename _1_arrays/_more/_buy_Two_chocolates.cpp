// https://leetcode.com/problems/buy-two-chocolates/description/

class Solution {
public:
    int buyChoco(vector<int>& prices, int money) {
        priority_queue<int> pq;
        
        for(int i=0;i<prices.size();i++) {
            pq.push(prices[i]);
            
            if(pq.size()>2) pq.pop();
        }
        
        int sum = 0;
        while(!pq.empty()) {
            sum += pq.top();
            pq.pop();
        }
        
        if(sum<=money) return money-sum;
        return money;
    }
};