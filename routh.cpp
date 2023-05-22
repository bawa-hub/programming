#include <iostream>
#include <queue>
using namespace std;


long long minCost(long long arr[], long long n) {
        priority_queue<long long, vector<long long>, greater<long long> > pq;
        for(int i=0;i<n;i++) pq.push(arr[i]);
        long long cost = 0;
        
        while(pq.size()>=2) {
             long long first = pq.top();
             cout<< "first: " << first;
             pq.pop();
             long long sec = pq.top();
             cout<< "sec: " << sec << endl;
             pq.pop();
             cost+=(first+sec);
             pq.push(first+sec);
        }
        
        return cost;
    }

    int main() {
        long long arr[] = {4,3,2,6};
        cout << minCost(arr, 4);

    }