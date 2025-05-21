#include <bits/stdc++.h>
using namespace std;

int main() {

    // max heap
    priority_queue<int> pq;
    // min heap
    priority_queue<int, vector<int>, greater<int>> pq2;
    // priority_queue with multiple keys, it will sort with first key
    priority_queue<pair<int, int>> pq3;

    // note: if to find kth minimum element use max heap else use min heap

    pq.push(10);
    pq.push(20);
    pq.push(30);

    pq2.push(10);
    pq2.push(20);
    pq2.push(30);

    pq3.push({1,2});
    pq3.push({5,9});
    pq3.push({3,6});

    cout << "pq top: " << pq.top() << endl; // 30
    cout << "pq2 top: " << pq2.top() << endl; // 10
    cout << "pq3 top: " << pq3.top().first << endl; // 5


    return 0;
}