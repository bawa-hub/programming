#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        long long x, y;
        cin >> x >> y;
        if ((2 * y + x) % 2 == 0 && y != 1)
        {
            cout << "YES" << endl;
        }
        else
            cout << "NO" << endl;
    }
}