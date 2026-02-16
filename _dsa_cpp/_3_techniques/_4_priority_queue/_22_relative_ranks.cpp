// https://leetcode.com/problems/relative-ranks/description/
class Solution {
public:
    vector<string> findRelativeRanks(vector<int>& score) {
        priority_queue<pair<int, int>> pq;

        for(int i=0;i<score.size();i++) {
            pq.push(make_pair(score[i], i));
        }

      vector<string> v(score.size());
        int i=1;
        while(!pq.empty()) {
            pair<int, int> p = pq.top();
            pq.pop();
           if(i==1) {
               v[p.second] = "Gold Medal";
           } else if(i==2) {
               v[p.second] = "Silver Medal";
           } else if(i==3) {
               v[p.second] = "Bronze Medal";
           } else {
               v[p.second] = to_string(i);
           }
           i++;
        }
        return v;
    }
};