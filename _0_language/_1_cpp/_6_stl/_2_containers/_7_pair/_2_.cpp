#include <bits/stdc++.h>
using namespace std;

int main()
{
    pair<int, int> p[3];
    p[0] = {1, 2};
    p[1] = {2, 3};
    p[2] = {3, 4};
    for (int i = 0; i < 3; ++i)
    {
        cout << p[i].first << " " << p[i].second << endl;
    }
}