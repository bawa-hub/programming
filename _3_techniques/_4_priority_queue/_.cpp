#include <bits/stdc++.h>
using namespace std;

int main() {

    // max heap
    priority_queue<int> pq;
    // min heap
    priority_queue<int, vector<int>, greater<int>> pq2;

    // note: if to find kth minimum element, use max heap else use min heap

    pq.push(10);
    pq.push(20);
    pq.push(30);

    pq2.push(10);
    pq2.push(20);
    pq2.push(30);

    cout << pq.top() << endl; // 30
    cout << pq2.top() << endl; // 10

    return 0;
}