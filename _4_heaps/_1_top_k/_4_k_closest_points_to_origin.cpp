// https://leetcode.com/problems/k-closest-points-to-origin/

class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        priority_queue<pair<int, pair<int, int>>> pq;

        for(int i=0;i<points.size();i++) {
            // no need to find sqrt 😎
            int dist = (points[i][0]*points[i][0])+(points[i][1]*points[i][1]);
            pq.push({dist, {points[i][0],points[i][1]}});

            if(pq.size()>k) pq.pop();
        } 

        vector<vector<int>> res;
        while(!pq.empty()) {
            res.push_back({pq.top().second.first,pq.top().second.second});
            pq.pop();
        }

        return res;
    }
};