#include <bits/stdc++.h>
using namespace std;

int main()
{
    pair<int, string> p, q;
    p = make_pair(1, "abc");
    q = {2, "abcd"};
    cout << p.first << " " << p.second << endl;
    cout << q.first << " " << q.second << endl;

    pair<int, string> p1 = p;
    p1.first = 3;
    cout << p1.first << " " << p1.second << endl;

    pair<int, string> &p2 = p;
    p2.first = 5;
    cout << p.first << " " << p.second << endl;
}