// https://practice.geeksforgeeks.org/problems/minimum-cost-of-ropes-1587115620/1

class Solution
{
    public:
    //Function to return the minimum cost of connecting the ropes.
    long long minCost(long long arr[], long long n) {
        priority_queue<long long, vector<long long>, greater<long long>> pq;
        for(int i=0;i<n;i++) pq.push(arr[i]);
        long long cost = 0;
        
        while(pq.size()>=2) {
             long long first = pq.top();
             pq.pop();
             long long sec = pq.top();
             pq.pop();
             cost+=(first+sec);
             pq.push(first+sec);
        }
        
        return cost;
    }
};